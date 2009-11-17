<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/zerrors_nacl_386.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/syscall/zerrors_nacl_386.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// mkerrors_nacl.sh /home/rsc/pub/nacl/native_client/src/trusted/service_runtime/include/sys/errno.h</span>
<a id="L2"></a><span class="comment">// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT</span>

<a id="L4"></a>package syscall

<a id="L6"></a>const (
    <a id="L7"></a>EPERM           = 1;
    <a id="L8"></a>ENOENT          = 2;
    <a id="L9"></a>ESRCH           = 3;
    <a id="L10"></a>EINTR           = 4;
    <a id="L11"></a>EIO             = 5;
    <a id="L12"></a>ENXIO           = 6;
    <a id="L13"></a>E2BIG           = 7;
    <a id="L14"></a>ENOEXEC         = 8;
    <a id="L15"></a>EBADF           = 9;
    <a id="L16"></a>ECHILD          = 10;
    <a id="L17"></a>EAGAIN          = 11;
    <a id="L18"></a>ENOMEM          = 12;
    <a id="L19"></a>EACCES          = 13;
    <a id="L20"></a>EFAULT          = 14;
    <a id="L21"></a>EBUSY           = 16;
    <a id="L22"></a>EEXIST          = 17;
    <a id="L23"></a>EXDEV           = 18;
    <a id="L24"></a>ENODEV          = 19;
    <a id="L25"></a>ENOTDIR         = 20;
    <a id="L26"></a>EISDIR          = 21;
    <a id="L27"></a>EINVAL          = 22;
    <a id="L28"></a>ENFILE          = 23;
    <a id="L29"></a>EMFILE          = 24;
    <a id="L30"></a>ENOTTY          = 25;
    <a id="L31"></a>EFBIG           = 27;
    <a id="L32"></a>ENOSPC          = 28;
    <a id="L33"></a>ESPIPE          = 29;
    <a id="L34"></a>EROFS           = 30;
    <a id="L35"></a>EMLINK          = 31;
    <a id="L36"></a>EPIPE           = 32;
    <a id="L37"></a>ENAMETOOLONG    = 36;
    <a id="L38"></a>ENOSYS          = 38;
    <a id="L39"></a>EDQUOT          = 122;
    <a id="L40"></a>EDOM            = 33;
    <a id="L41"></a>ERANGE          = 34;
    <a id="L42"></a>ENOMSG          = 35;
    <a id="L43"></a>ECHRNG          = 37;
    <a id="L44"></a>EL3HLT          = 39;
    <a id="L45"></a>EL3RST          = 40;
    <a id="L46"></a>ELNRNG          = 41;
    <a id="L47"></a>EUNATCH         = 42;
    <a id="L48"></a>ENOCSI          = 43;
    <a id="L49"></a>EL2HLT          = 44;
    <a id="L50"></a>EDEADLK         = 45;
    <a id="L51"></a>ENOLCK          = 46;
    <a id="L52"></a>EBADE           = 50;
    <a id="L53"></a>EBADR           = 51;
    <a id="L54"></a>EXFULL          = 52;
    <a id="L55"></a>ENOANO          = 53;
    <a id="L56"></a>EBADRQC         = 54;
    <a id="L57"></a>EBADSLT         = 55;
    <a id="L58"></a>EBFONT          = 57;
    <a id="L59"></a>ENOSTR          = 60;
    <a id="L60"></a>ENODATA         = 61;
    <a id="L61"></a>ETIME           = 62;
    <a id="L62"></a>ENOSR           = 63;
    <a id="L63"></a>ENONET          = 64;
    <a id="L64"></a>ENOPKG          = 65;
    <a id="L65"></a>EREMOTE         = 66;
    <a id="L66"></a>ENOLINK         = 67;
    <a id="L67"></a>EADV            = 68;
    <a id="L68"></a>ESRMNT          = 69;
    <a id="L69"></a>ECOMM           = 70;
    <a id="L70"></a>EPROTO          = 71;
    <a id="L71"></a>EMULTIHOP       = 74;
    <a id="L72"></a>ELBIN           = 75;
    <a id="L73"></a>EDOTDOT         = 76;
    <a id="L74"></a>EBADMSG         = 77;
    <a id="L75"></a>EFTYPE          = 79;
    <a id="L76"></a>ENOTUNIQ        = 80;
    <a id="L77"></a>EBADFD          = 81;
    <a id="L78"></a>EREMCHG         = 82;
    <a id="L79"></a>ELIBACC         = 83;
    <a id="L80"></a>ELIBBAD         = 84;
    <a id="L81"></a>ELIBSCN         = 85;
    <a id="L82"></a>ELIBMAX         = 86;
    <a id="L83"></a>ELIBEXEC        = 87;
    <a id="L84"></a>ENMFILE         = 89;
    <a id="L85"></a>ENOTEMPTY       = 90;
    <a id="L86"></a>ELOOP           = 92;
    <a id="L87"></a>EOPNOTSUPP      = 95;
    <a id="L88"></a>EPFNOSUPPORT    = 96;
    <a id="L89"></a>ECONNRESET      = 104;
    <a id="L90"></a>ENOBUFS         = 105;
    <a id="L91"></a>EAFNOSUPPORT    = 106;
    <a id="L92"></a>EPROTOTYPE      = 107;
    <a id="L93"></a>ENOTSOCK        = 108;
    <a id="L94"></a>ENOPROTOOPT     = 109;
    <a id="L95"></a>ESHUTDOWN       = 110;
    <a id="L96"></a>ECONNREFUSED    = 111;
    <a id="L97"></a>EADDRINUSE      = 112;
    <a id="L98"></a>ECONNABORTED    = 113;
    <a id="L99"></a>ENETUNREACH     = 114;
    <a id="L100"></a>ENETDOWN        = 115;
    <a id="L101"></a>ETIMEDOUT       = 116;
    <a id="L102"></a>EHOSTDOWN       = 117;
    <a id="L103"></a>EHOSTUNREACH    = 118;
    <a id="L104"></a>EINPROGRESS     = 119;
    <a id="L105"></a>EALREADY        = 120;
    <a id="L106"></a>EDESTADDRREQ    = 121;
    <a id="L107"></a>EPROTONOSUPPORT = 123;
    <a id="L108"></a>ESOCKTNOSUPPORT = 124;
    <a id="L109"></a>EADDRNOTAVAIL   = 125;
    <a id="L110"></a>ENETRESET       = 126;
    <a id="L111"></a>EISCONN         = 127;
    <a id="L112"></a>ENOTCONN        = 128;
    <a id="L113"></a>ETOOMANYREFS    = 129;
    <a id="L114"></a>EPROCLIM        = 130;
    <a id="L115"></a>EUSERS          = 131;
    <a id="L116"></a>ESTALE          = 133;
    <a id="L117"></a>ENOMEDIUM       = 135;
    <a id="L118"></a>ENOSHARE        = 136;
    <a id="L119"></a>ECASECLASH      = 137;
    <a id="L120"></a>EILSEQ          = 138;
    <a id="L121"></a>EOVERFLOW       = 139;
    <a id="L122"></a>ECANCELED       = 140;
    <a id="L123"></a>EL2NSYNC        = 88;
    <a id="L124"></a>EIDRM           = 91;
    <a id="L125"></a>EMSGSIZE        = 132;
    <a id="L126"></a>ENACL           = 99; <span class="comment">/* otherwise unused */</span>
<a id="L127"></a>)


