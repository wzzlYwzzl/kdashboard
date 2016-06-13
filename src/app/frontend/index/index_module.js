import config from './index_config';
import interceptor from './index_interceptor_config';

export default angular
    .module(
        'Dashboard',
        [
          'ng-admin',
        ])
    .config(config)
    .config(interceptor);
