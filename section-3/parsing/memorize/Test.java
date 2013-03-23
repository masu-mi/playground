public class Test {
    public static void main(String[] args) {
        BacktrackLexer lexer = new BacktrackLexer(args[0]);
        BacktrackParser parser = new BacktrackParser(lexer);
        parser.stat();
    }
}
