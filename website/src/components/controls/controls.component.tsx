import { goBack, resetBoard, reverseBoard } from '../../state/game/game.slice';
import { useAppDispatch } from '../../state/hooks';
import { BackIcon, EndIcon, ReverseIcon } from '../../utilities/icons.utility'

export default function Controls() {
    const dispatch = useAppDispatch();
    
    return (
        <div className="flex justify-center mt-4 pb-2 text-darkGreen bg-yellow rounded-full">
            <button className="pt-2 pl-4 pr-4 " onClick={() => dispatch(goBack())}>
                <BackIcon />
            </button>
            <button className="pt-2 pl-4 pr-4" onClick={() => dispatch(reverseBoard())}>
                <ReverseIcon />
            </button>
            <button className="pt-2 pl-4 pr-4" onClick={() => dispatch(resetBoard())}>
                <EndIcon />
            </button>
        </div>
    );
}
