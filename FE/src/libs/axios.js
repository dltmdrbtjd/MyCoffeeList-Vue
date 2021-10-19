import axios from "axios";

const instance = axios.create({
  baseURL: "http://localhost:5000",
  withCredentials: false
});

const APIS = {
  // postCreate: (contents) => instance.post('/api/post', contents),
};

export { APIS };
