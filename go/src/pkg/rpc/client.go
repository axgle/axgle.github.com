<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/rpc/client.go</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/rpc/client.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package rpc

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bufio&#34;;
    <a id="L9"></a>&#34;gob&#34;;
    <a id="L10"></a>&#34;http&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;log&#34;;
    <a id="L13"></a>&#34;net&#34;;
    <a id="L14"></a>&#34;os&#34;;
    <a id="L15"></a>&#34;sync&#34;;
<a id="L16"></a>)

<a id="L18"></a><span class="comment">// Call represents an active RPC.</span>
<a id="L19"></a>type Call struct {
    <a id="L20"></a>ServiceMethod string;      <span class="comment">// The name of the service and method to call.</span>
    <a id="L21"></a>Args          interface{}; <span class="comment">// The argument to the function (*struct).</span>
    <a id="L22"></a>Reply         interface{}; <span class="comment">// The reply from the function (*struct).</span>
    <a id="L23"></a>Error         os.Error;    <span class="comment">// After completion, the error status.</span>
    <a id="L24"></a>Done          chan *Call;  <span class="comment">// Strobes when call is complete; value is the error status.</span>
    <a id="L25"></a>seq           uint64;
<a id="L26"></a>}

<a id="L28"></a><span class="comment">// Client represents an RPC Client.</span>
<a id="L29"></a><span class="comment">// There may be multiple outstanding Calls associated</span>
<a id="L30"></a><span class="comment">// with a single Client.</span>
<a id="L31"></a>type Client struct {
    <a id="L32"></a>mutex    sync.Mutex; <span class="comment">// protects pending, seq</span>
    <a id="L33"></a>shutdown os.Error;   <span class="comment">// non-nil if the client is shut down</span>
    <a id="L34"></a>sending  sync.Mutex;
    <a id="L35"></a>seq      uint64;
    <a id="L36"></a>conn     io.ReadWriteCloser;
    <a id="L37"></a>enc      *gob.Encoder;
    <a id="L38"></a>dec      *gob.Decoder;
    <a id="L39"></a>pending  map[uint64]*Call;
<a id="L40"></a>}

