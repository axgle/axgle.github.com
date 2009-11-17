<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/tls/handshake_server_test.go</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/tls/handshake_server_test.go</h1>

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
    <a id="L9"></a>&#34;big&#34;;
    <a id="L10"></a>&#34;crypto/rsa&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;testing&#34;;
    <a id="L13"></a>&#34;testing/script&#34;;
<a id="L14"></a>)

<a id="L16"></a>type zeroSource struct{}

<a id="L18"></a>func (zeroSource) Read(b []byte) (n int, err os.Error) {
    <a id="L19"></a>for i := range b {
        <a id="L20"></a>b[i] = 0
    <a id="L21"></a>}

    <a id="L23"></a>return len(b), nil;
<a id="L24"></a>}

<a id="L26"></a>var testConfig *Config

<a id="L28"></a>func init() {
    <a id="L29"></a>testConfig = new(Config);
    <a id="L30"></a>testConfig.Time = func() int64 { return 0 };
    <a id="L31"></a>testConfig.Rand = zeroSource{};
    <a id="L32"></a>testConfig.Certificates = make([]Certificate, 1);
    <a id="L33"></a>testConfig.Certificates[0].Certificate = [][]byte{testCertificate};
    <a id="L34"></a>testConfig.Certificates[0].PrivateKey = testPrivateKey;
<a id="L35"></a>}

<a id="L37"></a>func setupServerHandshake() (writeChan chan interface{}, controlChan chan interface{}, msgChan chan interface{}) {
    <a id="L38"></a>sh := new(serverHandshake);
    <a id="L39"></a>writeChan = make(chan interface{});
    <a id="L40"></a>controlChan = make(chan interface{});
    <a id="L41"></a>msgChan = make(chan interface{});

    <a id="L43"></a>go sh.loop(writeChan, controlChan, msgChan, testConfig);
    <a id="L44"></a>return;
<a id="L45"></a>}

