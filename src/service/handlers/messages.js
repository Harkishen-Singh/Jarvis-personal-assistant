const AST = require('./ast').AST;
class Message {

  recMessage(req, res, next) {
    var username = req.query.username;
    var message = req.query.message;

    if (message.length >= 2) {
       var messageArr = message.split(" ")
      var searchKind = messageArr[0]
      var queryParam = messageArr[1]
      const ASTT = new AST();
      console.log("HEre")
      var resp = ASTT.serialize(searchKind, queryParam)
      console.log("resp:: ", resp)
    } 
    res.send(username);
    next();
  }

}

module.exports = { Message };
