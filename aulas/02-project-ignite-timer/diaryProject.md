# Infos e Anotações Gerais do Projeto

- ## Styled Components
    - Maneira de estilizar as nossas aplicações em React, utilizando um conceito chamado CSS in JS
        - CSS in JS, como o próprio nome já diz, iremos escrever o CSS da nossa aplicação no formato ou sitaxe
        da linguagem JS.
    - Como estamos utilizando o TS, o styled components, é um daqueles pacotes que ele não traz consigo as tipagens.
    - As declarações de tipagens deste pacote, elas ficam em um pacote separado, que seria o @types/styled-components
    - Styled components, basicamente são abstrações de CSS, ou seja, estilizações

        - Declarando uma constante, iniciando com letra em caixa alta, pois, é um componente.
        - Como esse button vai ser um componente estilizado, eu preciso sinalizar qual elemento HTML ele será e faremos isso da seguinte forma:
            - Como iremos fazer um button, então, importamos o próprio button do html
            - Para adicionar o CSS, iremos utilizar Template-literals, que é nada mais que ``, logo após o styled.button.

            >
                export const ButtonContainer = styled.button`
                    width: 100px;
                    height: 40px;
                `

    - Ao realizarmos uma interpolação, ou seja, incluir um código JS dentro de uma string 'maior', que é o que esta sinalizado pelas ``, dentro da criação do nosso button de exemplo, o styled-components irá executar como uma função e este vai enviar pra a função todas as propriedades do meu ButtonContainer, como podemos verificar logo abaixo:

        > 
            <ButtonContainer variant={variant}>
                Enviar
            </ButtonContainer>

    - As propriedades serão acessadas da seguinte forma:

        >
            ${ props => {
                return `background-color: ${buttonVariant[props.variant]}`
            } }



- ## Instalação
    - npm i @types/styled-components -D
    - instalando como dependencia de desenvolvimento, pois, não precisamos das tipagens em produção, 
    porque em prod, nosso código, vai ser convertido totalmente em JS


- ## Configuração/Gestão de temas ( Styled Components)
    - É possível ter quantos temas julgarmos necessário
    - Os temas são controlados por JS
    - Equivale as variáveis criadas no css puro

- ## Tipagem de temas
    - Quando usamos TS em meus projetos, eu tenho a possibilidade de criar arquivos de tipagem especificos minha aplicação, ou seja, customizados.

    - Criação de uma pasta @types
        - arquivo styled.d.ts
            - mas porque o d.ts?
                - Significa que eu so vou ter código de definição de tipos do TS e nunca código JS ou qualquer coisa assim.

    - typeof => Função específica do TS (type operator)

    - declare module => modo de criar uma tipagem para o modulo que esta sendo utilizado no momento
        - Isso irá inferir no momento em que importarmos o modulo em algum arquivo, a tipagem que será 'puxada', será a que for definida dentro do escopo em questão

- ## Global Styles

    - No style-components, os estilos globais serão com sintaxe JS e não CSS
    - inicialmente usamos função própria do style-components, chamada createGlobalStyle e apartir dela conseguimos definir todo o escopo de estilização mais global da nossa aplicação


- ## Configurando o ESLint
    - ESLint => EcmaScript Linting
        - Linting => Processo para validar que teu código segue padrões estipulados pelos criadores do projeto em si.

     - Instalação
        - npm i eslint -D

        - pacote especifico do eslint, orientado e criado pela rocketseat
            - npm i @rocketseat/eslint-config -D
                - esse pacote ele traz alguns padrões de escrita de código para JS e claro que pode ser criado o próprio padrão, mas hoje estaremos utilizando este padrão criado pela escola.

        - para criação de configs pessoais:
            - npx eslint --init
            - responder as perguntas feitas pelo vscode e assim, será criado sua própria config do eslint e ficando assim um padrão próprio do seu projeto.

        - criando arquivo eslintrc.json, para aplicar um extends com o nome do pacote e como o pacote da rocket, tem diversas configs, iremos pegar a própria do react:
            >
                {
                    "extends": "@rocketseat/eslint-config/react"
                }

        - Para testar:
            - rodar no terminal: 
                - npx eslint src --ext .ts,.tsx
                    - src: pasta aonde esta os arquivos
                    - --ext: a extensão dos arquivos que será verificado os erros
                    - logo mem seguida adicionar as extensões

        - Para corrigir os erros de maneira geral:
            - rodar o mesmo comando acima, porém, com --fix adicionado ao final do comando:
                - npx eslint src --ext .ts,.tsx --fix


