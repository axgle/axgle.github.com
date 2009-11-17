<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/zsysnum_nacl_386.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/syscall/zsysnum_nacl_386.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// mksysnum_nacl.sh /home/rsc/pub/nacl/native_client/src/trusted/service_runtime/include/bits/nacl_syscalls.h</span>
<a id="L2"></a><span class="comment">// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT</span>

<a id="L4"></a>package syscall

<a id="L6"></a>const (
    <a id="L7"></a>SYS_NULL                = 1;
    <a id="L8"></a>SYS_OPEN                = 10;
    <a id="L9"></a>SYS_CLOSE               = 11;
    <a id="L10"></a>SYS_READ                = 12;
    <a id="L11"></a>SYS_WRITE               = 13;
    <a id="L12"></a>SYS_LSEEK               = 14;
    <a id="L13"></a>SYS_IOCTL               = 15;
    <a id="L14"></a>SYS_STAT                = 16;
    <a id="L15"></a>SYS_FSTAT               = 17;
    <a id="L16"></a>SYS_CHMOD               = 18;
    <a id="L17"></a>SYS_SYSBRK              = 20;
    <a id="L18"></a>SYS_MMAP                = 21;
    <a id="L19"></a>SYS_MUNMAP              = 22;
    <a id="L20"></a>SYS_GETDENTS            = 23;
    <a id="L21"></a>SYS_EXIT                = 30;
    <a id="L22"></a>SYS_GETPID              = 31;
    <a id="L23"></a>SYS_SCHED_YIELD         = 32;
    <a id="L24"></a>SYS_SYSCONF             = 33;
    <a id="L25"></a>SYS_GETTIMEOFDAY        = 40;
    <a id="L26"></a>SYS_CLOCK               = 41;
    <a id="L27"></a>SYS_MULTIMEDIA_INIT     = 50;
    <a id="L28"></a>SYS_MULTIMEDIA_SHUTDOWN = 51;
    <a id="L29"></a>SYS_VIDEO_INIT          = 52;
    <a id="L30"></a>SYS_VIDEO_SHUTDOWN      = 53;
    <a id="L31"></a>SYS_VIDEO_UPDATE        = 54;
    <a id="L32"></a>SYS_VIDEO_POLL_EVENT    = 55;
    <a id="L33"></a>SYS_AUDIO_INIT          = 56;
    <a id="L34"></a>SYS_AUDIO_SHUTDOWN      = 57;
    <a id="L35"></a>SYS_AUDIO_STREAM        = 58;
    <a id="L36"></a>SYS_IMC_MAKEBOUNDSOCK   = 60;
    <a id="L37"></a>SYS_IMC_ACCEPT          = 61;
    <a id="L38"></a>SYS_IMC_CONNECT         = 62;
    <a id="L39"></a>SYS_IMC_SENDMSG         = 63;
    <a id="L40"></a>SYS_IMC_RECVMSG         = 64;
    <a id="L41"></a>SYS_IMC_MEM_OBJ_CREATE  = 65;
    <a id="L42"></a>SYS_IMC_SOCKETPAIR      = 66;
    <a id="L43"></a>SYS_MUTEX_CREATE        = 70;
    <a id="L44"></a>SYS_MUTEX_LOCK          = 71;
    <a id="L45"></a>SYS_MUTEX_TRYLOCK       = 72;
    <a id="L46"></a>SYS_MUTEX_UNLOCK        = 73;
    <a id="L47"></a>SYS_COND_CREATE         = 74;
    <a id="L48"></a>SYS_COND_WAIT           = 75;
    <a id="L49"></a>SYS_COND_SIGNAL         = 76;
    <a id="L50"></a>SYS_COND_BROADCAST      = 77;
    <a id="L51"></a>SYS_COND_TIMED_WAIT_ABS = 79;
    <a id="L52"></a>SYS_THREAD_CREATE       = 80;
    <a id="L53"></a>SYS_THREAD_EXIT         = 81;
    <a id="L54"></a>SYS_TLS_INIT            = 82;
    <a id="L55"></a>SYS_THREAD_NICE         = 83;
    <a id="L56"></a>SYS_SRPC_GET_FD         = 90;
    <a id="L57"></a>SYS_SEM_CREATE          = 100;
    <a id="L58"></a>SYS_SEM_WAIT            = 101;
    <a id="L59"></a>SYS_SEM_POST            = 102;
    <a id="L60"></a>SYS_SEM_GET_VALUE       = 103;
<a id="L61"></a>)
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
