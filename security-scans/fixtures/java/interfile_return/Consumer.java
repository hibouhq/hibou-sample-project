package fixtures;

public final class Consumer {
    public static Process inspect() throws java.io.IOException {
        return Runtime.getRuntime().exec(Source.read());
    }
}