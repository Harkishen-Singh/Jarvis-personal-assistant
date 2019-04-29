/* eslint-disable no-console */
/* eslint-disable no-unused-vars */
// eslint-disable-next-line no-console
console.warn('index script running');

// eslint-disable-next-line no-undef
const app = angular.module('jarvis-desktop', ['ngRoute', 'ngAnimate']),
	URL = 'http://127.0.0.1:3000',
	USER = 'default';
// eslint-disable-next-line no-console
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
});

app.controller('area-controller', function ($scope) {
	$scope.Initialize = function () {
		$scope.showJarvisBotArea = true;
	};
	$scope.AskButtonClick = function () {
		$scope.showJarvisBotArea = false;
		let ele = document.getElementById('jarvis-bot-area-id');
		ele.classList.toggle('hide');
	};
});
