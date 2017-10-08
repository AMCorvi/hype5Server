const gulp = require("gulp");
const babel = require("gulp-babel");
const sourcemaps = require("gulp-sourcemaps");
const gutil = require("gulp-util");
const shell = require("gulp-shell");

const paths = {
    dist: "./dist",
    distRegEx: ["<rootDir>/dist/"],
    src: "./src/**/*.js",
    server: "./src/server.js",
    jestconfig: "./jest.json",
    tests: "*.test.js"
}

gulp.task("test", shell.task([

    "./node_modules/.bin/jest" +
    " --no-cache"

]));

gulp.task("transpile", () => {
    return gulp.src([paths.src])
        .pipe(sourcemaps.init())
		.pipe(babel())
		.pipe(sourcemaps.write("."))
		.pipe(gulp.dest(paths.dist));
});

gulp.task("flow", shell.task(["./node_modules/.bin/flow"]) )

gulp.task("watch", ["flow","transpile"], () => {
    gulp.watch(paths.src, ["run"]);
})

/*Task sequence for watch.
Using a seperate task instead of listing the gulp. Watch call prevents async calls to functions
You learn something new everyday */
gulp.task("run", ["flow", "test", "transpile"]);

gulp.task("default",["watch"])
