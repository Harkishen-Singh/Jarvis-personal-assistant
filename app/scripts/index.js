// eslint-disable-next-line no-undef
const app = angular.module('jarvis-desktop', [ 'ngRoute', 'ngAnimate', 'ngStorage', ]),
	URL = 'http://127.0.0.1:3000',
	// eslint-disable-next-line no-unused-vars
	USER = 'default';

app.config(function($routeProvider) {

	$routeProvider
		.when('/', {
			templateUrl: '../templates/components/main-screen.html',
			controller : 'MainController',
			title      : 'Jarvis - personal assistant',
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

		if (serverResponse) {

			return serverResponse;

		}
		return null;

	};
	return {
		updateServiceStore: updateServiceStore,
		getStore          : getStore,
		getServerResponse : getServerResponse,
	};

});

app.factory('$recentlyUsed', function () {

	let usageArray = [];
	let updateUsageStore = function (tag, plainQuery, message) {

		let usageObject = {
			tag       : tag,
			plainQuery: plainQuery,
			message   : message,
		};
		usageArray.push(usageObject);
		// eslint-disable-next-line no-console
		console.log(usageArray);

	};
	let resetUsageStore = function () {

		usageArray = [];

	};
	let getUsageStore = function () {

		return usageArray;

	};
	return {
		updateUsageStore: updateUsageStore,
		resetUsageStore : resetUsageStore,
		getUsageStore   : getUsageStore,
	};

});

// controllers
app.controller('MainController', function() {
});

app.controller('area-controller', function ($scope, $http, responseService, $recentlyUsed) {

	let supportedTags = [ 'weather', 'google', 'bing', 'yahoo', ];
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

			// for supporting recently used functionality
			for (let v in supportedTags) {

				if (supportedTags[ v ] === query.substring(0, query.indexOf(' ', 0))) {

					$recentlyUsed.updateUsageStore(
						query.substring(
							0, query.indexOf(" ", 0)
						),
						query.substring(
							query.indexOf(" ", 0) + 1,
							query.length
						),
						query,
					);

				}

			}

			let data = 'username=' + USER + '&message=' + query;
			$http({
				url    : URL + '/message',
				method : 'POST',
				headers: {
					'Content-Type': 'application/x-www-form-urlencoded',
				},
				data: data,
			}).then((resp) => {

				let res = resp.data,
					message = res[ 'message' ],
					status = res[ 'status' ],
					result = res[ 'result' ];
				// eslint-disable-next-line no-console
				console.warn(res);

				// $scope types for handling different response types
				$scope.showWeatherScope = false;
				$scope.showQueryScope = false;
				$scope.showMedicineHealthScope = false;

				// response checks
				if (status && message.includes('current weather conditions')) {

					$scope.showWeatherScope = true;
					responseService.updateServiceStore(result, res);

				} else if (status && message.includes('top search results')) {

					$scope.showQueryScope = true;
					responseService.updateServiceStore(result, res);

				} else if (
					('success' === status || status) &&
					(
						'Information about the medicine : ' === message ||
						'Help on the given symptoms : ' === message
					)
				) {

					$scope.showMedicineHealthScope = true;
					responseService.updateServiceStore(res, res);

				}

			});

		} else {

			document.getElementById('user-input-area').value = '';
			$scope.showQueryScope = false;
			$scope.showWeatherScope = false;
			$scope.showMedicineHealthScope = false;

			// re-initialize services
			responseService.updateServiceStore(null, null);

		}

	};

});

app.controller('weather-view-controller', function ($scope, responseService) {

	let serviceStore = responseService.getStore();
	const { temperature: temperature1, } = serviceStore;
	let temperature = temperature1;
	switch (true) {

	case 50 < temperature :
		$scope.tempTag = 'head-weather-data-extreme';
		break;
	case 40 < temperature:
		$scope.tempTag = 'head-weather-data-very-high';
		break;
	case 30 < temperature:
		$scope.tempTag = 'head-weather-data-high';
		break;
	case 20 < temperature:
		$scope.tempTag = 'head-weather-data-mild';
		break;
	case 15 < temperature:
		$scope.tempTag = 'head-weather-data-cool';
		break;
	case 10 < temperature:
		$scope.tempTag = 'head-weather-data-cooler';
		break;
	case 0 < temperature:
		$scope.tempTag = 'head-weather-data-cold';
		break;
	case 0 > temperature:
		$scope.tempTag = 'head-weather-data-extreme-cold ';
		break;
	default:

	}
	$scope.weatherData = serviceStore;

});

app.controller('query-view-controller', function ($scope, responseService) {

	$scope.queryData = responseService.getStore();

});

app.controller('medicine-view-controller', function ($scope, responseService) {

	let serviceStore = responseService.getStore();
	$scope.messageInfo = serviceStore.message;
	$scope.messageResult = serviceStore.result;

});

app.controller('recent-usage-controller', function ($scope, $recentlyUsed) {

	$scope.recentUsageArray = $recentlyUsed.getUsageStore();

});
