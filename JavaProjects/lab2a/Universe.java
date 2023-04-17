public class Universe {

    private int including = 0;
    private static int count = 0;
    private String name;
    private Atom [] parts = new Atom[10000];
    public Point center = new Point();
    public Universe(String name1) {
        name = name1;
        count++;
        System.out.println("создание новой вселенной <<" + name + ">>");
    }
    public void setAtom(double a, double b, double c){
        Atom n = new Atom(a, b, c);
        System.out.println("создание тела во вселенной" + this.name + " с координатами: " + a + ", " + b + ", " + c);
        parts[including] = n;
        including++;
        center.setCoord(0, 0, 0);
        for (int i = 0; i < including; i++){
            center.setCoord(center.getX() + parts[i].getX(),
                    center.getY() + parts[i].getY(),
                    center.getZ() + parts[i].getZ());
        }
        center.setCoord(center.getX()/including, center.getY()/including, center.getZ()/including);
    }

    public Point getCenter(){
        return center;
    }

    public double dist(Universe a){
        return center.dist(a.center);
    }

}
