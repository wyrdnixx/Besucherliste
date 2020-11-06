import mysql from 'mysql-await';

class DB {

 dbcondata = {
    host: '127.0.0.1',
    user: 'mydb',
    password: 'mydb',
    database: 'mydb'
 };

 

 constructor() {


 }



 getVisitors(data, callback) {
 this.con.query('select * from visitors;', function(err, results) {
     if(err){throw err};

     console.log('Results in getVisitors: ', results[0]);
     return callback(results);
 })    
 };


 async insertVisitor(data){
 
    
        /** Create connection pool using loaded config */
        const connection = mysql.createConnection(this.dbcondata);
       
        connection.on(`error`, (err) => {
          console.error(`Connection error ${err.code}`);
        });



        var sql = 'INSERT INTO `mydb`.`visitors` (`surname`, `givenname`, `birthd`) VALUES ("' + data.Surname + '", "' +data.Givenname + '", "' + data.Birthd + '");'

        let result =  await connection.awaitQuery(sql);
        console.log("sql-Result: ", result);
    
    //con = new mysql.createConnection(this.connection);
    
    return result;
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
      
