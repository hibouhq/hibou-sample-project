package fixtures;

public final class Loop {
    public static Process inspect(int count) throws java.io.IOException {
        String command = "echo fixture";
        for (int index = 0; index < count; index++) {
            command = System.getenv("COMMAND");
        }
        return Runtime.getRuntime().exec(command);
    }
}