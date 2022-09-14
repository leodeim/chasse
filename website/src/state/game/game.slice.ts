import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Chess, ShortMove } from 'chess.js';
import { Orientation, START_POSITION } from '../../utilities/chess.utility';
import { getWindowProperties, WindowProperties } from '../../utilities/window.utility';
import { RootState } from '../store';

interface GameState {
    gameFenArray: any[],
    boardOrientation: Orientation,
    windowProperties: WindowProperties
}

const initialState: GameState = {
    gameFenArray: [START_POSITION],
    boardOrientation: Orientation.white,
    windowProperties: getWindowProperties()
};

export const gameSlice = createSlice({
    name: 'game',
    initialState,
    reducers: {
        makeMove(state, action: PayloadAction<ShortMove | string>) {
            let game = new Chess()
            game.load(state.gameFenArray.slice(-1)[0])
            const result = game.move(action.payload);

            if (result !== null) {
                state.gameFenArray.push(game.fen())
            }
        },
        goBack(state) {
            state.gameFenArray.pop()
        },
        resetBoard(state) {
            state.gameFenArray.push(START_POSITION)
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

export const selectGameFen = (state: RootState) => state.game.gameFenArray.slice(-1)[0]
export const selectBoardOrientation = (state: RootState) => state.game.boardOrientation;
export const selectWindowMinDimension = (state: RootState) => state.game.windowProperties.minDimension;
export const selectWindowPosition = (state: RootState) => state.game.windowProperties.position;

export default gameSlice.reducer;