enum Position {
    h = 'horizontal',
    v = 'vertical',
};

export type WindowProperties = {
    minDimension: number,
    position: Position
};

function getWindowProperties(): WindowProperties {
    const { innerWidth: width, innerHeight: height } = window;

    return {
        minDimension: Math.min(width, height),
        position: (width > height) ? Position.h : Position.v
    };
};

export { getWindowProperties, Position };