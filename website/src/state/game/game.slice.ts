import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Orientation, START_POSITION, PositionObject, START_POSITION_OBJECT } from '../../utilities/chess.utility';
import { peek, push, removeLast } from '../../utilities/stack.utility';
import { getWindowProperties, WindowProperties } from '../../utilities/window.utility';
import { RootState } from '../store';

export type MoveItem = {
    position: PositionObject
    sessionId: string
}
interface GameState {
    gameFen: PositionObject,
    history: PositionObject[],
    boardOrientation: Orientation,
    windowProperties: WindowProperties,
    sessionId: string,
    loading: boolean,
    wsState: boolean
}

const initialState: GameState = {
    gameFen: START_POSITION_OBJECT,
    history: [START_POSITION_OBJECT],
    boardOrientation: Orientation.white,
    windowProperties: getWindowProperties(),
    sessionId: 'a88f494f-50f5-475e-98cf-2b0d9e2f05f4',
    loading: false,
    wsState: false
};

export const gameSlice = createSlice({
    name: 'game',
    initialState,
    reducers: {
        makeMove(state, _: PayloadAction<MoveItem>) {
            state.loading = true;
        },
        makeMoveSuccessful(state, _: PayloadAction<MoveItem>) {
            state.loading = false;
            // let newPosition = action.payload.position;
            // state.gameFen = newPosition;
            // if (newPosition !== peek(state.history)) {
            //     state.history = push(state.history, newPosition)
            // }
        },
        historyPop(state) {
            state.history = removeLast(state.history);
        },
        updatePosition(state, action: PayloadAction<string>) {
            let obj = JSON.parse(action.payload)
            state.gameFen = obj;
            if (obj !== peek(state.history)) {
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
    },
});

export const {
    makeMove,
    reverseBoard,
    updateWindowProperties,
    updatePosition,
    historyPop,
    updateWsState,
    makeMoveSuccessful } = gameSlice.actions;

export const selectGameFen = (state: RootState) => state.game.gameFen
export const selectBoardOrientation = (state: RootState) => state.game.boardOrientation;
export const selectWindowMinDimension = (state: RootState) => state.game.windowProperties.minDimension;
export const selectWindowPosition = (state: RootState) => state.game.windowProperties.position;
export const selectSessionId = (state: RootState) => state.game.sessionId;
export const selectHistory = (state: RootState) => state.game.history;
export const selectWsState = (state: RootState) => state.game.wsState;

export default gameSlice.reducer;