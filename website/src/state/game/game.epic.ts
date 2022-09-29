import { Action } from '@reduxjs/toolkit'
import { Observable, interval } from 'rxjs'
import { mergeMap, filter, debounce } from 'rxjs/operators'
import { makeMove } from './game.slice'
import { SendWebsocketMove } from '../../socket/socket'

const writeMoveToSocket = async (position: string) => {
    console.log("writeMoveToSocket from epic")
    SendWebsocketMove({
        sessionId: 'a2e5b7a2-7a01-416a-be9a-40dd25bd0c7b',
        fen: position
    })
}

export const makeMoveEpic = (action$: Observable<Action>) =>
    action$.pipe(
        filter(makeMove.match),
        debounce(() => interval(1000)),
        mergeMap(({ payload }) =>
            writeMoveToSocket(payload)
        )
    )