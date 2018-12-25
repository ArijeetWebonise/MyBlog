import { combineReducers } from 'redux';
import ping from './ping';
import login from './login';

const reducers = {
  ping,
  login,
};

export default combineReducers(reducers);
