import { Square } from "chess.js";
import { Chessboard } from "react-chessboard";
import { customPieces } from "../../utilities/chess.utility";
import { useDispatch, useSelector } from "react-redux";
import { makeMove, selectBoardOrientation, selectGameFen, selectWindowMinDimension } from "../../state/game/game.slice";


export default function Board(props: any) {
    const dispatch = useDispatch();
    const game = useSelector(selectGameFen);
    const boardOrientation = useSelector(selectBoardOrientation);
    const windowMinDimensions = useSelector(selectWindowMinDimension);

    function onDrop(sourceSquare: Square, targetSquare: Square) {
        dispatch(makeMove({
            from: sourceSquare,
            to: targetSquare,
            promotion: 'q'
        }))

        return true;
    }

    return (
        <div className="border-8 border-solid border-yellow">
            <Chessboard
                position={game}
                onPieceDrop={onDrop}
                boardOrientation={boardOrientation}
                boardWidth={windowMinDimensions * 0.8}
                customDarkSquareStyle={{ backgroundColor: '' }}
                customPieces={customPieces()}
            />
        </div>
    );
}
