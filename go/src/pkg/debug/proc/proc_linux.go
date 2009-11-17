<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/proc/proc_linux.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/debug/proc/proc_linux.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package proc

<a id="L7"></a><span class="comment">// TODO(rsc): Imports here after to be in proc.go too in order</span>
<a id="L8"></a><span class="comment">// for deps.bash to get the right answer.</span>
<a id="L9"></a>import (
    <a id="L10"></a>&#34;container/vector&#34;;
    <a id="L11"></a>&#34;fmt&#34;;
    <a id="L12"></a>&#34;io&#34;;
    <a id="L13"></a>&#34;os&#34;;
    <a id="L14"></a>&#34;runtime&#34;;
    <a id="L15"></a>&#34;strconv&#34;;
    <a id="L16"></a>&#34;strings&#34;;
    <a id="L17"></a>&#34;sync&#34;;
    <a id="L18"></a>&#34;syscall&#34;;
<a id="L19"></a>)

<a id="L21"></a><span class="comment">// This is an implementation of the process tracing interface using</span>
<a id="L22"></a><span class="comment">// Linux&#39;s ptrace(2) interface.  The implementation is multi-threaded.</span>
<a id="L23"></a><span class="comment">// Each attached process has an associated monitor thread, and each</span>
<a id="L24"></a><span class="comment">// running attached thread has an associated &#34;wait&#34; thread.  The wait</span>
<a id="L25"></a><span class="comment">// thread calls wait4 on the thread&#39;s TID and reports any wait events</span>
<a id="L26"></a><span class="comment">// or errors via &#34;debug events&#34;.  The monitor thread consumes these</span>
<a id="L27"></a><span class="comment">// wait events and updates the internally maintained state of each</span>
<a id="L28"></a><span class="comment">// thread.  All ptrace calls must run in the monitor thread, so the</span>
<a id="L29"></a><span class="comment">// monitor executes closures received on the debugReq channel.</span>
<a id="L30"></a><span class="comment">//</span>
<a id="L31"></a><span class="comment">// As ptrace&#39;s documentation is somewhat light, this is heavily based</span>
<a id="L32"></a><span class="comment">// on information gleaned from the implementation of ptrace found at</span>
<a id="L33"></a><span class="comment">//   http://lxr.linux.no/linux+v2.6.30/kernel/ptrace.c</span>
<a id="L34"></a><span class="comment">//   http://lxr.linux.no/linux+v2.6.30/arch/x86/kernel/ptrace.c#L854</span>
<a id="L35"></a><span class="comment">// as well as experimentation and examination of gdb&#39;s behavior.</span>

<a id="L37"></a>const (
    <a id="L38"></a>trace    = false;
    <a id="L39"></a>traceIP  = false;
    <a id="L40"></a>traceMem = false;
<a id="L41"></a>)

<a id="L43"></a><span class="comment">/*</span>
<a id="L44"></a><span class="comment"> * Thread state</span>
<a id="L45"></a><span class="comment"> */</span>

<a id="L47"></a><span class="comment">// Each thread can be in one of the following set of states.</span>
<a id="L48"></a><span class="comment">// Each state satisfies</span>
<a id="L49"></a><span class="comment">//  isRunning() || isStopped() || isZombie() || isTerminal().</span>
<a id="L50"></a><span class="comment">//</span>
<a id="L51"></a><span class="comment">// Running threads can be sent signals and must be waited on, but they</span>
<a id="L52"></a><span class="comment">// cannot be inspected using ptrace.</span>
<a id="L53"></a><span class="comment">//</span>
<a id="L54"></a><span class="comment">// Stopped threads can be inspected and continued, but cannot be</span>
<a id="L55"></a><span class="comment">// meaningfully waited on.  They can be sent signals, but the signals</span>
<a id="L56"></a><span class="comment">// will be queued until they are running again.</span>
<a id="L57"></a><span class="comment">//</span>
<a id="L58"></a><span class="comment">// Zombie threads cannot be inspected, continued, or sent signals (and</span>
<a id="L59"></a><span class="comment">// therefore they cannot be stopped), but they must be waited on.</span>
<a id="L60"></a><span class="comment">//</span>
<a id="L61"></a><span class="comment">// Terminal threads no longer exist in the OS and thus you can&#39;t do</span>
<a id="L62"></a><span class="comment">// anything with them.</span>
<a id="L63"></a>type threadState string

<a id="L65"></a>const (
    <a id="L66"></a>running             threadState = &#34;Running&#34;;
    <a id="L67"></a>singleStepping      threadState = &#34;SingleStepping&#34;; <span class="comment">// Transient</span>
    <a id="L68"></a>stopping            threadState = &#34;Stopping&#34;;       <span class="comment">// Transient</span>
    <a id="L69"></a>stopped             threadState = &#34;Stopped&#34;;
    <a id="L70"></a>stoppedBreakpoint   threadState = &#34;StoppedBreakpoint&#34;;
    <a id="L71"></a>stoppedSignal       threadState = &#34;StoppedSignal&#34;;
    <a id="L72"></a>stoppedThreadCreate threadState = &#34;StoppedThreadCreate&#34;;
    <a id="L73"></a>stoppedExiting      threadState = &#34;StoppedExiting&#34;;
    <a id="L74"></a>exiting             threadState = &#34;Exiting&#34;; <span class="comment">// Transient (except main thread)</span>
    <a id="L75"></a>exited              threadState = &#34;Exited&#34;;
    <a id="L76"></a>detached            threadState = &#34;Detached&#34;;
<a id="L77"></a>)

<a id="L79"></a>func (ts threadState) isRunning() bool {
    <a id="L80"></a>return ts == running || ts == singleStepping || ts == stopping
<a id="L81"></a>}

<a id="L83"></a>func (ts threadState) isStopped() bool {
    <a id="L84"></a>return ts == stopped || ts == stoppedBreakpoint || ts == stoppedSignal || ts == stoppedThreadCreate || ts == stoppedExiting
<a id="L85"></a>}

<a id="L87"></a>func (ts threadState) isZombie() bool { return ts == exiting }

<a id="L89"></a>func (ts threadState) isTerminal() bool { return ts == exited || ts == detached }

<a id="L91"></a>func (ts threadState) String() string { return string(ts) }

<a id="L93"></a><span class="comment">/*</span>
<a id="L94"></a><span class="comment"> * Basic types</span>
<a id="L95"></a><span class="comment"> */</span>

<a id="L97"></a><span class="comment">// A breakpoint stores information about a single breakpoint,</span>
<a id="L98"></a><span class="comment">// including its program counter, the overwritten text if the</span>
<a id="L99"></a><span class="comment">// breakpoint is installed.</span>
<a id="L100"></a>type breakpoint struct {
    <a id="L101"></a>pc      uintptr;
    <a id="L102"></a>olddata []byte;
<a id="L103"></a>}

