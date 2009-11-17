<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/ztypes_linux_arm.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/syscall/ztypes_linux_arm.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// godefs -gsyscall -f-m32 types_linux.c</span>

<a id="L3"></a><span class="comment">// MACHINE GENERATED - DO NOT EDIT.</span>

<a id="L5"></a>package syscall

<a id="L7"></a><span class="comment">// Constants</span>
<a id="L8"></a>const (
    <a id="L9"></a>sizeofPtr               = 0x4;
    <a id="L10"></a>sizeofShort             = 0x2;
    <a id="L11"></a>sizeofInt               = 0x4;
    <a id="L12"></a>sizeofLong              = 0x4;
    <a id="L13"></a>sizeofLongLong          = 0x8;
    <a id="L14"></a>PathMax                 = 0x1000;
    <a id="L15"></a>O_RDONLY                = 0;
    <a id="L16"></a>O_WRONLY                = 0x1;
    <a id="L17"></a>O_RDWR                  = 0x2;
    <a id="L18"></a>O_APPEND                = 0x400;
    <a id="L19"></a>O_ASYNC                 = 0x2000;
    <a id="L20"></a>O_CREAT                 = 0x40;
    <a id="L21"></a>O_NOCTTY                = 0x100;
    <a id="L22"></a>O_NONBLOCK              = 0x800;
    <a id="L23"></a>O_SYNC                  = 0x1000;
    <a id="L24"></a>O_TRUNC                 = 0x200;
    <a id="L25"></a>O_CLOEXEC               = 0;
    <a id="L26"></a>F_GETFD                 = 0x1;
    <a id="L27"></a>F_SETFD                 = 0x2;
    <a id="L28"></a>F_GETFL                 = 0x3;
    <a id="L29"></a>F_SETFL                 = 0x4;
    <a id="L30"></a>FD_CLOEXEC              = 0x1;
    <a id="L31"></a>NAME_MAX                = 0xff;
    <a id="L32"></a>S_IFMT                  = 0xf000;
    <a id="L33"></a>S_IFIFO                 = 0x1000;
    <a id="L34"></a>S_IFCHR                 = 0x2000;
    <a id="L35"></a>S_IFDIR                 = 0x4000;
    <a id="L36"></a>S_IFBLK                 = 0x6000;
    <a id="L37"></a>S_IFREG                 = 0x8000;
    <a id="L38"></a>S_IFLNK                 = 0xa000;
    <a id="L39"></a>S_IFSOCK                = 0xc000;
    <a id="L40"></a>S_ISUID                 = 0x800;
    <a id="L41"></a>S_ISGID                 = 0x400;
    <a id="L42"></a>S_ISVTX                 = 0x200;
    <a id="L43"></a>S_IRUSR                 = 0x100;
    <a id="L44"></a>S_IWUSR                 = 0x80;
    <a id="L45"></a>S_IXUSR                 = 0x40;
    <a id="L46"></a>WNOHANG                 = 0x1;
    <a id="L47"></a>WUNTRACED               = 0x2;
    <a id="L48"></a>WEXITED                 = 0x4;
    <a id="L49"></a>WSTOPPED                = 0x2;
    <a id="L50"></a>WCONTINUED              = 0x8;
    <a id="L51"></a>WNOWAIT                 = 0x1000000;
    <a id="L52"></a>WCLONE                  = 0x80000000;
    <a id="L53"></a>WALL                    = 0x40000000;
    <a id="L54"></a>WNOTHREAD               = 0x20000000;
    <a id="L55"></a>AF_UNIX                 = 0x1;
    <a id="L56"></a>AF_INET                 = 0x2;
    <a id="L57"></a>AF_INET6                = 0xa;
    <a id="L58"></a>SOCK_STREAM             = 0x1;
    <a id="L59"></a>SOCK_DGRAM              = 0x2;
    <a id="L60"></a>SOCK_RAW                = 0x3;
    <a id="L61"></a>SOCK_SEQPACKET          = 0x5;
    <a id="L62"></a>SOL_SOCKET              = 0x1;
    <a id="L63"></a>SO_REUSEADDR            = 0x2;
    <a id="L64"></a>SO_KEEPALIVE            = 0x9;
    <a id="L65"></a>SO_DONTROUTE            = 0x5;
    <a id="L66"></a>SO_BROADCAST            = 0x6;
    <a id="L67"></a>SO_LINGER               = 0xd;
    <a id="L68"></a>SO_SNDBUF               = 0x7;
    <a id="L69"></a>SO_RCVBUF               = 0x8;
    <a id="L70"></a>SO_SNDTIMEO             = 0x15;
    <a id="L71"></a>SO_RCVTIMEO             = 0x14;
    <a id="L72"></a>IPPROTO_TCP             = 0x6;
    <a id="L73"></a>IPPROTO_UDP             = 0x11;
    <a id="L74"></a>TCP_NODELAY             = 0x1;
    <a id="L75"></a>SOMAXCONN               = 0x80;
    <a id="L76"></a>SizeofSockaddrInet4     = 0x10;
    <a id="L77"></a>SizeofSockaddrInet6     = 0x1c;
    <a id="L78"></a>SizeofSockaddrAny       = 0x1c;
    <a id="L79"></a>SizeofSockaddrUnix      = 0x6e;
    <a id="L80"></a>PTRACE_TRACEME          = 0;
    <a id="L81"></a>PTRACE_PEEKTEXT         = 0x1;
    <a id="L82"></a>PTRACE_PEEKDATA         = 0x2;
    <a id="L83"></a>PTRACE_PEEKUSER         = 0x3;
    <a id="L84"></a>PTRACE_POKETEXT         = 0x4;
    <a id="L85"></a>PTRACE_POKEDATA         = 0x5;
    <a id="L86"></a>PTRACE_POKEUSER         = 0x6;
    <a id="L87"></a>PTRACE_CONT             = 0x7;
    <a id="L88"></a>PTRACE_KILL             = 0x8;
    <a id="L89"></a>PTRACE_SINGLESTEP       = 0x9;
    <a id="L90"></a>PTRACE_GETREGS          = 0xc;
    <a id="L91"></a>PTRACE_SETREGS          = 0xd;
    <a id="L92"></a>PTRACE_GETFPREGS        = 0xe;
    <a id="L93"></a>PTRACE_SETFPREGS        = 0xf;
    <a id="L94"></a>PTRACE_ATTACH           = 0x10;
    <a id="L95"></a>PTRACE_DETACH           = 0x11;
    <a id="L96"></a>PTRACE_GETFPXREGS       = 0x12;
    <a id="L97"></a>PTRACE_SETFPXREGS       = 0x13;
    <a id="L98"></a>PTRACE_SYSCALL          = 0x18;
    <a id="L99"></a>PTRACE_SETOPTIONS       = 0x4200;
    <a id="L100"></a>PTRACE_GETEVENTMSG      = 0x4201;
    <a id="L101"></a>PTRACE_GETSIGINFO       = 0x4202;
    <a id="L102"></a>PTRACE_SETSIGINFO       = 0x4203;
    <a id="L103"></a>PTRACE_O_TRACESYSGOOD   = 0x1;
    <a id="L104"></a>PTRACE_O_TRACEFORK      = 0x2;
    <a id="L105"></a>PTRACE_O_TRACEVFORK     = 0x4;
    <a id="L106"></a>PTRACE_O_TRACECLONE     = 0x8;
    <a id="L107"></a>PTRACE_O_TRACEEXEC      = 0x10;
    <a id="L108"></a>PTRACE_O_TRACEVFORKDONE = 0x20;
    <a id="L109"></a>PTRACE_O_TRACEEXIT      = 0x40;
    <a id="L110"></a>PTRACE_O_MASK           = 0x7f;
    <a id="L111"></a>PTRACE_EVENT_FORK       = 0x1;
    <a id="L112"></a>PTRACE_EVENT_VFORK      = 0x2;
    <a id="L113"></a>PTRACE_EVENT_CLONE      = 0x3;
    <a id="L114"></a>PTRACE_EVENT_EXEC       = 0x4;
    <a id="L115"></a>PTRACE_EVENT_VFORK_DONE = 0x5;
    <a id="L116"></a>PTRACE_EVENT_EXIT       = 0x6;
    <a id="L117"></a>EPOLLIN                 = 0x1;
    <a id="L118"></a>EPOLLRDHUP              = 0x2000;
    <a id="L119"></a>EPOLLOUT                = 0x4;
    <a id="L120"></a>EPOLLONESHOT            = 0x40000000;
    <a id="L121"></a>EPOLL_CTL_MOD           = 0x3;
    <a id="L122"></a>EPOLL_CTL_ADD           = 0x1;
    <a id="L123"></a>EPOLL_CTL_DEL           = 0x2;
<a id="L124"></a>)

