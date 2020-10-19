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
	v12 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/coreos/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	"github.com/skulup/operator-helper/configs"
	v13 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:openapi-gen=true
// +kubebuilder:object:generate=true

// MetricSpec defines some properties use the create the ServiceMonitorSpec
// objects if prometheus metrics is supported by the platform
type MetricSpec struct {
	// SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.
	SampleLimit uint64 `json:"sampleLimit,omitempty"`
	// Timeout after which the scrape is ended
	ScrapeTimeout string `json:"scrapeTimeout,omitempty"`
	// ScrapeInterval defines the interval at which metrics should be scraped
	ScrapeInterval string `json:"scrapeInterval,omitempty"`
	// TlsConfig defines the TLS configuration to use when scraping the endpoint
	TlsConfig *v12.TLSConfig `json:"tlsConfig,omitempty"`
	// BasicAuth allow an endpoint to authenticate over basic authentication More info: https://prometheus.io/docs/operating/configuration/#endpoints
	BasicAuth *v12.BasicAuth `json:"basicAuth,omitempty"`
	// BearerTokenFile defines the file to read bearer token for scraping targets.
	BearerTokenFile string `json:"bearerTokenFile,omitempty"`
	// BearerTokenSecret defines the secret to mount to read bearer token for scraping targets. The secret needs to be in the same namespace as the service monitor and accessible by the Prometheus Operator.
	BearerTokenSecret v13.SecretKeySelector `json:"bearerTokenSecret,omitempty"`
	// RelabelConfigs to apply to samples before scraping. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config
	ReLabelings []*v12.RelabelConfig `json:"relabelings,omitempty"`
	// MetricRelabelConfigs to apply to samples before ingestion.
	MetricReLabelings []*v12.RelabelConfig `json:"metricRelabelings,omitempty"`
}

// NewServiceMonitor creates a ServiceMonitor from the MetricSpec
func (in *MetricSpec) NewServiceMonitor(name, namespace string, labels map[string]string, labelSector metav1.LabelSelector, port string) *v12.ServiceMonitor {
	return &v12.ServiceMonitor{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceMonitor",
			APIVersion: "monitoring.coreos.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: v12.ServiceMonitorSpec{
			Selector:    labelSector,
			SampleLimit: in.SampleLimit,
			Endpoints: []v12.Endpoint{
				{
					Port:                 port,
					Interval:             in.ScrapeInterval,
					ScrapeTimeout:        in.ScrapeTimeout,
					TLSConfig:            in.TlsConfig,
					BearerTokenFile:      in.BearerTokenFile,
					BearerTokenSecret:    in.BearerTokenSecret,
					BasicAuth:            in.BasicAuth,
					MetricRelabelConfigs: in.MetricReLabelings,
					RelabelConfigs:       in.ReLabelings,
				},
			},
		},
	}
}

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
