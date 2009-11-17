<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/tls/handshake_messages_test.go</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/tls/handshake_messages_test.go</h1>

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
    <a id="L8"></a>&#34;rand&#34;;
    <a id="L9"></a>&#34;reflect&#34;;
    <a id="L10"></a>&#34;testing&#34;;
    <a id="L11"></a>&#34;testing/quick&#34;;
<a id="L12"></a>)

<a id="L14"></a>var tests = []interface{}{
    <a id="L15"></a>&amp;clientHelloMsg{},
    <a id="L16"></a>&amp;clientKeyExchangeMsg{},
    <a id="L17"></a>&amp;finishedMsg{},
<a id="L18"></a>}

<a id="L20"></a>type testMessage interface {
    <a id="L21"></a>marshal() []byte;
    <a id="L22"></a>unmarshal([]byte) bool;
<a id="L23"></a>}

<a id="L25"></a>func TestMarshalUnmarshal(t *testing.T) {
    <a id="L26"></a>rand := rand.New(rand.NewSource(0));
    <a id="L27"></a>for i, iface := range tests {
        <a id="L28"></a>ty := reflect.NewValue(iface).Type();

        <a id="L30"></a>for j := 0; j &lt; 100; j++ {
            <a id="L31"></a>v, ok := quick.Value(ty, rand);
            <a id="L32"></a>if !ok {
                <a id="L33"></a>t.Errorf(&#34;#%d: failed to create value&#34;, i);
                <a id="L34"></a>break;
            <a id="L35"></a>}

            <a id="L37"></a>m1 := v.Interface().(testMessage);
            <a id="L38"></a>marshaled := m1.marshal();
            <a id="L39"></a>m2 := iface.(testMessage);
            <a id="L40"></a>if !m2.unmarshal(marshaled) {
                <a id="L41"></a>t.Errorf(&#34;#%d failed to unmarshal %#v&#34;, i, m1);
                <a id="L42"></a>break;
            <a id="L43"></a>}
            <a id="L44"></a>m2.marshal(); <span class="comment">// to fill any marshal cache in the message</span>

            <a id="L46"></a>if !reflect.DeepEqual(m1, m2) {
                <a id="L47"></a>t.Errorf(&#34;#%d got:%#v want:%#v&#34;, i, m1, m2);
                <a id="L48"></a>break;
            <a id="L49"></a>}

            <a id="L51"></a><span class="comment">// Now check that all prefixes are invalid.</span>
            <a id="L52"></a>for j := 0; j &lt; len(marshaled); j++ {
                <a id="L53"></a>if m2.unmarshal(marshaled[0:j]) {
                    <a id="L54"></a>t.Errorf(&#34;#%d unmarshaled a prefix of length %d of %#v&#34;, i, j, m1);
                    <a id="L55"></a>break;
                <a id="L56"></a>}
            <a id="L57"></a>}
        <a id="L58"></a>}
    <a id="L59"></a>}
<a id="L60"></a>}

<a id="L62"></a>func randomBytes(n int, rand *rand.Rand) []byte {
    <a id="L63"></a>r := make([]byte, n);
    <a id="L64"></a>for i := 0; i &lt; n; i++ {
        <a id="L65"></a>r[i] = byte(rand.Int31())
    <a id="L66"></a>}
    <a id="L67"></a>return r;
<a id="L68"></a>}

<a id="L70"></a>func (*clientHelloMsg) Generate(rand *rand.Rand, size int) reflect.Value {
    <a id="L71"></a>m := &amp;clientHelloMsg{};
    <a id="L72"></a>m.major = uint8(rand.Intn(256));
    <a id="L73"></a>m.minor = uint8(rand.Intn(256));
    <a id="L74"></a>m.random = randomBytes(32, rand);
    <a id="L75"></a>m.sessionId = randomBytes(rand.Intn(32), rand);
    <a id="L76"></a>m.cipherSuites = make([]uint16, rand.Intn(63)+1);
    <a id="L77"></a>for i := 0; i &lt; len(m.cipherSuites); i++ {
        <a id="L78"></a>m.cipherSuites[i] = uint16(rand.Int31())
    <a id="L79"></a>}
    <a id="L80"></a>m.compressionMethods = randomBytes(rand.Intn(63)+1, rand);

    <a id="L82"></a>return reflect.NewValue(m);
<a id="L83"></a>}

<a id="L85"></a>func (*clientKeyExchangeMsg) Generate(rand *rand.Rand, size int) reflect.Value {
    <a id="L86"></a>m := &amp;clientKeyExchangeMsg{};
    <a id="L87"></a>m.ciphertext = randomBytes(rand.Intn(1000), rand);
    <a id="L88"></a>return reflect.NewValue(m);
<a id="L89"></a>}

<a id="L91"></a>func (*finishedMsg) Generate(rand *rand.Rand, size int) reflect.Value {
    <a id="L92"></a>m := &amp;finishedMsg{};
    <a id="L93"></a>m.verifyData = randomBytes(12, rand);
    <a id="L94"></a>return reflect.NewValue(m);
<a id="L95"></a>}
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
