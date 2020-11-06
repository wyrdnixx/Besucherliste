import eventBus from './eventbus.js';
import VisitorController from '../controllers/visitorcontroller.js'

class Interface {

   visitorController = new VisitorController();
   constructor() {

    

       eventBus.on('NewVisitor', async (clientId, message) =>{
           
        console.log("--> Create new Visitor: " + clientId, message);
        let resr = await  this.visitorController.newVisitor(message);
               
        
       })
   }

}

export default Interface;