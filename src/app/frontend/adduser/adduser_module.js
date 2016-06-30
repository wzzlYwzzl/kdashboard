import stateConfig from './adduser_stateconfig';

export default angular
    .module(
        'kubernetesDashboard.adduser',
        [
          'ngMaterial',
          'ngResource',
          'ui.router',
        ])
    .config(stateConfig);
 