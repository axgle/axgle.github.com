<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/exec.go</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/syscall/exec.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Fork, exec, wait, etc.</span>

<a id="L7"></a>package syscall

<a id="L9"></a>import (
    <a id="L10"></a>&#34;sync&#34;;
    <a id="L11"></a>&#34;unsafe&#34;;
<a id="L12"></a>)

<a id="L14"></a><span class="comment">// Lock synchronizing creation of new file descriptors with fork.</span>
<a id="L15"></a><span class="comment">//</span>
<a id="L16"></a><span class="comment">// We want the child in a fork/exec sequence to inherit only the</span>
<a id="L17"></a><span class="comment">// file descriptors we intend.  To do that, we mark all file</span>
<a id="L18"></a><span class="comment">// descriptors close-on-exec and then, in the child, explicitly</span>
<a id="L19"></a><span class="comment">// unmark the ones we want the exec&#39;ed program to keep.</span>
<a id="L20"></a><span class="comment">// Unix doesn&#39;t make this easy: there is, in general, no way to</span>
<a id="L21"></a><span class="comment">// allocate a new file descriptor close-on-exec.  Instead you</span>
<a id="L22"></a><span class="comment">// have to allocate the descriptor and then mark it close-on-exec.</span>
<a id="L23"></a><span class="comment">// If a fork happens between those two events, the child&#39;s exec</span>
<a id="L24"></a><span class="comment">// will inherit an unwanted file descriptor.</span>
<a id="L25"></a><span class="comment">//</span>
<a id="L26"></a><span class="comment">// This lock solves that race: the create new fd/mark close-on-exec</span>
<a id="L27"></a><span class="comment">// operation is done holding ForkLock for reading, and the fork itself</span>
<a id="L28"></a><span class="comment">// is done holding ForkLock for writing.  At least, that&#39;s the idea.</span>
<a id="L29"></a><span class="comment">// There are some complications.</span>
<a id="L30"></a><span class="comment">//</span>
<a id="L31"></a><span class="comment">// Some system calls that create new file descriptors can block</span>
<a id="L32"></a><span class="comment">// for arbitrarily long times: open on a hung NFS server or named</span>
<a id="L33"></a><span class="comment">// pipe, accept on a socket, and so on.  We can&#39;t reasonably grab</span>
<a id="L34"></a><span class="comment">// the lock across those operations.</span>
<a id="L35"></a><span class="comment">//</span>
<a id="L36"></a><span class="comment">// It is worse to inherit some file descriptors than others.</span>
<a id="L37"></a><span class="comment">// If a non-malicious child accidentally inherits an open ordinary file,</span>
<a id="L38"></a><span class="comment">// that&#39;s not a big deal.  On the other hand, if a long-lived child</span>
<a id="L39"></a><span class="comment">// accidentally inherits the write end of a pipe, then the reader</span>
<a id="L40"></a><span class="comment">// of that pipe will not see EOF until that child exits, potentially</span>
<a id="L41"></a><span class="comment">// causing the parent program to hang.  This is a common problem</span>
<a id="L42"></a><span class="comment">// in threaded C programs that use popen.</span>
<a id="L43"></a><span class="comment">//</span>
<a id="L44"></a><span class="comment">// Luckily, the file descriptors that are most important not to</span>
<a id="L45"></a><span class="comment">// inherit are not the ones that can take an arbitrarily long time</span>
<a id="L46"></a><span class="comment">// to create: pipe returns instantly, and the net package uses</span>
<a id="L47"></a><span class="comment">// non-blocking I/O to accept on a listening socket.</span>
<a id="L48"></a><span class="comment">// The rules for which file descriptor-creating operations use the</span>
<a id="L49"></a><span class="comment">// ForkLock are as follows:</span>
<a id="L50"></a><span class="comment">//</span>
<a id="L51"></a><span class="comment">// 1) Pipe.    Does not block.  Use the ForkLock.</span>
<a id="L52"></a><span class="comment">// 2) Socket.  Does not block.  Use the ForkLock.</span>
<a id="L53"></a><span class="comment">// 3) Accept.  If using non-blocking mode, use the ForkLock.</span>
<a id="L54"></a><span class="comment">//             Otherwise, live with the race.</span>
<a id="L55"></a><span class="comment">// 4) Open.    Can block.  Use O_CLOEXEC if available (Linux).</span>
<a id="L56"></a><span class="comment">//             Otherwise, live with the race.</span>
<a id="L57"></a><span class="comment">// 5) Dup.     Does not block.  Use the ForkLock.</span>
<a id="L58"></a><span class="comment">//             On Linux, could use fcntl F_DUPFD_CLOEXEC</span>
<a id="L59"></a><span class="comment">//             instead of the ForkLock, but only for dup(fd, -1).</span>

