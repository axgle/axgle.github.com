<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/tls/record_process.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/tls/record_process.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package tls

<a id="L7"></a><span class="comment">// A recordProcessor accepts reassembled records, decrypts and verifies them</span>
<a id="L8"></a><span class="comment">// and routes them either to the handshake processor, to up to the application.</span>
<a id="L9"></a><span class="comment">// It also accepts requests from the application for the current connection</span>
<a id="L10"></a><span class="comment">// state, or for a notification when the state changes.</span>

<a id="L12"></a>import (
    <a id="L13"></a>&#34;bytes&#34;;
    <a id="L14"></a>&#34;container/list&#34;;
    <a id="L15"></a>&#34;crypto/subtle&#34;;
    <a id="L16"></a>&#34;hash&#34;;
<a id="L17"></a>)

<a id="L19"></a><span class="comment">// getConnectionState is a request from the application to get the current</span>
<a id="L20"></a><span class="comment">// ConnectionState.</span>
<a id="L21"></a>type getConnectionState struct {
    <a id="L22"></a>reply chan&lt;- ConnectionState;
<a id="L23"></a>}

<a id="L25"></a><span class="comment">// waitConnectionState is a request from the application to be notified when</span>
<a id="L26"></a><span class="comment">// the connection state changes.</span>
<a id="L27"></a>type waitConnectionState struct {
    <a id="L28"></a>reply chan&lt;- ConnectionState;
<a id="L29"></a>}

<a id="L31"></a><span class="comment">// connectionStateChange is a message from the handshake processor that the</span>
<a id="L32"></a><span class="comment">// connection state has changed.</span>
<a id="L33"></a>type connectionStateChange struct {
    <a id="L34"></a>connState ConnectionState;
<a id="L35"></a>}

<a id="L37"></a><span class="comment">// changeCipherSpec is a message send to the handshake processor to signal that</span>
<a id="L38"></a><span class="comment">// the peer is switching ciphers.</span>
<a id="L39"></a>type changeCipherSpec struct{}

<a id="L41"></a><span class="comment">// newCipherSpec is a message from the handshake processor that future</span>
<a id="L42"></a><span class="comment">// records should be processed with a new cipher and MAC function.</span>
<a id="L43"></a>type newCipherSpec struct {
    <a id="L44"></a>encrypt encryptor;
    <a id="L45"></a>mac     hash.Hash;
<a id="L46"></a>}

<a id="L48"></a>type recordProcessor struct {
    <a id="L49"></a>decrypt       encryptor;
    <a id="L50"></a>mac           hash.Hash;
    <a id="L51"></a>seqNum        uint64;
    <a id="L52"></a>handshakeBuf  []byte;
    <a id="L53"></a>appDataChan   chan&lt;- []byte;
    <a id="L54"></a>requestChan   &lt;-chan interface{};
    <a id="L55"></a>controlChan   &lt;-chan interface{};
    <a id="L56"></a>recordChan    &lt;-chan *record;
    <a id="L57"></a>handshakeChan chan&lt;- interface{};

    <a id="L59"></a><span class="comment">// recordRead is nil when we don&#39;t wish to read any more.</span>
    <a id="L60"></a>recordRead &lt;-chan *record;
    <a id="L61"></a><span class="comment">// appDataSend is nil when len(appData) == 0.</span>
    <a id="L62"></a>appDataSend chan&lt;- []byte;
    <a id="L63"></a><span class="comment">// appData contains any application data queued for upstream.</span>
    <a id="L64"></a>appData []byte;
    <a id="L65"></a><span class="comment">// A list of channels waiting for connState to change.</span>
    <a id="L66"></a>waitQueue *list.List;
    <a id="L67"></a>connState ConnectionState;
    <a id="L68"></a>shutdown  bool;
    <a id="L69"></a>header    [13]byte;
<a id="L70"></a>}

