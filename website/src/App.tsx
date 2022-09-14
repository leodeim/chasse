import { Routes, Route } from "react-router-dom";
import Home from './pages/home/home.page';
import Game from './pages/game/game.page';
import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { updateWindowProperties } from './state/game/game.slice';

export default function App() {
    const dispatch = useDispatch();

    useEffect(() => {
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
