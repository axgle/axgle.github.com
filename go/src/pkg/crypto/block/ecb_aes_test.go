<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/ecb_aes_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/ecb_aes_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// ECB AES test vectors.</span>

<a id="L7"></a><span class="comment">// See U.S. National Institute of Standards and Technology (NIST)</span>
<a id="L8"></a><span class="comment">// Special Publication 800-38A, ``Recommendation for Block Cipher</span>
<a id="L9"></a><span class="comment">// Modes of Operation,&#39;&#39; 2001 Edition, pp. 24-27.</span>

<a id="L11"></a>package block

<a id="L13"></a>import (
    <a id="L14"></a>&#34;bytes&#34;;
    <a id="L15"></a>&#34;crypto/aes&#34;;
    <a id="L16"></a>&#34;io&#34;;
    <a id="L17"></a>&#34;testing&#34;;
<a id="L18"></a>)

<a id="L20"></a>type ecbTest struct {
    <a id="L21"></a>name string;
    <a id="L22"></a>key  []byte;
    <a id="L23"></a>in   []byte;
    <a id="L24"></a>out  []byte;
<a id="L25"></a>}

<a id="L27"></a>var commonInput = []byte{
    <a id="L28"></a>0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
    <a id="L29"></a>0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
    <a id="L30"></a>0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11, 0xe5, 0xfb, 0xc1, 0x19, 0x1a, 0x0a, 0x52, 0xef,
    <a id="L31"></a>0xf6, 0x9f, 0x24, 0x45, 0xdf, 0x4f, 0x9b, 0x17, 0xad, 0x2b, 0x41, 0x7b, 0xe6, 0x6c, 0x37, 0x10,
<a id="L32"></a>}

<a id="L34"></a>var commonKey128 = []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}

<a id="L36"></a>var commonKey192 = []byte{
    <a id="L37"></a>0x8e, 0x73, 0xb0, 0xf7, 0xda, 0x0e, 0x64, 0x52, 0xc8, 0x10, 0xf3, 0x2b, 0x80, 0x90, 0x79, 0xe5,
    <a id="L38"></a>0x62, 0xf8, 0xea, 0xd2, 0x52, 0x2c, 0x6b, 0x7b,
<a id="L39"></a>}

<a id="L41"></a>var commonKey256 = []byte{
    <a id="L42"></a>0x60, 0x3d, 0xeb, 0x10, 0x15, 0xca, 0x71, 0xbe, 0x2b, 0x73, 0xae, 0xf0, 0x85, 0x7d, 0x77, 0x81,
    <a id="L43"></a>0x1f, 0x35, 0x2c, 0x07, 0x3b, 0x61, 0x08, 0xd7, 0x2d, 0x98, 0x10, 0xa3, 0x09, 0x14, 0xdf, 0xf4,
<a id="L44"></a>}

<a id="L46"></a>var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

