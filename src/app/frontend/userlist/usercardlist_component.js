export class UserCardListController {
  /**
 * @ngInject
 * @param {!ui.router.$state} $state
 */
  constructor($state) {
    /**
         * List of users. Initialized from the scope.
         * @export {!backendApi.UserList}
         */
    this.userList;

    /** @private {!ui.router.$state} */
    this.state_ = $state;
  }
}

/**
 * Definition object for the component that displays user list card.
 *
 * @type {!angular.Component}
 */
export const userCardListComponent = {
  templateUrl: 'userlist/usercardlist.html',
  controller: UserCardListController,
  bindings: {
    /** {!backendApi.UserList} */
    'userList': '<',
    /** {boolean} */
    'selectable': '<',
    /** {boolean} */
    'withStatuses': '<',
  },
};
