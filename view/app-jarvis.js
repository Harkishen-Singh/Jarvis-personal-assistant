/* eslint-disable no-console */
// eslint-disable-next-line no-undef
const app = angular.module('jarvis', ['ngRoute']),
	URL = 'http://127.0.0.1:3000',
	USER = 'default';

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
				console.log(recognizing);
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
				let res = resp.data;
				console.log(res);
			}).catch(e => {
				throw e;
			});

			$scope.message = 'Type a message ...';
			if (!recognizing) {
				console.log('adfa');
				setTimeout(function()
				{ 
					$scope.toggleStartStop();
				}, 
				2000);
			}

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
		$scope.toggleStartStop();
	};

	$scope.toggleStartStop = function() {
		recognition.continuous = true;

		recognition.onresult = function (event) {
			var n, submessage;
			var mess = document.getElementById('message-input');
			mess.value = '';
			for (var i = 0; i < event.results.length; i++) {
				if (event.results[i].isFinal) {
					mess.value += event.results[i][0].transcript;
					if (mess.value.endsWith('send')) {
						n = mess.value.lastIndexOf('send');
						submessage =  mess.value.substring(0,n);
						$scope.message = submessage;
						$scope.addMessagesToStack();
					} 
					else if (mess.value.includes('start jarvis')) {
						n = mess.value.lastIndexOf('start jarvis');
						submessage = mess.value.substring(n+12);
						mess.value = submessage;
						$scope.message = submessage; 
					}
					else {
						$scope.message = mess.value;
					}
					
				} else {
					mess.value += event.results[i][0].transcript;
					if (mess.value.includes('start jarvis')) {
						n = mess.value.lastIndexOf('start jarvis');
						submessage = mess.value.substring(n+12);
						mess.value = submessage;
					}
					$scope.message = mess.value;
				}
			}
		};

		if (recognizing) {
			recognition.stop();
			console.log(recognition);
			$scope.showLoaderListening = false;
			recognizing = false;
		} else {
			recognition.start();
			console.log(recognition);
			console.log(recognizing);
			$scope.showLoaderListening = true;
			console.log($scope.showLoaderListening);
			recognizing = true;
			$scope.message = '';
			console.log(recognizing);
		}
	};

});




