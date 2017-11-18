const express = require('express');
const path = require('path');
const open = require('open');

/* eslint-disable no-console */
const port = 4000;
const app = express();
// const compiler = webpack(config);

app.use(express.static('dist'));

// app.use(require('webpack-hot-middleware')(compiler));

app.get('*', (req, res) => {
  res.sendFile(path.join(__dirname, '../dist/index.html'));
});

app.listen(port, (err) => {
  if (err) {
    console.log(err);
  } else {
    open(`http://localhost:${port}`);
  }
});

