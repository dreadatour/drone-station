import {RouterAction} from 'connected-react-router';
import {Dispatch} from 'redux';

import {Error} from 'drone-station/models/error';

export enum ErrorActionTypes {
  ERROR = 'ERROR',
  NOT_FOUND = 'ERROR/NOT_FOUND',
  CLEANUP = 'ERROR/CLEANUP',
}

export type ErrorAction = {
  type: ErrorActionTypes.ERROR
  payload: Error
};

export type ErrorNotFoundAction = {
  type: ErrorActionTypes.NOT_FOUND
};

export type ErrorCleanupAction = {
  type: ErrorActionTypes.CLEANUP
};

export type ErrorActions = ErrorAction
                         | ErrorNotFoundAction
                         | ErrorCleanupAction;

export const errorActionCreator = (error: any): ErrorAction | RouterAction => {
  const payload: Error = {
    code: 'unknown',
    message: 'Unknown error',
  };

  if (error !== undefined && error.response !== undefined) {
    if (error.response.data !== undefined) {
      if (error.response.data.code !== undefined) {
        payload.code = error.response.data.code;
      }
      if (error.response.data.message !== undefined) {
        payload.message = error.response.data.message;
      }
      if (error.response.data.payload !== undefined) {
        payload.meta = error.response.data.payload;
      }
    } else if (error.response.message !== undefined) {
      payload.message = error.response.message;
    }
  }

  return {
    type: ErrorActionTypes.ERROR,
    payload: payload,
  };
};

export const errorNotFoundActionCreator = (): ErrorNotFoundAction => {
  return {
    type: ErrorActionTypes.NOT_FOUND,
  };
};

const errorCleanupActionCreator = (): ErrorCleanupAction => {
  return {
    type: ErrorActionTypes.CLEANUP,
  };
};

export const cleanupErrors = () => {
  return (dispatch: Dispatch) => {
    dispatch(errorCleanupActionCreator());
  };
};
