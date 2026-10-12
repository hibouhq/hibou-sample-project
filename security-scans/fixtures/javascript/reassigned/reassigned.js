export function inspect() {
    let command = process.env.COMMAND;
    command = "echo fixture";
    return require("child_process").exec(command);
}