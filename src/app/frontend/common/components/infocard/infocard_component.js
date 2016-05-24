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
 * Represents a card that can be used to display grouped detail info about any resource.
 * Usage:
 * <kd-info-card>
 *  <kd-info-card-header>Resource details</kd-info-card-header>
 *  <kd-info-card-group name="Details">
 *    <kd-info-card-entry title="Name">MyName</kd-info-card-entry>
 *    <kd-info-card-entry title="Namespace">MyNamespace</kd-info-card-entry>
 *  </kd-info-card-group>
 * </kd-info-card>
 *
 * @type {!angular.Component}
 */
export const infoCardComponent = {
  templateUrl: 'common/components/infocard/infocard.html',
  transclude: /** @type {undefined} TODO(bryk): Remove this when externs are fixed */ ({
    'header': 'kdInfoCardHeader',
    'section': 'kdInfoCardSection',
  }),
};
