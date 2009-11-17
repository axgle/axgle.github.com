<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/os/exec.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/os/exec.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package os

<a id="L7"></a>import (
    <a id="L8"></a>&#34;syscall&#34;;
<a id="L9"></a>)

<a id="L11"></a><span class="comment">// ForkExec forks the current process and invokes Exec with the file, arguments,</span>
<a id="L12"></a><span class="comment">// and environment specified by argv0, argv, and envv.  It returns the process</span>
<a id="L13"></a><span class="comment">// id of the forked process and an Error, if any.  The fd array specifies the</span>
<a id="L14"></a><span class="comment">// file descriptors to be set up in the new process: fd[0] will be Unix file</span>
<a id="L15"></a><span class="comment">// descriptor 0 (standard input), fd[1] descriptor 1, and so on.  A nil entry</span>
<a id="L16"></a><span class="comment">// will cause the child to have no open file descriptor with that index.</span>
<a id="L17"></a><span class="comment">// If dir is not empty, the child chdirs into the directory before execing the program.</span>
<a id="L18"></a>func ForkExec(argv0 string, argv []string, envv []string, dir string, fd []*File) (pid int, err Error) {
    <a id="L19"></a><span class="comment">// Create array of integer (system) fds.</span>
    <a id="L20"></a>intfd := make([]int, len(fd));
    <a id="L21"></a>for i, f := range fd {
        <a id="L22"></a>if f == nil {
            <a id="L23"></a>intfd[i] = -1
        <a id="L24"></a>} else {
            <a id="L25"></a>intfd[i] = f.Fd()
        <a id="L26"></a>}
    <a id="L27"></a>}

    <a id="L29"></a>p, e := syscall.ForkExec(argv0, argv, envv, dir, intfd);
    <a id="L30"></a>if e != 0 {
        <a id="L31"></a>return 0, &amp;PathError{&#34;fork/exec&#34;, argv0, Errno(e)}
    <a id="L32"></a>}
    <a id="L33"></a>return p, nil;
<a id="L34"></a>}

<a id="L36"></a><span class="comment">// Exec replaces the current process with an execution of the program</span>
<a id="L37"></a><span class="comment">// named by argv0, with arguments argv and environment envv.</span>
<a id="L38"></a><span class="comment">// If successful, Exec never returns.  If it fails, it returns an Error.</span>
<a id="L39"></a><span class="comment">// ForkExec is almost always a better way to execute a program.</span>
<a id="L40"></a>func Exec(argv0 string, argv []string, envv []string) Error {
    <a id="L41"></a>if envv == nil {
        <a id="L42"></a>envv = Environ()
    <a id="L43"></a>}
    <a id="L44"></a>e := syscall.Exec(argv0, argv, envv);
    <a id="L45"></a>if e != 0 {
        <a id="L46"></a>return &amp;PathError{&#34;exec&#34;, argv0, Errno(e)}
    <a id="L47"></a>}
    <a id="L48"></a>return nil;
<a id="L49"></a>}

<a id="L51"></a><span class="comment">// TODO(rsc): Should os implement its own syscall.WaitStatus</span>
<a id="L52"></a><span class="comment">// wrapper with the methods, or is exposing the underlying one enough?</span>
<a id="L53"></a><span class="comment">//</span>
<a id="L54"></a><span class="comment">// TODO(rsc): Certainly need to have Rusage struct,</span>
<a id="L55"></a><span class="comment">// since syscall one might have different field types across</span>
<a id="L56"></a><span class="comment">// different OS.</span>

<a id="L58"></a><span class="comment">// Waitmsg stores the information about an exited process as reported by Wait.</span>
<a id="L59"></a>type Waitmsg struct {
    <a id="L60"></a>Pid                 int;             <span class="comment">// The process&#39;s id.</span>
    <a id="L61"></a>syscall.WaitStatus;                  <span class="comment">// System-dependent status info.</span>
    <a id="L62"></a>Rusage              *syscall.Rusage; <span class="comment">// System-dependent resource usage info.</span>
<a id="L63"></a>}

<a id="L65"></a><span class="comment">// Options for Wait.</span>
<a id="L66"></a>const (
    <a id="L67"></a>WNOHANG   = syscall.WNOHANG;  <span class="comment">// Don&#39;t wait if no process has exited.</span>
    <a id="L68"></a>WSTOPPED  = syscall.WSTOPPED; <span class="comment">// If set, status of stopped subprocesses is also reported.</span>
    <a id="L69"></a>WUNTRACED = WSTOPPED;
    <a id="L70"></a>WRUSAGE   = 1 &lt;&lt; 20; <span class="comment">// Record resource usage.</span>
<a id="L71"></a>)

<a id="L73"></a><span class="comment">// WRUSAGE must not be too high a bit, to avoid clashing with Linux&#39;s</span>
<a id="L74"></a><span class="comment">// WCLONE, WALL, and WNOTHREAD flags, which sit in the top few bits of</span>
<a id="L75"></a><span class="comment">// the options</span>

<a id="L77"></a><span class="comment">// Wait waits for process pid to exit or stop, and then returns a</span>
<a id="L78"></a><span class="comment">// Waitmsg describing its status and an Error, if any. The options</span>
<a id="L79"></a><span class="comment">// (WNOHANG etc.) affect the behavior of the Wait call.</span>
<a id="L80"></a>func Wait(pid int, options int) (w *Waitmsg, err Error) {
    <a id="L81"></a>var status syscall.WaitStatus;
    <a id="L82"></a>var rusage *syscall.Rusage;
    <a id="L83"></a>if options&amp;WRUSAGE != 0 {
        <a id="L84"></a>rusage = new(syscall.Rusage);
        <a id="L85"></a>options ^= WRUSAGE;
    <a id="L86"></a>}
    <a id="L87"></a>pid1, e := syscall.Wait4(pid, &amp;status, options, rusage);
    <a id="L88"></a>if e != 0 {
        <a id="L89"></a>return nil, NewSyscallError(&#34;wait&#34;, e)
    <a id="L90"></a>}
    <a id="L91"></a>w = new(Waitmsg);
    <a id="L92"></a>w.Pid = pid1;
    <a id="L93"></a>w.WaitStatus = status;
    <a id="L94"></a>w.Rusage = rusage;
    <a id="L95"></a>return w, nil;
<a id="L96"></a>}

<a id="L98"></a><span class="comment">// Convert i to decimal string.</span>
<a id="L99"></a>func itod(i int) string {
    <a id="L100"></a>if i == 0 {
        <a id="L101"></a>return &#34;0&#34;
    <a id="L102"></a>}

    <a id="L104"></a>u := uint64(i);
    <a id="L105"></a>if i &lt; 0 {
        <a id="L106"></a>u = -u
    <a id="L107"></a>}

    <a id="L109"></a><span class="comment">// Assemble decimal in reverse order.</span>
    <a id="L110"></a>var b [32]byte;
    <a id="L111"></a>bp := len(b);
    <a id="L112"></a>for ; u &gt; 0; u /= 10 {
        <a id="L113"></a>bp--;
        <a id="L114"></a>b[bp] = byte(u%10) + &#39;0&#39;;
    <a id="L115"></a>}

    <a id="L117"></a>if i &lt; 0 {
        <a id="L118"></a>bp--;
        <a id="L119"></a>b[bp] = &#39;-&#39;;
    <a id="L120"></a>}

    <a id="L122"></a>return string(b[bp:len(b)]);
<a id="L123"></a>}

<a id="L125"></a>func (w Waitmsg) String() string {
    <a id="L126"></a><span class="comment">// TODO(austin) Use signal names when possible?</span>
    <a id="L127"></a>res := &#34;&#34;;
    <a id="L128"></a>switch {
    <a id="L129"></a>case w.Exited():
        <a id="L130"></a>res = &#34;exit status &#34; + itod(w.ExitStatus())
    <a id="L131"></a>case w.Signaled():
        <a id="L132"></a>res = &#34;signal &#34; + itod(w.Signal())
    <a id="L133"></a>case w.Stopped():
        <a id="L134"></a>res = &#34;stop signal &#34; + itod(w.StopSignal());
        <a id="L135"></a>if w.StopSignal() == syscall.SIGTRAP &amp;&amp; w.TrapCause() != 0 {
            <a id="L136"></a>res += &#34; (trap &#34; + itod(w.TrapCause()) + &#34;)&#34;
        <a id="L137"></a>}
    <a id="L138"></a>case w.Continued():
        <a id="L139"></a>res = &#34;continued&#34;
    <a id="L140"></a>}
    <a id="L141"></a>if w.CoreDump() {
        <a id="L142"></a>res += &#34; (core dumped)&#34;
    <a id="L143"></a>}
    <a id="L144"></a>return res;
<a id="L145"></a>}

<a id="L147"></a><span class="comment">// Getpid returns the process id of the caller.</span>
<a id="L148"></a>func Getpid() int { return syscall.Getpid() }

<a id="L150"></a><span class="comment">// Getppid returns the process id of the caller&#39;s parent.</span>
<a id="L151"></a>func Getppid() int { return syscall.Getppid() }
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
