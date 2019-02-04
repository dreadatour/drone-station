import * as React from 'react';

import {Drone} from 'drone-station/models/drone';

export type DronesListStateProps = {
  geohash: string
  drones: Drone[] | null
};

export type DronesListDispatchProps = {
  loadDronesList: (quadrant: string) => void
  cleanupDrones: () => void
};

type DronesListProps = DronesListStateProps & DronesListDispatchProps;

class DronesList extends React.Component<DronesListProps> {
  componentDidMount () {
    window.scrollTo(0, 0);
    this.props.loadDronesList(this.props.geohash);
  }

  componentWillUnmount () {
    this.props.cleanupDrones();
  }

  private renderDrones = () => {
    if (this.props.drones === null || this.props.drones.length === 0) {
      return (
          <tr>
            <td colSpan={2}><i>No drones found</i></td>
          </tr>
      );
    }

    return (
      <>
        {this.props.drones.map((drone, i) =>
          <tr key={i}>
            <td className='text-monospace'>{drone.id}</td>
            <td className='text-monospace'>{drone.x},{drone.y}</td>
          </tr>
        )}
      </>
    );
  }

  render () {
    return (
      <div className='p-3'>
        <table className='table table-sm table-borderless align-middle text-left'>
          <thead>
            <tr>
              <td className='align-bottom text-left'><h5>Drones:</h5></td>
              <td className='text-right'><button className='btn btn-primary w-100'><i className='fas fa-plus' /> Add</button></td>
            </tr>
          </thead>
          <tbody>
            {this.renderDrones()}
          </tbody>
        </table>
      </div>
    );
  }
}

export default DronesList;