<a id="L126"></a><span class="comment">// Types</span>

<a id="L128"></a>type _C_short int16

<a id="L130"></a>type _C_int int32

<a id="L132"></a>type _C_long int32

<a id="L134"></a>type _C_long_long int64

<a id="L136"></a>type Timespec struct {
    <a id="L137"></a>Sec  int32;
    <a id="L138"></a>Nsec int32;
<a id="L139"></a>}

<a id="L141"></a>type Timeval struct {
    <a id="L142"></a>Sec  int32;
    <a id="L143"></a>Usec int32;
<a id="L144"></a>}

<a id="L146"></a>type Timex struct {
    <a id="L147"></a>Modes     uint32;
    <a id="L148"></a>Offset    int32;
    <a id="L149"></a>Freq      int32;
    <a id="L150"></a>Maxerror  int32;
    <a id="L151"></a>Esterror  int32;
    <a id="L152"></a>Status    int32;
    <a id="L153"></a>Constant  int32;
    <a id="L154"></a>Precision int32;
    <a id="L155"></a>Tolerance int32;
    <a id="L156"></a>Time      Timeval;
    <a id="L157"></a>Tick      int32;
    <a id="L158"></a>Ppsfreq   int32;
    <a id="L159"></a>Jitter    int32;
    <a id="L160"></a>Shift     int32;
    <a id="L161"></a>Stabil    int32;
    <a id="L162"></a>Jitcnt    int32;
    <a id="L163"></a>Calcnt    int32;
    <a id="L164"></a>Errcnt    int32;
    <a id="L165"></a>Stbcnt    int32;
    <a id="L166"></a>Pad0      int32;
    <a id="L167"></a>Pad1      int32;
    <a id="L168"></a>Pad2      int32;
    <a id="L169"></a>Pad3      int32;
    <a id="L170"></a>Pad4      int32;
    <a id="L171"></a>Pad5      int32;
    <a id="L172"></a>Pad6      int32;
    <a id="L173"></a>Pad7      int32;
    <a id="L174"></a>Pad8      int32;
    <a id="L175"></a>Pad9      int32;
    <a id="L176"></a>Pad10     int32;
    <a id="L177"></a>Pad11     int32;
<a id="L178"></a>}

