<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/tls/record_write.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/tls/record_write.go</h1>

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
    <a id="L8"></a>&#34;fmt&#34;;
    <a id="L9"></a>&#34;hash&#34;;
    <a id="L10"></a>&#34;io&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">// writerEnableApplicationData is a message which instructs recordWriter to</span>
<a id="L14"></a><span class="comment">// start reading and transmitting data from the application data channel.</span>
<a id="L15"></a>type writerEnableApplicationData struct{}

<a id="L17"></a><span class="comment">// writerChangeCipherSpec updates the encryption and MAC functions and resets</span>
<a id="L18"></a><span class="comment">// the sequence count.</span>
<a id="L19"></a>type writerChangeCipherSpec struct {
    <a id="L20"></a>encryptor encryptor;
    <a id="L21"></a>mac       hash.Hash;
<a id="L22"></a>}

<a id="L24"></a><span class="comment">// writerSetVersion sets the version number bytes that we included in the</span>
<a id="L25"></a><span class="comment">// record header for future records.</span>
<a id="L26"></a>type writerSetVersion struct {
    <a id="L27"></a>major, minor uint8;
<a id="L28"></a>}

<a id="L30"></a><span class="comment">// A recordWriter accepts messages from the handshake processor and</span>
<a id="L31"></a><span class="comment">// application data. It writes them to the outgoing connection and blocks on</span>
<a id="L32"></a><span class="comment">// writing. It doesn&#39;t read from the application data channel until the</span>
<a id="L33"></a><span class="comment">// handshake processor has signaled that the handshake is complete.</span>
<a id="L34"></a>type recordWriter struct {
    <a id="L35"></a>writer       io.Writer;
    <a id="L36"></a>encryptor    encryptor;
    <a id="L37"></a>mac          hash.Hash;
    <a id="L38"></a>seqNum       uint64;
    <a id="L39"></a>major, minor uint8;
    <a id="L40"></a>shutdown     bool;
    <a id="L41"></a>appChan      &lt;-chan []byte;
    <a id="L42"></a>controlChan  &lt;-chan interface{};
    <a id="L43"></a>header       [13]byte;
<a id="L44"></a>}

<a id="L46"></a>func (w *recordWriter) loop(writer io.Writer, appChan &lt;-chan []byte, controlChan &lt;-chan interface{}) {
    <a id="L47"></a>w.writer = writer;
    <a id="L48"></a>w.encryptor = nop{};
    <a id="L49"></a>w.mac = nop{};
    <a id="L50"></a>w.appChan = appChan;
    <a id="L51"></a>w.controlChan = controlChan;

    <a id="L53"></a>for !w.shutdown {
        <a id="L54"></a>msg := &lt;-controlChan;
        <a id="L55"></a>if _, ok := msg.(writerEnableApplicationData); ok {
            <a id="L56"></a>break
        <a id="L57"></a>}
        <a id="L58"></a>w.processControlMessage(msg);
    <a id="L59"></a>}

    <a id="L61"></a>for !w.shutdown {
        <a id="L62"></a><span class="comment">// Always process control messages first.</span>
        <a id="L63"></a>if controlMsg, ok := &lt;-controlChan; ok {
            <a id="L64"></a>w.processControlMessage(controlMsg);
            <a id="L65"></a>continue;
        <a id="L66"></a>}

        <a id="L68"></a>select {
        <a id="L69"></a>case controlMsg := &lt;-controlChan:
            <a id="L70"></a>w.processControlMessage(controlMsg)
        <a id="L71"></a>case appMsg := &lt;-appChan:
            <a id="L72"></a>w.processAppMessage(appMsg)
        <a id="L73"></a>}
    <a id="L74"></a>}

    <a id="L76"></a>if !closed(appChan) {
        <a id="L77"></a>go func() {
            <a id="L78"></a>for _ = range appChan {
            <a id="L79"></a>}
        <a id="L80"></a>}()
    <a id="L81"></a>}
    <a id="L82"></a>if !closed(controlChan) {
        <a id="L83"></a>go func() {
            <a id="L84"></a>for _ = range controlChan {
            <a id="L85"></a>}
        <a id="L86"></a>}()
    <a id="L87"></a>}
<a id="L88"></a>}

