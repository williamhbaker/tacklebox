import React, { useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faEnvelope,
  faLock,
  faExclamationTriangle,
} from '@fortawesome/free-solid-svg-icons';

import { login } from 'api';

const Form = () => {
  const [fields, setFields] = useState({ email: '', password: '' });

  const handleInputChange = (e) => {
    setFields({
      ...fields,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    login(fields);
  };

  return (
    <form
      style={{ maxWidth: '30rem', margin: '0 auto' }}
      onSubmit={handleSubmit}
    >
      <div className="field">
        <label className="label">Email</label>
        <div className={'control has-icons-right has-icons-left'}>
          <input
            name={'email'}
            className={'input is-danger'}
            value={fields.email}
            placeholder={'user@domain.com'}
            type={'text'}
            onChange={handleInputChange}
          />
          <span className="icon is-left">
            <FontAwesomeIcon icon={faEnvelope} />
          </span>
          <span className="icon is-right">
            <FontAwesomeIcon icon={faExclamationTriangle} />
          </span>
        </div>
        <p className="help is-danger">{<>&nbsp;</>}</p>
      </div>

      <div className="field">
        <label className="label">Password</label>
        <div className={'control has-icons-right has-icons-left'}>
          <input
            name={'password'}
            className={'input is-danger'}
            value={fields.password}
            type={'password'}
            onChange={handleInputChange}
          />
          <span className="icon is-left">
            <FontAwesomeIcon icon={faLock} />
          </span>
          <span className="icon is-right">
            <FontAwesomeIcon icon={faExclamationTriangle} />
          </span>
        </div>
        <p className="help is-danger">{<>&nbsp;</>}</p>
      </div>

      <div className="field is-grouped">
        <p className="control is-expanded">
          <button
            className={'button is-primary'}
            style={{ width: '100%' }}
            type="submit"
          >
            Submit
          </button>
        </p>
        <p className="control is-expanded">
          <button
            type="button"
            className="button is-light"
            style={{ width: '100%' }}
          >
            Cancel
          </button>
        </p>
      </div>
    </form>
  );
};

export default Form;
