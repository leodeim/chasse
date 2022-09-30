import { Chess, Square } from "chess.js";
import { Chessboard } from "react-chessboard";
import { customPieces } from "../../utilities/chess.utility";
import { useDispatch, useSelector } from "react-redux";
import { makeMove, selectBoardOrientation, selectGameFen, selectSessionId, selectWindowMinDimension } from "../../state/game/game.slice";
import { SendWebsocketJoinRoom } from "../../socket/socket";
import { useEffect } from "react";


export default function Board(props: any) {
    const dispatch = useDispatch();
    const game = useSelector(selectGameFen);
    const boardOrientation = useSelector(selectBoardOrientation);
    const windowMinDimensions = useSelector(selectWindowMinDimension);
    const sessionId = useSelector(selectSessionId);

    useEffect(() => {
        SendWebsocketJoinRoom(sessionId)
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    function onDrop(sourceSquare: Square, targetSquare: Square) {
        const chess = new Chess()
        if (game !== undefined) chess.load(game)
        let result = chess.move({
            from: sourceSquare,
            to: targetSquare,
            promotion: 'q'
        })

        if (result != null) {
            dispatch(makeMove({
                position: chess.fen(),
                sessionId: sessionId
            }));
        }

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
