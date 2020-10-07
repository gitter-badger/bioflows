var gulp = require('gulp');

var shell = require('gulp-shell');


gulp.task('make-clean',function (){
   shell.task(['make clean']);
});

gulp.task('make-html',function (){
   shell.task(['make html']);
});

gulp.task('watch',function (){
   gulp.watch('*',gulp.series('make-clean','make-html'));
});

gulp.task('default',gulp.series('watch'))
