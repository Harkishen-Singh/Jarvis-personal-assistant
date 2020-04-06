/* eslint-disable no-invalid-this */
/* eslint-disable no-new-object */
const axios = require('axios');
const cheerio = require('cheerio');

const results = [];
class Message {
  async recMessage(req, res, next) {
    // const username = req.query.username;
    const message = req.query.message;

    if (message.length >= 2) {
      const messageArr = message.split(' ');
      const searchKind = messageArr[0];
      const queryParam = messageArr[1];
      switch (searchKind) {
        case 'google':
          const SCRAPING_URL = 'https://www.google.com/search?q=' + queryParam;
          axios
              .get(SCRAPING_URL)
              .then((resd) => {
                resd = resd.data;
                if (resd) {
                  const $ = cheerio.load(resd);
                  $('.BNeawe.vvjwJb.AP7Wnd').each(function() {
                    const details = new Object();
                    details.Head = $(this).text();
                    console.log('details:: ', details);
                    results.push(details);
                  });
                }
                response = {
                  Status: true,
                  Message: 'here are the top search results',
                  Result: results
                };
                res.send(response);
              })
              .catch((err) => console.log(err));
      }
    }

    next();
  }
}

module.exports = { Message };
