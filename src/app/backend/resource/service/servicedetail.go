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

	"k8s.io/kubernetes/pkg/api"
	k8sClient "k8s.io/kubernetes/pkg/client/unversioned"

	"github.com/kubernetes/dashboard/client"
	"github.com/kubernetes/dashboard/resource/common"
	"github.com/kubernetes/dashboard/resource/pod"
)

// Service is a representation of a service.
type ServiceDetail struct {
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

	// Type determines how the service will be exposed.  Valid options: ClusterIP, NodePort, LoadBalancer
	Type api.ServiceType `json:"type"`

	// ClusterIP is usually assigned by the master. Valid values are None, empty string (""), or
	// a valid IP address. None can be specified for headless services when proxying is not required
	ClusterIP string `json:"clusterIP"`

	// PodList represents list of pods targeted by same label selector as this service.
	PodList pod.PodList `json:"podList"`
}

// GetServiceDetail gets service details.
func GetServiceDetail(client k8sClient.Interface, heapsterClient client.HeapsterClient,
	namespace, name string) (*ServiceDetail, error) {

	log.Printf("Getting details of %s service in %s namespace", name, namespace)

	// TODO(maciaszczykm): Use channels.
	serviceData, err := client.Services(namespace).Get(name)
	if err != nil {
		return nil, err
	}

	podList, err := GetServicePods(client, heapsterClient, namespace, serviceData.Spec.Selector)
	if err != nil {
		return nil, err
	}

	service := ToServiceDetail(serviceData)
	service.PodList = *podList

	return &service, nil
}

// GetServicePods gets list of pods targeted by given label selector in given namespace.
func GetServicePods(client k8sClient.Interface, heapsterClient client.HeapsterClient,
	namespace string, serviceSelector map[string]string) (*pod.PodList, error) {

	channels := &common.ResourceChannels{
		PodList: common.GetPodListChannel(client, 1),
	}

	apiPodList := <-channels.PodList.List
	if err := <-channels.PodList.Error; err != nil {
		return nil, err
	}

	apiPods := common.FilterNamespacedPodsBySelector(apiPodList.Items, namespace, serviceSelector)
	podList := pod.CreatePodList(apiPods, heapsterClient)

	return &podList, nil
}
