function allowCommand(value) {
    return "echo fixture";
}

export function inspect() {
    return require("child_process").exec(allowCommand(process.env.COMMAND));
}