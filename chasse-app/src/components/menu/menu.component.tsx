import { useSelector } from 'react-redux';
import { Link } from 'react-router-dom';
import { selectSessionId } from '../../state/game/game.slice';
import { getApiUrl } from '../../utilities/environment.utility';
import { HomeIcon, ShareIcon } from '../../utilities/icons.utility';
import { copyToClipboard } from '../../utilities/window.utility';

export default function Menu() {
    const sessionId = useSelector(selectSessionId);
    
    let handleShare = () => { copyToClipboard(getApiUrl() + "board/" + sessionId) }

    return (
        <div className="flex sm:flex-col justify-center sm:mr-4 text-colorSecondary">
            <div className="flex sm:flex-col mb-4 pt-2 sm:mb-0 sm:pt-4 sm:pb-4 bg-colorMainDark rounded-full">
                <Link to="/" className="pb-2 pl-4 pr-4 sm:pl-2 sm:pr-2 sm:pt-2" onClick={() => { }}>
                    <HomeIcon />
                </Link>
                <button className="pb-2 pl-4 pr-4 sm:pl-2 sm:pr-2" onClick={() => handleShare()}>
                    <ShareIcon />
                </button>
            </div>
        </div>
    );
}
