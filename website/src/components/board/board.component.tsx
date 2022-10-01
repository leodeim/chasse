import Chessboard from "../../lib/Chessboard";
import { customPieces, Piece, Square } from "../../utilities/chess.utility";
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

    function onDrop(obj: { sourceSquare: Square, targetSquare: Square, piece: Piece }) {
        console.log(obj)
        // TODO: move logic

        // if (result != null) {
        //     dispatch(makeMove({
        //         position: chess.fen(),
        //         sessionId: sessionId
        //     }));
        // }

        return true;
    }

    function getPositionObject(position: any) {
        console.log(position)
    }

    console.log(JSON.stringify(game))

    return (
        <div>
            <Chessboard
                position={game}
                onDrop={onDrop}
                getPosition={getPositionObject}
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
