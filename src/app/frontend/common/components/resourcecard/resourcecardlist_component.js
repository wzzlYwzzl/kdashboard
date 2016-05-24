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
export class ResourceCardListController {
  /**
   * @ngInject
   */
  constructor() {
    /**
     * Whether to make list items selectable.
     * Initialized from a bingind.
     * @export {boolean|undefined}
     */
    this.selectable;

    /**
     * Whether to make show statuses for list items.
     * Initialized from a bingind.
     * @export {boolean|undefined}
     */
    this.withStatuses;

    /**
     * @private {!./resourcecardheadercolumns_component.ResourceCardHeaderColumnsController}
     */
    this.headerColumns_;
  }

  /**
   * @param {!./resourcecardheadercolumns_component.
   *     ResourceCardHeaderColumnsController} headerColumns
   */
  setHeaderColumns(headerColumns) {
    if (this.headerColumns_) {
      throw new Error('Header columns already set');
    }
    this.headerColumns_ = headerColumns;
  }

  /**
   * @param {!angular.JQLite} columnElement
   * @param {number} index
   */
  sizeBodyColumn(columnElement, index) { this.headerColumns_.sizeBodyColumn(columnElement, index); }
}

/**
 * Resource card list component. Used for displaying lists of resources.
 * It consists of a set of header columns which configure sizing of body columns. The number
 * of header and body columns must be exactly the same.
 *
 * Sample usage:
 *   <kd-resource-card-list selectable="true" with-statuses="true">
 *     <kd-resource-card-header-columns>
 *       <kd-resource-card-header-column size="medium" grow="2">
 *         Name
 *       </kd-resource-card-header-column>
 *       <kd-resource-card-header-column size="small" grow="nogrow">
 *         Age
 *       </kd-resource-card-header-column>
 *       <kd-resource-card-header-column>
 *         Labels
 *       </kd-resource-card-header-column>
 *     </kd-resource-card-header-columns>
 *
 *     <kd-resource-card ng-repeat="myResource in $ctrl.foo">
 *       <kd-resource-card-status>Foo status</kd-resource-card-status>
 *         <kd-resource-card-columns>
 *           <kd-resource-card-column>
 *              {{myResource.name}}
 *           </kd-resource-card-column>
 *           <kd-resource-card-column>
 *              {{myResource.age}}
 *           </kd-resource-card-column>
 *           <kd-resource-card-column>
 *              {{myResource.labels}}
 *           </kd-resource-card-column>
 *         </kd-resource-card-columns>
 *       <kd-resource-card-footer>Foo footer</kd-resource-card-footer>
 *     </kd-resource-card>
 *   </kd-resource-card-list>
 *
 * @type {!angular.Component}
 */
export const resourceCardListComponent = {
  templateUrl: 'common/components/resourcecard/resourcecardlist.html',
  transclude: true,
  controller: ResourceCardListController,
  bindings: {
    /** {boolean|undefined} whether to make list items selectable */
    'selectable': '<',
    /** {boolean|undefined} whether to show statuses for list items */
    'withStatuses': '<',
  },
  bindToController: true,
};
