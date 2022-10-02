import { Routes, Route } from "react-router-dom";
import Home from './pages/home/home.page';
import Game from './pages/game/game.page';
import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { updatePosition, updateWindowProperties, updateWsState } from './state/game/game.slice';
import './socket/socket'
import { wsClient } from "./socket/socket";

export default function App() {
    const dispatch = useDispatch();

    useEffect(() => {
        wsClient.onmessage = (message) => {
            dispatch(updatePosition(JSON.parse(message.data.toString()).position))
        };
        wsClient.onopen = () => {
            console.log('WebSocket Connected');
            dispatch(updateWsState(true));
        };
        wsClient.onclose = () => {
            console.log('WebSocket Disconnected');
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
                <Route path="board/:sessionId" element={<Game />} />
            </Routes>
        </div>
    );
}
