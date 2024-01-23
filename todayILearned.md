___________________________________________________________________________________________________ 


### Fundamentos React

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