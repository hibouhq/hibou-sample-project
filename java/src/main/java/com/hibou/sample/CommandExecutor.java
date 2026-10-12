package com.hibou.sample;

public final class CommandExecutor {

    private CommandExecutor() {
    }

    public static Process run(String command) throws java.io.IOException {
        return Runtime.getRuntime().exec(command);
    }
}
