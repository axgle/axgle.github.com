<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/zsyscall_nacl_386.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/syscall/zsyscall_nacl_386.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// mksyscall.sh -l32 syscall_nacl.go syscall_nacl_386.go</span>
<a id="L2"></a><span class="comment">// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT</span>

<a id="L4"></a>package syscall

<a id="L6"></a>import &#34;unsafe&#34;

<a id="L8"></a>func Chmod(path string, mode int) (errno int) {
    <a id="L9"></a>_, _, e1 := Syscall(SYS_CHMOD, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
    <a id="L10"></a>errno = int(e1);
    <a id="L11"></a>return;
<a id="L12"></a>}

<a id="L14"></a>func Clock() (clock int) {
    <a id="L15"></a>r0, _, _ := Syscall(SYS_CLOCK, 0, 0, 0);
    <a id="L16"></a>clock = int(r0);
    <a id="L17"></a>return;
<a id="L18"></a>}

<a id="L20"></a>func Close(fd int) (errno int) {
    <a id="L21"></a>_, _, e1 := Syscall(SYS_CLOSE, uintptr(fd), 0, 0);
    <a id="L22"></a>errno = int(e1);
    <a id="L23"></a>return;
<a id="L24"></a>}

<a id="L26"></a>func Exit(code int) {
    <a id="L27"></a>Syscall(SYS_EXIT, uintptr(code), 0, 0);
    <a id="L28"></a>return;
<a id="L29"></a>}

<a id="L31"></a>func Fstat(fd int, stat *Stat_t) (errno int) {
    <a id="L32"></a>_, _, e1 := Syscall(SYS_FSTAT, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0);
    <a id="L33"></a>errno = int(e1);
    <a id="L34"></a>return;
<a id="L35"></a>}

<a id="L37"></a>func Getdents(fd int, buf []byte) (n int, errno int) {
    <a id="L38"></a>var _p0 *byte;
    <a id="L39"></a>if len(buf) &gt; 0 {
        <a id="L40"></a>_p0 = &amp;buf[0]
    <a id="L41"></a>}
    <a id="L42"></a>r0, _, e1 := Syscall(SYS_GETDENTS, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)));
    <a id="L43"></a>n = int(r0);
    <a id="L44"></a>errno = int(e1);
    <a id="L45"></a>return;
<a id="L46"></a>}

<a id="L48"></a>func Getpid() (pid int) {
    <a id="L49"></a>r0, _, _ := Syscall(SYS_GETPID, 0, 0, 0);
    <a id="L50"></a>pid = int(r0);
    <a id="L51"></a>return;
<a id="L52"></a>}

<a id="L54"></a>func Gettimeofday(tv *Timeval) (errno int) {
    <a id="L55"></a>_, _, e1 := Syscall(SYS_GETTIMEOFDAY, uintptr(unsafe.Pointer(tv)), 0, 0);
    <a id="L56"></a>errno = int(e1);
    <a id="L57"></a>return;
<a id="L58"></a>}

<a id="L60"></a>func Open(path string, mode int, perm int) (fd int, errno int) {
    <a id="L61"></a>r0, _, e1 := Syscall(SYS_OPEN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(perm));
    <a id="L62"></a>fd = int(r0);
    <a id="L63"></a>errno = int(e1);
    <a id="L64"></a>return;
<a id="L65"></a>}

<a id="L67"></a>func Read(fd int, p []byte) (n int, errno int) {
    <a id="L68"></a>var _p0 *byte;
    <a id="L69"></a>if len(p) &gt; 0 {
        <a id="L70"></a>_p0 = &amp;p[0]
    <a id="L71"></a>}
    <a id="L72"></a>r0, _, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)));
    <a id="L73"></a>n = int(r0);
    <a id="L74"></a>errno = int(e1);
    <a id="L75"></a>return;
<a id="L76"></a>}

<a id="L78"></a>func read(fd int, buf *byte, nbuf int) (n int, errno int) {
    <a id="L79"></a>r0, _, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf));
    <a id="L80"></a>n = int(r0);
    <a id="L81"></a>errno = int(e1);
    <a id="L82"></a>return;
