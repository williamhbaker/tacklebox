import React from 'react';

import Header from 'components/Header';
import Section from 'components/Section';
import Container from 'components/Container';
import Form from './form';

const SignUp = () => {
  return (
    <Section>
      <Container>
        <Header>Sign Up</Header>
        <Form></Form>
      </Container>
    </Section>
  );
};

export default SignUp;
