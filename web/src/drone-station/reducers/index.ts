import {connectRouter} from 'connected-react-router';
import {History} from 'history';
import {combineReducers} from 'redux';

import drones from 'drone-station/reducers/drone';
import State from 'drone-station/state';

const rootReducer = (history: History) => combineReducers<State>({
  router: connectRouter(history),
  drones: drones,
});

export default rootReducer;
