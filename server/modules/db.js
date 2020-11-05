import mysql from 'mysql';

class DB {

 connection = {
    host: '127.0.0.1',
    user: 'mydb',
    password: 'mydb',
    database: 'mydb'
 };

 con = new mysql.createConnection(this.connection);

 constructor() {

    
 }



 getVisitors(data, callback) {
 this.con.query('select * from visitors;', function(err, results) {
     if(err){throw err};

     console.log('Results in getVisitors: ', results[0]);
     return callback(results);
 })    
 };


/* insertVisitor(data,callback){
    
    console.log('insertVisitor new Visitor: ', data.Surname);

    var sql = 'INSERT INTO `mydb`.`visitors` (`surname`, `givenname`, `birthd`) VALUES ("' + data.Surname + '", "' +data.Givenname + '", "' + data.Birthd + '");'
    console.log("sql query: ", sql);

    this.con.query(sql, function(err, results) {
        if(err) return callback(err)
        return callback(results);
    })    
} */

insertVisitor(data){
    
    console.log('insertVisitor new Visitor: ', data.Surname);

    var sql = 'INSERT INTO `mydb`.`visitors` (`surname`, `givenname`, `birthd`) VALUES ("' + data.Surname + '", "' +data.Givenname + '", "' + data.Birthd + '");'
    console.log("sql query: ", sql);

    this.con.query(sql, function(err) {
        
        return results;
    })    
}


 // ---------------


  query() {
    console.log('connecting...1')
    
    var pool = mysql.createPool({
        connectionLimit: 100,
        host: '127.0.0.1',
        user: 'mydb',
        password: 'mydb',
        database: 'mydb'
        
    });

    var sql = 'select * from visitors;';

    
    return new Promise(function(resolve, reject) {

        pool.getConnection(function(err, connection) {
            // do whatever you want with your connection here
        
            return  connection.query(sql, function (err, result, fields) {
              if (err) reject(err);
               
              else 
              console.log("sql-result: ", result)
                
                  connection.release();
              //res.json(result);
              resolve(result);
              
            })       
          });

    })

    /* 
    this.con.connect((err) => {
        console.log('connecting...2 : ', err)
        if(err) {
            console.log('DB connection error: ', err);            
            
        } else {
            console.log('DB connection ok');
        }
    })

    this.con.query('select * from visitors1;', (err, rows) =>
    {
        if(err) {
            console.log('query-err: ', err);     
            return err;       
        }
        console.log('res: ', rows);
        res =  rows;
    });
     */
    
    

    
}


}


export default DB;
      
