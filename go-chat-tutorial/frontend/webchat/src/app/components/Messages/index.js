import React from 'react';

export class Messages extends React.Component {
    render() {
        const { messages } = this.props;

        return messages && messages.map(message => <p key={message.id}><strong>{message.sender}:</strong> {message.body}</p>);
    }
}