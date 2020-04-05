class Handler {
  default(req, res, next) {
    res.send('hello user');
    return next();
  }

  echo(req, res, next) {
    res.send(req.params);
    return next();
  }
}

module.exports = { Handler };
