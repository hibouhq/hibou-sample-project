package fixtures;

public final class Direct {
    public static Process inspect() throws java.io.IOException {
        String command = System.getenv("COMMAND");
        return Runtime.getRuntime().exec(command);
    }
}