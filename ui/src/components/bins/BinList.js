import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';

import {
  getBins,
  selectBins,
  createBin,
  selectBinsLoadingInProgress,
  selectActiveBin,
} from 'features/bins/binsSlice';

import { selectHooksInProgress } from 'features/hooks/hooksSlice';

import Bin from './Bin';

const BinList = () => {
  const dispatch = useDispatch();
  const bins = useSelector(selectBins);
  const activeBin = useSelector(selectActiveBin);
  const binLoadingInProgress = useSelector(selectBinsLoadingInProgress);
  const hookLoadingInProgress = useSelector(selectHooksInProgress);

  useEffect(() => {
    dispatch(getBins());
  }, [dispatch]);

  const handleCreateNewBin = (e) => {
    e.preventDefault();
    dispatch(createBin());
  };

  const sortedBins = bins.slice().sort((a, b) => {
    const aDate = new Date(a.Created);
    const bDate = new Date(b.Created);
    return bDate - aDate;
  });

  return (
    <>
      <button
        className={`button mb-3 is-info ${
          binLoadingInProgress && 'is-loading'
        }`}
        onClick={handleCreateNewBin}
      >
        Create New Bin
      </button>
      <br />
      <table className="table">
        <tbody>
          {sortedBins.map((bin) => (
            <Bin
              key={bin.ID}
              id={bin.ID}
              created={bin.Created}
              active={bin.ID === activeBin}
              inProgress={hookLoadingInProgress}
            />
          ))}
        </tbody>
      </table>
    </>
  );
};

export default BinList;
