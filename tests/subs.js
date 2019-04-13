/**
 * CMOS society
 * logs google search engine code for the caller parent subprocess
 * author: codeZero
 */
const webdriver = require('selenium-webdriver'),
    chrome = require('selenium-webdriver/chrome'),
    By = webdriver.By;
const util = require('util')

require('chromedriver');
// require('geckodriver');

/**
 * reference link: https://stackoverflow.com/a/4351548
 * example below
 * $ node process-2.js one two=three four
    0: node
    1: /Users/mjr/work/node/process-2.js
    2: one
    3: two=three
    4: four
 */
let email = null,
    password = null,
    url = null;
process.argv.forEach((val, index, array) => {
    var path = require('chromedriver').path;

    var service = new chrome.ServiceBuilder(path).build();
    chrome.setDefaultService(service);
    if (index === 2) { // corresponds to the emailID heroku
        email = val;
    } else if (index == 3) { // corresponds to password heroku
        password = val;
    } else if(index == 3) { // corresponds to search url
        url = val;

        // keep the below block of code in the last part the else if block

        var options = new chrome.Options();
        options.addArguments("--no-sandbox");
        options.addArguments("--disable-dev-shm-usage");
        options.addArguments("--disable-gpu");
        // options.addArguments("--headless");
        var arrAnswer = [];

        var driver = new webdriver.Builder()
                            .setChromeOptions(options)
                            .forBrowser('chrome')
                            .build();

        driver.get(url).then(() => {
            
        });
    }
})

