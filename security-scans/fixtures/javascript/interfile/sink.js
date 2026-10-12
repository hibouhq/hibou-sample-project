export function execute(command) {
    return require("child_process").exec(command);
}