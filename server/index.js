import eventBus from './modules/eventbus.js';
import Server from './modules/server.js';

import eventbus from './modules/eventbus.js';
import Interface from './modules/interface.js';

const server = new Server(8080);

const inter = new Interface();

/* eventBus.on(
    'message',
    (clientId, message) => server.sendMessage(clientId, message),
    
); */

 
/* 
 eventBus.on(
    'NewVisitorCreated',
    (clientId, message) => server.sendMessage(clientId, message),
    
); 
 */



eventBus.on(
    'error',
    (error) => console.log(error),
);

server.on(
    'message',
    (clientId, message) => eventBus.consumeMessage(clientId, message),    
);



server.on(
    'error',
    (error) => console.log(error),
);