import { Routes, Route, useNavigate, NavigateFunction } from "react-router-dom";
import Home from './pages/home/home.page';
import Game from './pages/game/game.page';
import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { updatePosition, updateRecentSessionState, updateWindowProperties, updateWsState } from './state/game/game.slice';
import './socket/socket'
import { WebsocketAction, WebsocketMessage, WebsocketResponse, wsClient } from "./socket/socket";
import { clearRecentData, getRecentSession } from "./utilities/storage.utility";
import { getDevMode, getAppVersion, getApiUrl } from "./utilities/environment.utility";
import { AnyAction, Dispatch } from 'redux';
import axios, { AxiosResponse } from "axios";

export default function App() {
    const dispatch = useDispatch();
    const navigate = useNavigate();

    useEffect(() => {
        setupApplication(dispatch, navigate)
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    return (
        <div className="flex flex-col items-center justify-center bg-colorMain min-h-screen text-lg text-white">
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="board/" element={<Home />} />
                <Route path="board/:sessionId" element={<Game />} />
            </Routes>
        </div>
    );
}

function setupApplication(dispatch: Dispatch<AnyAction>, navigate: NavigateFunction) {
    console.log('APP VERSION: ', getAppVersion())

    const refreshPage = () => {
        navigate(0);
    }

    wsClient.onmessage = (message) => {
        let msg: WebsocketMessage = JSON.parse(message.data.toString())
        switch (msg.response) {
            case WebsocketResponse.BLANK:
                if (msg.action === WebsocketAction.MOVE && msg.position !== undefined) {
                    getDevMode() && console.log('WS - move received');
                    dispatch(updatePosition(msg.position))
                }
                break
            case WebsocketResponse.ERROR:
                getDevMode() && console.log('WS respond: ERROR');
                if (msg.action === WebsocketAction.JOIN_ROOM) {
                    clearRecentData()
                    navigate("/") // TODO: find better solution for restarting (maybe popup with button)
                }
                break
            case WebsocketResponse.OK:
                if (msg.action === WebsocketAction.CONNECT) {
                    dispatch(updateWsState(true));
                    getDevMode() && console.log('WS connection successful');
                    break
                }
                getDevMode() && console.log('WS respond: OK');
                break
        }
    };
    wsClient.onclose = () => {
        getDevMode() && console.log('WS disconnected');
        dispatch(updateWsState(false));
        refreshPage();
    };

    function checkLastSession() {
        let recentSessionId = getRecentSession()
        
        if (recentSessionId !== null) {
            axios
                .get(getApiUrl() + "api/v1/session/"+recentSessionId)
                .then((_: AxiosResponse) => {
                    dispatch(updateRecentSessionState(true))
                })
                .catch((_) => {
                    clearRecentData()
                });
        }
    }
    checkLastSession()

    function handleResize() {
        dispatch(updateWindowProperties());
    }
    window.addEventListener('resize', handleResize);
    return () => window.removeEventListener('resize', handleResize);
}
