<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/doc/doc.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/go/doc/doc.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The doc package extracts source code documentation from a Go AST.</span>
<a id="L6"></a>package doc

<a id="L8"></a>import (
    <a id="L9"></a>&#34;container/vector&#34;;
    <a id="L10"></a>&#34;go/ast&#34;;
    <a id="L11"></a>&#34;go/token&#34;;
    <a id="L12"></a>&#34;regexp&#34;;
    <a id="L13"></a>&#34;sort&#34;;
<a id="L14"></a>)


<a id="L17"></a><span class="comment">// ----------------------------------------------------------------------------</span>

<a id="L19"></a>type typeDoc struct {
    <a id="L20"></a><span class="comment">// len(decl.Specs) == 1, and the element type is *ast.TypeSpec</span>
    <a id="L21"></a><span class="comment">// if the type declaration hasn&#39;t been seen yet, decl is nil</span>
    <a id="L22"></a>decl *ast.GenDecl;
    <a id="L23"></a><span class="comment">// values, factory functions, and methods associated with the type</span>
    <a id="L24"></a>values    *vector.Vector; <span class="comment">// list of *ast.GenDecl (consts and vars)</span>
    <a id="L25"></a>factories map[string]*ast.FuncDecl;
    <a id="L26"></a>methods   map[string]*ast.FuncDecl;
<a id="L27"></a>}


<a id="L30"></a><span class="comment">// docReader accumulates documentation for a single package.</span>
<a id="L31"></a><span class="comment">// It modifies the AST: Comments (declaration documentation)</span>
<a id="L32"></a><span class="comment">// that have been collected by the DocReader are set to nil</span>
<a id="L33"></a><span class="comment">// in the respective AST nodes so that they are not printed</span>
<a id="L34"></a><span class="comment">// twice (once when printing the documentation and once when</span>
<a id="L35"></a><span class="comment">// printing the corresponding AST node).</span>
<a id="L36"></a><span class="comment">//</span>
<a id="L37"></a>type docReader struct {
    <a id="L38"></a>doc     *ast.CommentGroup; <span class="comment">// package documentation, if any</span>
    <a id="L39"></a>pkgName string;
    <a id="L40"></a>values  *vector.Vector; <span class="comment">// list of *ast.GenDecl (consts and vars)</span>
    <a id="L41"></a>types   map[string]*typeDoc;
    <a id="L42"></a>funcs   map[string]*ast.FuncDecl;
    <a id="L43"></a>bugs    *vector.Vector; <span class="comment">// list of *ast.CommentGroup</span>
<a id="L44"></a>}


<a id="L47"></a>func (doc *docReader) init(pkgName string) {
    <a id="L48"></a>doc.pkgName = pkgName;
    <a id="L49"></a>doc.values = vector.New(0);
    <a id="L50"></a>doc.types = make(map[string]*typeDoc);
    <a id="L51"></a>doc.funcs = make(map[string]*ast.FuncDecl);
    <a id="L52"></a>doc.bugs = vector.New(0);
<a id="L53"></a>}


<a id="L56"></a>func (doc *docReader) addType(decl *ast.GenDecl) {
    <a id="L57"></a>spec := decl.Specs[0].(*ast.TypeSpec);
    <a id="L58"></a>typ := doc.lookupTypeDoc(spec.Name.Value);
    <a id="L59"></a><span class="comment">// typ should always be != nil since declared types</span>
    <a id="L60"></a><span class="comment">// are always named - be conservative and check</span>
    <a id="L61"></a>if typ != nil {
        <a id="L62"></a><span class="comment">// a type should be added at most once, so typ.decl</span>
        <a id="L63"></a><span class="comment">// should be nil - if it isn&#39;t, simply overwrite it</span>
        <a id="L64"></a>typ.decl = decl
    <a id="L65"></a>}
<a id="L66"></a>}


<a id="L69"></a>func (doc *docReader) lookupTypeDoc(name string) *typeDoc {
    <a id="L70"></a>if name == &#34;&#34; {
        <a id="L71"></a>return nil <span class="comment">// no type docs for anonymous types</span>
    <a id="L72"></a>}
    <a id="L73"></a>if tdoc, found := doc.types[name]; found {
        <a id="L74"></a>return tdoc
    <a id="L75"></a>}
    <a id="L76"></a><span class="comment">// type wasn&#39;t found - add one without declaration</span>
    <a id="L77"></a>tdoc := &amp;typeDoc{nil, vector.New(0), make(map[string]*ast.FuncDecl), make(map[string]*ast.FuncDecl)};
    <a id="L78"></a>doc.types[name] = tdoc;
    <a id="L79"></a>return tdoc;
