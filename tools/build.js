/* eslint-disable no-console */
const webpack = require('webpack');
require('colors');
const webpackConfig = require('../webpack.config.prod');

process.env.NODE_ENV = 'production';

webpack(webpackConfig).run((err, stats) => {
  if (err) {
    console.log(err.bold.red);
    return 1;
  }

  const jsonStats = stats.toJson();

  if (jsonStats.hasErrors) {
    return jsonStats.errors.map(error => console.log(error.red));
  }

  if (jsonStats.hasWarnings) {
    console.logs('Webpack generated some warnings. '.bold.yellow);
    jsonStats.warnings.map(warning => console.log(warning.yellow));
  }

  console.log(`webpack stats: ${stats}`);
  console.log('compilation success'.green);
  return 0;
});
