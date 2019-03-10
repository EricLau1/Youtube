import React, { Component } from 'react';
import { Link } from 'react-router-dom';

import Header from '../../components/Header';

export default class Dashboard extends Component {

    constructor() {
        super();
        this.state = {
            user: {},
        }
    }

    componentDidMount() {
        const token = localStorage.getItem('token');
        fetch('http://localhost:9000/admin', { headers: new Headers({ 'Authorization': `Bearer ${token}` })})
        .then(response => {
            if(response.ok) {
                return response.json();
            }
            throw new Error("Oops! Ocorreu um erro. :(");
        })
        .then(user => this.setState({ user }))
        .catch(e => console.log(e));
    }

    render() {
        return (
            <div>
                <Header title="Dashboard" />
                <hr className="my-3" />
                <p>
                    <code> {this.state.user.nickname}, {this.state.user.email} logado com sucesso! ^-^  </code>
                </p>
                <div className="text-center">
                    <Link to="/logout" className="btn btn-outline-primary"> Log Out </Link>
                </div>
            </div>
        );
    }
}