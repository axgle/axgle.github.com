<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/zsysnum_darwin_386.go</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

</head>

<body>

  <script>
    // Catch 'enter' key down events and trigger the search form submission.
    function codesearchKeyDown(event) {
      if (event.which == 13) {
        var form = document.getElementById('codesearch');
        var query = document.getElementById('codesearchQuery');
        form.q.value = "lang:go package:go.googlecode.com " + query.value;
        document.getElementById('codesearch').submit();
}      return true;
}
    // Capture the submission event and construct the query parameter.
    function codeSearchSubmit() {
      var query = document.getElementById('codesearchQuery');
      var form = document.getElementById('codesearch');
      form.q.value = "lang:go package:go.googlecode.com " + query.value;
      return true;
}  </script>

<div id="topnav">
  <table summary="">
    <tr>
      <td id="headerImage">
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
      </td>
      <td>
        <div id="headerDocSetTitle">The Go Programming Language</div>
      </td>
      <td>
        <!-- <table>
          <tr>
            <td>
              <! The input box is outside of the form because we want to add
              a couple of restricts to the query before submitting. If we just
              add the restricts to the text box before submitting, then they
              appear in the box when the user presses 'back'. Thus we use a
              hidden field in the form. However, there's no way to stop the
              non-hidden text box from also submitting a value unless we move
              it outside of the form
              <input type="search" id="codesearchQuery" value="" size="30" onkeydown="return codesearchKeyDown(event);"/>
              <form method="GET" action="http://www.google.com/codesearch" id="codesearch" class="search" onsubmit="return codeSearchSubmit();" style="display:inline;">
                <input type="hidden" name="q" value=""/>
                <input type="submit" value="Code search" />
                <span style="color: red">(TODO: remove for now?)</span>
              </form>
            </td>
          </tr>
          <tr>
            <td>
              <span style="color: gray;">(e.g. &ldquo;pem&rdquo; or &ldquo;xml&rdquo;)</span>
            </td>
          </tr>
        </table> -->
      </td>
    </tr>
  </table>
</div>

<div id="linkList">
  <ul>
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Help</li>
    <li>#go-nuts on irc.freenode.net</li>
    <li><a href="http://groups.google.com/group/golang-nuts">Go Nuts mailing list</a></li>
    <li><a href="http://code.google.com/p/go/issues/list">Issue tracker</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Go code search</li>
    <form method="GET" action="http://golang.org/search" class="search">
    <input type="search" name="q" value="" size="25" style="width:80%; max-width:200px" />
    <input type="submit" value="Go" />
    </form>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Last update</li>
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/syscall/zsysnum_darwin_386.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// mksysnum_darwin.sh /home/rsc/pub/xnu-1228/bsd/kern/syscalls.master</span>
<a id="L2"></a><span class="comment">// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT</span>

<a id="L4"></a>package syscall

