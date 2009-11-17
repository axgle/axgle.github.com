<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/rsa/rsa_test.go</title>

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
	<li>Thu Nov 12 16:00:00 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/rsa/rsa_test.go</h1>

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
    <a id="L10"></a>&#34;crypto/sha1&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;testing&#34;;
<a id="L13"></a>)

<a id="L15"></a>func TestKeyGeneration(t *testing.T) {
    <a id="L16"></a>urandom, err := os.Open(&#34;/dev/urandom&#34;, os.O_RDONLY, 0);
    <a id="L17"></a>if err != nil {
        <a id="L18"></a>t.Errorf(&#34;failed to open /dev/urandom&#34;)
    <a id="L19"></a>}

    <a id="L21"></a>priv, err := GenerateKey(urandom, 16);
    <a id="L22"></a>if err != nil {
        <a id="L23"></a>t.Errorf(&#34;failed to generate key&#34;)
    <a id="L24"></a>}
    <a id="L25"></a>pub := &amp;priv.PublicKey;
    <a id="L26"></a>m := big.NewInt(42);
    <a id="L27"></a>c := encrypt(new(big.Int), pub, m);
    <a id="L28"></a>m2, err := decrypt(nil, priv, c);
    <a id="L29"></a>if err != nil {
        <a id="L30"></a>t.Errorf(&#34;error while decrypting: %s&#34;, err)
    <a id="L31"></a>}
    <a id="L32"></a>if m.Cmp(m2) != 0 {
        <a id="L33"></a>t.Errorf(&#34;got:%v, want:%v (%s)&#34;, m2, m, priv)
    <a id="L34"></a>}

    <a id="L36"></a>m3, err := decrypt(urandom, priv, c);
    <a id="L37"></a>if err != nil {
        <a id="L38"></a>t.Errorf(&#34;error while decrypting (blind): %s&#34;, err)
    <a id="L39"></a>}
    <a id="L40"></a>if m.Cmp(m3) != 0 {
        <a id="L41"></a>t.Errorf(&#34;(blind) got:%v, want:%v&#34;, m3, m)
    <a id="L42"></a>}
<a id="L43"></a>}

<a id="L45"></a>type testEncryptOAEPMessage struct {
    <a id="L46"></a>in   []byte;
    <a id="L47"></a>seed []byte;
    <a id="L48"></a>out  []byte;
<a id="L49"></a>}

<a id="L51"></a>type testEncryptOAEPStruct struct {
    <a id="L52"></a>modulus string;
    <a id="L53"></a>e       int;
    <a id="L54"></a>d       string;
    <a id="L55"></a>msgs    []testEncryptOAEPMessage;
<a id="L56"></a>}

<a id="L58"></a>func TestEncryptOAEP(t *testing.T) {
    <a id="L59"></a>sha1 := sha1.New();
    <a id="L60"></a>n := new(big.Int);
    <a id="L61"></a>for i, test := range testEncryptOAEPData {
        <a id="L62"></a>n.SetString(test.modulus, 16);
        <a id="L63"></a>public := PublicKey{n, test.e};

        <a id="L65"></a>for j, message := range test.msgs {
            <a id="L66"></a>randomSource := bytes.NewBuffer(message.seed);
            <a id="L67"></a>out, err := EncryptOAEP(sha1, randomSource, &amp;public, message.in, nil);
            <a id="L68"></a>if err != nil {
                <a id="L69"></a>t.Errorf(&#34;#%d,%d error: %s&#34;, i, j, err)
            <a id="L70"></a>}
            <a id="L71"></a>if bytes.Compare(out, message.out) != 0 {
                <a id="L72"></a>t.Errorf(&#34;#%d,%d bad result: %s (want %s)&#34;, i, j, out, message.out)
            <a id="L73"></a>}
        <a id="L74"></a>}
    <a id="L75"></a>}
<a id="L76"></a>}

<a id="L78"></a>func TestDecryptOAEP(t *testing.T) {
    <a id="L79"></a>urandom, err := os.Open(&#34;/dev/urandom&#34;, os.O_RDONLY, 0);
    <a id="L80"></a>if err != nil {
        <a id="L81"></a>t.Errorf(&#34;Failed to open /dev/urandom&#34;)
    <a id="L82"></a>}

    <a id="L84"></a>sha1 := sha1.New();
    <a id="L85"></a>n := new(big.Int);
    <a id="L86"></a>d := new(big.Int);
    <a id="L87"></a>for i, test := range testEncryptOAEPData {
        <a id="L88"></a>n.SetString(test.modulus, 16);
        <a id="L89"></a>d.SetString(test.d, 16);
        <a id="L90"></a>private := PrivateKey{PublicKey{n, test.e}, d, nil, nil};

        <a id="L92"></a>for j, message := range test.msgs {
            <a id="L93"></a>out, err := DecryptOAEP(sha1, nil, &amp;private, message.out, nil);
            <a id="L94"></a>if err != nil {
                <a id="L95"></a>t.Errorf(&#34;#%d,%d error: %s&#34;, i, j, err)
            <a id="L96"></a>} else if bytes.Compare(out, message.in) != 0 {
                <a id="L97"></a>t.Errorf(&#34;#%d,%d bad result: %#v (want %#v)&#34;, i, j, out, message.in)
            <a id="L98"></a>}

            <a id="L100"></a><span class="comment">// Decrypt with blinding.</span>
            <a id="L101"></a>out, err = DecryptOAEP(sha1, urandom, &amp;private, message.out, nil);
            <a id="L102"></a>if err != nil {
                <a id="L103"></a>t.Errorf(&#34;#%d,%d (blind) error: %s&#34;, i, j, err)
            <a id="L104"></a>} else if bytes.Compare(out, message.in) != 0 {
                <a id="L105"></a>t.Errorf(&#34;#%d,%d (blind) bad result: %#v (want %#v)&#34;, i, j, out, message.in)
            <a id="L106"></a>}
        <a id="L107"></a>}
    <a id="L108"></a>}
<a id="L109"></a>}

<a id="L111"></a><span class="comment">// testEncryptOAEPData contains a subset of the vectors from RSA&#39;s &#34;Test vectors for RSA-OAEP&#34;.</span>
<a id="L112"></a>var testEncryptOAEPData = []testEncryptOAEPStruct{
    <a id="L113"></a><span class="comment">// Key 1</span>
    <a id="L114"></a>testEncryptOAEPStruct{&#34;a8b3b284af8eb50b387034a860f146c4919f318763cd6c5598c8ae4811a1e0abc4c7e0b082d693a5e7fced675cf4668512772c0cbc64a742c6c630f533c8cc72f62ae833c40bf25842e984bb78bdbf97c0107d55bdb662f5c4e0fab9845cb5148ef7392dd3aaff93ae1e6b667bb3d4247616d4f5ba10d4cfd226de88d39f16fb&#34;,
        <a id="L115"></a>65537,
        <a id="L116"></a>&#34;53339cfdb79fc8466a655c7316aca85c55fd8f6dd898fdaf119517ef4f52e8fd8e258df93fee180fa0e4ab29693cd83b152a553d4ac4d1812b8b9fa5af0e7f55fe7304df41570926f3311f15c4d65a732c483116ee3d3d2d0af3549ad9bf7cbfb78ad884f84d5beb04724dc7369b31def37d0cf539e9cfcdd3de653729ead5d1&#34;,
        <a id="L117"></a>[]testEncryptOAEPMessage{
            <a id="L118"></a><span class="comment">// Example 1.1</span>
            <a id="L119"></a>testEncryptOAEPMessage{
                <a id="L120"></a>[]byte{0x66, 0x28, 0x19, 0x4e, 0x12, 0x07, 0x3d, 0xb0,
                    <a id="L121"></a>0x3b, 0xa9, 0x4c, 0xda, 0x9e, 0xf9, 0x53, 0x23, 0x97,
                    <a id="L122"></a>0xd5, 0x0d, 0xba, 0x79, 0xb9, 0x87, 0x00, 0x4a, 0xfe,
                    <a id="L123"></a>0xfe, 0x34,
                <a id="L124"></a>},
                <a id="L125"></a>[]byte{0x18, 0xb7, 0x76, 0xea, 0x21, 0x06, 0x9d, 0x69,
                    <a id="L126"></a>0x77, 0x6a, 0x33, 0xe9, 0x6b, 0xad, 0x48, 0xe1, 0xdd,
                    <a id="L127"></a>0xa0, 0xa5, 0xef,
                <a id="L128"></a>},
                <a id="L129"></a>[]byte{0x35, 0x4f, 0xe6, 0x7b, 0x4a, 0x12, 0x6d, 0x5d,
                    <a id="L130"></a>0x35, 0xfe, 0x36, 0xc7, 0x77, 0x79, 0x1a, 0x3f, 0x7b,
                    <a id="L131"></a>0xa1, 0x3d, 0xef, 0x48, 0x4e, 0x2d, 0x39, 0x08, 0xaf,
                    <a id="L132"></a>0xf7, 0x22, 0xfa, 0xd4, 0x68, 0xfb, 0x21, 0x69, 0x6d,
                    <a id="L133"></a>0xe9, 0x5d, 0x0b, 0xe9, 0x11, 0xc2, 0xd3, 0x17, 0x4f,
                    <a id="L134"></a>0x8a, 0xfc, 0xc2, 0x01, 0x03, 0x5f, 0x7b, 0x6d, 0x8e,
                    <a id="L135"></a>0x69, 0x40, 0x2d, 0xe5, 0x45, 0x16, 0x18, 0xc2, 0x1a,
                    <a id="L136"></a>0x53, 0x5f, 0xa9, 0xd7, 0xbf, 0xc5, 0xb8, 0xdd, 0x9f,
                    <a id="L137"></a>0xc2, 0x43, 0xf8, 0xcf, 0x92, 0x7d, 0xb3, 0x13, 0x22,
                    <a id="L138"></a>0xd6, 0xe8, 0x81, 0xea, 0xa9, 0x1a, 0x99, 0x61, 0x70,
                    <a id="L139"></a>0xe6, 0x57, 0xa0, 0x5a, 0x26, 0x64, 0x26, 0xd9, 0x8c,
                    <a id="L140"></a>0x88, 0x00, 0x3f, 0x84, 0x77, 0xc1, 0x22, 0x70, 0x94,
                    <a id="L141"></a>0xa0, 0xd9, 0xfa, 0x1e, 0x8c, 0x40, 0x24, 0x30, 0x9c,
                    <a id="L142"></a>0xe1, 0xec, 0xcc, 0xb5, 0x21, 0x00, 0x35, 0xd4, 0x7a,
                    <a id="L143"></a>0xc7, 0x2e, 0x8a,
                <a id="L144"></a>},
            <a id="L145"></a>},
            <a id="L146"></a><span class="comment">// Example 1.2</span>
            <a id="L147"></a>testEncryptOAEPMessage{
                <a id="L148"></a>[]byte{0x75, 0x0c, 0x40, 0x47, 0xf5, 0x47, 0xe8, 0xe4,
                    <a id="L149"></a>0x14, 0x11, 0x85, 0x65, 0x23, 0x29, 0x8a, 0xc9, 0xba,
                    <a id="L150"></a>0xe2, 0x45, 0xef, 0xaf, 0x13, 0x97, 0xfb, 0xe5, 0x6f,
                    <a id="L151"></a>0x9d, 0xd5,
                <a id="L152"></a>},
                <a id="L153"></a>[]byte{0x0c, 0xc7, 0x42, 0xce, 0x4a, 0x9b, 0x7f, 0x32,
                    <a id="L154"></a>0xf9, 0x51, 0xbc, 0xb2, 0x51, 0xef, 0xd9, 0x25, 0xfe,
                    <a id="L155"></a>0x4f, 0xe3, 0x5f,
                <a id="L156"></a>},
                <a id="L157"></a>[]byte{0x64, 0x0d, 0xb1, 0xac, 0xc5, 0x8e, 0x05, 0x68,
                    <a id="L158"></a>0xfe, 0x54, 0x07, 0xe5, 0xf9, 0xb7, 0x01, 0xdf, 0xf8,
                    <a id="L159"></a>0xc3, 0xc9, 0x1e, 0x71, 0x6c, 0x53, 0x6f, 0xc7, 0xfc,
                    <a id="L160"></a>0xec, 0x6c, 0xb5, 0xb7, 0x1c, 0x11, 0x65, 0x98, 0x8d,
                    <a id="L161"></a>0x4a, 0x27, 0x9e, 0x15, 0x77, 0xd7, 0x30, 0xfc, 0x7a,
                    <a id="L162"></a>0x29, 0x93, 0x2e, 0x3f, 0x00, 0xc8, 0x15, 0x15, 0x23,
                    <a id="L163"></a>0x6d, 0x8d, 0x8e, 0x31, 0x01, 0x7a, 0x7a, 0x09, 0xdf,
                    <a id="L164"></a>0x43, 0x52, 0xd9, 0x04, 0xcd, 0xeb, 0x79, 0xaa, 0x58,
                    <a id="L165"></a>0x3a, 0xdc, 0xc3, 0x1e, 0xa6, 0x98, 0xa4, 0xc0, 0x52,
                    <a id="L166"></a>0x83, 0xda, 0xba, 0x90, 0x89, 0xbe, 0x54, 0x91, 0xf6,
                    <a id="L167"></a>0x7c, 0x1a, 0x4e, 0xe4, 0x8d, 0xc7, 0x4b, 0xbb, 0xe6,
                    <a id="L168"></a>0x64, 0x3a, 0xef, 0x84, 0x66, 0x79, 0xb4, 0xcb, 0x39,
                    <a id="L169"></a>0x5a, 0x35, 0x2d, 0x5e, 0xd1, 0x15, 0x91, 0x2d, 0xf6,
                    <a id="L170"></a>0x96, 0xff, 0xe0, 0x70, 0x29, 0x32, 0x94, 0x6d, 0x71,
                    <a id="L171"></a>0x49, 0x2b, 0x44,
                <a id="L172"></a>},
            <a id="L173"></a>},
            <a id="L174"></a><span class="comment">// Example 1.3</span>
            <a id="L175"></a>testEncryptOAEPMessage{
                <a id="L176"></a>[]byte{0xd9, 0x4a, 0xe0, 0x83, 0x2e, 0x64, 0x45, 0xce,
                    <a id="L177"></a>0x42, 0x33, 0x1c, 0xb0, 0x6d, 0x53, 0x1a, 0x82, 0xb1,
                    <a id="L178"></a>0xdb, 0x4b, 0xaa, 0xd3, 0x0f, 0x74, 0x6d, 0xc9, 0x16,
                    <a id="L179"></a>0xdf, 0x24, 0xd4, 0xe3, 0xc2, 0x45, 0x1f, 0xff, 0x59,
                    <a id="L180"></a>0xa6, 0x42, 0x3e, 0xb0, 0xe1, 0xd0, 0x2d, 0x4f, 0xe6,
                    <a id="L181"></a>0x46, 0xcf, 0x69, 0x9d, 0xfd, 0x81, 0x8c, 0x6e, 0x97,
                    <a id="L182"></a>0xb0, 0x51,
                <a id="L183"></a>},
                <a id="L184"></a>[]byte{0x25, 0x14, 0xdf, 0x46, 0x95, 0x75, 0x5a, 0x67,
                    <a id="L185"></a>0xb2, 0x88, 0xea, 0xf4, 0x90, 0x5c, 0x36, 0xee, 0xc6,
                    <a id="L186"></a>0x6f, 0xd2, 0xfd,
                <a id="L187"></a>},
                <a id="L188"></a>[]byte{0x42, 0x37, 0x36, 0xed, 0x03, 0x5f, 0x60, 0x26,
                    <a id="L189"></a>0xaf, 0x27, 0x6c, 0x35, 0xc0, 0xb3, 0x74, 0x1b, 0x36,
                    <a id="L190"></a>0x5e, 0x5f, 0x76, 0xca, 0x09, 0x1b, 0x4e, 0x8c, 0x29,
                    <a id="L191"></a>0xe2, 0xf0, 0xbe, 0xfe, 0xe6, 0x03, 0x59, 0x5a, 0xa8,
                    <a id="L192"></a>0x32, 0x2d, 0x60, 0x2d, 0x2e, 0x62, 0x5e, 0x95, 0xeb,
                    <a id="L193"></a>0x81, 0xb2, 0xf1, 0xc9, 0x72, 0x4e, 0x82, 0x2e, 0xca,
                    <a id="L194"></a>0x76, 0xdb, 0x86, 0x18, 0xcf, 0x09, 0xc5, 0x34, 0x35,
                    <a id="L195"></a>0x03, 0xa4, 0x36, 0x08, 0x35, 0xb5, 0x90, 0x3b, 0xc6,
                    <a id="L196"></a>0x37, 0xe3, 0x87, 0x9f, 0xb0, 0x5e, 0x0e, 0xf3, 0x26,
                    <a id="L197"></a>0x85, 0xd5, 0xae, 0xc5, 0x06, 0x7c, 0xd7, 0xcc, 0x96,
                    <a id="L198"></a>0xfe, 0x4b, 0x26, 0x70, 0xb6, 0xea, 0xc3, 0x06, 0x6b,
                    <a id="L199"></a>0x1f, 0xcf, 0x56, 0x86, 0xb6, 0x85, 0x89, 0xaa, 0xfb,
                    <a id="L200"></a>0x7d, 0x62, 0x9b, 0x02, 0xd8, 0xf8, 0x62, 0x5c, 0xa3,
                    <a id="L201"></a>0x83, 0x36, 0x24, 0xd4, 0x80, 0x0f, 0xb0, 0x81, 0xb1,
                    <a id="L202"></a>0xcf, 0x94, 0xeb,
                <a id="L203"></a>},
            <a id="L204"></a>},
        <a id="L205"></a>},
    <a id="L206"></a>},
    <a id="L207"></a><span class="comment">// Key 10</span>
    <a id="L208"></a>testEncryptOAEPStruct{&#34;ae45ed5601cec6b8cc05f803935c674ddbe0d75c4c09fd7951fc6b0caec313a8df39970c518bffba5ed68f3f0d7f22a4029d413f1ae07e4ebe9e4177ce23e7f5404b569e4ee1bdcf3c1fb03ef113802d4f855eb9b5134b5a7c8085adcae6fa2fa1417ec3763be171b0c62b760ede23c12ad92b980884c641f5a8fac26bdad4a03381a22fe1b754885094c82506d4019a535a286afeb271bb9ba592de18dcf600c2aeeae56e02f7cf79fc14cf3bdc7cd84febbbf950ca90304b2219a7aa063aefa2c3c1980e560cd64afe779585b6107657b957857efde6010988ab7de417fc88d8f384c4e6e72c3f943e0c31c0c4a5cc36f879d8a3ac9d7d59860eaada6b83bb&#34;,
        <a id="L209"></a>65537,
        <a id="L210"></a>&#34;056b04216fe5f354ac77250a4b6b0c8525a85c59b0bd80c56450a22d5f438e596a333aa875e291dd43f48cb88b9d5fc0d499f9fcd1c397f9afc070cd9e398c8d19e61db7c7410a6b2675dfbf5d345b804d201add502d5ce2dfcb091ce9997bbebe57306f383e4d588103f036f7e85d1934d152a323e4a8db451d6f4a5b1b0f102cc150e02feee2b88dea4ad4c1baccb24d84072d14e1d24a6771f7408ee30564fb86d4393a34bcf0b788501d193303f13a2284b001f0f649eaf79328d4ac5c430ab4414920a9460ed1b7bc40ec653e876d09abc509ae45b525190116a0c26101848298509c1c3bf3a483e7274054e15e97075036e989f60932807b5257751e79&#34;,
        <a id="L211"></a>[]testEncryptOAEPMessage{
            <a id="L212"></a><span class="comment">// Example 10.1</span>
            <a id="L213"></a>testEncryptOAEPMessage{
                <a id="L214"></a>[]byte{0x8b, 0xba, 0x6b, 0xf8, 0x2a, 0x6c, 0x0f, 0x86,
                    <a id="L215"></a>0xd5, 0xf1, 0x75, 0x6e, 0x97, 0x95, 0x68, 0x70, 0xb0,
                    <a id="L216"></a>0x89, 0x53, 0xb0, 0x6b, 0x4e, 0xb2, 0x05, 0xbc, 0x16,
                    <a id="L217"></a>0x94, 0xee,
                <a id="L218"></a>},
                <a id="L219"></a>[]byte{0x47, 0xe1, 0xab, 0x71, 0x19, 0xfe, 0xe5, 0x6c,
                    <a id="L220"></a>0x95, 0xee, 0x5e, 0xaa, 0xd8, 0x6f, 0x40, 0xd0, 0xaa,
                    <a id="L221"></a>0x63, 0xbd, 0x33,
                <a id="L222"></a>},
                <a id="L223"></a>[]byte{0x53, 0xea, 0x5d, 0xc0, 0x8c, 0xd2, 0x60, 0xfb,
                    <a id="L224"></a>0x3b, 0x85, 0x85, 0x67, 0x28, 0x7f, 0xa9, 0x15, 0x52,
                    <a id="L225"></a>0xc3, 0x0b, 0x2f, 0xeb, 0xfb, 0xa2, 0x13, 0xf0, 0xae,
                    <a id="L226"></a>0x87, 0x70, 0x2d, 0x06, 0x8d, 0x19, 0xba, 0xb0, 0x7f,
                    <a id="L227"></a>0xe5, 0x74, 0x52, 0x3d, 0xfb, 0x42, 0x13, 0x9d, 0x68,
                    <a id="L228"></a>0xc3, 0xc5, 0xaf, 0xee, 0xe0, 0xbf, 0xe4, 0xcb, 0x79,
                    <a id="L229"></a>0x69, 0xcb, 0xf3, 0x82, 0xb8, 0x04, 0xd6, 0xe6, 0x13,
                    <a id="L230"></a>0x96, 0x14, 0x4e, 0x2d, 0x0e, 0x60, 0x74, 0x1f, 0x89,
                    <a id="L231"></a>0x93, 0xc3, 0x01, 0x4b, 0x58, 0xb9, 0xb1, 0x95, 0x7a,
                    <a id="L232"></a>0x8b, 0xab, 0xcd, 0x23, 0xaf, 0x85, 0x4f, 0x4c, 0x35,
                    <a id="L233"></a>0x6f, 0xb1, 0x66, 0x2a, 0xa7, 0x2b, 0xfc, 0xc7, 0xe5,
                    <a id="L234"></a>0x86, 0x55, 0x9d, 0xc4, 0x28, 0x0d, 0x16, 0x0c, 0x12,
                    <a id="L235"></a>0x67, 0x85, 0xa7, 0x23, 0xeb, 0xee, 0xbe, 0xff, 0x71,
                    <a id="L236"></a>0xf1, 0x15, 0x94, 0x44, 0x0a, 0xae, 0xf8, 0x7d, 0x10,
                    <a id="L237"></a>0x79, 0x3a, 0x87, 0x74, 0xa2, 0x39, 0xd4, 0xa0, 0x4c,
                    <a id="L238"></a>0x87, 0xfe, 0x14, 0x67, 0xb9, 0xda, 0xf8, 0x52, 0x08,
                    <a id="L239"></a>0xec, 0x6c, 0x72, 0x55, 0x79, 0x4a, 0x96, 0xcc, 0x29,
                    <a id="L240"></a>0x14, 0x2f, 0x9a, 0x8b, 0xd4, 0x18, 0xe3, 0xc1, 0xfd,
                    <a id="L241"></a>0x67, 0x34, 0x4b, 0x0c, 0xd0, 0x82, 0x9d, 0xf3, 0xb2,
                    <a id="L242"></a>0xbe, 0xc6, 0x02, 0x53, 0x19, 0x62, 0x93, 0xc6, 0xb3,
                    <a id="L243"></a>0x4d, 0x3f, 0x75, 0xd3, 0x2f, 0x21, 0x3d, 0xd4, 0x5c,
                    <a id="L244"></a>0x62, 0x73, 0xd5, 0x05, 0xad, 0xf4, 0xcc, 0xed, 0x10,
                    <a id="L245"></a>0x57, 0xcb, 0x75, 0x8f, 0xc2, 0x6a, 0xee, 0xfa, 0x44,
                    <a id="L246"></a>0x12, 0x55, 0xed, 0x4e, 0x64, 0xc1, 0x99, 0xee, 0x07,
                    <a id="L247"></a>0x5e, 0x7f, 0x16, 0x64, 0x61, 0x82, 0xfd, 0xb4, 0x64,
                    <a id="L248"></a>0x73, 0x9b, 0x68, 0xab, 0x5d, 0xaf, 0xf0, 0xe6, 0x3e,
                    <a id="L249"></a>0x95, 0x52, 0x01, 0x68, 0x24, 0xf0, 0x54, 0xbf, 0x4d,
                    <a id="L250"></a>0x3c, 0x8c, 0x90, 0xa9, 0x7b, 0xb6, 0xb6, 0x55, 0x32,
                    <a id="L251"></a>0x84, 0xeb, 0x42, 0x9f, 0xcc,
                <a id="L252"></a>},
            <a id="L253"></a>},
        <a id="L254"></a>},
    <a id="L255"></a>},
<a id="L256"></a>}
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
