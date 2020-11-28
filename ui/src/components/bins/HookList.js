import React from 'react';
import { useSelector } from 'react-redux';

import { selectHooks } from 'features/hooks/hooksSlice';

import Hook from './Hook';

const Hooklist = () => {
  const hooks = useSelector(selectHooks);

  return (
    <table>
      <tbody>
        {hooks.map((hook) => {
          console.log(hook.ID);
          // return <p>{hook.ID}</p>;
          return <Hook key={hook.ID} id={hook.ID} />;
        })}
      </tbody>
    </table>
  );
};

export default Hooklist;
