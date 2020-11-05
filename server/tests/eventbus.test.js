import test from 'ava';
import eventBus from '../modules/eventbus.js';
import Interface from '../modules/interface.js';

import uuid from 'uuid';

const testClientId = uuid.v4();


const inter = new Interface();

test.afterEach(() => {
    eventBus.removeAllListeners();
});

test.cb('rejects messages that are not JSON', (t) => {
    eventBus.on('error', (error) => {
        t.is(error.message, 'Message is not valid JSON');
        t.end();
    });

    eventBus.consumeMessage(
        testClientId,
        'test',
    );
});

test.cb('rejects messages that are not objects', (t) => {
    eventBus.on('error', (error) => {
        t.is(error.message, 'Message is not an object');
        t.end();
    });

    eventBus.consumeMessage(
        testClientId,
        '[]',
    );
});

test.cb('rejects messages that do not include a type', (t) => {
    eventBus.on('error', (error) => {
        t.is(error.message, 'Message does not include a type');
        t.end();
    });

    eventBus.consumeMessage(
        testClientId,
        '{}',
    );
});

test.cb('rejects messages that do not include data', (t) => {
    eventBus.on('error', (error) => {
        t.is(error.message, 'Message does not include data');
        t.end();
    });

    eventBus.consumeMessage(
        testClientId,
        '{ "type": "test" }',
    );
});

test.cb('emits events from messages', (t) => {
    const testEvent = {
        type: 'test',
        data: {
            test: 'test',
        },
    };

    eventBus.on('test', (clientId, data) => {
        t.is(clientId, testClientId);
        t.deepEqual(data, testEvent.data);
        t.end();
    });

    eventBus.consumeMessage(
        testClientId,
        JSON.stringify(testEvent),
    );
});

test.cb('rejects events that are not objects', (t) => {
    eventBus.on('error', (error) => {
        t.is(error.message, 'Event is not an object');
        t.end();
    });

    eventBus.produceMessage(
        testClientId,
        'test',
    );
});

test.cb('rejects events that do not include a type', (t) => {
    eventBus.on('error', (error) => {
        t.is(error.message, 'Event does not include a type');
        t.end();
    });

    eventBus.produceMessage(
        testClientId,
        {},
    );
});

test.cb('rejects events that do not include data', (t) => {
    eventBus.on('error', (error) => {
        t.is(error.message, 'Event does not include data');
        t.end();
    });

    eventBus.produceMessage(
        testClientId,
        { type: 'test' },
    );
});

test.cb('rejects events that do not valid json', (t) => {
    const testEvent = {
        type: 'test',
        data: {},
    };

    testEvent.data.event = testEvent;

    eventBus.on('error', (error) => {
        t.is(error.message, 'Event is not valid JSON');
        t.end();
    });

    eventBus.produceMessage(
        testClientId,
        testEvent,
    );
});

test.cb('emits messages from events', (t) => {
    const testEvent = {
        type: 'test',
        data: {
            test: 'test',
        },
    };


    eventBus.on('test', (clientId, data) => {
        t.is(clientId, testClientId);
        t.is(data, JSON.stringify(testEvent));
        t.end();
    });

    eventBus.produceMessage(
        testClientId,
        testEvent,
    );
});

test.cb('emits NewVisitor from events', (t) => {
    const testEvent = {
        type: 'NewVisitor',
        data: {
            test: 'Hans',
        },
    };

   /*  eventBus.on('NewVisitorCreated', (clientId, data) => {
        t.is(clientId, testClientId);
        t.is(data, JSON.stringify(testEvent)); */
        t.end();
/*     }); */

    eventBus.consumeMessage(
        testClientId,
        JSON.stringify(testEvent),
    );
});