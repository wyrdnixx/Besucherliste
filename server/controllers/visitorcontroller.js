import Visitor from '../models/visitor.js';
import ReturnMessage from '../models/ReturnMessage.js';
import DB from '../modules/db.js';
//import { result } from 'lodash';

class VisitorController {


    db = new DB();
    
    constructor() {
        
    };


 /*     errmessage = {
        type:"Error",
        data:{
            INFO:"invalid visitor JSON"}

            successmessage = {
        type:"Success",
        data:{
            INFO:"Visitor created"}
        }; */

       

    newVisitor(_message) {
        

        if(_message.Surname != null && _message.Givenname != null && _message.Birthd != null) {

            var v = new Visitor(_message.Surname, _message.Givenname, _message.Birthd );

            
                        
            
/* 
            this.db.insertVisitor(_message, function(result){
                 console.log('newVisitor Result: ', result.affectedRows);
                 return  new ReturnMessage(0,"Visitor Created: ", result.affectedRows);
             });
 */


            var x = this.db.insertVisitor(_message);
            
            console.log(x);

        };


        //console.log("newVisitor Result:" , res);
        
             /* 
            console.log(typeof v);
            if( v === typeof Visitor) {

                
               return  new ReturnMessage(0,"Visitor Created");

            } else return new ReturnMessage(1,"Visitor object could not be created");
            
        } else return new ReturnMessage(1,"Invaild data supplied"); */


        };            
}

export default VisitorController;