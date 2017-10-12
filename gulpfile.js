
const gulp = require('gulp');
const glob = require('glob');
const rollup = require('gulp-better-rollup');
const uglify = require('rollup-plugin-uglify');
const babel = require('rollup-plugin-babel');
const path = require('path');
const cleanCSS = require('gulp-clean-css');
const concatCSS = require('gulp-concat-css');
const fs = require('fs');

function getEntryPoints(appPath) {
    var taskNames = [];
    var workingDirectory = path.join(appPath, 'static/js/*.entry.js');

    var entryPoints = glob.sync(workingDirectory)
        .map((componentDir) => {
            return path.basename(componentDir);
        });

    entryPoints.forEach((name) => {
        const entry = path.join(appPath, `static/js/${name}`);
        const tskName = createTask(name, entry, appPath);
        
        gulp.watch(entry, [tskName]);
        taskNames.push(tskName);
    });

    return taskNames;
}

function createTask(name, entry, appPath) {
    const taskName = getTaskName(name, appPath);
    const dest = path.join(appPath, "static/dist/js");
    const rollOptions = getRollupOptions(entry, name);

    gulp.task(taskName, () => {
        gulp.src(entry)
            .pipe(rollup(rollOptions, 'iife'))
            .pipe(gulp.dest(dest))
    });

    return taskName;
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

function getTasks() {
    var rollupTasks = [];
    const appFolders = ['./app/', './api/'];

    for (let i = 0; i < appFolders.length; i++) {
        const currFolder = appFolders[i];

        glob.sync(currFolder + '*').forEach((filePath) => {
            const staticPath = path.join(filePath, 'static');
            var appTasks = [];

            if (fs.existsSync(staticPath)) {
                appTasks = getEntryPoints(filePath);
            }

            rollupTasks = rollupTasks.concat(appTasks);
        });
    }

    var otherTasks = [];//['css-clean'];

    for (let i = 0; i < otherTasks.length; i++) {
        rollupTasks.push(otherTasks[i]);
    }

    return rollupTasks;
}

function getTaskName(name, appPath) {
    const appName = appPath.replace('./', '').replace('/', '.');
    const cleanName = name.replace('.entry.js', '');

    return `${appName}-roll-${cleanName}`;
}

gulp.task('css-clean', () => {
    return gulp.src('static/css/*.css')
        .pipe(concatCSS('bundle.css'))
        .pipe(cleanCSS())
        .pipe(gulp.dest('static/dist/css'))
});

//gulp.watch('static/css/*.css', ['css-clean']);

gulp.task('default', getTasks());