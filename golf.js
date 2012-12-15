#!/usr/bin/env node
print =console.log;
assert = require('assert');
fs = require('fs');

fh = fs.readFileSync('./output.txt', 'utf-8');
ans = function(){
  return (eval('function test(){' + fs.readFileSync('./code.js', 'utf-8') + '}; test();'));
}();
print(ans);
assert.equal(ans, fh);
print(ans);
