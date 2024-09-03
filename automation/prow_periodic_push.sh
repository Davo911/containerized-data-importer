#!/bin/sh
set -e
build_date="$(date +%Y%m%d)"
cat "$QUAY_PASSWORD" | docker login --username $(cat "$QUAY_USER") --password-stdin=true quay.io

#TODO Fix this case
case $BUILD_ARCH in
  crossbuild-s390x)
    export DOCKER_TAG="${build_date}_$(git show -s --format=%h)-s390x"
    ;;
  crossbuild-aarch64)
    export DOCKER_TAG="${build_date}_$(git show -s --format=%h)-arm64"
    ;;
  *)
    echo "Invalid BUILD_ARCH value: ${BUILD_ARCH}"
    exit 1
    ;;
esac


make manifests
make bazel-push-images
bucket_dir="kubevirt-prow/devel/nightly/release/kubevirt/containerized-data-importer/${build_date}"
gsutil cp ./_out/manifests/release/cdi-operator.yaml gs://$bucket_dir/cdi-operator-${BUILD_ARCH}.yaml
gsutil cp ./_out/manifests/release/cdi-cr.yaml gs://$bucket_dir/cdi-cr-${BUILD_ARCH}.yaml
