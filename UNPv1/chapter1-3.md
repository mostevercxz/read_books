Ignored chapters:
2.8 SCTP

Questions:

1. It is remarkable how many network break-ins have occurred by a hacker sending data to cause a server's call to sprintf to overflow its buffer. How to do this?
2. Additional tips on writing secure network programs are found in Chapter 23 of [Garfinkel, Schwartz, and Spafford 2003]. How to write secure network programs?
3. If we want to be certain that a datagram reaches its destination, we can build lots of features into our application: acknowledgments from the other end, timeouts, retransmissions, and the like. How to make sure a datagram reaches its destination?
4. After a full-duplex connection is established, it can be turned into a simplex connection if desired (see Section 6.6). How to do it?
5. TIME_WAIT State, To allow old duplicate segments to expire in the network, how to understand this?  for a packet in one direction to be lost, and another MSL seconds for the reply to be lost, what does that mean? (至今还不解。。SO上的回答也不解，W. Richard Stevens is dead.. reply to be lost应该指的是使最后的ACK失效的时间, packet in one direction to be lost应该指的是server发过来的FIN消失的时间)
6. What's going on if the last ACK nerver received? The passive closer will retransmit the FIN segment, just like any other
segment that doesn't get ACKed. If the other end never ACKs it,and the 2MSL timer passes, the other end should then respond with RST. If it doesn't respond at all (maybe the other end has crashed), eventually the retransmit timer will be reached and the connection will be aborted.
7. how many times tcp retransmits? windows(TcpMaxDataRetransmissions), linux(man tcp, tcp_retries1, 3 tcp_retries2, 15, RTO, ss -i)
8. How to calculate RTO? ss -i [BLOG](http://sgros.blogspot.jp/2012/02/calculating-tcp-rto.html)

[TCP/IP illustrated volume 1](http://www.pcvr.nl/tcpip/)

Answers:
[2.7 Please explain the TIME_WAIT state.](http://www.softlab.ntua.gr/facilities/documentation/unix/unix-socket-faq/unix-socket-faq-2.html)

[time_wait's effect](http://www.isi.edu/~faber/pubs/time_wait.pdf)

It must wait a reasonable amount of time to see whether the FIN segment from Program B is retransmitted, indicating that Program B never received the final ACK segment from Program A. In that case, Program A must be able to retransmit the final ACK segment. The Program A socket cannot be freed until this time period has elapsed. The time period is defined as twice the maximum segment life time, normally in the range of 1 to 4 minutes, depending on the TCP implementation.

**When TCP sends a group of segments, it normally sets a single retransmission timer**, waiting for the other end to acknowledge reception. TCP does not set a different retransmission timer for every segment. Rather, it sets a timer when it sends a window of data and updates the timeout as ACKs arrive. **If an acknowledgment is not received in time, a segment is retransmitted.**

Given the MSL value for an implementation, the rule is: When TCP performs an active close and sends the final ACK, that connection must stay in the TIME_ WAIT state for twice the MSL. **This lets TCP resend the final ACK in case it is lost.**
The final ACK is resent not because the TCP retransmits ACKs (they do not consume sequence numbers and are not retransmitted by TCP), but **because the other side will retransmit its FIN (which does consume a sequence number)**. Indeed, TCP will always retransmit FINs until it receives a final ACK.

**For a packet in one direction to be lost, and another MSL seconds for the reply to be lost.**

TIME_WAIT State, I also have the same question about this. I thought of an assumption.
In extreme senario, suppose the last ACK from client end spends MSL to reach server end. At this point, the end point thinks that this ACK has already perished due to MSL timeout. So server end retransmit the FIN imediately. In order to assure this FIN can reach client end (or if not, we have to assure its perishment), we must have another MSL. 	
**The server waits for RTO(retransmission timeout), other than MSL to retransmit the FIN, right?** I think the server doesn't care anything about MSL, all it would do is retransmitting FIN until an ACK arrives or FIN retries exceed

Acknowledges:

1. International Organization for Standardization (ISO)
2. open systems interconnection (OSI)
3. Ethernet maximum transfer unit (MTU)
4. Application Programming Interfaces (APIs)
5. Computer Systems Research Group (CSRG)
6. Internet Engineering Task Force (IETF)
7. IPV4 header

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/ipv4.gif "ipv4")

# Chapter 1, introduction
A protocol, an agreement on how those programs will communicate. Before delving into the design details of a protocol, high-level decisions must be made about which program is expected to initiate communication and when responses are expected.

**Client and server on the same Ethernet communicating using TCP**
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/1/01fig03.gif "cs communicate layer")

Routers are the building blocks of WANs. The largest WAN today is the Internet.

## 1.2 a simple daytime application

	// Redhat 6.4
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

	struct in6_addr{
		union{
			uint8_t __u6_addr8[16];
		} __in6_u;
	};

	struct sockaddr_in6{
		__SOCKADDR_COMMON (sin6_);
		in_port_t sin6_port;
		uint32_t sin6_flowwinfo;// IPv6 flow infomation
		struct in6_addr sin6_addr;
		uint32_t sin6_scope_id;//IPv6 scope-id
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
2. The other end that receives the FIN performs the **passive close**. The received FIN is acknowledged by TCP. **The receipt of the FIN is also passed to the application as an end-of-file (after any data that may have already been queued for the application to receive)**, since the receipt(收到) of the FIN means the application will not receive any additional data on the connection.
3. Sometime later, the application that received the end-of-file will close its socket. This causes its TCP to send a FIN.
4. The TCP on the system that receives this final FIN (the end that did the active close) acknowledges the FIN.

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/tcp_termination.gif "tcp_termination")

Since a FIN and an ACK are required in each direction, four segments are normally required. We use the qualifier "normally" because in some scenarios, the FIN in Step 1 is sent with data. Also, the segments in Steps 2 and 3 are both from the end performing the passive close and could be combined into one segment.

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

Looking at Figure 2.5 and assuming that the final ACK is lost. The server will resend its final FIN, so the client must maintain state information, allowing it to resend the final ACK. If it did not maintain this information, it would respond with an RST (a different type of TCP segment), which would be interpreted by the server as an error. If TCP is performing all the work necessary to terminate both directions of data flow cleanly for a connection (its full-duplex close), then it must correctly handle the loss of any of these four segments. This example also shows why the end that performs the active close is the end that remains in the TIME_WAIT state: because that end is the one that might have to retransmit the final ACK.

Establish another connection with same ip and port. **This latter connection is called an incarnation(化身) of the previous connection** since the IP addresses and ports are the same. TCP must prevent old duplicates from a connection from reappearing at some later time and being misinterpreted as belonging to a new incarnation of the same connection. To do this, **TCP will not initiate a new incarnation of a connection that is currently in the TIME_WAIT state**. Since the duration of the TIME_WAIT state is twice the MSL, **this allows MSL seconds for a packet in one direction to be lost, and another MSL seconds for the reply to be lost**. By enforcing this rule, we are guaranteed that when we successfully establish a TCP connection, all old duplicates from previous incarnations of the connection have expired in the network.

[Read RFC 1337 to understand](https://www.ietf.org/rfc/rfc1337)

*(There is an exception to this rule. Berkeley-derived implementations will initiate a new incarnation of a connection that is currently in the TIME_WAIT state** if the arriving SYN has a sequence number that is "greater than" the ending sequence number from the previous incarnation**. Pages 958–959 of TCPv2 talk about this in more detail. This requires the server to perform the active close, since the TIME_WAIT state must exist on the end that receives the next SYN. This capability is used by the rsh command. RFC 1185 [Jacobson, Braden, and Zhang 1990] talks about some pitfalls in doing this.)*

## 2.8 SCTP Association Establishment and Termination
---
Four-way handshake

1. **server passive open**(socket, bind, listen)
2. The client **issues an active open** by calling connect or by sending a message, which implicitly opens the association. This causes the client SCTP to send an INIT message (which stands for "initialization") to tell the server the client's list of IP addresses, initial sequence number, initiation tag to identify all packets in this association, number of outbound streams the client is requesting, and number of inbound streams the client can support.
3. The server acknowledges the client's INIT message with an INIT-ACK message, which contains the server's list of IP addresses, initial sequence number, initiation tag, number of outbound streams the server is requesting, number of inbound streams the server can support, and a state cookie. The state cookie contains all of the state that the server needs to ensure that the association is valid, and is digitally signed to ensure its validity.
4. The client echos the server's state cookie with a COOKIE-ECHO message. This message may also contain user data bundled within the same packet.
5. The server acknowledges that the cookie was correct and that the association was established with a COOKIE-ACK message. This message may also contain user data bundled within the same packet.

**Figure 2.6. SCTP four-way handshake.**

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/sctp_handshake.gif "sctp handshake")

---
Association Termination

SCTP State Transition Diagram

Watching the Packets

SCTP Options

## 2.9 Port Numbers

The port numbers are divided into 3 ranges:

1. The well-known ports: 0 through 1023. These port numbers are controlled and assigned by the IANA.
2. The registered ports: 1024 through 49151
3. The dynamic or private ports, 49152 through 65535. The IANA says nothing about these ports. These are what we call ephemeral ports. (The magic number 49152 is three-fourths of 65536.)

**Figure 2.10. Allocation of port numbers.**

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/port_allocation.gif "port allocation")

Unix systems have the concept of a reserved port, which is any port less than 1024. These ports can only be assigned to a socket by an appropriately privileged process. All the IANA well-known ports are reserved ports; hence, the server allocating this port (such as the FTP server) must have superuser privileges when it starts.

There are a few clients (not servers) that require a reserved port as part of the client/server authentication: the rlogin and rsh clients are common examples. These clients call the library function rresvport to create a TCP socket and assign an unused port in the range 513–1023 to the socket. This function normally tries to bind port 1023, and if that fails, it tries to bind 1022, and so on, until it either succeeds or fails on port 513.(man rresvport)

**Socket Pair:**

The socket pair for a TCP connection is the four-tuple that defines the two endpoints of the connection: the local IP address, local port, foreign IP address, and foreign port.

A socket pair uniquely identifies every TCP connection on a network. For SCTP, an association is identified by a set of local IP addresses, a local port, a set of foreign IP addresses, and a foreign port.

The two values that identify each endpoint, an IP address and a port number, are often called a socket.

## 2.10 TCP Port Numbers and concurrent servers
When we specify the local IP address as an asterisk, it is called the wildcard character.

the server can specify that it wants only to accept incoming connections that arrive destined to one specific local interface. **This is a one-or-any choice for the server**. The server cannot specify a list of multiple addresses.

Notice from this example that TCP cannot demultiplex incoming segments by looking at just the destination port number. TCP must look at all four elements in the socket pair to determine which endpoint receives an arriving segment.

## 2.11 Buffer Sizes and Limitations

Some limits:

1. **The maximum size of an IPv4 datagram is 65,535 bytes**, including the IPv4 header.This is because of the 16-bit total length field in ipv4 header.
2. **The maximum size of an IPv6 datagram is 65,575 bytes**, including the 40-byte IPv6 header. This is because of the 16-bit payload length field in ipv6 header. Notice that the IPv6 payload length field does not include the size of the IPv6 header, while the IPv4 total length field does include the header size.
3. Many networks have an MTU which can be dictated by the hardware. For example, the Ethernet MTU is 1,500 bytes. Other datalinks, such as point-to-point links using the Point-to-Point Protocol (PPP), have a configurable MTU. Older SLIP links often used an MTU of 1,006 or 296 bytes. T**he minimum link MTU for IPv4 is 68 bytes. This permits a maximum-sized IPv4 header (20 bytes of fixed header, 30 bytes of options) and minimum-sized fragment (the fragment offset is in units of 8 bytes)**. The minimum link MTU for IPv6 is 1,280 bytes. IPv6 can run over links with a smaller MTU, but requires link-specific fragmentation and reassembly to make the link appear to have an MTU of at least 1,280 bytes.
4. The smallest MTU in the path between two hosts is called the path MTU. Today, the Ethernet MTU of 1,500 bytes is often the path MTU. The path MTU need not be the same in both directions between any two hosts because routing in the Internet is often asymmetric [Paxson 1996]. That is, the route from A to B can differ from the route from B to A.
5. When an IP datagram is to be sent out an interface, if the size of the datagram exceeds the link MTU, fragmentation is performed by both IPv4 and IPv6. The fragments are not normally reassembled until they reach the final destination. IPv4 hosts perform fragmentation on datagrams that they generate and IPv4 routers perform fragmentation on datagrams that they forward. But with IPv6, only hosts perform fragmentation on datagrams that they generate; IPv6 routers do not fragment datagrams that they are forwarding. (The IP datagrams generated by the router's Telnet server are generated by the router, not forwarded by the router.)
6. If the "don't fragment" (DF) bit is set in the IPv4 header, it specifies that this datagram must not be fragmented, either by the sending host or by any router. A router that receives an IPv4 datagram with the DF bit set whose size exceeds the outgoing link's MTU generates an ICMPv4 "destination unreachable, fragmentation needed but DF bit set" error message. Since IPv6 routers do not perform fragmentation, there is an implied DF bit with every IPv6 datagram. When an IPv6 router receives a datagram whose size exceeds the outgoing link's MTU, it generates an ICMPv6 "packet too big" error message.(The IPv4 DF bit and its implied IPv6 counterpart can be used for path MTU discovery, while Path MTU discovery is problematic in the Internet today; many firewalls drop all ICMP messages, including the fragmentation required message)
7. IPv4 and IPv6 define a minimum reassembly buffer size, the minimum datagram size that we are guaranteed any implementation must support.For IPv4, this is 576 bytes. IPv6 raises this to 1,500 bytes.
8. TCP has a maximum segment size (MSS) that announces to the peer TCP the maximum amount of TCP data that the peer can send per segment. **The goal of the MSS is to tell the peer the actual value of the reassembly buffer size and to try to avoid fragmentation**. The MSS is often set to the interface MTU minus the fixed sizes of the IP and TCP headers. On an Ethernet using IPv4, this would be 1,460, and on an Ethernet using IPv6, this would be 1,440. (The TCP header is 20 bytes for both, but the IPv4 header is 20 bytes and the IPv6 header is 40 bytes.) The MSS value in the TCP MSS option is a 16-bit field, limiting the value to 65,535. This is fine for IPv4, since the maximum amount of TCP data in an IPv4 datagram is 65,495 (65,535 minus the 20-byte IPv4 header and minus the 20-byte TCP header).

---
TCP output
**Figure 2.15. Steps and buffers involved when an application writes to a TCP socket.**

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/tcp_steps.gif "tcp steps")

Every TCP socket has a send buffer and we can change the size of this buffer with the SO_SNDBUF socket option. Therefore, the successful return from a write to a TCP socket only tells us that we can reuse our application buffer. It does not tell us that either the peer TCP has received the data or that the peer application has received the data. (We will talk about this more with the SO_LINGER socket option in Section 7.5.)

**TCP takes the data in the socket send buffer and sends it to the peer TCP based on all the rules of TCP data transmission** (Chapter 19 and 20 of TCPv1). The peer TCP must acknowledge the data, and as the ACKs arrive from the peer, only then can our TCP discard the acknowledged data from the socket send buffer. **TCP must keep a copy of our data until it is acknowledged by the peer.**

TCP sends the data to IP in MSS-sized or smaller chunks, prepending its TCP header to each segment, where the MSS is the value announced by the peer, **or 536(576-20-20) if the peer did not send an MSS option**. IP prepends its header, searches the routing table for the destination IP address (the matching routing table entry specifies the outgoing interface), and passes the datagram to the appropriate datalink.

---
UDP output

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/udp_steps.gif "udp steps")

