<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/ogle/event.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/ogle/event.go</h1>

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
    <a id="L9"></a>&#34;fmt&#34;;
    <a id="L10"></a>&#34;os&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">/*</span>
<a id="L14"></a><span class="comment"> * Hooks and events</span>
<a id="L15"></a><span class="comment"> */</span>

<a id="L17"></a><span class="comment">// An EventHandler is a function that takes an event and returns a</span>
<a id="L18"></a><span class="comment">// response to that event and possibly an error.  If an event handler</span>
<a id="L19"></a><span class="comment">// returns an error, the process stops and no other handlers for that</span>
<a id="L20"></a><span class="comment">// event are executed.</span>
<a id="L21"></a>type EventHandler func(e Event) (EventAction, os.Error)

<a id="L23"></a><span class="comment">// An EventAction is an event handler&#39;s response to an event.  If all</span>
<a id="L24"></a><span class="comment">// of an event&#39;s handlers execute without returning errors, their</span>
<a id="L25"></a><span class="comment">// results are combined as follows: If any handler returned</span>
<a id="L26"></a><span class="comment">// EAContinue, then the process resumes (without returning from</span>
<a id="L27"></a><span class="comment">// WaitStop); otherwise, if any handler returned EAStop, the process</span>
<a id="L28"></a><span class="comment">// remains stopped; otherwise, if all handlers returned EADefault, the</span>
<a id="L29"></a><span class="comment">// process resumes.  A handler may return EARemoveSelf bit-wise or&#39;d</span>
<a id="L30"></a><span class="comment">// with any other action to indicate that the handler should be</span>
<a id="L31"></a><span class="comment">// removed from the hook.</span>
<a id="L32"></a>type EventAction int

<a id="L34"></a>const (
    <a id="L35"></a>EARemoveSelf EventAction = 0x100;
    <a id="L36"></a>EADefault    EventAction = iota;
    <a id="L37"></a>EAStop;
    <a id="L38"></a>EAContinue;
<a id="L39"></a>)

<a id="L41"></a><span class="comment">// A EventHook allows event handlers to be added and removed.</span>
<a id="L42"></a>type EventHook interface {
    <a id="L43"></a>AddHandler(EventHandler);
    <a id="L44"></a>RemoveHandler(EventHandler);
    <a id="L45"></a>NumHandler() int;
    <a id="L46"></a>handle(e Event) (EventAction, os.Error);
    <a id="L47"></a>String() string;
<a id="L48"></a>}

<a id="L50"></a><span class="comment">// EventHook is almost, but not quite, suitable for user-defined</span>
<a id="L51"></a><span class="comment">// events.  If we want user-defined events, make EventHook a struct,</span>
<a id="L52"></a><span class="comment">// special-case adding and removing handlers in breakpoint hooks, and</span>
<a id="L53"></a><span class="comment">// provide a public interface for posting events to hooks.</span>

<a id="L55"></a>type Event interface {
    <a id="L56"></a>Process() *Process;
    <a id="L57"></a>Goroutine() *Goroutine;
    <a id="L58"></a>String() string;
<a id="L59"></a>}

<a id="L61"></a>type commonHook struct {
    <a id="L62"></a><span class="comment">// Head of handler chain</span>
    <a id="L63"></a>head *handler;
    <a id="L64"></a><span class="comment">// Number of non-internal handlers</span>
    <a id="L65"></a>len int;
<a id="L66"></a>}

<a id="L68"></a>type handler struct {
    <a id="L69"></a>eh  EventHandler;
    <a id="L70"></a><span class="comment">// True if this handler must be run before user-defined</span>
    <a id="L71"></a><span class="comment">// handlers in order to ensure correctness.</span>
    <a id="L72"></a>internal bool;
    <a id="L73"></a><span class="comment">// True if this handler has been removed from the chain.</span>
    <a id="L74"></a>removed bool;
    <a id="L75"></a>next    *handler;
<a id="L76"></a>}

<a id="L78"></a>func (h *commonHook) AddHandler(eh EventHandler) {
    <a id="L79"></a>h.addHandler(eh, false)
<a id="L80"></a>}

<a id="L82"></a>func (h *commonHook) addHandler(eh EventHandler, internal bool) {
    <a id="L83"></a><span class="comment">// Ensure uniqueness of handlers</span>
    <a id="L84"></a>h.RemoveHandler(eh);

    <a id="L86"></a>if !internal {
        <a id="L87"></a>h.len++
    <a id="L88"></a>}
    <a id="L89"></a><span class="comment">// Add internal handlers to the beginning</span>
    <a id="L90"></a>if internal || h.head == nil {
        <a id="L91"></a>h.head = &amp;handler{eh, internal, false, h.head};
        <a id="L92"></a>return;
    <a id="L93"></a>}
    <a id="L94"></a><span class="comment">// Add handler after internal handlers</span>
    <a id="L95"></a><span class="comment">// TODO(austin) This should probably go on the end instead</span>
    <a id="L96"></a>prev := h.head;
    <a id="L97"></a>for prev.next != nil &amp;&amp; prev.internal {
        <a id="L98"></a>prev = prev.next
    <a id="L99"></a>}
    <a id="L100"></a>prev.next = &amp;handler{eh, internal, false, prev.next};
<a id="L101"></a>}

