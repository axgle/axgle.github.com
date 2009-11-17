<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/eax_aes_test.go</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/eax_aes_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package block

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;crypto/aes&#34;;
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;testing&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">// Test vectors from http://www.cs.ucdavis.edu/~rogaway/papers/eax.pdf</span>

<a id="L17"></a>type eaxAESTest struct {
    <a id="L18"></a>msg    []byte;
    <a id="L19"></a>key    []byte;
    <a id="L20"></a>nonce  []byte;
    <a id="L21"></a>header []byte;
    <a id="L22"></a>cipher []byte;
<a id="L23"></a>}

<a id="L25"></a>var eaxAESTests = []eaxAESTest{
    <a id="L26"></a>eaxAESTest{
        <a id="L27"></a>[]byte{},
        <a id="L28"></a>[]byte{0x23, 0x39, 0x52, 0xDE, 0xE4, 0xD5, 0xED, 0x5F, 0x9B, 0x9C, 0x6D, 0x6F, 0xF8, 0x0F, 0xF4, 0x78},
        <a id="L29"></a>[]byte{0x62, 0xEC, 0x67, 0xF9, 0xC3, 0xA4, 0xA4, 0x07, 0xFC, 0xB2, 0xA8, 0xC4, 0x90, 0x31, 0xA8, 0xB3},
        <a id="L30"></a>[]byte{0x6B, 0xFB, 0x91, 0x4F, 0xD0, 0x7E, 0xAE, 0x6B},
        <a id="L31"></a>[]byte{0xE0, 0x37, 0x83, 0x0E, 0x83, 0x89, 0xF2, 0x7B, 0x02, 0x5A, 0x2D, 0x65, 0x27, 0xE7, 0x9D, 0x01},
    <a id="L32"></a>},
    <a id="L33"></a>eaxAESTest{
        <a id="L34"></a>[]byte{0xF7, 0xFB},
        <a id="L35"></a>[]byte{0x91, 0x94, 0x5D, 0x3F, 0x4D, 0xCB, 0xEE, 0x0B, 0xF4, 0x5E, 0xF5, 0x22, 0x55, 0xF0, 0x95, 0xA4},
        <a id="L36"></a>[]byte{0xBE, 0xCA, 0xF0, 0x43, 0xB0, 0xA2, 0x3D, 0x84, 0x31, 0x94, 0xBA, 0x97, 0x2C, 0x66, 0xDE, 0xBD},
        <a id="L37"></a>[]byte{0xFA, 0x3B, 0xFD, 0x48, 0x06, 0xEB, 0x53, 0xFA},
        <a id="L38"></a>[]byte{0x19, 0xDD, 0x5C, 0x4C, 0x93, 0x31, 0x04, 0x9D, 0x0B, 0xDA, 0xB0, 0x27, 0x74, 0x08, 0xF6, 0x79, 0x67, 0xE5},
    <a id="L39"></a>},
    <a id="L40"></a>eaxAESTest{
        <a id="L41"></a>[]byte{0x1A, 0x47, 0xCB, 0x49, 0x33},
        <a id="L42"></a>[]byte{0x01, 0xF7, 0x4A, 0xD6, 0x40, 0x77, 0xF2, 0xE7, 0x04, 0xC0, 0xF6, 0x0A, 0xDA, 0x3D, 0xD5, 0x23},
        <a id="L43"></a>[]byte{0x70, 0xC3, 0xDB, 0x4F, 0x0D, 0x26, 0x36, 0x84, 0x00, 0xA1, 0x0E, 0xD0, 0x5D, 0x2B, 0xFF, 0x5E},
        <a id="L44"></a>[]byte{0x23, 0x4A, 0x34, 0x63, 0xC1, 0x26, 0x4A, 0xC6},
        <a id="L45"></a>[]byte{0xD8, 0x51, 0xD5, 0xBA, 0xE0, 0x3A, 0x59, 0xF2, 0x38, 0xA2, 0x3E, 0x39, 0x19, 0x9D, 0xC9, 0x26, 0x66, 0x26, 0xC4, 0x0F, 0x80},
    <a id="L46"></a>},
    <a id="L47"></a>eaxAESTest{
        <a id="L48"></a>[]byte{0x48, 0x1C, 0x9E, 0x39, 0xB1},
        <a id="L49"></a>[]byte{0xD0, 0x7C, 0xF6, 0xCB, 0xB7, 0xF3, 0x13, 0xBD, 0xDE, 0x66, 0xB7, 0x27, 0xAF, 0xD3, 0xC5, 0xE8},
        <a id="L50"></a>[]byte{0x84, 0x08, 0xDF, 0xFF, 0x3C, 0x1A, 0x2B, 0x12, 0x92, 0xDC, 0x19, 0x9E, 0x46, 0xB7, 0xD6, 0x17},
        <a id="L51"></a>[]byte{0x33, 0xCC, 0xE2, 0xEA, 0xBF, 0xF5, 0xA7, 0x9D},
        <a id="L52"></a>[]byte{0x63, 0x2A, 0x9D, 0x13, 0x1A, 0xD4, 0xC1, 0x68, 0xA4, 0x22, 0x5D, 0x8E, 0x1F, 0xF7, 0x55, 0x93, 0x99, 0x74, 0xA7, 0xBE, 0xDE},
    <a id="L53"></a>},
    <a id="L54"></a>eaxAESTest{
        <a id="L55"></a>[]byte{0x40, 0xD0, 0xC0, 0x7D, 0xA5, 0xE4},
        <a id="L56"></a>[]byte{0x35, 0xB6, 0xD0, 0x58, 0x00, 0x05, 0xBB, 0xC1, 0x2B, 0x05, 0x87, 0x12, 0x45, 0x57, 0xD2, 0xC2},
        <a id="L57"></a>[]byte{0xFD, 0xB6, 0xB0, 0x66, 0x76, 0xEE, 0xDC, 0x5C, 0x61, 0xD7, 0x42, 0x76, 0xE1, 0xF8, 0xE8, 0x16},
        <a id="L58"></a>[]byte{0xAE, 0xB9, 0x6E, 0xAE, 0xBE, 0x29, 0x70, 0xE9},
        <a id="L59"></a>[]byte{0x07, 0x1D, 0xFE, 0x16, 0xC6, 0x75, 0xCB, 0x06, 0x77, 0xE5, 0x36, 0xF7, 0x3A, 0xFE, 0x6A, 0x14, 0xB7, 0x4E, 0xE4, 0x98, 0x44, 0xDD},
    <a id="L60"></a>},
    <a id="L61"></a>eaxAESTest{
        <a id="L62"></a>[]byte{0x4D, 0xE3, 0xB3, 0x5C, 0x3F, 0xC0, 0x39, 0x24, 0x5B, 0xD1, 0xFB, 0x7D},
        <a id="L63"></a>[]byte{0xBD, 0x8E, 0x6E, 0x11, 0x47, 0x5E, 0x60, 0xB2, 0x68, 0x78, 0x4C, 0x38, 0xC6, 0x2F, 0xEB, 0x22},
        <a id="L64"></a>[]byte{0x6E, 0xAC, 0x5C, 0x93, 0x07, 0x2D, 0x8E, 0x85, 0x13, 0xF7, 0x50, 0x93, 0x5E, 0x46, 0xDA, 0x1B},
        <a id="L65"></a>[]byte{0xD4, 0x48, 0x2D, 0x1C, 0xA7, 0x8D, 0xCE, 0x0F},
        <a id="L66"></a>[]byte{0x83, 0x5B, 0xB4, 0xF1, 0x5D, 0x74, 0x3E, 0x35, 0x0E, 0x72, 0x84, 0x14, 0xAB, 0xB8, 0x64, 0x4F, 0xD6, 0xCC, 0xB8, 0x69, 0x47, 0xC5, 0xE1, 0x05, 0x90, 0x21, 0x0A, 0x4F},
    <a id="L67"></a>},
    <a id="L68"></a>eaxAESTest{
        <a id="L69"></a>[]byte{0x8B, 0x0A, 0x79, 0x30, 0x6C, 0x9C, 0xE7, 0xED, 0x99, 0xDA, 0xE4, 0xF8, 0x7F, 0x8D, 0xD6, 0x16, 0x36},
        <a id="L70"></a>[]byte{0x7C, 0x77, 0xD6, 0xE8, 0x13, 0xBE, 0xD5, 0xAC, 0x98, 0xBA, 0xA4, 0x17, 0x47, 0x7A, 0x2E, 0x7D},
        <a id="L71"></a>[]byte{0x1A, 0x8C, 0x98, 0xDC, 0xD7, 0x3D, 0x38, 0x39, 0x3B, 0x2B, 0xF1, 0x56, 0x9D, 0xEE, 0xFC, 0x19},
        <a id="L72"></a>[]byte{0x65, 0xD2, 0x01, 0x79, 0x90, 0xD6, 0x25, 0x28},
        <a id="L73"></a>[]byte{0x02, 0x08, 0x3E, 0x39, 0x79, 0xDA, 0x01, 0x48, 0x12, 0xF5, 0x9F, 0x11, 0xD5, 0x26, 0x30, 0xDA, 0x30, 0x13, 0x73, 0x27, 0xD1, 0x06, 0x49, 0xB0, 0xAA, 0x6E, 0x1C, 0x18, 0x1D, 0xB6, 0x17, 0xD7, 0xF2},
    <a id="L74"></a>},
    <a id="L75"></a>eaxAESTest{
        <a id="L76"></a>[]byte{0x1B, 0xDA, 0x12, 0x2B, 0xCE, 0x8A, 0x8D, 0xBA, 0xF1, 0x87, 0x7D, 0x96, 0x2B, 0x85, 0x92, 0xDD, 0x2D, 0x56},
        <a id="L77"></a>[]byte{0x5F, 0xFF, 0x20, 0xCA, 0xFA, 0xB1, 0x19, 0xCA, 0x2F, 0xC7, 0x35, 0x49, 0xE2, 0x0F, 0x5B, 0x0D},
        <a id="L78"></a>[]byte{0xDD, 0xE5, 0x9B, 0x97, 0xD7, 0x22, 0x15, 0x6D, 0x4D, 0x9A, 0xFF, 0x2B, 0xC7, 0x55, 0x98, 0x26},
        <a id="L79"></a>[]byte{0x54, 0xB9, 0xF0, 0x4E, 0x6A, 0x09, 0x18, 0x9A},
        <a id="L80"></a>[]byte{0x2E, 0xC4, 0x7B, 0x2C, 0x49, 0x54, 0xA4, 0x89, 0xAF, 0xC7, 0xBA, 0x48, 0x97, 0xED, 0xCD, 0xAE, 0x8C, 0xC3, 0x3B, 0x60, 0x45, 0x05, 0x99, 0xBD, 0x02, 0xC9, 0x63, 0x82, 0x90, 0x2A, 0xEF, 0x7F, 0x83, 0x2A},
    <a id="L81"></a>},
    <a id="L82"></a>eaxAESTest{
        <a id="L83"></a>[]byte{0x6C, 0xF3, 0x67, 0x20, 0x87, 0x2B, 0x85, 0x13, 0xF6, 0xEA, 0xB1, 0xA8, 0xA4, 0x44, 0x38, 0xD5, 0xEF, 0x11},
        <a id="L84"></a>[]byte{0xA4, 0xA4, 0x78, 0x2B, 0xCF, 0xFD, 0x3E, 0xC5, 0xE7, 0xEF, 0x6D, 0x8C, 0x34, 0xA5, 0x61, 0x23},
        <a id="L85"></a>[]byte{0xB7, 0x81, 0xFC, 0xF2, 0xF7, 0x5F, 0xA5, 0xA8, 0xDE, 0x97, 0xA9, 0xCA, 0x48, 0xE5, 0x22, 0xEC},
        <a id="L86"></a>[]byte{0x89, 0x9A, 0x17, 0x58, 0x97, 0x56, 0x1D, 0x7E},
        <a id="L87"></a>[]byte{0x0D, 0xE1, 0x8F, 0xD0, 0xFD, 0xD9, 0x1E, 0x7A, 0xF1, 0x9F, 0x1D, 0x8E, 0xE8, 0x73, 0x39, 0x38, 0xB1, 0xE8, 0xE7, 0xF6, 0xD2, 0x23, 0x16, 0x18, 0x10, 0x2F, 0xDB, 0x7F, 0xE5, 0x5F, 0xF1, 0x99, 0x17, 0x00},
    <a id="L88"></a>},
    <a id="L89"></a>eaxAESTest{
        <a id="L90"></a>[]byte{0xCA, 0x40, 0xD7, 0x44, 0x6E, 0x54, 0x5F, 0xFA, 0xED, 0x3B, 0xD1, 0x2A, 0x74, 0x0A, 0x65, 0x9F, 0xFB, 0xBB, 0x3C, 0xEA, 0xB7},
        <a id="L91"></a>[]byte{0x83, 0x95, 0xFC, 0xF1, 0xE9, 0x5B, 0xEB, 0xD6, 0x97, 0xBD, 0x01, 0x0B, 0xC7, 0x66, 0xAA, 0xC3},
        <a id="L92"></a>[]byte{0x22, 0xE7, 0xAD, 0xD9, 0x3C, 0xFC, 0x63, 0x93, 0xC5, 0x7E, 0xC0, 0xB3, 0xC1, 0x7D, 0x6B, 0x44},
        <a id="L93"></a>[]byte{0x12, 0x67, 0x35, 0xFC, 0xC3, 0x20, 0xD2, 0x5A},
        <a id="L94"></a>[]byte{0xCB, 0x89, 0x20, 0xF8, 0x7A, 0x6C, 0x75, 0xCF, 0xF3, 0x96, 0x27, 0xB5, 0x6E, 0x3E, 0xD1, 0x97, 0xC5, 0x52, 0xD2, 0x95, 0xA7, 0xCF, 0xC4, 0x6A, 0xFC, 0x25, 0x3B, 0x46, 0x52, 0xB1, 0xAF, 0x37, 0x95, 0xB1, 0x24, 0xAB, 0x6E},
    <a id="L95"></a>},
<a id="L96"></a>}

