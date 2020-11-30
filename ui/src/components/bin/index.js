import React from 'react';

import { useParams } from 'react-router-dom';

import Header from 'components/Header';
import Section from 'components/Section';
import Container from 'components/Container';
import Hooklist from './HookList';

const Bin = () => {
  const { id } = useParams();

  return (
    <Section>
      <Container>
        <Header>Hooks for {id}</Header>
        <Hooklist id={id} />
      </Container>
    </Section>
  );
};

export default Bin;
