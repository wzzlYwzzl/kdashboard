import stateConfig from './user_stateconfig';
import Base64 from './user_base64';
import AuthenticationService from './user_authentication_service';

export default angular.module('kubernetesDashboard.user', ['ui.router', 'ngCookies', 'ngMessages', 'ngMaterial'])
    .config(stateConfig)
    .service('Base64', Base64)
    .service('AuthenticationService', AuthenticationService);
