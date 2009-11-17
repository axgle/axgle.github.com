<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/cfb.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/cfb.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Cipher feedback (CFB) mode.</span>

<a id="L7"></a><span class="comment">// CFB provides confidentiality by feeding a fraction of</span>
<a id="L8"></a><span class="comment">// the previous ciphertext in as the plaintext for the next</span>
<a id="L9"></a><span class="comment">// block operation.</span>

<a id="L11"></a><span class="comment">// See NIST SP 800-38A, pp 11-13</span>

<a id="L13"></a>package block

<a id="L15"></a>import (
    <a id="L16"></a>&#34;io&#34;;
<a id="L17"></a>)

<a id="L19"></a>type cfbCipher struct {
    <a id="L20"></a>c          Cipher;
    <a id="L21"></a>blockSize  int; <span class="comment">// our block size (s/8)</span>
    <a id="L22"></a>cipherSize int; <span class="comment">// underlying cipher block size</span>
    <a id="L23"></a>iv         []byte;
    <a id="L24"></a>tmp        []byte;
<a id="L25"></a>}

<a id="L27"></a>func newCFB(c Cipher, s int, iv []byte) *cfbCipher {
    <a id="L28"></a>if s == 0 || s%8 != 0 {
        <a id="L29"></a>panicln(&#34;crypto/block: invalid CFB mode&#34;, s)
    <a id="L30"></a>}
    <a id="L31"></a>b := c.BlockSize();
    <a id="L32"></a>x := new(cfbCipher);
    <a id="L33"></a>x.c = c;
    <a id="L34"></a>x.blockSize = s / 8;
    <a id="L35"></a>x.cipherSize = b;
    <a id="L36"></a>x.iv = copy(iv);
    <a id="L37"></a>x.tmp = make([]byte, b);
    <a id="L38"></a>return x;
<a id="L39"></a>}

<a id="L41"></a>func (x *cfbCipher) BlockSize() int { return x.blockSize }

<a id="L43"></a>func (x *cfbCipher) Encrypt(src, dst []byte) {
    <a id="L44"></a><span class="comment">// Encrypt old IV and xor prefix with src to make dst.</span>
    <a id="L45"></a>x.c.Encrypt(x.iv, x.tmp);
    <a id="L46"></a>for i := 0; i &lt; x.blockSize; i++ {
        <a id="L47"></a>dst[i] = src[i] ^ x.tmp[i]
    <a id="L48"></a>}

    <a id="L50"></a><span class="comment">// Slide unused IV pieces down and insert dst at end.</span>
    <a id="L51"></a>for i := 0; i &lt; x.cipherSize-x.blockSize; i++ {
        <a id="L52"></a>x.iv[i] = x.iv[i+x.blockSize]
    <a id="L53"></a>}
    <a id="L54"></a>off := x.cipherSize - x.blockSize;
    <a id="L55"></a>for i := off; i &lt; x.cipherSize; i++ {
        <a id="L56"></a>x.iv[i] = dst[i-off]
    <a id="L57"></a>}
<a id="L58"></a>}

<a id="L60"></a>func (x *cfbCipher) Decrypt(src, dst []byte) {
    <a id="L61"></a><span class="comment">// Encrypt [sic] old IV and xor prefix with src to make dst.</span>
    <a id="L62"></a>x.c.Encrypt(x.iv, x.tmp);
    <a id="L63"></a>for i := 0; i &lt; x.blockSize; i++ {
        <a id="L64"></a>dst[i] = src[i] ^ x.tmp[i]
    <a id="L65"></a>}

    <a id="L67"></a><span class="comment">// Slide unused IV pieces down and insert src at top.</span>
    <a id="L68"></a>for i := 0; i &lt; x.cipherSize-x.blockSize; i++ {
        <a id="L69"></a>x.iv[i] = x.iv[i+x.blockSize]
    <a id="L70"></a>}
    <a id="L71"></a>off := x.cipherSize - x.blockSize;
    <a id="L72"></a>for i := off; i &lt; x.cipherSize; i++ {
        <a id="L73"></a><span class="comment">// Reconstruct src = dst ^ x.tmp</span>
        <a id="L74"></a><span class="comment">// in case we overwrote src (src == dst).</span>
        <a id="L75"></a>x.iv[i] = dst[i-off] ^ x.tmp[i-off]
    <a id="L76"></a>}
<a id="L77"></a>}

<a id="L79"></a><span class="comment">// NewCFBDecrypter returns a reader that reads data from r and decrypts it using c</span>
<a id="L80"></a><span class="comment">// in s-bit cipher feedback (CFB) mode with the initialization vector iv.</span>
<a id="L81"></a><span class="comment">// The returned Reader does not buffer or read ahead except</span>
<a id="L82"></a><span class="comment">// as required by the cipher&#39;s block size.</span>
<a id="L83"></a><span class="comment">// Modes for s not a multiple of 8 are unimplemented.</span>
<a id="L84"></a>func NewCFBDecrypter(c Cipher, s int, iv []byte, r io.Reader) io.Reader {
    <a id="L85"></a>return NewECBDecrypter(newCFB(c, s, iv), r)
<a id="L86"></a>}

<a id="L88"></a><span class="comment">// NewCFBEncrypter returns a writer that encrypts data using c</span>
<a id="L89"></a><span class="comment">// in s-bit cipher feedback (CFB) mode with the initialization vector iv</span>
<a id="L90"></a><span class="comment">// and writes the encrypted data to w.</span>
<a id="L91"></a><span class="comment">// The returned Writer does no buffering except as required</span>
<a id="L92"></a><span class="comment">// by the cipher&#39;s block size, so there is no need for a Flush method.</span>
<a id="L93"></a><span class="comment">// Modes for s not a multiple of 8 are unimplemented.</span>
<a id="L94"></a>func NewCFBEncrypter(c Cipher, s int, iv []byte, w io.Writer) io.Writer {
    <a id="L95"></a>return NewECBEncrypter(newCFB(c, s, iv), w)
<a id="L96"></a>}
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
