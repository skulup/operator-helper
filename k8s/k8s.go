package k8s

// EnvVarPodIP holds the POD's IP
const EnvVarPodIP = "POD_IP"
const EnvVarEnvoySidecarStatus = "ENVOY_SIDECAR_STATUS"

const (
	// LabelAppName defines the app label
	LabelAppName = "app.kubernetes.io/name"
	// LabelAppManagedBy defines the managed-by label
	LabelAppManagedBy = "app.kubernetes.io/managed-by"
)

// ContainerShellCommand is helper factory method to create the shell command
func ContainerShellCommand() []string {
	return []string{
		"sh",
		"-c",
	}
}
