const queryController = require('./queryController').QueryController;

class AST {

    serialize(firstParam, secondParam) {
        console.log("first param: ", firstParam)
        const queryCont = new queryController();
        if (firstParam == "google") {
            var response = queryCont.HandleGoogleQuery(secondParam)
                    console.log("response:: ", response)
                    return response
        }

    }

}

module.exports = { AST }