<a id="L72"></a><span class="comment">// drainRequestChannel processes messages from the request channel until it&#39;s closed.</span>
<a id="L73"></a>func drainRequestChannel(requestChan &lt;-chan interface{}, c ConnectionState) {
    <a id="L74"></a>for v := range requestChan {
        <a id="L75"></a>if closed(requestChan) {
            <a id="L76"></a>return
        <a id="L77"></a>}
        <a id="L78"></a>switch r := v.(type) {
        <a id="L79"></a>case getConnectionState:
            <a id="L80"></a>r.reply &lt;- c
        <a id="L81"></a>case waitConnectionState:
            <a id="L82"></a>r.reply &lt;- c
        <a id="L83"></a>}
    <a id="L84"></a>}
<a id="L85"></a>}

<a id="L87"></a>func (p *recordProcessor) loop(appDataChan chan&lt;- []byte, requestChan &lt;-chan interface{}, controlChan &lt;-chan interface{}, recordChan &lt;-chan *record, handshakeChan chan&lt;- interface{}) {
    <a id="L88"></a>noop := nop{};
    <a id="L89"></a>p.decrypt = noop;
    <a id="L90"></a>p.mac = noop;
    <a id="L91"></a>p.waitQueue = list.New();

    <a id="L93"></a>p.appDataChan = appDataChan;
    <a id="L94"></a>p.requestChan = requestChan;
    <a id="L95"></a>p.controlChan = controlChan;
    <a id="L96"></a>p.recordChan = recordChan;
    <a id="L97"></a>p.handshakeChan = handshakeChan;
    <a id="L98"></a>p.recordRead = recordChan;

    <a id="L100"></a>for !p.shutdown {
        <a id="L101"></a>select {
        <a id="L102"></a>case p.appDataSend &lt;- p.appData:
            <a id="L103"></a>p.appData = nil;
            <a id="L104"></a>p.appDataSend = nil;
            <a id="L105"></a>p.recordRead = p.recordChan;
        <a id="L106"></a>case c := &lt;-controlChan:
            <a id="L107"></a>p.processControlMsg(c)
        <a id="L108"></a>case r := &lt;-requestChan:
            <a id="L109"></a>p.processRequestMsg(r)
        <a id="L110"></a>case r := &lt;-p.recordRead:
            <a id="L111"></a>p.processRecord(r)
        <a id="L112"></a>}
    <a id="L113"></a>}

    <a id="L115"></a>p.wakeWaiters();
    <a id="L116"></a>go drainRequestChannel(p.requestChan, p.connState);
    <a id="L117"></a>go func() {
        <a id="L118"></a>for _ = range controlChan {
        <a id="L119"></a>}
    <a id="L120"></a>}();

    <a id="L122"></a>close(handshakeChan);
    <a id="L123"></a>if len(p.appData) &gt; 0 {
        <a id="L124"></a>appDataChan &lt;- p.appData
    <a id="L125"></a>}
    <a id="L126"></a>close(appDataChan);
<a id="L127"></a>}

<a id="L129"></a>func (p *recordProcessor) processRequestMsg(requestMsg interface{}) {
    <a id="L130"></a>if closed(p.requestChan) {
        <a id="L131"></a>p.shutdown = true;
        <a id="L132"></a>return;
    <a id="L133"></a>}

    <a id="L135"></a>switch r := requestMsg.(type) {
    <a id="L136"></a>case getConnectionState:
        <a id="L137"></a>r.reply &lt;- p.connState
    <a id="L138"></a>case waitConnectionState:
        <a id="L139"></a>if p.connState.HandshakeComplete {
            <a id="L140"></a>r.reply &lt;- p.connState
        <a id="L141"></a>}
        <a id="L142"></a>p.waitQueue.PushBack(r.reply);
    <a id="L143"></a>}
<a id="L144"></a>}

<a id="L146"></a>func (p *recordProcessor) processControlMsg(msg interface{}) {
    <a id="L147"></a>connState, ok := msg.(ConnectionState);
    <a id="L148"></a>if !ok || closed(p.controlChan) {
        <a id="L149"></a>p.shutdown = true;
        <a id="L150"></a>return;
    <a id="L151"></a>}

    <a id="L153"></a>p.connState = connState;
    <a id="L154"></a>p.wakeWaiters();
<a id="L155"></a>}

