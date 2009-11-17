<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/tls/handshake_messages.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/tls/handshake_messages.go</h1>

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
<a id="L9"></a>)

<a id="L11"></a>type clientHelloMsg struct {
    <a id="L12"></a>raw                []byte;
    <a id="L13"></a>major, minor       uint8;
    <a id="L14"></a>random             []byte;
    <a id="L15"></a>sessionId          []byte;
    <a id="L16"></a>cipherSuites       []uint16;
    <a id="L17"></a>compressionMethods []uint8;
<a id="L18"></a>}

<a id="L20"></a>func (m *clientHelloMsg) marshal() []byte {
    <a id="L21"></a>if m.raw != nil {
        <a id="L22"></a>return m.raw
    <a id="L23"></a>}

    <a id="L25"></a>length := 2 + 32 + 1 + len(m.sessionId) + 2 + len(m.cipherSuites)*2 + 1 + len(m.compressionMethods);
    <a id="L26"></a>x := make([]byte, 4+length);
    <a id="L27"></a>x[0] = typeClientHello;
    <a id="L28"></a>x[1] = uint8(length &gt;&gt; 16);
    <a id="L29"></a>x[2] = uint8(length &gt;&gt; 8);
    <a id="L30"></a>x[3] = uint8(length);
    <a id="L31"></a>x[4] = m.major;
    <a id="L32"></a>x[5] = m.minor;
    <a id="L33"></a>bytes.Copy(x[6:38], m.random);
    <a id="L34"></a>x[38] = uint8(len(m.sessionId));
    <a id="L35"></a>bytes.Copy(x[39:39+len(m.sessionId)], m.sessionId);
    <a id="L36"></a>y := x[39+len(m.sessionId) : len(x)];
    <a id="L37"></a>y[0] = uint8(len(m.cipherSuites) &gt;&gt; 7);
    <a id="L38"></a>y[1] = uint8(len(m.cipherSuites) &lt;&lt; 1);
    <a id="L39"></a>for i, suite := range m.cipherSuites {
        <a id="L40"></a>y[2+i*2] = uint8(suite &gt;&gt; 8);
        <a id="L41"></a>y[3+i*2] = uint8(suite);
    <a id="L42"></a>}
    <a id="L43"></a>z := y[2+len(m.cipherSuites)*2 : len(y)];
    <a id="L44"></a>z[0] = uint8(len(m.compressionMethods));
    <a id="L45"></a>bytes.Copy(z[1:len(z)], m.compressionMethods);
    <a id="L46"></a>m.raw = x;

    <a id="L48"></a>return x;
<a id="L49"></a>}

<a id="L51"></a>func (m *clientHelloMsg) unmarshal(data []byte) bool {
    <a id="L52"></a>if len(data) &lt; 39 {
        <a id="L53"></a>return false
    <a id="L54"></a>}
    <a id="L55"></a>m.raw = data;
    <a id="L56"></a>m.major = data[4];
    <a id="L57"></a>m.minor = data[5];
    <a id="L58"></a>m.random = data[6:38];
    <a id="L59"></a>sessionIdLen := int(data[38]);
    <a id="L60"></a>if sessionIdLen &gt; 32 || len(data) &lt; 39+sessionIdLen {
        <a id="L61"></a>return false
    <a id="L62"></a>}
    <a id="L63"></a>m.sessionId = data[39 : 39+sessionIdLen];
    <a id="L64"></a>data = data[39+sessionIdLen : len(data)];
    <a id="L65"></a>if len(data) &lt; 2 {
        <a id="L66"></a>return false
    <a id="L67"></a>}
    <a id="L68"></a><span class="comment">// cipherSuiteLen is the number of bytes of cipher suite numbers. Since</span>
    <a id="L69"></a><span class="comment">// they are uint16s, the number must be even.</span>
    <a id="L70"></a>cipherSuiteLen := int(data[0])&lt;&lt;8 | int(data[1]);
    <a id="L71"></a>if cipherSuiteLen%2 == 1 || len(data) &lt; 2+cipherSuiteLen {
        <a id="L72"></a>return false
    <a id="L73"></a>}
    <a id="L74"></a>numCipherSuites := cipherSuiteLen / 2;
    <a id="L75"></a>m.cipherSuites = make([]uint16, numCipherSuites);
    <a id="L76"></a>for i := 0; i &lt; numCipherSuites; i++ {
        <a id="L77"></a>m.cipherSuites[i] = uint16(data[2+2*i])&lt;&lt;8 | uint16(data[3+2*i])
    <a id="L78"></a>}
    <a id="L79"></a>data = data[2+cipherSuiteLen : len(data)];
    <a id="L80"></a>if len(data) &lt; 2 {
        <a id="L81"></a>return false
    <a id="L82"></a>}
    <a id="L83"></a>compressionMethodsLen := int(data[0]);
    <a id="L84"></a>if len(data) &lt; 1+compressionMethodsLen {
        <a id="L85"></a>return false
    <a id="L86"></a>}
    <a id="L87"></a>m.compressionMethods = data[1 : 1+compressionMethodsLen];

    <a id="L89"></a><span class="comment">// A ClientHello may be following by trailing data: RFC 4346 section 7.4.1.2</span>
    <a id="L90"></a>return true;
<a id="L91"></a>}

