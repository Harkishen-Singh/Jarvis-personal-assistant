const fs = require('fs');
const _lock = require('lock').Lock;
const Sentence = require('./sentence').Sentence;
const exceptionsAST = require('./constance').exceptionsAST;
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
  }

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

      const lexer = sentence.getLexer();
      const { feature, position } = featurePositions[0];
      lexer.setHeadPosition(position);

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
          let twoWayParsing = false;
          while (true) {
            // The lexing starts with first parsing all the words
            // to the right (since meaning has mmore weight to the right
            // of the feature mostly) of the feature word (or node).
            // Break and reset after all right words are parsed. To do this,
            // the HEAD to initial position using cache and begin parsing
            // the left nodes. Once the left end is reached
            // stop the parsing since all words are parsed.
            const { status, value } = lexer.next();
            if (!status) {
              if (twoWayParsing) {
                // stop after entire query is parsed to
                // avoid endless loop and false cpu cycles.
                break;
              }

              twoWayParsing = true;
              lexer.initReverseHEAD();
              continue;
            }

            {
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
            }

            entity = value;
            break;
          }

          return this.taskFormat('meaning', entity);
        }
      };

      return new Promise(async (resolve) => {
        const task = ast(feature);
        const result = await this.taskInstance.run(task);
        resolve(result);
      });
    });
  }
}

module.exports = { Query };
