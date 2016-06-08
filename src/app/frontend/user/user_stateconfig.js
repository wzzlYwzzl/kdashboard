import {stateName, stateUrl} from './user_state';
import {UserController} from './user_controller';
import {actionbarViewName} from 'chrome/chrome_state';

/**
 * @param {!ui.router.$stateProvider} $stateProvider
 * @ngInject
 */
export default function stateConfig($stateProvider) {
  return $stateProvider.state(stateName, {
    url: stateUrl,
    view: {
      '': {
        controller: UserController,
        controllerAs: '$ctrl',
        templateUrl: 'user/user.html',
      },
      [actionbarViewName]: {},
    },
  });
}
