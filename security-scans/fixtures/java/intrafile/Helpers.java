package fixtures;

public final class Helpers {
    private static Process execute(String command) throws java.io.IOException {
        return Runtime.getRuntime().exec(command);
    }

    public static Process inspect() throws java.io.IOException {
        return execute(System.getenv("COMMAND"));
    }
}