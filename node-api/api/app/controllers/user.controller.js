const User = require('../models/User.model.js');

// Criar e salvar novo User
exports.create = (req, res) => {
    
    // Validar requisição 
    if(!req.body.nome) {
        return res.status(400).send({
            message: "Infelizmente ocorreu um erro... " + req.body.data
        });
    }

    // Criando um usuario
    const user = new User({
        nome: req.body.nome, 
        email: req.body.email || "email@email.com"
    });

    // salvando novo usuario no banco
    user.save()
    .then(data => {
        res.send(data);
    }).catch(err => {
        res.status(500).send({
            message: err.message || "Ocorreu um erro enquanto estava sendo salvo..."
        });
    });
};

// Listar todos os usuarios
exports.findAll = (req, res) => {
    User.find()
    .then(users => {
        res.send(users);
    }).catch(err => {
        res.status(500).send({
            message: err.message || "Infelizmente ocorreu um erro. Por favor tente mais tarde..."
        });
    });
};

// procurar pelo Id
exports.findOne = (req, res) => {
    User.findById(req.params.id)
    .then(user => {
        if(!user) {
            return res.status(404).send({
                message: "Erro ao recuperar o id: " + req.params.id
            });            
        }
        res.send(user);
    }).catch(err => {
        if(err.kind === 'ObjectId') {
            return res.status(404).send({
                message: "Não existe usuario com id: " + req.params.id
            });                
        }
        return res.status(500).send({
            message: "Erro ao recuperar o id: " + req.params.id
        });
    });
};

// atualizar pelo Id
exports.update = (req, res) => {
    // Validate Request
    if(!req.body.nome) {
        return res.status(400).send({
            message: "Necessário informar o campo nome..."
        });
    }

    // procura o usuario e atualiza com os novos dados
    User.findByIdAndUpdate(req.params.id, {
        nome: req.body.nome,
        email: req.body.email || "email@email.com"
    }, {new: true})
    .then(user => {
        if(!user) {
            return res.status(404).send({
                message: "usuario nao encontrado com id: " + req.params.id
            });
        }
        res.send(user);
    }).catch(err => {
        if(err.kind === 'ObjectId') {
            return res.status(404).send({
                message: "Usuario nao encontrado com id: " + req.params.id
            });                
        }
        return res.status(500).send({
            message: "Erro ao atualizar o usuario com id: " + req.params.id
        });
    });
};

// deletar pelo id
exports.delete = (req, res) => {

    User.findByIdAndRemove(req.params.id)
    .then(user => {
        if(!user) {
            return res.status(404).send({
                message: "Usuario não encontrado com id: " + req.params.id
            });
        }
        res.send({message: "usuario deletado com sucesso!"});
    }).catch(err => {
        if(err.kind === 'ObjectId' || err.name === 'NotFound') {
            return res.status(404).send({
                message: "Usuario não encontrado com id: " + req.params.id
            });                
        }
        return res.status(500).send({
            message: "Não possivel excluir usuario com id: " + req.params.id
        });
    });
};