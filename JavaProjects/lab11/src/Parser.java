public class Parser {

    private static String sym, res;
    private static int index, line, col;
    private static boolean prints = false;


    private static void setStartValues(){
        index = col = 0;
        line = 1;
        res = "<Stmt>";
    }

    private static void printRes(){
        if (prints == true) {
            System.out.println(res);
        }
    }

    private static void next(String str) {
        if (str.length() == index) {
            sym = "";
        } else if (Character.isLetter(str.charAt(index))) {
            int i = index;
            while (i < str.length() && (Character.isLetter(str.charAt(i)) || Character.isDigit(str.charAt(i)))) {
                i++;
            }
            sym = str.substring(index, i);
            index = i;
            col++;
        } else if (str.charAt(index) == '"') {
            int x;
            x = str.indexOf('"', index + 1);
            sym = str.substring(index, x + 1);
            index = x + 1;
            col++;
        } else if (str.substring(index, index + 1).equals("\n")) {
            line++;
            col = 0;
            index++;
            next(str);
        } else if (str.substring(index, index + 1).equals(" ")) {
            index++;
            next(str);
        } else {
            col++;
            sym = str.substring(index, ++index);
        }
    }

    public static int lexemCount(String str){
        setStartValues();
        next(str);
        int i = 0;
        while (!sym.equals("")){
            next(str);
            i++;
        }
        setStartValues();
        return i;
    }

    public static void lexemPrint(String str){
        setStartValues();
        int count = lexemCount(str);
        for (int i = 0; i < count; i++) {
            next(str);
            System.out.println(String.format("%s  (%d, %d)", sym, line, col));
        }
        setStartValues();
    }

    public static void parsePrints(String str) throws MySyntaxException {
        prints = true;
        parse(str);
        prints = false;
    }

    public static boolean parse(String str) throws MySyntaxException {
        setStartValues();
        printRes();
        next(str);
        parseStmt(str);

        if(!sym.equals("")){
            throw new MySyntaxException("illegal symbol or sequence of symbols started", line, col);
        } else {return true;}
    }

    private static void parseStmt(String str) throws MySyntaxException {
        res = "<Expr> = <Expr>";
        printRes();

        parseExpr(str);
        if (!sym.equals("=")) {
            throw new MySyntaxException("no \"=\", but it can to be", line, col);
        } else {
            next(str);
        }
        parseExpr(str);
    }

    private static void parseExpr(String str) throws MySyntaxException {
        if (!sym.equals("") && (sym.equals("(") || sym.substring(0, 1).equals("\"") || Character.isLetter(sym.charAt(0)))) {
            res = res.replaceFirst("<Expr>", "<Atom> <Expr>");
            printRes();

            parseAtom(str);
            parseExpr(str);
        } else {
            res = res.replaceFirst("<Expr>", "");
            printRes();
        }
    }

    private static void parseAtom(String str) throws MySyntaxException {
        if (sym.equals("(")) {
            res = res.replaceFirst("<Atom>", "( <Expr> )");
            printRes();

            next(str);
            parseExpr(str);
            if (!sym.equals(")")){
                throw new MySyntaxException("there should be \")\"", line, col);
            }
            next(str);
        } else if (sym.substring(0, 1).equals("\"") || Character.isLetter(sym.charAt(0))){
            res = res.replaceFirst("<Atom>", sym);
            printRes();

            next(str);
        } else {
            throw new MySyntaxException("illegal lexem", line, col);
        }
    }
}
