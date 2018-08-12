import styled from 'styled-components';


export const Wrapper = styled.button`
    width: 40px;
    height: 40px;
    display: flex;
    justify-content: center; /* align horizontal */
    align-items: center; /* align vertical */
    margin: 5px;
    border-radius: 7px;
    color: white;
    background: ${ props => props.selected ? `#3d1abd` : `#8361ff`};
    outline:0;
    font-size: 20px;
    font-family: "Avenir Next,Segoe UI",Helvetica,Arial,sans-serif,"Apple Color Emoji","Segoe UI Emoji","Segoe UI Symbol";
    font-weight: bold;
`