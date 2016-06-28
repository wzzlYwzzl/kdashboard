//import {stateName as adduser} from 'adduser/adduser_state';

/**
 * @final
 */
export class UserListActionBarController {
  /**
   * @param {!ui.router.$state} $state
   * @ngInject
   */
  constructor($state) {
    /** @private {!ui.router.$state} */
    this.state_ = $state;
  }

  /**
   * @export
   */
  redirectToAddUserPage() { this.state_.go(uselist); }
}
