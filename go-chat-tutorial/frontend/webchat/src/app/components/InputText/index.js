import React from 'react';

export class InputText extends React.Component {
    render() {
        const { type, placeholder, onChange, defaultValue } = this.props;
        
        return <input 
            type={type ? type : 'text'}
            placeholder={placeholder ? placeholder : ''}
            onChange={e => onChange(e.target.value)}
            value={defaultValue}
        />
    }
}