<a id="L80"></a>}


<a id="L83"></a>func baseTypeName(typ ast.Expr) string {
    <a id="L84"></a>switch t := typ.(type) {
    <a id="L85"></a>case *ast.Ident:
        <a id="L86"></a><span class="comment">// if the type is not exported, the effect to</span>
        <a id="L87"></a><span class="comment">// a client is as if there were no type name</span>
        <a id="L88"></a>if t.IsExported() {
            <a id="L89"></a>return string(t.Value)
        <a id="L90"></a>}
    <a id="L91"></a>case *ast.StarExpr:
        <a id="L92"></a>return baseTypeName(t.X)
    <a id="L93"></a>}
    <a id="L94"></a>return &#34;&#34;;
<a id="L95"></a>}


<a id="L98"></a>func (doc *docReader) addValue(decl *ast.GenDecl) {
    <a id="L99"></a><span class="comment">// determine if decl should be associated with a type</span>
    <a id="L100"></a><span class="comment">// Heuristic: For each typed entry, determine the type name, if any.</span>
    <a id="L101"></a><span class="comment">//            If there is exactly one type name that is sufficiently</span>
    <a id="L102"></a><span class="comment">//            frequent, associate the decl with the respective type.</span>
    <a id="L103"></a>domName := &#34;&#34;;
    <a id="L104"></a>domFreq := 0;
    <a id="L105"></a>prev := &#34;&#34;;
    <a id="L106"></a>for _, s := range decl.Specs {
        <a id="L107"></a>if v, ok := s.(*ast.ValueSpec); ok {
            <a id="L108"></a>name := &#34;&#34;;
            <a id="L109"></a>switch {
            <a id="L110"></a>case v.Type != nil:
                <a id="L111"></a><span class="comment">// a type is present; determine it&#39;s name</span>
                <a id="L112"></a>name = baseTypeName(v.Type)
            <a id="L113"></a>case decl.Tok == token.CONST:
                <a id="L114"></a><span class="comment">// no type is present but we have a constant declaration;</span>
                <a id="L115"></a><span class="comment">// use the previous type name (w/o more type information</span>
                <a id="L116"></a><span class="comment">// we cannot handle the case of unnamed variables with</span>
                <a id="L117"></a><span class="comment">// initializer expressions except for some trivial cases)</span>
                <a id="L118"></a>name = prev
            <a id="L119"></a>}
            <a id="L120"></a>if name != &#34;&#34; {
                <a id="L121"></a><span class="comment">// entry has a named type</span>
                <a id="L122"></a>if domName != &#34;&#34; &amp;&amp; domName != name {
                    <a id="L123"></a><span class="comment">// more than one type name - do not associate</span>
                    <a id="L124"></a><span class="comment">// with any type</span>
                    <a id="L125"></a>domName = &#34;&#34;;
                    <a id="L126"></a>break;
                <a id="L127"></a>}
                <a id="L128"></a>domName = name;
                <a id="L129"></a>domFreq++;
            <a id="L130"></a>}
            <a id="L131"></a>prev = name;
        <a id="L132"></a>}
    <a id="L133"></a>}

    <a id="L135"></a><span class="comment">// determine values list</span>
    <a id="L136"></a>const threshold = 0.75;
    <a id="L137"></a>values := doc.values;
    <a id="L138"></a>if domName != &#34;&#34; &amp;&amp; domFreq &gt;= int(float(len(decl.Specs))*threshold) {
        <a id="L139"></a><span class="comment">// typed entries are sufficiently frequent</span>
        <a id="L140"></a>typ := doc.lookupTypeDoc(domName);
        <a id="L141"></a>if typ != nil {
            <a id="L142"></a>values = typ.values <span class="comment">// associate with that type</span>
        <a id="L143"></a>}
    <a id="L144"></a>}

    <a id="L146"></a>values.Push(decl);
<a id="L147"></a>}


