class Task {
  constructor(name) {
    this.task = name;
  }

  log() {
    console.log(`running service with task: ${this.task} ...`);
  }

  run(task, ...args) {
    switch (task) {
      case 'weather':
        this.weather(args[0]);
        break;

      case 'meaning':
        this.meaning(args[0]);
    }
  }

  weather(location) {
    // Do weather operation
  }

  meaning(entity) {
    // Do meaning operation
  }
}

module.exports = { Task };
