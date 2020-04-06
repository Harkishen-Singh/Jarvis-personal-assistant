const Query = require('../query/query').Query;

class Handler {
  async execute(req, res, next) {
    const query = req.query.query;
    console.warn('query is ', query);
    const runner = new Query(query);
    const response = await runner.run();
    res.send(response);
    return next();
  }
}

module.exports = { Handler };
