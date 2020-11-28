import React from 'react';
import Header from 'components/Header';
import Section from 'components/Section';
import Container from 'components/Container';
import Form from './form';

const FormContainer = () => {
  return (
    <Section>
      <Container>
        <Header>Log In</Header>
        <Form></Form>
      </Container>
    </Section>
  );
};

export default FormContainer;
