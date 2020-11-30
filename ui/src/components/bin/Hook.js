import React from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrashAlt } from '@fortawesome/free-solid-svg-icons';

const Hook = ({ id }) => {
  return (
    <tr>
      <td>
        <button onClick={() => {}} className="button is-link is-light">
          {id}
        </button>
      </td>
      <td>
        <button className="button is-danger" onClick={() => {}}>
          <FontAwesomeIcon icon={faTrashAlt} />
        </button>
      </td>
    </tr>
  );
};

export default Hook;
