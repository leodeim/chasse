export function getApiUrl(): string {
    const DEV_MODE = process.env.REACT_APP_DEV_MODE;
    if (DEV_MODE !== undefined) {
        return "http://localhost:8085/"
    }

    return "https://chasse.fun/"
}

export function getWebsocketUrl(): string {
    const DEV_MODE = process.env.REACT_APP_DEV_MODE;
    if (DEV_MODE !== undefined) {
        return "ws://localhost:8085/"
    }

    return "wss://chasse.fun/"
}
