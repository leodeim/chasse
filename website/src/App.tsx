import Board from './components/board/board.component'
import { Routes, Route, Link } from "react-router-dom";
import Home from './pages/home/home.page';
import Game from './pages/game/game.page';

export default function App() {
    return (
        <div className="flex flex-col items-center justify-center bg-green min-h-screen text-lg text-white">
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="board/:sessionId" element={<Game />} />
            </Routes>
        </div>
    );
}
