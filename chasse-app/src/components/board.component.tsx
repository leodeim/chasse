import Chessboard from "@leonidasdeim/chessboardjsx";
import { calculateMove, customPieces, Piece, Square } from "../utilities/chess.utility";
import { makeMove, selectBoardOrientation, selectGamePosition, selectSessionId, selectTabletMode, selectWindowMinDimension } from "../state/game/game.slice";
import { useAppDispatch, useAppSelector } from "../state/hooks";


export default function GameBoard(props: any) {
    const dispatch = useAppDispatch();
    const gamePosition = useAppSelector(selectGamePosition);
    const boardOrientation = useAppSelector(selectBoardOrientation);
    const windowMinDimensions = useAppSelector(selectWindowMinDimension);
    const sessionId = useAppSelector(selectSessionId);
    const tabletMode = useAppSelector(selectTabletMode);


    let gamePositionCopy = {
        ...gamePosition
    };

    function onDrop(obj: { sourceSquare: Square, targetSquare: Square, piece: Piece }) {
        let newGamePosition = calculateMove(obj, gamePosition)

        if (newGamePosition !== null) {
            dispatch(makeMove({
                position: newGamePosition,
                sessionId: sessionId
            }));
        }

        return true;
    }

    return (
        <div>
            <Chessboard
                position={gamePositionCopy}
                onDrop={onDrop}
                orientation={boardOrientation}
                width={windowMinDimensions * 0.8}
                darkSquareStyle={{ backgroundColor: '' }}
                lightSquareStyle={{ backgroundColor: '#ede4e4' }}
                pieces={customPieces(boardOrientation, tabletMode)}
                sparePieces={true}
                dropOffBoard={'trash'}
                transitionDuration={200}
                showNotation={false}
                dropSquareStyle={{
                    boxShadow: 'inset 0 0 1px 10px #717d8f'
                }}
                boardStyle={{
                    boxShadow: '0 5px 30px rgba(0, 0, 0, 0.5)'
                  }}
            />
        </div>
    );
}
