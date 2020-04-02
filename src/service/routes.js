const restify = require('restify');

class Routes {
  constructor(port, logger) {
    this.server = restify.createServer({
      name: 'myapp',
      version: '0.0.1',
      log: logger
    });
    this.PORT = port;
    this.log = logger;

    this.applyMiddleWares();
    this.applyRoutes();
  }

  applyMiddleWares() {
    this.server.use(restify.plugins.acceptParser(this.server.acceptable));
    this.server.use(restify.plugins.queryParser());
    this.server.use(restify.plugins.bodyParser());
  }

  applyRoutes() {
    this.server.get('/', handlers.default);
    this.server.get('/echo/:name', handlers.echo);
  }

  listen() {
    this.server.listen(this.PORT, () => {
      console.log('%s listening at %s', this.server.name, this.server.url);
    });
  }
}

const handlers = {
  default: (req, res, next) => {
    res.send('hello user');
    return next();
  },
  echo: (req, res, next) => {
    res.send(req.params);
    return next();
  }
}

module.exports = { Routes };