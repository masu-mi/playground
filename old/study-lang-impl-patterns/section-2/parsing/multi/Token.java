public class Token {
    public int type;
    public String text;
    public Token(int type, String text) {this.type = type; this.text=text;}
    public String toString() {
        String tname = Lexer.getTokenName(type);
        return "<'" + text + "'," + tname + ">";
    }
}
