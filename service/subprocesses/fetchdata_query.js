/**
 * COSS society
 * logs google search engine code for the caller parent subprocess
 * author: codeZero
 */
const webdriver = require('selenium-webdriver'),
    chrome = require('selenium-webdriver/chrome'),
    By = webdriver.By;

require('chromedriver');
require('geckodriver');

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
let method = null,
    url = null;
process.argv.forEach((val, index, array) => {
    var path = require('chromedriver').path;

    var service = new chrome.ServiceBuilder(path).build();
    chrome.setDefaultService(service);
    if (index === 2) { // corresponds to the search method
        method = val;
    } else if(index == 3) { // corresponds to search url
        url = val;

        // keep the below block of code in the last part the else if block

        var options = new chrome.Options();
        options.addArguments("--no-sandbox");
        options.addArguments("--disable-dev-shm-usage");
        options.addArguments("--disable-gpu");
        options.addArguments("--headless");

        var driver = new webdriver.Builder()
                            .setChromeOptions(options)
                            .forBrowser('chrome')
                            .build();

        driver.get(url).then(() => {
            driver.findElement(By.tagName('body')).then(cc => {
                cc.getAttribute('innerHTML').then(result => {
                    console.log("result::: ", result);
                    driver.quit();
                });
            });
        });
    }
})