<a id="L42"></a>func (client *Client) send(c *Call) {
    <a id="L43"></a><span class="comment">// Register this call.</span>
    <a id="L44"></a>client.mutex.Lock();
    <a id="L45"></a>if client.shutdown != nil {
        <a id="L46"></a>c.Error = client.shutdown;
        <a id="L47"></a>client.mutex.Unlock();
        <a id="L48"></a>_ = c.Done &lt;- c; <span class="comment">// do not block</span>
        <a id="L49"></a>return;
    <a id="L50"></a>}
    <a id="L51"></a>c.seq = client.seq;
    <a id="L52"></a>client.seq++;
    <a id="L53"></a>client.pending[c.seq] = c;
    <a id="L54"></a>client.mutex.Unlock();

    <a id="L56"></a><span class="comment">// Encode and send the request.</span>
    <a id="L57"></a>request := new(Request);
    <a id="L58"></a>client.sending.Lock();
    <a id="L59"></a>request.Seq = c.seq;
    <a id="L60"></a>request.ServiceMethod = c.ServiceMethod;
    <a id="L61"></a>client.enc.Encode(request);
    <a id="L62"></a>err := client.enc.Encode(c.Args);
    <a id="L63"></a>if err != nil {
        <a id="L64"></a>panicln(&#34;rpc: client encode error:&#34;, err)
    <a id="L65"></a>}
    <a id="L66"></a>client.sending.Unlock();
<a id="L67"></a>}

<a id="L69"></a>func (client *Client) input() {
    <a id="L70"></a>var err os.Error;
    <a id="L71"></a>for err == nil {
        <a id="L72"></a>response := new(Response);
        <a id="L73"></a>err = client.dec.Decode(response);
        <a id="L74"></a>if err != nil {
            <a id="L75"></a>if err == os.EOF {
                <a id="L76"></a>err = io.ErrUnexpectedEOF
            <a id="L77"></a>}
            <a id="L78"></a>break;
        <a id="L79"></a>}
        <a id="L80"></a>seq := response.Seq;
        <a id="L81"></a>client.mutex.Lock();
        <a id="L82"></a>c := client.pending[seq];
        <a id="L83"></a>client.pending[seq] = c, false;
        <a id="L84"></a>client.mutex.Unlock();
        <a id="L85"></a>err = client.dec.Decode(c.Reply);
        <a id="L86"></a>c.Error = os.ErrorString(response.Error);
        <a id="L87"></a><span class="comment">// We don&#39;t want to block here.  It is the caller&#39;s responsibility to make</span>
        <a id="L88"></a><span class="comment">// sure the channel has enough buffer space. See comment in Go().</span>
        <a id="L89"></a>_ = c.Done &lt;- c; <span class="comment">// do not block</span>
    <a id="L90"></a>}
    <a id="L91"></a><span class="comment">// Terminate pending calls.</span>
    <a id="L92"></a>client.mutex.Lock();
    <a id="L93"></a>client.shutdown = err;
    <a id="L94"></a>for _, call := range client.pending {
        <a id="L95"></a>call.Error = err;
        <a id="L96"></a>_ = call.Done &lt;- call; <span class="comment">// do not block</span>
    <a id="L97"></a>}
    <a id="L98"></a>client.mutex.Unlock();
    <a id="L99"></a>log.Stderr(&#34;rpc: client protocol error:&#34;, err);
<a id="L100"></a>}

<a id="L102"></a><span class="comment">// NewClient returns a new Client to handle requests to the</span>
<a id="L103"></a><span class="comment">// set of services at the other end of the connection.</span>
<a id="L104"></a>func NewClient(conn io.ReadWriteCloser) *Client {
    <a id="L105"></a>client := new(Client);
    <a id="L106"></a>client.conn = conn;
    <a id="L107"></a>client.enc = gob.NewEncoder(conn);
    <a id="L108"></a>client.dec = gob.NewDecoder(conn);
    <a id="L109"></a>client.pending = make(map[uint64]*Call);
    <a id="L110"></a>go client.input();
    <a id="L111"></a>return client;
<a id="L112"></a>}

<a id="L114"></a><span class="comment">// DialHTTP connects to an HTTP RPC server at the specified network address.</span>
<a id="L115"></a>func DialHTTP(network, address string) (*Client, os.Error) {
    <a id="L116"></a>conn, err := net.Dial(network, &#34;&#34;, address);
    <a id="L117"></a>if err != nil {
        <a id="L118"></a>return nil, err
    <a id="L119"></a>}
    <a id="L120"></a>io.WriteString(conn, &#34;CONNECT &#34;+rpcPath+&#34; HTTP/1.0\n\n&#34;);

    <a id="L122"></a><span class="comment">// Require successful HTTP response</span>
    <a id="L123"></a><span class="comment">// before switching to RPC protocol.</span>
    <a id="L124"></a>resp, err := http.ReadResponse(bufio.NewReader(conn));
    <a id="L125"></a>if err == nil &amp;&amp; resp.Status == connected {
        <a id="L126"></a>return NewClient(conn), nil
    <a id="L127"></a>}
    <a id="L128"></a>if err == nil {
        <a id="L129"></a>err = os.ErrorString(&#34;unexpected HTTP response: &#34; + resp.Status)
    <a id="L130"></a>}
    <a id="L131"></a>conn.Close();
    <a id="L132"></a>return nil, &amp;net.OpError{&#34;dial-http&#34;, network + &#34; &#34; + address, nil, err};
<a id="L133"></a>}

<a id="L135"></a><span class="comment">// Dial connects to an RPC server at the specified network address.</span>
<a id="L136"></a>func Dial(network, address string) (*Client, os.Error) {
    <a id="L137"></a>conn, err := net.Dial(network, &#34;&#34;, address);
    <a id="L138"></a>if err != nil {
        <a id="L139"></a>return nil, err
    <a id="L140"></a>}
    <a id="L141"></a>return NewClient(conn), nil;
<a id="L142"></a>}

<a id="L144"></a><span class="comment">// Go invokes the function asynchronously.  It returns the Call structure representing</span>
<a id="L145"></a><span class="comment">// the invocation.  The done channel will signal when the call is complete by returning</span>
<a id="L146"></a><span class="comment">// the same Call object.  If done is nil, Go will allocate a new channel.</span>
<a id="L147"></a><span class="comment">// If non-nil, done must be buffered or Go will deliberately crash.</span>
<a id="L148"></a>func (client *Client) Go(serviceMethod string, args interface{}, reply interface{}, done chan *Call) *Call {
    <a id="L149"></a>c := new(Call);
    <a id="L150"></a>c.ServiceMethod = serviceMethod;
    <a id="L151"></a>c.Args = args;
    <a id="L152"></a>c.Reply = reply;
    <a id="L153"></a>if done == nil {
        <a id="L154"></a>done = make(chan *Call, 10) <span class="comment">// buffered.</span>
    <a id="L155"></a>} else {
        <a id="L156"></a><span class="comment">// If caller passes done != nil, it must arrange that</span>
        <a id="L157"></a><span class="comment">// done has enough buffer for the number of simultaneous</span>
        <a id="L158"></a><span class="comment">// RPCs that will be using that channel.  If the channel</span>
        <a id="L159"></a><span class="comment">// is totally unbuffered, it&#39;s best not to run at all.</span>
        <a id="L160"></a>if cap(done) == 0 {
            <a id="L161"></a>log.Crash(&#34;rpc: done channel is unbuffered&#34;)
        <a id="L162"></a>}
    <a id="L163"></a>}
    <a id="L164"></a>c.Done = done;
    <a id="L165"></a>if client.shutdown != nil {
        <a id="L166"></a>c.Error = client.shutdown;
        <a id="L167"></a>_ = c.Done &lt;- c; <span class="comment">// do not block</span>
        <a id="L168"></a>return c;
    <a id="L169"></a>}
    <a id="L170"></a>client.send(c);
    <a id="L171"></a>return c;
<a id="L172"></a>}

<a id="L174"></a><span class="comment">// Call invokes the named function, waits for it to complete, and returns its error status.</span>
<a id="L175"></a>func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) os.Error {
    <a id="L176"></a>if client.shutdown != nil {
        <a id="L177"></a>return client.shutdown
    <a id="L178"></a>}
    <a id="L179"></a>call := &lt;-client.Go(serviceMethod, args, reply, nil).Done;
    <a id="L180"></a>return call.Error;
<a id="L181"></a>}
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
