const electron = require('electron'),
	BrowserWindow = electron.BrowserWindow,
	App = electron.app,
	Tray = electron.Tray,
	Menu = electron.Menu;

let mainWindow = null,
	tray = null;

function MainWindow() {

	// initialise system tray
	// eslint-disable-next-line no-undef
	tray = new Tray(__dirname + '/assets/images/icon-jarvis1.png');
	tray.setToolTip('Jarvis - The personal assistant');
	const trayMenu = Menu.buildFromTemplate([
		{
			label: 'Exit',
			click: function () {

				App.quit();

			}
		},
		{
			label: 'Show App',
			click: function () {

				mainWindow.show();

			}
		},
		{
			label: 'Minimize',
			click: function () {

				mainWindow.minimize();

			}
		}
	]);
	tray.setContextMenu(trayMenu);
	mainWindow = new BrowserWindow({
		width          : 400,
		height         : 560,
		backgroundColor: '#fff',
		// eslint-disable-next-line no-undef
		icon      		   : __dirname + '/assets/images/icon-jarvis1.png',
		resizable      : true,
		autoHideMenuBar: true
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
	mainWindow.on('closed', (event) => {

		event.preventDefault();
		mainWindow.hide();

	});

	mainWindow.on('minimize', (event) => {

		event.preventDefault();
		mainWindow.minimize();

	});

	mainWindow.on('close', (event) => {

		event.preventDefault();
		mainWindow.hide();

	});

}

App.on('ready', MainWindow);
App.on('quit', function () {

	App.quit();

});
