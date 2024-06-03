#!/bin/bash
set -xe

cd /tmp

wget 'https://github.com/helmfile/helmfile/releases/download/v0.165.0/helmfile_0.165.0_linux_amd64.tar.gz' -O helmfile.tar.gz
tar -xvzf helmfile.tar.gz
mv helmfile /usr/local/bin

wget 'https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv5.4.2/kustomize_v5.4.2_linux_amd64.tar.gz' -O kustomize.tar.gz
tar -xvzf kustomize.tar.gz
mv kustomize /usr/local/bin