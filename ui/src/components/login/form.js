import React, { useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faEnvelope,
  faLock,
  faExclamationTriangle,
} from '@fortawesome/free-solid-svg-icons';

import {
  login,
  selectLoginInProgress,
  selectUser,
} from 'features/user/userSlice';

const Form = () => {
  const history = useHistory();
  const dispatch = useDispatch();
  const loginInProgress = useSelector(selectLoginInProgress);
  const user = useSelector(selectUser);

  const [fields, setFields] = useState({ email: '', password: '' });
  const [triedLogin, setTriedLogin] = useState(false);

  if (user) history.push('/');

  const handleInputChange = (e) => {
    setFields({
      ...fields,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    setTriedLogin(true);
    dispatch(login(fields));
  };

  const hasError = triedLogin && !user && !loginInProgress;

  return (
    <form
      style={{ maxWidth: '30rem', margin: '0 auto' }}
      onSubmit={handleSubmit}
    >
      <fieldset disabled={loginInProgress}>
        <div className="field">
          <label className="label">Email</label>
          <div className={'control has-icons-right has-icons-left'}>
            <input
              name={'email'}
              className={`input ${hasError && 'is-danger'}`}
              value={fields.email}
              placeholder={'user@domain.com'}
              type={'text'}
              onChange={handleInputChange}
            />
            <span className="icon is-left">
              <FontAwesomeIcon icon={faEnvelope} />
            </span>
            <span className="icon is-right">
              {hasError && <FontAwesomeIcon icon={faExclamationTriangle} />}
            </span>
          </div>
          <p className="help is-danger">
            {(hasError && 'invalid credentials') || <>&nbsp;</>}
          </p>
        </div>

        <div className="field">
          <label className="label">Password</label>
          <div className={'control has-icons-right has-icons-left'}>
            <input
              name={'password'}
              className={`input ${hasError && 'is-danger'}`}
              value={fields.password}
              type={'password'}
              onChange={handleInputChange}
            />
            <span className="icon is-left">
              <FontAwesomeIcon icon={faLock} />
            </span>
            <span className="icon is-right">
              {hasError && <FontAwesomeIcon icon={faExclamationTriangle} />}
            </span>
          </div>
          <p className="help is-danger">
            {(hasError && 'invalid credentials') || <>&nbsp;</>}
          </p>
        </div>

        <div className="field is-grouped">
          <p className="control is-expanded">
            <button
              className={`button is-primary ${loginInProgress && 'is-loading'}`}
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
