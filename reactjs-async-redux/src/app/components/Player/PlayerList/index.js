import React, { Component } from 'react';
import { connect } from 'react-redux';
import * as actions from '../../../store/actions';

class PlayerList extends Component {

  componentDidMount() {
    this.props.loadPlayers();
  }

  render() {
    const { players } = this.props;
    console.log(this.props);
    return (
      <table>
        <tbody>
          {
            players.map(player => (
              <tr key={player.id}>
                <td>{player.nickname}</td>
                <td className={player.online?'online':'offline'}>
                  {player.online?'online':'offline'}
                </td>
                <td>
                  <button onClick={e => this.props.disconnectPlayer(player.id)}> {player.online?'Desconectar': 'Conectar'} </button>
                </td>
                <td>
                  <button type="button" onClick={e => this.props.deletePlayer(player.id)}>Delete</button>
                </td>
              </tr>
            ))
          }
        </tbody>
      </table>
    );
  }
}

const mapStateToProps = state => ({
  players: state.data
});

export default connect(mapStateToProps, actions)(PlayerList);
