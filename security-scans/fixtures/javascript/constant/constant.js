export function inspect() {
    const ignored = process.env.COMMAND;
    return require("child_process").exec("echo fixture");
}