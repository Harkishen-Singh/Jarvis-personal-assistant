const AST = require('./ast').AST;
const cheerio = require('cheerio');
class Message {
  async recMessage(req, response, next) {
    const username = req.query.username;
    const message = req.query.message;

    if (message.length >= 2) {
      const messageArr = message.split(' ');
      const searchKind = messageArr[0];
      const queryParam = messageArr[1];
      const ASTT = new AST();
      console.log('HEre');
      let res = await ASTT.serialize(searchKind, queryParam);
      // console.log('resp:: ', res);
      console.log('I am hereeeeeee');
      const results = [];
      res = res.data;
      if (res) {
        const $ = cheerio.load(res);
        //   console.log('$ loaded:: ', $);
        $('.BNeawe.vvjwJb.AP7Wnd').each(function() {
          const details = new Object();
          details.title = $(this).text();
          // console.log('details:: ', details);
          results.push(details);
        });
        response.send(results);
      }
    }
    next();
  }
}

module.exports = { Message };
