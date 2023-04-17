import java.util.Iterator;

public class StringSpace implements Iterable{
    private String s;

    public StringSpace(String name){
        this.s = name;
    }

    public Iterator iterator() {
        return new PairIterator();
    }

    public void insert(int i, String cap){
        this.s = s.substring(0, i) + cap + s.substring(i, s.length());
    }

    public String getS() {
        return s;
    }

    private class PairIterator implements Iterator{
        private int pos;
        private String str;
        public PairIterator() {
            pos = 0;
            str = s.replaceAll("\\s", "");
        }

        public boolean hasNext() {
            return pos < str.length() - 1;
        }

        public String next() {
            pos++;
            return str.substring(pos - 1, pos + 1);
        }
    }
}
