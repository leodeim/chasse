import { Link } from 'react-router-dom';
import { selectSessionId } from '../state/game/game.slice';
import { getApiUrl } from '../utilities/environment.utility';
import { HomeIcon, ShareIcon } from '../utilities/icons.utility';
import { copyToClipboard } from '../utilities/window.utility';
import { useAppSelector } from '../state/hooks';
import QuickMenu, { Direction, QuickMenuButtonProps } from './quickmenu.component';

export default function Menu() {
    const sessionId = useAppSelector(selectSessionId);

    let shareButton: QuickMenuButtonProps[] = [
        {
            text: "Copy link",
            handler: () => copyToClipboard(getApiUrl() + "board/" + sessionId),
        },
    ]

    return (
        <div className="flex sm:flex-col justify-center sm:mr-4 text-colorSecondary">
            <div className="flex sm:flex-col mb-4 pt-2 sm:mb-0 sm:pt-4 sm:pb-4 bg-colorMainDark rounded-full">
                <Link to="/" className="sm:pb-4 pl-4 pr-4 sm:pl-2 sm:pr-2 sm:pt-2" onClick={() => { }}>
                    <HomeIcon />
                </Link>
                <div className="pl-4 pr-4 sm:pl-2 sm:pr-2">
                    <QuickMenu
                        direction={Direction.Down}
                        icon={<ShareIcon />}
                        buttons={shareButton}
                    />
                </div>
            </div>
        </div>
    );

}
