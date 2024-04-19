# React
- É uma biblioteca para construção de interfaces altamente interativas
- Interface:
  - É tudo que o usuário interagi quando ele utiliza a aplicação em si
  - Interfaces náo necessariamente estão presentes somente em aplicações web,
  o React pode ser encontrado hoje na construção de aplicações web, mobile,
  dispositivos de tv ou ate mesmo em dispositivos de Realidade virtual, usando
  o react VR
  - O react ele  é uma biblioteca e não um framework, até existem frameworks que
  utilizam o React 
    - Como: Remix, Next, Redwood, Blitz e muitos outros

- Conceitos
  - Rendering Patterns
    - SPA's
      - Single Page Applications
        - AngularJS, Vue, React, Svelte, Aurelia...
        - O Browser quando acessa uma rota específica dentro da aplicação, 
        faz com que o back-end busque dentro do banco de dados e retorne
        essas informações necessárias para que sejam mostradas em tela
        - Todas essas informações sáo retornadas em formato JSON (Javascript 
        Object Notation) ou seja em uma estrutura de dados.
        - Ao transformar toda essa parte de estrutura de dados em HTML, 
        CSS e JS, é aonde entra o React e outras tecnologias como mobile,
        por exemplo.
    - Server-Side-Rendering
      - Quando o usuário requisita uma page da nossa aplicação
      - A Página é recebida pelo nosso servidor e este servidor contêm
      tanto o código front quanto o back-end da aplicação

  - Bundlers
    - Basicamente é fazer a divisão da aplicação JS em vários pequenos arquivos e um vai importando ao outro e ao final ao ir somando esses arquivos resulta na aplicação em si, que tantos esperamos.
    - Ex. de bundler famoso:
      - Webpack

  - Compilers
    - Compiladores de códigos, ou seja, ferramentas que convertem nosso código de um formato para outro.
    - Ex.: Tentando executar um código React em nosso browser
    - Um dos compiladores mais conhecidos: Babel
      - Funciona da seguinte forma: pega sintaxes mais atualizadas do JS e converte para uma sintaxe mais compátiveis com os browsers.


  - Verificando o que os browsers atualmente suportam e não suportam
    - https://caniuse.com

  - Alternativas para criação de projetos (bibliotecas mais famosas)
    - Vite => https://vitejs.dev
    - Snowpack => https://www.snowpack.dev

    