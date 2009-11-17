<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/zerrors_linux_amd64.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/syscall/zerrors_linux_amd64.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// mkerrors.sh</span>
<a id="L2"></a><span class="comment">// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT</span>

<a id="L4"></a><span class="comment">// godefs -gsyscall _const.c</span>

<a id="L6"></a><span class="comment">// MACHINE GENERATED - DO NOT EDIT.</span>

<a id="L8"></a>package syscall

<a id="L10"></a><span class="comment">// Constants</span>
<a id="L11"></a>const (
    <a id="L12"></a>AF_APPLETALK                     = 0x5;
    <a id="L13"></a>AF_ASH                           = 0x12;
    <a id="L14"></a>AF_ATMPVC                        = 0x8;
    <a id="L15"></a>AF_ATMSVC                        = 0x14;
    <a id="L16"></a>AF_AX25                          = 0x3;
    <a id="L17"></a>AF_BLUETOOTH                     = 0x1f;
    <a id="L18"></a>AF_BRIDGE                        = 0x7;
    <a id="L19"></a>AF_DECnet                        = 0xc;
    <a id="L20"></a>AF_ECONET                        = 0x13;
    <a id="L21"></a>AF_FILE                          = 0x1;
    <a id="L22"></a>AF_INET                          = 0x2;
    <a id="L23"></a>AF_INET6                         = 0xa;
    <a id="L24"></a>AF_IPX                           = 0x4;
    <a id="L25"></a>AF_IRDA                          = 0x17;
    <a id="L26"></a>AF_IUCV                          = 0x20;
    <a id="L27"></a>AF_KEY                           = 0xf;
    <a id="L28"></a>AF_LOCAL                         = 0x1;
    <a id="L29"></a>AF_MAX                           = 0x22;
    <a id="L30"></a>AF_NETBEUI                       = 0xd;
    <a id="L31"></a>AF_NETLINK                       = 0x10;
    <a id="L32"></a>AF_NETROM                        = 0x6;
    <a id="L33"></a>AF_PACKET                        = 0x11;
    <a id="L34"></a>AF_PPPOX                         = 0x18;
    <a id="L35"></a>AF_ROSE                          = 0xb;
    <a id="L36"></a>AF_ROUTE                         = 0x10;
    <a id="L37"></a>AF_RXRPC                         = 0x21;
    <a id="L38"></a>AF_SECURITY                      = 0xe;
    <a id="L39"></a>AF_SNA                           = 0x16;
    <a id="L40"></a>AF_UNIX                          = 0x1;
    <a id="L41"></a>AF_UNSPEC                        = 0;
    <a id="L42"></a>AF_WANPIPE                       = 0x19;
    <a id="L43"></a>AF_X25                           = 0x9;
    <a id="L44"></a>E2BIG                            = 0x7;
    <a id="L45"></a>EACCES                           = 0xd;
    <a id="L46"></a>EADDRINUSE                       = 0x62;
    <a id="L47"></a>EADDRNOTAVAIL                    = 0x63;
    <a id="L48"></a>EADV                             = 0x44;
    <a id="L49"></a>EAFNOSUPPORT                     = 0x61;
    <a id="L50"></a>EAGAIN                           = 0xb;
    <a id="L51"></a>EALREADY                         = 0x72;
    <a id="L52"></a>EBADE                            = 0x34;
    <a id="L53"></a>EBADF                            = 0x9;
    <a id="L54"></a>EBADFD                           = 0x4d;
    <a id="L55"></a>EBADMSG                          = 0x4a;
    <a id="L56"></a>EBADR                            = 0x35;
    <a id="L57"></a>EBADRQC                          = 0x38;
    <a id="L58"></a>EBADSLT                          = 0x39;
    <a id="L59"></a>EBFONT                           = 0x3b;
    <a id="L60"></a>EBUSY                            = 0x10;
    <a id="L61"></a>ECANCELED                        = 0x7d;
    <a id="L62"></a>ECHILD                           = 0xa;
    <a id="L63"></a>ECHRNG                           = 0x2c;
    <a id="L64"></a>ECOMM                            = 0x46;
    <a id="L65"></a>ECONNABORTED                     = 0x67;
    <a id="L66"></a>ECONNREFUSED                     = 0x6f;
    <a id="L67"></a>ECONNRESET                       = 0x68;
    <a id="L68"></a>EDEADLK                          = 0x23;
    <a id="L69"></a>EDEADLOCK                        = 0x23;
    <a id="L70"></a>EDESTADDRREQ                     = 0x59;
    <a id="L71"></a>EDOM                             = 0x21;
    <a id="L72"></a>EDOTDOT                          = 0x49;
    <a id="L73"></a>EDQUOT                           = 0x7a;
    <a id="L74"></a>EEXIST                           = 0x11;
    <a id="L75"></a>EFAULT                           = 0xe;
    <a id="L76"></a>EFBIG                            = 0x1b;
    <a id="L77"></a>EHOSTDOWN                        = 0x70;
    <a id="L78"></a>EHOSTUNREACH                     = 0x71;
    <a id="L79"></a>EIDRM                            = 0x2b;
    <a id="L80"></a>EILSEQ                           = 0x54;
    <a id="L81"></a>EINPROGRESS                      = 0x73;
    <a id="L82"></a>EINTR                            = 0x4;
    <a id="L83"></a>EINVAL                           = 0x16;
    <a id="L84"></a>EIO                              = 0x5;
    <a id="L85"></a>EISCONN                          = 0x6a;
    <a id="L86"></a>EISDIR                           = 0x15;
    <a id="L87"></a>EISNAM                           = 0x78;
    <a id="L88"></a>EKEYEXPIRED                      = 0x7f;
    <a id="L89"></a>EKEYREJECTED                     = 0x81;
    <a id="L90"></a>EKEYREVOKED                      = 0x80;
    <a id="L91"></a>EL2HLT                           = 0x33;
    <a id="L92"></a>EL2NSYNC                         = 0x2d;
    <a id="L93"></a>EL3HLT                           = 0x2e;
    <a id="L94"></a>EL3RST                           = 0x2f;
    <a id="L95"></a>ELIBACC                          = 0x4f;
    <a id="L96"></a>ELIBBAD                          = 0x50;
    <a id="L97"></a>ELIBEXEC                         = 0x53;
    <a id="L98"></a>ELIBMAX                          = 0x52;
    <a id="L99"></a>ELIBSCN                          = 0x51;
    <a id="L100"></a>ELNRNG                           = 0x30;
    <a id="L101"></a>ELOOP                            = 0x28;
    <a id="L102"></a>EMEDIUMTYPE                      = 0x7c;
    <a id="L103"></a>EMFILE                           = 0x18;
    <a id="L104"></a>EMLINK                           = 0x1f;
    <a id="L105"></a>EMSGSIZE                         = 0x5a;
    <a id="L106"></a>EMULTIHOP                        = 0x48;
    <a id="L107"></a>ENAMETOOLONG                     = 0x24;
    <a id="L108"></a>ENAVAIL                          = 0x77;
    <a id="L109"></a>ENETDOWN                         = 0x64;
    <a id="L110"></a>ENETRESET                        = 0x66;
    <a id="L111"></a>ENETUNREACH                      = 0x65;
    <a id="L112"></a>ENFILE                           = 0x17;
    <a id="L113"></a>ENOANO                           = 0x37;
    <a id="L114"></a>ENOBUFS                          = 0x69;
    <a id="L115"></a>ENOCSI                           = 0x32;
    <a id="L116"></a>ENODATA                          = 0x3d;
    <a id="L117"></a>ENODEV                           = 0x13;
    <a id="L118"></a>ENOENT                           = 0x2;
    <a id="L119"></a>ENOEXEC                          = 0x8;
    <a id="L120"></a>ENOKEY                           = 0x7e;
    <a id="L121"></a>ENOLCK                           = 0x25;
    <a id="L122"></a>ENOLINK                          = 0x43;
    <a id="L123"></a>ENOMEDIUM                        = 0x7b;
    <a id="L124"></a>ENOMEM                           = 0xc;
    <a id="L125"></a>ENOMSG                           = 0x2a;
    <a id="L126"></a>ENONET                           = 0x40;
    <a id="L127"></a>ENOPKG                           = 0x41;
    <a id="L128"></a>ENOPROTOOPT                      = 0x5c;
    <a id="L129"></a>ENOSPC                           = 0x1c;
    <a id="L130"></a>ENOSR                            = 0x3f;
    <a id="L131"></a>ENOSTR                           = 0x3c;
    <a id="L132"></a>ENOSYS                           = 0x26;
    <a id="L133"></a>ENOTBLK                          = 0xf;
    <a id="L134"></a>ENOTCONN                         = 0x6b;
    <a id="L135"></a>ENOTDIR                          = 0x14;
    <a id="L136"></a>ENOTEMPTY                        = 0x27;
    <a id="L137"></a>ENOTNAM                          = 0x76;
    <a id="L138"></a>ENOTRECOVERABLE                  = 0x83;
    <a id="L139"></a>ENOTSOCK                         = 0x58;
    <a id="L140"></a>ENOTSUP                          = 0x5f;
    <a id="L141"></a>ENOTTY                           = 0x19;
    <a id="L142"></a>ENOTUNIQ                         = 0x4c;
    <a id="L143"></a>ENXIO                            = 0x6;
    <a id="L144"></a>EOPNOTSUPP                       = 0x5f;
    <a id="L145"></a>EOVERFLOW                        = 0x4b;
    <a id="L146"></a>EOWNERDEAD                       = 0x82;
    <a id="L147"></a>EPERM                            = 0x1;
    <a id="L148"></a>EPFNOSUPPORT                     = 0x60;
    <a id="L149"></a>EPIPE                            = 0x20;
    <a id="L150"></a>EPOLLERR                         = 0x8;
    <a id="L151"></a>EPOLLET                          = -0x80000000;
    <a id="L152"></a>EPOLLHUP                         = 0x10;
    <a id="L153"></a>EPOLLIN                          = 0x1;
    <a id="L154"></a>EPOLLMSG                         = 0x400;
    <a id="L155"></a>EPOLLONESHOT                     = 0x40000000;
    <a id="L156"></a>EPOLLOUT                         = 0x4;
    <a id="L157"></a>EPOLLPRI                         = 0x2;
    <a id="L158"></a>EPOLLRDBAND                      = 0x80;
    <a id="L159"></a>EPOLLRDHUP                       = 0x2000;
    <a id="L160"></a>EPOLLRDNORM                      = 0x40;
    <a id="L161"></a>EPOLLWRBAND                      = 0x200;
    <a id="L162"></a>EPOLLWRNORM                      = 0x100;
    <a id="L163"></a>EPOLL_CTL_ADD                    = 0x1;
    <a id="L164"></a>EPOLL_CTL_DEL                    = 0x2;
    <a id="L165"></a>EPOLL_CTL_MOD                    = 0x3;
    <a id="L166"></a>EPROTO                           = 0x47;
    <a id="L167"></a>EPROTONOSUPPORT                  = 0x5d;
    <a id="L168"></a>EPROTOTYPE                       = 0x5b;
    <a id="L169"></a>ERANGE                           = 0x22;
    <a id="L170"></a>EREMCHG                          = 0x4e;
    <a id="L171"></a>EREMOTE                          = 0x42;
    <a id="L172"></a>EREMOTEIO                        = 0x79;
    <a id="L173"></a>ERESTART                         = 0x55;
    <a id="L174"></a>EROFS                            = 0x1e;
    <a id="L175"></a>ESHUTDOWN                        = 0x6c;
    <a id="L176"></a>ESOCKTNOSUPPORT                  = 0x5e;
    <a id="L177"></a>ESPIPE                           = 0x1d;
    <a id="L178"></a>ESRCH                            = 0x3;
    <a id="L179"></a>ESRMNT                           = 0x45;
    <a id="L180"></a>ESTALE                           = 0x74;
    <a id="L181"></a>ESTRPIPE                         = 0x56;
    <a id="L182"></a>ETIME                            = 0x3e;
    <a id="L183"></a>ETIMEDOUT                        = 0x6e;
    <a id="L184"></a>ETOOMANYREFS                     = 0x6d;
    <a id="L185"></a>ETXTBSY                          = 0x1a;
    <a id="L186"></a>EUCLEAN                          = 0x75;
    <a id="L187"></a>EUNATCH                          = 0x31;
    <a id="L188"></a>EUSERS                           = 0x57;
    <a id="L189"></a>EWOULDBLOCK                      = 0xb;
    <a id="L190"></a>EXDEV                            = 0x12;
    <a id="L191"></a>EXFULL                           = 0x36;
    <a id="L192"></a>EXPR_NEST_MAX                    = 0x20;
    <a id="L193"></a>FD_CLOEXEC                       = 0x1;
    <a id="L194"></a>FD_SETSIZE                       = 0x400;
    <a id="L195"></a>F_DUPFD                          = 0;
    <a id="L196"></a>F_DUPFD_CLOEXEC                  = 0x406;
    <a id="L197"></a>F_EXLCK                          = 0x4;
    <a id="L198"></a>F_GETFD                          = 0x1;
    <a id="L199"></a>F_GETFL                          = 0x3;
    <a id="L200"></a>F_GETLEASE                       = 0x401;
    <a id="L201"></a>F_GETLK                          = 0x5;
    <a id="L202"></a>F_GETLK64                        = 0x5;
    <a id="L203"></a>F_GETOWN                         = 0x9;
    <a id="L204"></a>F_GETSIG                         = 0xb;
    <a id="L205"></a>F_LOCK                           = 0x1;
    <a id="L206"></a>F_NOTIFY                         = 0x402;
    <a id="L207"></a>F_OK                             = 0;
    <a id="L208"></a>F_RDLCK                          = 0;
    <a id="L209"></a>F_SETFD                          = 0x2;
    <a id="L210"></a>F_SETFL                          = 0x4;
    <a id="L211"></a>F_SETLEASE                       = 0x400;
    <a id="L212"></a>F_SETLK                          = 0x6;
    <a id="L213"></a>F_SETLK64                        = 0x6;
    <a id="L214"></a>F_SETLKW                         = 0x7;
    <a id="L215"></a>F_SETLKW64                       = 0x7;
    <a id="L216"></a>F_SETOWN                         = 0x8;
    <a id="L217"></a>F_SETSIG                         = 0xa;
    <a id="L218"></a>F_SHLCK                          = 0x8;
    <a id="L219"></a>F_TEST                           = 0x3;
    <a id="L220"></a>F_TLOCK                          = 0x2;
    <a id="L221"></a>F_ULOCK                          = 0;
    <a id="L222"></a>F_UNLCK                          = 0x2;
    <a id="L223"></a>F_WRLCK                          = 0x1;
    <a id="L224"></a>IPPROTO_AH                       = 0x33;
    <a id="L225"></a>IPPROTO_COMP                     = 0x6c;
    <a id="L226"></a>IPPROTO_DSTOPTS                  = 0x3c;
    <a id="L227"></a>IPPROTO_EGP                      = 0x8;
    <a id="L228"></a>IPPROTO_ENCAP                    = 0x62;
    <a id="L229"></a>IPPROTO_ESP                      = 0x32;
    <a id="L230"></a>IPPROTO_FRAGMENT                 = 0x2c;
    <a id="L231"></a>IPPROTO_GRE                      = 0x2f;
    <a id="L232"></a>IPPROTO_HOPOPTS                  = 0;
    <a id="L233"></a>IPPROTO_ICMP                     = 0x1;
    <a id="L234"></a>IPPROTO_ICMPV6                   = 0x3a;
    <a id="L235"></a>IPPROTO_IDP                      = 0x16;
    <a id="L236"></a>IPPROTO_IGMP                     = 0x2;
    <a id="L237"></a>IPPROTO_IP                       = 0;
    <a id="L238"></a>IPPROTO_IPIP                     = 0x4;
    <a id="L239"></a>IPPROTO_IPV6                     = 0x29;
    <a id="L240"></a>IPPROTO_MTP                      = 0x5c;
    <a id="L241"></a>IPPROTO_NONE                     = 0x3b;
    <a id="L242"></a>IPPROTO_PIM                      = 0x67;
    <a id="L243"></a>IPPROTO_PUP                      = 0xc;
    <a id="L244"></a>IPPROTO_RAW                      = 0xff;
    <a id="L245"></a>IPPROTO_ROUTING                  = 0x2b;
    <a id="L246"></a>IPPROTO_RSVP                     = 0x2e;
    <a id="L247"></a>IPPROTO_SCTP                     = 0x84;
    <a id="L248"></a>IPPROTO_TCP                      = 0x6;
    <a id="L249"></a>IPPROTO_TP                       = 0x1d;
    <a id="L250"></a>IPPROTO_UDP                      = 0x11;
    <a id="L251"></a>IP_ADD_MEMBERSHIP                = 0x23;
    <a id="L252"></a>IP_ADD_SOURCE_MEMBERSHIP         = 0x27;
    <a id="L253"></a>IP_BLOCK_SOURCE                  = 0x26;
    <a id="L254"></a>IP_DEFAULT_MULTICAST_LOOP        = 0x1;
    <a id="L255"></a>IP_DEFAULT_MULTICAST_TTL         = 0x1;
    <a id="L256"></a>IP_DROP_MEMBERSHIP               = 0x24;
    <a id="L257"></a>IP_DROP_SOURCE_MEMBERSHIP        = 0x28;
    <a id="L258"></a>IP_HDRINCL                       = 0x3;
    <a id="L259"></a>IP_MAX_MEMBERSHIPS               = 0x14;
    <a id="L260"></a>IP_MSFILTER                      = 0x29;
    <a id="L261"></a>IP_MTU_DISCOVER                  = 0xa;
    <a id="L262"></a>IP_MULTICAST_IF                  = 0x20;
    <a id="L263"></a>IP_MULTICAST_LOOP                = 0x22;
    <a id="L264"></a>IP_MULTICAST_TTL                 = 0x21;
    <a id="L265"></a>IP_OPTIONS                       = 0x4;
    <a id="L266"></a>IP_PKTINFO                       = 0x8;
    <a id="L267"></a>IP_PKTOPTIONS                    = 0x9;
    <a id="L268"></a>IP_PMTUDISC                      = 0xa;
    <a id="L269"></a>IP_PMTUDISC_DO                   = 0x2;
    <a id="L270"></a>IP_PMTUDISC_DONT                 = 0;
    <a id="L271"></a>IP_PMTUDISC_WANT                 = 0x1;
    <a id="L272"></a>IP_RECVERR                       = 0xb;
    <a id="L273"></a>IP_RECVOPTS                      = 0x6;
    <a id="L274"></a>IP_RECVRETOPTS                   = 0x7;
    <a id="L275"></a>IP_RECVTOS                       = 0xd;
    <a id="L276"></a>IP_RECVTTL                       = 0xc;
    <a id="L277"></a>IP_RETOPTS                       = 0x7;
    <a id="L278"></a>IP_ROUTER_ALERT                  = 0x5;
    <a id="L279"></a>IP_TOS                           = 0x1;
    <a id="L280"></a>IP_TTL                           = 0x2;
    <a id="L281"></a>IP_UNBLOCK_SOURCE                = 0x25;
    <a id="L282"></a>NAME_MAX                         = 0xff;
    <a id="L283"></a>O_ACCMODE                        = 0x3;
    <a id="L284"></a>O_APPEND                         = 0x400;
    <a id="L285"></a>O_ASYNC                          = 0x2000;
    <a id="L286"></a>O_CLOEXEC                        = 0x80000;
    <a id="L287"></a>O_CREAT                          = 0x40;
    <a id="L288"></a>O_DIRECT                         = 0x4000;
    <a id="L289"></a>O_DIRECTORY                      = 0x10000;
    <a id="L290"></a>O_DSYNC                          = 0x1000;
    <a id="L291"></a>O_EXCL                           = 0x80;
    <a id="L292"></a>O_FSYNC                          = 0x1000;
    <a id="L293"></a>O_LARGEFILE                      = 0;
    <a id="L294"></a>O_NDELAY                         = 0x800;
    <a id="L295"></a>O_NOATIME                        = 0x40000;
    <a id="L296"></a>O_NOCTTY                         = 0x100;
    <a id="L297"></a>O_NOFOLLOW                       = 0x20000;
    <a id="L298"></a>O_NONBLOCK                       = 0x800;
    <a id="L299"></a>O_RDONLY                         = 0;
    <a id="L300"></a>O_RDWR                           = 0x2;
    <a id="L301"></a>O_RSYNC                          = 0x1000;
    <a id="L302"></a>O_SYNC                           = 0x1000;
    <a id="L303"></a>O_TRUNC                          = 0x200;
    <a id="L304"></a>O_WRONLY                         = 0x1;
    <a id="L305"></a>PTRACE_ARCH_PRCTL                = 0x1e;
    <a id="L306"></a>PTRACE_ATTACH                    = 0x10;
    <a id="L307"></a>PTRACE_CONT                      = 0x7;
    <a id="L308"></a>PTRACE_DETACH                    = 0x11;
    <a id="L309"></a>PTRACE_EVENT_CLONE               = 0x3;
    <a id="L310"></a>PTRACE_EVENT_EXEC                = 0x4;
    <a id="L311"></a>PTRACE_EVENT_EXIT                = 0x6;
    <a id="L312"></a>PTRACE_EVENT_FORK                = 0x1;
    <a id="L313"></a>PTRACE_EVENT_VFORK               = 0x2;
    <a id="L314"></a>PTRACE_EVENT_VFORK_DONE          = 0x5;
    <a id="L315"></a>PTRACE_GETEVENTMSG               = 0x4201;
    <a id="L316"></a>PTRACE_GETFPREGS                 = 0xe;
    <a id="L317"></a>PTRACE_GETFPXREGS                = 0x12;
    <a id="L318"></a>PTRACE_GETREGS                   = 0xc;
    <a id="L319"></a>PTRACE_GETSIGINFO                = 0x4202;
    <a id="L320"></a>PTRACE_GET_THREAD_AREA           = 0x19;
    <a id="L321"></a>PTRACE_KILL                      = 0x8;
    <a id="L322"></a>PTRACE_OLDSETOPTIONS             = 0x15;
    <a id="L323"></a>PTRACE_O_MASK                    = 0x7f;
    <a id="L324"></a>PTRACE_O_TRACECLONE              = 0x8;
    <a id="L325"></a>PTRACE_O_TRACEEXEC               = 0x10;
    <a id="L326"></a>PTRACE_O_TRACEEXIT               = 0x40;
    <a id="L327"></a>PTRACE_O_TRACEFORK               = 0x2;
    <a id="L328"></a>PTRACE_O_TRACESYSGOOD            = 0x1;
    <a id="L329"></a>PTRACE_O_TRACEVFORK              = 0x4;
    <a id="L330"></a>PTRACE_O_TRACEVFORKDONE          = 0x20;
    <a id="L331"></a>PTRACE_PEEKDATA                  = 0x2;
    <a id="L332"></a>PTRACE_PEEKTEXT                  = 0x1;
    <a id="L333"></a>PTRACE_PEEKUSR                   = 0x3;
    <a id="L334"></a>PTRACE_POKEDATA                  = 0x5;
    <a id="L335"></a>PTRACE_POKETEXT                  = 0x4;
    <a id="L336"></a>PTRACE_POKEUSR                   = 0x6;
    <a id="L337"></a>PTRACE_SETFPREGS                 = 0xf;
    <a id="L338"></a>PTRACE_SETFPXREGS                = 0x13;
    <a id="L339"></a>PTRACE_SETOPTIONS                = 0x4200;
    <a id="L340"></a>PTRACE_SETREGS                   = 0xd;
    <a id="L341"></a>PTRACE_SETSIGINFO                = 0x4203;
    <a id="L342"></a>PTRACE_SET_THREAD_AREA           = 0x1a;
    <a id="L343"></a>PTRACE_SINGLESTEP                = 0x9;
    <a id="L344"></a>PTRACE_SYSCALL                   = 0x18;
    <a id="L345"></a>PTRACE_TRACEME                   = 0;
    <a id="L346"></a>SIGABRT                          = 0x6;
    <a id="L347"></a>SIGALRM                          = 0xe;
    <a id="L348"></a>SIGBUS                           = 0x7;
    <a id="L349"></a>SIGCHLD                          = 0x11;
    <a id="L350"></a>SIGCLD                           = 0x11;
    <a id="L351"></a>SIGCONT                          = 0x12;
    <a id="L352"></a>SIGFPE                           = 0x8;
    <a id="L353"></a>SIGHUP                           = 0x1;
    <a id="L354"></a>SIGILL                           = 0x4;
    <a id="L355"></a>SIGINT                           = 0x2;
    <a id="L356"></a>SIGIO                            = 0x1d;
    <a id="L357"></a>SIGIOT                           = 0x6;
    <a id="L358"></a>SIGKILL                          = 0x9;
    <a id="L359"></a>SIGPIPE                          = 0xd;
    <a id="L360"></a>SIGPOLL                          = 0x1d;
    <a id="L361"></a>SIGPROF                          = 0x1b;
    <a id="L362"></a>SIGPWR                           = 0x1e;
    <a id="L363"></a>SIGQUIT                          = 0x3;
    <a id="L364"></a>SIGSEGV                          = 0xb;
    <a id="L365"></a>SIGSTKFLT                        = 0x10;
    <a id="L366"></a>SIGSTOP                          = 0x13;
    <a id="L367"></a>SIGSYS                           = 0x1f;
    <a id="L368"></a>SIGTERM                          = 0xf;
    <a id="L369"></a>SIGTRAP                          = 0x5;
    <a id="L370"></a>SIGTSTP                          = 0x14;
    <a id="L371"></a>SIGTTIN                          = 0x15;
    <a id="L372"></a>SIGTTOU                          = 0x16;
    <a id="L373"></a>SIGUNUSED                        = 0x1f;
    <a id="L374"></a>SIGURG                           = 0x17;
    <a id="L375"></a>SIGUSR1                          = 0xa;
    <a id="L376"></a>SIGUSR2                          = 0xc;
    <a id="L377"></a>SIGVTALRM                        = 0x1a;
    <a id="L378"></a>SIGWINCH                         = 0x1c;
    <a id="L379"></a>SIGXCPU                          = 0x18;
    <a id="L380"></a>SIGXFSZ                          = 0x19;
    <a id="L381"></a>SOCK_DGRAM                       = 0x2;
    <a id="L382"></a>SOCK_PACKET                      = 0xa;
    <a id="L383"></a>SOCK_RAW                         = 0x3;
    <a id="L384"></a>SOCK_RDM                         = 0x4;
    <a id="L385"></a>SOCK_SEQPACKET                   = 0x5;
    <a id="L386"></a>SOCK_STREAM                      = 0x1;
    <a id="L387"></a>SOL_AAL                          = 0x109;
    <a id="L388"></a>SOL_ATM                          = 0x108;
    <a id="L389"></a>SOL_DECNET                       = 0x105;
    <a id="L390"></a>SOL_ICMPV6                       = 0x3a;
    <a id="L391"></a>SOL_IP                           = 0;
    <a id="L392"></a>SOL_IPV6                         = 0x29;
    <a id="L393"></a>SOL_IRDA                         = 0x10a;
    <a id="L394"></a>SOL_PACKET                       = 0x107;
    <a id="L395"></a>SOL_RAW                          = 0xff;
    <a id="L396"></a>SOL_SOCKET                       = 0x1;
    <a id="L397"></a>SOL_TCP                          = 0x6;
    <a id="L398"></a>SOL_X25                          = 0x106;
    <a id="L399"></a>SOMAXCONN                        = 0x80;
    <a id="L400"></a>SO_ACCEPTCONN                    = 0x1e;
    <a id="L401"></a>SO_ATTACH_FILTER                 = 0x1a;
    <a id="L402"></a>SO_BINDTODEVICE                  = 0x19;
    <a id="L403"></a>SO_BROADCAST                     = 0x6;
    <a id="L404"></a>SO_BSDCOMPAT                     = 0xe;
    <a id="L405"></a>SO_DEBUG                         = 0x1;
    <a id="L406"></a>SO_DETACH_FILTER                 = 0x1b;
    <a id="L407"></a>SO_DONTROUTE                     = 0x5;
    <a id="L408"></a>SO_ERROR                         = 0x4;
    <a id="L409"></a>SO_KEEPALIVE                     = 0x9;
    <a id="L410"></a>SO_LINGER                        = 0xd;
    <a id="L411"></a>SO_NO_CHECK                      = 0xb;
    <a id="L412"></a>SO_OOBINLINE                     = 0xa;
    <a id="L413"></a>SO_PASSCRED                      = 0x10;
    <a id="L414"></a>SO_PASSSEC                       = 0x22;
    <a id="L415"></a>SO_PEERCRED                      = 0x11;
    <a id="L416"></a>SO_PEERNAME                      = 0x1c;
    <a id="L417"></a>SO_PEERSEC                       = 0x1f;
    <a id="L418"></a>SO_PRIORITY                      = 0xc;
    <a id="L419"></a>SO_RCVBUF                        = 0x8;
    <a id="L420"></a>SO_RCVBUFFORCE                   = 0x21;
    <a id="L421"></a>SO_RCVLOWAT                      = 0x12;
    <a id="L422"></a>SO_RCVTIMEO                      = 0x14;
    <a id="L423"></a>SO_REUSEADDR                     = 0x2;
    <a id="L424"></a>SO_SECURITY_AUTHENTICATION       = 0x16;
    <a id="L425"></a>SO_SECURITY_ENCRYPTION_NETWORK   = 0x18;
    <a id="L426"></a>SO_SECURITY_ENCRYPTION_TRANSPORT = 0x17;
    <a id="L427"></a>SO_SNDBUF                        = 0x7;
    <a id="L428"></a>SO_SNDBUFFORCE                   = 0x20;
    <a id="L429"></a>SO_SNDLOWAT                      = 0x13;
    <a id="L430"></a>SO_SNDTIMEO                      = 0x15;
    <a id="L431"></a>SO_TIMESTAMP                     = 0x1d;
    <a id="L432"></a>SO_TIMESTAMPNS                   = 0x23;
    <a id="L433"></a>SO_TYPE                          = 0x3;
    <a id="L434"></a>S_BLKSIZE                        = 0x200;
    <a id="L435"></a>S_IEXEC                          = 0x40;
    <a id="L436"></a>S_IFBLK                          = 0x6000;
    <a id="L437"></a>S_IFCHR                          = 0x2000;
    <a id="L438"></a>S_IFDIR                          = 0x4000;
    <a id="L439"></a>S_IFIFO                          = 0x1000;
    <a id="L440"></a>S_IFLNK                          = 0xa000;
    <a id="L441"></a>S_IFMT                           = 0xf000;
    <a id="L442"></a>S_IFREG                          = 0x8000;
    <a id="L443"></a>S_IFSOCK                         = 0xc000;
    <a id="L444"></a>S_IREAD                          = 0x100;
    <a id="L445"></a>S_IRGRP                          = 0x20;
    <a id="L446"></a>S_IROTH                          = 0x4;
    <a id="L447"></a>S_IRUSR                          = 0x100;
    <a id="L448"></a>S_IRWXG                          = 0x38;
    <a id="L449"></a>S_IRWXO                          = 0x7;
    <a id="L450"></a>S_IRWXU                          = 0x1c0;
    <a id="L451"></a>S_ISGID                          = 0x400;
    <a id="L452"></a>S_ISUID                          = 0x800;
    <a id="L453"></a>S_ISVTX                          = 0x200;
    <a id="L454"></a>S_IWGRP                          = 0x10;
    <a id="L455"></a>S_IWOTH                          = 0x2;
    <a id="L456"></a>S_IWRITE                         = 0x80;
    <a id="L457"></a>S_IWUSR                          = 0x80;
    <a id="L458"></a>S_IXGRP                          = 0x8;
    <a id="L459"></a>S_IXOTH                          = 0x1;
    <a id="L460"></a>S_IXUSR                          = 0x40;
    <a id="L461"></a>TCP_CONGESTION                   = 0xd;
    <a id="L462"></a>TCP_CORK                         = 0x3;
    <a id="L463"></a>TCP_DEFER_ACCEPT                 = 0x9;
    <a id="L464"></a>TCP_INFO                         = 0xb;
    <a id="L465"></a>TCP_KEEPCNT                      = 0x6;
    <a id="L466"></a>TCP_KEEPIDLE                     = 0x4;
    <a id="L467"></a>TCP_KEEPINTVL                    = 0x5;
    <a id="L468"></a>TCP_LINGER2                      = 0x8;
    <a id="L469"></a>TCP_MAXSEG                       = 0x2;
    <a id="L470"></a>TCP_MAXWIN                       = 0xffff;
    <a id="L471"></a>TCP_MAX_WINSHIFT                 = 0xe;
    <a id="L472"></a>TCP_MD5SIG                       = 0xe;
    <a id="L473"></a>TCP_MD5SIG_MAXKEYLEN             = 0x50;
    <a id="L474"></a>TCP_MSS                          = 0x200;
    <a id="L475"></a>TCP_NODELAY                      = 0x1;
    <a id="L476"></a>TCP_QUICKACK                     = 0xc;
    <a id="L477"></a>TCP_SYNCNT                       = 0x7;
    <a id="L478"></a>TCP_WINDOW_CLAMP                 = 0xa;
    <a id="L479"></a>WALL                             = 0x40000000;
    <a id="L480"></a>WCLONE                           = 0x80000000;
    <a id="L481"></a>WCONTINUED                       = 0x8;
    <a id="L482"></a>WEXITED                          = 0x4;
    <a id="L483"></a>WNOHANG                          = 0x1;
    <a id="L484"></a>WNOTHREAD                        = 0x20000000;
    <a id="L485"></a>WNOWAIT                          = 0x1000000;
    <a id="L486"></a>WORDSIZE                         = 0x40;
    <a id="L487"></a>WSTOPPED                         = 0x2;
    <a id="L488"></a>WUNTRACED                        = 0x2;
<a id="L489"></a>)

