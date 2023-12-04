package org.example;
import org.eclipse.paho.client.mqttv3.*;
import java.util.Scanner;

public class Publisher {
    private static final String broker = "tcp://broker.emqx.io:1883";
    private static final String topic = "iu9/Penkin";
    public static void main(String[] args) {
        try {
            MqttClient client = new MqttClient(broker, MqttClient.generateClientId());
            client.connect();
            Scanner scanner = new Scanner(System.in);
            for (int i = 0; i < 3; i++) {
                System.out.println("Введите число х и степень n:");
                int x = scanner.nextInt();
                int n = scanner.nextInt();
                String res = "" + x + " " + n;
                MqttMessage message = new MqttMessage(res.getBytes());
                client.publish(topic, message);
                System.out.println("Данные успешно отправлены в топик " + topic);
            }
            client.disconnect();
        } catch (MqttException e) {
            e.printStackTrace();
        }
    }
}
