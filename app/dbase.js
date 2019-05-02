const JsonDB = require('node-json-db'),
	Config = require('node-json-db/lib/JsonDBConfig');

class DBase {
	constructor(dname) {
		this.dname = dname;
		this.db = new JsonDB(new Config(dname, true, true, '/'));
	}

	updateStore(path, data) {

	}
}