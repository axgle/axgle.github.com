<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/hmac/hmac_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/hmac/hmac_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package hmac

<a id="L7"></a>import (
    <a id="L8"></a>&#34;hash&#34;;
    <a id="L9"></a>&#34;fmt&#34;;
    <a id="L10"></a>&#34;strings&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)

<a id="L14"></a>type hmacTest struct {
    <a id="L15"></a>hash func([]byte) hash.Hash;
    <a id="L16"></a>key  []byte;
    <a id="L17"></a>in   []byte;
    <a id="L18"></a>out  string;
<a id="L19"></a>}

<a id="L21"></a><span class="comment">// Tests from US FIPS 198</span>
<a id="L22"></a><span class="comment">// http://csrc.nist.gov/publications/fips/fips198/fips-198a.pdf</span>
<a id="L23"></a>var hmacTests = []hmacTest{
    <a id="L24"></a>hmacTest{
        <a id="L25"></a>NewSHA1,
        <a id="L26"></a>[]byte{
            <a id="L27"></a>0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
            <a id="L28"></a>0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
            <a id="L29"></a>0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
            <a id="L30"></a>0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
            <a id="L31"></a>0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27,
            <a id="L32"></a>0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
            <a id="L33"></a>0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
            <a id="L34"></a>0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
        <a id="L35"></a>},
        <a id="L36"></a>strings.Bytes(&#34;Sample #1&#34;),
        <a id="L37"></a>&#34;4f4ca3d5d68ba7cc0a1208c9c61e9c5da0403c0a&#34;,
    <a id="L38"></a>},
    <a id="L39"></a>hmacTest{
        <a id="L40"></a>NewSHA1,
        <a id="L41"></a>[]byte{
            <a id="L42"></a>0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37,
            <a id="L43"></a>0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
            <a id="L44"></a>0x40, 0x41, 0x42, 0x43,
        <a id="L45"></a>},
        <a id="L46"></a>strings.Bytes(&#34;Sample #2&#34;),
        <a id="L47"></a>&#34;0922d3405faa3d194f82a45830737d5cc6c75d24&#34;,
    <a id="L48"></a>},
    <a id="L49"></a>hmacTest{
        <a id="L50"></a>NewSHA1,
        <a id="L51"></a>[]byte{
            <a id="L52"></a>0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57,
            <a id="L53"></a>0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f,
            <a id="L54"></a>0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67,
            <a id="L55"></a>0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f,
            <a id="L56"></a>0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77,
            <a id="L57"></a>0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f,
            <a id="L58"></a>0x80, 0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87,
            <a id="L59"></a>0x88, 0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f,
            <a id="L60"></a>0x90, 0x91, 0x92, 0x93, 0x94, 0x95, 0x96, 0x97,
            <a id="L61"></a>0x98, 0x99, 0x9a, 0x9b, 0x9c, 0x9d, 0x9e, 0x9f,
            <a id="L62"></a>0xa0, 0xa1, 0xa2, 0xa3, 0xa4, 0xa5, 0xa6, 0xa7,
            <a id="L63"></a>0xa8, 0xa9, 0xaa, 0xab, 0xac, 0xad, 0xae, 0xaf,
            <a id="L64"></a>0xb0, 0xb1, 0xb2, 0xb3,
        <a id="L65"></a>},
        <a id="L66"></a>strings.Bytes(&#34;Sample #3&#34;),
        <a id="L67"></a>&#34;bcf41eab8bb2d802f3d05caf7cb092ecf8d1a3aa&#34;,
    <a id="L68"></a>},

    <a id="L70"></a><span class="comment">// Test from Plan 9.</span>
    <a id="L71"></a>hmacTest{
        <a id="L72"></a>NewMD5,
        <a id="L73"></a>strings.Bytes(&#34;Jefe&#34;),
        <a id="L74"></a>strings.Bytes(&#34;what do ya want for nothing?&#34;),
        <a id="L75"></a>&#34;750c783e6ab0b503eaa86e310a5db738&#34;,
    <a id="L76"></a>},
<a id="L77"></a>}

<a id="L79"></a>func TestHMAC(t *testing.T) {
    <a id="L80"></a>for i, tt := range hmacTests {
        <a id="L81"></a>h := tt.hash(tt.key);
        <a id="L82"></a>for j := 0; j &lt; 2; j++ {
            <a id="L83"></a>n, err := h.Write(tt.in);
            <a id="L84"></a>if n != len(tt.in) || err != nil {
                <a id="L85"></a>t.Errorf(&#34;test %d.%d: Write(%d) = %d, %v&#34;, i, j, len(tt.in), n, err);
                <a id="L86"></a>continue;
            <a id="L87"></a>}
            <a id="L88"></a>sum := fmt.Sprintf(&#34;%x&#34;, h.Sum());
            <a id="L89"></a>if sum != tt.out {
                <a id="L90"></a>t.Errorf(&#34;test %d.%d: have %s want %s\n&#34;, i, j, sum, tt.out)
            <a id="L91"></a>}

            <a id="L93"></a><span class="comment">// Second iteration: make sure reset works.</span>
            <a id="L94"></a>h.Reset();
        <a id="L95"></a>}
    <a id="L96"></a>}
<a id="L97"></a>}
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
