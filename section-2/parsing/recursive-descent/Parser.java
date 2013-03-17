public abstract class Parser {
    Lexer input;
    Token lookahead;

    public Parser(Lexer input_) {
        input = input_;
        lookahead = input.nextToken();
    }

    public void match(int x) {
        if (lookahead.type == x) consume();
        else throw new Error("excepting " + input.getTokenName(x) +
                "; found " + lookahead);
    }
    public void consume() { lookahead = input.nextToken(); }
}
