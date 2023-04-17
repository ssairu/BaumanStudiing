public class Test {
    public static void main(String[] args) {
        int n = 5;
        System.out.println("пусть количество городов: 5");
        Cities loc1 = new Cities(n);
        System.out.println();
        System.out.println("начальная матрица:");
        loc1.outMatrix();
        System.out.println();
        System.out.println("запишем некоторые дороги");
        loc1.setRoad(1, 2, 25);
        loc1.setRoad(1, 0, 10);
        loc1.setRoad(1, 4, 65);
        loc1.setRoad(3, 4, 3);
        loc1.setRoad(0, 2, 1);
        loc1.setRoad(2, 3, 12);
        System.out.println();
        System.out.println("новая матрица городов: ");
        loc1.outMatrix();
        System.out.println();
        int [] a = new int[] {0, 3, 2, 3};
        int [] b = new int[] {0, 1, 2, 3};
        int path1 = loc1.path(a);
        if (path1 > 0)
            System.out.println("этот путь равен: " + path1);
        System.out.println();
        int path2 = loc1.path(b);
        if (path2 > 0)
            System.out.println("этот путь равен: " + path2);
    }

}
