import { reverseBoard, selectSessionId, selectHistory, MoveItem, makeMove, historyPop, toggleTabletMode } from '../state/game/game.slice';
import { useAppDispatch } from '../state/hooks';
import { START_POSITION_OBJECT } from '../utilities/chess.utility';
import { BackIcon, MenuIcon, ReverseIcon, TabletModeIcon } from '../utilities/icons.utility'
import { peek2 } from '../utilities/stack.utility';
import { useAppSelector } from '../state/hooks';
import QuickMenu, { Direction, QuickMenuButtonProps } from './quickmenu.component';

export default function Controls() {
    const dispatch = useAppDispatch();
    const sessionId = useAppSelector(selectSessionId);
    const moveHistory = useAppSelector(selectHistory);
 
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

    let boardActionButtons: QuickMenuButtonProps[] = [
        {
            text: "Clear",
            handler: () => {
                let moveItem: MoveItem = {
                    position: {},
                    sessionId: sessionId
                }
                dispatch(makeMove(moveItem));
            }
        },
        {
            text: "Reset",
            handler: () => {
                let moveItem: MoveItem = {
                    position: START_POSITION_OBJECT,
                    sessionId: sessionId
                }
                dispatch(makeMove(moveItem));
            }
        },
    ]
    
    return (
        <div className="flex sm:flex-col justify-center sm:ml-4 text-colorMain">
            <div className="flex sm:flex-col mt-4 pb-2 sm:mt-0 sm:pt-4 sm:pb-4 bg-colorSecondary rounded-full">
                <button className="pt-2 pl-4 pr-4 sm:pl-2 sm:pr-2" onClick={() => goBack()}>
                    <BackIcon />
                </button>
                <button className="pt-2 pl-4 pr-4 sm:pl-2 sm:pr-2" onClick={() => dispatch(reverseBoard())}>
                    <ReverseIcon />
                </button>
                <button className="pt-2 pl-4 pr-4 sm:pl-2 sm:pr-2" onClick={() => dispatch(toggleTabletMode())}>
                    <TabletModeIcon />
                </button>
                <div className="pt-2 pl-4 pr-4 sm:pl-2 sm:pr-2 text-colorRed">
                    <QuickMenu
                        direction={Direction.Up}
                        icon={<MenuIcon />}
                        buttons={boardActionButtons}
                    />
                </div>
            </div>
        </div>
    );
}
