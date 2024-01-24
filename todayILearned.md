___________________________________________________________________________________________________ 


### Fundamentos ReactJS - Aula 01

- Biblioteca para construção de interfaces altamente interativas.
    - Interfaces não são somente interfaces Web
    - Interface é tudo que o usuário interagi quando ele está utilizando sua aplicação

- Hoje o React pode ser utilizado para construção de interfaces na Web, no mobile, em dispositivos de tv, através do react TV e em realidade virtual também, utilizando o react VR.
- O react por si só, ele é uma biblioteca e não uma framework.


- #### SPA
    - Single Page Application -> conceito ou forma de trabalharmos com a parte visual de nossa aplicação, diferente das tradicionais.
        - O Browser, quando ele acessa uma rota, por exemplo /users, aonde será retornado a listagem de usuários
        - O Backend dessa vez, ele busca no banco de dados a listagem de usuários, porém, o back não tem mais dentro dele no servidor, todas as intruções para construção do HTML, JS e CSS da página, ou seja, o backend não fica mais responsável pela construção visual da tela, na verdade o back vai somente retornar quais são os dados dos usuários
        - Porém, essas informações não são retornadas no formato de HTML, como era no modelo tradicional de SSR, essas informações são retornadas no formato JSON
            - JSON -> Javascript Object Notation 
                - É uma estrutura de dados, um formato de disponibilizar dados, ou seja, trafegar dados entre duas fontes diferentes.. como uma origem e um destino.
                - O JSON é uma estrutura universal
                - Muitas linguagens e tecnologias, conseguem interpretar JSON
                    - é possível enviar dados de uma aplicação JAVA, para uma aplicação Python em JSON e as duas linguagens conseguem entender essa informação
                
        - Aplicação Frontend, é a aplicação responsável por obter os dados em JSON e converter essa informação em HTML,CSS e JS
            - Nesse momento que o react e outras tecnologias por exemplo entram em ação
        - Podemos ter também um frontend mobile escrito com react native, swift, kotlin, onde irão utilizar dos dados do mesmo backend, porém, agora para criar as interfaces nativas para o celular e não mais o HTML, CSS e JS

        - Apartir do momento que utilizamos um conceito como este de SPA a gente passa a ter a possibilidade de ter multiplos frontends, como por exemplo web, mobile, uma API pública, aplicativo para android, um para IOS...todos consumindo a mesma fonte de dados e consumindo os dados em JSON...pois, JSON é uma estrutura universal, tanto como IOS, quanto android...independente da linguagem, é possível entender JSON e apartir desses dados será criado a interface para o usuário 

        -  Desta forma a gente consegue ter ganhos em performance, manutenibilidade do nosso código.
        - 


    - Rendering Patterns
    - AngularJS, Vue, React, Svelte, Aurelia.


- #### SSR
    - Server Side Rendering
        - Grande maioria das aplicações na web utilizam esse pattern
        - Toda a vez que o usuário(browser) requisita uma página da nossa aplicação, ela é recebida pelo nosso servidor, onde o servidor contém todo nosso código frontend e backend 
        - O Backend interpreta e devolve 'O Browser está solicitando a rota de usuários '
        - Faz toda a busca no banco de dados e cria todo o html para mostrar a listagem de usuários para o browser e devolve o HTML, CSS e JS pronto do servidor para que o navegador possa carregar.

    - Rendering Patterns
    - Ex: Wordpress, Ruby & Rails(Github), Magento...
    - Padrão tradicional

___________________________________________________________________________________________________ 

### Bundlers & compilers - Aula 02

- Nem sempre os browsers, vão suportar as versões mais atualizadas do JS que estamos utilizando.
- Por isso, foram criadas várias ferramentas, para fazer com que a gente consiga converter nosso código que está escrito em um JS mais moderno, para versões específicas do JS que rodam em ambientes específicos.


- #### Compilers
    - São compiladores de código, ou seja, ferramentas que convertem nosso código de um formato para outro
    - Exemplo: Executar um código react em um navegador, porém, este navegador não suporta ainda algumas funcionalides do JS, sendo assim, eu preciso converter este meu código de uma versão mais atualizada do JS para uma versão que o navegador consiga compreender.

    - Compiler + famoso => Babel (https://babeljs.io/)

- #### Bundlers
    - Algo bem comum quando trabalhamos com arquivos escritos JS, é querer distribuir nossa aplicação em múltiplos arquivos, que ao ir importando um arquivo dentro do outro e juntando todas as funcionalidades existentes, resultaram na criação da aplicação com vários aquivos JS que montam a funcionalidade esperada da aplicação
    - Essa ação não é suportada nativamente pelos browsers
    - bundler mais famoso => webpack (https://webpack.js.org/)
    - #### Alternativas

        - #### Vite (https://vitejs.dev/)
            - O vite ele utiliza por padrão os ECMA Scripts models nativos, ou seja, não é necessário bundlers
            - efetua o processo de compiling de forma automática, ou seja, não precisa do babel para executar..ele tem um compilador interno do próprio Vite.
             - Desta forma ao utilizar o Vite, teremos o melhor dos dois (Bundlers e compilers), utilizando os ECM nativos dos próprios browser e assim conseguimos construir nossas aplicações front-end de uma maneira mais performática.

        - #### Snowpack (https://www.snowpack.dev/)
     
___________________________________________________________________________________________________ 

- ### Criação do primeiro projeto - Aula 03

    - #### versionamento do node 
        - n => (https://github.com/tj/n)
        - instalação => npm install -g n

    - #### inicializando o projeto novo
        - npm create vite@latest
        - escolher o framework
        - selecionar a variante
            - SWC(Speedy Web Compiler) é um compilador TypeScript/JavaScript super-rápido escrito em Rust.
        - start no projeto
            - cd 01-fundamentos-reactjs
            - npm install
            - npm run dev

    - #### Curiosidades
        - o vite ele já traz o fast refesh de maneira nativa, ou seja, todas as alterações refletem de maneira automática no browser

___________________________________________________________________________________________________ 