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
