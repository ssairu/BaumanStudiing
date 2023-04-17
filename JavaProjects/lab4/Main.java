import java.util.Iterator;
public class Main {
    public static void main(String[ ] args ) {
        StringSpace b = new StringSpace("qwer  ty");
        System.out.println("создали строку \"" + b.getS() + "\", переберём её");
        Iterator<String> it = b.iterator();
        while ( it.hasNext() ) {
            String s = it.next();
            System.out.println(s);
        }

        StringSpace a = new StringSpace("a  b  c  d ef");
        System.out.println("создали строку \"" + a.getS() + "\"");
        System.out.println("обновим строку \"" + a.getS() + "\",");
        System.out.println("вставив на 1 позицию строку \" х \"");
        a.insert(1, "x");

        System.out.println("теперь строка \"" + a.getS() + "\", переберём её");
        Iterator<String> it1 = a.iterator();
        while ( it1.hasNext() ) {
            String s = it1.next();
            System.out.println(s);
        }
    }
}