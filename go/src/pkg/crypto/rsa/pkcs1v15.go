<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/rsa/pkcs1v15.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/rsa/pkcs1v15.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package rsa

<a id="L7"></a>import (
    <a id="L8"></a>&#34;big&#34;;
    <a id="L9"></a>&#34;bytes&#34;;
    <a id="L10"></a>&#34;crypto/subtle&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;os&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">// This file implements encryption and decryption using PKCS#1 v1.5 padding.</span>

<a id="L17"></a><span class="comment">// EncryptPKCS1v15 encrypts the given message with RSA and the padding scheme from PKCS#1 v1.5.</span>
<a id="L18"></a><span class="comment">// The message must be no longer than the length of the public modulus minus 11 bytes.</span>
<a id="L19"></a><span class="comment">// WARNING: use of this function to encrypt plaintexts other than session keys</span>
<a id="L20"></a><span class="comment">// is dangerous. Use RSA OAEP in new protocols.</span>
<a id="L21"></a>func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) (out []byte, err os.Error) {
    <a id="L22"></a>k := (pub.N.Len() + 7) / 8;
    <a id="L23"></a>if len(msg) &gt; k-11 {
        <a id="L24"></a>err = MessageTooLongError{};
        <a id="L25"></a>return;
    <a id="L26"></a>}

    <a id="L28"></a><span class="comment">// EM = 0x02 || PS || 0x00 || M</span>
    <a id="L29"></a>em := make([]byte, k-1);
    <a id="L30"></a>em[0] = 2;
    <a id="L31"></a>ps, mm := em[1:len(em)-len(msg)-1], em[len(em)-len(msg):len(em)];
    <a id="L32"></a>err = nonZeroRandomBytes(ps, rand);
    <a id="L33"></a>if err != nil {
        <a id="L34"></a>return
    <a id="L35"></a>}
    <a id="L36"></a>em[len(em)-len(msg)-1] = 0;
    <a id="L37"></a>bytes.Copy(mm, msg);

    <a id="L39"></a>m := new(big.Int).SetBytes(em);
    <a id="L40"></a>c := encrypt(new(big.Int), pub, m);
    <a id="L41"></a>out = c.Bytes();
    <a id="L42"></a>return;
<a id="L43"></a>}

<a id="L45"></a><span class="comment">// DecryptPKCS1v15 decrypts a plaintext using RSA and the padding scheme from PKCS#1 v1.5.</span>
<a id="L46"></a><span class="comment">// If rand != nil, it uses RSA blinding to avoid timing side-channel attacks.</span>
<a id="L47"></a>func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) (out []byte, err os.Error) {
    <a id="L48"></a>valid, out, err := decryptPKCS1v15(rand, priv, ciphertext);
    <a id="L49"></a>if err == nil &amp;&amp; valid == 0 {
        <a id="L50"></a>err = DecryptionError{}
    <a id="L51"></a>}

    <a id="L53"></a>return;
<a id="L54"></a>}

<a id="L56"></a><span class="comment">// DecryptPKCS1v15SessionKey decrypts a session key using RSA and the padding scheme from PKCS#1 v1.5.</span>
<a id="L57"></a><span class="comment">// If rand != nil, it uses RSA blinding to avoid timing side-channel attacks.</span>
<a id="L58"></a><span class="comment">// It returns an error if the ciphertext is the wrong length or if the</span>
<a id="L59"></a><span class="comment">// ciphertext is greater than the public modulus. Otherwise, no error is</span>
<a id="L60"></a><span class="comment">// returned. If the padding is valid, the resulting plaintext message is copied</span>
<a id="L61"></a><span class="comment">// into key. Otherwise, key is unchanged. These alternatives occur in constant</span>
<a id="L62"></a><span class="comment">// time. It is intended that the user of this function generate a random</span>
<a id="L63"></a><span class="comment">// session key beforehand and continue the protocol with the resulting value.</span>
<a id="L64"></a><span class="comment">// This will remove any possibility that an attacker can learn any information</span>
<a id="L65"></a><span class="comment">// about the plaintext.</span>
<a id="L66"></a><span class="comment">// See ``Chosen Ciphertext Attacks Against Protocols Based on the RSA</span>
<a id="L67"></a><span class="comment">// Encryption Standard PKCS #1&#39;&#39;, Daniel Bleichenbacher, Advances in Cryptology</span>
<a id="L68"></a><span class="comment">// (Crypto &#39;98),</span>
<a id="L69"></a>func DecryptPKCS1v15SessionKey(rand io.Reader, priv *PrivateKey, ciphertext []byte, key []byte) (err os.Error) {
    <a id="L70"></a>k := (priv.N.Len() + 7) / 8;
    <a id="L71"></a>if k-(len(key)+3+8) &lt; 0 {
        <a id="L72"></a>err = DecryptionError{};
        <a id="L73"></a>return;
    <a id="L74"></a>}

    <a id="L76"></a>valid, msg, err := decryptPKCS1v15(rand, priv, ciphertext);
    <a id="L77"></a>if err != nil {
        <a id="L78"></a>return
    <a id="L79"></a>}

    <a id="L81"></a>valid &amp;= subtle.ConstantTimeEq(int32(len(msg)), int32(len(key)));
    <a id="L82"></a>subtle.ConstantTimeCopy(valid, key, msg);
    <a id="L83"></a>return;
<a id="L84"></a>}

