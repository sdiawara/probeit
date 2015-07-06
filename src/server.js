/*global require, define, console*/

function start() {
    'use strict';

    var express = require('express'),
        app = express(),
        server;

    app.get('/', function (req, res) {
        res.send('Hello probeit!');
    });

    server = app.listen(3000, function () {
        var host = server.address().address,
            port = server.address().port;
        
        console.log('Pollit listening at http://%s:%s', host, port);
    });
}

start();