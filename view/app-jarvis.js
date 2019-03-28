/* eslint-disable no-console */
// eslint-disable-next-line no-undef
const app = angular.module('jarvis', ['ngRoute']),
	URL = 'http://127.0.0.1:3000',
	USER = 'default';
console.warn('ggggg')

app.config(function($routeProvider) {
	$routeProvider
		.when('/', {
			templateUrl:'./components/main.html',
			controller:'MainController',
			title:'Jarvis - personal assistant',
		});
});

app.controller('MainController', function($scope,$location,$rootScope,$http) {

	// eslint-disable-next-line no-undef
	var recognition = new webkitSpeechRecognition();
	var recognizing;

	$scope.messageStack = [];
	$scope.showLoaderListening = false;

	$scope.addMessagesToStack = function() {
		if (!$scope.message.startsWith('Type a message')) {

			if ($scope.showLoaderListening) {
				$scope.showLoaderListening = false;
				recognition.stop();
				recognizing = false;
			}

			let message = $scope.message,
				date = new Date(),
				hrs = date.getHours(),
				mins = date.getMinutes(),
				messageObj = {
					message: '',
					sender: '',
					time: '',
					length: null
				},
				data = null;

			messageObj.message = message;
			messageObj.length = message.length;
			messageObj.time = String(hrs + ':' + mins);
			messageObj.sender = 'you';

			$scope.messageStack.push(messageObj);
			data = 'username='+USER+'&message='+messageObj.message;

			$http({
				url:URL+'/message',
				method:'POST',
				headers: {
					'Content-Type': 'application/x-www-form-urlencoded'
				},
				data:data
			}).then(resp => {
				let res = (resp.data),
					message = res['message'],
					status = res['status'],
					hrs2 = new Date().getHours(),
					mins2 = new Date().getMinutes();
				messageObj = {
					message: '',
					sender: '',
					time: '',
					length: null
				};

				console.log(res);
				if (status === 'success' || status) {
					messageObj.sender = 'jarvis-bot';
					messageObj.time = String(hrs2 + ':' + mins2);
					messageObj.length = message.length;
					messageObj.message = message;
				} else {
					console.error('[JARVIS] error fetching from service.');
				}

				// output view

			}).catch(e => {
				throw e;
			});

			$scope.message = 'Type a message ...';
		} else {
			alert('Please enter a message');
		}
	};

	$scope.removeMessage = function(){
		if($scope.message.startsWith('Type a message ...')){
			$scope.message = '';
		}
	};

	$scope.initStack = function() {
		$scope.message = 'Type a message ...';
	};

	$scope.toggleStartStop = function() {
		recognition.continuous = true;
		// recognition.onend = reset;

		recognition.onresult = function (event) {
			var mess = document.getElementById('message-input');
			mess.value = '';
			for (var i = 0; i < event.results.length; i++) {
				if (event.results[i].isFinal) {
					mess.value += event.results[i][0].transcript;
					if (mess.value.endsWith('send')) {
						var n = mess.value.lastIndexOf('send');
						var submessage =  mess.value.substring(0,n);
						$scope.message = submessage;
						$scope.addMessagesToStack();
					} else {
						$scope.message = mess.value;
					}
				} else {
					mess.value += event.results[i][0].transcript;
					$scope.message = mess.value;
				}
			}
		};

		var button = document.getElementById('button');
		if (recognizing) {
			recognition.stop();
			$scope.showLoaderListening = false;
			recognizing = false;
			// eslint-disable-next-line no-undef
			button.innerHTML = 'Click to Speak';
		} else {
			recognition.start();
			$scope.showLoaderListening = true;
			recognizing = true;
			$scope.message = '';
		}
	};

});




