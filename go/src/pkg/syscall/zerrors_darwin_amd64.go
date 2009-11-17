<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/zerrors_darwin_amd64.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/syscall/zerrors_darwin_amd64.go</h1>

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
    <a id="L12"></a>AF_APPLETALK              = 0x10;
    <a id="L13"></a>AF_CCITT                  = 0xa;
    <a id="L14"></a>AF_CHAOS                  = 0x5;
    <a id="L15"></a>AF_CNT                    = 0x15;
    <a id="L16"></a>AF_COIP                   = 0x14;
    <a id="L17"></a>AF_DATAKIT                = 0x9;
    <a id="L18"></a>AF_DECnet                 = 0xc;
    <a id="L19"></a>AF_DLI                    = 0xd;
    <a id="L20"></a>AF_E164                   = 0x1c;
    <a id="L21"></a>AF_ECMA                   = 0x8;
    <a id="L22"></a>AF_HYLINK                 = 0xf;
    <a id="L23"></a>AF_IMPLINK                = 0x3;
    <a id="L24"></a>AF_INET                   = 0x2;
    <a id="L25"></a>AF_INET6                  = 0x1e;
    <a id="L26"></a>AF_IPX                    = 0x17;
    <a id="L27"></a>AF_ISDN                   = 0x1c;
    <a id="L28"></a>AF_ISO                    = 0x7;
    <a id="L29"></a>AF_LAT                    = 0xe;
    <a id="L30"></a>AF_LINK                   = 0x12;
    <a id="L31"></a>AF_LOCAL                  = 0x1;
    <a id="L32"></a>AF_MAX                    = 0x25;
    <a id="L33"></a>AF_NATM                   = 0x1f;
    <a id="L34"></a>AF_NDRV                   = 0x1b;
    <a id="L35"></a>AF_NETBIOS                = 0x21;
    <a id="L36"></a>AF_NS                     = 0x6;
    <a id="L37"></a>AF_OSI                    = 0x7;
    <a id="L38"></a>AF_PPP                    = 0x22;
    <a id="L39"></a>AF_PUP                    = 0x4;
    <a id="L40"></a>AF_RESERVED_36            = 0x24;
    <a id="L41"></a>AF_ROUTE                  = 0x11;
    <a id="L42"></a>AF_SIP                    = 0x18;
    <a id="L43"></a>AF_SNA                    = 0xb;
    <a id="L44"></a>AF_SYSTEM                 = 0x20;
    <a id="L45"></a>AF_UNIX                   = 0x1;
    <a id="L46"></a>AF_UNSPEC                 = 0;
    <a id="L47"></a>E2BIG                     = 0x7;
    <a id="L48"></a>EACCES                    = 0xd;
    <a id="L49"></a>EADDRINUSE                = 0x30;
    <a id="L50"></a>EADDRNOTAVAIL             = 0x31;
    <a id="L51"></a>EAFNOSUPPORT              = 0x2f;
    <a id="L52"></a>EAGAIN                    = 0x23;
    <a id="L53"></a>EALREADY                  = 0x25;
    <a id="L54"></a>EAUTH                     = 0x50;
    <a id="L55"></a>EBADARCH                  = 0x56;
    <a id="L56"></a>EBADEXEC                  = 0x55;
    <a id="L57"></a>EBADF                     = 0x9;
    <a id="L58"></a>EBADMACHO                 = 0x58;
    <a id="L59"></a>EBADMSG                   = 0x5e;
    <a id="L60"></a>EBADRPC                   = 0x48;
    <a id="L61"></a>EBUSY                     = 0x10;
    <a id="L62"></a>ECANCELED                 = 0x59;
    <a id="L63"></a>ECHILD                    = 0xa;
    <a id="L64"></a>ECONNABORTED              = 0x35;
    <a id="L65"></a>ECONNREFUSED              = 0x3d;
    <a id="L66"></a>ECONNRESET                = 0x36;
    <a id="L67"></a>EDEADLK                   = 0xb;
    <a id="L68"></a>EDESTADDRREQ              = 0x27;
    <a id="L69"></a>EDEVERR                   = 0x53;
    <a id="L70"></a>EDOM                      = 0x21;
    <a id="L71"></a>EDQUOT                    = 0x45;
    <a id="L72"></a>EEXIST                    = 0x11;
    <a id="L73"></a>EFAULT                    = 0xe;
    <a id="L74"></a>EFBIG                     = 0x1b;
    <a id="L75"></a>EFTYPE                    = 0x4f;
    <a id="L76"></a>EHOSTDOWN                 = 0x40;
    <a id="L77"></a>EHOSTUNREACH              = 0x41;
    <a id="L78"></a>EIDRM                     = 0x5a;
    <a id="L79"></a>EILSEQ                    = 0x5c;
    <a id="L80"></a>EINPROGRESS               = 0x24;
    <a id="L81"></a>EINTR                     = 0x4;
    <a id="L82"></a>EINVAL                    = 0x16;
    <a id="L83"></a>EIO                       = 0x5;
    <a id="L84"></a>EISCONN                   = 0x38;
    <a id="L85"></a>EISDIR                    = 0x15;
    <a id="L86"></a>ELAST                     = 0x67;
    <a id="L87"></a>ELOOP                     = 0x3e;
    <a id="L88"></a>EMFILE                    = 0x18;
    <a id="L89"></a>EMLINK                    = 0x1f;
    <a id="L90"></a>EMSGSIZE                  = 0x28;
    <a id="L91"></a>EMULTIHOP                 = 0x5f;
    <a id="L92"></a>ENAMETOOLONG              = 0x3f;
    <a id="L93"></a>ENEEDAUTH                 = 0x51;
    <a id="L94"></a>ENETDOWN                  = 0x32;
    <a id="L95"></a>ENETRESET                 = 0x34;
    <a id="L96"></a>ENETUNREACH               = 0x33;
    <a id="L97"></a>ENFILE                    = 0x17;
    <a id="L98"></a>ENOATTR                   = 0x5d;
    <a id="L99"></a>ENOBUFS                   = 0x37;
    <a id="L100"></a>ENODATA                   = 0x60;
    <a id="L101"></a>ENODEV                    = 0x13;
    <a id="L102"></a>ENOENT                    = 0x2;
    <a id="L103"></a>ENOEXEC                   = 0x8;
    <a id="L104"></a>ENOLCK                    = 0x4d;
    <a id="L105"></a>ENOLINK                   = 0x61;
    <a id="L106"></a>ENOMEM                    = 0xc;
    <a id="L107"></a>ENOMSG                    = 0x5b;
    <a id="L108"></a>ENOPOLICY                 = 0x67;
    <a id="L109"></a>ENOPROTOOPT               = 0x2a;
    <a id="L110"></a>ENOSPC                    = 0x1c;
    <a id="L111"></a>ENOSR                     = 0x62;
    <a id="L112"></a>ENOSTR                    = 0x63;
    <a id="L113"></a>ENOSYS                    = 0x4e;
    <a id="L114"></a>ENOTBLK                   = 0xf;
    <a id="L115"></a>ENOTCONN                  = 0x39;
    <a id="L116"></a>ENOTDIR                   = 0x14;
    <a id="L117"></a>ENOTEMPTY                 = 0x42;
    <a id="L118"></a>ENOTSOCK                  = 0x26;
    <a id="L119"></a>ENOTSUP                   = 0x2d;
    <a id="L120"></a>ENOTTY                    = 0x19;
    <a id="L121"></a>ENXIO                     = 0x6;
    <a id="L122"></a>EOPNOTSUPP                = 0x66;
    <a id="L123"></a>EOVERFLOW                 = 0x54;
    <a id="L124"></a>EPERM                     = 0x1;
    <a id="L125"></a>EPFNOSUPPORT              = 0x2e;
    <a id="L126"></a>EPIPE                     = 0x20;
    <a id="L127"></a>EPROCLIM                  = 0x43;
    <a id="L128"></a>EPROCUNAVAIL              = 0x4c;
    <a id="L129"></a>EPROGMISMATCH             = 0x4b;
    <a id="L130"></a>EPROGUNAVAIL              = 0x4a;
    <a id="L131"></a>EPROTO                    = 0x64;
    <a id="L132"></a>EPROTONOSUPPORT           = 0x2b;
    <a id="L133"></a>EPROTOTYPE                = 0x29;
    <a id="L134"></a>EPWROFF                   = 0x52;
    <a id="L135"></a>ERANGE                    = 0x22;
    <a id="L136"></a>EREMOTE                   = 0x47;
    <a id="L137"></a>EROFS                     = 0x1e;
    <a id="L138"></a>ERPCMISMATCH              = 0x49;
    <a id="L139"></a>ESHLIBVERS                = 0x57;
    <a id="L140"></a>ESHUTDOWN                 = 0x3a;
    <a id="L141"></a>ESOCKTNOSUPPORT           = 0x2c;
    <a id="L142"></a>ESPIPE                    = 0x1d;
    <a id="L143"></a>ESRCH                     = 0x3;
    <a id="L144"></a>ESTALE                    = 0x46;
    <a id="L145"></a>ETIME                     = 0x65;
    <a id="L146"></a>ETIMEDOUT                 = 0x3c;
    <a id="L147"></a>ETOOMANYREFS              = 0x3b;
    <a id="L148"></a>ETXTBSY                   = 0x1a;
    <a id="L149"></a>EUSERS                    = 0x44;
    <a id="L150"></a>EVFILT_AIO                = -0x3;
    <a id="L151"></a>EVFILT_FS                 = -0x9;
    <a id="L152"></a>EVFILT_MACHPORT           = -0x8;
    <a id="L153"></a>EVFILT_PROC               = -0x5;
    <a id="L154"></a>EVFILT_READ               = -0x1;
    <a id="L155"></a>EVFILT_SIGNAL             = -0x6;
    <a id="L156"></a>EVFILT_SYSCOUNT           = 0x9;
    <a id="L157"></a>EVFILT_THREADMARKER       = 0x9;
    <a id="L158"></a>EVFILT_TIMER              = -0x7;
    <a id="L159"></a>EVFILT_VNODE              = -0x4;
    <a id="L160"></a>EVFILT_WRITE              = -0x2;
    <a id="L161"></a>EV_ADD                    = 0x1;
    <a id="L162"></a>EV_CLEAR                  = 0x20;
    <a id="L163"></a>EV_DELETE                 = 0x2;
    <a id="L164"></a>EV_DISABLE                = 0x8;
    <a id="L165"></a>EV_ENABLE                 = 0x4;
    <a id="L166"></a>EV_EOF                    = 0x8000;
    <a id="L167"></a>EV_ERROR                  = 0x4000;
    <a id="L168"></a>EV_FLAG0                  = 0x1000;
    <a id="L169"></a>EV_FLAG1                  = 0x2000;
    <a id="L170"></a>EV_ONESHOT                = 0x10;
    <a id="L171"></a>EV_OOBAND                 = 0x2000;
    <a id="L172"></a>EV_POLL                   = 0x1000;
    <a id="L173"></a>EV_RECEIPT                = 0x40;
    <a id="L174"></a>EV_SYSFLAGS               = 0xf000;
    <a id="L175"></a>EWOULDBLOCK               = 0x23;
    <a id="L176"></a>EXDEV                     = 0x12;
    <a id="L177"></a>FD_CLOEXEC                = 0x1;
    <a id="L178"></a>FD_SETSIZE                = 0x400;
    <a id="L179"></a>F_ADDSIGS                 = 0x3b;
    <a id="L180"></a>F_ALLOCATEALL             = 0x4;
    <a id="L181"></a>F_ALLOCATECONTIG          = 0x2;
    <a id="L182"></a>F_CHKCLEAN                = 0x29;
    <a id="L183"></a>F_DUPFD                   = 0;
    <a id="L184"></a>F_FREEZE_FS               = 0x35;
    <a id="L185"></a>F_FULLFSYNC               = 0x33;
    <a id="L186"></a>F_GETFD                   = 0x1;
    <a id="L187"></a>F_GETFL                   = 0x3;
    <a id="L188"></a>F_GETLK                   = 0x7;
    <a id="L189"></a>F_GETOWN                  = 0x5;
    <a id="L190"></a>F_GETPATH                 = 0x32;
    <a id="L191"></a>F_GLOBAL_NOCACHE          = 0x37;
    <a id="L192"></a>F_LOG2PHYS                = 0x31;
    <a id="L193"></a>F_MARKDEPENDENCY          = 0x3c;
    <a id="L194"></a>F_NOCACHE                 = 0x30;
    <a id="L195"></a>F_PATHPKG_CHECK           = 0x34;
    <a id="L196"></a>F_PEOFPOSMODE             = 0x3;
    <a id="L197"></a>F_PREALLOCATE             = 0x2a;
    <a id="L198"></a>F_RDADVISE                = 0x2c;
    <a id="L199"></a>F_RDAHEAD                 = 0x2d;
    <a id="L200"></a>F_RDLCK                   = 0x1;
    <a id="L201"></a>F_READBOOTSTRAP           = 0x2e;
    <a id="L202"></a>F_SETFD                   = 0x2;
    <a id="L203"></a>F_SETFL                   = 0x4;
    <a id="L204"></a>F_SETLK                   = 0x8;
    <a id="L205"></a>F_SETLKW                  = 0x9;
    <a id="L206"></a>F_SETOWN                  = 0x6;
    <a id="L207"></a>F_SETSIZE                 = 0x2b;
    <a id="L208"></a>F_THAW_FS                 = 0x36;
    <a id="L209"></a>F_UNLCK                   = 0x2;
    <a id="L210"></a>F_VOLPOSMODE              = 0x4;
    <a id="L211"></a>F_WRITEBOOTSTRAP          = 0x2f;
    <a id="L212"></a>F_WRLCK                   = 0x3;
    <a id="L213"></a>IPPROTO_3PC               = 0x22;
    <a id="L214"></a>IPPROTO_ADFS              = 0x44;
    <a id="L215"></a>IPPROTO_AH                = 0x33;
    <a id="L216"></a>IPPROTO_AHIP              = 0x3d;
    <a id="L217"></a>IPPROTO_APES              = 0x63;
    <a id="L218"></a>IPPROTO_ARGUS             = 0xd;
    <a id="L219"></a>IPPROTO_AX25              = 0x5d;
    <a id="L220"></a>IPPROTO_BHA               = 0x31;
    <a id="L221"></a>IPPROTO_BLT               = 0x1e;
    <a id="L222"></a>IPPROTO_BRSATMON          = 0x4c;
    <a id="L223"></a>IPPROTO_CFTP              = 0x3e;
    <a id="L224"></a>IPPROTO_CHAOS             = 0x10;
    <a id="L225"></a>IPPROTO_CMTP              = 0x26;
    <a id="L226"></a>IPPROTO_CPHB              = 0x49;
    <a id="L227"></a>IPPROTO_CPNX              = 0x48;
    <a id="L228"></a>IPPROTO_DDP               = 0x25;
    <a id="L229"></a>IPPROTO_DGP               = 0x56;
    <a id="L230"></a>IPPROTO_DIVERT            = 0xfe;
    <a id="L231"></a>IPPROTO_DONE              = 0x101;
    <a id="L232"></a>IPPROTO_DSTOPTS           = 0x3c;
    <a id="L233"></a>IPPROTO_EGP               = 0x8;
    <a id="L234"></a>IPPROTO_EMCON             = 0xe;
    <a id="L235"></a>IPPROTO_ENCAP             = 0x62;
    <a id="L236"></a>IPPROTO_EON               = 0x50;
    <a id="L237"></a>IPPROTO_ESP               = 0x32;
    <a id="L238"></a>IPPROTO_ETHERIP           = 0x61;
    <a id="L239"></a>IPPROTO_FRAGMENT          = 0x2c;
    <a id="L240"></a>IPPROTO_GGP               = 0x3;
    <a id="L241"></a>IPPROTO_GMTP              = 0x64;
    <a id="L242"></a>IPPROTO_GRE               = 0x2f;
    <a id="L243"></a>IPPROTO_HELLO             = 0x3f;
    <a id="L244"></a>IPPROTO_HMP               = 0x14;
    <a id="L245"></a>IPPROTO_HOPOPTS           = 0;
    <a id="L246"></a>IPPROTO_ICMP              = 0x1;
    <a id="L247"></a>IPPROTO_ICMPV6            = 0x3a;
    <a id="L248"></a>IPPROTO_IDP               = 0x16;
    <a id="L249"></a>IPPROTO_IDPR              = 0x23;
    <a id="L250"></a>IPPROTO_IDRP              = 0x2d;
    <a id="L251"></a>IPPROTO_IGMP              = 0x2;
    <a id="L252"></a>IPPROTO_IGP               = 0x55;
    <a id="L253"></a>IPPROTO_IGRP              = 0x58;
    <a id="L254"></a>IPPROTO_IL                = 0x28;
    <a id="L255"></a>IPPROTO_INLSP             = 0x34;
    <a id="L256"></a>IPPROTO_INP               = 0x20;
    <a id="L257"></a>IPPROTO_IP                = 0;
    <a id="L258"></a>IPPROTO_IPCOMP            = 0x6c;
    <a id="L259"></a>IPPROTO_IPCV              = 0x47;
    <a id="L260"></a>IPPROTO_IPEIP             = 0x5e;
    <a id="L261"></a>IPPROTO_IPIP              = 0x4;
    <a id="L262"></a>IPPROTO_IPPC              = 0x43;
    <a id="L263"></a>IPPROTO_IPV4              = 0x4;
    <a id="L264"></a>IPPROTO_IPV6              = 0x29;
    <a id="L265"></a>IPPROTO_IRTP              = 0x1c;
    <a id="L266"></a>IPPROTO_KRYPTOLAN         = 0x41;
    <a id="L267"></a>IPPROTO_LARP              = 0x5b;
    <a id="L268"></a>IPPROTO_LEAF1             = 0x19;
    <a id="L269"></a>IPPROTO_LEAF2             = 0x1a;
    <a id="L270"></a>IPPROTO_MAX               = 0x100;
    <a id="L271"></a>IPPROTO_MAXID             = 0x34;
    <a id="L272"></a>IPPROTO_MEAS              = 0x13;
    <a id="L273"></a>IPPROTO_MHRP              = 0x30;
    <a id="L274"></a>IPPROTO_MICP              = 0x5f;
    <a id="L275"></a>IPPROTO_MTP               = 0x5c;
    <a id="L276"></a>IPPROTO_MUX               = 0x12;
    <a id="L277"></a>IPPROTO_ND                = 0x4d;
    <a id="L278"></a>IPPROTO_NHRP              = 0x36;
    <a id="L279"></a>IPPROTO_NONE              = 0x3b;
    <a id="L280"></a>IPPROTO_NSP               = 0x1f;
    <a id="L281"></a>IPPROTO_NVPII             = 0xb;
    <a id="L282"></a>IPPROTO_OSPFIGP           = 0x59;
    <a id="L283"></a>IPPROTO_PGM               = 0x71;
    <a id="L284"></a>IPPROTO_PIGP              = 0x9;
    <a id="L285"></a>IPPROTO_PIM               = 0x67;
    <a id="L286"></a>IPPROTO_PRM               = 0x15;
    <a id="L287"></a>IPPROTO_PUP               = 0xc;
    <a id="L288"></a>IPPROTO_PVP               = 0x4b;
    <a id="L289"></a>IPPROTO_RAW               = 0xff;
    <a id="L290"></a>IPPROTO_RCCMON            = 0xa;
    <a id="L291"></a>IPPROTO_RDP               = 0x1b;
    <a id="L292"></a>IPPROTO_ROUTING           = 0x2b;
    <a id="L293"></a>IPPROTO_RSVP              = 0x2e;
    <a id="L294"></a>IPPROTO_RVD               = 0x42;
    <a id="L295"></a>IPPROTO_SATEXPAK          = 0x40;
    <a id="L296"></a>IPPROTO_SATMON            = 0x45;
    <a id="L297"></a>IPPROTO_SCCSP             = 0x60;
    <a id="L298"></a>IPPROTO_SDRP              = 0x2a;
    <a id="L299"></a>IPPROTO_SEP               = 0x21;
    <a id="L300"></a>IPPROTO_SRPC              = 0x5a;
    <a id="L301"></a>IPPROTO_ST                = 0x7;
    <a id="L302"></a>IPPROTO_SVMTP             = 0x52;
    <a id="L303"></a>IPPROTO_SWIPE             = 0x35;
    <a id="L304"></a>IPPROTO_TCF               = 0x57;
    <a id="L305"></a>IPPROTO_TCP               = 0x6;
    <a id="L306"></a>IPPROTO_TP                = 0x1d;
    <a id="L307"></a>IPPROTO_TPXX              = 0x27;
    <a id="L308"></a>IPPROTO_TRUNK1            = 0x17;
    <a id="L309"></a>IPPROTO_TRUNK2            = 0x18;
    <a id="L310"></a>IPPROTO_TTP               = 0x54;
    <a id="L311"></a>IPPROTO_UDP               = 0x11;
    <a id="L312"></a>IPPROTO_VINES             = 0x53;
    <a id="L313"></a>IPPROTO_VISA              = 0x46;
    <a id="L314"></a>IPPROTO_VMTP              = 0x51;
    <a id="L315"></a>IPPROTO_WBEXPAK           = 0x4f;
    <a id="L316"></a>IPPROTO_WBMON             = 0x4e;
    <a id="L317"></a>IPPROTO_WSN               = 0x4a;
    <a id="L318"></a>IPPROTO_XNET              = 0xf;
    <a id="L319"></a>IPPROTO_XTP               = 0x24;
    <a id="L320"></a>IP_ADD_MEMBERSHIP         = 0xc;
    <a id="L321"></a>IP_DEFAULT_MULTICAST_LOOP = 0x1;
    <a id="L322"></a>IP_DEFAULT_MULTICAST_TTL  = 0x1;
    <a id="L323"></a>IP_DROP_MEMBERSHIP        = 0xd;
    <a id="L324"></a>IP_DUMMYNET_CONFIGURE     = 0x3c;
    <a id="L325"></a>IP_DUMMYNET_DEL           = 0x3d;
    <a id="L326"></a>IP_DUMMYNET_FLUSH         = 0x3e;
    <a id="L327"></a>IP_DUMMYNET_GET           = 0x40;
    <a id="L328"></a>IP_FAITH                  = 0x16;
    <a id="L329"></a>IP_FW_ADD                 = 0x28;
    <a id="L330"></a>IP_FW_DEL                 = 0x29;
    <a id="L331"></a>IP_FW_FLUSH               = 0x2a;
    <a id="L332"></a>IP_FW_GET                 = 0x2c;
    <a id="L333"></a>IP_FW_RESETLOG            = 0x2d;
    <a id="L334"></a>IP_FW_ZERO                = 0x2b;
    <a id="L335"></a>IP_HDRINCL                = 0x2;
    <a id="L336"></a>IP_IPSEC_POLICY           = 0x15;
    <a id="L337"></a>IP_MAX_MEMBERSHIPS        = 0x14;
    <a id="L338"></a>IP_MULTICAST_IF           = 0x9;
    <a id="L339"></a>IP_MULTICAST_LOOP         = 0xb;
    <a id="L340"></a>IP_MULTICAST_TTL          = 0xa;
    <a id="L341"></a>IP_MULTICAST_VIF          = 0xe;
    <a id="L342"></a>IP_NAT__XXX               = 0x37;
    <a id="L343"></a>IP_OLD_FW_ADD             = 0x32;
    <a id="L344"></a>IP_OLD_FW_DEL             = 0x33;
    <a id="L345"></a>IP_OLD_FW_FLUSH           = 0x34;
    <a id="L346"></a>IP_OLD_FW_GET             = 0x36;
    <a id="L347"></a>IP_OLD_FW_RESETLOG        = 0x38;
    <a id="L348"></a>IP_OLD_FW_ZERO            = 0x35;
    <a id="L349"></a>IP_OPTIONS                = 0x1;
    <a id="L350"></a>IP_PORTRANGE              = 0x13;
    <a id="L351"></a>IP_PORTRANGE_DEFAULT      = 0;
    <a id="L352"></a>IP_PORTRANGE_HIGH         = 0x1;
    <a id="L353"></a>IP_PORTRANGE_LOW          = 0x2;
    <a id="L354"></a>IP_RECVDSTADDR            = 0x7;
    <a id="L355"></a>IP_RECVIF                 = 0x14;
    <a id="L356"></a>IP_RECVOPTS               = 0x5;
    <a id="L357"></a>IP_RECVRETOPTS            = 0x6;
    <a id="L358"></a>IP_RECVTTL                = 0x18;
    <a id="L359"></a>IP_RETOPTS                = 0x8;
    <a id="L360"></a>IP_RSVP_OFF               = 0x10;
    <a id="L361"></a>IP_RSVP_ON                = 0xf;
    <a id="L362"></a>IP_RSVP_VIF_OFF           = 0x12;
    <a id="L363"></a>IP_RSVP_VIF_ON            = 0x11;
    <a id="L364"></a>IP_STRIPHDR               = 0x17;
    <a id="L365"></a>IP_TOS                    = 0x3;
    <a id="L366"></a>IP_TRAFFIC_MGT_BACKGROUND = 0x41;
    <a id="L367"></a>IP_TTL                    = 0x4;
    <a id="L368"></a>O_ACCMODE                 = 0x3;
    <a id="L369"></a>O_ALERT                   = 0x20000000;
    <a id="L370"></a>O_APPEND                  = 0x8;
    <a id="L371"></a>O_ASYNC                   = 0x40;
    <a id="L372"></a>O_CREAT                   = 0x200;
    <a id="L373"></a>O_DIRECTORY               = 0x100000;
    <a id="L374"></a>O_EVTONLY                 = 0x8000;
    <a id="L375"></a>O_EXCL                    = 0x800;
    <a id="L376"></a>O_EXLOCK                  = 0x20;
    <a id="L377"></a>O_FSYNC                   = 0x80;
    <a id="L378"></a>O_NDELAY                  = 0x4;
    <a id="L379"></a>O_NOCTTY                  = 0x20000;
    <a id="L380"></a>O_NOFOLLOW                = 0x100;
    <a id="L381"></a>O_NONBLOCK                = 0x4;
    <a id="L382"></a>O_POPUP                   = 0x80000000;
    <a id="L383"></a>O_RDONLY                  = 0;
    <a id="L384"></a>O_RDWR                    = 0x2;
    <a id="L385"></a>O_SHLOCK                  = 0x10;
    <a id="L386"></a>O_SYMLINK                 = 0x200000;
    <a id="L387"></a>O_SYNC                    = 0x80;
    <a id="L388"></a>O_TRUNC                   = 0x400;
    <a id="L389"></a>O_WRONLY                  = 0x1;
    <a id="L390"></a>SIGABRT                   = 0x6;
    <a id="L391"></a>SIGALRM                   = 0xe;
    <a id="L392"></a>SIGBUS                    = 0xa;
    <a id="L393"></a>SIGCHLD                   = 0x14;
    <a id="L394"></a>SIGCONT                   = 0x13;
    <a id="L395"></a>SIGEMT                    = 0x7;
    <a id="L396"></a>SIGFPE                    = 0x8;
    <a id="L397"></a>SIGHUP                    = 0x1;
    <a id="L398"></a>SIGILL                    = 0x4;
    <a id="L399"></a>SIGINFO                   = 0x1d;
    <a id="L400"></a>SIGINT                    = 0x2;
    <a id="L401"></a>SIGIO                     = 0x17;
    <a id="L402"></a>SIGIOT                    = 0x6;
    <a id="L403"></a>SIGKILL                   = 0x9;
    <a id="L404"></a>SIGPIPE                   = 0xd;
    <a id="L405"></a>SIGPROF                   = 0x1b;
    <a id="L406"></a>SIGQUIT                   = 0x3;
    <a id="L407"></a>SIGSEGV                   = 0xb;
    <a id="L408"></a>SIGSTOP                   = 0x11;
    <a id="L409"></a>SIGSYS                    = 0xc;
    <a id="L410"></a>SIGTERM                   = 0xf;
    <a id="L411"></a>SIGTRAP                   = 0x5;
    <a id="L412"></a>SIGTSTP                   = 0x12;
    <a id="L413"></a>SIGTTIN                   = 0x15;
    <a id="L414"></a>SIGTTOU                   = 0x16;
    <a id="L415"></a>SIGURG                    = 0x10;
    <a id="L416"></a>SIGUSR1                   = 0x1e;
    <a id="L417"></a>SIGUSR2                   = 0x1f;
    <a id="L418"></a>SIGVTALRM                 = 0x1a;
    <a id="L419"></a>SIGWINCH                  = 0x1c;
    <a id="L420"></a>SIGXCPU                   = 0x18;
    <a id="L421"></a>SIGXFSZ                   = 0x19;
    <a id="L422"></a>SOCK_DGRAM                = 0x2;
    <a id="L423"></a>SOCK_MAXADDRLEN           = 0xff;
    <a id="L424"></a>SOCK_RAW                  = 0x3;
    <a id="L425"></a>SOCK_RDM                  = 0x4;
    <a id="L426"></a>SOCK_SEQPACKET            = 0x5;
    <a id="L427"></a>SOCK_STREAM               = 0x1;
    <a id="L428"></a>SOL_SOCKET                = 0xffff;
    <a id="L429"></a>SOMAXCONN                 = 0x80;
    <a id="L430"></a>SO_ACCEPTCONN             = 0x2;
    <a id="L431"></a>SO_BROADCAST              = 0x20;
    <a id="L432"></a>SO_DEBUG                  = 0x1;
    <a id="L433"></a>SO_DONTROUTE              = 0x10;
    <a id="L434"></a>SO_DONTTRUNC              = 0x2000;
    <a id="L435"></a>SO_ERROR                  = 0x1007;
    <a id="L436"></a>SO_KEEPALIVE              = 0x8;
    <a id="L437"></a>SO_LABEL                  = 0x1010;
    <a id="L438"></a>SO_LINGER                 = 0x80;
    <a id="L439"></a>SO_LINGER_SEC             = 0x1080;
    <a id="L440"></a>SO_NKE                    = 0x1021;
    <a id="L441"></a>SO_NOADDRERR              = 0x1023;
    <a id="L442"></a>SO_NOSIGPIPE              = 0x1022;
    <a id="L443"></a>SO_NOTIFYCONFLICT         = 0x1026;
    <a id="L444"></a>SO_NREAD                  = 0x1020;
    <a id="L445"></a>SO_NWRITE                 = 0x1024;
    <a id="L446"></a>SO_OOBINLINE              = 0x100;
    <a id="L447"></a>SO_PEERLABEL              = 0x1011;
    <a id="L448"></a>SO_RCVBUF                 = 0x1002;
    <a id="L449"></a>SO_RCVLOWAT               = 0x1004;
    <a id="L450"></a>SO_RCVTIMEO               = 0x1006;
    <a id="L451"></a>SO_RESTRICTIONS           = 0x1081;
    <a id="L452"></a>SO_RESTRICT_DENYIN        = 0x1;
    <a id="L453"></a>SO_RESTRICT_DENYOUT       = 0x2;
    <a id="L454"></a>SO_RESTRICT_DENYSET       = 0x80000000;
    <a id="L455"></a>SO_REUSEADDR              = 0x4;
    <a id="L456"></a>SO_REUSEPORT              = 0x200;
    <a id="L457"></a>SO_REUSESHAREUID          = 0x1025;
    <a id="L458"></a>SO_SNDBUF                 = 0x1001;
    <a id="L459"></a>SO_SNDLOWAT               = 0x1003;
    <a id="L460"></a>SO_SNDTIMEO               = 0x1005;
    <a id="L461"></a>SO_TIMESTAMP              = 0x400;
    <a id="L462"></a>SO_TYPE                   = 0x1008;
    <a id="L463"></a>SO_USELOOPBACK            = 0x40;
    <a id="L464"></a>SO_WANTMORE               = 0x4000;
    <a id="L465"></a>SO_WANTOOBFLAG            = 0x8000;
    <a id="L466"></a>S_IEXEC                   = 0x40;
    <a id="L467"></a>S_IFBLK                   = 0x6000;
    <a id="L468"></a>S_IFCHR                   = 0x2000;
    <a id="L469"></a>S_IFDIR                   = 0x4000;
    <a id="L470"></a>S_IFIFO                   = 0x1000;
    <a id="L471"></a>S_IFLNK                   = 0xa000;
    <a id="L472"></a>S_IFMT                    = 0xf000;
    <a id="L473"></a>S_IFREG                   = 0x8000;
    <a id="L474"></a>S_IFSOCK                  = 0xc000;
    <a id="L475"></a>S_IFWHT                   = 0xe000;
    <a id="L476"></a>S_IFXATTR                 = 0x10000;
    <a id="L477"></a>S_IREAD                   = 0x100;
    <a id="L478"></a>S_IRGRP                   = 0x20;
    <a id="L479"></a>S_IROTH                   = 0x4;
    <a id="L480"></a>S_IRUSR                   = 0x100;
    <a id="L481"></a>S_IRWXG                   = 0x38;
    <a id="L482"></a>S_IRWXO                   = 0x7;
    <a id="L483"></a>S_IRWXU                   = 0x1c0;
    <a id="L484"></a>S_ISGID                   = 0x400;
    <a id="L485"></a>S_ISTXT                   = 0x200;
    <a id="L486"></a>S_ISUID                   = 0x800;
    <a id="L487"></a>S_ISVTX                   = 0x200;
    <a id="L488"></a>S_IWGRP                   = 0x10;
    <a id="L489"></a>S_IWOTH                   = 0x2;
    <a id="L490"></a>S_IWRITE                  = 0x80;
    <a id="L491"></a>S_IWUSR                   = 0x80;
    <a id="L492"></a>S_IXGRP                   = 0x8;
    <a id="L493"></a>S_IXOTH                   = 0x1;
    <a id="L494"></a>S_IXUSR                   = 0x40;
    <a id="L495"></a>TCP_KEEPALIVE             = 0x10;
    <a id="L496"></a>TCP_MAXBURST              = 0x4;
    <a id="L497"></a>TCP_MAXHLEN               = 0x3c;
    <a id="L498"></a>TCP_MAXOLEN               = 0x28;
    <a id="L499"></a>TCP_MAXSEG                = 0x2;
    <a id="L500"></a>TCP_MAXWIN                = 0xffff;
    <a id="L501"></a>TCP_MAX_SACK              = 0x3;
    <a id="L502"></a>TCP_MAX_WINSHIFT          = 0xe;
    <a id="L503"></a>TCP_MINMSS                = 0xd8;
    <a id="L504"></a>TCP_MINMSSOVERLOAD        = 0x3e8;
    <a id="L505"></a>TCP_MSS                   = 0x200;
    <a id="L506"></a>TCP_NODELAY               = 0x1;
    <a id="L507"></a>TCP_NOOPT                 = 0x8;
    <a id="L508"></a>TCP_NOPUSH                = 0x4;
    <a id="L509"></a>WCONTINUED                = 0x10;
    <a id="L510"></a>WCOREFLAG                 = 0x80;
    <a id="L511"></a>WEXITED                   = 0x4;
    <a id="L512"></a>WNOHANG                   = 0x1;
    <a id="L513"></a>WNOWAIT                   = 0x20;
    <a id="L514"></a>WSTOPPED                  = 0x7f;
    <a id="L515"></a>WUNTRACED                 = 0x2;
