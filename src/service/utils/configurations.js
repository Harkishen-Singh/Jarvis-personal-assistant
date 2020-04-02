const date = new Date();
const configurations = {
  onFirstStart: {
    date: date.getDate(),
    month: date.getMonth(),
    year: date.getFullYear()
  },
  store: {
    gc: {}
  }
};

module.exports = {configurations};
