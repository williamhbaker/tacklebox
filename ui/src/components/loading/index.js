import React from 'react';

import Section from 'components/Section';
import Container from 'components/Container';

const Loading = () => (
  <Section>
    <Container>
      <h3 style={{ marginTop: '10rem' }}>Loading...</h3>
      <progress className="progress is-medium is-info" max="100" />
    </Container>
  </Section>
);

export default Loading;
