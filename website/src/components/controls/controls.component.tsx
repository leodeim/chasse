import { useSelector } from 'react-redux';
import { reverseBoard, selectSessionId, selectHistory, MoveItem, makeMove, historyPop } from '../../state/game/game.slice';
import { useAppDispatch } from '../../state/hooks';
import { START_POSITION_OBJECT } from '../../utilities/chess.utility';
import { BackIcon, EndIcon, ReverseIcon, StartIcon } from '../../utilities/icons.utility'
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
            position: START_POSITION_OBJECT,
            sessionId: sessionId
        }
        dispatch(makeMove(moveItem));
    }

    let clearBoard = () => {
        let moveItem: MoveItem = {
            position: {},
            sessionId: sessionId
        }
        dispatch(makeMove(moveItem));
    }
    
    return (
        <div className="flex sm:flex-col justify-center ml-4 text-green">
            <div className="flex sm:flex-col mt-4 pb-2 sm:mt-0 sm:pt-4 sm:pb-4 bg-yellow rounded-full">
                <button className="pt-2 pl-4 pr-4 sm:pl-2 sm:pr-2" onClick={() => goBack()}>
                    <BackIcon />
                </button>
                <button className="pt-2 pl-4 pr-4 sm:pl-2 sm:pr-2" onClick={() => dispatch(reverseBoard())}>
                    <ReverseIcon />
                </button>
                <button className="pt-2 pl-4 pr-4 sm:pl-2 sm:pr-2 text-blue" onClick={() => resetBoard()}>
                    <StartIcon />
                </button>
                <button className="pt-2 pl-4 pr-4 sm:pl-2 sm:pr-2 sm:mb-2 text-red" onClick={() => clearBoard()}>
                    <EndIcon />
                </button>
            </div>
        </div>
    );
}
