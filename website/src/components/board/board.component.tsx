import Chessboard from "../../lib/Chessboard";
import { customPieces, Piece, Square } from "../../utilities/chess.utility";
import { useDispatch, useSelector } from "react-redux";
import { makeMove, selectBoardOrientation, selectGameFen, selectSessionId, selectWindowMinDimension, selectWsState, updatePosition } from "../../state/game/game.slice";
import { SendWebsocketJoinRoom } from "../../socket/socket";
import { useEffect } from "react";
import { objectTraps } from "immer/dist/internal";


export default function GameBoard(props: any) {
    const dispatch = useDispatch();
    const game = useSelector(selectGameFen);
    const boardOrientation = useSelector(selectBoardOrientation);
    const windowMinDimensions = useSelector(selectWindowMinDimension);
    const sessionId = useSelector(selectSessionId);
    const wsState = useSelector(selectWsState)

    useEffect(() => {
        if (wsState === true) {
            SendWebsocketJoinRoom(sessionId)
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [wsState]);

    function onDrop(obj: { sourceSquare: Square, targetSquare: Square, piece: Piece }) {
        let newGame = {
            ...game
        };

        if (obj.sourceSquare === obj.targetSquare && newGame[obj.targetSquare] === obj.piece) {
            return true;
        }
        if (obj.targetSquare !== 'offBoard') {
            newGame[obj.targetSquare] = obj.piece;
        }
        delete newGame[obj.sourceSquare]
        
        dispatch(makeMove({
            position: newGame,
            sessionId: sessionId
        }));

        return true;
    }

    return (
        <div>
            <Chessboard
                position={game}
                onDrop={onDrop}
                orientation={boardOrientation}
                width={windowMinDimensions * 0.6}
                darkSquareStyle={{ backgroundColor: '' }}
                pieces={customPieces()}
                sparePieces={true}
                dropOffBoard={'trash'}
                boardStyle={{
                    borderRadius: "5px",
                    boxShadow: `0 5px 15px rgba(0, 0, 0, 0.5)`
                  }}
            />
        </div>
    );
}
