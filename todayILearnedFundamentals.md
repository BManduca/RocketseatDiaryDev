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

- ### Componentes - Aula 04

    - Componente em si é uma forma de 'desacoplar' um trecho do código ou da nossa interface para que isso se torne repetitivel ou que se torne muito mais fácil de aplicar uma manutenção futuramente em um trecho muito menor da interface em si.
    - O Componente a gente pode entender como um trecho ou uma parte da interface da nossa aplicação e que pode ser repetida várias vezes e inclusive com informações diferentes.
    - Basicamente Componente é uma função que retorna HTML
        - todos os componentes da aplicação react precisão tem extensão JSX
            - JSX = JavaScript + XML (HTML)
            - Basicamente é um arquivo JavaScript com HTML dentro dele 
___________________________________________________________________________________________________ 

- ### Propriedades - Aula 05

    - informações que são passadas para componentes
    - Assim como passamos para as tags HTML, os atributos necessários, nos componentes temos um comportamento parecido, porém o que passamos são propriedades.
    - Com os componentes conseguimos aproveitar o HTML, estilização, funcionamento com JS e acima de tudo diferenciação no uso ou na chamada do componente dentro da aplicação.
    - Propriedades alteram elementos visuais de um componente sem que seja preciso perder a flexibilidade de abstrair um componente em outro arquivo para reaproveitar elementos que sejam comuns dentro desses componentes 
___________________________________________________________________________________________________ 

- ### CSS Modules - Aula 06

    - Escopo: quando temos algumas propriedades atreladas a um componente e gostariamos que não afetasse outra parte por exemplo da aplicação, nos definimos esta dentro de um escopo, ou seja, em uma estrutura(espaço) aonde tudo que for designado ali, so será aplicado ao componente.
    - Scoped CSS: é quando temos um css especifico de um componente e não afete outras partes do código ou da aplicação, como mencionamos logo acima.
    - Vite CSS Modules => https://vitejs.dev/guide/features#css-modules
    - Para que um CSS fique totalmente atrelado a um componente e nao interfira em outros elementos, a extensão do arquivo deve ficar como (NomeDoArquivo).module.css
    - Quando trabalhamos com CSS modules, buscamos trabalhar essencialmente com Classes e não utilizar ID's ou qualquer outro tipo de identificação/seleção
___________________________________________________________________________________________________ 


- ### CSS Global - Aula 07

    - tabela cores(mostrando codigo hash) com base no tailwind => https://tailwindcss.com/docs/customizing-colors
    - -webkit-font-smoothing: antialiased
        - propriedade que faz as letras ficarem mais desenhadas, mais fininhas, com um certo padrão de detalhamento e isso pode ser usado geralmente quando não utilizamos as fontes que são padrão do sistema.

    - ![Link Preconnect font Google](/RocketseatDiaryDev/aulas/01-fundamentos-reactjs/img/linkRoboto.png)
        - Links que vão estabelecer digamos a conexão do nosso servidor com o servidor da Google, para buscar as fontes

    - Utilizar tamanho da fonte relativas (rem) => ao utilizar 1rem, por exemplo, quer dizer que eu vou estar colocando/utilizando 1 unidade relativa do tamanho pardão do HTML, que é 16px por padrão, caso o usuário não aumente e nem diminua, o tamanho da fonte será esse.

___________________________________________________________________________________________________ 

- ### Components: Sidebar - Aula 08
    - div por padrão já é display-block, ou seja, sempre terá 100% da largura.
    - ao utilizarmos o grid-template-columns, podemos colocar valores fixos para algumas colunas, porém, podemos também utilizar 1fr, por exemplo, onde fr é a unidade utilizada no grid quando queremos dizer que algo vai ter um tamanho flexível.
    - overflow: hidden => ele faz com que o elemento ocupe somente o espaço ou tamanho da sidebar que estamos utilizando na aplicação 
    - https://phosphoricons.com/ => biblioteca para icons
___________________________________________________________________________________________________ 

