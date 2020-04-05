const restify = require('restify');
const Default = require('./handlers/default').Handler;
const Message = require('./handlers/messages').Message;
const Handlers = {
  default: new Default(),
  messages: new Message()
};

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
    this.server.get('/', Handlers.default.default);
    this.server.get('/echo/:name', Handlers.default.echo);
    this.server.get('/messages', Handlers.messages.recMessage )
  }

  listen() {
    this.server.listen(this.PORT, () => {
      console.log('%s listening at %s', this.server.name, this.server.url);
    });
  }
}

module.exports = {Routes};
