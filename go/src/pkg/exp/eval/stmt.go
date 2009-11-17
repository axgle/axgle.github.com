<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/stmt.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/stmt.go</h1>

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
    <a id="L8"></a>&#34;bignum&#34;;
    <a id="L9"></a>&#34;log&#34;;
    <a id="L10"></a>&#34;go/ast&#34;;
    <a id="L11"></a>&#34;go/token&#34;;
<a id="L12"></a>)

<a id="L14"></a>const (
    <a id="L15"></a>returnPC = ^uint(0);
    <a id="L16"></a>badPC    = ^uint(1);
<a id="L17"></a>)

<a id="L19"></a><span class="comment">/*</span>
<a id="L20"></a><span class="comment"> * Statement compiler</span>
<a id="L21"></a><span class="comment"> */</span>

<a id="L23"></a>type stmtCompiler struct {
    <a id="L24"></a>*blockCompiler;
    <a id="L25"></a>pos token.Position;
    <a id="L26"></a><span class="comment">// This statement&#39;s label, or nil if it is not labeled.</span>
    <a id="L27"></a>stmtLabel *label;
<a id="L28"></a>}

<a id="L30"></a>func (a *stmtCompiler) diag(format string, args ...) {
    <a id="L31"></a>a.diagAt(&amp;a.pos, format, args)
<a id="L32"></a>}

<a id="L34"></a><span class="comment">/*</span>
<a id="L35"></a><span class="comment"> * Flow checker</span>
<a id="L36"></a><span class="comment"> */</span>

<a id="L38"></a>type flowEnt struct {
    <a id="L39"></a><span class="comment">// Whether this flow entry is conditional.  If true, flow can</span>
    <a id="L40"></a><span class="comment">// continue to the next PC.</span>
    <a id="L41"></a>cond bool;
    <a id="L42"></a><span class="comment">// True if this will terminate flow (e.g., a return statement).</span>
    <a id="L43"></a><span class="comment">// cond must be false and jumps must be nil if this is true.</span>
    <a id="L44"></a>term bool;
    <a id="L45"></a><span class="comment">// PC&#39;s that can be reached from this flow entry.</span>
    <a id="L46"></a>jumps []*uint;
    <a id="L47"></a><span class="comment">// Whether this flow entry has been visited by reachesEnd.</span>
    <a id="L48"></a>visited bool;
<a id="L49"></a>}

<a id="L51"></a>type flowBlock struct {
    <a id="L52"></a><span class="comment">// If this is a goto, the target label.</span>
    <a id="L53"></a>target string;
    <a id="L54"></a><span class="comment">// The inner-most block containing definitions.</span>
    <a id="L55"></a>block *block;
    <a id="L56"></a><span class="comment">// The numVars from each block leading to the root of the</span>
    <a id="L57"></a><span class="comment">// scope, starting at block.</span>
    <a id="L58"></a>numVars []int;
<a id="L59"></a>}

<a id="L61"></a>type flowBuf struct {
    <a id="L62"></a>cb  *codeBuf;
    <a id="L63"></a><span class="comment">// ents is a map from PC&#39;s to flow entries.  Any PC missing</span>
    <a id="L64"></a><span class="comment">// from this map is assumed to reach only PC+1.</span>
    <a id="L65"></a>ents map[uint]*flowEnt;
    <a id="L66"></a><span class="comment">// gotos is a map from goto positions to information on the</span>
    <a id="L67"></a><span class="comment">// block at the point of the goto.</span>
    <a id="L68"></a>gotos map[*token.Position]*flowBlock;
    <a id="L69"></a><span class="comment">// labels is a map from label name to information on the block</span>
    <a id="L70"></a><span class="comment">// at the point of the label.  labels are tracked by name,</span>
    <a id="L71"></a><span class="comment">// since mutliple labels at the same PC can have different</span>
    <a id="L72"></a><span class="comment">// blocks.</span>
    <a id="L73"></a>labels map[string]*flowBlock;
<a id="L74"></a>}

<a id="L76"></a>func newFlowBuf(cb *codeBuf) *flowBuf {
    <a id="L77"></a>return &amp;flowBuf{cb, make(map[uint]*flowEnt), make(map[*token.Position]*flowBlock), make(map[string]*flowBlock)}
<a id="L78"></a>}

