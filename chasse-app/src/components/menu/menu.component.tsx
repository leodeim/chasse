import { useSelector } from 'react-redux';
import { Link } from 'react-router-dom';
import { selectSessionId } from '../../state/game/game.slice';
import { getApiUrl } from '../../utilities/environment.utility';
import { HomeIcon, ShareIcon } from '../../utilities/icons.utility';
import { copyToClipboard } from '../../utilities/window.utility';
import { Popover, Transition } from '@headlessui/react'
import { Fragment } from 'react'

export default function Menu() {
    const sessionId = useSelector(selectSessionId);

    let handleShare = () => { copyToClipboard(getApiUrl() + "board/" + sessionId) }

    return (
        <div className="flex sm:flex-col justify-center sm:mr-4 text-colorSecondary">
            <div className="flex sm:flex-col mb-4 pt-2 sm:mb-0 sm:pt-4 sm:pb-4 bg-colorMainDark rounded-full">
                <Link to="/" className="sm:pb-4 pl-4 pr-4 sm:pl-2 sm:pr-2 sm:pt-2" onClick={() => { }}>
                    <HomeIcon />
                </Link>
                <button className="pl-4 pr-4 sm:pl-2 sm:pr-2">
                    <SharePopover
                        handle={handleShare}
                    />
                </button>
            </div>
        </div>
    );

}

function SharePopover(props) {
    return (
        <Popover className="relative">
            {({ open }) => (
                <>
                    <Popover.Button>
                        <ShareIcon />
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
                        <Popover.Panel className="absolute z-50 w-36">
                            <div className="overflow-hidden rounded-lg shadow-lg ring-1 ring-black ring-opacity-5">
                                <div className="relative grid gap-8 bg-colorMainLight p-6">
                                    <Popover.Button onClick={() => props.handle()}
                                        className="-m-3 flex items-center rounded-lg p-2 transition duration-150 ease-in-out hover:bg-gray-50 focus:outline-none focus-visible:ring focus-visible:ring-orange-500 focus-visible:ring-opacity-50"
                                    >
                                        <div className="ml-3">
                                            <p className="text-sm font-medium text-gray-900">
                                                Copy link
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
