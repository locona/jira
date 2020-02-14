const micromatch = require('micromatch');

console.log(micromatch(['release/M11', 'release/M2'], ['release/M*'])) //=> ['foo', 'bar', 'baz']
