export function getApiUrl(): string {
    const CUSTOM_URL = process.env.REACT_APP_API_URL;
    if (CUSTOM_URL === undefined) {
        return "http://localhost:8085/"
    }
    
    return CUSTOM_URL
}

export function getWebsocketUrl(): string {
    const CUSTOM_URL = process.env.REACT_APP_WS_URL;
    if (CUSTOM_URL === undefined) {
        return "ws://localhost:8085/"
    }

    return CUSTOM_URL
}
