# Chapter 4:Elementary TCP sockets

Questions:

1. What does the paragraph mean? We must be careful to distinguish between the interface on which a packet arrives versus the destination IP address of that packet. In Section 8.8, we will talk about the weak end system model and the strong end system model. Most implementations employ the former, meaning it is okay for a packet to arrive with a destination IP address that identifies an interface other than the interface on which the packet arrives. (This assumes a multihomed host.) Binding a non-wildcard IP address restricts the datagrams that will be delivered to the socket based only on the destination IP address. It says nothing about the arriving interface, unless the host employs the strong end system model.

## 4.1 introduction
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/4/sock_funcs.gif "sock_funcs")

## 4.2 socket func
basic funcs:

	#include <sys/socket.h>
	int socket(int family, int type, int protocol);//non-negative if OK,-1 error

    family : AF_INET, AF_INET6, AF_LOCAL(unix domain sockets,AF_UNIX), AF_ROUTE(Routing sockets), AF_KEY(key socket)
    type : SOCK_STREAM, SOCK_DGRAM, SOCK_SEQPACKET(sequenced packed socket), SOCK_RAW
    protocol : IPPROTO_TCP, IPPROTO_UDP, IPPROTO_SCTP

Combinations of family and type for the socket function.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/4/combination_socket.gif "combination_socket")

---
AF_xx VS PF_xx

The "AF_" prefix stands for "address family" and the "PF_" prefix stands for "protocol family." Historically, the intent was that a single protocol family might support multiple address families and that the PF_ value was used to create the socket and the AF_ value was used in socket address structures. But in actuality, a protocol family supporting multiple address families has never been supported and the <sys/socket.h> header defines the PF_ value for a given protocol to be equal to the AF_ value for that protocol. While there is no guarantee that this equality between the two will always be true, should anyone change this for existing protocols, lots of existing code would break. To conform to existing coding practice, we use only the AF_ constants in this text, although you may encounter the PF_ value, mainly in calls to socket.


## 4.3 connect func

    #include <sys/socket.h> 
    int connect(int sockfd, const struct sockaddr *servaddr, socklen_t addrlen);//0 if OK,-1 on error
 
In the case of a TCP socket, the connect function initiates TCP's three-way handshake (Section 2.6). The function returns only when the connection is established or an error occurs. Errors:

1. If the client TCP receives no response to its SYN segment, ETIMEDOUT is returned.
2. If the server's response to the client's SYN is a reset (RST), this indicates that no process is waiting for connections on the server host at the port specified (i.e., the server process is probably not running). This is a **hard error** and the error ECONNREFUSED is returned to the client as soon as the RST is received.
3. If the client's SYN elicits(引出) an ICMP "destination unreachable" from some intermediate router, this is considered a **soft error**. The client kernel saves the message but keeps sending SYNs with the same time between each SYN as in the first scenario. If no response is received after some fixed amount of time (75 seconds for 4.4BSD), the saved ICMP error is returned to the process as either **EHOSTUNREACH** or **ENETUNREACH**. (applications should just treat ENETUNREACH and EHOSTUNREACH as the same error.)

## 4.4 bind

    #include <sys/socket.h>
    int bind (int sockfd, const struct sockaddr *myaddr, socklen_t addrlen);//0 if OK,-1 on error

1. Servers bind their well-known port when they start. Exceptions to this rule are Remote Procedure Call (RPC) servers.
2. A process can bind a specific IP address to its socket. The IP address must belong to an interface on the host. If a TCP server does not bind an IP address to its socket, the kernel uses the destination IP address of the client's SYN as the server's source IP address.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/4/bind_results.gif "bind_results")

With IPv4, the wildcard address is specified by the constant INADDR_ANY.IPv6, we write

		struct sockaddr_in   servaddr;
        servaddr.sin_addr.s_addr = htonl (INADDR_ANY);     /* wildcard */
 
		struct sockaddr_in6    serv;
        serv.sin6_addr = in6addr_any;     /* wildcard */
The system allocates and initializes the `in6addr_any` variable to the constant `IN6ADDR_ANY_INIT`. The <netinet/in.h> header contains the extern declaration for in6addr_any.

The value of `INADDR_ANY` (0) is the same in either network or host byte order, so the use of htonl is not really required. But, since all the `INADDR_constants` defined by the <netinet/in.h> header are defined in host byte order, we should use htonl with any of these constants.(good habit!!)

If we tell the kernel to choose an ephemeral port number for our socket, notice that bind does not return the chosen value. To obtain the value of the ephemeral port assigned by the kernel, we must call *getsockname* to return the protocol address.


## 4.5 listen

    #include <sys/socket.h>
    #int listen (int sockfd, int backlog);//Returns: 0 if OK, -1 on error 


When a socket is created by the socket function, it is assumed to be an active socket, that is, a client socket that will issue a connect. The listen function converts an unconnected socket into a passive socket, indicating that the kernel should accept incoming connection requests directed to this socket. In terms of the TCP state transition diagram (Figure 2.4), the call to listen moves the socket from the CLOSED state to the LISTEN state.

backlog : the maximum number of connections the kernel should queue for this socket.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/4/listen_queues.gif "listen_queues")

## 4.7 fork and exec

    #include <unistd.h>
    pid_t fork(void);//Returns: 0 in child, process ID of child in parent, -1 on error
 