<a id="L150"></a>func (doc *docReader) addFunc(fun *ast.FuncDecl) {
    <a id="L151"></a>name := fun.Name.Value;

    <a id="L153"></a><span class="comment">// determine if it should be associated with a type</span>
    <a id="L154"></a>if fun.Recv != nil {
        <a id="L155"></a><span class="comment">// method</span>
        <a id="L156"></a>typ := doc.lookupTypeDoc(baseTypeName(fun.Recv.Type));
        <a id="L157"></a>if typ != nil {
            <a id="L158"></a><span class="comment">// exported receiver type</span>
            <a id="L159"></a>typ.methods[name] = fun
        <a id="L160"></a>}
        <a id="L161"></a><span class="comment">// otherwise don&#39;t show the method</span>
        <a id="L162"></a><span class="comment">// TODO(gri): There may be exported methods of non-exported types</span>
        <a id="L163"></a><span class="comment">// that can be called because of exported values (consts, vars, or</span>
        <a id="L164"></a><span class="comment">// function results) of that type. Could determine if that is the</span>
        <a id="L165"></a><span class="comment">// case and then show those methods in an appropriate section.</span>
        <a id="L166"></a>return;
    <a id="L167"></a>}

    <a id="L169"></a><span class="comment">// perhaps a factory function</span>
    <a id="L170"></a><span class="comment">// determine result type, if any</span>
    <a id="L171"></a>if len(fun.Type.Results) &gt;= 1 {
        <a id="L172"></a>res := fun.Type.Results[0];
        <a id="L173"></a>if len(res.Names) &lt;= 1 {
            <a id="L174"></a><span class="comment">// exactly one (named or anonymous) result associated</span>
            <a id="L175"></a><span class="comment">// with the first type in result signature (there may</span>
            <a id="L176"></a><span class="comment">// be more than one result)</span>
            <a id="L177"></a>tname := baseTypeName(res.Type);
            <a id="L178"></a>typ := doc.lookupTypeDoc(tname);
            <a id="L179"></a>if typ != nil {
                <a id="L180"></a><span class="comment">// named and exported result type</span>

                <a id="L182"></a><span class="comment">// Work-around for failure of heuristic: In package os</span>
                <a id="L183"></a><span class="comment">// too many functions are considered factory functions</span>
                <a id="L184"></a><span class="comment">// for the Error type. Eliminate manually for now as</span>
                <a id="L185"></a><span class="comment">// this appears to be the only important case in the</span>
                <a id="L186"></a><span class="comment">// current library where the heuristic fails.</span>
                <a id="L187"></a>if doc.pkgName == &#34;os&#34; &amp;&amp; tname == &#34;Error&#34; &amp;&amp;
                    <a id="L188"></a>name != &#34;NewError&#34; &amp;&amp; name != &#34;NewSyscallError&#34; {
                    <a id="L189"></a><span class="comment">// not a factory function for os.Error</span>
                    <a id="L190"></a>doc.funcs[name] = fun; <span class="comment">// treat as ordinary function</span>
                    <a id="L191"></a>return;
                <a id="L192"></a>}

                <a id="L194"></a>typ.factories[name] = fun;
                <a id="L195"></a>return;
            <a id="L196"></a>}
        <a id="L197"></a>}
    <a id="L198"></a>}

    <a id="L200"></a><span class="comment">// ordinary function</span>
    <a id="L201"></a>doc.funcs[name] = fun;
<a id="L202"></a>}


<a id="L205"></a>func (doc *docReader) addDecl(decl ast.Decl) {
    <a id="L206"></a>switch d := decl.(type) {
    <a id="L207"></a>case *ast.GenDecl:
        <a id="L208"></a>if len(d.Specs) &gt; 0 {
            <a id="L209"></a>switch d.Tok {
            <a id="L210"></a>case token.CONST, token.VAR:
                <a id="L211"></a><span class="comment">// constants and variables are always handled as a group</span>
                <a id="L212"></a>doc.addValue(d)
            <a id="L213"></a>case token.TYPE:
                <a id="L214"></a><span class="comment">// types are handled individually</span>
                <a id="L215"></a>var noPos token.Position;
                <a id="L216"></a>for _, spec := range d.Specs {
                    <a id="L217"></a><span class="comment">// make a (fake) GenDecl node for this TypeSpec</span>
                    <a id="L218"></a><span class="comment">// (we need to do this here - as opposed to just</span>
                    <a id="L219"></a><span class="comment">// for printing - so we don&#39;t lose the GenDecl</span>
                    <a id="L220"></a><span class="comment">// documentation)</span>
                    <a id="L221"></a><span class="comment">//</span>
                    <a id="L222"></a><span class="comment">// TODO(gri): Consider just collecting the TypeSpec</span>
                    <a id="L223"></a><span class="comment">// node (and copy in the GenDecl.doc if there is no</span>
                    <a id="L224"></a><span class="comment">// doc in the TypeSpec - this is currently done in</span>
                    <a id="L225"></a><span class="comment">// makeTypeDocs below). Simpler data structures, but</span>
                    <a id="L226"></a><span class="comment">// would lose GenDecl documentation if the TypeSpec</span>
                    <a id="L227"></a><span class="comment">// has documentation as well.</span>
                    <a id="L228"></a>doc.addType(&amp;ast.GenDecl{d.Doc, d.Pos(), token.TYPE, noPos, []ast.Spec{spec}, noPos})
                    <a id="L229"></a><span class="comment">// A new GenDecl node is created, no need to nil out d.Doc.</span>
                <a id="L230"></a>}
            <a id="L231"></a>}
        <a id="L232"></a>}
    <a id="L233"></a>case *ast.FuncDecl:
        <a id="L234"></a>doc.addFunc(d)
    <a id="L235"></a>}
<a id="L236"></a>}


