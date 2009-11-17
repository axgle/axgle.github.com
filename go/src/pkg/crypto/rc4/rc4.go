<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/rc4/rc4.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/rc4/rc4.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package implements RC4 encryption, as defined in Bruce Schneier&#39;s</span>
<a id="L6"></a><span class="comment">// Applied Cryptography.</span>
<a id="L7"></a>package rc4

<a id="L9"></a><span class="comment">// BUG(agl): RC4 is in common use but has design weaknesses that make</span>
<a id="L10"></a><span class="comment">// it a poor choice for new protocols.</span>

<a id="L12"></a>import (
    <a id="L13"></a>&#34;os&#34;;
    <a id="L14"></a>&#34;strconv&#34;;
<a id="L15"></a>)

<a id="L17"></a><span class="comment">// A Cipher is an instance of RC4 using a particular key.</span>
<a id="L18"></a>type Cipher struct {
    <a id="L19"></a>s    [256]byte;
    <a id="L20"></a>i, j uint8;
<a id="L21"></a>}

<a id="L23"></a>type KeySizeError int

<a id="L25"></a>func (k KeySizeError) String() string {
    <a id="L26"></a>return &#34;crypto/rc4: invalid key size &#34; + strconv.Itoa(int(k))
<a id="L27"></a>}

<a id="L29"></a><span class="comment">// NewCipher creates and returns a new Cipher.  The key argument should be the</span>
<a id="L30"></a><span class="comment">// RC4 key, at least 1 byte and at most 256 bytes.</span>
<a id="L31"></a>func NewCipher(key []byte) (*Cipher, os.Error) {
    <a id="L32"></a>k := len(key);
    <a id="L33"></a>if k &lt; 1 || k &gt; 256 {
        <a id="L34"></a>return nil, KeySizeError(k)
    <a id="L35"></a>}
    <a id="L36"></a>var c Cipher;
    <a id="L37"></a>for i := 0; i &lt; 256; i++ {
        <a id="L38"></a>c.s[i] = uint8(i)
    <a id="L39"></a>}
    <a id="L40"></a>var j uint8 = 0;
    <a id="L41"></a>for i := 0; i &lt; 256; i++ {
        <a id="L42"></a>j += c.s[i] + key[i%k];
        <a id="L43"></a>c.s[i], c.s[j] = c.s[j], c.s[i];
    <a id="L44"></a>}
    <a id="L45"></a>return &amp;c, nil;
<a id="L46"></a>}

<a id="L48"></a><span class="comment">// XORKeyStream will XOR each byte of the given buffer with a byte of the</span>
<a id="L49"></a><span class="comment">// generated keystream.</span>
<a id="L50"></a>func (c *Cipher) XORKeyStream(buf []byte) {
    <a id="L51"></a>for i := range buf {
        <a id="L52"></a>c.i += 1;
        <a id="L53"></a>c.j += c.s[c.i];
        <a id="L54"></a>c.s[c.i], c.s[c.j] = c.s[c.j], c.s[c.i];
        <a id="L55"></a>buf[i] ^= c.s[c.s[c.i]+c.s[c.j]];
    <a id="L56"></a>}
<a id="L57"></a>}

<a id="L59"></a><span class="comment">// Reset zeros the key data so that it will no longer appear in the</span>
<a id="L60"></a><span class="comment">// process&#39;s memory.</span>
<a id="L61"></a>func (c *Cipher) Reset() {
    <a id="L62"></a>for i := range c.s {
        <a id="L63"></a>c.s[i] = 0
    <a id="L64"></a>}
    <a id="L65"></a>c.i, c.j = 0, 0;
<a id="L66"></a>}
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
