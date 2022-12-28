import GameBoard from "../../components/board/board.component";
import Controls from "../../components/controls/controls.component";
import Menu from "../../components/menu/menu.component";
import { SendWebsocketJoinRoom } from "../../socket/socket";
import { useEffect } from "react";
import { selectWsState, updateSessionId } from "../../state/game/game.slice";
import { useDispatch, useSelector } from "react-redux";
import { useParams, useNavigate } from "react-router-dom";
import { storeSession } from "../../utilities/storage.utility";


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
            SendWebsocketJoinRoom(sessionId)
            storeSession(sessionId)
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [wsState]);

    return (
        <div className="flex sm:flex-row flex-col justify-around min-h-screen">
            <div className="flex justify-center">
                <Menu />
            </div>
            <div className="flex flex-col justify-center">
                <GameBoard />
            </div>
            <div className="flex justify-center">
                <Controls />
            </div>
        </div>
    );
}