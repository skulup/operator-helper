package pods

import (
	"github.com/skulup/operator-helper/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewPodSpec(cfg types.PodConfig, volumes []v1.Volume, initContainers []v1.Container, containers []v1.Container) v1.PodSpec {
	var activeDeadlineSeconds *int64
	if cfg.ActiveDeadlineSeconds > 0 {
		activeDeadlineSeconds = &cfg.ActiveDeadlineSeconds
	}
	return v1.PodSpec{
		Affinity:              &cfg.Affinity,
		Tolerations:           cfg.Tolerations,
		NodeSelector:          cfg.NodeSelector,
		RestartPolicy:         cfg.RestartPolicy,
		SecurityContext:       &cfg.SecurityContext,
		ActiveDeadlineSeconds: activeDeadlineSeconds,
		InitContainers:        initContainers,
		Containers:            containers,
		Volumes:               volumes,
	}
}

func NewPodTemplateSpec(name, generateName string, labels, annotations map[string]string, podSpec v1.PodSpec) v1.PodTemplateSpec {
	return v1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Name:         name,
			GenerateName: generateName,
			Labels:       labels,
			Annotations:  annotations,
		},
		Spec: podSpec,
	}
}
