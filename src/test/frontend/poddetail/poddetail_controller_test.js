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

import {PodDetailController} from 'poddetail/poddetail_controller';
import podDetailModule from 'poddetail/poddetail_module';

describe('Pod detail controller', () => {

  beforeEach(() => { angular.mock.module(podDetailModule.name); });

  it('should initialize controller', angular.mock.inject(($controller) => {
    let data = {podDetail: {}};
    /** @type {!PodDetailController} */
    let ctrl = $controller(PodDetailController, {podDetail: data});

    expect(ctrl.podDetail).toBe(data);
  }));
});
