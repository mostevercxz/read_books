Questions:

1. It is remarkable how many network break-ins have occurred by a hacker sending data to cause a server's call to sprintf to overflow its buffer. How to do this?
2. Additional tips on writing secure network programs are found in Chapter 23 of [Garfinkel, Schwartz, and Spafford 2003]. How to write secure network programs?
3. If we want to be certain that a datagram reaches its destination, we can build lots of features into our application: acknowledgments from the other end, timeouts, retransmissions, and the like. How to make sure a datagram reaches its destination?
4. After a full-duplex connection is established, it can be turned into a simplex connection if desired (see Section 6.6). How to do it?
5. TIME_WAIT State, To allow old duplicate segments to expire in the network, how to understand this?

Acknowledges:

1. International Organization for Standardization (ISO)
2. open systems interconnection (OSI)
3. Ethernet maximum transfer unit (MTU)
4. Application Programming Interfaces (APIs)
5. Computer Systems Research Group (CSRG)
6. Internet Engineering Task Force (IETF)

# Chapter 1, introduction
A protocol, an agreement on how those programs will communicate. Before delving into the design details of a protocol, high-level decisions must be made about which program is expected to initiate communication and when responses are expected.

**Client and server on the same Ethernet communicating using TCP**
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/1/01fig03.gif "cs communicate layer")

Routers are the building blocks of WANs. The largest WAN today is the Internet.

## 1.2 a simple daytime application

    /usr/include/bits/sockaddr.h
    typedef unsigned short int sa_family_t;
    #define __SOCKADDR_COMMON(sa_prefix) \
    sa_family_t sa_prefix##family
	#define __SOCKADDR_COMMON_SIZE (sizeof(unsigned short int))

	/usr/include/bits/socket.h
	struct sockaddr{
		__SOCKADDR_COMMON (sa_);// Common data : address family and length
		char sa_data[14];//address data
	};
	#define PF_INET 2//IP Protocol family
	#define PF_INET6 10//IP version 6
	#define AF_INET PF_INET
	#define AF_INET6 PF_INET6

    /usr/include/netinet/in.h
	typedef uint16_t in_port_t;
	typedef uint32_t in_addr_t;
	struct in_addr{
		in_addr_t s_addr;
	};
    struct sockaddr_in{
    __SOCKADDR_COMMON (sin_);
    in_port_t sin_port;
    struct in_addr sin_addr;    
    // Pad to size of 'struct sockaddr'
    unsigned char sin_zero[sizeof(struct sockaddr) - __SOCKADDR_COMMON_SIZE -
    	sizeof(in_port_t) -
    	sizeof(struct in_addr)];//16 - 2 - 2 - 4 = 8
    };

	/usr/include/bits/types.h
	#define __U32_TYPE unsigned int
	#if __WORDSIZE == 32
	#define __STD_TYPE __extension__ typedef
	#elif __WORDSIZE == 64
	#define __STD_TYPE typedef
	#endif
	__STD_TYPE __U32_TYPE __socklen_t;

	/usr/include/arpa/inet.h
	/usr/include/bits/socket.h
	typedef __socklen_t socklen_t;
	

We will encounter many different uses of the term "socket." First, the API that we are using is called the sockets API. In the preceding paragraph, we referred to a function named socket that is part of the sockets API. In the preceding paragraph, we also referred to a TCP socket, which is synonymous with a TCP endpoint.

inet_pton, presentation to numeric
where \r is the ASCII carriage return and \n is the ASCII linefeed

**Important-do not simply use read**:

With a byte-stream protocol, these 26 bytes can be returned in numerous ways: a single TCP segment containing all 26 bytes of data, in 26 TCP segments each containing 1 byte of data, or any other combination that totals to 26 bytes. Normally, a single segment containing all 26 bytes of data is returned, but with larger data sizes, **we cannot assume that the server's reply will be returned by a single read**. Therefore, when reading from a TCP socket, we always need to code the read in a loop and terminate the loop **when either read returns 0 (i.e., the other end closed the connection) or a value less than 0 (an error)**.

In this example, the end of the record is being denoted by the server closing the connection. This technique is also used by version 1.0 of the Hypertext Transfer Protocol (HTTP). Other techniques are available. 

