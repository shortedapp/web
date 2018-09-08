import styled from 'styled-components';

export const WrapperHeader = styled.div`
    display: grid;
    color: ${props => props.textColor};
    background: ${props => props.widgetRowBackgroundColor};
    @media (min-width: 901px) {
        grid-template-columns: repeat(5, 1fr);
        grid-template-areas: 'code name name name percentage';
    }
    @media (max-width: 900px) {
        grid-template-columns: repeat(2, 1fr);
        grid-template-areas: 'code percentage';
    }
    margin: 7px;
    height: 40px;
    margin-bottom: 20px;
    align-items: center;
    .code {
        grid-area: code;
        margin-left: 10px;
    }
    .company-name {
        grid-area: name;
        text-align: center;
        margin-right: 30px;
    }
    .percentage {
        grid-area: percentage;
        text-align: center;
    }
`;
export const Wrapper = styled.a`
    display: grid;
    z-index: 0;
    position: relative;
    color: ${props => props.textColor};
    background: ${props => props.widgetRowBackgroundColor};
    text-decoration: none !important;
    @media (min-width: 901px) {
        grid-template-columns: repeat(5, 1fr);
        grid-template-areas: 'code name name name percentage';
    }
    @media (max-width: 900px) {
        grid-template-columns: repeat(2, 1fr);
        grid-template-areas: 'code percentage';
    }
    margin: 4px;
    height: 100%;
    border-radius: 0 30px 30px 0;
    &:hover,
    &:visited,
    &:link,
    &:active {
        text-decoration: none !important;
        color: ${props => props.textColor};
    }
    align-items: center;
`;
export const WrapperHovered = styled.a`
    display: grid;
    position: relative;
    color: ${props => props.textColor};
    background: ${props => props.widgetRowBackgroundColor};
    text-decoration: none !important;
    z-index: 10;
    grid-template-columns: repeat(5, 1fr);
    grid-template-areas: 'code name name name percentage';
    margin: 4px;
    height: 100%;
    border-radius: 0 30px 30px 0;
    transform: scale(1.05);
    transition-duration: 0.1s;
    -webkit-box-shadow: -3px 4px 7px 0px rgba(0, 0, 0, 0.25);
    -moz-box-shadow: -3px 4px 7px 0px rgba(0, 0, 0, 0.25);
    box-shadow: -3px 4px 7px 0px rgba(0, 0, 0, 0.25);
    &:hover,
    &:visited,
    &:link,
    &:active {
        text-decoration: none !important;
        color: ${props => props.textColor};
    }
    align-items: center;
`;

export const Name = styled.div`
    grid-area: name;
    @media (max-width: 1300px) {
        display: none;
    }
    display: inline-block;
    vertical-align: middle;
    flex-wrap: wrap;
    flex: 1;
    flex-direction: column;
    justify-content: center;
    vertical-align: middle;
    font-size: 14px;
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
    min-width: 0;
    margin-right: 30px;
    -webkit-line-clamp: 1;
    -webkit-box-orient: vertical;
`;

export const Code = styled.div`
    grid-area: code;
    display: flex;
    flex-direction: column;
    justify-content: center;

    .code {
        width: 60px;
        height: 45px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        vertical-align: middle;
        text-align: center;
        font-size: 21px;
        font-weight: bold;
    }
`;

export const Percent = styled.div`
    grid-area: percentage;
    margin-left: auto;
    padding-right: 8px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    vertical-align: middle;

    .circle {
        background: #f98080;
        height: 45px;
        width: 70px;
        border-radius: 10px 50px 50px 10px;
        display: flex;
        flex-direction: column;
        justify-content: center;
        vertical-align: middle;
        text-align: center;
        font-size: 16px;
        font-weight: 400;
    }
`;
