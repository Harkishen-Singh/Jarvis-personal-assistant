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


process.argv.forEach((val, index, array) => {
    var path = require('chromedriver').path;

    var service = new chrome.ServiceBuilder(path).build();
    chrome.setDefaultService(service);
    if (index === 2) { // corresponds to the medicine name
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

        const BASEURL = 'https://www.medindia.net/doctors/drug_information/',
            ENDWARE = '.htm';
        driver.get(BASEURL + medicines + ENDWARE).then(() => {
            driver.findElements(By.className('drug-content')).then(loop => {
                var count = 0;
                loop.forEach(itr => {
                    var objFormat = {};
                    if (count == 0) {
                        count++;
                        return;
                    } else {
                        itr.getAttribute('innerText').then(text => {
                        console.warn(count++);
                        text = text.replace('•', '');
                        if (count == 1)
                            objFormat['trade_names'] = text;
                        else if (count == 2)
                            objFormat['prescribed'] = text;
                        else if (count == 3)
                            objFormat['contraindications'] = text;
                        else if (count == 4)
                            objFormat['dosage'] = text;
                        else if (count == 5)
                            objFormat['how_to_consume'] = text;
                        else if (count == 6)
                            objFormat['precautions'] = text;
                        else if (count == 7)
                            objFormat['side_effects'] = text;
                        else if (count == 8)
                            objFormat['relevant_info'] = text;
                        else if (count == 9) {

                            objFormat['storage_conditions'] = text;
                            let inString = JSON.stringify(objFormat);
                            inString = inString.replace("'", '"');
                            console.warn('medicine result');
                            console.warn(objFormat)
                        }
                        });
                    }
                });
            });
        });
    }
});

