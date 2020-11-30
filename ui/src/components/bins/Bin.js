import React from 'react';
import { useDispatch } from 'react-redux';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrashAlt } from '@fortawesome/free-solid-svg-icons';
import { useHistory } from 'react-router-dom';
import { destroyBin } from 'features/bins/binsSlice';

const Bin = ({ id, created, active, inProgress }) => {
  const dispatch = useDispatch();
  const history = useHistory();

  const handleDestroyClick = (e, id) => {
    e.preventDefault();
    dispatch(destroyBin(id));
  };

  const handleBinClick = async (e, id) => {
    e.preventDefault();
    history.push(`/bin/${id}`);
  };

  return (
    <tr>
      <td>
        <button
          onClick={(e) => handleBinClick(e, id)}
          className={`button 'is-link is-light`}
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
