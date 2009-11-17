<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/cmac_aes_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/cmac_aes_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// CMAC test vectors.  See NIST SP 800-38B, Appendix D.</span>

<a id="L7"></a>package block

<a id="L9"></a>import (
    <a id="L10"></a>&#34;crypto/aes&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)

<a id="L14"></a>type cmacAESTest struct {
    <a id="L15"></a>key    []byte;
    <a id="L16"></a>in     []byte;
    <a id="L17"></a>digest []byte;
<a id="L18"></a>}

<a id="L20"></a>var cmacAESTests = []cmacAESTest{
    <a id="L21"></a>cmacAESTest{
        <a id="L22"></a>commonKey128,
        <a id="L23"></a>nil,
        <a id="L24"></a>[]byte{0xbb, 0x1d, 0x69, 0x29, 0xe9, 0x59, 0x37, 0x28, 0x7f, 0xa3, 0x7d, 0x12, 0x9b, 0x75, 0x67, 0x46},
    <a id="L25"></a>},
    <a id="L26"></a>cmacAESTest{
        <a id="L27"></a>commonKey128,
        <a id="L28"></a>[]byte{0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a},
        <a id="L29"></a>[]byte{0x07, 0x0a, 0x16, 0xb4, 0x6b, 0x4d, 0x41, 0x44, 0xf7, 0x9b, 0xdd, 0x9d, 0xd0, 0x4a, 0x28, 0x7c},
    <a id="L30"></a>},
    <a id="L31"></a>cmacAESTest{
        <a id="L32"></a>commonKey128,
        <a id="L33"></a>[]byte{
            <a id="L34"></a>0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
            <a id="L35"></a>0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
            <a id="L36"></a>0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11,
        <a id="L37"></a>},
        <a id="L38"></a>[]byte{0xdf, 0xa6, 0x67, 0x47, 0xde, 0x9a, 0xe6, 0x30, 0x30, 0xca, 0x32, 0x61, 0x14, 0x97, 0xc8, 0x27},
    <a id="L39"></a>},
    <a id="L40"></a>cmacAESTest{
        <a id="L41"></a>commonKey128,
        <a id="L42"></a>[]byte{
            <a id="L43"></a>0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
            <a id="L44"></a>0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
            <a id="L45"></a>0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11, 0xe5, 0xfb, 0xc1, 0x19, 0x1a, 0x0a, 0x52, 0xef,
            <a id="L46"></a>0xf6, 0x9f, 0x24, 0x45, 0xdf, 0x4f, 0x9b, 0x17, 0xad, 0x2b, 0x41, 0x7b, 0xe6, 0x6c, 0x37, 0x10,
        <a id="L47"></a>},
        <a id="L48"></a>[]byte{0x51, 0xf0, 0xbe, 0xbf, 0x7e, 0x3b, 0x9d, 0x92, 0xfc, 0x49, 0x74, 0x17, 0x79, 0x36, 0x3c, 0xfe},
    <a id="L49"></a>},
    <a id="L50"></a>cmacAESTest{
        <a id="L51"></a>commonKey192,
        <a id="L52"></a>nil,
        <a id="L53"></a>[]byte{0xd1, 0x7d, 0xdf, 0x46, 0xad, 0xaa, 0xcd, 0xe5, 0x31, 0xca, 0xc4, 0x83, 0xde, 0x7a, 0x93, 0x67},
    <a id="L54"></a>},
    <a id="L55"></a>cmacAESTest{
        <a id="L56"></a>commonKey192,
        <a id="L57"></a>[]byte{0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a},
        <a id="L58"></a>[]byte{0x9e, 0x99, 0xa7, 0xbf, 0x31, 0xe7, 0x10, 0x90, 0x06, 0x62, 0xf6, 0x5e, 0x61, 0x7c, 0x51, 0x84},
    <a id="L59"></a>},
    <a id="L60"></a>cmacAESTest{
        <a id="L61"></a>commonKey192,
        <a id="L62"></a>[]byte{
            <a id="L63"></a>0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
            <a id="L64"></a>0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
            <a id="L65"></a>0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11,
        <a id="L66"></a>},
        <a id="L67"></a>[]byte{0x8a, 0x1d, 0xe5, 0xbe, 0x2e, 0xb3, 0x1a, 0xad, 0x08, 0x9a, 0x82, 0xe6, 0xee, 0x90, 0x8b, 0x0e},
    <a id="L68"></a>},
    <a id="L69"></a>cmacAESTest{
        <a id="L70"></a>commonKey192,
        <a id="L71"></a>[]byte{
            <a id="L72"></a>0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
            <a id="L73"></a>0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
            <a id="L74"></a>0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11, 0xe5, 0xfb, 0xc1, 0x19, 0x1a, 0x0a, 0x52, 0xef,
            <a id="L75"></a>0xf6, 0x9f, 0x24, 0x45, 0xdf, 0x4f, 0x9b, 0x17, 0xad, 0x2b, 0x41, 0x7b, 0xe6, 0x6c, 0x37, 0x10,
        <a id="L76"></a>},
        <a id="L77"></a>[]byte{0xa1, 0xd5, 0xdf, 0x0e, 0xed, 0x79, 0x0f, 0x79, 0x4d, 0x77, 0x58, 0x96, 0x59, 0xf3, 0x9a, 0x11},
    <a id="L78"></a>},
    <a id="L79"></a>cmacAESTest{
        <a id="L80"></a>commonKey256,
        <a id="L81"></a>nil,
        <a id="L82"></a>[]byte{0x02, 0x89, 0x62, 0xf6, 0x1b, 0x7b, 0xf8, 0x9e, 0xfc, 0x6b, 0x55, 0x1f, 0x46, 0x67, 0xd9, 0x83},
    <a id="L83"></a>},
    <a id="L84"></a>cmacAESTest{
        <a id="L85"></a>commonKey256,
        <a id="L86"></a>[]byte{0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a},
        <a id="L87"></a>[]byte{0x28, 0xa7, 0x02, 0x3f, 0x45, 0x2e, 0x8f, 0x82, 0xbd, 0x4b, 0xf2, 0x8d, 0x8c, 0x37, 0xc3, 0x5c},
    <a id="L88"></a>},
    <a id="L89"></a>cmacAESTest{
        <a id="L90"></a>commonKey256,
        <a id="L91"></a>[]byte{
            <a id="L92"></a>0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
            <a id="L93"></a>0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
            <a id="L94"></a>0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11,
        <a id="L95"></a>},
        <a id="L96"></a>[]byte{0xaa, 0xf3, 0xd8, 0xf1, 0xde, 0x56, 0x40, 0xc2, 0x32, 0xf5, 0xb1, 0x69, 0xb9, 0xc9, 0x11, 0xe6},
    <a id="L97"></a>},
    <a id="L98"></a>cmacAESTest{
        <a id="L99"></a>commonKey256,
        <a id="L100"></a>[]byte{
            <a id="L101"></a>0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
            <a id="L102"></a>0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
            <a id="L103"></a>0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11, 0xe5, 0xfb, 0xc1, 0x19, 0x1a, 0x0a, 0x52, 0xef,
            <a id="L104"></a>0xf6, 0x9f, 0x24, 0x45, 0xdf, 0x4f, 0x9b, 0x17, 0xad, 0x2b, 0x41, 0x7b, 0xe6, 0x6c, 0x37, 0x10,
        <a id="L105"></a>},
        <a id="L106"></a>[]byte{0xe1, 0x99, 0x21, 0x90, 0x54, 0x9f, 0x6e, 0xd5, 0x69, 0x6a, 0x2c, 0x05, 0x6c, 0x31, 0x54, 0x10},
    <a id="L107"></a>},
<a id="L108"></a>}

<a id="L110"></a>func TestCMAC_AES(t *testing.T) {
    <a id="L111"></a>for i, tt := range cmacAESTests {
        <a id="L112"></a>c, err := aes.NewCipher(tt.key);
        <a id="L113"></a>if err != nil {
            <a id="L114"></a>t.Errorf(&#34;test %d: NewCipher: %s&#34;, i, err);
            <a id="L115"></a>continue;
        <a id="L116"></a>}
        <a id="L117"></a>d := NewCMAC(c);
        <a id="L118"></a>n, err := d.Write(tt.in);
        <a id="L119"></a>if err != nil || n != len(tt.in) {
            <a id="L120"></a>t.Errorf(&#34;test %d: Write %d: %d, %s&#34;, i, len(tt.in), n, err);
            <a id="L121"></a>continue;
        <a id="L122"></a>}
        <a id="L123"></a>sum := d.Sum();
        <a id="L124"></a>if !same(sum, tt.digest) {
            <a id="L125"></a>x := d.(*cmac);
            <a id="L126"></a>t.Errorf(&#34;test %d: digest mismatch\n\twant %x\n\thave %x\n\tk1 %x\n\tk2 %x&#34;, i, tt.digest, sum, x.k1, x.k2);
            <a id="L127"></a>continue;
        <a id="L128"></a>}
    <a id="L129"></a>}
<a id="L130"></a>}
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