<a id="L105"></a>func (bp *breakpoint) String() string {
    <a id="L106"></a>if bp == nil {
        <a id="L107"></a>return &#34;&lt;nil&gt;&#34;
    <a id="L108"></a>}
    <a id="L109"></a>return fmt.Sprintf(&#34;%#x&#34;, bp.pc);
<a id="L110"></a>}

<a id="L112"></a><span class="comment">// bpinst386 is the breakpoint instruction used on 386 and amd64.</span>
<a id="L113"></a>var bpinst386 = []byte{0xcc}

<a id="L115"></a><span class="comment">// A debugEvent represents a reason a thread stopped or a wait error.</span>
<a id="L116"></a>type debugEvent struct {
    <a id="L117"></a>*os.Waitmsg;
    <a id="L118"></a>t   *thread;
    <a id="L119"></a>err os.Error;
<a id="L120"></a>}

<a id="L122"></a><span class="comment">// A debugReq is a request to execute a closure in the monitor thread.</span>
<a id="L123"></a>type debugReq struct {
    <a id="L124"></a>f   func() os.Error;
    <a id="L125"></a>res chan os.Error;
<a id="L126"></a>}

<a id="L128"></a><span class="comment">// A transitionHandler specifies a function to be called when a thread</span>
<a id="L129"></a><span class="comment">// changes state and a function to be called when an error occurs in</span>
<a id="L130"></a><span class="comment">// the monitor.  Both run in the monitor thread.  Before the monitor</span>
<a id="L131"></a><span class="comment">// invokes a handler, it removes the handler from the handler queue.</span>
<a id="L132"></a><span class="comment">// The handler should re-add itself if needed.</span>
<a id="L133"></a>type transitionHandler struct {
    <a id="L134"></a>handle func(*thread, threadState, threadState);
    <a id="L135"></a>onErr  func(os.Error);
<a id="L136"></a>}

<a id="L138"></a><span class="comment">// A process is a Linux process, which consists of a set of threads.</span>
<a id="L139"></a><span class="comment">// Each running process has one monitor thread, which processes</span>
<a id="L140"></a><span class="comment">// messages from the debugEvents, debugReqs, and stopReq channels and</span>
<a id="L141"></a><span class="comment">// calls transition handlers.</span>
<a id="L142"></a><span class="comment">//</span>
<a id="L143"></a><span class="comment">// To send a message to the monitor thread, first receive from the</span>
<a id="L144"></a><span class="comment">// ready channel.  If the ready channel returns true, the monitor is</span>
<a id="L145"></a><span class="comment">// still running and will accept a message.  If the ready channel</span>
<a id="L146"></a><span class="comment">// returns false, the monitor is not running (the ready channel has</span>
<a id="L147"></a><span class="comment">// been closed), and the reason it is not running will be stored in err.</span>
<a id="L148"></a>type process struct {
    <a id="L149"></a>pid                int;
    <a id="L150"></a>threads            map[int]*thread;
    <a id="L151"></a>breakpoints        map[uintptr]*breakpoint;
    <a id="L152"></a>ready              chan bool;
    <a id="L153"></a>debugEvents        chan *debugEvent;
    <a id="L154"></a>debugReqs          chan *debugReq;
    <a id="L155"></a>stopReq            chan os.Error;
    <a id="L156"></a>transitionHandlers *vector.Vector;
    <a id="L157"></a>err                os.Error;
<a id="L158"></a>}

<a id="L160"></a><span class="comment">// A thread represents a Linux thread in another process that is being</span>
<a id="L161"></a><span class="comment">// debugged.  Each running thread has an associated goroutine that</span>
<a id="L162"></a><span class="comment">// waits for thread updates and sends them to the process monitor.</span>
<a id="L163"></a>type thread struct {
    <a id="L164"></a>tid  int;
    <a id="L165"></a>proc *process;
    <a id="L166"></a><span class="comment">// Whether to ignore the next SIGSTOP received by wait.</span>
    <a id="L167"></a>ignoreNextSigstop bool;

    <a id="L169"></a><span class="comment">// Thread state.  Only modified via setState.</span>
    <a id="L170"></a>state threadState;
    <a id="L171"></a><span class="comment">// If state == StoppedBreakpoint</span>
    <a id="L172"></a>breakpoint *breakpoint;
    <a id="L173"></a><span class="comment">// If state == StoppedSignal or state == Exited</span>
    <a id="L174"></a>signal int;
    <a id="L175"></a><span class="comment">// If state == StoppedThreadCreate</span>
    <a id="L176"></a>newThread *thread;
    <a id="L177"></a><span class="comment">// If state == Exited</span>
    <a id="L178"></a>exitStatus int;
<a id="L179"></a>}

<a id="L181"></a><span class="comment">/*</span>
<a id="L182"></a><span class="comment"> * Errors</span>
<a id="L183"></a><span class="comment"> */</span>

<a id="L185"></a>type badState struct {
    <a id="L186"></a>thread  *thread;
    <a id="L187"></a>message string;
    <a id="L188"></a>state   threadState;
<a id="L189"></a>}

<a id="L191"></a>func (e *badState) String() string {
    <a id="L192"></a>return fmt.Sprintf(&#34;Thread %d %s from state %v&#34;, e.thread.tid, e.message, e.state)
<a id="L193"></a>}

<a id="L195"></a>type breakpointExistsError Word

<a id="L197"></a>func (e breakpointExistsError) String() string {
    <a id="L198"></a>return fmt.Sprintf(&#34;breakpoint already exists at PC %#x&#34;, e)
<a id="L199"></a>}

<a id="L201"></a>type noBreakpointError Word

<a id="L203"></a>func (e noBreakpointError) String() string { return fmt.Sprintf(&#34;no breakpoint at PC %#x&#34;, e) }

<a id="L205"></a>type newThreadError struct {
    <a id="L206"></a>*os.Waitmsg;
    <a id="L207"></a>wantPid int;
    <a id="L208"></a>wantSig int;
<a id="L209"></a>}

<a id="L211"></a>func (e *newThreadError) String() string {
    <a id="L212"></a>return fmt.Sprintf(&#34;newThread wait wanted pid %v and signal %v, got %v and %v&#34;, e.Pid, e.StopSignal(), e.wantPid, e.wantSig)
<a id="L213"></a>}

<a id="L215"></a>type ProcessExited struct{}

<a id="L217"></a>func (p ProcessExited) String() string { return &#34;process exited&#34; }

<a id="L219"></a><span class="comment">/*</span>
<a id="L220"></a><span class="comment"> * Ptrace wrappers</span>
<a id="L221"></a><span class="comment"> */</span>

<a id="L223"></a>func (t *thread) ptracePeekText(addr uintptr, out []byte) (int, os.Error) {
    <a id="L224"></a>c, err := syscall.PtracePeekText(t.tid, addr, out);
    <a id="L225"></a>if traceMem {
        <a id="L226"></a>fmt.Printf(&#34;peek(%#x) =&gt; %v, %v\n&#34;, addr, out, err)
    <a id="L227"></a>}
    <a id="L228"></a>return c, os.NewSyscallError(&#34;ptrace(PEEKTEXT)&#34;, err);
<a id="L229"></a>}

<a id="L231"></a>func (t *thread) ptracePokeText(addr uintptr, out []byte) (int, os.Error) {
    <a id="L232"></a>c, err := syscall.PtracePokeText(t.tid, addr, out);
    <a id="L233"></a>if traceMem {
        <a id="L234"></a>fmt.Printf(&#34;poke(%#x, %v) =&gt; %v\n&#34;, addr, out, err)
    <a id="L235"></a>}
    <a id="L236"></a>return c, os.NewSyscallError(&#34;ptrace(POKETEXT)&#34;, err);
<a id="L237"></a>}

<a id="L239"></a>func (t *thread) ptraceGetRegs(regs *syscall.PtraceRegs) os.Error {
    <a id="L240"></a>err := syscall.PtraceGetRegs(t.tid, regs);
    <a id="L241"></a>return os.NewSyscallError(&#34;ptrace(GETREGS)&#34;, err);
<a id="L242"></a>}

<a id="L244"></a>func (t *thread) ptraceSetRegs(regs *syscall.PtraceRegs) os.Error {
    <a id="L245"></a>err := syscall.PtraceSetRegs(t.tid, regs);
    <a id="L246"></a>return os.NewSyscallError(&#34;ptrace(SETREGS)&#34;, err);
<a id="L247"></a>}

<a id="L249"></a>func (t *thread) ptraceSetOptions(options int) os.Error {
    <a id="L250"></a>err := syscall.PtraceSetOptions(t.tid, options);
    <a id="L251"></a>return os.NewSyscallError(&#34;ptrace(SETOPTIONS)&#34;, err);
<a id="L252"></a>}

<a id="L254"></a>func (t *thread) ptraceGetEventMsg() (uint, os.Error) {
    <a id="L255"></a>msg, err := syscall.PtraceGetEventMsg(t.tid);
    <a id="L256"></a>return msg, os.NewSyscallError(&#34;ptrace(GETEVENTMSG)&#34;, err);
<a id="L257"></a>}

<a id="L259"></a>func (t *thread) ptraceCont() os.Error {
    <a id="L260"></a>err := syscall.PtraceCont(t.tid, 0);
    <a id="L261"></a>return os.NewSyscallError(&#34;ptrace(CONT)&#34;, err);
<a id="L262"></a>}

<a id="L264"></a>func (t *thread) ptraceContWithSignal(sig int) os.Error {
    <a id="L265"></a>err := syscall.PtraceCont(t.tid, sig);
    <a id="L266"></a>return os.NewSyscallError(&#34;ptrace(CONT)&#34;, err);
<a id="L267"></a>}

<a id="L269"></a>func (t *thread) ptraceStep() os.Error {
    <a id="L270"></a>err := syscall.PtraceSingleStep(t.tid);
    <a id="L271"></a>return os.NewSyscallError(&#34;ptrace(SINGLESTEP)&#34;, err);
<a id="L272"></a>}

<a id="L274"></a>func (t *thread) ptraceDetach() os.Error {
    <a id="L275"></a>err := syscall.PtraceDetach(t.tid);
    <a id="L276"></a>return os.NewSyscallError(&#34;ptrace(DETACH)&#34;, err);
<a id="L277"></a>}

<a id="L279"></a><span class="comment">/*</span>
<a id="L280"></a><span class="comment"> * Logging utilties</span>
<a id="L281"></a><span class="comment"> */</span>

<a id="L283"></a>var logLock sync.Mutex

<a id="L285"></a>func (t *thread) logTrace(format string, args ...) {
    <a id="L286"></a>if !trace {
        <a id="L287"></a>return
    <a id="L288"></a>}
    <a id="L289"></a>logLock.Lock();
    <a id="L290"></a>defer logLock.Unlock();
    <a id="L291"></a>fmt.Fprintf(os.Stderr, &#34;Thread %d&#34;, t.tid);
    <a id="L292"></a>if traceIP {
        <a id="L293"></a>var regs syscall.PtraceRegs;
        <a id="L294"></a>err := t.ptraceGetRegs(&amp;regs);
        <a id="L295"></a>if err == nil {
            <a id="L296"></a>fmt.Fprintf(os.Stderr, &#34;@%x&#34;, regs.PC())
        <a id="L297"></a>}
    <a id="L298"></a>}
    <a id="L299"></a>fmt.Fprint(os.Stderr, &#34;: &#34;);
    <a id="L300"></a>fmt.Fprintf(os.Stderr, format, args);
    <a id="L301"></a>fmt.Fprint(os.Stderr, &#34;\n&#34;);
<a id="L302"></a>}

<a id="L304"></a>func (t *thread) warn(format string, args ...) {
    <a id="L305"></a>logLock.Lock();
    <a id="L306"></a>defer logLock.Unlock();
    <a id="L307"></a>fmt.Fprintf(os.Stderr, &#34;Thread %d: WARNING &#34;, t.tid);
    <a id="L308"></a>fmt.Fprintf(os.Stderr, format, args);
    <a id="L309"></a>fmt.Fprint(os.Stderr, &#34;\n&#34;);
<a id="L310"></a>}

<a id="L312"></a>func (p *process) logTrace(format string, args ...) {
    <a id="L313"></a>if !trace {
        <a id="L314"></a>return
    <a id="L315"></a>}
    <a id="L316"></a>logLock.Lock();
    <a id="L317"></a>defer logLock.Unlock();
    <a id="L318"></a>fmt.Fprintf(os.Stderr, &#34;Process %d: &#34;, p.pid);
    <a id="L319"></a>fmt.Fprintf(os.Stderr, format, args);
    <a id="L320"></a>fmt.Fprint(os.Stderr, &#34;\n&#34;);
<a id="L321"></a>}

<a id="L323"></a><span class="comment">/*</span>
<a id="L324"></a><span class="comment"> * State utilities</span>
<a id="L325"></a><span class="comment"> */</span>

<a id="L327"></a><span class="comment">// someStoppedThread returns a stopped thread from the process.</span>
<a id="L328"></a><span class="comment">// Returns nil if no threads are stopped.</span>
<a id="L329"></a><span class="comment">//</span>
<a id="L330"></a><span class="comment">// Must be called from the monitor thread.</span>
<a id="L331"></a>func (p *process) someStoppedThread() *thread {
    <a id="L332"></a>for _, t := range p.threads {
        <a id="L333"></a>if t.state.isStopped() {
            <a id="L334"></a>return t
        <a id="L335"></a>}
    <a id="L336"></a>}
    <a id="L337"></a>return nil;
<a id="L338"></a>}

<a id="L340"></a><span class="comment">// someRunningThread returns a running thread from the process.</span>
<a id="L341"></a><span class="comment">// Returns nil if no threads are running.</span>
<a id="L342"></a><span class="comment">//</span>
<a id="L343"></a><span class="comment">// Must be called from the monitor thread.</span>
<a id="L344"></a>func (p *process) someRunningThread() *thread {
    <a id="L345"></a>for _, t := range p.threads {
        <a id="L346"></a>if t.state.isRunning() {
            <a id="L347"></a>return t
        <a id="L348"></a>}
    <a id="L349"></a>}
    <a id="L350"></a>return nil;
<a id="L351"></a>}

<a id="L353"></a><span class="comment">/*</span>
<a id="L354"></a><span class="comment"> * Breakpoint utilities</span>
<a id="L355"></a><span class="comment"> */</span>

<a id="L357"></a><span class="comment">// installBreakpoints adds breakpoints to the attached process.</span>
<a id="L358"></a><span class="comment">//</span>
<a id="L359"></a><span class="comment">// Must be called from the monitor thread.</span>
<a id="L360"></a>func (p *process) installBreakpoints() os.Error {
    <a id="L361"></a>n := 0;
    <a id="L362"></a>main := p.someStoppedThread();
    <a id="L363"></a>for _, b := range p.breakpoints {
        <a id="L364"></a>if b.olddata != nil {
            <a id="L365"></a>continue
        <a id="L366"></a>}

        <a id="L368"></a>b.olddata = make([]byte, len(bpinst386));
        <a id="L369"></a>_, err := main.ptracePeekText(uintptr(b.pc), b.olddata);
        <a id="L370"></a>if err != nil {
            <a id="L371"></a>b.olddata = nil;
            <a id="L372"></a>return err;
        <a id="L373"></a>}

        <a id="L375"></a>_, err = main.ptracePokeText(uintptr(b.pc), bpinst386);
        <a id="L376"></a>if err != nil {
            <a id="L377"></a>b.olddata = nil;
            <a id="L378"></a>return err;
        <a id="L379"></a>}
        <a id="L380"></a>n++;
    <a id="L381"></a>}
    <a id="L382"></a>if n &gt; 0 {
        <a id="L383"></a>p.logTrace(&#34;installed %d/%d breakpoints&#34;, n, len(p.breakpoints))
    <a id="L384"></a>}

    <a id="L386"></a>return nil;
<a id="L387"></a>}

<a id="L389"></a><span class="comment">// uninstallBreakpoints removes the installed breakpoints from p.</span>
<a id="L390"></a><span class="comment">//</span>
<a id="L391"></a><span class="comment">// Must be called from the monitor thread.</span>
<a id="L392"></a>func (p *process) uninstallBreakpoints() os.Error {
    <a id="L393"></a>if len(p.threads) == 0 {
        <a id="L394"></a>return nil
    <a id="L395"></a>}
    <a id="L396"></a>n := 0;
    <a id="L397"></a>main := p.someStoppedThread();
    <a id="L398"></a>for _, b := range p.breakpoints {
        <a id="L399"></a>if b.olddata == nil {
            <a id="L400"></a>continue
        <a id="L401"></a>}

        <a id="L403"></a>_, err := main.ptracePokeText(uintptr(b.pc), b.olddata);
        <a id="L404"></a>if err != nil {
            <a id="L405"></a>return err
        <a id="L406"></a>}
        <a id="L407"></a>b.olddata = nil;
        <a id="L408"></a>n++;
    <a id="L409"></a>}
    <a id="L410"></a>if n &gt; 0 {
        <a id="L411"></a>p.logTrace(&#34;uninstalled %d/%d breakpoints&#34;, n, len(p.breakpoints))
    <a id="L412"></a>}

    <a id="L414"></a>return nil;
<a id="L415"></a>}

<a id="L417"></a><span class="comment">/*</span>
<a id="L418"></a><span class="comment"> * Debug event handling</span>
<a id="L419"></a><span class="comment"> */</span>

<a id="L421"></a><span class="comment">// wait waits for a wait event from this thread and sends it on the</span>
<a id="L422"></a><span class="comment">// debug events channel for this thread&#39;s process.  This should be</span>
<a id="L423"></a><span class="comment">// started in its own goroutine when the attached thread enters a</span>
<a id="L424"></a><span class="comment">// running state.  The goroutine will exit as soon as it sends a debug</span>
<a id="L425"></a><span class="comment">// event.</span>
<a id="L426"></a>func (t *thread) wait() {
    <a id="L427"></a>for {
        <a id="L428"></a>var ev debugEvent;
        <a id="L429"></a>ev.t = t;
        <a id="L430"></a>t.logTrace(&#34;beginning wait&#34;);
        <a id="L431"></a>ev.Waitmsg, ev.err = os.Wait(t.tid, syscall.WALL);
        <a id="L432"></a>if ev.err == nil &amp;&amp; ev.Pid != t.tid {
            <a id="L433"></a>panic(&#34;Wait returned pid &#34;, ev.Pid, &#34; wanted &#34;, t.tid)
        <a id="L434"></a>}
        <a id="L435"></a>if ev.StopSignal() == syscall.SIGSTOP &amp;&amp; t.ignoreNextSigstop {
            <a id="L436"></a><span class="comment">// Spurious SIGSTOP.  See Thread.Stop().</span>
            <a id="L437"></a>t.ignoreNextSigstop = false;
            <a id="L438"></a>err := t.ptraceCont();
            <a id="L439"></a>if err == nil {
                <a id="L440"></a>continue
            <a id="L441"></a>}
            <a id="L442"></a><span class="comment">// If we failed to continue, just let</span>
            <a id="L443"></a><span class="comment">// the stop go through so we can</span>
            <a id="L444"></a><span class="comment">// update the thread&#39;s state.</span>
        <a id="L445"></a>}
        <a id="L446"></a>if !&lt;-t.proc.ready {
            <a id="L447"></a><span class="comment">// The monitor exited</span>
            <a id="L448"></a>break
        <a id="L449"></a>}
        <a id="L450"></a>t.proc.debugEvents &lt;- &amp;ev;
        <a id="L451"></a>break;
    <a id="L452"></a>}
<a id="L453"></a>}

<a id="L455"></a><span class="comment">// setState sets this thread&#39;s state, starts a wait thread if</span>
<a id="L456"></a><span class="comment">// necessary, and invokes state transition handlers.</span>
<a id="L457"></a><span class="comment">//</span>
<a id="L458"></a><span class="comment">// Must be called from the monitor thread.</span>
<a id="L459"></a>func (t *thread) setState(new threadState) {
    <a id="L460"></a>old := t.state;
    <a id="L461"></a>t.state = new;
    <a id="L462"></a>t.logTrace(&#34;state %v -&gt; %v&#34;, old, new);

    <a id="L464"></a>if !old.isRunning() &amp;&amp; (new.isRunning() || new.isZombie()) {
        <a id="L465"></a><span class="comment">// Start waiting on this thread</span>
        <a id="L466"></a>go t.wait()
    <a id="L467"></a>}

    <a id="L469"></a><span class="comment">// Invoke state change handlers</span>
    <a id="L470"></a>handlers := t.proc.transitionHandlers;
    <a id="L471"></a>if handlers.Len() == 0 {
        <a id="L472"></a>return
    <a id="L473"></a>}

    <a id="L475"></a>t.proc.transitionHandlers = vector.New(0);
    <a id="L476"></a>for _, h := range handlers.Data() {
        <a id="L477"></a>h := h.(*transitionHandler);
        <a id="L478"></a>h.handle(t, old, new);
    <a id="L479"></a>}
<a id="L480"></a>}

<a id="L482"></a><span class="comment">// sendSigstop sends a SIGSTOP to this thread.</span>
<a id="L483"></a>func (t *thread) sendSigstop() os.Error {
    <a id="L484"></a>t.logTrace(&#34;sending SIGSTOP&#34;);
    <a id="L485"></a>err := syscall.Tgkill(t.proc.pid, t.tid, syscall.SIGSTOP);
    <a id="L486"></a>return os.NewSyscallError(&#34;tgkill&#34;, err);
<a id="L487"></a>}

<a id="L489"></a><span class="comment">// stopAsync sends SIGSTOP to all threads in state &#39;running&#39;.</span>
<a id="L490"></a><span class="comment">//</span>
<a id="L491"></a><span class="comment">// Must be called from the monitor thread.</span>
<a id="L492"></a>func (p *process) stopAsync() os.Error {
    <a id="L493"></a>for _, t := range p.threads {
        <a id="L494"></a>if t.state == running {
            <a id="L495"></a>err := t.sendSigstop();
            <a id="L496"></a>if err != nil {
                <a id="L497"></a>return err
            <a id="L498"></a>}
            <a id="L499"></a>t.setState(stopping);
        <a id="L500"></a>}
    <a id="L501"></a>}
    <a id="L502"></a>return nil;
<a id="L503"></a>}

<a id="L505"></a><span class="comment">// doTrap handles SIGTRAP debug events with a cause of 0.  These can</span>
<a id="L506"></a><span class="comment">// be caused either by an installed breakpoint, a breakpoint in the</span>
<a id="L507"></a><span class="comment">// program text, or by single stepping.</span>
<a id="L508"></a><span class="comment">//</span>
<a id="L509"></a><span class="comment">// TODO(austin) I think we also get this on an execve syscall.</span>
<a id="L510"></a>func (ev *debugEvent) doTrap() (threadState, os.Error) {
    <a id="L511"></a>t := ev.t;

    <a id="L513"></a>if t.state == singleStepping {
        <a id="L514"></a>return stopped, nil
    <a id="L515"></a>}

    <a id="L517"></a><span class="comment">// Hit a breakpoint.  Linux leaves the program counter after</span>
    <a id="L518"></a><span class="comment">// the breakpoint.  If this is an installed breakpoint, we</span>
    <a id="L519"></a><span class="comment">// need to back the PC up to the breakpoint PC.</span>
    <a id="L520"></a>var regs syscall.PtraceRegs;
    <a id="L521"></a>err := t.ptraceGetRegs(&amp;regs);
    <a id="L522"></a>if err != nil {
        <a id="L523"></a>return stopped, err
    <a id="L524"></a>}

    <a id="L526"></a>b, ok := t.proc.breakpoints[uintptr(regs.PC())-uintptr(len(bpinst386))];
    <a id="L527"></a>if !ok {
        <a id="L528"></a><span class="comment">// We must have hit a breakpoint that was actually in</span>
        <a id="L529"></a><span class="comment">// the program.  Leave the IP where it is so we don&#39;t</span>
        <a id="L530"></a><span class="comment">// re-execute the breakpoint instruction.  Expose the</span>
        <a id="L531"></a><span class="comment">// fact that we stopped with a SIGTRAP.</span>
        <a id="L532"></a>return stoppedSignal, nil
    <a id="L533"></a>}

    <a id="L535"></a>t.breakpoint = b;
    <a id="L536"></a>t.logTrace(&#34;at breakpoint %v, backing up PC from %#x&#34;, b, regs.PC());

    <a id="L538"></a>regs.SetPC(uint64(b.pc));
    <a id="L539"></a>err = t.ptraceSetRegs(&amp;regs);
    <a id="L540"></a>if err != nil {
        <a id="L541"></a>return stopped, err
    <a id="L542"></a>}
    <a id="L543"></a>return stoppedBreakpoint, nil;
<a id="L544"></a>}

<a id="L546"></a><span class="comment">// doPtraceClone handles SIGTRAP debug events with a PTRACE_EVENT_CLONE</span>
<a id="L547"></a><span class="comment">// cause.  It initializes the new thread, adds it to the process, and</span>
<a id="L548"></a><span class="comment">// returns the appropriate thread state for the existing thread.</span>
<a id="L549"></a>func (ev *debugEvent) doPtraceClone() (threadState, os.Error) {
    <a id="L550"></a>t := ev.t;

    <a id="L552"></a><span class="comment">// Get the TID of the new thread</span>
    <a id="L553"></a>tid, err := t.ptraceGetEventMsg();
    <a id="L554"></a>if err != nil {
        <a id="L555"></a>return stopped, err
    <a id="L556"></a>}

    <a id="L558"></a>nt, err := t.proc.newThread(int(tid), syscall.SIGSTOP, true);
    <a id="L559"></a>if err != nil {
        <a id="L560"></a>return stopped, err
    <a id="L561"></a>}

    <a id="L563"></a><span class="comment">// Remember the thread</span>
    <a id="L564"></a>t.newThread = nt;

    <a id="L566"></a>return stoppedThreadCreate, nil;
<a id="L567"></a>}

<a id="L569"></a><span class="comment">// doPtraceExit handles SIGTRAP debug events with a PTRACE_EVENT_EXIT</span>
<a id="L570"></a><span class="comment">// cause.  It sets up the thread&#39;s state, but does not remove it from</span>
<a id="L571"></a><span class="comment">// the process.  A later WIFEXITED debug event will remove it from the</span>
<a id="L572"></a><span class="comment">// process.</span>
<a id="L573"></a>func (ev *debugEvent) doPtraceExit() (threadState, os.Error) {
    <a id="L574"></a>t := ev.t;

    <a id="L576"></a><span class="comment">// Get exit status</span>
    <a id="L577"></a>exitStatus, err := t.ptraceGetEventMsg();
    <a id="L578"></a>if err != nil {
        <a id="L579"></a>return stopped, err
    <a id="L580"></a>}
    <a id="L581"></a>ws := syscall.WaitStatus(exitStatus);
    <a id="L582"></a>t.logTrace(&#34;exited with %v&#34;, ws);
    <a id="L583"></a>switch {
    <a id="L584"></a>case ws.Exited():
        <a id="L585"></a>t.exitStatus = ws.ExitStatus()
    <a id="L586"></a>case ws.Signaled():
        <a id="L587"></a>t.signal = ws.Signal()
    <a id="L588"></a>}

    <a id="L590"></a><span class="comment">// We still need to continue this thread and wait on this</span>
    <a id="L591"></a><span class="comment">// thread&#39;s WIFEXITED event.  We&#39;ll delete it then.</span>
    <a id="L592"></a>return stoppedExiting, nil;
<a id="L593"></a>}

<a id="L595"></a><span class="comment">// process handles a debug event.  It modifies any thread or process</span>
<a id="L596"></a><span class="comment">// state as necessary, uninstalls breakpoints if necessary, and stops</span>
<a id="L597"></a><span class="comment">// any running threads.</span>
<a id="L598"></a>func (ev *debugEvent) process() os.Error {
    <a id="L599"></a>if ev.err != nil {
        <a id="L600"></a>return ev.err
    <a id="L601"></a>}

    <a id="L603"></a>t := ev.t;
    <a id="L604"></a>t.exitStatus = -1;
    <a id="L605"></a>t.signal = -1;

    <a id="L607"></a><span class="comment">// Decode wait status.</span>
    <a id="L608"></a>var state threadState;
    <a id="L609"></a>switch {
    <a id="L610"></a>case ev.Stopped():
        <a id="L611"></a>state = stoppedSignal;
        <a id="L612"></a>t.signal = ev.StopSignal();
        <a id="L613"></a>t.logTrace(&#34;stopped with %v&#34;, ev);
        <a id="L614"></a>if ev.StopSignal() == syscall.SIGTRAP {
            <a id="L615"></a><span class="comment">// What caused the debug trap?</span>
            <a id="L616"></a>var err os.Error;
            <a id="L617"></a>switch cause := ev.TrapCause(); cause {
            <a id="L618"></a>case 0:
                <a id="L619"></a><span class="comment">// Breakpoint or single stepping</span>
                <a id="L620"></a>state, err = ev.doTrap()

            <a id="L622"></a>case syscall.PTRACE_EVENT_CLONE:
                <a id="L623"></a>state, err = ev.doPtraceClone()

            <a id="L625"></a>case syscall.PTRACE_EVENT_EXIT:
                <a id="L626"></a>state, err = ev.doPtraceExit()

            <a id="L628"></a>default:
                <a id="L629"></a>t.warn(&#34;Unknown trap cause %d&#34;, cause)
            <a id="L630"></a>}

            <a id="L632"></a>if err != nil {
                <a id="L633"></a>t.setState(stopped);
                <a id="L634"></a>t.warn(&#34;failed to handle trap %v: %v&#34;, ev, err);
            <a id="L635"></a>}
        <a id="L636"></a>}

    <a id="L638"></a>case ev.Exited():
        <a id="L639"></a>state = exited;
        <a id="L640"></a>t.proc.threads[t.tid] = nil, false;
        <a id="L641"></a>t.logTrace(&#34;exited %v&#34;, ev);
        <a id="L642"></a><span class="comment">// We should have gotten the exit status in</span>
        <a id="L643"></a><span class="comment">// PTRACE_EVENT_EXIT, but just in case.</span>
        <a id="L644"></a>t.exitStatus = ev.ExitStatus();

    <a id="L646"></a>case ev.Signaled():
        <a id="L647"></a>state = exited;
        <a id="L648"></a>t.proc.threads[t.tid] = nil, false;
        <a id="L649"></a>t.logTrace(&#34;signaled %v&#34;, ev);
        <a id="L650"></a><span class="comment">// Again, this should be redundant.</span>
        <a id="L651"></a>t.signal = ev.Signal();

    <a id="L653"></a>default:
        <a id="L654"></a>panic(fmt.Sprintf(&#34;Unexpected wait status %v&#34;, ev.Waitmsg))
    <a id="L655"></a>}

    <a id="L657"></a><span class="comment">// If we sent a SIGSTOP to the thread (indicated by state</span>
    <a id="L658"></a><span class="comment">// Stopping), we might have raced with a different type of</span>
    <a id="L659"></a><span class="comment">// stop.  If we didn&#39;t get the stop we expected, then the</span>
    <a id="L660"></a><span class="comment">// SIGSTOP we sent is now queued up, so we should ignore the</span>
    <a id="L661"></a><span class="comment">// next one we get.</span>
    <a id="L662"></a>if t.state == stopping &amp;&amp; ev.StopSignal() != syscall.SIGSTOP {
        <a id="L663"></a>t.ignoreNextSigstop = true
    <a id="L664"></a>}

    <a id="L666"></a><span class="comment">// TODO(austin) If we&#39;re in state stopping and get a SIGSTOP,</span>
    <a id="L667"></a><span class="comment">// set state stopped instead of stoppedSignal.</span>

    <a id="L669"></a>t.setState(state);

    <a id="L671"></a>if t.proc.someRunningThread() == nil {
        <a id="L672"></a><span class="comment">// Nothing is running, uninstall breakpoints</span>
        <a id="L673"></a>return t.proc.uninstallBreakpoints()
    <a id="L674"></a>}
    <a id="L675"></a><span class="comment">// Stop any other running threads</span>
    <a id="L676"></a>return t.proc.stopAsync();
<a id="L677"></a>}

<a id="L679"></a><span class="comment">// onStop adds a handler for state transitions from running to</span>
<a id="L680"></a><span class="comment">// non-running states.  The handler will be called from the monitor</span>
<a id="L681"></a><span class="comment">// thread.</span>
<a id="L682"></a><span class="comment">//</span>
<a id="L683"></a><span class="comment">// Must be called from the monitor thread.</span>
<a id="L684"></a>func (t *thread) onStop(handle func(), onErr func(os.Error)) {
    <a id="L685"></a><span class="comment">// TODO(austin) This is rather inefficient for things like</span>
    <a id="L686"></a><span class="comment">// stepping all threads during a continue.  Maybe move</span>
    <a id="L687"></a><span class="comment">// transitionHandlers to the thread, or have both per-thread</span>
    <a id="L688"></a><span class="comment">// and per-process transition handlers.</span>
    <a id="L689"></a>h := &amp;transitionHandler{nil, onErr};
    <a id="L690"></a>h.handle = func(st *thread, old, new threadState) {
        <a id="L691"></a>if t == st &amp;&amp; old.isRunning() &amp;&amp; !new.isRunning() {
            <a id="L692"></a>handle()
        <a id="L693"></a>} else {
            <a id="L694"></a>t.proc.transitionHandlers.Push(h)
        <a id="L695"></a>}
    <a id="L696"></a>};
    <a id="L697"></a>t.proc.transitionHandlers.Push(h);
<a id="L698"></a>}

<a id="L700"></a><span class="comment">/*</span>
<a id="L701"></a><span class="comment"> * Event monitor</span>
<a id="L702"></a><span class="comment"> */</span>

<a id="L704"></a><span class="comment">// monitor handles debug events and debug requests for p, exiting when</span>
<a id="L705"></a><span class="comment">// there are no threads left in p.</span>
<a id="L706"></a>func (p *process) monitor() {
    <a id="L707"></a>var err os.Error;

    <a id="L709"></a><span class="comment">// Linux requires that all ptrace calls come from the thread</span>
    <a id="L710"></a><span class="comment">// that originally attached.  Prevent the Go scheduler from</span>
    <a id="L711"></a><span class="comment">// migrating us to other OS threads.</span>
    <a id="L712"></a>runtime.LockOSThread();
    <a id="L713"></a>defer runtime.UnlockOSThread();

    <a id="L715"></a>hadThreads := false;
    <a id="L716"></a>for err == nil {
        <a id="L717"></a>p.ready &lt;- true;
        <a id="L718"></a>select {
        <a id="L719"></a>case event := &lt;-p.debugEvents:
            <a id="L720"></a>err = event.process()

        <a id="L722"></a>case req := &lt;-p.debugReqs:
            <a id="L723"></a>req.res &lt;- req.f()

        <a id="L725"></a>case err = &lt;-p.stopReq:
            <a id="L726"></a>break
        <a id="L727"></a>}

        <a id="L729"></a>if len(p.threads) == 0 {
            <a id="L730"></a>if err == nil &amp;&amp; hadThreads {
                <a id="L731"></a>p.logTrace(&#34;no more threads; monitor exiting&#34;);
                <a id="L732"></a>err = ProcessExited{};
            <a id="L733"></a>}
        <a id="L734"></a>} else {
            <a id="L735"></a>hadThreads = true
        <a id="L736"></a>}
    <a id="L737"></a>}

    <a id="L739"></a><span class="comment">// Abort waiting handlers</span>
    <a id="L740"></a><span class="comment">// TODO(austin) How do I stop the wait threads?</span>
    <a id="L741"></a>for _, h := range p.transitionHandlers.Data() {
        <a id="L742"></a>h := h.(*transitionHandler);
        <a id="L743"></a>h.onErr(err);
    <a id="L744"></a>}

    <a id="L746"></a><span class="comment">// Indicate that the monitor cannot receive any more messages</span>
    <a id="L747"></a>p.err = err;
    <a id="L748"></a>close(p.ready);
<a id="L749"></a>}

<a id="L751"></a><span class="comment">// do executes f in the monitor thread (and, thus, atomically with</span>
<a id="L752"></a><span class="comment">// respect to thread state changes).  f must not block.</span>
<a id="L753"></a><span class="comment">//</span>
<a id="L754"></a><span class="comment">// Must NOT be called from the monitor thread.</span>
<a id="L755"></a>func (p *process) do(f func() os.Error) os.Error {
    <a id="L756"></a>if !&lt;-p.ready {
        <a id="L757"></a>return p.err
    <a id="L758"></a>}
    <a id="L759"></a>req := &amp;debugReq{f, make(chan os.Error)};
    <a id="L760"></a>p.debugReqs &lt;- req;
    <a id="L761"></a>return &lt;-req.res;
<a id="L762"></a>}

<a id="L764"></a><span class="comment">// stopMonitor stops the monitor with the given error.  If the monitor</span>
<a id="L765"></a><span class="comment">// is already stopped, does nothing.</span>
<a id="L766"></a>func (p *process) stopMonitor(err os.Error) {
    <a id="L767"></a>if err == nil {
        <a id="L768"></a>panic(&#34;cannot stop the monitor with no error&#34;)
    <a id="L769"></a>}
    <a id="L770"></a>if &lt;-p.ready {
        <a id="L771"></a>p.stopReq &lt;- err
    <a id="L772"></a>}
<a id="L773"></a>}

<a id="L775"></a><span class="comment">/*</span>
<a id="L776"></a><span class="comment"> * Public thread interface</span>
<a id="L777"></a><span class="comment"> */</span>

<a id="L779"></a>func (t *thread) Regs() (Regs, os.Error) {
    <a id="L780"></a>var regs syscall.PtraceRegs;

    <a id="L782"></a>err := t.proc.do(func() os.Error {
        <a id="L783"></a>if !t.state.isStopped() {
            <a id="L784"></a>return &amp;badState{t, &#34;cannot get registers&#34;, t.state}
        <a id="L785"></a>}
        <a id="L786"></a>return t.ptraceGetRegs(&amp;regs);
    <a id="L787"></a>});
    <a id="L788"></a>if err != nil {
        <a id="L789"></a>return nil, err
    <a id="L790"></a>}

    <a id="L792"></a>setter := func(r *syscall.PtraceRegs) os.Error {
        <a id="L793"></a>return t.proc.do(func() os.Error {
            <a id="L794"></a>if !t.state.isStopped() {
                <a id="L795"></a>return &amp;badState{t, &#34;cannot get registers&#34;, t.state}
            <a id="L796"></a>}
            <a id="L797"></a>return t.ptraceSetRegs(r);
        <a id="L798"></a>})
    <a id="L799"></a>};
    <a id="L800"></a>return newRegs(&amp;regs, setter), nil;
<a id="L801"></a>}

<a id="L803"></a>func (t *thread) Peek(addr Word, out []byte) (int, os.Error) {
    <a id="L804"></a>var c int;

    <a id="L806"></a>err := t.proc.do(func() os.Error {
        <a id="L807"></a>if !t.state.isStopped() {
            <a id="L808"></a>return &amp;badState{t, &#34;cannot peek text&#34;, t.state}
        <a id="L809"></a>}

        <a id="L811"></a>var err os.Error;
        <a id="L812"></a>c, err = t.ptracePeekText(uintptr(addr), out);
        <a id="L813"></a>return err;
    <a id="L814"></a>});

    <a id="L816"></a>return c, err;
<a id="L817"></a>}

<a id="L819"></a>func (t *thread) Poke(addr Word, out []byte) (int, os.Error) {
    <a id="L820"></a>var c int;

    <a id="L822"></a>err := t.proc.do(func() os.Error {
        <a id="L823"></a>if !t.state.isStopped() {
            <a id="L824"></a>return &amp;badState{t, &#34;cannot poke text&#34;, t.state}
        <a id="L825"></a>}

        <a id="L827"></a>var err os.Error;
        <a id="L828"></a>c, err = t.ptracePokeText(uintptr(addr), out);
        <a id="L829"></a>return err;
    <a id="L830"></a>});

    <a id="L832"></a>return c, err;
<a id="L833"></a>}

<a id="L835"></a><span class="comment">// stepAsync starts this thread single stepping.  When the single step</span>
<a id="L836"></a><span class="comment">// is complete, it will send nil on the given channel.  If an error</span>
<a id="L837"></a><span class="comment">// occurs while setting up the single step, it returns that error.  If</span>
<a id="L838"></a><span class="comment">// an error occurs while waiting for the single step to complete, it</span>
<a id="L839"></a><span class="comment">// sends that error on the channel.</span>
<a id="L840"></a>func (t *thread) stepAsync(ready chan os.Error) os.Error {
    <a id="L841"></a>if err := t.ptraceStep(); err != nil {
        <a id="L842"></a>return err
    <a id="L843"></a>}
    <a id="L844"></a>t.setState(singleStepping);
    <a id="L845"></a>t.onStop(func() { ready &lt;- nil },
        <a id="L846"></a>func(err os.Error) { ready &lt;- err });
    <a id="L847"></a>return nil;
<a id="L848"></a>}

<a id="L850"></a>func (t *thread) Step() os.Error {
    <a id="L851"></a>t.logTrace(&#34;Step {&#34;);
    <a id="L852"></a>defer t.logTrace(&#34;}&#34;);

    <a id="L854"></a>ready := make(chan os.Error);

    <a id="L856"></a>err := t.proc.do(func() os.Error {
        <a id="L857"></a>if !t.state.isStopped() {
            <a id="L858"></a>return &amp;badState{t, &#34;cannot single step&#34;, t.state}
        <a id="L859"></a>}
        <a id="L860"></a>return t.stepAsync(ready);
    <a id="L861"></a>});
    <a id="L862"></a>if err != nil {
        <a id="L863"></a>return err
    <a id="L864"></a>}

    <a id="L866"></a>err = &lt;-ready;
    <a id="L867"></a>return err;
<a id="L868"></a>}

<a id="L870"></a><span class="comment">// TODO(austin) We should probably get this via C&#39;s strsignal.</span>
<a id="L871"></a>var sigNames = [...]string{
    <a id="L872"></a>&#34;SIGEXIT&#34;, &#34;SIGHUP&#34;, &#34;SIGINT&#34;, &#34;SIGQUIT&#34;, &#34;SIGILL&#34;,
    <a id="L873"></a>&#34;SIGTRAP&#34;, &#34;SIGABRT&#34;, &#34;SIGBUS&#34;, &#34;SIGFPE&#34;, &#34;SIGKILL&#34;,
    <a id="L874"></a>&#34;SIGUSR1&#34;, &#34;SIGSEGV&#34;, &#34;SIGUSR2&#34;, &#34;SIGPIPE&#34;, &#34;SIGALRM&#34;,
    <a id="L875"></a>&#34;SIGTERM&#34;, &#34;SIGSTKFLT&#34;, &#34;SIGCHLD&#34;, &#34;SIGCONT&#34;, &#34;SIGSTOP&#34;,
    <a id="L876"></a>&#34;SIGTSTP&#34;, &#34;SIGTTIN&#34;, &#34;SIGTTOU&#34;, &#34;SIGURG&#34;, &#34;SIGXCPU&#34;,
    <a id="L877"></a>&#34;SIGXFSZ&#34;, &#34;SIGVTALRM&#34;, &#34;SIGPROF&#34;, &#34;SIGWINCH&#34;, &#34;SIGPOLL&#34;,
    <a id="L878"></a>&#34;SIGPWR&#34;, &#34;SIGSYS&#34;,
<a id="L879"></a>}

<a id="L881"></a><span class="comment">// sigName returns the symbolic name for the given signal number.  If</span>
<a id="L882"></a><span class="comment">// the signal number is invalid, returns &#34;&lt;invalid&gt;&#34;.</span>
<a id="L883"></a>func sigName(signal int) string {
    <a id="L884"></a>if signal &lt; 0 || signal &gt;= len(sigNames) {
        <a id="L885"></a>return &#34;&lt;invalid&gt;&#34;
    <a id="L886"></a>}
    <a id="L887"></a>return sigNames[signal];
<a id="L888"></a>}

<a id="L890"></a>func (t *thread) Stopped() (Cause, os.Error) {
    <a id="L891"></a>var c Cause;
    <a id="L892"></a>err := t.proc.do(func() os.Error {
        <a id="L893"></a>switch t.state {
        <a id="L894"></a>case stopped:
            <a id="L895"></a>c = Stopped{}

        <a id="L897"></a>case stoppedBreakpoint:
            <a id="L898"></a>c = Breakpoint(t.breakpoint.pc)

        <a id="L900"></a>case stoppedSignal:
            <a id="L901"></a>c = Signal(sigName(t.signal))

        <a id="L903"></a>case stoppedThreadCreate:
            <a id="L904"></a>c = &amp;ThreadCreate{t.newThread}

        <a id="L906"></a>case stoppedExiting, exiting, exited:
            <a id="L907"></a>if t.signal == -1 {
                <a id="L908"></a>c = &amp;ThreadExit{t.exitStatus, &#34;&#34;}
            <a id="L909"></a>} else {
                <a id="L910"></a>c = &amp;ThreadExit{t.exitStatus, sigName(t.signal)}
            <a id="L911"></a>}

        <a id="L913"></a>default:
            <a id="L914"></a>return &amp;badState{t, &#34;cannot get stop cause&#34;, t.state}
        <a id="L915"></a>}
        <a id="L916"></a>return nil;
    <a id="L917"></a>});
    <a id="L918"></a>if err != nil {
        <a id="L919"></a>return nil, err
    <a id="L920"></a>}

    <a id="L922"></a>return c, nil;
<a id="L923"></a>}

<a id="L925"></a>func (p *process) Threads() []Thread {
    <a id="L926"></a>var res []Thread;

    <a id="L928"></a>p.do(func() os.Error {
        <a id="L929"></a>res = make([]Thread, len(p.threads));
        <a id="L930"></a>i := 0;
        <a id="L931"></a>for _, t := range p.threads {
            <a id="L932"></a><span class="comment">// Exclude zombie threads.</span>
            <a id="L933"></a>st := t.state;
            <a id="L934"></a>if st == exiting || st == exited || st == detached {
                <a id="L935"></a>continue
            <a id="L936"></a>}

            <a id="L938"></a>res[i] = t;
            <a id="L939"></a>i++;
        <a id="L940"></a>}
        <a id="L941"></a>res = res[0:i];
        <a id="L942"></a>return nil;
    <a id="L943"></a>});
    <a id="L944"></a>return res;
<a id="L945"></a>}

<a id="L947"></a>func (p *process) AddBreakpoint(pc Word) os.Error {
    <a id="L948"></a>return p.do(func() os.Error {
        <a id="L949"></a>if t := p.someRunningThread(); t != nil {
            <a id="L950"></a>return &amp;badState{t, &#34;cannot add breakpoint&#34;, t.state}
        <a id="L951"></a>}
        <a id="L952"></a>if _, ok := p.breakpoints[uintptr(pc)]; ok {
            <a id="L953"></a>return breakpointExistsError(pc)
        <a id="L954"></a>}
        <a id="L955"></a>p.breakpoints[uintptr(pc)] = &amp;breakpoint{pc: uintptr(pc)};
        <a id="L956"></a>return nil;
    <a id="L957"></a>})
<a id="L958"></a>}

<a id="L960"></a>func (p *process) RemoveBreakpoint(pc Word) os.Error {
    <a id="L961"></a>return p.do(func() os.Error {
        <a id="L962"></a>if t := p.someRunningThread(); t != nil {
            <a id="L963"></a>return &amp;badState{t, &#34;cannot remove breakpoint&#34;, t.state}
        <a id="L964"></a>}
        <a id="L965"></a>if _, ok := p.breakpoints[uintptr(pc)]; !ok {
            <a id="L966"></a>return noBreakpointError(pc)
        <a id="L967"></a>}
        <a id="L968"></a>p.breakpoints[uintptr(pc)] = nil, false;
        <a id="L969"></a>return nil;
    <a id="L970"></a>})
<a id="L971"></a>}

<a id="L973"></a>func (p *process) Continue() os.Error {
    <a id="L974"></a><span class="comment">// Single step any threads that are stopped at breakpoints so</span>
    <a id="L975"></a><span class="comment">// we can reinstall breakpoints.</span>
    <a id="L976"></a>var ready chan os.Error;
    <a id="L977"></a>count := 0;

    <a id="L979"></a>err := p.do(func() os.Error {
        <a id="L980"></a><span class="comment">// We make the ready channel big enough to hold all</span>
        <a id="L981"></a><span class="comment">// ready message so we don&#39;t jam up the monitor if we</span>
        <a id="L982"></a><span class="comment">// stop listening (e.g., if there&#39;s an error).</span>
        <a id="L983"></a>ready = make(chan os.Error, len(p.threads));

        <a id="L985"></a>for _, t := range p.threads {
            <a id="L986"></a>if !t.state.isStopped() {
                <a id="L987"></a>continue
            <a id="L988"></a>}

            <a id="L990"></a><span class="comment">// We use the breakpoint map directly here</span>
            <a id="L991"></a><span class="comment">// instead of checking the stop cause because</span>
            <a id="L992"></a><span class="comment">// it could have been stopped at a breakpoint</span>
            <a id="L993"></a><span class="comment">// for some other reason, or the breakpoint</span>
            <a id="L994"></a><span class="comment">// could have been added since it was stopped.</span>
            <a id="L995"></a>var regs syscall.PtraceRegs;
            <a id="L996"></a>err := t.ptraceGetRegs(&amp;regs);
            <a id="L997"></a>if err != nil {
                <a id="L998"></a>return err
            <a id="L999"></a>}
            <a id="L1000"></a>if b, ok := p.breakpoints[uintptr(regs.PC())]; ok {
                <a id="L1001"></a>t.logTrace(&#34;stepping over breakpoint %v&#34;, b);
                <a id="L1002"></a>if err := t.stepAsync(ready); err != nil {
                    <a id="L1003"></a>return err
                <a id="L1004"></a>}
                <a id="L1005"></a>count++;
            <a id="L1006"></a>}
        <a id="L1007"></a>}
        <a id="L1008"></a>return nil;
    <a id="L1009"></a>});
    <a id="L1010"></a>if err != nil {
        <a id="L1011"></a>p.stopMonitor(err);
        <a id="L1012"></a>return err;
    <a id="L1013"></a>}

    <a id="L1015"></a><span class="comment">// Wait for single stepping threads</span>
    <a id="L1016"></a>for count &gt; 0 {
        <a id="L1017"></a>err = &lt;-ready;
        <a id="L1018"></a>if err != nil {
            <a id="L1019"></a>p.stopMonitor(err);
            <a id="L1020"></a>return err;
        <a id="L1021"></a>}
        <a id="L1022"></a>count--;
    <a id="L1023"></a>}

    <a id="L1025"></a><span class="comment">// Continue all threads</span>
    <a id="L1026"></a>err = p.do(func() os.Error {
        <a id="L1027"></a>if err := p.installBreakpoints(); err != nil {
            <a id="L1028"></a>return err
        <a id="L1029"></a>}

        <a id="L1031"></a>for _, t := range p.threads {
            <a id="L1032"></a>var err os.Error;
            <a id="L1033"></a>switch {
            <a id="L1034"></a>case !t.state.isStopped():
                <a id="L1035"></a>continue

            <a id="L1037"></a>case t.state == stoppedSignal &amp;&amp; t.signal != syscall.SIGSTOP &amp;&amp; t.signal != syscall.SIGTRAP:
                <a id="L1038"></a>t.logTrace(&#34;continuing with signal %d&#34;, t.signal);
                <a id="L1039"></a>err = t.ptraceContWithSignal(t.signal);

            <a id="L1041"></a>default:
                <a id="L1042"></a>t.logTrace(&#34;continuing&#34;);
                <a id="L1043"></a>err = t.ptraceCont();
            <a id="L1044"></a>}
            <a id="L1045"></a>if err != nil {
                <a id="L1046"></a>return err
            <a id="L1047"></a>}
            <a id="L1048"></a>if t.state == stoppedExiting {
                <a id="L1049"></a>t.setState(exiting)
            <a id="L1050"></a>} else {
                <a id="L1051"></a>t.setState(running)
            <a id="L1052"></a>}
        <a id="L1053"></a>}
        <a id="L1054"></a>return nil;
    <a id="L1055"></a>});
    <a id="L1056"></a>if err != nil {
        <a id="L1057"></a><span class="comment">// TODO(austin) Do we need to stop the monitor with</span>
        <a id="L1058"></a><span class="comment">// this error atomically with the do-routine above?</span>
        <a id="L1059"></a>p.stopMonitor(err);
        <a id="L1060"></a>return err;
    <a id="L1061"></a>}

    <a id="L1063"></a>return nil;
<a id="L1064"></a>}

<a id="L1066"></a>func (p *process) WaitStop() os.Error {
    <a id="L1067"></a><span class="comment">// We need a non-blocking ready channel for the case where all</span>
    <a id="L1068"></a><span class="comment">// threads are already stopped.</span>
    <a id="L1069"></a>ready := make(chan os.Error, 1);

    <a id="L1071"></a>err := p.do(func() os.Error {
        <a id="L1072"></a><span class="comment">// Are all of the threads already stopped?</span>
        <a id="L1073"></a>if p.someRunningThread() == nil {
            <a id="L1074"></a>ready &lt;- nil;
            <a id="L1075"></a>return nil;
        <a id="L1076"></a>}

        <a id="L1078"></a><span class="comment">// Monitor state transitions</span>
        <a id="L1079"></a>h := &amp;transitionHandler{};
        <a id="L1080"></a>h.handle = func(st *thread, old, new threadState) {
            <a id="L1081"></a>if !new.isRunning() {
                <a id="L1082"></a>if p.someRunningThread() == nil {
                    <a id="L1083"></a>ready &lt;- nil;
                    <a id="L1084"></a>return;
                <a id="L1085"></a>}
            <a id="L1086"></a>}
            <a id="L1087"></a>p.transitionHandlers.Push(h);
        <a id="L1088"></a>};
        <a id="L1089"></a>h.onErr = func(err os.Error) { ready &lt;- err };
        <a id="L1090"></a>p.transitionHandlers.Push(h);
        <a id="L1091"></a>return nil;
    <a id="L1092"></a>});
    <a id="L1093"></a>if err != nil {
        <a id="L1094"></a>return err
    <a id="L1095"></a>}

    <a id="L1097"></a>return &lt;-ready;
<a id="L1098"></a>}

<a id="L1100"></a>func (p *process) Stop() os.Error {
    <a id="L1101"></a>err := p.do(func() os.Error { return p.stopAsync() });
    <a id="L1102"></a>if err != nil {
        <a id="L1103"></a>return err
    <a id="L1104"></a>}

    <a id="L1106"></a>return p.WaitStop();
<a id="L1107"></a>}

<a id="L1109"></a>func (p *process) Detach() os.Error {
    <a id="L1110"></a>if err := p.Stop(); err != nil {
        <a id="L1111"></a>return err
    <a id="L1112"></a>}

    <a id="L1114"></a>err := p.do(func() os.Error {
        <a id="L1115"></a>if err := p.uninstallBreakpoints(); err != nil {
            <a id="L1116"></a>return err
        <a id="L1117"></a>}

        <a id="L1119"></a>for pid, t := range p.threads {
            <a id="L1120"></a>if t.state.isStopped() {
                <a id="L1121"></a><span class="comment">// We can&#39;t detach from zombies.</span>
                <a id="L1122"></a>if err := t.ptraceDetach(); err != nil {
                    <a id="L1123"></a>return err
                <a id="L1124"></a>}
            <a id="L1125"></a>}
            <a id="L1126"></a>t.setState(detached);
            <a id="L1127"></a>p.threads[pid] = nil, false;
        <a id="L1128"></a>}
        <a id="L1129"></a>return nil;
    <a id="L1130"></a>});
    <a id="L1131"></a><span class="comment">// TODO(austin) Wait for monitor thread to exit?</span>
    <a id="L1132"></a>return err;
<a id="L1133"></a>}

<a id="L1135"></a><span class="comment">// newThread creates a new thread object and waits for its initial</span>
<a id="L1136"></a><span class="comment">// signal.  If cloned is true, this thread was cloned from a thread we</span>
<a id="L1137"></a><span class="comment">// are already attached to.</span>
<a id="L1138"></a><span class="comment">//</span>
<a id="L1139"></a><span class="comment">// Must be run from the monitor thread.</span>
<a id="L1140"></a>func (p *process) newThread(tid int, signal int, cloned bool) (*thread, os.Error) {
    <a id="L1141"></a>t := &amp;thread{tid: tid, proc: p, state: stopped};

    <a id="L1143"></a><span class="comment">// Get the signal from the thread</span>
    <a id="L1144"></a><span class="comment">// TODO(austin) Thread might already be stopped if we&#39;re attaching.</span>
    <a id="L1145"></a>w, err := os.Wait(tid, syscall.WALL);
    <a id="L1146"></a>if err != nil {
        <a id="L1147"></a>return nil, err
    <a id="L1148"></a>}
    <a id="L1149"></a>if w.Pid != tid || w.StopSignal() != signal {
        <a id="L1150"></a>return nil, &amp;newThreadError{w, tid, signal}
    <a id="L1151"></a>}

    <a id="L1153"></a>if !cloned {
        <a id="L1154"></a>err = t.ptraceSetOptions(syscall.PTRACE_O_TRACECLONE | syscall.PTRACE_O_TRACEEXIT);
        <a id="L1155"></a>if err != nil {
            <a id="L1156"></a>return nil, err
        <a id="L1157"></a>}
    <a id="L1158"></a>}

    <a id="L1160"></a>p.threads[tid] = t;

    <a id="L1162"></a>return t, nil;
<a id="L1163"></a>}

<a id="L1165"></a><span class="comment">// attachThread attaches a running thread to the process.</span>
<a id="L1166"></a><span class="comment">//</span>
<a id="L1167"></a><span class="comment">// Must NOT be run from the monitor thread.</span>
<a id="L1168"></a>func (p *process) attachThread(tid int) (*thread, os.Error) {
    <a id="L1169"></a>p.logTrace(&#34;attaching to thread %d&#34;, tid);
    <a id="L1170"></a>var thr *thread;
    <a id="L1171"></a>err := p.do(func() os.Error {
        <a id="L1172"></a>errno := syscall.PtraceAttach(tid);
        <a id="L1173"></a>if errno != 0 {
            <a id="L1174"></a>return os.NewSyscallError(&#34;ptrace(ATTACH)&#34;, errno)
        <a id="L1175"></a>}

        <a id="L1177"></a>var err os.Error;
        <a id="L1178"></a>thr, err = p.newThread(tid, syscall.SIGSTOP, false);
        <a id="L1179"></a>return err;
    <a id="L1180"></a>});
    <a id="L1181"></a>return thr, err;
<a id="L1182"></a>}

<a id="L1184"></a><span class="comment">// attachAllThreads attaches to all threads in a process.</span>
<a id="L1185"></a>func (p *process) attachAllThreads() os.Error {
    <a id="L1186"></a>taskPath := &#34;/proc/&#34; + strconv.Itoa(p.pid) + &#34;/task&#34;;
    <a id="L1187"></a>taskDir, err := os.Open(taskPath, os.O_RDONLY, 0);
    <a id="L1188"></a>if err != nil {
        <a id="L1189"></a>return err
    <a id="L1190"></a>}
    <a id="L1191"></a>defer taskDir.Close();

    <a id="L1193"></a><span class="comment">// We stop threads as we attach to them; however, because new</span>
    <a id="L1194"></a><span class="comment">// threads can appear while we&#39;re looping over all of them, we</span>
    <a id="L1195"></a><span class="comment">// have to repeatly scan until we know we&#39;re attached to all</span>
    <a id="L1196"></a><span class="comment">// of them.</span>
    <a id="L1197"></a>for again := true; again; {
        <a id="L1198"></a>again = false;

        <a id="L1200"></a>tids, err := taskDir.Readdirnames(-1);
        <a id="L1201"></a>if err != nil {
            <a id="L1202"></a>return err
        <a id="L1203"></a>}

        <a id="L1205"></a>for _, tidStr := range tids {
            <a id="L1206"></a>tid, err := strconv.Atoi(tidStr);
            <a id="L1207"></a>if err != nil {
                <a id="L1208"></a>return err
            <a id="L1209"></a>}
            <a id="L1210"></a>if _, ok := p.threads[tid]; ok {
                <a id="L1211"></a>continue
            <a id="L1212"></a>}

            <a id="L1214"></a>_, err = p.attachThread(tid);
            <a id="L1215"></a>if err != nil {
                <a id="L1216"></a><span class="comment">// There could have been a race, or</span>
                <a id="L1217"></a><span class="comment">// this process could be a zobmie.</span>
                <a id="L1218"></a>statFile, err2 := io.ReadFile(taskPath + &#34;/&#34; + tidStr + &#34;/stat&#34;);
                <a id="L1219"></a>if err2 != nil {
                    <a id="L1220"></a>switch err2 := err2.(type) {
                    <a id="L1221"></a>case *os.PathError:
                        <a id="L1222"></a>if err2.Error == os.ENOENT {
                            <a id="L1223"></a><span class="comment">// Raced with thread exit</span>
                            <a id="L1224"></a>p.logTrace(&#34;raced with thread %d exit&#34;, tid);
                            <a id="L1225"></a>continue;
                        <a id="L1226"></a>}
                    <a id="L1227"></a>}
                    <a id="L1228"></a><span class="comment">// Return the original error</span>
                    <a id="L1229"></a>return err;
                <a id="L1230"></a>}

                <a id="L1232"></a>statParts := strings.Split(string(statFile), &#34; &#34;, 4);
                <a id="L1233"></a>if len(statParts) &gt; 2 &amp;&amp; statParts[2] == &#34;Z&#34; {
                    <a id="L1234"></a><span class="comment">// tid is a zombie</span>
                    <a id="L1235"></a>p.logTrace(&#34;thread %d is a zombie&#34;, tid);
                    <a id="L1236"></a>continue;
                <a id="L1237"></a>}

                <a id="L1239"></a><span class="comment">// Return the original error</span>
                <a id="L1240"></a>return err;
            <a id="L1241"></a>}
            <a id="L1242"></a>again = true;
        <a id="L1243"></a>}
    <a id="L1244"></a>}

    <a id="L1246"></a>return nil;
<a id="L1247"></a>}

<a id="L1249"></a><span class="comment">// newProcess creates a new process object and starts its monitor thread.</span>
<a id="L1250"></a>func newProcess(pid int) *process {
    <a id="L1251"></a>p := &amp;process{
        <a id="L1252"></a>pid: pid,
        <a id="L1253"></a>threads: make(map[int]*thread),
        <a id="L1254"></a>breakpoints: make(map[uintptr]*breakpoint),
        <a id="L1255"></a>ready: make(chan bool, 1),
        <a id="L1256"></a>debugEvents: make(chan *debugEvent),
        <a id="L1257"></a>debugReqs: make(chan *debugReq),
        <a id="L1258"></a>stopReq: make(chan os.Error),
        <a id="L1259"></a>transitionHandlers: vector.New(0),
    <a id="L1260"></a>};

    <a id="L1262"></a>go p.monitor();

    <a id="L1264"></a>return p;
<a id="L1265"></a>}

<a id="L1267"></a><span class="comment">// Attach attaches to process pid and stops all of its threads.</span>
<a id="L1268"></a>func Attach(pid int) (Process, os.Error) {
    <a id="L1269"></a>p := newProcess(pid);

    <a id="L1271"></a><span class="comment">// Attach to all threads</span>
    <a id="L1272"></a>err := p.attachAllThreads();
    <a id="L1273"></a>if err != nil {
        <a id="L1274"></a>p.Detach();
        <a id="L1275"></a><span class="comment">// TODO(austin) Detach stopped the monitor already</span>
        <a id="L1276"></a><span class="comment">//p.stopMonitor(err);</span>
        <a id="L1277"></a>return nil, err;
    <a id="L1278"></a>}

    <a id="L1280"></a>return p, nil;
<a id="L1281"></a>}

<a id="L1283"></a><span class="comment">// ForkExec forks the current process and execs argv0, stopping the</span>
<a id="L1284"></a><span class="comment">// new process after the exec syscall.  See os.ForkExec for additional</span>
<a id="L1285"></a><span class="comment">// details.</span>
<a id="L1286"></a>func ForkExec(argv0 string, argv []string, envv []string, dir string, fd []*os.File) (Process, os.Error) {
    <a id="L1287"></a>p := newProcess(-1);

    <a id="L1289"></a><span class="comment">// Create array of integer (system) fds.</span>
    <a id="L1290"></a>intfd := make([]int, len(fd));
    <a id="L1291"></a>for i, f := range fd {
        <a id="L1292"></a>if f == nil {
            <a id="L1293"></a>intfd[i] = -1
        <a id="L1294"></a>} else {
            <a id="L1295"></a>intfd[i] = f.Fd()
        <a id="L1296"></a>}
    <a id="L1297"></a>}

    <a id="L1299"></a><span class="comment">// Fork from the monitor thread so we get the right tracer pid.</span>
    <a id="L1300"></a>err := p.do(func() os.Error {
        <a id="L1301"></a>pid, errno := syscall.PtraceForkExec(argv0, argv, envv, dir, intfd);
        <a id="L1302"></a>if errno != 0 {
            <a id="L1303"></a>return &amp;os.PathError{&#34;fork/exec&#34;, argv0, os.Errno(errno)}
        <a id="L1304"></a>}
        <a id="L1305"></a>p.pid = pid;

        <a id="L1307"></a><span class="comment">// The process will raise SIGTRAP when it reaches execve.</span>
        <a id="L1308"></a>_, err := p.newThread(pid, syscall.SIGTRAP, false);
        <a id="L1309"></a>return err;
    <a id="L1310"></a>});
    <a id="L1311"></a>if err != nil {
        <a id="L1312"></a>p.stopMonitor(err);
        <a id="L1313"></a>return nil, err;
    <a id="L1314"></a>}

    <a id="L1316"></a>return p, nil;
<a id="L1317"></a>}
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
