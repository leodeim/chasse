import { Routes, Route, useNavigate } from "react-router-dom";
import Home from './pages/home/home.page';
import Game from './pages/game/game.page';
import { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { setupWsApp } from "./socket/setup";
import { clearRecentData, getRecentSession } from "./utilities/storage.utility";
import axios, { AxiosResponse } from "axios";
import { getApiUrl, getAppVersion } from "./utilities/environment.utility";
import { updateRecentSessionState, updateWindowProperties } from "./state/game/game.slice";


export default function App() {
    const dispatch = useDispatch();
    const navigate = useNavigate();

    useEffect(() => {
        console.log('APP VERSION:', getAppVersion())
        
        setupWsApp(dispatch, navigate)

        function checkLastSession() {
            let recentSessionId = getRecentSession()
    
            if (recentSessionId !== null) {
                axios
                    .get(getApiUrl() + "api/v1/session/" + recentSessionId)
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
