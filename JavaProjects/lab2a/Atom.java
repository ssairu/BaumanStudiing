public class Atom {
    private double x, y, z;
    public static int count = 0;
    public Atom(double a, double b, double c){
        x = a;
        y = b;
        z = c;
        count++;
    }
    public double getX(){
        return this.x;
    }
    public double getY(){
        return this.y;
    }
    public double getZ(){
        return this.z;
    }
}
