import { Routes, Route, useNavigate } from "react-router-dom";
import Home from './pages/home/home.page';
import Game from './pages/game/game.page';
import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { updatePosition, updateWindowProperties, updateWsState } from './state/game/game.slice';
import './socket/socket'
import { WebsocketAction, WebsocketMessage, WebsocketResponse, wsClient } from "./socket/socket";
import { clearRecentData } from "./utilities/storage.utility";

export default function App() {
    const dispatch = useDispatch();
    const navigate = useNavigate();

    const refreshPage = () => {
        navigate(0);
    }

    useEffect(() => {
        wsClient.onmessage = (message) => {
            let msg: WebsocketMessage = JSON.parse(message.data.toString())
            switch (msg.response) {
                case WebsocketResponse.BLANK:
                    if (msg.action === WebsocketAction.MOVE && msg.position !== undefined) {
                        console.log('WS - move received');
                        dispatch(updatePosition(msg.position))
                    }
                    break
                case WebsocketResponse.ERROR:
                    console.log('WS respond: ERROR');
                    if (msg.action === WebsocketAction.JOIN_ROOM) {
                        clearRecentData()
                        navigate("/") // TODO: find better solution for restarting (maybe popup with button)
                    }
                    break
                case WebsocketResponse.OK:
                    if (msg.action === WebsocketAction.CONNECT) {
                        dispatch(updateWsState(true));
                        console.log('WS connection successful');
                        break
                    }
                    console.log('WS respond: OK');
                    break
            }
        };
        wsClient.onclose = () => {
            console.log('WS disconnected');
            dispatch(updateWsState(false));
            refreshPage();
        };
        function handleResize() {
            dispatch(updateWindowProperties());
        }

        window.addEventListener('resize', handleResize);
        return () => window.removeEventListener('resize', handleResize);
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
