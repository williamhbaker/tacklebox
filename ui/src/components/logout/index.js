import React, { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';

import { logout } from 'features/user/userSlice';

import Section from 'components/Section';
import Container from 'components/Container';

const LogOut = () => {
  const history = useHistory();
  const dispatch = useDispatch();

  useEffect(() => {
    (async () => {
      await dispatch(logout());
      history.push('/');
    })();
  });

  return (
    <Section>
      <Container>
        <p>Logging you out...</p>
      </Container>
    </Section>
  );
};

export default LogOut;
