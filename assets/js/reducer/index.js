import { combineReducers, applyMiddleware, createStore } from 'redux';
import thunk from 'redux-thunk';
import Logger from 'redux-logger';
import promise from 'redux-promise-middleware';
import SkillReducer from './skill';
import ProjectReducer from './project';
import ExperienceReducer from './experience';
import EducationReducer from './education';

const reducers = combineReducers({
    SkillReducer,
    ProjectReducer,
    ExperienceReducer,
    EducationReducer
});

const middleware = applyMiddleware(thunk, Logger, promise);

const store = createStore(reducers, middleware);

export default store;
