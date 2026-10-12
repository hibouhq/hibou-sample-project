package fixtures;

public class Parent {
    public Process execute(String command) throws java.io.IOException {
        return Runtime.getRuntime().exec(command);
    }
}