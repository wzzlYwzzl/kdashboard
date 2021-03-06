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

import stateConfig from './user_stateconfig';
import UserLoginService from './user_service';

/**
 * Module with a view that displays resources categorized as workloads, e.g., Replica Sets or
 * Deployments.
 */
export default angular
    .module(
        'kubernetesDashboard.user',
        [
          'ngMaterial',
          'ui.router',
        ])
    .config(stateConfig)
    .run(userloginConfig)
    .service('UserLoginService', UserLoginService);

/**
 * Configures event catchers for the state change.
 *
 * @param {!angular.Scope} $rootScope
 * @param {!angular.$http} $http
 * @ngInject
 */
function userloginConfig($rootScope, $http, UserLoginService) {
  $rootScope.$on('$stateChangeStart', function(event, toState) {
    // if (fromState.name !== 'userlogin' && toState.name !== 'userlogin' &&
    //   UserLoginService.loginuser.name === undefined) {
    //   // $state.go('userlogin');
    //   event.preventDefault();
    //   $state.go('userlogin');
    // }
    if (toState.name === 'deploy') {
      $http.get('/api/v1/users/allinfo').success(function(response) {
        UserLoginService.loginuser.cpususe = response.cpususe;
        UserLoginService.loginuser.memoryuse = response.memoryuse;
        //$state.reload('deploy');
      });
    }
  });
}
