<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/cgo/ast.go</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/cgo/ast.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Parse input AST and prepare Prog structure.</span>

<a id="L7"></a>package main

<a id="L9"></a>import (
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;go/ast&#34;;
    <a id="L12"></a>&#34;go/doc&#34;;
    <a id="L13"></a>&#34;go/parser&#34;;
    <a id="L14"></a>&#34;go/scanner&#34;;
    <a id="L15"></a>&#34;os&#34;;
<a id="L16"></a>)

<a id="L18"></a><span class="comment">// A Cref refers to an expression of the form C.xxx in the AST.</span>
<a id="L19"></a>type Cref struct {
    <a id="L20"></a>Name     string;
    <a id="L21"></a>Expr     *ast.Expr;
    <a id="L22"></a>Context  string; <span class="comment">// &#34;type&#34;, &#34;expr&#34;, or &#34;call&#34;</span>
    <a id="L23"></a>TypeName bool;   <span class="comment">// whether xxx is a C type name</span>
    <a id="L24"></a>Type     *Type;  <span class="comment">// the type of xxx</span>
    <a id="L25"></a>FuncType *FuncType;
<a id="L26"></a>}

<a id="L28"></a><span class="comment">// A Prog collects information about a cgo program.</span>
<a id="L29"></a>type Prog struct {
    <a id="L30"></a>AST         *ast.File; <span class="comment">// parsed AST</span>
    <a id="L31"></a>Preamble    string;    <span class="comment">// C preamble (doc comment on import &#34;C&#34;)</span>
    <a id="L32"></a>PackagePath string;
    <a id="L33"></a>Package     string;
    <a id="L34"></a>Crefs       []*Cref;
    <a id="L35"></a>Typedef     map[string]ast.Expr;
    <a id="L36"></a>Vardef      map[string]*Type;
    <a id="L37"></a>Funcdef     map[string]*FuncType;
    <a id="L38"></a>PtrSize     int64;
    <a id="L39"></a>GccOptions  []string;
<a id="L40"></a>}

<a id="L42"></a><span class="comment">// A Type collects information about a type in both the C and Go worlds.</span>
<a id="L43"></a>type Type struct {
    <a id="L44"></a>Size  int64;
    <a id="L45"></a>Align int64;
    <a id="L46"></a>C     string;
    <a id="L47"></a>Go    ast.Expr;
<a id="L48"></a>}

<a id="L50"></a><span class="comment">// A FuncType collects information about a function type in both the C and Go worlds.</span>
<a id="L51"></a>type FuncType struct {
    <a id="L52"></a>Params []*Type;
    <a id="L53"></a>Result *Type;
    <a id="L54"></a>Go     *ast.FuncType;
<a id="L55"></a>}

