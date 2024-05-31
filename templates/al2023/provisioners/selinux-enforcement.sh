#!/usr/bin/env bash

set -o pipefail
set -o nounset
set -o errexit

if [[ "$SELINUX_ENFORCED" == "true" ]]; then
  sudo chcon -t bin_t /usr/bin/nodeadm
  sudo chcon -t bin_t /usr/bin/kubelet
fi
