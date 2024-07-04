# Hekura (Helm + Kustomize + Raw manifests)
A simple tool for combining Helmfile and Kustomize and Raw Manifests together to build highly customizable manifests

### Why Combine Tools?
using a customized chart can lead to a disconnect from the original and hinder updates. When needing to include additional manifests but unable to modify the original chart, an algorithm integrating Helm charts, Kustomize, and raw manifest files is proposed to create a highly customizable set of manifests.

### Configuration
The `hekura.yaml` configuration is like this:
```commandline
helmfile: 
  - sample-manifests/helmfile/

kustomize:
  - sample-manifests/kustomize/

raw-manifests:
  - sample-manifests/raw-manifests/
```
your manifest files can be structured likes this:
```commandline
└── sample-manifests
    ├── helmfile
    │   ├── helmfile.yaml
    │   └── values.yaml
    ├── kustomize
    │   ├── kustomization.yaml
    │   └── patch.yaml
    └── raw-manifests
        └── network-policy.yaml
```

The Hekura will start to generating and applying the manifests by the following order:

1- helmfiles

2- kustomize

3- raw manifests

in the above configuration you can pass multiple helmfile and kustomize and raw-manifests directories.

### Manifest Generation
you can generate the manifests using the following command:
```commandline
$ hekura template --config hekura.yaml
```

### Read more
the main idea behind this repository described at the following article: https://medium.com/itnext/helm-kustomize-raw-manifests-combination-570f81acf996
