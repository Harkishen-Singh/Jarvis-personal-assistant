const JsonDB = require('node-json-db').JsonDB;
const Config = require('node-json-db/dist/lib/JsonDBConfig').Config;
const _lock = require('lock').Lock;

class DBManager {
  constructor(name, isReadable) {
    this.db = new JsonDB(new Config(name, true, isReadable, '/'));
    this.stack = [];
  }

  getTransactionHistory() {
    return this.stack;
  }

  getMostRecentTransaction() {
    return this.stack[this.stack.length - 1];
  }

  clearHistory() {
    this.stack = [];
  }

  format(operation, path, data) {
    return {
      path,
      operation,
      data,
    };
  }

  fetch(path) {
    this.stack.push(this.format('fetch', path, null));
    return this.db.getData(path);
  }

  commit(path, data) {
    const lock = _lock();

    lock('rw', (release) => {
      this.stack.push(this.format('commit', path, data));
      this.db.push(path, data, true);
      release((err) => {
        if (err) throw err;
        console.log('released');
      });
    });
  }

  delete(path) {
    const lock = _lock();

    lock('rw', (release) => {
      this.stack.push(this.format('delete', path, data));
      this.db.delete(path);
      release();
    });
  }

  getMemorySizeofDatabaseinBytes() {
    let bytes = 0;

    function sizeOf(obj) {
      if (obj !== null && obj !== undefined) {
        switch (typeof obj) {
          case 'number':
            bytes += 8;
            break;
          case 'string':
            bytes += obj.length * 2;
            break;
          case 'boolean':
            bytes += 4;
            break;
          case 'object':
            const objClass = Object.prototype.toString.call(obj).slice(8, -1);
            if (objClass === 'Object' || objClass === 'Array') {
              for (const key in obj) {
                if (!obj.hasOwnProperty(key)) continue;
                sizeOf(obj[key]);
              }
            } else bytes += obj.toString().length * 2;
            break;
        }
      }
      return bytes;
    }

    function formatByteSize(bytes) {
      if (bytes < 1024) return bytes + ' bytes';
      else if (bytes < 1048576) return (bytes / 1024).toFixed(3) + ' KiB';
      else if (bytes < 1073741824) return (bytes / 1048576).toFixed(3) + ' MiB';
      else return (bytes / 1073741824).toFixed(3) + ' GiB';
    }

    return formatByteSize(sizeOf(obj));
  }
}

const DBService = new DBManager('db', true);

module.exports = { DBService };
