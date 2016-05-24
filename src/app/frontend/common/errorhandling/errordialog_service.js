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
 * Error dialog to display errors which could not be responded properly
 *
 * @final
 */
export default class ErrorDialog {
  /**
   * @param {!md.$dialog} $mdDialog
   * @ngInject
   */
  constructor($mdDialog) {
    /** @private {!md.$dialog} */
    this.mdDialog_ = $mdDialog;
  }

  /**
   * Opens a pop-up window that displays the error message
   *
   * @param {string} title
   * @param {string} text
   * @export
   */
  open(title, text) {
    let alert = this.mdDialog_.alert();
    alert.title(title);
    alert.textContent(text);
    alert.ok('Close');
    this.mdDialog_.show(alert);
  }
}
