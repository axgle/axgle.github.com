<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/ast/filter.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/go/ast/filter.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package ast

<a id="L7"></a>import &#34;go/token&#34;


<a id="L10"></a>func filterIdentList(list []*Ident) []*Ident {
    <a id="L11"></a>j := 0;
    <a id="L12"></a>for _, x := range list {
        <a id="L13"></a>if x.IsExported() {
            <a id="L14"></a>list[j] = x;
            <a id="L15"></a>j++;
        <a id="L16"></a>}
    <a id="L17"></a>}
    <a id="L18"></a>return list[0:j];
<a id="L19"></a>}


<a id="L22"></a><span class="comment">// isExportedType assumes that typ is a correct type.</span>
<a id="L23"></a>func isExportedType(typ Expr) bool {
    <a id="L24"></a>switch t := typ.(type) {
    <a id="L25"></a>case *Ident:
        <a id="L26"></a>return t.IsExported()
    <a id="L27"></a>case *ParenExpr:
        <a id="L28"></a>return isExportedType(t.X)
    <a id="L29"></a>case *SelectorExpr:
        <a id="L30"></a><span class="comment">// assume t.X is a typename</span>
        <a id="L31"></a>return t.Sel.IsExported()
    <a id="L32"></a>case *StarExpr:
        <a id="L33"></a>return isExportedType(t.X)
    <a id="L34"></a>}
    <a id="L35"></a>return false;
<a id="L36"></a>}


<a id="L39"></a>func filterFieldList(list []*Field, incomplete *bool) []*Field {
    <a id="L40"></a>j := 0;
    <a id="L41"></a>for _, f := range list {
        <a id="L42"></a>exported := false;
        <a id="L43"></a>if len(f.Names) == 0 {
            <a id="L44"></a><span class="comment">// anonymous field</span>
            <a id="L45"></a><span class="comment">// (Note that a non-exported anonymous field</span>
            <a id="L46"></a><span class="comment">// may still refer to a type with exported</span>
            <a id="L47"></a><span class="comment">// fields, so this is not absolutely correct.</span>
            <a id="L48"></a><span class="comment">// However, this cannot be done w/o complete</span>
            <a id="L49"></a><span class="comment">// type information.)</span>
            <a id="L50"></a>exported = isExportedType(f.Type)
        <a id="L51"></a>} else {
            <a id="L52"></a>n := len(f.Names);
            <a id="L53"></a>f.Names = filterIdentList(f.Names);
            <a id="L54"></a>if len(f.Names) &lt; n {
                <a id="L55"></a>*incomplete = true
            <a id="L56"></a>}
            <a id="L57"></a>exported = len(f.Names) &gt; 0;
        <a id="L58"></a>}
        <a id="L59"></a>if exported {
            <a id="L60"></a>filterType(f.Type);
            <a id="L61"></a>list[j] = f;
            <a id="L62"></a>j++;
        <a id="L63"></a>}
    <a id="L64"></a>}
    <a id="L65"></a>if j &lt; len(list) {
        <a id="L66"></a>*incomplete = true
    <a id="L67"></a>}
    <a id="L68"></a>return list[0:j];
<a id="L69"></a>}


<a id="L72"></a>func filterParamList(list []*Field) {
    <a id="L73"></a>for _, f := range list {
        <a id="L74"></a>filterType(f.Type)
    <a id="L75"></a>}
<a id="L76"></a>}


<a id="L79"></a>var noPos token.Position

<a id="L81"></a>func filterType(typ Expr) {
    <a id="L82"></a>switch t := typ.(type) {
    <a id="L83"></a>case *ArrayType:
        <a id="L84"></a>filterType(t.Elt)
    <a id="L85"></a>case *StructType:
        <a id="L86"></a>t.Fields = filterFieldList(t.Fields, &amp;t.Incomplete)
    <a id="L87"></a>case *FuncType:
        <a id="L88"></a>filterParamList(t.Params);
        <a id="L89"></a>filterParamList(t.Results);
    <a id="L90"></a>case *InterfaceType:
        <a id="L91"></a>t.Methods = filterFieldList(t.Methods, &amp;t.Incomplete)
    <a id="L92"></a>case *MapType:
        <a id="L93"></a>filterType(t.Key);
        <a id="L94"></a>filterType(t.Value);
    <a id="L95"></a>case *ChanType:
        <a id="L96"></a>filterType(t.Value)
    <a id="L97"></a>}
<a id="L98"></a>}


