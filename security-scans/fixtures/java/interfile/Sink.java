package fixtures;

public final class Sink {
    public static Process execute(String command) throws java.io.IOException {
        return Runtime.getRuntime().exec(command);
    }
}