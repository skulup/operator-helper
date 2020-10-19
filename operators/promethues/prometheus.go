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

package promethues

import (
	"github.com/coreos/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	"github.com/skulup/operator-helper/configs"
)

// NewAlertmanagerInterface creates new AlertmanagerInterface
func NewAlertmanagerInterface(namespace string) v1.AlertmanagerInterface {
	return NewMonitoringInterface().Alertmanagers(namespace)
}

// NewPodMonitorInterface creates new PodMonitorInterface
func NewPodMonitorInterface(namespace string) v1.PodMonitorInterface {
	return NewMonitoringInterface().PodMonitors(namespace)
}

// NewPrometheusInterface creates new PrometheusInterface
func NewPrometheusInterface(namespace string) v1.PrometheusInterface {
	return NewMonitoringInterface().Prometheuses(namespace)
}

// ServiceMonitorInterface creates new ServiceMonitorInterface
func NewServiceMonitorInterface(namespace string) v1.ServiceMonitorInterface {
	return NewMonitoringInterface().ServiceMonitors(namespace)
}

// NewThanosRulerInterface creates new ThanosRulerInterface
func NewThanosRulerInterface(namespace string) v1.ThanosRulerInterface {
	return NewMonitoringInterface().ThanosRulers(namespace)
}

// NewMonitoringInterface creates a new MonitoringV1Interface
func NewMonitoringInterface() v1.MonitoringV1Interface {
	return v1.New(configs.RequireRestClient())
}
