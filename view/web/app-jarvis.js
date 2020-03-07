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

app.filter('unsafe', function($sce) {
	return function(val) {
		return $sce.trustAsHtml(val);
	};
});

app.controller('video-view-controller', [ '$scope', '$sce', function ($scope, $sce) {

	let length = $scope.videoDetails.length;
	$scope.url = {};
	for (let i = 0; i < length; i ++ ) {

		let urlData = $scope.videoDetails[ i ].link.replace('watch?v=', 'embed/');
		$scope.url[ i ] = $sce.trustAsResourceUrl(urlData);

	}

}, ]);

app.controller('MainController', function($scope,$location,$rootScope,$http) {

	// eslint-disable-next-line no-undef
	var recognition = new webkitSpeechRecognition();
	var recognizing;

	var reminders = [];

	$scope.controlMainBanner = () => {
		$scope.mainBanner = true;
		setTimeout(() =>{
			$scope.mainBanner = false;
		}, 500);
	};
	$scope.messageStack = [];
	$scope.showLoaderListening = false;

	var input = document.getElementById('message-input');
	input.addEventListener('keyup', (event) => {
		if (event.keyCode === 13) {
			event.preventDefault();
			document.getElementById('message-bar-send').click();
		}
	});

	$scope.getResponse = (messageObj) => {
		const data = `username=${USER}&message=${messageObj.message}`;
		const index = $scope.messageStack.findIndex(obj => messageObj.fullDate === obj.fullDate && messageObj.text === obj.text);
		$http({
			url:`${URL}/message`,
			method:'POST',
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded'
			},
			data: data
		}).then(resp => {
			const res = (resp.data),
				message = res['message'],
				status = res['status'],
				result = res['result'],
				show = res['show'];
			let resSuccess = false;
			const resMessageObj = {
				message: message,
				sender: 'jarvis-bot',
				time: String(new Date().getHours() + ':' + new Date().getMinutes()),
				result: result,
				show: false,
				length: message.length
			};
			const successMessageList = [
				'here are the top search results',
				'here are the top search videos',
				'here are the searched images',
				'Enter Reminder details : ',
				'Enter Mail Details : ',
				'Information about the medicine : ',
				'Help on the given symptoms : ',
				'here are the current weather conditions'
			];
			if ((status === 'success' || status) && (successMessageList.includes(message) || !show)) {
				resSuccess = true;
			} else if (show) {
				resMessageObj.show = show;
				resSuccess = true;
			} else {
				console.error('[JARVIS] error fetching from service.');
				messageObj.hasError = true;
			}

			if (resSuccess) {
				setTimeout(() => {
					$scope.scrollDown();
				}, 100);
				if (index !== -1) {
					$scope.messageStack.splice(index+1, 0, resMessageObj);
				}
				else {
					$scope.messageStack.push(resMessageObj);
				}
			}
			messageObj.isLoading = false;
		}).catch(e => {
			messageObj.isLoading = false;
			messageObj.hasError = true;
			throw e;
		});
	};

	$scope.addMessagesToStack = () => {
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
					length: null,
					isLoading: true,
					fullDate: date
				};

			messageObj.message = message;
			messageObj.length = message.length;
			messageObj.time = String(hrs + ':' + mins);
			messageObj.sender = 'you';

			$scope.messageStack.push(messageObj);
			$scope.showLoading = true;
			setTimeout(() => {
				$scope.scrollDown();
			}, 100);

			$scope.getResponse(messageObj);
			$scope.message = '';
		}
	};

	$scope.retryMessage = (message) => {
		message.hasError = false;
		message.isLoading = true;
		$scope.getResponse(message);
	};

	$scope.formData = {};
	$scope.setReminder = () => {
		$scope.messageStack.pop();
		let reminder_title = $scope.formData.remTitle,
			reminder_description = $scope.formData.remDescription,
			reminder_time = $scope.formData.remTime,
			reminderObj = {
				title: '',
				description: '',
				time: '',
				cook: ''
			},
			data = null;
		reminderObj.title = reminder_title;
		reminderObj.description = reminder_description;
		reminderObj.time = reminder_time;
		document.cookie = reminderObj.title+'='+reminderObj.description+'; expires='+reminderObj.time.toUTCString();+'; path=/';
		reminderObj.cook = document.cookie;
		
		data = 'title='+reminderObj.title+'&description='+reminderObj.description+'&time='+reminderObj.time+'&cookie='+reminderObj.cook;
		
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
			console.log('res: ', res);
			console.log('message', message);
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

	$scope.formData = {};
	$scope.sendMail = () => {
		$scope.messageStack.pop();
		let mail_sender = $scope.formData.Sender,
			mail_to = $scope.formData.To,
			mail_cc = $scope.formData.CC,
			mail_bcc = $scope.formData.BCC,
			mail_subject = $scope.formData.Subject,
			mail_body = $scope.formData.Body,
			
			mailObj = {
				sender: '',
				to: '',
				cc: '',
				bcc: '',
				subject: '',
				body: ''
			},
			data = null;
		
		mailObj.sender = mail_sender;
		mailObj.to = mail_to;
		mailObj.cc = mail_cc;
		mailObj.bcc = mail_bcc;
		mailObj.subject = mail_subject;
		mailObj.body = mail_body;

		console.log(mailObj);
		
		data = 'sender='+mailObj.sender+'&to='+mailObj.to+'&subject='+mailObj.subject+'&body='+mailObj.body+
			'&cc='+mailObj.cc+'&body='+mailObj.bcc;
		
		console.log(data);
		
		$http({
			url:URL+'/email',
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
			if ((status === 'success' || status) && message === 'Mail sent Successfully') {
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
		$scope.formData.To = '';
		$scope.formData.Subject = '';
		$scope.formData.Body = '';
	};

	function reminderNotif() {
		var x = document.cookie;
		var allCookie = x.split(';');
		if (allCookie.length > reminders.length && allCookie !== '') {
			for (var i = reminders.length; i <allCookie.length; i++) {
				var oneCookie = allCookie[i].split('=');
				var rem = {
					title: '',
					desc: ''
				};
				rem.title = oneCookie[0];
				rem.desc = oneCookie[1];
				reminders.push(rem);
			}
			console.log('created');
			console.log(reminders);
		}
		if ((allCookie === '' && allCookie.length-1 < reminders.length) || (allCookie !== '' && allCookie.length < reminders.length)) {
			for (i = 0; i <allCookie.length; i++) {
				oneCookie = allCookie[i].split('=');
				var title = oneCookie[0];
				if (reminders[i].title !== title) {
					alert('\tReminder! \n\n\t'+reminders[i].title+'\n\n'+reminders[i].desc);
					reminders.splice(i,i+1);
				}
			}
			console.log('deleted');
			console.log(reminders);
		}		
	}
	setInterval(reminderNotif,10000);

	$scope.scrollDown = () => {
		var elem = document.getElementById('stackArea-parent');
		elem.scrollTop = elem.scrollHeight;
	};

	$scope.initStack = () => {
		$scope.message = '';
	};

	$scope.toggleStartStop = () => {
		recognition.continuous = true;

		recognition.onresult = (event) => {
			var i, n, submessage;
			var mess = document.getElementById('message-input');
			mess.value = '';
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
	$scope.initSidebar = () => {
		$scope.showHelp = false;
	};
	$scope.toggleHelp = () => {
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
