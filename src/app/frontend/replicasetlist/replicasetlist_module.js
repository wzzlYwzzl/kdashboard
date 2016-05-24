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

import stateConfig from './replicasetlist_stateconfig';
import filtersModule from 'common/filters/filters_module';
import componentsModule from 'common/components/components_module';
import {replicaSetCardComponent} from './replicasetcard_component';
import {replicaSetCardListComponent} from './replicasetcardlist_component';
import replicaSetDetailModule from 'replicasetdetail/replicasetdetail_module';

/**
 * Angular module for the Replica Set list view.
 *
 * The view shows Replica Set running in the cluster and allows to manage them.
 */
export default angular
    .module(
        'kubernetesDashboard.replicaSetList',
        [
          'ngMaterial',
          'ngResource',
          'ui.router',
          filtersModule.name,
          componentsModule.name,
          replicaSetDetailModule.name,
        ])
    .config(stateConfig)
    .component('kdReplicaSetCardList', replicaSetCardListComponent)
    .component('kdReplicaSetCard', replicaSetCardComponent);
