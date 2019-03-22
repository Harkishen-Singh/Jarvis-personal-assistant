const webdriver = require('selenium-webdriver'),
    firefox = require('selenium-webdriver/firefox'),
    By = webdriver.By;

{ describe, it, after, before } require('selenium-webdriver/testing')

require('chromedriver');
require('geckodriver');

var driver;

var options = new firefox.Options();
options.addArguments("--no-sandbox");
options.addArguments("--disable-dev-shm-usage");
options.addArguments("--disable-gpu");

describe("Firefox Environments", function() {
    this.timeout(60000);

    it("Creating virtual browser instances", done => {
        driver = new webdriver.Builder()
                    .setFirefoxOptions(options)
                    .forBrowser('firefox')
                    .build();
        driver.then(() => {
            done();
        });
    });

    it("Opening Jarvis assistant", done => {
        driver.get('http://127.0.0.1:8080').then(() => {
            done();
        });
    });

    describe("Jarvis assistant operations", function() {
        this.timeout(30000);

        it("Search for message bar", done => {
            driver.findElement(By.xpath('//*[@id="message-input"]')).then(() => {
                done();
            });
        });

        it("Insert message", done => {
            driver.findElement(By.xpath('//*[@id="message-input"]')).click();
            driver.findElement(By.xpath('//*[@id="message-input"]')).then(vals => {
                vals.sendKeys('Hi! This is selenium Bot.');
                driver.findElement(By.xpath('//*[@id="message-bar-send"]')).click().then( rr => {
                    done();
                });
            });
        });

        it("Show message in chat screen", done => {
            driver.findElement(By.xpath('//*[@id="stackArea"]/div')).then(() => {
                done();
            });
        });

        it("Preventing messages for empty mesages, using alert warning check", done => {
            driver.findElement(By.xpath('//*[@id="message-bar-send"]')).click().then( rr => {
                driver.switchTo().alert().then(() => {
                    done();
                });
            });
        });

        it("Closing browser", () => {
            driver.quit();
        });

    });
});