import {neighbours} from 'latlon-geohash';
import * as React from 'react';

import GeohashBtn from 'drone-station/components/geohash/GeohashBtn';

export type GeohashInfoStateProps = {
  geohash: string
};

export type GeohashInfoDispatchProps = {
  onChange: (geohash: string) => void
};

type GeohashInfoProps = GeohashInfoStateProps & GeohashInfoDispatchProps;

class GeohashInfo extends React.Component<GeohashInfoProps> {
  render () {
    const {geohash, onChange} = this.props;

    const geoNeighbours = neighbours(geohash);

    return (
      <div className='p-3 border-bottom'>
        <table className='table table-sm table-borderless align-middle text-center'>
          <thead>
            <tr>
              <td className='align-bottom text-left'><h5>Quadrant:</h5></td>
              <td><button className='btn btn-info w-100' onClick={() => onChange(`${geohash}s`)}><i className='fas fa-search-plus' /> Zoom In</button></td>
              <td><button className='btn btn-info w-100' onClick={() => onChange(geohash.substr(0, geohash.length - 1))}><i className='fas fa-search-minus' /> Zoom Out</button></td>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td><GeohashBtn geohash={geoNeighbours.nw} onClick={onChange} /></td>
              <td><GeohashBtn geohash={geoNeighbours.n} onClick={onChange} /></td>
              <td><GeohashBtn geohash={geoNeighbours.ne} onClick={onChange} /></td>
            </tr>
            <tr>
              <td><GeohashBtn geohash={geoNeighbours.w} onClick={onChange} /></td>
              <td><GeohashBtn geohash={geohash} active disabled onClick={onChange} /></td>
              <td><GeohashBtn geohash={geoNeighbours.e} onClick={onChange} /></td>
            </tr>
            <tr>
              <td><GeohashBtn geohash={geoNeighbours.sw} onClick={onChange} /></td>
              <td><GeohashBtn geohash={geoNeighbours.s} onClick={onChange} /></td>
              <td><GeohashBtn geohash={geoNeighbours.se} onClick={onChange} /></td>
            </tr>
          </tbody>
        </table>
      </div>
    );
  }
}

export default GeohashInfo;
