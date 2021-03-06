

import React from 'react';
import PropTypes from 'prop-types';

import { withStyles } from '@material-ui/core/styles';
import MenuList from '@material-ui/core/MenuList';
import MenuItem from '@material-ui/core/MenuItem';
import Divider from  '@material-ui/core/Divider'
import ListItemText from '@material-ui/core/ListItemText'
import Drawer from '@material-ui/core/Drawer';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import DashboardIcon from '@material-ui/icons/Dashboard';
import ViewListIcon from '@material-ui/icons/ViewList';
import CodeIcon from '@material-ui/icons/Code';
import { Link } from 'react-router-dom'
import MenuItemLink from './MenuItemLink';

// const styles = {
//     root: {
//         width: '200px',
//         overflowX: 'auto',
//     },
//     };

const drawerWidth = 200;

const styles = theme => ({
  root: {
    flexGrow: 1,
    height: 440,
    zIndex: 1,
    overflow: 'hidden',
    position: 'relative',
    display: 'flex',
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
  },
  drawerPaper: {
    position: 'relative',
    width: drawerWidth,
  },
  content: {
    flexGrow: 1,
    backgroundColor: theme.palette.background.default,
    padding: theme.spacing.unit * 3,
    minWidth: 0, // So the Typography noWrap works
  },
  toolbar: theme.mixins.toolbar,
  menuItem: {
    '&:focus': {
      backgroundColor: theme.palette.primary.main,
      '& $primary, & $icon': {
        color: theme.palette.common.white,
      },
    },
  },
  primary: {},
  icon: {},
});

class Menu extends React.Component {

    static contextTypes = {
        router: PropTypes.object
    }

    constructor(props, context) {
        super(props, context);
        this.classes = this.props.classes;
    }

    render() {
        return (
        <div>
            <Drawer
                    variant="permanent"
                    classes={{
                    paper: this.classes.drawerPaper,
                    }}
                >
                <MenuList>
                    {/* <MenuItemLink className={this.classes.menuItem} to="/" >
                        <ListItemIcon className={this.classes.icon}>
                            <DashboardIcon />
                        </ListItemIcon>
                        <ListItemText classes={{ primary: this.classes.primary }} inset primary="Dashboard" />
                    </MenuItemLink> */}
                    <MenuItemLink className={this.classes.menuItem} to="/alerts">
                        <ListItemIcon className={this.classes.icon}>
                            <ViewListIcon />
                        </ListItemIcon>
                        <ListItemText classes={{ primary:this.classes.primary }} inset primary="Alerts List" />
                    </MenuItemLink>
                    {/* <MenuItemLink className={this.classes.menuItem} to="/">
                        <ListItemIcon className={this.classes.icon}>
                            <CodeIcon />
                        </ListItemIcon>
                        <ListItemText classes={{ primary: this.classes.primary }} inset primary="..." />
                    </MenuItemLink> */}
                </MenuList>
            </Drawer>
        </div>
        )};
}

export default withStyles(styles)(Menu);

            {/* <List component="nav">
                <ListItem button>
                <ListItemText primary="Inbox" />
                </ListItem>
                <Divider />
                <ListItem button divider>
                <ListItemText primary="Drafts" />
                </ListItem>
                <ListItem button>
                <ListItemText primary="Trash" />
                </ListItem>
                <Divider light />
                <ListItem button>
                <ListItemText primary="Spam" />
                </ListItem>
            </List> */}