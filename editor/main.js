const { app, BrowserWindow } = require('electron');

function createWindow () {
  // Create the browser window.
  let win = new BrowserWindow({
    width: 1280,
    height: 768,
    webPreferences: {
      nodeIntegration: true
    },
    title:"Rufus Editor",
  });
  win.setMenuBarVisibility(false)
  win.loadFile('assets/editor.html');
}

app.on('ready', createWindow);