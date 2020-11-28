import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';

import {
  getBins,
  selectBins,
  createBin,
  selectBinsLoadingInProgress,
} from 'features/bins/binsSlice';

import Bin from './Bin';

const BinList = () => {
  const dispatch = useDispatch();
  const bins = useSelector(selectBins);
  const binLoadingInProgress = useSelector(selectBinsLoadingInProgress);

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
            <Bin key={bin.ID} id={bin.ID} created={bin.Created} />
          ))}
        </tbody>
      </table>
    </>
  );
};

export default BinList;
