const axios = require('axios');
const cheerio = require('cheerio');
const request = require('request')
const SCRAPING_URL = 'https://www.google.com/search?q=delhi';

class QueryController {

    HandleGoogleQuery(query) {
        axios.get(SCRAPING_URL)
            .then((res) => {
                console.log("I am hereeeeeee");
                const results = [];
                res = res.data;
                if (res) {
                    const $ = cheerio.load(res);
                    console.log("$ loaded:: ", $)
                    $('.BNeawe.vvjwJb.AP7Wnd').each(function () {
                        
                        var details = new Object();
                        details.title = $(this).text();
                        console.log("details:: ", details)
                        results.push(details)
                    });
                }
                return results
            })
            .catch(err => console.log(err));
    }

}

module.exports = { QueryController }