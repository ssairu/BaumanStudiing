public class Figure implements Comparable<Figure >{
    private int count;
    private double[] sizes = new double[10000];
    private double max;
    public Figure(double[] x){
        for (int i = 0; i < x.length; i++) {
            sizes[i] = x[i];
            if (i == 0){
                max = x[i];
            }
            else if (max < x[i]) {
                max = x[i];
            }
        }
        count = x.length;
    }

    public int compareTo(Figure o) {
        if (this.max > o.max){
            return 1;
        } else if (this.max < o.max) {
            return -1;
        } else {
            return 0;
        }
    }
    public String toString(){
        String res = "values:";
        for (int i = 0; i < count; i++){
            res += String.format(" %f", this.sizes[i]);
        }
        res +=String.format("; n: %d; ", this.count) + String.format("max: %f", max);
        return res;
    }
}
