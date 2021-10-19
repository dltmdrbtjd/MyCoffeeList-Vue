const express = require('express');
const app = express();

app.get('/', function (req, res) {
  res.send('Server set');
});

app.listen(5000, () => {
  console.log('server on');
});