<a id="L103"></a>func (h *commonHook) RemoveHandler(eh EventHandler) {
    <a id="L104"></a>plink := &amp;h.head;
    <a id="L105"></a>for l := *plink; l != nil; plink, l = &amp;l.next, l.next {
        <a id="L106"></a>if l.eh == eh {
            <a id="L107"></a>if !l.internal {
                <a id="L108"></a>h.len--
            <a id="L109"></a>}
            <a id="L110"></a>l.removed = true;
            <a id="L111"></a>*plink = l.next;
            <a id="L112"></a>break;
        <a id="L113"></a>}
    <a id="L114"></a>}
<a id="L115"></a>}

<a id="L117"></a>func (h *commonHook) NumHandler() int { return h.len }

<a id="L119"></a>func (h *commonHook) handle(e Event) (EventAction, os.Error) {
    <a id="L120"></a>action := EADefault;
    <a id="L121"></a>plink := &amp;h.head;
    <a id="L122"></a>for l := *plink; l != nil; plink, l = &amp;l.next, l.next {
        <a id="L123"></a>if l.removed {
            <a id="L124"></a>continue
        <a id="L125"></a>}
        <a id="L126"></a>a, err := l.eh(e);
        <a id="L127"></a>if a&amp;EARemoveSelf == EARemoveSelf {
            <a id="L128"></a>if !l.internal {
                <a id="L129"></a>h.len--
            <a id="L130"></a>}
            <a id="L131"></a>l.removed = true;
            <a id="L132"></a>*plink = l.next;
            <a id="L133"></a>a &amp;^= EARemoveSelf;
        <a id="L134"></a>}
        <a id="L135"></a>if err != nil {
            <a id="L136"></a>return EAStop, err
        <a id="L137"></a>}
        <a id="L138"></a>if a &gt; action {
            <a id="L139"></a>action = a
        <a id="L140"></a>}
    <a id="L141"></a>}
    <a id="L142"></a>return action, nil;
<a id="L143"></a>}

<a id="L145"></a>type commonEvent struct {
    <a id="L146"></a><span class="comment">// The process of this event</span>
    <a id="L147"></a>p   *Process;
    <a id="L148"></a><span class="comment">// The goroutine of this event.</span>
    <a id="L149"></a>t   *Goroutine;
<a id="L150"></a>}

<a id="L152"></a>func (e *commonEvent) Process() *Process { return e.p }

<a id="L154"></a>func (e *commonEvent) Goroutine() *Goroutine { return e.t }

<a id="L156"></a><span class="comment">/*</span>
<a id="L157"></a><span class="comment"> * Standard event handlers</span>
<a id="L158"></a><span class="comment"> */</span>