<a id="L157"></a>func (p *recordProcessor) wakeWaiters() {
    <a id="L158"></a>for i := p.waitQueue.Front(); i != nil; i = i.Next() {
        <a id="L159"></a>i.Value.(chan&lt;- ConnectionState) &lt;- p.connState
    <a id="L160"></a>}
    <a id="L161"></a>p.waitQueue.Init();
<a id="L162"></a>}

<a id="L164"></a>func (p *recordProcessor) processRecord(r *record) {
    <a id="L165"></a>if closed(p.recordChan) {
        <a id="L166"></a>p.shutdown = true;
        <a id="L167"></a>return;
    <a id="L168"></a>}

    <a id="L170"></a>p.decrypt.XORKeyStream(r.payload);
    <a id="L171"></a>if len(r.payload) &lt; p.mac.Size() {
        <a id="L172"></a>p.error(alertBadRecordMAC);
        <a id="L173"></a>return;
    <a id="L174"></a>}

    <a id="L176"></a>fillMACHeader(&amp;p.header, p.seqNum, len(r.payload)-p.mac.Size(), r);
    <a id="L177"></a>p.seqNum++;

    <a id="L179"></a>p.mac.Reset();
    <a id="L180"></a>p.mac.Write(p.header[0:13]);
    <a id="L181"></a>p.mac.Write(r.payload[0 : len(r.payload)-p.mac.Size()]);
    <a id="L182"></a>macBytes := p.mac.Sum();

    <a id="L184"></a>if subtle.ConstantTimeCompare(macBytes, r.payload[len(r.payload)-p.mac.Size():len(r.payload)]) != 1 {
        <a id="L185"></a>p.error(alertBadRecordMAC);
        <a id="L186"></a>return;
    <a id="L187"></a>}

    <a id="L189"></a>switch r.contentType {
    <a id="L190"></a>case recordTypeHandshake:
        <a id="L191"></a>p.processHandshakeRecord(r.payload[0 : len(r.payload)-p.mac.Size()])
    <a id="L192"></a>case recordTypeChangeCipherSpec:
        <a id="L193"></a>if len(r.payload) != 1 || r.payload[0] != 1 {
            <a id="L194"></a>p.error(alertUnexpectedMessage);
            <a id="L195"></a>return;
        <a id="L196"></a>}

        <a id="L198"></a>p.handshakeChan &lt;- changeCipherSpec{};
        <a id="L199"></a>newSpec, ok := (&lt;-p.controlChan).(*newCipherSpec);
        <a id="L200"></a>if !ok {
            <a id="L201"></a>p.connState.Error = alertUnexpectedMessage;
            <a id="L202"></a>p.shutdown = true;
            <a id="L203"></a>return;
        <a id="L204"></a>}
        <a id="L205"></a>p.decrypt = newSpec.encrypt;
        <a id="L206"></a>p.mac = newSpec.mac;
        <a id="L207"></a>p.seqNum = 0;
    <a id="L208"></a>case recordTypeApplicationData:
        <a id="L209"></a>if p.connState.HandshakeComplete == false {
            <a id="L210"></a>p.error(alertUnexpectedMessage);
            <a id="L211"></a>return;
        <a id="L212"></a>}
        <a id="L213"></a>p.recordRead = nil;
        <a id="L214"></a>p.appData = r.payload;
        <a id="L215"></a>p.appDataSend = p.appDataChan;
    <a id="L216"></a>default:
        <a id="L217"></a>p.error(alertUnexpectedMessage);
        <a id="L218"></a>return;
    <a id="L219"></a>}
<a id="L220"></a>}

