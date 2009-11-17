<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/ast/walk.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/go/ast/walk.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package ast

<a id="L7"></a>import &#34;fmt&#34;


<a id="L10"></a><span class="comment">// A Visitor&#39;s Visit method is invoked for each node encountered by Walk.</span>
<a id="L11"></a><span class="comment">// If Visit returns true, Walk is invoked for each of the node&#39;s children.</span>
<a id="L12"></a><span class="comment">//</span>
<a id="L13"></a>type Visitor interface {
    <a id="L14"></a>Visit(node interface{}) bool;
<a id="L15"></a>}


<a id="L18"></a>func walkIdent(v Visitor, x *Ident) {
    <a id="L19"></a>if x != nil {
        <a id="L20"></a>Walk(v, x)
    <a id="L21"></a>}
<a id="L22"></a>}


<a id="L25"></a>func walkCommentGroup(v Visitor, g *CommentGroup) {
    <a id="L26"></a>if g != nil {
        <a id="L27"></a>Walk(v, g)
    <a id="L28"></a>}
<a id="L29"></a>}


<a id="L32"></a>func walkFieldList(v Visitor, list []*Field) {
    <a id="L33"></a>for _, x := range list {
        <a id="L34"></a>Walk(v, x)
    <a id="L35"></a>}
<a id="L36"></a>}


<a id="L39"></a>func walkIdentList(v Visitor, list []*Ident) {
    <a id="L40"></a>for _, x := range list {
        <a id="L41"></a>Walk(v, x)
    <a id="L42"></a>}
<a id="L43"></a>}


<a id="L46"></a>func walkExprList(v Visitor, list []Expr) {
    <a id="L47"></a>for _, x := range list {
        <a id="L48"></a>Walk(v, x)
    <a id="L49"></a>}
<a id="L50"></a>}


<a id="L53"></a>func walkStmtList(v Visitor, list []Stmt) {
    <a id="L54"></a>for _, s := range list {
        <a id="L55"></a>Walk(v, s)
    <a id="L56"></a>}
<a id="L57"></a>}


<a id="L60"></a>func walkBlockStmt(v Visitor, b *BlockStmt) {
    <a id="L61"></a>if b != nil {
        <a id="L62"></a>Walk(v, b)
    <a id="L63"></a>}
<a id="L64"></a>}


