const date = new Date();
const data = {
  onFirstStart: {
    date: date.getDate(),
    month: date.getMonth(),
    year: date.getFullYear()
  },
  store: {
    gc: {}
  },
  personal: {
    location: 'delhi'
  }
};

module.exports = { data };
