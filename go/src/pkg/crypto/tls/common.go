<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/tls/common.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/tls/common.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package tls

<a id="L7"></a>import (
    <a id="L8"></a>&#34;crypto/rsa&#34;;
    <a id="L9"></a>&#34;io&#34;;
    <a id="L10"></a>&#34;os&#34;;
<a id="L11"></a>)

<a id="L13"></a>const (
    <a id="L14"></a><span class="comment">// maxTLSCiphertext is the maximum length of a plaintext payload.</span>
    <a id="L15"></a>maxTLSPlaintext = 16384;
    <a id="L16"></a><span class="comment">// maxTLSCiphertext is the maximum length payload after compression and encryption.</span>
    <a id="L17"></a>maxTLSCiphertext = 16384 + 2048;
    <a id="L18"></a><span class="comment">// maxHandshakeMsg is the largest single handshake message that we&#39;ll buffer.</span>
    <a id="L19"></a>maxHandshakeMsg = 65536;
<a id="L20"></a>)


<a id="L23"></a><span class="comment">// TLS record types.</span>
<a id="L24"></a>type recordType uint8

<a id="L26"></a>const (
    <a id="L27"></a>recordTypeChangeCipherSpec recordType = 20;
    <a id="L28"></a>recordTypeAlert            recordType = 21;
    <a id="L29"></a>recordTypeHandshake        recordType = 22;
    <a id="L30"></a>recordTypeApplicationData  recordType = 23;
<a id="L31"></a>)

<a id="L33"></a><span class="comment">// TLS handshake message types.</span>
<a id="L34"></a>const (
    <a id="L35"></a>typeClientHello       uint8 = 1;
    <a id="L36"></a>typeServerHello       uint8 = 2;
    <a id="L37"></a>typeCertificate       uint8 = 11;
    <a id="L38"></a>typeServerHelloDone   uint8 = 14;
    <a id="L39"></a>typeClientKeyExchange uint8 = 16;
    <a id="L40"></a>typeFinished          uint8 = 20;
<a id="L41"></a>)

<a id="L43"></a><span class="comment">// TLS cipher suites.</span>
<a id="L44"></a>var (
    <a id="L45"></a>TLS_RSA_WITH_RC4_128_SHA uint16 = 5;
<a id="L46"></a>)

<a id="L48"></a><span class="comment">// TLS compression types.</span>
<a id="L49"></a>var (
    <a id="L50"></a>compressionNone uint8 = 0;
<a id="L51"></a>)

<a id="L53"></a>type ConnectionState struct {
    <a id="L54"></a>HandshakeComplete bool;
    <a id="L55"></a>CipherSuite       string;
    <a id="L56"></a>Error             alertType;
<a id="L57"></a>}

<a id="L59"></a><span class="comment">// A Config structure is used to configure a TLS client or server. After one</span>
<a id="L60"></a><span class="comment">// has been passed to a TLS function it must not be modified.</span>
<a id="L61"></a>type Config struct {
    <a id="L62"></a><span class="comment">// Rand provides the source of entropy for nonces and RSA blinding.</span>
    <a id="L63"></a>Rand io.Reader;
    <a id="L64"></a><span class="comment">// Time returns the current time as the number of seconds since the epoch.</span>
    <a id="L65"></a>Time         func() int64;
    <a id="L66"></a>Certificates []Certificate;
<a id="L67"></a>}

<a id="L69"></a>type Certificate struct {
    <a id="L70"></a>Certificate [][]byte;
    <a id="L71"></a>PrivateKey  *rsa.PrivateKey;
<a id="L72"></a>}

<a id="L74"></a><span class="comment">// A TLS record.</span>
<a id="L75"></a>type record struct {
    <a id="L76"></a>contentType  recordType;
    <a id="L77"></a>major, minor uint8;
    <a id="L78"></a>payload      []byte;
<a id="L79"></a>}

<a id="L81"></a>type handshakeMessage interface {
    <a id="L82"></a>marshal() []byte;
<a id="L83"></a>}

<a id="L85"></a>type encryptor interface {
    <a id="L86"></a><span class="comment">// XORKeyStream xors the contents of the slice with bytes from the key stream.</span>
    <a id="L87"></a>XORKeyStream(buf []byte);
<a id="L88"></a>}

<a id="L90"></a><span class="comment">// mutualVersion returns the protocol version to use given the advertised</span>
<a id="L91"></a><span class="comment">// version of the peer.</span>
<a id="L92"></a>func mutualVersion(theirMajor, theirMinor uint8) (major, minor uint8, ok bool) {
    <a id="L93"></a><span class="comment">// We don&#39;t deal with peers &lt; TLS 1.0 (aka version 3.1).</span>
    <a id="L94"></a>if theirMajor &lt; 3 || theirMajor == 3 &amp;&amp; theirMinor &lt; 1 {
        <a id="L95"></a>return 0, 0, false
    <a id="L96"></a>}
    <a id="L97"></a>major = 3;
    <a id="L98"></a>minor = 2;
    <a id="L99"></a>if theirMinor &lt; minor {
        <a id="L100"></a>minor = theirMinor
    <a id="L101"></a>}
    <a id="L102"></a>ok = true;
    <a id="L103"></a>return;
<a id="L104"></a>}

<a id="L106"></a><span class="comment">// A nop implements the NULL encryption and MAC algorithms.</span>
<a id="L107"></a>type nop struct{}

<a id="L109"></a>func (nop) XORKeyStream(buf []byte) {}

<a id="L111"></a>func (nop) Write(buf []byte) (int, os.Error) { return len(buf), nil }

<a id="L113"></a>func (nop) Sum() []byte { return nil }

<a id="L115"></a>func (nop) Reset() {}

<a id="L117"></a>func (nop) Size() int { return 0 }
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