- ### Components: Post - Aula 09
    - Tag time => Ela permite que passemos um atríbuto chamado 'dateTime'
        - as propriedades que possuem a junção de mais de uma palavra, sempre utilizam camelCase
___________________________________________________________________________________________________ 

- ### Estilizações do Post - Aula 10
    - Quando utilizamos o seguinte padrão: .post > header {}
        - significa que iremos aplicar algum estilo na header que esta diretamente dentro do post
        (8:10)
        - focus-within => é utilizado quando basicamente é encontrado alguma situação aonde houver foco em algum input, text-area ou qualquer elemento dentro do local aonde estamos trabalhando, no qual será aplicado alguma estilização, dentro deste container que estamos trabalhando.
         
___________________________________________________________________________________________________ 

- ### Estilizações do Comentário - Aula 12

    - utilizando display: flex, para deixar o avatar e o comentário lado a lado no box de comentário
    - Quando presenciamos uma div, envolvendo o elemento que vamos estilizar e que tem como propriedade o display: flex, podemos utilizar o flex: 1 em qualquer elemento dentro dessa div 'pai' e isso faz com que a div seja 'esticada', ou seja, ela vai acabar utilizando todo o espaço ali presente
    - border: 0; -> para resetar o estilo padrão que os buttons tem no html 
___________________________________________________________________________________________________ 

- ### Componente: Avatar - Aula 13
    - Dentro do react existem dois grandes momentos onde criamos um componente 
        - 1º: Quando algo repete muito em tela 
        - 2º: Quando é possível tirar algo de um componente maior, sem que este mesmo pare de funcionar, deta forma, o componente maior fica mais "clean", com uma funcionalidade mais clara para receber a devida manutenção.
        - Quando queremos que todo Avatar, ou seja, o componente chamado para representar o avatar por exemplo presente em nossa aplicação, que não receber a propriedade harBorder como False, tenha por padrão ele com valor True e para isso podemos fazer da seguinte forma:
            - const hasBorder = props.hasBorder !== false; -> isso quer dizer que a borda esta presente e se caso a propriedade nem for enviada, automaticamente, ela já é diferente de False.

        - Desestruturação: é um recurso onde é possível extrair elementos de arrays ou propriedades de objetos em variáveis separadas, tudo em uma única linha, transformando assim o seu estilo de codificação, bem mais conciso e melhorando a legibilidade.
        - As unidades relativas com relação a responsividade, dão um poder gigantesco na hora de codificar

___________________________________________________________________________________________________ 


### Parte 3 -> Os motores do ReactW

- ## Iterando no JSX

    - Iteração: O Fato de repetir algo, ou seja, criar uma estrutura de repetição.
    - para realizar a iteração, iremos utilizar o map, pois este método ele da um return e sendo assim, irá mostrar as informações que estão sendo iteradas.
        - Para este caso de iteração, não devemos escolher o forEach, pois, ele não tem um retorno
            Ex.: se fizermos const algo = posts.forEach(post => {}), o algo sempre será void, ou seja ao usar o forEach como mencionado acima, nada será retornado em tela.
        - De maneira obrigatória, após chamar o método map, devemos aplicar um return, para que tudo seja retornado como esperado.

- ## Propriedades do Post
    - Para trabalharmos com data, estaremos utilizando o intl
        - https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Intl/DateTimeFormat

    - intl permite que seja feito formatação de datas, números, pluralização, listas e muito mais

    - Existem outras bibliotecas para se trabalhar com datas, como por exemplo:
        - date-fns -> npm i date-fns
            - doc: https://date-fns.org/v3.3.1/docs/format
            - formatDistanceToNow -> para realizar o cálculo da data tomando como parâmetro a distância da entre a data da publicação com relação a data de hoje, por exemplo.
                - parâmetros:
                    - 1º a data em si
                    - 2º o idioma (locale: ptBR)
                    - 3º addSuffix: boolean


        - No JS temos a seguinte forma também de trabalhar com dates
            - toISOString()
                - retorn a uma cadeia de caracteres (string) simplificada no formato ISO
                extendido (ISO 8601), que é sempre de 24 ou 27 caracteres de tamanho -> YYYY-MM-DDTHH:mm:ss.sssZ ou ±YYYYYY-MM-DDTHH:mm:ss.sssZ, respectivamente. O fuso horário é sempre o deslocamento zero UTC, como denotado pelo sufixo "z".


