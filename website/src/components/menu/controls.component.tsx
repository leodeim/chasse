import { Link } from 'react-router-dom';
import { HomeIcon, ShareIcon } from '../../utilities/icons.utility'

export default function Menu() {
    return (
        <div className="flex justify-center mb-4 pt-2 text-yellow bg-darkGreen rounded-full">
            <Link to="/" className="pb-2 pl-4 pr-4" onClick={() => { }}>
                <HomeIcon />
            </Link>
            <button className="pb-2 pl-4 pr-4" onClick={() => { }}>
                <ShareIcon />
            </button>
        </div>
    );
}
