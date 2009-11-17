<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/tls/prf_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/tls/prf_test.go</h1>

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
    <a id="L8"></a>&#34;encoding/hex&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a>type testSplitPreMasterSecretTest struct {
    <a id="L13"></a>in, out1, out2 string;
<a id="L14"></a>}

<a id="L16"></a>var testSplitPreMasterSecretTests = []testSplitPreMasterSecretTest{
    <a id="L17"></a>testSplitPreMasterSecretTest{&#34;&#34;, &#34;&#34;, &#34;&#34;},
    <a id="L18"></a>testSplitPreMasterSecretTest{&#34;00&#34;, &#34;00&#34;, &#34;00&#34;},
    <a id="L19"></a>testSplitPreMasterSecretTest{&#34;0011&#34;, &#34;00&#34;, &#34;11&#34;},
    <a id="L20"></a>testSplitPreMasterSecretTest{&#34;001122&#34;, &#34;0011&#34;, &#34;1122&#34;},
    <a id="L21"></a>testSplitPreMasterSecretTest{&#34;00112233&#34;, &#34;0011&#34;, &#34;2233&#34;},
<a id="L22"></a>}

<a id="L24"></a>func TestSplitPreMasterSecret(t *testing.T) {
    <a id="L25"></a>for i, test := range testSplitPreMasterSecretTests {
        <a id="L26"></a>in, _ := hex.DecodeString(test.in);
        <a id="L27"></a>out1, out2 := splitPreMasterSecret(in);
        <a id="L28"></a>s1 := hex.EncodeToString(out1);
        <a id="L29"></a>s2 := hex.EncodeToString(out2);
        <a id="L30"></a>if s1 != test.out1 || s2 != test.out2 {
            <a id="L31"></a>t.Errorf(&#34;#%d: got: (%s, %s) want: (%s, %s)&#34;, i, s1, s2, test.out1, test.out2)
        <a id="L32"></a>}
    <a id="L33"></a>}
<a id="L34"></a>}

<a id="L36"></a>type testKeysFromTest struct {
    <a id="L37"></a>preMasterSecret            string;
    <a id="L38"></a>clientRandom, serverRandom string;
    <a id="L39"></a>masterSecret               string;
    <a id="L40"></a>clientMAC, serverMAC       string;
    <a id="L41"></a>clientKey, serverKey       string;
    <a id="L42"></a>macLen, keyLen             int;
<a id="L43"></a>}

<a id="L45"></a>func TestKeysFromPreMasterSecret(t *testing.T) {
    <a id="L46"></a>for i, test := range testKeysFromTests {
        <a id="L47"></a>in, _ := hex.DecodeString(test.preMasterSecret);
        <a id="L48"></a>clientRandom, _ := hex.DecodeString(test.clientRandom);
        <a id="L49"></a>serverRandom, _ := hex.DecodeString(test.serverRandom);
        <a id="L50"></a>master, clientMAC, serverMAC, clientKey, serverKey := keysFromPreMasterSecret11(in, clientRandom, serverRandom, test.macLen, test.keyLen);
        <a id="L51"></a>masterString := hex.EncodeToString(master);
        <a id="L52"></a>clientMACString := hex.EncodeToString(clientMAC);
        <a id="L53"></a>serverMACString := hex.EncodeToString(serverMAC);
        <a id="L54"></a>clientKeyString := hex.EncodeToString(clientKey);
        <a id="L55"></a>serverKeyString := hex.EncodeToString(serverKey);
        <a id="L56"></a>if masterString != test.masterSecret ||
            <a id="L57"></a>clientMACString != test.clientMAC ||
            <a id="L58"></a>serverMACString != test.serverMAC ||
            <a id="L59"></a>clientKeyString != test.clientKey ||
            <a id="L60"></a>serverKeyString != test.serverKey {
            <a id="L61"></a>t.Errorf(&#34;#%d: got: (%s, %s, %s, %s, %s) want: (%s, %s, %s, %s %s)&#34;, i, masterString, clientMACString, serverMACString, clientKeyString, serverMACString, test.masterSecret, test.clientMAC, test.serverMAC, test.clientKey, test.serverKey)
        <a id="L62"></a>}
    <a id="L63"></a>}
<a id="L64"></a>}

<a id="L66"></a><span class="comment">// These test vectors were generated from GnuTLS using `gnutls-cli --insecure -d 9 `</span>
<a id="L67"></a>var testKeysFromTests = []testKeysFromTest{
    <a id="L68"></a>testKeysFromTest{
        <a id="L69"></a>&#34;0302cac83ad4b1db3b9ab49ad05957de2a504a634a386fc600889321e1a971f57479466830ac3e6f468e87f5385fa0c5&#34;,
        <a id="L70"></a>&#34;4ae66303755184a3917fcb44880605fcc53baa01912b22ed94473fc69cebd558&#34;,
        <a id="L71"></a>&#34;4ae663020ec16e6bb5130be918cfcafd4d765979a3136a5d50c593446e4e44db&#34;,
        <a id="L72"></a>&#34;3d851bab6e5556e959a16bc36d66cfae32f672bfa9ecdef6096cbb1b23472df1da63dbbd9827606413221d149ed08ceb&#34;,
        <a id="L73"></a>&#34;805aaa19b3d2c0a0759a4b6c9959890e08480119&#34;,
        <a id="L74"></a>&#34;2d22f9fe519c075c16448305ceee209fc24ad109&#34;,
        <a id="L75"></a>&#34;d50b5771244f850cd8117a9ccafe2cf1&#34;,
        <a id="L76"></a>&#34;e076e33206b30507a85c32855acd0919&#34;,
        <a id="L77"></a>20,
        <a id="L78"></a>16,
    <a id="L79"></a>},
    <a id="L80"></a>testKeysFromTest{
        <a id="L81"></a>&#34;03023f7527316bc12cbcd69e4b9e8275d62c028f27e65c745cfcddc7ce01bd3570a111378b63848127f1c36e5f9e4890&#34;,
        <a id="L82"></a>&#34;4ae66364b5ea56b20ce4e25555aed2d7e67f42788dd03f3fee4adae0459ab106&#34;,
        <a id="L83"></a>&#34;4ae66363ab815cbf6a248b87d6b556184e945e9b97fbdf247858b0bdafacfa1c&#34;,
        <a id="L84"></a>&#34;7d64be7c80c59b740200b4b9c26d0baaa1c5ae56705acbcf2307fe62beb4728c19392c83f20483801cce022c77645460&#34;,
        <a id="L85"></a>&#34;97742ed60a0554ca13f04f97ee193177b971e3b0&#34;,
        <a id="L86"></a>&#34;37068751700400e03a8477a5c7eec0813ab9e0dc&#34;,
        <a id="L87"></a>&#34;207cddbc600d2a200abac6502053ee5c&#34;,
        <a id="L88"></a>&#34;df3f94f6e1eacc753b815fe16055cd43&#34;,
        <a id="L89"></a>20,
        <a id="L90"></a>16,
    <a id="L91"></a>},
    <a id="L92"></a>testKeysFromTest{
        <a id="L93"></a>&#34;832d515f1d61eebb2be56ba0ef79879efb9b527504abb386fb4310ed5d0e3b1f220d3bb6b455033a2773e6d8bdf951d278a187482b400d45deb88a5d5a6bb7d6a7a1decc04eb9ef0642876cd4a82d374d3b6ff35f0351dc5d411104de431375355addc39bfb1f6329fb163b0bc298d658338930d07d313cd980a7e3d9196cac1&#34;,
        <a id="L94"></a>&#34;4ae663b2ee389c0de147c509d8f18f5052afc4aaf9699efe8cb05ece883d3a5e&#34;,
        <a id="L95"></a>&#34;4ae664d503fd4cff50cfc1fb8fc606580f87b0fcdac9554ba0e01d785bdf278e&#34;,
        <a id="L96"></a>&#34;1aff2e7a2c4279d0126f57a65a77a8d9d0087cf2733366699bec27eb53d5740705a8574bb1acc2abbe90e44f0dd28d6c&#34;,
        <a id="L97"></a>&#34;3c7647c93c1379a31a609542aa44e7f117a70085&#34;,
        <a id="L98"></a>&#34;0d73102994be74a575a3ead8532590ca32a526d4&#34;,
        <a id="L99"></a>&#34;ac7581b0b6c10d85bbd905ffbf36c65e&#34;,
        <a id="L100"></a>&#34;ff07edde49682b45466bd2e39464b306&#34;,
        <a id="L101"></a>20,
        <a id="L102"></a>16,
    <a id="L103"></a>},
<a id="L104"></a>}
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
