function execute(command) {
    return require("child_process").exec(command);
}

export function inspect() {
    return execute(process.env.COMMAND);
}