<a id="L90"></a><span class="comment">// fillMACHeader generates a MAC header. See RFC 4346, section 6.2.3.1.</span>
<a id="L91"></a>func fillMACHeader(header *[13]byte, seqNum uint64, length int, r *record) {
    <a id="L92"></a>header[0] = uint8(seqNum &gt;&gt; 56);
    <a id="L93"></a>header[1] = uint8(seqNum &gt;&gt; 48);
    <a id="L94"></a>header[2] = uint8(seqNum &gt;&gt; 40);
    <a id="L95"></a>header[3] = uint8(seqNum &gt;&gt; 32);
    <a id="L96"></a>header[4] = uint8(seqNum &gt;&gt; 24);
    <a id="L97"></a>header[5] = uint8(seqNum &gt;&gt; 16);
    <a id="L98"></a>header[6] = uint8(seqNum &gt;&gt; 8);
    <a id="L99"></a>header[7] = uint8(seqNum);
    <a id="L100"></a>header[8] = uint8(r.contentType);
    <a id="L101"></a>header[9] = r.major;
    <a id="L102"></a>header[10] = r.minor;
    <a id="L103"></a>header[11] = uint8(length &gt;&gt; 8);
    <a id="L104"></a>header[12] = uint8(length);
<a id="L105"></a>}

<a id="L107"></a>func (w *recordWriter) writeRecord(r *record) {
    <a id="L108"></a>w.mac.Reset();

    <a id="L110"></a>fillMACHeader(&amp;w.header, w.seqNum, len(r.payload), r);

    <a id="L112"></a>w.mac.Write(w.header[0:13]);
    <a id="L113"></a>w.mac.Write(r.payload);
    <a id="L114"></a>macBytes := w.mac.Sum();

    <a id="L116"></a>w.encryptor.XORKeyStream(r.payload);
    <a id="L117"></a>w.encryptor.XORKeyStream(macBytes);

    <a id="L119"></a>length := len(r.payload) + len(macBytes);
    <a id="L120"></a>w.header[11] = uint8(length &gt;&gt; 8);
    <a id="L121"></a>w.header[12] = uint8(length);
    <a id="L122"></a>w.writer.Write(w.header[8:13]);
    <a id="L123"></a>w.writer.Write(r.payload);
    <a id="L124"></a>w.writer.Write(macBytes);

    <a id="L126"></a>w.seqNum++;
<a id="L127"></a>}

<a id="L129"></a>func (w *recordWriter) processControlMessage(controlMsg interface{}) {
    <a id="L130"></a>if controlMsg == nil {
        <a id="L131"></a>w.shutdown = true;
        <a id="L132"></a>return;
    <a id="L133"></a>}

    <a id="L135"></a>switch msg := controlMsg.(type) {
    <a id="L136"></a>case writerChangeCipherSpec:
        <a id="L137"></a>w.writeRecord(&amp;record{recordTypeChangeCipherSpec, w.major, w.minor, []byte{0x01}});
        <a id="L138"></a>w.encryptor = msg.encryptor;
        <a id="L139"></a>w.mac = msg.mac;
        <a id="L140"></a>w.seqNum = 0;
    <a id="L141"></a>case writerSetVersion:
        <a id="L142"></a>w.major = msg.major;
        <a id="L143"></a>w.minor = msg.minor;
    <a id="L144"></a>case alert:
        <a id="L145"></a>w.writeRecord(&amp;record{recordTypeAlert, w.major, w.minor, []byte{byte(msg.level), byte(msg.error)}})
    <a id="L146"></a>case handshakeMessage:
        <a id="L147"></a><span class="comment">// TODO(agl): marshal may return a slice too large for a single record.</span>
        <a id="L148"></a>w.writeRecord(&amp;record{recordTypeHandshake, w.major, w.minor, msg.marshal()})
    <a id="L149"></a>default:
        <a id="L150"></a>fmt.Printf(&#34;processControlMessage: unknown %#v\n&#34;, msg)
    <a id="L151"></a>}
<a id="L152"></a>}

<a id="L154"></a>func (w *recordWriter) processAppMessage(appMsg []byte) {
    <a id="L155"></a>if closed(w.appChan) {
        <a id="L156"></a>w.writeRecord(&amp;record{recordTypeApplicationData, w.major, w.minor, []byte{byte(alertCloseNotify)}});
        <a id="L157"></a>w.shutdown = true;
        <a id="L158"></a>return;
    <a id="L159"></a>}

    <a id="L161"></a>var done int;
    <a id="L162"></a>for done &lt; len(appMsg) {
        <a id="L163"></a>todo := len(appMsg);
        <a id="L164"></a>if todo &gt; maxTLSPlaintext {
            <a id="L165"></a>todo = maxTLSPlaintext
        <a id="L166"></a>}
        <a id="L167"></a>w.writeRecord(&amp;record{recordTypeApplicationData, w.major, w.minor, appMsg[done : done+todo]});
        <a id="L168"></a>done += todo;
    <a id="L169"></a>}
<a id="L170"></a>}
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
