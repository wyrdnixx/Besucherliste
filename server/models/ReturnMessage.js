class ReturnMessage {
    
    
        type= "";
        data = {
            INFO: ""
        }
    
    
  

/**
 * 
 * @param {return code 0=success; >0 failed} _code 
 * @param {info text} _info 
 */
    constructor(_code,_info) {

        
        
        
        if (_code === 0) {
            
            this.type = "Success";   
            
            
        } else {
            this.type = "Error";            
        }
        
        this.data.INFO = _info
        
        //this.data.INFO = _info;
    }
}

export default  ReturnMessage;