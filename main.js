/*jshint esversion: 6 */
const {
  app,
  BrowserWindow
} = require('electron');
const os = require('os');
const exec = require('child_process').execFile;
const path = require('path');

function createWindow() {
  let win = new BrowserWindow({
    width: 1280,
    height: 768,
    webPreferences: {
      nodeIntegration: true
    },
    title: "RF Analyzer",
    icon: path.join(__dirname, 'assets/icon.png'),
  });
  win.setMenuBarVisibility(false);
  win.loadFile('assets/editor.html');
}
app.on('ready', () => {
  if (os.platform == 'win32') {
    exec("ProjetoCompilador.exe");
  } else if (os.platform == "linux") {
    exec("./ProjetoCompilador");
  }
  createWindow();
});