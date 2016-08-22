 /**
  *  @final
  */
export default class UserLoginService {

    /**
     * @ngInject
     */
  constructor() { 
/** @export  {Object}*/
    this.loginuser = {}; 
}

/**
   * @param {!string} username
   * @param {!string} password
   * @export
   */
  setUser(username, password) {
    this.loginuser = {
      name: username,
      password: password,
    };
  }
}