* For example, the Simple Mail Transfer Protocol (SMTP) marks the end of a record with the two-byte sequence of an ASCII carriage return followed by an ASCII linefeed. 
* Sun Remote Procedure Call (RPC) and the Domain Name System (DNS) place a binary count containing the record length in front of each record that is sent when using TCP. 

**The important concept here is that TCP itself provides no record markers**: If an application wants to delineate the ends of records, it must do so itself and there are a few common ways to accomplish this.

## 1.3 Protocol independence
It is better to make a program protocol-independent. Figure 11.11 will show a version of this client that is protocol-independent by using the getaddrinfo function (which is called by tcp_connect).

In Chapter 11, we will discuss the functions that convert between hostnames and IP addresses, and between service names and ports. We purposely put off the discussion of these functions and continue using IP addresses and port numbers so we know exactly what goes into the socket address structures that we must fill in and examine.

## 1.4 Error handling, wrapper functions
In any real-world program, it is essential to check every function call for an error return.

Since terminating on an error is the common case, we can shorten our programs by defining a wrapper function that performs the actual function call, tests the return value, and terminates on an error. The convention we use is to capitalize the name of the function, as in

    int
    Socket(int family, int type, int protocol)
    {
    	int n;
    
    	if ( (n = socket(family, type, protocol)) < 0)
    	err_sys("socket error");
    	return (n);
    }

We will find that thread functions do not set the standard Unix errno variable when an error occurs; instead, the errno value is the return value of the function. This means that every time we call one of the pthread_ functions, we must allocate a variable, save the return value in that variable, and then set errno to this value before calling err_sys.

    lib/wrappthread.c
    
	errno = n, err_sys("pthread_mutex_lock error");
    void
    Pthread_mutex_lock(pthread_mutex_t *mptr)
    {
    	int n;
	    
    	if ( (n = pthread_mutex_lock(mptr)) == 0)
	    return;
    	errno = n;
	    err_sys("pthread_mutex_lock error");
    }

With careful C coding, we could use macros instead of functions, providing a little run-time efficiency, but these wrapper functions are **rarely** the performance bottleneck of a program. This technique has the side benefit of checking for errors from functions **whose error returns are often ignored**: close and listen, for example.

---
**Unix errno variable**

When an error occurs in a Unix function (such as one of the socket functions), the global variable errno is set to a positive value indicating the type of error and the function normally returns –1.

The value of errno is set by a function only if an error occurs. Its value is undefined if the function does not return an error. All of the positive error values are constants with all-uppercase names beginning with "E," and are normally defined in the <sys/errno.h> header. No error has a value of 0.

Storing errno in a global variable does not work with multiple threads that share all global variables. We will talk about solutions to this problem in Chapter 26.

/usr/include/sys/errno.h
/usr/include/asm-generic/errno.h

## 1.5 A simple daytime server
**Figure 1.9 TCP daytime server.**

    #include <time.h>
	struct sockaddr_in serveraddr;
	bzero(&serveraddr, sizeof(serveraddr));
	servaddr.sin_family = AF_INET;
	servaddr.sin_addr.s_addr = htonl(INADDR_ANY);
	servaddr.sin_port = htons(13); /* daytime server */

	clientfd = Accept(listenfd, (sa*)NULL, NULL);
	time_t ticks = time(NULL);
    snprintf(buff, sizeof(buff), "%.24s\r\n", ctime(&ticks));
	Write(clientfd, buff, strlen(buff));

A TCP connection uses what is called a three-way handshake to establish a connection. When this handshake completes, accept returns, and the return value from the function is a new descriptor (connfd) that is called the connected descriptor. This new descriptor is used for communication with the new client. A new descriptor is returned by accept for each client that connects to our server.

It is remarkable how many network break-ins have occurred by a hacker sending data to cause a server's call to sprintf to overflow its buffer. Other functions that we should be careful with are gets, strcat, and strcpy, normally calling fgets, strncat, and strncpy instead. Even better are the more recently available functions strlcat and strlcpy, which ensure the result is a properly terminated string.

close() initiates the normal TCP connection termination sequence: a FIN is sent in each direction and each FIN is acknowledged by the other end. We will say much more about TCP's three-way handshake and the four TCP packets used to terminate a TCP connection in Section 2.6.

Some problems:

