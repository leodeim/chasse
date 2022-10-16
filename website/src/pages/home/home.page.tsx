import { useDispatch, useSelector } from 'react-redux';
import { selectWindowMinDimension, updateSessionId } from '../../state/game/game.slice';
import { FaChess } from 'react-icons/fa';
import { SiAddthis } from 'react-icons/si';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { useEffect } from 'react';

export default function Home() {
    const windowMinDimensions = useSelector(selectWindowMinDimension);
    const navigate = useNavigate();
    const dispatch = useDispatch();

    const squareStyle = {
        width: windowMinDimensions / 3,
        height: windowMinDimensions / 3,
    }

    useEffect(() => {
        dispatch(updateSessionId(""))
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    let createSession = () => {
        axios
            .get("http://localhost:8085/api/v1/session/new")
            .then((response: any) => {
                console.log(response)
                navigate("/board/" + response.data.sessionId)
            })
            .catch((err) => console.log(err));
    }

    return (
        <div style={{ boxShadow: '0 5px 30px rgba(0, 0, 0, 0.5)' }}>
            <div className='flex flex-row'>
                <div style={squareStyle} className='bg-yellow text-grey text-xs sm:text-lg flex flex-col items-center justify-center break-words text-center select-none'>
                    <FaChess />
                    <p>
                        ZEN CHESS
                    </p>
                </div>
                <div style={squareStyle}></div>
            </div>
            <div className='flex flex-row'>
                <div style={squareStyle}></div>
                <div onClick={ () => createSession() } style={squareStyle} className='bg-yellow hover:bg-darkYellow inner-shadow cursor-pointer text-darkDarkGreen font-bold text-sm sm:text-xl flex flex-col items-center justify-center break-words text-center select-none'>
                    <p className="pb-1">
                        CREATE SESSION
                    </p>
                    <SiAddthis />
                </div>
            </div>
        </div>
    );
}