export default class UserLoginService {
  constructor() { this.loginuser = {}; }

  setUser(username, password) {
    this.loginuser = {
      name: username,
      password: password,
    };
  }
}
