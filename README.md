# CS587 Homework-2

## University of Illinois at Chicago

## Instructions to run the project.

### To start the server:

```
ssh to the VM
ssh me@192.168.56.101
```

- Create a folder where the following files can be copied:
  bankingClient.go, bankingServer.go, myRpc.t and Makefile

- cd into that folder then run the following commands:

```
make clean
make install
cd client folder
sudo -E ethosRun
```

This will start the banking server instance on ethos.

### To start the client and make calls:

- In another terminal(s) ssh and cd into the same client folder with different users.
- Then run the following command to run the ethos VM terminal:

`etAl client.ethos`

In the new terminal, run:

`bankingClient`

This will then execute the bankingClient code which will make calls to the bankingServer.

- For logging: `ethosLog .`

- For allowing permissions to other users to a folder: `sudo chmod g=rwx /home/me`

- For network issues: `systemctl restart network`
