<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/zerrors_linux_arm.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/syscall/zerrors_linux_arm.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// mkerrors.sh</span>
<a id="L2"></a><span class="comment">// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT</span>

<a id="L4"></a><span class="comment">// godefs -gsyscall _errors.c</span>

<a id="L6"></a><span class="comment">// MACHINE GENERATED - DO NOT EDIT.</span>

<a id="L8"></a>package syscall

<a id="L10"></a><span class="comment">// Constants</span>
<a id="L11"></a>const (
    <a id="L12"></a>EMULTIHOP       = 0x48;
    <a id="L13"></a>EUNATCH         = 0x31;
    <a id="L14"></a>EAFNOSUPPORT    = 0x61;
    <a id="L15"></a>EREMCHG         = 0x4e;
    <a id="L16"></a>EACCES          = 0xd;
    <a id="L17"></a>EL3RST          = 0x2f;
    <a id="L18"></a>EDESTADDRREQ    = 0x59;
    <a id="L19"></a>EILSEQ          = 0x54;
    <a id="L20"></a>ESPIPE          = 0x1d;
    <a id="L21"></a>EMLINK          = 0x1f;
    <a id="L22"></a>EOWNERDEAD      = 0x82;
    <a id="L23"></a>ENOTTY          = 0x19;
    <a id="L24"></a>EBADE           = 0x34;
    <a id="L25"></a>EBADF           = 0x9;
    <a id="L26"></a>EBADR           = 0x35;
    <a id="L27"></a>EADV            = 0x44;
    <a id="L28"></a>ERANGE          = 0x22;
    <a id="L29"></a>ECANCELED       = 0x7d;
    <a id="L30"></a>ETXTBSY         = 0x1a;
    <a id="L31"></a>ENOMEM          = 0xc;
    <a id="L32"></a>EINPROGRESS     = 0x73;
    <a id="L33"></a>ENOTEMPTY       = 0x27;
    <a id="L34"></a>ENOTBLK         = 0xf;
    <a id="L35"></a>EPROTOTYPE      = 0x5b;
    <a id="L36"></a>ERESTART        = 0x55;
    <a id="L37"></a>EISNAM          = 0x78;
    <a id="L38"></a>ENOMSG          = 0x2a;
    <a id="L39"></a>EALREADY        = 0x72;
    <a id="L40"></a>ETIMEDOUT       = 0x6e;
    <a id="L41"></a>ENODATA         = 0x3d;
    <a id="L42"></a>EINTR           = 0x4;
    <a id="L43"></a>ENOLINK         = 0x43;
    <a id="L44"></a>EPERM           = 0x1;
    <a id="L45"></a>ELOOP           = 0x28;
    <a id="L46"></a>ENETDOWN        = 0x64;
    <a id="L47"></a>ESTALE          = 0x74;
    <a id="L48"></a>ENOTSOCK        = 0x58;
    <a id="L49"></a>ENOSR           = 0x3f;
    <a id="L50"></a>ECHILD          = 0xa;
    <a id="L51"></a>ELNRNG          = 0x30;
    <a id="L52"></a>EPIPE           = 0x20;
    <a id="L53"></a>EBADMSG         = 0x4a;
    <a id="L54"></a>EBFONT          = 0x3b;
    <a id="L55"></a>EREMOTE         = 0x42;
    <a id="L56"></a>ETOOMANYREFS    = 0x6d;
    <a id="L57"></a>EPFNOSUPPORT    = 0x60;
    <a id="L58"></a>ENONET          = 0x40;
    <a id="L59"></a>EXFULL          = 0x36;
    <a id="L60"></a>EBADSLT         = 0x39;
    <a id="L61"></a>ENOTNAM         = 0x76;
    <a id="L62"></a>ENOCSI          = 0x32;
    <a id="L63"></a>EADDRINUSE      = 0x62;
    <a id="L64"></a>ENETRESET       = 0x66;
    <a id="L65"></a>EISDIR          = 0x15;
    <a id="L66"></a>EIDRM           = 0x2b;
    <a id="L67"></a>ECOMM           = 0x46;
    <a id="L68"></a>EBADFD          = 0x4d;
    <a id="L69"></a>EL2HLT          = 0x33;
    <a id="L70"></a>ENOKEY          = 0x7e;
    <a id="L71"></a>EINVAL          = 0x16;
    <a id="L72"></a>ESHUTDOWN       = 0x6c;
    <a id="L73"></a>EKEYREJECTED    = 0x81;
    <a id="L74"></a>ELIBSCN         = 0x51;
    <a id="L75"></a>ENAVAIL         = 0x77;
    <a id="L76"></a>EOVERFLOW       = 0x4b;
    <a id="L77"></a>EUCLEAN         = 0x75;
    <a id="L78"></a>ENOMEDIUM       = 0x7b;
    <a id="L79"></a>EBUSY           = 0x10;
    <a id="L80"></a>EPROTO          = 0x47;
    <a id="L81"></a>ENODEV          = 0x13;
    <a id="L82"></a>EKEYEXPIRED     = 0x7f;
    <a id="L83"></a>EROFS           = 0x1e;
    <a id="L84"></a>ELIBACC         = 0x4f;
    <a id="L85"></a>E2BIG           = 0x7;
    <a id="L86"></a>EDEADLK         = 0x23;
    <a id="L87"></a>ENOTDIR         = 0x14;
    <a id="L88"></a>ECONNRESET      = 0x68;
    <a id="L89"></a>ENXIO           = 0x6;
    <a id="L90"></a>EBADRQC         = 0x38;
    <a id="L91"></a>ENAMETOOLONG    = 0x24;
    <a id="L92"></a>ESOCKTNOSUPPORT = 0x5e;
    <a id="L93"></a>ELIBEXEC        = 0x53;
    <a id="L94"></a>EDOTDOT         = 0x49;
    <a id="L95"></a>EADDRNOTAVAIL   = 0x63;
    <a id="L96"></a>ETIME           = 0x3e;
    <a id="L97"></a>EPROTONOSUPPORT = 0x5d;
    <a id="L98"></a>ENOTRECOVERABLE = 0x83;
    <a id="L99"></a>EIO             = 0x5;
    <a id="L100"></a>ENETUNREACH     = 0x65;
    <a id="L101"></a>EXDEV           = 0x12;
    <a id="L102"></a>EDQUOT          = 0x7a;
    <a id="L103"></a>EREMOTEIO       = 0x79;
    <a id="L104"></a>ENOSPC          = 0x1c;
    <a id="L105"></a>ENOEXEC         = 0x8;
    <a id="L106"></a>EMSGSIZE        = 0x5a;
    <a id="L107"></a>EDOM            = 0x21;
    <a id="L108"></a>ENOSTR          = 0x3c;
    <a id="L109"></a>EFBIG           = 0x1b;
    <a id="L110"></a>ESRCH           = 0x3;
    <a id="L111"></a>ECHRNG          = 0x2c;
    <a id="L112"></a>EHOSTDOWN       = 0x70;
    <a id="L113"></a>ENOLCK          = 0x25;
    <a id="L114"></a>ENFILE          = 0x17;
    <a id="L115"></a>ENOSYS          = 0x26;
    <a id="L116"></a>ENOTCONN        = 0x6b;
    <a id="L117"></a>ENOTSUP         = 0x5f;
    <a id="L118"></a>ESRMNT          = 0x45;
    <a id="L119"></a>EDEADLOCK       = 0x23;
    <a id="L120"></a>ECONNABORTED    = 0x67;
    <a id="L121"></a>ENOANO          = 0x37;
    <a id="L122"></a>EISCONN         = 0x6a;
    <a id="L123"></a>EUSERS          = 0x57;
    <a id="L124"></a>ENOPROTOOPT     = 0x5c;
    <a id="L125"></a>EMFILE          = 0x18;
    <a id="L126"></a>ENOBUFS         = 0x69;
    <a id="L127"></a>EL3HLT          = 0x2e;
    <a id="L128"></a>EFAULT          = 0xe;
    <a id="L129"></a>EWOULDBLOCK     = 0xb;
    <a id="L130"></a>ELIBBAD         = 0x50;
    <a id="L131"></a>ESTRPIPE        = 0x56;
    <a id="L132"></a>ECONNREFUSED    = 0x6f;
    <a id="L133"></a>EAGAIN          = 0xb;
    <a id="L134"></a>ELIBMAX         = 0x52;
    <a id="L135"></a>EEXIST          = 0x11;
    <a id="L136"></a>EL2NSYNC        = 0x2d;
    <a id="L137"></a>ENOENT          = 0x2;
    <a id="L138"></a>ENOPKG          = 0x41;
    <a id="L139"></a>EKEYREVOKED     = 0x80;
    <a id="L140"></a>EHOSTUNREACH    = 0x71;
    <a id="L141"></a>ENOTUNIQ        = 0x4c;
    <a id="L142"></a>EOPNOTSUPP      = 0x5f;
    <a id="L143"></a>EMEDIUMTYPE     = 0x7c;
    <a id="L144"></a>SIGBUS          = 0x7;
    <a id="L145"></a>SIGTTIN         = 0x15;
    <a id="L146"></a>SIGPROF         = 0x1b;
    <a id="L147"></a>SIGFPE          = 0x8;
    <a id="L148"></a>SIGHUP          = 0x1;
    <a id="L149"></a>SIGTTOU         = 0x16;
    <a id="L150"></a>SIGSTKFLT       = 0x10;
    <a id="L151"></a>SIGUSR1         = 0xa;
    <a id="L152"></a>SIGURG          = 0x17;
    <a id="L153"></a>SIGIO           = 0x1d;
    <a id="L154"></a>SIGQUIT         = 0x3;
    <a id="L155"></a>SIGCLD          = 0x11;
    <a id="L156"></a>SIGABRT         = 0x6;
    <a id="L157"></a>SIGTRAP         = 0x5;
    <a id="L158"></a>SIGVTALRM       = 0x1a;
    <a id="L159"></a>SIGPOLL         = 0x1d;
    <a id="L160"></a>SIGSEGV         = 0xb;
    <a id="L161"></a>SIGCONT         = 0x12;
    <a id="L162"></a>SIGPIPE         = 0xd;
    <a id="L163"></a>SIGWINCH        = 0x1c;
    <a id="L164"></a>SIGXFSZ         = 0x19;
    <a id="L165"></a>SIGCHLD         = 0x11;
    <a id="L166"></a>SIGSYS          = 0x1f;
    <a id="L167"></a>SIGSTOP         = 0x13;
    <a id="L168"></a>SIGALRM         = 0xe;
    <a id="L169"></a>SIGUSR2         = 0xc;
    <a id="L170"></a>SIGTSTP         = 0x14;
    <a id="L171"></a>SIGKILL         = 0x9;
    <a id="L172"></a>SIGXCPU         = 0x18;
    <a id="L173"></a>SIGUNUSED       = 0x1f;
    <a id="L174"></a>SIGPWR          = 0x1e;
    <a id="L175"></a>SIGILL          = 0x4;
    <a id="L176"></a>SIGINT          = 0x2;
    <a id="L177"></a>SIGIOT          = 0x6;
    <a id="L178"></a>SIGTERM         = 0xf;
    <a id="L179"></a>O_EXCL          = 0x80;
<a id="L180"></a>)

