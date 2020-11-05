import test from 'ava';
import VisitorController from '../controllers/visitorcontroller.js'


const visitorController = new VisitorController();


const successmessage = {
    type:"Success",
    data:{
        INFO:"Visitor created"}
    };

const invalidDataMessage = {
    type:"Error",
    data:{
        INFO:"Invaild data supplied"}
    };

test ('newVisitor err - incomplete JSON Data', t => {

    const testVisitor = {        
            name: 'Hans',        
    }
    
    //{"type":"NewVisitor","data":{"Surname":"Hans", "Givenname":"Mueller", "Birthd":"01.01.1990"}}    

    t.is(  JSON.stringify(invalidDataMessage),  JSON.stringify(visitorController.newVisitor(testVisitor)));
})

test ('newVisitor Success', t => {

    const testVisitor = {
        type:"NewVisitor",
        data:{
            Surname:"Hans",
            Givenname:"Mueller",
            Birthd:"01.01.1990"}
        };    

    t.is(  JSON.stringify(successmessage),  JSON.stringify(visitorController.newVisitor(testVisitor)));
})