import Visitor from '../models/visitor.js';
import ReturnMessage from '../models/ReturnMessage.js';
import DB from '../modules/db.js';


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

        rMSG = {
            type:'',
            data:{
                INFO:''
                }
            };


async testfunc(_a){
    console.log("--> a : ", _a);
    this.rMSG.type = "Success";
    this.rMSG.data.INFO = "Visitor created"
    return this.rMSG;
}

    async newVisitor(_message) {


        this.rMSG.type = "Success";
        this.rMSG.data.INFO = "Visitor created";
        return _message;
        
        if(_message.Surname != null && _message.Givenname != null && _message.Birthd != null) {

            const  x = await   this.db.insertVisitor(_message);
            console.log("Status: ",x.affectedRows);
            
            if (x.affectedRows != 1) {
               //rMSG = await new ReturnMessage(1,"Visitor object could not be created");
                this.rMSG.type = "Error";
                this.rMSG.data.INFO = "Visitor not created";

            } else {
               //rMSG =  await new ReturnMessage(0,"Visitor Created"); 

               this.rMSG.type = "Success";
               this.rMSG.data.INFO = "Visitor created";

            }
            
        };
        console.log("--------->>>> ", this.rMSG);
        return this.rMSG;
        };            
}

export default VisitorController;