<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/printer/printer_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/go/printer/printer_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package printer

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;flag&#34;;
    <a id="L10"></a>&#34;io&#34;;
    <a id="L11"></a>&#34;go/ast&#34;;
    <a id="L12"></a>&#34;go/parser&#34;;
    <a id="L13"></a>&#34;path&#34;;
    <a id="L14"></a>&#34;testing&#34;;
<a id="L15"></a>)


<a id="L18"></a>const (
    <a id="L19"></a>dataDir  = &#34;testdata&#34;;
    <a id="L20"></a>tabwidth = 8;
<a id="L21"></a>)


<a id="L24"></a>var update = flag.Bool(&#34;update&#34;, false, &#34;update golden files&#34;)


<a id="L27"></a>func lineString(text []byte, i int) string {
    <a id="L28"></a>i0 := i;
    <a id="L29"></a>for i &lt; len(text) &amp;&amp; text[i] != &#39;\n&#39; {
        <a id="L30"></a>i++
    <a id="L31"></a>}
    <a id="L32"></a>return string(text[i0:i]);
<a id="L33"></a>}


<a id="L36"></a>type checkMode uint

<a id="L38"></a>const (
    <a id="L39"></a>export checkMode = 1 &lt;&lt; iota;
    <a id="L40"></a>rawFormat;
<a id="L41"></a>)


<a id="L44"></a>func check(t *testing.T, source, golden string, mode checkMode) {
    <a id="L45"></a><span class="comment">// parse source</span>
    <a id="L46"></a>prog, err := parser.ParseFile(source, nil, parser.ParseComments);
    <a id="L47"></a>if err != nil {
        <a id="L48"></a>t.Error(err);
        <a id="L49"></a>return;
    <a id="L50"></a>}

    <a id="L52"></a><span class="comment">// filter exports if necessary</span>
    <a id="L53"></a>if mode&amp;export != 0 {
        <a id="L54"></a>ast.FileExports(prog); <span class="comment">// ignore result</span>
        <a id="L55"></a>prog.Comments = nil;   <span class="comment">// don&#39;t print comments that are not in AST</span>
    <a id="L56"></a>}

    <a id="L58"></a><span class="comment">// determine printer configuration</span>
    <a id="L59"></a>cfg := Config{Tabwidth: tabwidth};
    <a id="L60"></a>if mode&amp;rawFormat != 0 {
        <a id="L61"></a>cfg.Mode |= RawFormat
    <a id="L62"></a>}

    <a id="L64"></a><span class="comment">// format source</span>
    <a id="L65"></a>var buf bytes.Buffer;
    <a id="L66"></a>if _, err := cfg.Fprint(&amp;buf, prog); err != nil {
        <a id="L67"></a>t.Error(err)
    <a id="L68"></a>}
    <a id="L69"></a>res := buf.Bytes();

    <a id="L71"></a><span class="comment">// update golden files if necessary</span>
    <a id="L72"></a>if *update {
        <a id="L73"></a>if err := io.WriteFile(golden, res, 0644); err != nil {
            <a id="L74"></a>t.Error(err)
        <a id="L75"></a>}
        <a id="L76"></a>return;
    <a id="L77"></a>}

    <a id="L79"></a><span class="comment">// get golden</span>
    <a id="L80"></a>gld, err := io.ReadFile(golden);
    <a id="L81"></a>if err != nil {
        <a id="L82"></a>t.Error(err);
        <a id="L83"></a>return;
    <a id="L84"></a>}

    <a id="L86"></a><span class="comment">// compare lengths</span>
    <a id="L87"></a>if len(res) != len(gld) {
        <a id="L88"></a>t.Errorf(&#34;len = %d, expected %d (= len(%s))&#34;, len(res), len(gld), golden)
    <a id="L89"></a>}

    <a id="L91"></a><span class="comment">// compare contents</span>
    <a id="L92"></a>for i, line, offs := 0, 1, 0; i &lt; len(res) &amp;&amp; i &lt; len(gld); i++ {
        <a id="L93"></a>ch := res[i];
        <a id="L94"></a>if ch != gld[i] {
            <a id="L95"></a>t.Errorf(&#34;%s:%d:%d: %s&#34;, source, line, i-offs+1, lineString(res, offs));
            <a id="L96"></a>t.Errorf(&#34;%s:%d:%d: %s&#34;, golden, line, i-offs+1, lineString(gld, offs));
            <a id="L97"></a>t.Error();
            <a id="L98"></a>return;
        <a id="L99"></a>}
        <a id="L100"></a>if ch == &#39;\n&#39; {
            <a id="L101"></a>line++;
            <a id="L102"></a>offs = i + 1;
        <a id="L103"></a>}
    <a id="L104"></a>}
<a id="L105"></a>}


<a id="L108"></a>type entry struct {
    <a id="L109"></a>source, golden string;
    <a id="L110"></a>mode           checkMode;
<a id="L111"></a>}

<a id="L113"></a><span class="comment">// Use gotest -update to create/update the respective golden files.</span>
<a id="L114"></a>var data = []entry{
    <a id="L115"></a>entry{&#34;empty.input&#34;, &#34;empty.golden&#34;, 0},
    <a id="L116"></a>entry{&#34;comments.input&#34;, &#34;comments.golden&#34;, 0},
    <a id="L117"></a>entry{&#34;comments.input&#34;, &#34;comments.x&#34;, export},
    <a id="L118"></a>entry{&#34;linebreaks.input&#34;, &#34;linebreaks.golden&#34;, 0},
    <a id="L119"></a>entry{&#34;expressions.input&#34;, &#34;expressions.golden&#34;, 0},
    <a id="L120"></a>entry{&#34;expressions.input&#34;, &#34;expressions.raw&#34;, rawFormat},
    <a id="L121"></a>entry{&#34;declarations.input&#34;, &#34;declarations.golden&#34;, 0},
    <a id="L122"></a>entry{&#34;statements.input&#34;, &#34;statements.golden&#34;, 0},
<a id="L123"></a>}


<a id="L126"></a>func Test(t *testing.T) {
    <a id="L127"></a>for _, e := range data {
        <a id="L128"></a>source := path.Join(dataDir, e.source);
        <a id="L129"></a>golden := path.Join(dataDir, e.golden);
        <a id="L130"></a>check(t, source, golden, e.mode);
        <a id="L131"></a><span class="comment">// TODO(gri) check that golden is idempotent</span>
        <a id="L132"></a><span class="comment">//check(t, golden, golden, e.mode);</span>
    <a id="L133"></a>}
<a id="L134"></a>}
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
