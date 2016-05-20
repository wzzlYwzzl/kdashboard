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

package service

import (
	"log"

	client "k8s.io/kubernetes/pkg/client/unversioned"

	"github.com/kubernetes/dashboard/resource/common"
)

// Service is a representation of a service.
type Service struct {
	ObjectMeta common.ObjectMeta `json:"objectMeta"`
	TypeMeta   common.TypeMeta   `json:"typeMeta"`

	// InternalEndpoint of all Kubernetes services that have the same label selector as connected Replication
	// Controller. Endpoint is DNS name merged with ports.
	InternalEndpoint common.Endpoint `json:"internalEndpoint"`

	// ExternalEndpoints of all Kubernetes services that have the same label selector as connected Replication
	// Controller. Endpoint is external IP address name merged with ports.
	ExternalEndpoints []common.Endpoint `json:"externalEndpoints"`

	// Label selector of the service.
	Selector map[string]string `json:"selector"`

	// ClusterIP is usually assigned by the master. Valid values are None, empty string (""), or
	// a valid IP address. None can be specified for headless services when proxying is not required
	ClusterIP string `json:"clusterIP"`
}

// ServiceList contains a list of services in the cluster.
type ServiceList struct {
	// Unordered list of services.
	Services []Service `json:"services"`
}

// GetServiceList returns a list of all services in the cluster.
func GetServiceList(client client.Interface) (*ServiceList, error) {
	log.Printf("Getting list of all services in the cluster")

	channels := &common.ResourceChannels{
		ServiceList: common.GetServiceListChannel(client, 1),
	}

	services := <-channels.ServiceList.List
	if err := <-channels.ServiceList.Error; err != nil {
		return nil, err
	}

	serviceList := &ServiceList{Services: make([]Service, 0)}
	for _, service := range services.Items {
		serviceList.Services = append(serviceList.Services, ToService(&service))
	}

	return serviceList, nil
}