1. not protocol-independent 
2. Our server handles only one client at a time. If multiple client connections arrive at about the same time, the kernel queues them, up to some limit, and returns them to accept one at a time.
3. The server that we show in Figure 1.9 is called an iterative server because it iterates through each client, one at a time. There are numerous techniques for writing a concurrent server, one that handles multiple clients at the same time. The simplest technique for a concurrent server is to call the Unix fork function (Section 4.7), creating one child process for each client. Other techniques are to use threads instead of fork (Section 26.4), or to pre-fork a fixed number of children when the server starts (Section 30.6).
4. If we start a server like this from a shell command line, we might want the server to run for a long time, since servers often run for as long as the system is up. This requires that we add code to the server to run correctly as a Unix daemon: a process that can run in the background, unattached to a terminal

## 1.6 road map

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/1/01fig10.gif "road map")
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/1/01fig11.gif "road map")
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/1/01fig12.gif "road map")
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/1/01fig13.gif "road map")

## 1.7 OSI model
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/1/01fig14.gif "OSI model layers")

We consider the bottom two layers of the OSI model as the device driver and networking hardware that are supplied with the system. Normally, we need not concern ourselves with these layers other than being aware of some properties of the datalink, such as the 1500-byte Ethernet maximum transfer unit (MTU).

The network layer is handled by the IPv4 and IPv6 protocols, both of which we will describe in Appendix A. The transport layers that we can choose from are TCP and UDP, and we will describe these in Chapter 2. **We show a gap between TCP and UDP in Figure 1.14 to indicate that it is possible for an application to bypass the transport layer and use IPv4 or IPv6 directly**. This is called a raw socket, and we will talk about this in Chapter 28.

The sockets programming interfaces described in this book are interfaces from the upper three layers (the "application") into the transport layer. This is the focus of this book: **how to write applications using sockets that use either TCP or UDP**. We already mentioned raw sockets, and in Chapter 29 we will see that we can even bypass the IP layer completely to read and write our own datalink-layer frames.

Why do sockets provide the interface from the upper three layers of the OSI model into the transport layer?
 
1. (knows little about each other)the upper three layers handle all the details of the application (FTP, Telnet, or HTTP, for example) and know little about the communication details. The lower four layers know little about the application, but handle all the communication details: sending data, waiting for acknowledgments, sequencing data that arrives out of order, calculating and verifying checksums, and so on
2. (User process and OS kernel)the upper three layers often form what is called a user process while the lower four layers are normally provided as part of the operating system (OS) kernel. Unix provides this separation between the user process and the kernel, as do many other contemporary operating systems.

Linux's networking code and sockets API were developed from scratch. More information on the various BSD releases, and on the history of the various Unix systems in general, can be found in Chapter 01 of [McKusick et al. 1996].(uploaded to pan.baidu.com)

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/1/mc_reference.png "mc reference")

## 1.9 Test networks
    netstat -ni
    netstat -nr
    ifconfig eth0
    ifconfig -a
    The MULTICAST flag is often an indication that the host supports multicasting.
    ping -b Bcast_address

## 1.10 Unix standards
In this text, we will refer to this standard as simply The POSIX Specification.

POSIX is an acronym for Portable Operating System Interface. POSIX is not a single standard, but a family of standards being developed by the Institute for Electrical and Electronics Engineers, Inc., normally called the IEEE. The POSIX standards have also been adopted as international standards by ISO and the International Electrotechnical Commission (IEC), called ISO/IEC.

The focus of this book is on The Single Unix Specification Version 3, with our main focus on the sockets API. Whenever possible we will use the standard functions. The above brief backgrounds on POSIX and The Open Group both continue with The Austin Group's publication of [The Single Unix Specification Version 3](http://www.UNIX.org/version3), as mentioned at the beginning of this section. 


The Internet standards process is documented in RFC 2026 [Bradner 1996]. Internet standards normally deal with protocol issues and not with programming APIs.

## 1.11 64-bit architectures
The model that is becoming most prevalent for 64-bit Unix systems is called the LP64 model, meaning only long integers (L) and pointers (P) require 64 bits.(char 1-bit, short 2-bits, int 4-bits).

The networking API problem is that some drafts of POSIX.1g specified that function arguments containing the size of a socket address structures have the size_t datatype (e.g., the third argument to bind and connect).

The solution is to use datatypes designed specifically to handle these scenarios. The sockets API uses the *socklen_t* datatype for lengths of socket address structures, and [XTI](www.hob.de/support/xti_forum/unix_network_programming.pdf) (X/Open Transport Interface) uses the t_scalar_t and t_uscalar_t datatypes. The reason for not changing these values from 32 bits to 64 bits is to make it easier to provide binary compatibility on the new 64-bit systems for applications compiled under 32-bit systems.

