const fs = require('fs');
const cheerio = require('cheerio');
require('colors');

/* eslint-disable no-console */

fs.readFile('src/index.ejs', 'utf8', (err, markup) => {
  if (err) {
    return console.log(err);
  }

  const $ = cheerio.load(markup);

  $('head').prepend('<link rel="stylesheet" href="/style.css">');

  fs.writeFile('dist/index.html', $.html(), 'utf8', (errInner) => {
    if (errInner) {
      return console.log(errInner);
    }
    return console.log('index.html built'.green);
  });
  return console.log('index.html completed build'.green);
});
