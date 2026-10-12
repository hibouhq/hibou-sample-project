package fixtures;

public final class Sanitized {
    private static String allowCommand(String input) {
        return "echo fixture";
    }

    public static Process inspect() throws java.io.IOException {
        return Runtime.getRuntime().exec(allowCommand(System.getenv("COMMAND")));
    }
}