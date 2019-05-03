/**
 * COSS society
 * logs google search engine code for the caller parent subprocess
 * author: codeZero
 */
const webdriver = require('selenium-webdriver'),
    chrome = require('selenium-webdriver/chrome'),
    By = webdriver.By;

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

process.argv.forEach((val, index, array) => {
    var path = require('chromedriver').path;

    var service = new chrome.ServiceBuilder(path).build();
    chrome.setDefaultService(service);
    if (index === 2) { // corresponds to the medicine name
        // var medicines = val;
        var sympLink = val
        console.log('symptom received -> ' + sympLink);
        // keep the below block of code in the last part the else if block

        var options = new chrome.Options(), data
        options.addArguments("--no-sandbox");
        options.addArguments("--disable-dev-shm-usage");
        options.addArguments("--disable-gpu");
        options.addArguments("--headless");
        var driver = new webdriver.Builder()
                            .setChromeOptions(options)
                            .forBrowser('chrome')
                            .build();

        const BASEURL = ' https://www.medindia.net/drugs/medical-condition/';
        driver.get(BASEURL + sympLink).then(() => {
            console.log('searching')
            driver.findElements(By.tagName('article')).then(loop => {
                loop.forEach((itr, ind) => {
                    itr.getAttribute('innerText').then(attr => {
                        attr = attr.replace('•', '');
                        data += attr + "\n";
                        if (ind === (loop.length - 1)){
                            console.log('data -> ' + data)
                        }
                    })
                })
            });
        });
    }
});