<a id="L83"></a>}

<a id="L85"></a>func Stat(path string, stat *Stat_t) (errno int) {
    <a id="L86"></a>_, _, e1 := Syscall(SYS_STAT, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(stat)), 0);
    <a id="L87"></a>errno = int(e1);
    <a id="L88"></a>return;
<a id="L89"></a>}

<a id="L91"></a>func Write(fd int, p []byte) (n int, errno int) {
    <a id="L92"></a>var _p0 *byte;
    <a id="L93"></a>if len(p) &gt; 0 {
        <a id="L94"></a>_p0 = &amp;p[0]
    <a id="L95"></a>}
    <a id="L96"></a>r0, _, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)));
    <a id="L97"></a>n = int(r0);
    <a id="L98"></a>errno = int(e1);
    <a id="L99"></a>return;
<a id="L100"></a>}

<a id="L102"></a>func MultimediaInit(subsys int) (errno int) {
    <a id="L103"></a>_, _, e1 := Syscall(SYS_MULTIMEDIA_INIT, uintptr(subsys), 0, 0);
    <a id="L104"></a>errno = int(e1);
    <a id="L105"></a>return;
<a id="L106"></a>}

<a id="L108"></a>func MultimediaShutdown() (errno int) {
    <a id="L109"></a>_, _, e1 := Syscall(SYS_MULTIMEDIA_SHUTDOWN, 0, 0, 0);
    <a id="L110"></a>errno = int(e1);
    <a id="L111"></a>return;
<a id="L112"></a>}

<a id="L114"></a>func CondCreate() (cv int, errno int) {
    <a id="L115"></a>r0, _, e1 := Syscall(SYS_COND_CREATE, 0, 0, 0);
    <a id="L116"></a>cv = int(r0);
    <a id="L117"></a>errno = int(e1);
    <a id="L118"></a>return;
<a id="L119"></a>}

<a id="L121"></a>func CondWait(cv int, mutex int) (errno int) {
    <a id="L122"></a>_, _, e1 := Syscall(SYS_COND_WAIT, uintptr(cv), uintptr(mutex), 0);
    <a id="L123"></a>errno = int(e1);
    <a id="L124"></a>return;
<a id="L125"></a>}

<a id="L127"></a>func CondSignal(cv int) (errno int) {
    <a id="L128"></a>_, _, e1 := Syscall(SYS_COND_SIGNAL, uintptr(cv), 0, 0);
    <a id="L129"></a>errno = int(e1);
    <a id="L130"></a>return;
<a id="L131"></a>}

<a id="L133"></a>func CondBroadcast(cv int) (errno int) {
    <a id="L134"></a>_, _, e1 := Syscall(SYS_COND_BROADCAST, uintptr(cv), 0, 0);
    <a id="L135"></a>errno = int(e1);
    <a id="L136"></a>return;
<a id="L137"></a>}

<a id="L139"></a>func CondTimedWaitAbs(cv int, mutex int, abstime *Timespec) (errno int) {
    <a id="L140"></a>_, _, e1 := Syscall(SYS_COND_TIMED_WAIT_ABS, uintptr(cv), uintptr(mutex), uintptr(unsafe.Pointer(abstime)));
    <a id="L141"></a>errno = int(e1);
    <a id="L142"></a>return;
<a id="L143"></a>}

<a id="L145"></a>func MutexCreate() (mutex int, errno int) {
    <a id="L146"></a>r0, _, e1 := Syscall(SYS_MUTEX_CREATE, 0, 0, 0);
    <a id="L147"></a>mutex = int(r0);
    <a id="L148"></a>errno = int(e1);
    <a id="L149"></a>return;
<a id="L150"></a>}

<a id="L152"></a>func MutexLock(mutex int) (errno int) {
    <a id="L153"></a>_, _, e1 := Syscall(SYS_MUTEX_LOCK, uintptr(mutex), 0, 0);
    <a id="L154"></a>errno = int(e1);
    <a id="L155"></a>return;
<a id="L156"></a>}

