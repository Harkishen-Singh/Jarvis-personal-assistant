/**
 * CMOS society
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


// process.argv.forEach((val, index, array) => {
    var path = require('chromedriver').path;

    var service = new chrome.ServiceBuilder(path).build();
    chrome.setDefaultService(service);
    // if (index === 2) { // corresponds to the medicine name
        // var medicines = val;
        var medicines = 'Rebamipide'

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

        const BASEURL = ' https://www.medindia.net/drugs/medical-condition/',
            ENDWARE = '.htm';
        driver.get(BASEURL + medicines + ENDWARE).then(() => {
            driver.findElements(By.className('drug-content')).then(loop => {
                var count = -1;
                var data='';
                loop.forEach((itr, i) => {
                    if (count == -1) {
                        count++;
                        return;
                    } else {
                        itr.getAttribute('innerText').then(text => {
                        text = text.replace('•', '');
                        console.warn(count++);
                        data += text;
                        if (i === loop.length - 1){
                            console.warn('data -> ' + data)
                        }
                        });
                    }
                });
            });
        });
    // }
// });

