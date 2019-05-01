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
let imported0 = document.createElement('script');
imported0.src = '../scripts/connect.js';
document.head.appendChild(imported0);

app.config(function($routeProvider) {
	$routeProvider
		.when('/', {
			templateUrl: '../templates/components/main-screen.html',
			controller: 'MainController',
			title: 'Jarvis - personal assistant',
		});
});

// services

app.factory('weatherResponseService', function () {
	let serverResponse = {},
		serviceStore = {};
	let updateServiceStore = function (object, response = {}) {
		serverResponse = response;
		serviceStore = JSON.parse(object);
		return true;
	};
	let getStore = function () {
		return serviceStore;
	};
	let getServerResponse = function () {
		if (serverResponse)
			return serverResponse;
		return null;
	};
	return {
		updateServiceStore: updateServiceStore,
		getStore: getStore,
		getServerResponse : getServerResponse
	};
});

// controllers

app.controller('MainController', function($scope) {
});

app.controller('area-controller', function ($scope, $http, weatherResponseService) {
	$scope.Initialize = function () {
		$scope.showJarvisBotArea = true;
		$scope.showLabel = true;
		$scope.showReset = false;
	};
	$scope.AskButtonClick = function (query) {
		$scope.showJarvisBotArea = false;
		$scope.showLabel = !$scope.showLabel;
		$scope.showReset = !$scope.showReset;
		console.warn($scope);
		let ele0 = document.getElementById('jarvis-bot-area-id'),
			ele1 = document.getElementById('ele1-move'),
			ele2 = document.getElementById('ele2-move'),
			ele3 = document.getElementById('userButton'),
			ele4 = document.getElementById('ele4-move'),
			ele6 = document.getElementById('jarvis-message-id');
		ele0.classList.toggle('hide');
		ele1.classList.toggle('logo-post-query');
		ele2.classList.toggle('user-area-post-query');
		ele3.classList.toggle('user-input-ask-button-post-query');
		ele4.classList.toggle('user-input-outer-layer-post-query');
		ele6.classList.toggle('message-jarvis-bot-post-query');
		let data = 'username=' + USER + '&message=' + query;
		$http({
			url:URL+'/message',
			method:'POST',
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded'
			},
			data: data
		}).then(resp => {
			let res = (resp.data),
				message = res['message'],
				status = res['status'],
				result = res['result'],
				show = res['show'],
				hrs2 = new Date().getHours(),
				mins2 = new Date().getMinutes(),
				messageObj = {
					message: '',
					sender: '',
					time: '',
					result: '',
					show: false,
					length: null
				};
			console.log(res);
			// $scope types for handling different response types
			$scope.showWeatherScope = false;

			// response checks
			if (status && message.includes('current weather conditions')) {
				$scope.showWeatherScope = true;
				weatherResponseService.updateServiceStore(result, res);
			}
		});
	};
});

app.controller('weather-view-controller', function ($scope, weatherResponseService) {
	console.warn('weather-view-controller called');
	let response = {},
		serviceStore = {};
	response = weatherResponseService.getServerResponse();
	serviceStore = weatherResponseService.getStore();
	console.warn('from weather controller');
	console.warn(serviceStore);
	console.warn(response);
})
