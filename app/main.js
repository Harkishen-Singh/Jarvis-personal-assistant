const electron = require('electron'),
	path = require('path'),
	BrowserWindow = electron.BrowserWindow,
	App = electron.app,
	Menu = electron.Menu;

let mainWindow = null;

function MainWindow() {

	mainWindow = new BrowserWindow({
		width          : 400,
		height         : 560,
		backgroundColor: '#fff',
		// eslint-disable-next-line no-undef
		icon      		   : __dirname + '/assets/images/icon-jarvis1.png',
		resizable      : true,
		autoHideMenuBar: true,
	});
	Menu.setApplicationMenu(null);
	// eslint-disable-next-line no-undef
	// mainWindow.setIcon(path.join(__dirname, '/assets/images/icon-jarvis.ico'));
	// eslint-disable-next-line no-undef
	mainWindow.loadURL('file://' + __dirname + '/templates/index.html').then(() => {

		// eslint-disable-next-line no-console
		console.warn('main configuration window found. starting the application');

	});
	mainWindow.webContents.openDevTools();
	mainWindow.on('closed', () => {

		mainWindow = null;

	});

}

App.on('ready', MainWindow);