<a id="L80"></a><span class="comment">// put creates a flow control point for the next PC in the code buffer.</span>
<a id="L81"></a><span class="comment">// This should be done before pushing the instruction into the code buffer.</span>
<a id="L82"></a>func (f *flowBuf) put(cond bool, term bool, jumps []*uint) {
    <a id="L83"></a>pc := f.cb.nextPC();
    <a id="L84"></a>if ent, ok := f.ents[pc]; ok {
        <a id="L85"></a>log.Crashf(&#34;Flow entry already exists at PC %d: %+v&#34;, pc, ent)
    <a id="L86"></a>}
    <a id="L87"></a>f.ents[pc] = &amp;flowEnt{cond, term, jumps, false};
<a id="L88"></a>}

<a id="L90"></a><span class="comment">// putTerm creates a flow control point at the next PC that</span>
<a id="L91"></a><span class="comment">// unconditionally terminates execution.</span>
<a id="L92"></a>func (f *flowBuf) putTerm() { f.put(false, true, nil) }

<a id="L94"></a><span class="comment">// put1 creates a flow control point at the next PC that jumps to one</span>
<a id="L95"></a><span class="comment">// PC and, if cond is true, can also continue to the PC following the</span>
<a id="L96"></a><span class="comment">// next PC.</span>
<a id="L97"></a>func (f *flowBuf) put1(cond bool, jumpPC *uint) {
    <a id="L98"></a>f.put(cond, false, []*uint{jumpPC})
<a id="L99"></a>}

<a id="L101"></a>func newFlowBlock(target string, b *block) *flowBlock {
    <a id="L102"></a><span class="comment">// Find the inner-most block containing definitions</span>
    <a id="L103"></a>for b.numVars == 0 &amp;&amp; b.outer != nil &amp;&amp; b.outer.scope == b.scope {
        <a id="L104"></a>b = b.outer
    <a id="L105"></a>}

    <a id="L107"></a><span class="comment">// Count parents leading to the root of the scope</span>
    <a id="L108"></a>n := 0;
    <a id="L109"></a>for bp := b; bp.scope == b.scope; bp = bp.outer {
        <a id="L110"></a>n++
    <a id="L111"></a>}

    <a id="L113"></a><span class="comment">// Capture numVars from each block to the root of the scope</span>
    <a id="L114"></a>numVars := make([]int, n);
    <a id="L115"></a>i := 0;
    <a id="L116"></a>for bp := b; i &lt; n; bp = bp.outer {
        <a id="L117"></a>numVars[i] = bp.numVars;
        <a id="L118"></a>i++;
    <a id="L119"></a>}

    <a id="L121"></a>return &amp;flowBlock{target, b, numVars};
<a id="L122"></a>}

<a id="L124"></a><span class="comment">// putGoto captures the block at a goto statement.  This should be</span>
<a id="L125"></a><span class="comment">// called in addition to putting a flow control point.</span>
<a id="L126"></a>func (f *flowBuf) putGoto(pos token.Position, target string, b *block) {
    <a id="L127"></a>f.gotos[&amp;pos] = newFlowBlock(target, b)
<a id="L128"></a>}

<a id="L130"></a><span class="comment">// putLabel captures the block at a label.</span>
<a id="L131"></a>func (f *flowBuf) putLabel(name string, b *block) {
    <a id="L132"></a>f.labels[name] = newFlowBlock(&#34;&#34;, b)
<a id="L133"></a>}

<a id="L135"></a><span class="comment">// reachesEnd returns true if the end of f&#39;s code buffer can be</span>
<a id="L136"></a><span class="comment">// reached from the given program counter.  Error reporting is the</span>
<a id="L137"></a><span class="comment">// caller&#39;s responsibility.</span>
<a id="L138"></a>func (f *flowBuf) reachesEnd(pc uint) bool {
    <a id="L139"></a>endPC := f.cb.nextPC();
    <a id="L140"></a>if pc &gt; endPC {
        <a id="L141"></a>log.Crashf(&#34;Reached bad PC %d past end PC %d&#34;, pc, endPC)
    <a id="L142"></a>}

    <a id="L144"></a>for ; pc &lt; endPC; pc++ {
        <a id="L145"></a>ent, ok := f.ents[pc];
        <a id="L146"></a>if !ok {
            <a id="L147"></a>continue
        <a id="L148"></a>}

        <a id="L150"></a>if ent.visited {
            <a id="L151"></a>return false
        <a id="L152"></a>}
        <a id="L153"></a>ent.visited = true;

        <a id="L155"></a>if ent.term {
            <a id="L156"></a>return false
        <a id="L157"></a>}

        <a id="L159"></a><span class="comment">// If anything can reach the end, we can reach the end</span>
        <a id="L160"></a><span class="comment">// from pc.</span>
        <a id="L161"></a>for _, j := range ent.jumps {
            <a id="L162"></a>if f.reachesEnd(*j) {
                <a id="L163"></a>return true
            <a id="L164"></a>}
        <a id="L165"></a>}
        <a id="L166"></a><span class="comment">// If the jump was conditional, we can reach the next</span>
        <a id="L167"></a><span class="comment">// PC, so try reaching the end from it.</span>
        <a id="L168"></a>if ent.cond {
            <a id="L169"></a>continue
        <a id="L170"></a>}
        <a id="L171"></a>return false;
    <a id="L172"></a>}
    <a id="L173"></a>return true;
<a id="L174"></a>}

<a id="L176"></a><span class="comment">// gotosObeyScopes returns true if no goto statement causes any</span>
<a id="L177"></a><span class="comment">// variables to come into scope that were not in scope at the point of</span>
<a id="L178"></a><span class="comment">// the goto.  Reports any errors using the given compiler.</span>
<a id="L179"></a>func (f *flowBuf) gotosObeyScopes(a *compiler) {
    <a id="L180"></a>for pos, src := range f.gotos {
        <a id="L181"></a>tgt := f.labels[src.target];

        <a id="L183"></a><span class="comment">// The target block must be a parent of this block</span>
        <a id="L184"></a>numVars := src.numVars;
        <a id="L185"></a>b := src.block;
        <a id="L186"></a>for len(numVars) &gt; 0 &amp;&amp; b != tgt.block {
            <a id="L187"></a>b = b.outer;
            <a id="L188"></a>numVars = numVars[1:len(numVars)];
        <a id="L189"></a>}
        <a id="L190"></a>if b != tgt.block {
            <a id="L191"></a><span class="comment">// We jumped into a deeper block</span>
            <a id="L192"></a>a.diagAt(pos, &#34;goto causes variables to come into scope&#34;);
            <a id="L193"></a>return;
        <a id="L194"></a>}

        <a id="L196"></a><span class="comment">// There must be no variables in the target block that</span>
        <a id="L197"></a><span class="comment">// did not exist at the jump</span>
        <a id="L198"></a>tgtNumVars := tgt.numVars;
        <a id="L199"></a>for i := range numVars {
            <a id="L200"></a>if tgtNumVars[i] &gt; numVars[i] {
                <a id="L201"></a>a.diagAt(pos, &#34;goto causes variables to come into scope&#34;);
                <a id="L202"></a>return;
            <a id="L203"></a>}
        <a id="L204"></a>}
    <a id="L205"></a>}
<a id="L206"></a>}

<a id="L208"></a><span class="comment">/*</span>
<a id="L209"></a><span class="comment"> * Statement generation helpers</span>
<a id="L210"></a><span class="comment"> */</span>

<a id="L212"></a>func (a *stmtCompiler) defineVar(ident *ast.Ident, t Type) *Variable {
    <a id="L213"></a>v, prev := a.block.DefineVar(ident.Value, ident.Pos(), t);
    <a id="L214"></a>if prev != nil {
        <a id="L215"></a><span class="comment">// TODO(austin) It&#39;s silly that we have to capture</span>
        <a id="L216"></a><span class="comment">// Pos() in a variable.</span>
        <a id="L217"></a>pos := prev.Pos();
        <a id="L218"></a>if pos.IsValid() {
            <a id="L219"></a>a.diagAt(ident, &#34;variable %s redeclared in this block\n\tprevious declaration at %s&#34;, ident.Value, &amp;pos)
        <a id="L220"></a>} else {
            <a id="L221"></a>a.diagAt(ident, &#34;variable %s redeclared in this block&#34;, ident.Value)
        <a id="L222"></a>}
        <a id="L223"></a>return nil;
    <a id="L224"></a>}

    <a id="L226"></a><span class="comment">// Initialize the variable</span>
    <a id="L227"></a>index := v.Index;
    <a id="L228"></a>if v.Index &gt;= 0 {
        <a id="L229"></a>a.push(func(v *Thread) { v.f.Vars[index] = t.Zero() })
    <a id="L230"></a>}
    <a id="L231"></a>return v;
<a id="L232"></a>}

<a id="L234"></a><span class="comment">// TODO(austin) Move doAssign to here</span>

<a id="L236"></a><span class="comment">/*</span>
<a id="L237"></a><span class="comment"> * Statement compiler</span>
<a id="L238"></a><span class="comment"> */</span>

<a id="L240"></a>func (a *stmtCompiler) compile(s ast.Stmt) {
    <a id="L241"></a>if a.block.inner != nil {
        <a id="L242"></a>log.Crash(&#34;Child scope still entered&#34;)
    <a id="L243"></a>}

    <a id="L245"></a>notimpl := false;
    <a id="L246"></a>switch s := s.(type) {
    <a id="L247"></a>case *ast.BadStmt:
        <a id="L248"></a><span class="comment">// Error already reported by parser.</span>
        <a id="L249"></a>a.silentErrors++

    <a id="L251"></a>case *ast.DeclStmt:
        <a id="L252"></a>a.compileDeclStmt(s)

    <a id="L254"></a>case *ast.EmptyStmt:
        <a id="L255"></a><span class="comment">// Do nothing.</span>

    <a id="L257"></a>case *ast.LabeledStmt:
        <a id="L258"></a>a.compileLabeledStmt(s)

    <a id="L260"></a>case *ast.ExprStmt:
        <a id="L261"></a>a.compileExprStmt(s)

    <a id="L263"></a>case *ast.IncDecStmt:
        <a id="L264"></a>a.compileIncDecStmt(s)

    <a id="L266"></a>case *ast.AssignStmt:
        <a id="L267"></a>a.compileAssignStmt(s)

    <a id="L269"></a>case *ast.GoStmt:
        <a id="L270"></a>notimpl = true

    <a id="L272"></a>case *ast.DeferStmt:
        <a id="L273"></a>notimpl = true

    <a id="L275"></a>case *ast.ReturnStmt:
        <a id="L276"></a>a.compileReturnStmt(s)

    <a id="L278"></a>case *ast.BranchStmt:
        <a id="L279"></a>a.compileBranchStmt(s)

    <a id="L281"></a>case *ast.BlockStmt:
        <a id="L282"></a>a.compileBlockStmt(s)

    <a id="L284"></a>case *ast.IfStmt:
        <a id="L285"></a>a.compileIfStmt(s)

    <a id="L287"></a>case *ast.CaseClause:
        <a id="L288"></a>a.diag(&#34;case clause outside switch&#34;)

    <a id="L290"></a>case *ast.SwitchStmt:
        <a id="L291"></a>a.compileSwitchStmt(s)

    <a id="L293"></a>case *ast.TypeCaseClause:
        <a id="L294"></a>notimpl = true

    <a id="L296"></a>case *ast.TypeSwitchStmt:
        <a id="L297"></a>notimpl = true

    <a id="L299"></a>case *ast.CommClause:
        <a id="L300"></a>notimpl = true

    <a id="L302"></a>case *ast.SelectStmt:
        <a id="L303"></a>notimpl = true

    <a id="L305"></a>case *ast.ForStmt:
        <a id="L306"></a>a.compileForStmt(s)

    <a id="L308"></a>case *ast.RangeStmt:
        <a id="L309"></a>notimpl = true

    <a id="L311"></a>default:
        <a id="L312"></a>log.Crashf(&#34;unexpected ast node type %T&#34;, s)
    <a id="L313"></a>}

    <a id="L315"></a>if notimpl {
        <a id="L316"></a>a.diag(&#34;%T statment node not implemented&#34;, s)
    <a id="L317"></a>}

    <a id="L319"></a>if a.block.inner != nil {
        <a id="L320"></a>log.Crash(&#34;Forgot to exit child scope&#34;)
    <a id="L321"></a>}
<a id="L322"></a>}

<a id="L324"></a>func (a *stmtCompiler) compileDeclStmt(s *ast.DeclStmt) {
    <a id="L325"></a>switch decl := s.Decl.(type) {
    <a id="L326"></a>case *ast.BadDecl:
        <a id="L327"></a><span class="comment">// Do nothing.  Already reported by parser.</span>
        <a id="L328"></a>a.silentErrors++

    <a id="L330"></a>case *ast.FuncDecl:
        <a id="L331"></a>if !a.block.global {
            <a id="L332"></a>log.Crash(&#34;FuncDecl at statement level&#34;)
        <a id="L333"></a>}

    <a id="L335"></a>case *ast.GenDecl:
        <a id="L336"></a>if decl.Tok == token.IMPORT &amp;&amp; !a.block.global {
            <a id="L337"></a>log.Crash(&#34;import at statement level&#34;)
        <a id="L338"></a>}

    <a id="L340"></a>default:
        <a id="L341"></a>log.Crashf(&#34;Unexpected Decl type %T&#34;, s.Decl)
    <a id="L342"></a>}
    <a id="L343"></a>a.compileDecl(s.Decl);
<a id="L344"></a>}

<a id="L346"></a>func (a *stmtCompiler) compileVarDecl(decl *ast.GenDecl) {
    <a id="L347"></a>for _, spec := range decl.Specs {
        <a id="L348"></a>spec := spec.(*ast.ValueSpec);
        <a id="L349"></a>if spec.Values == nil {
            <a id="L350"></a><span class="comment">// Declaration without assignment</span>
            <a id="L351"></a>if spec.Type == nil {
                <a id="L352"></a><span class="comment">// Parser should have caught</span>
                <a id="L353"></a>log.Crash(&#34;Type and Values nil&#34;)
            <a id="L354"></a>}
            <a id="L355"></a>t := a.compileType(a.block, spec.Type);
            <a id="L356"></a><span class="comment">// Define placeholders even if type compile failed</span>
            <a id="L357"></a>for _, n := range spec.Names {
                <a id="L358"></a>a.defineVar(n, t)
            <a id="L359"></a>}
        <a id="L360"></a>} else {
            <a id="L361"></a><span class="comment">// Declaration with assignment</span>
            <a id="L362"></a>lhs := make([]ast.Expr, len(spec.Names));
            <a id="L363"></a>for i, n := range spec.Names {
                <a id="L364"></a>lhs[i] = n
            <a id="L365"></a>}
            <a id="L366"></a>a.doAssign(lhs, spec.Values, decl.Tok, spec.Type);
        <a id="L367"></a>}
    <a id="L368"></a>}
<a id="L369"></a>}

<a id="L371"></a>func (a *stmtCompiler) compileDecl(decl ast.Decl) {
    <a id="L372"></a>switch d := decl.(type) {
    <a id="L373"></a>case *ast.BadDecl:
        <a id="L374"></a><span class="comment">// Do nothing.  Already reported by parser.</span>
        <a id="L375"></a>a.silentErrors++

    <a id="L377"></a>case *ast.FuncDecl:
        <a id="L378"></a>decl := a.compileFuncType(a.block, d.Type);
        <a id="L379"></a>if decl == nil {
            <a id="L380"></a>return
        <a id="L381"></a>}
        <a id="L382"></a><span class="comment">// Declare and initialize v before compiling func</span>
        <a id="L383"></a><span class="comment">// so that body can refer to itself.</span>
        <a id="L384"></a>c, prev := a.block.DefineConst(d.Name.Value, a.pos, decl.Type, decl.Type.Zero());
        <a id="L385"></a>if prev != nil {
            <a id="L386"></a>pos := prev.Pos();
            <a id="L387"></a>if pos.IsValid() {
                <a id="L388"></a>a.diagAt(d.Name, &#34;identifier %s redeclared in this block\n\tprevious declaration at %s&#34;, d.Name.Value, &amp;pos)
            <a id="L389"></a>} else {
                <a id="L390"></a>a.diagAt(d.Name, &#34;identifier %s redeclared in this block&#34;, d.Name.Value)
            <a id="L391"></a>}
        <a id="L392"></a>}
        <a id="L393"></a>fn := a.compileFunc(a.block, decl, d.Body);
        <a id="L394"></a>if c == nil || fn == nil {
            <a id="L395"></a>return
        <a id="L396"></a>}
        <a id="L397"></a>var zeroThread Thread;
        <a id="L398"></a>c.Value.(FuncValue).Set(nil, fn(&amp;zeroThread));

    <a id="L400"></a>case *ast.GenDecl:
        <a id="L401"></a>switch d.Tok {
        <a id="L402"></a>case token.IMPORT:
            <a id="L403"></a>log.Crashf(&#34;%v not implemented&#34;, d.Tok)
        <a id="L404"></a>case token.CONST:
            <a id="L405"></a>log.Crashf(&#34;%v not implemented&#34;, d.Tok)
        <a id="L406"></a>case token.TYPE:
            <a id="L407"></a>a.compileTypeDecl(a.block, d)
        <a id="L408"></a>case token.VAR:
            <a id="L409"></a>a.compileVarDecl(d)
        <a id="L410"></a>}

    <a id="L412"></a>default:
        <a id="L413"></a>log.Crashf(&#34;Unexpected Decl type %T&#34;, decl)
    <a id="L414"></a>}
<a id="L415"></a>}

<a id="L417"></a>func (a *stmtCompiler) compileLabeledStmt(s *ast.LabeledStmt) {
    <a id="L418"></a><span class="comment">// Define label</span>
    <a id="L419"></a>l, ok := a.labels[s.Label.Value];
    <a id="L420"></a>if ok {
        <a id="L421"></a>if l.resolved.IsValid() {
            <a id="L422"></a>a.diag(&#34;label %s redeclared in this block\n\tprevious declaration at %s&#34;, s.Label.Value, &amp;l.resolved)
        <a id="L423"></a>}
    <a id="L424"></a>} else {
        <a id="L425"></a>pc := badPC;
        <a id="L426"></a>l = &amp;label{name: s.Label.Value, gotoPC: &amp;pc};
        <a id="L427"></a>a.labels[l.name] = l;
    <a id="L428"></a>}
    <a id="L429"></a>l.desc = &#34;regular label&#34;;
    <a id="L430"></a>l.resolved = s.Pos();

    <a id="L432"></a><span class="comment">// Set goto PC</span>
    <a id="L433"></a>*l.gotoPC = a.nextPC();

    <a id="L435"></a><span class="comment">// Define flow entry so we can check for jumps over declarations.</span>
    <a id="L436"></a>a.flow.putLabel(l.name, a.block);

    <a id="L438"></a><span class="comment">// Compile the statement.  Reuse our stmtCompiler for simplicity.</span>
    <a id="L439"></a>sc := &amp;stmtCompiler{a.blockCompiler, s.Stmt.Pos(), l};
    <a id="L440"></a>sc.compile(s.Stmt);
<a id="L441"></a>}

<a id="L443"></a>func (a *stmtCompiler) compileExprStmt(s *ast.ExprStmt) {
    <a id="L444"></a>bc := a.enterChild();
    <a id="L445"></a>defer bc.exit();

    <a id="L447"></a>e := a.compileExpr(bc.block, false, s.X);
    <a id="L448"></a>if e == nil {
        <a id="L449"></a>return
    <a id="L450"></a>}

    <a id="L452"></a>if e.exec == nil {
        <a id="L453"></a>a.diag(&#34;%s cannot be used as expression statement&#34;, e.desc);
        <a id="L454"></a>return;
    <a id="L455"></a>}

    <a id="L457"></a>a.push(e.exec);
<a id="L458"></a>}

<a id="L460"></a>func (a *stmtCompiler) compileIncDecStmt(s *ast.IncDecStmt) {
    <a id="L461"></a><span class="comment">// Create temporary block for extractEffect</span>
    <a id="L462"></a>bc := a.enterChild();
    <a id="L463"></a>defer bc.exit();

    <a id="L465"></a>l := a.compileExpr(bc.block, false, s.X);
    <a id="L466"></a>if l == nil {
        <a id="L467"></a>return
    <a id="L468"></a>}

    <a id="L470"></a>if l.evalAddr == nil {
        <a id="L471"></a>l.diag(&#34;cannot assign to %s&#34;, l.desc);
        <a id="L472"></a>return;
    <a id="L473"></a>}
    <a id="L474"></a>if !(l.t.isInteger() || l.t.isFloat()) {
        <a id="L475"></a>l.diagOpType(s.Tok, l.t);
        <a id="L476"></a>return;
    <a id="L477"></a>}

    <a id="L479"></a>var op token.Token;
    <a id="L480"></a>var desc string;
    <a id="L481"></a>switch s.Tok {
    <a id="L482"></a>case token.INC:
        <a id="L483"></a>op = token.ADD;
        <a id="L484"></a>desc = &#34;increment statement&#34;;
    <a id="L485"></a>case token.DEC:
        <a id="L486"></a>op = token.SUB;
        <a id="L487"></a>desc = &#34;decrement statement&#34;;
    <a id="L488"></a>default:
        <a id="L489"></a>log.Crashf(&#34;Unexpected IncDec token %v&#34;, s.Tok)
    <a id="L490"></a>}

    <a id="L492"></a>effect, l := l.extractEffect(bc.block, desc);

    <a id="L494"></a>one := l.newExpr(IdealIntType, &#34;constant&#34;);
    <a id="L495"></a>one.pos = s.Pos();
    <a id="L496"></a>one.eval = func() *bignum.Integer { return bignum.Int(1) };

    <a id="L498"></a>binop := l.compileBinaryExpr(op, l, one);
    <a id="L499"></a>if binop == nil {
        <a id="L500"></a>return
    <a id="L501"></a>}

    <a id="L503"></a>assign := a.compileAssign(s.Pos(), bc.block, l.t, []*expr{binop}, &#34;&#34;, &#34;&#34;);
    <a id="L504"></a>if assign == nil {
        <a id="L505"></a>log.Crashf(&#34;compileAssign type check failed&#34;)
    <a id="L506"></a>}

    <a id="L508"></a>lf := l.evalAddr;
    <a id="L509"></a>a.push(func(v *Thread) {
        <a id="L510"></a>effect(v);
        <a id="L511"></a>assign(lf(v), v);
    <a id="L512"></a>});
<a id="L513"></a>}

<a id="L515"></a>func (a *stmtCompiler) doAssign(lhs []ast.Expr, rhs []ast.Expr, tok token.Token, declTypeExpr ast.Expr) {
    <a id="L516"></a>nerr := a.numError();

    <a id="L518"></a><span class="comment">// Compile right side first so we have the types when</span>
    <a id="L519"></a><span class="comment">// compiling the left side and so we don&#39;t see definitions</span>
    <a id="L520"></a><span class="comment">// made on the left side.</span>
    <a id="L521"></a>rs := make([]*expr, len(rhs));
    <a id="L522"></a>for i, re := range rhs {
        <a id="L523"></a>rs[i] = a.compileExpr(a.block, false, re)
    <a id="L524"></a>}

    <a id="L526"></a>errOp := &#34;assignment&#34;;
    <a id="L527"></a>if tok == token.DEFINE || tok == token.VAR {
        <a id="L528"></a>errOp = &#34;declaration&#34;
    <a id="L529"></a>}
    <a id="L530"></a>ac, ok := a.checkAssign(a.pos, rs, errOp, &#34;value&#34;);
    <a id="L531"></a>ac.allowMapForms(len(lhs));

    <a id="L533"></a><span class="comment">// If this is a definition and the LHS is too big, we won&#39;t be</span>
    <a id="L534"></a><span class="comment">// able to produce the usual error message because we can&#39;t</span>
    <a id="L535"></a><span class="comment">// begin to infer the types of the LHS.</span>
    <a id="L536"></a>if (tok == token.DEFINE || tok == token.VAR) &amp;&amp; len(lhs) &gt; len(ac.rmt.Elems) {
        <a id="L537"></a>a.diag(&#34;not enough values for definition&#34;)
    <a id="L538"></a>}

    <a id="L540"></a><span class="comment">// Compile left type if there is one</span>
    <a id="L541"></a>var declType Type;
    <a id="L542"></a>if declTypeExpr != nil {
        <a id="L543"></a>declType = a.compileType(a.block, declTypeExpr)
    <a id="L544"></a>}

    <a id="L546"></a><span class="comment">// Compile left side</span>
    <a id="L547"></a>ls := make([]*expr, len(lhs));
    <a id="L548"></a>nDefs := 0;
    <a id="L549"></a>for i, le := range lhs {
        <a id="L550"></a><span class="comment">// If this is a definition, get the identifier and its type</span>
        <a id="L551"></a>var ident *ast.Ident;
        <a id="L552"></a>var lt Type;
        <a id="L553"></a>switch tok {
        <a id="L554"></a>case token.DEFINE:
            <a id="L555"></a><span class="comment">// Check that it&#39;s an identifier</span>
            <a id="L556"></a>ident, ok = le.(*ast.Ident);
            <a id="L557"></a>if !ok {
                <a id="L558"></a>a.diagAt(le, &#34;left side of := must be a name&#34;);
                <a id="L559"></a><span class="comment">// Suppress new defitions errors</span>
                <a id="L560"></a>nDefs++;
                <a id="L561"></a>continue;
            <a id="L562"></a>}

            <a id="L564"></a><span class="comment">// Is this simply an assignment?</span>
            <a id="L565"></a>if _, ok := a.block.defs[ident.Value]; ok {
                <a id="L566"></a>ident = nil;
                <a id="L567"></a>break;
            <a id="L568"></a>}
            <a id="L569"></a>nDefs++;

        <a id="L571"></a>case token.VAR:
            <a id="L572"></a>ident = le.(*ast.Ident)
        <a id="L573"></a>}

        <a id="L575"></a><span class="comment">// If it&#39;s a definition, get or infer its type.</span>
        <a id="L576"></a>if ident != nil {
            <a id="L577"></a><span class="comment">// Compute the identifier&#39;s type from the RHS</span>
            <a id="L578"></a><span class="comment">// type.  We use the computed MultiType so we</span>
            <a id="L579"></a><span class="comment">// don&#39;t have to worry about unpacking.</span>
            <a id="L580"></a>switch {
            <a id="L581"></a>case declTypeExpr != nil:
                <a id="L582"></a><span class="comment">// We have a declaration type, use it.</span>
                <a id="L583"></a><span class="comment">// If declType is nil, we gave an</span>
                <a id="L584"></a><span class="comment">// error when we compiled it.</span>
                <a id="L585"></a>lt = declType

            <a id="L587"></a>case i &gt;= len(ac.rmt.Elems):
                <a id="L588"></a><span class="comment">// Define a placeholder.  We already</span>
                <a id="L589"></a><span class="comment">// gave the &#34;not enough&#34; error above.</span>
                <a id="L590"></a>lt = nil

            <a id="L592"></a>case ac.rmt.Elems[i] == nil:
                <a id="L593"></a><span class="comment">// We gave the error when we compiled</span>
                <a id="L594"></a><span class="comment">// the RHS.</span>
                <a id="L595"></a>lt = nil

            <a id="L597"></a>case ac.rmt.Elems[i].isIdeal():
                <a id="L598"></a><span class="comment">// If the type is absent and the</span>
                <a id="L599"></a><span class="comment">// corresponding expression is a</span>
                <a id="L600"></a><span class="comment">// constant expression of ideal</span>
                <a id="L601"></a><span class="comment">// integer or ideal float type, the</span>
                <a id="L602"></a><span class="comment">// type of the declared variable is</span>
                <a id="L603"></a><span class="comment">// int or float respectively.</span>
                <a id="L604"></a>switch {
                <a id="L605"></a>case ac.rmt.Elems[i].isInteger():
                    <a id="L606"></a>lt = IntType
                <a id="L607"></a>case ac.rmt.Elems[i].isFloat():
                    <a id="L608"></a>lt = FloatType
                <a id="L609"></a>default:
                    <a id="L610"></a>log.Crashf(&#34;unexpected ideal type %v&#34;, rs[i].t)
                <a id="L611"></a>}

            <a id="L613"></a>default:
                <a id="L614"></a>lt = ac.rmt.Elems[i]
            <a id="L615"></a>}
        <a id="L616"></a>}

        <a id="L618"></a><span class="comment">// If it&#39;s a definition, define the identifier</span>
        <a id="L619"></a>if ident != nil {
            <a id="L620"></a>if a.defineVar(ident, lt) == nil {
                <a id="L621"></a>continue
            <a id="L622"></a>}
        <a id="L623"></a>}

        <a id="L625"></a><span class="comment">// Compile LHS</span>
        <a id="L626"></a>ls[i] = a.compileExpr(a.block, false, le);
        <a id="L627"></a>if ls[i] == nil {
            <a id="L628"></a>continue
        <a id="L629"></a>}

        <a id="L631"></a>if ls[i].evalMapValue != nil {
            <a id="L632"></a><span class="comment">// Map indexes are not generally addressable,</span>
            <a id="L633"></a><span class="comment">// but they are assignable.</span>
            <a id="L634"></a><span class="comment">//</span>
            <a id="L635"></a><span class="comment">// TODO(austin) Now that the expression</span>
            <a id="L636"></a><span class="comment">// compiler uses semantic values, this might</span>
            <a id="L637"></a><span class="comment">// be easier to implement as a function call.</span>
            <a id="L638"></a>sub := ls[i];
            <a id="L639"></a>ls[i] = ls[i].newExpr(sub.t, sub.desc);
            <a id="L640"></a>ls[i].evalMapValue = sub.evalMapValue;
            <a id="L641"></a>mvf := sub.evalMapValue;
            <a id="L642"></a>et := sub.t;
            <a id="L643"></a>ls[i].evalAddr = func(t *Thread) Value {
                <a id="L644"></a>m, k := mvf(t);
                <a id="L645"></a>e := m.Elem(t, k);
                <a id="L646"></a>if e == nil {
                    <a id="L647"></a>e = et.Zero();
                    <a id="L648"></a>m.SetElem(t, k, e);
                <a id="L649"></a>}
                <a id="L650"></a>return e;
            <a id="L651"></a>};
        <a id="L652"></a>} else if ls[i].evalAddr == nil {
            <a id="L653"></a>ls[i].diag(&#34;cannot assign to %s&#34;, ls[i].desc);
            <a id="L654"></a>continue;
        <a id="L655"></a>}
    <a id="L656"></a>}

    <a id="L658"></a><span class="comment">// A short variable declaration may redeclare variables</span>
    <a id="L659"></a><span class="comment">// provided they were originally declared in the same block</span>
    <a id="L660"></a><span class="comment">// with the same type, and at least one of the variables is</span>
    <a id="L661"></a><span class="comment">// new.</span>
    <a id="L662"></a>if tok == token.DEFINE &amp;&amp; nDefs == 0 {
        <a id="L663"></a>a.diag(&#34;at least one new variable must be declared&#34;);
        <a id="L664"></a>return;
    <a id="L665"></a>}

    <a id="L667"></a><span class="comment">// If there have been errors, our arrays are full of nil&#39;s so</span>
    <a id="L668"></a><span class="comment">// get out of here now.</span>
    <a id="L669"></a>if nerr != a.numError() {
        <a id="L670"></a>return
    <a id="L671"></a>}

    <a id="L673"></a><span class="comment">// Check for &#39;a[x] = r, ok&#39;</span>
    <a id="L674"></a>if len(ls) == 1 &amp;&amp; len(rs) == 2 &amp;&amp; ls[0].evalMapValue != nil {
        <a id="L675"></a>a.diag(&#34;a[x] = r, ok form not implemented&#34;);
        <a id="L676"></a>return;
    <a id="L677"></a>}

    <a id="L679"></a><span class="comment">// Create assigner</span>
    <a id="L680"></a>var lt Type;
    <a id="L681"></a>n := len(lhs);
    <a id="L682"></a>if n == 1 {
        <a id="L683"></a>lt = ls[0].t
    <a id="L684"></a>} else {
        <a id="L685"></a>lts := make([]Type, len(ls));
        <a id="L686"></a>for i, l := range ls {
            <a id="L687"></a>if l != nil {
                <a id="L688"></a>lts[i] = l.t
            <a id="L689"></a>}
        <a id="L690"></a>}
        <a id="L691"></a>lt = NewMultiType(lts);
    <a id="L692"></a>}
    <a id="L693"></a>bc := a.enterChild();
    <a id="L694"></a>defer bc.exit();
    <a id="L695"></a>assign := ac.compile(bc.block, lt);
    <a id="L696"></a>if assign == nil {
        <a id="L697"></a>return
    <a id="L698"></a>}

    <a id="L700"></a><span class="comment">// Compile</span>
    <a id="L701"></a>if n == 1 {
        <a id="L702"></a><span class="comment">// Don&#39;t need temporaries and can avoid []Value.</span>
        <a id="L703"></a>lf := ls[0].evalAddr;
        <a id="L704"></a>a.push(func(t *Thread) { assign(lf(t), t) });
    <a id="L705"></a>} else if tok == token.VAR || (tok == token.DEFINE &amp;&amp; nDefs == n) {
        <a id="L706"></a><span class="comment">// Don&#39;t need temporaries</span>
        <a id="L707"></a>lfs := make([]func(*Thread) Value, n);
        <a id="L708"></a>for i, l := range ls {
            <a id="L709"></a>lfs[i] = l.evalAddr
        <a id="L710"></a>}
        <a id="L711"></a>a.push(func(t *Thread) {
            <a id="L712"></a>dest := make([]Value, n);
            <a id="L713"></a>for i, lf := range lfs {
                <a id="L714"></a>dest[i] = lf(t)
            <a id="L715"></a>}
            <a id="L716"></a>assign(multiV(dest), t);
        <a id="L717"></a>});
    <a id="L718"></a>} else {
        <a id="L719"></a><span class="comment">// Need temporaries</span>
        <a id="L720"></a>lmt := lt.(*MultiType);
        <a id="L721"></a>lfs := make([]func(*Thread) Value, n);
        <a id="L722"></a>for i, l := range ls {
            <a id="L723"></a>lfs[i] = l.evalAddr
        <a id="L724"></a>}
        <a id="L725"></a>a.push(func(t *Thread) {
            <a id="L726"></a>temp := lmt.Zero().(multiV);
            <a id="L727"></a>assign(temp, t);
            <a id="L728"></a><span class="comment">// Copy to destination</span>
            <a id="L729"></a>for i := 0; i &lt; n; i++ {
                <a id="L730"></a><span class="comment">// TODO(austin) Need to evaluate LHS</span>
                <a id="L731"></a><span class="comment">// before RHS</span>
                <a id="L732"></a>lfs[i](t).Assign(t, temp[i])
            <a id="L733"></a>}
        <a id="L734"></a>});
    <a id="L735"></a>}
<a id="L736"></a>}

<a id="L738"></a>var assignOpToOp = map[token.Token]token.Token{
    <a id="L739"></a>token.ADD_ASSIGN: token.ADD,
    <a id="L740"></a>token.SUB_ASSIGN: token.SUB,
    <a id="L741"></a>token.MUL_ASSIGN: token.MUL,
    <a id="L742"></a>token.QUO_ASSIGN: token.QUO,
    <a id="L743"></a>token.REM_ASSIGN: token.REM,

    <a id="L745"></a>token.AND_ASSIGN: token.AND,
    <a id="L746"></a>token.OR_ASSIGN: token.OR,
    <a id="L747"></a>token.XOR_ASSIGN: token.XOR,
    <a id="L748"></a>token.SHL_ASSIGN: token.SHL,
    <a id="L749"></a>token.SHR_ASSIGN: token.SHR,
    <a id="L750"></a>token.AND_NOT_ASSIGN: token.AND_NOT,
<a id="L751"></a>}

<a id="L753"></a>func (a *stmtCompiler) doAssignOp(s *ast.AssignStmt) {
    <a id="L754"></a>if len(s.Lhs) != 1 || len(s.Rhs) != 1 {
        <a id="L755"></a>a.diag(&#34;tuple assignment cannot be combined with an arithmetic operation&#34;);
        <a id="L756"></a>return;
    <a id="L757"></a>}

    <a id="L759"></a><span class="comment">// Create temporary block for extractEffect</span>
    <a id="L760"></a>bc := a.enterChild();
    <a id="L761"></a>defer bc.exit();

    <a id="L763"></a>l := a.compileExpr(bc.block, false, s.Lhs[0]);
    <a id="L764"></a>r := a.compileExpr(bc.block, false, s.Rhs[0]);
    <a id="L765"></a>if l == nil || r == nil {
        <a id="L766"></a>return
    <a id="L767"></a>}

    <a id="L769"></a>if l.evalAddr == nil {
        <a id="L770"></a>l.diag(&#34;cannot assign to %s&#34;, l.desc);
        <a id="L771"></a>return;
    <a id="L772"></a>}

    <a id="L774"></a>effect, l := l.extractEffect(bc.block, &#34;operator-assignment&#34;);

    <a id="L776"></a>binop := r.compileBinaryExpr(assignOpToOp[s.Tok], l, r);
    <a id="L777"></a>if binop == nil {
        <a id="L778"></a>return
    <a id="L779"></a>}

    <a id="L781"></a>assign := a.compileAssign(s.Pos(), bc.block, l.t, []*expr{binop}, &#34;assignment&#34;, &#34;value&#34;);
    <a id="L782"></a>if assign == nil {
        <a id="L783"></a>log.Crashf(&#34;compileAssign type check failed&#34;)
    <a id="L784"></a>}

    <a id="L786"></a>lf := l.evalAddr;
    <a id="L787"></a>a.push(func(t *Thread) {
        <a id="L788"></a>effect(t);
        <a id="L789"></a>assign(lf(t), t);
    <a id="L790"></a>});
<a id="L791"></a>}

<a id="L793"></a>func (a *stmtCompiler) compileAssignStmt(s *ast.AssignStmt) {
    <a id="L794"></a>switch s.Tok {
    <a id="L795"></a>case token.ASSIGN, token.DEFINE:
        <a id="L796"></a>a.doAssign(s.Lhs, s.Rhs, s.Tok, nil)

    <a id="L798"></a>default:
        <a id="L799"></a>a.doAssignOp(s)
    <a id="L800"></a>}
<a id="L801"></a>}

<a id="L803"></a>func (a *stmtCompiler) compileReturnStmt(s *ast.ReturnStmt) {
    <a id="L804"></a>if a.fnType == nil {
        <a id="L805"></a>a.diag(&#34;cannot return at the top level&#34;);
        <a id="L806"></a>return;
    <a id="L807"></a>}

    <a id="L809"></a>if len(s.Results) == 0 &amp;&amp; (len(a.fnType.Out) == 0 || a.outVarsNamed) {
        <a id="L810"></a><span class="comment">// Simple case.  Simply exit from the function.</span>
        <a id="L811"></a>a.flow.putTerm();
        <a id="L812"></a>a.push(func(v *Thread) { v.pc = returnPC });
        <a id="L813"></a>return;
    <a id="L814"></a>}

    <a id="L816"></a>bc := a.enterChild();
    <a id="L817"></a>defer bc.exit();

    <a id="L819"></a><span class="comment">// Compile expressions</span>
    <a id="L820"></a>bad := false;
    <a id="L821"></a>rs := make([]*expr, len(s.Results));
    <a id="L822"></a>for i, re := range s.Results {
        <a id="L823"></a>rs[i] = a.compileExpr(bc.block, false, re);
        <a id="L824"></a>if rs[i] == nil {
            <a id="L825"></a>bad = true
        <a id="L826"></a>}
    <a id="L827"></a>}
    <a id="L828"></a>if bad {
        <a id="L829"></a>return
    <a id="L830"></a>}

    <a id="L832"></a><span class="comment">// Create assigner</span>

    <a id="L834"></a><span class="comment">// However, if the expression list in the &#34;return&#34; statement</span>
    <a id="L835"></a><span class="comment">// is a single call to a multi-valued function, the values</span>
    <a id="L836"></a><span class="comment">// returned from the called function will be returned from</span>
    <a id="L837"></a><span class="comment">// this one.</span>
    <a id="L838"></a>assign := a.compileAssign(s.Pos(), bc.block, NewMultiType(a.fnType.Out), rs, &#34;return&#34;, &#34;value&#34;);

    <a id="L840"></a><span class="comment">// XXX(Spec) &#34;The result types of the current function and the</span>
    <a id="L841"></a><span class="comment">// called function must match.&#34;  Match is fuzzy.  It should</span>
    <a id="L842"></a><span class="comment">// say that they must be assignment compatible.</span>

    <a id="L844"></a><span class="comment">// Compile</span>
    <a id="L845"></a>start := len(a.fnType.In);
    <a id="L846"></a>nout := len(a.fnType.Out);
    <a id="L847"></a>a.flow.putTerm();
    <a id="L848"></a>a.push(func(t *Thread) {
        <a id="L849"></a>assign(multiV(t.f.Vars[start:start+nout]), t);
        <a id="L850"></a>t.pc = returnPC;
    <a id="L851"></a>});
<a id="L852"></a>}

<a id="L854"></a>func (a *stmtCompiler) findLexicalLabel(name *ast.Ident, pred func(*label) bool, errOp, errCtx string) *label {
    <a id="L855"></a>bc := a.blockCompiler;
    <a id="L856"></a>for ; bc != nil; bc = bc.parent {
        <a id="L857"></a>if bc.label == nil {
            <a id="L858"></a>continue
        <a id="L859"></a>}
        <a id="L860"></a>l := bc.label;
        <a id="L861"></a>if name == nil &amp;&amp; pred(l) {
            <a id="L862"></a>return l
        <a id="L863"></a>}
        <a id="L864"></a>if name != nil &amp;&amp; l.name == name.Value {
            <a id="L865"></a>if !pred(l) {
                <a id="L866"></a>a.diag(&#34;cannot %s to %s %s&#34;, errOp, l.desc, l.name);
                <a id="L867"></a>return nil;
            <a id="L868"></a>}
            <a id="L869"></a>return l;
        <a id="L870"></a>}
    <a id="L871"></a>}
    <a id="L872"></a>if name == nil {
        <a id="L873"></a>a.diag(&#34;%s outside %s&#34;, errOp, errCtx)
    <a id="L874"></a>} else {
        <a id="L875"></a>a.diag(&#34;%s label %s not defined&#34;, errOp, name.Value)
    <a id="L876"></a>}
    <a id="L877"></a>return nil;
<a id="L878"></a>}

<a id="L880"></a>func (a *stmtCompiler) compileBranchStmt(s *ast.BranchStmt) {
    <a id="L881"></a>var pc *uint;

    <a id="L883"></a>switch s.Tok {
    <a id="L884"></a>case token.BREAK:
        <a id="L885"></a>l := a.findLexicalLabel(s.Label, func(l *label) bool { return l.breakPC != nil }, &#34;break&#34;, &#34;for loop, switch, or select&#34;);
        <a id="L886"></a>if l == nil {
            <a id="L887"></a>return
        <a id="L888"></a>}
        <a id="L889"></a>pc = l.breakPC;

    <a id="L891"></a>case token.CONTINUE:
        <a id="L892"></a>l := a.findLexicalLabel(s.Label, func(l *label) bool { return l.continuePC != nil }, &#34;continue&#34;, &#34;for loop&#34;);
        <a id="L893"></a>if l == nil {
            <a id="L894"></a>return
        <a id="L895"></a>}
        <a id="L896"></a>pc = l.continuePC;

    <a id="L898"></a>case token.GOTO:
        <a id="L899"></a>l, ok := a.labels[s.Label.Value];
        <a id="L900"></a>if !ok {
            <a id="L901"></a>pc := badPC;
            <a id="L902"></a>l = &amp;label{name: s.Label.Value, desc: &#34;unresolved label&#34;, gotoPC: &amp;pc, used: s.Pos()};
            <a id="L903"></a>a.labels[l.name] = l;
        <a id="L904"></a>}

        <a id="L906"></a>pc = l.gotoPC;
        <a id="L907"></a>a.flow.putGoto(s.Pos(), l.name, a.block);

    <a id="L909"></a>case token.FALLTHROUGH:
        <a id="L910"></a>a.diag(&#34;fallthrough outside switch&#34;);
        <a id="L911"></a>return;

    <a id="L913"></a>default:
        <a id="L914"></a>log.Crash(&#34;Unexpected branch token %v&#34;, s.Tok)
    <a id="L915"></a>}

    <a id="L917"></a>a.flow.put1(false, pc);
    <a id="L918"></a>a.push(func(v *Thread) { v.pc = *pc });
<a id="L919"></a>}

<a id="L921"></a>func (a *stmtCompiler) compileBlockStmt(s *ast.BlockStmt) {
    <a id="L922"></a>bc := a.enterChild();
    <a id="L923"></a>bc.compileStmts(s);
    <a id="L924"></a>bc.exit();
<a id="L925"></a>}

<a id="L927"></a>func (a *stmtCompiler) compileIfStmt(s *ast.IfStmt) {
    <a id="L928"></a><span class="comment">// The scope of any variables declared by [the init] statement</span>
    <a id="L929"></a><span class="comment">// extends to the end of the &#34;if&#34; statement and the variables</span>
    <a id="L930"></a><span class="comment">// are initialized once before the statement is entered.</span>
    <a id="L931"></a><span class="comment">//</span>
    <a id="L932"></a><span class="comment">// XXX(Spec) What this really wants to say is that there&#39;s an</span>
    <a id="L933"></a><span class="comment">// implicit scope wrapping every if, for, and switch</span>
    <a id="L934"></a><span class="comment">// statement.  This is subtly different from what it actually</span>
    <a id="L935"></a><span class="comment">// says when there&#39;s a non-block else clause, because that</span>
    <a id="L936"></a><span class="comment">// else claus has to execute in a scope that is *not* the</span>
    <a id="L937"></a><span class="comment">// surrounding scope.</span>
    <a id="L938"></a>bc := a.enterChild();
    <a id="L939"></a>defer bc.exit();

    <a id="L941"></a><span class="comment">// Compile init statement, if any</span>
    <a id="L942"></a>if s.Init != nil {
        <a id="L943"></a>bc.compileStmt(s.Init)
    <a id="L944"></a>}

    <a id="L946"></a>elsePC := badPC;
    <a id="L947"></a>endPC := badPC;

    <a id="L949"></a><span class="comment">// Compile condition, if any.  If there is no condition, we</span>
    <a id="L950"></a><span class="comment">// fall through to the body.</span>
    <a id="L951"></a>if s.Cond != nil {
        <a id="L952"></a>e := bc.compileExpr(bc.block, false, s.Cond);
        <a id="L953"></a>switch {
        <a id="L954"></a>case e == nil:
            <a id="L955"></a><span class="comment">// Error reported by compileExpr</span>
        <a id="L956"></a>case !e.t.isBoolean():
            <a id="L957"></a>e.diag(&#34;&#39;if&#39; condition must be boolean\n\t%v&#34;, e.t)
        <a id="L958"></a>default:
            <a id="L959"></a>eval := e.asBool();
            <a id="L960"></a>a.flow.put1(true, &amp;elsePC);
            <a id="L961"></a>a.push(func(t *Thread) {
                <a id="L962"></a>if !eval(t) {
                    <a id="L963"></a>t.pc = elsePC
                <a id="L964"></a>}
            <a id="L965"></a>});
        <a id="L966"></a>}
    <a id="L967"></a>}

    <a id="L969"></a><span class="comment">// Compile body</span>
    <a id="L970"></a>body := bc.enterChild();
    <a id="L971"></a>body.compileStmts(s.Body);
    <a id="L972"></a>body.exit();

    <a id="L974"></a><span class="comment">// Compile else</span>
    <a id="L975"></a>if s.Else != nil {
        <a id="L976"></a><span class="comment">// Skip over else if we executed the body</span>
        <a id="L977"></a>a.flow.put1(false, &amp;endPC);
        <a id="L978"></a>a.push(func(v *Thread) { v.pc = endPC });
        <a id="L979"></a>elsePC = a.nextPC();
        <a id="L980"></a>bc.compileStmt(s.Else);
    <a id="L981"></a>} else {
        <a id="L982"></a>elsePC = a.nextPC()
    <a id="L983"></a>}
    <a id="L984"></a>endPC = a.nextPC();
<a id="L985"></a>}

<a id="L987"></a>func (a *stmtCompiler) compileSwitchStmt(s *ast.SwitchStmt) {
    <a id="L988"></a><span class="comment">// Create implicit scope around switch</span>
    <a id="L989"></a>bc := a.enterChild();
    <a id="L990"></a>defer bc.exit();

    <a id="L992"></a><span class="comment">// Compile init statement, if any</span>
    <a id="L993"></a>if s.Init != nil {
        <a id="L994"></a>bc.compileStmt(s.Init)
    <a id="L995"></a>}

    <a id="L997"></a><span class="comment">// Compile condition, if any, and extract its effects</span>
    <a id="L998"></a>var cond *expr;
    <a id="L999"></a>condbc := bc.enterChild();
    <a id="L1000"></a>if s.Tag != nil {
        <a id="L1001"></a>e := condbc.compileExpr(condbc.block, false, s.Tag);
        <a id="L1002"></a>if e != nil {
            <a id="L1003"></a>var effect func(*Thread);
            <a id="L1004"></a>effect, cond = e.extractEffect(condbc.block, &#34;switch&#34;);
            <a id="L1005"></a>a.push(effect);
        <a id="L1006"></a>}
    <a id="L1007"></a>}

    <a id="L1009"></a><span class="comment">// Count cases</span>
    <a id="L1010"></a>ncases := 0;
    <a id="L1011"></a>hasDefault := false;
    <a id="L1012"></a>for _, c := range s.Body.List {
        <a id="L1013"></a>clause, ok := c.(*ast.CaseClause);
        <a id="L1014"></a>if !ok {
            <a id="L1015"></a>a.diagAt(clause, &#34;switch statement must contain case clauses&#34;);
            <a id="L1016"></a>continue;
        <a id="L1017"></a>}
        <a id="L1018"></a>if clause.Values == nil {
            <a id="L1019"></a>if hasDefault {
                <a id="L1020"></a>a.diagAt(clause, &#34;switch statement contains more than one default case&#34;)
            <a id="L1021"></a>}
            <a id="L1022"></a>hasDefault = true;
        <a id="L1023"></a>} else {
            <a id="L1024"></a>ncases += len(clause.Values)
        <a id="L1025"></a>}
    <a id="L1026"></a>}

    <a id="L1028"></a><span class="comment">// Compile case expressions</span>
    <a id="L1029"></a>cases := make([]func(*Thread) bool, ncases);
    <a id="L1030"></a>i := 0;
    <a id="L1031"></a>for _, c := range s.Body.List {
        <a id="L1032"></a>clause, ok := c.(*ast.CaseClause);
        <a id="L1033"></a>if !ok {
            <a id="L1034"></a>continue
        <a id="L1035"></a>}
        <a id="L1036"></a>for _, v := range clause.Values {
            <a id="L1037"></a>e := condbc.compileExpr(condbc.block, false, v);
            <a id="L1038"></a>switch {
            <a id="L1039"></a>case e == nil:
                <a id="L1040"></a><span class="comment">// Error reported by compileExpr</span>
            <a id="L1041"></a>case cond == nil &amp;&amp; !e.t.isBoolean():
                <a id="L1042"></a>a.diagAt(v, &#34;&#39;case&#39; condition must be boolean&#34;)
            <a id="L1043"></a>case cond == nil:
                <a id="L1044"></a>cases[i] = e.asBool()
            <a id="L1045"></a>case cond != nil:
                <a id="L1046"></a><span class="comment">// Create comparison</span>
                <a id="L1047"></a><span class="comment">// TOOD(austin) This produces bad error messages</span>
                <a id="L1048"></a>compare := e.compileBinaryExpr(token.EQL, cond, e);
                <a id="L1049"></a>if compare != nil {
                    <a id="L1050"></a>cases[i] = compare.asBool()
                <a id="L1051"></a>}
            <a id="L1052"></a>}
            <a id="L1053"></a>i++;
        <a id="L1054"></a>}
    <a id="L1055"></a>}

    <a id="L1057"></a><span class="comment">// Emit condition</span>
    <a id="L1058"></a>casePCs := make([]*uint, ncases+1);
    <a id="L1059"></a>endPC := badPC;

    <a id="L1061"></a>a.flow.put(false, false, casePCs);
    <a id="L1062"></a>a.push(func(t *Thread) {
        <a id="L1063"></a>for i, c := range cases {
            <a id="L1064"></a>if c(t) {
                <a id="L1065"></a>t.pc = *casePCs[i];
                <a id="L1066"></a>return;
            <a id="L1067"></a>}
        <a id="L1068"></a>}
        <a id="L1069"></a>t.pc = *casePCs[ncases];
    <a id="L1070"></a>});
    <a id="L1071"></a>condbc.exit();

    <a id="L1073"></a><span class="comment">// Compile cases</span>
    <a id="L1074"></a>i = 0;
    <a id="L1075"></a>for _, c := range s.Body.List {
        <a id="L1076"></a>clause, ok := c.(*ast.CaseClause);
        <a id="L1077"></a>if !ok {
            <a id="L1078"></a>continue
        <a id="L1079"></a>}

        <a id="L1081"></a><span class="comment">// Save jump PC&#39;s</span>
        <a id="L1082"></a>pc := a.nextPC();
        <a id="L1083"></a>if clause.Values != nil {
            <a id="L1084"></a>for _ = range clause.Values {
                <a id="L1085"></a>casePCs[i] = &amp;pc;
                <a id="L1086"></a>i++;
            <a id="L1087"></a>}
        <a id="L1088"></a>} else {
            <a id="L1089"></a><span class="comment">// Default clause</span>
            <a id="L1090"></a>casePCs[ncases] = &amp;pc
        <a id="L1091"></a>}

        <a id="L1093"></a><span class="comment">// Compile body</span>
        <a id="L1094"></a>fall := false;
        <a id="L1095"></a>for j, s := range clause.Body {
            <a id="L1096"></a>if br, ok := s.(*ast.BranchStmt); ok &amp;&amp; br.Tok == token.FALLTHROUGH {
                <a id="L1097"></a><span class="comment">// println(&#34;Found fallthrough&#34;);</span>
                <a id="L1098"></a><span class="comment">// It may be used only as the final</span>
                <a id="L1099"></a><span class="comment">// non-empty statement in a case or</span>
                <a id="L1100"></a><span class="comment">// default clause in an expression</span>
                <a id="L1101"></a><span class="comment">// &#34;switch&#34; statement.</span>
                <a id="L1102"></a>for _, s2 := range clause.Body[j+1 : len(clause.Body)] {
                    <a id="L1103"></a><span class="comment">// XXX(Spec) 6g also considers</span>
                    <a id="L1104"></a><span class="comment">// empty blocks to be empty</span>
                    <a id="L1105"></a><span class="comment">// statements.</span>
                    <a id="L1106"></a>if _, ok := s2.(*ast.EmptyStmt); !ok {
                        <a id="L1107"></a>a.diagAt(s, &#34;fallthrough statement must be final statement in case&#34;);
                        <a id="L1108"></a>break;
                    <a id="L1109"></a>}
                <a id="L1110"></a>}
                <a id="L1111"></a>fall = true;
            <a id="L1112"></a>} else {
                <a id="L1113"></a>bc.compileStmt(s)
            <a id="L1114"></a>}
        <a id="L1115"></a>}
        <a id="L1116"></a><span class="comment">// Jump out of switch, unless there was a fallthrough</span>
        <a id="L1117"></a>if !fall {
            <a id="L1118"></a>a.flow.put1(false, &amp;endPC);
            <a id="L1119"></a>a.push(func(v *Thread) { v.pc = endPC });
        <a id="L1120"></a>}
    <a id="L1121"></a>}

    <a id="L1123"></a><span class="comment">// Get end PC</span>
    <a id="L1124"></a>endPC = a.nextPC();
    <a id="L1125"></a>if !hasDefault {
        <a id="L1126"></a>casePCs[ncases] = &amp;endPC
    <a id="L1127"></a>}
<a id="L1128"></a>}

<a id="L1130"></a>func (a *stmtCompiler) compileForStmt(s *ast.ForStmt) {
    <a id="L1131"></a><span class="comment">// Wrap the entire for in a block.</span>
    <a id="L1132"></a>bc := a.enterChild();
    <a id="L1133"></a>defer bc.exit();

    <a id="L1135"></a><span class="comment">// Compile init statement, if any</span>
    <a id="L1136"></a>if s.Init != nil {
        <a id="L1137"></a>bc.compileStmt(s.Init)
    <a id="L1138"></a>}

    <a id="L1140"></a>bodyPC := badPC;
    <a id="L1141"></a>postPC := badPC;
    <a id="L1142"></a>checkPC := badPC;
    <a id="L1143"></a>endPC := badPC;

    <a id="L1145"></a><span class="comment">// Jump to condition check.  We generate slightly less code by</span>
    <a id="L1146"></a><span class="comment">// placing the condition check after the body.</span>
    <a id="L1147"></a>a.flow.put1(false, &amp;checkPC);
    <a id="L1148"></a>a.push(func(v *Thread) { v.pc = checkPC });

    <a id="L1150"></a><span class="comment">// Compile body</span>
    <a id="L1151"></a>bodyPC = a.nextPC();
    <a id="L1152"></a>body := bc.enterChild();
    <a id="L1153"></a>if a.stmtLabel != nil {
        <a id="L1154"></a>body.label = a.stmtLabel
    <a id="L1155"></a>} else {
        <a id="L1156"></a>body.label = &amp;label{resolved: s.Pos()}
    <a id="L1157"></a>}
    <a id="L1158"></a>body.label.desc = &#34;for loop&#34;;
    <a id="L1159"></a>body.label.breakPC = &amp;endPC;
    <a id="L1160"></a>body.label.continuePC = &amp;postPC;
    <a id="L1161"></a>body.compileStmts(s.Body);
    <a id="L1162"></a>body.exit();

    <a id="L1164"></a><span class="comment">// Compile post, if any</span>
    <a id="L1165"></a>postPC = a.nextPC();
    <a id="L1166"></a>if s.Post != nil {
        <a id="L1167"></a><span class="comment">// TODO(austin) Does the parser disallow short</span>
        <a id="L1168"></a><span class="comment">// declarations in s.Post?</span>
        <a id="L1169"></a>bc.compileStmt(s.Post)
    <a id="L1170"></a>}

    <a id="L1172"></a><span class="comment">// Compile condition check, if any</span>
    <a id="L1173"></a>checkPC = a.nextPC();
    <a id="L1174"></a>if s.Cond == nil {
        <a id="L1175"></a><span class="comment">// If the condition is absent, it is equivalent to true.</span>
        <a id="L1176"></a>a.flow.put1(false, &amp;bodyPC);
        <a id="L1177"></a>a.push(func(v *Thread) { v.pc = bodyPC });
    <a id="L1178"></a>} else {
        <a id="L1179"></a>e := bc.compileExpr(bc.block, false, s.Cond);
        <a id="L1180"></a>switch {
        <a id="L1181"></a>case e == nil:
            <a id="L1182"></a><span class="comment">// Error reported by compileExpr</span>
        <a id="L1183"></a>case !e.t.isBoolean():
            <a id="L1184"></a>a.diag(&#34;&#39;for&#39; condition must be boolean\n\t%v&#34;, e.t)
        <a id="L1185"></a>default:
            <a id="L1186"></a>eval := e.asBool();
            <a id="L1187"></a>a.flow.put1(true, &amp;bodyPC);
            <a id="L1188"></a>a.push(func(t *Thread) {
                <a id="L1189"></a>if eval(t) {
                    <a id="L1190"></a>t.pc = bodyPC
                <a id="L1191"></a>}
            <a id="L1192"></a>});
        <a id="L1193"></a>}
    <a id="L1194"></a>}

    <a id="L1196"></a>endPC = a.nextPC();
<a id="L1197"></a>}

<a id="L1199"></a><span class="comment">/*</span>
<a id="L1200"></a><span class="comment"> * Block compiler</span>
<a id="L1201"></a><span class="comment"> */</span>

<a id="L1203"></a>func (a *blockCompiler) compileStmt(s ast.Stmt) {
    <a id="L1204"></a>sc := &amp;stmtCompiler{a, s.Pos(), nil};
    <a id="L1205"></a>sc.compile(s);
<a id="L1206"></a>}

<a id="L1208"></a>func (a *blockCompiler) compileStmts(block *ast.BlockStmt) {
    <a id="L1209"></a>for _, sub := range block.List {
        <a id="L1210"></a>a.compileStmt(sub)
    <a id="L1211"></a>}
<a id="L1212"></a>}

<a id="L1214"></a>func (a *blockCompiler) enterChild() *blockCompiler {
    <a id="L1215"></a>block := a.block.enterChild();
    <a id="L1216"></a>return &amp;blockCompiler{
        <a id="L1217"></a>funcCompiler: a.funcCompiler,
        <a id="L1218"></a>block: block,
        <a id="L1219"></a>parent: a,
    <a id="L1220"></a>};
<a id="L1221"></a>}

<a id="L1223"></a>func (a *blockCompiler) exit() { a.block.exit() }

<a id="L1225"></a><span class="comment">/*</span>
<a id="L1226"></a><span class="comment"> * Function compiler</span>
<a id="L1227"></a><span class="comment"> */</span>

<a id="L1229"></a>func (a *compiler) compileFunc(b *block, decl *FuncDecl, body *ast.BlockStmt) (func(*Thread) Func) {
    <a id="L1230"></a><span class="comment">// Create body scope</span>
    <a id="L1231"></a><span class="comment">//</span>
    <a id="L1232"></a><span class="comment">// The scope of a parameter or result is the body of the</span>
    <a id="L1233"></a><span class="comment">// corresponding function.</span>
    <a id="L1234"></a>bodyScope := b.ChildScope();
    <a id="L1235"></a>defer bodyScope.exit();
    <a id="L1236"></a>for i, t := range decl.Type.In {
        <a id="L1237"></a>if decl.InNames[i] != nil {
            <a id="L1238"></a>bodyScope.DefineVar(decl.InNames[i].Value, decl.InNames[i].Pos(), t)
        <a id="L1239"></a>} else {
            <a id="L1240"></a>bodyScope.DefineTemp(t)
        <a id="L1241"></a>}
    <a id="L1242"></a>}
    <a id="L1243"></a>for i, t := range decl.Type.Out {
        <a id="L1244"></a>if decl.OutNames[i] != nil {
            <a id="L1245"></a>bodyScope.DefineVar(decl.OutNames[i].Value, decl.OutNames[i].Pos(), t)
        <a id="L1246"></a>} else {
            <a id="L1247"></a>bodyScope.DefineTemp(t)
        <a id="L1248"></a>}
    <a id="L1249"></a>}

    <a id="L1251"></a><span class="comment">// Create block context</span>
    <a id="L1252"></a>cb := newCodeBuf();
    <a id="L1253"></a>fc := &amp;funcCompiler{
        <a id="L1254"></a>compiler: a,
        <a id="L1255"></a>fnType: decl.Type,
        <a id="L1256"></a>outVarsNamed: len(decl.OutNames) &gt; 0 &amp;&amp; decl.OutNames[0] != nil,
        <a id="L1257"></a>codeBuf: cb,
        <a id="L1258"></a>flow: newFlowBuf(cb),
        <a id="L1259"></a>labels: make(map[string]*label),
    <a id="L1260"></a>};
    <a id="L1261"></a>bc := &amp;blockCompiler{
        <a id="L1262"></a>funcCompiler: fc,
        <a id="L1263"></a>block: bodyScope.block,
    <a id="L1264"></a>};

    <a id="L1266"></a><span class="comment">// Compile body</span>
    <a id="L1267"></a>nerr := a.numError();
    <a id="L1268"></a>bc.compileStmts(body);
    <a id="L1269"></a>fc.checkLabels();
    <a id="L1270"></a>if nerr != a.numError() {
        <a id="L1271"></a>return nil
    <a id="L1272"></a>}

    <a id="L1274"></a><span class="comment">// Check that the body returned if necessary.  We only check</span>
    <a id="L1275"></a><span class="comment">// this if there were no errors compiling the body.</span>
    <a id="L1276"></a>if len(decl.Type.Out) &gt; 0 &amp;&amp; fc.flow.reachesEnd(0) {
        <a id="L1277"></a><span class="comment">// XXX(Spec) Not specified.</span>
        <a id="L1278"></a>a.diagAt(&amp;body.Rbrace, &#34;function ends without a return statement&#34;);
        <a id="L1279"></a>return nil;
    <a id="L1280"></a>}

    <a id="L1282"></a>code := fc.get();
    <a id="L1283"></a>maxVars := bodyScope.maxVars;
    <a id="L1284"></a>return func(t *Thread) Func { return &amp;evalFunc{t.f, maxVars, code} };
<a id="L1285"></a>}

<a id="L1287"></a><span class="comment">// Checks that labels were resolved and that all jumps obey scoping</span>
<a id="L1288"></a><span class="comment">// rules.  Reports an error and set fc.err if any check fails.</span>
<a id="L1289"></a>func (a *funcCompiler) checkLabels() {
    <a id="L1290"></a>nerr := a.numError();
    <a id="L1291"></a>for _, l := range a.labels {
        <a id="L1292"></a>if !l.resolved.IsValid() {
            <a id="L1293"></a>a.diagAt(&amp;l.used, &#34;label %s not defined&#34;, l.name)
        <a id="L1294"></a>}
    <a id="L1295"></a>}
    <a id="L1296"></a>if nerr != a.numError() {
        <a id="L1297"></a><span class="comment">// Don&#39;t check scopes if we have unresolved labels</span>
        <a id="L1298"></a>return
    <a id="L1299"></a>}

    <a id="L1301"></a><span class="comment">// Executing the &#34;goto&#34; statement must not cause any variables</span>
    <a id="L1302"></a><span class="comment">// to come into scope that were not already in scope at the</span>
    <a id="L1303"></a><span class="comment">// point of the goto.</span>
    <a id="L1304"></a>a.flow.gotosObeyScopes(a.compiler);
<a id="L1305"></a>}
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
