const URL = process.env.REACT_APP_ENDPOINT;

export const login = async (data) => {
  let json = JSON.stringify(data);

  let init = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: json,
  };

  let response = await fetch(`${URL}/login`, init);

  console.log(response);
};