This time, we show the socket send buffer as a dashed box because it doesn't really exist. A UDP socket has a send buffer size (which we can change with the SO_SNDBUF socket option, Section 7.5), but this is simply an upper limit on the maximum-sized UDP datagram that can be written to the socket. If an application writes a datagram larger than the socket send buffer size, EMSGSIZE is returned.

## 2.12 Standard internet services
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/standard_services.gif "standard services")

telnet baidu.com http

These service names are mapped into the port numbers shown in Figure 2.18 by the /etc/services file.

These "simple services" are often disabled by default on modern systems due to denial-of-service and other resource utilization attacks against them.

## 2.13 Protocol Usage by Common Internet Applications
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/2/protocol_usage.gif "protocol usage")

## 2.14 Summary
TCP establishes connections using a three-way handshake and terminates connections using a four-packet exchange. When a TCP connection is established, it goes from the CLOSED state to the ESTABLISHED state, and when it is terminated, it goes back to the CLOSED state. There are 11 states in which a TCP connection can be, and a state transition diagram gives the rules on how to go between the states. Understanding this diagram is essential to diagnosing problems using the netstat command and understanding what happens when an application calls functions such as connect, accept, and close.

TCP's TIME_WAIT state is a continual source of confusion with network programmers. This state exists to implement TCP's full-duplex connection termination (i.e., to handle the case of the final ACK being lost), and to allow old duplicate segments to expire in the network.

