const queryController = require('./queryController').QueryController;

class AST {
  serialize(firstParam, secondParam) {
    return new Promise(async (resolve, reject) => {
      const queryCont = new queryController();
      if (firstParam == 'google') {
        console.log('first param: ', firstParam);
        const response = await queryCont.HandleGoogleQuery(secondParam);
        // console.log('response:: ', response);
        // // return response;


        resolve(response);
      }
    });
  }
}

module.exports = { AST };