<a id="L101"></a>func filterSpec(spec Spec) bool {
    <a id="L102"></a>switch s := spec.(type) {
    <a id="L103"></a>case *ValueSpec:
        <a id="L104"></a>s.Names = filterIdentList(s.Names);
        <a id="L105"></a>if len(s.Names) &gt; 0 {
            <a id="L106"></a>filterType(s.Type);
            <a id="L107"></a>return true;
        <a id="L108"></a>}
    <a id="L109"></a>case *TypeSpec:
        <a id="L110"></a><span class="comment">// TODO(gri) consider stripping forward declarations</span>
        <a id="L111"></a><span class="comment">//           of structs, interfaces, functions, and methods</span>
        <a id="L112"></a>if s.Name.IsExported() {
            <a id="L113"></a>filterType(s.Type);
            <a id="L114"></a>return true;
        <a id="L115"></a>}
    <a id="L116"></a>}
    <a id="L117"></a>return false;
<a id="L118"></a>}


<a id="L121"></a>func filterSpecList(list []Spec) []Spec {
    <a id="L122"></a>j := 0;
    <a id="L123"></a>for _, s := range list {
        <a id="L124"></a>if filterSpec(s) {
            <a id="L125"></a>list[j] = s;
            <a id="L126"></a>j++;
        <a id="L127"></a>}
    <a id="L128"></a>}
    <a id="L129"></a>return list[0:j];
<a id="L130"></a>}


<a id="L133"></a>func filterDecl(decl Decl) bool {
    <a id="L134"></a>switch d := decl.(type) {
    <a id="L135"></a>case *GenDecl:
        <a id="L136"></a>d.Specs = filterSpecList(d.Specs);
        <a id="L137"></a>return len(d.Specs) &gt; 0;
    <a id="L138"></a>case *FuncDecl:
        <a id="L139"></a><span class="comment">// TODO consider removing function declaration altogether if</span>
        <a id="L140"></a><span class="comment">//      forward declaration (i.e., if d.Body == nil) because</span>
        <a id="L141"></a><span class="comment">//      in that case the actual declaration will come later.</span>
        <a id="L142"></a>d.Body = nil; <span class="comment">// strip body</span>
        <a id="L143"></a>return d.Name.IsExported();
    <a id="L144"></a>}
    <a id="L145"></a>return false;
<a id="L146"></a>}


<a id="L149"></a><span class="comment">// FileExports trims the AST for a Go source file in place such that only</span>
<a id="L150"></a><span class="comment">// exported nodes remain: all top-level identifiers which are not exported</span>
<a id="L151"></a><span class="comment">// and their associated information (such as type, initial value, or function</span>
<a id="L152"></a><span class="comment">// body) are removed. Non-exported fields and methods of exported types are</span>
<a id="L153"></a><span class="comment">// stripped, and the function bodies of exported functions are set to nil.</span>
<a id="L154"></a><span class="comment">// The File.comments list is not changed.</span>
<a id="L155"></a><span class="comment">//</span>
<a id="L156"></a><span class="comment">// FileExports returns true if there is an exported declaration; it returns</span>
<a id="L157"></a><span class="comment">// false otherwise.</span>
<a id="L158"></a><span class="comment">//</span>
<a id="L159"></a>func FileExports(src *File) bool {
    <a id="L160"></a>j := 0;
    <a id="L161"></a>for _, d := range src.Decls {
        <a id="L162"></a>if filterDecl(d) {
            <a id="L163"></a>src.Decls[j] = d;
            <a id="L164"></a>j++;
        <a id="L165"></a>}
    <a id="L166"></a>}
    <a id="L167"></a>src.Decls = src.Decls[0:j];
    <a id="L168"></a>return j &gt; 0;
<a id="L169"></a>}


<a id="L172"></a><span class="comment">// PackageExports trims the AST for a Go package in place such that only</span>
<a id="L173"></a><span class="comment">// exported nodes remain. The pkg.Files list is not changed, so that file</span>
<a id="L174"></a><span class="comment">// names and top-level package comments don&#39;t get lost.</span>
<a id="L175"></a><span class="comment">//</span>
<a id="L176"></a><span class="comment">// PackageExports returns true if there is an exported declaration; it</span>
<a id="L177"></a><span class="comment">// returns false otherwise.</span>
<a id="L178"></a><span class="comment">//</span>
<a id="L179"></a>func PackageExports(pkg *Package) bool {
    <a id="L180"></a>hasExports := false;
    <a id="L181"></a>for _, f := range pkg.Files {
        <a id="L182"></a>if FileExports(f) {
            <a id="L183"></a>hasExports = true
        <a id="L184"></a>}
    <a id="L185"></a>}
    <a id="L186"></a>return hasExports;
<a id="L187"></a>}


