package fixtures;

public final class Reassigned {
    public static Process inspect() throws java.io.IOException {
        String command = System.getenv("COMMAND");
        command = "echo fixture";
        return Runtime.getRuntime().exec(command);
    }
}