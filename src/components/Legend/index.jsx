import React from 'react';

import { Wrapper } from './style';
/**
 * Renders a shorted.com.au logo
 */
const Legend = (props) => (
    <Wrapper>
        <p>legend goes here</p>
        <p>code: {props.code}</p>
    </Wrapper>
)

export default Legend;