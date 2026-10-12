package fixtures;

public final class Caller {
    public static Process inspect() throws java.io.IOException {
        return new Child().execute(System.getenv("COMMAND"));
    }
}