package fixtures;

public final class Source {
    public static String read() {
        return System.getenv("COMMAND");
    }
}