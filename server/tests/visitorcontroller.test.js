import test from 'ava';
import VisitorController from '../controllers/visitorcontroller.js'



const vc = new VisitorController();    


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

test.skip ('newVisitor err - incomplete JSON Data', t => {

    const testVisitor = {        
            name: 'Hans',        
    }
    
    //{"type":"NewVisitor","data":{"Surname":"Hans", "Givenname":"Mueller", "Birthd":"01.01.1990"}}    

    t.is(  JSON.stringify(invalidDataMessage),  JSON.stringify(visitorController.newVisitor(testVisitor)));
})

test ('newVisitor Success', async t => {


    const testVisitor = {
        type:"NewVisitor",
        data:{
            Surname:"Hans",
            Givenname:"Mueller",
            Birthd:"01.01.1990"}
        };    
        
        //const value = await vc.newVisitor(testVisitor);
        //console.log("---> TEST VALUE: ", value);
        
        //t.is(  JSON.stringify(successmessage), vc.newVisitor(testVisitor));
        t.is(  JSON.stringify(successmessage), vc.newVisitor("Hallo"));
        
})

test.skip ('async test' ,async t => {
    
    const a = 'xallo';

    //const val = await vc.testfunc(a);    
    t.is (successmessage, await vc.testfunc(a));
}) 

