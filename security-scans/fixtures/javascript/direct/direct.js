export function inspect() {
    const command = process.env.COMMAND;
    return require("child_process").exec(command);
}