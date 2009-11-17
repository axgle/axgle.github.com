<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/hmac/hmac.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/hmac/hmac.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The hmac package implements the Keyed-Hash Message Authentication Code (HMAC)</span>
<a id="L6"></a><span class="comment">// as defined in U.S. Federal Information Processing Standards Publication 198.</span>
<a id="L7"></a><span class="comment">// An HMAC is a cryptographic hash that uses a key to sign a message.</span>
<a id="L8"></a><span class="comment">// The receiver verifies the hash by recomputing it using the same key.</span>
<a id="L9"></a>package hmac

<a id="L11"></a>import (
    <a id="L12"></a>&#34;crypto/md5&#34;;
    <a id="L13"></a>&#34;crypto/sha1&#34;;
    <a id="L14"></a>&#34;hash&#34;;
    <a id="L15"></a>&#34;os&#34;;
<a id="L16"></a>)

<a id="L18"></a><span class="comment">// FIPS 198:</span>
<a id="L19"></a><span class="comment">// http://csrc.nist.gov/publications/fips/fips198/fips-198a.pdf</span>

<a id="L21"></a><span class="comment">// key is zero padded to 64 bytes</span>
<a id="L22"></a><span class="comment">// ipad = 0x36 byte repeated to 64 bytes</span>
<a id="L23"></a><span class="comment">// opad = 0x5c byte repeated to 64 bytes</span>
<a id="L24"></a><span class="comment">// hmac = H([key ^ opad] H([key ^ ipad] text))</span>

<a id="L26"></a>const (
    <a id="L27"></a><span class="comment">// NOTE(rsc): This constant is actually the</span>
    <a id="L28"></a><span class="comment">// underlying hash function&#39;s block size.</span>
    <a id="L29"></a><span class="comment">// HMAC is only conventionally used with</span>
    <a id="L30"></a><span class="comment">// MD5 and SHA1, and both use 64-byte blocks.</span>
    <a id="L31"></a><span class="comment">// The hash.Hash interface doesn&#39;t provide a</span>
    <a id="L32"></a><span class="comment">// way to find out the block size.</span>
    <a id="L33"></a>padSize = 64;
<a id="L34"></a>)

<a id="L36"></a>type hmac struct {
    <a id="L37"></a>size  int;
    <a id="L38"></a>key   []byte;
    <a id="L39"></a>tmp   []byte;
    <a id="L40"></a>inner hash.Hash;
<a id="L41"></a>}

<a id="L43"></a>func (h *hmac) tmpPad(xor byte) {
    <a id="L44"></a>for i, k := range h.key {
        <a id="L45"></a>h.tmp[i] = xor ^ k
    <a id="L46"></a>}
    <a id="L47"></a>for i := len(h.key); i &lt; padSize; i++ {
        <a id="L48"></a>h.tmp[i] = xor
    <a id="L49"></a>}
<a id="L50"></a>}

<a id="L52"></a>func (h *hmac) Sum() []byte {
    <a id="L53"></a>h.tmpPad(0x5c);
    <a id="L54"></a>sum := h.inner.Sum();
    <a id="L55"></a>for i, b := range sum {
        <a id="L56"></a>h.tmp[padSize+i] = b
    <a id="L57"></a>}
    <a id="L58"></a>h.inner.Reset();
    <a id="L59"></a>h.inner.Write(h.tmp);
    <a id="L60"></a>return h.inner.Sum();
<a id="L61"></a>}

<a id="L63"></a>func (h *hmac) Write(p []byte) (n int, err os.Error) {
    <a id="L64"></a>return h.inner.Write(p)
<a id="L65"></a>}

<a id="L67"></a>func (h *hmac) Size() int { return h.size }

<a id="L69"></a>func (h *hmac) Reset() {
    <a id="L70"></a>h.inner.Reset();
    <a id="L71"></a>h.tmpPad(0x36);
    <a id="L72"></a>h.inner.Write(h.tmp[0:padSize]);
<a id="L73"></a>}

<a id="L75"></a><span class="comment">// New returns a new HMAC hash using the given hash and key.</span>
<a id="L76"></a>func New(h hash.Hash, key []byte) hash.Hash {
    <a id="L77"></a>if len(key) &gt; padSize {
        <a id="L78"></a><span class="comment">// If key is too big, hash it.</span>
        <a id="L79"></a>h.Write(key);
        <a id="L80"></a>key = h.Sum();
    <a id="L81"></a>}
    <a id="L82"></a>hm := new(hmac);
    <a id="L83"></a>hm.inner = h;
    <a id="L84"></a>hm.size = h.Size();
    <a id="L85"></a>hm.key = make([]byte, len(key));
    <a id="L86"></a>for i, k := range key {
        <a id="L87"></a>hm.key[i] = k
    <a id="L88"></a>}
    <a id="L89"></a>hm.tmp = make([]byte, padSize+hm.size);
    <a id="L90"></a>hm.Reset();
    <a id="L91"></a>return hm;
<a id="L92"></a>}

<a id="L94"></a><span class="comment">// NewMD5 returns a new HMAC-MD5 hash using the given key.</span>
<a id="L95"></a>func NewMD5(key []byte) hash.Hash { return New(md5.New(), key) }

<a id="L97"></a><span class="comment">// NewSHA1 returns a new HMAC-SHA1 hash using the given key.</span>
<a id="L98"></a>func NewSHA1(key []byte) hash.Hash { return New(sha1.New(), key) }
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
