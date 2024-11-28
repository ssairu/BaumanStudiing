#include "painlessMesh.h"

#define   MESH_PREFIX     "Monitor"
#define   MESH_PASSWORD   "12345678"
#define   MESH_PORT       5555

Scheduler userScheduler;
painlessMesh  mesh;

void checkStart();
void checkEnd();
void parseMessage();
void sendMessage();

Task taskSendMessage( TASK_IMMEDIATE , TASK_FOREVER, &checkStart);

String msg = "";
char s1 = 0;
char s2 = 0;
char s3 = 0;

void updateTOP(char c){
  s1 = s2;
  s2 = s3;
  s3 = c;
}

void clearBUF(){
  s1 = 0;
  s2 = 0;
  s3 = 0;
}



void checkStart(){
  if (Serial.available()){
    // Serial.print("AS ");
    char s = char(Serial.read());
    // Serial.println(s);
    updateTOP(s);

    if (s1 == '&' && s2 == '$' && s3 == '&'){
      msg = "&$&";
      clearBUF();
      Serial.printf("msg is %s\n", msg);
      taskSendMessage.setCallback(&parseMessage);
      Serial.print("go to parse \n");
    }
    // else {
    //   Serial.print("check start further \n");
    // }
  }
  // else {
  //   taskSendMessage.delay(50);
  //   Serial.print("-");
  // }
}

void parseMessage(){
  char character;
  char i = 0;
  while (Serial.available() && i < 50){
    // Serial.print("AP ");
    character = char(Serial.read());
    // Serial.println(character);
    updateTOP(character);

    msg = msg + character;
    i++;
    // Serial.printf("msg = %s \n", msg);

    if (s1 == '%' && s2 == '@' && s3 == '%'){
      taskSendMessage.setCallback(&sendMessage);
      Serial.print("go to send \n");
      break;
    }
  }  
}

void sendMessage(){
  mesh.sendBroadcast( msg );
  Serial.printf("SENDING %d \n", msg.length());
  msg = "";
  taskSendMessage.setCallback(&checkStart);
}




void receivedCallback( uint32_t from, String &msg ) {
  // Serial.printf("startHere: Received from %u msg=%s\n", from, msg.c_str());
}

void newConnectionCallback(uint32_t nodeId) {
    // Serial.printf("--> startHere: New Connection, nodeId = %u\n", nodeId);
}

void changedConnectionCallback() {
  // Serial.printf("Changed connections\n");
}

void nodeTimeAdjustedCallback(int32_t offset) {
    // Serial.printf("Adjusted time %u. Offset = %d\n", mesh.getNodeTime(),offset);
}

void setup() {
  Serial.setRxBufferSize(16384);
  Serial.begin(115200);
  Serial.setTimeout(100000);

  // mesh.setDebugMsgTypes( ERROR | MESH_STATUS | CONNECTION | SYNC | COMMUNICATION | GENERAL | MSG_TYPES | REMOTE ); // all types on
 mesh.setDebugMsgTypes( ERROR );  // set before init() so that you can see startup messages

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

// void sendMessage() {
//   // String msg = "qwertyuiop";
//   // mesh.sendBroadcast( msg );
//   // Serial.println(msg);
//   // String msg2 = Serial.readString();
//   // Serial.println(msg2);
//   // mesh.sendBroadcast( msg2 );
//   String counter = "";
//   while (Serial.available()){
//         char character = 0;
//         while (character != '&'){
//           character = char(Serial.read());
//         }
//         char s1 = char(Serial.read());
//         char s2 = char(Serial.read());
//         if (s1 != '$' || s2 != '&'){
//           continue;
//         }
//         counter = counter + "+";

//         String content = "";
        

//         for (int i = 0; i < 20; i++){
//           character = char(Serial.read());
//           Serial.print(character);
//           content = content + character;
//         }
        
//         while (character != '%'){
//           character = char(Serial.read());
//           Serial.print(character);
//           content = content + character;
//         }
        

//         String msg2 = "&$&" + content;
//         char t1 = char(Serial.read());
//         char t2 = char(Serial.read());
//         while (t1 != '@' || t2 != '%'){

//           content = "";
//           character = t2;
//           while (character != '%'){
//             character = char(Serial.read());
//             Serial.print(character);
//             content = content + character;
//           }

//           msg2 = msg2 + t1 + t2 + content;
//           t1 = (char)Serial.read();
//           t2 = (char)Serial.read();
//         }
//         msg2 = msg2 + "@%"; // + counter;
        
//         Serial.print(msg2); 
//         mesh.sendBroadcast( msg2 );
//       }
// }

