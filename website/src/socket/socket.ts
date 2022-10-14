import { w3cwebsocket } from "websocket";

export const wsClient = new w3cwebsocket('ws://127.0.0.1:8085/ws');

export enum WebsocketAction {
    BLANK = 0,
    MOVE = 1,
	GO_BACK = 2,
	RESET = 3,
    JOIN_ROOM = 4,
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
    console.log("WS move sent")
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
    console.log("WS join room sent")
    let msg: WebsocketMessage = {
        action: WebsocketAction.JOIN_ROOM,
        sessionId: sessionId
    }
    wsClient.send(JSON.stringify(msg));
    return true;
}