*fork* (including the variants of it provided by some systems) is the only way in Unix to create a new process. It returns once in the calling process (called the **parent) with a return value that is the process ID of the newly created process** (the child). It also returns once in the child, **with a return value of 0**. Hence, the return value tells the process whether it is the parent or the child.

The reason fork returns 0 in the child, instead of the parent's process ID, is because a child has only one parent and it can always obtain the parent's process ID by calling *getppid*. A parent, on the other hand, can have any number of children, and** there is no way to obtain the process IDs of its children**. If a parent wants to keep track of the process IDs of all its children, it must record the return values from fork.

Two typical uses of fork:
1. A process makes a copy of itself so that one copy can handle one operation while the other copy does another task.
2. A process wants to execute another program. Since the only way to create a new process is by calling fork, the process first calls fork to make a copy of itself, and then one of the copies (typically the child process) calls exec (described next) to replace itself with the new program. This is typical for programs such as shells.

The only way in which **an executable program file on disk can be executed by Unix** is for an existing process to call one of the six **exec functions**.

*exec* replaces the current process image with the new program file, and this new program normally starts at the main function. The process ID does not change.(So, no new process is created.)

Differences in six exec funcs are:
1. whether the program file to execute is specified by a filename or a pathname
2. whether the arguments to the new program are listed one by one or referenced through an array of pointers
3. whether the environment of the calling process is passed to the new program or whether a new environment is specified


    #include <unistd.h>
     
    int execl (const char *pathname, const char *arg0, ... /* (char *) 0 */ );
    int execv (const char *pathname, char *const argv[]);
    int execle (const char *pathname, const char *arg0, .../* (char *) 0, char *const envp[] */ );
     
    int execve (const char *pathname, char *const argv[], char *const envp[]);
    int execlp (const char *filename, const char *arg0, ... /* (char *) 0 */ );
    int execvp (const char *filename, char *const argv[]);
     
    All six return: -1 on error, no return on success

Relationship among the six exec functions:

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/4/six_exec_relations.gif "six exec")
 
Look at the Figure above:
1. The three functions **in the top row** specify each argument string as a separate argument to the exec function, **with a null pointer terminating the variable number of arguments**. The three functions in the second row have an argv array, containing pointers to the argument strings. **This argv array must contain a null pointer to specify its end, since a count is not specified**.
2. The two functions in the left column specify a filename argument. This is converted into a pathname using the current PATH environment variable. If the filename argument to execlp or execvp contains a slash (/) anywhere in the string, the PATH variable is not used.
3. The four functions in the left two columns do not specify an explicit environment pointer. Instead, the current value of the **external variable environ** is used for building an environment list that is passed to the new program. The two functions in the right column specify an explicit environment list. **The envp array of pointers must be terminated by a null pointer.**

Descriptors open in the process before calling exec normally remain open across the exec. We use the qualifier "normally" because this can be disabled using fcntl to set the FD_CLOEXEC descriptor flag. The inetd server uses this feature.

## 4.8 concurrent servers

Outline for typical concurrent server:

    pid_t pid;
    int listenfd;
    int connfd;
    
    listenfd = socket(...);
    Bind(listenfd, ...);
    Listen(listenfd, LISTENQ);
    
    for(;;){
    connfd = accept(listenfd, ...);
    	if( (pid == fork()) == 0 ){//error occurs,exit
    	close(listenfd);
    	do_something(connfd);
    	close(connfd);
    	exit(0)
    	}
    close(connfd);//decrements the reference count
    }

We must understand that every file or socket has a reference count. The reference count is maintained in the file table entry. This is a count of the number of descriptors that are currently open that refer to this file or socket. The actual cleanup and de-allocation of the socket does not happen until the reference count reaches 0.

## 4.9 close()

    #include <unistd.h>
    int close (int sockfd);//Returns: 0 if OK, -1 on error

The default action of close with a TCP socket is to mark the socket as closed and return to the process immediately. The socket descriptor is no longer usable by the process: It cannot be used as an argument to read or write. But, TCP will try to send any data that is already queued to be sent to the other end.

If we really want to send a FIN on a TCP connection, the shutdown function can be used (Section 6.6) instead of close.

We must also be aware of what happens in our concurrent server if the parent does not call close for each connected socket returned by accept.
 
1. the parent will eventually run out of descriptors, as there is usually a limit to the number of descriptors that any process can have open at any time. 
2. None of the client connections will be terminated.

## 4.10 getsockname, getpeername

#include <sys/socket.h>
 
    int getsockname(int sockfd, struct sockaddr *localaddr, socklen_t *addrlen);     
    int getpeername(int sockfd, struct sockaddr *peeraddr, socklen_t *addrlen);
    Both return: 0 if OK, -1 on error

1. After *connect* successfully returns in a TCP client that does not call bind, *getsockname* returns the local IP address and local port number assigned to the connection by the kernel
2. After calling bind with a port number of 0 (telling the kernel to choose the local port number), getsockname returns the local port number that was assigned.
3. getsockname can be called to obtain the address family of a socket(sockaddr_storage)
4. In a TCP server that binds the wildcard IP address, once a connection is established with a client (accept returns successfully), **the server can call getsockname to obtain the local IP address assigned to the connection**. The socket descriptor argument in this call must be that of the connected socket, and not the listening socket