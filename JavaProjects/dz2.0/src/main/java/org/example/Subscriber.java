package org.example;
import org.eclipse.paho.client.mqttv3.*;
import org.eclipse.paho.client.mqttv3.persist.MemoryPersistence;


public class Subscriber {
    private static final String broker = "tcp://broker.emqx.io:1883";
    private static final String topic = "iu9/Penkin";
    public static void main(String[] args) {
        try {
            MqttClient client = new MqttClient(broker, MqttClient.generateClientId());
            client.connect();

            client.subscribe(topic, (topic, message) -> {
                String parameters = new String(message.getPayload());
                int x = Integer.parseInt(parameters.substring(0,1));
                int n = Integer.parseInt(parameters.substring(2, 3));
                int res = 1;
                for(int i = 0; i < n - 1; i++){
                    res *= x;
                }
                res *= n;
                System.out.println("Производная (" + x + "^" + n + ")' = " + res);
            });
        } catch (MqttException e) {
            e.printStackTrace();
        }
    }
}
