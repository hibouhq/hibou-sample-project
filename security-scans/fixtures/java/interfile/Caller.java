package fixtures;

public final class Caller {
    public static Process inspect() throws java.io.IOException {
        return Sink.execute(System.getenv("COMMAND"));
    }
}