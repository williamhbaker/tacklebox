const URL = process.env.REACT_APP_ENDPOINT;

export const login = async (data) => {
  let json = JSON.stringify(data);

  let init = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: json,
  };

  const response = await fetch(`${URL}/login`, init);
  return response.ok ? response.json() : null;
};

export const logout = async () => {
  let init = {
    method: 'POST',
    credentials: 'include',
  };

  const response = await fetch(`${URL}/logout`, init);
  return response.ok;
};

export const checkStatus = async () => {
  let init = {
    method: 'GET',
    credentials: 'include',
  };

  const response = await fetch(`${URL}/user`, init);
  return response.ok ? response.json() : null;
};

export const getBins = async () => {
  let init = {
    method: 'GET',
    credentials: 'include',
  };

  const response = await fetch(`${URL}/user/bins`, init);
  return response.ok ? response.json() : null;
};

export const createBin = async (data) => {
  let init = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
  };

  const response = await fetch(`${URL}/bin`, init);
  return response.ok ? response.json() : null;
};

export const destroyBin = async (binID) => {
  let init = {
    method: 'DELETE',
    credentials: 'include',
  };

  const response = await fetch(`${URL}/bin/${binID}`, init);
  return response.ok ? response.json() : null;
};

export const getHooks = async (binID) => {
  let init = {
    method: 'GET',
    credentials: 'include',
  };

  const response = await fetch(`${URL}/bin/${binID}`, init);
  return response.ok ? response.json() : null;
};
