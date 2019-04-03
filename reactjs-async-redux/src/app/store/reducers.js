import { ActionTypes } from './types';

const INITIAL_STATE = {
    data: [],
    loading: false,
    error: false,
};

const reducer = (state = INITIAL_STATE, action ) => {
    switch(action.type) {
        case ActionTypes.FETCH_PLAYERS_LOADING:
            return {
                ...state,
                loading: true,
            };
        case ActionTypes.FETCH_PLAYERS_SUCCESS:
            return {
                data: action.payload,
                loading: false,
                error: false,
            }
        case ActionTypes.FETCH_PLAYER_DISCONNECT:
            return {
                data: action.payload,
                loading: false,
                error: false,
            }
        case ActionTypes.FETCH_PLAYERS_FAILURE:
            return {
                ...state,
                loading: false,
                error: true,
            }
        default: 
            return state;
    }
}

export default reducer;