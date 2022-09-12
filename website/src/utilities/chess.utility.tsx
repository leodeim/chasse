
export function customPieces() {
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
