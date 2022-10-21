const KEY_RECENT = 'recent'

export const getRecentSession = (): null | string => {
    if (!localStorage) return null
    try {
        return localStorage.getItem(KEY_RECENT)
    } catch (error) {
        console.error(`Error getting ${KEY_RECENT} from localStorage`, error)
    }
    return null
}

export const storeSession = (sessionId: string) => {
    if (!localStorage) return
    try {
        localStorage.setItem(KEY_RECENT, sessionId)
    } catch (error) {
        console.error(`Error storing ${KEY_RECENT} to localStorage`, error)
    }
}

export const clearRecentData = () => {
    if (!localStorage) return
    localStorage.removeItem(KEY_RECENT)
}
