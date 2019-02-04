import * as React from 'react';
import {Map, Marker, Pane, Popup, Rectangle, TileLayer} from 'react-leaflet';

import {Drone, DroneAdd} from 'drone-station/models/drone';
import {LeafletMouseEvent} from 'leaflet';

export type Bounds = [number, number][];

export type DronesMapStateProps = {
  bounds: Bounds
  geohash: string
  drones: Drone[] | null
};

export type DronesMapDispatchProps = {
  addDrone: (drone: DroneAdd) => void
};

type DronesMapProps = DronesMapStateProps & DronesMapDispatchProps;

class DronesMap extends React.Component<DronesMapProps> {
  private marker = (drone: Drone) => {
    const {bounds} = this.props;

    const latitude = bounds[0][0] + (bounds[1][0] - bounds[0][0]) * parseFloat(drone.y) / 100;
    const longitude = bounds[0][1] + (bounds[1][1] - bounds[0][1]) * parseFloat(drone.x) / 100;

    return (
      <Marker key={drone.id} position={[latitude, longitude]}>
        <Popup>
          ID: {drone.id}<br />
          Quadrant: {this.props.geohash}<br />
          X: {drone.x}<br />
          Y: {drone.y}<br />
          Latitude: {latitude.toFixed(6)}<br />
          Longitude: {longitude.toFixed(6)}<br />
        </Popup>
      </Marker>
    );
  }

  private addDrone = (e: LeafletMouseEvent) => {
    const {lat, lng} = e.latlng;
    const {bounds} = this.props;

    if (lat <= bounds[0][0] || lat >= bounds[1][0] || lng <= bounds[0][1] || lng >= bounds[1][1]) {
      return;
    }

    const x = 100 * (lng - bounds[0][1]) / (bounds[1][1] - bounds[0][1]);
    const y = 100 * (lat - bounds[0][0]) / (bounds[1][0] - bounds[0][0]);

    this.props.addDrone({
      quadrant: this.props.geohash,
      x: x.toFixed(2),
      y: y.toFixed(2),
    });
  }

  render () {
    const {bounds, drones} = this.props;

    return (
      <Map
        bounds={bounds}
        animate={false}
        boxZoom={false}
        doubleClickZoom={false}
        dragging={false}
        keyboard={false}
        scrollWheelZoom={false}
        touchZoom={false}
        zoomControl={false}
        onClick={this.addDrone}
      >
        <TileLayer
          attribution='&amp;copy <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
          url='//maps05.a-cdn.net/tiles/housinganywhere/{z}/{x}/{y}.png?language=en&scale=2'
        />
        <Pane name='quadrant' style={{zIndex: 499}}>
          <Rectangle
            bounds={bounds}
            color='#dc3545'
            fill={false}
            weight={1}
            attribution='foo'
            interactive={false}
          />
        </Pane>
        {drones !== null && drones.map((drone) =>
          this.marker(drone)
        )}
      </Map>
    );
  }
}

export default DronesMap;
