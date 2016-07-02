/**
 * @fileoverview Root configuration file of the Gulp build system. It loads child modules which
 * define specific Gulp tasks.
 *
 * Learn more at: http://gulpjs.com
 */
import './build/check';
import './build/cluster';
import './build/backend';
import './build/build';
import './build/dependencies';
import './build/deploy';
import './build/index';
import './build/script';
import './build/serve';
import './build/style';
import './build/test';

// No business logic in this file.
