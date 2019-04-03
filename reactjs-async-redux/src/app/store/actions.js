import { ActionTypes } from './types';

const fetchPlayersLoading = () => {
    return {
        type: ActionTypes.FETCH_PLAYERS_LOADING
    }
}

const fetchPlayersSuccess = data => {
    return {
        type: ActionTypes.FETCH_PLAYERS_SUCCESS,
        payload: data,
    };
}

const fetchPlayersFailure = error => {
    return {
        type: ActionTypes.FETCH_PLAYERS_FAILURE,
        payload: error
    }
}

export const loadPlayers = () => {
    return dispatch => {
        dispatch(fetchPlayersLoading())
        return fetch('http://localhost:9000/players')
                .then(response => {
                    if(response.ok) {
                        return response.json();
                    }
                    throw new Error("Oops! :(");
                })
                .then(json => {
                    dispatch(fetchPlayersSuccess(json))
                    return json;
                })
                .catch(error => dispatch(fetchPlayersFailure(error)));
    }
}

const fetchDisconnectPlayer = data => {
    return {
        type: ActionTypes.FETCH_PLAYER_DISCONNECT,
        payload: data
    };
};

export const disconnectPlayer = id => {
    return dispatch => {
        dispatch(fetchPlayersLoading())
        return fetch(`http://localhost:9000/players/${id}`, {method: 'PUT'})
                .then(response => {
                    if(response.ok) {
                        return response.json();
                    }
                    throw new Error("Oops! :(");
                })
                .then(json => {
                    console.log(json)
                    dispatch(fetchDisconnectPlayer(json))
                    return json;
                })
                .catch(error => dispatch(fetchPlayersFailure(error)));
    }
}