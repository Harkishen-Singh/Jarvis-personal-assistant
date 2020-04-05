const fs = require('fs');
const _lock = require('lock').Lock;
const Sentence = require('./sentence').Sentence;
const exceptionsAST = require('./constance').exceptionsAST;
const db = require('../utils/db-manager').DBService;

class QueryEngine {
  constructor() {
    this.query = '';
    this.features = {}; // parsing-priority-1
    this.featuresList = [];

    this.initializeFeatures();
  }

  initializeFeatures() {
    const raw = fs.readFileSync('data/features.json');
    this.features = JSON.parse(raw);

    for (const value of Object.values(this.features)) {
      this.featuresList.push(value);
    }
  }

  insertQuery(query) {
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
          let location = db.fetch('/personal/location');
          if (location === undefined) {
            // fetch the location. for now, "bhubaneswar"
            location = 'bhubaneswar';
          }

          // TODO: launch service for fetching weather.
        } else if (feature in this.features.meaning) {
          let entity;

          while (true) {
            const { value } = lexer.next();
            if (value in stopwords) {
              continue;
            }

            // first count
            if (value in exceptionsAST.meaning.continueFirstCount) {
              lexer.getIterValue();
            }
          }
        }
      };
    });
  }
}

const QueryService = new QueryEngine();

module.exports = { QueryService };
