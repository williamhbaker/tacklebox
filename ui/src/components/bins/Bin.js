import React from 'react';
import { useDispatch } from 'react-redux';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrashAlt } from '@fortawesome/free-solid-svg-icons';

import { destroyBin } from 'features/bins/binsSlice';
import { getHooks } from 'features/hooks/hooksSlice';

const Bin = ({ id, created }) => {
  const dispatch = useDispatch();

  const handleDestroyClick = (e, id) => {
    e.preventDefault();
    dispatch(destroyBin(id));
  };

  const handleBinClick = (e, id) => {
    e.preventDefault();
    dispatch(getHooks(id));
  };

  return (
    <tr>
      <td>
        <button
          onClick={(e) => handleBinClick(e, id)}
          className="button is-link is-light"
        >
          {id}
        </button>
      </td>
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