<a id="L239"></a>func copyCommentList(list []*ast.Comment) []*ast.Comment {
    <a id="L240"></a>copy := make([]*ast.Comment, len(list));
    <a id="L241"></a>for i, c := range list {
        <a id="L242"></a>copy[i] = c
    <a id="L243"></a>}
    <a id="L244"></a>return copy;
<a id="L245"></a>}


<a id="L248"></a>var (
    <a id="L249"></a>bug_markers = regexp.MustCompile(&#34;^/[/*][ \t]*BUG\\(.*\\):[ \t]*&#34;); <span class="comment">// BUG(uid):</span>
    <a id="L250"></a>bug_content = regexp.MustCompile(&#34;[^ \n\r\t]+&#34;);                    <span class="comment">// at least one non-whitespace char</span>
<a id="L251"></a>)


<a id="L254"></a><span class="comment">// addFile adds the AST for a source file to the docReader.</span>
<a id="L255"></a><span class="comment">// Adding the same AST multiple times is a no-op.</span>
<a id="L256"></a><span class="comment">//</span>
<a id="L257"></a>func (doc *docReader) addFile(src *ast.File) {
    <a id="L258"></a><span class="comment">// add package documentation</span>
    <a id="L259"></a>if src.Doc != nil {
        <a id="L260"></a><span class="comment">// TODO(gri) This won&#39;t do the right thing if there is more</span>
        <a id="L261"></a><span class="comment">//           than one file with package comments. Consider</span>
        <a id="L262"></a><span class="comment">//           using ast.MergePackageFiles which handles these</span>
        <a id="L263"></a><span class="comment">//           comments correctly (but currently looses BUG(...)</span>
        <a id="L264"></a><span class="comment">//           comments).</span>
        <a id="L265"></a>doc.doc = src.Doc;
        <a id="L266"></a>src.Doc = nil; <span class="comment">// doc consumed - remove from ast.File node</span>
    <a id="L267"></a>}

    <a id="L269"></a><span class="comment">// add all declarations</span>
    <a id="L270"></a>for _, decl := range src.Decls {
        <a id="L271"></a>doc.addDecl(decl)
    <a id="L272"></a>}

    <a id="L274"></a><span class="comment">// collect BUG(...) comments</span>
    <a id="L275"></a>for c := src.Comments; c != nil; c = c.Next {
        <a id="L276"></a>text := c.List[0].Text;
        <a id="L277"></a>cstr := string(text);
        <a id="L278"></a>if m := bug_markers.ExecuteString(cstr); len(m) &gt; 0 {
            <a id="L279"></a><span class="comment">// found a BUG comment; maybe empty</span>
            <a id="L280"></a>if bstr := cstr[m[1]:len(cstr)]; bug_content.MatchString(bstr) {
                <a id="L281"></a><span class="comment">// non-empty BUG comment; collect comment without BUG prefix</span>
                <a id="L282"></a>list := copyCommentList(c.List);
                <a id="L283"></a>list[0].Text = text[m[1]:len(text)];
                <a id="L284"></a>doc.bugs.Push(&amp;ast.CommentGroup{list, nil});
            <a id="L285"></a>}
        <a id="L286"></a>}
    <a id="L287"></a>}
    <a id="L288"></a>src.Comments = nil; <span class="comment">// consumed unassociated comments - remove from ast.File node</span>
<a id="L289"></a>}


<a id="L292"></a>func NewFileDoc(file *ast.File) *PackageDoc {
    <a id="L293"></a>var r docReader;
    <a id="L294"></a>r.init(file.Name.Value);
    <a id="L295"></a>r.addFile(file);
    <a id="L296"></a>return r.newDoc(&#34;&#34;, &#34;&#34;, nil);
<a id="L297"></a>}


<a id="L300"></a>func NewPackageDoc(pkg *ast.Package, importpath string) *PackageDoc {
    <a id="L301"></a>var r docReader;
    <a id="L302"></a>r.init(pkg.Name);
    <a id="L303"></a>filenames := make([]string, len(pkg.Files));
    <a id="L304"></a>i := 0;
    <a id="L305"></a>for filename, f := range pkg.Files {
        <a id="L306"></a>r.addFile(f);
        <a id="L307"></a>filenames[i] = filename;
        <a id="L308"></a>i++;
    <a id="L309"></a>}
    <a id="L310"></a>return r.newDoc(importpath, pkg.Path, filenames);
<a id="L311"></a>}


<a id="L314"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L315"></a><span class="comment">// Conversion to external representation</span>

<a id="L317"></a><span class="comment">// ValueDoc is the documentation for a group of declared</span>
<a id="L318"></a><span class="comment">// values, either vars or consts.</span>
<a id="L319"></a><span class="comment">//</span>
<a id="L320"></a>type ValueDoc struct {
    <a id="L321"></a>Doc   string;
    <a id="L322"></a>Decl  *ast.GenDecl;
    <a id="L323"></a>order int;
<a id="L324"></a>}

<a id="L326"></a>type sortValueDoc []*ValueDoc

<a id="L328"></a>func (p sortValueDoc) Len() int      { return len(p) }
<a id="L329"></a>func (p sortValueDoc) Swap(i, j int) { p[i], p[j] = p[j], p[i] }


<a id="L332"></a>func declName(d *ast.GenDecl) string {
    <a id="L333"></a>if len(d.Specs) != 1 {
        <a id="L334"></a>return &#34;&#34;
    <a id="L335"></a>}

    <a id="L337"></a>switch v := d.Specs[0].(type) {
    <a id="L338"></a>case *ast.ValueSpec:
        <a id="L339"></a>return v.Names[0].Value
    <a id="L340"></a>case *ast.TypeSpec:
        <a id="L341"></a>return v.Name.Value
    <a id="L342"></a>}

    <a id="L344"></a>return &#34;&#34;;
<a id="L345"></a>}


<a id="L348"></a>func (p sortValueDoc) Less(i, j int) bool {
    <a id="L349"></a><span class="comment">// sort by name</span>
    <a id="L350"></a><span class="comment">// pull blocks (name = &#34;&#34;) up to top</span>
    <a id="L351"></a><span class="comment">// in original order</span>
    <a id="L352"></a>if ni, nj := declName(p[i].Decl), declName(p[j].Decl); ni != nj {
        <a id="L353"></a>return ni &lt; nj
    <a id="L354"></a>}
    <a id="L355"></a>return p[i].order &lt; p[j].order;
<a id="L356"></a>}


<a id="L359"></a>func makeValueDocs(v *vector.Vector, tok token.Token) []*ValueDoc {
    <a id="L360"></a>d := make([]*ValueDoc, v.Len()); <span class="comment">// big enough in any case</span>
    <a id="L361"></a>n := 0;
    <a id="L362"></a>for i := range d {
        <a id="L363"></a>decl := v.At(i).(*ast.GenDecl);
        <a id="L364"></a>if decl.Tok == tok {
            <a id="L365"></a>d[n] = &amp;ValueDoc{CommentText(decl.Doc), decl, i};
            <a id="L366"></a>n++;
            <a id="L367"></a>decl.Doc = nil; <span class="comment">// doc consumed - removed from AST</span>
        <a id="L368"></a>}
    <a id="L369"></a>}
    <a id="L370"></a>d = d[0:n];
    <a id="L371"></a>sort.Sort(sortValueDoc(d));
    <a id="L372"></a>return d;
<a id="L373"></a>}


<a id="L376"></a><span class="comment">// FuncDoc is the documentation for a func declaration,</span>
<a id="L377"></a><span class="comment">// either a top-level function or a method function.</span>
<a id="L378"></a><span class="comment">//</span>
<a id="L379"></a>type FuncDoc struct {
    <a id="L380"></a>Doc  string;
    <a id="L381"></a>Recv ast.Expr; <span class="comment">// TODO(rsc): Would like string here</span>
    <a id="L382"></a>Name string;
    <a id="L383"></a>Decl *ast.FuncDecl;
<a id="L384"></a>}

<a id="L386"></a>type sortFuncDoc []*FuncDoc

<a id="L388"></a>func (p sortFuncDoc) Len() int           { return len(p) }
<a id="L389"></a>func (p sortFuncDoc) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
<a id="L390"></a>func (p sortFuncDoc) Less(i, j int) bool { return p[i].Name &lt; p[j].Name }


<a id="L393"></a>func makeFuncDocs(m map[string]*ast.FuncDecl) []*FuncDoc {
    <a id="L394"></a>d := make([]*FuncDoc, len(m));
    <a id="L395"></a>i := 0;
    <a id="L396"></a>for _, f := range m {
        <a id="L397"></a>doc := new(FuncDoc);
        <a id="L398"></a>doc.Doc = CommentText(f.Doc);
        <a id="L399"></a>f.Doc = nil; <span class="comment">// doc consumed - remove from ast.FuncDecl node</span>
        <a id="L400"></a>if f.Recv != nil {
            <a id="L401"></a>doc.Recv = f.Recv.Type
        <a id="L402"></a>}
        <a id="L403"></a>doc.Name = f.Name.Value;
        <a id="L404"></a>doc.Decl = f;
        <a id="L405"></a>d[i] = doc;
        <a id="L406"></a>i++;
    <a id="L407"></a>}
    <a id="L408"></a>sort.Sort(sortFuncDoc(d));
    <a id="L409"></a>return d;
<a id="L410"></a>}


<a id="L413"></a><span class="comment">// TypeDoc is the documentation for a declared type.</span>
<a id="L414"></a><span class="comment">// Consts and Vars are sorted lists of constants and variables of (mostly) that type.</span>
<a id="L415"></a><span class="comment">// Factories is a sorted list of factory functions that return that type.</span>
<a id="L416"></a><span class="comment">// Methods is a sorted list of method functions on that type.</span>
<a id="L417"></a>type TypeDoc struct {
    <a id="L418"></a>Doc       string;
    <a id="L419"></a>Type      *ast.TypeSpec;
    <a id="L420"></a>Consts    []*ValueDoc;
    <a id="L421"></a>Vars      []*ValueDoc;
    <a id="L422"></a>Factories []*FuncDoc;
    <a id="L423"></a>Methods   []*FuncDoc;
    <a id="L424"></a>Decl      *ast.GenDecl;
    <a id="L425"></a>order     int;
<a id="L426"></a>}

<a id="L428"></a>type sortTypeDoc []*TypeDoc

<a id="L430"></a>func (p sortTypeDoc) Len() int      { return len(p) }
<a id="L431"></a>func (p sortTypeDoc) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
<a id="L432"></a>func (p sortTypeDoc) Less(i, j int) bool {
    <a id="L433"></a><span class="comment">// sort by name</span>
    <a id="L434"></a><span class="comment">// pull blocks (name = &#34;&#34;) up to top</span>
    <a id="L435"></a><span class="comment">// in original order</span>
    <a id="L436"></a>if ni, nj := p[i].Type.Name.Value, p[j].Type.Name.Value; ni != nj {
        <a id="L437"></a>return ni &lt; nj
    <a id="L438"></a>}
    <a id="L439"></a>return p[i].order &lt; p[j].order;
<a id="L440"></a>}


<a id="L443"></a><span class="comment">// NOTE(rsc): This would appear not to be correct for type ( )</span>
<a id="L444"></a><span class="comment">// blocks, but the doc extractor above has split them into</span>
<a id="L445"></a><span class="comment">// individual declarations.</span>
<a id="L446"></a>func (doc *docReader) makeTypeDocs(m map[string]*typeDoc) []*TypeDoc {
    <a id="L447"></a>d := make([]*TypeDoc, len(m));
    <a id="L448"></a>i := 0;
    <a id="L449"></a>for _, old := range m {
        <a id="L450"></a><span class="comment">// all typeDocs should have a declaration associated with</span>
        <a id="L451"></a><span class="comment">// them after processing an entire package - be conservative</span>
        <a id="L452"></a><span class="comment">// and check</span>
        <a id="L453"></a>if decl := old.decl; decl != nil {
            <a id="L454"></a>typespec := decl.Specs[0].(*ast.TypeSpec);
            <a id="L455"></a>t := new(TypeDoc);
            <a id="L456"></a>doc := typespec.Doc;
            <a id="L457"></a>typespec.Doc = nil; <span class="comment">// doc consumed - remove from ast.TypeSpec node</span>
            <a id="L458"></a>if doc == nil {
                <a id="L459"></a><span class="comment">// no doc associated with the spec, use the declaration doc, if any</span>
                <a id="L460"></a>doc = decl.Doc
            <a id="L461"></a>}
            <a id="L462"></a>decl.Doc = nil; <span class="comment">// doc consumed - remove from ast.Decl node</span>
            <a id="L463"></a>t.Doc = CommentText(doc);
            <a id="L464"></a>t.Type = typespec;
            <a id="L465"></a>t.Consts = makeValueDocs(old.values, token.CONST);
            <a id="L466"></a>t.Vars = makeValueDocs(old.values, token.VAR);
            <a id="L467"></a>t.Factories = makeFuncDocs(old.factories);
            <a id="L468"></a>t.Methods = makeFuncDocs(old.methods);
            <a id="L469"></a>t.Decl = old.decl;
            <a id="L470"></a>t.order = i;
            <a id="L471"></a>d[i] = t;
            <a id="L472"></a>i++;
        <a id="L473"></a>} else {
            <a id="L474"></a><span class="comment">// no corresponding type declaration found - move any associated</span>
            <a id="L475"></a><span class="comment">// values, factory functions, and methods back to the top-level</span>
            <a id="L476"></a><span class="comment">// so that they are not lost (this should only happen if a package</span>
            <a id="L477"></a><span class="comment">// file containing the explicit type declaration is missing or if</span>
            <a id="L478"></a><span class="comment">// an unqualified type name was used after a &#34;.&#34; import)</span>
            <a id="L479"></a><span class="comment">// 1) move values</span>
            <a id="L480"></a>doc.values.AppendVector(old.values);
            <a id="L481"></a><span class="comment">// 2) move factory functions</span>
            <a id="L482"></a>for name, f := range old.factories {
                <a id="L483"></a>doc.funcs[name] = f
            <a id="L484"></a>}
            <a id="L485"></a><span class="comment">// 3) move methods</span>
            <a id="L486"></a>for name, f := range old.methods {
                <a id="L487"></a><span class="comment">// don&#39;t overwrite functions with the same name</span>
                <a id="L488"></a>if _, found := doc.funcs[name]; !found {
                    <a id="L489"></a>doc.funcs[name] = f
                <a id="L490"></a>}
            <a id="L491"></a>}
        <a id="L492"></a>}
    <a id="L493"></a>}
    <a id="L494"></a>d = d[0:i]; <span class="comment">// some types may have been ignored</span>
    <a id="L495"></a>sort.Sort(sortTypeDoc(d));
    <a id="L496"></a>return d;
<a id="L497"></a>}


<a id="L500"></a>func makeBugDocs(v *vector.Vector) []string {
    <a id="L501"></a>d := make([]string, v.Len());
    <a id="L502"></a>for i := 0; i &lt; v.Len(); i++ {
        <a id="L503"></a>d[i] = CommentText(v.At(i).(*ast.CommentGroup))
    <a id="L504"></a>}
    <a id="L505"></a>return d;
<a id="L506"></a>}


<a id="L509"></a><span class="comment">// PackageDoc is the documentation for an entire package.</span>
<a id="L510"></a><span class="comment">//</span>
<a id="L511"></a>type PackageDoc struct {
    <a id="L512"></a>PackageName string;
    <a id="L513"></a>ImportPath  string;
    <a id="L514"></a>FilePath    string;
    <a id="L515"></a>Filenames   []string;
    <a id="L516"></a>Doc         string;
    <a id="L517"></a>Consts      []*ValueDoc;
    <a id="L518"></a>Types       []*TypeDoc;
    <a id="L519"></a>Vars        []*ValueDoc;
    <a id="L520"></a>Funcs       []*FuncDoc;
    <a id="L521"></a>Bugs        []string;
<a id="L522"></a>}


<a id="L525"></a><span class="comment">// newDoc returns the accumulated documentation for the package.</span>
<a id="L526"></a><span class="comment">//</span>
<a id="L527"></a>func (doc *docReader) newDoc(importpath, filepath string, filenames []string) *PackageDoc {
    <a id="L528"></a>p := new(PackageDoc);
    <a id="L529"></a>p.PackageName = doc.pkgName;
    <a id="L530"></a>p.ImportPath = importpath;
    <a id="L531"></a>p.FilePath = filepath;
    <a id="L532"></a>sort.SortStrings(filenames);
    <a id="L533"></a>p.Filenames = filenames;
    <a id="L534"></a>p.Doc = CommentText(doc.doc);
    <a id="L535"></a><span class="comment">// makeTypeDocs may extend the list of doc.values and</span>
    <a id="L536"></a><span class="comment">// doc.funcs and thus must be called before any other</span>
    <a id="L537"></a><span class="comment">// function consuming those lists</span>
    <a id="L538"></a>p.Types = doc.makeTypeDocs(doc.types);
    <a id="L539"></a>p.Consts = makeValueDocs(doc.values, token.CONST);
    <a id="L540"></a>p.Vars = makeValueDocs(doc.values, token.VAR);
    <a id="L541"></a>p.Funcs = makeFuncDocs(doc.funcs);
    <a id="L542"></a>p.Bugs = makeBugDocs(doc.bugs);
    <a id="L543"></a>return p;
<a id="L544"></a>}


<a id="L547"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L548"></a><span class="comment">// Filtering by name</span>

<a id="L550"></a><span class="comment">// Does s look like a regular expression?</span>
<a id="L551"></a>func isRegexp(s string) bool {
    <a id="L552"></a>metachars := &#34;.(|)*+?^$[]&#34;;
    <a id="L553"></a>for _, c := range s {
        <a id="L554"></a>for _, m := range metachars {
            <a id="L555"></a>if c == m {
                <a id="L556"></a>return true
            <a id="L557"></a>}
        <a id="L558"></a>}
    <a id="L559"></a>}
    <a id="L560"></a>return false;
<a id="L561"></a>}


<a id="L564"></a>func match(s string, a []string) bool {
    <a id="L565"></a>for _, t := range a {
        <a id="L566"></a>if isRegexp(t) {
            <a id="L567"></a>if matched, _ := regexp.MatchString(t, s); matched {
                <a id="L568"></a>return true
            <a id="L569"></a>}
        <a id="L570"></a>}
        <a id="L571"></a>if s == t {
            <a id="L572"></a>return true
        <a id="L573"></a>}
    <a id="L574"></a>}
    <a id="L575"></a>return false;
<a id="L576"></a>}


<a id="L579"></a>func matchDecl(d *ast.GenDecl, names []string) bool {
    <a id="L580"></a>for _, d := range d.Specs {
        <a id="L581"></a>switch v := d.(type) {
        <a id="L582"></a>case *ast.ValueSpec:
            <a id="L583"></a>for _, name := range v.Names {
                <a id="L584"></a>if match(name.Value, names) {
                    <a id="L585"></a>return true
                <a id="L586"></a>}
            <a id="L587"></a>}
        <a id="L588"></a>case *ast.TypeSpec:
            <a id="L589"></a>if match(v.Name.Value, names) {
                <a id="L590"></a>return true
            <a id="L591"></a>}
        <a id="L592"></a>}
    <a id="L593"></a>}
    <a id="L594"></a>return false;
<a id="L595"></a>}


<a id="L598"></a>func filterValueDocs(a []*ValueDoc, names []string) []*ValueDoc {
    <a id="L599"></a>w := 0;
    <a id="L600"></a>for _, vd := range a {
        <a id="L601"></a>if matchDecl(vd.Decl, names) {
            <a id="L602"></a>a[w] = vd;
            <a id="L603"></a>w++;
        <a id="L604"></a>}
    <a id="L605"></a>}
    <a id="L606"></a>return a[0:w];
<a id="L607"></a>}


<a id="L610"></a>func filterFuncDocs(a []*FuncDoc, names []string) []*FuncDoc {
    <a id="L611"></a>w := 0;
    <a id="L612"></a>for _, fd := range a {
        <a id="L613"></a>if match(fd.Name, names) {
            <a id="L614"></a>a[w] = fd;
            <a id="L615"></a>w++;
        <a id="L616"></a>}
    <a id="L617"></a>}
    <a id="L618"></a>return a[0:w];
<a id="L619"></a>}


<a id="L622"></a>func filterTypeDocs(a []*TypeDoc, names []string) []*TypeDoc {
    <a id="L623"></a>w := 0;
    <a id="L624"></a>for _, td := range a {
        <a id="L625"></a>match := false;
        <a id="L626"></a>if matchDecl(td.Decl, names) {
            <a id="L627"></a>match = true
        <a id="L628"></a>} else {
            <a id="L629"></a><span class="comment">// type name doesn&#39;t match, but we may have matching factories or methods</span>
            <a id="L630"></a>td.Factories = filterFuncDocs(td.Factories, names);
            <a id="L631"></a>td.Methods = filterFuncDocs(td.Methods, names);
            <a id="L632"></a>match = len(td.Factories) &gt; 0 || len(td.Methods) &gt; 0;
        <a id="L633"></a>}
        <a id="L634"></a>if match {
            <a id="L635"></a>a[w] = td;
            <a id="L636"></a>w++;
        <a id="L637"></a>}
    <a id="L638"></a>}
    <a id="L639"></a>return a[0:w];
<a id="L640"></a>}


<a id="L643"></a><span class="comment">// Filter eliminates information from d that is not</span>
<a id="L644"></a><span class="comment">// about one of the given names.</span>
<a id="L645"></a><span class="comment">// TODO: Recognize &#34;Type.Method&#34; as a name.</span>
<a id="L646"></a><span class="comment">// TODO(r): maybe precompile the regexps.</span>
<a id="L647"></a><span class="comment">//</span>
<a id="L648"></a>func (p *PackageDoc) Filter(names []string) {
    <a id="L649"></a>p.Consts = filterValueDocs(p.Consts, names);
    <a id="L650"></a>p.Vars = filterValueDocs(p.Vars, names);
    <a id="L651"></a>p.Types = filterTypeDocs(p.Types, names);
    <a id="L652"></a>p.Funcs = filterFuncDocs(p.Funcs, names);
    <a id="L653"></a>p.Doc = &#34;&#34;; <span class="comment">// don&#39;t show top-level package doc</span>
<a id="L654"></a>}
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