- ## Estado (useState)
    - Estados -> são basicamente variáveis que eu gostaria que o componente monitorasse
    - o useState não retorna somente o valor da variável, na verdade são retornado duas 'coisas' dentro de um array, por isso geralmente usamos o conceito de desestruturação.
    - spread Operator -> faz a leitura da variável, que muitas vezes vem ser um array e cópia os valores existentes na váriavel


- ## Inserindo comentários (Programação declarativa)
    - Programação imperativa
        - O que deve ser feito (passo-a-passo)
        -  ex: no nosso código temos => event.target.comment.value = '';
            - é considerado imperativo, pois, estamos dando dizendo que é preciso limpar o valor da textarea do meu comentário.

        - Mas dentro do react, geralmente evita utilziar a programação imperativa

    - Programação declarativa
        - ao invês de falarmos como eu espero atingir o meu resultado final (passo-a-passo), 
        eu so declaro qual é o resultado que eu espero
        - Quais seriam as condições para eu ter o resultado final.
        - dentro do react geralmente iremos usar esse tipo de programação
        - Através da forma declarativa no React, basicamente conseguimos repassar quais são as condições para as "coisas" acontecerem, mas não executamos as mesmas de forma imperativa.
        - Dentro da nossa textarea, como escolhemos pela forma declarativa, através do onChang, sempre estaremos setando a function handleNewCommentChange e ao enviar o formulario, estaremos chamando a function handleCreateNewComment, que ira fazer de forma declarativa a reinicialização do campo textarea, refletindo assim sempre o valor que vai estar no estado e também ira gravar o comentário como esperado.

- ## Entendendo a key
    - No React quando fazemos o uso de uma lista(iteração com map), cada component ou ate mesmo elemento, é importante que ele tenha uma propriedade chamada key e essa propriedade recebe qualquer valor, seja número, strings...so que é preciso passar pra ela, qual a informação única de cada um dos posts , ou seja, de cada um dos itens da minha lista 
    - Porque a key precisa ser única?
        - 3 principais momentos em que o component é renderizado novamente no React
            Exemplo: Relembrando da aplicação.. toda a vez que alteramos a lista de comentários, é "chamado" o método setComments, colocando assim um novo comentário no final e ai um comentário aparece em tela
        => Momentos em que o componente é totalmente renderizado
        - Quando o state é alterado, todo o código do component, quando uma das varíaveis por exemplo, newCommentText ou comments é alterada, é executado novamente.
        - Quando uma propriedade é alterada
            - todo o código do componente é executado novamente
        - Quando um componente pai renderiza novamente.

    - Quando falamos sobre lista
        -  no react, quando temos um post novo, é analisado através da key, qual post foi alterado por exemplo ou se tem um posto novo realmente e ai ele "recebe" o valor das key's e verifica após a atualização... por exemplo: 
            Posts: 1 2 3 4
                   1 2 3 4 5

                   ao renderizar novamente, o código vai reparar que temos um post novo e que o restante já existia

            - então desta forma ajuda o React a entender o que já existia em tela e o que não existia em tela.


    - Porque não devemos usar o índice do array?
        - const posts = [1,2,3,4]
              => índices 0,1,2,3
                - se caso eu alterasse o post 1 com o post 3, iria acarretar na renderização do component, poŕem, os índices não trocariam de lugar, então, basicamente foi so trocado os posts e não as posições... e isso pode confundir o codigo todo... porque desta forma, sempre vai parecer que os posts estão mudando de posição...
        - Por isso é preciso usar uma informação, onde é possível identificar cada item da lista unicamente, mesmo que seja dado um f5 na tela, o valor precisa se manter o mesmo..

