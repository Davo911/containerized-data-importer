/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2018 Red Hat, Inc.
 *
 */

// this file directly copy/pasted from KubeVirt project here:
// https://github.com/kubevirt/kubevirt/blob/main/tools/util/marshaller.go

package util

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/ghodss/yaml"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// MarshallObject marshalls an object to yaml appropriate for kubectl
func MarshallObject(obj interface{}, writer io.Writer) error {
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	var r unstructured.Unstructured
	if err := json.Unmarshal(jsonBytes, &r.Object); err != nil {
		return err
	}

	unstructured.RemoveNestedField(r.Object, "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(r.Object, "spec", "template", "metadata", "creationTimestamp")
	unstructured.RemoveNestedField(r.Object, "status")

	deployments, exists, err := unstructured.NestedSlice(r.Object, "spec", "install", "spec", "deployments")
	if err != nil {
		return err
	}
	// remove timestamp and status from deployments in CSV objects
	// remove status and metadata.creationTimestamp
	if exists {
		for _, obj := range deployments {
			deployment := obj.(map[string]interface{})
			unstructured.RemoveNestedField(deployment, "metadata", "creationTimestamp")
			unstructured.RemoveNestedField(deployment, "spec", "template", "metadata", "creationTimestamp")
			unstructured.RemoveNestedField(deployment, "status")
		}
		if err = unstructured.SetNestedSlice(r.Object, deployments, "spec", "install", "spec", "deployments"); err != nil {
			return err
		}
	}

	jsonBytes, err = json.Marshal(r.Object)
	if err != nil {
		return err
	}

	yamlBytes, err := yaml.JSONToYAML(jsonBytes)
	if err != nil {
		return err
	}

	// fix templates by removing quotes...
	s := string(yamlBytes)
	s = strings.Replace(s, "'{{", "{{", -1)
	s = strings.Replace(s, "}}'", "}}", -1)
	yamlBytes = []byte(s)

	_, err = writer.Write([]byte("---\n"))
	if err != nil {
		return err
	}

	_, err = writer.Write(yamlBytes)
	if err != nil {
		return err
	}

	return nil
}
