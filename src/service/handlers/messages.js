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
                    details.head = $(this).text();
                    console.log('details:: ', details);
                    results.push(details);
                  });
                }
                const response = {
                  status: true,
                  message: 'here are the top search results',
                  result: results
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
                      details.meaning = def;
                      details.example = example;
                      results.push(details);
                    }
                  });
                }
                const response = {
                  status: true,
                  message: 'here are the top search results',
                  result: results
                };
                res.send(response);
              })
              .catch((err) => console.log(err));
          break;
        case 'yahoo':
          SCRAPING_URL =
            'https://search.yahoo.com/search?nojs=1&p=' + queryParam;
          axios
              .get(SCRAPING_URL)
              .then((resd) => {
                resd = resd.data;
                if (resd) {
                  const $ = cheerio.load(resd);
                  $('ol > li > div > div.compTitle').each(
                      function() {
                        const currentNode = $(this);
                        console.log(currentNode.text());
                        const titleNode = currentNode.children('h3');
                        let link;
                        let title;

                        if (titleNode) {
                          title = titleNode.text();
                        }

                        const linkNode = currentNode.children('div');
                        if (linkNode) {
                          link = linkNode.text();
                        }

                        if (link) {
                          const details = new Object();
                          details.head = title;
                          details.link = link;
                          results.push(details);
                        }
                      }
                  );
                }
                const response = {
                  status: true,
                  message: 'here are the top search results',
                  result: results
                };
                res.send(response);
              })
              .catch((err) => console.log(err));
          break;
      }
    }

    next();
  }
}

module.exports = { Message };
