<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/parser/interface.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/go/parser/interface.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This file contains the exported entry points for invoking the parser.</span>

<a id="L7"></a>package parser

<a id="L9"></a>import (
    <a id="L10"></a>&#34;bytes&#34;;
    <a id="L11"></a>&#34;fmt&#34;;
    <a id="L12"></a>&#34;go/ast&#34;;
    <a id="L13"></a>&#34;go/scanner&#34;;
    <a id="L14"></a>&#34;io&#34;;
    <a id="L15"></a>&#34;os&#34;;
    <a id="L16"></a>pathutil &#34;path&#34;;
    <a id="L17"></a>&#34;strings&#34;;
<a id="L18"></a>)


<a id="L21"></a><span class="comment">// If src != nil, readSource converts src to a []byte if possible;</span>
<a id="L22"></a><span class="comment">// otherwise it returns an error. If src == nil, readSource returns</span>
<a id="L23"></a><span class="comment">// the result of reading the file specified by filename.</span>
<a id="L24"></a><span class="comment">//</span>
<a id="L25"></a>func readSource(filename string, src interface{}) ([]byte, os.Error) {
    <a id="L26"></a>if src != nil {
        <a id="L27"></a>switch s := src.(type) {
        <a id="L28"></a>case string:
            <a id="L29"></a>return strings.Bytes(s), nil
        <a id="L30"></a>case []byte:
            <a id="L31"></a>return s, nil
        <a id="L32"></a>case *bytes.Buffer:
            <a id="L33"></a><span class="comment">// is io.Reader, but src is already available in []byte form</span>
            <a id="L34"></a>if s != nil {
                <a id="L35"></a>return s.Bytes(), nil
            <a id="L36"></a>}
        <a id="L37"></a>case io.Reader:
            <a id="L38"></a>var buf bytes.Buffer;
            <a id="L39"></a>_, err := io.Copy(&amp;buf, s);
            <a id="L40"></a>if err != nil {
                <a id="L41"></a>return nil, err
            <a id="L42"></a>}
            <a id="L43"></a>return buf.Bytes(), nil;
        <a id="L44"></a>default:
            <a id="L45"></a>return nil, os.ErrorString(&#34;invalid source&#34;)
        <a id="L46"></a>}
    <a id="L47"></a>}

    <a id="L49"></a>return io.ReadFile(filename);
<a id="L50"></a>}


<a id="L53"></a><span class="comment">// ParseExpr parses a Go expression and returns the corresponding</span>
<a id="L54"></a><span class="comment">// AST node. The filename and src arguments have the same interpretation</span>
<a id="L55"></a><span class="comment">// as for ParseFile. If there is an error, the result expression</span>
<a id="L56"></a><span class="comment">// may be nil or contain a partial AST.</span>
<a id="L57"></a><span class="comment">//</span>
<a id="L58"></a>func ParseExpr(filename string, src interface{}) (ast.Expr, os.Error) {
    <a id="L59"></a>data, err := readSource(filename, src);
    <a id="L60"></a>if err != nil {
        <a id="L61"></a>return nil, err
    <a id="L62"></a>}

    <a id="L64"></a>var p parser;
    <a id="L65"></a>p.init(filename, data, 0);
    <a id="L66"></a>return p.parseExpr(), p.GetError(scanner.Sorted);
<a id="L67"></a>}


<a id="L70"></a><span class="comment">// ParseStmtList parses a list of Go statements and returns the list</span>
<a id="L71"></a><span class="comment">// of corresponding AST nodes. The filename and src arguments have the same</span>
<a id="L72"></a><span class="comment">// interpretation as for ParseFile. If there is an error, the node</span>
<a id="L73"></a><span class="comment">// list may be nil or contain partial ASTs.</span>
<a id="L74"></a><span class="comment">//</span>
<a id="L75"></a>func ParseStmtList(filename string, src interface{}) ([]ast.Stmt, os.Error) {
    <a id="L76"></a>data, err := readSource(filename, src);
    <a id="L77"></a>if err != nil {
        <a id="L78"></a>return nil, err
    <a id="L79"></a>}

    <a id="L81"></a>var p parser;
    <a id="L82"></a>p.init(filename, data, 0);
    <a id="L83"></a>return p.parseStmtList(), p.GetError(scanner.Sorted);
<a id="L84"></a>}


<a id="L87"></a><span class="comment">// ParseDeclList parses a list of Go declarations and returns the list</span>
<a id="L88"></a><span class="comment">// of corresponding AST nodes.  The filename and src arguments have the same</span>
<a id="L89"></a><span class="comment">// interpretation as for ParseFile. If there is an error, the node</span>
<a id="L90"></a><span class="comment">// list may be nil or contain partial ASTs.</span>
<a id="L91"></a><span class="comment">//</span>
<a id="L92"></a>func ParseDeclList(filename string, src interface{}) ([]ast.Decl, os.Error) {
    <a id="L93"></a>data, err := readSource(filename, src);
    <a id="L94"></a>if err != nil {
        <a id="L95"></a>return nil, err
    <a id="L96"></a>}

    <a id="L98"></a>var p parser;
    <a id="L99"></a>p.init(filename, data, 0);
    <a id="L100"></a>return p.parseDeclList(), p.GetError(scanner.Sorted);
<a id="L101"></a>}


<a id="L104"></a><span class="comment">// ParseFile parses a Go source file and returns a File node.</span>
<a id="L105"></a><span class="comment">//</span>
<a id="L106"></a><span class="comment">// If src != nil, ParseFile parses the file source from src. src may</span>
<a id="L107"></a><span class="comment">// be provided in a variety of formats. At the moment the following types</span>
<a id="L108"></a><span class="comment">// are supported: string, []byte, and io.Reader. In this case, filename is</span>
<a id="L109"></a><span class="comment">// only used for source position information and error messages.</span>
<a id="L110"></a><span class="comment">//</span>
<a id="L111"></a><span class="comment">// If src == nil, ParseFile parses the file specified by filename.</span>
<a id="L112"></a><span class="comment">//</span>
<a id="L113"></a><span class="comment">// The mode parameter controls the amount of source text parsed and other</span>
<a id="L114"></a><span class="comment">// optional parser functionality.</span>
<a id="L115"></a><span class="comment">//</span>
<a id="L116"></a><span class="comment">// If the source couldn&#39;t be read, the returned AST is nil and the error</span>
<a id="L117"></a><span class="comment">// indicates the specific failure. If the source was read but syntax</span>
<a id="L118"></a><span class="comment">// errors were found, the result is a partial AST (with ast.BadX nodes</span>
<a id="L119"></a><span class="comment">// representing the fragments of erroneous source code). Multiple errors</span>
<a id="L120"></a><span class="comment">// are returned via a scanner.ErrorList which is sorted by file position.</span>
<a id="L121"></a><span class="comment">//</span>
<a id="L122"></a>func ParseFile(filename string, src interface{}, mode uint) (*ast.File, os.Error) {
    <a id="L123"></a>data, err := readSource(filename, src);
    <a id="L124"></a>if err != nil {
        <a id="L125"></a>return nil, err
    <a id="L126"></a>}

    <a id="L128"></a>var p parser;
    <a id="L129"></a>p.init(filename, data, mode);
    <a id="L130"></a>return p.parseFile(), p.GetError(scanner.NoMultiples);
<a id="L131"></a>}


<a id="L134"></a><span class="comment">// ParsePkgFile parses the file specified by filename and returns the</span>
<a id="L135"></a><span class="comment">// corresponding AST. If the file cannot be read, has syntax errors, or</span>
<a id="L136"></a><span class="comment">// does not belong to the package (i.e., pkgname != &#34;&#34; and the package</span>
<a id="L137"></a><span class="comment">// name in the file doesn&#39;t match pkkname), an error is returned. Mode</span>
<a id="L138"></a><span class="comment">// flags that control the amount of source text parsed are ignored.</span>
<a id="L139"></a><span class="comment">//</span>
<a id="L140"></a>func ParsePkgFile(pkgname, filename string, mode uint) (*ast.File, os.Error) {
    <a id="L141"></a>src, err := io.ReadFile(filename);
    <a id="L142"></a>if err != nil {
        <a id="L143"></a>return nil, err
    <a id="L144"></a>}

    <a id="L146"></a>if pkgname != &#34;&#34; {
        <a id="L147"></a>prog, err := ParseFile(filename, src, PackageClauseOnly);
        <a id="L148"></a>if err != nil {
            <a id="L149"></a>return nil, err
        <a id="L150"></a>}
        <a id="L151"></a>if prog.Name.Value != pkgname {
            <a id="L152"></a>return nil, os.NewError(fmt.Sprintf(&#34;multiple packages found: %s, %s&#34;, prog.Name.Value, pkgname))
        <a id="L153"></a>}
    <a id="L154"></a>}

    <a id="L156"></a><span class="comment">// ignore flags that control partial parsing</span>
    <a id="L157"></a>return ParseFile(filename, src, mode&amp;^(PackageClauseOnly|ImportsOnly));
<a id="L158"></a>}


<a id="L161"></a><span class="comment">// ParsePackage parses all files in the directory specified by path and</span>
<a id="L162"></a><span class="comment">// returns an AST representing the package found. The set of files may be</span>
<a id="L163"></a><span class="comment">// restricted by providing a non-nil filter function; only the files with</span>
<a id="L164"></a><span class="comment">// os.Dir entries passing through the filter are considered.</span>
<a id="L165"></a><span class="comment">// If ParsePackage does not find exactly one package, it returns an error.</span>
<a id="L166"></a><span class="comment">// Mode flags that control the amount of source text parsed are ignored.</span>
<a id="L167"></a><span class="comment">//</span>
<a id="L168"></a>func ParsePackage(path string, filter func(*os.Dir) bool, mode uint) (*ast.Package, os.Error) {
    <a id="L169"></a>fd, err := os.Open(path, os.O_RDONLY, 0);
    <a id="L170"></a>if err != nil {
        <a id="L171"></a>return nil, err
    <a id="L172"></a>}
    <a id="L173"></a>defer fd.Close();

    <a id="L175"></a>list, err := fd.Readdir(-1);
    <a id="L176"></a>if err != nil {
        <a id="L177"></a>return nil, err
    <a id="L178"></a>}

    <a id="L180"></a>name := &#34;&#34;;
    <a id="L181"></a>files := make(map[string]*ast.File);
    <a id="L182"></a>for i := 0; i &lt; len(list); i++ {
        <a id="L183"></a>entry := &amp;list[i];
        <a id="L184"></a>if filter == nil || filter(entry) {
            <a id="L185"></a>src, err := ParsePkgFile(name, pathutil.Join(path, entry.Name), mode);
            <a id="L186"></a>if err != nil {
                <a id="L187"></a>return nil, err
            <a id="L188"></a>}
            <a id="L189"></a>files[entry.Name] = src;
            <a id="L190"></a>if name == &#34;&#34; {
                <a id="L191"></a>name = src.Name.Value
            <a id="L192"></a>}
        <a id="L193"></a>}
    <a id="L194"></a>}

    <a id="L196"></a>if len(files) == 0 {
        <a id="L197"></a>return nil, os.NewError(path + &#34;: no package found&#34;)
    <a id="L198"></a>}

    <a id="L200"></a>return &amp;ast.Package{name, path, files}, nil;
<a id="L201"></a>}
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
