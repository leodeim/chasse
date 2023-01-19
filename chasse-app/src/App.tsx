import { Routes, Route, useNavigate } from "react-router-dom";
import Home from './pages/home.page';
import Game from './pages/game.page';
import { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { setupWsApp } from "./socket/socket.setup";
import { clearRecentData, getRecentSession } from "./utilities/storage.utility";
import axios, { AxiosResponse } from "axios";
import { getApiUrl, getAppVersion } from "./utilities/environment.utility";
import { selectWsState, updateRecentSessionState, updateWindowProperties } from "./state/game/game.slice";
import { Dialog, Transition } from '@headlessui/react'
import { Fragment } from 'react'


export default function App() {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const wsState = useSelector(selectWsState);

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
            <ReconnectingDialog
                isOpen={!wsState}
            />
            <Routes>
                <Route path="/" element={<Home />} />
                <Route path="board/" element={<Home />} />
                <Route path="board/:sessionId" element={<Game />} />
            </Routes>
        </div>
    );
}

function ReconnectingDialog(props) {
    return (
        <Transition appear show={props.isOpen} as={Fragment}>
            <Dialog as="div" className="relative z-10" onClose={() => { }}>
                <Transition.Child
                    as={Fragment}
                    enter="ease-out duration-300"
                    enterFrom="opacity-0"
                    enterTo="opacity-100"
                    leave="ease-in duration-200"
                    leaveFrom="opacity-100"
                    leaveTo="opacity-0"
                >
                    <div className="fixed inset-0 bg-black bg-opacity-25" />
                </Transition.Child>

                <div className="fixed inset-0 overflow-y-auto">
                    <div className="flex min-h-full items-center justify-center p-4 text-center">
                        <Transition.Child
                            as={Fragment}
                            enter="ease-out duration-300"
                            enterFrom="opacity-0 scale-95"
                            enterTo="opacity-100 scale-100"
                            leave="ease-in duration-200"
                            leaveFrom="opacity-100 scale-100"
                            leaveTo="opacity-0 scale-95"
                        >
                            <Dialog.Panel className="w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all">
                                <Dialog.Title
                                    as="h3"
                                    className="text-lg font-medium leading-6 text-gray-900"
                                >
                                    Connecting...
                                </Dialog.Title>
                                <div className="mt-2">
                                    <p className="text-sm text-gray-500">
                                        Please wait, trying to reconnect
                                    </p>
                                </div>
                            </Dialog.Panel>
                        </Transition.Child>
                    </div>
                </div>
            </Dialog>
        </Transition>
    )
}
