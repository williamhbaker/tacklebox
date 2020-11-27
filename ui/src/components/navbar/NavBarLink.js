import React from 'react';
import { NavLink } from 'react-router-dom';

const NavBarLink = ({ path, icon, text, onClick }) => (
  <NavLink
    to={path}
    onClick={onClick}
    className="navbar-item"
    activeClassName="is-active"
  >
    <span className="icon">{icon}</span>
    <span>{text}</span>
  </NavLink>
);

export default NavBarLink;
