<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/aes/aes_test.go</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/aes/aes_test.go</h1>

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
    <a id="L8"></a>&#34;testing&#34;;
<a id="L9"></a>)

<a id="L11"></a><span class="comment">// See const.go for overview of math here.</span>

<a id="L13"></a><span class="comment">// Test that powx is initialized correctly.</span>
<a id="L14"></a><span class="comment">// (Can adapt this code to generate it too.)</span>
<a id="L15"></a>func TestPowx(t *testing.T) {
    <a id="L16"></a>p := 1;
    <a id="L17"></a>for i := 0; i &lt; len(powx); i++ {
        <a id="L18"></a>if powx[i] != byte(p) {
            <a id="L19"></a>t.Errorf(&#34;powx[%d] = %#x, want %#x&#34;, i, powx[i], p)
        <a id="L20"></a>}
        <a id="L21"></a>p &lt;&lt;= 1;
        <a id="L22"></a>if p&amp;0x100 != 0 {
            <a id="L23"></a>p ^= poly
        <a id="L24"></a>}
    <a id="L25"></a>}
<a id="L26"></a>}

<a id="L28"></a><span class="comment">// Multiply b and c as GF(2) polynomials modulo poly</span>
<a id="L29"></a>func mul(b, c uint32) uint32 {
    <a id="L30"></a>i := b;
    <a id="L31"></a>j := c;
    <a id="L32"></a>s := uint32(0);
    <a id="L33"></a>for k := uint32(1); k &lt; 0x100 &amp;&amp; j != 0; k &lt;&lt;= 1 {
        <a id="L34"></a><span class="comment">// Invariant: k == 1&lt;&lt;n, i == b * xⁿ</span>

        <a id="L36"></a>if j&amp;k != 0 {
            <a id="L37"></a><span class="comment">// s += i in GF(2); xor in binary</span>
            <a id="L38"></a>s ^= i;
            <a id="L39"></a>j ^= k; <span class="comment">// turn off bit to end loop early</span>
        <a id="L40"></a>}

        <a id="L42"></a><span class="comment">// i *= x in GF(2) modulo the polynomial</span>
        <a id="L43"></a>i &lt;&lt;= 1;
        <a id="L44"></a>if i&amp;0x100 != 0 {
            <a id="L45"></a>i ^= poly
        <a id="L46"></a>}
    <a id="L47"></a>}
    <a id="L48"></a>return s;
<a id="L49"></a>}

<a id="L51"></a><span class="comment">// Test all mul inputs against bit-by-bit n² algorithm.</span>
<a id="L52"></a>func TestMul(t *testing.T) {
    <a id="L53"></a>for i := uint32(0); i &lt; 256; i++ {
        <a id="L54"></a>for j := uint32(0); j &lt; 256; j++ {
            <a id="L55"></a><span class="comment">// Multiply i, j bit by bit.</span>
            <a id="L56"></a>s := uint8(0);
            <a id="L57"></a>for k := uint(0); k &lt; 8; k++ {
                <a id="L58"></a>for l := uint(0); l &lt; 8; l++ {
                    <a id="L59"></a>if i&amp;(1&lt;&lt;k) != 0 &amp;&amp; j&amp;(1&lt;&lt;l) != 0 {
                        <a id="L60"></a>s ^= powx[k+l]
                    <a id="L61"></a>}
                <a id="L62"></a>}
            <a id="L63"></a>}
            <a id="L64"></a>if x := mul(i, j); x != uint32(s) {
                <a id="L65"></a>t.Fatalf(&#34;mul(%#x, %#x) = %#x, want %#x&#34;, i, j, x, s)
            <a id="L66"></a>}
        <a id="L67"></a>}
    <a id="L68"></a>}
<a id="L69"></a>}

<a id="L71"></a><span class="comment">// Check that S-boxes are inverses of each other.</span>
<a id="L72"></a><span class="comment">// They have more structure that we could test,</span>
<a id="L73"></a><span class="comment">// but if this sanity check passes, we&#39;ll assume</span>
<a id="L74"></a><span class="comment">// the cut and paste from the FIPS PDF worked.</span>
<a id="L75"></a>func TestSboxes(t *testing.T) {
    <a id="L76"></a>for i := 0; i &lt; 256; i++ {
        <a id="L77"></a>if j := sbox0[sbox1[i]]; j != byte(i) {
            <a id="L78"></a>t.Errorf(&#34;sbox0[sbox1[%#x]] = %#x&#34;, i, j)
        <a id="L79"></a>}
        <a id="L80"></a>if j := sbox1[sbox0[i]]; j != byte(i) {
            <a id="L81"></a>t.Errorf(&#34;sbox1[sbox0[%#x]] = %#x&#34;, i, j)
        <a id="L82"></a>}
    <a id="L83"></a>}
<a id="L84"></a>}

<a id="L86"></a><span class="comment">// Test that encryption tables are correct.</span>
<a id="L87"></a><span class="comment">// (Can adapt this code to generate them too.)</span>
<a id="L88"></a>func TestTe(t *testing.T) {
    <a id="L89"></a>for i := 0; i &lt; 256; i++ {
        <a id="L90"></a>s := uint32(sbox0[i]);
        <a id="L91"></a>s2 := mul(s, 2);
        <a id="L92"></a>s3 := mul(s, 3);
        <a id="L93"></a>w := s2&lt;&lt;24 | s&lt;&lt;16 | s&lt;&lt;8 | s3;
        <a id="L94"></a>for j := 0; j &lt; 4; j++ {
            <a id="L95"></a>if x := te[j][i]; x != w {
                <a id="L96"></a>t.Fatalf(&#34;te[%d][%d] = %#x, want %#x&#34;, j, i, x, w)
            <a id="L97"></a>}
            <a id="L98"></a>w = w&lt;&lt;24 | w&gt;&gt;8;
        <a id="L99"></a>}
    <a id="L100"></a>}
<a id="L101"></a>}

<a id="L103"></a><span class="comment">// Test that decryption tables are correct.</span>
<a id="L104"></a><span class="comment">// (Can adapt this code to generate them too.)</span>
<a id="L105"></a>func TestTd(t *testing.T) {
    <a id="L106"></a>for i := 0; i &lt; 256; i++ {
        <a id="L107"></a>s := uint32(sbox1[i]);
        <a id="L108"></a>s9 := mul(s, 0x9);
        <a id="L109"></a>sb := mul(s, 0xb);
        <a id="L110"></a>sd := mul(s, 0xd);
        <a id="L111"></a>se := mul(s, 0xe);
        <a id="L112"></a>w := se&lt;&lt;24 | s9&lt;&lt;16 | sd&lt;&lt;8 | sb;
        <a id="L113"></a>for j := 0; j &lt; 4; j++ {
            <a id="L114"></a>if x := td[j][i]; x != w {
                <a id="L115"></a>t.Fatalf(&#34;td[%d][%d] = %#x, want %#x&#34;, j, i, x, w)
            <a id="L116"></a>}
            <a id="L117"></a>w = w&lt;&lt;24 | w&gt;&gt;8;
        <a id="L118"></a>}
    <a id="L119"></a>}
<a id="L120"></a>}

<a id="L122"></a><span class="comment">// Test vectors are from FIPS 197:</span>
<a id="L123"></a><span class="comment">//	http://www.csrc.nist.gov/publications/fips/fips197/fips-197.pdf</span>

<a id="L125"></a><span class="comment">// Appendix A of FIPS 197: Key expansion examples</span>
<a id="L126"></a>type KeyTest struct {
    <a id="L127"></a>key []byte;
    <a id="L128"></a>enc []uint32;
    <a id="L129"></a>dec []uint32; <span class="comment">// decryption expansion; not in FIPS 197, computed from C implementation.</span>
<a id="L130"></a>}

<a id="L132"></a>var keyTests = []KeyTest{
    <a id="L133"></a>KeyTest{
        <a id="L134"></a><span class="comment">// A.1.  Expansion of a 128-bit Cipher Key</span>
        <a id="L135"></a>[]byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c},
        <a id="L136"></a>[]uint32{
            <a id="L137"></a>0x2b7e1516, 0x28aed2a6, 0xabf71588, 0x09cf4f3c,
            <a id="L138"></a>0xa0fafe17, 0x88542cb1, 0x23a33939, 0x2a6c7605,
            <a id="L139"></a>0xf2c295f2, 0x7a96b943, 0x5935807a, 0x7359f67f,
            <a id="L140"></a>0x3d80477d, 0x4716fe3e, 0x1e237e44, 0x6d7a883b,
            <a id="L141"></a>0xef44a541, 0xa8525b7f, 0xb671253b, 0xdb0bad00,
            <a id="L142"></a>0xd4d1c6f8, 0x7c839d87, 0xcaf2b8bc, 0x11f915bc,
            <a id="L143"></a>0x6d88a37a, 0x110b3efd, 0xdbf98641, 0xca0093fd,
            <a id="L144"></a>0x4e54f70e, 0x5f5fc9f3, 0x84a64fb2, 0x4ea6dc4f,
            <a id="L145"></a>0xead27321, 0xb58dbad2, 0x312bf560, 0x7f8d292f,
            <a id="L146"></a>0xac7766f3, 0x19fadc21, 0x28d12941, 0x575c006e,
            <a id="L147"></a>0xd014f9a8, 0xc9ee2589, 0xe13f0cc8, 0xb6630ca6,
        <a id="L148"></a>},
        <a id="L149"></a>[]uint32{
            <a id="L150"></a>0xd014f9a8, 0xc9ee2589, 0xe13f0cc8, 0xb6630ca6,
            <a id="L151"></a>0xc7b5a63, 0x1319eafe, 0xb0398890, 0x664cfbb4,
            <a id="L152"></a>0xdf7d925a, 0x1f62b09d, 0xa320626e, 0xd6757324,
            <a id="L153"></a>0x12c07647, 0xc01f22c7, 0xbc42d2f3, 0x7555114a,
            <a id="L154"></a>0x6efcd876, 0xd2df5480, 0x7c5df034, 0xc917c3b9,
            <a id="L155"></a>0x6ea30afc, 0xbc238cf6, 0xae82a4b4, 0xb54a338d,
            <a id="L156"></a>0x90884413, 0xd280860a, 0x12a12842, 0x1bc89739,
            <a id="L157"></a>0x7c1f13f7, 0x4208c219, 0xc021ae48, 0x969bf7b,
            <a id="L158"></a>0xcc7505eb, 0x3e17d1ee, 0x82296c51, 0xc9481133,
            <a id="L159"></a>0x2b3708a7, 0xf262d405, 0xbc3ebdbf, 0x4b617d62,
            <a id="L160"></a>0x2b7e1516, 0x28aed2a6, 0xabf71588, 0x9cf4f3c,
        <a id="L161"></a>},
    <a id="L162"></a>},
    <a id="L163"></a>KeyTest{
        <a id="L164"></a><span class="comment">// A.2.  Expansion of a 192-bit Cipher Key</span>
        <a id="L165"></a>[]byte{
            <a id="L166"></a>0x8e, 0x73, 0xb0, 0xf7, 0xda, 0x0e, 0x64, 0x52, 0xc8, 0x10, 0xf3, 0x2b, 0x80, 0x90, 0x79, 0xe5,
            <a id="L167"></a>0x62, 0xf8, 0xea, 0xd2, 0x52, 0x2c, 0x6b, 0x7b,
        <a id="L168"></a>},
        <a id="L169"></a>[]uint32{
            <a id="L170"></a>0x8e73b0f7, 0xda0e6452, 0xc810f32b, 0x809079e5,
            <a id="L171"></a>0x62f8ead2, 0x522c6b7b, 0xfe0c91f7, 0x2402f5a5,
            <a id="L172"></a>0xec12068e, 0x6c827f6b, 0x0e7a95b9, 0x5c56fec2,
            <a id="L173"></a>0x4db7b4bd, 0x69b54118, 0x85a74796, 0xe92538fd,
            <a id="L174"></a>0xe75fad44, 0xbb095386, 0x485af057, 0x21efb14f,
            <a id="L175"></a>0xa448f6d9, 0x4d6dce24, 0xaa326360, 0x113b30e6,
            <a id="L176"></a>0xa25e7ed5, 0x83b1cf9a, 0x27f93943, 0x6a94f767,
            <a id="L177"></a>0xc0a69407, 0xd19da4e1, 0xec1786eb, 0x6fa64971,
            <a id="L178"></a>0x485f7032, 0x22cb8755, 0xe26d1352, 0x33f0b7b3,
            <a id="L179"></a>0x40beeb28, 0x2f18a259, 0x6747d26b, 0x458c553e,
            <a id="L180"></a>0xa7e1466c, 0x9411f1df, 0x821f750a, 0xad07d753,
            <a id="L181"></a>0xca400538, 0x8fcc5006, 0x282d166a, 0xbc3ce7b5,
            <a id="L182"></a>0xe98ba06f, 0x448c773c, 0x8ecc7204, 0x01002202,
        <a id="L183"></a>},
        <a id="L184"></a>nil,
    <a id="L185"></a>},
    <a id="L186"></a>KeyTest{
        <a id="L187"></a><span class="comment">// A.3.  Expansion of a 256-bit Cipher Key</span>
        <a id="L188"></a>[]byte{
            <a id="L189"></a>0x60, 0x3d, 0xeb, 0x10, 0x15, 0xca, 0x71, 0xbe, 0x2b, 0x73, 0xae, 0xf0, 0x85, 0x7d, 0x77, 0x81,
            <a id="L190"></a>0x1f, 0x35, 0x2c, 0x07, 0x3b, 0x61, 0x08, 0xd7, 0x2d, 0x98, 0x10, 0xa3, 0x09, 0x14, 0xdf, 0xf4,
        <a id="L191"></a>},
        <a id="L192"></a>[]uint32{
            <a id="L193"></a>0x603deb10, 0x15ca71be, 0x2b73aef0, 0x857d7781,
            <a id="L194"></a>0x1f352c07, 0x3b6108d7, 0x2d9810a3, 0x0914dff4,
            <a id="L195"></a>0x9ba35411, 0x8e6925af, 0xa51a8b5f, 0x2067fcde,
            <a id="L196"></a>0xa8b09c1a, 0x93d194cd, 0xbe49846e, 0xb75d5b9a,
            <a id="L197"></a>0xd59aecb8, 0x5bf3c917, 0xfee94248, 0xde8ebe96,
            <a id="L198"></a>0xb5a9328a, 0x2678a647, 0x98312229, 0x2f6c79b3,
            <a id="L199"></a>0x812c81ad, 0xdadf48ba, 0x24360af2, 0xfab8b464,
            <a id="L200"></a>0x98c5bfc9, 0xbebd198e, 0x268c3ba7, 0x09e04214,
            <a id="L201"></a>0x68007bac, 0xb2df3316, 0x96e939e4, 0x6c518d80,
            <a id="L202"></a>0xc814e204, 0x76a9fb8a, 0x5025c02d, 0x59c58239,
            <a id="L203"></a>0xde136967, 0x6ccc5a71, 0xfa256395, 0x9674ee15,
            <a id="L204"></a>0x5886ca5d, 0x2e2f31d7, 0x7e0af1fa, 0x27cf73c3,
            <a id="L205"></a>0x749c47ab, 0x18501dda, 0xe2757e4f, 0x7401905a,
            <a id="L206"></a>0xcafaaae3, 0xe4d59b34, 0x9adf6ace, 0xbd10190d,
            <a id="L207"></a>0xfe4890d1, 0xe6188d0b, 0x046df344, 0x706c631e,
        <a id="L208"></a>},
        <a id="L209"></a>nil,
    <a id="L210"></a>},
<a id="L211"></a>}

<a id="L213"></a><span class="comment">// Test key expansion against FIPS 197 examples.</span>
<a id="L214"></a>func TestExpandKey(t *testing.T) {
<a id="L215"></a>L:
    <a id="L216"></a>for i, tt := range keyTests {
        <a id="L217"></a>enc := make([]uint32, len(tt.enc));
        <a id="L218"></a>var dec []uint32;
        <a id="L219"></a>if tt.dec != nil {
            <a id="L220"></a>dec = make([]uint32, len(tt.dec))
        <a id="L221"></a>}
        <a id="L222"></a>expandKey(tt.key, enc, dec);
        <a id="L223"></a>for j, v := range enc {
            <a id="L224"></a>if v != tt.enc[j] {
                <a id="L225"></a>t.Errorf(&#34;key %d: enc[%d] = %#x, want %#x&#34;, i, j, v, tt.enc[j]);
                <a id="L226"></a>continue L;
            <a id="L227"></a>}
        <a id="L228"></a>}
        <a id="L229"></a>if dec != nil {
            <a id="L230"></a>for j, v := range dec {
                <a id="L231"></a>if v != tt.dec[j] {
                    <a id="L232"></a>t.Errorf(&#34;key %d: dec[%d] = %#x, want %#x&#34;, i, j, v, tt.dec[j]);
                    <a id="L233"></a>continue L;
                <a id="L234"></a>}
            <a id="L235"></a>}
        <a id="L236"></a>}
    <a id="L237"></a>}
<a id="L238"></a>}

<a id="L240"></a><span class="comment">// Appendix B, C of FIPS 197: Cipher examples, Example vectors.</span>
<a id="L241"></a>type CryptTest struct {
    <a id="L242"></a>key []byte;
    <a id="L243"></a>in  []byte;
    <a id="L244"></a>out []byte;
<a id="L245"></a>}

<a id="L247"></a>var encryptTests = []CryptTest{
    <a id="L248"></a>CryptTest{
        <a id="L249"></a><span class="comment">// Appendix B.</span>
        <a id="L250"></a>[]byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c},
        <a id="L251"></a>[]byte{0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34},
        <a id="L252"></a>[]byte{0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b, 0x32},
    <a id="L253"></a>},
    <a id="L254"></a>CryptTest{
        <a id="L255"></a><span class="comment">// Appendix C.1.  AES-128</span>
        <a id="L256"></a>[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
        <a id="L257"></a>[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
        <a id="L258"></a>[]byte{0x69, 0xc4, 0xe0, 0xd8, 0x6a, 0x7b, 0x04, 0x30, 0xd8, 0xcd, 0xb7, 0x80, 0x70, 0xb4, 0xc5, 0x5a},
    <a id="L259"></a>},
    <a id="L260"></a>CryptTest{
        <a id="L261"></a><span class="comment">// Appendix C.2.  AES-192</span>
        <a id="L262"></a>[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
            <a id="L263"></a>0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
        <a id="L264"></a>},
        <a id="L265"></a>[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
        <a id="L266"></a>[]byte{0xdd, 0xa9, 0x7c, 0xa4, 0x86, 0x4c, 0xdf, 0xe0, 0x6e, 0xaf, 0x70, 0xa0, 0xec, 0x0d, 0x71, 0x91},
    <a id="L267"></a>},
    <a id="L268"></a>CryptTest{
        <a id="L269"></a><span class="comment">// Appendix C.3.  AES-256</span>
        <a id="L270"></a>[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
            <a id="L271"></a>0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
        <a id="L272"></a>},
        <a id="L273"></a>[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
        <a id="L274"></a>[]byte{0x8e, 0xa2, 0xb7, 0xca, 0x51, 0x67, 0x45, 0xbf, 0xea, 0xfc, 0x49, 0x90, 0x4b, 0x49, 0x60, 0x89},
    <a id="L275"></a>},
<a id="L276"></a>}

<a id="L278"></a><span class="comment">// Test encryptBlock against FIPS 197 examples.</span>
<a id="L279"></a>func TestEncryptBlock(t *testing.T) {
    <a id="L280"></a>for i, tt := range encryptTests {
        <a id="L281"></a>n := len(tt.key) + 28;
        <a id="L282"></a>enc := make([]uint32, n);
        <a id="L283"></a>dec := make([]uint32, n);
        <a id="L284"></a>expandKey(tt.key, enc, dec);
        <a id="L285"></a>out := make([]byte, len(tt.in));
        <a id="L286"></a>encryptBlock(enc, tt.in, out);
        <a id="L287"></a>for j, v := range out {
            <a id="L288"></a>if v != tt.out[j] {
                <a id="L289"></a>t.Errorf(&#34;encryptBlock %d: out[%d] = %#x, want %#x&#34;, i, j, v, tt.out[j]);
                <a id="L290"></a>break;
            <a id="L291"></a>}
        <a id="L292"></a>}
    <a id="L293"></a>}
<a id="L294"></a>}

<a id="L296"></a><span class="comment">// Test decryptBlock against FIPS 197 examples.</span>
<a id="L297"></a>func TestDecryptBlock(t *testing.T) {
    <a id="L298"></a>for i, tt := range encryptTests {
        <a id="L299"></a>n := len(tt.key) + 28;
        <a id="L300"></a>enc := make([]uint32, n);
        <a id="L301"></a>dec := make([]uint32, n);
        <a id="L302"></a>expandKey(tt.key, enc, dec);
        <a id="L303"></a>plain := make([]byte, len(tt.in));
        <a id="L304"></a>decryptBlock(dec, tt.out, plain);
        <a id="L305"></a>for j, v := range plain {
            <a id="L306"></a>if v != tt.in[j] {
                <a id="L307"></a>t.Errorf(&#34;decryptBlock %d: plain[%d] = %#x, want %#x&#34;, i, j, v, tt.in[j]);
                <a id="L308"></a>break;
            <a id="L309"></a>}
        <a id="L310"></a>}
    <a id="L311"></a>}
<a id="L312"></a>}

<a id="L314"></a><span class="comment">// Test Cipher Encrypt method against FIPS 197 examples.</span>
<a id="L315"></a>func TestCipherEncrypt(t *testing.T) {
    <a id="L316"></a>for i, tt := range encryptTests {
        <a id="L317"></a>c, err := NewCipher(tt.key);
        <a id="L318"></a>if err != nil {
            <a id="L319"></a>t.Errorf(&#34;NewCipher(%d bytes) = %s&#34;, len(tt.key), err);
            <a id="L320"></a>continue;
        <a id="L321"></a>}
        <a id="L322"></a>out := make([]byte, len(tt.in));
        <a id="L323"></a>c.Encrypt(tt.in, out);
        <a id="L324"></a>for j, v := range out {
            <a id="L325"></a>if v != tt.out[j] {
                <a id="L326"></a>t.Errorf(&#34;Cipher.Encrypt %d: out[%d] = %#x, want %#x&#34;, i, j, v, tt.out[j]);
                <a id="L327"></a>break;
            <a id="L328"></a>}
        <a id="L329"></a>}
    <a id="L330"></a>}
<a id="L331"></a>}

<a id="L333"></a><span class="comment">// Test Cipher Decrypt against FIPS 197 examples.</span>
<a id="L334"></a>func TestCipherDecrypt(t *testing.T) {
    <a id="L335"></a>for i, tt := range encryptTests {
        <a id="L336"></a>c, err := NewCipher(tt.key);
        <a id="L337"></a>if err != nil {
            <a id="L338"></a>t.Errorf(&#34;NewCipher(%d bytes) = %s&#34;, len(tt.key), err);
            <a id="L339"></a>continue;
        <a id="L340"></a>}
        <a id="L341"></a>plain := make([]byte, len(tt.in));
        <a id="L342"></a>c.Decrypt(tt.out, plain);
        <a id="L343"></a>for j, v := range plain {
            <a id="L344"></a>if v != tt.in[j] {
                <a id="L345"></a>t.Errorf(&#34;decryptBlock %d: plain[%d] = %#x, want %#x&#34;, i, j, v, tt.in[j]);
                <a id="L346"></a>break;
            <a id="L347"></a>}
        <a id="L348"></a>}
    <a id="L349"></a>}
<a id="L350"></a>}
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
