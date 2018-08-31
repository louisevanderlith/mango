const gulp = require('gulp');
const glob = require('glob');
const rollup = require('gulp-better-rollup');
const uglify = require('rollup-plugin-uglify');
const babel = require('rollup-plugin-babel');
const path = require('path');
const cleanCSS = require('gulp-clean-css');
const concatCSS = require('gulp-concat-css');
const fs = require('fs');
const ugh = require('uglify-es').minify;

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

        gulp.watch(entry, gulp.series([tskName]));
        taskNames.push(tskName);
    });

    return taskNames;
}

function getFontsPoints(appPath){
    var taskNames = [];
    var workingDirectory = path.join(appPath, 'static/fonts/glyphicons-halflings-regular.{eot,svg,ttf,woff,woff2}');

    var entryPoints = glob.sync(workingDirectory)
        .map((componentDir) => {
            return path.basename(componentDir);
        });

    entryPoints.forEach((name) => {
        const entry = path.join(appPath, `static/fonts/${name}`);
        const tskName = createFontsTask(name, entry, appPath);

        gulp.watch(entry, gulp.series([tskName]));
        taskNames.push(tskName);
    });

    return taskNames;
}

function createJSTask(name, entry, appPath) {
    const taskName = getJSTaskName(name, appPath);
    const dest = path.join(appPath, "static/dist/js");
    const rollOptions = getRollupOptions(entry, name);

    gulp.task(taskName, gulp.series(() => {
        gulp.src(entry)
            .pipe(rollup(rollOptions, 'iife'))
            .on('error', (err) => {
                console.error("Entry: %s, Task: %s. Details: %s", entry, taskName, err);
            })
            .pipe(gulp.dest(dest));
    }));

    return taskName;
}

function createCSSTask(appPath) {
    const name = getNameFromPath(appPath);
    const taskName = `${name}-css`;
    const fullPath = path.join(appPath, 'static/css/*.css');
    const destPath = path.join(appPath, 'static/dist/css');

    gulp.task(taskName, gulp.series(() => {
        gulp.src(fullPath)
            .pipe(concatCSS('bundle.css'))
            .pipe(cleanCSS())
            .pipe(gulp.dest(destPath))
    }));

    gulp.watch(fullPath, gulp.series([taskName]));

    return taskName;
}

function createColorTask(appPath) {
    const name = getNameFromPath(appPath);
    const taskName = `${name}-colorcss`;
    const fullPath = path.join(appPath, 'static/css/color/*.css');
    const destPath = path.join(appPath, 'static/dist/css/color');

    gulp.task(taskName, gulp.series(() => {
        gulp.src(fullPath)
            .pipe(cleanCSS())
            .pipe(gulp.dest(destPath))
    }));

    gulp.watch(fullPath, gulp.series([taskName]));

    return taskName;
}

function createFontsTask(appPath) {
    const name = getNameFromPath(appPath);
    const taskName = `glyphicons-halflings-regular`;
    const fullPath = path.join(appPath, 'static/fonts/glyphicons-halflings-regular.{eot,svg,ttf,woff,woff2}');

    gulp.task(taskName, () => {
        gulp.src(fullPath)
            .pipe(gulp.dest('static/_shared/fonts/'));
    });
}

function createSharedTasks(destinations) {
    const cssTask = createSharedCSSTask(destinations);
    const jsTask = createSharedJSTask(destinations);
    const htmlTask = createSharedHTMLTask(destinations);
    const fontsTask = createSharedFontsTasks(destinations);

    return [cssTask, jsTask, htmlTask, fontsTask];
}

function createSharedCSSTask(destinations) {
    const taskName = '_shared.CSS';
    const fullPath = 'app/_shared/css/*css';

    gulp.task(taskName, gulp.series(() => {
        let pipeline = gulp.src(fullPath)
            .pipe(cleanCSS());

        queueDestinations(pipeline, 'CSS', destinations);
    }));

    gulp.watch(fullPath, gulp.series([taskName]));

    return taskName;
}

function createSharedJSTask(destinations) {
    const taskName = '_shared.JS';
    const fullPath = 'app/_shared/js/*.js';

    gulp.task(taskName, gulp.series(() => {
        let pipeline = gulp.src(fullPath);
        // Is Rollup needed?

        queueDestinations(pipeline, 'JS', destinations);
    }));

    gulp.watch(fullPath, gulp.series([taskName]));

    return taskName;
}

function createSharedHTMLTask(destinations) {
    const taskName = '_shared.HTML';
    const fullPath = 'app/_shared/*.html';

    gulp.task(taskName, gulp.series(() => {
        let pipeline = gulp.src(fullPath);

        queueDestinations(pipeline, 'HTML', destinations);
    }));

    gulp.watch(fullPath, gulp.series([taskName]));

    return taskName;
}

function createSharedFontsTasks(destinations) {
    const taskName = '_shared.FONTS';
    const fullPath = 'app/_shared/fonts/*';

    gulp.task(taskName, gulp.series(() => {
        let pipeline = gulp.src(fullPath);

        queueDestinations(pipeline, 'FONTS', destinations);
    }));

    gulp.watch(fullPath, gulp.series([taskName]));

    return taskName
}

function queueDestinations(pipeline, sectionName, destinations) {
    const sections = {
        'CSS': 'static/_shared/css',
        'JS': 'static/_shared/js',
        'HTML': 'views/_shared',
        'FONTS': 'static/_shared/fonts',
    };

    const currSection = sections[sectionName];

    for (var i = 0; i < destinations.length; i++) {
        const d = destinations[i];

        if (d !== './app/_shared' && d !== './app/gate') {
            const destFolder = path.join(d, currSection);
            pipeline = pipeline.pipe(gulp.dest(destFolder));
        }
    }

    return pipeline;
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
            uglify({}, ugh)
        ]
    };
}

function getTasks() {
    var rollupTasks = [];
    const appFolders = ['./app/', './api/'];

    for (let i = 0; i < appFolders.length; i++) {
        const currFolder = appFolders[i];
        const children = glob.sync(currFolder + '*');

        if (currFolder === './app/') {
            children.push('./api/secure'); // it's not a app, but it has a UI.
            const sharedTasks = createSharedTasks(children);
            rollupTasks = rollupTasks.concat(sharedTasks);
        }

        children.forEach((filePath) => {
            const staticPath = path.join(filePath, 'static');
            var appTasks = [];

            if (fs.existsSync(staticPath)) {
                appTasks = getEntryPoints(filePath);

                let fontTask = getFontsPoints(filePath);
                let cssTask = createCSSTask(filePath);
                let colorTask = createColorTask(filePath);

                appTasks.push(fontTask);
                appTasks.push(cssTask);
                appTasks.push(colorTask);
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

function getNameFromPath(appPath) {
    return appPath.replace('./', '').replace('/', '.');
}

gulp.task('default', gulp.parallel(getTasks()));