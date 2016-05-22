/**
 *  @fileOverview gulp tasks for compile backend applications.
 */
import gulp from 'gulp';
import lodash from 'lodash';
import path from 'path';

import conf from './conf';
import goCommand from './gocommand';

/**
 * Compiles backend application in development mode and places the binary in the serve
 * directory.
 */
gulp.task('backend', ['package-backend-source'], function(doneFn) {
  goCommand(
      [
        'build',
        // Install dependencies to speed up subsequent compilations.
        '-i',
        '-o',
        path.join(conf.paths.serve, conf.backend.binaryName),
        conf.backend.mainPackageName,
      ],
      doneFn);
});

/**
 * Compiles backend application in production mode for the current architecture and places the
 * binary in the dist directory.
 *
 * The production binary difference from development binary is only that it contains all
 * dependencies inside it and is targeted for a specific architecture.
 */
gulp.task('backend:prod', ['package-backend-source', 'clean-dist'], function() {
  let outputBinaryPath = path.join(conf.paths.dist, conf.backend.binaryName);
  return backendProd([[outputBinaryPath, conf.arch.default]]);
});

/**
 * Compiles backend application in production mode for all architectures and places the
 * binary in the dist directory.
 *
 * The production binary difference from development binary is only that it contains all
 * dependencies inside it and is targeted specific architecture.
 */
gulp.task('backend:prod:cross', ['package-backend-source', 'clean-dist'], function() {
  let outputBinaryPaths =
      conf.paths.distCross.map((dir) => path.join(dir, conf.backend.binaryName));
  return backendProd(lodash.zip(outputBinaryPaths, conf.arch.list));
});

/**
 * Moves all backend source files (app and tests) to a temporary package directory where it can be
 * applied go commands.
 *
 * This is required to consolidate test and app files into single directories and to make packaging
 * work.
 */
gulp.task('package- -source', function() {
  return gulp
      .src([path.join(conf.paths.backendSrc, '**/*'), path.join(conf.paths.backendTest, '**/*')])
      .pipe(gulp.dest(conf.paths.backendTmpSrc));
});

/**
 * @param {!Array<!Array<string>>} outputBinaryPathsAndArchs array of
 *    (output binary path, architecture) pairs
 * @return {!Promise}
 */
function backendProd(outputBinaryPathsAndArchs) {
  let promiseFn = (path, arch) => {
    return (resolve, reject) => {
      goCommand(
          [
            'build',
            '-a',
            '-installsuffix',
            'cgo',
            '-o',
            path,
            conf.backend.mainPackageName,
          ],
          (err) => {
            if (err) {
              reject(err);
            } else {
              resolve();
            }
          },
          {
            // Disable cgo package. Required to run on scratch docker image.
            CGO_ENABLED: '0',
            GOARCH: arch,
          });
    };
  };

  let goCommandPromises = outputBinaryPathsAndArchs.map(
      (pathAndArch) => new Promise(promiseFn(pathAndArch[0], pathAndArch[1])));

  return Promise.all(goCommandPromises);
}
