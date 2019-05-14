Version 5 - Implemented Programming Synchronization

Server provides read and write features. It assigns semaphore for write operation (critical region). When the in-memory data is getting
written, the read operations are stopped. Once the semaphore is released, then read operations are allowed.

There are two clients who are sending read and write requests simultaneously. The output for both the clients can be found in:
1. send_write_request_client_output.txt
2. send_read_client_request_output.txt

The output for the server can be found in : server_output.txt


To run the experiments:

execute these statements:
go run tcp_server_with_semaphore.go
go run tcp_client_with_semaphore.go
go run tcp_client_add_requests.go
