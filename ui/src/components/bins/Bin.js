import React from 'react';
import { useDispatch } from 'react-redux';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrashAlt } from '@fortawesome/free-solid-svg-icons';

import { destroyBin } from 'features/bins/binsSlice';

const Bin = ({ id, created }) => {
  const dispatch = useDispatch();

  const handleDestroyClick = (e, id) => {
    e.preventDefault();
    dispatch(destroyBin(id));
  };

  return (
    <tr>
      <td>{id}</td>
      <td>
        <button
          className="button is-danger"
          onClick={(e) => handleDestroyClick(e, id)}
        >
          <FontAwesomeIcon icon={faTrashAlt} />
        </button>
      </td>
    </tr>
  );
};

export default Bin;
