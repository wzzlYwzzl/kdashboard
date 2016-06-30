export default class UserLoginService {
  constructor($http, $rootScope) {
    this.http_ = $http;
    this.rootScope_ = $rootScope;
    this.loginuser;
  }

  setUser(username, password) {
    this.loginuser = {
      name: username,
      password: password,
    };
  }
}
