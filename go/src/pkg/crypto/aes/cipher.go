<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/aes/cipher.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/aes/cipher.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package aes

<a id="L7"></a>import (
    <a id="L8"></a>&#34;os&#34;;
    <a id="L9"></a>&#34;strconv&#34;;
<a id="L10"></a>)

<a id="L12"></a><span class="comment">// The AES block size in bytes.</span>
<a id="L13"></a>const BlockSize = 16

<a id="L15"></a><span class="comment">// A Cipher is an instance of AES encryption using a particular key.</span>
<a id="L16"></a>type Cipher struct {
    <a id="L17"></a>enc []uint32;
    <a id="L18"></a>dec []uint32;
<a id="L19"></a>}

<a id="L21"></a>type KeySizeError int

<a id="L23"></a>func (k KeySizeError) String() string {
    <a id="L24"></a>return &#34;crypto/aes: invalid key size &#34; + strconv.Itoa(int(k))
<a id="L25"></a>}

<a id="L27"></a><span class="comment">// NewCipher creates and returns a new Cipher.</span>
<a id="L28"></a><span class="comment">// The key argument should be the AES key,</span>
<a id="L29"></a><span class="comment">// either 16, 24, or 32 bytes to select</span>
<a id="L30"></a><span class="comment">// AES-128, AES-192, or AES-256.</span>
<a id="L31"></a>func NewCipher(key []byte) (*Cipher, os.Error) {
    <a id="L32"></a>k := len(key);
    <a id="L33"></a>switch k {
    <a id="L34"></a>default:
        <a id="L35"></a>return nil, KeySizeError(k)
    <a id="L36"></a>case 16, 24, 32:
        <a id="L37"></a>break
    <a id="L38"></a>}

    <a id="L40"></a>n := k + 28;
    <a id="L41"></a>c := &amp;Cipher{make([]uint32, n), make([]uint32, n)};
    <a id="L42"></a>expandKey(key, c.enc, c.dec);
    <a id="L43"></a>return c, nil;
<a id="L44"></a>}

<a id="L46"></a><span class="comment">// BlockSize returns the AES block size, 16 bytes.</span>
<a id="L47"></a><span class="comment">// It is necessary to satisfy the Key interface in the</span>
<a id="L48"></a><span class="comment">// package &#34;crypto/modes&#34;.</span>
<a id="L49"></a>func (c *Cipher) BlockSize() int { return BlockSize }

<a id="L51"></a><span class="comment">// Encrypt encrypts the 16-byte buffer src using the key k</span>
<a id="L52"></a><span class="comment">// and stores the result in dst.</span>
<a id="L53"></a><span class="comment">// Note that for amounts of data larger than a block,</span>
<a id="L54"></a><span class="comment">// it is not safe to just call Encrypt on successive blocks;</span>
<a id="L55"></a><span class="comment">// instead, use an encryption mode like AESCBC (see modes.go).</span>
<a id="L56"></a>func (c *Cipher) Encrypt(src, dst []byte) { encryptBlock(c.enc, src, dst) }

<a id="L58"></a><span class="comment">// Decrypt decrypts the 16-byte buffer src using the key k</span>
<a id="L59"></a><span class="comment">// and stores the result in dst.</span>
<a id="L60"></a>func (c *Cipher) Decrypt(src, dst []byte) { decryptBlock(c.dec, src, dst) }

<a id="L62"></a><span class="comment">// Reset zeros the key data, so that it will no longer</span>
<a id="L63"></a><span class="comment">// appear in the process&#39;s memory.</span>
<a id="L64"></a>func (c *Cipher) Reset() {
    <a id="L65"></a>for i := 0; i &lt; len(c.enc); i++ {
        <a id="L66"></a>c.enc[i] = 0
    <a id="L67"></a>}
    <a id="L68"></a>for i := 0; i &lt; len(c.dec); i++ {
        <a id="L69"></a>c.dec[i] = 0
    <a id="L70"></a>}
<a id="L71"></a>}
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
