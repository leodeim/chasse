import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Orientation, PositionObject, START_POSITION_OBJECT } from '../../utilities/chess.utility';
import { peek, push, removeLast } from '../../utilities/stack.utility';
import { getWindowProperties, WindowProperties } from '../../utilities/window.utility';
import { RootState } from '../store';
import isEqual from 'lodash.isequal';

export type MoveItem = {
    position: PositionObject
    sessionId: string
}
interface GameState {
    gamePosition: PositionObject,
    history: PositionObject[],
    boardOrientation: Orientation,
    windowProperties: WindowProperties,
    sessionId: string,
    loading: boolean,
    wsState: boolean
}

const initialState: GameState = {
    gamePosition: START_POSITION_OBJECT,
    history: [START_POSITION_OBJECT],
    boardOrientation: Orientation.white,
    windowProperties: getWindowProperties(),
    sessionId: '',
    loading: false,
    wsState: false
};

export const gameSlice = createSlice({
    name: 'game',
    initialState,
    reducers: {
        makeMove(state, action: PayloadAction<MoveItem>) {
            state.loading = true;
            let obj = action.payload.position;
            let lastObj = peek(state.history);
            state.gamePosition = obj;
            if (!isEqual(obj, lastObj)) {
                state.history = push(state.history, obj)
            }
        },
        makeMoveSuccessful(state, _: PayloadAction<MoveItem>) {
            state.loading = false;
        },
        historyPop(state) {
            state.history = removeLast(state.history);
        },
        updatePosition(state, action: PayloadAction<string>) {
            let obj = JSON.parse(action.payload);
            let lastObj = peek(state.history);
            state.gamePosition = obj;
            if (!isEqual(obj, lastObj)) {
                state.history = push(state.history, obj)
            }
        },
        reverseBoard(state) {
            state.boardOrientation =
                (state.boardOrientation === Orientation.white) ? Orientation.black : Orientation.white;
        },
        updateWindowProperties(state) {
            state.windowProperties = getWindowProperties();
        },
        updateWsState(state, action: PayloadAction<boolean>) {
            state.wsState = action.payload;
        },
        updateSessionId(state, action: PayloadAction<string>) {
            state.sessionId = action.payload;
        },
    },
});

export const {
    makeMove,
    reverseBoard,
    updateWindowProperties,
    updatePosition,
    historyPop,
    updateWsState,
    makeMoveSuccessful,
    updateSessionId} = gameSlice.actions;

export const selectGamePosition = (state: RootState) => state.game.gamePosition
export const selectBoardOrientation = (state: RootState) => state.game.boardOrientation;
export const selectWindowMinDimension = (state: RootState) => state.game.windowProperties.minDimension;
export const selectWindowPosition = (state: RootState) => state.game.windowProperties.position;
export const selectSessionId = (state: RootState) => state.game.sessionId;
export const selectHistory = (state: RootState) => state.game.history;
export const selectWsState = (state: RootState) => state.game.wsState;

export default gameSlice.reducer;