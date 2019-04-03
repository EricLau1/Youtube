import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import { createPlayer } from '../../../store/actions';

class PlayerForm extends React.Component {

    render() {
        return (
            <form>
                <div>
                    <label>Nickname: </label>
                    <input type="text" onChange={e => this.props.model.nickname = e.target.value } required /> 
                </div>
                <button type="button" onClick={e => this.props.createPlayer(this.props.model)}> Salvar </button>
            </form>
        );
    }
}

const mapStateToProps = state => ({
    model: state.model
});

const mapDispatchToProps = dispatch => bindActionCreators({ createPlayer }, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(PlayerForm);
