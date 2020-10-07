var gulp = require("gulp");
var shell = require('gulp-shell');
// This compiles new binary with source change
gulp.task("make-clean", shell.task([
   'make clean'
]));
gulp.task("make-html", ["make-clean"], shell.task([
   'make html'
]))
gulp.task('watch', function() {
// Watch the source code for all changes
   gulp.watch("*", ['make-clean', 'make-html']);
});
gulp.task('default', ['watch']);