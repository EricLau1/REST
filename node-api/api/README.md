Node JS Api Rest

comandos

iniciar o projeto:

*npm init -y* 

pacotes necessários para instalar:

*npm install express body-parser mongoose --save*

necessario ter o MongoDB instalado 

criar o arquivo 'server.js'

criar a pasta "config" e o arquivo 'database.config.js' dentro dela.

criar a pasta "app" e dentro dela a pasta "models". Crie o arquivo 'User.model.js' dentro da pasta "models".

criar a pasta "routes" dentro da pasta "app" e o arquivo 'user.routes.js' dentro da pasta "routes".

criar a pasta "controllers" dentro da pasta "app" e o arquivo 'user.controller.js' dentro dela.

Após configurar todos os documentos.

é necessário mudar o cabeçalho do postman para application/json para que as requisições funcionem.

Quando for passar algum parametro por Id informar o _Id gerado pelo mongodb