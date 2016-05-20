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

// PodInfo represents aggregate information about controller's pods.
type PodInfo struct {
	// Number of pods that are created.
	Current int `json:"current"`

	// Number of pods that are desired.
	Desired int `json:"desired"`

	// Number of pods that are currently running.
	Running int `json:"running"`

	// Number of pods that are currently waiting.
	Pending int `json:"pending"`

	// Number of pods that are failed.
	Failed int `json:"failed"`

	// Unique warning messages related to pods in this resource.
	Warnings []Event `json:"warnings"`
}

// GetPodInfo returns aggregate information about a group of pods.
func GetPodInfo(current int, desired int, pods []api.Pod) PodInfo {
	result := PodInfo{
		Current:  current,
		Desired:  desired,
		Warnings: make([]Event, 0),
	}

	for _, pod := range pods {
		switch pod.Status.Phase {
		case api.PodRunning:
			result.Running++
		case api.PodPending:
			result.Pending++
		case api.PodFailed:
			result.Failed++
		}
	}

	return result
}
