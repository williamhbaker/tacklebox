import React, { useEffect, useState } from 'react';
import { useSelector } from 'react-redux';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faBox,
  faSignInAlt,
  faSignOutAlt,
  faUserPlus,
  faAngleDoubleRight,
} from '@fortawesome/free-solid-svg-icons';

import { selectUser } from 'features/user/userSlice';
import { selectActiveBin } from 'features/bins/binsSlice';

import NavBarLink from './NavBarLink';

const NavBar = () => {
  useEffect(() => {
    document.body.classList.add('has-navbar-fixed-top');
  });

  const user = useSelector(selectUser);
  const activeBin = useSelector(selectActiveBin);

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
            text={'All Bins'}
            onClick={closeMenu}
          ></NavBarLink>
          {activeBin && (
            <NavBarLink
              path={`/bin/${activeBin}`}
              icon={<FontAwesomeIcon icon={faAngleDoubleRight} />}
              text={activeBin}
              onClick={closeMenu}
            ></NavBarLink>
          )}
        </div>
        <div className="navbar-end">
          {user ? (
            <>
              <NavBarLink
                path={'/logout'}
                icon={<FontAwesomeIcon icon={faSignOutAlt} />}
                text={'Log Out'}
                onClick={closeMenu}
              ></NavBarLink>
            </>
          ) : (
            <>
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
            </>
          )}
        </div>
      </div>
    </nav>
  );
};

export default NavBar;
