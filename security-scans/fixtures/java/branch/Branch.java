package fixtures;

public final class Branch {
    public static Process inspect(boolean useInput) throws java.io.IOException {
        String command = "echo fixture";
        if (useInput) {
            command = System.getenv("COMMAND");
        }
        return Runtime.getRuntime().exec(command);
    }
}