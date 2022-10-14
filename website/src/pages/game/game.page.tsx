import GameBoard from "../../components/board/board.component";
import Controls from "../../components/controls/controls.component";
import Menu from "../../components/menu/menu.component";
import { SendWebsocketJoinRoom } from "../../socket/socket";
import { useEffect } from "react";
import { selectWsState } from "../../state/game/game.slice";
import { useSelector } from "react-redux";
import { useParams, useNavigate } from "react-router-dom";


export default function Game() {
    const navigate = useNavigate();
    let { sessionId } = useParams();
    const wsState = useSelector(selectWsState)

    useEffect(() => {
        console.log(wsState === true)
        console.log(sessionId !== undefined)
        console.log(wsState === true && sessionId !== undefined)

        if (sessionId === undefined) {
            navigate("/");
        }
        if (wsState === true && sessionId !== undefined) {
            SendWebsocketJoinRoom(sessionId)
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [wsState]);

    return (
        <div className="sm:flex sm:flex-row">
            <div className="flex justify-center">
                <Menu />
            </div>
            <GameBoard />
            <div className="flex justify-center">
                <Controls />
            </div>
        </div>
    );
}