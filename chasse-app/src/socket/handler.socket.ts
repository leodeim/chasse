import { w3cwebsocket, IMessageEvent, ICloseEvent } from "websocket";
import { getWebsocketUrl, getDevMode } from "../utilities/environment.utility";

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

export class SocketHandler {
    public client: w3cwebsocket | undefined;
   
    constructor(callbacks: SocketCallbacks) {
        this.connect(callbacks)
    }

    connect(callbacks: SocketCallbacks) {
        getDevMode() && console.log('WS CONNECTING...');
        this.client = new w3cwebsocket(getWebsocketUrl() + 'api/ws');
        this.registerCallbacks(callbacks)
    }

    registerCallbacks(callbacks: SocketCallbacks) {
        if (callbacks.message !== undefined) this.client!.onmessage = callbacks.message;
        if (callbacks.open !== undefined) this.client!.onopen = callbacks.open;
        if (callbacks.close !== undefined) this.client!.onclose = callbacks.close;
        if (callbacks.error !== undefined) this.client!.onerror = callbacks.error;
    }

    private clientIsActive(): boolean {
        if (!this.client === undefined) {
            return false
        }
        return true
    }
   
    sendMove(data: MoveMessage) {
        if (!this.clientIsActive()) {
            return
        }

        getDevMode() && console.log("WS <- MOVE")
        let msg: WebsocketMessage = {
            action: WebsocketAction.MOVE,
            position: data.position,
            sessionId: data.sessionId
        }
        this.client!.send(JSON.stringify(msg));
    }
    
    sendJoinRoom(sessionId: string): boolean {
        if (!this.clientIsActive() || (sessionId === '' || sessionId === undefined)) {
            return false;
        }

        getDevMode() && console.log("WS <- JOIN ROOM")
        let msg: WebsocketMessage = {
            action: WebsocketAction.JOIN_ROOM,
            sessionId: sessionId
        }
        this.client!.send(JSON.stringify(msg));
        return true;
    }
}

interface messageReceiver {
    (message: IMessageEvent): void;
};

interface openReceiver {
    (): void;
};

interface closeReceiver {
    (event: ICloseEvent): void;
};

interface errorReceiver {
    (error: Error): void;
};

export type SocketCallbacks = {
    message?: messageReceiver
    open?: openReceiver
    close?: closeReceiver
    error?: errorReceiver
}
