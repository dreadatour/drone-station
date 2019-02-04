import * as React from 'react';
import {Map, Pane, Rectangle, TileLayer} from 'react-leaflet';

export type Bounds = [number, number][];

export type DroneMapStateProps = {
  bounds: Bounds
};

export type DroneMapDispatchProps = {
};

type DroneMapProps = DroneMapStateProps & DroneMapDispatchProps;

class DroneMap extends React.Component<DroneMapProps> {
  render () {
    const {bounds} = this.props;
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
      </Map>
    );
  }
}

export default DroneMap;
