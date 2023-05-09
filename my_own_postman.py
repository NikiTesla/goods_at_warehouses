"""
This is thing to test sockets. It would be easier to test requests to json-rpc server with golang.
But interest here - to make cross-language tool that allows you to pull the handles

Methods:
1. Goods.Add
2. Goods.Create
3. Warehouses.Create 
4. Goods.Reserve
5. Goods.CancelReservation
6. Warehouses.GetAmount
7. repeat __doc__
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


def CreateGood():
    """CreateGood reads input name, size, code, amount. Call procedurewith proper format of Good{}"""
    name = input("Enter name: ")
    size = float(input("Enter size: "))
    code = int(input("Enter code: "))
    amount = int(input("Enter amount: "))

    jsonrpc_call("Goods.Create", [{
        "name": name,
        "size": size,
        "code": code,
        "amount": amount
    }])

def AddGood():
    """AddGood reads good_code, warehouse_id and amount of goods to add. Adding this into warehosue_goods"""
    code = int(input("Enter good uniqe code: "))
    warehouse_id = int(input("Enter warehouse id: "))
    amount = int(input("Enter amount of goods: "))

    jsonrpc_call("Goods.Add", [{
        "good_code": code,
        "warehouse_id": warehouse_id,
        "amount": amount
    }])

def CreateWarehouse():
    """CreateWarehouse reads input name, size, code, amount. Call remote procedure with proper format of Warehouse{}"""

    name = input("Enter name: ")
    availability = True if input("Enter true/false for availability: ") == "true" else False

    jsonrpc_call("Warehouses.Create", [{
        "name": name,
        "availability": availability
    }])

def ReserveGood():
    """ReserveGood reads input numbers, split it by spaces and call Goods.Reserve with this list"""
    goods = []
    good = input("Enter good for reservation (good_code warehouse_id amount): ").split()
    while good:
        goods.append({"good_code": int(good[0]), "warehouse_id": int(good[1]), "amount": int(good[2])})
        good = input("Enter good for reservation (good_code warehouse_id amount): ").split()

    jsonrpc_call("Goods.Reserve", goods)

def CancelGoodReservation():
    """CancelGoodReservation reads input numbers, split it by spaces and call Goods. CancelReseration with this list"""
    goods = []
    good = input("Enter good to cancel reservation (good_code warehouse_id amount): ").split()
    while good:
        goods.append({"good_code": int(good[0]), "warehouse_id": int(good[1]), "amount": int(good[2])})
        good = input("Enter good for cancel reservation (good_code warehouse_id amount): ")

    jsonrpc_call("Goods.CancelReservation", goods)

def GetAmount():
    """Get amount reads good code and warehouse id. Call procedure Warehouses.GetAmount to get amount of good on warehouse"""
    good_code = int(input("Enter good code: "))
    warehouse_id = int(input("Enter warehouse id: "))

    jsonrpc_call("Warehouses.GetAmount", {
        "good_code": good_code,
        "warehouse_id": warehouse_id
    })

def match_choice():
    match choice:
        case 1:
            AddGood()
        case 2:
            CreateGood()
        case 3:
            CreateWarehouse()
        case 4:
            ReserveGood()
        case 5:
            CancelGoodReservation()
        case 6:
            GetAmount()
        case 7:
            print(__doc__)

if __name__ == "__main__":
    print(__doc__)
    while True:
        try:
            choice = int(input("Enter choice: "))
        except Exception:
            print("Should be integer")
        else:
            match_choice()
            print()