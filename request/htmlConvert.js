process.env.GLOBAL_AGENT_HTTP_PROXY = 'http://127.0.0.1:7890';
require('global-agent/bootstrap');

const axios = require('axios');
const fs = require('fs');

axios.get('https://v2ex.com/t/1019827')
  .then(response => {
    fs.writeFileSync('/Users/sure/Documents/page.html', response.data);
    console.log('HTML saved to page.html');
  })
  .catch(error => {
    console.error('Error fetching page:', error.message);
  });