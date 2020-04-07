/* eslint-disable no-invalid-this */
/* eslint-disable no-new-object */
const axios = require('axios');
const cheerio = require('cheerio');

class Message {
  async recMessage(req, res, next) {
    // const username = req.query.username;
    res.setHeader('Access-Control-Allow-Origin', '*');
    const message = req.query.message;
    const results = [];

    if (message.length >= 2) {
      const messageArr = message.split(' ');
      const searchKind = messageArr[0];
      const queryParam = messageArr[1];
      let SCRAPING_URL;
      switch (searchKind) {
        case 'google':
          SCRAPING_URL = 'https://www.google.com/search?q=' + queryParam;
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
                const response = {
                  Status: true,
                  Message: 'here are the top search results',
                  Result: results
                };
                res.send(response);
              })
              .catch((err) => console.log(err));
          break;
        case 'meaning':
          SCRAPING_URL =
            'https://www.oxfordlearnersdictionaries.com/definition/english/' +
            queryParam;
          axios
              .get(SCRAPING_URL)
              .then((resd) => {
                resd = resd.data;
                if (resd) {
                  const $ = cheerio.load(resd);
                  $('ol > li').each(function() {
                    const currentNode = $(this);
                    const defNode = currentNode.find('span.def');
                    let def;
                    let example;

                    if (defNode) {
                      def = defNode.text();
                    }

                    const exampleNode = currentNode
                        .find('ul.examples > li')
                        .first();
                    if (exampleNode) {
                      example = exampleNode.text();
                    }

                    if (def) {
                      const details = new Object();
                      details.Meaning = def;
                      details.Example = example;
                      results.push(details);
                    }
                  });
                }
                const response = {
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
