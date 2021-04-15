const fetch = require('node-fetch');
const PORT = 8080;

var http = require('http');
var fs = require('fs');

// init http server
fs.readFile('./index.html', function (err, html) {

    if (err) throw err;

    http.createServer(function (request, response) {
        response.writeHeader(200, { "Content-Type": "text/html" });
        response.write(html);
        response.end();
    }).listen(PORT);
});