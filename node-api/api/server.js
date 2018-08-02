
const express = require('express');
const bodyParser = require('body-parser');

// create express app
const app = express();

// parse requests of content-type - application/x-www-form-urlencoded
app.use(bodyParser.urlencoded({ extended: true }))

// parse requests of content-type - application/json
app.use(bodyParser.json())

// Configuring the database
const dbConfig = require('./config/database.config.js');
const mongoose = require('mongoose');

mongoose.Promise = global.Promise;

// Conectando ao banco de dados
mongoose.connect(dbConfig.url, { useNewUrlParser: true })
.then(() => {
    console.log("Banco de dados conectado com sucesso!");    
}).catch(err => {
    console.log('Erro ao conectar no banco. Saido agora...');
    process.exit();
});

// definindo uma rota simples
app.get('/', (req, res) => {
    res.json({"message": "Olá mundo!"});
});

// chamando as rotas para usuarios
require('./app/routes/user.routes')(app);

// listen for requests
app.listen(3000, () => {
    console.log("Servidor iniciado na porta 3000");
});