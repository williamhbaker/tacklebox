import React, { useEffect, useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faBox,
  faSignInAlt,
  faUserPlus,
} from '@fortawesome/free-solid-svg-icons';

import NavBarLink from './NavBarLink';

const NavBar = () => {
  useEffect(() => {
    document.body.classList.add('has-navbar-fixed-top');
  });

  const [menuOpen, setMenuOpen] = useState(false);

  const toggleMenu = () => {
    setMenuOpen((prev) => !prev);
  };

  const closeMenu = () => {
    setMenuOpen(false);
  };

  return (
    <nav className="navbar is-primary is-fixed-top">
      <div className="navbar-brand">
        <div
          className={`navbar-burger burger ${menuOpen && 'is-active'}`}
          onClick={toggleMenu}
        >
          <span></span>
          <span></span>
          <span></span>
        </div>
      </div>

      <div className={`navbar-menu ${menuOpen && 'is-active'}`}>
        <div className="navbar-start">
          <NavBarLink
            path={'/bins'}
            icon={<FontAwesomeIcon icon={faBox} />}
            text={'Bins'}
            onClick={closeMenu}
          ></NavBarLink>
        </div>
        <div className="navbar-end">
          <NavBarLink
            path={'/login'}
            icon={<FontAwesomeIcon icon={faSignInAlt} />}
            text={'Log In'}
            onClick={closeMenu}
          ></NavBarLink>
          <NavBarLink
            path={'/signUp'}
            icon={<FontAwesomeIcon icon={faUserPlus} />}
            text={'Sign Up'}
            onClick={closeMenu}
          ></NavBarLink>
        </div>
      </div>
    </nav>
  );
};

export default NavBar;
