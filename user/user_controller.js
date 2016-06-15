// import AuthenticationService from './user_authentication_service';

/**
 * @final
 */
export class UserController {
  /**
   * [constructor description]
   * @ngInject
   */
  constructor(user) {
    this.user = user;
  }

/*
  login() {
    scope_.dataLoading = true;
    auth_.login(scope_.username, scope_.password, function(response) {
      if (response.success) {
        auth_.setCredentials(scope_.username, scope_.password);
        state_.go('workloads');
      } else {
        scope_.error = response.message;
        scope_.dataLoading = false;
      }
    });
  }
  */
}
