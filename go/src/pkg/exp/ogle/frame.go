<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/ogle/frame.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/ogle/frame.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package ogle

<a id="L7"></a>import (
    <a id="L8"></a>&#34;debug/gosym&#34;;
    <a id="L9"></a>&#34;debug/proc&#34;;
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;os&#34;;
<a id="L12"></a>)

<a id="L14"></a><span class="comment">// A Frame represents a single frame on a remote call stack.</span>
<a id="L15"></a>type Frame struct {
    <a id="L16"></a><span class="comment">// pc is the PC of the next instruction that will execute in</span>
    <a id="L17"></a><span class="comment">// this frame.  For lower frames, this is the instruction</span>
    <a id="L18"></a><span class="comment">// following the CALL instruction.</span>
    <a id="L19"></a>pc, sp, fp proc.Word;
    <a id="L20"></a><span class="comment">// The runtime.Stktop of the active stack segment</span>
    <a id="L21"></a>stk remoteStruct;
    <a id="L22"></a><span class="comment">// The function this stack frame is in</span>
    <a id="L23"></a>fn  *gosym.Func;
    <a id="L24"></a><span class="comment">// The path and line of the CALL or current instruction.  Note</span>
    <a id="L25"></a><span class="comment">// that this differs slightly from the meaning of Frame.pc.</span>
    <a id="L26"></a>path string;
    <a id="L27"></a>line int;
    <a id="L28"></a><span class="comment">// The inner and outer frames of this frame.  outer is filled</span>
    <a id="L29"></a><span class="comment">// in lazily.</span>
    <a id="L30"></a>inner, outer *Frame;
<a id="L31"></a>}

<a id="L33"></a><span class="comment">// newFrame returns the top-most Frame of the given g&#39;s thread.</span>
<a id="L34"></a>func newFrame(g remoteStruct) (*Frame, os.Error) {
    <a id="L35"></a>var f *Frame;
    <a id="L36"></a>err := try(func(a aborter) { f = aNewFrame(a, g) });
    <a id="L37"></a>return f, err;
<a id="L38"></a>}

<a id="L40"></a>func aNewFrame(a aborter, g remoteStruct) *Frame {
    <a id="L41"></a>p := g.r.p;
    <a id="L42"></a>var pc, sp proc.Word;

    <a id="L44"></a><span class="comment">// Is this G alive?</span>
    <a id="L45"></a>switch g.field(p.f.G.Status).(remoteInt).aGet(a) {
    <a id="L46"></a>case p.runtime.Gidle, p.runtime.Gmoribund, p.runtime.Gdead:
        <a id="L47"></a>return nil
    <a id="L48"></a>}

    <a id="L50"></a><span class="comment">// Find the OS thread for this G</span>

    <a id="L52"></a><span class="comment">// TODO(austin) Ideally, we could look at the G&#39;s state and</span>
    <a id="L53"></a><span class="comment">// figure out if it&#39;s on an OS thread or not.  However, this</span>
    <a id="L54"></a><span class="comment">// is difficult because the state isn&#39;t updated atomically</span>
    <a id="L55"></a><span class="comment">// with scheduling changes.</span>
    <a id="L56"></a>for _, t := range p.proc.Threads() {
        <a id="L57"></a>regs, err := t.Regs();
        <a id="L58"></a>if err != nil {
            <a id="L59"></a><span class="comment">// TODO(austin) What to do?</span>
            <a id="L60"></a>continue
        <a id="L61"></a>}
        <a id="L62"></a>thisg := p.G(regs);
        <a id="L63"></a>if thisg == g.addr().base {
            <a id="L64"></a><span class="comment">// Found this G&#39;s OS thread</span>
            <a id="L65"></a>pc = regs.PC();
            <a id="L66"></a>sp = regs.SP();

            <a id="L68"></a><span class="comment">// If this thread crashed, try to recover it</span>
            <a id="L69"></a>if pc == 0 {
                <a id="L70"></a>pc = p.peekUintptr(a, pc);
                <a id="L71"></a>sp += 8;
            <a id="L72"></a>}

            <a id="L74"></a>break;
        <a id="L75"></a>}
    <a id="L76"></a>}

    <a id="L78"></a>if pc == 0 &amp;&amp; sp == 0 {
        <a id="L79"></a><span class="comment">// G is not mapped to an OS thread.  Use the</span>
        <a id="L80"></a><span class="comment">// scheduler&#39;s stored PC and SP.</span>
        <a id="L81"></a>sched := g.field(p.f.G.Sched).(remoteStruct);
        <a id="L82"></a>pc = proc.Word(sched.field(p.f.Gobuf.Pc).(remoteUint).aGet(a));
        <a id="L83"></a>sp = proc.Word(sched.field(p.f.Gobuf.Sp).(remoteUint).aGet(a));
    <a id="L84"></a>}

    <a id="L86"></a><span class="comment">// Get Stktop</span>
    <a id="L87"></a>stk := g.field(p.f.G.Stackbase).(remotePtr).aGet(a).(remoteStruct);

    <a id="L89"></a>return prepareFrame(a, pc, sp, stk, nil);
<a id="L90"></a>}

