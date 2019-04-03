import React, { Component } from 'react';
import { connect } from 'react-redux';
import * as actions from './store/actions';

class App extends Component {

  componentDidMount() {
    this.props.loadPlayers();
  }

  render() {
    const { players } = this.props;
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
                  <button onClick={e => this.props.disconnectPlayer(player.id)}> Desconectar </button>
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

export default connect(mapStateToProps, actions)(App);