<a id="L98"></a>func TestEAXEncrypt_AES(t *testing.T) {
    <a id="L99"></a>b := new(bytes.Buffer);
    <a id="L100"></a>for i, tt := range eaxAESTests {
        <a id="L101"></a>test := fmt.Sprintf(&#34;test %d&#34;, i);
        <a id="L102"></a>c, err := aes.NewCipher(tt.key);
        <a id="L103"></a>if err != nil {
            <a id="L104"></a>t.Fatalf(&#34;%s: NewCipher(%d bytes) = %s&#34;, test, len(tt.key), err)
        <a id="L105"></a>}
        <a id="L106"></a>b.Reset();
        <a id="L107"></a>enc := NewEAXEncrypter(c, tt.nonce, tt.header, 16, b);
        <a id="L108"></a>n, err := io.Copy(enc, bytes.NewBuffer(tt.msg));
        <a id="L109"></a>if n != int64(len(tt.msg)) || err != nil {
            <a id="L110"></a>t.Fatalf(&#34;%s: io.Copy into encrypter: %d, %s&#34;, test, n, err)
        <a id="L111"></a>}
        <a id="L112"></a>err = enc.Close();
        <a id="L113"></a>if err != nil {
            <a id="L114"></a>t.Fatalf(&#34;%s: enc.Close: %s&#34;, test, err)
        <a id="L115"></a>}
        <a id="L116"></a>if d := b.Bytes(); !same(d, tt.cipher) {
            <a id="L117"></a>t.Fatalf(&#34;%s: got %x want %x&#34;, test, d, tt.cipher)
        <a id="L118"></a>}
    <a id="L119"></a>}
<a id="L120"></a>}

<a id="L122"></a>func TestEAXDecrypt_AES(t *testing.T) {
    <a id="L123"></a>b := new(bytes.Buffer);
    <a id="L124"></a>for i, tt := range eaxAESTests {
        <a id="L125"></a>test := fmt.Sprintf(&#34;test %d&#34;, i);
        <a id="L126"></a>c, err := aes.NewCipher(tt.key);
        <a id="L127"></a>if err != nil {
            <a id="L128"></a>t.Fatalf(&#34;%s: NewCipher(%d bytes) = %s&#34;, test, len(tt.key), err)
        <a id="L129"></a>}
        <a id="L130"></a>b.Reset();
        <a id="L131"></a>dec := NewEAXDecrypter(c, tt.nonce, tt.header, 16, bytes.NewBuffer(tt.cipher));
        <a id="L132"></a>n, err := io.Copy(b, dec);
        <a id="L133"></a>if n != int64(len(tt.msg)) || err != nil {
            <a id="L134"></a>t.Fatalf(&#34;%s: io.Copy into decrypter: %d, %s&#34;, test, n, err)
        <a id="L135"></a>}
        <a id="L136"></a>if d := b.Bytes(); !same(d, tt.msg) {
            <a id="L137"></a>t.Fatalf(&#34;%s: got %x want %x&#34;, test, d, tt.msg)
        <a id="L138"></a>}
    <a id="L139"></a>}
<a id="L140"></a>}
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
