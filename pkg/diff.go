package pkg

// Diff currently returns a placeholder string prefixed with "hi" and the
// path to the built resource file.
//
// TODO: Implement actual diff functionality. This function should take the
// path to a file containing generated Kubernetes manifests and compare it
// against the currently applied manifests in a Kubernetes cluster.
// It might use `kubectl diff -f <file>` or similar mechanisms.
func Diff(builtResourceFile string) string {
	return "hi" + builtResourceFile
}
