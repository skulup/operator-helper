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
