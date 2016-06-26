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

/**
 * @final
 */
export class UserController {
  /**
   * @param {!backendApi.Workloads} workloads
   * @ngInject
   */
  constructor($state, $scope, $rootScope) {
    /** @export {!backendApi.Workloads} */
    this.user;
    this.rootScope_ = $rootScope;
    this.state_ = $state;
    this.scope_ = $scope;
  }

  login() {
    // this.user.username = 'abc';
    // this.state_.go('workloads');
    this.message = this.user.username;
  }
}
