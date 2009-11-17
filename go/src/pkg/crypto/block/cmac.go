<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/cmac.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/cmac.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// CMAC message authentication code, defined in</span>
<a id="L6"></a><span class="comment">// NIST Special Publication SP 800-38B.</span>

<a id="L8"></a>package block

<a id="L10"></a>import (
    <a id="L11"></a>&#34;hash&#34;;
    <a id="L12"></a>&#34;os&#34;;
<a id="L13"></a>)

<a id="L15"></a>const (
    <a id="L16"></a><span class="comment">// minimal irreducible polynomial of degree b</span>
    <a id="L17"></a>r64  = 0x1b;
    <a id="L18"></a>r128 = 0x87;
<a id="L19"></a>)

<a id="L21"></a>type cmac struct {
    <a id="L22"></a>k1, k2, ci, digest []byte;
    <a id="L23"></a>p                  int; <span class="comment">// position in ci</span>
    <a id="L24"></a>c                  Cipher;
<a id="L25"></a>}

<a id="L27"></a><span class="comment">// TODO(rsc): Should this return an error instead of panic?</span>

<a id="L29"></a><span class="comment">// NewCMAC returns a new instance of a CMAC message authentication code</span>
<a id="L30"></a><span class="comment">// digest using the given Cipher.</span>
<a id="L31"></a>func NewCMAC(c Cipher) hash.Hash {
    <a id="L32"></a>var r byte;
    <a id="L33"></a>n := c.BlockSize();
    <a id="L34"></a>switch n {
    <a id="L35"></a>case 64 / 8:
        <a id="L36"></a>r = r64
    <a id="L37"></a>case 128 / 8:
        <a id="L38"></a>r = r128
    <a id="L39"></a>default:
        <a id="L40"></a>panic(&#34;crypto/block: NewCMAC: invalid cipher block size&#34;, n)
    <a id="L41"></a>}

    <a id="L43"></a>d := new(cmac);
    <a id="L44"></a>d.c = c;
    <a id="L45"></a>d.k1 = make([]byte, n);
    <a id="L46"></a>d.k2 = make([]byte, n);
    <a id="L47"></a>d.ci = make([]byte, n);
    <a id="L48"></a>d.digest = make([]byte, n);

    <a id="L50"></a><span class="comment">// Subkey generation, p. 7</span>
    <a id="L51"></a>c.Encrypt(d.k1, d.k1);
    <a id="L52"></a>if shift1(d.k1, d.k1) != 0 {
        <a id="L53"></a>d.k1[n-1] ^= r
    <a id="L54"></a>}
    <a id="L55"></a>if shift1(d.k1, d.k2) != 0 {
        <a id="L56"></a>d.k2[n-1] ^= r
    <a id="L57"></a>}

    <a id="L59"></a>return d;
<a id="L60"></a>}

<a id="L62"></a><span class="comment">// Reset clears the digest state, starting a new digest.</span>
<a id="L63"></a>func (d *cmac) Reset() {
    <a id="L64"></a>for i := range d.ci {
        <a id="L65"></a>d.ci[i] = 0
    <a id="L66"></a>}
    <a id="L67"></a>d.p = 0;
<a id="L68"></a>}

<a id="L70"></a><span class="comment">// Write adds the given data to the digest state.</span>
<a id="L71"></a>func (d *cmac) Write(p []byte) (n int, err os.Error) {
    <a id="L72"></a><span class="comment">// Xor input into ci.</span>
    <a id="L73"></a>for _, c := range p {
        <a id="L74"></a><span class="comment">// If ci is full, encrypt and start over.</span>
        <a id="L75"></a>if d.p &gt;= len(d.ci) {
            <a id="L76"></a>d.c.Encrypt(d.ci, d.ci);
            <a id="L77"></a>d.p = 0;
        <a id="L78"></a>}
        <a id="L79"></a>d.ci[d.p] ^= c;
        <a id="L80"></a>d.p++;
    <a id="L81"></a>}
    <a id="L82"></a>return len(p), nil;
<a id="L83"></a>}

<a id="L85"></a><span class="comment">// Sum returns the CMAC digest, one cipher block in length,</span>
<a id="L86"></a><span class="comment">// of the data written with Write.</span>
<a id="L87"></a>func (d *cmac) Sum() []byte {
    <a id="L88"></a><span class="comment">// Finish last block, mix in key, encrypt.</span>
    <a id="L89"></a><span class="comment">// Don&#39;t edit ci, in case caller wants</span>
    <a id="L90"></a><span class="comment">// to keep digesting after call to Sum.</span>
    <a id="L91"></a>k := d.k1;
    <a id="L92"></a>if d.p &lt; len(d.digest) {
        <a id="L93"></a>k = d.k2
    <a id="L94"></a>}
    <a id="L95"></a>for i := 0; i &lt; len(d.ci); i++ {
        <a id="L96"></a>d.digest[i] = d.ci[i] ^ k[i]
    <a id="L97"></a>}
    <a id="L98"></a>if d.p &lt; len(d.digest) {
        <a id="L99"></a>d.digest[d.p] ^= 0x80
    <a id="L100"></a>}
    <a id="L101"></a>d.c.Encrypt(d.digest, d.digest);
    <a id="L102"></a>return d.digest;
<a id="L103"></a>}

<a id="L105"></a>func (d *cmac) Size() int { return len(d.digest) }
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
