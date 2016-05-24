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
 * Controller for the port mappings directive.
 *
 * @final
 */
export default class PortMappingsController {
  /** @ngInject */
  constructor() {
    /**
     * Two way data binding from the scope.
     * @export {!Array<!backendApi.PortMapping>}
     */
    this.portMappings = [this.newEmptyPortMapping_(this.protocols[0])];

    /**
     * Initialized from the scope.
     * @export {!Array<string>}
     */
    this.protocols;

    /**
     * Initialized from the scope.
     * @export {boolean}
     */
    this.isExternal;
  }

  /**
   * Call checks on port mapping:
   *  - adds new port mapping when last empty port mapping has been filled
   *  - validates port mapping
   * @param {!angular.FormController|undefined} portMappingForm
   * @param {number} portMappingIndex
   * @export
   */
  checkPortMapping(portMappingForm, portMappingIndex) {
    this.addProtocolIfNeeed_();
    this.validatePortMapping_(portMappingForm, portMappingIndex);
  }

  /**
   * @param {string} defaultProtocol
   * @return {!backendApi.PortMapping}
   * @private
   */
  newEmptyPortMapping_(defaultProtocol) {
    return {port: null, targetPort: null, protocol: defaultProtocol};
  }

  /**
   * @export
   */
  addProtocolIfNeeed_() {
    let lastPortMapping = this.portMappings[this.portMappings.length - 1];
    if (this.isPortMappingFilled_(lastPortMapping)) {
      this.portMappings.push(this.newEmptyPortMapping_(this.protocols[0]));
    }
  }

  /**
   * Validates port mapping. In case when only one port is specified it is considered as invalid.
   * @param {!angular.FormController|undefined} portMappingForm
   * @param {number} portIndex
   * @private
   */
  validatePortMapping_(portMappingForm, portIndex) {
    if (angular.isDefined(portMappingForm)) {
      /** @type {!backendApi.PortMapping} */
      let portMapping = this.portMappings[portIndex];

      /** @type {!angular.NgModelController} */
      let portElem = portMappingForm['port'];
      /** @type {!angular.NgModelController} */
      let targetPortElem = portMappingForm['targetPort'];

      /** @type {boolean} */
      let isValidPort = this.isPortMappingFilledOrEmpty_(portMapping) || !!portMapping.port;
      /** @type {boolean} */
      let isValidTargetPort =
          this.isPortMappingFilledOrEmpty_(portMapping) || !!portMapping.targetPort;

      portElem.$setValidity('empty', isValidPort);
      targetPortElem.$setValidity('empty', isValidTargetPort);
    }
  }

  /**
   * @param {number} index
   * @return {boolean}
   * @export
   */
  isRemovable(index) { return index !== (this.portMappings.length - 1); }

  /**
   * @param {number} index
   * @export
   */
  remove(index) { this.portMappings.splice(index, 1); }

  /**
   * Returns true when the given port mapping is filled by the user, i.e., is not empty.
   * @param {!backendApi.PortMapping} portMapping
   * @return {boolean}
   * @private
   */
  isPortMappingFilled_(portMapping) { return !!portMapping.port && !!portMapping.targetPort; }

  /**
   * Returns true when the given port mapping is filled or empty (both ports), false otherwise.
   * @param {!backendApi.PortMapping} portMapping
   * @return {boolean}
   * @private
   */
  isPortMappingFilledOrEmpty_(portMapping) { return !portMapping.port === !portMapping.targetPort; }
}
