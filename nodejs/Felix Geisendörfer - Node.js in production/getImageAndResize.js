var http = require('http');
var spawn = require('child_process').spawn;

http.createServer(function(req, res){
	var path = __dirname + '/img/' + req.url;
	var args = [path, '-resize', '1000x1000', '-'];
	var convert = spawn('convert', args);

	convert.stdout.pipe(res);
	convert.stderr.pipe(process.stderr);
}).listen(3000);