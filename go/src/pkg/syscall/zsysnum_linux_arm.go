<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/zsysnum_linux_arm.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/syscall/zsysnum_linux_arm.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// hand generated</span>

<a id="L3"></a>package syscall

<a id="L5"></a>const (
    <a id="L6"></a>SYS_SYSCALL_BASE = 0;

    <a id="L8"></a>SYS_RESTART_SYSCALL        = (SYS_SYSCALL_BASE + 0);
    <a id="L9"></a>SYS_EXIT                   = (SYS_SYSCALL_BASE + 1);
    <a id="L10"></a>SYS_FORK                   = (SYS_SYSCALL_BASE + 2);
    <a id="L11"></a>SYS_READ                   = (SYS_SYSCALL_BASE + 3);
    <a id="L12"></a>SYS_WRITE                  = (SYS_SYSCALL_BASE + 4);
    <a id="L13"></a>SYS_OPEN                   = (SYS_SYSCALL_BASE + 5);
    <a id="L14"></a>SYS_CLOSE                  = (SYS_SYSCALL_BASE + 6);
    <a id="L15"></a>SYS_CREAT                  = (SYS_SYSCALL_BASE + 8);
    <a id="L16"></a>SYS_LINK                   = (SYS_SYSCALL_BASE + 9);
    <a id="L17"></a>SYS_UNLINK                 = (SYS_SYSCALL_BASE + 10);
    <a id="L18"></a>SYS_EXECVE                 = (SYS_SYSCALL_BASE + 11);
    <a id="L19"></a>SYS_CHDIR                  = (SYS_SYSCALL_BASE + 12);
    <a id="L20"></a>SYS_TIME                   = (SYS_SYSCALL_BASE + 13);
    <a id="L21"></a>SYS_MKNOD                  = (SYS_SYSCALL_BASE + 14);
    <a id="L22"></a>SYS_CHMOD                  = (SYS_SYSCALL_BASE + 15);
    <a id="L23"></a>SYS_LCHOWN                 = (SYS_SYSCALL_BASE + 16);
    <a id="L24"></a>SYS_LSEEK                  = (SYS_SYSCALL_BASE + 19);
    <a id="L25"></a>SYS_GETPID                 = (SYS_SYSCALL_BASE + 20);
    <a id="L26"></a>SYS_MOUNT                  = (SYS_SYSCALL_BASE + 21);
    <a id="L27"></a>SYS_UMOUNT                 = (SYS_SYSCALL_BASE + 22);
    <a id="L28"></a>SYS_SETUID                 = (SYS_SYSCALL_BASE + 23);
    <a id="L29"></a>SYS_GETUID                 = (SYS_SYSCALL_BASE + 24);
    <a id="L30"></a>SYS_STIME                  = (SYS_SYSCALL_BASE + 25);
    <a id="L31"></a>SYS_PTRACE                 = (SYS_SYSCALL_BASE + 26);
    <a id="L32"></a>SYS_ALARM                  = (SYS_SYSCALL_BASE + 27);
    <a id="L33"></a>SYS_PAUSE                  = (SYS_SYSCALL_BASE + 29);
    <a id="L34"></a>SYS_UTIME                  = (SYS_SYSCALL_BASE + 30);
    <a id="L35"></a>SYS_ACCESS                 = (SYS_SYSCALL_BASE + 33);
    <a id="L36"></a>SYS_NICE                   = (SYS_SYSCALL_BASE + 34);
    <a id="L37"></a>SYS_SYNC                   = (SYS_SYSCALL_BASE + 36);
    <a id="L38"></a>SYS_KILL                   = (SYS_SYSCALL_BASE + 37);
    <a id="L39"></a>SYS_RENAME                 = (SYS_SYSCALL_BASE + 38);
    <a id="L40"></a>SYS_MKDIR                  = (SYS_SYSCALL_BASE + 39);
    <a id="L41"></a>SYS_RMDIR                  = (SYS_SYSCALL_BASE + 40);
    <a id="L42"></a>SYS_DUP                    = (SYS_SYSCALL_BASE + 41);
    <a id="L43"></a>SYS_PIPE                   = (SYS_SYSCALL_BASE + 42);
    <a id="L44"></a>SYS_TIMES                  = (SYS_SYSCALL_BASE + 43);
    <a id="L45"></a>SYS_BRK                    = (SYS_SYSCALL_BASE + 45);
    <a id="L46"></a>SYS_SETGID                 = (SYS_SYSCALL_BASE + 46);
    <a id="L47"></a>SYS_GETGID                 = (SYS_SYSCALL_BASE + 47);
    <a id="L48"></a>SYS_GETEUID                = (SYS_SYSCALL_BASE + 49);
    <a id="L49"></a>SYS_GETEGID                = (SYS_SYSCALL_BASE + 50);
    <a id="L50"></a>SYS_ACCT                   = (SYS_SYSCALL_BASE + 51);
    <a id="L51"></a>SYS_UMOUNT2                = (SYS_SYSCALL_BASE + 52);
    <a id="L52"></a>SYS_IOCTL                  = (SYS_SYSCALL_BASE + 54);
    <a id="L53"></a>SYS_FCNTL                  = (SYS_SYSCALL_BASE + 55);
    <a id="L54"></a>SYS_SETPGID                = (SYS_SYSCALL_BASE + 57);
    <a id="L55"></a>SYS_UMASK                  = (SYS_SYSCALL_BASE + 60);
    <a id="L56"></a>SYS_CHROOT                 = (SYS_SYSCALL_BASE + 61);
    <a id="L57"></a>SYS_USTAT                  = (SYS_SYSCALL_BASE + 62);
    <a id="L58"></a>SYS_DUP2                   = (SYS_SYSCALL_BASE + 63);
    <a id="L59"></a>SYS_GETPPID                = (SYS_SYSCALL_BASE + 64);
    <a id="L60"></a>SYS_GETPGRP                = (SYS_SYSCALL_BASE + 65);
    <a id="L61"></a>SYS_SETSID                 = (SYS_SYSCALL_BASE + 66);
    <a id="L62"></a>SYS_SIGACTION              = (SYS_SYSCALL_BASE + 67);
    <a id="L63"></a>SYS_SETREUID               = (SYS_SYSCALL_BASE + 70);
    <a id="L64"></a>SYS_SETREGID               = (SYS_SYSCALL_BASE + 71);
    <a id="L65"></a>SYS_SIGSUSPEND             = (SYS_SYSCALL_BASE + 72);
    <a id="L66"></a>SYS_SIGPENDING             = (SYS_SYSCALL_BASE + 73);
    <a id="L67"></a>SYS_SETHOSTNAME            = (SYS_SYSCALL_BASE + 74);
    <a id="L68"></a>SYS_SETRLIMIT              = (SYS_SYSCALL_BASE + 75);
    <a id="L69"></a>SYS_GETRLIMIT              = (SYS_SYSCALL_BASE + 76);
    <a id="L70"></a>SYS_GETRUSAGE              = (SYS_SYSCALL_BASE + 77);
    <a id="L71"></a>SYS_GETTIMEOFDAY           = (SYS_SYSCALL_BASE + 78);
    <a id="L72"></a>SYS_SETTIMEOFDAY           = (SYS_SYSCALL_BASE + 79);
    <a id="L73"></a>SYS_GETGROUPS              = (SYS_SYSCALL_BASE + 80);
    <a id="L74"></a>SYS_SETGROUPS              = (SYS_SYSCALL_BASE + 81);
    <a id="L75"></a>SYS_SELECT                 = (SYS_SYSCALL_BASE + 82);
    <a id="L76"></a>SYS_SYMLINK                = (SYS_SYSCALL_BASE + 83);
    <a id="L77"></a>SYS_READLINK               = (SYS_SYSCALL_BASE + 85);
    <a id="L78"></a>SYS_USELIB                 = (SYS_SYSCALL_BASE + 86);
    <a id="L79"></a>SYS_SWAPON                 = (SYS_SYSCALL_BASE + 87);
    <a id="L80"></a>SYS_REBOOT                 = (SYS_SYSCALL_BASE + 88);
    <a id="L81"></a>SYS_READDIR                = (SYS_SYSCALL_BASE + 89);
    <a id="L82"></a>SYS_MMAP                   = (SYS_SYSCALL_BASE + 90);
    <a id="L83"></a>SYS_MUNMAP                 = (SYS_SYSCALL_BASE + 91);
    <a id="L84"></a>SYS_TRUNCATE               = (SYS_SYSCALL_BASE + 92);
    <a id="L85"></a>SYS_FTRUNCATE              = (SYS_SYSCALL_BASE + 93);
    <a id="L86"></a>SYS_FCHMOD                 = (SYS_SYSCALL_BASE + 94);
    <a id="L87"></a>SYS_FCHOWN                 = (SYS_SYSCALL_BASE + 95);
    <a id="L88"></a>SYS_GETPRIORITY            = (SYS_SYSCALL_BASE + 96);
    <a id="L89"></a>SYS_SETPRIORITY            = (SYS_SYSCALL_BASE + 97);
    <a id="L90"></a>SYS_STATFS                 = (SYS_SYSCALL_BASE + 99);
    <a id="L91"></a>SYS_FSTATFS                = (SYS_SYSCALL_BASE + 100);
    <a id="L92"></a>SYS_SOCKETCALL             = (SYS_SYSCALL_BASE + 102);
    <a id="L93"></a>SYS_SYSLOG                 = (SYS_SYSCALL_BASE + 103);
    <a id="L94"></a>SYS_SETITIMER              = (SYS_SYSCALL_BASE + 104);
    <a id="L95"></a>SYS_GETITIMER              = (SYS_SYSCALL_BASE + 105);
    <a id="L96"></a>SYS_STAT                   = (SYS_SYSCALL_BASE + 106);
    <a id="L97"></a>SYS_LSTAT                  = (SYS_SYSCALL_BASE + 107);
    <a id="L98"></a>SYS_FSTAT                  = (SYS_SYSCALL_BASE + 108);
    <a id="L99"></a>SYS_VHANGUP                = (SYS_SYSCALL_BASE + 111);
    <a id="L100"></a>SYS_SYSCALL                = (SYS_SYSCALL_BASE + 113);
    <a id="L101"></a>SYS_WAIT4                  = (SYS_SYSCALL_BASE + 114);
    <a id="L102"></a>SYS_SWAPOFF                = (SYS_SYSCALL_BASE + 115);
    <a id="L103"></a>SYS_SYSINFO                = (SYS_SYSCALL_BASE + 116);
    <a id="L104"></a>SYS_IPC                    = (SYS_SYSCALL_BASE + 117);
    <a id="L105"></a>SYS_FSYNC                  = (SYS_SYSCALL_BASE + 118);
    <a id="L106"></a>SYS_SIGRETURN              = (SYS_SYSCALL_BASE + 119);
    <a id="L107"></a>SYS_CLONE                  = (SYS_SYSCALL_BASE + 120);
    <a id="L108"></a>SYS_SETDOMAINNAME          = (SYS_SYSCALL_BASE + 121);
    <a id="L109"></a>SYS_UNAME                  = (SYS_SYSCALL_BASE + 122);
    <a id="L110"></a>SYS_ADJTIMEX               = (SYS_SYSCALL_BASE + 124);
    <a id="L111"></a>SYS_MPROTECT               = (SYS_SYSCALL_BASE + 125);
    <a id="L112"></a>SYS_SIGPROCMASK            = (SYS_SYSCALL_BASE + 126);
    <a id="L113"></a>SYS_INIT_MODULE            = (SYS_SYSCALL_BASE + 128);
    <a id="L114"></a>SYS_DELETE_MODULE          = (SYS_SYSCALL_BASE + 129);
    <a id="L115"></a>SYS_QUOTACTL               = (SYS_SYSCALL_BASE + 131);
    <a id="L116"></a>SYS_GETPGID                = (SYS_SYSCALL_BASE + 132);
    <a id="L117"></a>SYS_FCHDIR                 = (SYS_SYSCALL_BASE + 133);
    <a id="L118"></a>SYS_BDFLUSH                = (SYS_SYSCALL_BASE + 134);
    <a id="L119"></a>SYS_SYSFS                  = (SYS_SYSCALL_BASE + 135);
    <a id="L120"></a>SYS_PERSONALITY            = (SYS_SYSCALL_BASE + 136);
    <a id="L121"></a>SYS_SETFSUID               = (SYS_SYSCALL_BASE + 138);
    <a id="L122"></a>SYS_SETFSGID               = (SYS_SYSCALL_BASE + 139);
    <a id="L123"></a>SYS__LLSEEK                = (SYS_SYSCALL_BASE + 140);
    <a id="L124"></a>SYS_GETDENTS               = (SYS_SYSCALL_BASE + 141);
    <a id="L125"></a>SYS__NEWSELECT             = (SYS_SYSCALL_BASE + 142);
    <a id="L126"></a>SYS_FLOCK                  = (SYS_SYSCALL_BASE + 143);
    <a id="L127"></a>SYS_MSYNC                  = (SYS_SYSCALL_BASE + 144);
    <a id="L128"></a>SYS_READV                  = (SYS_SYSCALL_BASE + 145);
    <a id="L129"></a>SYS_WRITEV                 = (SYS_SYSCALL_BASE + 146);
    <a id="L130"></a>SYS_GETSID                 = (SYS_SYSCALL_BASE + 147);
    <a id="L131"></a>SYS_FDATASYNC              = (SYS_SYSCALL_BASE + 148);
    <a id="L132"></a>SYS__SYSCTL                = (SYS_SYSCALL_BASE + 149);
    <a id="L133"></a>SYS_MLOCK                  = (SYS_SYSCALL_BASE + 150);
    <a id="L134"></a>SYS_MUNLOCK                = (SYS_SYSCALL_BASE + 151);
    <a id="L135"></a>SYS_MLOCKALL               = (SYS_SYSCALL_BASE + 152);
    <a id="L136"></a>SYS_MUNLOCKALL             = (SYS_SYSCALL_BASE + 153);
    <a id="L137"></a>SYS_SCHED_SETPARAM         = (SYS_SYSCALL_BASE + 154);
    <a id="L138"></a>SYS_SCHED_GETPARAM         = (SYS_SYSCALL_BASE + 155);
    <a id="L139"></a>SYS_SCHED_SETSCHEDULER     = (SYS_SYSCALL_BASE + 156);
    <a id="L140"></a>SYS_SCHED_GETSCHEDULER     = (SYS_SYSCALL_BASE + 157);
    <a id="L141"></a>SYS_SCHED_YIELD            = (SYS_SYSCALL_BASE + 158);
    <a id="L142"></a>SYS_SCHED_GET_PRIORITY_MAX = (SYS_SYSCALL_BASE + 159);
    <a id="L143"></a>SYS_SCHED_GET_PRIORITY_MIN = (SYS_SYSCALL_BASE + 160);
    <a id="L144"></a>SYS_SCHED_RR_GET_INTERVAL  = (SYS_SYSCALL_BASE + 161);
    <a id="L145"></a>SYS_NANOSLEEP              = (SYS_SYSCALL_BASE + 162);
    <a id="L146"></a>SYS_MREMAP                 = (SYS_SYSCALL_BASE + 163);
    <a id="L147"></a>SYS_SETRESUID              = (SYS_SYSCALL_BASE + 164);
    <a id="L148"></a>SYS_GETRESUID              = (SYS_SYSCALL_BASE + 165);
    <a id="L149"></a>SYS_POLL                   = (SYS_SYSCALL_BASE + 168);
    <a id="L150"></a>SYS_NFSSERVCTL             = (SYS_SYSCALL_BASE + 169);
    <a id="L151"></a>SYS_SETRESGID              = (SYS_SYSCALL_BASE + 170);
    <a id="L152"></a>SYS_GETRESGID              = (SYS_SYSCALL_BASE + 171);
    <a id="L153"></a>SYS_PRCTL                  = (SYS_SYSCALL_BASE + 172);
    <a id="L154"></a>SYS_RT_SIGRETURN           = (SYS_SYSCALL_BASE + 173);
    <a id="L155"></a>SYS_RT_SIGACTION           = (SYS_SYSCALL_BASE + 174);
    <a id="L156"></a>SYS_RT_SIGPROCMASK         = (SYS_SYSCALL_BASE + 175);
    <a id="L157"></a>SYS_RT_SIGPENDING          = (SYS_SYSCALL_BASE + 176);
    <a id="L158"></a>SYS_RT_SIGTIMEDWAIT        = (SYS_SYSCALL_BASE + 177);
    <a id="L159"></a>SYS_RT_SIGQUEUEINFO        = (SYS_SYSCALL_BASE + 178);
    <a id="L160"></a>SYS_RT_SIGSUSPEND          = (SYS_SYSCALL_BASE + 179);
    <a id="L161"></a>SYS_PREAD64                = (SYS_SYSCALL_BASE + 180);
    <a id="L162"></a>SYS_PWRITE64               = (SYS_SYSCALL_BASE + 181);
    <a id="L163"></a>SYS_CHOWN                  = (SYS_SYSCALL_BASE + 182);
    <a id="L164"></a>SYS_GETCWD                 = (SYS_SYSCALL_BASE + 183);
    <a id="L165"></a>SYS_CAPGET                 = (SYS_SYSCALL_BASE + 184);
    <a id="L166"></a>SYS_CAPSET                 = (SYS_SYSCALL_BASE + 185);
    <a id="L167"></a>SYS_SIGALTSTACK            = (SYS_SYSCALL_BASE + 186);
    <a id="L168"></a>SYS_SENDFILE               = (SYS_SYSCALL_BASE + 187);
    <a id="L169"></a>SYS_VFORK                  = (SYS_SYSCALL_BASE + 190);
    <a id="L170"></a>SYS_UGETRLIMIT             = (SYS_SYSCALL_BASE + 191);
    <a id="L171"></a>SYS_MMAP2                  = (SYS_SYSCALL_BASE + 192);
    <a id="L172"></a>SYS_TRUNCATE64             = (SYS_SYSCALL_BASE + 193);
    <a id="L173"></a>SYS_FTRUNCATE64            = (SYS_SYSCALL_BASE + 194);
    <a id="L174"></a>SYS_STAT64                 = (SYS_SYSCALL_BASE + 195);
    <a id="L175"></a>SYS_LSTAT64                = (SYS_SYSCALL_BASE + 196);
    <a id="L176"></a>SYS_FSTAT64                = (SYS_SYSCALL_BASE + 197);
    <a id="L177"></a>SYS_LCHOWN32               = (SYS_SYSCALL_BASE + 198);
    <a id="L178"></a>SYS_GETUID32               = (SYS_SYSCALL_BASE + 199);
    <a id="L179"></a>SYS_GETGID32               = (SYS_SYSCALL_BASE + 200);
    <a id="L180"></a>SYS_GETEUID32              = (SYS_SYSCALL_BASE + 201);
    <a id="L181"></a>SYS_GETEGID32              = (SYS_SYSCALL_BASE + 202);
    <a id="L182"></a>SYS_SETREUID32             = (SYS_SYSCALL_BASE + 203);
    <a id="L183"></a>SYS_SETREGID32             = (SYS_SYSCALL_BASE + 204);
    <a id="L184"></a>SYS_GETGROUPS32            = (SYS_SYSCALL_BASE + 205);
    <a id="L185"></a>SYS_SETGROUPS32            = (SYS_SYSCALL_BASE + 206);
    <a id="L186"></a>SYS_FCHOWN32               = (SYS_SYSCALL_BASE + 207);
    <a id="L187"></a>SYS_SETRESUID32            = (SYS_SYSCALL_BASE + 208);
    <a id="L188"></a>SYS_GETRESUID32            = (SYS_SYSCALL_BASE + 209);
    <a id="L189"></a>SYS_SETRESGID32            = (SYS_SYSCALL_BASE + 210);
    <a id="L190"></a>SYS_GETRESGID32            = (SYS_SYSCALL_BASE + 211);
    <a id="L191"></a>SYS_CHOWN32                = (SYS_SYSCALL_BASE + 212);
    <a id="L192"></a>SYS_SETUID32               = (SYS_SYSCALL_BASE + 213);
    <a id="L193"></a>SYS_SETGID32               = (SYS_SYSCALL_BASE + 214);
    <a id="L194"></a>SYS_SETFSUID32             = (SYS_SYSCALL_BASE + 215);
    <a id="L195"></a>SYS_SETFSGID32             = (SYS_SYSCALL_BASE + 216);
    <a id="L196"></a>SYS_GETDENTS64             = (SYS_SYSCALL_BASE + 217);
    <a id="L197"></a>SYS_PIVOT_ROOT             = (SYS_SYSCALL_BASE + 218);
    <a id="L198"></a>SYS_MINCORE                = (SYS_SYSCALL_BASE + 219);
    <a id="L199"></a>SYS_MADVISE                = (SYS_SYSCALL_BASE + 220);
    <a id="L200"></a>SYS_FCNTL64                = (SYS_SYSCALL_BASE + 221);
    <a id="L201"></a>SYS_GETTID                 = (SYS_SYSCALL_BASE + 224);
    <a id="L202"></a>SYS_READAHEAD              = (SYS_SYSCALL_BASE + 225);
    <a id="L203"></a>SYS_SETXATTR               = (SYS_SYSCALL_BASE + 226);
    <a id="L204"></a>SYS_LSETXATTR              = (SYS_SYSCALL_BASE + 227);
    <a id="L205"></a>SYS_FSETXATTR              = (SYS_SYSCALL_BASE + 228);
    <a id="L206"></a>SYS_GETXATTR               = (SYS_SYSCALL_BASE + 229);
    <a id="L207"></a>SYS_LGETXATTR              = (SYS_SYSCALL_BASE + 230);
    <a id="L208"></a>SYS_FGETXATTR              = (SYS_SYSCALL_BASE + 231);
    <a id="L209"></a>SYS_LISTXATTR              = (SYS_SYSCALL_BASE + 232);
    <a id="L210"></a>SYS_LLISTXATTR             = (SYS_SYSCALL_BASE + 233);
    <a id="L211"></a>SYS_FLISTXATTR             = (SYS_SYSCALL_BASE + 234);
    <a id="L212"></a>SYS_REMOVEXATTR            = (SYS_SYSCALL_BASE + 235);
    <a id="L213"></a>SYS_LREMOVEXATTR           = (SYS_SYSCALL_BASE + 236);
    <a id="L214"></a>SYS_FREMOVEXATTR           = (SYS_SYSCALL_BASE + 237);
    <a id="L215"></a>SYS_TKILL                  = (SYS_SYSCALL_BASE + 238);
    <a id="L216"></a>SYS_SENDFILE64             = (SYS_SYSCALL_BASE + 239);
    <a id="L217"></a>SYS_FUTEX                  = (SYS_SYSCALL_BASE + 240);
    <a id="L218"></a>SYS_SCHED_SETAFFINITY      = (SYS_SYSCALL_BASE + 241);
    <a id="L219"></a>SYS_SCHED_GETAFFINITY      = (SYS_SYSCALL_BASE + 242);
    <a id="L220"></a>SYS_IO_SETUP               = (SYS_SYSCALL_BASE + 243);
    <a id="L221"></a>SYS_IO_DESTROY             = (SYS_SYSCALL_BASE + 244);
    <a id="L222"></a>SYS_IO_GETEVENTS           = (SYS_SYSCALL_BASE + 245);
    <a id="L223"></a>SYS_IO_SUBMIT              = (SYS_SYSCALL_BASE + 246);
    <a id="L224"></a>SYS_IO_CANCEL              = (SYS_SYSCALL_BASE + 247);
    <a id="L225"></a>SYS_EXIT_GROUP             = (SYS_SYSCALL_BASE + 248);
    <a id="L226"></a>SYS_LOOKUP_DCOOKIE         = (SYS_SYSCALL_BASE + 249);
    <a id="L227"></a>SYS_EPOLL_CREATE           = (SYS_SYSCALL_BASE + 250);
    <a id="L228"></a>SYS_EPOLL_CTL              = (SYS_SYSCALL_BASE + 251);
    <a id="L229"></a>SYS_EPOLL_WAIT             = (SYS_SYSCALL_BASE + 252);
    <a id="L230"></a>SYS_REMAP_FILE_PAGES       = (SYS_SYSCALL_BASE + 253);
    <a id="L231"></a>SYS_SET_TID_ADDRESS        = (SYS_SYSCALL_BASE + 256);
    <a id="L232"></a>SYS_TIMER_CREATE           = (SYS_SYSCALL_BASE + 257);
    <a id="L233"></a>SYS_TIMER_SETTIME          = (SYS_SYSCALL_BASE + 258);
    <a id="L234"></a>SYS_TIMER_GETTIME          = (SYS_SYSCALL_BASE + 259);
    <a id="L235"></a>SYS_TIMER_GETOVERRUN       = (SYS_SYSCALL_BASE + 260);
    <a id="L236"></a>SYS_TIMER_DELETE           = (SYS_SYSCALL_BASE + 261);
    <a id="L237"></a>SYS_CLOCK_SETTIME          = (SYS_SYSCALL_BASE + 262);
    <a id="L238"></a>SYS_CLOCK_GETTIME          = (SYS_SYSCALL_BASE + 263);
    <a id="L239"></a>SYS_CLOCK_GETRES           = (SYS_SYSCALL_BASE + 264);
    <a id="L240"></a>SYS_CLOCK_NANOSLEEP        = (SYS_SYSCALL_BASE + 265);
    <a id="L241"></a>SYS_STATFS64               = (SYS_SYSCALL_BASE + 266);
    <a id="L242"></a>SYS_FSTATFS64              = (SYS_SYSCALL_BASE + 267);
    <a id="L243"></a>SYS_TGKILL                 = (SYS_SYSCALL_BASE + 268);
    <a id="L244"></a>SYS_UTIMES                 = (SYS_SYSCALL_BASE + 269);
    <a id="L245"></a>SYS_ARM_FADVISE64_64       = (SYS_SYSCALL_BASE + 270);
    <a id="L246"></a>SYS_PCICONFIG_IOBASE       = (SYS_SYSCALL_BASE + 271);
    <a id="L247"></a>SYS_PCICONFIG_READ         = (SYS_SYSCALL_BASE + 272);
    <a id="L248"></a>SYS_PCICONFIG_WRITE        = (SYS_SYSCALL_BASE + 273);
    <a id="L249"></a>SYS_MQ_OPEN                = (SYS_SYSCALL_BASE + 274);
    <a id="L250"></a>SYS_MQ_UNLINK              = (SYS_SYSCALL_BASE + 275);
    <a id="L251"></a>SYS_MQ_TIMEDSEND           = (SYS_SYSCALL_BASE + 276);
    <a id="L252"></a>SYS_MQ_TIMEDRECEIVE        = (SYS_SYSCALL_BASE + 277);
    <a id="L253"></a>SYS_MQ_NOTIFY              = (SYS_SYSCALL_BASE + 278);
    <a id="L254"></a>SYS_MQ_GETSETATTR          = (SYS_SYSCALL_BASE + 279);
    <a id="L255"></a>SYS_WAITID                 = (SYS_SYSCALL_BASE + 280);
    <a id="L256"></a>SYS_SOCKET                 = (SYS_SYSCALL_BASE + 281);
    <a id="L257"></a>SYS_BIND                   = (SYS_SYSCALL_BASE + 282);
    <a id="L258"></a>SYS_CONNECT                = (SYS_SYSCALL_BASE + 283);
    <a id="L259"></a>SYS_LISTEN                 = (SYS_SYSCALL_BASE + 284);
    <a id="L260"></a>SYS_ACCEPT                 = (SYS_SYSCALL_BASE + 285);
    <a id="L261"></a>SYS_GETSOCKNAME            = (SYS_SYSCALL_BASE + 286);
    <a id="L262"></a>SYS_GETPEERNAME            = (SYS_SYSCALL_BASE + 287);
    <a id="L263"></a>SYS_SOCKETPAIR             = (SYS_SYSCALL_BASE + 288);
    <a id="L264"></a>SYS_SEND                   = (SYS_SYSCALL_BASE + 289);
    <a id="L265"></a>SYS_SENDTO                 = (SYS_SYSCALL_BASE + 290);
    <a id="L266"></a>SYS_RECV                   = (SYS_SYSCALL_BASE + 291);
    <a id="L267"></a>SYS_RECVFROM               = (SYS_SYSCALL_BASE + 292);
    <a id="L268"></a>SYS_SHUTDOWN               = (SYS_SYSCALL_BASE + 293);
    <a id="L269"></a>SYS_SETSOCKOPT             = (SYS_SYSCALL_BASE + 294);
    <a id="L270"></a>SYS_GETSOCKOPT             = (SYS_SYSCALL_BASE + 295);
    <a id="L271"></a>SYS_SENDMSG                = (SYS_SYSCALL_BASE + 296);
    <a id="L272"></a>SYS_RECVMSG                = (SYS_SYSCALL_BASE + 297);
    <a id="L273"></a>SYS_SEMOP                  = (SYS_SYSCALL_BASE + 298);
    <a id="L274"></a>SYS_SEMGET                 = (SYS_SYSCALL_BASE + 299);
    <a id="L275"></a>SYS_SEMCTL                 = (SYS_SYSCALL_BASE + 300);
    <a id="L276"></a>SYS_MSGSND                 = (SYS_SYSCALL_BASE + 301);
    <a id="L277"></a>SYS_MSGRCV                 = (SYS_SYSCALL_BASE + 302);
    <a id="L278"></a>SYS_MSGGET                 = (SYS_SYSCALL_BASE + 303);
    <a id="L279"></a>SYS_MSGCTL                 = (SYS_SYSCALL_BASE + 304);
    <a id="L280"></a>SYS_SHMAT                  = (SYS_SYSCALL_BASE + 305);
    <a id="L281"></a>SYS_SHMDT                  = (SYS_SYSCALL_BASE + 306);
    <a id="L282"></a>SYS_SHMGET                 = (SYS_SYSCALL_BASE + 307);
    <a id="L283"></a>SYS_SHMCTL                 = (SYS_SYSCALL_BASE + 308);
    <a id="L284"></a>SYS_ADD_KEY                = (SYS_SYSCALL_BASE + 309);
    <a id="L285"></a>SYS_REQUEST_KEY            = (SYS_SYSCALL_BASE + 310);
    <a id="L286"></a>SYS_KEYCTL                 = (SYS_SYSCALL_BASE + 311);
    <a id="L287"></a>SYS_SEMTIMEDOP             = (SYS_SYSCALL_BASE + 312);
    <a id="L288"></a>SYS_VSERVER                = (SYS_SYSCALL_BASE + 313);
    <a id="L289"></a>SYS_IOPRIO_SET             = (SYS_SYSCALL_BASE + 314);
    <a id="L290"></a>SYS_IOPRIO_GET             = (SYS_SYSCALL_BASE + 315);
    <a id="L291"></a>SYS_INOTIFY_INIT           = (SYS_SYSCALL_BASE + 316);
    <a id="L292"></a>SYS_INOTIFY_ADD_WATCH      = (SYS_SYSCALL_BASE + 317);
    <a id="L293"></a>SYS_INOTIFY_RM_WATCH       = (SYS_SYSCALL_BASE + 318);
    <a id="L294"></a>SYS_MBIND                  = (SYS_SYSCALL_BASE + 319);
    <a id="L295"></a>SYS_GET_MEMPOLICY          = (SYS_SYSCALL_BASE + 320);
    <a id="L296"></a>SYS_SET_MEMPOLICY          = (SYS_SYSCALL_BASE + 321);
    <a id="L297"></a>SYS_OPENAT                 = (SYS_SYSCALL_BASE + 322);
    <a id="L298"></a>SYS_MKDIRAT                = (SYS_SYSCALL_BASE + 323);
    <a id="L299"></a>SYS_MKNODAT                = (SYS_SYSCALL_BASE + 324);
    <a id="L300"></a>SYS_FCHOWNAT               = (SYS_SYSCALL_BASE + 325);
    <a id="L301"></a>SYS_FUTIMESAT              = (SYS_SYSCALL_BASE + 326);
    <a id="L302"></a>SYS_FSTATAT64              = (SYS_SYSCALL_BASE + 327);
    <a id="L303"></a>SYS_UNLINKAT               = (SYS_SYSCALL_BASE + 328);
    <a id="L304"></a>SYS_RENAMEAT               = (SYS_SYSCALL_BASE + 329);
    <a id="L305"></a>SYS_LINKAT                 = (SYS_SYSCALL_BASE + 330);
    <a id="L306"></a>SYS_SYMLINKAT              = (SYS_SYSCALL_BASE + 331);
    <a id="L307"></a>SYS_READLINKAT             = (SYS_SYSCALL_BASE + 332);
    <a id="L308"></a>SYS_FCHMODAT               = (SYS_SYSCALL_BASE + 333);
    <a id="L309"></a>SYS_FACCESSAT              = (SYS_SYSCALL_BASE + 334);
    <a id="L310"></a>SYS_UNSHARE                = (SYS_SYSCALL_BASE + 337);
    <a id="L311"></a>SYS_SET_ROBUST_LIST        = (SYS_SYSCALL_BASE + 338);
    <a id="L312"></a>SYS_GET_ROBUST_LIST        = (SYS_SYSCALL_BASE + 339);
    <a id="L313"></a>SYS_SPLICE                 = (SYS_SYSCALL_BASE + 340);
    <a id="L314"></a>SYS_ARM_SYNC_FILE_RANGE    = (SYS_SYSCALL_BASE + 341);
    <a id="L315"></a>SYS_SYNC_FILE_RANGE2       = SYS_ARM_SYNC_FILE_RANGE;
    <a id="L316"></a>SYS_TEE                    = (SYS_SYSCALL_BASE + 342);
    <a id="L317"></a>SYS_VMSPLICE               = (SYS_SYSCALL_BASE + 343);
    <a id="L318"></a>SYS_MOVE_PAGES             = (SYS_SYSCALL_BASE + 344);
    <a id="L319"></a>SYS_GETCPU                 = (SYS_SYSCALL_BASE + 345);
    <a id="L320"></a>SYS_KEXEC_LOAD             = (SYS_SYSCALL_BASE + 347);
    <a id="L321"></a>SYS_UTIMENSAT              = (SYS_SYSCALL_BASE + 348);
    <a id="L322"></a>SYS_SIGNALFD               = (SYS_SYSCALL_BASE + 349);
    <a id="L323"></a>SYS_TIMERFD_CREATE         = (SYS_SYSCALL_BASE + 350);
    <a id="L324"></a>SYS_EVENTFD                = (SYS_SYSCALL_BASE + 351);
    <a id="L325"></a>SYS_FALLOCATE              = (SYS_SYSCALL_BASE + 352);
    <a id="L326"></a>SYS_TIMERFD_SETTIME        = (SYS_SYSCALL_BASE + 353);
    <a id="L327"></a>SYS_TIMERFD_GETTIME        = (SYS_SYSCALL_BASE + 354);
    <a id="L328"></a>SYS_SIGNALFD4              = (SYS_SYSCALL_BASE + 355);
    <a id="L329"></a>SYS_EVENTFD2               = (SYS_SYSCALL_BASE + 356);
    <a id="L330"></a>SYS_EPOLL_CREATE1          = (SYS_SYSCALL_BASE + 357);
    <a id="L331"></a>SYS_DUP3                   = (SYS_SYSCALL_BASE + 358);
    <a id="L332"></a>SYS_PIPE2                  = (SYS_SYSCALL_BASE + 359);
    <a id="L333"></a>SYS_INOTIFY_INIT1          = (SYS_SYSCALL_BASE + 360);
<a id="L334"></a>)

<a id="L336"></a>func _darwin_system_call_conflict() {}
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
