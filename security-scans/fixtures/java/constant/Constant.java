package fixtures;

public final class Constant {
    public static Process inspect() throws java.io.IOException {
        String ignored = System.getenv("COMMAND");
        return Runtime.getRuntime().exec("echo fixture");
    }
}