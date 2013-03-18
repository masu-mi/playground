public class LookaheadParser extends Parser {
    public LookaheadParser(LookaheadLexer input) { super(input, 2); }

    public void list() {
        match(LookaheadLexer.LBRACK); elements(); match(LookaheadLexer.RBRACK);
    }

    void elements() {
        element();
        while ( LA(1) == LookaheadLexer.COMMA) {
            match(LookaheadLexer.COMMA); element();
        }
    }

    void element() {

        if ( LA(1) == LookaheadLexer.NAME && LA(2) == LookaheadLexer.EQUALS) {
            match(LookaheadLexer.NAME);
            match(LookaheadLexer.EQUALS);
            match(LookaheadLexer.NAME);
        }
        else if ( LA(1) == LookaheadLexer.NAME) match(LookaheadLexer.NAME);
        else if ( LA(1) == LookaheadLexer.LBRACK) list();
        else throw new Error("expecting name or list; found " + lookahead);
    }
}
