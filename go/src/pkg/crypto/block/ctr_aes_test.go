<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/ctr_aes_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/ctr_aes_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// CTR AES test vectors.</span>

<a id="L7"></a><span class="comment">// See U.S. National Institute of Standards and Technology (NIST)</span>
<a id="L8"></a><span class="comment">// Special Publication 800-38A, ``Recommendation for Block Cipher</span>
<a id="L9"></a><span class="comment">// Modes of Operation,&#39;&#39; 2001 Edition, pp. 55-58.</span>

<a id="L11"></a>package block

<a id="L13"></a>import (
    <a id="L14"></a>&#34;bytes&#34;;
    <a id="L15"></a>&#34;crypto/aes&#34;;
    <a id="L16"></a>&#34;io&#34;;
    <a id="L17"></a>&#34;testing&#34;;
<a id="L18"></a>)

<a id="L20"></a>type ctrTest struct {
    <a id="L21"></a>name string;
    <a id="L22"></a>key  []byte;
    <a id="L23"></a>iv   []byte;
    <a id="L24"></a>in   []byte;
    <a id="L25"></a>out  []byte;
<a id="L26"></a>}

<a id="L28"></a>var commonCounter = []byte{0xf0, 0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7, 0xf8, 0xf9, 0xfa, 0xfb, 0xfc, 0xfd, 0xfe, 0xff}

<a id="L30"></a>var ctrAESTests = []ctrTest{
    <a id="L31"></a><span class="comment">// NIST SP 800-38A pp 55-58</span>
    <a id="L32"></a>ctrTest{
        <a id="L33"></a>&#34;CTR-AES128&#34;,
        <a id="L34"></a>commonKey128,
        <a id="L35"></a>commonCounter,
        <a id="L36"></a>commonInput,
        <a id="L37"></a>[]byte{
            <a id="L38"></a>0x87, 0x4d, 0x61, 0x91, 0xb6, 0x20, 0xe3, 0x26, 0x1b, 0xef, 0x68, 0x64, 0x99, 0x0d, 0xb6, 0xce,
            <a id="L39"></a>0x98, 0x06, 0xf6, 0x6b, 0x79, 0x70, 0xfd, 0xff, 0x86, 0x17, 0x18, 0x7b, 0xb9, 0xff, 0xfd, 0xff,
            <a id="L40"></a>0x5a, 0xe4, 0xdf, 0x3e, 0xdb, 0xd5, 0xd3, 0x5e, 0x5b, 0x4f, 0x09, 0x02, 0x0d, 0xb0, 0x3e, 0xab,
            <a id="L41"></a>0x1e, 0x03, 0x1d, 0xda, 0x2f, 0xbe, 0x03, 0xd1, 0x79, 0x21, 0x70, 0xa0, 0xf3, 0x00, 0x9c, 0xee,
        <a id="L42"></a>},
    <a id="L43"></a>},
    <a id="L44"></a>ctrTest{
        <a id="L45"></a>&#34;CTR-AES192&#34;,
        <a id="L46"></a>commonKey192,
        <a id="L47"></a>commonCounter,
        <a id="L48"></a>commonInput,
        <a id="L49"></a>[]byte{
            <a id="L50"></a>0x1a, 0xbc, 0x93, 0x24, 0x17, 0x52, 0x1c, 0xa2, 0x4f, 0x2b, 0x04, 0x59, 0xfe, 0x7e, 0x6e, 0x0b,
            <a id="L51"></a>0x09, 0x03, 0x39, 0xec, 0x0a, 0xa6, 0xfa, 0xef, 0xd5, 0xcc, 0xc2, 0xc6, 0xf4, 0xce, 0x8e, 0x94,
            <a id="L52"></a>0x1e, 0x36, 0xb2, 0x6b, 0xd1, 0xeb, 0xc6, 0x70, 0xd1, 0xbd, 0x1d, 0x66, 0x56, 0x20, 0xab, 0xf7,
            <a id="L53"></a>0x4f, 0x78, 0xa7, 0xf6, 0xd2, 0x98, 0x09, 0x58, 0x5a, 0x97, 0xda, 0xec, 0x58, 0xc6, 0xb0, 0x50,
        <a id="L54"></a>},
    <a id="L55"></a>},
    <a id="L56"></a>ctrTest{
        <a id="L57"></a>&#34;CTR-AES256&#34;,
        <a id="L58"></a>commonKey256,
        <a id="L59"></a>commonCounter,
        <a id="L60"></a>commonInput,
        <a id="L61"></a>[]byte{
            <a id="L62"></a>0x60, 0x1e, 0xc3, 0x13, 0x77, 0x57, 0x89, 0xa5, 0xb7, 0xa7, 0xf5, 0x04, 0xbb, 0xf3, 0xd2, 0x28,
            <a id="L63"></a>0xf4, 0x43, 0xe3, 0xca, 0x4d, 0x62, 0xb5, 0x9a, 0xca, 0x84, 0xe9, 0x90, 0xca, 0xca, 0xf5, 0xc5,
            <a id="L64"></a>0x2b, 0x09, 0x30, 0xda, 0xa2, 0x3d, 0xe9, 0x4c, 0xe8, 0x70, 0x17, 0xba, 0x2d, 0x84, 0x98, 0x8d,
            <a id="L65"></a>0xdf, 0xc9, 0xc5, 0x8d, 0xb6, 0x7a, 0xad, 0xa6, 0x13, 0xc2, 0xdd, 0x08, 0x45, 0x79, 0x41, 0xa6,
        <a id="L66"></a>},
    <a id="L67"></a>},
<a id="L68"></a>}

