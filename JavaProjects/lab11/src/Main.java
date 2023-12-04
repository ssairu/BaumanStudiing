public class Main {
    public static void main(String[] args) throws MySyntaxException {
        String b = "\" \"hg=((zxfnv))";
        String a = "\"abc\" x (\"a\" \"bd\" y) =\n" +
                "(x y () (\"qwerty\"))";
        System.out.println("string a:\n" + a + "\n");
        System.out.println("string b:\n" + b + "\n");

        System.out.println("lexems in a: " + Parser.lexemCount(a) + "\nand them:");
        Parser.lexemPrint(a);


        System.out.println("\n\n\nparse a:");
        if(Parser.parse(a)){
            Parser.parsePrints(a);
        }


        System.out.println("\n\n\nparse b:");
        if(Parser.parse(b)){
            Parser.parsePrints(b);
        }

    }
}