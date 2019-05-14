								TCP for server and client, server in go, client in python




----------------------Output for server------------------------

Shubhams-MacBook-Pro:Shubham_Assignment_2 shubhamadep$ go run udp_server_dictionary.go 
Launching server... getting the dictionary ready...
Dictionary is ready.
Word recieved :  dog

Word recieved :  happy

Word recieved :  dogmatic

Word recieved :  askdws

Word recieved :  arsenal

Client closed the connection.
Word recieved :  n


----------------------Output for client------------------------

Shubhams-MacBook-Pro:Shubham_Assignment_2 shubhamadep$ python python_udp_client.py 
Enter a word : dog
Meaning of the word  dog  is :  in the manger  n. Person who stops others using a thing for which he or she has no use.

Do you want to continue ? y/n ?
y
Enter a word : happy
Meaning of the word  happy  is :  medium  n. Compromise; avoidance of extremes.

Do you want to continue ? y/n ?
y
Enter a word : dogmatic
Meaning of the word  dogmatic  is :   adj. Asserting or imposing personal opinions; intolerantly authoritative; arrogant.  dogmatically adv.

Do you want to continue ? y/n ?
y
Enter a word : askdws
Please check the spelling.
Enter a word : arsenal
Meaning of the word  arsenal  is :   n. 1 store, esp. Of weapons. 2 place for the storage and manufacture of weapons and ammunition. [arabic, = workshop]

Do you want to continue ? y/n ?
n

