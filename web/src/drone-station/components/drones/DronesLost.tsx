import * as React from 'react';

import {Drone} from 'drone-station/models/drone';

export type DronesLostStateProps = {
  drones: Drone[]
};

export type DronesLostDispatchProps = {};

type DronesLostProps = DronesLostStateProps & DronesLostDispatchProps;

class DronesLost extends React.Component<DronesLostProps> {
  render () {
    if (this.props.drones.length === 0) {
      return null;
    }

    return (
      <div className='p-3 border-top'>
        <table className='table table-sm table-borderless align-middle text-left'>
          <thead>
            <tr>
              <td><h5 className='text-danger'>Drones lost:</h5></td>
              <td></td>
            </tr>
          </thead>
          <tbody>
            {this.props.drones.map((drone, i) =>
              <tr key={i}>
                <td className='text-danger text-monospace'>{drone.id}</td>
                <td className='text-monospace'>{drone.x},{drone.y}</td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    );
  }
}

export default DronesLost;
