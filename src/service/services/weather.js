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
      date: 0
    };
  }

  skeleton() {
    return {
      temperature: 0,
      feelsLike: 0,
      humidity: 0,
      pressure: 0,
      dewPoint: 0,
      visibility: '',
      wind: '',
      condition: '',
      set: []
    };
  }

  scrape(stream) {
    const $ = cherrio.load(stream);
    const response = this.skeleton();
    response.temperature = $('span[class=current]').text();

    const filters = {
      filter: (state, txt, re=null) => {
        if (re !== null) return state.replace(txt, re);
        return state.replace(txt, '');
      },
      feelsLike: (txt) => {
        txt = txt.replace('Feels Like ', '');
        const garbagePosition = txt.indexOf('&');
        return txt.substring(0, garbagePosition);
      },
      dewPoint: (txt) => {
        txt = filters.filter(txt, 'Dew Point ');
        const position = txt.indexOf('&');
        return txt.substring(0, position);
      },
      skipFirstSpace: (txt) => {
        const position = txt.indexOf(' ');
        return txt.substring(position + 1);
      }
    };

    const details = $('.weather-info').html();
    const arr = details.split('</li>');
    arr.forEach((ele, i) => {
      ele = filters.filter(ele, '<span>');
      ele = filters.filter(ele, '</span>');
      ele = filters.filter(ele, '<span>');
      ele = filters.filter(ele, '</span>');
      ele = filters.filter(ele, '\n<ul>\n\n', '.');
      ele = filters.filter(ele, '<li>');
      arr[i] = filters.filter(ele, '\n');
    });

    const firstSet = arr[0].split('.');
    response.condition = firstSet[0];
    response.feelsLike = filters.feelsLike(firstSet[1]);

    response.wind = filters.skipFirstSpace(arr[1]);
    response.pressure = filters.skipFirstSpace(arr[2]);
    response.visibility = filters.skipFirstSpace(arr[3]);
    response.humidity = filters.skipFirstSpace(arr[4]);
    response.dewPoint = filters.dewPoint(arr[5]);

    return response;
  }

  fetch() {
    return new Promise((resolve, reject) => {
      https.get(this._base + this.formatInputs() + this.postfix, (response) => {
        let chunks = '';

        response.on('data', (chunk) => {
          console.log('receiving packets...');
          chunks += chunk; // sequence of byte streams being added each time.
        });

        response.on('end', () => {
          resolve(this.scrape(chunks));
        });
      }).on('error', (err) => {
        reject(err);
      });
    });
  }
}

module.exports = { Weather };
// const obj = new Weather('bhubaneswar', 'orissa', 'india');
// obj.fetch().then((result) => {
//   console.warn('the result is ');
//   console.warn(result);
// });
