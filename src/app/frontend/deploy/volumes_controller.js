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
 * Controller for the volume mount directive.
 *
 * @final
 */
export default class VolumesMountController {
  /** @ngInject */
  constructor() {
    /**
     * Two way data binding from the scope.
     * @export {!Array<!backendApi.VolumeMount>}
     */
    this.volumesMount = [this.newVolumeMount_(this.volumeTypes[0])];

    /**
     * Initialized from the scope.
     * @export {!Array<string>}
     */
    this.volumeTypes;

    /**
     * Whether the mount is readonly
     * @export {boolean}
     */
    this.readonly;
  }

  /**
   * Call checks on volume mount:
   *  - adds new volume mount when last empty volume mount has been filled
   *  - validates volume mount
   * @param {!angular.FormController|undefined} volumeMountForm
   * @param {number} volumeMountIndex
   * @export
   */
  checkVolumeMount(volumeMountForm, volumeMountIndex) {
    this.addVolumeMountIfNeeed_();
    this.validateVolumeMount_(volumeMountForm, volumeMountIndex);
  }

  /**
   * @param {string} defaultType
   * @return {!backendApi.VolumeMount}
   * @private
   */
  newVolumeMount_(defaultType) {
    return {volumeType: defaultType, hostPath: null, containerPath: null};
  }

  /**
   * @export
   */
  addVolumeMountIfNeeed_() {
    let lastVolumeMount = this.volumesMount[this.volumesMount.length - 1];
    if (this.isVolumeMountFilled_(lastVolumeMount)) {
      this.volumesMount.push(this.newVolumeMount_(this.volumeTypes[0]));
    }
  }

  /**
   * Validates volume mount. In case when only one path is specified it is considered as invalid.
   * @param {!angular.FormController|undefined} volumeMountForm
   * @param {number} volumeMountIndex
   * @private
   */
  validateVolumeMount_(volumeMountForm, volumeMountIndex) {
    if (angular.isDefined(volumeMountForm)) {
      /** @type {!backendApi.VolumeMount} */
      let volumeMount = this.volumesMount[volumeMountIndex];

      /** @type {!angular.NgModelController} */
      let hostPathElem = volumeMountForm['hostpath'];
      /** @type {!angular.NgModelController} */
      let containerPathElem = volumeMountForm['containerpath'];

      /** @type {boolean} */
      let isValidHostPath = this.isVolumeMountFilledOrEmpty_(volumeMount) || !!volumeMount.hostPath;
      /** @type {boolean} */
      let isValidContainerPath =
          this.isVolumeMountFilledOrEmpty_(volumeMount) || !!volumeMount.containerPath;

      hostPathElem.$setValidity('empty', isValidHostPath);
      containerPathElem.$setValidity('empty', isValidContainerPath);
    }
  }

  /**
   * @param {number} index
   * @return {boolean}
   * @export
   */
  isRemovable(index) { return index !== (this.volumesMount.length - 1); }

  /**
   * @param {number} index
   * @export
   */
  remove(index) { this.volumesMount.splice(index, 1); }

  /**
   * Returns true when the given volume mount is filled by the user, i.e., is not empty.
   * @param {!backendApi.VolumeMount} volumeMount
   * @return {boolean}
   * @private
   */
  isVolumeMountFilled_(volumeMount) {
    return !!volumeMount.hostPath && !!volumeMount.containerPath;
  }

  /**
   * Returns true when the given volume mount is filled or empty (both path), false otherwise.
   * @param {!backendApi.VolumeMount} volumeMount
   * @return {boolean}
   * @private
   */
  isVolumeMountFilledOrEmpty_(volumeMount) {
    return !volumeMount.hostPath === !volumeMount.containerPath;
  }
}