<a id="L48"></a>var ecbAESTests = []ecbTest{
    <a id="L49"></a><span class="comment">// FIPS 197, Appendix B, C</span>
    <a id="L50"></a>ecbTest{
        <a id="L51"></a>&#34;FIPS-197 Appendix B&#34;,
        <a id="L52"></a>commonKey128,
        <a id="L53"></a>[]byte{0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34},
        <a id="L54"></a>[]byte{0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b, 0x32},
    <a id="L55"></a>},

    <a id="L57"></a><span class="comment">// NIST SP 800-38A pp 24-27</span>
    <a id="L58"></a>ecbTest{
        <a id="L59"></a>&#34;ECB-AES128&#34;,
        <a id="L60"></a>commonKey128,
        <a id="L61"></a>commonInput,
        <a id="L62"></a>[]byte{
            <a id="L63"></a>0x3a, 0xd7, 0x7b, 0xb4, 0x0d, 0x7a, 0x36, 0x60, 0xa8, 0x9e, 0xca, 0xf3, 0x24, 0x66, 0xef, 0x97,
            <a id="L64"></a>0xf5, 0xd3, 0xd5, 0x85, 0x03, 0xb9, 0x69, 0x9d, 0xe7, 0x85, 0x89, 0x5a, 0x96, 0xfd, 0xba, 0xaf,
            <a id="L65"></a>0x43, 0xb1, 0xcd, 0x7f, 0x59, 0x8e, 0xce, 0x23, 0x88, 0x1b, 0x00, 0xe3, 0xed, 0x03, 0x06, 0x88,
            <a id="L66"></a>0x7b, 0x0c, 0x78, 0x5e, 0x27, 0xe8, 0xad, 0x3f, 0x82, 0x23, 0x20, 0x71, 0x04, 0x72, 0x5d, 0xd4,
        <a id="L67"></a>},
    <a id="L68"></a>},
    <a id="L69"></a>ecbTest{
        <a id="L70"></a>&#34;ECB-AES192&#34;,
        <a id="L71"></a>commonKey192,
        <a id="L72"></a>commonInput,
        <a id="L73"></a>[]byte{
            <a id="L74"></a>0xbd, 0x33, 0x4f, 0x1d, 0x6e, 0x45, 0xf2, 0x5f, 0xf7, 0x12, 0xa2, 0x14, 0x57, 0x1f, 0xa5, 0xcc,
            <a id="L75"></a>0x97, 0x41, 0x04, 0x84, 0x6d, 0x0a, 0xd3, 0xad, 0x77, 0x34, 0xec, 0xb3, 0xec, 0xee, 0x4e, 0xef,
            <a id="L76"></a>0xef, 0x7a, 0xfd, 0x22, 0x70, 0xe2, 0xe6, 0x0a, 0xdc, 0xe0, 0xba, 0x2f, 0xac, 0xe6, 0x44, 0x4e,
            <a id="L77"></a>0x9a, 0x4b, 0x41, 0xba, 0x73, 0x8d, 0x6c, 0x72, 0xfb, 0x16, 0x69, 0x16, 0x03, 0xc1, 0x8e, 0x0e,
        <a id="L78"></a>},
    <a id="L79"></a>},
    <a id="L80"></a>ecbTest{
        <a id="L81"></a>&#34;ECB-AES256&#34;,
        <a id="L82"></a>commonKey256,
        <a id="L83"></a>commonInput,
        <a id="L84"></a>[]byte{
            <a id="L85"></a>0xf3, 0xee, 0xd1, 0xbd, 0xb5, 0xd2, 0xa0, 0x3c, 0x06, 0x4b, 0x5a, 0x7e, 0x3d, 0xb1, 0x81, 0xf8,
            <a id="L86"></a>0x59, 0x1c, 0xcb, 0x10, 0xd4, 0x10, 0xed, 0x26, 0xdc, 0x5b, 0xa7, 0x4a, 0x31, 0x36, 0x28, 0x70,
            <a id="L87"></a>0xb6, 0xed, 0x21, 0xb9, 0x9c, 0xa6, 0xf4, 0xf9, 0xf1, 0x53, 0xe7, 0xb1, 0xbe, 0xaf, 0xed, 0x1d,
            <a id="L88"></a>0x23, 0x30, 0x4b, 0x7a, 0x39, 0xf9, 0xf3, 0xff, 0x06, 0x7d, 0x8d, 0x8f, 0x9e, 0x24, 0xec, 0xc7,
        <a id="L89"></a>},
    <a id="L90"></a>},
<a id="L91"></a>}

<a id="L93"></a>func TestECB_AES(t *testing.T) {
    <a id="L94"></a>for _, tt := range ecbAESTests {
        <a id="L95"></a>test := tt.name;

        <a id="L97"></a>c, err := aes.NewCipher(tt.key);
        <a id="L98"></a>if err != nil {
            <a id="L99"></a>t.Errorf(&#34;%s: NewCipher(%d bytes) = %s&#34;, test, len(tt.key), err);
            <a id="L100"></a>continue;
        <a id="L101"></a>}

        <a id="L103"></a>var crypt bytes.Buffer;
        <a id="L104"></a>w := NewECBEncrypter(c, &amp;crypt);
        <a id="L105"></a>var r io.Reader = bytes.NewBuffer(tt.in);
        <a id="L106"></a>n, err := io.Copy(w, r);
        <a id="L107"></a>if n != int64(len(tt.in)) || err != nil {
            <a id="L108"></a>t.Errorf(&#34;%s: ECBReader io.Copy = %d, %v want %d, nil&#34;, test, n, err, len(tt.in))
        <a id="L109"></a>} else if d := crypt.Bytes(); !same(tt.out, d) {
            <a id="L110"></a>t.Errorf(&#34;%s: ECBReader\nhave %x\nwant %x&#34;, test, d, tt.out)
        <a id="L111"></a>}

        <a id="L113"></a>var plain bytes.Buffer;
        <a id="L114"></a>r = NewECBDecrypter(c, bytes.NewBuffer(tt.out));
        <a id="L115"></a>w = &amp;plain;
        <a id="L116"></a>n, err = io.Copy(w, r);
        <a id="L117"></a>if n != int64(len(tt.out)) || err != nil {
            <a id="L118"></a>t.Errorf(&#34;%s: ECBWriter io.Copy = %d, %v want %d, nil&#34;, test, n, err, len(tt.out))
        <a id="L119"></a>} else if d := plain.Bytes(); !same(tt.in, d) {
            <a id="L120"></a>t.Errorf(&#34;%s: ECBWriter\nhave %x\nwant %x&#34;, test, d, tt.in)
        <a id="L121"></a>}

        <a id="L123"></a>if t.Failed() {
            <a id="L124"></a>break
        <a id="L125"></a>}
    <a id="L126"></a>}
<a id="L127"></a>}
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
