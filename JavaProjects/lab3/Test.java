import java.util.Arrays ;

public class Test {
    public static void main(String[] args) {
        Figure[] array1 = new Figure[3];
        Figure A = new Figure(new double[] {1, 2, 2, 4});
        Figure B = new Figure(new double[] {1, 2, 2, 3});
        Figure C = new Figure(new double[] {2, 3});
        array1[0] = A;
        array1[1] = B;
        array1[2] = C;
        Arrays.sort(array1);
        for (int i = 0; i < array1.length; i++){
            System.out.println(array1[i].toString());
        }
    }
}
