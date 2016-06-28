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

package workload

import (
	"log"

	"github.com/kubernetes/dashboard/client"
	"github.com/kubernetes/dashboard/resource/common"
	"github.com/kubernetes/dashboard/resource/deployment"
	"github.com/kubernetes/dashboard/resource/pod"
	"github.com/kubernetes/dashboard/resource/replicaset"
	"github.com/kubernetes/dashboard/resource/replicationcontroller"
	"github.com/kubernetes/dashboard/resource/user"
	httpdbuser "github.com/wzzlYwzzl/httpdatabase/resource/user"
	k8sClient "k8s.io/kubernetes/pkg/client/unversioned"
)

// Workloads stucture contains all resource lists grouped into the workloads category.
type Workloads struct {
	DeploymentList deployment.DeploymentList `json:"deploymentList"`

	ReplicaSetList replicaset.ReplicaSetList `json:"replicaSetList"`

	ReplicationControllerList replicationcontroller.ReplicationControllerList `json:"replicationControllerList"`

	PodList pod.PodList `json:"podList"`

	UserList httpdbuser.UserList `json:"userList"`
}

// GetWorkloads returns a list of all workloads in the cluster.
func GetWorkloads(client k8sClient.Interface,
	heapsterClient client.HeapsterClient, httpdbClient *client.HttpDBClient, namespaces []string) (*Workloads, error) {

	log.Printf("Getting lists of all workloads")
	channels := &common.ResourceChannels{
		ReplicationControllerList: common.GetReplicationControllerListChannel(client, 1),
		ReplicaSetList:            common.GetReplicaSetListChannel(client.Extensions(), 1),
		DeploymentList:            common.GetDeploymentListChannel(client.Extensions(), 1),
		ServiceList:               common.GetServiceListChannel(client, 3),
		PodList:                   common.GetPodListChannel(client, 4),
		EventList:                 common.GetEventListChannel(client, 3),
		NodeList:                  common.GetNodeListChannel(client, 3),
		UserList:                  common.GetUserListChannel(httpdbClient, 2),
	}

	return GetWorkloadsFromChannels(channels, heapsterClient, httpdbClient, namespaces)
}

// GetWorkloadsFromChannels returns a list of all workloads in the cluster, from the
// channel sources.
func GetWorkloadsFromChannels(channels *common.ResourceChannels,
	heapsterClient client.HeapsterClient, httpdbClient *client.HttpDBClient, namespaces []string) (*Workloads, error) {

	rsChan := make(chan *replicaset.ReplicaSetList)
	deploymentChan := make(chan *deployment.DeploymentList)
	rcChan := make(chan *replicationcontroller.ReplicationControllerList)
	podChan := make(chan *pod.PodList)
	errChan := make(chan error, 5)
	userChan := make(chan *httpdbuser.UserList)

	go func() {
		rcList, err := replicationcontroller.GetReplicationControllerListFromChannels(channels, namespaces)
		errChan <- err
		rcChan <- rcList
	}()

	go func() {
		rsList, err := replicaset.GetReplicaSetListFromChannels(channels, namespaces)
		errChan <- err
		rsChan <- rsList
	}()

	go func() {
		deploymentList, err := deployment.GetDeploymentListFromChannels(channels, namespaces)
		errChan <- err
		deploymentChan <- deploymentList
	}()

	go func() {
		podList, err := pod.GetPodListFromChannels(channels, heapsterClient, namespaces)
		errChan <- err
		podChan <- podList
	}()

	go func() {
		userList, err := user.GetUserListFromChannels(channels, httpdbClient, namespaces)
		userChan <- userList
		errChan <- err
	}()

	rcList := <-rcChan
	err := <-errChan
	if err != nil {
		return nil, err
	}

	podList := <-podChan
	err = <-errChan
	if err != nil {
		return nil, err
	}

	rsList := <-rsChan
	err = <-errChan
	if err != nil {
		return nil, err
	}

	deploymentList := <-deploymentChan
	err = <-errChan
	if err != nil {
		return nil, err
	}

	userList := <-userChan
	err = <-errChan
	if err != nil {
		return nil, err
	}

	workloads := &Workloads{
		ReplicaSetList:            *rsList,
		ReplicationControllerList: *rcList,
		DeploymentList:            *deploymentList,
		PodList:                   *podList,
		UserList:                  *userList,
	}

	return workloads, nil
}