## 1.12 Summary
The Single Unix Specification Version 3, known by several other names and called simply The POSIX Specification by us, is the confluence of two long-running standards efforts, finally drawn together by The Austin Group.

Readers interested in the history of Unix networking should consult [Salus 1994] for a description of Unix history, and [Salus 1995] for the history of TCP/IP and the Internet.


## Chapter 2, the transport Layer : TCP,UDP,SCTP
We cover various topics in this chapter that fall into this category: TCP's three-way handshake, TCP's connection termination sequence, and TCP's TIME_WAIT state; SCTP's four-way handshake and SCTP's connection termination; plus SCTP, TCP, and UDP buffering by the socket layer, and so on.

## 2.2 the big picture
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/02fig01.gif "all protocols")

tcpdump, communicates directly with the datalink using either the BSD packet filter (BPF) or the datalink provider interface (DLPI). We mark the dashed line beneath the nine applications on the right as the API, which is normally sockets or XTI. The interface to either BPF or DLPI does not use sockets or XTI. (There is an exception to this, which we will describe in more detail in Chapter 28: Linux provides access to the datalink using a special type of socket called SOCK_PACKET.)

Explanations of the protocols:

1. IPv4, Internet Protocol version 4. IPv4, which we often denote as just IP, has been the workhorse protocol of the IP suite since the early 1980s. It uses 32-bit addresses. IPv4 provides packet delivery service for TCP, UDP, SCTP, ICMP, and IGMP.
2. IPv6, Internet Protocol version 6. IPv6 was designed in the mid-1990s as a replacement for IPv4. The major change is a larger address comprising 128 bits, to deal with the explosive growth of the Internet in the 1990s. IPv6 provides packet delivery service for TCP, UDP, SCTP, and ICMPv6. We often use the word "IP" as an adjective, as in IP layer and IP address, when the distinction between IPv4 and IPv6 is not needed.
3. TCP, Transmission Control Protocol. TCP is a connection-oriented protocol that provides a reliable, full-duplex(双重的) byte stream to its users. TCP sockets are an example of stream sockets. TCP takes care of details such as acknowledgments, timeouts, retransmissions, and the like. Most Internet application programs use TCP. Notice that TCP can use either IPv4 or IPv6.
4. UDP, User Datagram Protocol. UDP is a connectionless protocol, and UDP sockets are an example of datagram sockets. There is no guarantee that UDP datagrams ever reach their intended destination. As with TCP, UDP can use either IPv4 or IPv6.
5. SCTP, Stream Control Transmission Protocol. SCTP is a connection-oriented protocol that provides a reliable full-duplex association. **The word "association" is used when referring to a connection in SCTP** because SCTP is **multihomed, involving a set of IP addresses and a single port for each side of an association**. SCTP provides a message service, which maintains record boundaries. As with TCP and UDP, SCTP can use either IPv4 or IPv6, but it can also use both IPv4 and IPv6 simultaneously on the same association.
6. ICMP, Internet Control Message Protocol. ICMP handles error and control information between routers and hosts. These messages are normally generated by and processed by the TCP/IP networking software itself, not user processes, although we show the ping and traceroute programs, which use ICMP. We sometimes refer to this protocol as ICMPv4 to distinguish it from ICMPv6.
7. IGMP, Internet Group Management Protocol. IGMP is used with multicasting (Chapter 21), which is optional with IPv4.
8. ARP, Address Resolution Protocol. ARP maps an IPv4 address into a hardware address (such as an Ethernet address). ARP is normally used on broadcast networks such as Ethernet, token ring, and FDDI, and is not needed on point-to-point networks.
9. RARP, Reverse Address Resolution Protocol. RARP maps a hardware address into an IPv4 address. It is sometimes used when a diskless node is booting.
10. ICMPv6, Internet Control Message Protocol version 6. ICMPv6 combines the functionality of ICMPv4, IGMP, and ARP.
11. BPF, BSD packet filter. This interface provides access to the datalink layer. It is normally found on Berkeley-derived kernels.
12. DLPI, Datalink provider interface. This interface also provides access to the datalink layer. It is normally provided with SVR4.

