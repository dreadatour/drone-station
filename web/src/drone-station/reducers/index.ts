import {connectRouter} from 'connected-react-router';
import {History} from 'history';
import {combineReducers} from 'redux';

import State from 'drone-station/state';

const rootReducer = (history: History) => combineReducers<State>({
  router: connectRouter(history),
});

export default rootReducer;
