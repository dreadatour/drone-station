import * as React from 'react';

class Header extends React.PureComponent {
  render () {
    return (
      <header>
        <nav className='navbar navbar-expand sticky-top bg-light border-bottom'>
          <h3 className='navbar-brand mb-0'>Drone Station</h3>
        </nav>
      </header>
    );
  }
}

export default Header;
