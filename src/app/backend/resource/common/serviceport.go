// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"k8s.io/kubernetes/pkg/api"
)

// ServicePort is a pair of port and protocol, e.g. a service endpoint.
type ServicePort struct {
	// Positive port number.
	Port int `json:"port"`

	// Protocol name, e.g., TCP or UDP.
	Protocol api.Protocol `json:"protocol"`
}

// GetServicePorts returns human readable name for the given service ports list.
func GetServicePorts(apiPorts []api.ServicePort) []ServicePort {
	var ports []ServicePort
	for _, port := range apiPorts {
		ports = append(ports, ServicePort{port.Port, port.Protocol})
	}
	return ports
}
