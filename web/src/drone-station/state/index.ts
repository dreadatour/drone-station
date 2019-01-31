import {RouterState} from 'connected-react-router';
import {AnyAction} from 'redux';
import {ThunkDispatch} from 'redux-thunk';

export type State = {
  router: RouterState
};

export type StateDispatch = ThunkDispatch<State, void, AnyAction>;

export default State;
