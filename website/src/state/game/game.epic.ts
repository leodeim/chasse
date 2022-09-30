import { Action } from '@reduxjs/toolkit'
import { Observable } from 'rxjs'
import { filter, mergeAll, map } from 'rxjs/operators'
import { makeMove, makeMoveSuccessful, MoveItem } from './game.slice'
import { SendWebsocketMove } from '../../socket/socket'

const writeMoveToSocket = async (move: MoveItem): Promise<MoveItem> => {
    SendWebsocketMove({
        sessionId: move.sessionId,
        fen: move.position
    })

    return move
}

export const makeMoveEpic = (action$: Observable<Action>) =>
    action$.pipe(
        filter(makeMove.match),
        map(({ payload }) =>
            writeMoveToSocket(payload)
            .then(makeMoveSuccessful)
        ),
        mergeAll()
    )
