<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/tls/tls.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/tls/tls.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package partially implements the TLS 1.1 protocol, as specified in RFC 4346.</span>
<a id="L6"></a>package tls

<a id="L8"></a>import (
    <a id="L9"></a>&#34;bytes&#34;;
    <a id="L10"></a>&#34;io&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;net&#34;;
    <a id="L13"></a>&#34;time&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// A Conn represents a secure connection.</span>
<a id="L17"></a>type Conn struct {
    <a id="L18"></a>net.Conn;
    <a id="L19"></a>writeChan                 chan&lt;- []byte;
    <a id="L20"></a>readChan                  &lt;-chan []byte;
    <a id="L21"></a>requestChan               chan&lt;- interface{};
    <a id="L22"></a>readBuf                   []byte;
    <a id="L23"></a>eof                       bool;
    <a id="L24"></a>readTimeout, writeTimeout int64;
<a id="L25"></a>}

<a id="L27"></a>func timeout(c chan&lt;- bool, nsecs int64) {
    <a id="L28"></a>time.Sleep(nsecs);
    <a id="L29"></a>c &lt;- true;
<a id="L30"></a>}

<a id="L32"></a>func (tls *Conn) Read(p []byte) (int, os.Error) {
    <a id="L33"></a>if len(tls.readBuf) == 0 {
        <a id="L34"></a>if tls.eof {
            <a id="L35"></a>return 0, os.EOF
        <a id="L36"></a>}

        <a id="L38"></a>var timeoutChan chan bool;
        <a id="L39"></a>if tls.readTimeout &gt; 0 {
            <a id="L40"></a>timeoutChan = make(chan bool);
            <a id="L41"></a>go timeout(timeoutChan, tls.readTimeout);
        <a id="L42"></a>}

        <a id="L44"></a>select {
        <a id="L45"></a>case b := &lt;-tls.readChan:
            <a id="L46"></a>tls.readBuf = b
        <a id="L47"></a>case &lt;-timeoutChan:
            <a id="L48"></a>return 0, os.EAGAIN
        <a id="L49"></a>}

        <a id="L51"></a><span class="comment">// TLS distinguishes between orderly closes and truncations. An</span>
        <a id="L52"></a><span class="comment">// orderly close is represented by a zero length slice.</span>
        <a id="L53"></a>if closed(tls.readChan) {
            <a id="L54"></a>return 0, io.ErrUnexpectedEOF
        <a id="L55"></a>}
        <a id="L56"></a>if len(tls.readBuf) == 0 {
            <a id="L57"></a>tls.eof = true;
            <a id="L58"></a>return 0, os.EOF;
        <a id="L59"></a>}
    <a id="L60"></a>}

    <a id="L62"></a>n := bytes.Copy(p, tls.readBuf);
    <a id="L63"></a>tls.readBuf = tls.readBuf[n:len(tls.readBuf)];
    <a id="L64"></a>return n, nil;
<a id="L65"></a>}

<a id="L67"></a>func (tls *Conn) Write(p []byte) (int, os.Error) {
    <a id="L68"></a>if tls.eof || closed(tls.readChan) {
        <a id="L69"></a>return 0, os.EOF
    <a id="L70"></a>}

    <a id="L72"></a>var timeoutChan chan bool;
    <a id="L73"></a>if tls.writeTimeout &gt; 0 {
        <a id="L74"></a>timeoutChan = make(chan bool);
        <a id="L75"></a>go timeout(timeoutChan, tls.writeTimeout);
    <a id="L76"></a>}

    <a id="L78"></a>select {
    <a id="L79"></a>case tls.writeChan &lt;- p:
    <a id="L80"></a>case &lt;-timeoutChan:
        <a id="L81"></a>return 0, os.EAGAIN
    <a id="L82"></a>}

    <a id="L84"></a>return len(p), nil;
<a id="L85"></a>}

<a id="L87"></a>func (tls *Conn) Close() os.Error {
    <a id="L88"></a>close(tls.writeChan);
    <a id="L89"></a>close(tls.requestChan);
    <a id="L90"></a>tls.eof = true;
    <a id="L91"></a>return nil;
<a id="L92"></a>}

<a id="L94"></a>func (tls *Conn) SetTimeout(nsec int64) os.Error {
    <a id="L95"></a>tls.readTimeout = nsec;
    <a id="L96"></a>tls.writeTimeout = nsec;
    <a id="L97"></a>return nil;
<a id="L98"></a>}

<a id="L100"></a>func (tls *Conn) SetReadTimeout(nsec int64) os.Error {
    <a id="L101"></a>tls.readTimeout = nsec;
    <a id="L102"></a>return nil;
<a id="L103"></a>}

<a id="L105"></a>func (tls *Conn) SetWriteTimeout(nsec int64) os.Error {
    <a id="L106"></a>tls.writeTimeout = nsec;
    <a id="L107"></a>return nil;
<a id="L108"></a>}

<a id="L110"></a>func (tls *Conn) GetConnectionState() ConnectionState {
    <a id="L111"></a>replyChan := make(chan ConnectionState);
    <a id="L112"></a>tls.requestChan &lt;- getConnectionState{replyChan};
    <a id="L113"></a>return &lt;-replyChan;
<a id="L114"></a>}

<a id="L116"></a><span class="comment">// Server establishes a secure connection over the given connection and acts</span>
<a id="L117"></a><span class="comment">// as a TLS server.</span>
<a id="L118"></a>func Server(conn net.Conn, config *Config) *Conn {
    <a id="L119"></a>tls := new(Conn);
    <a id="L120"></a>tls.Conn = conn;

    <a id="L122"></a>writeChan := make(chan []byte);
    <a id="L123"></a>readChan := make(chan []byte);
    <a id="L124"></a>requestChan := make(chan interface{});

    <a id="L126"></a>tls.writeChan = writeChan;
    <a id="L127"></a>tls.readChan = readChan;
    <a id="L128"></a>tls.requestChan = requestChan;

    <a id="L130"></a>handshakeWriterChan := make(chan interface{});
    <a id="L131"></a>processorHandshakeChan := make(chan interface{});
    <a id="L132"></a>handshakeProcessorChan := make(chan interface{});
    <a id="L133"></a>readerProcessorChan := make(chan *record);

    <a id="L135"></a>go new(recordWriter).loop(conn, writeChan, handshakeWriterChan);
    <a id="L136"></a>go recordReader(readerProcessorChan, conn);
    <a id="L137"></a>go new(recordProcessor).loop(readChan, requestChan, handshakeProcessorChan, readerProcessorChan, processorHandshakeChan);
    <a id="L138"></a>go new(serverHandshake).loop(handshakeWriterChan, handshakeProcessorChan, processorHandshakeChan, config);

    <a id="L140"></a>return tls;
<a id="L141"></a>}

<a id="L143"></a>type Listener struct {
    <a id="L144"></a>listener net.Listener;
    <a id="L145"></a>config   *Config;
<a id="L146"></a>}

<a id="L148"></a>func (l Listener) Accept() (c net.Conn, err os.Error) {
    <a id="L149"></a>c, err = l.listener.Accept();
    <a id="L150"></a>if err != nil {
        <a id="L151"></a>return
    <a id="L152"></a>}

    <a id="L154"></a>c = Server(c, l.config);
    <a id="L155"></a>return;
<a id="L156"></a>}

<a id="L158"></a>func (l Listener) Close() os.Error { return l.listener.Close() }

<a id="L160"></a>func (l Listener) Addr() net.Addr { return l.listener.Addr() }

<a id="L162"></a><span class="comment">// NewListener creates a Listener which accepts connections from an inner</span>
<a id="L163"></a><span class="comment">// Listener and wraps each connection with Server.</span>
<a id="L164"></a>func NewListener(listener net.Listener, config *Config) (l Listener) {
    <a id="L165"></a>l.listener = listener;
    <a id="L166"></a>l.config = config;
    <a id="L167"></a>return;
<a id="L168"></a>}
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