<a id="L160"></a><span class="comment">// EventPrint is a standard event handler that prints events as they</span>
<a id="L161"></a><span class="comment">// occur.  It will not cause the process to stop.</span>
<a id="L162"></a>func EventPrint(ev Event) (EventAction, os.Error) {
    <a id="L163"></a><span class="comment">// TODO(austin) Include process name here?</span>
    <a id="L164"></a>fmt.Fprintf(os.Stderr, &#34;*** %v\n&#34;, ev.String());
    <a id="L165"></a>return EADefault, nil;
<a id="L166"></a>}

<a id="L168"></a><span class="comment">// EventStop is a standard event handler that causes the process to stop.</span>
<a id="L169"></a>func EventStop(ev Event) (EventAction, os.Error) {
    <a id="L170"></a>return EAStop, nil
<a id="L171"></a>}

<a id="L173"></a><span class="comment">/*</span>
<a id="L174"></a><span class="comment"> * Breakpoints</span>
<a id="L175"></a><span class="comment"> */</span>

<a id="L177"></a>type breakpointHook struct {
    <a id="L178"></a>commonHook;
    <a id="L179"></a>p   *Process;
    <a id="L180"></a>pc  proc.Word;
<a id="L181"></a>}

<a id="L183"></a><span class="comment">// A Breakpoint event occurs when a process reaches a particular</span>
<a id="L184"></a><span class="comment">// program counter.  When this event is handled, the current goroutine</span>
<a id="L185"></a><span class="comment">// will be the goroutine that reached the program counter.</span>
<a id="L186"></a>type Breakpoint struct {
    <a id="L187"></a>commonEvent;
    <a id="L188"></a>osThread proc.Thread;
    <a id="L189"></a>pc       proc.Word;
<a id="L190"></a>}

<a id="L192"></a>func (h *breakpointHook) AddHandler(eh EventHandler) {
    <a id="L193"></a>h.addHandler(eh, false)
<a id="L194"></a>}

<a id="L196"></a>func (h *breakpointHook) addHandler(eh EventHandler, internal bool) {
    <a id="L197"></a><span class="comment">// We register breakpoint events lazily to avoid holding</span>
    <a id="L198"></a><span class="comment">// references to breakpoints without handlers.  Be sure to use</span>
    <a id="L199"></a><span class="comment">// the &#34;canonical&#34; breakpoint if there is one.</span>
    <a id="L200"></a>if cur, ok := h.p.breakpointHooks[h.pc]; ok {
        <a id="L201"></a>h = cur
    <a id="L202"></a>}
    <a id="L203"></a>oldhead := h.head;
    <a id="L204"></a>h.commonHook.addHandler(eh, internal);
    <a id="L205"></a>if oldhead == nil &amp;&amp; h.head != nil {
        <a id="L206"></a>h.p.proc.AddBreakpoint(h.pc);
        <a id="L207"></a>h.p.breakpointHooks[h.pc] = h;
    <a id="L208"></a>}
<a id="L209"></a>}

<a id="L211"></a>func (h *breakpointHook) RemoveHandler(eh EventHandler) {
    <a id="L212"></a>oldhead := h.head;
    <a id="L213"></a>h.commonHook.RemoveHandler(eh);
    <a id="L214"></a>if oldhead != nil &amp;&amp; h.head == nil {
        <a id="L215"></a>h.p.proc.RemoveBreakpoint(h.pc);
        <a id="L216"></a>h.p.breakpointHooks[h.pc] = nil, false;
    <a id="L217"></a>}
<a id="L218"></a>}

<a id="L220"></a>func (h *breakpointHook) String() string {
    <a id="L221"></a><span class="comment">// TODO(austin) Include process name?</span>
    <a id="L222"></a><span class="comment">// TODO(austin) Use line:pc or at least sym+%#x</span>
    <a id="L223"></a>return fmt.Sprintf(&#34;breakpoint at %#x&#34;, h.pc)
<a id="L224"></a>}

<a id="L226"></a>func (b *Breakpoint) PC() proc.Word { return b.pc }

<a id="L228"></a>func (b *Breakpoint) String() string {
    <a id="L229"></a><span class="comment">// TODO(austin) Include process name and goroutine</span>
    <a id="L230"></a><span class="comment">// TODO(austin) Use line:pc or at least sym+%#x</span>
    <a id="L231"></a>return fmt.Sprintf(&#34;breakpoint at %#x&#34;, b.pc)
<a id="L232"></a>}

<a id="L234"></a><span class="comment">/*</span>
<a id="L235"></a><span class="comment"> * Goroutine create/exit</span>
<a id="L236"></a><span class="comment"> */</span>

<a id="L238"></a>type goroutineCreateHook struct {
    <a id="L239"></a>commonHook;
<a id="L240"></a>}

<a id="L242"></a>func (h *goroutineCreateHook) String() string { return &#34;goroutine create&#34; }

<a id="L244"></a><span class="comment">// A GoroutineCreate event occurs when a process creates a new</span>
<a id="L245"></a><span class="comment">// goroutine.  When this event is handled, the current goroutine will</span>
<a id="L246"></a><span class="comment">// be the newly created goroutine.</span>
<a id="L247"></a>type GoroutineCreate struct {
    <a id="L248"></a>commonEvent;
    <a id="L249"></a>parent *Goroutine;
<a id="L250"></a>}

<a id="L252"></a><span class="comment">// Parent returns the goroutine that created this goroutine.  May be</span>
<a id="L253"></a><span class="comment">// nil if this event is the creation of the first goroutine.</span>
<a id="L254"></a>func (e *GoroutineCreate) Parent() *Goroutine { return e.parent }

<a id="L256"></a>func (e *GoroutineCreate) String() string {
    <a id="L257"></a><span class="comment">// TODO(austin) Include process name</span>
    <a id="L258"></a>if e.parent == nil {
        <a id="L259"></a>return fmt.Sprintf(&#34;%v created&#34;, e.t)
    <a id="L260"></a>}
    <a id="L261"></a>return fmt.Sprintf(&#34;%v created by %v&#34;, e.t, e.parent);
<a id="L262"></a>}

<a id="L264"></a>type goroutineExitHook struct {
    <a id="L265"></a>commonHook;
<a id="L266"></a>}

<a id="L268"></a>func (h *goroutineExitHook) String() string { return &#34;goroutine exit&#34; }

<a id="L270"></a><span class="comment">// A GoroutineExit event occurs when a Go goroutine exits.</span>
<a id="L271"></a>type GoroutineExit struct {
    <a id="L272"></a>commonEvent;
<a id="L273"></a>}

<a id="L275"></a>func (e *GoroutineExit) String() string {
    <a id="L276"></a><span class="comment">// TODO(austin) Include process name</span>
    <a id="L277"></a><span class="comment">//return fmt.Sprintf(&#34;%v exited&#34;, e.t);</span>
    <a id="L278"></a><span class="comment">// For debugging purposes</span>
    <a id="L279"></a>return fmt.Sprintf(&#34;goroutine %#x exited&#34;, e.t.g.addr().base)
<a id="L280"></a>}
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
