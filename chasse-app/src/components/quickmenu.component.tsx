import { Popover, Transition } from "@headlessui/react";
import { Fragment, useCallback } from "react";

export default function QuickMenu(props: QuickMenuProps) {
    const list = props.buttons.map((button, key)=> 
        <QuickMenuButton
            key={button.id}
            id={button.id}
            handler={button.handler}
            text={button.text}
        />)

    return (
        <Popover className="relative flex">
            {() => (
                <>
                    <Popover.Button>
                        {props.icon}
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
                        <Popover.Panel className={props.direction}>
                            <div className="overflow-hidden rounded-lg shadow-lg ring-1 ring-black ring-opacity-5">
                                {list}
                            </div>
                        </Popover.Panel>
                    </Transition>
                </>
            )}
        </Popover>
    )
}

function QuickMenuButton(props: QuickMenuButtonProps) {
    let handle = useCallback(() => { props.handler() }, [props])

    return (
        <div className="relative grid gap-8 bg-colorMainLight p-4">
            <Popover.Button onClick={handle}
                className="-m-3 flex items-center rounded-lg p-2 pl-0 transition duration-150 ease-in-out hover:bg-gray-50 focus:outline-none focus-visible:ring focus-visible:ring-orange-500 focus-visible:ring-opacity-50"
            >
                <div className="ml-3">
                    <p className="text-sm font-medium text-gray-900">
                        {props.text}
                    </p>
                </div>
            </Popover.Button>
        </div>
    )
}

export enum Direction {
    UpRight = "absolute left-full bottom-full z-50 w-28",
    UpLeft = "absolute right-full bottom-full z-50 w-28",
    DownRight = "absolute left-full top-full z-50 w-28",
    DownLeft = "absolute right-full top-full z-50 w-28"
}

export type QuickMenuProps = {
    direction: Direction
    icon: JSX.Element
    buttons: QuickMenuButtonProps[]
}

export type QuickMenuButtonProps = {
    id: number
    text: string
    handler: Function
}
