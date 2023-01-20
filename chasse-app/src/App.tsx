import { Routes, Route, useNavigate } from "react-router-dom";
import { Dispatch, useEffect } from 'react';
import { AnyAction } from "@reduxjs/toolkit";
import Home from './pages/home.page';
import Game from './pages/game.page';
import { setupWsApp } from "./socket/socket.setup";
import { clearRecentData, getRecentSession } from "./utilities/storage.utility";
import axios, { AxiosResponse } from "axios";
import { getApiUrl, getAppVersion } from "./utilities/environment.utility";
import { selectWsState, updateRecentSessionState, updateWindowProperties } from "./state/game/game.slice";
import { useAppDispatch, useAppSelector } from "./state/hooks";
import InfoDialog from "./components/dialog.component";


export default function App() {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();
    const wsState = useAppSelector(selectWsState);

    useEffect(() => {
        console.log('APP VERSION:', getAppVersion())

        setupWsApp(dispatch, navigate)
        checkLastSession(dispatch)

        let handleResize = () => dispatch(updateWindowProperties());
        window.addEventListener('resize', handleResize);
        return () => window.removeEventListener('resize', handleResize);
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    return (
        <div className="flex flex-col items-center justify-center bg-colorMain min-h-screen text-lg text-white">
            <InfoDialog
                isOpen={!wsState}
                title="Connecting..."
                text="Please wait, trying to reconnect"
            />
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="board/" element={<Home />} />
                <Route path="board/:sessionId" element={<Game />} />
            </Routes>
        </div>
    );
}

function checkLastSession(dispatch: Dispatch<AnyAction>) {
    let recentSessionId = getRecentSession()

    // TODO: move all API requests
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
