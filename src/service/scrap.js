const axios = require('axios');
const cheerio = require('cheerio');
const request = require('request');
const SCRAPING_URL = 'https://www.google.com/search?q=delhi';

const results = [];

axios.get(SCRAPING_URL)
    .then((res) => {
      const results = [];
      res = res.data;
      if (res) {
        const $ = cheerio.load(res);

        $('.BNeawe.vvjwJb.AP7Wnd').each(function() {
          const details = new Object();
          details.title = $(this).text();
          results.push(details);
        });
      }
      console.log(results);
    })
    .catch((err) => console.log(err));


