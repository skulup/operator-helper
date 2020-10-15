package pod

import (
	"github.com/skulup/operator-helper/k8s"
	v1 "k8s.io/api/core/v1"
)

// DecorateContainerEnvVars generate the pod environment variables
func DecorateContainerEnvVars(envoySideCarStatus bool, sources ...v1.EnvVar) []v1.EnvVar {
	sources = append(sources, v1.EnvVar{
		Name: k8s.EnvVarPodIP,
		ValueFrom: &v1.EnvVarSource{
			FieldRef: &v1.ObjectFieldSelector{
				FieldPath: "status.podIP",
			},
		},
	})
	if envoySideCarStatus {
		sources = append(sources, v1.EnvVar{
			Name: k8s.EnvVarEnvoySidecarStatus,
			ValueFrom: &v1.EnvVarSource{
				FieldRef: &v1.ObjectFieldSelector{
					FieldPath: `metadata.annotations['sidecar.istio.io/status']`,
				},
			},
		})
	}
	return sources
}
