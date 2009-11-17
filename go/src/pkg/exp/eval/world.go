<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/world.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/world.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package is the beginning of an interpreter for Go.</span>
<a id="L6"></a><span class="comment">// It can run simple Go programs but does not implement</span>
<a id="L7"></a><span class="comment">// interface values or packages.</span>
<a id="L8"></a>package eval

<a id="L10"></a>import (
    <a id="L11"></a>&#34;go/ast&#34;;
    <a id="L12"></a>&#34;go/parser&#34;;
    <a id="L13"></a>&#34;go/scanner&#34;;
    <a id="L14"></a>&#34;go/token&#34;;
    <a id="L15"></a>&#34;os&#34;;
<a id="L16"></a>)

<a id="L18"></a>type World struct {
    <a id="L19"></a>scope *Scope;
    <a id="L20"></a>frame *Frame;
<a id="L21"></a>}

<a id="L23"></a>func NewWorld() *World {
    <a id="L24"></a>w := new(World);
    <a id="L25"></a>w.scope = universe.ChildScope();
    <a id="L26"></a>w.scope.global = true; <span class="comment">// this block&#39;s vars allocate directly</span>
    <a id="L27"></a>return w;
<a id="L28"></a>}

<a id="L30"></a>type Code interface {
    <a id="L31"></a><span class="comment">// The type of the value Run returns, or nil if Run returns nil.</span>
    <a id="L32"></a>Type() Type;

    <a id="L34"></a><span class="comment">// Run runs the code; if the code is a single expression</span>
    <a id="L35"></a><span class="comment">// with a value, it returns the value; otherwise it returns nil.</span>
    <a id="L36"></a>Run() (Value, os.Error);
<a id="L37"></a>}

<a id="L39"></a>type stmtCode struct {
    <a id="L40"></a>w    *World;
    <a id="L41"></a>code code;
<a id="L42"></a>}

<a id="L44"></a>func (w *World) CompileStmtList(stmts []ast.Stmt) (Code, os.Error) {
    <a id="L45"></a>if len(stmts) == 1 {
        <a id="L46"></a>if s, ok := stmts[0].(*ast.ExprStmt); ok {
            <a id="L47"></a>return w.CompileExpr(s.X)
        <a id="L48"></a>}
    <a id="L49"></a>}
    <a id="L50"></a>errors := scanner.NewErrorVector();
    <a id="L51"></a>cc := &amp;compiler{errors, 0, 0};
    <a id="L52"></a>cb := newCodeBuf();
    <a id="L53"></a>fc := &amp;funcCompiler{
        <a id="L54"></a>compiler: cc,
        <a id="L55"></a>fnType: nil,
        <a id="L56"></a>outVarsNamed: false,
        <a id="L57"></a>codeBuf: cb,
        <a id="L58"></a>flow: newFlowBuf(cb),
        <a id="L59"></a>labels: make(map[string]*label),
    <a id="L60"></a>};
    <a id="L61"></a>bc := &amp;blockCompiler{
        <a id="L62"></a>funcCompiler: fc,
        <a id="L63"></a>block: w.scope.block,
    <a id="L64"></a>};
    <a id="L65"></a>nerr := cc.numError();
    <a id="L66"></a>for _, stmt := range stmts {
        <a id="L67"></a>bc.compileStmt(stmt)
    <a id="L68"></a>}
    <a id="L69"></a>fc.checkLabels();
    <a id="L70"></a>if nerr != cc.numError() {
        <a id="L71"></a>return nil, errors.GetError(scanner.Sorted)
    <a id="L72"></a>}
    <a id="L73"></a>return &amp;stmtCode{w, fc.get()}, nil;
<a id="L74"></a>}

<a id="L76"></a>func (w *World) CompileDeclList(decls []ast.Decl) (Code, os.Error) {
    <a id="L77"></a>stmts := make([]ast.Stmt, len(decls));
    <a id="L78"></a>for i, d := range decls {
        <a id="L79"></a>stmts[i] = &amp;ast.DeclStmt{d}
    <a id="L80"></a>}
    <a id="L81"></a>return w.CompileStmtList(stmts);
<a id="L82"></a>}

<a id="L84"></a>func (s *stmtCode) Type() Type { return nil }

<a id="L86"></a>func (s *stmtCode) Run() (Value, os.Error) {
    <a id="L87"></a>t := new(Thread);
    <a id="L88"></a>t.f = s.w.scope.NewFrame(nil);
    <a id="L89"></a>return nil, t.Try(func(t *Thread) { s.code.exec(t) });
<a id="L90"></a>}

<a id="L92"></a>type exprCode struct {
    <a id="L93"></a>w    *World;
    <a id="L94"></a>e    *expr;
    <a id="L95"></a>eval func(Value, *Thread);
<a id="L96"></a>}

