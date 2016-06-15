import {stateName, stateUrl} from './user_state';
//import {UserController} from './user_controller';
import {actionbarViewName} from 'chrome/chrome_state';
//import {breadcrumbsConfig} from 'common/components/breadcrumbs/breadcrumbs_service';

/**
 * @param {!ui.router.$stateProvider} $stateProvider
 * @ngInject
 */
export default function stateConfig($stateProvider) {
  return $stateProvider.state(stateName, {
    url: stateUrl,
 //   resolve: {
 //     user : resolveUsers,
  //  },
    view: {
      '': {
       // controller: UserController,
        //controllerAs: 'ctrl',
        templateUrl: 'user/user.html',
      },
      [actionbarViewName]: {},
    },
  });
}
/*
export function resolveUsers() {
  let user = {username: 'test', password: 'test'};
  return  user;
}
*/