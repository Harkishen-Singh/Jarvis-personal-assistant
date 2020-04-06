const bunyan = require('bunyan');

const PORT = process.env.PORT || 5000;
const bunyanLogger = bunyan.createLogger({ name: 'Jarvis' });
const initData = require('./utils/configurations').data;
const WebManager = require('./web').WebManager;
const db = require('./utils/db-manager').DBService;

// Initialize database.
db.commit('/init', initData.onFirstStart);
db.commit('/store', initData.store);
db.commit('/personal/location', initData.personal.location);

const server = new WebManager(PORT, bunyanLogger);

// Run service.
server.listen();
