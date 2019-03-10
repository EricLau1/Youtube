import { Component } from 'react';

export default class Logout extends Component {

    componentWillMount() {
        localStorage.removeItem('token');
        this.props.history.push('/');
    }

    render() {
        return null;
    }
}