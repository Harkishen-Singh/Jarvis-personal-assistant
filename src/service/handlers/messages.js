const AST = require('./ast').AST;
class Message {
  recMessage(req, res, next) {
    const username = req.query.username;
    const message = req.query.message;

    if (message.length >= 2) {
      const messageArr = message.split(' ');
      const searchKind = messageArr[0];
      const queryParam = messageArr[1];
      const ASTT = new AST();
      console.log('HEre');
      const resp = ASTT.serialize(searchKind, queryParam);
      console.log('resp:: ', resp);
    }
    res.send(username);
    next();
  }
}

module.exports = { Message };
