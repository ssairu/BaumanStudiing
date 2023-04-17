import static java.lang.Math.*;
public class Point
{
    private String name;
    private double x;
    private double y;
    private double z;
    public Point(){}
    public Point(double varX, double varY, double varZ)
    {
        x=varX;
        y=varY;
        z=varZ;
    }
    public void setCoord(double varX, double varY, double varZ)
    {
        this.x=varX;
        this.y=varY;
        this.z=varZ;
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
    public double getR()
    {
        return pow(pow(this.x,2)+pow(this.y,2)+pow(this.z,2),0.5);
    }
    public double dist(Point a){
        return pow(pow((this.x - a.getX()),2)+pow((this.y - a.getY()),2)+pow((this.z - a.getZ()),2),0.5);
    }
}