<a id="L182"></a><span class="comment">// Types</span>


<a id="L185"></a><span class="comment">// Error table</span>
<a id="L186"></a>var errors = [...]string{
    <a id="L187"></a>72: &#34;multihop attempted&#34;,
    <a id="L188"></a>49: &#34;protocol driver not attached&#34;,
    <a id="L189"></a>97: &#34;address family not supported by protocol&#34;,
    <a id="L190"></a>78: &#34;remote address changed&#34;,
    <a id="L191"></a>13: &#34;permission denied&#34;,
    <a id="L192"></a>47: &#34;level 3 reset&#34;,
    <a id="L193"></a>89: &#34;destination address required&#34;,
    <a id="L194"></a>84: &#34;invalid or incomplete multibyte or wide character&#34;,
    <a id="L195"></a>29: &#34;illegal seek&#34;,
    <a id="L196"></a>31: &#34;too many links&#34;,
    <a id="L197"></a>130: &#34;owner died&#34;,
    <a id="L198"></a>25: &#34;inappropriate ioctl for device&#34;,
    <a id="L199"></a>52: &#34;invalid exchange&#34;,
    <a id="L200"></a>9: &#34;bad file descriptor&#34;,
    <a id="L201"></a>53: &#34;invalid request descriptor&#34;,
    <a id="L202"></a>68: &#34;advertise error&#34;,
    <a id="L203"></a>34: &#34;numerical result out of range&#34;,
    <a id="L204"></a>125: &#34;operation canceled&#34;,
    <a id="L205"></a>26: &#34;text file busy&#34;,
    <a id="L206"></a>12: &#34;cannot allocate memory&#34;,
    <a id="L207"></a>115: &#34;operation now in progress&#34;,
    <a id="L208"></a>39: &#34;directory not empty&#34;,
    <a id="L209"></a>15: &#34;block device required&#34;,
    <a id="L210"></a>91: &#34;protocol wrong type for socket&#34;,
    <a id="L211"></a>85: &#34;interrupted system call should be restarted&#34;,
    <a id="L212"></a>120: &#34;is a named type file&#34;,
    <a id="L213"></a>42: &#34;no message of desired type&#34;,
    <a id="L214"></a>114: &#34;operation already in progress&#34;,
    <a id="L215"></a>110: &#34;connection timed out&#34;,
    <a id="L216"></a>61: &#34;no data available&#34;,
    <a id="L217"></a>4: &#34;interrupted system call&#34;,
    <a id="L218"></a>67: &#34;link has been severed&#34;,
    <a id="L219"></a>1: &#34;operation not permitted&#34;,
    <a id="L220"></a>40: &#34;too many levels of symbolic links&#34;,
    <a id="L221"></a>100: &#34;network is down&#34;,
    <a id="L222"></a>116: &#34;stale NFS file handle&#34;,
    <a id="L223"></a>88: &#34;socket operation on non-socket&#34;,
    <a id="L224"></a>63: &#34;out of streams resources&#34;,
    <a id="L225"></a>10: &#34;no child processes&#34;,
    <a id="L226"></a>48: &#34;link number out of range&#34;,
    <a id="L227"></a>32: &#34;broken pipe&#34;,
    <a id="L228"></a>74: &#34;bad message&#34;,
    <a id="L229"></a>59: &#34;bad font file format&#34;,
    <a id="L230"></a>66: &#34;object is remote&#34;,
    <a id="L231"></a>109: &#34;too many references: cannot splice&#34;,
    <a id="L232"></a>96: &#34;protocol family not supported&#34;,
    <a id="L233"></a>64: &#34;machine is not on the network&#34;,
    <a id="L234"></a>54: &#34;exchange full&#34;,
    <a id="L235"></a>57: &#34;invalid slot&#34;,
    <a id="L236"></a>118: &#34;not a XENIX named type file&#34;,
    <a id="L237"></a>50: &#34;no CSI structure available&#34;,
    <a id="L238"></a>98: &#34;address already in use&#34;,
    <a id="L239"></a>102: &#34;network dropped connection on reset&#34;,
    <a id="L240"></a>21: &#34;is a directory&#34;,
    <a id="L241"></a>43: &#34;identifier removed&#34;,
    <a id="L242"></a>70: &#34;communication error on send&#34;,
    <a id="L243"></a>77: &#34;file descriptor in bad state&#34;,
    <a id="L244"></a>51: &#34;level 2 halted&#34;,
    <a id="L245"></a>126: &#34;required key not available&#34;,
    <a id="L246"></a>22: &#34;invalid argument&#34;,
    <a id="L247"></a>108: &#34;cannot send after transport endpoint shutdown&#34;,
    <a id="L248"></a>129: &#34;key was rejected by service&#34;,
    <a id="L249"></a>81: &#34;.lib section in a.out corrupted&#34;,
    <a id="L250"></a>119: &#34;no XENIX semaphores available&#34;,
    <a id="L251"></a>75: &#34;value too large for defined data type&#34;,
    <a id="L252"></a>117: &#34;structure needs cleaning&#34;,
    <a id="L253"></a>123: &#34;no medium found&#34;,
    <a id="L254"></a>16: &#34;device or resource busy&#34;,
    <a id="L255"></a>71: &#34;protocol error&#34;,
    <a id="L256"></a>19: &#34;no such device&#34;,
    <a id="L257"></a>127: &#34;key has expired&#34;,
    <a id="L258"></a>30: &#34;read-only file system&#34;,
    <a id="L259"></a>79: &#34;can not access a needed shared library&#34;,
    <a id="L260"></a>7: &#34;argument list too long&#34;,
    <a id="L261"></a>35: &#34;resource deadlock avoided&#34;,
    <a id="L262"></a>20: &#34;not a directory&#34;,
    <a id="L263"></a>104: &#34;connection reset by peer&#34;,
    <a id="L264"></a>6: &#34;no such device or address&#34;,
    <a id="L265"></a>56: &#34;invalid request code&#34;,
    <a id="L266"></a>36: &#34;file name too long&#34;,
    <a id="L267"></a>94: &#34;socket type not supported&#34;,
    <a id="L268"></a>83: &#34;cannot exec a shared library directly&#34;,
    <a id="L269"></a>73: &#34;RFS specific error&#34;,
    <a id="L270"></a>99: &#34;cannot assign requested address&#34;,
    <a id="L271"></a>62: &#34;timer expired&#34;,
    <a id="L272"></a>93: &#34;protocol not supported&#34;,
    <a id="L273"></a>131: &#34;state not recoverable&#34;,
    <a id="L274"></a>5: &#34;input/output error&#34;,
    <a id="L275"></a>101: &#34;network is unreachable&#34;,
    <a id="L276"></a>18: &#34;invalid cross-device link&#34;,
    <a id="L277"></a>122: &#34;disk quota exceeded&#34;,
    <a id="L278"></a>121: &#34;remote I/O error&#34;,
    <a id="L279"></a>28: &#34;no space left on device&#34;,
    <a id="L280"></a>8: &#34;exec format error&#34;,
    <a id="L281"></a>90: &#34;message too long&#34;,
    <a id="L282"></a>33: &#34;numerical argument out of domain&#34;,
    <a id="L283"></a>60: &#34;device not a stream&#34;,
    <a id="L284"></a>27: &#34;file too large&#34;,
    <a id="L285"></a>3: &#34;no such process&#34;,
    <a id="L286"></a>44: &#34;channel number out of range&#34;,
    <a id="L287"></a>112: &#34;host is down&#34;,
    <a id="L288"></a>37: &#34;no locks available&#34;,
    <a id="L289"></a>23: &#34;too many open files in system&#34;,
    <a id="L290"></a>38: &#34;function not implemented&#34;,
    <a id="L291"></a>107: &#34;transport endpoint is not connected&#34;,
    <a id="L292"></a>95: &#34;operation not supported&#34;,
    <a id="L293"></a>69: &#34;srmount error&#34;,
    <a id="L294"></a>103: &#34;software caused connection abort&#34;,
    <a id="L295"></a>55: &#34;no anode&#34;,
    <a id="L296"></a>106: &#34;transport endpoint is already connected&#34;,
    <a id="L297"></a>87: &#34;too many users&#34;,
    <a id="L298"></a>92: &#34;protocol not available&#34;,
    <a id="L299"></a>24: &#34;too many open files&#34;,
    <a id="L300"></a>105: &#34;no buffer space available&#34;,
    <a id="L301"></a>46: &#34;level 3 halted&#34;,
    <a id="L302"></a>14: &#34;bad address&#34;,
    <a id="L303"></a>11: &#34;resource temporarily unavailable&#34;,
    <a id="L304"></a>80: &#34;accessing a corrupted shared library&#34;,
    <a id="L305"></a>86: &#34;streams pipe error&#34;,
    <a id="L306"></a>111: &#34;connection refused&#34;,
    <a id="L307"></a>82: &#34;attempting to link in too many shared libraries&#34;,
    <a id="L308"></a>17: &#34;file exists&#34;,
    <a id="L309"></a>45: &#34;level 2 not synchronized&#34;,
    <a id="L310"></a>2: &#34;no such file or directory&#34;,
    <a id="L311"></a>65: &#34;package not installed&#34;,
    <a id="L312"></a>128: &#34;key has been revoked&#34;,
    <a id="L313"></a>113: &#34;no route to host&#34;,
    <a id="L314"></a>76: &#34;name not unique on network&#34;,
    <a id="L315"></a>124: &#34;wrong medium type&#34;,
<a id="L316"></a>}
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
