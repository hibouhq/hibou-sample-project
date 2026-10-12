export function inspect(useInput) {
    let command = "echo fixture";
    if (useInput) {
        command = process.env.COMMAND;
    }
    return require("child_process").exec(command);
}