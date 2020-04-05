const https = require('https');

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

  fetch() {
    return new Promise((resolve, reject) => {
      https.get(this.base + this.formatInputs() + this.postfix, (response) => {
        let chunks = '';

        response.on('data', (chunk) => {
          console.log('receiving packets...');
          chunks += chunk; // sequence of byte streams being added each time.
        });

        response.on('end', () => {
          resolve(chunks);
        });
      }).on('error', (err) => {
        reject(err);
      });
    });
  }
}

module.exports = { Weather };
