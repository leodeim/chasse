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

export function getAppVersion(): string {
    const VERSION = process.env.REACT_APP_VERSION;
    if (VERSION === undefined) {
        return "no version"
    }

    return VERSION
}

export function getDevMode(): boolean {
    const DEV_MODE = process.env.REACT_APP_DEV_MODE;
    if (DEV_MODE === undefined) {
        return false
    }

    return true
}
