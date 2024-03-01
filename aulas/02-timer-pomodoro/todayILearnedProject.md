# Project 03 - Timer Pomodoro

## - Styled Components
- Styled Components -> Maneira de estilizar nossas aplicações em react, utilizando um conceito chamado CSS in JS
- instalação: npm i @types/styled-components -D
    - instalando como dependencia de desenvolvimento, pois, não precisamos das tipagens em produção, porque em prod, nosso código, vai ser convertido totalmente em JS

- Template literals -> é identificada pelos ``, exemplo:
> 
    import styled from 'styled-components'

    // ao fazer styled.button, estamos herdando o elemento button do proprio HTML
    export const ButtonContainer = styled.button`
        // todo css vem aqui dentro
    `

## - Configurando temas
- É possível criar quantos temas for preciso dentro da aplicação