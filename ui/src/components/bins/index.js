import React from 'react';

import Header from 'components/Header';
import Section from 'components/Section';
import Container from 'components/Container';
import BinList from './BinList';

const Bins = () => (
  <Section>
    <Container>
      <Header>Bins</Header>
      <BinList />
    </Container>
  </Section>
);

export default Bins;
