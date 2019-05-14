import socket

msgFromClient       = "dog\n"
bytesToSend         = str.encode(msgFromClient)
addr   = ("127.0.0.1", 8081)
bufferSize          = 1024
goahead = "y" 
word = ""
while goahead == "y" :

    word = input("Enter a word : ")
    bytesToSend = str.encode(word+"\n")
    UDPClientSocket = socket.socket(family=socket.AF_INET, type=socket.SOCK_DGRAM)
    UDPClientSocket.sendto(bytesToSend, addr)
    msgFromServer = UDPClientSocket.recvfrom(bufferSize)
    message = msgFromServer[0].decode("utf-8")
    if message != "incorrect" :
        print("Meaning of the word ", word, " is : ", message) #msgFromServer[0]
        print("Do you want to continue ? y/n ?")
        goahead = input()
    else :
        print("Please check the spelling.")

UDPClientSocket.sendto(str.encode("n"), addr)