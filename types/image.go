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

package types

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
)

// +k8s:openapi-gen=true
// +kubebuilder:object:generate=false

// Image represents the container image of a pod
type Image struct {
	// The container repository
	Repository string `json:"repository,omitempty"`
	// The container tag
	Tag string `json:"tag,omitempty"`

	PullPolicy v1.PullPolicy `json:"imagePullPolicy,omitempty"`
}

// Name returns the actual docker image name in the format <repository>:<tag>
// Deprecated. New code should use ToString
func (in Image) Name() string {
	return in.ToString()
}

// ToString returns the actual docker image name in the format <repository>:<tag>
func (in Image) ToString() string {
	return fmt.Sprintf("%s:%s", in.Repository, in.Tag)
}

func (in *Image) SetDefaults(repository, tag string, pullPolicy v1.PullPolicy) (changed bool) {
	if in.Repository == "" {
		changed = true
		in.Repository = repository
	}
	if in.Tag == "" {
		changed = true
		in.Tag = tag
	}
	if in.PullPolicy == "" {
		changed = true
		in.PullPolicy = pullPolicy
	}
	return
}
