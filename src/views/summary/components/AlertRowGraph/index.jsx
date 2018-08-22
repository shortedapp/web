import React from 'react';
import {
    VictoryChart,
    VictoryAxis,
    VictoryArea,
    VictoryContainer,
} from 'victory';
// import { LineChart, Line, ResponsiveContainer } from 'recharts';
import {Wrapper} from './style';

/**
 * LegendCompanyMarketCap
 *
 */
class AlertRowGraph extends React.Component {
    constructor(props) {
        super(props);
        this.state = {};
    }
    getTickValues() {
        return [
            new Date(2002, 1, 1),
            new Date(2017, 1, 1),
            new Date(2018, 1, 1),
        ];
    }

    render() {
        return (
            <Wrapper>
                <VictoryChart
                    padding={{top: 0, left: 0, right: 0, bottom: 0}}
                    // height={100}
                    width={850}
                    // style={{parent: {maxWidth: 245}}}
                    // containerComponent={<VictoryContainer responsive />}
                    >
                    <VictoryArea
                        interpolation="natural"
                        style={{
                            data: {
                                fill: '#c0caff',
                                strokeWidth: 2,
                                fillOpacity: '0.4',
                                stroke: '#8396ff',
                            },
                        }}
                        data={this.props.data}
                    />
                    <VictoryAxis
                        tickFormat={() => ''}
                        style={{axis: {stroke: 'none'}}}
                    />
                </VictoryChart>
            </Wrapper>
        );
    }
}

export default AlertRowGraph;