<a id="L6"></a>const (
    <a id="L7"></a><span class="comment">// SYS_NOSYS = 0;  // { int nosys(void); }   { indirect syscall }</span>
    <a id="L8"></a>SYS_EXIT  = 1; <span class="comment">// { void exit(int rval); }</span>
    <a id="L9"></a>SYS_FORK  = 2; <span class="comment">// { int fork(void); }</span>
    <a id="L10"></a>SYS_READ  = 3; <span class="comment">// { user_ssize_t read(int fd, user_addr_t cbuf, user_size_t nbyte); }</span>
    <a id="L11"></a>SYS_WRITE = 4; <span class="comment">// { user_ssize_t write(int fd, user_addr_t cbuf, user_size_t nbyte); }</span>
    <a id="L12"></a>SYS_OPEN  = 5; <span class="comment">// { int open(user_addr_t path, int flags, int mode); }</span>
    <a id="L13"></a>SYS_CLOSE = 6; <span class="comment">// { int close(int fd); }</span>
    <a id="L14"></a>SYS_WAIT4 = 7; <span class="comment">// { int wait4(int pid, user_addr_t status, int options, user_addr_t rusage); }</span>
    <a id="L15"></a><span class="comment">// SYS_NOSYS = 8;  // { int nosys(void); }   { old creat }</span>
    <a id="L16"></a>SYS_LINK   = 9;  <span class="comment">// { int link(user_addr_t path, user_addr_t link); }</span>
    <a id="L17"></a>SYS_UNLINK = 10; <span class="comment">// { int unlink(user_addr_t path); }</span>
    <a id="L18"></a><span class="comment">// SYS_NOSYS = 11;  // { int nosys(void); }   { old execv }</span>
    <a id="L19"></a>SYS_CHDIR      = 12; <span class="comment">// { int chdir(user_addr_t path); }</span>
    <a id="L20"></a>SYS_FCHDIR     = 13; <span class="comment">// { int fchdir(int fd); }</span>
    <a id="L21"></a>SYS_MKNOD      = 14; <span class="comment">// { int mknod(user_addr_t path, int mode, int dev); }</span>
    <a id="L22"></a>SYS_CHMOD      = 15; <span class="comment">// { int chmod(user_addr_t path, int mode); }</span>
    <a id="L23"></a>SYS_CHOWN      = 16; <span class="comment">// { int chown(user_addr_t path, int uid, int gid); }</span>
    <a id="L24"></a>SYS_OGETFSSTAT = 18; <span class="comment">// { int ogetfsstat(user_addr_t buf, int bufsize, int flags); }</span>
    <a id="L25"></a>SYS_GETFSSTAT  = 18; <span class="comment">// { int getfsstat(user_addr_t buf, int bufsize, int flags); }</span>
    <a id="L26"></a><span class="comment">// SYS_NOSYS = 19;  // { int nosys(void); }   { old lseek }</span>
    <a id="L27"></a>SYS_GETPID = 20; <span class="comment">// { int getpid(void); }</span>
    <a id="L28"></a><span class="comment">// SYS_NOSYS = 21;  // { int nosys(void); }   { old mount }</span>
    <a id="L29"></a><span class="comment">// SYS_NOSYS = 22;  // { int nosys(void); }   { old umount }</span>
    <a id="L30"></a>SYS_SETUID      = 23; <span class="comment">// { int setuid(uid_t uid); }</span>
    <a id="L31"></a>SYS_GETUID      = 24; <span class="comment">// { int getuid(void); }</span>
    <a id="L32"></a>SYS_GETEUID     = 25; <span class="comment">// { int geteuid(void); }</span>
    <a id="L33"></a>SYS_PTRACE      = 26; <span class="comment">// { int ptrace(int req, pid_t pid, caddr_t addr, int data); }</span>
    <a id="L34"></a>SYS_RECVMSG     = 27; <span class="comment">// { int recvmsg(int s, struct msghdr *msg, int flags); }</span>
    <a id="L35"></a>SYS_SENDMSG     = 28; <span class="comment">// { int sendmsg(int s, caddr_t msg, int flags); }</span>
    <a id="L36"></a>SYS_RECVFROM    = 29; <span class="comment">// { int recvfrom(int s, void *buf, size_t len, int flags, struct sockaddr *from, int *fromlenaddr); }</span>
    <a id="L37"></a>SYS_ACCEPT      = 30; <span class="comment">// { int accept(int s, caddr_t name, socklen_t	*anamelen); }</span>
    <a id="L38"></a>SYS_GETPEERNAME = 31; <span class="comment">// { int getpeername(int fdes, caddr_t asa, socklen_t *alen); }</span>
    <a id="L39"></a>SYS_GETSOCKNAME = 32; <span class="comment">// { int getsockname(int fdes, caddr_t asa, socklen_t *alen); }</span>
    <a id="L40"></a><span class="comment">// SYS_NOSYS = 27;  // { int nosys(void); }</span>
    <a id="L41"></a><span class="comment">// SYS_NOSYS = 28;  // { int nosys(void); }</span>
    <a id="L42"></a><span class="comment">// SYS_NOSYS = 29;  // { int nosys(void); }</span>
    <a id="L43"></a><span class="comment">// SYS_NOSYS = 30;  // { int nosys(void); }</span>
    <a id="L44"></a><span class="comment">// SYS_NOSYS = 31;  // { int nosys(void); }</span>
    <a id="L45"></a><span class="comment">// SYS_NOSYS = 32;  // { int nosys(void); }</span>
    <a id="L46"></a>SYS_ACCESS   = 33; <span class="comment">// { int access(user_addr_t path, int flags); }</span>
    <a id="L47"></a>SYS_CHFLAGS  = 34; <span class="comment">// { int chflags(char *path, int flags); }</span>
    <a id="L48"></a>SYS_FCHFLAGS = 35; <span class="comment">// { int fchflags(int fd, int flags); }</span>
    <a id="L49"></a>SYS_SYNC     = 36; <span class="comment">// { int sync(void); }</span>
    <a id="L50"></a>SYS_KILL     = 37; <span class="comment">// { int kill(int pid, int signum, int posix); }</span>
    <a id="L51"></a><span class="comment">// SYS_NOSYS = 38;  // { int nosys(void); }   { old stat  }</span>
    <a id="L52"></a>SYS_GETPPID = 39; <span class="comment">// { int getppid(void); }</span>
    <a id="L53"></a><span class="comment">// SYS_NOSYS = 40;  // { int nosys(void); }   { old lstat }</span>
    <a id="L54"></a>SYS_DUP     = 41; <span class="comment">// { int dup(u_int fd); }</span>
    <a id="L55"></a>SYS_PIPE    = 42; <span class="comment">// { int pipe(void); }</span>
    <a id="L56"></a>SYS_GETEGID = 43; <span class="comment">// { int getegid(void); }</span>
    <a id="L57"></a>SYS_PROFIL  = 44; <span class="comment">// { int profil(short *bufbase, size_t bufsize, u_long pcoffset, u_int pcscale); }</span>
    <a id="L58"></a><span class="comment">// SYS_NOSYS = 45;  // { int nosys(void); } { old ktrace }</span>
    <a id="L59"></a>SYS_SIGACTION   = 46; <span class="comment">// { int sigaction(int signum, struct __sigaction *nsa, struct sigaction *osa); }</span>
    <a id="L60"></a>SYS_GETGID      = 47; <span class="comment">// { int getgid(void); }</span>
    <a id="L61"></a>SYS_SIGPROCMASK = 48; <span class="comment">// { int sigprocmask(int how, user_addr_t mask, user_addr_t omask); }</span>
    <a id="L62"></a>SYS_GETLOGIN    = 49; <span class="comment">// { int getlogin(char *namebuf, u_int namelen); }</span>
    <a id="L63"></a>SYS_SETLOGIN    = 50; <span class="comment">// { int setlogin(char *namebuf); }</span>
    <a id="L64"></a>SYS_ACCT        = 51; <span class="comment">// { int acct(char *path); }</span>
    <a id="L65"></a>SYS_SIGPENDING  = 52; <span class="comment">// { int sigpending(struct sigvec *osv); }</span>
    <a id="L66"></a>SYS_SIGALTSTACK = 53; <span class="comment">// { int sigaltstack(struct sigaltstack *nss, struct sigaltstack *oss); }</span>
    <a id="L67"></a>SYS_IOCTL       = 54; <span class="comment">// { int ioctl(int fd, u_long com, caddr_t data); }</span>
    <a id="L68"></a>SYS_REBOOT      = 55; <span class="comment">// { int reboot(int opt, char *command); }</span>
    <a id="L69"></a>SYS_REVOKE      = 56; <span class="comment">// { int revoke(char *path); }</span>
    <a id="L70"></a>SYS_SYMLINK     = 57; <span class="comment">// { int symlink(char *path, char *link); }</span>
    <a id="L71"></a>SYS_READLINK    = 58; <span class="comment">// { int readlink(char *path, char *buf, int count); }</span>
    <a id="L72"></a>SYS_EXECVE      = 59; <span class="comment">// { int execve(char *fname, char **argp, char **envp); }</span>
    <a id="L73"></a>SYS_UMASK       = 60; <span class="comment">// { int umask(int newmask); }</span>
    <a id="L74"></a>SYS_CHROOT      = 61; <span class="comment">// { int chroot(user_addr_t path); }</span>
    <a id="L75"></a><span class="comment">// SYS_NOSYS = 62;  // { int nosys(void); }   { old fstat }</span>
    <a id="L76"></a><span class="comment">// SYS_NOSYS = 63;  // { int nosys(void); }   { used internally, reserved }</span>
    <a id="L77"></a><span class="comment">// SYS_NOSYS = 64;  // { int nosys(void); }   { old getpagesize }</span>
    <a id="L78"></a>SYS_MSYNC = 65; <span class="comment">// { int msync(caddr_t addr, size_t len, int flags); }</span>
    <a id="L79"></a>SYS_VFORK = 66; <span class="comment">// { int vfork(void); }</span>
    <a id="L80"></a><span class="comment">// SYS_NOSYS = 67;  // { int nosys(void); }   { old vread }</span>
    <a id="L81"></a><span class="comment">// SYS_NOSYS = 68;  // { int nosys(void); }   { old vwrite }</span>
    <a id="L82"></a>SYS_SBRK = 69; <span class="comment">// { int sbrk(int incr) NO_SYSCALL_STUB; }</span>
    <a id="L83"></a>SYS_SSTK = 70; <span class="comment">// { int sstk(int incr) NO_SYSCALL_STUB; }</span>
    <a id="L84"></a><span class="comment">// SYS_NOSYS = 71;  // { int nosys(void); }   { old mmap }</span>
    <a id="L85"></a>SYS_OVADVISE = 72; <span class="comment">// { int ovadvise(void) NO_SYSCALL_STUB; }   { old vadvise }</span>
    <a id="L86"></a>SYS_MUNMAP   = 73; <span class="comment">// { int munmap(caddr_t addr, size_t len); }</span>
    <a id="L87"></a>SYS_MPROTECT = 74; <span class="comment">// { int mprotect(caddr_t addr, size_t len, int prot); }</span>
    <a id="L88"></a>SYS_MADVISE  = 75; <span class="comment">// { int madvise(caddr_t addr, size_t len, int behav); }</span>
    <a id="L89"></a><span class="comment">// SYS_NOSYS = 76;  // { int nosys(void); }   { old vhangup }</span>
    <a id="L90"></a><span class="comment">// SYS_NOSYS = 77;  // { int nosys(void); }   { old vlimit }</span>
    <a id="L91"></a>SYS_MINCORE   = 78; <span class="comment">// { int mincore(user_addr_t addr, user_size_t len, user_addr_t vec); }</span>
    <a id="L92"></a>SYS_GETGROUPS = 79; <span class="comment">// { int getgroups(u_int gidsetsize, gid_t *gidset); }</span>
    <a id="L93"></a>SYS_SETGROUPS = 80; <span class="comment">// { int setgroups(u_int gidsetsize, gid_t *gidset); }</span>
    <a id="L94"></a>SYS_GETPGRP   = 81; <span class="comment">// { int getpgrp(void); }</span>
    <a id="L95"></a>SYS_SETPGID   = 82; <span class="comment">// { int setpgid(int pid, int pgid); }</span>
    <a id="L96"></a>SYS_SETITIMER = 83; <span class="comment">// { int setitimer(u_int which, struct itimerval *itv, struct itimerval *oitv); }</span>
    <a id="L97"></a><span class="comment">// SYS_NOSYS = 84;  // { int nosys(void); }   { old wait }</span>
    <a id="L98"></a>SYS_SWAPON    = 85; <span class="comment">// { int swapon(void); }</span>
    <a id="L99"></a>SYS_GETITIMER = 86; <span class="comment">// { int getitimer(u_int which, struct itimerval *itv); }</span>
    <a id="L100"></a><span class="comment">// SYS_NOSYS = 87;  // { int nosys(void); }   { old gethostname }</span>
    <a id="L101"></a><span class="comment">// SYS_NOSYS = 88;  // { int nosys(void); }   { old sethostname }</span>
    <a id="L102"></a>SYS_GETDTABLESIZE = 89; <span class="comment">// { int getdtablesize(void); }</span>
    <a id="L103"></a>SYS_DUP2          = 90; <span class="comment">// { int dup2(u_int from, u_int to); }</span>
    <a id="L104"></a><span class="comment">// SYS_NOSYS = 91;  // { int nosys(void); }   { old getdopt }</span>
    <a id="L105"></a>SYS_FCNTL  = 92; <span class="comment">// { int fcntl(int fd, int cmd, long arg); }</span>
    <a id="L106"></a>SYS_SELECT = 93; <span class="comment">// { int select(int nd, u_int32_t *in, u_int32_t *ou, u_int32_t *ex, struct timeval *tv); }</span>
    <a id="L107"></a><span class="comment">// SYS_NOSYS = 94;  // { int nosys(void); }   { old setdopt }</span>
    <a id="L108"></a>SYS_FSYNC       = 95; <span class="comment">// { int fsync(int fd); }</span>
    <a id="L109"></a>SYS_SETPRIORITY = 96; <span class="comment">// { int setpriority(int which, id_t who, int prio); }</span>
    <a id="L110"></a>SYS_SOCKET      = 97; <span class="comment">// { int socket(int domain, int type, int protocol); }</span>
    <a id="L111"></a>SYS_CONNECT     = 98; <span class="comment">// { int connect(int s, caddr_t name, socklen_t namelen); }</span>
    <a id="L112"></a><span class="comment">// SYS_NOSYS = 97;  // { int nosys(void); }</span>
    <a id="L113"></a><span class="comment">// SYS_NOSYS = 98;  // { int nosys(void); }</span>
    <a id="L114"></a><span class="comment">// SYS_NOSYS = 99;  // { int nosys(void); }   { old accept }</span>
    <a id="L115"></a>SYS_GETPRIORITY = 100; <span class="comment">// { int getpriority(int which, id_t who); }</span>
    <a id="L116"></a><span class="comment">// SYS_NOSYS = 101;  // { int nosys(void); }   { old send }</span>
    <a id="L117"></a><span class="comment">// SYS_NOSYS = 102;  // { int nosys(void); }   { old recv }</span>
    <a id="L118"></a><span class="comment">// SYS_NOSYS = 103;  // { int nosys(void); }   { old sigreturn }</span>
    <a id="L119"></a>SYS_BIND       = 104; <span class="comment">// { int bind(int s, caddr_t name, socklen_t namelen); }</span>
    <a id="L120"></a>SYS_SETSOCKOPT = 105; <span class="comment">// { int setsockopt(int s, int level, int name, caddr_t val, socklen_t valsize); }</span>
    <a id="L121"></a>SYS_LISTEN     = 106; <span class="comment">// { int listen(int s, int backlog); }</span>
    <a id="L122"></a><span class="comment">// SYS_NOSYS = 104;  // { int nosys(void); }</span>
    <a id="L123"></a><span class="comment">// SYS_NOSYS = 105;  // { int nosys(void); }</span>
    <a id="L124"></a><span class="comment">// SYS_NOSYS = 106;  // { int nosys(void); }</span>
    <a id="L125"></a><span class="comment">// SYS_NOSYS = 107;  // { int nosys(void); }   { old vtimes }</span>
    <a id="L126"></a><span class="comment">// SYS_NOSYS = 108;  // { int nosys(void); }   { old sigvec }</span>
    <a id="L127"></a><span class="comment">// SYS_NOSYS = 109;  // { int nosys(void); }   { old sigblock }</span>
    <a id="L128"></a><span class="comment">// SYS_NOSYS = 110;  // { int nosys(void); }   { old sigsetmask }</span>
    <a id="L129"></a>SYS_SIGSUSPEND = 111; <span class="comment">// { int sigsuspend(sigset_t mask); }</span>
    <a id="L130"></a><span class="comment">// SYS_NOSYS = 112;  // { int nosys(void); }   { old sigstack }</span>
    <a id="L131"></a><span class="comment">// SYS_NOSYS = 113;  // { int nosys(void); }   { old recvmsg }</span>
    <a id="L132"></a><span class="comment">// SYS_NOSYS = 114;  // { int nosys(void); }   { old sendmsg }</span>
    <a id="L133"></a><span class="comment">// SYS_NOSYS = 113;  // { int nosys(void); }</span>
    <a id="L134"></a><span class="comment">// SYS_NOSYS = 114;  // { int nosys(void); }</span>
    <a id="L135"></a><span class="comment">// SYS_NOSYS = 115;  // { int nosys(void); }   { old vtrace }</span>
    <a id="L136"></a>SYS_GETTIMEOFDAY = 116; <span class="comment">// { int gettimeofday(struct timeval *tp, struct timezone *tzp); }</span>
    <a id="L137"></a>SYS_GETRUSAGE    = 117; <span class="comment">// { int getrusage(int who, struct rusage *rusage); }</span>
    <a id="L138"></a>SYS_GETSOCKOPT   = 118; <span class="comment">// { int getsockopt(int s, int level, int name, caddr_t val, socklen_t *avalsize); }</span>
    <a id="L139"></a><span class="comment">// SYS_NOSYS = 118;  // { int nosys(void); }</span>
    <a id="L140"></a><span class="comment">// SYS_NOSYS = 119;  // { int nosys(void); }   { old resuba }</span>
    <a id="L141"></a>SYS_READV        = 120; <span class="comment">// { user_ssize_t readv(int fd, struct iovec *iovp, u_int iovcnt); }</span>
    <a id="L142"></a>SYS_WRITEV       = 121; <span class="comment">// { user_ssize_t writev(int fd, struct iovec *iovp, u_int iovcnt); }</span>
    <a id="L143"></a>SYS_SETTIMEOFDAY = 122; <span class="comment">// { int settimeofday(struct timeval *tv, struct timezone *tzp); }</span>
    <a id="L144"></a>SYS_FCHOWN       = 123; <span class="comment">// { int fchown(int fd, int uid, int gid); }</span>
    <a id="L145"></a>SYS_FCHMOD       = 124; <span class="comment">// { int fchmod(int fd, int mode); }</span>
    <a id="L146"></a><span class="comment">// SYS_NOSYS = 125;  // { int nosys(void); }   { old recvfrom }</span>
    <a id="L147"></a>SYS_SETREUID = 126; <span class="comment">// { int setreuid(uid_t ruid, uid_t euid); }</span>
    <a id="L148"></a>SYS_SETREGID = 127; <span class="comment">// { int setregid(gid_t rgid, gid_t egid); }</span>
    <a id="L149"></a>SYS_RENAME   = 128; <span class="comment">// { int rename(char *from, char *to); }</span>
    <a id="L150"></a><span class="comment">// SYS_NOSYS = 129;  // { int nosys(void); }   { old truncate }</span>
    <a id="L151"></a><span class="comment">// SYS_NOSYS = 130;  // { int nosys(void); }   { old ftruncate }</span>
    <a id="L152"></a>SYS_FLOCK      = 131; <span class="comment">// { int flock(int fd, int how); }</span>
    <a id="L153"></a>SYS_MKFIFO     = 132; <span class="comment">// { int mkfifo(user_addr_t path, int mode); }</span>
    <a id="L154"></a>SYS_SENDTO     = 133; <span class="comment">// { int sendto(int s, caddr_t buf, size_t len, int flags, caddr_t to, socklen_t tolen); }</span>
    <a id="L155"></a>SYS_SHUTDOWN   = 134; <span class="comment">// { int shutdown(int s, int how); }</span>
    <a id="L156"></a>SYS_SOCKETPAIR = 135; <span class="comment">// { int socketpair(int domain, int type, int protocol, int *rsv); }</span>
    <a id="L157"></a><span class="comment">// SYS_NOSYS = 133;  // { int nosys(void); }</span>
    <a id="L158"></a><span class="comment">// SYS_NOSYS = 134;  // { int nosys(void); }</span>
    <a id="L159"></a><span class="comment">// SYS_NOSYS = 135;  // { int nosys(void); }</span>
    <a id="L160"></a>SYS_MKDIR   = 136; <span class="comment">// { int mkdir(user_addr_t path, int mode); }</span>
    <a id="L161"></a>SYS_RMDIR   = 137; <span class="comment">// { int rmdir(char *path); }</span>
    <a id="L162"></a>SYS_UTIMES  = 138; <span class="comment">// { int utimes(char *path, struct timeval *tptr); }</span>
    <a id="L163"></a>SYS_FUTIMES = 139; <span class="comment">// { int futimes(int fd, struct timeval *tptr); }</span>
    <a id="L164"></a>SYS_ADJTIME = 140; <span class="comment">// { int adjtime(struct timeval *delta, struct timeval *olddelta); }</span>
    <a id="L165"></a><span class="comment">// SYS_NOSYS = 141;  // { int nosys(void); }   { old getpeername }</span>
    <a id="L166"></a>SYS_GETHOSTUUID = 142; <span class="comment">// { int gethostuuid(unsigned char *uuid_buf, const struct timespec *timeoutp); }</span>
    <a id="L167"></a><span class="comment">// SYS_NOSYS = 143;  // { int nosys(void); }   { old sethostid 	}</span>
    <a id="L168"></a><span class="comment">// SYS_NOSYS = 144;  // { int nosys(void); }   { old getrlimit }</span>
    <a id="L169"></a><span class="comment">// SYS_NOSYS = 145;  // { int nosys(void); }   { old setrlimit }</span>
    <a id="L170"></a><span class="comment">// SYS_NOSYS = 146;  // { int nosys(void); }   { old killpg }</span>
    <a id="L171"></a>SYS_SETSID = 147; <span class="comment">// { int setsid(void); }</span>
    <a id="L172"></a><span class="comment">// SYS_NOSYS = 148;  // { int nosys(void); }   { old setquota }</span>
    <a id="L173"></a><span class="comment">// SYS_NOSYS = 149;  // { int nosys(void); }   { old qquota }</span>
    <a id="L174"></a><span class="comment">// SYS_NOSYS = 150;  // { int nosys(void); }   { old getsockname }</span>
    <a id="L175"></a>SYS_GETPGID     = 151; <span class="comment">// { int getpgid(pid_t pid); }</span>
    <a id="L176"></a>SYS_SETPRIVEXEC = 152; <span class="comment">// { int setprivexec(int flag); }</span>
    <a id="L177"></a>SYS_PREAD       = 153; <span class="comment">// { user_ssize_t pread(int fd, user_addr_t buf, user_size_t nbyte, off_t offset); }</span>
    <a id="L178"></a>SYS_PWRITE      = 154; <span class="comment">// { user_ssize_t pwrite(int fd, user_addr_t buf, user_size_t nbyte, off_t offset); }</span>
    <a id="L179"></a>SYS_NFSSVC      = 155; <span class="comment">// { int nfssvc(int flag, caddr_t argp); }</span>
    <a id="L180"></a><span class="comment">// SYS_NOSYS = 155;  // { int nosys(void); }</span>
    <a id="L181"></a><span class="comment">// SYS_NOSYS = 156;  // { int nosys(void); }   { old getdirentries }</span>
    <a id="L182"></a>SYS_STATFS  = 157; <span class="comment">// { int statfs(char *path, struct statfs *buf); }</span>
    <a id="L183"></a>SYS_FSTATFS = 158; <span class="comment">// { int fstatfs(int fd, struct statfs *buf); }</span>
    <a id="L184"></a>SYS_UNMOUNT = 159; <span class="comment">// { int unmount(user_addr_t path, int flags); }</span>
    <a id="L185"></a><span class="comment">// SYS_NOSYS = 160;  // { int nosys(void); }   { old async_daemon }</span>
    <a id="L186"></a>SYS_GETFH = 161; <span class="comment">// { int getfh(char *fname, fhandle_t *fhp); }</span>
    <a id="L187"></a><span class="comment">// SYS_NOSYS = 161;  // { int nosys(void); }</span>
    <a id="L188"></a><span class="comment">// SYS_NOSYS = 162;  // { int nosys(void); }   { old getdomainname }</span>
    <a id="L189"></a><span class="comment">// SYS_NOSYS = 163;  // { int nosys(void); }   { old setdomainname }</span>
    <a id="L190"></a><span class="comment">// SYS_NOSYS = 164;  // { int nosys(void); }</span>
    <a id="L191"></a>SYS_QUOTACTL = 165; <span class="comment">// { int quotactl(const char *path, int cmd, int uid, caddr_t arg); }</span>
    <a id="L192"></a><span class="comment">// SYS_NOSYS = 166;  // { int nosys(void); }   { old exportfs }</span>
    <a id="L193"></a>SYS_MOUNT = 167; <span class="comment">// { int mount(char *type, char *path, int flags, caddr_t data); }</span>
    <a id="L194"></a><span class="comment">// SYS_NOSYS = 168;  // { int nosys(void); }   { old ustat }</span>
    <a id="L195"></a>SYS_CSOPS = 169; <span class="comment">// { int csops(pid_t pid, uint32_t ops, user_addr_t useraddr, user_size_t usersize); }</span>
    <a id="L196"></a><span class="comment">// SYS_NOSYS = 171;  // { int nosys(void); }   { old wait3 }</span>
    <a id="L197"></a><span class="comment">// SYS_NOSYS = 172;  // { int nosys(void); }   { old rpause	}</span>
    <a id="L198"></a>SYS_WAITID = 173; <span class="comment">// { int waitid(idtype_t idtype, id_t id, siginfo_t *infop, int options); }</span>
    <a id="L199"></a><span class="comment">// SYS_NOSYS = 174;  // { int nosys(void); }   { old getdents }</span>
    <a id="L200"></a><span class="comment">// SYS_NOSYS = 175;  // { int nosys(void); }   { old gc_control }</span>
    <a id="L201"></a>SYS_ADD_PROFIL = 176; <span class="comment">// { int add_profil(short *bufbase, size_t bufsize, u_long pcoffset, u_int pcscale); }</span>
    <a id="L202"></a><span class="comment">// SYS_NOSYS = 177;  // { int nosys(void); }</span>
    <a id="L203"></a><span class="comment">// SYS_NOSYS = 178;  // { int nosys(void); }</span>
    <a id="L204"></a><span class="comment">// SYS_NOSYS = 179;  // { int nosys(void); }</span>
    <a id="L205"></a>SYS_KDEBUG_TRACE = 180; <span class="comment">// { int kdebug_trace(int code, int arg1, int arg2, int arg3, int arg4, int arg5) NO_SYSCALL_STUB; }</span>
    <a id="L206"></a>SYS_SETGID       = 181; <span class="comment">// { int setgid(gid_t gid); }</span>
    <a id="L207"></a>SYS_SETEGID      = 182; <span class="comment">// { int setegid(gid_t egid); }</span>
    <a id="L208"></a>SYS_SETEUID      = 183; <span class="comment">// { int seteuid(uid_t euid); }</span>
    <a id="L209"></a>SYS_SIGRETURN    = 184; <span class="comment">// { int sigreturn(struct ucontext *uctx, int infostyle); }</span>
    <a id="L210"></a><span class="comment">// SYS_NOSYS = 186;  // { int nosys(void); }</span>
    <a id="L211"></a><span class="comment">// SYS_NOSYS = 187;  // { int nosys(void); }</span>
    <a id="L212"></a>SYS_STAT      = 188; <span class="comment">// { int stat(user_addr_t path, user_addr_t ub); }</span>
    <a id="L213"></a>SYS_FSTAT     = 189; <span class="comment">// { int fstat(int fd, user_addr_t ub); }</span>
    <a id="L214"></a>SYS_LSTAT     = 190; <span class="comment">// { int lstat(user_addr_t path, user_addr_t ub); }</span>
    <a id="L215"></a>SYS_PATHCONF  = 191; <span class="comment">// { int pathconf(char *path, int name); }</span>
    <a id="L216"></a>SYS_FPATHCONF = 192; <span class="comment">// { int fpathconf(int fd, int name); }</span>
    <a id="L217"></a><span class="comment">// SYS_NOSYS = 193;  // { int nosys(void); }</span>
    <a id="L218"></a>SYS_GETRLIMIT     = 194; <span class="comment">// { int getrlimit(u_int which, struct rlimit *rlp); }</span>
    <a id="L219"></a>SYS_SETRLIMIT     = 195; <span class="comment">// { int setrlimit(u_int which, struct rlimit *rlp); }</span>
    <a id="L220"></a>SYS_GETDIRENTRIES = 196; <span class="comment">// { int getdirentries(int fd, char *buf, u_int count, long *basep); }</span>
    <a id="L221"></a>SYS_MMAP          = 197; <span class="comment">// { user_addr_t mmap(caddr_t addr, size_t len, int prot, int flags, int fd, off_t pos); }</span>
    <a id="L222"></a><span class="comment">// SYS_NOSYS = 198;  // { int nosys(void); } 	{ __syscall }</span>
    <a id="L223"></a>SYS_LSEEK     = 199; <span class="comment">// { off_t lseek(int fd, off_t offset, int whence); }</span>
    <a id="L224"></a>SYS_TRUNCATE  = 200; <span class="comment">// { int truncate(char *path, off_t length); }</span>
    <a id="L225"></a>SYS_FTRUNCATE = 201; <span class="comment">// { int ftruncate(int fd, off_t length); }</span>
    <a id="L226"></a>SYS___SYSCTL  = 202; <span class="comment">// { int __sysctl(int *name, u_int namelen, void *old, size_t *oldlenp, void *new, size_t newlen); }</span>
    <a id="L227"></a>SYS_MLOCK     = 203; <span class="comment">// { int mlock(caddr_t addr, size_t len); }</span>
    <a id="L228"></a>SYS_MUNLOCK   = 204; <span class="comment">// { int munlock(caddr_t addr, size_t len); }</span>
    <a id="L229"></a>SYS_UNDELETE  = 205; <span class="comment">// { int undelete(user_addr_t path); }</span>
    <a id="L230"></a>SYS_ATSOCKET  = 206; <span class="comment">// { int ATsocket(int proto); }</span>
    <a id="L231"></a><span class="comment">// SYS_NOSYS = 213;  // { int nosys(void); } 	{ Reserved for AppleTalk }</span>
    <a id="L232"></a><span class="comment">// SYS_NOSYS = 206;  // { int nosys(void); }</span>
    <a id="L233"></a><span class="comment">// SYS_NOSYS = 207;  // { int nosys(void); }</span>
    <a id="L234"></a><span class="comment">// SYS_NOSYS = 208;  // { int nosys(void); }</span>
    <a id="L235"></a><span class="comment">// SYS_NOSYS = 209;  // { int nosys(void); }</span>
    <a id="L236"></a><span class="comment">// SYS_NOSYS = 210;  // { int nosys(void); }</span>
    <a id="L237"></a><span class="comment">// SYS_NOSYS = 211;  // { int nosys(void); }</span>
    <a id="L238"></a><span class="comment">// SYS_NOSYS = 212;  // { int nosys(void); }</span>
    <a id="L239"></a><span class="comment">// SYS_NOSYS = 213;  // { int nosys(void); } 	{ Reserved for AppleTalk }</span>
    <a id="L240"></a>SYS_KQUEUE_FROM_PORTSET_NP = 214; <span class="comment">// { int kqueue_from_portset_np(int portset); }</span>
    <a id="L241"></a>SYS_KQUEUE_PORTSET_NP      = 215; <span class="comment">// { int kqueue_portset_np(int fd); }</span>
    <a id="L242"></a>SYS_GETATTRLIST            = 220; <span class="comment">// { int getattrlist(const char *path, struct attrlist *alist, void *attributeBuffer, size_t bufferSize, u_long options); }</span>
    <a id="L243"></a>SYS_SETATTRLIST            = 221; <span class="comment">// { int setattrlist(const char *path, struct attrlist *alist, void *attributeBuffer, size_t bufferSize, u_long options); }</span>
    <a id="L244"></a>SYS_GETDIRENTRIESATTR      = 222; <span class="comment">// { int getdirentriesattr(int fd, struct attrlist *alist, void *buffer, size_t buffersize, u_long *count, u_long *basep, u_long *newstate, u_long options); }</span>
    <a id="L245"></a>SYS_EXCHANGEDATA           = 223; <span class="comment">// { int exchangedata(const char *path1, const char *path2, u_long options); }</span>
    <a id="L246"></a><span class="comment">// SYS_NOSYS = 224;  // { int nosys(void); } { was checkuseraccess }</span>
    <a id="L247"></a>SYS_SEARCHFS = 225; <span class="comment">// { int searchfs(const char *path, struct fssearchblock *searchblock, u_long *nummatches, u_long scriptcode, u_long options, struct searchstate *state); }</span>
    <a id="L248"></a>SYS_DELETE   = 226; <span class="comment">// { int delete(user_addr_t path) NO_SYSCALL_STUB; }       { private delete (Carbon semantics) }</span>
    <a id="L249"></a>SYS_COPYFILE = 227; <span class="comment">// { int copyfile(char *from, char *to, int mode, int flags) NO_SYSCALL_STUB; }</span>
    <a id="L250"></a><span class="comment">// SYS_NOSYS = 228;  // { int nosys(void); }</span>
    <a id="L251"></a><span class="comment">// SYS_NOSYS = 229;  // { int nosys(void); }</span>
    <a id="L252"></a>SYS_POLL         = 230; <span class="comment">// { int poll(struct pollfd *fds, u_int nfds, int timeout); }</span>
    <a id="L253"></a>SYS_WATCHEVENT   = 231; <span class="comment">// { int watchevent(struct eventreq *u_req, int u_eventmask); }</span>
    <a id="L254"></a>SYS_WAITEVENT    = 232; <span class="comment">// { int waitevent(struct eventreq *u_req, struct timeval *tv); }</span>
    <a id="L255"></a>SYS_MODWATCH     = 233; <span class="comment">// { int modwatch(struct eventreq *u_req, int u_eventmask); }</span>
    <a id="L256"></a>SYS_GETXATTR     = 234; <span class="comment">// { user_ssize_t getxattr(user_addr_t path, user_addr_t attrname, user_addr_t value, size_t size, uint32_t position, int options); }</span>
    <a id="L257"></a>SYS_FGETXATTR    = 235; <span class="comment">// { user_ssize_t fgetxattr(int fd, user_addr_t attrname, user_addr_t value, size_t size, uint32_t position, int options); }</span>
    <a id="L258"></a>SYS_SETXATTR     = 236; <span class="comment">// { int setxattr(user_addr_t path, user_addr_t attrname, user_addr_t value, size_t size, uint32_t position, int options); }</span>
    <a id="L259"></a>SYS_FSETXATTR    = 237; <span class="comment">// { int fsetxattr(int fd, user_addr_t attrname, user_addr_t value, size_t size, uint32_t position, int options); }</span>
    <a id="L260"></a>SYS_REMOVEXATTR  = 238; <span class="comment">// { int removexattr(user_addr_t path, user_addr_t attrname, int options); }</span>
    <a id="L261"></a>SYS_FREMOVEXATTR = 239; <span class="comment">// { int fremovexattr(int fd, user_addr_t attrname, int options); }</span>
    <a id="L262"></a>SYS_LISTXATTR    = 240; <span class="comment">// { user_ssize_t listxattr(user_addr_t path, user_addr_t namebuf, size_t bufsize, int options); }</span>
    <a id="L263"></a>SYS_FLISTXATTR   = 241; <span class="comment">// { user_ssize_t flistxattr(int fd, user_addr_t namebuf, size_t bufsize, int options); }</span>
    <a id="L264"></a>SYS_FSCTL        = 242; <span class="comment">// { int fsctl(const char *path, u_long cmd, caddr_t data, u_long options); }</span>
    <a id="L265"></a>SYS_INITGROUPS   = 243; <span class="comment">// { int initgroups(u_int gidsetsize, gid_t *gidset, int gmuid); }</span>
    <a id="L266"></a>SYS_POSIX_SPAWN  = 244; <span class="comment">// { int posix_spawn(pid_t *pid, const char *path, const struct _posix_spawn_args_desc *adesc, char **argv, char **envp); }</span>
    <a id="L267"></a><span class="comment">// SYS_NOSYS = 245;  // { int nosys(void); }</span>
    <a id="L268"></a><span class="comment">// SYS_NOSYS = 246;  // { int nosys(void); }</span>
    <a id="L269"></a>SYS_NFSCLNT = 247; <span class="comment">// { int nfsclnt(int flag, caddr_t argp); }</span>
    <a id="L270"></a><span class="comment">// SYS_NOSYS = 247;  // { int nosys(void); }</span>
    <a id="L271"></a>SYS_FHOPEN = 248; <span class="comment">// { int fhopen(const struct fhandle *u_fhp, int flags); }</span>
    <a id="L272"></a><span class="comment">// SYS_NOSYS = 248;  // { int nosys(void); }</span>
    <a id="L273"></a><span class="comment">// SYS_NOSYS = 249;  // { int nosys(void); }</span>
    <a id="L274"></a>SYS_MINHERIT = 250; <span class="comment">// { int minherit(void *addr, size_t len, int inherit); }</span>
    <a id="L275"></a>SYS_SEMSYS   = 251; <span class="comment">// { int semsys(u_int which, int a2, int a3, int a4, int a5); }</span>
    <a id="L276"></a><span class="comment">// SYS_NOSYS = 251;  // { int nosys(void); }</span>
    <a id="L277"></a>SYS_MSGSYS = 252; <span class="comment">// { int msgsys(u_int which, int a2, int a3, int a4, int a5); }</span>
    <a id="L278"></a><span class="comment">// SYS_NOSYS = 252;  // { int nosys(void); }</span>
    <a id="L279"></a>SYS_SHMSYS = 253; <span class="comment">// { int shmsys(u_int which, int a2, int a3, int a4); }</span>
    <a id="L280"></a><span class="comment">// SYS_NOSYS = 253;  // { int nosys(void); }</span>
    <a id="L281"></a>SYS_SEMCTL = 254; <span class="comment">// { int semctl(int semid, int semnum, int cmd, semun_t arg); }</span>
    <a id="L282"></a>SYS_SEMGET = 255; <span class="comment">// { int semget(key_t key, int	nsems, int semflg); }</span>
    <a id="L283"></a>SYS_SEMOP  = 256; <span class="comment">// { int semop(int semid, struct sembuf *sops, int nsops); }</span>
    <a id="L284"></a><span class="comment">// SYS_NOSYS = 257;  // { int nosys(void); }</span>
    <a id="L285"></a><span class="comment">// SYS_NOSYS = 254;  // { int nosys(void); }</span>
    <a id="L286"></a><span class="comment">// SYS_NOSYS = 255;  // { int nosys(void); }</span>
    <a id="L287"></a><span class="comment">// SYS_NOSYS = 256;  // { int nosys(void); }</span>
    <a id="L288"></a><span class="comment">// SYS_NOSYS = 257;  // { int nosys(void); }</span>
    <a id="L289"></a>SYS_MSGCTL = 258; <span class="comment">// { int msgctl(int msqid, int cmd, struct	msqid_ds *buf); }</span>
    <a id="L290"></a>SYS_MSGGET = 259; <span class="comment">// { int msgget(key_t key, int msgflg); }</span>
    <a id="L291"></a>SYS_MSGSND = 260; <span class="comment">// { int msgsnd(int msqid, void *msgp, size_t msgsz, int msgflg); }</span>
    <a id="L292"></a>SYS_MSGRCV = 261; <span class="comment">// { user_ssize_t msgrcv(int msqid, void *msgp, size_t msgsz, long msgtyp, int msgflg); }</span>
    <a id="L293"></a><span class="comment">// SYS_NOSYS = 258;  // { int nosys(void); }</span>
    <a id="L294"></a><span class="comment">// SYS_NOSYS = 259;  // { int nosys(void); }</span>
    <a id="L295"></a><span class="comment">// SYS_NOSYS = 260;  // { int nosys(void); }</span>
    <a id="L296"></a><span class="comment">// SYS_NOSYS = 261;  // { int nosys(void); }</span>
    <a id="L297"></a>SYS_SHMAT  = 262; <span class="comment">// { user_addr_t shmat(int shmid, void *shmaddr, int shmflg); }</span>
    <a id="L298"></a>SYS_SHMCTL = 263; <span class="comment">// { int shmctl(int shmid, int cmd, struct shmid_ds *buf); }</span>
    <a id="L299"></a>SYS_SHMDT  = 264; <span class="comment">// { int shmdt(void *shmaddr); }</span>
    <a id="L300"></a>SYS_SHMGET = 265; <span class="comment">// { int shmget(key_t key, size_t size, int shmflg); }</span>
    <a id="L301"></a><span class="comment">// SYS_NOSYS = 262;  // { int nosys(void); }</span>
    <a id="L302"></a><span class="comment">// SYS_NOSYS = 263;  // { int nosys(void); }</span>
    <a id="L303"></a><span class="comment">// SYS_NOSYS = 264;  // { int nosys(void); }</span>
    <a id="L304"></a><span class="comment">// SYS_NOSYS = 265;  // { int nosys(void); }</span>
    <a id="L305"></a>SYS_SHM_OPEN               = 266; <span class="comment">// { int shm_open(const char *name, int oflag, int mode); }</span>
    <a id="L306"></a>SYS_SHM_UNLINK             = 267; <span class="comment">// { int shm_unlink(const char *name); }</span>
    <a id="L307"></a>SYS_SEM_OPEN               = 268; <span class="comment">// { user_addr_t sem_open(const char *name, int oflag, int mode, int value); }</span>
    <a id="L308"></a>SYS_SEM_CLOSE              = 269; <span class="comment">// { int sem_close(sem_t *sem); }</span>
    <a id="L309"></a>SYS_SEM_UNLINK             = 270; <span class="comment">// { int sem_unlink(const char *name); }</span>
    <a id="L310"></a>SYS_SEM_WAIT               = 271; <span class="comment">// { int sem_wait(sem_t *sem); }</span>
    <a id="L311"></a>SYS_SEM_TRYWAIT            = 272; <span class="comment">// { int sem_trywait(sem_t *sem); }</span>
    <a id="L312"></a>SYS_SEM_POST               = 273; <span class="comment">// { int sem_post(sem_t *sem); }</span>
    <a id="L313"></a>SYS_SEM_GETVALUE           = 274; <span class="comment">// { int sem_getvalue(sem_t *sem, int *sval); }</span>
    <a id="L314"></a>SYS_SEM_INIT               = 275; <span class="comment">// { int sem_init(sem_t *sem, int phsared, u_int value); }</span>
    <a id="L315"></a>SYS_SEM_DESTROY            = 276; <span class="comment">// { int sem_destroy(sem_t *sem); }</span>
    <a id="L316"></a>SYS_OPEN_EXTENDED          = 277; <span class="comment">// { int open_extended(user_addr_t path, int flags, uid_t uid, gid_t gid, int mode, user_addr_t xsecurity) NO_SYSCALL_STUB; }</span>
    <a id="L317"></a>SYS_UMASK_EXTENDED         = 278; <span class="comment">// { int umask_extended(int newmask, user_addr_t xsecurity) NO_SYSCALL_STUB; }</span>
    <a id="L318"></a>SYS_STAT_EXTENDED          = 279; <span class="comment">// { int stat_extended(user_addr_t path, user_addr_t ub, user_addr_t xsecurity, user_addr_t xsecurity_size) NO_SYSCALL_STUB; }</span>
    <a id="L319"></a>SYS_LSTAT_EXTENDED         = 280; <span class="comment">// { int lstat_extended(user_addr_t path, user_addr_t ub, user_addr_t xsecurity, user_addr_t xsecurity_size) NO_SYSCALL_STUB; }</span>
    <a id="L320"></a>SYS_FSTAT_EXTENDED         = 281; <span class="comment">// { int fstat_extended(int fd, user_addr_t ub, user_addr_t xsecurity, user_addr_t xsecurity_size) NO_SYSCALL_STUB; }</span>
    <a id="L321"></a>SYS_CHMOD_EXTENDED         = 282; <span class="comment">// { int chmod_extended(user_addr_t path, uid_t uid, gid_t gid, int mode, user_addr_t xsecurity) NO_SYSCALL_STUB; }</span>
    <a id="L322"></a>SYS_FCHMOD_EXTENDED        = 283; <span class="comment">// { int fchmod_extended(int fd, uid_t uid, gid_t gid, int mode, user_addr_t xsecurity) NO_SYSCALL_STUB; }</span>
    <a id="L323"></a>SYS_ACCESS_EXTENDED        = 284; <span class="comment">// { int access_extended(user_addr_t entries, size_t size, user_addr_t results, uid_t uid) NO_SYSCALL_STUB; }</span>
    <a id="L324"></a>SYS_SETTID                 = 285; <span class="comment">// { int settid(uid_t uid, gid_t gid) NO_SYSCALL_STUB; }</span>
    <a id="L325"></a>SYS_GETTID                 = 286; <span class="comment">// { int gettid(uid_t *uidp, gid_t *gidp) NO_SYSCALL_STUB; }</span>
    <a id="L326"></a>SYS_SETSGROUPS             = 287; <span class="comment">// { int setsgroups(int setlen, user_addr_t guidset) NO_SYSCALL_STUB; }</span>
    <a id="L327"></a>SYS_GETSGROUPS             = 288; <span class="comment">// { int getsgroups(user_addr_t setlen, user_addr_t guidset) NO_SYSCALL_STUB; }</span>
    <a id="L328"></a>SYS_SETWGROUPS             = 289; <span class="comment">// { int setwgroups(int setlen, user_addr_t guidset) NO_SYSCALL_STUB; }</span>
    <a id="L329"></a>SYS_GETWGROUPS             = 290; <span class="comment">// { int getwgroups(user_addr_t setlen, user_addr_t guidset) NO_SYSCALL_STUB; }</span>
    <a id="L330"></a>SYS_MKFIFO_EXTENDED        = 291; <span class="comment">// { int mkfifo_extended(user_addr_t path, uid_t uid, gid_t gid, int mode, user_addr_t xsecurity) NO_SYSCALL_STUB; }</span>
    <a id="L331"></a>SYS_MKDIR_EXTENDED         = 292; <span class="comment">// { int mkdir_extended(user_addr_t path, uid_t uid, gid_t gid, int mode, user_addr_t xsecurity) NO_SYSCALL_STUB; }</span>
    <a id="L332"></a>SYS_IDENTITYSVC            = 293; <span class="comment">// { int identitysvc(int opcode, user_addr_t message) NO_SYSCALL_STUB; }</span>
    <a id="L333"></a>SYS_SHARED_REGION_CHECK_NP = 294; <span class="comment">// { int shared_region_check_np(uint64_t *start_address) NO_SYSCALL_STUB; }</span>
    <a id="L334"></a>SYS_SHARED_REGION_MAP_NP   = 295; <span class="comment">// { int shared_region_map_np(int fd, uint32_t count, const struct shared_file_mapping_np *mappings) NO_SYSCALL_STUB; }</span>
    <a id="L335"></a><span class="comment">// SYS_NOSYS = 296;  // { int nosys(void); } { old load_shared_file }</span>
    <a id="L336"></a><span class="comment">// SYS_NOSYS = 297;  // { int nosys(void); } { old reset_shared_file }</span>
    <a id="L337"></a><span class="comment">// SYS_NOSYS = 298;  // { int nosys(void); } { old new_system_shared_regions }</span>
    <a id="L338"></a><span class="comment">// SYS_ENOSYS = 299;  // { int enosys(void); } { old shared_region_map_file_np }</span>
    <a id="L339"></a><span class="comment">// SYS_ENOSYS = 300;  // { int enosys(void); } { old shared_region_make_private_np }</span>
    <a id="L340"></a>SYS___PTHREAD_MUTEX_DESTROY  = 301; <span class="comment">// { int __pthread_mutex_destroy(int mutexid); }</span>
    <a id="L341"></a>SYS___PTHREAD_MUTEX_INIT     = 302; <span class="comment">// { int __pthread_mutex_init(user_addr_t  mutex, user_addr_t attr); }</span>
    <a id="L342"></a>SYS___PTHREAD_MUTEX_LOCK     = 303; <span class="comment">// { int __pthread_mutex_lock(int mutexid); }</span>
    <a id="L343"></a>SYS___PTHREAD_MUTEX_TRYLOCK  = 304; <span class="comment">// { int __pthread_mutex_trylock(int mutexid); }</span>
    <a id="L344"></a>SYS___PTHREAD_MUTEX_UNLOCK   = 305; <span class="comment">// { int __pthread_mutex_unlock(int mutexid); }</span>
    <a id="L345"></a>SYS___PTHREAD_COND_INIT      = 306; <span class="comment">// { int __pthread_cond_init(user_addr_t cond, user_addr_t attr); }</span>
    <a id="L346"></a>SYS___PTHREAD_COND_DESTROY   = 307; <span class="comment">// { int __pthread_cond_destroy(int condid); }</span>
    <a id="L347"></a>SYS___PTHREAD_COND_BROADCAST = 308; <span class="comment">// { int __pthread_cond_broadcast(int condid); }</span>
    <a id="L348"></a>SYS___PTHREAD_COND_SIGNAL    = 309; <span class="comment">// { int __pthread_cond_signal(int condid); }</span>
    <a id="L349"></a>SYS_GETSID                   = 310; <span class="comment">// { int getsid(pid_t pid); }</span>
    <a id="L350"></a>SYS_SETTID_WITH_PID          = 311; <span class="comment">// { int settid_with_pid(pid_t pid, int assume) NO_SYSCALL_STUB; }</span>
    <a id="L351"></a>SYS___PTHREAD_COND_TIMEDWAIT = 312; <span class="comment">// { int __pthread_cond_timedwait(int condid, int mutexid, user_addr_t abstime); }</span>
    <a id="L352"></a>SYS_AIO_FSYNC                = 313; <span class="comment">// { int aio_fsync(int op, user_addr_t aiocbp); }</span>
    <a id="L353"></a>SYS_AIO_RETURN               = 314; <span class="comment">// { user_ssize_t aio_return(user_addr_t aiocbp); }</span>
    <a id="L354"></a>SYS_AIO_SUSPEND              = 315; <span class="comment">// { int aio_suspend(user_addr_t aiocblist, int nent, user_addr_t timeoutp); }</span>
    <a id="L355"></a>SYS_AIO_CANCEL               = 316; <span class="comment">// { int aio_cancel(int fd, user_addr_t aiocbp); }</span>
    <a id="L356"></a>SYS_AIO_ERROR                = 317; <span class="comment">// { int aio_error(user_addr_t aiocbp); }</span>
    <a id="L357"></a>SYS_AIO_READ                 = 318; <span class="comment">// { int aio_read(user_addr_t aiocbp); }</span>
    <a id="L358"></a>SYS_AIO_WRITE                = 319; <span class="comment">// { int aio_write(user_addr_t aiocbp); }</span>
    <a id="L359"></a>SYS_LIO_LISTIO               = 320; <span class="comment">// { int lio_listio(int mode, user_addr_t aiocblist, int nent, user_addr_t sigp); }</span>
    <a id="L360"></a>SYS___PTHREAD_COND_WAIT      = 321; <span class="comment">// { int __pthread_cond_wait(int condid, int mutexid); }</span>
    <a id="L361"></a>SYS_IOPOLICYSYS              = 322; <span class="comment">// { int iopolicysys(int cmd, void *arg) NO_SYSCALL_STUB; }</span>
    <a id="L362"></a><span class="comment">// SYS_NOSYS = 323;  // { int nosys(void); }</span>
    <a id="L363"></a>SYS_MLOCKALL   = 324; <span class="comment">// { int mlockall(int how); }</span>
    <a id="L364"></a>SYS_MUNLOCKALL = 325; <span class="comment">// { int munlockall(int how); }</span>
    <a id="L365"></a><span class="comment">// SYS_NOSYS = 326;  // { int nosys(void); }</span>
    <a id="L366"></a>SYS_ISSETUGID              = 327; <span class="comment">// { int issetugid(void); }</span>
    <a id="L367"></a>SYS___PTHREAD_KILL         = 328; <span class="comment">// { int __pthread_kill(int thread_port, int sig); }</span>
    <a id="L368"></a>SYS___PTHREAD_SIGMASK      = 329; <span class="comment">// { int __pthread_sigmask(int how, user_addr_t set, user_addr_t oset); }</span>
    <a id="L369"></a>SYS___SIGWAIT              = 330; <span class="comment">// { int __sigwait(user_addr_t set, user_addr_t sig); }</span>
    <a id="L370"></a>SYS___DISABLE_THREADSIGNAL = 331; <span class="comment">// { int __disable_threadsignal(int value); }</span>
    <a id="L371"></a>SYS___PTHREAD_MARKCANCEL   = 332; <span class="comment">// { int __pthread_markcancel(int thread_port); }</span>
    <a id="L372"></a>SYS___PTHREAD_CANCELED     = 333; <span class="comment">// { int __pthread_canceled(int  action); }</span>
    <a id="L373"></a>SYS___SEMWAIT_SIGNAL       = 334; <span class="comment">// { int __semwait_signal(int cond_sem, int mutex_sem, int timeout, int relative, time_t tv_sec, int32_t tv_nsec); }</span>
    <a id="L374"></a><span class="comment">// SYS_NOSYS = 335;  // { int nosys(void); }   { old utrace }</span>
    <a id="L375"></a>SYS_PROC_INFO = 336; <span class="comment">// { int proc_info(int32_t callnum,int32_t pid,uint32_t flavor, uint64_t arg,user_addr_t buffer,int32_t buffersize) NO_SYSCALL_STUB; }</span>
    <a id="L376"></a>SYS_SENDFILE  = 337; <span class="comment">// { int sendfile(int fd, int s, off_t offset, off_t *nbytes, struct sf_hdtr *hdtr, int flags); }</span>
    <a id="L377"></a><span class="comment">// SYS_NOSYS = 337;  // { int nosys(void); }</span>
    <a id="L378"></a>SYS_STAT64           = 338; <span class="comment">// { int stat64(user_addr_t path, user_addr_t ub); }</span>
    <a id="L379"></a>SYS_FSTAT64          = 339; <span class="comment">// { int fstat64(int fd, user_addr_t ub); }</span>
    <a id="L380"></a>SYS_LSTAT64          = 340; <span class="comment">// { int lstat64(user_addr_t path, user_addr_t ub); }</span>
    <a id="L381"></a>SYS_STAT64_EXTENDED  = 341; <span class="comment">// { int stat64_extended(user_addr_t path, user_addr_t ub, user_addr_t xsecurity, user_addr_t xsecurity_size) NO_SYSCALL_STUB; }</span>
    <a id="L382"></a>SYS_LSTAT64_EXTENDED = 342; <span class="comment">// { int lstat64_extended(user_addr_t path, user_addr_t ub, user_addr_t xsecurity, user_addr_t xsecurity_size) NO_SYSCALL_STUB; }</span>
    <a id="L383"></a>SYS_FSTAT64_EXTENDED = 343; <span class="comment">// { int fstat64_extended(int fd, user_addr_t ub, user_addr_t xsecurity, user_addr_t xsecurity_size) NO_SYSCALL_STUB; }</span>
    <a id="L384"></a>SYS_GETDIRENTRIES64  = 344; <span class="comment">// { user_ssize_t getdirentries64(int fd, void *buf, user_size_t bufsize, off_t *position) NO_SYSCALL_STUB; }</span>
    <a id="L385"></a>SYS_STATFS64         = 345; <span class="comment">// { int statfs64(char *path, struct statfs64 *buf); }</span>
    <a id="L386"></a>SYS_FSTATFS64        = 346; <span class="comment">// { int fstatfs64(int fd, struct statfs64 *buf); }</span>
    <a id="L387"></a>SYS_GETFSSTAT64      = 347; <span class="comment">// { int getfsstat64(user_addr_t buf, int bufsize, int flags); }</span>
    <a id="L388"></a>SYS___PTHREAD_CHDIR  = 348; <span class="comment">// { int __pthread_chdir(user_addr_t path); }</span>
    <a id="L389"></a>SYS___PTHREAD_FCHDIR = 349; <span class="comment">// { int __pthread_fchdir(int fd); }</span>
    <a id="L390"></a>SYS_AUDIT            = 350; <span class="comment">// { int audit(void *record, int length); }</span>
    <a id="L391"></a>SYS_AUDITON          = 351; <span class="comment">// { int auditon(int cmd, void *data, int length); }</span>
    <a id="L392"></a><span class="comment">// SYS_NOSYS = 352;  // { int nosys(void); }</span>
    <a id="L393"></a>SYS_GETAUID       = 353; <span class="comment">// { int getauid(au_id_t *auid); }</span>
    <a id="L394"></a>SYS_SETAUID       = 354; <span class="comment">// { int setauid(au_id_t *auid); }</span>
    <a id="L395"></a>SYS_GETAUDIT      = 355; <span class="comment">// { int getaudit(struct auditinfo *auditinfo); }</span>
    <a id="L396"></a>SYS_SETAUDIT      = 356; <span class="comment">// { int setaudit(struct auditinfo *auditinfo); }</span>
    <a id="L397"></a>SYS_GETAUDIT_ADDR = 357; <span class="comment">// { int getaudit_addr(struct auditinfo_addr *auditinfo_addr, int length); }</span>
    <a id="L398"></a>SYS_SETAUDIT_ADDR = 358; <span class="comment">// { int setaudit_addr(struct auditinfo_addr *auditinfo_addr, int length); }</span>
    <a id="L399"></a>SYS_AUDITCTL      = 359; <span class="comment">// { int auditctl(char *path); }</span>
    <a id="L400"></a><span class="comment">// SYS_NOSYS = 350;  // { int nosys(void); }</span>
    <a id="L401"></a><span class="comment">// SYS_NOSYS = 351;  // { int nosys(void); }</span>
    <a id="L402"></a><span class="comment">// SYS_NOSYS = 352;  // { int nosys(void); }</span>
    <a id="L403"></a><span class="comment">// SYS_NOSYS = 353;  // { int nosys(void); }</span>
    <a id="L404"></a><span class="comment">// SYS_NOSYS = 354;  // { int nosys(void); }</span>
    <a id="L405"></a><span class="comment">// SYS_NOSYS = 355;  // { int nosys(void); }</span>
    <a id="L406"></a><span class="comment">// SYS_NOSYS = 356;  // { int nosys(void); }</span>
    <a id="L407"></a><span class="comment">// SYS_NOSYS = 357;  // { int nosys(void); }</span>
    <a id="L408"></a><span class="comment">// SYS_NOSYS = 358;  // { int nosys(void); }</span>
    <a id="L409"></a><span class="comment">// SYS_NOSYS = 359;  // { int nosys(void); }</span>
    <a id="L410"></a>SYS_BSDTHREAD_CREATE    = 360; <span class="comment">// { user_addr_t bsdthread_create(user_addr_t func, user_addr_t func_arg, user_addr_t stack, user_addr_t pthread, uint32_t flags) NO_SYSCALL_STUB; }</span>
    <a id="L411"></a>SYS_BSDTHREAD_TERMINATE = 361; <span class="comment">// { int bsdthread_terminate(user_addr_t stackaddr, size_t freesize, uint32_t port, uint32_t sem) NO_SYSCALL_STUB; }</span>
    <a id="L412"></a>SYS_KQUEUE              = 362; <span class="comment">// { int kqueue(void); }</span>
    <a id="L413"></a>SYS_KEVENT              = 363; <span class="comment">// { int kevent(int fd, const struct kevent *changelist, int nchanges, struct kevent *eventlist, int nevents, const struct timespec *timeout); }</span>
    <a id="L414"></a>SYS_LCHOWN              = 364; <span class="comment">// { int lchown(user_addr_t path, uid_t owner, gid_t group); }</span>
    <a id="L415"></a>SYS_STACK_SNAPSHOT      = 365; <span class="comment">// { int stack_snapshot(pid_t pid, user_addr_t tracebuf, uint32_t tracebuf_size, uint32_t options) NO_SYSCALL_STUB; }</span>
    <a id="L416"></a>SYS_BSDTHREAD_REGISTER  = 366; <span class="comment">// { int bsdthread_register(user_addr_t threadstart, user_addr_t wqthread, int pthsize) NO_SYSCALL_STUB; }</span>
    <a id="L417"></a>SYS_WORKQ_OPEN          = 367; <span class="comment">// { int workq_open(void) NO_SYSCALL_STUB; }</span>
    <a id="L418"></a>SYS_WORKQ_OPS           = 368; <span class="comment">// { int workq_ops(int options, user_addr_t item, int prio) NO_SYSCALL_STUB; }</span>
    <a id="L419"></a><span class="comment">// SYS_NOSYS = 369;  // { int nosys(void); }</span>
    <a id="L420"></a><span class="comment">// SYS_NOSYS = 370;  // { int nosys(void); }</span>
    <a id="L421"></a><span class="comment">// SYS_NOSYS = 371;  // { int nosys(void); }</span>
    <a id="L422"></a><span class="comment">// SYS_NOSYS = 372;  // { int nosys(void); }</span>
    <a id="L423"></a><span class="comment">// SYS_NOSYS = 373;  // { int nosys(void); }</span>
    <a id="L424"></a><span class="comment">// SYS_NOSYS = 374;  // { int nosys(void); }</span>
    <a id="L425"></a><span class="comment">// SYS_NOSYS = 375;  // { int nosys(void); }</span>
    <a id="L426"></a><span class="comment">// SYS_NOSYS = 376;  // { int nosys(void); }</span>
    <a id="L427"></a><span class="comment">// SYS_NOSYS = 377;  // { int nosys(void); }</span>
    <a id="L428"></a><span class="comment">// SYS_NOSYS = 378;  // { int nosys(void); }</span>
    <a id="L429"></a><span class="comment">// SYS_NOSYS = 379;  // { int nosys(void); }</span>
    <a id="L430"></a>SYS___MAC_EXECVE      = 380; <span class="comment">// { int __mac_execve(char *fname, char **argp, char **envp, struct mac *mac_p); }</span>
    <a id="L431"></a>SYS___MAC_SYSCALL     = 381; <span class="comment">// { int __mac_syscall(char *policy, int call, user_addr_t arg); }</span>
    <a id="L432"></a>SYS___MAC_GET_FILE    = 382; <span class="comment">// { int __mac_get_file(char *path_p, struct mac *mac_p); }</span>
    <a id="L433"></a>SYS___MAC_SET_FILE    = 383; <span class="comment">// { int __mac_set_file(char *path_p, struct mac *mac_p); }</span>
    <a id="L434"></a>SYS___MAC_GET_LINK    = 384; <span class="comment">// { int __mac_get_link(char *path_p, struct mac *mac_p); }</span>
    <a id="L435"></a>SYS___MAC_SET_LINK    = 385; <span class="comment">// { int __mac_set_link(char *path_p, struct mac *mac_p); }</span>
    <a id="L436"></a>SYS___MAC_GET_PROC    = 386; <span class="comment">// { int __mac_get_proc(struct mac *mac_p); }</span>
    <a id="L437"></a>SYS___MAC_SET_PROC    = 387; <span class="comment">// { int __mac_set_proc(struct mac *mac_p); }</span>
    <a id="L438"></a>SYS___MAC_GET_FD      = 388; <span class="comment">// { int __mac_get_fd(int fd, struct mac *mac_p); }</span>
    <a id="L439"></a>SYS___MAC_SET_FD      = 389; <span class="comment">// { int __mac_set_fd(int fd, struct mac *mac_p); }</span>
    <a id="L440"></a>SYS___MAC_GET_PID     = 390; <span class="comment">// { int __mac_get_pid(pid_t pid, struct mac *mac_p); }</span>
    <a id="L441"></a>SYS___MAC_GET_LCID    = 391; <span class="comment">// { int __mac_get_lcid(pid_t lcid, struct mac *mac_p); }</span>
    <a id="L442"></a>SYS___MAC_GET_LCTX    = 392; <span class="comment">// { int __mac_get_lctx(struct mac *mac_p); }</span>
    <a id="L443"></a>SYS___MAC_SET_LCTX    = 393; <span class="comment">// { int __mac_set_lctx(struct mac *mac_p); }</span>
    <a id="L444"></a>SYS_SETLCID           = 394; <span class="comment">// { int setlcid(pid_t pid, pid_t lcid) NO_SYSCALL_STUB; }</span>
    <a id="L445"></a>SYS_GETLCID           = 395; <span class="comment">// { int getlcid(pid_t pid) NO_SYSCALL_STUB; }</span>
    <a id="L446"></a>SYS_READ_NOCANCEL     = 396; <span class="comment">// { user_ssize_t read_nocancel(int fd, user_addr_t cbuf, user_size_t nbyte) NO_SYSCALL_STUB; }</span>
    <a id="L447"></a>SYS_WRITE_NOCANCEL    = 397; <span class="comment">// { user_ssize_t write_nocancel(int fd, user_addr_t cbuf, user_size_t nbyte) NO_SYSCALL_STUB; }</span>
    <a id="L448"></a>SYS_OPEN_NOCANCEL     = 398; <span class="comment">// { int open_nocancel(user_addr_t path, int flags, int mode) NO_SYSCALL_STUB; }</span>
    <a id="L449"></a>SYS_CLOSE_NOCANCEL    = 399; <span class="comment">// { int close_nocancel(int fd) NO_SYSCALL_STUB; }</span>
    <a id="L450"></a>SYS_WAIT4_NOCANCEL    = 400; <span class="comment">// { int wait4_nocancel(int pid, user_addr_t status, int options, user_addr_t rusage) NO_SYSCALL_STUB; }</span>
    <a id="L451"></a>SYS_RECVMSG_NOCANCEL  = 401; <span class="comment">// { int recvmsg_nocancel(int s, struct msghdr *msg, int flags) NO_SYSCALL_STUB; }</span>
    <a id="L452"></a>SYS_SENDMSG_NOCANCEL  = 402; <span class="comment">// { int sendmsg_nocancel(int s, caddr_t msg, int flags) NO_SYSCALL_STUB; }</span>
    <a id="L453"></a>SYS_RECVFROM_NOCANCEL = 403; <span class="comment">// { int recvfrom_nocancel(int s, void *buf, size_t len, int flags, struct sockaddr *from, int *fromlenaddr) NO_SYSCALL_STUB; }</span>
    <a id="L454"></a>SYS_ACCEPT_NOCANCEL   = 404; <span class="comment">// { int accept_nocancel(int s, caddr_t name, socklen_t	*anamelen) NO_SYSCALL_STUB; }</span>
    <a id="L455"></a><span class="comment">// SYS_NOSYS = 401;  // { int nosys(void); }</span>
    <a id="L456"></a><span class="comment">// SYS_NOSYS = 402;  // { int nosys(void); }</span>
    <a id="L457"></a><span class="comment">// SYS_NOSYS = 403;  // { int nosys(void); }</span>
    <a id="L458"></a><span class="comment">// SYS_NOSYS = 404;  // { int nosys(void); }</span>
    <a id="L459"></a>SYS_MSYNC_NOCANCEL   = 405; <span class="comment">// { int msync_nocancel(caddr_t addr, size_t len, int flags) NO_SYSCALL_STUB; }</span>
    <a id="L460"></a>SYS_FCNTL_NOCANCEL   = 406; <span class="comment">// { int fcntl_nocancel(int fd, int cmd, long arg) NO_SYSCALL_STUB; }</span>
    <a id="L461"></a>SYS_SELECT_NOCANCEL  = 407; <span class="comment">// { int select_nocancel(int nd, u_int32_t *in, u_int32_t *ou, u_int32_t *ex, struct timeval *tv) NO_SYSCALL_STUB; }</span>
    <a id="L462"></a>SYS_FSYNC_NOCANCEL   = 408; <span class="comment">// { int fsync_nocancel(int fd) NO_SYSCALL_STUB; }</span>
    <a id="L463"></a>SYS_CONNECT_NOCANCEL = 409; <span class="comment">// { int connect_nocancel(int s, caddr_t name, socklen_t namelen) NO_SYSCALL_STUB; }</span>
    <a id="L464"></a><span class="comment">// SYS_NOSYS = 409;  // { int nosys(void); }</span>
    <a id="L465"></a>SYS_SIGSUSPEND_NOCANCEL = 410; <span class="comment">// { int sigsuspend_nocancel(sigset_t mask) NO_SYSCALL_STUB; }</span>
    <a id="L466"></a>SYS_READV_NOCANCEL      = 411; <span class="comment">// { user_ssize_t readv_nocancel(int fd, struct iovec *iovp, u_int iovcnt) NO_SYSCALL_STUB; }</span>
    <a id="L467"></a>SYS_WRITEV_NOCANCEL     = 412; <span class="comment">// { user_ssize_t writev_nocancel(int fd, struct iovec *iovp, u_int iovcnt) NO_SYSCALL_STUB; }</span>
    <a id="L468"></a>SYS_SENDTO_NOCANCEL     = 413; <span class="comment">// { int sendto_nocancel(int s, caddr_t buf, size_t len, int flags, caddr_t to, socklen_t tolen) NO_SYSCALL_STUB; }</span>
    <a id="L469"></a><span class="comment">// SYS_NOSYS = 413;  // { int nosys(void); }</span>
    <a id="L470"></a>SYS_PREAD_NOCANCEL  = 414; <span class="comment">// { user_ssize_t pread_nocancel(int fd, user_addr_t buf, user_size_t nbyte, off_t offset) NO_SYSCALL_STUB; }</span>
    <a id="L471"></a>SYS_PWRITE_NOCANCEL = 415; <span class="comment">// { user_ssize_t pwrite_nocancel(int fd, user_addr_t buf, user_size_t nbyte, off_t offset) NO_SYSCALL_STUB; }</span>
    <a id="L472"></a>SYS_WAITID_NOCANCEL = 416; <span class="comment">// { int waitid_nocancel(idtype_t idtype, id_t id, siginfo_t *infop, int options) NO_SYSCALL_STUB; }</span>
    <a id="L473"></a>SYS_POLL_NOCANCEL   = 417; <span class="comment">// { int poll_nocancel(struct pollfd *fds, u_int nfds, int timeout) NO_SYSCALL_STUB; }</span>
    <a id="L474"></a>SYS_MSGSND_NOCANCEL = 418; <span class="comment">// { int msgsnd_nocancel(int msqid, void *msgp, size_t msgsz, int msgflg) NO_SYSCALL_STUB; }</span>
    <a id="L475"></a>SYS_MSGRCV_NOCANCEL = 419; <span class="comment">// { user_ssize_t msgrcv_nocancel(int msqid, void *msgp, size_t msgsz, long msgtyp, int msgflg) NO_SYSCALL_STUB; }</span>
    <a id="L476"></a><span class="comment">// SYS_NOSYS = 418;  // { int nosys(void); }</span>
    <a id="L477"></a><span class="comment">// SYS_NOSYS = 419;  // { int nosys(void); }</span>
    <a id="L478"></a>SYS_SEM_WAIT_NOCANCEL         = 420; <span class="comment">// { int sem_wait_nocancel(sem_t *sem) NO_SYSCALL_STUB; }</span>
    <a id="L479"></a>SYS_AIO_SUSPEND_NOCANCEL      = 421; <span class="comment">// { int aio_suspend_nocancel(user_addr_t aiocblist, int nent, user_addr_t timeoutp) NO_SYSCALL_STUB; }</span>
    <a id="L480"></a>SYS___SIGWAIT_NOCANCEL        = 422; <span class="comment">// { int __sigwait_nocancel(user_addr_t set, user_addr_t sig) NO_SYSCALL_STUB; }</span>
    <a id="L481"></a>SYS___SEMWAIT_SIGNAL_NOCANCEL = 423; <span class="comment">// { int __semwait_signal_nocancel(int cond_sem, int mutex_sem, int timeout, int relative, time_t tv_sec, int32_t tv_nsec) NO_SYSCALL_STUB; }</span>
    <a id="L482"></a>SYS___MAC_MOUNT               = 424; <span class="comment">// { int __mac_mount(char *type, char *path, int flags, caddr_t data, struct mac *mac_p); }</span>
    <a id="L483"></a>SYS___MAC_GET_MOUNT           = 425; <span class="comment">// { int __mac_get_mount(char *path, struct mac *mac_p); }</span>
    <a id="L484"></a>SYS___MAC_GETFSSTAT           = 426; <span class="comment">// { int __mac_getfsstat(user_addr_t buf, int bufsize, user_addr_t mac, int macsize, int flags); }</span>
<a id="L485"></a>)
</pre>

</div>

<div id="footer">
<p>Except as noted, this content is
   licensed under <a href="http://creativecommons.org/licenses/by/3.0/">
   Creative Commons Attribution 3.0</a>.
</div>

<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-11222381-2");
pageTracker._trackPageview();
</script>
</body>
</html>
<!-- generated at Thu Nov 12 15:42:51 PST 2009 -->