<a id="L491"></a><span class="comment">// Types</span>


<a id="L494"></a><span class="comment">// Error table</span>
<a id="L495"></a>var errors = [...]string{
    <a id="L496"></a>72: &#34;multihop attempted&#34;,
    <a id="L497"></a>49: &#34;protocol driver not attached&#34;,
    <a id="L498"></a>97: &#34;address family not supported by protocol&#34;,
    <a id="L499"></a>78: &#34;remote address changed&#34;,
    <a id="L500"></a>13: &#34;permission denied&#34;,
    <a id="L501"></a>47: &#34;level 3 reset&#34;,
    <a id="L502"></a>89: &#34;destination address required&#34;,
    <a id="L503"></a>84: &#34;invalid or incomplete multibyte or wide character&#34;,
    <a id="L504"></a>29: &#34;illegal seek&#34;,
    <a id="L505"></a>31: &#34;too many links&#34;,
    <a id="L506"></a>130: &#34;owner died&#34;,
    <a id="L507"></a>25: &#34;inappropriate ioctl for device&#34;,
    <a id="L508"></a>52: &#34;invalid exchange&#34;,
    <a id="L509"></a>9: &#34;bad file descriptor&#34;,
    <a id="L510"></a>53: &#34;invalid request descriptor&#34;,
    <a id="L511"></a>68: &#34;advertise error&#34;,
    <a id="L512"></a>34: &#34;numerical result out of range&#34;,
    <a id="L513"></a>125: &#34;operation canceled&#34;,
    <a id="L514"></a>26: &#34;text file busy&#34;,
    <a id="L515"></a>12: &#34;cannot allocate memory&#34;,
    <a id="L516"></a>115: &#34;operation now in progress&#34;,
    <a id="L517"></a>15: &#34;block device required&#34;,
    <a id="L518"></a>91: &#34;protocol wrong type for socket&#34;,
    <a id="L519"></a>85: &#34;interrupted system call should be restarted&#34;,
    <a id="L520"></a>120: &#34;is a named type file&#34;,
    <a id="L521"></a>42: &#34;no message of desired type&#34;,
    <a id="L522"></a>114: &#34;operation already in progress&#34;,
    <a id="L523"></a>110: &#34;connection timed out&#34;,
    <a id="L524"></a>61: &#34;no data available&#34;,
    <a id="L525"></a>4: &#34;interrupted system call&#34;,
    <a id="L526"></a>67: &#34;link has been severed&#34;,
    <a id="L527"></a>1: &#34;operation not permitted&#34;,
    <a id="L528"></a>40: &#34;too many levels of symbolic links&#34;,
    <a id="L529"></a>100: &#34;network is down&#34;,
    <a id="L530"></a>116: &#34;stale NFS file handle&#34;,
    <a id="L531"></a>88: &#34;socket operation on non-socket&#34;,
    <a id="L532"></a>63: &#34;out of streams resources&#34;,
    <a id="L533"></a>10: &#34;no child processes&#34;,
    <a id="L534"></a>48: &#34;link number out of range&#34;,
    <a id="L535"></a>32: &#34;broken pipe&#34;,
    <a id="L536"></a>74: &#34;bad message&#34;,
    <a id="L537"></a>59: &#34;bad font file format&#34;,
    <a id="L538"></a>66: &#34;object is remote&#34;,
    <a id="L539"></a>109: &#34;too many references: cannot splice&#34;,
    <a id="L540"></a>96: &#34;protocol family not supported&#34;,
    <a id="L541"></a>64: &#34;machine is not on the network&#34;,
    <a id="L542"></a>54: &#34;exchange full&#34;,
    <a id="L543"></a>57: &#34;invalid slot&#34;,
    <a id="L544"></a>118: &#34;not a XENIX named type file&#34;,
    <a id="L545"></a>83: &#34;cannot exec a shared library directly&#34;,
    <a id="L546"></a>50: &#34;no CSI structure available&#34;,
    <a id="L547"></a>39: &#34;directory not empty&#34;,
    <a id="L548"></a>98: &#34;address already in use&#34;,
    <a id="L549"></a>102: &#34;network dropped connection on reset&#34;,
    <a id="L550"></a>21: &#34;is a directory&#34;,
    <a id="L551"></a>43: &#34;identifier removed&#34;,
    <a id="L552"></a>70: &#34;communication error on send&#34;,
    <a id="L553"></a>77: &#34;file descriptor in bad state&#34;,
    <a id="L554"></a>51: &#34;level 2 halted&#34;,
    <a id="L555"></a>126: &#34;required key not available&#34;,
    <a id="L556"></a>22: &#34;invalid argument&#34;,
    <a id="L557"></a>108: &#34;cannot send after transport endpoint shutdown&#34;,
    <a id="L558"></a>129: &#34;key was rejected by service&#34;,
    <a id="L559"></a>81: &#34;.lib section in a.out corrupted&#34;,
    <a id="L560"></a>119: &#34;no XENIX semaphores available&#34;,
    <a id="L561"></a>60: &#34;device not a stream&#34;,
    <a id="L562"></a>75: &#34;value too large for defined data type&#34;,
    <a id="L563"></a>117: &#34;structure needs cleaning&#34;,
    <a id="L564"></a>123: &#34;no medium found&#34;,
    <a id="L565"></a>16: &#34;device or resource busy&#34;,
    <a id="L566"></a>71: &#34;protocol error&#34;,
    <a id="L567"></a>19: &#34;no such device&#34;,
    <a id="L568"></a>127: &#34;key has expired&#34;,
    <a id="L569"></a>30: &#34;read-only file system&#34;,
    <a id="L570"></a>79: &#34;can not access a needed shared library&#34;,
    <a id="L571"></a>7: &#34;argument list too long&#34;,
    <a id="L572"></a>35: &#34;resource deadlock avoided&#34;,
    <a id="L573"></a>104: &#34;connection reset by peer&#34;,
    <a id="L574"></a>6: &#34;no such device or address&#34;,
    <a id="L575"></a>56: &#34;invalid request code&#34;,
    <a id="L576"></a>36: &#34;file name too long&#34;,
    <a id="L577"></a>94: &#34;socket type not supported&#34;,
    <a id="L578"></a>73: &#34;RFS specific error&#34;,
    <a id="L579"></a>99: &#34;cannot assign requested address&#34;,
    <a id="L580"></a>62: &#34;timer expired&#34;,
    <a id="L581"></a>93: &#34;protocol not supported&#34;,
    <a id="L582"></a>131: &#34;state not recoverable&#34;,
    <a id="L583"></a>5: &#34;input/output error&#34;,
    <a id="L584"></a>101: &#34;network is unreachable&#34;,
    <a id="L585"></a>18: &#34;invalid cross-device link&#34;,
    <a id="L586"></a>122: &#34;disk quota exceeded&#34;,
    <a id="L587"></a>121: &#34;remote I/O error&#34;,
    <a id="L588"></a>28: &#34;no space left on device&#34;,
    <a id="L589"></a>8: &#34;exec format error&#34;,
    <a id="L590"></a>90: &#34;message too long&#34;,
    <a id="L591"></a>33: &#34;numerical argument out of domain&#34;,
    <a id="L592"></a>27: &#34;file too large&#34;,
    <a id="L593"></a>3: &#34;no such process&#34;,
    <a id="L594"></a>44: &#34;channel number out of range&#34;,
    <a id="L595"></a>112: &#34;host is down&#34;,
    <a id="L596"></a>37: &#34;no locks available&#34;,
    <a id="L597"></a>23: &#34;too many open files in system&#34;,
    <a id="L598"></a>38: &#34;function not implemented&#34;,
    <a id="L599"></a>107: &#34;transport endpoint is not connected&#34;,
    <a id="L600"></a>95: &#34;operation not supported&#34;,
    <a id="L601"></a>69: &#34;srmount error&#34;,
    <a id="L602"></a>103: &#34;software caused connection abort&#34;,
    <a id="L603"></a>55: &#34;no anode&#34;,
    <a id="L604"></a>106: &#34;transport endpoint is already connected&#34;,
    <a id="L605"></a>87: &#34;too many users&#34;,
    <a id="L606"></a>92: &#34;protocol not available&#34;,
    <a id="L607"></a>24: &#34;too many open files&#34;,
    <a id="L608"></a>105: &#34;no buffer space available&#34;,
    <a id="L609"></a>46: &#34;level 3 halted&#34;,
    <a id="L610"></a>14: &#34;bad address&#34;,
    <a id="L611"></a>11: &#34;resource temporarily unavailable&#34;,
    <a id="L612"></a>80: &#34;accessing a corrupted shared library&#34;,
    <a id="L613"></a>86: &#34;streams pipe error&#34;,
    <a id="L614"></a>111: &#34;connection refused&#34;,
    <a id="L615"></a>82: &#34;attempting to link in too many shared libraries&#34;,
    <a id="L616"></a>17: &#34;file exists&#34;,
    <a id="L617"></a>45: &#34;level 2 not synchronized&#34;,
    <a id="L618"></a>2: &#34;no such file or directory&#34;,
    <a id="L619"></a>65: &#34;package not installed&#34;,
    <a id="L620"></a>128: &#34;key has been revoked&#34;,
    <a id="L621"></a>113: &#34;no route to host&#34;,
    <a id="L622"></a>76: &#34;name not unique on network&#34;,
    <a id="L623"></a>20: &#34;not a directory&#34;,
    <a id="L624"></a>124: &#34;wrong medium type&#34;,
<a id="L625"></a>}
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
