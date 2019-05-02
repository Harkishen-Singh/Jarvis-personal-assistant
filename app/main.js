/* eslint-disable no-console */
const electron = require('electron'),
	BrowserWindow = electron.BrowserWindow,
	App = electron.app,
	Menu = electron.Menu;

let mainWindow;

function MainWindow() {
	mainWindow = new BrowserWindow({
		width: 400,
		height: 560,
		backgroundColor: '#fff',
		resizable: true,
		autoHideMenuBar: true
	});
	Menu.setApplicationMenu(null);
	// eslint-disable-next-line no-undef
	mainWindow.loadURL('file://'+__dirname+'/templates/index.html').then(success => {
		console.warn('main configuration window found. starting the application');
		console.warn(success);
	}).catch(e => {
		console.error(e);
	});
	mainWindow.webContents.openDevTools();
	mainWindow.on('closed', () => {
		mainWindow = null;
	});
}

App.on('ready', MainWindow);