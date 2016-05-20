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
	"k8s.io/kubernetes/pkg/api"

	"github.com/kubernetes/dashboard/resource/common"
)

// ToService returns api service object based on kubernetes service object
func ToService(service *api.Service) Service {
	return Service{
		ObjectMeta:       common.NewObjectMeta(service.ObjectMeta),
		TypeMeta:         common.NewTypeMeta(common.ResourceKindService),
		InternalEndpoint: common.GetInternalEndpoint(service.Name, service.Namespace, service.Spec.Ports),
		// TODO(maciaszczykm): Fill ExternalEndpoints with data.
		Selector:  service.Spec.Selector,
		ClusterIP: service.Spec.ClusterIP,
	}
}

// ToServiceDetails returns api service object based on kubernetes service object
func ToServiceDetail(service *api.Service) ServiceDetail {
	return ServiceDetail{
		ObjectMeta:       common.NewObjectMeta(service.ObjectMeta),
		TypeMeta:         common.NewTypeMeta(common.ResourceKindService),
		InternalEndpoint: common.GetInternalEndpoint(service.Name, service.Namespace, service.Spec.Ports),
		// TODO(maciaszczykm): Fill ExternalEndpoints with data.
		Selector:  service.Spec.Selector,
		ClusterIP: service.Spec.ClusterIP,
		Type:      service.Spec.Type,
	}
}
