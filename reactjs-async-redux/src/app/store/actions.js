import { ActionTypes } from './types';

const fetchPlayersLoading = () => {
    return {
        type: ActionTypes.FETCH_PLAYERS_LOADING
    };
};

const fetchPlayersSuccess = data => {
    return {
        type: ActionTypes.FETCH_PLAYERS_SUCCESS,
        payload: data,
    };
};

const fetchPlayersFailure = error => {
    return {
        type: ActionTypes.FETCH_PLAYERS_FAILURE,
        payload: error
    };
};

const fetchCreatePlayer = data => {
    return {
        type: ActionTypes.FETCH_PLAYER_CREATE,
        payload: data,
    };
};

const fetchDisconnectPlayer = data => {
    return {
        type: ActionTypes.FETCH_PLAYER_DISCONNECT,
        payload: data
    };
};

const fetchDeletePlayer = data => {
    return {
        type: ActionTypes.FETCH_PLAYER_DESTROY,
        payload: data,
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
                    dispatch(fetchDisconnectPlayer(json))
                    return json;
                })
                .catch(error => dispatch(fetchPlayersFailure(error)));
    }
}

export const deletePlayer = id => {
    console.log(id);
    return dispatch => {
        dispatch(fetchPlayersLoading())
        return fetch(`http://localhost:9000/players/${id}`, {method: 'DELETE'})
                .then(response => {
                    if(response.ok) {
                        return response.json();
                    }
                    throw new Error("Oops! :(");
                })
                .then(json => {
                    dispatch(fetchDeletePlayer(json))
                    return json;
                })
                .catch(error => dispatch(fetchPlayersFailure(error)));
    }
}



export const createPlayer = player => {
    return dispatch => {
        dispatch(fetchPlayersLoading())
        return fetch(`http://localhost:9000/players`, {method: 'POST', body: JSON.stringify(player), headers: new Headers({'Content-Type': 'application/json'})})
                .then(response => {
                    if(response.ok) {
                        return response.json();
                    }
                    throw new Error("Oops! :(");
                })
                .then(json => {
                    console.log(json);
                    dispatch(fetchCreatePlayer(json))
                    return json;
                })
                .catch(error => dispatch(fetchPlayersFailure(error)));
    }
}