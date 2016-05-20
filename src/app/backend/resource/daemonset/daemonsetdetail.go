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

package daemonset

import (
	"log"

	"github.com/kubernetes/dashboard/client"
	"github.com/kubernetes/dashboard/resource/common"
	"github.com/kubernetes/dashboard/resource/pod"
	resourceService "github.com/kubernetes/dashboard/resource/service"
	"k8s.io/kubernetes/pkg/api"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/extensions"
	k8sClient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
)

// DaemonSeDetail represents detailed information about a Daemon Set.
type DaemonSetDetail struct {
	ObjectMeta common.ObjectMeta `json:"objectMeta"`
	TypeMeta   common.TypeMeta   `json:"typeMeta"`

	// Label selector of the Daemon Set.
	LabelSelector *unversioned.LabelSelector `json:"labelSelector,omitempty"`

	// Container image list of the pod template specified by this Daemon Set.
	ContainerImages []string `json:"containerImages"`

	// Aggregate information about pods of this daemon set.
	PodInfo common.PodInfo `json:"podInfo"`

	// Detailed information about Pods belonging to this Daemon Set.
	Pods pod.PodList `json:"pods"`

	// Detailed information about service related to Daemon Set.
	ServiceList resourceService.ServiceList `json:"serviceList"`

	// True when the data contains at least one pod with metrics information, false otherwise.
	HasMetrics bool `json:"hasMetrics"`
}

// Returns detailed information about the given daemon set in the given namespace.
func GetDaemonSetDetail(client k8sClient.Interface, heapsterClient client.HeapsterClient,
	namespace, name string) (*DaemonSetDetail, error) {
	log.Printf("Getting details of %s daemon set in %s namespace", name, namespace)

	daemonSetWithPods, err := getRawDaemonSetWithPods(client, namespace, name)
	if err != nil {
		return nil, err
	}
	daemonSet := daemonSetWithPods.DaemonSet
	pods := daemonSetWithPods.Pods

	services, err := client.Services(namespace).List(api.ListOptions{
		LabelSelector: labels.Everything(),
		FieldSelector: fields.Everything(),
	})

	nodes, err := client.Nodes().List(api.ListOptions{
		LabelSelector: labels.Everything(),
		FieldSelector: fields.Everything(),
	})

	if err != nil {
		return nil, err
	}

	daemonSetDetail := &DaemonSetDetail{
		ObjectMeta:    common.NewObjectMeta(daemonSet.ObjectMeta),
		TypeMeta:      common.NewTypeMeta(common.ResourceKindDaemonSet),
		LabelSelector: daemonSet.Spec.Selector,
		PodInfo:       getDaemonSetPodInfo(daemonSet, pods.Items),
		ServiceList:   resourceService.ServiceList{Services: make([]resourceService.Service, 0)},
	}

	matchingServices := getMatchingServicesforDS(services.Items, daemonSet)

	for _, service := range matchingServices {
		daemonSetDetail.ServiceList.Services = append(daemonSetDetail.ServiceList.Services,
			getService(service, *daemonSet, pods.Items, nodes.Items))
	}

	for _, container := range daemonSet.Spec.Template.Spec.Containers {
		daemonSetDetail.ContainerImages = append(daemonSetDetail.ContainerImages,
			container.Image)
	}

	daemonSetDetail.Pods = pod.CreatePodList(pods.Items, heapsterClient)

	return daemonSetDetail, nil
}

// TODO(floreks): This should be transactional to make sure that DS will not be deleted without pods
// Deletes daemon set with given name in given namespace and related pods.
// Also deletes services related to daemon set if deleteServices is true.
func DeleteDaemonSet(client k8sClient.Interface, namespace, name string,
	deleteServices bool) error {

	log.Printf("Deleting %s daemon set from %s namespace", name, namespace)

	if deleteServices {
		if err := DeleteDaemonSetServices(client, namespace, name); err != nil {
			return err
		}
	}

	pods, err := getRawDaemonSetPods(client, namespace, name)
	if err != nil {
		return err
	}

	if err := client.Extensions().DaemonSets(namespace).Delete(name); err != nil {
		return err
	}

	for _, pod := range pods.Items {
		if err := client.Pods(namespace).Delete(pod.Name, &api.DeleteOptions{}); err != nil {
			return err
		}
	}

	log.Printf("Successfully deleted %s daemon set from %s namespace", name, namespace)

	return nil
}

// Deletes services related to daemon set with given name in given namespace.
func DeleteDaemonSetServices(client k8sClient.Interface, namespace, name string) error {
	log.Printf("Deleting services related to %s daemon set from %s namespace", name,
		namespace)

	daemonSet, err := client.Extensions().DaemonSets(namespace).Get(name)
	if err != nil {
		return err
	}

	labelSelector, err := unversioned.LabelSelectorAsSelector(daemonSet.Spec.Selector)
	if err != nil {
		return err
	}

	services, err := getServicesForDSDeletion(client, labelSelector, namespace)
	if err != nil {
		return err
	}

	for _, service := range services {
		if err := client.Services(namespace).Delete(service.Name); err != nil {
			return err
		}
	}

	log.Printf("Successfully deleted services related to %s daemon set from %s namespace",
		name, namespace)

	return nil
}

// Returns detailed information about service from given service
func getService(service api.Service, daemonSet extensions.DaemonSet,
	pods []api.Pod, nodes []api.Node) resourceService.Service {

	result := resourceService.ToService(&service)
	// TODO: This may be wrong as we dont use all attributes from selector
	result.ExternalEndpoints = common.GetExternalEndpoints(daemonSet.Spec.Selector.MatchLabels, pods, service, nodes)

	return result
}
