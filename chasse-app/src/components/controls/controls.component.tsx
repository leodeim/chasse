import { useSelector } from 'react-redux';
import { reverseBoard, selectSessionId, selectHistory, MoveItem, makeMove, historyPop, toggleTabletMode } from '../../state/game/game.slice';
import { useAppDispatch } from '../../state/hooks';
import { START_POSITION_OBJECT } from '../../utilities/chess.utility';
import { BackIcon, EndIcon, MenuIcon, ReverseIcon, StartIcon, TabletModeIcon } from '../../utilities/icons.utility'
import { peek2 } from '../../utilities/stack.utility';
import { Popover, Transition } from '@headlessui/react'
import { Fragment } from 'react'

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
                <button className="pt-2 pl-4 pr-4 sm:pl-2 sm:pr-2 text-colorRed">
                    <BoardActionsPopover
                        reset={resetBoard}
                        clear={clearBoard}
                    />
                </button>
            </div>
        </div>
    );
}

function BoardActionsPopover(props) {
    return (
        <Popover className="relative flex">
            {({ open }) => (
                <>
                    <Popover.Button>
                        <MenuIcon />
                    </Popover.Button>
                    <Transition
                        as={Fragment}
                        enter="transition ease-out duration-200"
                        enterFrom="opacity-0 translate-y-1"
                        enterTo="opacity-100 translate-y-0"
                        leave="transition ease-in duration-150"
                        leaveFrom="opacity-100 translate-y-0"
                        leaveTo="opacity-0 translate-y-1"
                    >
                        <Popover.Panel className="absolute right-full bottom-full z-50 w-24">
                        <div className="overflow-hidden rounded-lg shadow-lg ring-1 ring-black ring-opacity-5">
                                <div className="relative grid bg-colorMainLight p-5">
                                    <Popover.Button onClick={() => props.clear()}
                                        className="-m-3 flex items-center rounded-lg p-2 transition duration-150 ease-in-out hover:bg-gray-50 focus:outline-none focus-visible:ring focus-visible:ring-orange-500 focus-visible:ring-opacity-50"
                                    >
                                        <div className="ml-2">
                                            <p className="text-sm font-medium text-gray-900">
                                                Clear
                                            </p>
                                        </div>
                                    </Popover.Button>
                                </div>
                                <div className="relative grid bg-colorMainLight p-5">
                                    <Popover.Button onClick={() => props.reset()}
                                        className="-m-3 flex items-center rounded-lg p-2 transition duration-150 ease-in-out hover:bg-gray-50 focus:outline-none focus-visible:ring focus-visible:ring-orange-500 focus-visible:ring-opacity-50"
                                    >
                                        <div className="ml-2">
                                            <p className="text-sm font-medium text-gray-900">
                                                Reset
                                            </p>
                                        </div>
                                    </Popover.Button>
                                </div>
                            </div>
                        </Popover.Panel>
                    </Transition>
                </>
            )}
        </Popover>
    )
}
