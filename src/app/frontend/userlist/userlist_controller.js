/**
 * Controller for the user list view.
 *
 * @final
 */
export class UserListController {
  /**
 * @param {!backendApi.UserList} userList
 * @ngInject
 */
  constructor(userList) {
    /** @export {!backendApi.UserList} */
    this.userList = userList;
  }
}
