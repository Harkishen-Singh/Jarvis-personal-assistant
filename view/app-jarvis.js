/* eslint-disable linebreak-style */
/* eslint-disable no-console */
// eslint-disable-next-line no-undef
const app = angular.module('jarvis', ['ngRoute']),
	URL = 'http://127.0.0.1:3000',
	USER = 'default';
console.log('From app-jarvis.js');
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

	$scope.controlMainBanner = function() {
		$scope.mainBanner = true;
		setTimeout(() =>{
			$scope.mainBanner = false;
		}, 500);
	};
	$scope.messageStack = [];
	$scope.showLoaderListening = false;

	var input = document.getElementById('message-input');
	input.addEventListener('keyup', function(event) {
		if (event.keyCode === 13) {
			event.preventDefault();
			document.getElementById('message-bar-send').click();
		}
	});

	$scope.addMessagesToStack = function() {
		if ($scope.message.length) {

			if ($scope.showLoaderListening) {
				$scope.showLoaderListening = false;
				recognition.stop();
				recognizing = false;
			}

			var mess = document.getElementById('message-input');
			mess.value = '';
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
			$scope.showLoading = true;
			setTimeout(() => {
				$scope.scrollDown();
			}, 100);
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
					result = res['result'],
					show = res['show'],
					hrs2 = new Date().getHours(),
					mins2 = new Date().getMinutes();
				messageObj = {
					message: '',
					sender: '',
					time: '',
					result: '',
					show: false,
					length: null
				};
				console.log(messageObj);
				console.log(res);
				setTimeout(() => {
					$scope.scrollDown();
				}, 100);
				if (status && message === 'here are the current weather conditions') {
					messageObj.sender = 'jarvis-bot';
					messageObj.time = String(new Date().getHours() + ':' + new Date().getMinutes());
					messageObj.length = message.length;
					messageObj.message = message;
					messageObj.result = JSON.parse(result);
					$scope.messageStack.push(messageObj);
					$scope.showLoading = false;
				} else if (status && message === 'here is the meaning of the searched word') {
					messageObj.sender = 'jarvis-bot';
					messageObj.time = String(new Date().getHours() + ':' + new Date().getMinutes());
					messageObj.length = message.length;
					messageObj.message = message;
					messageObj.result = result;
					$scope.messageStack.push(messageObj);
					$scope.showLoading = false;
				} else if ((status === 'success' || status) && message === 'here are the top search results' ) {
					messageObj.sender = 'jarvis-bot';
					messageObj.time = String(new Date().getHours() + ':' + new Date().getMinutes());
					messageObj.length = message.length;
					messageObj.message = message;
					messageObj.result = result;
					$scope.messageStack.push(messageObj);
					$scope.showLoading = false;
				} else if ((status === 'success' || status) && message === 'here are the searched images' ) {
					messageObj.sender = 'jarvis-bot';
					messageObj.time = String(hrs2 + ':' + mins2);
					messageObj.length = message.length;
					messageObj.message = message;
					messageObj.result = result;
					$scope.messageStack.push(messageObj);
					$scope.showLoading = false;
				} else if ((status === 'success' || status) && message === 'Enter Reminder details : ') {
					messageObj.sender = 'jarvis-bot';
					messageObj.time = String(new Date().getHours() + ':' + new Date().getMinutes());
					messageObj.length = message.length;
					messageObj.message = message;
					$scope.messageStack.push(messageObj);
				} else if ((status === 'success' || status) && message === 'Here are your reminders : ') {
					messageObj.sender = 'jarvis-bot';
					messageObj.time = String(new Date().getHours() + ':' + new Date().getMinutes());
					messageObj.length = message.length;
					messageObj.message = message;
					messageObj.result = result;
					$scope.messageStack.push(messageObj);
				} else if ((status === 'success' || status) && (message === 'Information about the medicine : ' || message === 'Help on the given symptoms : ')) {
					messageObj.sender = 'jarvis-bot';
					messageObj.time = String(hrs2 + ':' + mins2);
					messageObj.length = message.length;
					messageObj.message = message;
					messageObj.result = result;
					$scope.messageStack.push(messageObj);
					$scope.showLoading = false;
				} else if ((status === 'success' || status) && !show) {
					messageObj.sender = 'jarvis-bot';
					messageObj.time = String(new Date().getHours() + ':' + new Date().getMinutes());
					messageObj.length = message.length;
					messageObj.message = message;
					$scope.messageStack.push(messageObj);
					$scope.showLoading = false;
				} else if (show) {
					messageObj.sender = 'jarvis-bot';
					messageObj.time = String(new Date().getHours() + ':' + new Date().getMinutes());
					messageObj.length = message.length;
					messageObj.message = message;
					messageObj.show = show;
					$scope.messageStack.push(messageObj);
					$scope.showLoading = false;
				} else {
					console.error('[JARVIS] error fetching from service.');
				}
			}).catch(e => {
				throw e;
			});
			$scope.message = '';
			// if (!recognizing) {
			// 	setTimeout(() => {
			// 		$scope.toggleStartStop(0);
			// 	}, 2000);
			// }

		}
	};

	$scope.formData = {};
	$scope.setReminder = function() {
		$scope.messageStack.pop();
		let reminder_title = $scope.formData.remTitle,
			reminder_description = $scope.formData.remDescription,
			reminder_time = $scope.formData.remTime,
			reminderObj = {
				title: '',
				description: '',
				time: ''
			},
			data = null;
		reminderObj.title = reminder_title;
		reminderObj.description = reminder_description;
		reminderObj.time = reminder_time;
		data = 'title='+reminderObj.title+'&description='+reminderObj.description+'&time='+reminderObj.time;
		console.log(data);
		$http({
			url:URL+'/reminder',
			method:'POST',
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded'
			},
			data:data
		}).then(resp => {
			let res = (resp.data),
				message = res['message'],
				status = res['status'],
				messageObj = {
					message: '',
					sender: '',
					time: '',
					show: false,
					length: null
				};
			setTimeout(() => {
				$scope.scrollDown();
			}, 100);
			if ((status === 'success' || status) && message === 'Reminder has been set !') {
				messageObj.sender = 'jarvis-bot';
				messageObj.time = String(new Date().getHours() + ':' + new Date().getMinutes());
				messageObj.length = message.length;
				messageObj.message = message;
				$scope.messageStack.push(messageObj);
			} else {
				console.error('[JARVIS] error fetching from service.');
			}
		}).catch(e => {
			throw e;
		});
		$scope.formData.remTitle = '';
		$scope.formData.remDescription = '';
	};

	$scope.scrollDown = function() {
		var elem = document.getElementById('stackArea-parent');
		elem.scrollTop = elem.scrollHeight;
	};

	$scope.initStack = function() {
		$scope.message = '';
		// $scope.toggleStartStop(0);
	};

	$scope.toggleStartStop = function () {
		recognition.continuous = true;

		recognition.onresult = function (event) {
			var i, n, submessage;
			// var m, text;
			var mess = document.getElementById('message-input');
			mess.value = '';
			// text = '';
			// if (check === 0) {
			// 	for (i = 0; i < event.results.length; i++) {
			// 		if (event.results[i].isFinal) {
			// 			text += event.results[i][0].transcript;
			// 			console.log(text);
			// 			if (text.includes('start Jarvis')) {
			// 				m = text.lastIndexOf('start Jarvis');
			// 				submessage = text.substring(m+12);
			// 				mess.value = submessage;
			// 				$scope.message = submessage;
			// 				if (text.endsWith('send')) {
			// 					mess.value = text;
			// 					n = mess.value.lastIndexOf('send');
			// 					submessage =  mess.value.substring(m+12,n);
			// 					$scope.message = submessage;
			// 					$scope.addMessagesToStack();
			// 				}
			// 			}
			// 		} else {
			// 			text += event.results[i][0].transcript;
			// 			if (mess.value.includes('start jarvis')) {
			// 				mess.value += event.results[i][0].transcript;
			// 				n = mess.value.lastIndexOf('start jarvis');
			// 				submessage = mess.value.substring(n+12);
			// 				$scope.message = submessage;
			// 			}
			// 		}
			// 	}
			// } else if (check === 1) {
			// if (check === 0) {
			for (i = 0; i < event.results.length; i++) {
				if (event.results[i].isFinal) {
					mess.value += event.results[i][0].transcript;
					if (mess.value.endsWith('send')) {
						n = mess.value.lastIndexOf('send');
						submessage =  mess.value.substring(0,n);
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
			// }
			// }
		};

		if (recognizing) {
			recognition.stop();
			$scope.showLoaderListening = false;
			recognizing = false;
		} else {
			recognition.start();
			$scope.showLoaderListening = true;
			recognizing = true;
			$scope.message = '';
		}
	};
});

app.controller('sidebarController', function($scope) {

	console.warn('sidebar controller');
	$scope.initSidebar = function() {
		$scope.showHelp = false;
	};
	$scope.toggleHelp = function() {
		$scope.showHelp = !$scope.showHelp;
	};

});


$(document).ready(() => {
	$('#webQueries').hide();
	$('#videosYoutube').hide();
	$('#imagesBody').hide();
	$('#weatherBody').hide();
	$('#meaningBody').hide();
	$('#medicineBody').hide();
	$('#symptomsBody').hide();

	$('#videoHead').click(() => {
		$('#videosYoutube').show(1000);
		$('#webQueries').hide(500);
		$('#imagesBody').hide(500);
		$('#weatherBody').hide(500);
		$('#meaningBody').hide(500);
		$('#medicineBody').hide(500);
		$('#symptomsBody').hide(500);
	});
	$('#webHead').click(() => {
		$('#webQueries').show(1000);
		$('#videosYoutube').hide(500);
		$('#imagesBody').hide(500);
		$('#weatherBody').hide(500);
		$('#meaningBody').hide(500);
		$('#medicineBody').hide(500);
		$('#symptomsBody').hide(500);
	});
	$('#imageHead').click(() => {
		$('#webQueries').hide(500);
		$('#videosYoutube').hide(500);
		$('#imagesBody').show(1000);
		$('#weatherBody').hide(500);
		$('#meaningBody').hide(500);
		$('#medicineBody').hide(500);
		$('#symptomsBody').hide(500);
	});
	$('#weatherHead').click(() => {
		$('#webQueries').hide(500);
		$('#videosYoutube').hide(500);
		$('#imagesBody').hide(500);
		$('#weatherBody').show(1000);
		$('#meaningBody').hide(500);
		$('#medicineBody').hide(500);
		$('#symptomsBody').hide(500);
	});
	$('#meaningHead').click(() => {
		$('#webQueries').hide(500);
		$('#videosYoutube').hide(500);
		$('#imagesBody').hide(500);
		$('#weatherBody').hide(500);
		$('#meaningBody').show(1000);
		$('#medicineBody').hide(500);
		$('#symptomsBody').hide(500);
	});
	$('#medicineHead').click(() => {
		$('#webQueries').hide(500);
		$('#videosYoutube').hide(500);
		$('#imagesBody').hide(500);
		$('#weatherBody').hide(500);
		$('#meaningBody').hide(500);
		$('#medicineBody').show(1000);
		$('#symptomsBody').hide(500);
	});
	$('#symptomsHead').click(() => {
		$('#webQueries').hide(500);
		$('#videosYoutube').hide(500);
		$('#imagesBody').hide(500);
		$('#weatherBody').hide(500);
		$('#meaningBody').hide(500);
		$('#medicineBody').hide(500);
		$('#symptomsBody').show(1000);
	});
});
