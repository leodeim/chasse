import { Square } from "chess.js";
import { Chessboard } from "react-chessboard";
import { customPieces } from "../../utilities/chess.utility";
import { useDispatch, useSelector } from "react-redux";
import { makeMove, selectBoardOrientation, selectGameFen, selectWindowMinDimension } from "../../state/game/game.slice";
import { client, w3cwebsocket as W3CWebSocket } from "websocket";
import { useEffect } from "react";

let wsClient = new W3CWebSocket('ws://127.0.0.1:8085/ws/123');

export default function Board(props: any) {
    const dispatch = useDispatch();
    const game = useSelector(selectGameFen);
    const boardOrientation = useSelector(selectBoardOrientation);
    const windowMinDimensions = useSelector(selectWindowMinDimension);

    useEffect(() => {
        wsClient.onopen = () => {
            console.log('WebSocket Client Connected');
        };
        wsClient.onmessage = (message) => {
            console.log(message);
        };
    }, []);

    function onDrop(sourceSquare: Square, targetSquare: Square) {
        dispatch(makeMove({
            from: sourceSquare,
            to: targetSquare,
            promotion: 'q'
        }))

        wsClient.send(JSON.stringify({
            sessionId: "sadsadsa",
            position: "userevent"
        }))

        console.log(wsClient.readyState)

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
