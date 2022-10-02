import { useEffect } from "react";
import { useDispatch } from "react-redux";
import GameBoard from "../../components/board/board.component";
import Controls from "../../components/controls/controls.component";
import Menu from "../../components/menu/menu.component";
import { wsClient } from "../../socket/socket";
import { updatePosition, updateWsState } from "../../state/game/game.slice";

export default function Game() {
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