/*global module*/
module.exports = function (grunt) {
    'use strict';

    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),
        uglify: {
            options: {
                banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n'
            },
            build: {
                src: 'src/*.js',
                dest: 'build/<%= pkg.name %>.min.js'
            }
        },
        coveralls: {
            options: {
                force: false
            },
            coveralls: {
                src: 'build/reports/coverage/*.info'
            }
        },
        jasmine_node: {
            test: {
                options: {
                    coverage: {
                        reportDir: './build/reports/coverage/'
                    },
                    forceExit: true,
                    match: '.',
                    matchAll: false,
                    specFolders: ['test'],
                    extensions: 'js',
                    specNameMatcher: 'Spec',
                    captureExceptions: true,
                    junitreport: {
                        report: false,
                        savePath : './build/reports/jasmine/',
                        useDotNotation: true,
                        consolidate: true
                    }
                },
                src: ['**/*.js']
            }
        }
    });

    grunt.loadNpmTasks('grunt-contrib-uglify');
    grunt.loadNpmTasks('grunt-coveralls');
    grunt.loadNpmTasks('grunt-jasmine-node-coverage');


    grunt.registerTask('default', ['jasmine_node', 'uglify']);

};