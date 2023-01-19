import GameBoard from "../../components/board/board.component";
import Controls from "../../components/controls/controls.component";
import Menu from "../../components/menu/menu.component";
import { useEffect } from "react";
import { selectWsState, updateRecentSessionState, updateSessionId } from "../../state/game/game.slice";
import { useDispatch, useSelector } from "react-redux";
import { useParams, useNavigate } from "react-router-dom";
import { storeSession } from "../../utilities/storage.utility";
import { wsHandler } from "../../socket/setup.socket";


export default function Game() {
    const navigate = useNavigate();
    let { sessionId } = useParams();
    const wsState = useSelector(selectWsState);
    const dispatch = useDispatch();

    useEffect(() => {
        if (sessionId === undefined) {
            navigate("/");
        }
        if (wsState === true && sessionId !== undefined) {
            dispatch(updateSessionId(sessionId))
            wsHandler.sendJoinRoom(sessionId)
            storeSession(sessionId)
            dispatch(updateRecentSessionState(true))
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [wsState]);

    return (
        <div className="sm:flex sm:flex-row">
            <div className="flex justify-center pb-7 sm:pb-0">
                <Menu />
            </div>
            <GameBoard />
            <div className="flex justify-center pt-7 sm:pt-0">
                <Controls />
            </div>
        </div>
    );
}