<a id="L180"></a>type Time_t int32

<a id="L182"></a>type Tms struct {
    <a id="L183"></a>Utime  int32;
    <a id="L184"></a>Stime  int32;
    <a id="L185"></a>Cutime int32;
    <a id="L186"></a>Cstime int32;
<a id="L187"></a>}

<a id="L189"></a>type Utimbuf struct {
    <a id="L190"></a>Actime  int32;
    <a id="L191"></a>Modtime int32;
<a id="L192"></a>}

<a id="L194"></a>type Rusage struct {
    <a id="L195"></a>Utime    Timeval;
    <a id="L196"></a>Stime    Timeval;
    <a id="L197"></a>Maxrss   int32;
    <a id="L198"></a>Ixrss    int32;
    <a id="L199"></a>Idrss    int32;
    <a id="L200"></a>Isrss    int32;
    <a id="L201"></a>Minflt   int32;
    <a id="L202"></a>Majflt   int32;
    <a id="L203"></a>Nswap    int32;
    <a id="L204"></a>Inblock  int32;
    <a id="L205"></a>Oublock  int32;
    <a id="L206"></a>Msgsnd   int32;
    <a id="L207"></a>Msgrcv   int32;
    <a id="L208"></a>Nsignals int32;
    <a id="L209"></a>Nvcsw    int32;
    <a id="L210"></a>Nivcsw   int32;
<a id="L211"></a>}

