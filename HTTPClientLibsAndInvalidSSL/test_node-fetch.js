const https = require('https');

async function fetchData(url) {
  const fetch = (await import('node-fetch')).default;

  try {
    const response = await fetch(url);
    const data = await response.text();
    console.log(data.substring(0, 500));
  } catch (error) {
    console.error('Fetch error:', error);
  }
}

const testUrl = 'https://expired.badssl.com/';

fetchData(testUrl);
