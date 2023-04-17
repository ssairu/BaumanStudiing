public class Main {
    public static void main(String[] args) {
        Universe milkyway = new Universe("milkyway");
        milkyway.setAtom(0, 0, 0);
        milkyway.setAtom(2, 2, 2);

        Universe juiceway = new Universe("juiceway");
        juiceway.setAtom(3,3,3);
        juiceway.setAtom(5,-1,-1);


        System.out.print("расстояние между вселенными milkyway и juiceway: ");
        System.out.println(milkyway.dist(juiceway));

        System.out.print("расстояние между вселенными juiceway и milkyway: ");
        System.out.println(juiceway.dist(milkyway));
    }
}