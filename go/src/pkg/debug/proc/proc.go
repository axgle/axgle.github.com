<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/proc/proc.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

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
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/debug/proc/proc.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Package ptrace provides a platform-independent interface for</span>
<a id="L6"></a><span class="comment">// tracing and controlling running processes.  It supports</span>
<a id="L7"></a><span class="comment">// multi-threaded processes and provides typical low-level debugging</span>
<a id="L8"></a><span class="comment">// controls such as breakpoints, single stepping, and manipulating</span>
<a id="L9"></a><span class="comment">// memory and registers.</span>
<a id="L10"></a>package proc

<a id="L12"></a><span class="comment">// TODO(rsc): Have to import everything that proc_linux.go</span>
<a id="L13"></a><span class="comment">// and proc_darwin.go do, because deps.bash only looks at</span>
<a id="L14"></a><span class="comment">// this file.</span>
<a id="L15"></a>import (
    <a id="L16"></a>_ &#34;container/vector&#34;;
    <a id="L17"></a>_ &#34;fmt&#34;;
    <a id="L18"></a>_ &#34;io&#34;;
    <a id="L19"></a>&#34;os&#34;;
    <a id="L20"></a>_ &#34;runtime&#34;;
    <a id="L21"></a>&#34;strconv&#34;;
    <a id="L22"></a>_ &#34;strings&#34;;
    <a id="L23"></a>_ &#34;sync&#34;;
    <a id="L24"></a>_ &#34;syscall&#34;;
<a id="L25"></a>)

<a id="L27"></a>type Word uint64

<a id="L29"></a><span class="comment">// A Cause explains why a thread is stopped.</span>
<a id="L30"></a>type Cause interface {
    <a id="L31"></a>String() string;
<a id="L32"></a>}

<a id="L34"></a><span class="comment">// Regs is a set of named machine registers, including a program</span>
<a id="L35"></a><span class="comment">// counter, link register, and stack pointer.</span>
<a id="L36"></a><span class="comment">//</span>
<a id="L37"></a><span class="comment">// TODO(austin) There&#39;s quite a proliferation of methods here.  We</span>
<a id="L38"></a><span class="comment">// could make a Reg interface with Get and Set and make this just PC,</span>
<a id="L39"></a><span class="comment">// Link, SP, Names, and Reg.  We could also put Index in Reg and that</span>
<a id="L40"></a><span class="comment">// makes it easy to get the index of things like the PC (currently</span>
<a id="L41"></a><span class="comment">// there&#39;s just no way to know that).  This would also let us include</span>
<a id="L42"></a><span class="comment">// other per-register information like how to print it.</span>
<a id="L43"></a>type Regs interface {
    <a id="L44"></a><span class="comment">// PC returns the value of the program counter.</span>
    <a id="L45"></a>PC() Word;

    <a id="L47"></a><span class="comment">// SetPC sets the program counter to val.</span>
    <a id="L48"></a>SetPC(val Word) os.Error;

    <a id="L50"></a><span class="comment">// Link returns the link register, if any.</span>
    <a id="L51"></a>Link() Word;

    <a id="L53"></a><span class="comment">// SetLink sets the link register to val.</span>
    <a id="L54"></a>SetLink(val Word) os.Error;

    <a id="L56"></a><span class="comment">// SP returns the value of the stack pointer.</span>
    <a id="L57"></a>SP() Word;

    <a id="L59"></a><span class="comment">// SetSP sets the stack pointer register to val.</span>
    <a id="L60"></a>SetSP(val Word) os.Error;

    <a id="L62"></a><span class="comment">// Names returns the names of all of the registers.</span>
    <a id="L63"></a>Names() []string;

    <a id="L65"></a><span class="comment">// Get returns the value of a register, where i corresponds to</span>
    <a id="L66"></a><span class="comment">// the index of the register&#39;s name in the array returned by</span>
    <a id="L67"></a><span class="comment">// Names.</span>
    <a id="L68"></a>Get(i int) Word;

    <a id="L70"></a><span class="comment">// Set sets the value of a register.</span>
    <a id="L71"></a>Set(i int, val Word) os.Error;
<a id="L72"></a>}