## Exercies
2.1 What about other IP versions ?

    Decimal,Keyword,Version,Reference
    0-1,,Reserved,[Jon_Postel][RFC4928]
    2-3,,Unassigned,[Jon_Postel]
    4,IP,Internet Protocol,[RFC791][Jon_Postel]
    5,ST,ST Datagram Mode,[RFC1190][Jim_Forgie]
    6,IPv6,Internet Protocol version 6,[RFC1752]
    7,TP/IX,TP/IX: The Next Internet,[RFC1475]
    8,PIP,The P Internet Protocol,[RFC1621]
    9,TUBA,TUBA,[RFC1347]
    10-14,,Unassigned,[Jon_Postel]
    15,Reserved,[Jon_Postel]

2.2 google IP version 5
2.3 536 = 576 - 20 TCP headers - 20 IPv4 headers
2.4 The server performs the active close.
2.5 TCP cannot exceed the MSS announced by the other end, but it can always send less than this amount.
2.6 (Wow, google...)The "Protocol Numbers" section of the Assigned Numbers Web page (http://www.iana.org/numbers.htm) shows a value of 89 for OSPF


# Chapter 3, Sockets introduction
## 3.2 Socket addresss structure
Each supported protocol suite defines its own socket address structure. The names of these structures begin with sockaddr_ and end with a unique suffix for each protocol suite.

1. The POSIX specification requires only three members in the structure: sin_family, sin_addr, and sin_port.Almost all implementations add the sin_zero member so that all socket address structures are at least 16 bytes in size.
2. Even if the length field is present, we need never set it and need never examine it, unless we are dealing with routing sockets (Chapter 18). It is used within the kernel by the routines that deal with socket address structures from various protocol families. (The four socket functions that pass a socket address structure from the process to the kernel, bind, connect, sendto, and sendmsg, all go through the sockargs function in a Berkeley-derived implementation; The five socket functions that pass a socket address structure from the kernel to the process, accept, recvfrom, recvmsg, getpeername, and getsockname, all set the sin_len member before returning to the process)
3. sa_family_t is normally an 8-bit unsigned integer if the implementation supports the length field, or **an unsigned 16-bit integer if the length field is not supported**(my ubuntu 15.04 is 16-bit)
4. Both the IPv4 address and the TCP or UDP port number are always stored in the structure in network byte order.
5. The reason the sin_addr member is a structure, and not just an in_addr_t, is historical. Earlier releases (4.2BSD) defined the in_addr structure as a union of various structures, to allow access to each of the 4 bytes and to both of the 16-bit values contained within the 32-bit IPv4 address.
6. The sin_zero member is unused, but we always set it to 0 when filling in one of these structures. By convention, we always set the entire structure to 0 before filling it in, not just the sin_zero member.
7. Socket address structures are used only on a given host: The structure itself is not communicated between different hosts, although certain fields (e.g., the IP address and port) are used for communication.

Comparison of Socket address structures, see /usr/include/netinet/in.h linux/socket.h linux/un.h
![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/3/comparison_structure.png "comparison")


## 3.3 value-result arguments
The reason that the size changes from an integer to be a pointer to an integer is because the size is both a value when the function is called (it tells the kernel the size of the structure so that the kernel does not write past the end of the structure when filling it in) and a result when the function returns (it tells the process how much information the kernel actually stored in the structure). This type of argument is called a value-result argument. 

The most common value-result argument is :

* the length of a returned socket address structure
* The middle three arguments for the select function
* The length argument for the getsockopt function
* The msg_namelen and msg_controllen members of the msghdr structure, when used with recvmsg
* The ifc_len member of the ifconf structure
* The first of the two length arguments for the sysctl function

code:

    #include <sys/select.h>
    #include <sys/time.h> 
    int select(int maxfdp1, fd_set *readset, fd_set *writeset, fd_set *exceptset, const struct timeval *timeout);
	struct timeval {long tv_sec; long tv_usec;}

There are three possibilities:

1. Wait forever— Return only when one of the specified descriptors is ready for I/O. For this, we specify the timeout argument as a null pointer.
2. Wait up to a fixed amount of time— Return when one of the specified descriptors is ready for I/O, but do not wait beyond the number of seconds and microseconds specified in the timeval structure pointed to by the timeout argument.
3. Do not wait at all— Return immediately after checking the descriptors. This is called polling. To specify this, the timeout argument must point to a timeval structure and the timer value (the number of seconds and microseconds specified by the structure) must be 0.

**The wait in the first two scenarios is normally interrupted if the process catches a signal and returns from the signal handler.**

The three middle arguments, readset, writeset, and exceptset, specify the descriptors that we want the kernel to test for reading, writing, and exception conditions. **There are only two exception conditions currently supported**:

1. The arrival of out-of-band data for a socket. We will describe this in more detail in Chapter 24.
2. The presence of control status information to be read from the master side of a pseudo-terminal that has been put into packet mode. We do not talk about pseudo-terminals in this book.

code:

    <linux/posix_types.h>
	#define __FD_SETSIZE 1024
	<sys/select.h>
    typedef long int __fd_mask;
    #define _NFDBITS (8 * (int)sizeof(__fd_mask))
	typedef struct
	{
		__fd_mask __fds_bits[__FD_SETSIZE / _NFDBITS];
	}fd_set;

## 3.4 byte-order functions
MSB(most significant)
LSB(least significant)
little-endian: LSB comes first(at the starting address)

Write a program to determine Big-or-Little Endian

    union 
    {
      short s;
      char c[sizeof(short)];
    }un;
    un.s = 0x0102;

Host byte order <--> Network byte order
But, both history and the POSIX specification say that certain fields in the socket address structures must be maintained in network byte order. So, we have to do it ourselves.

    #include <netinet/in.h>
     
    uint16_t htons(uint16_t host16bitvalue) ;//host -> network short     
    uint32_t htonl(uint32_t host32bitvalue) ;//host -> network long
     
    uint16_t ntohs(uint16_t net16bitvalue) ;
    uint32_t ntohl(uint32_t net32bitvalue) ;
 
We use the term "byte" to mean an 8-bit quantity since almost all current computer systems use 8-bit bytes. Most Internet standards use the term **octet** instead of **byte** to mean an 8-bit quantity. This started in the early days of TCP/IP because much of the early work was done on systems such as the DEC-10, which did not use 8-bit bytes.

Another important convention in Internet standards is bit ordering. This represents four bytes in the order in which they appear on the wire; the leftmost bit is the most significant. However, **the numbering starts with zero assigned to the most significant bit(MSB 0)**.[Wikipedia, bit numbering.](https://en.wikipedia.org/wiki/Bit_numbering) (11110011 00110011, network)This is a notation that you should become familiar with to make it easier to read protocol definitions in RFCs.

Little-endian uses LSB 0;big-endian uses both LSB 0 and MSB 0.

LSB:

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/3/Lsb0.svg.png "LSB")

MSB :

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/3/Msb0.svg.png "MSB")

