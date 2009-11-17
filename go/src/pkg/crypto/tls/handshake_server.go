<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/tls/handshake_server.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/tls/handshake_server.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package tls

<a id="L7"></a><span class="comment">// The handshake goroutine reads handshake messages from the record processor</span>
<a id="L8"></a><span class="comment">// and outputs messages to be written on another channel. It updates the record</span>
<a id="L9"></a><span class="comment">// processor with the state of the connection via the control channel. In the</span>
<a id="L10"></a><span class="comment">// case of handshake messages that need synchronous processing (because they</span>
<a id="L11"></a><span class="comment">// affect the handling of the next record) the record processor knows about</span>
<a id="L12"></a><span class="comment">// them and either waits for a control message (Finished) or includes a reply</span>
<a id="L13"></a><span class="comment">// channel in the message (ChangeCipherSpec).</span>

<a id="L15"></a>import (
    <a id="L16"></a>&#34;crypto/hmac&#34;;
    <a id="L17"></a>&#34;crypto/rc4&#34;;
    <a id="L18"></a>&#34;crypto/rsa&#34;;
    <a id="L19"></a>&#34;crypto/sha1&#34;;
    <a id="L20"></a>&#34;crypto/subtle&#34;;
    <a id="L21"></a>&#34;io&#34;;
<a id="L22"></a>)

<a id="L24"></a>type cipherSuite struct {
    <a id="L25"></a>id                          uint16; <span class="comment">// The number of this suite on the wire.</span>
    <a id="L26"></a>hashLength, cipherKeyLength int;
    <a id="L27"></a><span class="comment">// TODO(agl): need a method to create the cipher and hash interfaces.</span>
<a id="L28"></a>}

<a id="L30"></a>var cipherSuites = []cipherSuite{
    <a id="L31"></a>cipherSuite{TLS_RSA_WITH_RC4_128_SHA, 20, 16},
<a id="L32"></a>}

<a id="L34"></a><span class="comment">// A serverHandshake performs the server side of the TLS 1.1 handshake protocol.</span>
<a id="L35"></a>type serverHandshake struct {
    <a id="L36"></a>writeChan   chan&lt;- interface{};
    <a id="L37"></a>controlChan chan&lt;- interface{};
    <a id="L38"></a>msgChan     &lt;-chan interface{};
    <a id="L39"></a>config      *Config;
<a id="L40"></a>}

