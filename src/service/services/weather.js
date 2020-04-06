const https = require('https');
const cherrio = require('cheerio');

class Weather {
  constructor(city, state, country) {
    this._base = 'https://www.msn.com/en-in/weather/today/';
    this.postfix = '/we-city?weadegreetype=C';
    this.city = city;
    this.state = state;
    this.country = country;
  }

  formatInputs() {
    return this.city + ',' + this.state + ',' + this.country;
  }

  set() {
    return {
      min: 0,
      max: 0,
      day: 0,
      date: 0,
    };
  }

  skeleton() {
    return {
      temperature: 0,
      feelsLike: 0,
      humidity: 0,
      pressure: 0,
      dewPoint: 0,
      condition: '',
      set: [],
    };
  }

  scrape(stream) {
    const $ = cherrio.load(stream);
    const response = this.skeleton();
    response.temperature = $('span[class=current]').text();
    console.warn('tmpr: ', response.temperature);
    return response;
  }

  fetch() {
    return new Promise((resolve, reject) => {
      https
        .get(this._base + this.formatInputs() + this.postfix, (response) => {
          let chunks = '';

          response.on('data', (chunk) => {
            console.log('receiving packets...');
            chunks += chunk; // sequence of byte streams being added each time.
          });

          response.on('end', () => {
            resolve(this.scrape(chunks));
          });
        })
        .on('error', (err) => {
          reject(err);
        });
    });
  }
}

// module.exports = { Weather };
const obj = new Weather('bhubaneswar', 'orissa', 'india');
obj.fetch().then((result) => {
  console.warn('the result is ');
  console.warn(result);
});
