/*global describe, it, require, expect, console,process*/

describe('Probebit', function () {
    'use strict';
    it('respond', function (done) {
        require(process.cwd() + '/src/server');
        require('request').get('http://localhost:3000', function (error, response, body) {
            expect(body).toBe('Hello probeit!');
            done();
        });
    });
});