<a id="L516"></a>)

<a id="L518"></a><span class="comment">// Types</span>


<a id="L521"></a><span class="comment">// Error table</span>
<a id="L522"></a>var errors = [...]string{
    <a id="L523"></a>95: &#34;EMULTIHOP (Reserved)&#34;,
    <a id="L524"></a>47: &#34;address family not supported by protocol family&#34;,
    <a id="L525"></a>13: &#34;permission denied&#34;,
    <a id="L526"></a>39: &#34;destination address required&#34;,
    <a id="L527"></a>92: &#34;illegal byte sequence&#34;,
    <a id="L528"></a>29: &#34;illegal seek&#34;,
    <a id="L529"></a>31: &#34;too many links&#34;,
    <a id="L530"></a>74: &#34;RPC prog. not avail&#34;,
    <a id="L531"></a>25: &#34;inappropriate ioctl for device&#34;,
    <a id="L532"></a>9: &#34;bad file descriptor&#34;,
    <a id="L533"></a>34: &#34;result too large&#34;,
    <a id="L534"></a>89: &#34;operation canceled&#34;,
    <a id="L535"></a>26: &#34;text file busy&#34;,
    <a id="L536"></a>12: &#34;cannot allocate memory&#34;,
    <a id="L537"></a>36: &#34;operation now in progress&#34;,
    <a id="L538"></a>66: &#34;directory not empty&#34;,
    <a id="L539"></a>15: &#34;block device required&#34;,
    <a id="L540"></a>41: &#34;protocol wrong type for socket&#34;,
    <a id="L541"></a>91: &#34;no message of desired type&#34;,
    <a id="L542"></a>73: &#34;RPC version wrong&#34;,
    <a id="L543"></a>20: &#34;not a directory&#34;,
    <a id="L544"></a>37: &#34;operation already in progress&#34;,
    <a id="L545"></a>60: &#34;operation timed out&#34;,
    <a id="L546"></a>81: &#34;need authenticator&#34;,
    <a id="L547"></a>96: &#34;no message available on STREAM&#34;,
    <a id="L548"></a>4: &#34;interrupted system call&#34;,
    <a id="L549"></a>97: &#34;ENOLINK (Reserved)&#34;,
    <a id="L550"></a>1: &#34;operation not permitted&#34;,
    <a id="L551"></a>50: &#34;network is down&#34;,
    <a id="L552"></a>70: &#34;stale NFS file handle&#34;,
    <a id="L553"></a>38: &#34;socket operation on non-socket&#34;,
    <a id="L554"></a>98: &#34;no STREAM resources&#34;,
    <a id="L555"></a>80: &#34;authentication error&#34;,
    <a id="L556"></a>10: &#34;no child processes&#34;,
    <a id="L557"></a>32: &#34;broken pipe&#34;,
    <a id="L558"></a>93: &#34;attribute not found&#34;,
    <a id="L559"></a>94: &#34;bad message&#34;,
    <a id="L560"></a>71: &#34;too many levels of remote in path&#34;,
    <a id="L561"></a>59: &#34;too many references: can&#39;t splice&#34;,
    <a id="L562"></a>46: &#34;protocol family not supported&#34;,
    <a id="L563"></a>76: &#34;bad procedure for program&#34;,
    <a id="L564"></a>48: &#34;address already in use&#34;,
    <a id="L565"></a>52: &#34;network dropped connection on reset&#34;,
    <a id="L566"></a>21: &#34;is a directory&#34;,
    <a id="L567"></a>90: &#34;identifier removed&#34;,
    <a id="L568"></a>83: &#34;device error&#34;,
    <a id="L569"></a>22: &#34;invalid argument&#34;,
    <a id="L570"></a>58: &#34;can&#39;t send after socket shutdown&#34;,
    <a id="L571"></a>82: &#34;device power is off&#34;,
    <a id="L572"></a>84: &#34;value too large to be stored in data type&#34;,
    <a id="L573"></a>16: &#34;resource busy&#34;,
    <a id="L574"></a>67: &#34;too many processes&#34;,
    <a id="L575"></a>100: &#34;protocol error&#34;,
    <a id="L576"></a>19: &#34;operation not supported by device&#34;,
    <a id="L577"></a>30: &#34;read-only file system&#34;,
    <a id="L578"></a>7: &#34;argument list too long&#34;,
    <a id="L579"></a>11: &#34;resource deadlock avoided&#34;,
    <a id="L580"></a>54: &#34;connection reset by peer&#34;,
    <a id="L581"></a>88: &#34;malformed Mach-o file&#34;,
    <a id="L582"></a>6: &#34;device not configured&#34;,
    <a id="L583"></a>72: &#34;RPC struct is bad&#34;,
    <a id="L584"></a>63: &#34;file name too long&#34;,
    <a id="L585"></a>103: &#34;policy not found&#34;,
    <a id="L586"></a>44: &#34;socket type not supported&#34;,
    <a id="L587"></a>49: &#34;can&#39;t assign requested address&#34;,
    <a id="L588"></a>101: &#34;STREAM ioctl timeout&#34;,
    <a id="L589"></a>43: &#34;protocol not supported&#34;,
    <a id="L590"></a>5: &#34;input/output error&#34;,
    <a id="L591"></a>51: &#34;network is unreachable&#34;,
    <a id="L592"></a>18: &#34;cross-device link&#34;,
    <a id="L593"></a>69: &#34;disc quota exceeded&#34;,
    <a id="L594"></a>28: &#34;no space left on device&#34;,
    <a id="L595"></a>8: &#34;exec format error&#34;,
    <a id="L596"></a>40: &#34;message too long&#34;,
    <a id="L597"></a>79: &#34;inappropriate file type or format&#34;,
    <a id="L598"></a>33: &#34;numerical argument out of domain&#34;,
    <a id="L599"></a>99: &#34;not a STREAM&#34;,
    <a id="L600"></a>27: &#34;file too large&#34;,
    <a id="L601"></a>3: &#34;no such process&#34;,
    <a id="L602"></a>64: &#34;host is down&#34;,
    <a id="L603"></a>77: &#34;no locks available&#34;,
    <a id="L604"></a>23: &#34;too many open files in system&#34;,
    <a id="L605"></a>78: &#34;function not implemented&#34;,
    <a id="L606"></a>86: &#34;bad CPU type in executable&#34;,
    <a id="L607"></a>57: &#34;socket is not connected&#34;,
    <a id="L608"></a>45: &#34;operation not supported&#34;,
    <a id="L609"></a>53: &#34;software caused connection abort&#34;,
    <a id="L610"></a>56: &#34;socket is already connected&#34;,
    <a id="L611"></a>87: &#34;shared library version mismatch&#34;,
    <a id="L612"></a>68: &#34;too many users&#34;,
    <a id="L613"></a>42: &#34;protocol not available&#34;,
    <a id="L614"></a>24: &#34;too many open files&#34;,
    <a id="L615"></a>62: &#34;too many levels of symbolic links&#34;,
    <a id="L616"></a>55: &#34;no buffer space available&#34;,
    <a id="L617"></a>14: &#34;bad address&#34;,
    <a id="L618"></a>35: &#34;resource temporarily unavailable&#34;,
    <a id="L619"></a>85: &#34;bad executable (or shared library)&#34;,
    <a id="L620"></a>61: &#34;connection refused&#34;,
    <a id="L621"></a>17: &#34;file exists&#34;,
    <a id="L622"></a>75: &#34;program version wrong&#34;,
    <a id="L623"></a>2: &#34;no such file or directory&#34;,
    <a id="L624"></a>65: &#34;no route to host&#34;,
    <a id="L625"></a>102: &#34;operation not supported on socket&#34;,
<a id="L626"></a>}
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
