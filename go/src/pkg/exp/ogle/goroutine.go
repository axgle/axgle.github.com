<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/ogle/goroutine.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/ogle/goroutine.go</h1>

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
    <a id="L8"></a>&#34;debug/proc&#34;;
    <a id="L9"></a>&#34;exp/eval&#34;;
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;os&#34;;
<a id="L12"></a>)

<a id="L14"></a><span class="comment">// A Goroutine represents a goroutine in a remote process.</span>
<a id="L15"></a>type Goroutine struct {
    <a id="L16"></a>g     remoteStruct;
    <a id="L17"></a>frame *Frame;
    <a id="L18"></a>dead  bool;
<a id="L19"></a>}

<a id="L21"></a>func (t *Goroutine) String() string {
    <a id="L22"></a>if t.dead {
        <a id="L23"></a>return &#34;&lt;dead thread&gt;&#34;
    <a id="L24"></a>}
    <a id="L25"></a><span class="comment">// TODO(austin) Give threads friendly ID&#39;s, possibly including</span>
    <a id="L26"></a><span class="comment">// the name of the entry function.</span>
    <a id="L27"></a>return fmt.Sprintf(&#34;thread %#x&#34;, t.g.addr().base);
<a id="L28"></a>}

<a id="L30"></a><span class="comment">// isG0 returns true if this thread if the internal idle thread</span>
<a id="L31"></a>func (t *Goroutine) isG0() bool { return t.g.addr().base == t.g.r.p.sys.g0.addr().base }

<a id="L33"></a>func (t *Goroutine) resetFrame() (err os.Error) {
    <a id="L34"></a><span class="comment">// TODO(austin) Reuse any live part of the current frame stack</span>
    <a id="L35"></a><span class="comment">// so existing references to Frame&#39;s keep working.</span>
    <a id="L36"></a>t.frame, err = newFrame(t.g);
    <a id="L37"></a>return;
<a id="L38"></a>}

<a id="L40"></a><span class="comment">// Out selects the caller frame of the current frame.</span>
<a id="L41"></a>func (t *Goroutine) Out() os.Error {
    <a id="L42"></a>f, err := t.frame.Outer();
    <a id="L43"></a>if f != nil {
        <a id="L44"></a>t.frame = f
    <a id="L45"></a>}
    <a id="L46"></a>return err;
<a id="L47"></a>}

<a id="L49"></a><span class="comment">// In selects the frame called by the current frame.</span>
<a id="L50"></a>func (t *Goroutine) In() os.Error {
    <a id="L51"></a>f := t.frame.Inner();
    <a id="L52"></a>if f != nil {
        <a id="L53"></a>t.frame = f
    <a id="L54"></a>}
    <a id="L55"></a>return nil;
<a id="L56"></a>}

<a id="L58"></a>func readylockedBP(ev Event) (EventAction, os.Error) {
    <a id="L59"></a>b := ev.(*Breakpoint);
    <a id="L60"></a>p := b.Process();

    <a id="L62"></a><span class="comment">// The new g is the only argument to this function, so the</span>
    <a id="L63"></a><span class="comment">// stack will have the return address, then the G*.</span>
    <a id="L64"></a>regs, err := b.osThread.Regs();
    <a id="L65"></a>if err != nil {
        <a id="L66"></a>return EAStop, err
    <a id="L67"></a>}
    <a id="L68"></a>sp := regs.SP();
    <a id="L69"></a>addr := sp + proc.Word(p.PtrSize());
    <a id="L70"></a>arg := remotePtr{remote{addr, p}, p.runtime.G};
    <a id="L71"></a>var gp eval.Value;
    <a id="L72"></a>err = try(func(a aborter) { gp = arg.aGet(a) });
    <a id="L73"></a>if err != nil {
        <a id="L74"></a>return EAStop, err
    <a id="L75"></a>}
    <a id="L76"></a>if gp == nil {
        <a id="L77"></a>return EAStop, UnknownGoroutine{b.osThread, 0}
    <a id="L78"></a>}
    <a id="L79"></a>gs := gp.(remoteStruct);
    <a id="L80"></a>g := &amp;Goroutine{gs, nil, false};
    <a id="L81"></a>p.goroutines[gs.addr().base] = g;

    <a id="L83"></a><span class="comment">// Enqueue goroutine creation event</span>
    <a id="L84"></a>parent := b.Goroutine();
    <a id="L85"></a>if parent.isG0() {
        <a id="L86"></a>parent = nil
    <a id="L87"></a>}
    <a id="L88"></a>p.postEvent(&amp;GoroutineCreate{commonEvent{p, g}, parent});

    <a id="L90"></a><span class="comment">// If we don&#39;t have any thread selected, select this one</span>
    <a id="L91"></a>if p.curGoroutine == nil {
        <a id="L92"></a>p.curGoroutine = g
    <a id="L93"></a>}

    <a id="L95"></a>return EADefault, nil;
<a id="L96"></a>}

<a id="L98"></a>func goexitBP(ev Event) (EventAction, os.Error) {
    <a id="L99"></a>b := ev.(*Breakpoint);
    <a id="L100"></a>p := b.Process();

    <a id="L102"></a>g := b.Goroutine();
    <a id="L103"></a>g.dead = true;

    <a id="L105"></a>addr := g.g.addr().base;
    <a id="L106"></a>p.goroutines[addr] = nil, false;

    <a id="L108"></a><span class="comment">// Enqueue thread exit event</span>
    <a id="L109"></a>p.postEvent(&amp;GoroutineExit{commonEvent{p, g}});

    <a id="L111"></a><span class="comment">// If we just exited our selected goroutine, selected another</span>
    <a id="L112"></a>if p.curGoroutine == g {
        <a id="L113"></a>p.selectSomeGoroutine()
    <a id="L114"></a>}

    <a id="L116"></a>return EADefault, nil;
<a id="L117"></a>}
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