<a id="L222"></a>func (p *recordProcessor) processHandshakeRecord(data []byte) {
    <a id="L223"></a>if p.handshakeBuf == nil {
        <a id="L224"></a>p.handshakeBuf = data
    <a id="L225"></a>} else {
        <a id="L226"></a>if len(p.handshakeBuf) &gt; maxHandshakeMsg {
            <a id="L227"></a>p.error(alertInternalError);
            <a id="L228"></a>return;
        <a id="L229"></a>}
        <a id="L230"></a>newBuf := make([]byte, len(p.handshakeBuf)+len(data));
        <a id="L231"></a>bytes.Copy(newBuf, p.handshakeBuf);
        <a id="L232"></a>bytes.Copy(newBuf[len(p.handshakeBuf):len(newBuf)], data);
        <a id="L233"></a>p.handshakeBuf = newBuf;
    <a id="L234"></a>}

    <a id="L236"></a>for len(p.handshakeBuf) &gt;= 4 {
        <a id="L237"></a>handshakeLen := int(p.handshakeBuf[1])&lt;&lt;16 |
            <a id="L238"></a>int(p.handshakeBuf[2])&lt;&lt;8 |
            <a id="L239"></a>int(p.handshakeBuf[3]);
        <a id="L240"></a>if handshakeLen+4 &gt; len(p.handshakeBuf) {
            <a id="L241"></a>break
        <a id="L242"></a>}

        <a id="L244"></a>bytes := p.handshakeBuf[0 : handshakeLen+4];
        <a id="L245"></a>p.handshakeBuf = p.handshakeBuf[handshakeLen+4 : len(p.handshakeBuf)];
        <a id="L246"></a>if bytes[0] == typeFinished {
            <a id="L247"></a><span class="comment">// Special case because Finished is synchronous: the</span>
            <a id="L248"></a><span class="comment">// handshake handler has to tell us if it&#39;s ok to start</span>
            <a id="L249"></a><span class="comment">// forwarding application data.</span>
            <a id="L250"></a>m := new(finishedMsg);
            <a id="L251"></a>if !m.unmarshal(bytes) {
                <a id="L252"></a>p.error(alertUnexpectedMessage)
            <a id="L253"></a>}
            <a id="L254"></a>p.handshakeChan &lt;- m;
            <a id="L255"></a>var ok bool;
            <a id="L256"></a>p.connState, ok = (&lt;-p.controlChan).(ConnectionState);
            <a id="L257"></a>if !ok || p.connState.Error != 0 {
                <a id="L258"></a>p.shutdown = true;
                <a id="L259"></a>return;
            <a id="L260"></a>}
        <a id="L261"></a>} else {
            <a id="L262"></a>msg, ok := parseHandshakeMsg(bytes);
            <a id="L263"></a>if !ok {
                <a id="L264"></a>p.error(alertUnexpectedMessage);
                <a id="L265"></a>return;
            <a id="L266"></a>}
            <a id="L267"></a>p.handshakeChan &lt;- msg;
        <a id="L268"></a>}
    <a id="L269"></a>}
<a id="L270"></a>}

<a id="L272"></a>func (p *recordProcessor) error(err alertType) {
    <a id="L273"></a>close(p.handshakeChan);
    <a id="L274"></a>p.connState.Error = err;
    <a id="L275"></a>p.wakeWaiters();
    <a id="L276"></a>p.shutdown = true;
<a id="L277"></a>}

<a id="L279"></a>func parseHandshakeMsg(data []byte) (interface{}, bool) {
    <a id="L280"></a>var m interface {
        <a id="L281"></a>unmarshal([]byte) bool;
    <a id="L282"></a>}

    <a id="L284"></a>switch data[0] {
    <a id="L285"></a>case typeClientHello:
        <a id="L286"></a>m = new(clientHelloMsg)
    <a id="L287"></a>case typeClientKeyExchange:
        <a id="L288"></a>m = new(clientKeyExchangeMsg)
    <a id="L289"></a>default:
        <a id="L290"></a>return nil, false
    <a id="L291"></a>}

    <a id="L293"></a>ok := m.unmarshal(data);
    <a id="L294"></a>return m, ok;
<a id="L295"></a>}
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
