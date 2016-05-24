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
	"reflect"
	"testing"

	"github.com/kubernetes/dashboard/resource/common"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/client/unversioned/testclient"
	"k8s.io/kubernetes/pkg/labels"
)

func TestGetDaemonSetPodInfo(t *testing.T) {
	cases := []struct {
		daemonset *extensions.DaemonSet
		pods      []api.Pod
		expected  common.PodInfo
	}{
		{
			&extensions.DaemonSet{
				Status: extensions.DaemonSetStatus{},
				Spec:   extensions.DaemonSetSpec{},
			},
			[]api.Pod{
				{
					Status: api.PodStatus{
						Phase: api.PodRunning,
					},
				},
			},
			common.PodInfo{
				Running: 1,
				Pending: 0,
				Failed:  0,
			},
		},
	}

	for _, c := range cases {
		actual := getDaemonSetPodInfo(c.daemonset, c.pods)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("getReplicaSetPodInfo(%#v, %#v) == \n%#v\nexpected \n%#v\n",
				c.daemonset, c.pods, actual, c.expected)
		}
	}
}

func TestGetServicesForDeletionforDS(t *testing.T) {
	cases := []struct {
		labelSelector   labels.Selector
		DaemonSetList   *extensions.DaemonSetList
		expected        *api.ServiceList
		expectedActions []string
	}{
		{
			labels.SelectorFromSet(map[string]string{"app": "test"}),
			&extensions.DaemonSetList{
				Items: []extensions.DaemonSet{
					{Spec: extensions.DaemonSetSpec{
						Selector: &unversioned.LabelSelector{
							MatchLabels: map[string]string{"app": "test"},
						},
					},
					},
				},
			},
			&api.ServiceList{
				Items: []api.Service{
					{Spec: api.ServiceSpec{Selector: map[string]string{"app": "test"}}},
				},
			},
			[]string{"list", "list"},
		},
		{
			labels.SelectorFromSet(map[string]string{"app": "test"}),
			&extensions.DaemonSetList{
				Items: []extensions.DaemonSet{
					{Spec: extensions.DaemonSetSpec{
						Selector: &unversioned.LabelSelector{
							MatchLabels: map[string]string{"app": "test"},
						},
					},
					},
					{Spec: extensions.DaemonSetSpec{
						Selector: &unversioned.LabelSelector{
							MatchLabels: map[string]string{"app": "test"},
						},
					},
					},
				},
			},
			&api.ServiceList{
				Items: []api.Service{
					{Spec: api.ServiceSpec{Selector: map[string]string{"app": "test"}}},
				},
			},
			[]string{"list"},
		},
		{
			labels.SelectorFromSet(map[string]string{"app": "test"}),
			&extensions.DaemonSetList{},
			&api.ServiceList{
				Items: []api.Service{
					{Spec: api.ServiceSpec{Selector: map[string]string{"app": "test"}}},
				},
			},
			[]string{"list"},
		},
	}

	for _, c := range cases {
		fakeClient := testclient.NewSimpleFake(c.DaemonSetList, c.expected)

		getServicesForDSDeletion(fakeClient, c.labelSelector, "mock")

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
	}
}
