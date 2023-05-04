"""
This is thing to test sockets. It would be easier to test requests to json-rpc server with golang.
But interest here - to make cross-language tool that allows you to pull the handles

Methods:
1. Goods.Add
2. Warehouses.Add 
"""

import socket
import json

HOST = "127.0.0.1"
# change in case of debug(5050)/release(8000)
PORT = 8000

def jsonrpc_call(method, params):
    payload = {
        "method": method,
        "params": [params],
        "id": 1
    }

    request = bytes(json.dumps(payload), 'utf-8')

    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.connect((HOST, PORT))
        s.sendall(request)

        response = s.recv(1024)
        print("recieved:", response)

def AddGoods():
    """AddGoods reads input text, split it by spaces and call Goods.Add with the list of splitted words"""
    goods = input("Enter goods divided by space: ").split()

    jsonrpc_call("Goods.Add", goods)

def AddWarehouses():
    """AddWarehouses reads input text, split it by spaces and call Warehouses.Add with the list of splitted words"""
    wh = input("Enter warehouses divided by space: ").split()

    jsonrpc_call("Warehouses.Add", wh)

if __name__ == "__main__":
    print(__doc__)
    while True:
        try:
            choice = int(input())
        except Exception:
            print("Should be integer")
        else:
            break
    
    match choice:
        case 1:
            AddGoods()
        case 2:
            AddWarehouses()
