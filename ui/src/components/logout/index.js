import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';

import { logout, selectUser } from 'features/user/userSlice';

import Section from 'components/Section';
import Container from 'components/Container';

const LogOut = () => {
  const history = useHistory();
  const dispatch = useDispatch();
  const user = useSelector(selectUser);

  if (!user) history.push('/');

  useEffect(() => {
    dispatch(logout());
  }, []);

  return (
    <Section>
      <Container>
        <p>Logging you out...</p>
      </Container>
    </Section>
  );
};

export default LogOut;
