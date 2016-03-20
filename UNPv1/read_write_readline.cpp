#include <unistd.h>
#include <errno.h>

#define ssize_t int

ssize_t //read "n" bytes from a dscriptor
readn(int fd, void *vptr, size_t n)
{
  size_t nLeft = n;
  ssize_t nReadCount = 0;
  char *ptr = (char *)vptr;

  //read until no bytes left
  while (nLeft > 0)
  {
    if ((nReadCount = read(fd, ptr, nLeft)) < 0)
    {
      if (errno == EINTR)
      {
        nReadCount = 0; // call read again
      }
      else
      {
        return -1;
      }
    }
    else if (0 == nReadCount) // met EOF
    {
      break;
    }
    nLeft -= nReadCount;
    ptr += nReadCount;
  }

  return n - nLeft;//this is because EOF exists
}

ssize_t //Write n bytes to a fd
writen(int fd, const void *vptr, size_t n)
{
  size_t nLeft = n;
  ssize_t nWriteCount = 0;
  const char *ptr = (const char *)vptr;//Note, this is const

  while (nLeft > 0)
  {
    if ((nWriteCount = write(fd, ptr, nLeft)) <= 0) // note, <=0, not < 0
    {
      // errors occur
      if (nWriteCount < 0 && errno == EINTR)
      {
        nWriteCount = 0;
      }
      else
      {
        return -1;
      }
    }
    nLeft -= nWriteCount;
    ptr += nWriteCount;
  }

  return n;
}

//Fucking slow version of readline
ssize_t
fucking_slow_readline(int fd, void *vptr, size_t maxlen)
{
  char c;
  ssize_t nReadCount;
  char *ptr = (char*)vptr;
  ssize_t i = 1;
  for (; i < maxlen; i++)
  {
  again:
    if ((nReadCount = read(fd, &c, 1)) == 1)//read successfully
    {
      *ptr++ = c;
      if (c == '\n')
      {
        break;
      }
    }
    else if (0 == nReadCount) // encounter EOF, i-1 bytes were read
    {
      *ptr = 0;
      return i - 1;
    }
    else//read error,check EINTR
    {
      if (errno == EINTR) { goto again; }
      return -1;
    }
  }

  *ptr = 0;//null terminate
  return i;
}


// Better version of readline
#define MAX_BUFF 1024
static int readCount;
static char * readPtr = NULL;// pointer to valid data buffer
static char readBuf[MAX_BUFF];// stores the buffer

static ssize_t
my_read(int fd, char *ptr)//ptr is a pointer to a char
{
  if (readCount <= 0)// need to fetch data again
  {
    again:
    if ((readCount = read(fd, readBuf, sizeof(readBuf))) < 0)// error occurs
    {
      if (errno == EINTR)
        goto again;
      return -1;
    }
    else if(0 == readCount)// EOF
    {
      return 0;
    }

    readPtr = readBuf;
  }

  readCount--;
  *ptr = *readPtr++;
  return 1;
}

ssize_t
readline(int fd, void *vptr, size_t maxlen)
{
  ssize_t nReadCount = 0;
  char *ptr = (char*)vptr;
  char c;
  ssize_t i = 1;
  for (; i < maxlen; i++)
  {
    if ((nReadCount = my_read(fd, &c)) == 1)
    {
      *ptr++ = c;
      if (c == '\n') { break; }
    }
    else if(0 == nReadCount)//met EOF
    {
      *ptr = 0;
      return i - 1;
    }
    else
    {
      return -1;
    }
  }

  *ptr = 0;
  return i;
}

ssize_t
readlinebuf(void **vptrptr)
{
  if (readCount)
  {
    *vptrptr = readPtr;
  }

  return readCount;
}