Each Internet protocol is defined by one or more documents called a Request for Comments (RFC), which are their formal specifications. The solution to Exercise 2.1 shows how to obtain RFCs.
We use the terms "IPv4/IPv6 host" and "dual-stack host" to denote hosts that support both IPv4 and IPv6.

## 2.3 User Datagram Protocol (UDP)
UDP is a simple transport-layer protocol. It is described in RFC 768 [Postel 1980]. The application writes a message to a UDP socket, which is then encapsulated in a UDP datagram, which is then further encapsulated as an IP datagram, which is then sent to its destination. There is no guarantee that a UDP datagram will ever reach its final destination, that order will be preserved across the network, or that datagrams arrive only once.

If we want to be certain that a datagram reaches its destination, we can build lots of features into our application: acknowledgments from the other end, timeouts, retransmissions, and the like.

Each UDP datagram has a length. The length of a datagram is passed to the receiving application along with the data. We have already mentioned that TCP is a byte-stream protocol, without any record boundaries at all (Section 1.2), which differs from UDP.

Connectionless. We also say that UDP provides a connectionless service, as there need not be any long-term relationship between a UDP client and server. (client -> multiple servers, server receive from different clients)

## 2.4 Transmission Control Protocol (TCP)
A TCP client establishes a connection with a given server, exchanges data with that server across the connection, and then terminates the connection.

1. reliability. When TCP sends data to the other end, it requires an acknowledgment in return. If an acknowledgment is not received, TCP automatically retransmits the data and waits a longer amount of time. After some number of retransmissions, TCP will give up, with the total amount of time spent trying to send data typically between 4 and 10 minutes. Note that **TCP does not guarantee that the data will be received by the other endpoint, as this is impossible**. It delivers data to the other endpoint if possible, and notifies the user (by giving up on retransmissions and breaking the connection) if it is not possible. Therefore, **TCP cannot be described as a 100% reliable protocol; it provides reliable delivery of data or reliable notification of failure**.
2. estimate the **round-trip time (RTT)** between a client and server dynamically so that it knows how long to wait for an acknowledgment, continuously estimates the RTT of a given connection
3. sequences the data by associating a sequence number with every byte that it sends.There is no reliability provided by UDP. UDP itself does not provide anything like acknowledgments, sequence numbers, RTT estimation, timeouts, or retransmissions. If a UDP datagram is duplicated in the network, two copies can be delivered to the receiving host. Also, if a UDP client sends two datagrams to the same destination, **they can be reordered by the network and arrive out of order**. UDP applications must handle all these cases, as we will show in Section 22.5.
4. provides flow control. TCP always tells its peer exactly how many bytes of data it is willing to accept from the peer at any one time. This is called the advertised window. UDP provides no flow control. It is easy for a fast UDP sender to transmit datagrams at a rate that the UDP receiver cannot keep up with, as we will show in Section 8.13.
5. a TCP connection is full-duplex. After a full-duplex connection is established, it can be turned into a simplex connection if desired (see Section 6.6). UDP can be full-duplex.

## 2.5 Stream Control Transmission Protocol (SCTP)
SCTP provides** associations** between clients and servers. SCTP also provides applications with reliability, sequencing, flow control, and full-duplex data transfer, like TCP. 
(The word "association" is used in SCTP instead of "connection" to avoid the connotation that a connection involves communication between only two IP addresses. An association refers to a communication between two systems, which may involve more than two addresses due to multihoming.)

SCTP is message-oriented. It provides sequenced delivery of individual records. Like UDP, the length of a record written by the sender is passed to the receiving application.

SCTP can provide multiple streams between connection endpoints, each with its own reliable sequenced delivery of messages. A lost message in one of these streams does not block delivery of messages in **any of the other streams**. This approach is in contrast to TCP, where a loss at any point in the single stream of bytes blocks delivery of all future data on the connection until the loss is repaired.

Similar robustness can be obtained from TCP with help from routing protocols. For example, BGP connections within a domain (iBGP) often use addresses that are assigned to a virtual interface within the router as the endpoints of the TCP connection. The domain's routing protocol ensures that if there is a route between two routers, it can be used, which would not be possible if the addresses used belonged to an interface that went down. (SCTP's multihoming feature allows hosts to multihome, not just routers, and allows this multihoming to occur across different service providers, which the routing-based TCP method cannot allow.)

## 2.6 TCP Connection Establishment and Termination

---
Three way handshake:

