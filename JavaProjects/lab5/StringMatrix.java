import java.util.Arrays;

public class StringMatrix {
    private long str, sumBinary;
    int count;
    public StringMatrix(int[] a) {
        str = 0;
        Arrays.stream(a).forEach(y-> {
            str = (str << 1) + y;
            sumBinary += y;});
        count = a.length;
    }

    public long getSumBinary() {
        return sumBinary;
    }

    public long getStr() {
        return str;
    }
    public String getString(){
        String a = Long.toBinaryString(str);
        if (a.length() < count){
            a = "0".repeat(count - a.length()) + a;
        }
        return a;
    }
}
