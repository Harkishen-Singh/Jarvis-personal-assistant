/* eslint-disable max-len */
/* eslint-disable new-cap */
const restify = require('restify');
const Default = require('./handlers/default').Handler;
const Message = require('./handlers/messages').Message;
const Handlers = {
  default: new Default(),
  messages: new Message()
};
const corsMiddleware = require('restify-cors-middleware');
const cors = corsMiddleware({
  origins: ['*']
});

class WebManager {
  constructor(port, logger) {
    this.server = restify.createServer({
      name: 'myapp',
      version: '0.0.1',
      log: logger
    });
    this.PORT = port;
    this.log = logger;
    this.server.pre(cors.preflight);
    this.server.use(cors.actual);
    this.applyMiddleWares();
    this.applyRoutes();
  }

  applyMiddleWares() {
    this.server.use(restify.plugins.acceptParser(this.server.acceptable));
    this.server.use(restify.plugins.queryParser());
    this.server.use(restify.plugins.bodyParser());
    // this.server.use(
    //     function crossOrigin(req, res, next) {
    //       res.header('Access-Control-Allow-Origin', '*');
    //       res.header('Access-Control-Allow-Headers', 'X-Requested-With');
    //       return next();
    //     }
    // );
  }

  applyRoutes() {
    this.server.get('/', Handlers.default.default);
    this.server.get('/echo/:name', Handlers.default.echo);
    this.server.get('/message', Handlers.messages.recMessage );
  }

  listen() {
    this.server.listen(this.PORT, () => {
      console.log('%s listening at %s', this.server.name, this.server.url);
    });
  }
}

module.exports = { WebManager };
