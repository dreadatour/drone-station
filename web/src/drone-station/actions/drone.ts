import Axios from 'axios';
import {Dispatch} from 'redux';

import {errorActionCreator} from 'drone-station/actions/error';
import {API_URL} from 'drone-station/config';
import {Drone, DroneAdd} from 'drone-station/models/drone';

export enum DroneActionTypes {
  LIST = 'DRONE/LIST',
  ADD = 'DRONE/ADD',
  REMOVE = 'DRONE/REMOVE',
  CLEANUP = 'DRONE/CLEANUP',
}

type GetDronesListAction = {
  type: DroneActionTypes.LIST
  payload: Drone[]
};

type AddDroneAction = {
  type: DroneActionTypes.ADD
  payload: Drone
};

type RemoveDroneAction = {
  type: DroneActionTypes.REMOVE
  payload: Drone
};

type CleanupDronesAction = {
  type: DroneActionTypes.CLEANUP
};

export type DronesActions = GetDronesListAction
                          | AddDroneAction
                          | RemoveDroneAction
                          | CleanupDronesAction;

const getDronesListActionCreator = (drones: Drone[]): GetDronesListAction => ({
  type: DroneActionTypes.LIST,
  payload: drones,
});

const addDroneActionCreator = (drone: Drone): AddDroneAction => ({
  type: DroneActionTypes.ADD,
  payload: drone,
});

const removeDroneActionCreator = (drone: Drone): RemoveDroneAction => ({
  type: DroneActionTypes.REMOVE,
  payload: drone,
});

const cleanupDronesActionCreator = (): CleanupDronesAction => ({
  type: DroneActionTypes.CLEANUP,
});

export const loadDronesList = () => {
  return (dispatch: Dispatch) => {
    Axios.get(`${API_URL}/v1/drones`)
      .then((response) => {
        dispatch(getDronesListActionCreator(response.data.payload));
      })
      .catch ((error) => {
        dispatch(errorActionCreator(error));
      });
  };
};

export const addDrone = (drone: DroneAdd) => {
  return (dispatch: Dispatch) => {
    Axios.post(`${API_URL}/v1/drones`, drone)
      .then((response) => {
        dispatch(addDroneActionCreator(response.data.payload));
      })
      .catch ((error) => {
        dispatch(errorActionCreator(error));
      });
  };
};

export const removeDrone = (drone: Drone) => {
  return (dispatch: Dispatch) => {
    Axios.post(`${API_URL}/v1/drones/${drone.id}`)
      .then(() => {
        dispatch(removeDroneActionCreator(drone));
      })
      .catch ((error) => {
        dispatch(errorActionCreator(error));
      });
  };
};

export const cleanupDrones = () => {
  return (dispatch: Dispatch) => {
    dispatch(cleanupDronesActionCreator());
  };
};
