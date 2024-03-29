import styled, { css } from 'styled-components';

export type ButtonVariantColor = 'primary' | 'secondary' | 'danger' | 'success';

interface ButtonContainerProps {
    variant: ButtonVariantColor;
}

const buttonVariants = {
    primary: 'purple',
    secondary: 'orange',
    danger: 'red',
    success: 'green'
}

export const ButtonContainer = styled.button<ButtonContainerProps>`
    width: 100px;
    height: 40px;
    border-radius: 4px;
    border: 0;
    margin: 0 5px;

    background-color: ${props => props.theme.primary};
    color: ${props => props.theme.white};

    //interpolação de string
    /* ${props => {
        return css`background-color: ${buttonVariants[props.variant]}`
    }} */
`