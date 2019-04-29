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
		let ele0 = document.getElementById('jarvis-bot-area-id'),
			ele1 = document.getElementById('ele1-move'),
			ele2 = document.getElementById('ele2-move');
		ele0.classList.toggle('hide');
		ele1.classList.toggle('logo-post-query');
		ele2.classList.toggle('user-area-post-query');
	};
});
