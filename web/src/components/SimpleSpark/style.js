import styled from 'styled-components';

export const Wrapper = styled.div`
    background: transparent;
    border-radius: 5px;
    display: flex;
    height: 100%;
`;
export const Header = styled.div`
    grid-area: header;
    padding-top: 3px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    text-align: center;
`;
export const Chart = styled.div`
    grid-area: chart;
    display: flex;
    flex-direction: column;
    justify-content: center;
    text-align: center;
    height: 82px;
    width: 100%;
`;
