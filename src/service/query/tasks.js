const Weather = require('../services/weather').Weather;

class Task {
  constructor(name) {
    this.task = name;
  }

  log() {
    console.log(`running service with task: ${this.task} ...`);
  }

  run(task, object) {
    console.warn(`task-runner: args: ${JSON.stringify(object)}`);
    switch (task) {
      case 'weather':
        if (Object.keys(object).length < 3) {
          throw new Error(`task: invalid number fo args: ${args.toString()}`);
        }
        const { city, state, country } = object;
        return this.weather(city, state, country);
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
