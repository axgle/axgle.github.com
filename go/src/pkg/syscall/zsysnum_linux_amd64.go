<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/zsysnum_linux_amd64.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/syscall/zsysnum_linux_amd64.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// mksysnum_linux.sh /usr/include/asm/unistd_64.h</span>
<a id="L2"></a><span class="comment">// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT</span>

<a id="L4"></a>package syscall

<a id="L6"></a>const (
    <a id="L7"></a>SYS_READ                   = 0;
    <a id="L8"></a>SYS_WRITE                  = 1;
    <a id="L9"></a>SYS_OPEN                   = 2;
    <a id="L10"></a>SYS_CLOSE                  = 3;
    <a id="L11"></a>SYS_STAT                   = 4;
    <a id="L12"></a>SYS_FSTAT                  = 5;
    <a id="L13"></a>SYS_LSTAT                  = 6;
    <a id="L14"></a>SYS_POLL                   = 7;
    <a id="L15"></a>SYS_LSEEK                  = 8;
    <a id="L16"></a>SYS_MMAP                   = 9;
    <a id="L17"></a>SYS_MPROTECT               = 10;
    <a id="L18"></a>SYS_MUNMAP                 = 11;
    <a id="L19"></a>SYS_BRK                    = 12;
    <a id="L20"></a>SYS_RT_SIGACTION           = 13;
    <a id="L21"></a>SYS_RT_SIGPROCMASK         = 14;
    <a id="L22"></a>SYS_RT_SIGRETURN           = 15;
    <a id="L23"></a>SYS_IOCTL                  = 16;
    <a id="L24"></a>SYS_PREAD64                = 17;
    <a id="L25"></a>SYS_PWRITE64               = 18;
    <a id="L26"></a>SYS_READV                  = 19;
    <a id="L27"></a>SYS_WRITEV                 = 20;
    <a id="L28"></a>SYS_ACCESS                 = 21;
    <a id="L29"></a>SYS_PIPE                   = 22;
    <a id="L30"></a>SYS_SELECT                 = 23;
    <a id="L31"></a>SYS_SCHED_YIELD            = 24;
    <a id="L32"></a>SYS_MREMAP                 = 25;
    <a id="L33"></a>SYS_MSYNC                  = 26;
    <a id="L34"></a>SYS_MINCORE                = 27;
    <a id="L35"></a>SYS_MADVISE                = 28;
    <a id="L36"></a>SYS_SHMGET                 = 29;
    <a id="L37"></a>SYS_SHMAT                  = 30;
    <a id="L38"></a>SYS_SHMCTL                 = 31;
    <a id="L39"></a>SYS_DUP                    = 32;
    <a id="L40"></a>SYS_DUP2                   = 33;
    <a id="L41"></a>SYS_PAUSE                  = 34;
    <a id="L42"></a>SYS_NANOSLEEP              = 35;
    <a id="L43"></a>SYS_GETITIMER              = 36;
    <a id="L44"></a>SYS_ALARM                  = 37;
    <a id="L45"></a>SYS_SETITIMER              = 38;
    <a id="L46"></a>SYS_GETPID                 = 39;
    <a id="L47"></a>SYS_SENDFILE               = 40;
    <a id="L48"></a>SYS_SOCKET                 = 41;
    <a id="L49"></a>SYS_CONNECT                = 42;
    <a id="L50"></a>SYS_ACCEPT                 = 43;
    <a id="L51"></a>SYS_SENDTO                 = 44;
    <a id="L52"></a>SYS_RECVFROM               = 45;
    <a id="L53"></a>SYS_SENDMSG                = 46;
    <a id="L54"></a>SYS_RECVMSG                = 47;
    <a id="L55"></a>SYS_SHUTDOWN               = 48;
    <a id="L56"></a>SYS_BIND                   = 49;
    <a id="L57"></a>SYS_LISTEN                 = 50;
    <a id="L58"></a>SYS_GETSOCKNAME            = 51;
    <a id="L59"></a>SYS_GETPEERNAME            = 52;
    <a id="L60"></a>SYS_SOCKETPAIR             = 53;
    <a id="L61"></a>SYS_SETSOCKOPT             = 54;
    <a id="L62"></a>SYS_GETSOCKOPT             = 55;
    <a id="L63"></a>SYS_CLONE                  = 56;
    <a id="L64"></a>SYS_FORK                   = 57;
    <a id="L65"></a>SYS_VFORK                  = 58;
    <a id="L66"></a>SYS_EXECVE                 = 59;
    <a id="L67"></a>SYS_EXIT                   = 60;
    <a id="L68"></a>SYS_WAIT4                  = 61;
    <a id="L69"></a>SYS_KILL                   = 62;
    <a id="L70"></a>SYS_UNAME                  = 63;
    <a id="L71"></a>SYS_SEMGET                 = 64;
    <a id="L72"></a>SYS_SEMOP                  = 65;
    <a id="L73"></a>SYS_SEMCTL                 = 66;
    <a id="L74"></a>SYS_SHMDT                  = 67;
    <a id="L75"></a>SYS_MSGGET                 = 68;
    <a id="L76"></a>SYS_MSGSND                 = 69;
    <a id="L77"></a>SYS_MSGRCV                 = 70;
    <a id="L78"></a>SYS_MSGCTL                 = 71;
    <a id="L79"></a>SYS_FCNTL                  = 72;
    <a id="L80"></a>SYS_FLOCK                  = 73;
    <a id="L81"></a>SYS_FSYNC                  = 74;
    <a id="L82"></a>SYS_FDATASYNC              = 75;
    <a id="L83"></a>SYS_TRUNCATE               = 76;
    <a id="L84"></a>SYS_FTRUNCATE              = 77;
    <a id="L85"></a>SYS_GETDENTS               = 78;
    <a id="L86"></a>SYS_GETCWD                 = 79;
    <a id="L87"></a>SYS_CHDIR                  = 80;
    <a id="L88"></a>SYS_FCHDIR                 = 81;
    <a id="L89"></a>SYS_RENAME                 = 82;
    <a id="L90"></a>SYS_MKDIR                  = 83;
    <a id="L91"></a>SYS_RMDIR                  = 84;
    <a id="L92"></a>SYS_CREAT                  = 85;
    <a id="L93"></a>SYS_LINK                   = 86;
    <a id="L94"></a>SYS_UNLINK                 = 87;
    <a id="L95"></a>SYS_SYMLINK                = 88;
    <a id="L96"></a>SYS_READLINK               = 89;
    <a id="L97"></a>SYS_CHMOD                  = 90;
    <a id="L98"></a>SYS_FCHMOD                 = 91;
    <a id="L99"></a>SYS_CHOWN                  = 92;
    <a id="L100"></a>SYS_FCHOWN                 = 93;
    <a id="L101"></a>SYS_LCHOWN                 = 94;
    <a id="L102"></a>SYS_UMASK                  = 95;
    <a id="L103"></a>SYS_GETTIMEOFDAY           = 96;
    <a id="L104"></a>SYS_GETRLIMIT              = 97;
    <a id="L105"></a>SYS_GETRUSAGE              = 98;
    <a id="L106"></a>SYS_SYSINFO                = 99;
    <a id="L107"></a>SYS_TIMES                  = 100;
    <a id="L108"></a>SYS_PTRACE                 = 101;
    <a id="L109"></a>SYS_GETUID                 = 102;
    <a id="L110"></a>SYS_SYSLOG                 = 103;
    <a id="L111"></a>SYS_GETGID                 = 104;
    <a id="L112"></a>SYS_SETUID                 = 105;
    <a id="L113"></a>SYS_SETGID                 = 106;
    <a id="L114"></a>SYS_GETEUID                = 107;
    <a id="L115"></a>SYS_GETEGID                = 108;
    <a id="L116"></a>SYS_SETPGID                = 109;
    <a id="L117"></a>SYS_GETPPID                = 110;
    <a id="L118"></a>SYS_GETPGRP                = 111;
    <a id="L119"></a>SYS_SETSID                 = 112;
    <a id="L120"></a>SYS_SETREUID               = 113;
    <a id="L121"></a>SYS_SETREGID               = 114;
    <a id="L122"></a>SYS_GETGROUPS              = 115;
    <a id="L123"></a>SYS_SETGROUPS              = 116;
    <a id="L124"></a>SYS_SETRESUID              = 117;
    <a id="L125"></a>SYS_GETRESUID              = 118;
    <a id="L126"></a>SYS_SETRESGID              = 119;
    <a id="L127"></a>SYS_GETRESGID              = 120;
    <a id="L128"></a>SYS_GETPGID                = 121;
    <a id="L129"></a>SYS_SETFSUID               = 122;
    <a id="L130"></a>SYS_SETFSGID               = 123;
    <a id="L131"></a>SYS_GETSID                 = 124;
    <a id="L132"></a>SYS_CAPGET                 = 125;
    <a id="L133"></a>SYS_CAPSET                 = 126;
    <a id="L134"></a>SYS_RT_SIGPENDING          = 127;
    <a id="L135"></a>SYS_RT_SIGTIMEDWAIT        = 128;
    <a id="L136"></a>SYS_RT_SIGQUEUEINFO        = 129;
    <a id="L137"></a>SYS_RT_SIGSUSPEND          = 130;
    <a id="L138"></a>SYS_SIGALTSTACK            = 131;
    <a id="L139"></a>SYS_UTIME                  = 132;
    <a id="L140"></a>SYS_MKNOD                  = 133;
    <a id="L141"></a>SYS_USELIB                 = 134;
    <a id="L142"></a>SYS_PERSONALITY            = 135;
    <a id="L143"></a>SYS_USTAT                  = 136;
    <a id="L144"></a>SYS_STATFS                 = 137;
    <a id="L145"></a>SYS_FSTATFS                = 138;
    <a id="L146"></a>SYS_SYSFS                  = 139;
    <a id="L147"></a>SYS_GETPRIORITY            = 140;
    <a id="L148"></a>SYS_SETPRIORITY            = 141;
    <a id="L149"></a>SYS_SCHED_SETPARAM         = 142;
    <a id="L150"></a>SYS_SCHED_GETPARAM         = 143;
    <a id="L151"></a>SYS_SCHED_SETSCHEDULER     = 144;
    <a id="L152"></a>SYS_SCHED_GETSCHEDULER     = 145;
    <a id="L153"></a>SYS_SCHED_GET_PRIORITY_MAX = 146;
    <a id="L154"></a>SYS_SCHED_GET_PRIORITY_MIN = 147;
    <a id="L155"></a>SYS_SCHED_RR_GET_INTERVAL  = 148;
    <a id="L156"></a>SYS_MLOCK                  = 149;
    <a id="L157"></a>SYS_MUNLOCK                = 150;
    <a id="L158"></a>SYS_MLOCKALL               = 151;
    <a id="L159"></a>SYS_MUNLOCKALL             = 152;
    <a id="L160"></a>SYS_VHANGUP                = 153;
    <a id="L161"></a>SYS_MODIFY_LDT             = 154;
    <a id="L162"></a>SYS_PIVOT_ROOT             = 155;
    <a id="L163"></a>SYS__SYSCTL                = 156;
    <a id="L164"></a>SYS_PRCTL                  = 157;
    <a id="L165"></a>SYS_ARCH_PRCTL             = 158;
    <a id="L166"></a>SYS_ADJTIMEX               = 159;
    <a id="L167"></a>SYS_SETRLIMIT              = 160;
    <a id="L168"></a>SYS_CHROOT                 = 161;
    <a id="L169"></a>SYS_SYNC                   = 162;
    <a id="L170"></a>SYS_ACCT                   = 163;
    <a id="L171"></a>SYS_SETTIMEOFDAY           = 164;
    <a id="L172"></a>SYS_MOUNT                  = 165;
    <a id="L173"></a>SYS_UMOUNT2                = 166;
    <a id="L174"></a>SYS_SWAPON                 = 167;
    <a id="L175"></a>SYS_SWAPOFF                = 168;
    <a id="L176"></a>SYS_REBOOT                 = 169;
    <a id="L177"></a>SYS_SETHOSTNAME            = 170;
    <a id="L178"></a>SYS_SETDOMAINNAME          = 171;
    <a id="L179"></a>SYS_IOPL                   = 172;
    <a id="L180"></a>SYS_IOPERM                 = 173;
    <a id="L181"></a>SYS_CREATE_MODULE          = 174;
    <a id="L182"></a>SYS_INIT_MODULE            = 175;
    <a id="L183"></a>SYS_DELETE_MODULE          = 176;
    <a id="L184"></a>SYS_GET_KERNEL_SYMS        = 177;
    <a id="L185"></a>SYS_QUERY_MODULE           = 178;
    <a id="L186"></a>SYS_QUOTACTL               = 179;
    <a id="L187"></a>SYS_NFSSERVCTL             = 180;
    <a id="L188"></a>SYS_GETPMSG                = 181;
    <a id="L189"></a>SYS_PUTPMSG                = 182;
    <a id="L190"></a>SYS_AFS_SYSCALL            = 183;
    <a id="L191"></a>SYS_TUXCALL                = 184;
    <a id="L192"></a>SYS_SECURITY               = 185;
    <a id="L193"></a>SYS_GETTID                 = 186;
    <a id="L194"></a>SYS_READAHEAD              = 187;
    <a id="L195"></a>SYS_SETXATTR               = 188;
    <a id="L196"></a>SYS_LSETXATTR              = 189;
    <a id="L197"></a>SYS_FSETXATTR              = 190;
    <a id="L198"></a>SYS_GETXATTR               = 191;
    <a id="L199"></a>SYS_LGETXATTR              = 192;
    <a id="L200"></a>SYS_FGETXATTR              = 193;
    <a id="L201"></a>SYS_LISTXATTR              = 194;
    <a id="L202"></a>SYS_LLISTXATTR             = 195;
    <a id="L203"></a>SYS_FLISTXATTR             = 196;
    <a id="L204"></a>SYS_REMOVEXATTR            = 197;
    <a id="L205"></a>SYS_LREMOVEXATTR           = 198;
    <a id="L206"></a>SYS_FREMOVEXATTR           = 199;
    <a id="L207"></a>SYS_TKILL                  = 200;
    <a id="L208"></a>SYS_TIME                   = 201;
    <a id="L209"></a>SYS_FUTEX                  = 202;
    <a id="L210"></a>SYS_SCHED_SETAFFINITY      = 203;
    <a id="L211"></a>SYS_SCHED_GETAFFINITY      = 204;
    <a id="L212"></a>SYS_SET_THREAD_AREA        = 205;
    <a id="L213"></a>SYS_IO_SETUP               = 206;
    <a id="L214"></a>SYS_IO_DESTROY             = 207;
    <a id="L215"></a>SYS_IO_GETEVENTS           = 208;
    <a id="L216"></a>SYS_IO_SUBMIT              = 209;
    <a id="L217"></a>SYS_IO_CANCEL              = 210;
    <a id="L218"></a>SYS_GET_THREAD_AREA        = 211;
    <a id="L219"></a>SYS_LOOKUP_DCOOKIE         = 212;
    <a id="L220"></a>SYS_EPOLL_CREATE           = 213;
    <a id="L221"></a>SYS_EPOLL_CTL_OLD          = 214;
    <a id="L222"></a>SYS_EPOLL_WAIT_OLD         = 215;
    <a id="L223"></a>SYS_REMAP_FILE_PAGES       = 216;
    <a id="L224"></a>SYS_GETDENTS64             = 217;
    <a id="L225"></a>SYS_SET_TID_ADDRESS        = 218;
    <a id="L226"></a>SYS_RESTART_SYSCALL        = 219;
    <a id="L227"></a>SYS_SEMTIMEDOP             = 220;
    <a id="L228"></a>SYS_FADVISE64              = 221;
    <a id="L229"></a>SYS_TIMER_CREATE           = 222;
    <a id="L230"></a>SYS_TIMER_SETTIME          = 223;
    <a id="L231"></a>SYS_TIMER_GETTIME          = 224;
    <a id="L232"></a>SYS_TIMER_GETOVERRUN       = 225;
    <a id="L233"></a>SYS_TIMER_DELETE           = 226;
    <a id="L234"></a>SYS_CLOCK_SETTIME          = 227;
    <a id="L235"></a>SYS_CLOCK_GETTIME          = 228;
    <a id="L236"></a>SYS_CLOCK_GETRES           = 229;
    <a id="L237"></a>SYS_CLOCK_NANOSLEEP        = 230;
    <a id="L238"></a>SYS_EXIT_GROUP             = 231;
    <a id="L239"></a>SYS_EPOLL_WAIT             = 232;
    <a id="L240"></a>SYS_EPOLL_CTL              = 233;
    <a id="L241"></a>SYS_TGKILL                 = 234;
    <a id="L242"></a>SYS_UTIMES                 = 235;
    <a id="L243"></a>SYS_VSERVER                = 236;
    <a id="L244"></a>SYS_MBIND                  = 237;
    <a id="L245"></a>SYS_SET_MEMPOLICY          = 238;
    <a id="L246"></a>SYS_GET_MEMPOLICY          = 239;
    <a id="L247"></a>SYS_MQ_OPEN                = 240;
    <a id="L248"></a>SYS_MQ_UNLINK              = 241;
    <a id="L249"></a>SYS_MQ_TIMEDSEND           = 242;
    <a id="L250"></a>SYS_MQ_TIMEDRECEIVE        = 243;
    <a id="L251"></a>SYS_MQ_NOTIFY              = 244;
    <a id="L252"></a>SYS_MQ_GETSETATTR          = 245;
    <a id="L253"></a>SYS_KEXEC_LOAD             = 246;
    <a id="L254"></a>SYS_WAITID                 = 247;
    <a id="L255"></a>SYS_ADD_KEY                = 248;
    <a id="L256"></a>SYS_REQUEST_KEY            = 249;
    <a id="L257"></a>SYS_KEYCTL                 = 250;
    <a id="L258"></a>SYS_IOPRIO_SET             = 251;
    <a id="L259"></a>SYS_IOPRIO_GET             = 252;
    <a id="L260"></a>SYS_INOTIFY_INIT           = 253;
    <a id="L261"></a>SYS_INOTIFY_ADD_WATCH      = 254;
    <a id="L262"></a>SYS_INOTIFY_RM_WATCH       = 255;
    <a id="L263"></a>SYS_MIGRATE_PAGES          = 256;
    <a id="L264"></a>SYS_OPENAT                 = 257;
    <a id="L265"></a>SYS_MKDIRAT                = 258;
    <a id="L266"></a>SYS_MKNODAT                = 259;
    <a id="L267"></a>SYS_FCHOWNAT               = 260;
    <a id="L268"></a>SYS_FUTIMESAT              = 261;
    <a id="L269"></a>SYS_NEWFSTATAT             = 262;
    <a id="L270"></a>SYS_UNLINKAT               = 263;
    <a id="L271"></a>SYS_RENAMEAT               = 264;
    <a id="L272"></a>SYS_LINKAT                 = 265;
    <a id="L273"></a>SYS_SYMLINKAT              = 266;
    <a id="L274"></a>SYS_READLINKAT             = 267;
    <a id="L275"></a>SYS_FCHMODAT               = 268;
    <a id="L276"></a>SYS_FACCESSAT              = 269;
    <a id="L277"></a>SYS_PSELECT6               = 270;
    <a id="L278"></a>SYS_PPOLL                  = 271;
    <a id="L279"></a>SYS_UNSHARE                = 272;
    <a id="L280"></a>SYS_SET_ROBUST_LIST        = 273;
    <a id="L281"></a>SYS_GET_ROBUST_LIST        = 274;
    <a id="L282"></a>SYS_SPLICE                 = 275;
    <a id="L283"></a>SYS_TEE                    = 276;
    <a id="L284"></a>SYS_SYNC_FILE_RANGE        = 277;
    <a id="L285"></a>SYS_VMSPLICE               = 278;
    <a id="L286"></a>SYS_MOVE_PAGES             = 279;
    <a id="L287"></a>SYS_UTIMENSAT              = 280;
    <a id="L288"></a>SYS_EPOLL_PWAIT            = 281;
    <a id="L289"></a>SYS_SIGNALFD               = 282;
    <a id="L290"></a>SYS_TIMERFD                = 283;
    <a id="L291"></a>SYS_EVENTFD                = 284;
    <a id="L292"></a>SYS_FALLOCATE              = 285;
<a id="L293"></a>)

<a id="L295"></a>func _darwin_system_call_conflict() {}
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