<a id="L93"></a>type serverHelloMsg struct {
    <a id="L94"></a>raw               []byte;
    <a id="L95"></a>major, minor      uint8;
    <a id="L96"></a>random            []byte;
    <a id="L97"></a>sessionId         []byte;
    <a id="L98"></a>cipherSuite       uint16;
    <a id="L99"></a>compressionMethod uint8;
<a id="L100"></a>}

<a id="L102"></a>func (m *serverHelloMsg) marshal() []byte {
    <a id="L103"></a>if m.raw != nil {
        <a id="L104"></a>return m.raw
    <a id="L105"></a>}

    <a id="L107"></a>length := 38 + len(m.sessionId);
    <a id="L108"></a>x := make([]byte, 4+length);
    <a id="L109"></a>x[0] = typeServerHello;
    <a id="L110"></a>x[1] = uint8(length &gt;&gt; 16);
    <a id="L111"></a>x[2] = uint8(length &gt;&gt; 8);
    <a id="L112"></a>x[3] = uint8(length);
    <a id="L113"></a>x[4] = m.major;
    <a id="L114"></a>x[5] = m.minor;
    <a id="L115"></a>bytes.Copy(x[6:38], m.random);
    <a id="L116"></a>x[38] = uint8(len(m.sessionId));
    <a id="L117"></a>bytes.Copy(x[39:39+len(m.sessionId)], m.sessionId);
    <a id="L118"></a>z := x[39+len(m.sessionId) : len(x)];
    <a id="L119"></a>z[0] = uint8(m.cipherSuite &gt;&gt; 8);
    <a id="L120"></a>z[1] = uint8(m.cipherSuite);
    <a id="L121"></a>z[2] = uint8(m.compressionMethod);
    <a id="L122"></a>m.raw = x;

    <a id="L124"></a>return x;
<a id="L125"></a>}

<a id="L127"></a>type certificateMsg struct {
    <a id="L128"></a>raw          []byte;
    <a id="L129"></a>certificates [][]byte;
<a id="L130"></a>}

<a id="L132"></a>func (m *certificateMsg) marshal() (x []byte) {
    <a id="L133"></a>if m.raw != nil {
        <a id="L134"></a>return m.raw
    <a id="L135"></a>}

    <a id="L137"></a>var i int;
    <a id="L138"></a>for _, slice := range m.certificates {
        <a id="L139"></a>i += len(slice)
    <a id="L140"></a>}

    <a id="L142"></a>length := 3 + 3*len(m.certificates) + i;
    <a id="L143"></a>x = make([]byte, 4+length);
    <a id="L144"></a>x[0] = typeCertificate;
    <a id="L145"></a>x[1] = uint8(length &gt;&gt; 16);
    <a id="L146"></a>x[2] = uint8(length &gt;&gt; 8);
    <a id="L147"></a>x[3] = uint8(length);

    <a id="L149"></a>certificateOctets := length - 3;
    <a id="L150"></a>x[4] = uint8(certificateOctets &gt;&gt; 16);
    <a id="L151"></a>x[5] = uint8(certificateOctets &gt;&gt; 8);
    <a id="L152"></a>x[6] = uint8(certificateOctets);

    <a id="L154"></a>y := x[7:len(x)];
    <a id="L155"></a>for _, slice := range m.certificates {
        <a id="L156"></a>y[0] = uint8(len(slice) &gt;&gt; 16);
        <a id="L157"></a>y[1] = uint8(len(slice) &gt;&gt; 8);
        <a id="L158"></a>y[2] = uint8(len(slice));
        <a id="L159"></a>bytes.Copy(y[3:len(y)], slice);
        <a id="L160"></a>y = y[3+len(slice) : len(y)];
    <a id="L161"></a>}

    <a id="L163"></a>m.raw = x;
    <a id="L164"></a>return;
<a id="L165"></a>}

