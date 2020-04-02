const bunyan = require('bunyan'),
  JsonDB = require('node-json-db').JsonDB,
  Config = require('node-json-db/dist/lib/JsonDBConfig').Config;

const PORT = process.env.PORT || 5000,
  bunyanLogger = bunyan.createLogger({ name: 'Jarvis' }),
  Routes = require('./routes').Routes;


const db = new JsonDB(new Config('database', true, true, '/')),
  server = new Routes(PORT, bunyanLogger);

db.push('/init', configurationOnFirstStart());
db.push('/store', )

// Run service.
server.listen();

function configurationOnFirstStart() {
  const date = new Date();

  return {
    timeFirstStart: {
      date: date.getDate(),
      month: date.getMonth(),
      year: date.getFullYear()
    }
  };
}