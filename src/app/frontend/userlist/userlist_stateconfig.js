import {actionbarViewName} from 'chrome/chrome_state';
import {breadcrumbsConfig} from 'common/components/breadcrumbs/breadcrumbs_service';
import {UserListController} from './userlist_controller';
import {stateName, stateUrl} from './userlist_state';
//import {stateName as userlogin} from 'userlogin/userlogin_state';
import {UserListActionBarController} from './userlistactionbar_controller';

/**
 * Configures states for the service view.
 *
 * @param {!ui.router.$stateProvider} $stateProvider
 * @ngInject
 */
export default function stateConfig($stateProvider) {
  $stateProvider.state(stateName, {
    url: stateUrl,
    resolve: {
      'userList': resolveUserList,
    },
    data: {
      [breadcrumbsConfig]: {
        'label': '用户',
      },
    },
    views: {
      '': {
        controller: UserListController,
        controllerAs: '$ctrl',
        templateUrl: 'userlist/userlist.html',
      },
      [actionbarViewName]: {
        controller: UserListActionBarController,
        controllerAs: 'ctrl',
        templateUrl: 'userlist/userlistactionbar.html',
      },
    },
  });
}

/**
 * @param {!angular.$resource} $resource
 * @return {!angular.$q.Promise}
 * @ngInject
 */
export function resolveUserList($resource) {
  /** @type {!angular.Resource<!backendApi.UserList>} */
  let resource = $resource('api/v1/users');
  return resource.get().$promise;
}