<a id="L167"></a>type serverHelloDoneMsg struct{}

<a id="L169"></a>func (m *serverHelloDoneMsg) marshal() []byte {
    <a id="L170"></a>x := make([]byte, 4);
    <a id="L171"></a>x[0] = typeServerHelloDone;
    <a id="L172"></a>return x;
<a id="L173"></a>}

<a id="L175"></a>type clientKeyExchangeMsg struct {
    <a id="L176"></a>raw        []byte;
    <a id="L177"></a>ciphertext []byte;
<a id="L178"></a>}

<a id="L180"></a>func (m *clientKeyExchangeMsg) marshal() []byte {
    <a id="L181"></a>if m.raw != nil {
        <a id="L182"></a>return m.raw
    <a id="L183"></a>}
    <a id="L184"></a>length := len(m.ciphertext) + 2;
    <a id="L185"></a>x := make([]byte, length+4);
    <a id="L186"></a>x[0] = typeClientKeyExchange;
    <a id="L187"></a>x[1] = uint8(length &gt;&gt; 16);
    <a id="L188"></a>x[2] = uint8(length &gt;&gt; 8);
    <a id="L189"></a>x[3] = uint8(length);
    <a id="L190"></a>x[4] = uint8(len(m.ciphertext) &gt;&gt; 8);
    <a id="L191"></a>x[5] = uint8(len(m.ciphertext));
    <a id="L192"></a>bytes.Copy(x[6:len(x)], m.ciphertext);

    <a id="L194"></a>m.raw = x;
    <a id="L195"></a>return x;
<a id="L196"></a>}

<a id="L198"></a>func (m *clientKeyExchangeMsg) unmarshal(data []byte) bool {
    <a id="L199"></a>m.raw = data;
    <a id="L200"></a>if len(data) &lt; 7 {
        <a id="L201"></a>return false
    <a id="L202"></a>}
    <a id="L203"></a>cipherTextLen := int(data[4])&lt;&lt;8 | int(data[5]);
    <a id="L204"></a>if len(data) != 6+cipherTextLen {
        <a id="L205"></a>return false
    <a id="L206"></a>}
    <a id="L207"></a>m.ciphertext = data[6:len(data)];
    <a id="L208"></a>return true;
<a id="L209"></a>}

<a id="L211"></a>type finishedMsg struct {
    <a id="L212"></a>raw        []byte;
    <a id="L213"></a>verifyData []byte;
<a id="L214"></a>}

<a id="L216"></a>func (m *finishedMsg) marshal() (x []byte) {
    <a id="L217"></a>if m.raw != nil {
        <a id="L218"></a>return m.raw
    <a id="L219"></a>}

    <a id="L221"></a>x = make([]byte, 16);
    <a id="L222"></a>x[0] = typeFinished;
    <a id="L223"></a>x[3] = 12;
    <a id="L224"></a>bytes.Copy(x[4:len(x)], m.verifyData);
    <a id="L225"></a>m.raw = x;
    <a id="L226"></a>return;
<a id="L227"></a>}

<a id="L229"></a>func (m *finishedMsg) unmarshal(data []byte) bool {
    <a id="L230"></a>m.raw = data;
    <a id="L231"></a>if len(data) != 4+12 {
        <a id="L232"></a>return false
    <a id="L233"></a>}
    <a id="L234"></a>m.verifyData = data[4:len(data)];
    <a id="L235"></a>return true;
<a id="L236"></a>}
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
