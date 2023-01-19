import { useDispatch, useSelector } from 'react-redux';
import { selectRecentSessionStatus, selectWindowMinDimension, updateSessionId } from '../state/game/game.slice';
import { SiAddthis } from 'react-icons/si';
import { MdOutlineOpenInNew } from 'react-icons/md';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { useEffect } from 'react';
import { getRecentSession } from '../utilities/storage.utility';
import './home.style.css';
import { getApiUrl } from '../utilities/environment.utility';

export default function Home() {
    const windowMinDimensions = useSelector(selectWindowMinDimension);
    const recentSessionAvailable = useSelector(selectRecentSessionStatus);
    const navigate = useNavigate();
    const dispatch = useDispatch();
    const recentSession = getRecentSession();

    const squareStyle = {
        width: windowMinDimensions / 3,
        height: windowMinDimensions / 3,
    }

    useEffect(() => {
        dispatch(updateSessionId(""))
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    return (
        <div style={{ boxShadow: '0 5px 30px rgba(0, 0, 0, 0.5)' }}>
            <div className='flex flex-row'>
                <div style={squareStyle} className='bg-colorSecondary text-colorDetails text-3xl sm:text-4xl md:text-5xl lg:text-6xl flex flex-col items-center justify-center break-words text-center select-none title'>
                    <p>
                        chasse
                    </p>
                </div>
                <div style={squareStyle} className='inner-shadow text-colorSecondary font-bold text-2xl sm:text-3xl md:text-4xl lg:text-5xl flex flex-col items-center justify-center break-words text-center select-none'>
                    <div className='w-1/2'>
                        <img src="/logo.png" alt="chasse" />   
                    </div>
                </div>
            </div>
            <div className='flex flex-row'>
                {
                    !recentSessionAvailable &&
                    <div style={squareStyle}></div>
                }
                {
                    recentSessionAvailable &&
                    <RecentSessionSquare
                        style={squareStyle}
                        navigate={navigate}
                        recent={recentSession}
                    />
                }
                <CreateSessionSquare
                    navigate={navigate}
                    style={squareStyle}
                />
            </div>
        </div>
    );
}

function CreateSessionSquare(props) {
    let createSession = () => {
        axios
            .get(getApiUrl() + "api/v1/session/new")
            .then((response: any) => {
                props.navigate("/board/" + response.data.sessionId)
            })
            .catch((err) => console.log(err));
    }

    return (
        <div onClick={() => createSession()} style={props.style} className='bg-colorSecondary hover:bg-colorSecondaryDark inner-shadow cursor-pointer text-colorMainDark font-bold text-sm sm:text-xl flex flex-col items-center justify-center break-words text-center select-none'>
            <p className="pb-1">
                NEW BOARD
            </p>
            <SiAddthis />
        </div>
    )
}

function RecentSessionSquare(props) {
    let openSession = () => {
        props.navigate("/board/" + props.recent)
    }

    return (
        <div onClick={() => openSession()} style={props.style} className='hover:bg-colorMainDark inner-shadow cursor-pointer text-colorSecondary font-bold text-sm sm:text-xl flex flex-col items-center justify-center break-words text-center select-none'>
            <p className="pb-1">
                OPEN LAST BOARD
            </p>
            <MdOutlineOpenInNew />
        </div>
    )
}
