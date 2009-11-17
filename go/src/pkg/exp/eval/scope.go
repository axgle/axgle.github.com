<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/scope.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/scope.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package eval

<a id="L7"></a>import (
    <a id="L8"></a>&#34;go/token&#34;;
    <a id="L9"></a>&#34;log&#34;;
<a id="L10"></a>)

<a id="L12"></a><span class="comment">/*</span>
<a id="L13"></a><span class="comment"> * Blocks and scopes</span>
<a id="L14"></a><span class="comment"> */</span>

<a id="L16"></a><span class="comment">// A definition can be a *Variable, *Constant, or Type.</span>
<a id="L17"></a>type Def interface {
    <a id="L18"></a>Pos() token.Position;
<a id="L19"></a>}

<a id="L21"></a>type Variable struct {
    <a id="L22"></a>token.Position;
    <a id="L23"></a><span class="comment">// Index of this variable in the Frame structure</span>
    <a id="L24"></a>Index int;
    <a id="L25"></a><span class="comment">// Static type of this variable</span>
    <a id="L26"></a>Type Type;
    <a id="L27"></a><span class="comment">// Value of this variable.  This is only used by Scope.NewFrame;</span>
    <a id="L28"></a><span class="comment">// therefore, it is useful for global scopes but cannot be used</span>
    <a id="L29"></a><span class="comment">// in function scopes.</span>
    <a id="L30"></a>Init Value;
<a id="L31"></a>}

<a id="L33"></a>type Constant struct {
    <a id="L34"></a>token.Position;
    <a id="L35"></a>Type  Type;
    <a id="L36"></a>Value Value;
<a id="L37"></a>}

<a id="L39"></a><span class="comment">// A block represents a definition block in which a name may not be</span>
<a id="L40"></a><span class="comment">// defined more than once.</span>
<a id="L41"></a>type block struct {
    <a id="L42"></a><span class="comment">// The block enclosing this one, including blocks in other</span>
    <a id="L43"></a><span class="comment">// scopes.</span>
    <a id="L44"></a>outer *block;
    <a id="L45"></a><span class="comment">// The nested block currently being compiled, or nil.</span>
    <a id="L46"></a>inner *block;
    <a id="L47"></a><span class="comment">// The Scope containing this block.</span>
    <a id="L48"></a>scope *Scope;
    <a id="L49"></a><span class="comment">// The Variables, Constants, and Types defined in this block.</span>
    <a id="L50"></a>defs map[string]Def;
    <a id="L51"></a><span class="comment">// The index of the first variable defined in this block.</span>
    <a id="L52"></a><span class="comment">// This must be greater than the index of any variable defined</span>
    <a id="L53"></a><span class="comment">// in any parent of this block within the same Scope at the</span>
    <a id="L54"></a><span class="comment">// time this block is entered.</span>
    <a id="L55"></a>offset int;
    <a id="L56"></a><span class="comment">// The number of Variables defined in this block.</span>
    <a id="L57"></a>numVars int;
    <a id="L58"></a><span class="comment">// If global, do not allocate new vars and consts in</span>
    <a id="L59"></a><span class="comment">// the frame; assume that the refs will be compiled in</span>
    <a id="L60"></a><span class="comment">// using defs[name].Init.</span>
    <a id="L61"></a>global bool;
<a id="L62"></a>}

<a id="L64"></a><span class="comment">// A Scope is the compile-time analogue of a Frame, which captures</span>
<a id="L65"></a><span class="comment">// some subtree of blocks.</span>
<a id="L66"></a>type Scope struct {
    <a id="L67"></a><span class="comment">// The root block of this scope.</span>
    <a id="L68"></a>*block;
    <a id="L69"></a><span class="comment">// The maximum number of variables required at any point in</span>
    <a id="L70"></a><span class="comment">// this Scope.  This determines the number of slots needed in</span>
    <a id="L71"></a><span class="comment">// Frame&#39;s created from this Scope at run-time.</span>
    <a id="L72"></a>maxVars int;
<a id="L73"></a>}

