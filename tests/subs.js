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
        var arrAnswer = [];

        var driver = new webdriver.Builder()
                            .setChromeOptions(options)
                            .forBrowser('chrome')
                            .build();

        driver.get(url).then(() => {
            driver.findElements(By.className('list-item')).then(cc => {
                var count = 0;
                cc.forEach(each => {
                    if (count %2 !== 0) {
                        count++;
                        console.log(count + ' here')
                        return;
                    } else count++;
                    each.getAttribute('innerHTML').then(ee => {
                        // console.log(ee)
                        var text = ee, got = false;
                        var len = text.length, c=0;
                        for(var i = 0; i < len; i++) {
                            if(text[i] === '>')
                                c++;
                            if(c === 2) {
                                c = 0;
                                for (var j=1; ; j++) {
                                    if (text.substring(i+j, i+j+4) === '</a>') {
                                        let stringss = text.substring(i+1, i+j-1)
                                        console.log('this -> ' + stringss + ' >==<ends here')
                                        arrAnswer.push(stringss.trim())
                                        got = true
                                        break;
                                    }
                                }
                            }
                            if (got)
                                break
                        }
                        console.log(arrAnswer)
                    })
                })
            });
            // driver.findElements(By.js("document.querySelector('body > div.container.mi-container > div.mi-container__left > div > div.related-links.top-gray.col-list.clear-fix > ul:nth-child(16) > li:nth-child(121) > h4 > a')")).then(cc => {
            //     cc.forEach(each => {
            //         each.getAttribute('innerHTML').then(ee => {
            //             console.log(ee)
            //         })
            //     })
            // })
        });
    }
})

