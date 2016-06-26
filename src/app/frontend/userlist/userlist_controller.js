/**
 * Controller for the pod list view.
 *
 * @final
 */
export class UserListController {
    /**
   * @param {!backendApi.PodList} podList
   * @ngInject
   */
    constructor(userList) {
        /** @export {!backendApi.UserList} */
        this.userList = userList;
    }
}