- ## Comunicações entre componentes
    - A única forma de comunicar um componente com outro, é através das suas propriedades
    - Hoje não é somente necessário passar como propriedades para nossos componentes, as strings, números... podemos passar também Funções e é assim que se comunica um componente com outro dentro do react
    - Da maneira como realizamos a comunicação entre componentes, buscamos criar a função deleteComment e enviamos como parâmetro para a função Comment e desta forma, o componente de comentário ele terá uma forma de se comunicar com o componente Pai, chamando essa função.

- ## Removendo comentários (imutabilidade)
    - Imutabilidade: as variáveis não sofrem mutação, ou seja, nunca alteramos o valor de uma variável na memória, nos criamos um novo espaço na memória.
    - Imutabilidade permite ser muito mais performático, principalmente dentro do React, que precisa ficar comparando as variáveis, para verificar se as mesmas foram alteradas.

- ## Validando formulário
    - propriedade onInvalid, é sempre chamada sempre que o HTML identifica que tentamos realizar um submit de um form, só que o texto do campo era inválido, ou seja, pode estar vazio.
    - setCustomValidity(): método usado para identificar qual mensagem de validação que queremos para o campo envolvido 

- ## Realizando likes nos comentários
    - Usaremos um estado para ficar observando o número de 'likes'
    - Iniciar o estado com alguma informação que seja do mesmo tipo que a informação que vai ser armazenada
    - ao adicionarmos o evento de onClick() para um button, ele sempre estará aguardando uma função, por isso devemos fazer o seguinte, por exemplo:
        > 
            <button onClick(handleLikeComment)> 
                <ThumbsUp />
                Aplaudir <span>{likeCount}</span>
            </button>
        
        - porém, se passar dessa forma:
            >
                <button onClick(setLikeCount(likeCount+1))> 

            -> Já estaria incorreto, pois, está sendo passado a execução de uma função em si e não a chamada da função.

        - A única forma de realizar a chamada de uma função diretamente no onClick(), é fazendo através de uma arrowFunction:
            > 
                <button onClick(() => setLikeCount(likeCount+1))> 

- ## Entendendo closures no React
    - Ao chamarmos por exemplo a função de alterar estado e como se o React criasse outro contexto para execução desse component, so que todo código presente no bloco da função por exemplo, sempre vai estar trabalhando no mesmo contexto, independente de quantas formas for acionada a chamada do método dentro da função
    - O react permite que a gente passe para qualquer função de atualização de estado, uma função, através de uma arrow Function e dentro da mesma, devemos retornar o novo valor atualizado. Através deste modo o react permite que a gente possa acessar o valor mais recente da propriedade, como um argumento da função que esta sendo utilizada/chamada
    - Sempre que for atualizar uma informação e essa informação depender do valor que ela tinha anteriormente, é sempre bom fazer atualização, usando esse padrão de funções.


- ## Fundamentos do TypeScript
    - TS nasceu como um superset, ou seja, um conjunto ferramental baseado na linguagem Javascript, para adicionar tipagem estática em cima de uma linguagem como Javascript que tem linguagem dinâmica.
    
- ## TypeScript no React
    - Generics: tipagem dentro do TS

- ## Extensão de interfaces
    - Quando temos um component e o retorno dele é uma tag HTMl como image ou button e queremos permitir que os outros componentes que chamarem este mesmo como estamos fazendo dentro da nossa aplicação de exemplo...
        - quando utilizamos o componente Comment, dentro dele é acionado o component Avatar e que eles possam além de ter propriedades que nós criamos como hasNorder por exemplo, as outras propriedades que possam ser acionadas pelo component e isso pode ser feito através das extensões presentes dentro do TS

    - Para evitar de ter que ficar chamando todas as propriedades necessárias ou que pretendemos utilizar no momento de dar um export por exemplo em uma function, podemos utilizar um pequeno hack existente, que é o restOperator, que geralmente vem da seguinte forma:
        >
            ...nomedaprops

    