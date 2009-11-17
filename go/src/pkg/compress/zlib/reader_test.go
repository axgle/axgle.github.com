<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/compress/zlib/reader_test.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/compress/zlib/reader_test.go</h1>

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
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;io&#34;;
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)

<a id="L14"></a>type zlibTest struct {
    <a id="L15"></a>desc       string;
    <a id="L16"></a>raw        string;
    <a id="L17"></a>compressed []byte;
    <a id="L18"></a>err        os.Error;
<a id="L19"></a>}

<a id="L21"></a><span class="comment">// Compare-to-golden test data was generated by the ZLIB example program at</span>
<a id="L22"></a><span class="comment">// http://www.zlib.net/zpipe.c</span>

<a id="L24"></a>var zlibTests = []zlibTest{
    <a id="L25"></a>zlibTest{
        <a id="L26"></a>&#34;empty&#34;,
        <a id="L27"></a>&#34;&#34;,
        <a id="L28"></a>[]byte{0x78, 0x9c, 0x03, 0x00, 0x00, 0x00, 0x00, 0x01},
        <a id="L29"></a>nil,
    <a id="L30"></a>},
    <a id="L31"></a>zlibTest{
        <a id="L32"></a>&#34;goodbye&#34;,
        <a id="L33"></a>&#34;goodbye, world&#34;,
        <a id="L34"></a>[]byte{
            <a id="L35"></a>0x78, 0x9c, 0x4b, 0xcf, 0xcf, 0x4f, 0x49, 0xaa,
            <a id="L36"></a>0x4c, 0xd5, 0x51, 0x28, 0xcf, 0x2f, 0xca, 0x49,
            <a id="L37"></a>0x01, 0x00, 0x28, 0xa5, 0x05, 0x5e,
        <a id="L38"></a>},
        <a id="L39"></a>nil,
    <a id="L40"></a>},
    <a id="L41"></a>zlibTest{
        <a id="L42"></a>&#34;bad header&#34;,
        <a id="L43"></a>&#34;&#34;,
        <a id="L44"></a>[]byte{0x78, 0x9f, 0x03, 0x00, 0x00, 0x00, 0x00, 0x01},
        <a id="L45"></a>HeaderError,
    <a id="L46"></a>},
    <a id="L47"></a>zlibTest{
        <a id="L48"></a>&#34;bad checksum&#34;,
        <a id="L49"></a>&#34;&#34;,
        <a id="L50"></a>[]byte{0x78, 0x9c, 0x03, 0x00, 0x00, 0x00, 0x00, 0xff},
        <a id="L51"></a>ChecksumError,
    <a id="L52"></a>},
    <a id="L53"></a>zlibTest{
        <a id="L54"></a>&#34;not enough data&#34;,
        <a id="L55"></a>&#34;&#34;,
        <a id="L56"></a>[]byte{0x78, 0x9c, 0x03, 0x00, 0x00, 0x00},
        <a id="L57"></a>io.ErrUnexpectedEOF,
    <a id="L58"></a>},
    <a id="L59"></a>zlibTest{
        <a id="L60"></a>&#34;excess data is silently ignored&#34;,
        <a id="L61"></a>&#34;&#34;,
        <a id="L62"></a>[]byte{
            <a id="L63"></a>0x78, 0x9c, 0x03, 0x00, 0x00, 0x00, 0x00, 0x01,
            <a id="L64"></a>0x78, 0x9c, 0xff,
        <a id="L65"></a>},
        <a id="L66"></a>nil,
    <a id="L67"></a>},
<a id="L68"></a>}

<a id="L70"></a>func TestInflater(t *testing.T) {
    <a id="L71"></a>b := new(bytes.Buffer);
    <a id="L72"></a>for _, tt := range zlibTests {
        <a id="L73"></a>in := bytes.NewBuffer(tt.compressed);
        <a id="L74"></a>zlib, err := NewInflater(in);
        <a id="L75"></a>if err != nil {
            <a id="L76"></a>if err != tt.err {
                <a id="L77"></a>t.Errorf(&#34;%s: NewInflater: %s&#34;, tt.desc, err)
            <a id="L78"></a>}
            <a id="L79"></a>continue;
        <a id="L80"></a>}
        <a id="L81"></a>defer zlib.Close();
        <a id="L82"></a>b.Reset();
        <a id="L83"></a>n, err := io.Copy(b, zlib);
        <a id="L84"></a>if err != nil {
            <a id="L85"></a>if err != tt.err {
                <a id="L86"></a>t.Errorf(&#34;%s: io.Copy: %v want %v&#34;, tt.desc, err, tt.err)
            <a id="L87"></a>}
            <a id="L88"></a>continue;
        <a id="L89"></a>}
        <a id="L90"></a>s := b.String();
        <a id="L91"></a>if s != tt.raw {
            <a id="L92"></a>t.Errorf(&#34;%s: got %d-byte %q want %d-byte %q&#34;, tt.desc, n, s, len(tt.raw), tt.raw)
        <a id="L93"></a>}
    <a id="L94"></a>}
<a id="L95"></a>}
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
