#include "painlessMesh.h"

#define   MESH_PREFIX     "Monitor"
#define   MESH_PASSWORD   "12345678"
#define   MESH_PORT       5555

Scheduler userScheduler;
painlessMesh  mesh;

void prmsg();

Task taskPrintMsg(TASK_IMMEDIATE , TASK_FOREVER, &prmsg);


String out = "";
void prmsg(){
  if (out != ""){
    if (out.length() > 50){
      Serial.print(out.substring(0, 50));
      out = out.substring(50);
    }
    else{
      Serial.print(out.substring(0, out.length()));
      out = "";
    }
  }
}


void receivedCallback( uint32_t from, String &msg ) {
  out += msg;
}

void newConnectionCallback(uint32_t nodeId) {
}

void changedConnectionCallback() {
  
}

void nodeTimeAdjustedCallback(int32_t offset) {
}

void setup() {
  Serial.begin(115200);

// mesh.setDebugMsgTypes( ERROR | MESH_STATUS | CONNECTION | SYNC | COMMUNICATION | GENERAL | MSG_TYPES | REMOTE ); // all types on
   mesh.setDebugMsgTypes( ERROR );  // set before init() so that you can see startup messages

  mesh.init( MESH_PREFIX, MESH_PASSWORD, &userScheduler, MESH_PORT );
  mesh.onReceive(&receivedCallback);
  mesh.onNewConnection(&newConnectionCallback);
  mesh.onChangedConnections(&changedConnectionCallback);
  mesh.onNodeTimeAdjusted(&nodeTimeAdjustedCallback);

  mesh.setContainsRoot(true);

  userScheduler.addTask( taskPrintMsg );
  taskPrintMsg.enable();
}

void loop() {
  mesh.update();
}