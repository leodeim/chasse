import { w3cwebsocket } from "websocket";
import { getWebsocketUrl, getDevMode } from "../utilities/environment.utility";

export const wsClient = new w3cwebsocket(getWebsocketUrl() + 'api/ws');

export enum WebsocketAction {
    BLANK = 0,
    MOVE = 1,
	GO_BACK = 2,
	RESET = 3,
    JOIN_ROOM = 4,
    CONNECT = 5,
}

export enum WebsocketResponse {
    BLANK = 0,
    OK = 1,
    ERROR = 2,
}

export type WebsocketMessage = {
    action: WebsocketAction,
    response?: WebsocketResponse,
    position?: string,
    sessionId: string
}

export type MoveMessage = {
    position: string,
    sessionId: string
}

export function SendWebsocketMove(data: MoveMessage) {
    getDevMode() && console.log("WS move sent")
    let msg: WebsocketMessage = {
        action: WebsocketAction.MOVE,
        position: data.position,
        sessionId: data.sessionId
    }
    wsClient.send(JSON.stringify(msg));
}

export function SendWebsocketJoinRoom(sessionId: string): boolean {
    if (sessionId === '' || sessionId === undefined) {
        return false;
    }
    getDevMode() && console.log("WS join room sent")
    let msg: WebsocketMessage = {
        action: WebsocketAction.JOIN_ROOM,
        sessionId: sessionId
    }
    wsClient.send(JSON.stringify(msg));
    return true;
}