<a id="L190"></a><span class="comment">// separator is an empty //-style comment that is interspersed between</span>
<a id="L191"></a><span class="comment">// different comment groups when they are concatenated into a single group</span>
<a id="L192"></a><span class="comment">//</span>
<a id="L193"></a>var separator = &amp;Comment{noPos, []byte{&#39;/&#39;, &#39;/&#39;}}


<a id="L196"></a><span class="comment">// MergePackageFiles creates a file AST by merging the ASTs of the</span>
<a id="L197"></a><span class="comment">// files belonging to a package.</span>
<a id="L198"></a><span class="comment">//</span>
<a id="L199"></a>func MergePackageFiles(pkg *Package) *File {
    <a id="L200"></a><span class="comment">// Count the number of package comments and declarations across</span>
    <a id="L201"></a><span class="comment">// all package files.</span>
    <a id="L202"></a>ncomments := 0;
    <a id="L203"></a>ndecls := 0;
    <a id="L204"></a>for _, f := range pkg.Files {
        <a id="L205"></a>if f.Doc != nil {
            <a id="L206"></a>ncomments += len(f.Doc.List) + 1 <span class="comment">// +1 for separator</span>
        <a id="L207"></a>}
        <a id="L208"></a>ndecls += len(f.Decls);
    <a id="L209"></a>}

    <a id="L211"></a><span class="comment">// Collect package comments from all package files into a single</span>
    <a id="L212"></a><span class="comment">// CommentGroup - the collected package documentation. The order</span>
    <a id="L213"></a><span class="comment">// is unspecified. In general there should be only one file with</span>
    <a id="L214"></a><span class="comment">// a package comment; but it&#39;s better to collect extra comments</span>
    <a id="L215"></a><span class="comment">// than drop them on the floor.</span>
    <a id="L216"></a>var doc *CommentGroup;
    <a id="L217"></a>if ncomments &gt; 0 {
        <a id="L218"></a>list := make([]*Comment, ncomments-1); <span class="comment">// -1: no separator before first group</span>
        <a id="L219"></a>i := 0;
        <a id="L220"></a>for _, f := range pkg.Files {
            <a id="L221"></a>if f.Doc != nil {
                <a id="L222"></a>if i &gt; 0 {
                    <a id="L223"></a><span class="comment">// not the first group - add separator</span>
                    <a id="L224"></a>list[i] = separator;
                    <a id="L225"></a>i++;
                <a id="L226"></a>}
                <a id="L227"></a>for _, c := range f.Doc.List {
                    <a id="L228"></a>list[i] = c;
                    <a id="L229"></a>i++;
                <a id="L230"></a>}
            <a id="L231"></a>}
        <a id="L232"></a>}
        <a id="L233"></a>doc = &amp;CommentGroup{list, nil};
    <a id="L234"></a>}

    <a id="L236"></a><span class="comment">// Collect declarations from all package files.</span>
    <a id="L237"></a>var decls []Decl;
    <a id="L238"></a>if ndecls &gt; 0 {
        <a id="L239"></a>decls = make([]Decl, ndecls);
        <a id="L240"></a>i := 0;
        <a id="L241"></a>for _, f := range pkg.Files {
            <a id="L242"></a>for _, d := range f.Decls {
                <a id="L243"></a>decls[i] = d;
                <a id="L244"></a>i++;
            <a id="L245"></a>}
        <a id="L246"></a>}
    <a id="L247"></a>}

    <a id="L249"></a><span class="comment">// TODO(gri) Should collect comments as well. For that the comment</span>
    <a id="L250"></a><span class="comment">//           list should be changed back into a []*CommentGroup,</span>
    <a id="L251"></a><span class="comment">//           otherwise need to modify the existing linked list.</span>
    <a id="L252"></a>return &amp;File{doc, noPos, &amp;Ident{noPos, pkg.Name}, decls, nil};
<a id="L253"></a>}
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
