module.exports = [
  {breaking: true, release: 'major'},
  {revert: true, release: 'patch'},
  // Default angular
  {type: 'feat', release: 'minor'},
  {type: 'fix', release: 'patch'},
  {type: 'improvement', release: 'patch'},
  {type: 'docs', release: 'patch'},
  {type: 'refactor', release: 'patch'},
  {type: 'ci', release: 'patch'}
];
