var app = angular.module('jarvis', ['ngRoute']);

app.config(function($routeProvider,$locationProvider) {
    $routeProvider
    .when("/", {
        templateUrl:'./components/main.html',
        controller:'MainController',
        title:'Jarvis - personal assistant',
    })
});

app.controller('MainController', function($scope,$location,$rootScope,$http) {
    // var messageStack = [];
    const date = new Date();

    $scope.messageStack = [];
    $scope.addMessagesToStack = function() {
        let message = $scope.message,
            hrs = date.getHours(),
            mins = date.getMinutes(),
            messageObj = {
                message: '',
                sender:'',
                time:'',
                length: null
            };

        messageObj.message = message;
        messageObj.length = message.length;
        messageObj.time = String(hrs + ':' + mins);
        messageObj.sender = 'you';

        $scope.messageStack.push(messageObj);

        console.warn($scope.messageStack);
    }

    $scope.initStack = function() {
        $scope.message = "Type a message ...";
    }

});