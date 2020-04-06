const Weather = require('../services/weather').Weather;

class Task {
  constructor(name) {
    this.task = name;
  }

  log() {
    console.log(`running service with task: ${this.task} ...`);
  }

  run(task, ...args) {
    console.warn(`task-runner: args: ${args.toString()}`);
    switch (task) {
      case 'weather':
        if (args.length < 3) {
          throw new Error(`task: invalid number fo args: ${args.toString()}`);
        }

        return this.weather(args[0], args[1], args[2]);
        break;

      case 'meaning':
        this.meaning(args[0]);
    }
  }

  weather(city, state, country) {
    const weather = new Weather(city, state, country);
    return weather.fetch();
  }

  meaning(entity) {
    // Do meaning operation
  }
}

module.exports = { Task };
