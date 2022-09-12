
export function getWindowMinDimension(): number {
    const { innerWidth: width, innerHeight: height } = window;
    return Math.min(width, height);
}