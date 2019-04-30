/*
* Aims to develop a modular way of connecting jarvis service and in turn serve any type of request
* */

class Connect {
	constructor(username = 'default') {
		this.username = username;
		this.url = 'http://0.0.0.0:3000'; // default listener for jarvis service
		this.query = '';
		this.method = 'POST';
		this.urlString = '';
	}

	setQuery(query) {
		this.query = query;
		this.urlString = this.url + '?username=' + this.username + '&message=' + this.query;
		return true;
	}

	send() {
		return new Promise((resolve, reject) => {
			const xhttp = new XMLHttpRequest();
			xhttp.onreadystatechange = function () {
				if (this.readyState === 4 && this.status === 200) {
					// eslint-disable-next-line no-console
					console.warn('received from jarvis service');
					// eslint-disable-next-line no-console
					console.warn(this.responseText);
					resolve(this.responseText);
				}
			};
			xhttp.open(this.method, this.urlString);
			xhttp.send();
		});
	}
}

module.exports = Connect;