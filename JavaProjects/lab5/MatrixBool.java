import java.util.Arrays;
import java.util.*;
import java.util.stream.Stream;


public class MatrixBool {
    private ArrayList<StringMatrix> A;
    long n;
    public MatrixBool(int[][] A1) {
        A = new ArrayList<>(1);
        Arrays.stream(A1)
                .map(x -> new StringMatrix(x))
                .forEach(x -> A.add(x));
        n = A.stream()
                .map(x -> x.getSumBinary())
                .reduce((x, y) -> x + y).get();
    }

    public StringMatrix getStringMatrix(int i){
        return A.get(i);
    }

    public long getN() {
        return n;
    }

    public void printMatrix(){
        A.stream().map(x -> x.getString())
                .map(x -> x.split(""))
                .forEach(x -> {
                    Arrays.stream(x)
                            .forEach(y -> System.out.print(y + " "));
                    System.out.println("");
                });
    }

    public Stream<Long> xorStream() {
        ArrayList<Long> result = new ArrayList<>();
        A.stream()
                .map(x -> x.getSumBinary() % 2)
                .forEach(x -> result.add(x));
        return result.stream();
    }

    public Optional<StringMatrix> getBestString() {
        Optional<StringMatrix> result = Optional.empty();
        Optional<StringMatrix> tmp = A
                .stream()
                .filter(x -> x.getSumBinary() > n - x.getSumBinary())
                .findFirst();
        if (tmp.isPresent()) {
            result = Optional.ofNullable(tmp.get());
        }
        return result;
    }
}
