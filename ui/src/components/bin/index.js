import React from 'react';

import { useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';

import Header from 'components/Header';
import Section from 'components/Section';
import Container from 'components/Container';
import HookList from './HookList';
import HookDetails from './HookDetails';

import { selectUser } from 'features/user/userSlice';

const Bin = () => {
  const { id } = useParams();
  const user = useSelector(selectUser);

  return (
    <Section>
      <Container>
        {user ? (
          <>
            <Header>Hooks for {id}</Header>
            <p>{`${window.location.protocol}//${window.location.hostname}/api/hook/${id}`}</p>
            <div className="columns">
              <div className="column">
                <HookList id={id} />
              </div>
              <div className="column">
                <HookDetails />
              </div>
            </div>
          </>
        ) : (
          <p>You must be logged in to view this page.</p>
        )}
      </Container>
    </Section>
  );
};

export default Bin;
