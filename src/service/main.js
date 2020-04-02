const bunyan = require('bunyan');
const JsonDB = require('node-json-db').JsonDB;
const Config = require('node-json-db/dist/lib/JsonDBConfig').Config;

const PORT = process.env.PORT || 5000;
const bunyanLogger = bunyan.createLogger({name: 'Jarvis'});
const Configurations = require('./utils/configurations').configurations;
Routes = require('./routes').Routes;

const db = new JsonDB(new Config('db', true, true, '/'));
const server = new Routes(PORT, bunyanLogger);

// Initialize database.
db.push('/init', Configurations.onFirstStart);
db.push('/store', Configurations.store);

// Run service.
server.listen();
