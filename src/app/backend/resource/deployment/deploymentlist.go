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

package deployment

import (
	"log"

	"github.com/kubernetes/dashboard/resource/common"

	"k8s.io/kubernetes/pkg/api"
	k8serrors "k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/apis/extensions"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

// ReplicationSetList contains a list of Deployments in the cluster.
type DeploymentList struct {
	// Unordered list of Deployments.
	Deployments []Deployment `json:"deployments"`
}

// Deployment is a presentation layer view of Kubernetes Deployment resource. This means
// it is Deployment plus additional augumented data we can get from other sources
// (like services that target the same pods).
type Deployment struct {
	ObjectMeta common.ObjectMeta `json:"objectMeta"`
	TypeMeta   common.TypeMeta   `json:"typeMeta"`

	// Aggregate information about pods belonging to this Deployment.
	Pods common.PodInfo `json:"pods"`

	// Container images of the Deployment.
	ContainerImages []string `json:"containerImages"`
}

// GetDeploymentList returns a list of all Deployments in the cluster.
func GetDeploymentList(client client.Interface, namespaces []string) (*DeploymentList, error) {
	log.Printf("Getting list of all deployments in the cluster")

	channels := &common.ResourceChannels{
		DeploymentList: common.GetDeploymentListChannel(client.Extensions(), 1),
		ServiceList:    common.GetServiceListChannel(client, 1),
		PodList:        common.GetPodListChannel(client, 1),
		EventList:      common.GetEventListChannel(client, 1),
		NodeList:       common.GetNodeListChannel(client, 1),
	}

	return GetDeploymentListFromChannels(channels, namespaces)
}

// GetDeploymentList returns a list of all Deployments in the cluster
// reading required resource list once from the channels.
func GetDeploymentListFromChannels(channels *common.ResourceChannels, namespaces []string) (
	*DeploymentList, error) {

	deployments := <-channels.DeploymentList.List
	if err := <-channels.DeploymentList.Error; err != nil {
		statusErr, ok := err.(*k8serrors.StatusError)
		if ok && statusErr.ErrStatus.Reason == "NotFound" {
			// NotFound - this means that the server does not support Deployment objects, which
			// is fine.
			emptyList := &DeploymentList{
				Deployments: make([]Deployment, 0),
			}
			return emptyList, nil
		}
		return nil, err
	}

	services := <-channels.ServiceList.List
	if err := <-channels.ServiceList.Error; err != nil {
		return nil, err
	}

	pods := <-channels.PodList.List
	if err := <-channels.PodList.Error; err != nil {
		return nil, err
	}

	events := <-channels.EventList.List
	if err := <-channels.EventList.Error; err != nil {
		return nil, err
	}

	nodes := <-channels.NodeList.List
	if err := <-channels.NodeList.Error; err != nil {
		return nil, err
	}

	result := getDeploymentList(deployments.Items, services.Items, pods.Items, events.Items,
		nodes.Items)

	if namespaces != nil {
		return FilterByNamespace(namespaces, result), nil
	}

	return result, nil
}

func getDeploymentList(deployments []extensions.Deployment,
	services []api.Service, pods []api.Pod, events []api.Event,
	nodes []api.Node) *DeploymentList {

	deploymentList := &DeploymentList{
		Deployments: make([]Deployment, 0),
	}

	for _, deployment := range deployments {

		matchingPods := common.FilterNamespacedPodsBySelector(pods, deployment.ObjectMeta.Namespace,
			deployment.Spec.Selector.MatchLabels)
		podInfo := getPodInfo(&deployment, matchingPods)

		deploymentList.Deployments = append(deploymentList.Deployments,
			Deployment{
				ObjectMeta:      common.NewObjectMeta(deployment.ObjectMeta),
				TypeMeta:        common.NewTypeMeta(common.ResourceKindDeployment),
				ContainerImages: common.GetContainerImages(&deployment.Spec.Template.Spec),
				Pods:            podInfo,
			})
	}

	return deploymentList
}

func FilterByNamespace(namespaces []string, dp *DeploymentList) *DeploymentList {
	res := make([]Deployment, 0)
	result := new(DeploymentList)
	for _, v1 := range dp.Deployments {
		for _, v2 := range namespaces {
			if v2 == v1.ObjectMeta.Namespace {
				res = append(res, v1)
				break
			}
		}
	}

	result.Deployments = res
	return result
}
