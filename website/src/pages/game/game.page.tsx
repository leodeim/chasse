import { useEffect } from "react";
import { useDispatch } from "react-redux";
import GameBoard from "../../components/board/board.component";
import Controls from "../../components/controls/controls.component";
import Menu from "../../components/menu/menu.component";
import { wsClient } from "../../socket/socket";
import { updatePosition, updateWsState } from "../../state/game/game.slice";

export default function Game() {
    const dispatch = useDispatch();

    useEffect(() => {
        wsClient.onmessage = (message) => {
            dispatch(updatePosition(JSON.parse(message.data.toString()).fen))
        };
        wsClient.onopen = () => {
            console.log('WebSocket Connected');
            dispatch(updateWsState(true));
        };
        
        wsClient.onclose = () => {
            console.log('WebSocket Disconnected');
            dispatch(updateWsState(false));
        };
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);


    return (
        <div>
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