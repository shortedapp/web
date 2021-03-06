import React from 'react';
import Transition from 'react-transition-group/Transition';
import {ThemeContext} from '../../../../theme-context';
import MoversListRow from '../MoversListRow';
import {Wrapper, Header, More, duration, transitionStyles} from './style';
/**
 * Renders a list of TopShortsListRow components. These rows will display the stock code, stock name, and percentage shorted
 * for the top say 20 stocks. With a "show more" button present at the bottom, taking them to a different view which will me dedicated to showing more short position
 * information for more stocks (maybe top 100). Will perhaps a more verbose set of properties and graphics.
 */
class MoversList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            inside: false,
            selected: false,
        };
    }
    componentDidMount() {
        this.toggleEnterState();
    }

    toggleEnterState() {
        this.setState({inside: true});
    }
    handleRowSelect(value) {
        this.setState({selected: value});
    }
    handleMouseLeave() {
        this.setState({
            selected: false,
        });
    }

    render() {
        const rows = this.props.data.data
            .slice(0, 5)
            .map(row_data => (
                <MoversListRow
                    handleSelect={() => this.handleRowSelect(row_data.code)}
                    selectedRow={this.state.selected}
                    row
                    key={row_data.code}
                    {...row_data}
                />
            ));
        return (
            <ThemeContext.Consumer>
                {theme => (
                    <Transition timeout={duration} in appear>
                        {state => {
                            return (
                                <Wrapper
                                    onMouseLeave={() => this.handleMouseLeave()}
                                    {...theme}
                                    duration={duration}
                                    {...transitionStyles[state]}>
                                    <Header>Top Movers</Header>
                                    <MoversListRow header key={'row-header'} />
                                    {rows}
                                    <More {...theme}>show more</More>
                                </Wrapper>
                            );
                        }}
                    </Transition>
                )}
            </ThemeContext.Consumer>
        );
    }
}

export default MoversList;
