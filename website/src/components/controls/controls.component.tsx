import { BackIcon, EndIcon, ReverseIcon } from '../../utilities/icons.utility'

export default function Controls(props: ControlsProps) {
    return (
        <div className="flex justify-center mt-4 pb-2 text-darkGreen bg-yellow rounded-full">
            <button className="pt-2 pl-4 pr-4 " onClick={() => props.back()}>
                <BackIcon />
            </button>
            <button className="pt-2 pl-4 pr-4" onClick={() => props.reverse()}>
                <ReverseIcon />
            </button>
            <button className="pt-2 pl-4 pr-4" onClick={() => props.reset()}>
                <EndIcon />
            </button>
        </div>
    );
}

interface ControlsProps {
    back(): any;
    reverse(): any;
    reset(): any;
}