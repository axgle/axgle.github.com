<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/rc4/rc4_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/rc4/rc4_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package rc4

<a id="L7"></a>import (
    <a id="L8"></a>&#34;testing&#34;;
<a id="L9"></a>)

<a id="L11"></a>type rc4Test struct {
    <a id="L12"></a>key, keystream []byte;
<a id="L13"></a>}

<a id="L15"></a>var golden = []rc4Test{
    <a id="L16"></a><span class="comment">// Test vectors from the original cypherpunk posting of ARC4:</span>
    <a id="L17"></a><span class="comment">//   http://groups.google.com/group/sci.crypt/msg/10a300c9d21afca0?pli=1</span>
    <a id="L18"></a>rc4Test{
        <a id="L19"></a>[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
        <a id="L20"></a>[]byte{0x74, 0x94, 0xc2, 0xe7, 0x10, 0x4b, 0x08, 0x79},
    <a id="L21"></a>},
    <a id="L22"></a>rc4Test{
        <a id="L23"></a>[]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
        <a id="L24"></a>[]byte{0xde, 0x18, 0x89, 0x41, 0xa3, 0x37, 0x5d, 0x3a},
    <a id="L25"></a>},
    <a id="L26"></a>rc4Test{
        <a id="L27"></a>[]byte{0xef, 0x01, 0x23, 0x45},
        <a id="L28"></a>[]byte{0xd6, 0xa1, 0x41, 0xa7, 0xec, 0x3c, 0x38, 0xdf, 0xbd, 0x61},
    <a id="L29"></a>},

    <a id="L31"></a><span class="comment">// Test vectors from the Wikipedia page: http://en.wikipedia.org/wiki/RC4</span>
    <a id="L32"></a>rc4Test{
        <a id="L33"></a>[]byte{0x4b, 0x65, 0x79},
        <a id="L34"></a>[]byte{0xeb, 0x9f, 0x77, 0x81, 0xb7, 0x34, 0xca, 0x72, 0xa7, 0x19},
    <a id="L35"></a>},
    <a id="L36"></a>rc4Test{
        <a id="L37"></a>[]byte{0x57, 0x69, 0x6b, 0x69},
        <a id="L38"></a>[]byte{0x60, 0x44, 0xdb, 0x6d, 0x41, 0xb7},
    <a id="L39"></a>},
<a id="L40"></a>}

<a id="L42"></a>func TestGolden(t *testing.T) {
    <a id="L43"></a>for i := 0; i &lt; len(golden); i++ {
        <a id="L44"></a>g := golden[i];
        <a id="L45"></a>c, err := NewCipher(g.key);
        <a id="L46"></a>if err != nil {
            <a id="L47"></a>t.Errorf(&#34;Failed to create cipher at golden index %d&#34;, i);
            <a id="L48"></a>return;
        <a id="L49"></a>}
        <a id="L50"></a>keystream := make([]byte, len(g.keystream));
        <a id="L51"></a>c.XORKeyStream(keystream);
        <a id="L52"></a>for j, v := range keystream {
            <a id="L53"></a>if g.keystream[j] != v {
                <a id="L54"></a>t.Errorf(&#34;Failed at golden index %d&#34;, i);
                <a id="L55"></a>break;
            <a id="L56"></a>}
        <a id="L57"></a>}
    <a id="L58"></a>}
<a id="L59"></a>}
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