<a id="L61"></a>var ForkLock sync.RWMutex

<a id="L63"></a><span class="comment">// Convert array of string to array</span>
<a id="L64"></a><span class="comment">// of NUL-terminated byte pointer.</span>
<a id="L65"></a>func StringArrayPtr(ss []string) []*byte {
    <a id="L66"></a>bb := make([]*byte, len(ss)+1);
    <a id="L67"></a>for i := 0; i &lt; len(ss); i++ {
        <a id="L68"></a>bb[i] = StringBytePtr(ss[i])
    <a id="L69"></a>}
    <a id="L70"></a>bb[len(ss)] = nil;
    <a id="L71"></a>return bb;
<a id="L72"></a>}

<a id="L74"></a>func CloseOnExec(fd int) { fcntl(fd, F_SETFD, FD_CLOEXEC) }

<a id="L76"></a>func SetNonblock(fd int, nonblocking bool) (errno int) {
    <a id="L77"></a>flag, err := fcntl(fd, F_GETFL, 0);
    <a id="L78"></a>if err != 0 {
        <a id="L79"></a>return err
    <a id="L80"></a>}
    <a id="L81"></a>if nonblocking {
        <a id="L82"></a>flag |= O_NONBLOCK
    <a id="L83"></a>} else {
        <a id="L84"></a>flag &amp;= ^O_NONBLOCK
    <a id="L85"></a>}
    <a id="L86"></a>_, err = fcntl(fd, F_SETFL, flag);
    <a id="L87"></a>return err;
<a id="L88"></a>}


