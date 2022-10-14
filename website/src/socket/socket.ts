import { w3cwebsocket } from "websocket";

export const wsClient = new w3cwebsocket('ws://127.0.0.1:8085/ws');

export enum WebsocketAction {
	MOVE = 0,
	GO_BACK = 1,
	RESET = 2,
    JOIN_ROOM = 3,
    ERROR = 4,
    OK = 5,
}

export type WebsocketMessage = {
    action: WebsocketAction,
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