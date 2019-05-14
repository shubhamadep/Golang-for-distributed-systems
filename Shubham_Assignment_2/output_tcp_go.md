								TCP for server and client, implemented in Golang




----------------------Output for server------------------------

Shubhams-MacBook-Pro:Shubham_Assignment_2 shubhamadep$ go run tcp_server_dictionary.go 
Launching server... getting the dictionary ready...
Dictionary is ready.
Word Received:dog
Word Received:happy
Word Received:dogmatic
Word Received:asdkwie
Word Received:submission
Word Received:n
Client closed the connection.


----------------------Output for client------------------------

Shubhams-MacBook-Pro:Shubham_Assignment_2 shubhamadep$ go run tcp_client_dictionary.go 
Enter a word: dog
Meaning of the word dog
 is : in the manger  n. Person who stops others using a thing for which he or she has no use.
Do you want to continue ? y/n
y
Enter a word: happy
Meaning of the word happy
 is : medium  n. Compromise; avoidance of extremes.
Do you want to continue ? y/n
y
Enter a word: dogmatic
Meaning of the word dogmatic
 is :  adj. Asserting or imposing personal opinions; intolerantly authoritative; arrogant.  dogmatically adv.
Do you want to continue ? y/n
y
Enter a word: asdkwie
Please check your spelling.
Enter a word: submission
Meaning of the word submission
 is :  n. 1 a submitting or being submitted. B thing submitted. 2 submissiveness. [latin submissio: related to *submit]
Do you want to continue ? y/n
n
