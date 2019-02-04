import * as React from 'react';

import Footer from 'drone-station/components/Footer';
import Header from 'drone-station/components/Header';
import DronesMapPageContainer from 'drone-station/containers/drones/DronesMapPageContainer';

class App extends React.Component {
  render () {
    return (
      <>
        <Header />
        <main role='main' className='d-flex flex-column h-100'>
          <DronesMapPageContainer quadrant='u15pmu' />
        </main>
        <Footer />
      </>
    );
  }
}

export default App;
