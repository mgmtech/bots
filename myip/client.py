# TODO: Create a gopy configuration file to expose the registry

# from gobots import bots
# bots['parrot'].frontend

frontend = "ipc://parrotbackend.ipc"

import zmq
context = zmq.Context()
socket = context.socket(zmq.SUB)
identity = "Client-GoBurt-Test"
socket.setsockopt(zmq.IDENTITY, identity) #Set client identity. Makes tracing easier
socket.setsockopt(zmq.SUBSCRIBE, '') #Set client identity. Makes tracing easier
socket.connect(frontend)

print "Started parrot client"
while True:
    print "{0}".format(socket.recv())
