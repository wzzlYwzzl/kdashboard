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

package validation

import (
	"testing"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/unversioned/testclient"
	"k8s.io/kubernetes/pkg/runtime"
)

func TestValidateName(t *testing.T) {
	spec := &AppNameValiditySpec{
		Namespace: "foo-namespace",
		Name:      "foo-name",
	}
	cases := []struct {
		spec     *AppNameValiditySpec
		objects  []runtime.Object
		expected bool
	}{
		{
			spec,
			nil,
			true,
		},
		{
			spec,
			[]runtime.Object{&api.ReplicationController{}},
			false,
		},
		{
			spec,
			[]runtime.Object{&api.Service{}},
			false,
		},
		{
			spec,
			[]runtime.Object{&api.ReplicationController{}, &api.Service{}},
			false,
		},
	}

	for _, c := range cases {
		testClient := testclient.NewSimpleFake(c.objects...)
		validity, _ := ValidateAppName(c.spec, testClient)
		if validity.Valid != c.expected {
			t.Errorf("Expected %#v validity to be %#v for objects %#v, but was %#v\n",
				c.spec, c.expected, c.objects, validity)
		}
	}
}
