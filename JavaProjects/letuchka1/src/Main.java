import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.stream.Collectors;

public class Main {
    public static void main(String[] args) {

        System.out.println("Пенкин Артём   Вариант №5:");

        ArrayList<Sphere> spheres = new ArrayList<>();

        spheres.add(new Sphere("A", 1));
        spheres.add(new Sphere("B", 3));
        spheres.add(new Sphere("C", 2));
        spheres.add(new Sphere("D", 4));
        spheres.add(new Sphere("E", 7));
        spheres.add(new Sphere("F", 5));
        spheres.add(new Sphere("G", 1));
        System.out.println("исходное множество сфер:");
        spheres.stream()
                .forEach(x -> System.out.print(x.name + "=" + x.radius + "; "));


        int MAX1 = 4;
        System.out.println("\n\nсферы радиуса меньше 4");
        spheres.stream()
                .filter(x -> x.radius <= MAX1)
                .forEach(x -> System.out.print(x.name + " "));

        int MAX2 = 8;
        System.out.println("\n\nсферы радиуса меньше 8");
        spheres.stream()
                .filter(x -> x.radius <= MAX2)
                .forEach(x -> System.out.print(x.name + " "));

        int MAX3 = 2;
        System.out.println("\n\nсферы радиуса меньше 2");
        spheres.stream()
                .filter(x -> x.radius <= MAX3)
                .forEach(x -> System.out.print(x.name + " "));
    }
}

class Sphere {
    public int radius;
    public String name;

    public Sphere(String name, int radius){
        this.name = name;
        this.radius = radius;
    }

}