<a id="L74"></a><span class="comment">// Thread is a thread in the process being traced.</span>
<a id="L75"></a>type Thread interface {
    <a id="L76"></a><span class="comment">// Step steps this thread by a single instruction.  The thread</span>
    <a id="L77"></a><span class="comment">// must be stopped.  If the thread is currently stopped on a</span>
    <a id="L78"></a><span class="comment">// breakpoint, this will step over the breakpoint.</span>
    <a id="L79"></a><span class="comment">//</span>
    <a id="L80"></a><span class="comment">// XXX What if it&#39;s stopped because of a signal?</span>
    <a id="L81"></a>Step() os.Error;

    <a id="L83"></a><span class="comment">// Stopped returns the reason that this thread is stopped.  It</span>
    <a id="L84"></a><span class="comment">// is an error is the thread not stopped.</span>
    <a id="L85"></a>Stopped() (Cause, os.Error);

    <a id="L87"></a><span class="comment">// Regs retrieves the current register values from this</span>
    <a id="L88"></a><span class="comment">// thread.  The thread must be stopped.</span>
    <a id="L89"></a>Regs() (Regs, os.Error);

    <a id="L91"></a><span class="comment">// Peek reads len(out) bytes from the address addr in this</span>
    <a id="L92"></a><span class="comment">// thread into out.  The thread must be stopped.  It returns</span>
    <a id="L93"></a><span class="comment">// the number of bytes successfully read.  If an error occurs,</span>
    <a id="L94"></a><span class="comment">// such as attempting to read unmapped memory, this count</span>
    <a id="L95"></a><span class="comment">// could be short and an error will be returned.  If this does</span>
    <a id="L96"></a><span class="comment">// encounter unmapped memory, it will read up to the byte</span>
    <a id="L97"></a><span class="comment">// preceding the unmapped area.</span>
    <a id="L98"></a>Peek(addr Word, out []byte) (int, os.Error);

    <a id="L100"></a><span class="comment">// Poke writes b to the address addr in this thread.  The</span>
    <a id="L101"></a><span class="comment">// thread must be stopped.  It returns the number of bytes</span>
    <a id="L102"></a><span class="comment">// successfully written.  If an error occurs, such as</span>
    <a id="L103"></a><span class="comment">// attempting to write to unmapped memory, this count could be</span>
    <a id="L104"></a><span class="comment">// short and an error will be returned.  If this does</span>
    <a id="L105"></a><span class="comment">// encounter unmapped memory, it will write up to the byte</span>
    <a id="L106"></a><span class="comment">// preceding the unmapped area.</span>
    <a id="L107"></a>Poke(addr Word, b []byte) (int, os.Error);
<a id="L108"></a>}

<a id="L110"></a><span class="comment">// Process is a process being traced.  It consists of a set of</span>
<a id="L111"></a><span class="comment">// threads.  A process can be running, stopped, or terminated.  The</span>
<a id="L112"></a><span class="comment">// process&#39;s state extends to all of its threads.</span>
<a id="L113"></a>type Process interface {
    <a id="L114"></a><span class="comment">// Threads returns an array of all threads in this process.</span>
    <a id="L115"></a>Threads() []Thread;

    <a id="L117"></a><span class="comment">// AddBreakpoint creates a new breakpoint at program counter</span>
    <a id="L118"></a><span class="comment">// pc.  Breakpoints can only be created when the process is</span>
    <a id="L119"></a><span class="comment">// stopped.  It is an error if a breakpoint already exists at</span>
    <a id="L120"></a><span class="comment">// pc.</span>
    <a id="L121"></a>AddBreakpoint(pc Word) os.Error;

    <a id="L123"></a><span class="comment">// RemoveBreakpoint removes the breakpoint at the program</span>
    <a id="L124"></a><span class="comment">// counter pc.  It is an error if no breakpoint exists at pc.</span>
    <a id="L125"></a>RemoveBreakpoint(pc Word) os.Error;

    <a id="L127"></a><span class="comment">// Stop stops all running threads in this process before</span>
    <a id="L128"></a><span class="comment">// returning.</span>
    <a id="L129"></a>Stop() os.Error;

    <a id="L131"></a><span class="comment">// Continue resumes execution of all threads in this process.</span>
    <a id="L132"></a><span class="comment">// Any thread that is stopped on a breakpoint will be stepped</span>
    <a id="L133"></a><span class="comment">// over that breakpoint.  Any thread that is stopped because</span>
    <a id="L134"></a><span class="comment">// of a signal (other than SIGSTOP or SIGTRAP) will receive</span>
    <a id="L135"></a><span class="comment">// the pending signal.</span>
    <a id="L136"></a>Continue() os.Error;

    <a id="L138"></a><span class="comment">// WaitStop waits until all threads in process p are stopped</span>
    <a id="L139"></a><span class="comment">// as a result of some thread hitting a breakpoint, receiving</span>
    <a id="L140"></a><span class="comment">// a signal, creating a new thread, or exiting.</span>
    <a id="L141"></a>WaitStop() os.Error;

    <a id="L143"></a><span class="comment">// Detach detaches from this process.  All stopped threads</span>
    <a id="L144"></a><span class="comment">// will be resumed.</span>
    <a id="L145"></a>Detach() os.Error;
<a id="L146"></a>}

