import { Routes, Route, useNavigate, NavigateFunction } from "react-router-dom";
import { useEffect } from 'react';
import Home from './pages/home.page';
import Game from './pages/game.page';
import { setupWsApp } from "./socket/socket.setup";
import { clearRecentData, getRecentSession } from "./utilities/storage.utility";
import { getAppVersion } from "./utilities/environment.utility";
import { selectWsState, updateRecentSessionState, updateWindowProperties } from "./state/game/game.slice";
import { useAppDispatch, useAppSelector } from "./state/hooks";
import InfoDialog from "./components/dialog.component";
import { getSession, SessionData } from "./api/api.session";
import { AppDispatch } from "./state/store";


export default function App() {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();
    const wsState = useAppSelector(selectWsState);

    useEffect(() => {
        prepareApplication(dispatch, navigate)
        let handleResize = () => dispatch(updateWindowProperties());
        window.addEventListener('resize', handleResize);
        return () => window.removeEventListener('resize', handleResize);
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    return (
        <div className="flex flex-col items-center justify-center bg-colorMain min-h-screen text-lg text-white">
            {
                !wsState &&
                <InfoDialog
                    isOpen={!wsState}
                    title="Connecting..."
                    text="Please wait, trying to reconnect"
                />
            }

            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="board/" element={<Home />} />
                <Route path="board/:sessionId" element={<Game />} />
            </Routes>
        </div>
    );
}

function prepareApplication(dispatch: AppDispatch, navigate: NavigateFunction) {
    console.log(`APP VERSION:`, getAppVersion())

    setupWsApp(dispatch, navigate)
    
    let recentSessionId = getRecentSession()
    if (recentSessionId !== null) {
        getSession(recentSessionId, (status: number, _: SessionData) => {
            if (status === 200) dispatch(updateRecentSessionState(true))
            else clearRecentData()
        }, () => {
            clearRecentData()
        })
    }
}