<a id="L98"></a>func (w *World) CompileExpr(e ast.Expr) (Code, os.Error) {
    <a id="L99"></a>errors := scanner.NewErrorVector();
    <a id="L100"></a>cc := &amp;compiler{errors, 0, 0};

    <a id="L102"></a>ec := cc.compileExpr(w.scope.block, false, e);
    <a id="L103"></a>if ec == nil {
        <a id="L104"></a>return nil, errors.GetError(scanner.Sorted)
    <a id="L105"></a>}
    <a id="L106"></a>var eval func(Value, *Thread);
    <a id="L107"></a>switch t := ec.t.(type) {
    <a id="L108"></a>case *idealIntType:
        <a id="L109"></a><span class="comment">// nothing</span>
    <a id="L110"></a>case *idealFloatType:
        <a id="L111"></a><span class="comment">// nothing</span>
    <a id="L112"></a>default:
        <a id="L113"></a>if tm, ok := t.(*MultiType); ok &amp;&amp; len(tm.Elems) == 0 {
            <a id="L114"></a>return &amp;stmtCode{w, code{ec.exec}}, nil
        <a id="L115"></a>}
        <a id="L116"></a>eval = genAssign(ec.t, ec);
    <a id="L117"></a>}
    <a id="L118"></a>return &amp;exprCode{w, ec, eval}, nil;
<a id="L119"></a>}

<a id="L121"></a>func (e *exprCode) Type() Type { return e.e.t }

<a id="L123"></a>func (e *exprCode) Run() (Value, os.Error) {
    <a id="L124"></a>t := new(Thread);
    <a id="L125"></a>t.f = e.w.scope.NewFrame(nil);
    <a id="L126"></a>switch e.e.t.(type) {
    <a id="L127"></a>case *idealIntType:
        <a id="L128"></a>return &amp;idealIntV{e.e.asIdealInt()()}, nil
    <a id="L129"></a>case *idealFloatType:
        <a id="L130"></a>return &amp;idealFloatV{e.e.asIdealFloat()()}, nil
    <a id="L131"></a>}
    <a id="L132"></a>v := e.e.t.Zero();
    <a id="L133"></a>eval := e.eval;
    <a id="L134"></a>err := t.Try(func(t *Thread) { eval(v, t) });
    <a id="L135"></a>return v, err;
<a id="L136"></a>}

<a id="L138"></a>func (w *World) Compile(text string) (Code, os.Error) {
    <a id="L139"></a>stmts, err := parser.ParseStmtList(&#34;input&#34;, text);
    <a id="L140"></a>if err == nil {
        <a id="L141"></a>return w.CompileStmtList(stmts)
    <a id="L142"></a>}

    <a id="L144"></a><span class="comment">// Otherwise try as DeclList.</span>
    <a id="L145"></a>decls, err1 := parser.ParseDeclList(&#34;input&#34;, text);
    <a id="L146"></a>if err1 == nil {
        <a id="L147"></a>return w.CompileDeclList(decls)
    <a id="L148"></a>}

    <a id="L150"></a><span class="comment">// Have to pick an error.</span>
    <a id="L151"></a><span class="comment">// Parsing as statement list admits more forms,</span>
    <a id="L152"></a><span class="comment">// its error is more likely to be useful.</span>
    <a id="L153"></a>return nil, err;
<a id="L154"></a>}

<a id="L156"></a>type RedefinitionError struct {
    <a id="L157"></a>Name string;
    <a id="L158"></a>Prev Def;
<a id="L159"></a>}

<a id="L161"></a>func (e *RedefinitionError) String() string {
    <a id="L162"></a>res := &#34;identifier &#34; + e.Name + &#34; redeclared&#34;;
    <a id="L163"></a>pos := e.Prev.Pos();
    <a id="L164"></a>if pos.IsValid() {
        <a id="L165"></a>res += &#34;; previous declaration at &#34; + pos.String()
    <a id="L166"></a>}
    <a id="L167"></a>return res;
<a id="L168"></a>}

<a id="L170"></a>func (w *World) DefineConst(name string, t Type, val Value) os.Error {
    <a id="L171"></a>_, prev := w.scope.DefineConst(name, token.Position{}, t, val);
    <a id="L172"></a>if prev != nil {
        <a id="L173"></a>return &amp;RedefinitionError{name, prev}
    <a id="L174"></a>}
    <a id="L175"></a>return nil;
<a id="L176"></a>}

<a id="L178"></a>func (w *World) DefineVar(name string, t Type, val Value) os.Error {
    <a id="L179"></a>v, prev := w.scope.DefineVar(name, token.Position{}, t);
    <a id="L180"></a>if prev != nil {
        <a id="L181"></a>return &amp;RedefinitionError{name, prev}
    <a id="L182"></a>}
    <a id="L183"></a>v.Init = val;
    <a id="L184"></a>return nil;
<a id="L185"></a>}
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
