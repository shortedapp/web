import React from 'react';
import Transition from 'react-transition-group/Transition';
import {Menu, Icon, Switch, Button} from 'antd';
import 'antd/dist/antd.css';
import {ThemeContext, themes} from '../../theme-context';
import Logo from '../../components/Logo';
import ErrorBoundary from '../../components/ErrorBoundary';
import ThemeSwitch from '../../components/ThemeSwitch';
import Sectors from '../../views/sectors';
import Alerts from '../../views/alerts';
import Seasonality from '../../views/seasonality';
import Movers from '../../views/movers';
import Summary from '../../views/summary';
import {
    DashboardWrapper,
    ContentWrapper,
    DashboardNavbarWrapper,
    ThemeWrapper,
    HeaderWrapper,
    NavBarCollapseButton,
} from './style';

const SubMenu = Menu.SubMenu;
class Dashboard extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            theme: 'dark',
            current: 'SUMMARY',
            collapsed: false,
        };
    }
    changeTheme = value => {
        this.setState({
            theme: value ? 'dark' : 'light',
        });
    };
    handleClick = e => {
        console.log('click ', e);
        this.setState({
            current: e.key,
        });
    };
    toggleCollapsed = () => {
        this.setState({
            collapsed: !this.state.collapsed,
        });
    };
    getView(selection, theme) {
        switch (selection) {
            case 'SECTORS':
                return (
                    <ErrorBoundary>
                        <Sectors theme={themes[theme]} />
                    </ErrorBoundary>
                );
            case 'MOVERS':
                return (
                    <ErrorBoundary>
                        <Movers theme={themes[theme]} />
                    </ErrorBoundary>
                );
            case 'ALERTS':
                return (
                    <ErrorBoundary>
                        <Alerts theme={themes[theme]} />
                    </ErrorBoundary>
                );
            case 'SEASONALITY':
                return (
                    <ErrorBoundary>
                        <Seasonality theme={themes[theme]} />
                    </ErrorBoundary>
                );
            case 'SUMMARY':
                return (
                    <ErrorBoundary>
                        <Summary theme={themes[theme]} />
                    </ErrorBoundary>
                );
            default:
                return (
                    <ErrorBoundary>
                        <Summary theme={themes[theme]} />
                    </ErrorBoundary>
                );
        }
    }

    render() {
        return (
            <ThemeContext.Consumer>
                {theme => (
                    <DashboardWrapper
                        width={this.state.collapsed ? `80px` : `200px`}>
                        <HeaderWrapper {...theme.style} />
                        <Logo
                            collapsed={this.state.collapsed}
                            theme={theme.style}
                        />
                        <NavBarCollapseButton
                            {...theme.style}
                            width={this.state.collapsed ? `80px` : `200px`}>
                            <Button
                                type="primary"
                                onClick={this.toggleCollapsed}
                                style={{width: 50}}>
                                <Icon
                                    type={
                                        this.state.collapsed
                                            ? 'menu-unfold'
                                            : 'menu-fold'
                                    }
                                />
                            </Button>
                        </NavBarCollapseButton>
                        <ThemeWrapper
                            {...theme.style}
                            width={this.state.collapsed ? `80px` : `200px`}>
                            <ThemeSwitch
                                theme={theme.name}
                                checked={this.state.theme}
                                changeTheme={value => this.changeTheme(value)}
                            />
                        </ThemeWrapper>

                        <DashboardNavbarWrapper
                            {...theme.style}
                            width={this.state.collapsed ? `80px` : `200px`}>
                            <Menu
                                className="menu"
                                inlineCollapsed={this.state.collapsed}
                                theme={this.state.theme}
                                onClick={this.handleClick}
                                defaultOpenKeys={['sub1']}
                                selectedKeys={[this.state.current]}
                                mode="inline">
                                <Menu.Item key="SUMMARY">
                                    <Icon type="pie-chart" />
                                    <span>Summary</span>
                                </Menu.Item>
                                <Menu.Item key="SECTORS">
                                    <Icon type="desktop" />
                                    <span>Sector Breakdown</span>
                                </Menu.Item>
                                <Menu.Item key="SEASONALITY">
                                    <Icon type="inbox" />
                                    <span>Seasonality</span>
                                </Menu.Item>
                                <Menu.Item key="ALERTS">
                                    <Icon type="warning" />
                                    <span>Alerts</span>
                                </Menu.Item>
                                <Menu.Item key="MOVERS">
                                    <Icon type="line-chart" />
                                    <span>Movers</span>
                                </Menu.Item>
                            </Menu>
                        </DashboardNavbarWrapper>
                        <ContentWrapper {...themes[this.state.theme].style}>
                            <ThemeContext.Provider value={theme.style}>
                                {this.getView(this.state.current, theme)}
                            </ThemeContext.Provider>
                        </ContentWrapper>
                    </DashboardWrapper>
                )}
            </ThemeContext.Consumer>
        );
    }
}

export default Dashboard;
