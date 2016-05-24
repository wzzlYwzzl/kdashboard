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

import NamespaceDialogController from 'deploy/createnamespace_controller';
import deployModule from 'deploy/deploy_module';

describe('Create-Namespace dialog', () => {

  let ctrl;
  let httpBackend;
  beforeEach(() => {
    angular.mock.module(deployModule.name);
    angular.mock.inject(($controller, $httpBackend, $mdDialog, $log, _errorDialog_) => {
      ctrl = $controller(NamespaceDialogController, {
        'namespaces': [],
        'mdDialog_': $mdDialog,
        'log_': $log,
        'errorDialog_': _errorDialog_,
      });
      httpBackend = $httpBackend;
    });
  });

  it('should disable ok-button if namespace already exists', () => {
    // given one existing namespace
    ctrl.namespaces = ['my-namespace'];

    // when entering the same name agian
    ctrl.namespace = 'my-namespace';

    // then button is disabled
    let result = ctrl.isDisabled();
    expect(result).toBeTruthy();
  });

  it('should be enable ok-button after some input', () => {

    // when entering something correct
    ctrl.namespace = 'my-namespace';

    // then button is enabled
    let result = ctrl.isDisabled();
    expect(result).toBeFalsy();
  });

  /**
   * Test for the namespace validation pattern
   *
   * The pattern is copied from the error message of the API server. The beginning is matched with ^
   * the end with $. Full pattern used for validation:
   *
   * ^<pattern from error message>$
   */
  it('should validate proper names as correct', () => {
    // given the k8s rule for namespace names
    let rule = ctrl.namespacePattern;

    // then the following names should be accepted
    expect('mynamspace'.match(rule)).toBeDefined();
    expect('my-namspace-with-dashes'.match(rule)).toBeDefined();
    expect('my-namspace-with-numbers-234'.match(rule)).toBeDefined();
  });

  it('should validate inproper names as incorrect', () => {
    // given the k8s rule for namespace names
    let rule = ctrl.namespacePattern;

    // then the following names should be rejected
    expect('mynamspace-with-illegal-chars-§$'.match(rule)).toBeNull();
    expect('-mynamspace-with-dash-prefix'.match(rule)).toBeNull();
    expect('mynamspace-with-dash-suffix-'.match(rule)).toBeNull();
    expect('mynamspace-with-german-umlaut-ÖÄ'.match(rule)).toBeNull();
    expect('my namspace with spaces'.match(rule)).toBeNull();
    expect('  '.match(rule)).toBeNull();
  });

  it('should not submit if the form has validation errors', () => {
    // given a validation error in the html form
    ctrl.namespaceForm = {};
    ctrl.namespaceForm.$valid = false;

    // when trying to submit
    ctrl.createNamespace();

    // then form data was not sent to backend (thus flush will throw error)
    expect(httpBackend.flush).toThrow();
  });

  it('should hide creation dialog and open an error dialog if namespace cannot be created', () => {
    spyOn(ctrl.errorDialog_, 'open');
    spyOn(ctrl.mdDialog_, 'hide');
    spyOn(ctrl.log_, 'info');
    ctrl.namespaceForm = {};
    ctrl.namespaceForm.$valid = true;
    /** @type {string} */
    let errorMessage = 'Something bad happened';
    // return an erranous response
    httpBackend.expectPOST('api/v1/namespaces').respond(500, errorMessage);
    // when
    ctrl.createNamespace();
    httpBackend.flush();
    // expect
    expect(ctrl.mdDialog_.hide).toHaveBeenCalled();
    expect(ctrl.errorDialog_.open).toHaveBeenCalledWith('Error creating namespace', errorMessage);
    expect(ctrl.log_.info).toHaveBeenCalled();

  });
});
