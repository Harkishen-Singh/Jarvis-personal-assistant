const bunyan = require('bunyan');

const PORT = process.env.PORT || 5000;
const bunyanLogger = bunyan.createLogger({ name: 'Jarvis' });
const Configurations = require('./utils/configurations').configurations;
const WebManager = require('./routes').WebManager;
const db = require('./utils/db-manager').DBService;

// Initialize database.
db.commit('/init', Configurations.onFirstStart);
db.commit('/store', Configurations.store);

const server = new WebManager(PORT, bunyanLogger);

// Run service.
server.listen();
