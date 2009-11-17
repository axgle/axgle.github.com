<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/ofb.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/ofb.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Output feedback (OFB) mode.</span>

<a id="L7"></a><span class="comment">// OFB converts a block cipher into a stream cipher by</span>
<a id="L8"></a><span class="comment">// repeatedly encrypting an initialization vector and</span>
<a id="L9"></a><span class="comment">// xoring the resulting stream of data with the input.</span>

<a id="L11"></a><span class="comment">// See NIST SP 800-38A, pp 13-15</span>

<a id="L13"></a>package block

<a id="L15"></a>import (
    <a id="L16"></a>&#34;io&#34;;
<a id="L17"></a>)

<a id="L19"></a>type ofbStream struct {
    <a id="L20"></a>c   Cipher;
    <a id="L21"></a>iv  []byte;
<a id="L22"></a>}

<a id="L24"></a>func newOFBStream(c Cipher, iv []byte) *ofbStream {
    <a id="L25"></a>x := new(ofbStream);
    <a id="L26"></a>x.c = c;
    <a id="L27"></a>n := len(iv);
    <a id="L28"></a>if n != c.BlockSize() {
        <a id="L29"></a>panicln(&#34;crypto/block: newOFBStream: invalid iv size&#34;, n, &#34;!=&#34;, c.BlockSize())
    <a id="L30"></a>}
    <a id="L31"></a>x.iv = copy(iv);
    <a id="L32"></a>return x;
<a id="L33"></a>}

<a id="L35"></a>func (x *ofbStream) Next() []byte {
    <a id="L36"></a>x.c.Encrypt(x.iv, x.iv);
    <a id="L37"></a>return x.iv;
<a id="L38"></a>}

<a id="L40"></a><span class="comment">// NewOFBReader returns a reader that reads data from r, decrypts (or encrypts)</span>
<a id="L41"></a><span class="comment">// it using c in output feedback (OFB) mode with the initialization vector iv.</span>
<a id="L42"></a><span class="comment">// The returned Reader does not buffer and has no block size.</span>
<a id="L43"></a><span class="comment">// In OFB mode, encryption and decryption are the same operation:</span>
<a id="L44"></a><span class="comment">// an OFB reader applied to an encrypted stream produces a decrypted</span>
<a id="L45"></a><span class="comment">// stream and vice versa.</span>
<a id="L46"></a>func NewOFBReader(c Cipher, iv []byte, r io.Reader) io.Reader {
    <a id="L47"></a>return newXorReader(newOFBStream(c, iv), r)
<a id="L48"></a>}

<a id="L50"></a><span class="comment">// NewOFBWriter returns a writer that encrypts (or decrypts) data using c</span>
<a id="L51"></a><span class="comment">// in cipher feedback (OFB) mode with the initialization vector iv</span>
<a id="L52"></a><span class="comment">// and writes the encrypted data to w.</span>
<a id="L53"></a><span class="comment">// The returned Writer does not buffer and has no block size.</span>
<a id="L54"></a><span class="comment">// In OFB mode, encryption and decryption are the same operation:</span>
<a id="L55"></a><span class="comment">// an OFB writer applied to an decrypted stream produces an encrypted</span>
<a id="L56"></a><span class="comment">// stream and vice versa.</span>
<a id="L57"></a>func NewOFBWriter(c Cipher, iv []byte, w io.Writer) io.Writer {
    <a id="L58"></a>return newXorWriter(newOFBStream(c, iv), w)
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
