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

import coresFilter from './cores_filter';
import memoryFilter from './memory_filter';
import middleEllipsisFilter from './middleellipsis_filter';
import relativeTimeFilter from './relativetime_filter';
import appConfigModule from '../appconfig/appconfig_module';

/**
 * Module containing common filters for the application.
 */
export default angular
    .module(
        'kubernetesDashboard.common.filters',
        [
          'ngMaterial',
          appConfigModule.name,
        ])
    .filter('middleEllipsis', middleEllipsisFilter)
    .filter('kdMemory', memoryFilter)
    .filter('kdCores', coresFilter)
    .filter('relativeTime', relativeTimeFilter);
