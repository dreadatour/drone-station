import {bounds, Neighbours, neighbours} from 'latlon-geohash';
import * as React from 'react';

import GeohashInfo from 'drone-station/components/geohash/GeohashInfo';
import DroneMap, {Bounds} from 'drone-station/components/map/DroneMap';
import DronesListContainer from 'drone-station/containers/drones/DronesListContainer';

type DroneMapPageState = {
  geohash: string
  bounds: Bounds
  neighbours: Neighbours
};

const defaultState = {
  geohash: 'u15pmu',
};

export type DroneMapPageStateProps = {
  geohash?: string
};

export type DroneMapPageDispatchProps = {
};

type DroneMapPageProps = DroneMapPageStateProps & DroneMapPageDispatchProps;

class DroneMapPage extends React.Component<DroneMapPageProps, DroneMapPageState> {
  constructor (props: DroneMapPageProps) {
    super(props);

    const geohash = props.geohash || defaultState.geohash;

    this.state = {...defaultState, geohash: geohash, bounds: this.geohashToBounds(geohash), neighbours: neighbours(geohash)};
  }

  private geohashToBounds = (geohash: string): Bounds => {
    const b = bounds(geohash);
    return [[b.sw.lat, b.sw.lon], [b.ne.lat, b.ne.lon]];
  }

  private onGeohashChange = (geohash: string) => {
    this.setState({geohash: geohash, bounds: this.geohashToBounds(geohash)});
  }

  render () {
    return (
      <div className='row m-0 p-0 h-100'>
        <div className='col-8 m-0 p-0 border-right map-container'>
          <DroneMap bounds={this.state.bounds} />
        </div>
        <div className='col-4 m-0 p-0'>
          <GeohashInfo geohash={this.state.geohash} onChange={this.onGeohashChange} />
          <DronesListContainer quadrant={this.state.geohash} />
        </div>
      </div>
    );
  }
}

export default DroneMapPage;
