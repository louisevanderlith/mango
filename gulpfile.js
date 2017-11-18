
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
        const tskName = createJSTask(name, entry, appPath);

        gulp.watch(entry, [tskName]);
        taskNames.push(tskName);
    });

    return taskNames;
}

function createJSTask(name, entry, appPath) {
    const taskName = getJSTaskName(name, appPath);
    const dest = path.join(appPath, "static/dist/js");
    const rollOptions = getRollupOptions(entry, name);

    gulp.task(taskName, () => {
        gulp.src(entry)
            .pipe(rollup(rollOptions, 'iife'))
            .pipe(gulp.dest(dest))
    });

    return taskName;
}

function createCSSTask(appPath) {
    const name = getNameFromPath(appPath);
    const taskName = `${name}-css`;
    const fullPath = path.join(appPath, 'static/css/*.css');
    const destPath = path.join(appPath, 'static/dist/css');

    gulp.task(taskName, () => {
        gulp.src(fullPath)
            .pipe(concatCSS('bundle.css'))
            .pipe(cleanCSS())
            .pipe(gulp.dest(destPath))
    });

    gulp.watch(fullPath, [taskName]);

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
            jquery: 'https://code.jquery.com/jquery-3.2.1.min.js'
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
                var cssTask = createCSSTask(filePath);

                appTasks.push(cssTask);
            }

            rollupTasks = rollupTasks.concat(appTasks);
        });
    }

    return rollupTasks;
}

function getJSTaskName(name, appPath) {
    const appName = getNameFromPath(appPath);
    const cleanName = name.replace('.entry.js', '');

    return `${appName}-roll-${cleanName}`;
}

function getNameFromPath(appPath){
    return appPath.replace('./', '').replace('/', '.');
}

gulp.task('default', getTasks());