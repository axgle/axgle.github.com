<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/http/server.go</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

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
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

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
  <h1 id="generatedHeader">Source file /src/pkg/http/server.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// HTTP server.  See RFC 2616.</span>

<a id="L7"></a><span class="comment">// TODO(rsc):</span>
<a id="L8"></a><span class="comment">//	logging</span>
<a id="L9"></a><span class="comment">//	cgi support</span>
<a id="L10"></a><span class="comment">//	post support</span>

<a id="L12"></a>package http

<a id="L14"></a>import (
    <a id="L15"></a>&#34;bufio&#34;;
    <a id="L16"></a>&#34;fmt&#34;;
    <a id="L17"></a>&#34;io&#34;;
    <a id="L18"></a>&#34;log&#34;;
    <a id="L19"></a>&#34;net&#34;;
    <a id="L20"></a>&#34;os&#34;;
    <a id="L21"></a>&#34;path&#34;;
    <a id="L22"></a>&#34;strconv&#34;;
    <a id="L23"></a>&#34;strings&#34;;
<a id="L24"></a>)

<a id="L26"></a><span class="comment">// Errors introduced by the HTTP server.</span>
<a id="L27"></a>var (
    <a id="L28"></a>ErrWriteAfterFlush = os.NewError(&#34;Conn.Write called after Flush&#34;);
    <a id="L29"></a>ErrHijacked        = os.NewError(&#34;Conn has been hijacked&#34;);
<a id="L30"></a>)

<a id="L32"></a><span class="comment">// Objects implementing the Handler interface can be</span>
<a id="L33"></a><span class="comment">// registered to serve a particular path or subtree</span>
<a id="L34"></a><span class="comment">// in the HTTP server.</span>
<a id="L35"></a>type Handler interface {
    <a id="L36"></a>ServeHTTP(*Conn, *Request);
<a id="L37"></a>}

<a id="L39"></a><span class="comment">// A Conn represents the server side of a single active HTTP connection.</span>
<a id="L40"></a>type Conn struct {
    <a id="L41"></a>RemoteAddr string;   <span class="comment">// network address of remote side</span>
    <a id="L42"></a>Req        *Request; <span class="comment">// current HTTP request</span>

    <a id="L44"></a>rwc      io.ReadWriteCloser; <span class="comment">// i/o connection</span>
    <a id="L45"></a>buf      *bufio.ReadWriter;  <span class="comment">// buffered rwc</span>
    <a id="L46"></a>handler  Handler;            <span class="comment">// request handler</span>
    <a id="L47"></a>hijacked bool;               <span class="comment">// connection has been hijacked by handler</span>

    <a id="L49"></a><span class="comment">// state for the current reply</span>
    <a id="L50"></a>closeAfterReply bool;              <span class="comment">// close connection after this reply</span>
    <a id="L51"></a>chunking        bool;              <span class="comment">// using chunked transfer encoding for reply body</span>
    <a id="L52"></a>wroteHeader     bool;              <span class="comment">// reply header has been written</span>
    <a id="L53"></a>header          map[string]string; <span class="comment">// reply header parameters</span>
    <a id="L54"></a>written         int64;             <span class="comment">// number of bytes written in body</span>
    <a id="L55"></a>status          int;               <span class="comment">// status code passed to WriteHeader</span>
<a id="L56"></a>}

<a id="L58"></a><span class="comment">// Create new connection from rwc.</span>
<a id="L59"></a>func newConn(rwc net.Conn, handler Handler) (c *Conn, err os.Error) {
    <a id="L60"></a>c = new(Conn);
    <a id="L61"></a>if a := rwc.RemoteAddr(); a != nil {
        <a id="L62"></a>c.RemoteAddr = a.String()
    <a id="L63"></a>}
    <a id="L64"></a>c.handler = handler;
    <a id="L65"></a>c.rwc = rwc;
    <a id="L66"></a>br := bufio.NewReader(rwc);
    <a id="L67"></a>bw := bufio.NewWriter(rwc);
    <a id="L68"></a>c.buf = bufio.NewReadWriter(br, bw);
    <a id="L69"></a>return c, nil;
<a id="L70"></a>}

<a id="L72"></a><span class="comment">// Read next request from connection.</span>
<a id="L73"></a>func (c *Conn) readRequest() (req *Request, err os.Error) {
    <a id="L74"></a>if c.hijacked {
        <a id="L75"></a>return nil, ErrHijacked
    <a id="L76"></a>}
    <a id="L77"></a>if req, err = ReadRequest(c.buf.Reader); err != nil {
        <a id="L78"></a>return nil, err
    <a id="L79"></a>}

    <a id="L81"></a><span class="comment">// Reset per-request connection state.</span>
    <a id="L82"></a>c.header = make(map[string]string);
    <a id="L83"></a>c.wroteHeader = false;
    <a id="L84"></a>c.Req = req;

    <a id="L86"></a><span class="comment">// Default output is HTML encoded in UTF-8.</span>
    <a id="L87"></a>c.SetHeader(&#34;Content-Type&#34;, &#34;text/html; charset=utf-8&#34;);

    <a id="L89"></a>if req.ProtoAtLeast(1, 1) {
        <a id="L90"></a><span class="comment">// HTTP/1.1 or greater: use chunked transfer encoding</span>
        <a id="L91"></a><span class="comment">// to avoid closing the connection at EOF.</span>
        <a id="L92"></a>c.chunking = true;
        <a id="L93"></a>c.SetHeader(&#34;Transfer-Encoding&#34;, &#34;chunked&#34;);
    <a id="L94"></a>} else {
        <a id="L95"></a><span class="comment">// HTTP version &lt; 1.1: cannot do chunked transfer</span>
        <a id="L96"></a><span class="comment">// encoding, so signal EOF by closing connection.</span>
        <a id="L97"></a><span class="comment">// Could avoid closing the connection if there is</span>
        <a id="L98"></a><span class="comment">// a Content-Length: header in the response,</span>
        <a id="L99"></a><span class="comment">// but everyone who expects persistent connections</span>
        <a id="L100"></a><span class="comment">// does HTTP/1.1 now.</span>
        <a id="L101"></a>c.closeAfterReply = true;
        <a id="L102"></a>c.chunking = false;
    <a id="L103"></a>}

    <a id="L105"></a>return req, nil;
<a id="L106"></a>}

<a id="L108"></a><span class="comment">// SetHeader sets a header line in the eventual reply.</span>
<a id="L109"></a><span class="comment">// For example, SetHeader(&#34;Content-Type&#34;, &#34;text/html; charset=utf-8&#34;)</span>
<a id="L110"></a><span class="comment">// will result in the header line</span>
<a id="L111"></a><span class="comment">//</span>
<a id="L112"></a><span class="comment">//	Content-Type: text/html; charset=utf-8</span>
<a id="L113"></a><span class="comment">//</span>
<a id="L114"></a><span class="comment">// being sent.  UTF-8 encoded HTML is the default setting for</span>
<a id="L115"></a><span class="comment">// Content-Type in this library, so users need not make that</span>
<a id="L116"></a><span class="comment">// particular call.  Calls to SetHeader after WriteHeader (or Write)</span>
<a id="L117"></a><span class="comment">// are ignored.</span>
<a id="L118"></a>func (c *Conn) SetHeader(hdr, val string) { c.header[CanonicalHeaderKey(hdr)] = val }

<a id="L120"></a><span class="comment">// WriteHeader sends an HTTP response header with status code.</span>
<a id="L121"></a><span class="comment">// If WriteHeader is not called explicitly, the first call to Write</span>
<a id="L122"></a><span class="comment">// will trigger an implicit WriteHeader(http.StatusOK).</span>
<a id="L123"></a><span class="comment">// Thus explicit calls to WriteHeader are mainly used to</span>
<a id="L124"></a><span class="comment">// send error codes.</span>
<a id="L125"></a>func (c *Conn) WriteHeader(code int) {
    <a id="L126"></a>if c.hijacked {
        <a id="L127"></a>log.Stderr(&#34;http: Conn.WriteHeader on hijacked connection&#34;);
        <a id="L128"></a>return;
    <a id="L129"></a>}
    <a id="L130"></a>if c.wroteHeader {
        <a id="L131"></a>log.Stderr(&#34;http: multiple Conn.WriteHeader calls&#34;);
        <a id="L132"></a>return;
    <a id="L133"></a>}
    <a id="L134"></a>c.wroteHeader = true;
    <a id="L135"></a>c.status = code;
    <a id="L136"></a>c.written = 0;
    <a id="L137"></a>if !c.Req.ProtoAtLeast(1, 0) {
        <a id="L138"></a>return
    <a id="L139"></a>}
    <a id="L140"></a>proto := &#34;HTTP/1.0&#34;;
    <a id="L141"></a>if c.Req.ProtoAtLeast(1, 1) {
        <a id="L142"></a>proto = &#34;HTTP/1.1&#34;
    <a id="L143"></a>}
    <a id="L144"></a>codestring := strconv.Itoa(code);
    <a id="L145"></a>text, ok := statusText[code];
    <a id="L146"></a>if !ok {
        <a id="L147"></a>text = &#34;status code &#34; + codestring
    <a id="L148"></a>}
    <a id="L149"></a>io.WriteString(c.buf, proto+&#34; &#34;+codestring+&#34; &#34;+text+&#34;\r\n&#34;);
    <a id="L150"></a>for k, v := range c.header {
        <a id="L151"></a>io.WriteString(c.buf, k+&#34;: &#34;+v+&#34;\r\n&#34;)
    <a id="L152"></a>}
    <a id="L153"></a>io.WriteString(c.buf, &#34;\r\n&#34;);
<a id="L154"></a>}

<a id="L156"></a><span class="comment">// Write writes the data to the connection as part of an HTTP reply.</span>
<a id="L157"></a><span class="comment">// If WriteHeader has not yet been called, Write calls WriteHeader(http.StatusOK)</span>
<a id="L158"></a><span class="comment">// before writing the data.</span>
<a id="L159"></a>func (c *Conn) Write(data []byte) (n int, err os.Error) {
    <a id="L160"></a>if c.hijacked {
        <a id="L161"></a>log.Stderr(&#34;http: Conn.Write on hijacked connection&#34;);
        <a id="L162"></a>return 0, ErrHijacked;
    <a id="L163"></a>}
    <a id="L164"></a>if !c.wroteHeader {
        <a id="L165"></a>c.WriteHeader(StatusOK)
    <a id="L166"></a>}
    <a id="L167"></a>if len(data) == 0 {
        <a id="L168"></a>return 0, nil
    <a id="L169"></a>}

    <a id="L171"></a>c.written += int64(len(data)); <span class="comment">// ignoring errors, for errorKludge</span>

    <a id="L173"></a><span class="comment">// TODO(rsc): if chunking happened after the buffering,</span>
    <a id="L174"></a><span class="comment">// then there would be fewer chunk headers.</span>
    <a id="L175"></a><span class="comment">// On the other hand, it would make hijacking more difficult.</span>
    <a id="L176"></a>if c.chunking {
        <a id="L177"></a>fmt.Fprintf(c.buf, &#34;%x\r\n&#34;, len(data)) <span class="comment">// TODO(rsc): use strconv not fmt</span>
    <a id="L178"></a>}
    <a id="L179"></a>n, err = c.buf.Write(data);
    <a id="L180"></a>if err == nil &amp;&amp; c.chunking {
        <a id="L181"></a>if n != len(data) {
            <a id="L182"></a>err = io.ErrShortWrite
        <a id="L183"></a>}
        <a id="L184"></a>if err == nil {
            <a id="L185"></a>io.WriteString(c.buf, &#34;\r\n&#34;)
        <a id="L186"></a>}
    <a id="L187"></a>}

    <a id="L189"></a>return n, err;
<a id="L190"></a>}

<a id="L192"></a><span class="comment">// If this is an error reply (4xx or 5xx)</span>
<a id="L193"></a><span class="comment">// and the handler wrote some data explaining the error,</span>
<a id="L194"></a><span class="comment">// some browsers (i.e., Chrome, Internet Explorer)</span>
<a id="L195"></a><span class="comment">// will show their own error instead unless the error is</span>
<a id="L196"></a><span class="comment">// long enough.  The minimum lengths used in those</span>
<a id="L197"></a><span class="comment">// browsers are in the 256-512 range.</span>
<a id="L198"></a><span class="comment">// Pad to 1024 bytes.</span>
<a id="L199"></a>func errorKludge(c *Conn, req *Request) {
    <a id="L200"></a>const min = 1024;

    <a id="L202"></a><span class="comment">// Is this an error?</span>
    <a id="L203"></a>if kind := c.status / 100; kind != 4 &amp;&amp; kind != 5 {
        <a id="L204"></a>return
    <a id="L205"></a>}

    <a id="L207"></a><span class="comment">// Did the handler supply any info?  Enough?</span>
    <a id="L208"></a>if c.written == 0 || c.written &gt;= min {
        <a id="L209"></a>return
    <a id="L210"></a>}

    <a id="L212"></a><span class="comment">// Is it a broken browser?</span>
    <a id="L213"></a>var msg string;
    <a id="L214"></a>switch agent := req.UserAgent; {
    <a id="L215"></a>case strings.Index(agent, &#34;MSIE&#34;) &gt;= 0:
        <a id="L216"></a>msg = &#34;Internet Explorer&#34;
    <a id="L217"></a>case strings.Index(agent, &#34;Chrome/&#34;) &gt;= 0:
        <a id="L218"></a>msg = &#34;Chrome&#34;
    <a id="L219"></a>default:
        <a id="L220"></a>return
    <a id="L221"></a>}
    <a id="L222"></a>msg += &#34; would ignore this error page if this text weren&#39;t here.\n&#34;;

    <a id="L224"></a><span class="comment">// Is it text?  (&#34;Content-Type&#34; is always in the map)</span>
    <a id="L225"></a>baseType := strings.Split(c.header[&#34;Content-Type&#34;], &#34;;&#34;, 2)[0];
    <a id="L226"></a>switch baseType {
    <a id="L227"></a>case &#34;text/html&#34;:
        <a id="L228"></a>io.WriteString(c, &#34;&lt;!-- &#34;);
        <a id="L229"></a>for c.written &lt; min {
            <a id="L230"></a>io.WriteString(c, msg)
        <a id="L231"></a>}
        <a id="L232"></a>io.WriteString(c, &#34; --&gt;&#34;);
    <a id="L233"></a>case &#34;text/plain&#34;:
        <a id="L234"></a>io.WriteString(c, &#34;\n&#34;);
        <a id="L235"></a>for c.written &lt; min {
            <a id="L236"></a>io.WriteString(c, msg)
        <a id="L237"></a>}
    <a id="L238"></a>}
<a id="L239"></a>}

<a id="L241"></a>func (c *Conn) flush() {
    <a id="L242"></a>if !c.wroteHeader {
        <a id="L243"></a>c.WriteHeader(StatusOK)
    <a id="L244"></a>}
    <a id="L245"></a>errorKludge(c, c.Req);
    <a id="L246"></a>if c.chunking {
        <a id="L247"></a>io.WriteString(c.buf, &#34;0\r\n&#34;);
        <a id="L248"></a><span class="comment">// trailer key/value pairs, followed by blank line</span>
        <a id="L249"></a>io.WriteString(c.buf, &#34;\r\n&#34;);
    <a id="L250"></a>}
    <a id="L251"></a>c.buf.Flush();
<a id="L252"></a>}

<a id="L254"></a><span class="comment">// Close the connection.</span>
<a id="L255"></a>func (c *Conn) close() {
    <a id="L256"></a>if c.buf != nil {
        <a id="L257"></a>c.buf.Flush();
        <a id="L258"></a>c.buf = nil;
    <a id="L259"></a>}
    <a id="L260"></a>if c.rwc != nil {
        <a id="L261"></a>c.rwc.Close();
        <a id="L262"></a>c.rwc = nil;
    <a id="L263"></a>}
<a id="L264"></a>}

<a id="L266"></a><span class="comment">// Serve a new connection.</span>
<a id="L267"></a>func (c *Conn) serve() {
    <a id="L268"></a>for {
        <a id="L269"></a>req, err := c.readRequest();
        <a id="L270"></a>if err != nil {
            <a id="L271"></a>break
        <a id="L272"></a>}
        <a id="L273"></a><span class="comment">// HTTP cannot have multiple simultaneous active requests.</span>
        <a id="L274"></a><span class="comment">// Until the server replies to this request, it can&#39;t read another,</span>
        <a id="L275"></a><span class="comment">// so we might as well run the handler in this goroutine.</span>
        <a id="L276"></a>c.handler.ServeHTTP(c, req);
        <a id="L277"></a>if c.hijacked {
            <a id="L278"></a>return
        <a id="L279"></a>}
        <a id="L280"></a>c.flush();
        <a id="L281"></a>if c.closeAfterReply {
            <a id="L282"></a>break
        <a id="L283"></a>}
    <a id="L284"></a>}
    <a id="L285"></a>c.close();
<a id="L286"></a>}

<a id="L288"></a><span class="comment">// Hijack lets the caller take over the connection.</span>
<a id="L289"></a><span class="comment">// After a call to c.Hijack(), the HTTP server library</span>
<a id="L290"></a><span class="comment">// will not do anything else with the connection.</span>
<a id="L291"></a><span class="comment">// It becomes the caller&#39;s responsibility to manage</span>
<a id="L292"></a><span class="comment">// and close the connection.</span>
<a id="L293"></a>func (c *Conn) Hijack() (rwc io.ReadWriteCloser, buf *bufio.ReadWriter, err os.Error) {
    <a id="L294"></a>if c.hijacked {
        <a id="L295"></a>return nil, nil, ErrHijacked
    <a id="L296"></a>}
    <a id="L297"></a>c.hijacked = true;
    <a id="L298"></a>rwc = c.rwc;
    <a id="L299"></a>buf = c.buf;
    <a id="L300"></a>c.rwc = nil;
    <a id="L301"></a>c.buf = nil;
    <a id="L302"></a>return;
<a id="L303"></a>}

<a id="L305"></a><span class="comment">// The HandlerFunc type is an adapter to allow the use of</span>
<a id="L306"></a><span class="comment">// ordinary functions as HTTP handlers.  If f is a function</span>
<a id="L307"></a><span class="comment">// with the appropriate signature, HandlerFunc(f) is a</span>
<a id="L308"></a><span class="comment">// Handler object that calls f.</span>
<a id="L309"></a>type HandlerFunc func(*Conn, *Request)

<a id="L311"></a><span class="comment">// ServeHTTP calls f(c, req).</span>
<a id="L312"></a>func (f HandlerFunc) ServeHTTP(c *Conn, req *Request) {
    <a id="L313"></a>f(c, req)
<a id="L314"></a>}

<a id="L316"></a><span class="comment">// Helper handlers</span>

<a id="L318"></a><span class="comment">// NotFound replies to the request with an HTTP 404 not found error.</span>
<a id="L319"></a>func NotFound(c *Conn, req *Request) {
    <a id="L320"></a>c.SetHeader(&#34;Content-Type&#34;, &#34;text/plain; charset=utf-8&#34;);
    <a id="L321"></a>c.WriteHeader(StatusNotFound);
    <a id="L322"></a>io.WriteString(c, &#34;404 page not found\n&#34;);
<a id="L323"></a>}

<a id="L325"></a><span class="comment">// NotFoundHandler returns a simple request handler</span>
<a id="L326"></a><span class="comment">// that replies to each request with a ``404 page not found&#39;&#39; reply.</span>
<a id="L327"></a>func NotFoundHandler() Handler { return HandlerFunc(NotFound) }

<a id="L329"></a><span class="comment">// Redirect replies to the request with a redirect to url,</span>
<a id="L330"></a><span class="comment">// which may be a path relative to the request path.</span>
<a id="L331"></a>func Redirect(c *Conn, url string, code int) {
    <a id="L332"></a><span class="comment">// RFC2616 recommends that a short note &#34;SHOULD&#34; be included in the</span>
    <a id="L333"></a><span class="comment">// response because older user agents may not understand 301/307.</span>
    <a id="L334"></a>note := &#34;&lt;a href=\&#34;%v\&#34;&gt;&#34; + statusText[code] + &#34;&lt;/a&gt;.\n&#34;;
    <a id="L335"></a>if c.Req.Method == &#34;POST&#34; {
        <a id="L336"></a>note = &#34;&#34;
    <a id="L337"></a>}

    <a id="L339"></a>u, err := ParseURL(url);
    <a id="L340"></a>if err != nil {
        <a id="L341"></a>goto finish
    <a id="L342"></a>}

    <a id="L344"></a><span class="comment">// If url was relative, make absolute by</span>
    <a id="L345"></a><span class="comment">// combining with request path.</span>
    <a id="L346"></a><span class="comment">// The browser would probably do this for us,</span>
    <a id="L347"></a><span class="comment">// but doing it ourselves is more reliable.</span>

    <a id="L349"></a><span class="comment">// NOTE(rsc): RFC 2616 says that the Location</span>
    <a id="L350"></a><span class="comment">// line must be an absolute URI, like</span>
    <a id="L351"></a><span class="comment">// &#34;http://www.google.com/redirect/&#34;,</span>
    <a id="L352"></a><span class="comment">// not a path like &#34;/redirect/&#34;.</span>
    <a id="L353"></a><span class="comment">// Unfortunately, we don&#39;t know what to</span>
    <a id="L354"></a><span class="comment">// put in the host name section to get the</span>
    <a id="L355"></a><span class="comment">// client to connect to us again, so we can&#39;t</span>
    <a id="L356"></a><span class="comment">// know the right absolute URI to send back.</span>
    <a id="L357"></a><span class="comment">// Because of this problem, no one pays attention</span>
    <a id="L358"></a><span class="comment">// to the RFC; they all send back just a new path.</span>
    <a id="L359"></a><span class="comment">// So do we.</span>
    <a id="L360"></a>oldpath := c.Req.URL.Path;
    <a id="L361"></a>if oldpath == &#34;&#34; { <span class="comment">// should not happen, but avoid a crash if it does</span>
        <a id="L362"></a>oldpath = &#34;/&#34;
    <a id="L363"></a>}
    <a id="L364"></a>if u.Scheme == &#34;&#34; {
        <a id="L365"></a><span class="comment">// no leading http://server</span>
        <a id="L366"></a>if url == &#34;&#34; || url[0] != &#39;/&#39; {
            <a id="L367"></a><span class="comment">// make relative path absolute</span>
            <a id="L368"></a>olddir, _ := path.Split(oldpath);
            <a id="L369"></a>url = olddir + url;
        <a id="L370"></a>}

        <a id="L372"></a><span class="comment">// clean up but preserve trailing slash</span>
        <a id="L373"></a>trailing := url[len(url)-1] == &#39;/&#39;;
        <a id="L374"></a>url = path.Clean(url);
        <a id="L375"></a>if trailing &amp;&amp; url[len(url)-1] != &#39;/&#39; {
            <a id="L376"></a>url += &#34;/&#34;
        <a id="L377"></a>}
    <a id="L378"></a>}

<a id="L380"></a>finish:
    <a id="L381"></a>c.SetHeader(&#34;Location&#34;, url);
    <a id="L382"></a>c.WriteHeader(code);
    <a id="L383"></a>fmt.Fprintf(c, note, url);
<a id="L384"></a>}

<a id="L386"></a><span class="comment">// Redirect to a fixed URL</span>
<a id="L387"></a>type redirectHandler struct {
    <a id="L388"></a>url  string;
    <a id="L389"></a>code int;
<a id="L390"></a>}

<a id="L392"></a>func (rh *redirectHandler) ServeHTTP(c *Conn, req *Request) {
    <a id="L393"></a>Redirect(c, rh.url, rh.code)
<a id="L394"></a>}

<a id="L396"></a><span class="comment">// RedirectHandler returns a request handler that redirects</span>
<a id="L397"></a><span class="comment">// each request it receives to the given url using the given</span>
<a id="L398"></a><span class="comment">// status code.</span>
<a id="L399"></a>func RedirectHandler(url string, code int) Handler {
    <a id="L400"></a>return &amp;redirectHandler{url, code}
<a id="L401"></a>}

<a id="L403"></a><span class="comment">// ServeMux is an HTTP request multiplexer.</span>
<a id="L404"></a><span class="comment">// It matches the URL of each incoming request against a list of registered</span>
<a id="L405"></a><span class="comment">// patterns and calls the handler for the pattern that</span>
<a id="L406"></a><span class="comment">// most closely matches the URL.</span>
<a id="L407"></a><span class="comment">//</span>
<a id="L408"></a><span class="comment">// Patterns named fixed paths, like &#34;/favicon.ico&#34;,</span>
<a id="L409"></a><span class="comment">// or subtrees, like &#34;/images/&#34; (note the trailing slash).</span>
<a id="L410"></a><span class="comment">// Patterns must begin with /.</span>
<a id="L411"></a><span class="comment">// Longer patterns take precedence over shorter ones, so that</span>
<a id="L412"></a><span class="comment">// if there are handlers registered for both &#34;/images/&#34;</span>
<a id="L413"></a><span class="comment">// and &#34;/images/thumbnails/&#34;, the latter handler will be</span>
<a id="L414"></a><span class="comment">// called for paths beginning &#34;/images/thumbnails/&#34; and the</span>
<a id="L415"></a><span class="comment">// former will receiver requests for any other paths in the</span>
<a id="L416"></a><span class="comment">// &#34;/images/&#34; subtree.</span>
<a id="L417"></a><span class="comment">//</span>
<a id="L418"></a><span class="comment">// In the future, the pattern syntax may be relaxed to allow</span>
<a id="L419"></a><span class="comment">// an optional host-name at the beginning of the pattern,</span>
<a id="L420"></a><span class="comment">// so that a handler might register for the two patterns</span>
<a id="L421"></a><span class="comment">// &#34;/codesearch&#34; and &#34;codesearch.google.com/&#34;</span>
<a id="L422"></a><span class="comment">// without taking over requests for http://www.google.com/.</span>
<a id="L423"></a><span class="comment">//</span>
<a id="L424"></a><span class="comment">// ServeMux also takes care of sanitizing the URL request path,</span>
<a id="L425"></a><span class="comment">// redirecting any request containing . or .. elements to an</span>
<a id="L426"></a><span class="comment">// equivalent .- and ..-free URL.</span>
<a id="L427"></a>type ServeMux struct {
    <a id="L428"></a>m map[string]Handler;
<a id="L429"></a>}

<a id="L431"></a><span class="comment">// NewServeMux allocates and returns a new ServeMux.</span>
<a id="L432"></a>func NewServeMux() *ServeMux { return &amp;ServeMux{make(map[string]Handler)} }

<a id="L434"></a><span class="comment">// DefaultServeMux is the default ServeMux used by Serve.</span>
<a id="L435"></a>var DefaultServeMux = NewServeMux()

<a id="L437"></a><span class="comment">// Does path match pattern?</span>
<a id="L438"></a>func pathMatch(pattern, path string) bool {
    <a id="L439"></a>if len(pattern) == 0 {
        <a id="L440"></a><span class="comment">// should not happen</span>
        <a id="L441"></a>return false
    <a id="L442"></a>}
    <a id="L443"></a>n := len(pattern);
    <a id="L444"></a>if pattern[n-1] != &#39;/&#39; {
        <a id="L445"></a>return pattern == path
    <a id="L446"></a>}
    <a id="L447"></a>return len(path) &gt;= n &amp;&amp; path[0:n] == pattern;
<a id="L448"></a>}

<a id="L450"></a><span class="comment">// Return the canonical path for p, eliminating . and .. elements.</span>
<a id="L451"></a>func cleanPath(p string) string {
    <a id="L452"></a>if p == &#34;&#34; {
        <a id="L453"></a>return &#34;/&#34;
    <a id="L454"></a>}
    <a id="L455"></a>if p[0] != &#39;/&#39; {
        <a id="L456"></a>p = &#34;/&#34; + p
    <a id="L457"></a>}
    <a id="L458"></a>np := path.Clean(p);
    <a id="L459"></a><span class="comment">// path.Clean removes trailing slash except for root;</span>
    <a id="L460"></a><span class="comment">// put the trailing slash back if necessary.</span>
    <a id="L461"></a>if p[len(p)-1] == &#39;/&#39; &amp;&amp; np != &#34;/&#34; {
        <a id="L462"></a>np += &#34;/&#34;
    <a id="L463"></a>}
    <a id="L464"></a>return np;
<a id="L465"></a>}

<a id="L467"></a><span class="comment">// ServeHTTP dispatches the request to the handler whose</span>
<a id="L468"></a><span class="comment">// pattern most closely matches the request URL.</span>
<a id="L469"></a>func (mux *ServeMux) ServeHTTP(c *Conn, req *Request) {
    <a id="L470"></a><span class="comment">// Clean path to canonical form and redirect.</span>
    <a id="L471"></a>if p := cleanPath(req.URL.Path); p != req.URL.Path {
        <a id="L472"></a>c.SetHeader(&#34;Location&#34;, p);
        <a id="L473"></a>c.WriteHeader(StatusMovedPermanently);
        <a id="L474"></a>return;
    <a id="L475"></a>}

    <a id="L477"></a><span class="comment">// Most-specific (longest) pattern wins.</span>
    <a id="L478"></a>var h Handler;
    <a id="L479"></a>var n = 0;
    <a id="L480"></a>for k, v := range mux.m {
        <a id="L481"></a>if !pathMatch(k, req.URL.Path) {
            <a id="L482"></a>continue
        <a id="L483"></a>}
        <a id="L484"></a>if h == nil || len(k) &gt; n {
            <a id="L485"></a>n = len(k);
            <a id="L486"></a>h = v;
        <a id="L487"></a>}
    <a id="L488"></a>}
    <a id="L489"></a>if h == nil {
        <a id="L490"></a>h = NotFoundHandler()
    <a id="L491"></a>}
    <a id="L492"></a>h.ServeHTTP(c, req);
<a id="L493"></a>}

<a id="L495"></a><span class="comment">// Handle registers the handler for the given pattern.</span>
<a id="L496"></a>func (mux *ServeMux) Handle(pattern string, handler Handler) {
    <a id="L497"></a>if pattern == &#34;&#34; || pattern[0] != &#39;/&#39; {
        <a id="L498"></a>panicln(&#34;http: invalid pattern&#34;, pattern)
    <a id="L499"></a>}

    <a id="L501"></a>mux.m[pattern] = handler;

    <a id="L503"></a><span class="comment">// Helpful behavior:</span>
    <a id="L504"></a><span class="comment">// If pattern is /tree/, insert permanent redirect for /tree.</span>
    <a id="L505"></a>n := len(pattern);
    <a id="L506"></a>if n &gt; 0 &amp;&amp; pattern[n-1] == &#39;/&#39; {
        <a id="L507"></a>mux.m[pattern[0:n-1]] = RedirectHandler(pattern, StatusMovedPermanently)
    <a id="L508"></a>}
<a id="L509"></a>}

<a id="L511"></a><span class="comment">// Handle registers the handler for the given pattern</span>
<a id="L512"></a><span class="comment">// in the DefaultServeMux.</span>
<a id="L513"></a>func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }

<a id="L515"></a><span class="comment">// Serve accepts incoming HTTP connections on the listener l,</span>
<a id="L516"></a><span class="comment">// creating a new service thread for each.  The service threads</span>
<a id="L517"></a><span class="comment">// read requests and then call handler to reply to them.</span>
<a id="L518"></a><span class="comment">// Handler is typically nil, in which case the DefaultServeMux is used.</span>
<a id="L519"></a>func Serve(l net.Listener, handler Handler) os.Error {
    <a id="L520"></a>if handler == nil {
        <a id="L521"></a>handler = DefaultServeMux
    <a id="L522"></a>}
    <a id="L523"></a>for {
        <a id="L524"></a>rw, e := l.Accept();
        <a id="L525"></a>if e != nil {
            <a id="L526"></a>return e
        <a id="L527"></a>}
        <a id="L528"></a>c, err := newConn(rw, handler);
        <a id="L529"></a>if err != nil {
            <a id="L530"></a>continue
        <a id="L531"></a>}
        <a id="L532"></a>go c.serve();
    <a id="L533"></a>}
    <a id="L534"></a>panic(&#34;not reached&#34;);
<a id="L535"></a>}

<a id="L537"></a><span class="comment">// ListenAndServe listens on the TCP network address addr</span>
<a id="L538"></a><span class="comment">// and then calls Serve with handler to handle requests</span>
<a id="L539"></a><span class="comment">// on incoming connections.  Handler is typically nil,</span>
<a id="L540"></a><span class="comment">// in which case the DefaultServeMux is used.</span>
<a id="L541"></a><span class="comment">//</span>
<a id="L542"></a><span class="comment">// A trivial example server is:</span>
<a id="L543"></a><span class="comment">//</span>
<a id="L544"></a><span class="comment">//	package main</span>
<a id="L545"></a><span class="comment">//</span>
<a id="L546"></a><span class="comment">//	import (</span>
<a id="L547"></a><span class="comment">//		&#34;http&#34;;</span>
<a id="L548"></a><span class="comment">//		&#34;io&#34;;</span>
<a id="L549"></a><span class="comment">//	)</span>
<a id="L550"></a><span class="comment">//</span>
<a id="L551"></a><span class="comment">//	// hello world, the web server</span>
<a id="L552"></a><span class="comment">//	func HelloServer(c *http.Conn, req *http.Request) {</span>
<a id="L553"></a><span class="comment">//		io.WriteString(c, &#34;hello, world!\n&#34;);</span>
<a id="L554"></a><span class="comment">//	}</span>
<a id="L555"></a><span class="comment">//</span>
<a id="L556"></a><span class="comment">//	func main() {</span>
<a id="L557"></a><span class="comment">//		http.Handle(&#34;/hello&#34;, http.HandlerFunc(HelloServer));</span>
<a id="L558"></a><span class="comment">//		err := http.ListenAndServe(&#34;:12345&#34;, nil);</span>
<a id="L559"></a><span class="comment">//		if err != nil {</span>
<a id="L560"></a><span class="comment">//			panic(&#34;ListenAndServe: &#34;, err.String())</span>
<a id="L561"></a><span class="comment">//		}</span>
<a id="L562"></a><span class="comment">//	}</span>
<a id="L563"></a>func ListenAndServe(addr string, handler Handler) os.Error {
    <a id="L564"></a>l, e := net.Listen(&#34;tcp&#34;, addr);
    <a id="L565"></a>if e != nil {
        <a id="L566"></a>return e
    <a id="L567"></a>}
    <a id="L568"></a>e = Serve(l, handler);
    <a id="L569"></a>l.Close();
    <a id="L570"></a>return e;
<a id="L571"></a>}
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
