import message_pb2
import socket
import struct
from matplotlib import pyplot as plt
import time

HOST = "localhost"  # Standard loopback interface address (localhost)
PORT = 20112  # Port to listen on (non-privileged ports are > 1023)

# A Map : <attribute name : List<value>>
dict = {}
plotted = False

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.bind((HOST, PORT))
s.listen()
while True:
    conn, addr = s.accept()
    print("Listening")
    message = conn.recv(1024)
    print(message)
    message_body = message_pb2.AggregateData()
    message_body.ParseFromString(message)

    name = message_body.AGGREGATE_NAME
    value = message_body.AGGREGATE_VALUE

    if name in dict.keys():
        list = dict.get(name)
        list.append(message_body.AGGREGATE_VALUE)
    else:
        dict[name] = []


    for name in dict.keys():
        if(not plotted and len(dict[name]) >= 50):
            plt.plot(dict[name])
            plt.xlabel(name)
            plt.ylabel('value')
            plt.show()
            plotted = True


