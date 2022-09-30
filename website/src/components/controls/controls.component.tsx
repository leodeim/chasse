import { useSelector } from 'react-redux';
import { reverseBoard, selectSessionId, selectHistory, MoveItem, makeMove, historyPop } from '../../state/game/game.slice';
import { useAppDispatch } from '../../state/hooks';
import { START_POSITION } from '../../utilities/chess.utility';
import { BackIcon, EndIcon, ReverseIcon } from '../../utilities/icons.utility'
import { peek2 } from '../../utilities/stack.utility';

export default function Controls() {
    const dispatch = useAppDispatch();
    const sessionId = useSelector(selectSessionId);
    const moveHistory = useSelector(selectHistory);
 
    let goBack = () => {
        let lastPosition = peek2(moveHistory);
        dispatch(historyPop())
        if (lastPosition !== undefined) {
            let moveItem: MoveItem = {
                position: lastPosition,
                sessionId: sessionId
            }
            dispatch(makeMove(moveItem));
        }
    }

    let resetBoard = () => {
        let moveItem: MoveItem = {
            position: START_POSITION,
            sessionId: sessionId
        }
        dispatch(makeMove(moveItem));
    }
    
    return (
        <div className="flex justify-center mt-4 pb-2 text-darkGreen bg-yellow rounded-full">
            <button className="pt-2 pl-4 pr-4 " onClick={() => goBack()}>
                <BackIcon />
            </button>
            <button className="pt-2 pl-4 pr-4" onClick={() => dispatch(reverseBoard())}>
                <ReverseIcon />
            </button>
            <button className="pt-2 pl-4 pr-4" onClick={() => resetBoard()}>
                <EndIcon />
            </button>
        </div>
    );
}