![alt text](http://7xp1jz.com1.z0.glb.clouddn.com/unpv1/3/bit_order_example.gif "bit order")

## 3.5 Byte manipulation funcs
There are two groups of functions that operate on multibyte fields, without interpreting the data, and without assuming that the data is a null-terminated C string. The functions beginning with str (for string), defined by including the <string.h> header, deal with null-terminated C character strings.

The first group of functions, whose names begin with b (for byte), are from 4.2BSD and are still provided by almost any system that supports the socket functions.

The second group of functions, whose names begin with mem (for memory), are from the ANSI C standard and are provided with any system that supports an ANSI C library.

    #include <strings.h>
     
    void bzero(void *dest, size_t nbytes);     
    void bcopy(const void *src, void *dest, size_t nbytes);//bcopy is not back, src --> dest
    int bcmp(const void *ptr1, const void *ptr2, size_t nbytes);
 
    #include <string.h>
     
    void *memset(void *dest, int c, size_t len);     
    void *memcpy(void *dest, const void *src, size_t nbytes);
    int memcmp(const void *ptr1, const void *ptr2, size_t nbytes);
 
*bcopy* correctly handles overlapping fields, while the behavior of *memcpy* is undefined if the source and destination overlap. The ANSI C memmove function must be used when the fields overlap.

One way to remember the order of the two pointers for *memcpy* is to remember that they are written in the same left-to-right order as an assignment statement in C:

dest = src;

One way to remember the order of the final two arguments to *memset* is to realize that all of the ANSI C memXXX functions require a length argument, and it is always the final argument.

*memcmp*, the comparison is done assuming the two unequal bytes are unsigned chars.

## 3.6 funcs:
`inet_aton, inet_addr, inet_ntoa`

`inet_aton, inet_ntoa, inet_addr` convert an IPv4 address from a dotted-decimal string (e.g., "206.168.112.96") to its 32-bit network byte ordered binary value.

`inet_pton and inet_ntop`, handle both IPv4 and IPv6 addresses.

    #include <arpa/inet.h>
     
    int inet_aton(const char *strptr, struct in_addr *addrptr);//Returns: 1 if string was valid, 0 on error
     
    in_addr_t inet_addr(const char *strptr);//Deprecated, Returns: 32-bit binary network byte ordered IPv4 address; INADDR_NONE if error
     
    char *inet_ntoa(struct in_addr inaddr);//Do not use,Returns: pointer to dotted-decimal string

The problem with *inet_addr* is that all 2^32 possible binary values are valid IP addresses (0.0.0.0 through 255.255.255.255), but the function returns the constant **INADDR_NONE** (typically 32 one-bits) on an error. This means the dotted-decimal string 255.255.255.255 (the IPv4 limited broadcast address, Section 20.2) cannot be handled by this function since its binary value appears to indicate failure of the function.
 
## 3.7 funcs
    inet_pton, inet_ntop

The letters "p" and "n" stand for presentation and numeric. The presentation format for an address is often an ASCII string and the numeric format is the binary value that goes into a socket address structure

    #include <arpa/inet.h>
     
    int inet_pton(int family, const char *strptr, void *addrptr);//Returns: 1 if OK, 0 if input not a valid presentation format, -1 on error
     
    const char *inet_ntop(int family, const void *addrptr, char *strptr, size_t len);//Returns: pointer to result if OK, NULL on error

The family argument for both functions is either `AF_INET or AF_INET6`. If family is not supported, both functions return an error with errno set to EAFNOSUPPORT.

    defined in the <netinet/in.h> :
    #define INET_ADDRSTRLEN   16   /* for IPv4 dotted-decimal */
    #define INET6_ADDRSTRLEN  46   /* for IPv6 hex string */

If len is too small to hold the resulting presentation format, including the terminating null, **a null pointer is returned and errno is set to ENOSPC**.

The *strptr* argument to **inet_ntop** cannot be a null pointer. The caller must allocate memory for the destination and specify its size. On success, this pointer is the return value of the function.

    foo.sin_addr.s_addr = inet_addr(cp);
     -->
    inet_pton(AF_INET, cp, &foo.sin_addr);
    
    ptr = inet_ntoa(foo.sin_addr);
     -->
    char str[INET_ADDRSTRLEN];
    ptr = inet_ntop(AF_INET, &foo.sin_addr, str, sizeof(str));

    int
    inet_pton(int family, const char *strptr, void *addrptr)
    {
    if (family == AF_INET) {
    	struct in_addr in_val;
    
    	if (inet_aton(strptr, &in_val)) {
    		memcpy(addrptr, &in_val, sizeof(struct in_addr));
    		return (1);
    	}
    	return (0);
    }
    errno = EAFNOSUPPORT;
    return (-1);
    }


    const char *
    inet_ntop(int family, const void *addrptr, char *strptr, size_t len)
    {
    	const u_char *p = (const u_char *) addrptr;
    
    	if (family == AF_INET) {
    		char temp[INET_ADDRSTRLEN];
    
    		snprintf(temp, sizeof(temp), "%d.%d.%d.%d", p[0], p[1], p[2], p[3]);
    		if (strlen(temp) >= len) {
    		errno = ENOSPC;
    		return (NULL);
    	}
    	strcpy(strptr, temp);
    	return (strptr);
    	}
    	errno = EAFNOSUPPORT;
    	return (NULL);
    }

## 3.8 sock_ntop and related funcs

Problem : 

    const char * inet_ntop(family, const void *addrptr, char *strptr, size_t len)// need write struct sockaddr_in or sockaddr_in6, protocol-dependent

    char * sock_ntop(const struct sockaddr *addr, socklen_t addrLen);//return non-null pointer if OK,NULL on error

Notice that using static storage for the result prevents the function from being re-entrant or thread-safe. We made this design decision for this function to allow us to easily call it from the simple examples in the book.(I could use std::string)


    // used to replace inet_ntop, skipping the step of declaring sockaddr structures
    // currently used only in IPv4, IPv6
	// 0 indicates no error,we can use the presentation string; -1 means error
    int sock_ntop(const struct sockaddr *sa, socklen_t len, std::string &presentation)
    {
      char portBuffer[7]={0};
      char buffer[35] = {0};// the max length of IPv6 28 + 7;
      switch(sa->family)
      {
    	case AF_INET:
		{
      		struct sockaddr_in *ipv4Addr = (struct sockaddr_in *)sa;
      		if (inet_ntop(sa->family, &(ipv4Addr->sin_addr), buffer, sizeof(buffer)) )
      		{
    			uint16_t port = ntohs(ipv4Addr->sin_port);
	    		if (port != 0)
    			{
    				snprintf(portBuffer, sizeof(portBuffer), ":%d", port);
    				strcat(buffer, portBuffer);
    			}
    			presentation = buffer;
    	  	}
    	  	else
    	  	{
    			return -1;
    	  	}
		}// end case AF_INET
		break;
		case AF_INET6:
		{
			struct sockaddr_in6 *ipv6Addr = (struct sockaddr_in6 *)sa;
      		if (inet_ntop(sa->family, &(ipv6Addr->sin6_addr), buffer, sizeof(buffer)) )
      		{
    			uint16_t port = ntohs(ipv6Addr->sin6_port);
	    		if (port != 0)
    			{
    				snprintf(portBuffer, sizeof(portBuffer), ":%d", port);
    				strcat(buffer, portBuffer);
    			}
    			presentation = buffer;
    	  	}
    	  	else
    	  	{
    			return -1;
    	  	}
		}//end case AF_INET6
		break;
		default:
			return -1;
			break;
      }//end of switch

		return 0;
    }// end of func

## 3.9 Heading to readn, writen and readline funcs

A read or write on a stream socket might input or output fewer bytes than requested, but this is not an error condition. The reason is that buffer limits might be reached for the socket in the kernel. All that is required to input or output the remaining bytes is for the caller to invoke the read or write function again.

Nevertheless, we always call our writen function instead of write, in case the implementation returns a short count.

    ssize_t readn(int filedes, void *buff, size_t nbytes); 
    ssize_t writen(int filedes, const void *buff, size_t nbytes);
    ssize_t readline(int filedes, void *buff, size_t maxlen);
     
    All return: number of bytes read or written, –1 on error

Our three functions look for the error **EINTR** (the system call was interrupted by a caught signal, which we will discuss in more detail in Section 5.9) and continue reading or writing if the error occurs.


Good "defensive programming" techniques require these programs to not only expect their counterparts to follow the network protocol, but to check for unexpected network traffic as well. Using stdio to buffer data for performance flies in the face of these goals since the application has no way to tell if unexpected data is being held in the stdio buffers at any given time.

The desire to operate on lines comes up again and again. But our advice is to think in terms of buffers and not lines. Write your code to read buffers of data, and if a line is expected, check the buffer to see if it contains that line.

Source code inclued in read_write_readline.cpp

## 3.10 Summary

* value-result arguments
* self-defining, family field in socket structure
* inet_ntop, inet_pton, sock_* functions, getaddrinfo, getpeername, protocol-independent
* TCP sockets provide a byte stream to an application: There are no record markers. The return value from a read can be less than what we asked for, but this does not indicate an error.

## Exercises
3.1 if passed by value, the value will be destroyed when the function ends.
3.2 ISO C++ forbids incrementing a pointer of type 'const void*'
3.3 

	int inet_aton(const char *strptr, struct in_addr *addrptr);//Returns: 1 if string was valid, 0 on error
     
    in_addr_t inet_addr(const char *strptr);//Deprecated, Returns: 32-bit binary network byte ordered IPv4 address; INADDR_NONE if error

The inet_aton and inet_addr functions have traditionally been liberal in what they accept as a dotted-decimal IPv4 address string: allowing from one to four numbers separated by decimal points, and allowing a leading 0x to specify a hexadecimal number, or a leading 0 to specify an octal number. (telnet 0xe, telnet 011110111)