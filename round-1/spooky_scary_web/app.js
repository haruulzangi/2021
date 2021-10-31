const express = require('express')
const port = 9010
const host = '0.0.0.0'
const fs = require('fs')
const indexhtml = fs.readFileSync('./index.html', 'utf-8')
const flag = fs.readFileSync('./flag', 'utf-8')
const _ = require('lodash')

app = express()
app.use(express.static(__dirname+'/images'));
app.use(express.json())

app.use(express.urlencoded({
    extended: true
  }));
const notsus = {}
app.get('/', (req, res) => {
    res.sendFile(__dirname+'/index.html');
})

app.post('/skoopy', (req, res) => {
    const user = {
        userip: req.connection.remoteAddress,
    };
    parsedBody = _.defaultsDeep({
        admin: false, 
        cоntent: '', 
    }, req.body)
    if(user.admin) {
        res.send(flag)
        process.exit(0)
    }
    else {
        res.send('Aay chi yah geed bgan admin bish ym baaaij!'+user.userip);
    }
    
})


app.listen(port, host)
console.log('listening on... '+port+' '+host)
