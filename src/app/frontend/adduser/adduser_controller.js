import {stateName as userlist} from 'userlist/userlist_state';

export default class AddUserController {
  /**
   * @param  {!ui.router.$state} $state
   * @ngInject
   */
  constructor($state, $resource) {
    /** @private {!ui.router.$state} */
    this.state_ = $state;

    /** @private {!angular.$resource} */
    this.resource_ = $resource;

    /** @export {!angular.FormController} */
    this.adduserForm;

    /** @export {string} */
    this.username;

    /** @export {string} */
    this.password;

    /** @export {number} */
    this.cpus;

    /** @export {number} */
    this.memory;

    /** @export {string}*/
    this.nameMaxLength = '50';

    /**@export {string}*/
    this.passwordMinLength = '5';
  }

  addUser() {
    if (!this.adduserForm.$valid) return;

    /** @type {!backendApi.UserCreate} */
    let userCreate = {
      name: this.username,
      password: this.password,
      cpus: this.cpus,
      memory: this.memory,
    };

    /** @type {!angular.Resource<!backendApi.UserCreate>} */
    let resource = this.resource_('api/v1/users');

    resource.save(userCreate, () => {
        this.state_.reload(userlist);
        this.state_.go(userlist);
    });
    
     
  }

  cancel() { this.state_.go(userlist); }
}
