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

package replicaset

import (
	"log"

	"github.com/kubernetes/dashboard/resource/common"
	"github.com/kubernetes/dashboard/resource/event"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

// GetReplicaSetEvents gets events associated to replica set.
func GetReplicaSetEvents(client client.Interface, namespace, replicaSetName string) (
	*common.EventList, error) {

	log.Printf("Getting events related to %s replica set in %s namespace", replicaSetName,
		namespace)

	// Get events for replica set.
	rsEvents, err := event.GetEvents(client, namespace, replicaSetName)

	if err != nil {
		return nil, err
	}

	// Get events for pods in replica set.
	podEvents, err := GetReplicaSetPodsEvents(client, namespace, replicaSetName)

	if err != nil {
		return nil, err
	}

	apiEvents := append(rsEvents, podEvents...)

	if !event.IsTypeFilled(apiEvents) {
		apiEvents = event.FillEventsType(apiEvents)
	}

	events := event.AppendEvents(apiEvents, common.EventList{
		Namespace: namespace,
		Events:    make([]common.Event, 0),
	})

	log.Printf("Found %d events related to %s replica set in %s namespace",
		len(events.Events), replicaSetName, namespace)

	return &events, nil
}

// GetReplicaSetPodsEvents gets events associated to pods in replica set.
func GetReplicaSetPodsEvents(client client.Interface, namespace, replicaSetName string) (
	[]api.Event, error) {

	replicaSet, err := client.Extensions().ReplicaSets(namespace).Get(replicaSetName)

	if err != nil {
		return nil, err
	}

	podEvents, err := event.GetPodsEvents(client, namespace, replicaSet.Spec.Selector.MatchLabels)

	if err != nil {
		return nil, err
	}

	return podEvents, nil
}
