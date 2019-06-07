/*jshint esversion: 6 */
const fs = require('fs');
const net = require('net');
$(document).ready(function () {
    let lastData = "";
    class Token {
        constructor(token, lexeme, line) {
            this.token = token;
            this.lexeme = lexeme;
            this.line = line;
        }
    }
    let client = new net.Socket();
    client.connect(9261, '127.0.0.1', function () {
        console.log("connected");
    });
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
        $("#table tbody > tr").remove();
        $("#syn-errs").text('');
        $("#lex-errs").text('');
        let data = $("#program-text").val();
        if (lastData != data) {
            client.write("getTokens " + data);
            client.write("synAnalyze " + data);
            lastData = data;
        }
    });
    client.on('data', function (data) {
        let dataString = data.toString('utf8');
        let _data = dataString.split(" ");
        if (_data[0] == "sendToken") {
            let tokens = [];
            let synError = false;
            for (let i = 1; i < _data.length; i += 4) {
                if (_data[i - 1] == "synErr") {
                    let expected = "";
                    for (let x = i + 10; x < _data.length; x++) {
                        if (_data[x] != "$") {
                            if (x == _data.length - 1)
                                expected += " " + _data[x];
                            else
                                expected += " '" + _data[x] + "'";
                        }
                    }
                    _line = _data[i + 5].split("=");
                    $("#syn-errs").append(
                        "Token(s) esperado(s): " + expected +
                        "\nEncontrado: " + _data[i + 3].replace(",", "") +
                        "\nLinha " + _line[1].replace(",", ""));
                    synError = true;
                    break;
                }
                let _token = new Token(_data[i], _data[i + 1], _data[i + 2]);
                tokens.push(_token);
            }
            if (!synError) {
                $("#syn-errs").append(
                    "Nenhum erro encontrado"
                );
            }
            let lexError = false;
            tokens.forEach((token) => {
                $("#table tbody").append(
                    "<tr>" +
                    "<td>" + token.token + "</td>" +
                    "<td>" + token.lexeme + "</td>" +
                    "<td>" + token.line + "</td>" +
                    "</tr>"
                );

                if (token.token == "ERROR") {
                    lexError = true;
                    $("#lex-errs").append(
                        " Lexema ' " + token.lexeme +
                        " ' não reconhecido na linha " + token.line
                    );
                }
            });
            if (!lexError) {
                $("#lex-errs").append(
                    "Nenhum erro encontrado"
                );
            }
        }
    });
});
