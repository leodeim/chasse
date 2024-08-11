import axios, { AxiosResponse } from "axios";
import { getApiUrl } from "../utilities/environment.utility";
import * as Joi from 'joi';

export type SessionData = {
    action?: number,
    response?: number,
    position: string,
    sessionId: string
}

export function getSession(id: string, ok: Function, nok: Function) {
    axios
        .get(getApiUrl() + `api/v1/session/` + id)
        .then((response: AxiosResponse) => {
            const schema = Joi.object({
                action: Joi.number().optional(),
                response: Joi.number().optional(),
                position: Joi.string().required(),
                sessionId: Joi.string().required()
            });
            const { error } = schema.validate(response.data);

            if (error) nok(error, response.status);
            else ok(response.status, response.data);
        })
        .catch((err) => {
            nok(err);
        });
}

export function newSession(ok: Function, nok: Function) {
    axios
        .get(getApiUrl() + `api/v1/session`)
        .then((response: AxiosResponse) => {
            const schema = Joi.object({
                action: Joi.number().optional(),
                response: Joi.number().optional(),
                position: Joi.string().required(),
                sessionId: Joi.string().required()
            });
            const { error } = schema.validate(response.data);

            if (error) nok(error, response.status);
            else ok(response.status, response.data);
        })
        .catch((err) => nok(err));
}