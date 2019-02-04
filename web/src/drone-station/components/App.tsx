import * as React from 'react';

import Footer from 'drone-station/components/Footer';
import Header from 'drone-station/components/Header';
import DroneMapPage from 'drone-station/components/layout/DroneMapPage';

class App extends React.Component {
  render () {
    return (
      <>
        <Header />
        <main role='main' className='d-flex flex-column h-100'>
          <DroneMapPage />
        </main>
        <Footer />
      </>
    );
  }
}

export default App;
