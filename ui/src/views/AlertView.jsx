
import React from "react";
import Alert from "../components/Alert";
import TopBar from "../components/TopBar"
import Menu from "../components/Menu"
import { withStyles } from '@material-ui/core/styles';

const styles = theme => ({
    root: {
      flexGrow: 1,
    //   height: 440,
      zIndex: 1,
      overflow: 'hidden',
      position: 'relative',
      display: 'flex',
      height: '100%',
    },

    content: {
      flexGrow: 1,
      backgroundColor: theme.palette.background.default,
      padding: theme.spacing.unit * 3,
      minWidth: 0, // So the Typography noWrap works
    },
});

class AlertView extends React.Component {

    constructor(props){
        super(props);
        this.classes = this.props.classes;
    };
    
    render() {
        return (
        <div >
            <div>
                <TopBar />
            </div>
            <div className={this.classes.root}>
                <Menu />
                <Alert id={this.props.match.params.id} className={this.classes.content}/>
            </div>
        </div>
    )}
}

export default withStyles(styles)(AlertView);
