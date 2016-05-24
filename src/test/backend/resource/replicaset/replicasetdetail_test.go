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
	"reflect"
	"testing"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/client/restclient"
	k8sClient "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/testclient"

	"github.com/kubernetes/dashboard/client"
	"github.com/kubernetes/dashboard/resource/common"
	"github.com/kubernetes/dashboard/resource/pod"
)

type FakeHeapsterClient struct {
	client k8sClient.Interface
}

func (c FakeHeapsterClient) Get(path string) client.RequestInterface {
	return &restclient.Request{}
}

func TestGetReplicaSetDetail(t *testing.T) {
	eventList := &api.EventList{}
	podList := &api.PodList{}

	cases := []struct {
		namespace, name string
		expectedActions []string
		replicaSet      *extensions.ReplicaSet
		expected        *ReplicaSetDetail
	}{
		{
			"test-namespace", "test-name",
			[]string{"get", "list", "list", "get", "list", "list"},
			&extensions.ReplicaSet{
				ObjectMeta: api.ObjectMeta{Name: "test-replicaset"},
				Spec: extensions.ReplicaSetSpec{
					Selector: &unversioned.LabelSelector{
						MatchLabels: map[string]string{},
					}},
			},
			&ReplicaSetDetail{
				ObjectMeta: common.ObjectMeta{Name: "test-replicaset"},
				TypeMeta:   common.TypeMeta{Kind: common.ResourceKindReplicaSet},
				PodInfo:    common.PodInfo{Warnings: []common.Event{}},
				PodList:    pod.PodList{Pods: []pod.Pod{}},
				EventList:  common.EventList{Events: []common.Event{}},
			},
		},
	}

	for _, c := range cases {
		fakeClient := testclient.NewSimpleFake(c.replicaSet, podList, eventList, c.replicaSet,
			podList, eventList)
		fakeHeapsterClient := FakeHeapsterClient{client: testclient.NewSimpleFake()}

		actual, _ := GetReplicaSetDetail(fakeClient, fakeHeapsterClient, c.namespace, c.name)

		actions := fakeClient.Actions()
		if len(actions) != len(c.expectedActions) {
			t.Errorf("Unexpected actions: %v, expected %d actions got %d", actions,
				len(c.expectedActions), len(actions))
			continue
		}

		for i, verb := range c.expectedActions {
			if actions[i].GetVerb() != verb {
				t.Errorf("Unexpected action: %+v, expected %s",
					actions[i], verb)
			}
		}

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("GetEvents(client,heapsterClient,%#v, %#v) == \ngot: %#v, \nexpected %#v",
				c.namespace, c.name, actual, c.expected)
		}
	}
}