<a id="L75"></a>func (b *block) enterChild() *block {
    <a id="L76"></a>if b.inner != nil &amp;&amp; b.inner.scope == b.scope {
        <a id="L77"></a>log.Crash(&#34;Failed to exit child block before entering another child&#34;)
    <a id="L78"></a>}
    <a id="L79"></a>sub := &amp;block{
        <a id="L80"></a>outer: b,
        <a id="L81"></a>scope: b.scope,
        <a id="L82"></a>defs: make(map[string]Def),
        <a id="L83"></a>offset: b.offset + b.numVars,
    <a id="L84"></a>};
    <a id="L85"></a>b.inner = sub;
    <a id="L86"></a>return sub;
<a id="L87"></a>}

<a id="L89"></a>func (b *block) exit() {
    <a id="L90"></a>if b.outer == nil {
        <a id="L91"></a>log.Crash(&#34;Cannot exit top-level block&#34;)
    <a id="L92"></a>}
    <a id="L93"></a>if b.outer.scope == b.scope {
        <a id="L94"></a>if b.outer.inner != b {
            <a id="L95"></a>log.Crash(&#34;Already exited block&#34;)
        <a id="L96"></a>}
        <a id="L97"></a>if b.inner != nil &amp;&amp; b.inner.scope == b.scope {
            <a id="L98"></a>log.Crash(&#34;Exit of parent block without exit of child block&#34;)
        <a id="L99"></a>}
    <a id="L100"></a>}
    <a id="L101"></a>b.outer.inner = nil;
<a id="L102"></a>}

<a id="L104"></a>func (b *block) ChildScope() *Scope {
    <a id="L105"></a>if b.inner != nil &amp;&amp; b.inner.scope == b.scope {
        <a id="L106"></a>log.Crash(&#34;Failed to exit child block before entering a child scope&#34;)
    <a id="L107"></a>}
    <a id="L108"></a>sub := b.enterChild();
    <a id="L109"></a>sub.offset = 0;
    <a id="L110"></a>sub.scope = &amp;Scope{sub, 0};
    <a id="L111"></a>return sub.scope;
<a id="L112"></a>}

<a id="L114"></a>func (b *block) DefineVar(name string, pos token.Position, t Type) (*Variable, Def) {
    <a id="L115"></a>if prev, ok := b.defs[name]; ok {
        <a id="L116"></a>return nil, prev
    <a id="L117"></a>}
    <a id="L118"></a>v := b.defineSlot(t, false);
    <a id="L119"></a>v.Position = pos;
    <a id="L120"></a>b.defs[name] = v;
    <a id="L121"></a>return v, nil;
<a id="L122"></a>}

<a id="L124"></a>func (b *block) DefineTemp(t Type) *Variable { return b.defineSlot(t, true) }

<a id="L126"></a>func (b *block) defineSlot(t Type, temp bool) *Variable {
    <a id="L127"></a>if b.inner != nil &amp;&amp; b.inner.scope == b.scope {
        <a id="L128"></a>log.Crash(&#34;Failed to exit child block before defining variable&#34;)
    <a id="L129"></a>}
    <a id="L130"></a>index := -1;
    <a id="L131"></a>if !b.global || temp {
        <a id="L132"></a>index = b.offset + b.numVars;
        <a id="L133"></a>b.numVars++;
        <a id="L134"></a>if index &gt;= b.scope.maxVars {
            <a id="L135"></a>b.scope.maxVars = index + 1
        <a id="L136"></a>}
    <a id="L137"></a>}
    <a id="L138"></a>v := &amp;Variable{token.Position{}, index, t, nil};
    <a id="L139"></a>return v;
<a id="L140"></a>}

<a id="L142"></a>func (b *block) DefineConst(name string, pos token.Position, t Type, v Value) (*Constant, Def) {
    <a id="L143"></a>if prev, ok := b.defs[name]; ok {
        <a id="L144"></a>return nil, prev
    <a id="L145"></a>}
    <a id="L146"></a>c := &amp;Constant{pos, t, v};
    <a id="L147"></a>b.defs[name] = c;
    <a id="L148"></a>return c, nil;
<a id="L149"></a>}

<a id="L151"></a>func (b *block) DefineType(name string, pos token.Position, t Type) Type {
    <a id="L152"></a>if _, ok := b.defs[name]; ok {
        <a id="L153"></a>return nil
    <a id="L154"></a>}
    <a id="L155"></a>nt := &amp;NamedType{pos, name, nil, true, make(map[string]Method)};
    <a id="L156"></a>if t != nil {
        <a id="L157"></a>nt.Complete(t)
    <a id="L158"></a>}
    <a id="L159"></a>b.defs[name] = nt;
    <a id="L160"></a>return nt;
<a id="L161"></a>}

<a id="L163"></a>func (b *block) Lookup(name string) (bl *block, level int, def Def) {
    <a id="L164"></a>for b != nil {
        <a id="L165"></a>if d, ok := b.defs[name]; ok {
            <a id="L166"></a>return b, level, d
        <a id="L167"></a>}
        <a id="L168"></a>if b.outer != nil &amp;&amp; b.scope != b.outer.scope {
            <a id="L169"></a>level++
        <a id="L170"></a>}
        <a id="L171"></a>b = b.outer;
    <a id="L172"></a>}
    <a id="L173"></a>return nil, 0, nil;
<a id="L174"></a>}

<a id="L176"></a>func (s *Scope) NewFrame(outer *Frame) *Frame { return outer.child(s.maxVars) }

<a id="L178"></a><span class="comment">/*</span>
<a id="L179"></a><span class="comment"> * Frames</span>
<a id="L180"></a><span class="comment"> */</span>

<a id="L182"></a>type Frame struct {
    <a id="L183"></a>Outer *Frame;
    <a id="L184"></a>Vars  []Value;
<a id="L185"></a>}

<a id="L187"></a>func (f *Frame) Get(level int, index int) Value {
    <a id="L188"></a>for ; level &gt; 0; level-- {
        <a id="L189"></a>f = f.Outer
    <a id="L190"></a>}
    <a id="L191"></a>return f.Vars[index];
<a id="L192"></a>}

<a id="L194"></a>func (f *Frame) child(numVars int) *Frame {
    <a id="L195"></a><span class="comment">// TODO(austin) This is probably rather expensive.  All values</span>
    <a id="L196"></a><span class="comment">// require heap allocation and zeroing them when we execute a</span>
    <a id="L197"></a><span class="comment">// definition typically requires some computation.</span>
    <a id="L198"></a>return &amp;Frame{f, make([]Value, numVars)}
<a id="L199"></a>}
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
