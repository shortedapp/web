import React from 'react';

import LegendCompanyHeader from '../../components/LegendCompanyHeader';
import LegendCompanyMarketCap from '../../components/LegendCompanyMarketCap';
import {
    Wrapper,
    UnselectedWrapper,
    CompanyName } from './style';
import ShortedAPI from '../../services/sapi/client';
/**
 * Renders a shorted.com.au logo
 * TODO: add data fetch here, async or prefetch based of top-short positions
 */
class Legend extends React.Component {
    constructor(props) {
        super(props)
        this.apiClient = new ShortedAPI()
    }
    render() {
        const data = this.apiClient.getStockSummary(this.props.code)
        const profile = this.props.code ? (
            <Wrapper>
                <LegendCompanyHeader code={this.props.code}/>
                    <CompanyName>{data.metadata.name}</CompanyName>
                <LegendCompanyMarketCap data={data.data} />
            </Wrapper>
            ) : (<UnselectedWrapper><p>hover over graph to show profile</p></UnselectedWrapper>);
        return profile
    }
}

export default Legend;