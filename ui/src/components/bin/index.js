import React from 'react';

import { useParams } from 'react-router-dom';

import Header from 'components/Header';
import Section from 'components/Section';
import Container from 'components/Container';
import HookList from './HookList';
import HookDetails from './HookDetails';

const Bin = () => {
  const { id } = useParams();

  return (
    <Section>
      <Container>
        <Header>Hooks for {id}</Header>
        <div className="columns">
          <div className="column">
            <HookList id={id} />
          </div>
          <div className="column">
            <HookDetails />
          </div>
        </div>
      </Container>
    </Section>
  );
};

export default Bin;