<a id="L42"></a>func (h *serverHandshake) loop(writeChan chan&lt;- interface{}, controlChan chan&lt;- interface{}, msgChan &lt;-chan interface{}, config *Config) {
    <a id="L43"></a>h.writeChan = writeChan;
    <a id="L44"></a>h.controlChan = controlChan;
    <a id="L45"></a>h.msgChan = msgChan;
    <a id="L46"></a>h.config = config;

    <a id="L48"></a>defer close(writeChan);
    <a id="L49"></a>defer close(controlChan);

    <a id="L51"></a>clientHello, ok := h.readHandshakeMsg().(*clientHelloMsg);
    <a id="L52"></a>if !ok {
        <a id="L53"></a>h.error(alertUnexpectedMessage);
        <a id="L54"></a>return;
    <a id="L55"></a>}
    <a id="L56"></a>major, minor, ok := mutualVersion(clientHello.major, clientHello.minor);
    <a id="L57"></a>if !ok {
        <a id="L58"></a>h.error(alertProtocolVersion);
        <a id="L59"></a>return;
    <a id="L60"></a>}

    <a id="L62"></a>finishedHash := newFinishedHash();
    <a id="L63"></a>finishedHash.Write(clientHello.marshal());

    <a id="L65"></a>hello := new(serverHelloMsg);

    <a id="L67"></a><span class="comment">// We only support a single ciphersuite so we look for it in the list</span>
    <a id="L68"></a><span class="comment">// of client supported suites.</span>
    <a id="L69"></a><span class="comment">//</span>
    <a id="L70"></a><span class="comment">// TODO(agl): Add additional cipher suites.</span>
    <a id="L71"></a>var suite *cipherSuite;

    <a id="L73"></a>for _, id := range clientHello.cipherSuites {
        <a id="L74"></a>for _, supported := range cipherSuites {
            <a id="L75"></a>if supported.id == id {
                <a id="L76"></a>suite = &amp;supported;
                <a id="L77"></a>break;
            <a id="L78"></a>}
        <a id="L79"></a>}
    <a id="L80"></a>}

    <a id="L82"></a>foundCompression := false;
    <a id="L83"></a><span class="comment">// We only support null compression, so check that the client offered it.</span>
    <a id="L84"></a>for _, compression := range clientHello.compressionMethods {
        <a id="L85"></a>if compression == compressionNone {
            <a id="L86"></a>foundCompression = true;
            <a id="L87"></a>break;
        <a id="L88"></a>}
    <a id="L89"></a>}

    <a id="L91"></a>if suite == nil || !foundCompression {
        <a id="L92"></a>h.error(alertHandshakeFailure);
        <a id="L93"></a>return;
    <a id="L94"></a>}

    <a id="L96"></a>hello.major = major;
    <a id="L97"></a>hello.minor = minor;
    <a id="L98"></a>hello.cipherSuite = suite.id;
    <a id="L99"></a>currentTime := uint32(config.Time());
    <a id="L100"></a>hello.random = make([]byte, 32);
    <a id="L101"></a>hello.random[0] = byte(currentTime &gt;&gt; 24);
    <a id="L102"></a>hello.random[1] = byte(currentTime &gt;&gt; 16);
    <a id="L103"></a>hello.random[2] = byte(currentTime &gt;&gt; 8);
    <a id="L104"></a>hello.random[3] = byte(currentTime);
    <a id="L105"></a>_, err := io.ReadFull(config.Rand, hello.random[4:len(hello.random)]);
    <a id="L106"></a>if err != nil {
        <a id="L107"></a>h.error(alertInternalError);
        <a id="L108"></a>return;
    <a id="L109"></a>}
    <a id="L110"></a>hello.compressionMethod = compressionNone;

    <a id="L112"></a>finishedHash.Write(hello.marshal());
    <a id="L113"></a>writeChan &lt;- writerSetVersion{major, minor};
    <a id="L114"></a>writeChan &lt;- hello;

    <a id="L116"></a>if len(config.Certificates) == 0 {
        <a id="L117"></a>h.error(alertInternalError);
        <a id="L118"></a>return;
    <a id="L119"></a>}

    <a id="L121"></a>certMsg := new(certificateMsg);
    <a id="L122"></a>certMsg.certificates = config.Certificates[0].Certificate;
    <a id="L123"></a>finishedHash.Write(certMsg.marshal());
    <a id="L124"></a>writeChan &lt;- certMsg;

    <a id="L126"></a>helloDone := new(serverHelloDoneMsg);
    <a id="L127"></a>finishedHash.Write(helloDone.marshal());
    <a id="L128"></a>writeChan &lt;- helloDone;

    <a id="L130"></a>ckx, ok := h.readHandshakeMsg().(*clientKeyExchangeMsg);
    <a id="L131"></a>if !ok {
        <a id="L132"></a>h.error(alertUnexpectedMessage);
        <a id="L133"></a>return;
    <a id="L134"></a>}
    <a id="L135"></a>finishedHash.Write(ckx.marshal());

    <a id="L137"></a>preMasterSecret := make([]byte, 48);
    <a id="L138"></a>_, err = io.ReadFull(config.Rand, preMasterSecret[2:len(preMasterSecret)]);
    <a id="L139"></a>if err != nil {
        <a id="L140"></a>h.error(alertInternalError);
        <a id="L141"></a>return;
    <a id="L142"></a>}

    <a id="L144"></a>err = rsa.DecryptPKCS1v15SessionKey(config.Rand, config.Certificates[0].PrivateKey, ckx.ciphertext, preMasterSecret);
    <a id="L145"></a>if err != nil {
        <a id="L146"></a>h.error(alertHandshakeFailure);
        <a id="L147"></a>return;
    <a id="L148"></a>}
    <a id="L149"></a><span class="comment">// We don&#39;t check the version number in the premaster secret. For one,</span>
    <a id="L150"></a><span class="comment">// by checking it, we would leak information about the validity of the</span>
    <a id="L151"></a><span class="comment">// encrypted pre-master secret. Secondly, it provides only a small</span>
    <a id="L152"></a><span class="comment">// benefit against a downgrade attack and some implementations send the</span>
    <a id="L153"></a><span class="comment">// wrong version anyway. See the discussion at the end of section</span>
    <a id="L154"></a><span class="comment">// 7.4.7.1 of RFC 4346.</span>

    <a id="L156"></a>masterSecret, clientMAC, serverMAC, clientKey, serverKey :=
        <a id="L157"></a>keysFromPreMasterSecret11(preMasterSecret, clientHello.random, hello.random, suite.hashLength, suite.cipherKeyLength);

    <a id="L159"></a>_, ok = h.readHandshakeMsg().(changeCipherSpec);
    <a id="L160"></a>if !ok {
        <a id="L161"></a>h.error(alertUnexpectedMessage);
        <a id="L162"></a>return;
    <a id="L163"></a>}

    <a id="L165"></a>cipher, _ := rc4.NewCipher(clientKey);
    <a id="L166"></a>controlChan &lt;- &amp;newCipherSpec{cipher, hmac.New(sha1.New(), clientMAC)};

    <a id="L168"></a>clientFinished, ok := h.readHandshakeMsg().(*finishedMsg);
    <a id="L169"></a>if !ok {
        <a id="L170"></a>h.error(alertUnexpectedMessage);
        <a id="L171"></a>return;
    <a id="L172"></a>}

    <a id="L174"></a>verify := finishedHash.clientSum(masterSecret);
    <a id="L175"></a>if len(verify) != len(clientFinished.verifyData) ||
        <a id="L176"></a>subtle.ConstantTimeCompare(verify, clientFinished.verifyData) != 1 {
        <a id="L177"></a>h.error(alertHandshakeFailure);
        <a id="L178"></a>return;
    <a id="L179"></a>}

    <a id="L181"></a>controlChan &lt;- ConnectionState{true, &#34;TLS_RSA_WITH_RC4_128_SHA&#34;, 0};

    <a id="L183"></a>finishedHash.Write(clientFinished.marshal());

    <a id="L185"></a>cipher2, _ := rc4.NewCipher(serverKey);
    <a id="L186"></a>writeChan &lt;- writerChangeCipherSpec{cipher2, hmac.New(sha1.New(), serverMAC)};

    <a id="L188"></a>finished := new(finishedMsg);
    <a id="L189"></a>finished.verifyData = finishedHash.serverSum(masterSecret);
    <a id="L190"></a>writeChan &lt;- finished;

    <a id="L192"></a>writeChan &lt;- writerEnableApplicationData{};

    <a id="L194"></a>for {
        <a id="L195"></a>_, ok := h.readHandshakeMsg().(*clientHelloMsg);
        <a id="L196"></a>if !ok {
            <a id="L197"></a>h.error(alertUnexpectedMessage);
            <a id="L198"></a>return;
        <a id="L199"></a>}
        <a id="L200"></a><span class="comment">// We reject all renegotication requests.</span>
        <a id="L201"></a>writeChan &lt;- alert{alertLevelWarning, alertNoRenegotiation};
    <a id="L202"></a>}
<a id="L203"></a>}

<a id="L205"></a>func (h *serverHandshake) readHandshakeMsg() interface{} {
    <a id="L206"></a>v := &lt;-h.msgChan;
    <a id="L207"></a>if closed(h.msgChan) {
        <a id="L208"></a><span class="comment">// If the channel closed then the processor received an error</span>
        <a id="L209"></a><span class="comment">// from the peer and we don&#39;t want to echo it back to them.</span>
        <a id="L210"></a>h.msgChan = nil;
        <a id="L211"></a>return 0;
    <a id="L212"></a>}
    <a id="L213"></a>if _, ok := v.(alert); ok {
        <a id="L214"></a><span class="comment">// We got an alert from the processor. We forward to the writer</span>
        <a id="L215"></a><span class="comment">// and shutdown.</span>
        <a id="L216"></a>h.writeChan &lt;- v;
        <a id="L217"></a>h.msgChan = nil;
        <a id="L218"></a>return 0;
    <a id="L219"></a>}
    <a id="L220"></a>return v;
<a id="L221"></a>}

<a id="L223"></a>func (h *serverHandshake) error(e alertType) {
    <a id="L224"></a>if h.msgChan != nil {
        <a id="L225"></a><span class="comment">// If we didn&#39;t get an error from the processor, then we need</span>
        <a id="L226"></a><span class="comment">// to tell it about the error.</span>
        <a id="L227"></a>h.controlChan &lt;- ConnectionState{false, &#34;&#34;, e};
        <a id="L228"></a>close(h.controlChan);
        <a id="L229"></a>go func() {
            <a id="L230"></a>for _ = range h.msgChan {
            <a id="L231"></a>}
        <a id="L232"></a>}();
        <a id="L233"></a>h.writeChan &lt;- alert{alertLevelError, e};
    <a id="L234"></a>}
<a id="L235"></a>}
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