1. The server must be prepared to accept an incoming connection. This is normally done by calling socket, bind, and listen and is called a passive open.
2. The client issues an active open by calling connect. This causes the client TCP to send a "synchronize" (SYN) segment, which tells the server the client's initial sequence number for the data that the client will send on the connection. Normally, there is no data sent with the SYN; it just contains an IP header, a TCP header, and possible TCP options.
3. The server must acknowledge (ACK) the client's SYN and the server must also send its own SYN containing the initial sequence number for the data that the server will send on the connection. The server sends its SYN and the ACK of the client's SYN in a single segment.
4. The client must acknowledge the server's SYN.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/tcp_handshake.gif "handshake")

The acknowledgment number in an ACK is the next expected sequence number for the end sending the ACK. Since a SYN occupies one byte of the sequence number space, the acknowledgment number in the ACK of each SYN is the initial sequence number plus one. Similarly, the ACK of each FIN is the sequence number of the FIN plus one.

The *socket* function is the equivalent of having a telephone to use.
 
*bind* is telling other people your telephone number so that they can call you.
 
*listen* is turning on the ringer so that you will hear when an incoming call arrives. 

*connect* requires that we know the other person's phone number and dial it. 

*accept* is when the person being called answers the phone. Having the client's identity returned by accept (where the identify is the client's IP address and port number) is similar to *having the caller ID feature show the caller's phone number*. One difference, however, is that accept returns the client's identity only after the connection has been established, whereas the caller ID feature shows the caller's phone number before we choose whether to answer the phone or not. 

If the DNS is used (Chapter 11), it provides a service analogous to a telephone book. 

*getaddrinfo* is similar to looking up a person's phone number in the phone book. 

*getnameinfo* would be the equivalent of having a phone book sorted by telephone numbers that we could search, instead of a book sorted by name.

---
TCP options

1. MSS option. With this option, the TCP **sending the SYN announces its maximum segment size**, **the maximum amount of data that it is willing to accept in each TCP segment**, on this connection. The sending TCP uses the receiver's MSS value as the maximum size of a segment that it sends. We will see how to fetch and set this TCP option with the TCP_MAXSEG socket option
2. Window scale option. The maximum window that either TCP can advertise to the other TCP is 65,535, because the corresponding field in the TCP header occupies 16 bits. Both end-systems must support this option for the window scale to be used on a connection. We will see how to affect this option with the SO_RCVBUF socket option.  TCP can send the option with its SYN as part of an active open. But, it can scale its windows only if the other end also sends the option with its SYN. Similarly, the server's TCP can send this option only if it receives the option with the client's SYN. This logic assumes that implementations ignore options that they do not understand, which is required and common, but** unfortunately, not guaranteed with all implementations**.
3. Timestamp option.This option is needed for high-speed connections to prevent possible data corruption caused by old, delayed, or duplicated segments.

---
TCP connection termination

1. One application calls close first, and we say that this end performs the active close. This end's TCP sends a FIN segment, which means it is finished sending data.
2. The other end that receives the FIN performs the **passive close**. The received FIN is acknowledged by TCP. **The receipt of the FIN is also passed to the application as an end-of-file (after any data that may have already been queued for the application to receive)**, since the receipt of the FIN means the application will not receive any additional data on the connection.
3. Sometime later, the application that received the end-of-file will close its socket. This causes its TCP to send a FIN.
4. The TCP on the system that receives this final FIN (the end that did the active close) acknowledges the FIN.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/tcp_termination.gif "tcp_termination")

A FIN occupies one byte of sequence number space just like a SYN. Therefore, the ACK of each FIN is the sequence number of the FIN plus one.

Between Steps 2 and 3 it is possible for data to flow from the end doing the passive close to the end doing the active close. This is called a half-close and we will talk about this in detail with the shutdown function.

The sending of each FIN occurs when a socket is closed. We indicated that the application calls close for this to happen, but realize that when a Unix process terminates, either voluntarily (calling exit or having the main function return) or involuntarily (receiving a signal that terminates the process), all open descriptors are closed, which will also cause a FIN to be sent on any TCP connection that is still open.

Either end—the client or the server—can perform the active close. Often the client performs the active close, but with some protocols (notably HTTP/1.0), the server performs the active close.

---
TCP State transition diagram

