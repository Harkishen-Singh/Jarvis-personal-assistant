const fs = require('fs');
const Lexer = require('./lexer').Lexer;
let stopwords = [];

(() => {
  const raw = fs.readFileSync('data/stopwords.json');
  stopwords = JSON.parse(raw);
})();

class Sentence {
  constructor(sentence) {
    this.sentence = sentence;
    this.sentenceTokenized = [];
    this.stopwordsValue = [];

    // bools
    this.isTokenized = false;
    this.isStopwordsCalculated = false;

    // lexer
    this.tokenize();
    this.lexer = new Lexer(this.sentenceTokenized);
  }

  tokenize() {
    this.sentenceTokenized = this.sentence.split(' ');
    this.isTokenized = true;
    return this.sentenceTokenized;
  }

  stopwords() {
    const words = [];
    let tmp;
    if (!this.isTokenized) {
      tmp = this.tokenize();
    } else {
      tmp = this.sentenceTokenized;
    }

    for (const word in tmp) {
      if (word in stopwords) {
        words.push(word);
      }
    }

    this.isStopwordsCalculated = true;
    this.stopwordsValue = tmp;
    return words;
  }

  getLexerInstance() {
    return this.lexer;
  }
}

module.exports = { Sentence };
