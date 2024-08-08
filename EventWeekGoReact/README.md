# Projeto AMA - Ask me anything

- A aplicação será composto de: 
  - Um servidor web escrito em GO
  - Front-end será feito em React


## Funcionamento da aplicação

- Basicamente existirá várias salas, aonde cada uma tera um ID
- Teremos a posibilidade de compartilhar a sala, para mais usuários que quiserem entrar
- A idéia principal é que as pessoas possam fazem perguntas na sala em tempo real
- Cada pergunta feita, pode ser curtida caso seja de interesse da pessoa
- O diferencial é no uso de WebSockets, aonde é feita uma requisição para o servidor e o mesmo ficará retornando respostas para a aplicação, atualizando em tempo real, quando chegar novas perguntas ou novas interações com as perguntas já efetuadas.