import { Routes, Route } from "react-router-dom";
import Home from './pages/home/home.page';
import Game from './pages/game/game.page';
import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { updatePosition, updateWindowProperties, updateWsState } from './state/game/game.slice';
import './socket/socket'
import { WebsocketAction, WebsocketMessage, wsClient } from "./socket/socket";

export default function App() {
    const dispatch = useDispatch();

    useEffect(() => {
        wsClient.onmessage = (message) => {
            let msg : WebsocketMessage = JSON.parse(message.data.toString())
            switch (msg.action) {
                case WebsocketAction.MOVE:
                    if (msg.position !== undefined) {
                        console.log('WS - move received');
                        dispatch(updatePosition(msg.position))
                    }
                    break
                case WebsocketAction.ERROR:
                    console.log('WS respond: ERROR');
                    break
                case WebsocketAction.OK:
                    console.log('WS respond: OK');
                    break
                default:
                    console.log('WS - bad message received');
            }
        };
        wsClient.onopen = () => {
            console.log('WS connected');
            dispatch(updateWsState(true));
        };
        wsClient.onclose = () => {
            console.log('WS disconnected');
            dispatch(updateWsState(false));
        };
        function handleResize() {
            dispatch(updateWindowProperties());
        }

        window.addEventListener('resize', handleResize);
        return () => window.removeEventListener('resize', handleResize);
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    return (
        <div className="flex flex-col items-center justify-center bg-green min-h-screen text-lg text-white">
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="board/" element={<Home />} />
                <Route path="board/:sessionId" element={<Game />} />
            </Routes>
        </div>
    );
}
