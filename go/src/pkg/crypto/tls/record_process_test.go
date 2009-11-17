<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/tls/record_process_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/tls/record_process_test.go</h1>

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
    <a id="L10"></a>&#34;testing/script&#34;;
<a id="L11"></a>)

<a id="L13"></a>func setup() (appDataChan chan []byte, requestChan chan interface{}, controlChan chan interface{}, recordChan chan *record, handshakeChan chan interface{}) {
    <a id="L14"></a>rp := new(recordProcessor);
    <a id="L15"></a>appDataChan = make(chan []byte);
    <a id="L16"></a>requestChan = make(chan interface{});
    <a id="L17"></a>controlChan = make(chan interface{});
    <a id="L18"></a>recordChan = make(chan *record);
    <a id="L19"></a>handshakeChan = make(chan interface{});

    <a id="L21"></a>go rp.loop(appDataChan, requestChan, controlChan, recordChan, handshakeChan);
    <a id="L22"></a>return;
<a id="L23"></a>}

<a id="L25"></a>func fromHex(s string) []byte {
    <a id="L26"></a>b, _ := hex.DecodeString(s);
    <a id="L27"></a>return b;
<a id="L28"></a>}

<a id="L30"></a>func TestNullConnectionState(t *testing.T) {
    <a id="L31"></a>_, requestChan, controlChan, recordChan, _ := setup();
    <a id="L32"></a>defer close(requestChan);
    <a id="L33"></a>defer close(controlChan);
    <a id="L34"></a>defer close(recordChan);

    <a id="L36"></a><span class="comment">// Test a simple request for the connection state.</span>
    <a id="L37"></a>replyChan := make(chan ConnectionState);
    <a id="L38"></a>sendReq := script.NewEvent(&#34;send request&#34;, nil, script.Send{requestChan, getConnectionState{replyChan}});
    <a id="L39"></a>getReply := script.NewEvent(&#34;get reply&#34;, []*script.Event{sendReq}, script.Recv{replyChan, ConnectionState{false, &#34;&#34;, 0}});

    <a id="L41"></a>err := script.Perform(0, []*script.Event{sendReq, getReply});
    <a id="L42"></a>if err != nil {
        <a id="L43"></a>t.Errorf(&#34;Got error: %s&#34;, err)
    <a id="L44"></a>}
<a id="L45"></a>}

<a id="L47"></a>func TestWaitConnectionState(t *testing.T) {
    <a id="L48"></a>_, requestChan, controlChan, recordChan, _ := setup();
    <a id="L49"></a>defer close(requestChan);
    <a id="L50"></a>defer close(controlChan);
    <a id="L51"></a>defer close(recordChan);

    <a id="L53"></a><span class="comment">// Test that waitConnectionState doesn&#39;t get a reply until the connection state changes.</span>
    <a id="L54"></a>replyChan := make(chan ConnectionState);
    <a id="L55"></a>sendReq := script.NewEvent(&#34;send request&#34;, nil, script.Send{requestChan, waitConnectionState{replyChan}});
    <a id="L56"></a>replyChan2 := make(chan ConnectionState);
    <a id="L57"></a>sendReq2 := script.NewEvent(&#34;send request 2&#34;, []*script.Event{sendReq}, script.Send{requestChan, getConnectionState{replyChan2}});
    <a id="L58"></a>getReply2 := script.NewEvent(&#34;get reply 2&#34;, []*script.Event{sendReq2}, script.Recv{replyChan2, ConnectionState{false, &#34;&#34;, 0}});
    <a id="L59"></a>sendState := script.NewEvent(&#34;send state&#34;, []*script.Event{getReply2}, script.Send{controlChan, ConnectionState{true, &#34;test&#34;, 1}});
    <a id="L60"></a>getReply := script.NewEvent(&#34;get reply&#34;, []*script.Event{sendState}, script.Recv{replyChan, ConnectionState{true, &#34;test&#34;, 1}});

    <a id="L62"></a>err := script.Perform(0, []*script.Event{sendReq, sendReq2, getReply2, sendState, getReply});
    <a id="L63"></a>if err != nil {
        <a id="L64"></a>t.Errorf(&#34;Got error: %s&#34;, err)
    <a id="L65"></a>}
<a id="L66"></a>}

<a id="L68"></a>func TestHandshakeAssembly(t *testing.T) {
    <a id="L69"></a>_, requestChan, controlChan, recordChan, handshakeChan := setup();
    <a id="L70"></a>defer close(requestChan);
    <a id="L71"></a>defer close(controlChan);
    <a id="L72"></a>defer close(recordChan);

    <a id="L74"></a><span class="comment">// Test the reassembly of a fragmented handshake message.</span>
    <a id="L75"></a>send1 := script.NewEvent(&#34;send 1&#34;, nil, script.Send{recordChan, &amp;record{recordTypeHandshake, 0, 0, fromHex(&#34;10000003&#34;)}});
    <a id="L76"></a>send2 := script.NewEvent(&#34;send 2&#34;, []*script.Event{send1}, script.Send{recordChan, &amp;record{recordTypeHandshake, 0, 0, fromHex(&#34;0001&#34;)}});
    <a id="L77"></a>send3 := script.NewEvent(&#34;send 3&#34;, []*script.Event{send2}, script.Send{recordChan, &amp;record{recordTypeHandshake, 0, 0, fromHex(&#34;42&#34;)}});
    <a id="L78"></a>recvMsg := script.NewEvent(&#34;recv&#34;, []*script.Event{send3}, script.Recv{handshakeChan, &amp;clientKeyExchangeMsg{fromHex(&#34;10000003000142&#34;), fromHex(&#34;42&#34;)}});

    <a id="L80"></a>err := script.Perform(0, []*script.Event{send1, send2, send3, recvMsg});
    <a id="L81"></a>if err != nil {
        <a id="L82"></a>t.Errorf(&#34;Got error: %s&#34;, err)
    <a id="L83"></a>}
<a id="L84"></a>}

<a id="L86"></a>func TestEarlyApplicationData(t *testing.T) {
    <a id="L87"></a>_, requestChan, controlChan, recordChan, handshakeChan := setup();
    <a id="L88"></a>defer close(requestChan);
    <a id="L89"></a>defer close(controlChan);
    <a id="L90"></a>defer close(recordChan);

    <a id="L92"></a><span class="comment">// Test that applicaton data received before the handshake has completed results in an error.</span>
    <a id="L93"></a>send := script.NewEvent(&#34;send&#34;, nil, script.Send{recordChan, &amp;record{recordTypeApplicationData, 0, 0, fromHex(&#34;&#34;)}});
    <a id="L94"></a>recv := script.NewEvent(&#34;recv&#34;, []*script.Event{send}, script.Closed{handshakeChan});

    <a id="L96"></a>err := script.Perform(0, []*script.Event{send, recv});
    <a id="L97"></a>if err != nil {
        <a id="L98"></a>t.Errorf(&#34;Got error: %s&#34;, err)
    <a id="L99"></a>}
<a id="L100"></a>}

<a id="L102"></a>func TestApplicationData(t *testing.T) {
    <a id="L103"></a>appDataChan, requestChan, controlChan, recordChan, handshakeChan := setup();
    <a id="L104"></a>defer close(requestChan);
    <a id="L105"></a>defer close(controlChan);
    <a id="L106"></a>defer close(recordChan);

    <a id="L108"></a><span class="comment">// Test that the application data is forwarded after a successful Finished message.</span>
    <a id="L109"></a>send1 := script.NewEvent(&#34;send 1&#34;, nil, script.Send{recordChan, &amp;record{recordTypeHandshake, 0, 0, fromHex(&#34;1400000c000000000000000000000000&#34;)}});
    <a id="L110"></a>recv1 := script.NewEvent(&#34;recv finished&#34;, []*script.Event{send1}, script.Recv{handshakeChan, &amp;finishedMsg{fromHex(&#34;1400000c000000000000000000000000&#34;), fromHex(&#34;000000000000000000000000&#34;)}});
    <a id="L111"></a>send2 := script.NewEvent(&#34;send connState&#34;, []*script.Event{recv1}, script.Send{controlChan, ConnectionState{true, &#34;&#34;, 0}});
    <a id="L112"></a>send3 := script.NewEvent(&#34;send 2&#34;, []*script.Event{send2}, script.Send{recordChan, &amp;record{recordTypeApplicationData, 0, 0, fromHex(&#34;0102&#34;)}});
    <a id="L113"></a>recv2 := script.NewEvent(&#34;recv data&#34;, []*script.Event{send3}, script.Recv{appDataChan, []byte{0x01, 0x02}});

    <a id="L115"></a>err := script.Perform(0, []*script.Event{send1, recv1, send2, send3, recv2});
    <a id="L116"></a>if err != nil {
        <a id="L117"></a>t.Errorf(&#34;Got error: %s&#34;, err)
    <a id="L118"></a>}
<a id="L119"></a>}

<a id="L121"></a>func TestInvalidChangeCipherSpec(t *testing.T) {
    <a id="L122"></a>appDataChan, requestChan, controlChan, recordChan, handshakeChan := setup();
    <a id="L123"></a>defer close(requestChan);
    <a id="L124"></a>defer close(controlChan);
    <a id="L125"></a>defer close(recordChan);

    <a id="L127"></a>send1 := script.NewEvent(&#34;send 1&#34;, nil, script.Send{recordChan, &amp;record{recordTypeChangeCipherSpec, 0, 0, []byte{1}}});
    <a id="L128"></a>recv1 := script.NewEvent(&#34;recv 1&#34;, []*script.Event{send1}, script.Recv{handshakeChan, changeCipherSpec{}});
    <a id="L129"></a>send2 := script.NewEvent(&#34;send 2&#34;, []*script.Event{recv1}, script.Send{controlChan, ConnectionState{false, &#34;&#34;, 42}});
    <a id="L130"></a>close := script.NewEvent(&#34;close 1&#34;, []*script.Event{send2}, script.Closed{appDataChan});
    <a id="L131"></a>close2 := script.NewEvent(&#34;close 2&#34;, []*script.Event{send2}, script.Closed{handshakeChan});

    <a id="L133"></a>err := script.Perform(0, []*script.Event{send1, recv1, send2, close, close2});
    <a id="L134"></a>if err != nil {
        <a id="L135"></a>t.Errorf(&#34;Got error: %s&#34;, err)
    <a id="L136"></a>}
<a id="L137"></a>}
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