<a id="L92"></a><span class="comment">// prepareFrame creates a Frame from the PC and SP within that frame,</span>
<a id="L93"></a><span class="comment">// as well as the active stack segment.  This function takes care of</span>
<a id="L94"></a><span class="comment">// traversing stack breaks and unwinding closures.</span>
<a id="L95"></a>func prepareFrame(a aborter, pc, sp proc.Word, stk remoteStruct, inner *Frame) *Frame {
    <a id="L96"></a><span class="comment">// Based on src/pkg/runtime/amd64/traceback.c:traceback</span>
    <a id="L97"></a>p := stk.r.p;
    <a id="L98"></a>top := inner == nil;

    <a id="L100"></a><span class="comment">// Get function</span>
    <a id="L101"></a>var path string;
    <a id="L102"></a>var line int;
    <a id="L103"></a>var fn *gosym.Func;

    <a id="L105"></a>for i := 0; i &lt; 100; i++ {
        <a id="L106"></a><span class="comment">// Traverse segmented stack breaks</span>
        <a id="L107"></a>if p.sys.lessstack != nil &amp;&amp; pc == proc.Word(p.sys.lessstack.Value) {
            <a id="L108"></a><span class="comment">// Get stk-&gt;gobuf.pc</span>
            <a id="L109"></a>pc = proc.Word(stk.field(p.f.Stktop.Gobuf).(remoteStruct).field(p.f.Gobuf.Pc).(remoteUint).aGet(a));
            <a id="L110"></a><span class="comment">// Get stk-&gt;gobuf.sp</span>
            <a id="L111"></a>sp = proc.Word(stk.field(p.f.Stktop.Gobuf).(remoteStruct).field(p.f.Gobuf.Sp).(remoteUint).aGet(a));
            <a id="L112"></a><span class="comment">// Get stk-&gt;stackbase</span>
            <a id="L113"></a>stk = stk.field(p.f.Stktop.Stackbase).(remotePtr).aGet(a).(remoteStruct);
            <a id="L114"></a>continue;
        <a id="L115"></a>}

        <a id="L117"></a><span class="comment">// Get the PC of the call instruction</span>
        <a id="L118"></a>callpc := pc;
        <a id="L119"></a>if !top &amp;&amp; (p.sys.goexit == nil || pc != proc.Word(p.sys.goexit.Value)) {
            <a id="L120"></a>callpc--
        <a id="L121"></a>}

        <a id="L123"></a><span class="comment">// Look up function</span>
        <a id="L124"></a>path, line, fn = p.syms.PCToLine(uint64(callpc));
        <a id="L125"></a>if fn != nil {
            <a id="L126"></a>break
        <a id="L127"></a>}

        <a id="L129"></a><span class="comment">// Closure?</span>
        <a id="L130"></a>var buf = make([]byte, p.ClosureSize());
        <a id="L131"></a>if _, err := p.Peek(pc, buf); err != nil {
            <a id="L132"></a>break
        <a id="L133"></a>}
        <a id="L134"></a>spdelta, ok := p.ParseClosure(buf);
        <a id="L135"></a>if ok {
            <a id="L136"></a>sp += proc.Word(spdelta);
            <a id="L137"></a>pc = p.peekUintptr(a, sp-proc.Word(p.PtrSize()));
        <a id="L138"></a>}
    <a id="L139"></a>}
    <a id="L140"></a>if fn == nil {
        <a id="L141"></a>return nil
    <a id="L142"></a>}

    <a id="L144"></a><span class="comment">// Compute frame pointer</span>
    <a id="L145"></a>var fp proc.Word;
    <a id="L146"></a>if fn.FrameSize &lt; p.PtrSize() {
        <a id="L147"></a>fp = sp + proc.Word(p.PtrSize())
    <a id="L148"></a>} else {
        <a id="L149"></a>fp = sp + proc.Word(fn.FrameSize)
    <a id="L150"></a>}
    <a id="L151"></a><span class="comment">// TODO(austin) To really figure out if we&#39;re in the prologue,</span>
    <a id="L152"></a><span class="comment">// we need to disassemble the function and look for the call</span>
    <a id="L153"></a><span class="comment">// to morestack.  For now, just special case the entry point.</span>
    <a id="L154"></a><span class="comment">//</span>
    <a id="L155"></a><span class="comment">// TODO(austin) What if we&#39;re in the call to morestack in the</span>
    <a id="L156"></a><span class="comment">// prologue?  Then top == false.</span>
    <a id="L157"></a>if top &amp;&amp; pc == proc.Word(fn.Entry) {
        <a id="L158"></a><span class="comment">// We&#39;re in the function prologue, before SP</span>
        <a id="L159"></a><span class="comment">// has been adjusted for the frame.</span>
        <a id="L160"></a>fp -= proc.Word(fn.FrameSize - p.PtrSize())
    <a id="L161"></a>}

    <a id="L163"></a>return &amp;Frame{pc, sp, fp, stk, fn, path, line, inner, nil};
<a id="L164"></a>}

