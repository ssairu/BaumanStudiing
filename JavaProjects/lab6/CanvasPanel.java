import javax.swing.*;
import java.awt.*;

public class CanvasPanel extends JPanel {
    private int diameter = 7, hight = 20, percent = 95;

    public void setDiameter(int newd){
        diameter = newd;
        repaint();
    }
    public void setHight(int newd){
        hight = newd;
        repaint();
    }
    public void setPercent(int newd){
        percent = newd;
        repaint();
    }
    protected void paintComponent (Graphics g) {
        super.paintComponent(g);
        g.setColor(Color.green);
        g.fillRoundRect(185, 10, 30, 30, 1, 1);
        g.fillRoundRect(200 - diameter * 5, 40,
                diameter * 10, hight * 10, 10, 5);
        g.setColor(Color.pink);
        g.fillRoundRect(200 - diameter * 5,
                40 + (int)((float)hight * ((float)(100 - percent) / 10)),
                diameter * 10,
                (int)((float)hight * percent / 10), 10 ,5);
    }
}