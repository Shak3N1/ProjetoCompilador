/*jshint esversion: 6 */
const fs = require('fs');
const net = require('net');
const exec = require('child_process').execFile;
const os = require('os');
$(document).ready(function () {
    if (os.platform == 'win32') {
        exec("ProjetoCompilador.exe");
    } else if (os.platform == "linux") {
        exec("./ProjetoCompilador");
    }
    class Token {
        constructor(token, lexeme, line) {
            this.token = token;
            this.lexeme = lexeme;
            this.line = line;
        }
    }
    let client = new net.Socket();
    $("#open-file").click(function () {
        fs.readFile(document.getElementById("upload").files[0].path, (err, data) => {
            if (err) {
                alert(err);
                throw err;
            }
            $("#program-text").text(data);
        });
    });
    $("#save-file").click(function () {
        let data = $("#program-text").val();
        let url = document.getElementById("upload").files[0].path;
        fs.writeFile(url, data, (err) => {
            if (err) {
                alert(err);
                throw err;
            } else alert("Arquivo salvo!");
        });
    });
    $("#analyze-file").click(function () {
        let data = $("#program-text").val();
        client.write('getTokens ' + data);
    });

    client.connect(9261, '127.0.0.1', function () {
        console.log('Connected');
    });

    client.on('data', function (data) {
        console.log('Received: ' + data);
        let dataString = data.toString('utf8');
        let _data = dataString.split(" ");
        if (_data[0] == "sendToken") {
            let tokens = [];
            for (let i = 1; i < _data.length; i += 4) {
                let _token = new Token(_data[i], _data[i + 1], _data[i + 2]);
                tokens.push(_token);
            }
            $("#table tbody > tr").remove();
            tokens.forEach((token) => {
                $("#table tbody").append(
                    "<tr>" +
                    "<td>" + token.token + "</td>" +
                    "<td>" + token.lexeme + "</td>" +
                    "<td>" + token.line + "</td>" +
                    "</tr>"
                );
                $("#lex-errs").text('');
                if (token.token == "ERROR") {
                    $("#lex-errs").append(
                        token.token + " " +
                        " Lexema: " + token.lexeme +
                        " Linha: " + token.line
                    );
                }
            });
        }
    });
});
