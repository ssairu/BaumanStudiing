import javax.swing.*;
import javax.swing.event.ChangeEvent;
import javax.swing.event.ChangeListener;


public class PictureForm {
    private JPanel mainPanel;
    private JSpinner ширинаSpinner;
    private JSpinner высотаSpinner;
    private JSpinner процентЗаполненностиSpinner;
    private JTextPane percentPanel;
    private CanvasPanel canvasPanel1;

    public PictureForm() {
        процентЗаполненностиSpinner.addChangeListener(new ChangeListener() {
            @Override
            public void stateChanged(ChangeEvent e) {
                int percent = (int)процентЗаполненностиSpinner.getValue();
                if (percent >= 100) {
                    процентЗаполненностиSpinner.setValue(100);
                    percent=100;
                }
                if (percent <= 0) {
                    процентЗаполненностиSpinner.setValue(0);
                    percent=0;
                }
                int V = (int)((double) percent / 400 * 3.1415 *
                        (int)ширинаSpinner.getValue() *
                        (int)ширинаSpinner.getValue() *
                        (int) высотаSpinner.getValue());
                percentPanel.setText(String.format("Объём: %d", V));
                canvasPanel1.setPercent(percent);
            }
        });
        высотаSpinner.addChangeListener(new ChangeListener() {
            @Override
            public void stateChanged(ChangeEvent e) {
                int hight = (int)высотаSpinner.getValue();
                if (hight <= 1){
                    высотаSpinner.setValue(1);
                    hight = 1;
                }
                if (hight >= 45){
                    высотаSpinner.setValue(45);
                    hight = 45;
                }
                canvasPanel1.setHight(hight);
                int V = (int)((double) (int)процентЗаполненностиSpinner.getValue()
                        / 400 * 3.1415 *
                        (int)ширинаSpinner.getValue() *
                        (int)ширинаSpinner.getValue() *
                        (int) высотаSpinner.getValue());
                percentPanel.setText(String.format("Объём: %d", V));
            }
        });
        ширинаSpinner.addChangeListener(new ChangeListener() {
            @Override
            public void stateChanged(ChangeEvent e) {
                int width = (int)ширинаSpinner.getValue();
                if (width <= 1){
                    ширинаSpinner.setValue(1);
                    width = 1;
                }
                if (width >= 38){
                    ширинаSpinner.setValue(38);
                    width = 38;
                }
                int V = (int)((double) (int)процентЗаполненностиSpinner.getValue()
                        / 400 * 3.1415 *
                        (int)ширинаSpinner.getValue() *
                        (int)ширинаSpinner.getValue() *
                        (int) высотаSpinner.getValue());
                percentPanel.setText(String.format("Объём: %d", V));
                canvasPanel1.setDiameter(width);
            }
        });

        ширинаSpinner.setValue(7);
        canvasPanel1.setDiameter(7);
        высотаSpinner.setValue(20);
        процентЗаполненностиSpinner.setValue(95);

    }

    public static void main(String[] args) {
        JFrame frame = new JFrame("Bottle");
        frame.setContentPane(new PictureForm().mainPanel);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.pack();
        frame.setVisible(true);
    }

    private void createUIComponents() {
        canvasPanel1 = new CanvasPanel();
    }
}
