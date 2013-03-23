import java.util.*;


public abstract class Parser {

    Lexer input;
    List<Integer> markers;
    List<Token> lookahead;
    int p = 0;
    public static int FAILED = -1;

    public Parser(Lexer input) {
        this.input = input;
        lookahead = new ArrayList<Token>();
        markers   = new ArrayList<Integer>();

        Token.setLexer(input);
    }


    public Token LT(int i) { sync(i); return lookahead.get(p+i-1); }
    public int LA(int i) { return LT(i).type; }
    public void match(int x) {
        if (LA(1) == x) consume();
        else throw new MismatchedTokenException("expecting " + input.getTokenName(x) + "; found " + LT(1));
    }

    public void sync(int i) {
        if ( p+i-1 > (lookahead.size()-1) ) {
            int n = (p+i-1) - (lookahead.size()-1);
            fill(n);
        }
    }
    public void fill(int n) {
        for (int i=1; i<=n; i++) {
            lookahead.add(input.nextToken());
        }
    }

    public void consume() {
        p++;
        if ( p==lookahead.size() && !isSpeculating() ) {
            p=0;
            lookahead.clear();
            clearMemo();
        }
        sync(1);
    }
    public int mark() { markers.add(p); return p; }
    public void release() {
        int marker = markers.get(markers.size()-1);
        markers.remove(markers.size()-1);
        seek(marker);
    }
    public void seek(int index) { p=index; }
    public boolean isSpeculating() { return markers.size() > 0; }


    public boolean alreadyParsedRule(Map<Integer, Integer> memorization)
        throws PreviousParseFailedException
    {
        Integer memoI = memorization.get(index());
        if ( memoI==null ) { return false; }
        int memo = memoI.intValue();
        System.out.println("parsed list before at index " + index() +
                "; skip ahead to token index " + memo + ": " + lookahead.get(memo).text);
        if (memo==FAILED) throw new PreviousParseFailedException("parsed failed previously");
        seek(memo);
        return true;
    }

    public void memorize(Map<Integer, Integer> memorization,
            int startTokenIndex, boolean failed)
    {
        int stopTokenIndex = failed?FAILED:index();
        memorization.put(startTokenIndex, stopTokenIndex);
    }
    public int index() { return p; }

    public abstract void clearMemo();
}
