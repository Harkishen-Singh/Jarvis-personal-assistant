class Lexer {
  constructor(sentenceTokenized) {
    this.DEFAULT_POSITION = 0;
    this.currentValue = '';
    this.sentenceTokenized = sentenceTokenized;
    this.length = sentenceTokenized.length;
    this.headPosition = this.DEFAULT_POSITION;
  }

  setHeadPosition(position) {
    if (typeof position !== 'number') {
      throw new Error('qe: lexer: lexer position is not a number.');
    }

    if (position >= this.sentenceTokenized.length) {
      throw new Error(
          'qe: lexer: invalid position:' +
          'position cannot be greater than the tokenized sentence size'
      );
    }

    this.headPosition = position;
  }

  current() {
    this.currentValue = this.sentenceTokenized[this.headPosition];
    return this.sentenceTokenized[this.headPosition];
  }

  currentHead() {
    return {
      position: this.headPosition,
      value: this.sentenceTokenized[this.headPosition]
    };
  }

  next() {
    if (this.headPosition + 1 > this.length) {
      // false signifies that the position of the lexer cannot be incremented.
      // This denotes that the current position is already on the last
      // word of the sentence.
      return { status: false, value: null };
    }

    this.headPosition += 1;
    return { status: true, value: this.sentenceTokenized[this.headPosition] };
  }

  previous() {
    if (this.headPosition == 0) {
      // false signifies that the position of lexer cannot be decremented.
      // This denotes that the current position is already on the first
      // word of the sentence.
      return { status: false, value: null };
    }

    this.headPosition -= 1;
    return { status: true, word: this.sentenceTokenized[this.headPosition] };
  }

  reset() {
    this.headPosition = this.DEFAULT_POSITION;
    return { status: true, word: this.sentenceTokenized[this.headPosition] };
  }
}

module.exports = { Lexer };
