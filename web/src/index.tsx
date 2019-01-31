import * as React from 'react';
import * as ReactDOM from 'react-dom';

class App extends React.Component {
  render() {
    return (
      <div>
        <div className='d-flex flex-column flex-md-row align-items-center p-3 px-md-4 mb-3 bg-white border-bottom shadow-sm'>
          <h5 className='my-0 mr-md-auto font-weight-normal'>Drone Station</h5>
        </div>
      </div>
    );
  }
}

const render = () => ReactDOM.render(
  <App />,
  document.getElementById('root')
);

render();
