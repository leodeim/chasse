import GameBoard from "../../components/board/board.component";
import Controls from "../../components/controls/controls.component";
import Menu from "../../components/menu/menu.component";

export default function Game() {
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