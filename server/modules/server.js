import EventEmitter from 'events';
import forEach from 'lodash/forEach.js';
import uuid from 'uuid';
import WebSocket from 'ws';

class Server extends EventEmitter {
    clients = {};

    constructor(port = 8080, pingInterval = 10000) {
        super();

        this.webSocketServer = new WebSocket.Server({
            port: port,
        });

        this.webSocketServer.on(
            'listening',
            () => this.emit('listening'),
            console.log("Server is listening... ")
        );

        this.webSocketServer.on(
            'connection',
            (webSocket) => this.handleConnection(webSocket),
            
        );

        this.webSocketServer.on(
            'error',
            (error) => this.emit('error', error),
        );               

        this.pingInterval = setInterval(
            () => this.pingClients(),
            pingInterval,
        );        
    }

    close() {
        this.webSocketServer.close();

        clearInterval(this.pingInterval);
    }

    handleConnection(webSocket) {
        const clientId = uuid.v4();

        webSocket.alive = true;

        webSocket.on(
            'message',            
            (message) => this.receiveMessage(clientId, message),            
        );

        webSocket.on(
            'pong',
            () => {
                webSocket.alive = true;
                this.emit('pong', clientId);
            },
        );

        webSocket.on(
            'close',
            () => {
                delete this.clients[clientId];
                this.emit('close', clientId);
            },
        );

        this.clients[clientId] = webSocket;
    }

    receiveMessage(clientId, message) {
        console.log("server got: ", clientId, " MSG: ", message)
        /* console.log(this) */
        this.emit(
            'message',
            clientId,
            message,
        );
        
        //Jojo
        //this.sendMessage(clientId, message)

    }

    sendMessage(clientId, message) {
        //console.log("msg-test : " + message);
        console.log("server sendmessage: ", clientId, " MSG: ", message)
        this.clients[clientId].send(message);
    }

    pingClients() {
        forEach(
            this.clients,
            (webSocket, sessionId) => {
                if (!webSocket.alive) {
                    delete this.clients[sessionId];
                    return webSocket.terminate();
                }

                webSocket.alive = false;

                if (webSocket.readyState === WebSocket.OPEN) {
                    webSocket.ping();
                }
            },
        );
    }
}

export default Server;