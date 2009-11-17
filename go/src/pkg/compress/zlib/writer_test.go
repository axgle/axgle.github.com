<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/compress/zlib/writer_test.go</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/compress/zlib/writer_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package zlib

<a id="L7"></a>import (
    <a id="L8"></a>&#34;io&#34;;
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>&#34;testing&#34;;
<a id="L11"></a>)

<a id="L13"></a>var filenames = []string{
    <a id="L14"></a>&#34;testdata/e.txt&#34;,
    <a id="L15"></a>&#34;testdata/pi.txt&#34;,
<a id="L16"></a>}

<a id="L18"></a><span class="comment">// Tests that compressing and then decompressing the given file at the given compression level</span>
<a id="L19"></a><span class="comment">// yields equivalent bytes to the original file.</span>
<a id="L20"></a>func testFileLevel(t *testing.T, fn string, level int) {
    <a id="L21"></a><span class="comment">// Read the file, as golden output.</span>
    <a id="L22"></a>golden, err := os.Open(fn, os.O_RDONLY, 0444);
    <a id="L23"></a>if err != nil {
        <a id="L24"></a>t.Errorf(&#34;%s (level=%d): %v&#34;, fn, level, err);
        <a id="L25"></a>return;
    <a id="L26"></a>}
    <a id="L27"></a>defer golden.Close();

    <a id="L29"></a><span class="comment">// Read the file again, and push it through a pipe that compresses at the write end, and decompresses at the read end.</span>
    <a id="L30"></a>raw, err := os.Open(fn, os.O_RDONLY, 0444);
    <a id="L31"></a>if err != nil {
        <a id="L32"></a>t.Errorf(&#34;%s (level=%d): %v&#34;, fn, level, err);
        <a id="L33"></a>return;
    <a id="L34"></a>}
    <a id="L35"></a>piper, pipew := io.Pipe();
    <a id="L36"></a>defer piper.Close();
    <a id="L37"></a>go func() {
        <a id="L38"></a>defer raw.Close();
        <a id="L39"></a>defer pipew.Close();
        <a id="L40"></a>zlibw, err := NewDeflaterLevel(pipew, level);
        <a id="L41"></a>if err != nil {
            <a id="L42"></a>t.Errorf(&#34;%s (level=%d): %v&#34;, fn, level, err);
            <a id="L43"></a>return;
        <a id="L44"></a>}
        <a id="L45"></a>defer zlibw.Close();
        <a id="L46"></a>var b [1024]byte;
        <a id="L47"></a>for {
            <a id="L48"></a>n, err0 := raw.Read(&amp;b);
            <a id="L49"></a>if err0 != nil &amp;&amp; err0 != os.EOF {
                <a id="L50"></a>t.Errorf(&#34;%s (level=%d): %v&#34;, fn, level, err0);
                <a id="L51"></a>return;
            <a id="L52"></a>}
            <a id="L53"></a>_, err1 := zlibw.Write(b[0:n]);
            <a id="L54"></a>if err1 == os.EPIPE {
                <a id="L55"></a><span class="comment">// Fail, but do not report the error, as some other (presumably reportable) error broke the pipe.</span>
                <a id="L56"></a>return
            <a id="L57"></a>}
            <a id="L58"></a>if err1 != nil {
                <a id="L59"></a>t.Errorf(&#34;%s (level=%d): %v&#34;, fn, level, err1);
                <a id="L60"></a>return;
            <a id="L61"></a>}
            <a id="L62"></a>if err0 == os.EOF {
                <a id="L63"></a>break
            <a id="L64"></a>}
        <a id="L65"></a>}
    <a id="L66"></a>}();
    <a id="L67"></a>zlibr, err := NewInflater(piper);
    <a id="L68"></a>if err != nil {
        <a id="L69"></a>t.Errorf(&#34;%s (level=%d): %v&#34;, fn, level, err);
        <a id="L70"></a>return;
    <a id="L71"></a>}
    <a id="L72"></a>defer zlibr.Close();

    <a id="L74"></a><span class="comment">// Compare the two.</span>
    <a id="L75"></a>b0, err0 := io.ReadAll(golden);
    <a id="L76"></a>b1, err1 := io.ReadAll(zlibr);
    <a id="L77"></a>if err0 != nil {
        <a id="L78"></a>t.Errorf(&#34;%s (level=%d): %v&#34;, fn, level, err0);
        <a id="L79"></a>return;
    <a id="L80"></a>}
    <a id="L81"></a>if err1 != nil {
        <a id="L82"></a>t.Errorf(&#34;%s (level=%d): %v&#34;, fn, level, err1);
        <a id="L83"></a>return;
    <a id="L84"></a>}
    <a id="L85"></a>if len(b0) != len(b1) {
        <a id="L86"></a>t.Errorf(&#34;%s (level=%d): length mismatch %d versus %d&#34;, fn, level, len(b0), len(b1));
        <a id="L87"></a>return;
    <a id="L88"></a>}
    <a id="L89"></a>for i := 0; i &lt; len(b0); i++ {
        <a id="L90"></a>if b0[i] != b1[i] {
            <a id="L91"></a>t.Errorf(&#34;%s (level=%d): mismatch at %d, 0x%02x versus 0x%02x\n&#34;, fn, level, i, b0[i], b1[i]);
            <a id="L92"></a>return;
        <a id="L93"></a>}
    <a id="L94"></a>}
<a id="L95"></a>}

<a id="L97"></a>func TestWriter(t *testing.T) {
    <a id="L98"></a>for _, fn := range filenames {
        <a id="L99"></a>testFileLevel(t, fn, DefaultCompression);
        <a id="L100"></a>testFileLevel(t, fn, NoCompression);
        <a id="L101"></a>for level := BestSpeed; level &lt;= BestCompression; level++ {
            <a id="L102"></a>testFileLevel(t, fn, level)
        <a id="L103"></a>}
    <a id="L104"></a>}
<a id="L105"></a>}
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
