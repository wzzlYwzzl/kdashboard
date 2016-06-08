export default class AuthenticationService {
  /**
   * @param  {!angular.$http} $http
   * @param  {!angular.$cookieStore} $cookieStore
   * @param  {!angular.$rootScope} $rootScope
   * @param  {!angular.$rootScope} $timeout
   * @ngInject
   */
  constructor($http, $cookieStore, $rootScope, $timeout, Base64) {
    this.http_ = $http;
    this.cookieStore_ = $cookieStore;
    this.rootScope_ = $rootScope;
    this.timeout_ = $timeout;
    this.base64_ = Base64;
  }

  login(username, password, callback) {
    /* Dummy authentication for testing, uses $timeout to simulate api call
             ----------------------------------------------*/
    $timeout(function() {
      let response = {success: username === 'test' && password === 'test'};
      if (!response.success) {
        response.message = 'Username or password is incorrect';
      }
      callback(response);
    }, 1000);

    /* Use this for real authentication
     ----------------------------------------------*/
    //$http.post('/api/authenticate', { username: username, password: password })
    //    .success(function (response) {
    //        callback(response);
    //    });
  }

  letCredentials(username, password) {
    let authdata = base64_.encode(`${username}:${password}`);

    rootScope_.globals = {
      currentUser: {
        username: username,
        authdata: authdata,
      },
    };

    http_.defaults.headers.common['Authorization'] = `Basic ${authdata}`;
    cookieStore_.put('globals', rootScope_.globals);
  }

  clearCredentials() {
    rootScope_.globals = {};
    cookieStore_.remove('globals');
    http_.defaults.headers.common.Authorization = 'Basic';
  }
}
