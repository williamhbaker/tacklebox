import React, { useState } from 'react';

import NavBar from './NavBar';

const NavControl = () => {
  const [menuOpen, setMenuOpen] = useState(false);

  const toggleMenu = () => {
    setMenuOpen((prev) => !prev);
  };

  const closeMenu = () => {
    setMenuOpen(false);
  };

  return (
    <NavBar
      menuOpen={menuOpen}
      onToggleClick={toggleMenu}
      onCloseClick={closeMenu}
    ></NavBar>
  );
};

export default NavControl;
