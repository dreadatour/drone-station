import {connect} from 'react-redux';

import {addDrone, cleanupDrones, loadDronesList} from 'drone-station/actions/drone';
import DronesMapPage, {DronesMapPageDispatchProps, DronesMapPageStateProps} from 'drone-station/components/DronesMapPage';
import {Drone} from 'drone-station/models/drone';
import State from 'drone-station/state';

type DronesMapPageContainerOwnProps = {
  quadrant: string
};

const mapStateToProps = (state: State, ownProps: DronesMapPageContainerOwnProps): DronesMapPageStateProps => {
  let drones: Drone[] | null = null;

  if (state.drones.list !== null) {
    drones = Array
    .from(state.drones.list.values())
    .sort((a, b) => {
      return a.id.localeCompare(b.id);
    });
  }

  return {
    geohash: ownProps.quadrant,
    drones: drones,
  };
};

const mapDispatchToProps: DronesMapPageDispatchProps = {
  loadDronesList: loadDronesList,
  addDrone: addDrone,
  cleanupDrones: cleanupDrones,
};

const DronesMapPageContainer = connect(mapStateToProps, mapDispatchToProps, null, {pure: false})(DronesMapPage);

export default DronesMapPageContainer;
