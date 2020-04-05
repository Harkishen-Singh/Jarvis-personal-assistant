const fs = require('fs');
const _lock = require('lock').Lock;
const Sentence = require('./sentence').Sentence;
const exceptionsAST = require('./constance').exceptionsAST;
const db = require('../utils/db-manager').DBService;
const Task = require('./tasks').Task;

let features;

(() => {
  features = JSON.parse(fs.readFileSync('data/features.json'));
})();

class Query {
  constructor(query) {
    this.query = query;
    this.features = features; // parsing-priority-1
    this.featuresList = [];
    this.taskInstance = new Task('qe-main');

    for (const value of Object.values(this.features)) {
      this.featuresList.push(value);
    }

    const sentence = new Sentence(query);
    sentence.tokenize();

    this.process(sentence, sentence.sentenceTokenized);
  }

  bindFeaturePosition(feature, position) {
    return {
      feature,
      position
    };
  }

  getFeaturesAlongWithPositions(fList, queryTokenized) {
    const features = [];
    for (const i in queryTokenized) {
      if (queryTokenized[i]) {
        const word = queryTokenized[i];
        if (word in fList) {
          features.push(this.bindFeaturePosition(word, i));
        }
      }
    }

    return features;
  };

  /**
   * Formats the input into an acceptable form for
   * the task service.
   * @param {string} task
   * @param  {...any} args
   * @return {Promise}
   */
  taskFormat(task, ...args) {
    return {
      task,
      ...args
    };
  }

  run(query) {
    const lock = _lock();

    lock('ops', (release) => {
      const sentence = new Sentence(query);
      sentence.tokenize();
      const stopwords = sentence.stopwords();

      const featurePositions = this.getFeaturesAlongWithPositions(
          this.featuresList,
          sentence.sentenceTokenized
      );
      if (featurePositions > 1) {
        // TODO
        throw new Error(
            'eq: multi-features not supported at the moment.' +
            'Support would be provided in the following versions.'
        );
      }

      const lexer = sentence.getLexerInstance();
      const featurePosition = featurePositions[0].position;
      lexer.setHeadPosition(featurePosition);

      const ast = (feature) => {
        if (feature in this.features.weather) {
          // weather task

          let location;
          if (location === undefined) {
            // if not found, set localtion to "bhubaneswar"
            location = 'bhubaneswar';
          }

          return this.taskFormat('weather', location);
        } else if (feature in this.features.meaning) {
          // meaning task

          let entity;
          while (true) {
            const { status, value } = lexer.next();
            if (!status) {
              lexer.initReverseHEAD();
              continue;
            }

            if (value in stopwords) {
              continue;
            }

            // first count
            if (value in exceptionsAST.meaning.continueFirstCount) {
              if (lexer.getIterValue() === 1) {
                continue;
              }

              // since first skip is already done,
              // the user wants to know the meaning
              // of the word in exception.
              entity = value;
              break;
            }

            entity = value;
            break;
          }

          return this.taskFormat('meaning', entity);
        }
      };

      return new Promise((resolve, reject) => {
        resolve(1);
        reject(new Error('a temporary rejection'));
      });
    });
  }
}

module.exports = { Query };
