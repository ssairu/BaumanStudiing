public class MySyntaxException extends Exception{
    public MySyntaxException(String message, int line, int col){
        super(String.format("Error %s in LINE: %d, COL: %d     <3", message, line, col));
    }
}
