#include "painlessMesh.h"

#define   MESH_PREFIX     "Monitor"
#define   MESH_PASSWORD   "12345678"
#define   MESH_PORT       5555

Scheduler userScheduler;
painlessMesh  mesh;


void sendMessage();

Task taskSendMessage( TASK_SECOND * 1 , TASK_FOREVER, &sendMessage );

void sendMessage() {
  // String msg = "5:2.42:-47.4:";
  // mesh.sendBroadcast( msg );
  // Serial.println(msg);
  String msg2 = Serial.readString();
  Serial.println(msg2);
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
  Serial.setTimeout(15000);

//mesh.setDebugMsgTypes( ERROR | MESH_STATUS | CONNECTION | SYNC | COMMUNICATION | GENERAL | MSG_TYPES | REMOTE ); // all types on
 mesh.setDebugMsgTypes( ERROR | STARTUP | CONNECTION );  // set before init() so that you can see startup messages

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
