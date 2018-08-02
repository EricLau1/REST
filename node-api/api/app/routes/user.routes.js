module.exports = (app) => {
    const users = require('../controllers/user.controller.js');

    //  Criar novo usuario
    app.post('/users', users.create);

    // Listar usuarios
    app.get('/users', users.findAll);

    // Buscar usuario pelo Id
    app.get('/users/:id', users.findOne);

    // Atualizar usuario pelo Id
    app.put('/users/:id', users.update);

    // Deletar usuario pelo Id
    app.delete('/users/:id', users.delete);
}