public class BacktrackParser extends Parser {
    public BacktrackParser(BacktrackLexer input) { super(input); }


    public void stat() throws RecognitionException {
        if (speculate_stat_alt1()) {
            list(); match(Lexer.EOF_TYPE);
        }
        else if (speculate_stat_alt2()) {
            assign(); match(Lexer.EOF_TYPE);
        }
        else throw new NoViableAltException("expecting stat found " + LT(1));

    }

    public boolean speculate_stat_alt1() {
        boolean success = true;
        mark();
        try {list(); match(Lexer.EOF_TYPE); }
        catch (RecognitionException e) { success = false; }
        release();
        return success;
    }

    public boolean speculate_stat_alt2() {
        boolean success = true;
        mark();
        try {assign(); match(Lexer.EOF_TYPE); }
        catch (RecognitionException e) { success = false; }
        release();
        return success;
    }

    public void list() {
        match(BacktrackLexer.LBRACK); elements(); match(BacktrackLexer.RBRACK);
    }

    public void assign() {
        list(); match(BacktrackLexer.EQUALS); list();
    }

    void elements() {
        element();
        while ( LA(1) == BacktrackLexer.COMMA) {
            match(BacktrackLexer.COMMA); element();
        }
    }

    void element() {

        if ( LA(1) == BacktrackLexer.NAME &&
                LA(2) == BacktrackLexer.EQUALS &&
                LA(3) == BacktrackLexer.LBRACK) {
            match(BacktrackLexer.NAME);
            match(BacktrackLexer.EQUALS);
            list();
        }
        else if ( LA(1) == BacktrackLexer.NAME && LA(2) == BacktrackLexer.EQUALS) {
            match(BacktrackLexer.NAME);
            match(BacktrackLexer.EQUALS);
            match(BacktrackLexer.NAME);
        }
        else if ( LA(1) == BacktrackLexer.NAME) match(BacktrackLexer.NAME);
        else if ( LA(1) == BacktrackLexer.LBRACK) list();
        else throw new Error("expecting name or list; found " + lookahead);
    }
}
