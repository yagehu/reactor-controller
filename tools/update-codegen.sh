#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
bash "${CODEGEN_PKG}"/generate-groups.sh all \
    github.com/yagehu/reactor-controller/pkg/generated \
    github.com/yagehu/reactor-controller/pkg/apis \
    reactor:v1alpha1 \
    --output-base "$(dirname "${BASH_SOURCE[0]}")/../../../.." \
    --go-header-file "${SCRIPT_ROOT}"/tools/boilerplate.go.txt