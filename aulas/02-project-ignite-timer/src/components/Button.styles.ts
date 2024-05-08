import styled , { css } from 'styled-components'

export type ButtonVariant = 'primary' | 'secondary' | 'danger' | 'success';

interface ButtonContainerProps {
    variant: ButtonVariant;
}

const buttonVariant = {
    primary: 'purple',
    secondary: 'blue',
    danger: 'red',
    success: 'green'
}

export const ButtonContainer = styled.button<ButtonContainerProps>`
    width: 100px;
    height: 40px;
    border-radius: 4px;
    margin: 8px;
    border: 0;

    background-color: ${props => props.theme.secondary};
    color: ${props => props.theme.white};

    /* ${ props => {
        return css`
            background-color: ${buttonVariant[props.variant]}
        `
    } } */


`