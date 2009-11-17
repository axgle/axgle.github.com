<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/unicode/digit_test.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/unicode/digit_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package unicode_test

<a id="L7"></a>import (
    <a id="L8"></a>&#34;testing&#34;;
    <a id="L9"></a>. &#34;unicode&#34;;
<a id="L10"></a>)

<a id="L12"></a>var testDigit = []int{
    <a id="L13"></a>0x0030,
    <a id="L14"></a>0x0039,
    <a id="L15"></a>0x0661,
    <a id="L16"></a>0x06F1,
    <a id="L17"></a>0x07C9,
    <a id="L18"></a>0x0966,
    <a id="L19"></a>0x09EF,
    <a id="L20"></a>0x0A66,
    <a id="L21"></a>0x0AEF,
    <a id="L22"></a>0x0B66,
    <a id="L23"></a>0x0B6F,
    <a id="L24"></a>0x0BE6,
    <a id="L25"></a>0x0BEF,
    <a id="L26"></a>0x0C66,
    <a id="L27"></a>0x0CEF,
    <a id="L28"></a>0x0D66,
    <a id="L29"></a>0x0D6F,
    <a id="L30"></a>0x0E50,
    <a id="L31"></a>0x0E59,
    <a id="L32"></a>0x0ED0,
    <a id="L33"></a>0x0ED9,
    <a id="L34"></a>0x0F20,
    <a id="L35"></a>0x0F29,
    <a id="L36"></a>0x1040,
    <a id="L37"></a>0x1049,
    <a id="L38"></a>0x1090,
    <a id="L39"></a>0x1091,
    <a id="L40"></a>0x1099,
    <a id="L41"></a>0x17E0,
    <a id="L42"></a>0x17E9,
    <a id="L43"></a>0x1810,
    <a id="L44"></a>0x1819,
    <a id="L45"></a>0x1946,
    <a id="L46"></a>0x194F,
    <a id="L47"></a>0x19D0,
    <a id="L48"></a>0x19D9,
    <a id="L49"></a>0x1B50,
    <a id="L50"></a>0x1B59,
    <a id="L51"></a>0x1BB0,
    <a id="L52"></a>0x1BB9,
    <a id="L53"></a>0x1C40,
    <a id="L54"></a>0x1C49,
    <a id="L55"></a>0x1C50,
    <a id="L56"></a>0x1C59,
    <a id="L57"></a>0xA620,
    <a id="L58"></a>0xA629,
    <a id="L59"></a>0xA8D0,
    <a id="L60"></a>0xA8D9,
    <a id="L61"></a>0xA900,
    <a id="L62"></a>0xA909,
    <a id="L63"></a>0xAA50,
    <a id="L64"></a>0xAA59,
    <a id="L65"></a>0xFF10,
    <a id="L66"></a>0xFF19,
    <a id="L67"></a>0x104A1,
    <a id="L68"></a>0x1D7CE,
<a id="L69"></a>}

<a id="L71"></a>var testLetter = []int{
    <a id="L72"></a>0x0041,
    <a id="L73"></a>0x0061,
    <a id="L74"></a>0x00AA,
    <a id="L75"></a>0x00BA,
    <a id="L76"></a>0x00C8,
    <a id="L77"></a>0x00DB,
    <a id="L78"></a>0x00F9,
    <a id="L79"></a>0x02EC,
    <a id="L80"></a>0x0535,
    <a id="L81"></a>0x06E6,
    <a id="L82"></a>0x093D,
    <a id="L83"></a>0x0A15,
    <a id="L84"></a>0x0B99,
    <a id="L85"></a>0x0DC0,
    <a id="L86"></a>0x0EDD,
    <a id="L87"></a>0x1000,
    <a id="L88"></a>0x1200,
    <a id="L89"></a>0x1312,
    <a id="L90"></a>0x1401,
    <a id="L91"></a>0x1885,
    <a id="L92"></a>0x2C00,
    <a id="L93"></a>0xA800,
    <a id="L94"></a>0xF900,
    <a id="L95"></a>0xFA30,
    <a id="L96"></a>0xFFDA,
    <a id="L97"></a>0xFFDC,
    <a id="L98"></a>0x10000,
    <a id="L99"></a>0x10300,
    <a id="L100"></a>0x10400,
    <a id="L101"></a>0x20000,
    <a id="L102"></a>0x2F800,
    <a id="L103"></a>0x2FA1D,
<a id="L104"></a>}

<a id="L106"></a>func TestDigit(t *testing.T) {
    <a id="L107"></a>for _, r := range testDigit {
        <a id="L108"></a>if !IsDigit(r) {
            <a id="L109"></a>t.Errorf(&#34;IsDigit(U+%04X) = false, want true\n&#34;, r)
        <a id="L110"></a>}
    <a id="L111"></a>}
    <a id="L112"></a>for _, r := range testLetter {
        <a id="L113"></a>if IsDigit(r) {
            <a id="L114"></a>t.Errorf(&#34;IsDigit(U+%04X) = true, want false\n&#34;, r)
        <a id="L115"></a>}
    <a id="L116"></a>}
<a id="L117"></a>}

<a id="L119"></a><span class="comment">// Test that the special case in IsDigit agrees with the table</span>
<a id="L120"></a>func TestDigitOptimization(t *testing.T) {
    <a id="L121"></a>for i := 0; i &lt; 0x100; i++ {
        <a id="L122"></a>if Is(Digit, i) != IsDigit(i) {
            <a id="L123"></a>t.Errorf(&#34;IsDigit(U+%04X) disagrees with Is(Digit)&#34;, i)
        <a id="L124"></a>}
    <a id="L125"></a>}
<a id="L126"></a>}
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
