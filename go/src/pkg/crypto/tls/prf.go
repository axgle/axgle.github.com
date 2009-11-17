<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/tls/prf.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/tls/prf.go</h1>

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
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;crypto/hmac&#34;;
    <a id="L10"></a>&#34;crypto/md5&#34;;
    <a id="L11"></a>&#34;crypto/sha1&#34;;
    <a id="L12"></a>&#34;hash&#34;;
    <a id="L13"></a>&#34;os&#34;;
    <a id="L14"></a>&#34;strings&#34;;
<a id="L15"></a>)

<a id="L17"></a><span class="comment">// Split a premaster secret in two as specified in RFC 4346, section 5.</span>
<a id="L18"></a>func splitPreMasterSecret(secret []byte) (s1, s2 []byte) {
    <a id="L19"></a>s1 = secret[0 : (len(secret)+1)/2];
    <a id="L20"></a>s2 = secret[len(secret)/2 : len(secret)];
    <a id="L21"></a>return;
<a id="L22"></a>}

<a id="L24"></a><span class="comment">// pHash implements the P_hash function, as defined in RFC 4346, section 5.</span>
<a id="L25"></a>func pHash(result, secret, seed []byte, hash hash.Hash) {
    <a id="L26"></a>h := hmac.New(hash, secret);
    <a id="L27"></a>h.Write(seed);
    <a id="L28"></a>a := h.Sum();

    <a id="L30"></a>j := 0;
    <a id="L31"></a>for j &lt; len(result) {
        <a id="L32"></a>h.Reset();
        <a id="L33"></a>h.Write(a);
        <a id="L34"></a>h.Write(seed);
        <a id="L35"></a>b := h.Sum();
        <a id="L36"></a>todo := len(b);
        <a id="L37"></a>if j+todo &gt; len(result) {
            <a id="L38"></a>todo = len(result) - j
        <a id="L39"></a>}
        <a id="L40"></a>bytes.Copy(result[j:j+todo], b);
        <a id="L41"></a>j += todo;

        <a id="L43"></a>h.Reset();
        <a id="L44"></a>h.Write(a);
        <a id="L45"></a>a = h.Sum();
    <a id="L46"></a>}
<a id="L47"></a>}

<a id="L49"></a><span class="comment">// pRF11 implements the TLS 1.1 pseudo-random function, as defined in RFC 4346, section 5.</span>
<a id="L50"></a>func pRF11(result, secret, label, seed []byte) {
    <a id="L51"></a>hashSHA1 := sha1.New();
    <a id="L52"></a>hashMD5 := md5.New();

    <a id="L54"></a>labelAndSeed := make([]byte, len(label)+len(seed));
    <a id="L55"></a>bytes.Copy(labelAndSeed, label);
    <a id="L56"></a>bytes.Copy(labelAndSeed[len(label):len(labelAndSeed)], seed);

    <a id="L58"></a>s1, s2 := splitPreMasterSecret(secret);
    <a id="L59"></a>pHash(result, s1, labelAndSeed, hashMD5);
    <a id="L60"></a>result2 := make([]byte, len(result));
    <a id="L61"></a>pHash(result2, s2, labelAndSeed, hashSHA1);

    <a id="L63"></a>for i, b := range result2 {
        <a id="L64"></a>result[i] ^= b
    <a id="L65"></a>}
<a id="L66"></a>}

<a id="L68"></a>const (
    <a id="L69"></a>tlsRandomLength      = 32; <span class="comment">// Length of a random nonce in TLS 1.1.</span>
    <a id="L70"></a>masterSecretLength   = 48; <span class="comment">// Length of a master secret in TLS 1.1.</span>
    <a id="L71"></a>finishedVerifyLength = 12; <span class="comment">// Length of verify_data in a Finished message.</span>
<a id="L72"></a>)