<a id="L86"></a>func decryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) (valid int, msg []byte, err os.Error) {
    <a id="L87"></a>k := (priv.N.Len() + 7) / 8;
    <a id="L88"></a>if k &lt; 11 {
        <a id="L89"></a>err = DecryptionError{};
        <a id="L90"></a>return;
    <a id="L91"></a>}

    <a id="L93"></a>c := new(big.Int).SetBytes(ciphertext);
    <a id="L94"></a>m, err := decrypt(rand, priv, c);
    <a id="L95"></a>if err != nil {
        <a id="L96"></a>return
    <a id="L97"></a>}

    <a id="L99"></a>em := leftPad(m.Bytes(), k);
    <a id="L100"></a>firstByteIsZero := subtle.ConstantTimeByteEq(em[0], 0);
    <a id="L101"></a>secondByteIsTwo := subtle.ConstantTimeByteEq(em[1], 2);

    <a id="L103"></a><span class="comment">// The remainder of the plaintext must be a string of non-zero random</span>
    <a id="L104"></a><span class="comment">// octets, followed by a 0, followed by the message.</span>
    <a id="L105"></a><span class="comment">//   lookingForIndex: 1 iff we are still looking for the zero.</span>
    <a id="L106"></a><span class="comment">//   index: the offset of the first zero byte.</span>
    <a id="L107"></a>var lookingForIndex, index int;
    <a id="L108"></a>lookingForIndex = 1;

    <a id="L110"></a>for i := 2; i &lt; len(em); i++ {
        <a id="L111"></a>equals0 := subtle.ConstantTimeByteEq(em[i], 0);
        <a id="L112"></a>index = subtle.ConstantTimeSelect(lookingForIndex&amp;equals0, i, index);
        <a id="L113"></a>lookingForIndex = subtle.ConstantTimeSelect(equals0, 0, lookingForIndex);
    <a id="L114"></a>}

    <a id="L116"></a>valid = firstByteIsZero &amp; secondByteIsTwo &amp; (^lookingForIndex &amp; 1);
    <a id="L117"></a>msg = em[index+1 : len(em)];
    <a id="L118"></a>return;
<a id="L119"></a>}

<a id="L121"></a><span class="comment">// nonZeroRandomBytes fills the given slice with non-zero random octets.</span>
<a id="L122"></a>func nonZeroRandomBytes(s []byte, rand io.Reader) (err os.Error) {
    <a id="L123"></a>_, err = io.ReadFull(rand, s);
    <a id="L124"></a>if err != nil {
        <a id="L125"></a>return
    <a id="L126"></a>}

    <a id="L128"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L129"></a>for s[i] == 0 {
            <a id="L130"></a>_, err = rand.Read(s[i : i+1]);
            <a id="L131"></a>if err != nil {
                <a id="L132"></a>return
            <a id="L133"></a>}
        <a id="L134"></a>}
    <a id="L135"></a>}

    <a id="L137"></a>return;
<a id="L138"></a>}
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
