import React from 'react';

const Heading = (props) => (
  <header>
    <h1 className="title mb-5">{props.children}</h1>
  </header>
);

export default Heading;
