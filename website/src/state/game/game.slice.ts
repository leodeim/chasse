import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Chess, ShortMove } from 'chess.js';
import { Orientation, START_POSITION } from '../../utilities/chess.utility';
import { getWindowProperties, WindowProperties } from '../../utilities/window.utility';
import { RootState } from '../store';

interface GameState {
    gameFen: string,
    boardOrientation: Orientation,
    windowProperties: WindowProperties
}

const initialState: GameState = {
    gameFen: START_POSITION,
    boardOrientation: Orientation.white,
    windowProperties: getWindowProperties()
};

export const gameSlice = createSlice({
    name: 'game',
    initialState,
    reducers: {
        makeMove(state, action: PayloadAction<string>) {
            state.gameFen = action.payload
        },
        goBack(state) {
        },
        resetBoard(state) {
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