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

import {actionbarViewName} from 'chrome/chrome_state';
import {stateName} from './user_state';
import {stateUrl} from './user_state';
import {UserController} from './user_controller';
// import {WorkloadsActionBarController} from './workloadsactionbar_controller';

/**
 * @param {!ui.router.$stateProvider} $stateProvider
 * @ngInject
 */
export default function stateConfig($stateProvider) {
  $stateProvider.state(stateName, {
    url: stateUrl,
    data: {
      'kdBreadcrumbs': {
        'label': ' ',
      },
    },
    views: {
      '': {
        controller: UserController,
        controllerAs: 'ctrl',
        templateUrl: 'userlogin/user.html',
      },
      [actionbarViewName]: {},
    },
  });
}
