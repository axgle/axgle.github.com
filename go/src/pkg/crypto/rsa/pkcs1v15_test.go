<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/rsa/pkcs1v15_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/rsa/pkcs1v15_test.go</h1>

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
    <a id="L10"></a>&#34;encoding/base64&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;io&#34;;
    <a id="L13"></a>&#34;strings&#34;;
    <a id="L14"></a>&#34;testing&#34;;
    <a id="L15"></a>&#34;testing/quick&#34;;
<a id="L16"></a>)

<a id="L18"></a>func decodeBase64(in string) []byte {
    <a id="L19"></a>out := make([]byte, base64.StdEncoding.DecodedLen(len(in)));
    <a id="L20"></a>n, err := base64.StdEncoding.Decode(out, strings.Bytes(in));
    <a id="L21"></a>if err != nil {
        <a id="L22"></a>return nil
    <a id="L23"></a>}
    <a id="L24"></a>return out[0:n];
<a id="L25"></a>}

<a id="L27"></a>type DecryptPKCS1v15Test struct {
    <a id="L28"></a>in, out string;
<a id="L29"></a>}

<a id="L31"></a><span class="comment">// These test vectors were generated with `openssl rsautl -pkcs -encrypt`</span>
<a id="L32"></a>var decryptPKCS1v15Tests = []DecryptPKCS1v15Test{
    <a id="L33"></a>DecryptPKCS1v15Test{
        <a id="L34"></a>&#34;gIcUIoVkD6ATMBk/u/nlCZCCWRKdkfjCgFdo35VpRXLduiKXhNz1XupLLzTXAybEq15juc+EgY5o0DHv/nt3yg==&#34;,
        <a id="L35"></a>&#34;x&#34;,
    <a id="L36"></a>},
    <a id="L37"></a>DecryptPKCS1v15Test{
        <a id="L38"></a>&#34;Y7TOCSqofGhkRb+jaVRLzK8xw2cSo1IVES19utzv6hwvx+M8kFsoWQm5DzBeJCZTCVDPkTpavUuEbgp8hnUGDw==&#34;,
        <a id="L39"></a>&#34;testing.&#34;,
    <a id="L40"></a>},
    <a id="L41"></a>DecryptPKCS1v15Test{
        <a id="L42"></a>&#34;arReP9DJtEVyV2Dg3dDp4c/PSk1O6lxkoJ8HcFupoRorBZG+7+1fDAwT1olNddFnQMjmkb8vxwmNMoTAT/BFjQ==&#34;,
        <a id="L43"></a>&#34;testing.\n&#34;,
    <a id="L44"></a>},
    <a id="L45"></a>DecryptPKCS1v15Test{
        <a id="L46"></a>&#34;WtaBXIoGC54+vH0NH0CHHE+dRDOsMc/6BrfFu2lEqcKL9+uDuWaf+Xj9mrbQCjjZcpQuX733zyok/jsnqe/Ftw==&#34;,
        <a id="L47"></a>&#34;01234567890123456789012345678901234567890123456789012&#34;,
    <a id="L48"></a>},
<a id="L49"></a>}

