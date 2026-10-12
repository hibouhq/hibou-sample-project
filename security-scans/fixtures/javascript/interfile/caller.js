import { execute } from "./sink.js";

export function inspect() {
    return execute(process.env.COMMAND);
}