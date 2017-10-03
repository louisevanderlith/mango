

function setupTasks(gulp){
    gulp.task('')
}

function cleanCSS(){
    return gulp.src('css/*.css')
    .pipe(concatCSS('bundle.css'))
    .pipe(cleanCSS())
    .pipe(gulp.dest('static/dist/css'))
}

function cleanJS(){

}

function cleanHTML(){

}