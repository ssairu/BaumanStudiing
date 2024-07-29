import socket
import threading

def handle_client(client_socket):
    while True:
        data = client_socket.recv(1024).decode()
        if not data:
            break
        print(data + "    ._.\n")
    client_socket.close()

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind(('', 8081))
server.listen()

while True:
    client, addr = server.accept()
    client_thread = threading.Thread(target=handle_client, args=(client,))
    client_thread.start()
