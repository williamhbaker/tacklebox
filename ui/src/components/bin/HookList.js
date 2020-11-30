import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';

import {
  selectHooks,
  activateBin,
  selectActiveHook,
} from 'features/hooks/hooksSlice';

import HookLink from './HookLink';

const Hooklist = ({ id }) => {
  const dispatch = useDispatch();
  const hooks = useSelector(selectHooks);
  const activeHook = useSelector(selectActiveHook);

  useEffect(() => {
    dispatch(activateBin(id));
  }, [dispatch, id]);

  return (
    <table className="table">
      <tbody>
        {hooks.map((hook) => {
          return <HookLink key={hook} id={hook} active={activeHook === hook} />;
        })}
      </tbody>
    </table>
  );
};

export default Hooklist;
