import { useParams } from "react-router-dom";
import Board from "../../components/board/board.component";
import Controls from "../../components/controls/controls.component";
import Menu from "../../components/menu/menu.component";

export default function Game() {
    let { sessionId } = useParams();
    console.log(sessionId)

    return (
        <div>
            <div className="flex justify-center">
                <Menu />
            </div>
            <Board />
            <div className="flex justify-center">
                <Controls />
            </div>
        </div>
    );
}