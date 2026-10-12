package com.hibou.sample;

public final class Request {

    private final String command;

    public Request(String command) {
        this.command = command;
    }

    public String getParameter(String name) {
        return command;
    }
}
