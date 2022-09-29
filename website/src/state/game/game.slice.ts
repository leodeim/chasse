import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Orientation, START_POSITION } from '../../utilities/chess.utility';
import { IStack, Stack } from '../../utilities/stack.utility';
import { getWindowProperties, WindowProperties } from '../../utilities/window.utility';
import { RootState } from '../store';

interface GameState {
    gameFen: string,
    history: IStack<string>,
    boardOrientation: Orientation,
    windowProperties: WindowProperties
}

const initialState: GameState = {
    gameFen: START_POSITION,
    history: new Stack(100),
    boardOrientation: Orientation.white,
    windowProperties: getWindowProperties()
};

export const gameSlice = createSlice({
    name: 'game',
    initialState,
    reducers: {
        makeMove(state, action: PayloadAction<string>) {
            state.gameFen = action.payload;
            state.history.push(action.payload);
        },
        goBack(state) {
            state.history.pop();
            let newFen = state.history.pop();
            if (newFen !== undefined) makeMove(newFen);
        },
        resetBoard() {
            makeMove(START_POSITION);
        },
        reverseBoard(state) {
            state.boardOrientation =
                (state.boardOrientation === Orientation.white) ? Orientation.black : Orientation.white;
        },
        updateWindowProperties(state) {
            state.windowProperties = getWindowProperties();
        },
    },
});

export const {
    makeMove,
    goBack,
    resetBoard,
    reverseBoard,
    updateWindowProperties } = gameSlice.actions;

export const selectGameFen = (state: RootState) => state.game.gameFen
export const selectBoardOrientation = (state: RootState) => state.game.boardOrientation;
export const selectWindowMinDimension = (state: RootState) => state.game.windowProperties.minDimension;
export const selectWindowPosition = (state: RootState) => state.game.windowProperties.position;

export default gameSlice.reducer;