<a id="L166"></a><span class="comment">// Outer returns the Frame that called this Frame, or nil if this is</span>
<a id="L167"></a><span class="comment">// the outermost frame.</span>
<a id="L168"></a>func (f *Frame) Outer() (*Frame, os.Error) {
    <a id="L169"></a>var fr *Frame;
    <a id="L170"></a>err := try(func(a aborter) { fr = f.aOuter(a) });
    <a id="L171"></a>return fr, err;
<a id="L172"></a>}

<a id="L174"></a>func (f *Frame) aOuter(a aborter) *Frame {
    <a id="L175"></a><span class="comment">// Is there a cached outer frame</span>
    <a id="L176"></a>if f.outer != nil {
        <a id="L177"></a>return f.outer
    <a id="L178"></a>}

    <a id="L180"></a>p := f.stk.r.p;

    <a id="L182"></a>sp := f.fp;
    <a id="L183"></a>if f.fn == p.sys.newproc &amp;&amp; f.fn == p.sys.deferproc {
        <a id="L184"></a><span class="comment">// TODO(rsc) The compiler inserts two push/pop&#39;s</span>
        <a id="L185"></a><span class="comment">// around calls to go and defer.  Russ says this</span>
        <a id="L186"></a><span class="comment">// should get fixed in the compiler, but we account</span>
        <a id="L187"></a><span class="comment">// for it for now.</span>
        <a id="L188"></a>sp += proc.Word(2 * p.PtrSize())
    <a id="L189"></a>}

    <a id="L191"></a>pc := p.peekUintptr(a, f.fp-proc.Word(p.PtrSize()));
    <a id="L192"></a>if pc &lt; 0x1000 {
        <a id="L193"></a>return nil
    <a id="L194"></a>}

    <a id="L196"></a><span class="comment">// TODO(austin) Register this frame for shoot-down.</span>

    <a id="L198"></a>f.outer = prepareFrame(a, pc, sp, f.stk, f);
    <a id="L199"></a>return f.outer;
<a id="L200"></a>}

<a id="L202"></a><span class="comment">// Inner returns the Frame called by this Frame, or nil if this is the</span>
<a id="L203"></a><span class="comment">// innermost frame.</span>
<a id="L204"></a>func (f *Frame) Inner() *Frame { return f.inner }

<a id="L206"></a>func (f *Frame) String() string {
    <a id="L207"></a>res := f.fn.Name;
    <a id="L208"></a>if f.pc &gt; proc.Word(f.fn.Value) {
        <a id="L209"></a>res += fmt.Sprintf(&#34;+%#x&#34;, f.pc-proc.Word(f.fn.Entry))
    <a id="L210"></a>}
    <a id="L211"></a>return res + fmt.Sprintf(&#34; %s:%d&#34;, f.path, f.line);
<a id="L212"></a>}
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
