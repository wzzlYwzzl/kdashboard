/**
 * test for gulp.
 */

import gulp from 'gulp';
import path from 'path';

gulp.task('test-gulp', function() {
    console.log(path.join(__dirname,'../'));
});
