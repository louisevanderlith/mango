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
const flatten = require('gulp-flatten');

gulp.task("default", gulp.parallel(...watchApplications()));

//gather applications
//create watchers for given applications

//create _shared tasks
// watcher for .html
// watcher for .js
// watcher for .css
// write to applications found previously

function watchApplications() {
    var result = [];

    const appFolder = "./app/";
    const secureFolder = "./api/secure";

    const children = glob.sync(appFolder + '*');
    children.push(secureFolder); // it's not a app, but it has a UI.

    children.forEach((filePath) => {
        const staticPath = path.join(filePath, 'static');

        if (fs.existsSync(staticPath)) {
            var appjs = getEntries("js", "entry.js", staticPath);
            result.push(appjs);
        }
    });

    // _shared folder for every app
    const sharedTasks = createSharedTasks(children);
    result = result.concat(sharedTasks);

    return result;
}

function getEntries(folder, taskType, staticPath) {
    let results = [];

    let pattrn = `${folder}/*.${taskType}`;
    let pattrnPath = path.join(staticPath, pattrn);
    let filenames = getFileNames(pattrnPath);

    filenames.forEach((name) => {
        const entry = path.join(staticPath, `${folder}/${name}`);
        let taskName = getNameFromPath(staticPath, name);

        switch (folder) {
            case 'js':
            createJSTask(name, entry, staticPath, taskName);
            break;
            case 'css':
            createCSSTask(staticPath, pattrnPath, taskName);
            break;
            default:
            console.error("I don't know %s", folder);
        }

        gulp.watch(entry, gulp.series(taskName));
        results.push(taskName);
    });

    return results;
}

function getFileNames(pattern) {
    return glob.sync(pattern)
        .map((componentDir) => {
            return path.basename(componentDir);
        });
}

function createJSTask(name, entry, staticPath, taskName) {
    const dest = path.join(staticPath, "dist/js");
    let rollupOpts = getRollupOptions(entry, name);

    gulp.task(taskName, () => {
        gulp.src(entry)
            .pipe(rollup(rollupOpts, 'iife'))
            .on('error', (err) => {
                console.error("Entry: %s, Task: %s. Details: %s", entry, taskName, err);
            })
            .pipe(flatten())
            .pipe(gulp.dest(dest));
    });
}

function createCSSTask(staticPath, pattrnPath, taskName) {
    const dest = path.join(staticPath, "dist/css");

    gulp.task(taskName, () => {
        gulp.src(pattrnPath)
            .pipe(concatCSS('bundle.css'))
            .pipe(cleanCSS())
            .pipe(gulp.dest(dest))
    });
}

function getNameFromPath(staticPath, file) {
    const parts = staticPath.split('\\') || staticPath.split('/');

    let folder = parts[0];
    let app = parts[1];

    return `${folder}.${app}-${file}`;
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