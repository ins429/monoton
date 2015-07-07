var gulp = require('gulp');
var sourcemaps = require('gulp-sourcemaps');
var babel = require('gulp-babel');
var concat = require('gulp-concat');
var react = require('gulp-react');
var browserify = require('gulp-browserify');

gulp.task('default', function () {
    return gulp.src('src/**/*.jsx')
        .pipe(sourcemaps.init())
        .pipe(babel())
        .pipe(react({harmony: true, es6module: true}))
        .pipe(browserify({
            insertGlobals : true,
            debug : !gulp.env.production
        }))
        .pipe(concat('app.js'))
        .pipe(sourcemaps.write('.'))
        .pipe(gulp.dest('dist'));
});
