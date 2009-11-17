<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /doc/progs/cat_rot13.go</title>

  <link rel="stylesheet" type="text/css" href="../style.css">
  <script type="text/javascript" src="../godocs.js"></script>

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
        <a href="../../index.html"><img src="../logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../go_tutorial.html">Tutorial</a></li>
    <li><a href="../effective_go.html">Effective Go</a></li>
    <li><a href="../go_faq.html">FAQ</a></li>
    <li><a href="../go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../go_spec.html">Language Specification</a></li>
    <li><a href="../go_mem.html">Memory Model</a></li>
    <li><a href="../go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../install.html">Install Go</a></li>
    <li><a href="../contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../src/index.html">Source files</a></li>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /doc/progs/cat_rot13.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package main

<a id="L7"></a>import (
    <a id="L8"></a>&#34;./file&#34;;
    <a id="L9"></a>&#34;flag&#34;;
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;os&#34;;
<a id="L12"></a>)

<a id="L14"></a>var rot13Flag = flag.Bool(&#34;rot13&#34;, false, &#34;rot13 the input&#34;)

<a id="L16"></a>func rot13(b byte) byte {
    <a id="L17"></a>if &#39;a&#39; &lt;= b &amp;&amp; b &lt;= &#39;z&#39; {
        <a id="L18"></a>b = &#39;a&#39; + ((b-&#39;a&#39;)+13)%26
    <a id="L19"></a>}
    <a id="L20"></a>if &#39;A&#39; &lt;= b &amp;&amp; b &lt;= &#39;Z&#39; {
        <a id="L21"></a>b = &#39;A&#39; + ((b-&#39;A&#39;)+13)%26
    <a id="L22"></a>}
    <a id="L23"></a>return b;
<a id="L24"></a>}

<a id="L26"></a>type reader interface {
    <a id="L27"></a>Read(b []byte) (ret int, err os.Error);
    <a id="L28"></a>String() string;
<a id="L29"></a>}

<a id="L31"></a>type rotate13 struct {
    <a id="L32"></a>source reader;
<a id="L33"></a>}

<a id="L35"></a>func newRotate13(source reader) *rotate13 { <a id="L36"></a>return &amp;rotate13{source} <a id="L37"></a>}

<a id="L39"></a>func (r13 *rotate13) Read(b []byte) (ret int, err os.Error) {
    <a id="L40"></a>r, e := r13.source.Read(b);
    <a id="L41"></a>for i := 0; i &lt; r; i++ {
        <a id="L42"></a>b[i] = rot13(b[i])
    <a id="L43"></a>}
    <a id="L44"></a>return r, e;
<a id="L45"></a>}

<a id="L47"></a>func (r13 *rotate13) String() string { <a id="L48"></a>return r13.source.String() <a id="L49"></a>}
<a id="L50"></a><span class="comment">// end of rotate13 implementation</span>

<a id="L52"></a>func cat(r reader) {
    <a id="L53"></a>const NBUF = 512;
    <a id="L54"></a>var buf [NBUF]byte;

    <a id="L56"></a>if *rot13Flag {
        <a id="L57"></a>r = newRotate13(r)
    <a id="L58"></a>}
    <a id="L59"></a>for {
        <a id="L60"></a>switch nr, er := r.Read(&amp;buf); {
        <a id="L61"></a>case nr &lt; 0:
            <a id="L62"></a>fmt.Fprintf(os.Stderr, &#34;cat: error reading from %s: %s\n&#34;, r.String(), er.String());
            <a id="L63"></a>os.Exit(1);
        <a id="L64"></a>case nr == 0: <span class="comment">// EOF</span>
            <a id="L65"></a>return
        <a id="L66"></a>case nr &gt; 0:
            <a id="L67"></a>nw, ew := file.Stdout.Write(buf[0:nr]);
            <a id="L68"></a>if nw != nr {
                <a id="L69"></a>fmt.Fprintf(os.Stderr, &#34;cat: error writing from %s: %s\n&#34;, r.String(), ew.String())
            <a id="L70"></a>}
        <a id="L71"></a>}
    <a id="L72"></a>}
<a id="L73"></a>}

<a id="L75"></a>func main() {
    <a id="L76"></a>flag.Parse(); <span class="comment">// Scans the arg list and sets up flags</span>
    <a id="L77"></a>if flag.NArg() == 0 {
        <a id="L78"></a>cat(file.Stdin)
    <a id="L79"></a>}
    <a id="L80"></a>for i := 0; i &lt; flag.NArg(); i++ {
        <a id="L81"></a>f, err := file.Open(flag.Arg(i), 0, 0);
        <a id="L82"></a>if f == nil {
            <a id="L83"></a>fmt.Fprintf(os.Stderr, &#34;cat: can&#39;t open %s: error %s\n&#34;, flag.Arg(i), err);
            <a id="L84"></a>os.Exit(1);
        <a id="L85"></a>}
        <a id="L86"></a>cat(f);
        <a id="L87"></a>f.Close();
    <a id="L88"></a>}
<a id="L89"></a>}
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
