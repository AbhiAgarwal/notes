// Source: http://jsperf.com/string-vs-array-concat/2

// String w/o flattening
var dest = 'apple';
while (dest.length < 100) {
  var offs = dest.length - 5;
  for (var i = offs; i < offs + 10; i++) {
    dest += precompiled.charAt(i);
  }
}

// String
var dest = 'apple';
while (dest.length < 100) {
  var offs = dest.length - 5;
  for (var i = offs; i < offs + 10; i++) {
    dest += dest.charAt(i);
  }
}

// String w/o flattening by index (IE > 7)
var dest = 'apple';
while (dest.length < 100) {
  var offs = dest.length - 5;
  for (var i = offs; i < offs + 10; i++) {
    dest += precompiled[i];
  }
}

// Array
var dest = ['a', 'p', 'p', 'l', 'e'];
while (dest.length < 100) {
  var offs = dest.length - 5;
  for (var i = offs; i < offs + 10; i++) {
    dest[dest.length] = dest[i];
  }
}
dest = dest.join('');

// Prep code
var precompiled = 'apple';
while (precompiled.length < 100) {
  var offs = precompiled.length - 5;
  for (var i = offs; i < offs + 10; ++i) {
    // calling charAt(i) on `precompiled` will
    // flatten the string in most engines
    precompiled += precompiled.charAt(i);
  }
}