<a id="L148"></a><span class="comment">// Stopped is a stop cause used for threads that are stopped either by</span>
<a id="L149"></a><span class="comment">// user request (e.g., from the Stop method or after single stepping),</span>
<a id="L150"></a><span class="comment">// or that are stopped because some other thread caused the program to</span>
<a id="L151"></a><span class="comment">// stop.</span>
<a id="L152"></a>type Stopped struct{}

<a id="L154"></a>func (c Stopped) String() string { return &#34;stopped&#34; }

<a id="L156"></a><span class="comment">// Breakpoint is a stop cause resulting from a thread reaching a set</span>
<a id="L157"></a><span class="comment">// breakpoint.</span>
<a id="L158"></a>type Breakpoint Word

<a id="L160"></a><span class="comment">// PC returns the program counter that the program is stopped at.</span>
<a id="L161"></a>func (c Breakpoint) PC() Word { return Word(c) }

<a id="L163"></a>func (c Breakpoint) String() string {
    <a id="L164"></a>return &#34;breakpoint at 0x&#34; + strconv.Uitob64(uint64(c.PC()), 16)
<a id="L165"></a>}

<a id="L167"></a><span class="comment">// Signal is a stop cause resulting from a thread receiving a signal.</span>
<a id="L168"></a><span class="comment">// When the process is continued, the signal will be delivered.</span>
<a id="L169"></a>type Signal string

<a id="L171"></a><span class="comment">// Signal returns the signal being delivered to the thread.</span>
<a id="L172"></a>func (c Signal) Name() string { return string(c) }

<a id="L174"></a>func (c Signal) String() string { return c.Name() }

<a id="L176"></a><span class="comment">// ThreadCreate is a stop cause returned from an existing thread when</span>
<a id="L177"></a><span class="comment">// it creates a new thread.  The new thread exists in a primordial</span>
<a id="L178"></a><span class="comment">// form at this point and will begin executing in earnest when the</span>
<a id="L179"></a><span class="comment">// process is continued.</span>
<a id="L180"></a>type ThreadCreate struct {
    <a id="L181"></a>thread Thread;
<a id="L182"></a>}

<a id="L184"></a>func (c *ThreadCreate) NewThread() Thread { return c.thread }

<a id="L186"></a>func (c *ThreadCreate) String() string { return &#34;thread create&#34; }

<a id="L188"></a><span class="comment">// ThreadExit is a stop cause resulting from a thread exiting.  When</span>
<a id="L189"></a><span class="comment">// this cause first arises, the thread will still be in the list of</span>
<a id="L190"></a><span class="comment">// process threads and its registers and memory will still be</span>
<a id="L191"></a><span class="comment">// accessible.</span>
<a id="L192"></a>type ThreadExit struct {
    <a id="L193"></a>exitStatus int;
    <a id="L194"></a>signal     string;
<a id="L195"></a>}

<a id="L197"></a><span class="comment">// Exited returns true if the thread exited normally.</span>
<a id="L198"></a>func (c *ThreadExit) Exited() bool { return c.exitStatus != -1 }

<a id="L200"></a><span class="comment">// ExitStatus returns the exit status of the thread if it exited</span>
<a id="L201"></a><span class="comment">// normally or -1 otherwise.</span>
<a id="L202"></a>func (c *ThreadExit) ExitStatus() int { return c.exitStatus }

<a id="L204"></a><span class="comment">// Signaled returns true if the thread was terminated by a signal.</span>
<a id="L205"></a>func (c *ThreadExit) Signaled() bool { return c.exitStatus == -1 }

<a id="L207"></a><span class="comment">// StopSignal returns the signal that terminated the thread, or &#34;&#34; if</span>
<a id="L208"></a><span class="comment">// it was not terminated by a signal.</span>
<a id="L209"></a>func (c *ThreadExit) StopSignal() string { return c.signal }

<a id="L211"></a>func (c *ThreadExit) String() string {
    <a id="L212"></a>res := &#34;thread exited &#34;;
    <a id="L213"></a>switch {
    <a id="L214"></a>case c.Exited():
        <a id="L215"></a>res += &#34;with status &#34; + strconv.Itoa(c.ExitStatus())
    <a id="L216"></a>case c.Signaled():
        <a id="L217"></a>res += &#34;from signal &#34; + c.StopSignal()
    <a id="L218"></a>default:
        <a id="L219"></a>res += &#34;from unknown cause&#34;
    <a id="L220"></a>}
    <a id="L221"></a>return res;
<a id="L222"></a>}
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
