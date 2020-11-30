import React from 'react';
import { useDispatch } from 'react-redux';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrashAlt } from '@fortawesome/free-solid-svg-icons';

import { activateHook, deleteHook } from 'features/hooks/hooksSlice';

const HookLink = ({ id, active }) => {
  const dispatch = useDispatch();

  const handleHookClick = (e) => {
    e.preventDefault();
    dispatch(activateHook(id));
  };

  const handleDeleteClick = (e) => {
    e.preventDefault();
    dispatch(deleteHook(id));
  };

  return (
    <tr>
      <td>
        <button
          onClick={handleHookClick}
          className={`button ${active ? 'is-primary' : 'is-link is-light'}`}
        >
          {id}
        </button>
      </td>
      <td>
        <button className="button is-danger" onClick={handleDeleteClick}>
          <FontAwesomeIcon icon={faTrashAlt} />
        </button>
      </td>
    </tr>
  );
};

export default HookLink;
