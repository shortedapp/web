import React from 'react';
import Transition from 'react-transition-group/Transition';
import {
    TopShortsWrapper,
    CentreGraphWrapper,
    PickerWrapper,
    ChartWrapper,
 } from './style';
import headerBackground from '../../assets/images/header-background.svg';

import AppViewWrapper from './../../components/AppViewWrapper';
import WindowPicker from './../../components/WindowPicker';
import TopChart from './../../components/TopChart';

const duration = 300;

const defaultStyle = {
  width: 100,
  height: 100,
  backgroundImage: `url(${headerBackground})`,
  transition: `opacity ${duration}ms ease-in-out`,
  opacity: 0,
};

const transitionStyles = {
  entering: { opacity: 0 },
  entered: { opacity: 1 },
  exited: { opacity: 0}
};
/**
 * View:TopShorts
 * Overarching container for the top short view.
 * Showing the following key widgets/displays:
 *  * top 10 short positions graphically.
 *  * table of top short position changes
 *  * alerts/anomalies
 *  * reactive-widget on hover/select of a given graph
 * TODO:
 * * add Transitions of components such as window picker, graph and background etc.
 * * add graph integration via recharts etc.
 * * add legend component for on-select animation/effect
 * 
 */

class TopShorts extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            pickerOptions: {
                values: ['d', 'w', 'm', 'y'],
            },
            selectedWindow: false,
            inside: false
        }
    }
    handleWindowSeleted(value) {
        console.log(value)
        this.setState({selectedWindow: value})
    }
    componentDidMount() {
        this.toggleEnterState();
    }
    
    toggleEnterState() {
        this.setState({ inside: true });
    }

    render() {
        const { inside } = this.state;
        return (
            <Transition timeout={duration} in={true} appear={true}>
            {
                state => {
                    return (
                    <AppViewWrapper
                        background={headerBackground}
                        opacity={transitionStyles[state].opacity}
                        duration={duration}
                        >
                        <TopShortsWrapper>
                            <CentreGraphWrapper >
                                <PickerWrapper >
                                    <WindowPicker
                                        options={this.state.pickerOptions}
                                        selectedOption={this.state.selectedWindow}
                                        handleSelect={(e) => this.handleWindowSeleted(e)}
                                    />
                                </PickerWrapper>
                                <ChartWrapper >
                                    <TopChart />
                                </ChartWrapper>      
                            </CentreGraphWrapper>
                        </TopShortsWrapper>
                    </AppViewWrapper>
                    )
                }
            }
            </Transition>
            )
            
    }
}


export default TopShorts;
