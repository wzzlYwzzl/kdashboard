import AddUserController from './adduser_controller';
import {stateName, stateUrl} from './adduser_state';

/**
 * Configures states for the adduser view.
 *
 * @param {!ui.router.$stateProvider} $stateProvider
 * @ngInject
 */
export default function stateConfig($stateProvider) {
  $stateProvider.state(stateName, {
    controller: AddUserController,
    controllerAs: 'ctrl',
    url: stateUrl,
    templateUrl: 'adduser/adduser.html',
  });
}
