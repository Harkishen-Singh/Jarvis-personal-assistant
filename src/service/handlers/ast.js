/* eslint-disable new-cap */
const queryController = require('./queryController').QueryController;

class AST {
  async serialize(firstParam, secondParam) {
    console.log('first param: ', firstParam);
    const queryCont = new queryController();
    if (firstParam == 'google') {
      const response = await queryCont.HandleGoogleQuery(secondParam);
      console.log('response:: ', response);
      return response;
    }
  }
}

module.exports = { AST };