- ## Páginas e rotas
    - React Router Dom

        - inicialmente para trabalhar com as rotas, iremos utilizar no projeto a famosa biblioteca react-router e como trbalharemos com aplicações Web, iremos usar especificamente o react-router-dom

        - instalação:
            - npm i react-router-dom

    - BrowserRouter
        - Ao analisarmos a criação e importação de qualquer rota/page na nossa aplicação, veremos que no console, vai dar vários problemas inicialmente, desta forma, se analisarmos a [documentação do react-router](https://github.com/remix-run/react-router) e for no [getting-started](https://github.com/remix-run/react-router/blob/main/docs/start/_tutorial.md), veremos que a gente precisa importar o BrowserRouter do react-router-dom e realizar um wrap do nosso app dentro BrowserRouter 

    - Contexts-Providers
        - São components que basicamente não tem efeito visual nenhum, mas eles produzem um contexto para os components que estão dentro deles, ou seja, informações para estes components saberem do contexto de 'fora' ou digamos do mundo externo, para que estes obtenham informações do que eles já possuem 'naturalmente' ao serem criados

- ## Layout de Rotas

    - No [Layout](https://www.figma.com/design/ELte8mBmDArgUI3z3IxNwW/Ignite-Timer-(Community)?node-id=313-1874&t=PQkB3aBu26wW0jf7-0) que temos hoje, podemos verificar que temos partes em comuns nas pages da aplicação e sendo assim, para simplesmente ter essa parte 'constante' nas pages(routes), podemos criar um component, como por exemplo, um Header que será uma parte constante na aplicação e o conteúdo realmente vindo logo na sequência e está questão de transição por exemplo entre routes(pages), podemos fazer com base na utilização do component Outlet que vem de dentro do react-router-dom, que é nada mais que um espaço ao qual será inserido um conteúdo.

    - Dentro do react-router-dom, temos um component chamado NavLink e este nos auxiliará na hora da transição entre as rotas, este será utilizado no lugar das ancoras que estavamos utilizando no inicio do projeto.

        >
            <HeaderContainer>
                <img src={logoIgnite} alt="" />
                <nav>
                    <a href="">
                        <Timer size={24} />
                    </a>
                    <a href="">
                        <Scroll size={24} />
                    </a>
                </nav>
            </HeaderContainer>

        >
            <HeaderContainer>
                <img src={logoIgnite} alt="" />
                <nav>
                    <NavLink to="">
                        <Timer size={24} />
                    </NavLink>
                    <NavLink to="">
                        <Scroll size={24} />
                    </NavLink>
                </nav>
            </HeaderContainer>

        - obs.: Na ancora usamos a propriedade href, já no NavLink usamos o to


- ## Aprimorando Inputs
    - Elemento HTML datalist
        - Lista de sugestões para um input

- ## Página: History
    - Em aplicações mobile não existe uma forma de representar uma tabela de forma satisfatoria, por isso na aplicação envolvemos a tabela com uma div e quando for acessado em um dispositivo mobile e a tela estiver menor, será feito com que o usuário final consiga dar scroll na tabela, ou seja, consiga arrastar de um lado para o outro na tela, vendo assim a tabela como um todo.
    - Overflow: auto
        - Faz com que  se o tamanho da tabela, for maior do que o cantainer que ela esta inserida, gere automaticamente uma barra de rolagem 


    - border-collapse: collapse

- ## Status (Component)
    - Nosso component de status na aplicação, com base no styled-component, não precisa necessariamente ser um arquivo separado, so porque ele é visualmente diferente e sim pode ser somente um elemento estilizado, 

    - 