<a id="L158"></a>func MutexUnlock(mutex int) (errno int) {
    <a id="L159"></a>_, _, e1 := Syscall(SYS_MUTEX_UNLOCK, uintptr(mutex), 0, 0);
    <a id="L160"></a>errno = int(e1);
    <a id="L161"></a>return;
<a id="L162"></a>}

<a id="L164"></a>func MutexTryLock(mutex int) (errno int) {
    <a id="L165"></a>_, _, e1 := Syscall(SYS_MUTEX_TRYLOCK, uintptr(mutex), 0, 0);
    <a id="L166"></a>errno = int(e1);
    <a id="L167"></a>return;
<a id="L168"></a>}

<a id="L170"></a>func SemCreate() (sema int, errno int) {
    <a id="L171"></a>r0, _, e1 := Syscall(SYS_SEM_CREATE, 0, 0, 0);
    <a id="L172"></a>sema = int(r0);
    <a id="L173"></a>errno = int(e1);
    <a id="L174"></a>return;
<a id="L175"></a>}

<a id="L177"></a>func SemWait(sema int) (errno int) {
    <a id="L178"></a>_, _, e1 := Syscall(SYS_SEM_WAIT, uintptr(sema), 0, 0);
    <a id="L179"></a>errno = int(e1);
    <a id="L180"></a>return;
<a id="L181"></a>}

<a id="L183"></a>func SemPost(sema int) (errno int) {
    <a id="L184"></a>_, _, e1 := Syscall(SYS_SEM_POST, uintptr(sema), 0, 0);
    <a id="L185"></a>errno = int(e1);
    <a id="L186"></a>return;
<a id="L187"></a>}

<a id="L189"></a>func VideoInit(dx int, dy int) (errno int) {
    <a id="L190"></a>_, _, e1 := Syscall(SYS_VIDEO_INIT, uintptr(dx), uintptr(dy), 0);
    <a id="L191"></a>errno = int(e1);
    <a id="L192"></a>return;
<a id="L193"></a>}

<a id="L195"></a>func VideoUpdate(data *uint32) (errno int) {
    <a id="L196"></a>_, _, e1 := Syscall(SYS_VIDEO_UPDATE, uintptr(unsafe.Pointer(data)), 0, 0);
    <a id="L197"></a>errno = int(e1);
    <a id="L198"></a>return;
<a id="L199"></a>}

<a id="L201"></a>func VideoPollEvent(ev *byte) (errno int) {
    <a id="L202"></a>_, _, e1 := Syscall(SYS_VIDEO_POLL_EVENT, uintptr(unsafe.Pointer(ev)), 0, 0);
    <a id="L203"></a>errno = int(e1);
    <a id="L204"></a>return;
<a id="L205"></a>}

<a id="L207"></a>func VideoShutdown() (errno int) {
    <a id="L208"></a>_, _, e1 := Syscall(SYS_VIDEO_SHUTDOWN, 0, 0, 0);
    <a id="L209"></a>errno = int(e1);
    <a id="L210"></a>return;
<a id="L211"></a>}

<a id="L213"></a>func AudioInit(fmt int, nreq int, data *int) (errno int) {
    <a id="L214"></a>_, _, e1 := Syscall(SYS_AUDIO_INIT, uintptr(fmt), uintptr(nreq), uintptr(unsafe.Pointer(data)));
    <a id="L215"></a>errno = int(e1);
    <a id="L216"></a>return;
<a id="L217"></a>}

<a id="L219"></a>func AudioShutdown() (errno int) {
    <a id="L220"></a>_, _, e1 := Syscall(SYS_AUDIO_SHUTDOWN, 0, 0, 0);
    <a id="L221"></a>errno = int(e1);
    <a id="L222"></a>return;
<a id="L223"></a>}

<a id="L225"></a>func AudioStream(data *uint16, size *uintptr) (errno int) {
    <a id="L226"></a>_, _, e1 := Syscall(SYS_AUDIO_STREAM, uintptr(unsafe.Pointer(data)), uintptr(unsafe.Pointer(size)), 0);
    <a id="L227"></a>errno = int(e1);
    <a id="L228"></a>return;
<a id="L229"></a>}
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