<a id="L74"></a>var masterSecretLabel = strings.Bytes(&#34;master secret&#34;)
<a id="L75"></a>var keyExpansionLabel = strings.Bytes(&#34;key expansion&#34;)
<a id="L76"></a>var clientFinishedLabel = strings.Bytes(&#34;client finished&#34;)
<a id="L77"></a>var serverFinishedLabel = strings.Bytes(&#34;server finished&#34;)

<a id="L79"></a><span class="comment">// keysFromPreMasterSecret generates the connection keys from the pre master</span>
<a id="L80"></a><span class="comment">// secret, given the lengths of the MAC and cipher keys, as defined in RFC</span>
<a id="L81"></a><span class="comment">// 4346, section 6.3.</span>
<a id="L82"></a>func keysFromPreMasterSecret11(preMasterSecret, clientRandom, serverRandom []byte, macLen, keyLen int) (masterSecret, clientMAC, serverMAC, clientKey, serverKey []byte) {
    <a id="L83"></a>var seed [tlsRandomLength * 2]byte;
    <a id="L84"></a>bytes.Copy(seed[0:len(clientRandom)], clientRandom);
    <a id="L85"></a>bytes.Copy(seed[len(clientRandom):len(seed)], serverRandom);
    <a id="L86"></a>masterSecret = make([]byte, masterSecretLength);
    <a id="L87"></a>pRF11(masterSecret, preMasterSecret, masterSecretLabel, seed[0:len(seed)]);

    <a id="L89"></a>bytes.Copy(seed[0:len(clientRandom)], serverRandom);
    <a id="L90"></a>bytes.Copy(seed[len(serverRandom):len(seed)], clientRandom);

    <a id="L92"></a>n := 2*macLen + 2*keyLen;
    <a id="L93"></a>keyMaterial := make([]byte, n);
    <a id="L94"></a>pRF11(keyMaterial, masterSecret, keyExpansionLabel, seed[0:len(seed)]);
    <a id="L95"></a>clientMAC = keyMaterial[0:macLen];
    <a id="L96"></a>serverMAC = keyMaterial[macLen : macLen*2];
    <a id="L97"></a>clientKey = keyMaterial[macLen*2 : macLen*2+keyLen];
    <a id="L98"></a>serverKey = keyMaterial[macLen*2+keyLen : len(keyMaterial)];
    <a id="L99"></a>return;
<a id="L100"></a>}

<a id="L102"></a><span class="comment">// A finishedHash calculates the hash of a set of handshake messages suitable</span>
<a id="L103"></a><span class="comment">// for including in a Finished message.</span>
<a id="L104"></a>type finishedHash struct {
    <a id="L105"></a>clientMD5  hash.Hash;
    <a id="L106"></a>clientSHA1 hash.Hash;
    <a id="L107"></a>serverMD5  hash.Hash;
    <a id="L108"></a>serverSHA1 hash.Hash;
<a id="L109"></a>}

<a id="L111"></a>func newFinishedHash() finishedHash {
    <a id="L112"></a>return finishedHash{md5.New(), sha1.New(), md5.New(), sha1.New()}
<a id="L113"></a>}

<a id="L115"></a>func (h finishedHash) Write(msg []byte) (n int, err os.Error) {
    <a id="L116"></a>h.clientMD5.Write(msg);
    <a id="L117"></a>h.clientSHA1.Write(msg);
    <a id="L118"></a>h.serverMD5.Write(msg);
    <a id="L119"></a>h.serverSHA1.Write(msg);
    <a id="L120"></a>return len(msg), nil;
<a id="L121"></a>}

<a id="L123"></a><span class="comment">// finishedSum calculates the contents of the verify_data member of a Finished</span>
<a id="L124"></a><span class="comment">// message given the MD5 and SHA1 hashes of a set of handshake messages.</span>
<a id="L125"></a>func finishedSum(md5, sha1, label, masterSecret []byte) []byte {
    <a id="L126"></a>seed := make([]byte, len(md5)+len(sha1));
    <a id="L127"></a>bytes.Copy(seed, md5);
    <a id="L128"></a>bytes.Copy(seed[len(md5):len(seed)], sha1);
    <a id="L129"></a>out := make([]byte, finishedVerifyLength);
    <a id="L130"></a>pRF11(out, masterSecret, label, seed);
    <a id="L131"></a>return out;
<a id="L132"></a>}

<a id="L134"></a><span class="comment">// clientSum returns the contents of the verify_data member of a client&#39;s</span>
<a id="L135"></a><span class="comment">// Finished message.</span>
<a id="L136"></a>func (h finishedHash) clientSum(masterSecret []byte) []byte {
    <a id="L137"></a>md5 := h.clientMD5.Sum();
    <a id="L138"></a>sha1 := h.clientSHA1.Sum();
    <a id="L139"></a>return finishedSum(md5, sha1, clientFinishedLabel, masterSecret);
<a id="L140"></a>}

<a id="L142"></a><span class="comment">// serverSum returns the contents of the verify_data member of a server&#39;s</span>
<a id="L143"></a><span class="comment">// Finished message.</span>
<a id="L144"></a>func (h finishedHash) serverSum(masterSecret []byte) []byte {
    <a id="L145"></a>md5 := h.serverMD5.Sum();
    <a id="L146"></a>sha1 := h.serverSHA1.Sum();
    <a id="L147"></a>return finishedSum(md5, sha1, serverFinishedLabel, masterSecret);
<a id="L148"></a>}
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
