/**
 * COSS society
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
let email = 'harkishensingh@hotmail.com',
    githubRepo = null,
    password = '',
    url = 'heroku.com';
process.argv.forEach((val, index, array) => {
    var path = require('chromedriver').path;

    var service = new chrome.ServiceBuilder(path).build();
    chrome.setDefaultService(service);
    if (index === 2) { // corresponds to the emailID heroku
        githubRepo = val;
        // url = val;

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

        driver.get('https://id.heroku.com/login').then(() => {
            console.log('opened1')
            driver.findElement(By.xpath('//*[@id="email"]')).sendKeys(email).then(() => {
                driver.findElement(By.xpath('//*[@id="password"]')).sendKeys(password).then(() => {
                    driver.findElement(By.xpath('//*[@id="login"]/form/button')).click().then(() => {
                        driver.get('https://id.heroku.com/login').then(() => {
            console.log('opened2')
            driver.findElement(By.xpath('//*[@id="email"]')).sendKeys(email).then(() => {
                driver.findElement(By.xpath('//*[@id="password"]')).sendKeys(password).then(() => { 
                    driver.findElement(By.xpath('//*[@id="login"]/form/button')).click().then(() => {
                        driver.wait(() => {
                            return driver.executeScript('return document.readyState').then(function(readyState) {
                                return readyState === 'complete';
                            });
                        })
                        driver.sleep(5000).then(() => {
                            console.log('wait over. page loaded');
                            driver.findElement(By.xpath('//*[@id="ember159"]/button')).click().then(() => {
                                driver.wait(() => {
                                    return driver.executeScript('return document.readyState').then(function(readyState) {
                                        return readyState === 'complete';
                                    });
                                })
                                console.log('page loaded;')
                                driver.findElement(By.xpath('//*[@id="ember183"]')).click().then(() => {
                                    driver.wait(() => {
                                        return driver.executeScript('return document.readyState').then(function(readyState) {
                                            return readyState === 'complete';
                                        });
                                    });
                                    driver.sleep(3000).then(() => {
                                        // send app name
                                        driver.findElement(By.xpath('//*[@id="ember202"]')).click().then(() => {
                                            console.log('app created!');
                                            driver.sleep(4000).then(() => {
                                                driver.findElement(By.className('deploy-tab tab-github')).click().then(() => {
                                                    driver.sleep(3000).then(() => {
                                                        driver.findElement(By.xpath('//*[@id="search-term"]')).sendKeys(githubRepo).then(() => {
                                                            driver.findElement(By.className('br--right bl-0 async-button default hk-button--primary ember-view')).click().then(() => {
                                                                driver.sleep(4000).then(() => {
                                                                    driver.findElement(By.className('async-button default hk-button-sm--secondary ember-view')).click().then(() => {
                                                                        console.log('connectd to github repo!');
                                                                        driver.sleep(5000).then(() => {
                                                                            driver.findElements(By.className('btn btn-primary btn-github')).then(eleLoop => {
                                                                                eleLoop.forEach((item, index) => {
                                                                                    item.click();
                                                                                    if (index == 1) {
                                                                                        driver.sleep(3000).then(() => {
                                                                                            driver.wait(webdriver.until.elementLocated(By.className('btn btn-default btn-sm'))).then(full => {
                                                                                                driver.findElement(By.className('btn btn-default btn-sm')).then(elLink => {
                                                                                                    elLink.getAttribute('href').then((link) => {
                                                                                                        console.log('link to the hosted app ', link);
                                                                                                        driver.quit();
                                                                                                    });
                                                                                                });
                                                                                            })
                                                                                        });
                                                                                    }
                                                                                });
                                                                            });
                                                                        });
                                                                    });
                                                                });
                                                            });
                                                        });
                                                    });
                                                });
                                            });
                                        });
                                    });
                                });
                            });
                        });
                    });
                });
            });
        });
    });
                });
            });
        });
    }
})