<a id="L67"></a><span class="comment">// Walk traverses an AST in depth-first order and invokes v.Visit(n) for each</span>
<a id="L68"></a><span class="comment">// non-nil node n encountered, starting with node. If v.Visit(n) returns true,</span>
<a id="L69"></a><span class="comment">// Walk visits each of the children of n.</span>
<a id="L70"></a><span class="comment">//</span>
<a id="L71"></a>func Walk(v Visitor, node interface{}) {
    <a id="L72"></a>if node == nil || !v.Visit(node) {
        <a id="L73"></a>return
    <a id="L74"></a>}

    <a id="L76"></a><span class="comment">// walk children</span>
    <a id="L77"></a><span class="comment">// (the order of the cases matches the order</span>
    <a id="L78"></a><span class="comment">// of the corresponding declaration in ast.go)</span>
    <a id="L79"></a>switch n := node.(type) {
    <a id="L80"></a><span class="comment">// Comments and fields</span>
    <a id="L81"></a>case *Comment:
        <a id="L82"></a><span class="comment">// nothing to do</span>

    <a id="L84"></a>case *CommentGroup:
        <a id="L85"></a>for _, c := range n.List {
            <a id="L86"></a>Walk(v, c)
        <a id="L87"></a>}
        <a id="L88"></a><span class="comment">// TODO(gri): Keep comments in a list/vector instead</span>
        <a id="L89"></a><span class="comment">// of linking them via Next. Following next will lead</span>
        <a id="L90"></a><span class="comment">// to multiple visits and potentially n^2 behavior</span>
        <a id="L91"></a><span class="comment">// since Doc and Comments fields point into the global</span>
        <a id="L92"></a><span class="comment">// comments list.</span>

    <a id="L94"></a>case *Field:
        <a id="L95"></a>walkCommentGroup(v, n.Doc);
        <a id="L96"></a>walkIdentList(v, n.Names);
        <a id="L97"></a>Walk(v, n.Type);
        <a id="L98"></a>for _, x := range n.Tag {
            <a id="L99"></a>Walk(v, x)
        <a id="L100"></a>}
        <a id="L101"></a>walkCommentGroup(v, n.Comment);

    <a id="L103"></a><span class="comment">// Expressions</span>
    <a id="L104"></a>case *BadExpr, *Ident, *Ellipsis, *BasicLit:
        <a id="L105"></a><span class="comment">// nothing to do</span>

    <a id="L107"></a>case *StringList:
        <a id="L108"></a>for _, x := range n.Strings {
            <a id="L109"></a>Walk(v, x)
        <a id="L110"></a>}

    <a id="L112"></a>case *FuncLit:
        <a id="L113"></a>if n != nil {
            <a id="L114"></a>Walk(v, n.Type)
        <a id="L115"></a>}
        <a id="L116"></a>walkBlockStmt(v, n.Body);

    <a id="L118"></a>case *CompositeLit:
        <a id="L119"></a>Walk(v, n.Type);
        <a id="L120"></a>walkExprList(v, n.Elts);

    <a id="L122"></a>case *ParenExpr:
        <a id="L123"></a>Walk(v, n.X)

    <a id="L125"></a>case *SelectorExpr:
        <a id="L126"></a>Walk(v, n.X);
        <a id="L127"></a>walkIdent(v, n.Sel);

    <a id="L129"></a>case *IndexExpr:
        <a id="L130"></a>Walk(v, n.X);
        <a id="L131"></a>Walk(v, n.Index);
        <a id="L132"></a>Walk(v, n.End);

    <a id="L134"></a>case *TypeAssertExpr:
        <a id="L135"></a>Walk(v, n.X);
        <a id="L136"></a>Walk(v, n.Type);

    <a id="L138"></a>case *CallExpr:
        <a id="L139"></a>Walk(v, n.Fun);
        <a id="L140"></a>walkExprList(v, n.Args);

    <a id="L142"></a>case *StarExpr:
        <a id="L143"></a>Walk(v, n.X)

    <a id="L145"></a>case *UnaryExpr:
        <a id="L146"></a>Walk(v, n.X)

    <a id="L148"></a>case *BinaryExpr:
        <a id="L149"></a>Walk(v, n.X);
        <a id="L150"></a>Walk(v, n.Y);

    <a id="L152"></a>case *KeyValueExpr:
        <a id="L153"></a>Walk(v, n.Key);
        <a id="L154"></a>Walk(v, n.Value);

    <a id="L156"></a><span class="comment">// Types</span>
    <a id="L157"></a>case *ArrayType:
        <a id="L158"></a>Walk(v, n.Len);
        <a id="L159"></a>Walk(v, n.Elt);

    <a id="L161"></a>case *StructType:
        <a id="L162"></a>walkFieldList(v, n.Fields)

    <a id="L164"></a>case *FuncType:
        <a id="L165"></a>walkFieldList(v, n.Params);
        <a id="L166"></a>walkFieldList(v, n.Results);

    <a id="L168"></a>case *InterfaceType:
        <a id="L169"></a>walkFieldList(v, n.Methods)

    <a id="L171"></a>case *MapType:
        <a id="L172"></a>Walk(v, n.Key);
        <a id="L173"></a>Walk(v, n.Value);

    <a id="L175"></a>case *ChanType:
        <a id="L176"></a>Walk(v, n.Value)

    <a id="L178"></a><span class="comment">// Statements</span>
    <a id="L179"></a>case *BadStmt:
        <a id="L180"></a><span class="comment">// nothing to do</span>

    <a id="L182"></a>case *DeclStmt:
        <a id="L183"></a>Walk(v, n.Decl)

    <a id="L185"></a>case *EmptyStmt:
        <a id="L186"></a><span class="comment">// nothing to do</span>

    <a id="L188"></a>case *LabeledStmt:
        <a id="L189"></a>walkIdent(v, n.Label);
        <a id="L190"></a>Walk(v, n.Stmt);

    <a id="L192"></a>case *ExprStmt:
        <a id="L193"></a>Walk(v, n.X)

    <a id="L195"></a>case *IncDecStmt:
        <a id="L196"></a>Walk(v, n.X)

    <a id="L198"></a>case *AssignStmt:
        <a id="L199"></a>walkExprList(v, n.Lhs);
        <a id="L200"></a>walkExprList(v, n.Rhs);

    <a id="L202"></a>case *GoStmt:
        <a id="L203"></a>if n.Call != nil {
            <a id="L204"></a>Walk(v, n.Call)
        <a id="L205"></a>}

    <a id="L207"></a>case *DeferStmt:
        <a id="L208"></a>if n.Call != nil {
            <a id="L209"></a>Walk(v, n.Call)
        <a id="L210"></a>}

    <a id="L212"></a>case *ReturnStmt:
        <a id="L213"></a>walkExprList(v, n.Results)

    <a id="L215"></a>case *BranchStmt:
        <a id="L216"></a>walkIdent(v, n.Label)

    <a id="L218"></a>case *BlockStmt:
        <a id="L219"></a>walkStmtList(v, n.List)

    <a id="L221"></a>case *IfStmt:
        <a id="L222"></a>Walk(v, n.Init);
        <a id="L223"></a>Walk(v, n.Cond);
        <a id="L224"></a>walkBlockStmt(v, n.Body);
        <a id="L225"></a>Walk(v, n.Else);

    <a id="L227"></a>case *CaseClause:
        <a id="L228"></a>walkExprList(v, n.Values);
        <a id="L229"></a>walkStmtList(v, n.Body);

    <a id="L231"></a>case *SwitchStmt:
        <a id="L232"></a>Walk(v, n.Init);
        <a id="L233"></a>Walk(v, n.Tag);
        <a id="L234"></a>walkBlockStmt(v, n.Body);

    <a id="L236"></a>case *TypeCaseClause:
        <a id="L237"></a>walkExprList(v, n.Types);
        <a id="L238"></a>walkStmtList(v, n.Body);

    <a id="L240"></a>case *TypeSwitchStmt:
        <a id="L241"></a>Walk(v, n.Init);
        <a id="L242"></a>Walk(v, n.Assign);
        <a id="L243"></a>walkBlockStmt(v, n.Body);

    <a id="L245"></a>case *CommClause:
        <a id="L246"></a>Walk(v, n.Lhs);
        <a id="L247"></a>Walk(v, n.Rhs);
        <a id="L248"></a>walkStmtList(v, n.Body);

    <a id="L250"></a>case *SelectStmt:
        <a id="L251"></a>walkBlockStmt(v, n.Body)

    <a id="L253"></a>case *ForStmt:
        <a id="L254"></a>Walk(v, n.Init);
        <a id="L255"></a>Walk(v, n.Cond);
        <a id="L256"></a>Walk(v, n.Post);
        <a id="L257"></a>walkBlockStmt(v, n.Body);

    <a id="L259"></a>case *RangeStmt:
        <a id="L260"></a>Walk(v, n.Key);
        <a id="L261"></a>Walk(v, n.Value);
        <a id="L262"></a>Walk(v, n.X);
        <a id="L263"></a>walkBlockStmt(v, n.Body);

    <a id="L265"></a><span class="comment">// Declarations</span>
    <a id="L266"></a>case *ImportSpec:
        <a id="L267"></a>walkCommentGroup(v, n.Doc);
        <a id="L268"></a>walkIdent(v, n.Name);
        <a id="L269"></a>for _, x := range n.Path {
            <a id="L270"></a>Walk(v, x)
        <a id="L271"></a>}
        <a id="L272"></a>walkCommentGroup(v, n.Comment);

    <a id="L274"></a>case *ValueSpec:
        <a id="L275"></a>walkCommentGroup(v, n.Doc);
        <a id="L276"></a>walkIdentList(v, n.Names);
        <a id="L277"></a>Walk(v, n.Type);
        <a id="L278"></a>walkExprList(v, n.Values);
        <a id="L279"></a>walkCommentGroup(v, n.Comment);

    <a id="L281"></a>case *TypeSpec:
        <a id="L282"></a>walkCommentGroup(v, n.Doc);
        <a id="L283"></a>walkIdent(v, n.Name);
        <a id="L284"></a>Walk(v, n.Type);
        <a id="L285"></a>walkCommentGroup(v, n.Comment);

    <a id="L287"></a>case *BadDecl:
        <a id="L288"></a><span class="comment">// nothing to do</span>

    <a id="L290"></a>case *GenDecl:
        <a id="L291"></a>walkCommentGroup(v, n.Doc);
        <a id="L292"></a>for _, s := range n.Specs {
            <a id="L293"></a>Walk(v, s)
        <a id="L294"></a>}

    <a id="L296"></a>case *FuncDecl:
        <a id="L297"></a>walkCommentGroup(v, n.Doc);
        <a id="L298"></a>if n.Recv != nil {
            <a id="L299"></a>Walk(v, n.Recv)
        <a id="L300"></a>}
        <a id="L301"></a>walkIdent(v, n.Name);
        <a id="L302"></a>if n.Type != nil {
            <a id="L303"></a>Walk(v, n.Type)
        <a id="L304"></a>}
        <a id="L305"></a>walkBlockStmt(v, n.Body);

    <a id="L307"></a><span class="comment">// Files and packages</span>
    <a id="L308"></a>case *File:
        <a id="L309"></a>walkCommentGroup(v, n.Doc);
        <a id="L310"></a>walkIdent(v, n.Name);
        <a id="L311"></a>for _, d := range n.Decls {
            <a id="L312"></a>Walk(v, d)
        <a id="L313"></a>}
        <a id="L314"></a>walkCommentGroup(v, n.Comments);

    <a id="L316"></a>case *Package:
        <a id="L317"></a>for _, f := range n.Files {
            <a id="L318"></a>Walk(v, f)
        <a id="L319"></a>}

    <a id="L321"></a>default:
        <a id="L322"></a>fmt.Printf(&#34;ast.Walk: unexpected type %T&#34;, n);
        <a id="L323"></a>panic();
    <a id="L324"></a>}
<a id="L325"></a>}
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