<a id="L130"></a><span class="comment">// Error table</span>
<a id="L131"></a>var errors = [...]string{
    <a id="L132"></a>EPERM: &#34;operation not permitted&#34;,
    <a id="L133"></a>ENOENT: &#34;no such file or directory&#34;,
    <a id="L134"></a>ESRCH: &#34;no such process&#34;,
    <a id="L135"></a>EINTR: &#34;interrupted system call&#34;,
    <a id="L136"></a>EIO: &#34;I/O error&#34;,
    <a id="L137"></a>ENXIO: &#34;no such device or address&#34;,
    <a id="L138"></a>E2BIG: &#34;argument list too long&#34;,
    <a id="L139"></a>ENOEXEC: &#34;exec format error&#34;,
    <a id="L140"></a>EBADF: &#34;bad file number&#34;,
    <a id="L141"></a>ECHILD: &#34;no child processes&#34;,
    <a id="L142"></a>EAGAIN: &#34;try again&#34;,
    <a id="L143"></a>ENOMEM: &#34;out of memory&#34;,
    <a id="L144"></a>EACCES: &#34;permission denied&#34;,
    <a id="L145"></a>EFAULT: &#34;bad address&#34;,
    <a id="L146"></a>EBUSY: &#34;device or resource busy&#34;,
    <a id="L147"></a>EEXIST: &#34;file exists&#34;,
    <a id="L148"></a>EXDEV: &#34;cross-device link&#34;,
    <a id="L149"></a>ENODEV: &#34;no such device&#34;,
    <a id="L150"></a>ENOTDIR: &#34;not a directory&#34;,
    <a id="L151"></a>EISDIR: &#34;is a directory&#34;,
    <a id="L152"></a>EINVAL: &#34;invalid argument&#34;,
    <a id="L153"></a>ENFILE: &#34;file table overflow&#34;,
    <a id="L154"></a>EMFILE: &#34;too many open files&#34;,
    <a id="L155"></a>ENOTTY: &#34;not a typewriter&#34;,
    <a id="L156"></a>EFBIG: &#34;file too large&#34;,
    <a id="L157"></a>ENOSPC: &#34;no space left on device&#34;,
    <a id="L158"></a>ESPIPE: &#34;illegal seek&#34;,
    <a id="L159"></a>EROFS: &#34;read-only file system&#34;,
    <a id="L160"></a>EMLINK: &#34;too many links&#34;,
    <a id="L161"></a>EPIPE: &#34;broken pipe&#34;,
    <a id="L162"></a>ENAMETOOLONG: &#34;file name too long&#34;,
    <a id="L163"></a>ENOSYS: &#34;function not implemented&#34;,
    <a id="L164"></a>EDQUOT: &#34;quota exceeded&#34;,
    <a id="L165"></a>EDOM: &#34;math arg out of domain of func&#34;,
    <a id="L166"></a>ERANGE: &#34;math result not representable&#34;,
    <a id="L167"></a>ENOMSG: &#34;no message of desired type&#34;,
    <a id="L168"></a>ECHRNG: &#34;channel number out of range&#34;,
    <a id="L169"></a>EL3HLT: &#34;level 3 halted&#34;,
    <a id="L170"></a>EL3RST: &#34;level 3 reset&#34;,
    <a id="L171"></a>ELNRNG: &#34;link number out of range&#34;,
    <a id="L172"></a>EUNATCH: &#34;protocol driver not attached&#34;,
    <a id="L173"></a>ENOCSI: &#34;no CSI structure available&#34;,
    <a id="L174"></a>EL2HLT: &#34;level 2 halted&#34;,
    <a id="L175"></a>EDEADLK: &#34;deadlock condition&#34;,
    <a id="L176"></a>ENOLCK: &#34;no record locks available&#34;,
    <a id="L177"></a>EBADE: &#34;invalid exchange&#34;,
    <a id="L178"></a>EBADR: &#34;invalid request descriptor&#34;,
    <a id="L179"></a>EXFULL: &#34;exchange full&#34;,
    <a id="L180"></a>ENOANO: &#34;no anode&#34;,
    <a id="L181"></a>EBADRQC: &#34;invalid request code&#34;,
    <a id="L182"></a>EBADSLT: &#34;invalid slot&#34;,
    <a id="L183"></a>EBFONT: &#34;bad font file fmt&#34;,
    <a id="L184"></a>ENOSTR: &#34;device not a stream&#34;,
    <a id="L185"></a>ENODATA: &#34;no data (for no delay io)&#34;,
    <a id="L186"></a>ETIME: &#34;timer expired&#34;,
    <a id="L187"></a>ENOSR: &#34;out of streams resources&#34;,
    <a id="L188"></a>ENONET: &#34;machine is not on the network&#34;,
    <a id="L189"></a>ENOPKG: &#34;package not installed&#34;,
    <a id="L190"></a>EREMOTE: &#34;the object is remote&#34;,
    <a id="L191"></a>ENOLINK: &#34;the link has been severed&#34;,
    <a id="L192"></a>EADV: &#34;advertise error&#34;,
    <a id="L193"></a>ESRMNT: &#34;srmount error&#34;,
    <a id="L194"></a>ECOMM: &#34;communication error on send&#34;,
    <a id="L195"></a>EPROTO: &#34;protocol error&#34;,
    <a id="L196"></a>EMULTIHOP: &#34;multihop attempted&#34;,
    <a id="L197"></a>ELBIN: &#34;inode is remote (not really error)&#34;,
    <a id="L198"></a>EDOTDOT: &#34;cross mount point (not really error)&#34;,
    <a id="L199"></a>EBADMSG: &#34;trying to read unreadable message&#34;,
    <a id="L200"></a>EFTYPE: &#34;inappropriate file type or format&#34;,
    <a id="L201"></a>ENOTUNIQ: &#34;given log. name not unique&#34;,
    <a id="L202"></a>EBADFD: &#34;f.d. invalid for this operation&#34;,
    <a id="L203"></a>EREMCHG: &#34;remote address changed&#34;,
    <a id="L204"></a>ELIBACC: &#34;can&#39;t access a needed shared lib&#34;,
    <a id="L205"></a>ELIBBAD: &#34;accessing a corrupted shared lib&#34;,
    <a id="L206"></a>ELIBSCN: &#34;.lib section in a.out corrupted&#34;,
    <a id="L207"></a>ELIBMAX: &#34;attempting to link in too many libs&#34;,
    <a id="L208"></a>ELIBEXEC: &#34;attempting to exec a shared library&#34;,
    <a id="L209"></a>ENMFILE: &#34;no more files&#34;,
    <a id="L210"></a>ENOTEMPTY: &#34;directory not empty&#34;,
    <a id="L211"></a>ELOOP: &#34;too many symbolic links&#34;,
    <a id="L212"></a>EOPNOTSUPP: &#34;operation not supported on transport endpoint&#34;,
    <a id="L213"></a>EPFNOSUPPORT: &#34;protocol family not supported&#34;,
    <a id="L214"></a>ECONNRESET: &#34;connection reset by peer&#34;,
    <a id="L215"></a>ENOBUFS: &#34;no buffer space available&#34;,
    <a id="L216"></a>EAFNOSUPPORT: &#34;address family not supported by protocol family&#34;,
    <a id="L217"></a>EPROTOTYPE: &#34;protocol wrong type for socket&#34;,
    <a id="L218"></a>ENOTSOCK: &#34;socket operation on non-socket&#34;,
    <a id="L219"></a>ENOPROTOOPT: &#34;protocol not available&#34;,
    <a id="L220"></a>ESHUTDOWN: &#34;can&#39;t send after socket shutdown&#34;,
    <a id="L221"></a>ECONNREFUSED: &#34;connection refused&#34;,
    <a id="L222"></a>EADDRINUSE: &#34;address already in use&#34;,
    <a id="L223"></a>ECONNABORTED: &#34;connection aborted&#34;,
    <a id="L224"></a>ENETUNREACH: &#34;network is unreachable&#34;,
    <a id="L225"></a>ENETDOWN: &#34;network interface is not configured&#34;,
    <a id="L226"></a>ETIMEDOUT: &#34;connection timed out&#34;,
    <a id="L227"></a>EHOSTDOWN: &#34;host is down&#34;,
    <a id="L228"></a>EHOSTUNREACH: &#34;host is unreachable&#34;,
    <a id="L229"></a>EINPROGRESS: &#34;connection already in progress&#34;,
    <a id="L230"></a>EALREADY: &#34;socket already connected&#34;,
    <a id="L231"></a>EDESTADDRREQ: &#34;destination address required&#34;,
    <a id="L232"></a>EPROTONOSUPPORT: &#34;unknown protocol&#34;,
    <a id="L233"></a>ESOCKTNOSUPPORT: &#34;socket type not supported&#34;,
    <a id="L234"></a>EADDRNOTAVAIL: &#34;address not available&#34;,
    <a id="L235"></a>EISCONN: &#34;socket is already connected&#34;,
    <a id="L236"></a>ENOTCONN: &#34;socket is not connected&#34;,
    <a id="L237"></a>ENOMEDIUM: &#34;no medium (in tape drive)&#34;,
    <a id="L238"></a>ENOSHARE: &#34;no such host or network path&#34;,
    <a id="L239"></a>ECASECLASH: &#34;filename exists with different case&#34;,
    <a id="L240"></a>EOVERFLOW: &#34;value too large for defined data type&#34;,
    <a id="L241"></a>ECANCELED: &#34;operation canceled.&#34;,
    <a id="L242"></a>EL2NSYNC: &#34;level 2 not synchronized&#34;,
    <a id="L243"></a>EIDRM: &#34;identifier removed&#34;,
    <a id="L244"></a>EMSGSIZE: &#34;message too long&#34;,
    <a id="L245"></a>ENACL: &#34;not supported by native client&#34;,
<a id="L246"></a>}
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
