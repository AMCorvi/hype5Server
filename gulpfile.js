const gulp = require("gulp");
const babel = require("gulp-babel");
const eslint = require("gulp-eslint");
const concat = require("gulp-concat");
const sourcemaps = require("gulp-sourcemaps");
const gutil = require("gulp-util");
const exec = require("child_process").exec

const paths = {
    dist: "./dist",
    src: "./src/**/*.js",
    server: "./src/server.js"
}

gulp.task("test", cb => {
    exec("./node_modules/.bin/jest -o ", (err,stdout, lo)=>{
        if (err) return cb(err);
        gutil.log("\n\n", stdout,stdin );
        cb();
    });
});

gulp.task("transpile", () => {
    return gulp.src([paths.src])
        .pipe(sourcemaps.init())
		.pipe(babel({
			presets: ["env", "flow"]
		}))
		.pipe(sourcemaps.write("."))
		.pipe(gulp.dest(paths.dist));
});

gulp.task("flow", cb => {
    exec("./node_modules/.bin/flow", (err, stdout, stdin) => {
        if (err) return cb(err);
        gutil.log(gutil.colors.green("\n\n", stdout, stdin));
        cb();
    });
})

gulp.task("watch", ["flow","transpile"], () => {
    gulp.watch(paths.src, ["flow","transpile"]);
})

gulp.task("default",["watch"])
