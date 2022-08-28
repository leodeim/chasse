import { useEffect, useState } from "react";
import { Chess, ShortMove, Square } from "chess.js";
import { Chessboard } from "react-chessboard";

export default function Board() {
    const [windowDimensions, setWindowDimensions] = useState(getWindowMinDimension());
    const [game, setGame] = useState(new Chess());

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

    console.log("Board render")

    return (
        <div style={{ border: "5px solid #F0D9B5" }}>
            <Chessboard
                position={game.fen()}
                onPieceDrop={onDrop}
                boardWidth={windowDimensions * 0.8}
                customDarkSquareStyle={{ backgroundColor: '#8ba28c' }}
                customPieces={customPieces()}
            />
        </div>
    );
}

function getWindowMinDimension(): number {
    const { innerWidth: width, innerHeight: height } = window;
    return Math.min(width, height);
}

function customPieces() {
    const pieces = ['wP', 'wN', 'wB', 'wR', 'wQ', 'wK', 'bP', 'bN', 'bB', 'bR', 'bQ', 'bK'];
    const returnPieces: any[string] = [];
    pieces.map((p) => {
        returnPieces[p] = ({ squareWidth } : { squareWidth: number }) => (
            <div
                style={{
                    width: squareWidth,
                    height: squareWidth,
                    backgroundImage: `url(/media/${p}.png)`,
                    backgroundSize: '100%',
                }}
            />
        );
        return null;
    });
    return returnPieces;
};