<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/cbc_aes_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/cbc_aes_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// CBC AES test vectors.</span>

<a id="L7"></a><span class="comment">// See U.S. National Institute of Standards and Technology (NIST)</span>
<a id="L8"></a><span class="comment">// Special Publication 800-38A, ``Recommendation for Block Cipher</span>
<a id="L9"></a><span class="comment">// Modes of Operation,&#39;&#39; 2001 Edition, pp. 24-29.</span>

<a id="L11"></a>package block

<a id="L13"></a>import (
    <a id="L14"></a>&#34;bytes&#34;;
    <a id="L15"></a>&#34;crypto/aes&#34;;
    <a id="L16"></a>&#34;io&#34;;
    <a id="L17"></a>&#34;testing&#34;;
<a id="L18"></a>)

<a id="L20"></a>type cbcTest struct {
    <a id="L21"></a>name string;
    <a id="L22"></a>key  []byte;
    <a id="L23"></a>iv   []byte;
    <a id="L24"></a>in   []byte;
    <a id="L25"></a>out  []byte;
<a id="L26"></a>}

<a id="L28"></a>var cbcAESTests = []cbcTest{
    <a id="L29"></a><span class="comment">// NIST SP 800-38A pp 27-29</span>
    <a id="L30"></a>cbcTest{
        <a id="L31"></a>&#34;CBC-AES128&#34;,
        <a id="L32"></a>commonKey128,
        <a id="L33"></a>commonIV,
        <a id="L34"></a>commonInput,
        <a id="L35"></a>[]byte{
            <a id="L36"></a>0x76, 0x49, 0xab, 0xac, 0x81, 0x19, 0xb2, 0x46, 0xce, 0xe9, 0x8e, 0x9b, 0x12, 0xe9, 0x19, 0x7d,
            <a id="L37"></a>0x50, 0x86, 0xcb, 0x9b, 0x50, 0x72, 0x19, 0xee, 0x95, 0xdb, 0x11, 0x3a, 0x91, 0x76, 0x78, 0xb2,
            <a id="L38"></a>0x73, 0xbe, 0xd6, 0xb8, 0xe3, 0xc1, 0x74, 0x3b, 0x71, 0x16, 0xe6, 0x9e, 0x22, 0x22, 0x95, 0x16,
            <a id="L39"></a>0x3f, 0xf1, 0xca, 0xa1, 0x68, 0x1f, 0xac, 0x09, 0x12, 0x0e, 0xca, 0x30, 0x75, 0x86, 0xe1, 0xa7,
        <a id="L40"></a>},
    <a id="L41"></a>},
    <a id="L42"></a>cbcTest{
        <a id="L43"></a>&#34;CBC-AES192&#34;,
        <a id="L44"></a>commonKey192,
        <a id="L45"></a>commonIV,
        <a id="L46"></a>commonInput,
        <a id="L47"></a>[]byte{
            <a id="L48"></a>0x4f, 0x02, 0x1d, 0xb2, 0x43, 0xbc, 0x63, 0x3d, 0x71, 0x78, 0x18, 0x3a, 0x9f, 0xa0, 0x71, 0xe8,
            <a id="L49"></a>0xb4, 0xd9, 0xad, 0xa9, 0xad, 0x7d, 0xed, 0xf4, 0xe5, 0xe7, 0x38, 0x76, 0x3f, 0x69, 0x14, 0x5a,
            <a id="L50"></a>0x57, 0x1b, 0x24, 0x20, 0x12, 0xfb, 0x7a, 0xe0, 0x7f, 0xa9, 0xba, 0xac, 0x3d, 0xf1, 0x02, 0xe0,
            <a id="L51"></a>0x08, 0xb0, 0xe2, 0x79, 0x88, 0x59, 0x88, 0x81, 0xd9, 0x20, 0xa9, 0xe6, 0x4f, 0x56, 0x15, 0xcd,
        <a id="L52"></a>},
    <a id="L53"></a>},
    <a id="L54"></a>cbcTest{
        <a id="L55"></a>&#34;CBC-AES256&#34;,
        <a id="L56"></a>commonKey256,
        <a id="L57"></a>commonIV,
        <a id="L58"></a>commonInput,
        <a id="L59"></a>[]byte{
            <a id="L60"></a>0xf5, 0x8c, 0x4c, 0x04, 0xd6, 0xe5, 0xf1, 0xba, 0x77, 0x9e, 0xab, 0xfb, 0x5f, 0x7b, 0xfb, 0xd6,
            <a id="L61"></a>0x9c, 0xfc, 0x4e, 0x96, 0x7e, 0xdb, 0x80, 0x8d, 0x67, 0x9f, 0x77, 0x7b, 0xc6, 0x70, 0x2c, 0x7d,
            <a id="L62"></a>0x39, 0xf2, 0x33, 0x69, 0xa9, 0xd9, 0xba, 0xcf, 0xa5, 0x30, 0xe2, 0x63, 0x04, 0x23, 0x14, 0x61,
            <a id="L63"></a>0xb2, 0xeb, 0x05, 0xe2, 0xc3, 0x9b, 0xe9, 0xfc, 0xda, 0x6c, 0x19, 0x07, 0x8c, 0x6a, 0x9d, 0x1b,
        <a id="L64"></a>},
    <a id="L65"></a>},
<a id="L66"></a>}

<a id="L68"></a>func TestCBC_AES(t *testing.T) {
    <a id="L69"></a>for _, tt := range cbcAESTests {
        <a id="L70"></a>test := tt.name;

        <a id="L72"></a>c, err := aes.NewCipher(tt.key);
        <a id="L73"></a>if err != nil {
            <a id="L74"></a>t.Errorf(&#34;%s: NewCipher(%d bytes) = %s&#34;, test, len(tt.key), err);
            <a id="L75"></a>continue;
        <a id="L76"></a>}

        <a id="L78"></a>var crypt bytes.Buffer;
        <a id="L79"></a>w := NewCBCEncrypter(c, tt.iv, &amp;crypt);
        <a id="L80"></a>var r io.Reader = bytes.NewBuffer(tt.in);
        <a id="L81"></a>n, err := io.Copy(w, r);
        <a id="L82"></a>if n != int64(len(tt.in)) || err != nil {
            <a id="L83"></a>t.Errorf(&#34;%s: CBCEncrypter io.Copy = %d, %v want %d, nil&#34;, test, n, err, len(tt.in))
        <a id="L84"></a>} else if d := crypt.Bytes(); !same(tt.out, d) {
            <a id="L85"></a>t.Errorf(&#34;%s: CBCEncrypter\nhave %x\nwant %x&#34;, test, d, tt.out)
        <a id="L86"></a>}

        <a id="L88"></a>var plain bytes.Buffer;
        <a id="L89"></a>r = NewCBCDecrypter(c, tt.iv, bytes.NewBuffer(tt.out));
        <a id="L90"></a>w = &amp;plain;
        <a id="L91"></a>n, err = io.Copy(w, r);
        <a id="L92"></a>if n != int64(len(tt.out)) || err != nil {
            <a id="L93"></a>t.Errorf(&#34;%s: CBCDecrypter io.Copy = %d, %v want %d, nil&#34;, test, n, err, len(tt.out))
        <a id="L94"></a>} else if d := plain.Bytes(); !same(tt.in, d) {
            <a id="L95"></a>t.Errorf(&#34;%s: CBCDecrypter\nhave %x\nwant %x&#34;, test, d, tt.in)
        <a id="L96"></a>}

        <a id="L98"></a>if t.Failed() {
            <a id="L99"></a>break
        <a id="L100"></a>}
    <a id="L101"></a>}
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