<a id="L213"></a>type Rlimit struct {
    <a id="L214"></a>Cur uint64;
    <a id="L215"></a>Max uint64;
<a id="L216"></a>}

<a id="L218"></a>type _Gid_t uint32

<a id="L220"></a>type Stat_t struct {
    <a id="L221"></a>Dev      uint64;
    <a id="L222"></a>__pad1   uint16;
    <a id="L223"></a>Pad0     [2]byte;
    <a id="L224"></a>__st_ino uint32;
    <a id="L225"></a>Mode     uint32;
    <a id="L226"></a>Nlink    uint32;
    <a id="L227"></a>Uid      uint32;
    <a id="L228"></a>Gid      uint32;
    <a id="L229"></a>Rdev     uint64;
    <a id="L230"></a>__pad2   uint16;
    <a id="L231"></a>Pad1     [2]byte;
    <a id="L232"></a>Size     int64;
    <a id="L233"></a>Blksize  int32;
    <a id="L234"></a>Blocks   int64;
    <a id="L235"></a>Atim     Timespec;
    <a id="L236"></a>Mtim     Timespec;
    <a id="L237"></a>Ctim     Timespec;
    <a id="L238"></a>Ino      uint64;
<a id="L239"></a>}

<a id="L241"></a>type Statfs_t struct {
    <a id="L242"></a>Type    int32;
    <a id="L243"></a>Bsize   int32;
    <a id="L244"></a>Blocks  uint64;
    <a id="L245"></a>Bfree   uint64;
    <a id="L246"></a>Bavail  uint64;
    <a id="L247"></a>Files   uint64;
    <a id="L248"></a>Ffree   uint64;
    <a id="L249"></a>Fsid    [8]byte; <span class="comment">/* __fsid_t */</span>
    <a id="L250"></a>Namelen int32;
    <a id="L251"></a>Frsize  int32;
    <a id="L252"></a>Spare   [5]int32;
<a id="L253"></a>}

<a id="L255"></a>type Dirent struct {
    <a id="L256"></a>Ino    uint64;
    <a id="L257"></a>Off    int64;
    <a id="L258"></a>Reclen uint16;
    <a id="L259"></a>Type   uint8;
    <a id="L260"></a>Name   [256]int8;
    <a id="L261"></a>Pad0   [1]byte;
<a id="L262"></a>}

<a id="L264"></a>type RawSockaddrInet4 struct {
    <a id="L265"></a>Family uint16;
    <a id="L266"></a>Port   uint16;
    <a id="L267"></a>Addr   [4]byte; <span class="comment">/* in_addr */</span>
    <a id="L268"></a>Zero   [8]uint8;
<a id="L269"></a>}

<a id="L271"></a>type RawSockaddrInet6 struct {
    <a id="L272"></a>Family   uint16;
    <a id="L273"></a>Port     uint16;
    <a id="L274"></a>Flowinfo uint32;
    <a id="L275"></a>Addr     [16]byte; <span class="comment">/* in6_addr */</span>
    <a id="L276"></a>Scope_id uint32;
<a id="L277"></a>}