<a id="L57"></a>func openProg(name string) *Prog {
    <a id="L58"></a>p := new(Prog);
    <a id="L59"></a>var err os.Error;
    <a id="L60"></a>p.AST, err = parser.ParsePkgFile(&#34;&#34;, name, parser.ParseComments);
    <a id="L61"></a>if err != nil {
        <a id="L62"></a>if list, ok := err.(scanner.ErrorList); ok {
            <a id="L63"></a><span class="comment">// If err is a scanner.ErrorList, its String will print just</span>
            <a id="L64"></a><span class="comment">// the first error and then (+n more errors).</span>
            <a id="L65"></a><span class="comment">// Instead, turn it into a new Error that will return</span>
            <a id="L66"></a><span class="comment">// details for all the errors.</span>
            <a id="L67"></a>for _, e := range list {
                <a id="L68"></a>fmt.Fprintln(os.Stderr, e)
            <a id="L69"></a>}
            <a id="L70"></a>os.Exit(2);
        <a id="L71"></a>}
        <a id="L72"></a>fatal(&#34;parsing %s: %s&#34;, name, err);
    <a id="L73"></a>}
    <a id="L74"></a>p.Package = p.AST.Name.Value;

    <a id="L76"></a><span class="comment">// Find the import &#34;C&#34; line and get any extra C preamble.</span>
    <a id="L77"></a><span class="comment">// Delete the import &#34;C&#34; line along the way.</span>
    <a id="L78"></a>sawC := false;
    <a id="L79"></a>w := 0;
    <a id="L80"></a>for _, decl := range p.AST.Decls {
        <a id="L81"></a>d, ok := decl.(*ast.GenDecl);
        <a id="L82"></a>if !ok {
            <a id="L83"></a>p.AST.Decls[w] = decl;
            <a id="L84"></a>w++;
            <a id="L85"></a>continue;
        <a id="L86"></a>}
        <a id="L87"></a>ws := 0;
        <a id="L88"></a>for _, spec := range d.Specs {
            <a id="L89"></a>s, ok := spec.(*ast.ImportSpec);
            <a id="L90"></a>if !ok || len(s.Path) != 1 || string(s.Path[0].Value) != `&#34;C&#34;` {
                <a id="L91"></a>d.Specs[ws] = spec;
                <a id="L92"></a>ws++;
                <a id="L93"></a>continue;
            <a id="L94"></a>}
            <a id="L95"></a>sawC = true;
            <a id="L96"></a>if s.Name != nil {
                <a id="L97"></a>error(s.Path[0].Pos(), `cannot rename import &#34;C&#34;`)
            <a id="L98"></a>}
            <a id="L99"></a>if s.Doc != nil {
                <a id="L100"></a>p.Preamble += doc.CommentText(s.Doc) + &#34;\n&#34;
            <a id="L101"></a>} else if len(d.Specs) == 1 &amp;&amp; d.Doc != nil {
                <a id="L102"></a>p.Preamble += doc.CommentText(d.Doc) + &#34;\n&#34;
            <a id="L103"></a>}
        <a id="L104"></a>}
        <a id="L105"></a>if ws == 0 {
            <a id="L106"></a>continue
        <a id="L107"></a>}
        <a id="L108"></a>d.Specs = d.Specs[0:ws];
        <a id="L109"></a>p.AST.Decls[w] = d;
        <a id="L110"></a>w++;
    <a id="L111"></a>}
    <a id="L112"></a>p.AST.Decls = p.AST.Decls[0:w];

    <a id="L114"></a>if !sawC {
        <a id="L115"></a>error(noPos, `cannot find import &#34;C&#34;`)
    <a id="L116"></a>}

    <a id="L118"></a><span class="comment">// Accumulate pointers to uses of C.x.</span>
    <a id="L119"></a>p.Crefs = make([]*Cref, 0, 8);
    <a id="L120"></a>walk(p.AST, p, &#34;prog&#34;);
    <a id="L121"></a>return p;
<a id="L122"></a>}

<a id="L124"></a>func walk(x interface{}, p *Prog, context string) {
    <a id="L125"></a>switch n := x.(type) {
    <a id="L126"></a>case *ast.Expr:
        <a id="L127"></a>if sel, ok := (*n).(*ast.SelectorExpr); ok {
            <a id="L128"></a><span class="comment">// For now, assume that the only instance of capital C is</span>
            <a id="L129"></a><span class="comment">// when used as the imported package identifier.</span>
            <a id="L130"></a><span class="comment">// The parser should take care of scoping in the future,</span>
            <a id="L131"></a><span class="comment">// so that we will be able to distinguish a &#34;top-level C&#34;</span>
            <a id="L132"></a><span class="comment">// from a local C.</span>
            <a id="L133"></a>if l, ok := sel.X.(*ast.Ident); ok &amp;&amp; l.Value == &#34;C&#34; {
                <a id="L134"></a>i := len(p.Crefs);
                <a id="L135"></a>if i &gt;= cap(p.Crefs) {
                    <a id="L136"></a>new := make([]*Cref, 2*i);
                    <a id="L137"></a>for j, v := range p.Crefs {
                        <a id="L138"></a>new[j] = v
                    <a id="L139"></a>}
                    <a id="L140"></a>p.Crefs = new;
                <a id="L141"></a>}
                <a id="L142"></a>p.Crefs = p.Crefs[0 : i+1];
                <a id="L143"></a>p.Crefs[i] = &amp;Cref{
                    <a id="L144"></a>Name: sel.Sel.Value,
                    <a id="L145"></a>Expr: n,
                    <a id="L146"></a>Context: context,
                <a id="L147"></a>};
                <a id="L148"></a>break;
            <a id="L149"></a>}
        <a id="L150"></a>}
        <a id="L151"></a>walk(*n, p, context);

    <a id="L153"></a><span class="comment">// everything else just recurs</span>
    <a id="L154"></a>default:
        <a id="L155"></a>error(noPos, &#34;unexpected type %T in walk&#34;, x);
        <a id="L156"></a>panic();

    <a id="L158"></a>case nil:

    <a id="L160"></a><span class="comment">// These are ordered and grouped to match ../../pkg/go/ast/ast.go</span>
    <a id="L161"></a>case *ast.Field:
        <a id="L162"></a>walk(&amp;n.Type, p, &#34;type&#34;)
    <a id="L163"></a>case *ast.BadExpr:
    <a id="L164"></a>case *ast.Ident:
    <a id="L165"></a>case *ast.Ellipsis:
    <a id="L166"></a>case *ast.BasicLit:
    <a id="L167"></a>case *ast.StringList:
    <a id="L168"></a>case *ast.FuncLit:
        <a id="L169"></a>walk(n.Type, p, &#34;type&#34;);
        <a id="L170"></a>walk(n.Body, p, &#34;stmt&#34;);
    <a id="L171"></a>case *ast.CompositeLit:
        <a id="L172"></a>walk(&amp;n.Type, p, &#34;type&#34;);
        <a id="L173"></a>walk(n.Elts, p, &#34;expr&#34;);
    <a id="L174"></a>case *ast.ParenExpr:
        <a id="L175"></a>walk(&amp;n.X, p, context)
    <a id="L176"></a>case *ast.SelectorExpr:
        <a id="L177"></a>walk(&amp;n.X, p, &#34;selector&#34;)
    <a id="L178"></a>case *ast.IndexExpr:
        <a id="L179"></a>walk(&amp;n.X, p, &#34;expr&#34;);
        <a id="L180"></a>walk(&amp;n.Index, p, &#34;expr&#34;);
        <a id="L181"></a>if n.End != nil {
            <a id="L182"></a>walk(&amp;n.End, p, &#34;expr&#34;)
        <a id="L183"></a>}
    <a id="L184"></a>case *ast.TypeAssertExpr:
        <a id="L185"></a>walk(&amp;n.X, p, &#34;expr&#34;);
        <a id="L186"></a>walk(&amp;n.Type, p, &#34;type&#34;);
    <a id="L187"></a>case *ast.CallExpr:
        <a id="L188"></a>walk(&amp;n.Fun, p, &#34;call&#34;);
        <a id="L189"></a>walk(n.Args, p, &#34;expr&#34;);
    <a id="L190"></a>case *ast.StarExpr:
        <a id="L191"></a>walk(&amp;n.X, p, context)
    <a id="L192"></a>case *ast.UnaryExpr:
        <a id="L193"></a>walk(&amp;n.X, p, &#34;expr&#34;)
    <a id="L194"></a>case *ast.BinaryExpr:
        <a id="L195"></a>walk(&amp;n.X, p, &#34;expr&#34;);
        <a id="L196"></a>walk(&amp;n.Y, p, &#34;expr&#34;);
    <a id="L197"></a>case *ast.KeyValueExpr:
        <a id="L198"></a>walk(&amp;n.Key, p, &#34;expr&#34;);
        <a id="L199"></a>walk(&amp;n.Value, p, &#34;expr&#34;);

    <a id="L201"></a>case *ast.ArrayType:
        <a id="L202"></a>walk(&amp;n.Len, p, &#34;expr&#34;);
        <a id="L203"></a>walk(&amp;n.Elt, p, &#34;type&#34;);
    <a id="L204"></a>case *ast.StructType:
        <a id="L205"></a>walk(n.Fields, p, &#34;field&#34;)
    <a id="L206"></a>case *ast.FuncType:
        <a id="L207"></a>walk(n.Params, p, &#34;field&#34;);
        <a id="L208"></a>walk(n.Results, p, &#34;field&#34;);
    <a id="L209"></a>case *ast.InterfaceType:
        <a id="L210"></a>walk(n.Methods, p, &#34;field&#34;)
    <a id="L211"></a>case *ast.MapType:
        <a id="L212"></a>walk(&amp;n.Key, p, &#34;type&#34;);
        <a id="L213"></a>walk(&amp;n.Value, p, &#34;type&#34;);
    <a id="L214"></a>case *ast.ChanType:
        <a id="L215"></a>walk(&amp;n.Value, p, &#34;type&#34;)

    <a id="L217"></a>case *ast.BadStmt:
    <a id="L218"></a>case *ast.DeclStmt:
        <a id="L219"></a>walk(n.Decl, p, &#34;decl&#34;)
    <a id="L220"></a>case *ast.EmptyStmt:
    <a id="L221"></a>case *ast.LabeledStmt:
        <a id="L222"></a>walk(n.Stmt, p, &#34;stmt&#34;)
    <a id="L223"></a>case *ast.ExprStmt:
        <a id="L224"></a>walk(&amp;n.X, p, &#34;expr&#34;)
    <a id="L225"></a>case *ast.IncDecStmt:
        <a id="L226"></a>walk(&amp;n.X, p, &#34;expr&#34;)
    <a id="L227"></a>case *ast.AssignStmt:
        <a id="L228"></a>walk(n.Lhs, p, &#34;expr&#34;);
        <a id="L229"></a>walk(n.Rhs, p, &#34;expr&#34;);
    <a id="L230"></a>case *ast.GoStmt:
        <a id="L231"></a>walk(n.Call, p, &#34;expr&#34;)
    <a id="L232"></a>case *ast.DeferStmt:
        <a id="L233"></a>walk(n.Call, p, &#34;expr&#34;)
    <a id="L234"></a>case *ast.ReturnStmt:
        <a id="L235"></a>walk(n.Results, p, &#34;expr&#34;)
    <a id="L236"></a>case *ast.BranchStmt:
    <a id="L237"></a>case *ast.BlockStmt:
        <a id="L238"></a>walk(n.List, p, &#34;stmt&#34;)
    <a id="L239"></a>case *ast.IfStmt:
        <a id="L240"></a>walk(n.Init, p, &#34;stmt&#34;);
        <a id="L241"></a>walk(&amp;n.Cond, p, &#34;expr&#34;);
        <a id="L242"></a>walk(n.Body, p, &#34;stmt&#34;);
        <a id="L243"></a>walk(n.Else, p, &#34;stmt&#34;);
    <a id="L244"></a>case *ast.CaseClause:
        <a id="L245"></a>walk(n.Values, p, &#34;expr&#34;);
        <a id="L246"></a>walk(n.Body, p, &#34;stmt&#34;);
    <a id="L247"></a>case *ast.SwitchStmt:
        <a id="L248"></a>walk(n.Init, p, &#34;stmt&#34;);
        <a id="L249"></a>walk(&amp;n.Tag, p, &#34;expr&#34;);
        <a id="L250"></a>walk(n.Body, p, &#34;stmt&#34;);
    <a id="L251"></a>case *ast.TypeCaseClause:
        <a id="L252"></a>walk(n.Types, p, &#34;type&#34;);
        <a id="L253"></a>walk(n.Body, p, &#34;stmt&#34;);
    <a id="L254"></a>case *ast.TypeSwitchStmt:
        <a id="L255"></a>walk(n.Init, p, &#34;stmt&#34;);
        <a id="L256"></a>walk(n.Assign, p, &#34;stmt&#34;);
        <a id="L257"></a>walk(n.Body, p, &#34;stmt&#34;);
    <a id="L258"></a>case *ast.CommClause:
        <a id="L259"></a>walk(n.Lhs, p, &#34;expr&#34;);
        <a id="L260"></a>walk(n.Rhs, p, &#34;expr&#34;);
        <a id="L261"></a>walk(n.Body, p, &#34;stmt&#34;);
    <a id="L262"></a>case *ast.SelectStmt:
        <a id="L263"></a>walk(n.Body, p, &#34;stmt&#34;)
    <a id="L264"></a>case *ast.ForStmt:
        <a id="L265"></a>walk(n.Init, p, &#34;stmt&#34;);
        <a id="L266"></a>walk(&amp;n.Cond, p, &#34;expr&#34;);
        <a id="L267"></a>walk(n.Post, p, &#34;stmt&#34;);
        <a id="L268"></a>walk(n.Body, p, &#34;stmt&#34;);
    <a id="L269"></a>case *ast.RangeStmt:
        <a id="L270"></a>walk(&amp;n.Key, p, &#34;expr&#34;);
        <a id="L271"></a>walk(&amp;n.Value, p, &#34;expr&#34;);
        <a id="L272"></a>walk(&amp;n.X, p, &#34;expr&#34;);
        <a id="L273"></a>walk(n.Body, p, &#34;stmt&#34;);

    <a id="L275"></a>case *ast.ImportSpec:
    <a id="L276"></a>case *ast.ValueSpec:
        <a id="L277"></a>walk(&amp;n.Type, p, &#34;type&#34;);
        <a id="L278"></a>walk(n.Values, p, &#34;expr&#34;);
    <a id="L279"></a>case *ast.TypeSpec:
        <a id="L280"></a>walk(&amp;n.Type, p, &#34;type&#34;)

    <a id="L282"></a>case *ast.BadDecl:
    <a id="L283"></a>case *ast.GenDecl:
        <a id="L284"></a>walk(n.Specs, p, &#34;spec&#34;)
    <a id="L285"></a>case *ast.FuncDecl:
        <a id="L286"></a>if n.Recv != nil {
            <a id="L287"></a>walk(n.Recv, p, &#34;field&#34;)
        <a id="L288"></a>}
        <a id="L289"></a>walk(n.Type, p, &#34;type&#34;);
        <a id="L290"></a>if n.Body != nil {
            <a id="L291"></a>walk(n.Body, p, &#34;stmt&#34;)
        <a id="L292"></a>}

    <a id="L294"></a>case *ast.File:
        <a id="L295"></a>walk(n.Decls, p, &#34;decl&#34;)

    <a id="L297"></a>case *ast.Package:
        <a id="L298"></a>for _, f := range n.Files {
            <a id="L299"></a>walk(f, p, &#34;file&#34;)
        <a id="L300"></a>}

    <a id="L302"></a>case []ast.Decl:
        <a id="L303"></a>for _, d := range n {
            <a id="L304"></a>walk(d, p, context)
        <a id="L305"></a>}
    <a id="L306"></a>case []ast.Expr:
        <a id="L307"></a>for i := range n {
            <a id="L308"></a>walk(&amp;n[i], p, context)
        <a id="L309"></a>}
    <a id="L310"></a>case []*ast.Field:
        <a id="L311"></a>for _, f := range n {
            <a id="L312"></a>walk(f, p, context)
        <a id="L313"></a>}
    <a id="L314"></a>case []ast.Stmt:
        <a id="L315"></a>for _, s := range n {
            <a id="L316"></a>walk(s, p, context)
        <a id="L317"></a>}
    <a id="L318"></a>case []ast.Spec:
        <a id="L319"></a>for _, s := range n {
            <a id="L320"></a>walk(s, p, context)
        <a id="L321"></a>}
    <a id="L322"></a>}
<a id="L323"></a>}
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
