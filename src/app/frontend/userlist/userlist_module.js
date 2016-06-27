import stateConfig from './userlist_stateconfig';
import {userCardListComponent} from './usercardlist_component';

/**
 * Angular module for the Pods list view.
 *
 * The view shows Pods running in the cluster and allows to manage them.
 */
export default angular
    .module(
        'kubernetesDashboard.userList',
        [
          'ngMaterial',
          'ngResource',
          'ui.router',
        ])
    .config(stateConfig)
    .component('kdUserCardList', userCardListComponent);
