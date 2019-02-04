import * as React from 'react';

export type GeohashBtnStateProps = {
  geohash: string
  active?: boolean
  disabled?: boolean
};

export type GeohashBtnDispatchProps = {
  onClick: (geohash: string) => void
};

type GeohashBtnProps = GeohashBtnStateProps & GeohashBtnDispatchProps;

class GeohashBtn extends React.Component<GeohashBtnProps> {
  private onClick = () => {
    const {geohash, disabled, onClick} = this.props;

    if (geohash && !disabled) {
      onClick(geohash);
    }
  }

  render () {
    const {geohash, active, disabled, children} = this.props;

    if (!geohash) {
      return null;
    }

    if (active) {
      return <button className='btn text-danger w-100 text-monospace' disabled={true}>{geohash}</button>;
    }

    return <button className='btn btn-light w-100 text-monospace' disabled={disabled} onClick={this.onClick}>{children || geohash}</button>;
  }
}

export default GeohashBtn;