<a id="L70"></a>func TestCTR_AES(t *testing.T) {
    <a id="L71"></a>for _, tt := range ctrAESTests {
        <a id="L72"></a>test := tt.name;

        <a id="L74"></a>c, err := aes.NewCipher(tt.key);
        <a id="L75"></a>if err != nil {
            <a id="L76"></a>t.Errorf(&#34;%s: NewCipher(%d bytes) = %s&#34;, test, len(tt.key), err);
            <a id="L77"></a>continue;
        <a id="L78"></a>}

        <a id="L80"></a>for j := 0; j &lt;= 5; j += 5 {
            <a id="L81"></a>var crypt bytes.Buffer;
            <a id="L82"></a>in := tt.in[0 : len(tt.in)-j];
            <a id="L83"></a>w := NewCTRWriter(c, tt.iv, &amp;crypt);
            <a id="L84"></a>var r io.Reader = bytes.NewBuffer(in);
            <a id="L85"></a>n, err := io.Copy(w, r);
            <a id="L86"></a>if n != int64(len(in)) || err != nil {
                <a id="L87"></a>t.Errorf(&#34;%s/%d: CTRWriter io.Copy = %d, %v want %d, nil&#34;, test, len(in), n, err, len(in))
            <a id="L88"></a>} else if d, out := crypt.Bytes(), tt.out[0:len(in)]; !same(out, d) {
                <a id="L89"></a>t.Errorf(&#34;%s/%d: CTRWriter\ninpt %x\nhave %x\nwant %x&#34;, test, len(in), in, d, out)
            <a id="L90"></a>}
        <a id="L91"></a>}

        <a id="L93"></a>for j := 0; j &lt;= 7; j += 7 {
            <a id="L94"></a>var plain bytes.Buffer;
            <a id="L95"></a>out := tt.out[0 : len(tt.out)-j];
            <a id="L96"></a>r := NewCTRReader(c, tt.iv, bytes.NewBuffer(out));
            <a id="L97"></a>w := &amp;plain;
            <a id="L98"></a>n, err := io.Copy(w, r);
            <a id="L99"></a>if n != int64(len(out)) || err != nil {
                <a id="L100"></a>t.Errorf(&#34;%s/%d: CTRReader io.Copy = %d, %v want %d, nil&#34;, test, len(out), n, err, len(out))
            <a id="L101"></a>} else if d, in := plain.Bytes(), tt.in[0:len(out)]; !same(in, d) {
                <a id="L102"></a>t.Errorf(&#34;%s/%d: CTRReader\nhave %x\nwant %x&#34;, test, len(out), d, in)
            <a id="L103"></a>}
        <a id="L104"></a>}

        <a id="L106"></a>if t.Failed() {
            <a id="L107"></a>break
        <a id="L108"></a>}
    <a id="L109"></a>}
<a id="L110"></a>}
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
