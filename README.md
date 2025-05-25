# Hekura (Helm + Kustomize + Raw manifests)

A simple tool for combining Helmfile, Kustomize, and Raw Manifests together to build highly customizable Kubernetes manifests.

## Why Combine Tools?

Using a customized Helm chart can lead to a disconnect from the original chart, making updates difficult. When you need to include additional manifests but cannot modify the original chart, Hekura provides a solution by integrating Helm charts, Kustomize overlays, and raw manifest files to create a highly customizable set of manifests.

## Installation

To build Hekura from source, ensure you have Go installed on your system. Then, run the following command from the root of the project directory:

```bash
go build ./cmd/hekura/
```
This will create the `hekura` binary in your current directory. You can then move this binary to a directory in your `PATH`, such as `/usr/local/bin/`.

## Configuration

Hekura uses a `hekura.yaml` file to define the sources for your manifests. The configuration structure is as follows:

```yaml
helmfile: 
  - sample-manifests/helmfile/  # Path to a directory containing helmfile.yaml

kustomize:
  - sample-manifests/kustomize/ # Path to a directory containing kustomization.yaml

raw-manifests:
  - sample-manifests/raw-manifests/ # Path to a directory containing raw Kubernetes YAML files
```

You can specify multiple directories for each type of source. Hekura processes these sources in the following order:

1.  **Helmfile**: Generates manifests using Helmfile.
2.  **Kustomize**: Applies Kustomize overlays to the generated manifests.
3.  **Raw Manifests**: Includes any additional raw Kubernetes YAML files.

An example directory structure could be:

```
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

## Usage

Hekura provides the following commands:

### `template`

Generates the final Kubernetes manifests and prints them to standard output.

```bash
hekura template --config hekura.yaml
```

### `diff`

Generates the final Kubernetes manifests and shows the differences compared to the currently applied manifests in your cluster (requires `kubectl` to be configured).

```bash
hekura diff --config hekura.yaml
```

Both commands require a `--config` flag (or `-c`) to specify the path to your `hekura.yaml` file. If not provided, it defaults to `hekura.yaml` in the current directory.

## Contributing

Contributions are welcome! Please follow these guidelines:

1.  **Fork the repository.**
2.  **Create a new branch** for your feature or bug fix (e.g., `feature/my-new-feature` or `bugfix/issue-123`).
3.  **Write tests** for your changes to ensure they work as expected.
4.  **Format your code** using `go fmt` or `goimports` before committing.
5.  **Run the linter** to check for code style issues using the script: `./scripts/lint.sh`.
6.  **Create a pull request** against the `main` branch. Provide a clear description of your changes.

## Read More

The main idea behind this repository is described in the following article: [Helm + Kustomize + Raw Manifests Combination](https://medium.com/itnext/helm-kustomize-raw-manifests-combination-570f81acf996)
