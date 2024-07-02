# Hekura (Helm + Kustomize + Raw manifests)
Hekura is a simple tool that combines Helmfile, Kustomize, and Raw Manifests to build highly customizable manifests. It currently serves as a template engine over the Helmfile and Kustomize tools to combine them.

### Why Combine Tools?
Using a customized chart can lead to a disconnect from the original, making updates more challenging. When there is a need to include additional manifests but modifications to the original chart are not possible, an algorithm that integrates Helm charts, Kustomize, and raw manifest files is proposed. This approach allows for the creation of a highly customizable set of manifests.

### Installation
To install Hekura, follow these steps:
- install `kubectl` from https://kubernetes.io/docs/tasks/tools/#kubectl 
- install `helmfile` from https://github.com/helmfile/helmfile/releases
- install `kustomize` from https://github.com/kubernetes-sigs/kustomize/releases 
- install `hekura` from https://github.com/mlkmhd/hekura/releases 

### Usage
To use Hekura, you need to have a `hekura.yaml` configuration file structured as follows:
```commandline
helmfile: 
  - sample-manifests/helmfile/

kustomize:
  - sample-manifests/kustomize/

raw-manifests:
  - sample-manifests/raw-manifests/
```
Your manifest files should be organized like this:
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

Hekura will generate and apply the manifests in the following order:

1- Build the `helmfile`.

2- Apply the `kustomize` path files to the manifests generated in the previous step.

3- Add `raw manifests` to the existing generated manifests.

You can generate and apply the manifests using the following command:
```commandline
$ hekura template --config hekura.yaml -o all.yaml
$ kubectl apply -f all.yaml
```

### Read more
the main idea behind this repository described at the following article: https://medium.com/itnext/helm-kustomize-raw-manifests-combination-570f81acf996
