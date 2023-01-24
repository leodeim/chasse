import { NavigateFunction } from "react-router-dom";
import { IMessageEvent, ICloseEvent } from "websocket";
import { WebsocketAction, WebsocketMessage, WebsocketResponse, SocketHandler, SocketCallbacks } from "./socket.handler";
import { updatePosition, updateWsState } from '../state/game/game.slice';
import { clearRecentData } from "../utilities/storage.utility";
import { getDevMode } from "../utilities/environment.utility";
import './socket.handler'
import { AppDispatch } from '../state/store';

let wsCallbacks: SocketCallbacks = {}

export const wsHandler = new SocketHandler(wsCallbacks)

export function setupWsApp(dispatch: AppDispatch, navigate: NavigateFunction) {
    wsCallbacks.message = (message: IMessageEvent) => {
        onMessage(message, dispatch, navigate);
    }
    wsCallbacks.open = () => {
        onOpen(dispatch);
    }
    wsCallbacks.close = (event: ICloseEvent) => {
        onClose(event, dispatch);
    }
    wsCallbacks.error = onError;

    wsHandler.registerCallbacks(wsCallbacks);
}

function onMessage(message: IMessageEvent, dispatch: AppDispatch, navigate: NavigateFunction) {
    getDevMode() && console.log(`WS CALLBACK: MESSAGE`);

    let msg: WebsocketMessage = JSON.parse(message.data.toString())
    switch (msg.response) {
        case WebsocketResponse.BLANK:
            if (msg.action === WebsocketAction.MOVE && msg.position !== undefined) {
                getDevMode() && console.log(`WS -> MOVE`);
                dispatch(updatePosition(msg.position));
            }
            break;
        case WebsocketResponse.ERROR:
            getDevMode() && console.log(`WS -> ERROR`);
            if (msg.action === WebsocketAction.JOIN_ROOM) {
                clearRecentData();
                navigate("/"); // TODO: find better solution for error handling
            }
            break;
        case WebsocketResponse.OK:
            if (msg.action === WebsocketAction.CONNECT) {
                dispatch(updateWsState(true));
                getDevMode() && console.log(`WS -> CONNECTION SUCCESSFUL`);
                break;
            }
            getDevMode() && console.log(`WS -> OK`);
            break;
    }
}

function onOpen(dispatch: AppDispatch) {
    getDevMode() && console.log(`WS CALLBACK: OPEN`);
    dispatch(updateWsState(true))
}

function onClose(event: ICloseEvent, dispatch: AppDispatch) {
    getDevMode() && console.log(`WS CALLBACK: CLOSE`);
    getDevMode() && console.log(`WS -> ` + event.code);
    dispatch(updateWsState(false));
    setTimeout(function () {
        wsHandler!.connect(wsCallbacks);
        getDevMode() && console.log(`WS TRY RECONNECT`);
    }, 1000);
}

function onError(error: Error) {
    getDevMode() && console.log(`WS CALLBACK: ERROR`);
    getDevMode() && console.log(`WS -> ` + error.message);
}
