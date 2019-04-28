const electron = require('electron'),
	BrowserWindow = electron.BrowserWindow,
	App = electron.app,
	Menu = electron.Menu;

let mainWindow;

function MainWindow() {
	mainWindow = new BrowserWindow({
		width: 400,
		height: 600,
		backgroundColor: '#fff',
		resizable: true,
		autoHideMenuBar: true
	});
	Menu.setApplicationMenu(null);
	// eslint-disable-next-line no-undef
	mainWindow.loadURL('file://'+__dirname+'/templates/index.html');
	mainWindow.webContents.openDevTools();
}

App.on('ready', MainWindow);