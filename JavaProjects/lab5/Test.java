public class Test {
    public static void main(String[] args) {
        MatrixBool A = new MatrixBool(new int[][]{
                {1, 1, 1, 0},
                {1, 0, 1, 0},
                {0, 1, 1, 1},
                {0, 1, 1, 0},
                {0, 0, 0, 1}});
        MatrixBool B = new MatrixBool(new int[][]{
                {1, 1, 1, 1, 1, 1, 0, 1},
                {0, 0, 0, 1, 0, 0, 0, 1},
                {0, 0, 0, 0, 0, 0, 1, 0},
                {0, 0, 1, 0, 0, 0, 0, 0}});

        System.out.println("Создадим матрицу 1:");
        A.printMatrix();
        System.out.println("");
        System.out.println("Создадим матрицу 2:");
        B.printMatrix();
        System.out.println("");

        System.out.println("Выведем через поток исключающее или");
        System.out.println("для каждой из строк матрицы 1 и 2:");
        System.out.print("**");
        A.xorStream().forEach(x -> System.out.print(" " + x + " "));
        System.out.println("**");
        System.out.print("**");
        B.xorStream().forEach(x -> System.out.print(" " + x + " "));
        System.out.println("**");

        System.out.println("Посчитаем количество \"1\" и \"0\" в этих потоках:");
        System.out.print("\"1\" - " + A.xorStream().filter(x -> x == 1).count());
        System.out.println("; \"0\" - " + A.xorStream().filter(x -> x == 0).count());
        System.out.print("\"1\" - " + B.xorStream().filter(x -> x == 1).count());
        System.out.println("; \"0\" - " + B.xorStream().filter(x -> x == 0).count() + "\n");

        System.out.println("Найдём строку, матрицы 1, в которой \"1\"");
        System.out.println("больше чем в остальных строках вместе взятых:");
        if (A.getBestString().isPresent()) {
            System.out.println(A.getBestString().get().getString() + "\n");
        }
        else {
            System.out.println("такой строки не существует.\n");
        }

        System.out.println("Найдём строку, матрицы 2, в которой \"1\"");
        System.out.println("больше чем в остальных строках вместе взятых:");
        if (B.getBestString().isPresent()) {
            System.out.println(B.getBestString().get().getString() + "\n");
        }
        else {
            System.out.println("такой строки не существует.");
        }
    }
}