<a id="L91"></a><span class="comment">// Fork, dup fd onto 0..len(fd), and exec(argv0, argvv, envv) in child.</span>
<a id="L92"></a><span class="comment">// If a dup or exec fails, write the errno int to pipe.</span>
<a id="L93"></a><span class="comment">// (Pipe is close-on-exec so if exec succeeds, it will be closed.)</span>
<a id="L94"></a><span class="comment">// In the child, this function must not acquire any locks, because</span>
<a id="L95"></a><span class="comment">// they might have been locked at the time of the fork.  This means</span>
<a id="L96"></a><span class="comment">// no rescheduling, no malloc calls, and no new stack segments.</span>
<a id="L97"></a><span class="comment">// The calls to RawSyscall are okay because they are assembly</span>
<a id="L98"></a><span class="comment">// functions that do not grow the stack.</span>
<a id="L99"></a>func forkAndExecInChild(argv0 *byte, argv []*byte, envv []*byte, traceme bool, dir *byte, fd []int, pipe int) (pid int, err int) {
    <a id="L100"></a><span class="comment">// Declare all variables at top in case any</span>
    <a id="L101"></a><span class="comment">// declarations require heap allocation (e.g., err1).</span>
    <a id="L102"></a>var r1, r2, err1 uintptr;
    <a id="L103"></a>var nextfd int;
    <a id="L104"></a>var i int;

    <a id="L106"></a>darwin := OS == &#34;darwin&#34;;

    <a id="L108"></a><span class="comment">// About to call fork.</span>
    <a id="L109"></a><span class="comment">// No more allocation or calls of non-assembly functions.</span>
    <a id="L110"></a>r1, r2, err1 = RawSyscall(SYS_FORK, 0, 0, 0);
    <a id="L111"></a>if err1 != 0 {
        <a id="L112"></a>return 0, int(err1)
    <a id="L113"></a>}

    <a id="L115"></a><span class="comment">// On Darwin:</span>
    <a id="L116"></a><span class="comment">//	r1 = child pid in both parent and child.</span>
    <a id="L117"></a><span class="comment">//	r2 = 0 in parent, 1 in child.</span>
    <a id="L118"></a><span class="comment">// Convert to normal Unix r1 = 0 in child.</span>
    <a id="L119"></a>if darwin &amp;&amp; r2 == 1 {
        <a id="L120"></a>r1 = 0
    <a id="L121"></a>}

    <a id="L123"></a>if r1 != 0 {
        <a id="L124"></a><span class="comment">// parent; return PID</span>
        <a id="L125"></a>return int(r1), 0
    <a id="L126"></a>}

    <a id="L128"></a><span class="comment">// Fork succeeded, now in child.</span>

    <a id="L130"></a><span class="comment">// Enable tracing if requested.</span>
    <a id="L131"></a>if traceme {
        <a id="L132"></a>_, _, err1 = RawSyscall(SYS_PTRACE, uintptr(PTRACE_TRACEME), 0, 0);
        <a id="L133"></a>if err1 != 0 {
            <a id="L134"></a>goto childerror
        <a id="L135"></a>}
    <a id="L136"></a>}

    <a id="L138"></a><span class="comment">// Chdir</span>
    <a id="L139"></a>if dir != nil {
        <a id="L140"></a>_, _, err1 = RawSyscall(SYS_CHDIR, uintptr(unsafe.Pointer(dir)), 0, 0);
        <a id="L141"></a>if err1 != 0 {
            <a id="L142"></a>goto childerror
        <a id="L143"></a>}
    <a id="L144"></a>}

    <a id="L146"></a><span class="comment">// Pass 1: look for fd[i] &lt; i and move those up above len(fd)</span>
    <a id="L147"></a><span class="comment">// so that pass 2 won&#39;t stomp on an fd it needs later.</span>
    <a id="L148"></a>nextfd = int(len(fd));
    <a id="L149"></a>if pipe &lt; nextfd {
        <a id="L150"></a>_, _, err1 = RawSyscall(SYS_DUP2, uintptr(pipe), uintptr(nextfd), 0);
        <a id="L151"></a>if err1 != 0 {
            <a id="L152"></a>goto childerror
        <a id="L153"></a>}
        <a id="L154"></a>RawSyscall(SYS_FCNTL, uintptr(nextfd), F_SETFD, FD_CLOEXEC);
        <a id="L155"></a>pipe = nextfd;
        <a id="L156"></a>nextfd++;
    <a id="L157"></a>}
    <a id="L158"></a>for i = 0; i &lt; len(fd); i++ {
        <a id="L159"></a>if fd[i] &gt;= 0 &amp;&amp; fd[i] &lt; int(i) {
            <a id="L160"></a>_, _, err1 = RawSyscall(SYS_DUP2, uintptr(fd[i]), uintptr(nextfd), 0);
            <a id="L161"></a>if err1 != 0 {
                <a id="L162"></a>goto childerror
            <a id="L163"></a>}
            <a id="L164"></a>RawSyscall(SYS_FCNTL, uintptr(nextfd), F_SETFD, FD_CLOEXEC);
            <a id="L165"></a>fd[i] = nextfd;
            <a id="L166"></a>nextfd++;
            <a id="L167"></a>if nextfd == pipe { <span class="comment">// don&#39;t stomp on pipe</span>
                <a id="L168"></a>nextfd++
            <a id="L169"></a>}
        <a id="L170"></a>}
    <a id="L171"></a>}

    <a id="L173"></a><span class="comment">// Pass 2: dup fd[i] down onto i.</span>
    <a id="L174"></a>for i = 0; i &lt; len(fd); i++ {
        <a id="L175"></a>if fd[i] == -1 {
            <a id="L176"></a>RawSyscall(SYS_CLOSE, uintptr(i), 0, 0);
            <a id="L177"></a>continue;
        <a id="L178"></a>}
        <a id="L179"></a>if fd[i] == int(i) {
            <a id="L180"></a><span class="comment">// dup2(i, i) won&#39;t clear close-on-exec flag on Linux,</span>
            <a id="L181"></a><span class="comment">// probably not elsewhere either.</span>
            <a id="L182"></a>_, _, err1 = RawSyscall(SYS_FCNTL, uintptr(fd[i]), F_SETFD, 0);
            <a id="L183"></a>if err1 != 0 {
                <a id="L184"></a>goto childerror
            <a id="L185"></a>}
            <a id="L186"></a>continue;
        <a id="L187"></a>}
        <a id="L188"></a><span class="comment">// The new fd is created NOT close-on-exec,</span>
        <a id="L189"></a><span class="comment">// which is exactly what we want.</span>
        <a id="L190"></a>_, _, err1 = RawSyscall(SYS_DUP2, uintptr(fd[i]), uintptr(i), 0);
        <a id="L191"></a>if err1 != 0 {
            <a id="L192"></a>goto childerror
        <a id="L193"></a>}
    <a id="L194"></a>}

    <a id="L196"></a><span class="comment">// By convention, we don&#39;t close-on-exec the fds we are</span>
    <a id="L197"></a><span class="comment">// started with, so if len(fd) &lt; 3, close 0, 1, 2 as needed.</span>
    <a id="L198"></a><span class="comment">// Programs that know they inherit fds &gt;= 3 will need</span>
    <a id="L199"></a><span class="comment">// to set them close-on-exec.</span>
    <a id="L200"></a>for i = len(fd); i &lt; 3; i++ {
        <a id="L201"></a>RawSyscall(SYS_CLOSE, uintptr(i), 0, 0)
    <a id="L202"></a>}

    <a id="L204"></a><span class="comment">// Time to exec.</span>
    <a id="L205"></a>_, _, err1 = RawSyscall(SYS_EXECVE,
        <a id="L206"></a>uintptr(unsafe.Pointer(argv0)),
        <a id="L207"></a>uintptr(unsafe.Pointer(&amp;argv[0])),
        <a id="L208"></a>uintptr(unsafe.Pointer(&amp;envv[0])));

<a id="L210"></a>childerror:
    <a id="L211"></a><span class="comment">// send error code on pipe</span>
    <a id="L212"></a>RawSyscall(SYS_WRITE, uintptr(pipe), uintptr(unsafe.Pointer(&amp;err1)), uintptr(unsafe.Sizeof(err1)));
    <a id="L213"></a>for {
        <a id="L214"></a>RawSyscall(SYS_EXIT, 253, 0, 0)
    <a id="L215"></a>}

    <a id="L217"></a><span class="comment">// Calling panic is not actually safe,</span>
    <a id="L218"></a><span class="comment">// but the for loop above won&#39;t break</span>
    <a id="L219"></a><span class="comment">// and this shuts up the compiler.</span>
    <a id="L220"></a>panic(&#34;unreached&#34;);
<a id="L221"></a>}

<a id="L223"></a>func forkExec(argv0 string, argv []string, envv []string, traceme bool, dir string, fd []int) (pid int, err int) {
    <a id="L224"></a>var p [2]int;
    <a id="L225"></a>var n int;
    <a id="L226"></a>var err1 uintptr;
    <a id="L227"></a>var wstatus WaitStatus;

    <a id="L229"></a>p[0] = -1;
    <a id="L230"></a>p[1] = -1;

    <a id="L232"></a><span class="comment">// Convert args to C form.</span>
    <a id="L233"></a>argv0p := StringBytePtr(argv0);
    <a id="L234"></a>argvp := StringArrayPtr(argv);
    <a id="L235"></a>envvp := StringArrayPtr(envv);
    <a id="L236"></a>var dirp *byte;
    <a id="L237"></a>if len(dir) &gt; 0 {
        <a id="L238"></a>dirp = StringBytePtr(dir)
    <a id="L239"></a>}

    <a id="L241"></a><span class="comment">// Acquire the fork lock so that no other threads</span>
    <a id="L242"></a><span class="comment">// create new fds that are not yet close-on-exec</span>
    <a id="L243"></a><span class="comment">// before we fork.</span>
    <a id="L244"></a>ForkLock.Lock();

    <a id="L246"></a><span class="comment">// Allocate child status pipe close on exec.</span>
    <a id="L247"></a>if err = Pipe(&amp;p); err != 0 {
        <a id="L248"></a>goto error
    <a id="L249"></a>}
    <a id="L250"></a>if _, err = fcntl(p[0], F_SETFD, FD_CLOEXEC); err != 0 {
        <a id="L251"></a>goto error
    <a id="L252"></a>}
    <a id="L253"></a>if _, err = fcntl(p[1], F_SETFD, FD_CLOEXEC); err != 0 {
        <a id="L254"></a>goto error
    <a id="L255"></a>}

    <a id="L257"></a><span class="comment">// Kick off child.</span>
    <a id="L258"></a>pid, err = forkAndExecInChild(argv0p, argvp, envvp, traceme, dirp, fd, p[1]);
    <a id="L259"></a>if err != 0 {
    <a id="L260"></a>error:
        <a id="L261"></a>if p[0] &gt;= 0 {
            <a id="L262"></a>Close(p[0]);
            <a id="L263"></a>Close(p[1]);
        <a id="L264"></a>}
        <a id="L265"></a>ForkLock.Unlock();
        <a id="L266"></a>return 0, err;
    <a id="L267"></a>}
    <a id="L268"></a>ForkLock.Unlock();

    <a id="L270"></a><span class="comment">// Read child error status from pipe.</span>
    <a id="L271"></a>Close(p[1]);
    <a id="L272"></a>n, err = read(p[0], (*byte)(unsafe.Pointer(&amp;err1)), unsafe.Sizeof(err1));
    <a id="L273"></a>Close(p[0]);
    <a id="L274"></a>if err != 0 || n != 0 {
        <a id="L275"></a>if n == unsafe.Sizeof(err1) {
            <a id="L276"></a>err = int(err1)
        <a id="L277"></a>}
        <a id="L278"></a>if err == 0 {
            <a id="L279"></a>err = EPIPE
        <a id="L280"></a>}

        <a id="L282"></a><span class="comment">// Child failed; wait for it to exit, to make sure</span>
        <a id="L283"></a><span class="comment">// the zombies don&#39;t accumulate.</span>
        <a id="L284"></a>_, err1 := Wait4(pid, &amp;wstatus, 0, nil);
        <a id="L285"></a>for err1 == EINTR {
            <a id="L286"></a>_, err1 = Wait4(pid, &amp;wstatus, 0, nil)
        <a id="L287"></a>}
        <a id="L288"></a>return 0, err;
    <a id="L289"></a>}

    <a id="L291"></a><span class="comment">// Read got EOF, so pipe closed on exec, so exec succeeded.</span>
    <a id="L292"></a>return pid, 0;
<a id="L293"></a>}

<a id="L295"></a><span class="comment">// Combination of fork and exec, careful to be thread safe.</span>
<a id="L296"></a>func ForkExec(argv0 string, argv []string, envv []string, dir string, fd []int) (pid int, err int) {
    <a id="L297"></a>return forkExec(argv0, argv, envv, false, dir, fd)
<a id="L298"></a>}

<a id="L300"></a><span class="comment">// PtraceForkExec is like ForkExec, but starts the child in a traced state.</span>
<a id="L301"></a>func PtraceForkExec(argv0 string, argv []string, envv []string, dir string, fd []int) (pid int, err int) {
    <a id="L302"></a>return forkExec(argv0, argv, envv, true, dir, fd)
<a id="L303"></a>}

<a id="L305"></a><span class="comment">// Ordinary exec.</span>
<a id="L306"></a>func Exec(argv0 string, argv []string, envv []string) (err int) {
    <a id="L307"></a>_, _, err1 := RawSyscall(SYS_EXECVE,
        <a id="L308"></a>uintptr(unsafe.Pointer(StringBytePtr(argv0))),
        <a id="L309"></a>uintptr(unsafe.Pointer(&amp;StringArrayPtr(argv)[0])),
        <a id="L310"></a>uintptr(unsafe.Pointer(&amp;StringArrayPtr(envv)[0])));
    <a id="L311"></a>return int(err1);
<a id="L312"></a>}
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
