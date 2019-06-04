/*jshint esversion: 6 */
const {
  app,
  BrowserWindow
} = require('electron');
const os = require('os');
const exec = require('child_process').execFile;
function createWindow() {
  // Create the browser window.
  let win = new BrowserWindow({
    width: 1280,
    height: 768,
    webPreferences: {
      nodeIntegration: true
    },
    title: "Rufus Editor",
  });
  win.setMenuBarVisibility(false);
  win.loadFile('assets/editor.html');
}
app.on('ready', () => {
  if (os.platform == 'win32') {
    exec("ProjetoCompilador.exe");
  } else if (os.platform == "linux") {
    //exec("./ProjetoCompilador");
  }
  createWindow();
});