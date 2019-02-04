import {DroneActionTypes, DronesActions} from 'drone-station/actions/drone';
import {Drone} from 'drone-station/models/drone';
import DronesState from 'drone-station/state/drone';

const initialState: DronesState = {
  list: null,
};

function drones (state: DronesState = initialState, action: DronesActions): DronesState {
  let list = state.list === null ? new Map<string, Drone>() : state.list;

  switch (action.type) {
  case DroneActionTypes.LIST:
    list = new Map<string, Drone>();
    action.payload.forEach((drone) => list.set(drone.id, drone));
    return {...state, list: list};

  case DroneActionTypes.ADD:
    list.set(action.payload.id, action.payload);
    return {...state, list: list};

  case DroneActionTypes.CLEANUP:
    return initialState;

  default:
    return state;
  }
}

export default drones;
