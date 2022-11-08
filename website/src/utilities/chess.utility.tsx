export type Piece =
    'wP' | 'wN' | 'wB' | 'wR' | 'wQ' | 'wK' |
    'bP' | 'bN' | 'bB' | 'bR' | 'bQ' | 'bK';

export type Square =
    | 'a8' | 'b8' | 'c8' | 'd8' | 'e8' | 'f8' | 'g8' | 'h8'
    | 'a7' | 'b7' | 'c7' | 'd7' | 'e7' | 'f7' | 'g7' | 'h7'
    | 'a6' | 'b6' | 'c6' | 'd6' | 'e6' | 'f6' | 'g6' | 'h6'
    | 'a5' | 'b5' | 'c5' | 'd5' | 'e5' | 'f5' | 'g5' | 'h5'
    | 'a4' | 'b4' | 'c4' | 'd4' | 'e4' | 'f4' | 'g4' | 'h4'
    | 'a3' | 'b3' | 'c3' | 'd3' | 'e3' | 'f3' | 'g3' | 'h3'
    | 'a2' | 'b2' | 'c2' | 'd2' | 'e2' | 'f2' | 'g2' | 'h2'
    | 'a1' | 'b1' | 'c1' | 'd1' | 'e1' | 'f1' | 'g1' | 'h1'
    | 'offBoard' | 'spare';

export enum Orientation {
    white = "white",
    black = "black",
}

export function customPieces(orientation: Orientation, tabletMode: boolean) {
    const wPieces = ['wP', 'wN', 'wB', 'wR', 'wQ', 'wK'];
    const bPieces = ['bP', 'bN', 'bB', 'bR', 'bQ', 'bK'];
    const returnPieces: any[string] = [];

    let normalPiece = (piece: string, squareWidth: number) => {
        return (
            <div
            style={{
                width: squareWidth,
                height: squareWidth,
                backgroundImage: `url(/media/${piece}.png)`,
                backgroundSize: '100%',
            }}
            />
            )
        }
        
        let rotatedPiece = (piece: string, squareWidth: number) => {
        return (
            <div
                style={{
                    width: squareWidth,
                    height: squareWidth,
                    backgroundImage: `url(/media/${piece}.png)`,
                    backgroundSize: '100%',
                    transform: 'rotate(180deg)',
                }}
            />
        )
    }

    wPieces.map((p) => {
        let shouldRotate: boolean = tabletMode && orientation === Orientation.black
        if (shouldRotate) {
            returnPieces[p] = ({ squareWidth }: { squareWidth: number }) => {
                return rotatedPiece(p, squareWidth);
            };
        } else {
            returnPieces[p] = ({ squareWidth }: { squareWidth: number }) => {
                return normalPiece(p, squareWidth);
            };
        }
        return null;
    });
    bPieces.map((p) => {
        let shouldRotate: boolean = tabletMode && orientation === Orientation.white
        if (shouldRotate) {
            returnPieces[p] = ({ squareWidth }: { squareWidth: number }) => {
                return rotatedPiece(p, squareWidth);
            };
        } else {
            returnPieces[p] = ({ squareWidth }: { squareWidth: number }) => {
                return normalPiece(p, squareWidth);
            };
        }
        return null;
    });

    return returnPieces;
};

export function calculateMove(obj: { sourceSquare: Square, targetSquare: Square, piece: Piece }, currentPosition: PositionObject): PositionObject | null {
    if (obj.sourceSquare === obj.targetSquare && currentPosition[obj.targetSquare] === obj.piece) {
        return null;
    };
    if (obj.sourceSquare === 'spare' && obj.targetSquare === 'offBoard') {
        return null;
    };

    let newGamePosition = {
        ...currentPosition
    };

    if (obj.targetSquare !== 'offBoard') {
        newGamePosition[obj.targetSquare] = obj.piece;
    }
    delete newGamePosition[obj.sourceSquare];

    return newGamePosition;
}

export type PositionObject = {
    [key in Square]?: Piece;
};

export const START_POSITION_OBJECT: PositionObject = {
    'a8': 'bR',
    'b8': 'bN',
    'c8': 'bB',
    'd8': 'bQ',
    'e8': 'bK',
    'f8': 'bB',
    'g8': 'bN',
    'h8': 'bR',
    'a7': 'bP',
    'b7': 'bP',
    'c7': 'bP',
    'd7': 'bP',
    'e7': 'bP',
    'f7': 'bP',
    'g7': 'bP',
    'h7': 'bP',
    'd2': 'wP',
    'a2': 'wP',
    'b2': 'wP',
    'c2': 'wP',
    'e2': 'wP',
    'f2': 'wP',
    'g2': 'wP',
    'h2': 'wP',
    'a1': 'wR',
    'b1': 'wN',
    'c1': 'wB',
    'd1': 'wQ',
    'e1': 'wK',
    'f1': 'wB',
    'g1': 'wN',
    'h1': 'wR'
}