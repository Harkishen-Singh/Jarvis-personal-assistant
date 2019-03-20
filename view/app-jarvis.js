const app = angular.module('jarvis', ['ngRoute']);

app.config(function($routeProvider,$locationProvider) {
    $routeProvider
    .when("/", {
        templateUrl:'./components/main.html',
        controller:'MainController',
        title:'Jarvis - personal assistant',
    })
});

app.controller('MainController', function($scope,$location,$rootScope,$http) {
    console.warn('main controller called');
});