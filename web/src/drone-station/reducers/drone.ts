import {DroneActionTypes, DronesActions} from 'drone-station/actions/drone';
import DronesState from 'drone-station/state/drone';

const initialState: DronesState = {
  list: {},
};

function drones (state: DronesState = initialState, action: DronesActions): DronesState {
  let list = state.list;

  switch (action.type) {
  case DroneActionTypes.LIST:
    list = {};
    action.payload.forEach((drone) => list[drone.id] = drone);
    return {...state, list: list};

  case DroneActionTypes.ADD:
    list[action.payload.id] = action.payload;
    return {...state, list: list};

  case DroneActionTypes.CLEANUP:
    return initialState;

  default:
    return state;
  }
}

export default drones;
