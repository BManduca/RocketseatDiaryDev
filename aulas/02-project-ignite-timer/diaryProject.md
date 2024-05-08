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
                - Significa que eu so vou ter código de definição de tipos do TS e nunca código TS ou qualquer coisa assim.
                