[IBM netstat usage](https://www-01.ibm.com/support/knowledgecenter/SSLTBW_1.13.0/com.ibm.zos.r13.halu101/concepts.htm%23concepts)

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/tcp_state.png "tcp state")

**Figure 2.4. TCP state transition diagram.**

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/tcpstates.gif "tcp states")

We denote the normal client transitions with a darker solid line and the normal server transitions with a darker dashed line. We also note that there are two transitions that we have not talked about: **a simultaneous open** (when both ends send SYNs at about the same time and the SYNs cross in the network, SYN_RCVD, look carefully in the picture) and a simultaneous close (when both ends send FINs at the same time).

---
Watch the packets

**Figure 2.5. Packet exchange for TCP connection.**

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/tcp_packet.gif "tcp packet")

The client in this example announces an MSS of 536 (indicating that it implements only the minimum reassembly buffer size) and the server announces an MSS of 1,460 (typical for IPv4 on an Ethernet). It is okay for the MSS to be different in each direction.

Notice that the acknowledgment of the client's request is sent with the server's reply. This is called *piggybacking* and will normally happen when the time it takes the server to process the request and generate the reply is less than around 200 ms.

It is important to notice in Figure 2.5 that if the entire purpose of this connection was to send a one-segment request and receive a one-segment reply, there would be **eight segments of overhead** involved when using TCP. If UDP was used instead, only two packets would be exchanged: the request and the reply.

Another important feature provided by TCP is congestion control(拥塞控制). It is important to understand that many applications are built using UDP because the **application exchanges small amounts of data and UDP avoids the overhead of TCP connection establishment and connection termination**.

## 2.7 TIME_WAIT State
We can see in Figure 2.4 that the end that performs the active close goes through this state. The duration that this endpoint remains in this state is twice the maximum segment lifetime (MSL), sometimes called 2MSL.

Every implementation of TCP must choose a value for the MSL. The recommended value in RFC 1122 is 2 minutes, although Berkeley-derived implementations have traditionally used a value of 30 seconds instead.

*(The way in which a packet gets "lost" in a network is usually **the result of routing anomalies（异常）**. A router crashes or a link between two routers goes down and it takes the routing protocols seconds or minutes to stabilize and find an alternate path. During that time period, **routing loops** can occur (router A sends packets to router B, and B sends them back to A) and packets can get caught in these loops. In the meantime, assuming the lost packet is a TCP segment, the sending TCP times out and retransmits the packet, and the retransmitted packet gets to the final destination by some alternate path. But sometime later (up to MSL seconds after the lost packet started on its journey), the routing loop is corrected and the packet that was lost in the loop is sent to the final destination. **This original packet is called a lost duplicate or a wandering duplicate**)*

There are two reasons for the TIME_WAIT state:

1. To implement TCP's full-duplex connection termination reliably
2. To allow old duplicate segments to expire in the network

Looking at Figure 2.5 and assuming that the final ACK is lost. The server will resend its final FIN, so the client must maintain state information, allowing it to resend the final ACK. This example also shows why the end that performs the active close is the end that remains in the TIME_WAIT state: because that end is the one that might have to retransmit the final ACK.

Establish another connection with same ip and port. **This latter connection is called an incarnation(化身) of the previous connection** since the IP addresses and ports are the same. TCP must prevent old duplicates from a connection from reappearing at some later time and being misinterpreted as belonging to a new incarnation of the same connection. To do this, **TCP will not initiate a new incarnation of a connection that is currently in the TIME_WAIT state**. Since the duration of the TIME_WAIT state is twice the MSL, **this allows MSL seconds for a packet in one direction to be lost, and another MSL seconds for the reply to be lost**. By enforcing this rule, we are guaranteed that when we successfully establish a TCP connection, all old duplicates from previous incarnations of the connection have expired in the network.

[Read RFC 1337 to understand](https://www.ietf.org/rfc/rfc1337)

*(There is an exception to this rule. Berkeley-derived implementations will initiate a new incarnation of a connection that is currently in the TIME_WAIT state** if the arriving SYN has a sequence number that is "greater than" the ending sequence number from the previous incarnation**. Pages 958–959 of TCPv2 talk about this in more detail. This requires the server to perform the active close, since the TIME_WAIT state must exist on the end that receives the next SYN. This capability is used by the rsh command. RFC 1185 [Jacobson, Braden, and Zhang 1990] talks about some pitfalls in doing this.)*

## 2.8 SCTP Association Establishment and Termination
