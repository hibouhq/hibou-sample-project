package com.hibou.sample;

public final class CommandController {

    public void handle(Request request) throws java.io.IOException {
        CommandExecutor.run(request.getParameter("command"));
    }

    public void handleConstant() throws java.io.IOException {
        CommandExecutor.run("echo fixture");
    }
}
