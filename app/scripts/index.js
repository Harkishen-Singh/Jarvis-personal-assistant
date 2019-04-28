console.warn('index script running');

const app = angular.module('jarvis-desktop', ['ngRoute']),
	URL = 'http://127.0.0.1:3000',
	USER = 'default';
console.log('From app-jarvis.js');

app.config(function($routeProvider) {
	$routeProvider
		.when('/', {
			templateUrl: '../templates/components/main-screen.html',
			controller: 'MainController',
			title: 'Jarvis - personal assistant',
		});
});

app.controller('MainController', function($scope) {
    console.warn('angular is working')
})