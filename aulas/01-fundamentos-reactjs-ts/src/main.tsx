import React from 'react'
import ReactDOM from 'react-dom/client'
// ReactDOM => faz a integração do React(Core) com a DOM
// DOM -> Document Object Model => representação do HTML através do JS
// Ao importar o ReactDOM, estamos integrando o React para funcionar em um ambiente WEB(Browser)
import { App } from './App.js'

/* 

  através do método createRoot, recebemos um parâmetro, que é o elemento raiz da nossa página HTML 

  nesse caso nosso elemento raiz, é o root.

  O react apartir de agora, vai criar o HTML, o CSS e todo o JS da nossa aplicação, dentro
  dessa div root, que é o elemento raiz do html
  
  - quando construimos uma aplicação SPA, toda a interface é construída apartir do JS, ou seja, é o JS 
que constrói a interface da nossa aplicação

  - A interface ela não fica direto nos arquivos HTML
  - Ela fica dentro do JS e este tem total controle sobre essa interface, desta forma fica muito 
  mais fácil de manipular essa interface.

  - Método render
    - tem presente components do react, que funcionam semelhantes a tags do html
    - renderiza(mostra em tela) algo que é interno do react
    - logo abaixo vemos que o App (arquivo importado) esta sendo renderizado.
*/ 
ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    {/*
      o react basicamente esta 'jogando' este conteúdo do JS la para 
      dentro do nosso elemento com id='root' que esta dentro do HTML
    */}
    <App />
  </React.StrictMode>,
)
