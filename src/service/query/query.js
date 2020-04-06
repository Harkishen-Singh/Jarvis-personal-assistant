const fs = require('fs');
const Sentence = require('./sentence').Sentence;
const exceptionsAST = require('./constance').exceptionsAST;
const Task = require('./tasks').Task;

let features;

(() => {
  features = JSON.parse(fs.readFileSync(`${__dirname}/data/features.json`));
})();

class Query {
  constructor(query) {
    this.query = query;
    this.features = features; // parsing-priority-1
    this.featuresList = [];
    this.taskInstance = new Task('qe-main');

    for (const [, values] of Object.entries(this.features)) {
      for (const value of values) {
        this.featuresList.push(value);
      }
    }

    // this.process(sentence, sentence.sentenceTokenized);
  }

  bindFeaturePosition(feature, position) {
    return {
      feature,
      position: parseInt(position)
    };
  }

  getFeaturesAlongWithPositions(fList, queryTokenized) {
    const features = [];
    for (const i in queryTokenized) {
      if (queryTokenized[i]) {
        const word = queryTokenized[i];
        if (fList.includes(word)) {
          features.push(this.bindFeaturePosition(word, i));
        }
      }
    }

    return features;
  }

  taskFormat(task, object) {
    return {
      task,
      object
    };
  }

  run(query=this.query) {
    const sentence = new Sentence(query);
    sentence.tokenize();
    const stopwords = sentence.stopwords();

    const featurePositions = this.getFeaturesAlongWithPositions(
        this.featuresList,
        sentence.sentenceTokenized
    );
    if (featurePositions > 1) {
      // TODO: features can be sorted based on priority in the later stage.
      throw new Error(
          'eq: multi-features not supported at the moment.' +
            'Support would be provided in the following versions.'
      );
    }

    const lexer = sentence.getLexer();
    const { feature, position } = featurePositions[0];
    lexer.setHeadPosition(position);

    const filters = {
      weather: (txt) => {
        txt = txt.replace('weather ', '');
        const arr = txt.split(',');
        const city = arr[0];
        const state = arr[1];
        const country = arr[2];
        return { city, state, country };
      }
    };

    const ast = (feature) => {
      if (this.features.weather.includes(feature)) {
        // weather task
        // format: weather city,state,country
        // example: weather bhubaneswar,odisha,india

        const taskInput = filters.weather(sentence.getOriginalSentece());

        return this.taskFormat('weather', taskInput);
      } else if (this.features.meaning.includes(feature)) {
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
            if (stopwords.includes(valiue)) {
              continue;
            }

            // first count
            if (exceptionsAST.meaning.continueFirstCount.includes(value)) {
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
      const result = await this.taskInstance.run(task.task, task.object);
      resolve(result);
    });
  }
}

module.exports = { Query };
