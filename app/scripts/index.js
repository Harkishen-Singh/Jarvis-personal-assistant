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

app.factory('responseService', function () {
	let serverResponse = {},
		serviceStore = {};
	let updateServiceStore = function (object, response = {}) {
		serverResponse = response;
		try {
			serviceStore = JSON.parse(object);
		} catch (e) {
			serviceStore = object;
		}

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
app.controller('MainController', function() {
});

app.controller('area-controller', function ($scope, $http, responseService) {
	$scope.Initialize = function () {
		$scope.showJarvisBotArea = true;
		$scope.showLabel = true;
		$scope.showReset = false;
	};
	$scope.AskButtonClick = function (query) {
		$scope.showJarvisBotArea = false;
		$scope.showLabel = !$scope.showLabel;
		$scope.showReset = !$scope.showReset;
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
		if (query) {
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
					result = res['result'];
				// eslint-disable-next-line no-console
				console.warn(res);

				// $scope types for handling different response types
				$scope.showWeatherScope = false;
				$scope.showQueryScope = false;
				$scope.showMedicine_HealthScope = false;

				// response checks
				if (status && message.includes('current weather conditions')) {
					$scope.showWeatherScope = true;
					responseService.updateServiceStore(result, res);
				} else if (
					(status === 'success' || status) &&
					(
						message === 'Information about the medicine : ' ||
						message === 'Help on the given symptoms : '
					)
				) {
					$scope.showMedicine_HealthScope = true;
					responseService.updateServiceStore(res, res);
				}
			});
		} else {
			document.getElementById('user-input-area').value = '';
			$scope.showWeatherScope = false;
			$scope.showMedicine_HealthScope = false;

			// re-initialize services
			responseService.updateServiceStore(null, null);
		}
	};
});


app.controller('weather-view-controller', function ($scope, responseService) {
	let serviceStore = responseService.getStore();
	let temperature = serviceStore.temperature;
	switch (true) {
	case temperature > 50 :
		$scope.tempTag = 'head-weather-data-extreme';
		break;
	case temperature > 40:
		$scope.tempTag = 'head-weather-data-very-high';
		break;
	case temperature > 30:
		$scope.tempTag = 'head-weather-data-high';
		break;
	case temperature > 20:
		$scope.tempTag = 'head-weather-data-mild';
		break;
	case temperature > 15:
		$scope.tempTag = 'head-weather-data-cool';
		break;
	case temperature > 10:
		$scope.tempTag = 'head-weather-data-cooler';
		break;
	case temperature > 0:
		$scope.tempTag = 'head-weather-data-cold';
		break;
	case temperature < 0:
		$scope.tempTag = 'head-weather-data-extreme-cold ';
		break;
	default:
		// eslint-disable-next-line no-mixed-spaces-and-tabs
 	}
	$scope.weatherData = serviceStore;
});

app.controller('medicine-view-controller', function ($scope, responseService) {
	let serviceStore = responseService.getStore();
	$scope.messageInfo = serviceStore.message;
	$scope.messageResult = serviceStore.result;
});
