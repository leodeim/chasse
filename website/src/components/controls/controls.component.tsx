import { BackIcon, EndIcon, ReverseIcon, ShareIcon } from '../../utilities/icons.utility'

export default function Controls(props: ControlsProps) {
    return (
        <div className="flex justify-center pt-2">
            <button className="pt-2 pl-4 pr-4 text-zensquare" onClick={() => props.back()}>
                <BackIcon />
            </button>
            <button className="pt-2 pl-4 pr-4 text-zensquare" onClick={() => props.reverse()}>
                <ReverseIcon />
            </button>
            <button className="pt-2 pl-4 pr-4 text-zensquare" onClick={() => { }}>
                <ShareIcon />
            </button>
            <button className="pt-2 pl-4 pr-4 text-zensquare" onClick={() => props.reset()}>
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