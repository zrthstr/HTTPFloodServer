const axios = require('axios');
const https = require('https');

const url = 'https://expired.badssl.com/';

axios.get(url)
  .then(response => {
    console.log(`Status Code: ${response.status}`);
    console.log(`Response Body: ${response.data.substring(0, 500)}`);
  })
  .catch(error => {
    console.error(error);
  });

