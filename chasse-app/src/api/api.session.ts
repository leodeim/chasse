import axios, { AxiosResponse } from "axios";
import { getApiUrl } from "../utilities/environment.utility";

export function getSession(id: string, ok: Function, nok: Function) {
    axios
        .get(getApiUrl() + "api/v1/session/" + id)
        .then((response: AxiosResponse) => {
            if (response.status === 200) {
                ok(response.data);
            } else {
                nok();
            }
        })
        .catch((err) => {
            nok(err);
        });
}

export function newSession(ok: Function, nok: Function) {
    axios
        .get(getApiUrl() + "api/v1/session/new")
        .then((response: AxiosResponse) => {
            if (response.data.sessionId !== undefined) {
                ok(response.data.sessionId)
            } else {
                nok()
            }
        })
        .catch((err) => nok(err));
}