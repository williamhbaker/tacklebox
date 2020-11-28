import React from 'react';

const Bin = ({ id, created }) => (
  <tr>
    <td>
      {id} - {created}
    </td>
  </tr>
);

export default Bin;