<a id="L51"></a>func TestDecryptPKCS1v15(t *testing.T) {
    <a id="L52"></a>for i, test := range decryptPKCS1v15Tests {
        <a id="L53"></a>out, err := DecryptPKCS1v15(nil, rsaPrivateKey, decodeBase64(test.in));
        <a id="L54"></a>if err != nil {
            <a id="L55"></a>t.Errorf(&#34;#%d error decrypting&#34;, i)
        <a id="L56"></a>}
        <a id="L57"></a>want := strings.Bytes(test.out);
        <a id="L58"></a>if bytes.Compare(out, want) != 0 {
            <a id="L59"></a>t.Errorf(&#34;#%d got:%#v want:%#v&#34;, i, out, want)
        <a id="L60"></a>}
    <a id="L61"></a>}
<a id="L62"></a>}

<a id="L64"></a>func TestEncryptPKCS1v15(t *testing.T) {
    <a id="L65"></a>urandom, err := os.Open(&#34;/dev/urandom&#34;, os.O_RDONLY, 0);
    <a id="L66"></a>if err != nil {
        <a id="L67"></a>t.Errorf(&#34;Failed to open /dev/urandom&#34;)
    <a id="L68"></a>}
    <a id="L69"></a>k := (rsaPrivateKey.N.Len() + 7) / 8;

    <a id="L71"></a>tryEncryptDecrypt := func(in []byte, blind bool) bool {
        <a id="L72"></a>if len(in) &gt; k-11 {
            <a id="L73"></a>in = in[0 : k-11]
        <a id="L74"></a>}

        <a id="L76"></a>ciphertext, err := EncryptPKCS1v15(urandom, &amp;rsaPrivateKey.PublicKey, in);
        <a id="L77"></a>if err != nil {
            <a id="L78"></a>t.Errorf(&#34;error encrypting: %s&#34;, err);
            <a id="L79"></a>return false;
        <a id="L80"></a>}

        <a id="L82"></a>var rand io.Reader;
        <a id="L83"></a>if !blind {
            <a id="L84"></a>rand = nil
        <a id="L85"></a>} else {
            <a id="L86"></a>rand = urandom
        <a id="L87"></a>}
        <a id="L88"></a>plaintext, err := DecryptPKCS1v15(rand, rsaPrivateKey, ciphertext);
        <a id="L89"></a>if err != nil {
            <a id="L90"></a>t.Errorf(&#34;error decrypting: %s&#34;, err);
            <a id="L91"></a>return false;
        <a id="L92"></a>}

        <a id="L94"></a>if bytes.Compare(plaintext, in) != 0 {
            <a id="L95"></a>t.Errorf(&#34;output mismatch: %#v %#v&#34;, plaintext, in);
            <a id="L96"></a>return false;
        <a id="L97"></a>}
        <a id="L98"></a>return true;
    <a id="L99"></a>};

    <a id="L101"></a>quick.Check(tryEncryptDecrypt, nil);
<a id="L102"></a>}

<a id="L104"></a><span class="comment">// These test vectors were generated with `openssl rsautl -pkcs -encrypt`</span>
<a id="L105"></a>var decryptPKCS1v15SessionKeyTests = []DecryptPKCS1v15Test{
    <a id="L106"></a>DecryptPKCS1v15Test{
        <a id="L107"></a>&#34;e6ukkae6Gykq0fKzYwULpZehX+UPXYzMoB5mHQUDEiclRbOTqas4Y0E6nwns1BBpdvEJcilhl5zsox/6DtGsYg==&#34;,
        <a id="L108"></a>&#34;1234&#34;,
    <a id="L109"></a>},
    <a id="L110"></a>DecryptPKCS1v15Test{
        <a id="L111"></a>&#34;Dtis4uk/q/LQGGqGk97P59K03hkCIVFMEFZRgVWOAAhxgYpCRG0MX2adptt92l67IqMki6iVQyyt0TtX3IdtEw==&#34;,
        <a id="L112"></a>&#34;FAIL&#34;,
    <a id="L113"></a>},
    <a id="L114"></a>DecryptPKCS1v15Test{
        <a id="L115"></a>&#34;LIyFyCYCptPxrvTxpol8F3M7ZivlMsf53zs0vHRAv+rDIh2YsHS69ePMoPMe3TkOMZ3NupiL3takPxIs1sK+dw==&#34;,
        <a id="L116"></a>&#34;abcd&#34;,
    <a id="L117"></a>},
    <a id="L118"></a>DecryptPKCS1v15Test{
        <a id="L119"></a>&#34;bafnobel46bKy76JzqU/RIVOH0uAYvzUtauKmIidKgM0sMlvobYVAVQPeUQ/oTGjbIZ1v/6Gyi5AO4DtHruGdw==&#34;,
        <a id="L120"></a>&#34;FAIL&#34;,
    <a id="L121"></a>},
<a id="L122"></a>}

<a id="L124"></a>func TestEncryptPKCS1v15SessionKey(t *testing.T) {
    <a id="L125"></a>for i, test := range decryptPKCS1v15SessionKeyTests {
        <a id="L126"></a>key := strings.Bytes(&#34;FAIL&#34;);
        <a id="L127"></a>err := DecryptPKCS1v15SessionKey(nil, rsaPrivateKey, decodeBase64(test.in), key);
        <a id="L128"></a>if err != nil {
            <a id="L129"></a>t.Errorf(&#34;#%d error decrypting&#34;, i)
        <a id="L130"></a>}
        <a id="L131"></a>want := strings.Bytes(test.out);
        <a id="L132"></a>if bytes.Compare(key, want) != 0 {
            <a id="L133"></a>t.Errorf(&#34;#%d got:%#v want:%#v&#34;, i, key, want)
        <a id="L134"></a>}
    <a id="L135"></a>}
<a id="L136"></a>}

<a id="L138"></a>func TestNonZeroRandomBytes(t *testing.T) {
    <a id="L139"></a>urandom, err := os.Open(&#34;/dev/urandom&#34;, os.O_RDONLY, 0);
    <a id="L140"></a>if err != nil {
        <a id="L141"></a>t.Errorf(&#34;Failed to open /dev/urandom&#34;)
    <a id="L142"></a>}

    <a id="L144"></a>b := make([]byte, 512);
    <a id="L145"></a>err = nonZeroRandomBytes(b, urandom);
    <a id="L146"></a>if err != nil {
        <a id="L147"></a>t.Errorf(&#34;returned error: %s&#34;, err)
    <a id="L148"></a>}
    <a id="L149"></a>for _, b := range b {
        <a id="L150"></a>if b == 0 {
            <a id="L151"></a>t.Errorf(&#34;Zero octet found&#34;);
            <a id="L152"></a>return;
        <a id="L153"></a>}
    <a id="L154"></a>}
<a id="L155"></a>}

<a id="L157"></a>func bigFromString(s string) *big.Int {
    <a id="L158"></a>ret := new(big.Int);
    <a id="L159"></a>ret.SetString(s, 10);
    <a id="L160"></a>return ret;
<a id="L161"></a>}

<a id="L163"></a><span class="comment">// In order to generate new test vectors you&#39;ll need the PEM form of this key:</span>
<a id="L164"></a><span class="comment">// -----BEGIN RSA PRIVATE KEY-----</span>
<a id="L165"></a><span class="comment">// MIIBOgIBAAJBALKZD0nEffqM1ACuak0bijtqE2QrI/KLADv7l3kK3ppMyCuLKoF0</span>
<a id="L166"></a><span class="comment">// fd7Ai2KW5ToIwzFofvJcS/STa6HA5gQenRUCAwEAAQJBAIq9amn00aS0h/CrjXqu</span>
<a id="L167"></a><span class="comment">// /ThglAXJmZhOMPVn4eiu7/ROixi9sex436MaVeMqSNf7Ex9a8fRNfWss7Sqd9eWu</span>
<a id="L168"></a><span class="comment">// RTUCIQDasvGASLqmjeffBNLTXV2A5g4t+kLVCpsEIZAycV5GswIhANEPLmax0ME/</span>
<a id="L169"></a><span class="comment">// EO+ZJ79TJKN5yiGBRsv5yvx5UiHxajEXAiAhAol5N4EUyq6I9w1rYdhPMGpLfk7A</span>
<a id="L170"></a><span class="comment">// IU2snfRJ6Nq2CQIgFrPsWRCkV+gOYcajD17rEqmuLrdIRexpg8N1DOSXoJ8CIGlS</span>
<a id="L171"></a><span class="comment">// tAboUGBxTDq3ZroNism3DaMIbKPyYrAqhKov1h5V</span>
<a id="L172"></a><span class="comment">// -----END RSA PRIVATE KEY-----</span>

<a id="L174"></a>var rsaPrivateKey = &amp;PrivateKey{
    <a id="L175"></a>PublicKey: PublicKey{
        <a id="L176"></a>N: bigFromString(&#34;9353930466774385905609975137998169297361893554149986716853295022578535724979677252958524466350471210367835187480748268864277464700638583474144061408845077&#34;),
        <a id="L177"></a>E: 65537,
    <a id="L178"></a>},
    <a id="L179"></a>D: bigFromString(&#34;7266398431328116344057699379749222532279343923819063639497049039389899328538543087657733766554155839834519529439851673014800261285757759040931985506583861&#34;),
    <a id="L180"></a>P: bigFromString(&#34;98920366548084643601728869055592650835572950932266967461790948584315647051443&#34;),
    <a id="L181"></a>Q: bigFromString(&#34;94560208308847015747498523884063394671606671904944666360068158221458669711639&#34;),
<a id="L182"></a>}
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
