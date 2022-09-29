import { combineReducers, configureStore } from '@reduxjs/toolkit';
import { combineEpics, createEpicMiddleware } from 'redux-observable';
import { makeMoveEpic } from './game/game.epic';
import gameReducer from './game/game.slice';

const rootEpic = combineEpics<any>(makeMoveEpic)

const rootReducer = combineReducers({
    game: gameReducer,
})

const dependencies = {
    // axios: axios,
}

const epicMiddleware = createEpicMiddleware({ dependencies })

export type EpicDependenciesType = typeof dependencies

export const store = configureStore({
    reducer: rootReducer,
    middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(epicMiddleware),
})

epicMiddleware.run(rootEpic)


export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
