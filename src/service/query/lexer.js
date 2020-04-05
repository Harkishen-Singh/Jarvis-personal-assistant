class Lexer {
  constructor(sentenceTokenized) {
    this.DEFAULT_POSITION = 0;
    this.DEFAULT_ITER_COUNT = 0;
    this.currentValue = '';
    this.sentenceTokenized = sentenceTokenized;
    this.length = sentenceTokenized.length;
    this.headPosition = this.DEFAULT_POSITION;
    this.iterCount = this.DEFAULT_ITER_COUNT;
    this.cache = this.DEFAULT_POSITION;
    this.forwardDirection = true;
  }

  setHeadPosition(position) {
    if (typeof position !== 'number') {
      throw new Error('qe: lexer: lexer position is not a number.');
    }

    this.cache = position;

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

  /**
   * Returns the value of the number of positions the HEAD is
   * away from the initial position.
   * +ve -> HEAD is right of the initial position.
   * -ve -> HEAD is left of the initial position.
   * @return {number}
   */
  getIterValue() {
    return this.iterCount;
  }

  /**
   * Increments the position of the lexer HEAD by one position
   * and returns the value of the HEAD along with a status.
   * If the status is true, the next operation is successful,
   * else the value could not be incremented due to overflow.
   * @param {Number} sign
   * @return {{Boolean, String}} status, value
   */
  next() {
    if (!this.forwardDirection) this.previous;
    if (this.headPosition + 1 > this.length) {
      // false signifies that the position of the lexer cannot be incremented.
      // This denotes that the current position is already on the last
      // word of the sentence.
      return { status: false, value: null };
    }

    this.iterCount += 1;
    this.headPosition += 1;
    return { status: true, value: this.sentenceTokenized[this.headPosition] };
  }

  initReverseHEAD() {
    this.headPosition = this.cache;
    this.forwardDirection = !this.forwardDirection;
  }

  previous() {
    if (this.headPosition == 0) {
      // false signifies that the position of lexer cannot be decremented.
      // This denotes that the current position is already on the first
      // word of the sentence.
      return { status: false, value: null };
    }

    this.iterCount -= 1;
    this.headPosition -= 1;
    return { status: true, word: this.sentenceTokenized[this.headPosition] };
  }

  reset() {
    this.iterCount = this.DEFAULT_ITER_COUNT;
    this.headPosition = this.DEFAULT_POSITION;
    return { status: true, word: this.sentenceTokenized[this.headPosition] };
  }
}

module.exports = { Lexer };
