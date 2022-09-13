import { useEffect, useState } from "react";
import { Chess, ShortMove, Square } from "chess.js";
import { Chessboard } from "react-chessboard";
import { getWindowMinDimension } from "../../utilities/window.utility";
import { customPieces } from "../../utilities/chess.utility";
import Controls from "../controls/controls.component";

enum Orientation {
    white = "white",
    black = "black",
}
  
export default function Board() {
    const [windowDimensions, setWindowDimensions] = useState(getWindowMinDimension());
    const [game, setGame] = useState(new Chess());
    const [orientation, setOrientation] = useState(Orientation.white);

    useEffect(() => {
        function handleResize() {
            setWindowDimensions(getWindowMinDimension());
        }

        window.addEventListener('resize', handleResize);
        return () => window.removeEventListener('resize', handleResize);
    }, []);

    function makeAMove(move: ShortMove | string) {
        const gameCopy = { ...game };
        const result = gameCopy.move(move);
        setGame(gameCopy);
        return result; // null if the move was illegal, the move object if the move was legal
    }

    function onDrop(sourceSquare: Square, targetSquare: Square) {
        const move = makeAMove({
            from: sourceSquare,
            to: targetSquare,
            promotion: 'q' // always promote to a queen for example simplicity
        });

        // illegal move
        if (move === null) return false;

        return true;
    }

    function goBack() {
        const gameCopy = { ...game };
        gameCopy.undo()
        setGame(gameCopy);
    }

    function resetBoard() {
        const gameCopy = { ...game };
        gameCopy.reset()
        setGame(gameCopy);
    }

    function reverseBoard() {
        if (orientation === Orientation.white) {
            setOrientation(Orientation.black)
        } else {
            setOrientation(Orientation.white)
        }
    }

    return (
        <div>
            <div className="border-8 border-solid border-zensquare">
                <Chessboard
                    position={game.fen()}
                    onPieceDrop={onDrop}
                    boardOrientation={orientation}
                    boardWidth={windowDimensions * 0.8}
                    customDarkSquareStyle={{ backgroundColor: '' }}
                    customPieces={customPieces()}
                />
            </div>
            <div className="flex justify-center">
                <Controls back={goBack} reverse={reverseBoard} reset={resetBoard} />
            </div>
        </div>
    );
}
