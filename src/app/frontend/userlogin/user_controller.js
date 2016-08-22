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
   * @param {!ui.router.$state} $state
   * @param {!angular.$http} $http
   * @param {!./../userlogin/user_service} UserLoginService
   * @ngInject
   */
  constructor($state, $http, UserLoginService) {
    this.user = {};
    this.user.username = '';
    this.user.password = '';
    this.state_ = $state;
    this.http_ = $http;
    this.LoginService = UserLoginService;
  }

/**
   * @export
   */
  login() {
    let username = this.user.username;
    let password = this.user.password;
    this.LoginService.setUser(username, password);
    let that = this;
    that.http_.post('/api/v1/login', {name: username, password: password})
        .success(function(response) {
          that.LoginService.loginuser.cpus = response.cpus;
          that.LoginService.loginuser.memory = response.memory;
          that.LoginService.loginuser.cpususe = response.cpususe;
          that.LoginService.loginuser.memoryuse = response.memoryuse;

          if (that.LoginService.loginuser.name === 'admin') {
            that.state_.go(users);
          } else {
            that.state_.go(workloads, {name: that.LoginService.loginuser.name});
          }
          // that.rootScope_.user = that.user;
        });
  }
}
