package pods

import (
	"github.com/skulup/operator-helper/types"
	v1 "k8s.io/api/core/v1"
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
