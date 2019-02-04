import {bounds, Neighbours, neighbours} from 'latlon-geohash';
import * as React from 'react';

import DronesList from 'drone-station/components/drones/DronesList';
import DronesLost from 'drone-station/components/drones/DronesLost';
import DronesMap, {Bounds} from 'drone-station/components/drones/DronesMap';
import GeohashInfo from 'drone-station/components/geohash/GeohashInfo';
import {Drone, DroneAdd} from 'drone-station/models/drone';

type DronesMapPageState = {
  geohash: string
  bounds: Bounds
  neighbours: Neighbours
  dronesMissed: Drone[]
};

export type DronesMapPageStateProps = {
  geohash: string
  drones: Drone[] | null
};

export type DronesMapPageDispatchProps = {
  loadDronesList: (quadrant: string) => void
  addDrone: (drone: DroneAdd) => void
  cleanupDrones: () => void
};

type DronesMapPageProps = DronesMapPageStateProps & DronesMapPageDispatchProps;

class DronesMapPage extends React.Component<DronesMapPageProps, DronesMapPageState> {
  timer?: number;

  constructor (props: DronesMapPageProps) {
    super(props);

    const geohash = props.geohash;

    this.state = {geohash: geohash, bounds: this.geohashToBounds(geohash), neighbours: neighbours(geohash), dronesMissed: []};
  }

  componentDidMount () {
    this.setTimer(this.props.geohash);
  }

  componentWillUnmount () {
    if (this.timer !== undefined) {
      clearTimeout(this.timer);
    }
    this.props.cleanupDrones();
  }

  componentWillReceiveProps (nextProps: DronesMapPageProps) {
    if (this.props.drones === null || nextProps.drones === null) {
      return;
    }

    const ids = this.props.drones.map((drone) => drone.id);
    nextProps.drones.forEach((drone) => {
      const found = ids.indexOf(drone.id);
      if (found > -1) {
        ids.splice(found, 1);
      }
    });

    if (ids.length > 0) {
      const dronesMissed = this.state.dronesMissed;

      this.props.drones.forEach((drone) => {
        if (ids.indexOf(drone.id) > -1) {
          dronesMissed.push(drone);
          alert(`Drone ${drone.id} missed!`);
        }
      });

      this.setState({dronesMissed: dronesMissed});
    }
  }

  private setTimer = (geohash: string) => {
    if (this.timer !== undefined) {
      clearTimeout(this.timer);
    }
    this.timer = setInterval(
      () => {
        this.props.loadDronesList(geohash);
      },
      500
    );
  }

  private geohashToBounds = (geohash: string): Bounds => {
    const b = bounds(geohash);
    return [[b.sw.lat, b.sw.lon], [b.ne.lat, b.ne.lon]];
  }

  private onGeohashChange = (geohash: string) => {
    this.setState({geohash: geohash, bounds: this.geohashToBounds(geohash)}, () => {
      this.setTimer(geohash);
    });
  }

  render () {
    return (
      <div className='row m-0 p-0 h-100'>
        <div className='col-8 m-0 p-0 border-right map-container'>
          <DronesMap
            bounds={this.state.bounds}
            geohash={this.state.geohash}
            drones={this.props.drones}
            addDrone={this.props.addDrone}
          />
        </div>
        <div className='col-4 m-0 p-0 list-container'>
          <GeohashInfo geohash={this.state.geohash} onChange={this.onGeohashChange} />
          <DronesLost drones={this.state.dronesMissed} />
          <DronesList geohash={this.state.geohash} drones={this.props.drones} />
        </div>
      </div>
    );
  }
}

export default DronesMapPage;
