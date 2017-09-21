
const glob = require('glob');
const gulp = require('gulp');
const rollup = require('gulp-better-rollup');
const uglify = require('rollup-plugin-uglify');
const babel = require('rollup-plugin-babel');
const path = require('path');
const cleanCSS = require('gulp-clean-css');
const concatCSS = require('gulp-concat-css');

var entryPoints = [];

function getEntryPoints() {
    var entryPoints = glob.sync('static/js/*.entry.js')
        .map((componentDir) => {
            return path.basename(componentDir);
        });

    entryPoints.forEach((name) => {
        console.log(`Assembling task for ${name}`);
        const entry = `./static/js/${name}`;

        createTask(name, entry);
        createWatch(name, entry);
    });
}

function createTask(name, entry) {
    const taskName = getTaskName(name);
    const dest = "./static/dist/js";
    const rollOptions = getRollupOptions(entry, name);

    gulp.task(taskName, () => {
        gulp.src(entry)
            .pipe(rollup(rollOptions, 'iife'))
            .pipe(gulp.dest(dest))
    });
}

function getRollupOptions(entry, name) {
    return {
        entry: entry,
        format: 'iife',
        moduleName: name,
        globals: {
            jquery: 'jquery'
        },
        external: ['jquery'],
        paths: {
            jquery: 'http://code.jquery.com/jquery-3.2.1.min.js'
        },
        plugins: [
            babel({
                exclude: 'node_modules/**'
            }),
            uglify()
        ]
    };
}

function createWatch(name, entry) {
    gulp.watch(entry, [getTaskName(name)]);
}

function getTasks() {
    getEntryPoints();

    var rollupTasks = entryPoints.map((name) => {
        return getTaskName(name);
    });

    var otherTasks = ['css-clean'];

    for (let i = 0; i < otherTasks.length; i++) {
        rollupTasks.push(otherTasks[i]);
    }

    return rollupTasks;
}

function getTaskName(name) {
    const cleanName = name.replace('.entry.js', '');

    return `roll-${cleanName}`;
}

gulp.task('css-clean', () => {
    return gulp.src('static/css/*.css')
        .pipe(concatCSS('bundle.css'))
        .pipe(cleanCSS())
        .pipe(gulp.dest('static/dist/css'))
});

gulp.watch('static/css/*.css', ['css-clean'])

gulp.task('default', getTasks());