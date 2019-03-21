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
    var messageStack = [];

    $scope.addMessagesToStack = function() {
        let message = $scope.message;
        messageStack.push(message);
        console.warn(messageStack);
    }

    $scope.initStack = () => {
        $scope.message = "Type a message ...";
    }

});