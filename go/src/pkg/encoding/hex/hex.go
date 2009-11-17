<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/encoding/hex/hex.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/encoding/hex/hex.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package implements hexadecimal encoding and decoding.</span>
<a id="L6"></a>package hex

<a id="L8"></a>import (
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>&#34;strconv&#34;;
    <a id="L11"></a>&#34;strings&#34;;
<a id="L12"></a>)

<a id="L14"></a>const hextable = &#34;0123456789abcdef&#34;

<a id="L16"></a><span class="comment">// EncodedLen returns the length of an encoding of n source bytes.</span>
<a id="L17"></a>func EncodedLen(n int) int { return n * 2 }

<a id="L19"></a><span class="comment">// Encode encodes src into EncodedLen(len(src))</span>
<a id="L20"></a><span class="comment">// bytes of dst.  As a convenience, it returns the number</span>
<a id="L21"></a><span class="comment">// of bytes written to dst, but this value is always EncodedLen(len(src)).</span>
<a id="L22"></a><span class="comment">// Encode implements hexadecimal encoding.</span>
<a id="L23"></a>func Encode(dst, src []byte) int {
    <a id="L24"></a>for i, v := range src {
        <a id="L25"></a>dst[i*2] = hextable[v&gt;&gt;4];
        <a id="L26"></a>dst[i*2+1] = hextable[v&amp;0x0f];
    <a id="L27"></a>}

    <a id="L29"></a>return len(src) * 2;
<a id="L30"></a>}

<a id="L32"></a><span class="comment">// OddLengthInputError results from decoding an odd length slice.</span>
<a id="L33"></a>type OddLengthInputError struct{}

<a id="L35"></a>func (OddLengthInputError) String() string { return &#34;odd length hex string&#34; }

<a id="L37"></a><span class="comment">// InvalidHexCharError results from finding an invalid character in a hex string.</span>
<a id="L38"></a>type InvalidHexCharError byte

<a id="L40"></a>func (e InvalidHexCharError) String() string {
    <a id="L41"></a>return &#34;invalid hex char: &#34; + strconv.Itoa(int(e))
<a id="L42"></a>}


<a id="L45"></a>func DecodedLen(x int) int { return x / 2 }

<a id="L47"></a><span class="comment">// Decode decodes src into DecodedLen(len(src)) bytes, returning the actual</span>
<a id="L48"></a><span class="comment">// number of bytes written to dst.</span>
<a id="L49"></a><span class="comment">//</span>
<a id="L50"></a><span class="comment">// If Decode encounters invalid input, it returns an OddLengthInputError or an</span>
<a id="L51"></a><span class="comment">// InvalidHexCharError.</span>
<a id="L52"></a>func Decode(dst, src []byte) (int, os.Error) {
    <a id="L53"></a>if len(src)%2 == 1 {
        <a id="L54"></a>return 0, OddLengthInputError{}
    <a id="L55"></a>}

    <a id="L57"></a>for i := 0; i &lt; len(src)/2; i++ {
        <a id="L58"></a>a, ok := fromHexChar(src[i*2]);
        <a id="L59"></a>if !ok {
            <a id="L60"></a>return 0, InvalidHexCharError(src[i*2])
        <a id="L61"></a>}
        <a id="L62"></a>b, ok := fromHexChar(src[i*2+1]);
        <a id="L63"></a>if !ok {
            <a id="L64"></a>return 0, InvalidHexCharError(src[i*2+1])
        <a id="L65"></a>}
        <a id="L66"></a>dst[i] = (a &lt;&lt; 4) | b;
    <a id="L67"></a>}

    <a id="L69"></a>return len(src) / 2, nil;
<a id="L70"></a>}

<a id="L72"></a><span class="comment">// fromHexChar converts a hex character into its value and a success flag.</span>
<a id="L73"></a>func fromHexChar(c byte) (byte, bool) {
    <a id="L74"></a>switch {
    <a id="L75"></a>case 0 &lt;= c &amp;&amp; c &lt;= &#39;9&#39;:
        <a id="L76"></a>return c - &#39;0&#39;, true
    <a id="L77"></a>case &#39;a&#39; &lt;= c &amp;&amp; c &lt;= &#39;f&#39;:
        <a id="L78"></a>return c - &#39;a&#39; + 10, true
    <a id="L79"></a>case &#39;A&#39; &lt;= c &amp;&amp; c &lt;= &#39;F&#39;:
        <a id="L80"></a>return c - &#39;A&#39; + 10, true
    <a id="L81"></a>}

    <a id="L83"></a>return 0, false;
<a id="L84"></a>}

<a id="L86"></a><span class="comment">// EncodeToString returns the hexadecimal encoding of src.</span>
<a id="L87"></a>func EncodeToString(src []byte) string {
    <a id="L88"></a>dst := make([]byte, EncodedLen(len(src)));
    <a id="L89"></a>Encode(dst, src);
    <a id="L90"></a>return string(dst);
<a id="L91"></a>}

<a id="L93"></a><span class="comment">// DecodeString returns the bytes represented by the hexadecimal string s.</span>
<a id="L94"></a>func DecodeString(s string) ([]byte, os.Error) {
    <a id="L95"></a>src := strings.Bytes(s);
    <a id="L96"></a>dst := make([]byte, DecodedLen(len(src)));
    <a id="L97"></a>_, err := Decode(dst, src);
    <a id="L98"></a>if err != nil {
        <a id="L99"></a>return nil, err
    <a id="L100"></a>}
    <a id="L101"></a>return dst, nil;
<a id="L102"></a>}
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
