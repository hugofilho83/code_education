const express = require('express')
const app = express()
const port = 3000
const config = {
    host: 'db_mysql',
    user: 'root',
    password:'root',
    database: 'nodedb'
};

const mysql = require('mysql')
const connection = mysql.createConnection(config)

const sql = `INSERT INTO people (name) VALUES ('hugo tavares')`
connection.query(sql)

app.get('/', (req, res)=>{

    connection.query( 'SELECT * FROM people', function(error, result, fields) {

        if(error)
            res.json(error);
        else
            res.json(result);

        connection.end();

    });
    
})

app.listen(port,()=>{
    console.log('Rodando na porta ' + port)
})