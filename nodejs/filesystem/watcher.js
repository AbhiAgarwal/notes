'use strict'

var fs = require('fs');

fs.watch('testfile.txt', function(){
	console.log('File has changed');
});

console.log('Now watching');