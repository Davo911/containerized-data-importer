//go:build arm64
// +build arm64

package importer

import (
	"errors"
	"net/url"

	v1 "k8s.io/api/core/v1"

	"kubevirt.io/containerized-data-importer/pkg/common"
)

// VDDKDataSource is the data provider for vddk.
type VDDKDataSource struct {
}

func (V VDDKDataSource) Info() (ProcessingPhase, error) {
	panic("not support")
}

func (V VDDKDataSource) Transfer(path string, preallocation bool) (ProcessingPhase, error) {
	panic("not support")
}

func (V VDDKDataSource) TransferFile(fileName string, preallocation bool) (ProcessingPhase, error) {
	panic("not support")
}

func (V VDDKDataSource) GetURL() *url.URL {
	panic("not support")
}

func (V VDDKDataSource) GetTerminationMessage() *common.TerminationMessage {
	panic("not support")
}

func (V VDDKDataSource) Close() error {
	panic("not support")
}

func (V VDDKDataSource) IsDeltaCopy() bool {
	return false
}

func NewVDDKDataSource(endpoint string, accessKey string, secKey string, thumbprint string, uuid string, backingFile string, currentCheckpoint string, previousCheckpoint string, finalCheckpoint string, volumeMode v1.PersistentVolumeMode) (*VDDKDataSource, error) {
	return nil, errors.New("the arrch64 architecture does not support VDDK")
}

var _ DataSourceInterface = &VDDKDataSource{}
