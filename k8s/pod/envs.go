/*
 * Copyright 2020 Skulup Ltd, Open Collaborators
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
