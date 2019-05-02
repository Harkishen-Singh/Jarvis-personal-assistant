const JsonDB = require('node-json-db'),
	Config = require('node-json-db/lib/JsonDBConfig');

class DBase {

	constructor(dname) {

		this.dname = dname;
		this.db = new JsonDB(new Config(dname, true, true, '/'));

	}

	updateStore(path, data) {

		if (path[ 0 ] === '/') {

			this.db.push(path, data);
			return true;

		}
		throw new Error('invalid path. Path not starting with "/".');

	}

	getData(path) {

		if (path[ 0 ] === '/') {

			this.db.getData(path);

		}
		throw new Error('invalid path. Path not starting with "/".');

	}

	deleteData(path) {

		if (path[ 0 ] === '/') {

			this.db.delete(path);

		}
		throw new Error('invalid path. Path not starting with "/".');

	}

	reload() {

		this.db.reload();

	}

}

module.exports = DBase;
