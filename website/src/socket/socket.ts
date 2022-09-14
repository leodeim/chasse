import { useDispatch } from "react-redux";
import { w3cwebsocket } from "websocket";

export const wsClient = new w3cwebsocket('ws://127.0.0.1:8085/ws');

console.log("SOCKET TEST TEST TEST")

export enum WebsocketAction {
	MOVE = 0,
	GO_BACK = 1,
	RESET = 2,
	JOIN_ROOM = 3
}

type WebsocketMessage = {
    action: WebsocketAction,
    fen?: string,
    sessionId: string
}

export type MoveMessage = {
    fen: string,
    sessionId: string
}

wsClient.onopen = () => {
    console.log('WebSocket Client Connected');
};

export function SendWebsocketMove(data: MoveMessage) {
    console.log("SendWebsocketMove")
    let msg: WebsocketMessage = {
        action: WebsocketAction.MOVE,
        fen: data.fen,
        sessionId: data.sessionId
    }
    wsClient.send(JSON.stringify(msg));
}

export function SendWebsocketJoinRoom(sessionId: string) {
    console.log("SendWebsocketJoinRoom")
    let msg: WebsocketMessage = {
        action: WebsocketAction.JOIN_ROOM,
        sessionId: sessionId
    }
    wsClient.send(JSON.stringify(msg));
}