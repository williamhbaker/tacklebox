import axios from 'axios';

const URL = process.env.REACT_APP_ENDPOINT;

export const login = async (data) => {
  // let json = JSON.stringify(data);

  // let init = {
  //   method: 'POST',
  //   headers: { 'Content-Type': 'application/json' },
  //   credentials: 'include',
  //   body: json,
  // };

  // let response = await fetch(`${URL}/login`, init);

  // let response = await axios.get(`${URL}`);

  let response = await axios({
    method: 'post',
    url: `${URL}/login`,
    headers: { 'Content-Type': 'application/json' },
    // withCredentials: true,
    data,
  });

  console.log(response);
};