<a id="L279"></a>type RawSockaddrUnix struct {
    <a id="L280"></a>Family uint16;
    <a id="L281"></a>Path   [108]int8;
<a id="L282"></a>}

<a id="L284"></a>type RawSockaddr struct {
    <a id="L285"></a>Family uint16;
    <a id="L286"></a>Data   [14]int8;
<a id="L287"></a>}

<a id="L289"></a>type RawSockaddrAny struct {
    <a id="L290"></a>Addr RawSockaddr;
    <a id="L291"></a>Pad  [12]int8;
<a id="L292"></a>}

<a id="L294"></a>type _Socklen uint32

<a id="L296"></a>type Linger struct {
    <a id="L297"></a>Onoff  int32;
    <a id="L298"></a>Linger int32;
<a id="L299"></a>}

<a id="L301"></a>type PtraceRegs struct {
    <a id="L302"></a>Ebx      int32;
    <a id="L303"></a>Ecx      int32;
    <a id="L304"></a>Edx      int32;
    <a id="L305"></a>Esi      int32;
    <a id="L306"></a>Edi      int32;
    <a id="L307"></a>Ebp      int32;
    <a id="L308"></a>Eax      int32;
    <a id="L309"></a>Ds       uint16;
    <a id="L310"></a>__ds     uint16;
    <a id="L311"></a>Es       uint16;
    <a id="L312"></a>__es     uint16;
    <a id="L313"></a>Fs       uint16;
    <a id="L314"></a>__fs     uint16;
    <a id="L315"></a>Gs       uint16;
    <a id="L316"></a>__gs     uint16;
    <a id="L317"></a>Orig_eax int32;
    <a id="L318"></a>Eip      int32;
    <a id="L319"></a>Cs       uint16;
    <a id="L320"></a>__cs     uint16;
    <a id="L321"></a>Eflags   int32;
    <a id="L322"></a>Esp      int32;
    <a id="L323"></a>Ss       uint16;
    <a id="L324"></a>__ss     uint16;
<a id="L325"></a>}

<a id="L327"></a>type FdSet struct {
    <a id="L328"></a>Bits [32]int32;
<a id="L329"></a>}

<a id="L331"></a>type Sysinfo_t struct {
    <a id="L332"></a>Uptime    int32;
    <a id="L333"></a>Loads     [3]uint32;
    <a id="L334"></a>Totalram  uint32;
    <a id="L335"></a>Freeram   uint32;
    <a id="L336"></a>Sharedram uint32;
    <a id="L337"></a>Bufferram uint32;
    <a id="L338"></a>Totalswap uint32;
    <a id="L339"></a>Freeswap  uint32;
    <a id="L340"></a>Procs     uint16;
    <a id="L341"></a>Pad       uint16;
    <a id="L342"></a>Totalhigh uint32;
    <a id="L343"></a>Freehigh  uint32;
    <a id="L344"></a>Unit      uint32;
    <a id="L345"></a>_f        [8]int8;
<a id="L346"></a>}

<a id="L348"></a>type Utsname struct {
    <a id="L349"></a>Sysname    [65]int8;
    <a id="L350"></a>Nodename   [65]int8;
    <a id="L351"></a>Release    [65]int8;
    <a id="L352"></a>Version    [65]int8;
    <a id="L353"></a>Machine    [65]int8;
    <a id="L354"></a>Domainname [65]int8;
<a id="L355"></a>}

<a id="L357"></a>type Ustat_t struct {
    <a id="L358"></a>Tfree  int32;
    <a id="L359"></a>Tinode uint32;
    <a id="L360"></a>Fname  [6]int8;
    <a id="L361"></a>Fpack  [6]int8;
<a id="L362"></a>}

<a id="L364"></a>type EpollEvent struct {
    <a id="L365"></a>Events uint32;
    <a id="L366"></a>Fd     int32;
    <a id="L367"></a>Pad    int32;
<a id="L368"></a>}
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
