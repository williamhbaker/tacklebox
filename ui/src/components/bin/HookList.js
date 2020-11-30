import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';

import { selectHooks } from 'features/hooks/hooksSlice';
import { activateBin } from 'features/hooks/hooksSlice';

import Hook from './Hook';

const Hooklist = ({ id }) => {
  const dispatch = useDispatch();
  const hooks = useSelector(selectHooks);

  useEffect(() => {
    dispatch(activateBin(id));
  }, [dispatch, id]);

  return (
    <table>
      <tbody>
        {hooks.map((hook) => {
          return <Hook key={hook.ID} id={hook.ID} />;
        })}
      </tbody>
    </table>
  );
};

export default Hooklist;
