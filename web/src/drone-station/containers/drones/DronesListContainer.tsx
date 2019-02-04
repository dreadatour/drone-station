import {connect} from 'react-redux';

import {cleanupDrones, loadDronesList} from 'drone-station/actions/drone';
import DronesList, {DronesListDispatchProps, DronesListStateProps} from 'drone-station/components/drones/DronesList';
import {Drone} from 'drone-station/models/drone';
import State from 'drone-station/state';

type DronesListContainerOwnProps = {
  quadrant: string
};

const mapStateToProps = (state: State, ownProps: DronesListContainerOwnProps): DronesListStateProps => {
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

const mapDispatchToProps: DronesListDispatchProps = {
  loadDronesList: loadDronesList,
  cleanupDrones: cleanupDrones,
};

const DronesListContainer = connect(mapStateToProps, mapDispatchToProps, null, {pure: false})(DronesList);

export default DronesListContainer;
