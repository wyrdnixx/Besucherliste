import test from 'ava';
import Server from '../modules/server.js';
import WebSocket from 'ws';

let server;
let testClientId;
let webSocket;

const testMessage = JSON.stringify({
    type: 'test',
    data: {
        test: 'test',
    },
});

const testMessageNewVisitor = JSON.stringify({
    type: 'NewVisitor',
    data: {
        Name: 'Hans',
    },
});



test.before.cb((t) => {
    server = new Server(8080, 10);
    server.on(
        'listening',
        () => {
            webSocket = new WebSocket('ws://127.0.0.1:8080');
            webSocket.on('open', t.end);
        },
    );
});

test.after(() => {
    server.close();
    webSocket.close();
});

test.cb('receives messages', (t) => {
    server.on(
        'message',
        (clientId, message) => {
            testClientId = clientId;
            t.is(message, testMessage);
            t.end();
        },
    );

    webSocket.send(testMessage);
});

test.cb('sends messages', (t) => {
    webSocket.on(
        'message',
        (message) => {
            //console.warn(testMessage);
            t.is(message, testMessage);
            t.end();
        },
    );

    server.sendMessage(testClientId, testMessage);
});

test.cb('receives pongs', (t) => {
    server.on(
        'pong',
        (clientId) => {
            t.is(clientId, testClientId);
            t.end();
        },
    );

    webSocket.pong();
});

test.cb('sends pings', (t) => {
    webSocket.on(
        'ping',
        () => t.end(),
    );
});


test('adds active clients', (t) => {
    t.is(Object.keys(server.clients).length, 1);
});







// Last test

test.cb('removes inactive clients', (t) => {
    server.on(
        'close',
        (clientId) => {
            t.is(clientId, testClientId);
            t.is(Object.keys(server.clients).length, 0);
            t.end();
        },
    );

    webSocket.terminate();
});

