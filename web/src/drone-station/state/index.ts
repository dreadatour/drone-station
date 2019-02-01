import {RouterState} from 'connected-react-router';
import {AnyAction} from 'redux';
import {ThunkDispatch} from 'redux-thunk';

import DronesState from 'drone-station/state/drone';

export type State = {
  router: RouterState
  drones: DronesState
};

export type StateDispatch = ThunkDispatch<State, void, AnyAction>;

export default State;
