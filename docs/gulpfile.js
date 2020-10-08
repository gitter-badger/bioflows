const {series,watch} = require('gulp');
var shell = require('gulp-shell');

function make_clean(cb){
   shell(['make clean']);
   cb();
}

function make_html(cb){
   shell(['make html']);
   cb();
}


exports.default = function(){
   watch("source/*",series(shell.task('make clean'),shell.task('make html')));
}