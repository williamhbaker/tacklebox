import React from 'react';
import { useSelector } from 'react-redux';

import { selectActiveHookDetails } from 'features/hooks/hooksSlice';

const HookDetails = () => {
  const details = useSelector(selectActiveHookDetails);

  return (
    <>
      {details && (
        <>
          <p className="has-text-weight-semibold mb-3">
            {new Date(details.created).toString()}
          </p>
          <p className="mb-3">
            <span className="has-text-weight-semibold">ID:</span> {details.id}
          </p>
          <pre>{details.content}</pre>
        </>
      )}
    </>
  );
};

export default HookDetails;
