import React, { useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { signUp } from 'features/user/userSlice';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faEnvelope,
  faLock,
  faExclamationTriangle,
} from '@fortawesome/free-solid-svg-icons';

import { selectLoginInProgress, setMessage } from 'features/user/userSlice';

const Form = () => {
  const dispatch = useDispatch();
  const history = useHistory();
  const inProgress = useSelector(selectLoginInProgress);

  const [fields, setFields] = useState({ email: '', password: '' });
  const [errs, setErrs] = useState({ email: false, password: false });

  const handleInputChange = (e) => {
    setFields({
      ...fields,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (validateInputs()) {
      const res = await dispatch(signUp(fields));
      if (res.payload && res.payload.message) {
        dispatch(setMessage('Sign up successful - please log in!'));
        history.push('/login');
      }
    }
  };

  const validateInputs = () => {
    const validEmail = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/.test(
      fields.email
    );

    const validPassword = fields.password.length >= 6;

    setErrs({
      email: !validEmail,
      password: !validPassword,
    });

    return validEmail && validPassword;
  };

  return (
    <form
      style={{ maxWidth: '30rem', margin: '0 auto' }}
      onSubmit={handleSubmit}
    >
      <fieldset disabled={false}>
        <div className="field">
          <label className="label">Email</label>
          <div className={'control has-icons-right has-icons-left'}>
            <input
              name={'email'}
              className={`input ${errs.email && 'is-danger'}`}
              value={fields.email}
              placeholder={'user@domain.com'}
              type={'text'}
              onChange={handleInputChange}
            />
            <span className="icon is-left">
              <FontAwesomeIcon icon={faEnvelope} />
            </span>
            <span className="icon is-right">
              {errs.email && <FontAwesomeIcon icon={faExclamationTriangle} />}
            </span>
          </div>
          <p className="help is-danger">
            {(errs.email && 'invalid email') || <>&nbsp;</>}
          </p>
        </div>

        <div className="field">
          <label className="label">Password</label>
          <div className={'control has-icons-right has-icons-left'}>
            <input
              name={'password'}
              className={`input ${errs.password && 'is-danger'}`}
              value={fields.password}
              type={'password'}
              onChange={handleInputChange}
            />
            <span className="icon is-left">
              <FontAwesomeIcon icon={faLock} />
            </span>
            <span className="icon is-right">
              {errs.password && (
                <FontAwesomeIcon icon={faExclamationTriangle} />
              )}
            </span>
          </div>
          <p className="help is-danger">
            {(errs.password &&
              'password must be at least 6 characters long') || <>&nbsp;</>}
          </p>
        </div>

        <div className="field is-grouped">
          <p className="control is-expanded">
            <button
              className={`button is-primary ${inProgress && 'is-loading'}`}
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
      </fieldset>
    </form>
  );
};

export default Form;
