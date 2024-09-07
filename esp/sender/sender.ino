#include "painlessMesh.h"

#define   MESH_PREFIX     "Monitor"
#define   MESH_PASSWORD   "12345678"
#define   MESH_PORT       5555

Scheduler userScheduler;
painlessMesh  mesh;


void sendMessage();

Task taskSendMessage( TASK_SECOND * 10 , TASK_FOREVER, &sendMessage );


void sendMessage() {
  // String msg = "qwertyuiop";
  // mesh.sendBroadcast( msg );
  // Serial.println(msg);
  // String msg2 = Serial.readString();
  // Serial.println(msg2);
  // mesh.sendBroadcast( msg2 );
      while(Serial.available() > 0) {
        String trash = Serial.readStringUntil('&');
        if ((char)Serial.read() != '$' || (char)Serial.read() != '&'){
          continue;
        }
        
        String msg2 = "&$&" + Serial.readStringUntil('%');
        char t1 = (char)Serial.read();
        char t2 = (char)Serial.read();
        while (t1 != '@' || t2 != '%'){
          msg2 = msg2 + "%" + t1 + t2 + Serial.readStringUntil('%');
        }
        msg2 = msg2 + "%@%";
        Serial.print(msg2); 
        mesh.sendBroadcast( msg2 );
      }
    }


void receivedCallback( uint32_t from, String &msg ) {
  Serial.printf("startHere: Received from %u msg=%s\n", from, msg.c_str());
}

void newConnectionCallback(uint32_t nodeId) {
    Serial.printf("--> startHere: New Connection, nodeId = %u\n", nodeId);
}

void changedConnectionCallback() {
  Serial.printf("Changed connections\n");
}

void nodeTimeAdjustedCallback(int32_t offset) {
    Serial.printf("Adjusted time %u. Offset = %d\n", mesh.getNodeTime(),offset);
}

void setup() {
  Serial.begin(115200);
  Serial.setTimeout(1000);

  mesh.setDebugMsgTypes( ERROR | MESH_STATUS | CONNECTION | SYNC | COMMUNICATION | GENERAL | MSG_TYPES | REMOTE ); // all types on
//  mesh.setDebugMsgTypes( STARTUP | CONNECTION );  // set before init() so that you can see startup messages

  mesh.init( MESH_PREFIX, MESH_PASSWORD, &userScheduler, MESH_PORT );
  mesh.onReceive(&receivedCallback);
  mesh.onNewConnection(&newConnectionCallback);
  mesh.onChangedConnections(&changedConnectionCallback);
  mesh.onNodeTimeAdjusted(&nodeTimeAdjustedCallback);

  mesh.setContainsRoot(true);

  userScheduler.addTask( taskSendMessage );
  taskSendMessage.enable();
}

void loop() {
  mesh.update();
}