<a id="L47"></a>func testClientHelloFailure(t *testing.T, clientHello interface{}, expectedAlert alertType) {
    <a id="L48"></a>writeChan, controlChan, msgChan := setupServerHandshake();
    <a id="L49"></a>defer close(msgChan);

    <a id="L51"></a>send := script.NewEvent(&#34;send&#34;, nil, script.Send{msgChan, clientHello});
    <a id="L52"></a>recvAlert := script.NewEvent(&#34;recv alert&#34;, []*script.Event{send}, script.Recv{writeChan, alert{alertLevelError, expectedAlert}});
    <a id="L53"></a>close1 := script.NewEvent(&#34;msgChan close&#34;, []*script.Event{recvAlert}, script.Closed{writeChan});
    <a id="L54"></a>recvState := script.NewEvent(&#34;recv state&#34;, []*script.Event{send}, script.Recv{controlChan, ConnectionState{false, &#34;&#34;, expectedAlert}});
    <a id="L55"></a>close2 := script.NewEvent(&#34;controlChan close&#34;, []*script.Event{recvState}, script.Closed{controlChan});

    <a id="L57"></a>err := script.Perform(0, []*script.Event{send, recvAlert, close1, recvState, close2});
    <a id="L58"></a>if err != nil {
        <a id="L59"></a>t.Errorf(&#34;Got error: %s&#34;, err)
    <a id="L60"></a>}
<a id="L61"></a>}

<a id="L63"></a>func TestSimpleError(t *testing.T) {
    <a id="L64"></a>testClientHelloFailure(t, &amp;serverHelloDoneMsg{}, alertUnexpectedMessage)
<a id="L65"></a>}

<a id="L67"></a>var badProtocolVersions = []uint8{0, 0, 0, 5, 1, 0, 1, 5, 2, 0, 2, 5, 3, 0}

<a id="L69"></a>func TestRejectBadProtocolVersion(t *testing.T) {
    <a id="L70"></a>clientHello := new(clientHelloMsg);

    <a id="L72"></a>for i := 0; i &lt; len(badProtocolVersions); i += 2 {
        <a id="L73"></a>clientHello.major = badProtocolVersions[i];
        <a id="L74"></a>clientHello.minor = badProtocolVersions[i+1];

        <a id="L76"></a>testClientHelloFailure(t, clientHello, alertProtocolVersion);
    <a id="L77"></a>}
<a id="L78"></a>}

<a id="L80"></a>func TestNoSuiteOverlap(t *testing.T) {
    <a id="L81"></a>clientHello := &amp;clientHelloMsg{nil, 3, 1, nil, nil, []uint16{0xff00}, []uint8{0}};
    <a id="L82"></a>testClientHelloFailure(t, clientHello, alertHandshakeFailure);

<a id="L84"></a>}

<a id="L86"></a>func TestNoCompressionOverlap(t *testing.T) {
    <a id="L87"></a>clientHello := &amp;clientHelloMsg{nil, 3, 1, nil, nil, []uint16{TLS_RSA_WITH_RC4_128_SHA}, []uint8{0xff}};
    <a id="L88"></a>testClientHelloFailure(t, clientHello, alertHandshakeFailure);
<a id="L89"></a>}

<a id="L91"></a>func matchServerHello(v interface{}) bool {
    <a id="L92"></a>serverHello, ok := v.(*serverHelloMsg);
    <a id="L93"></a>if !ok {
        <a id="L94"></a>return false
    <a id="L95"></a>}
    <a id="L96"></a>return serverHello.major == 3 &amp;&amp;
        <a id="L97"></a>serverHello.minor == 2 &amp;&amp;
        <a id="L98"></a>serverHello.cipherSuite == TLS_RSA_WITH_RC4_128_SHA &amp;&amp;
        <a id="L99"></a>serverHello.compressionMethod == compressionNone;
<a id="L100"></a>}

<a id="L102"></a>func TestAlertForwarding(t *testing.T) {
    <a id="L103"></a>writeChan, controlChan, msgChan := setupServerHandshake();
    <a id="L104"></a>defer close(msgChan);

    <a id="L106"></a>a := alert{alertLevelError, alertNoRenegotiation};
    <a id="L107"></a>sendAlert := script.NewEvent(&#34;send alert&#34;, nil, script.Send{msgChan, a});
    <a id="L108"></a>recvAlert := script.NewEvent(&#34;recv alert&#34;, []*script.Event{sendAlert}, script.Recv{writeChan, a});
    <a id="L109"></a>closeWriter := script.NewEvent(&#34;close writer&#34;, []*script.Event{recvAlert}, script.Closed{writeChan});
    <a id="L110"></a>closeControl := script.NewEvent(&#34;close control&#34;, []*script.Event{recvAlert}, script.Closed{controlChan});

    <a id="L112"></a>err := script.Perform(0, []*script.Event{sendAlert, recvAlert, closeWriter, closeControl});
    <a id="L113"></a>if err != nil {
        <a id="L114"></a>t.Errorf(&#34;Got error: %s&#34;, err)
    <a id="L115"></a>}
<a id="L116"></a>}

<a id="L118"></a>func TestClose(t *testing.T) {
    <a id="L119"></a>writeChan, controlChan, msgChan := setupServerHandshake();

    <a id="L121"></a>close := script.NewEvent(&#34;close&#34;, nil, script.Close{msgChan});
    <a id="L122"></a>closed1 := script.NewEvent(&#34;closed1&#34;, []*script.Event{close}, script.Closed{writeChan});
    <a id="L123"></a>closed2 := script.NewEvent(&#34;closed2&#34;, []*script.Event{close}, script.Closed{controlChan});

    <a id="L125"></a>err := script.Perform(0, []*script.Event{close, closed1, closed2});
    <a id="L126"></a>if err != nil {
        <a id="L127"></a>t.Errorf(&#34;Got error: %s&#34;, err)
    <a id="L128"></a>}
<a id="L129"></a>}

<a id="L131"></a>func matchCertificate(v interface{}) bool {
    <a id="L132"></a>cert, ok := v.(*certificateMsg);
    <a id="L133"></a>if !ok {
        <a id="L134"></a>return false
    <a id="L135"></a>}
    <a id="L136"></a>return len(cert.certificates) == 1 &amp;&amp;
        <a id="L137"></a>bytes.Compare(cert.certificates[0], testCertificate) == 0;
<a id="L138"></a>}

<a id="L140"></a>func matchSetCipher(v interface{}) bool {
    <a id="L141"></a>_, ok := v.(writerChangeCipherSpec);
    <a id="L142"></a>return ok;
<a id="L143"></a>}

<a id="L145"></a>func matchDone(v interface{}) bool {
    <a id="L146"></a>_, ok := v.(*serverHelloDoneMsg);
    <a id="L147"></a>return ok;
<a id="L148"></a>}

<a id="L150"></a>func matchFinished(v interface{}) bool {
    <a id="L151"></a>finished, ok := v.(*finishedMsg);
    <a id="L152"></a>if !ok {
        <a id="L153"></a>return false
    <a id="L154"></a>}
    <a id="L155"></a>return bytes.Compare(finished.verifyData, fromHex(&#34;29122ae11453e631487b02ed&#34;)) == 0;
<a id="L156"></a>}

<a id="L158"></a>func matchNewCipherSpec(v interface{}) bool {
    <a id="L159"></a>_, ok := v.(*newCipherSpec);
    <a id="L160"></a>return ok;
<a id="L161"></a>}

<a id="L163"></a>func TestFullHandshake(t *testing.T) {
    <a id="L164"></a>writeChan, controlChan, msgChan := setupServerHandshake();
    <a id="L165"></a>defer close(msgChan);

    <a id="L167"></a><span class="comment">// The values for this test were obtained from running `gnutls-cli --insecure --debug 9`</span>
    <a id="L168"></a>clientHello := &amp;clientHelloMsg{fromHex(&#34;0100007603024aef7d77e4686d5dfd9d953dfe280788759ffd440867d687670216da45516b310000340033004500390088001600320044003800870013006600900091008f008e002f004100350084000a00050004008c008d008b008a01000019000900030200010000000e000c0000093132372e302e302e31&#34;), 3, 2, fromHex(&#34;4aef7d77e4686d5dfd9d953dfe280788759ffd440867d687670216da45516b31&#34;), nil, []uint16{0x33, 0x45, 0x39, 0x88, 0x16, 0x32, 0x44, 0x38, 0x87, 0x13, 0x66, 0x90, 0x91, 0x8f, 0x8e, 0x2f, 0x41, 0x35, 0x84, 0xa, 0x5, 0x4, 0x8c, 0x8d, 0x8b, 0x8a}, []uint8{0x0}};

    <a id="L170"></a>sendHello := script.NewEvent(&#34;send hello&#34;, nil, script.Send{msgChan, clientHello});
    <a id="L171"></a>setVersion := script.NewEvent(&#34;set version&#34;, []*script.Event{sendHello}, script.Recv{writeChan, writerSetVersion{3, 2}});
    <a id="L172"></a>recvHello := script.NewEvent(&#34;recv hello&#34;, []*script.Event{setVersion}, script.RecvMatch{writeChan, matchServerHello});
    <a id="L173"></a>recvCert := script.NewEvent(&#34;recv cert&#34;, []*script.Event{recvHello}, script.RecvMatch{writeChan, matchCertificate});
    <a id="L174"></a>recvDone := script.NewEvent(&#34;recv done&#34;, []*script.Event{recvCert}, script.RecvMatch{writeChan, matchDone});

    <a id="L176"></a>ckx := &amp;clientKeyExchangeMsg{nil, fromHex(&#34;872e1fee5f37dd86f3215938ac8de20b302b90074e9fb93097e6b7d1286d0f45abf2daf179deb618bb3c70ed0afee6ee24476ee4649e5a23358143c0f1d9c251&#34;)};
    <a id="L177"></a>sendCKX := script.NewEvent(&#34;send ckx&#34;, []*script.Event{recvDone}, script.Send{msgChan, ckx});

    <a id="L179"></a>sendCCS := script.NewEvent(&#34;send ccs&#34;, []*script.Event{sendCKX}, script.Send{msgChan, changeCipherSpec{}});
    <a id="L180"></a>recvNCS := script.NewEvent(&#34;recv done&#34;, []*script.Event{sendCCS}, script.RecvMatch{controlChan, matchNewCipherSpec});

    <a id="L182"></a>finished := &amp;finishedMsg{nil, fromHex(&#34;c8faca5d242f4423325c5b1a&#34;)};
    <a id="L183"></a>sendFinished := script.NewEvent(&#34;send finished&#34;, []*script.Event{recvNCS}, script.Send{msgChan, finished});
    <a id="L184"></a>recvFinished := script.NewEvent(&#34;recv finished&#34;, []*script.Event{sendFinished}, script.RecvMatch{writeChan, matchFinished});
    <a id="L185"></a>setCipher := script.NewEvent(&#34;set cipher&#34;, []*script.Event{sendFinished}, script.RecvMatch{writeChan, matchSetCipher});
    <a id="L186"></a>recvConnectionState := script.NewEvent(&#34;recv state&#34;, []*script.Event{sendFinished}, script.Recv{controlChan, ConnectionState{true, &#34;TLS_RSA_WITH_RC4_128_SHA&#34;, 0}});

    <a id="L188"></a>err := script.Perform(0, []*script.Event{sendHello, setVersion, recvHello, recvCert, recvDone, sendCKX, sendCCS, recvNCS, sendFinished, setCipher, recvConnectionState, recvFinished});
    <a id="L189"></a>if err != nil {
        <a id="L190"></a>t.Errorf(&#34;Got error: %s&#34;, err)
    <a id="L191"></a>}
<a id="L192"></a>}

<a id="L194"></a>var testCertificate = fromHex(&#34;3082025930820203a003020102020900c2ec326b95228959300d06092a864886f70d01010505003054310b3009060355040613024155311330110603550408130a536f6d652d53746174653121301f060355040a1318496e7465726e6574205769646769747320507479204c7464310d300b0603550403130474657374301e170d3039313032303232323434355a170d3130313032303232323434355a3054310b3009060355040613024155311330110603550408130a536f6d652d53746174653121301f060355040a1318496e7465726e6574205769646769747320507479204c7464310d300b0603550403130474657374305c300d06092a864886f70d0101010500034b003048024100b2990f49c47dfa8cd400ae6a4d1b8a3b6a13642b23f28b003bfb97790ade9a4cc82b8b2a81747ddec08b6296e53a08c331687ef25c4bf4936ba1c0e6041e9d150203010001a381b73081b4301d0603551d0e0416041478a06086837c9293a8c9b70c0bdabdb9d77eeedf3081840603551d23047d307b801478a06086837c9293a8c9b70c0bdabdb9d77eeedfa158a4563054310b3009060355040613024155311330110603550408130a536f6d652d53746174653121301f060355040a1318496e7465726e6574205769646769747320507479204c7464310d300b0603550403130474657374820900c2ec326b95228959300c0603551d13040530030101ff300d06092a864886f70d0101050500034100ac23761ae1349d85a439caad4d0b932b09ea96de1917c3e0507c446f4838cb3076fb4d431db8c1987e96f1d7a8a2054dea3a64ec99a3f0eda4d47a163bf1f6ac&#34;)

<a id="L196"></a>func bigFromString(s string) *big.Int {
    <a id="L197"></a>ret := new(big.Int);
    <a id="L198"></a>ret.SetString(s, 10);
    <a id="L199"></a>return ret;
<a id="L200"></a>}

<a id="L202"></a>var testPrivateKey = &amp;rsa.PrivateKey{
    <a id="L203"></a>PublicKey: rsa.PublicKey{
        <a id="L204"></a>N: bigFromString(&#34;9353930466774385905609975137998169297361893554149986716853295022578535724979677252958524466350471210367835187480748268864277464700638583474144061408845077&#34;),
        <a id="L205"></a>E: 65537,
    <a id="L206"></a>},
    <a id="L207"></a>D: bigFromString(&#34;7266398431328116344057699379749222532279343923819063639497049039389899328538543087657733766554155839834519529439851673014800261285757759040931985506583861&#34;),
    <a id="L208"></a>P: bigFromString(&#34;98920366548084643601728869055592650835572950932266967461790948584315647051443&#34;),
    <a id="L209"></a>Q: bigFromString(&#34;94560208308847015747498523884063394671606671904944666360068158221458669711639&#34;),
<a id="L210"></a>}
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
