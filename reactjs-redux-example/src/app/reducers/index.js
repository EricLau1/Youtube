import { combineReducers } from 'redux';
import { INCREMENT, DECREMENT } from '../actions/types';

export const counter = (state = {number: 0}, action) => {
    switch(action.type) {
        case INCREMENT:
            return { ...state, number: ++state.number };
        case DECREMENT:
            if(state.number > 0) {
                return { ...state, number: --state.number };
            }
            return state;
        default:
            return state;
    }
};

export default combineReducers({
    counter,
});