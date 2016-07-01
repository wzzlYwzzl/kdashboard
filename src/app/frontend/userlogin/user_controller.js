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
import {stateName as users} from 'userlist/userlist_state';
import {stateName as workloads} from 'workloads/workloads_state';

/**
 * @final
 */
export class UserController {
  /**
   * @param {!backendApi.Workloads} workloads
   * @ngInject
   */
  constructor($scope, $state, $rootScope, $http, UserLoginService) {
    this.user;
    this.rootScope_ = $rootScope;
    this.state_ = $state;
    this.http_ = $http;
    this.scope_ = $scope;
    this.UserLoginService = UserLoginService;
  }

  login() {
    let username = this.user.username;
    let password = this.user.password;
    this.UserLoginService.setUser(username, password);
    let that = this;
    that.http_.post('/api/v1/login', {name: username, password: password}).success(function(response) {
       that.UserLoginService.loginuser.cpus = response.cpus;
       that.UserLoginService.loginuser.memory = response.memory;
       that.UserLoginService.loginuser.cpususe = response.cpususe;
       that.UserLoginService.loginuser.memoryuse = response.memoryuse;

      if (that.UserLoginService.loginuser.name === 'admin') {
        that.state_.go(users);
      } else {
        that.state_.go(workloads, {name: that.UserLoginService.loginuser.name});
      }
      // that.rootScope_.user = that.user;
    });
  }
}
