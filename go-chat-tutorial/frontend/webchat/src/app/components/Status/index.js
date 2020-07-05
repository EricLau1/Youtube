import React from 'react';
import './styles.css';

export class Status extends React.Component {
    render() {
        const { status } = this.props;

        return (
            <span className="status">
                <span className={`status-icon ${status}`}></span>
                <span className="status-text">{status}</span>
            </span>
        );
    }
}