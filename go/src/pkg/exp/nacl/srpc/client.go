<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/nacl/srpc/client.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../../doc/style.css">
  <script type="text/javascript" src="../../../../../doc/godocs.js"></script>

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
        <a href="../../../../../index.html"><img src="../../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../../index.html">Source files</a></li>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/nacl/srpc/client.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package implements Native Client&#39;s simple RPC (SRPC).</span>
<a id="L6"></a>package srpc

<a id="L8"></a>import (
    <a id="L9"></a>&#34;bytes&#34;;
    <a id="L10"></a>&#34;log&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;sync&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">// A Client represents the client side of an SRPC connection.</span>
<a id="L16"></a>type Client struct {
    <a id="L17"></a>fd      int; <span class="comment">// fd to server</span>
    <a id="L18"></a>r       msgReceiver;
    <a id="L19"></a>s       msgSender;
    <a id="L20"></a>service map[string]srv; <span class="comment">// services by name</span>
    <a id="L21"></a>out     chan *msg;      <span class="comment">// send to out to write to connection</span>

    <a id="L23"></a>mu      sync.Mutex; <span class="comment">// protects pending, idGen</span>
    <a id="L24"></a>pending map[uint64]*RPC;
    <a id="L25"></a>idGen   uint64; <span class="comment">// generator for request IDs</span>
<a id="L26"></a>}

<a id="L28"></a><span class="comment">// A srv is a single method that the server offers.</span>
<a id="L29"></a>type srv struct {
    <a id="L30"></a>num uint32; <span class="comment">// method number</span>
    <a id="L31"></a>fmt string; <span class="comment">// argument format</span>
<a id="L32"></a>}

<a id="L34"></a><span class="comment">// An RPC represents a single RPC issued by a client.</span>
<a id="L35"></a>type RPC struct {
    <a id="L36"></a>Ret   []interface{}; <span class="comment">// Return values</span>
    <a id="L37"></a>Done  chan *RPC;     <span class="comment">// Channel where notification of done arrives</span>
    <a id="L38"></a>Errno Errno;         <span class="comment">// Status code</span>
    <a id="L39"></a>c     *Client;
    <a id="L40"></a>id    uint64; <span class="comment">// request id</span>
<a id="L41"></a>}

<a id="L43"></a><span class="comment">// NewClient allocates a new client using the file descriptor fd.</span>
<a id="L44"></a>func NewClient(fd int) (c *Client, err os.Error) {
    <a id="L45"></a>c = new(Client);
    <a id="L46"></a>c.fd = fd;
    <a id="L47"></a>c.r.fd = fd;
    <a id="L48"></a>c.s.fd = fd;
    <a id="L49"></a>c.service = make(map[string]srv);
    <a id="L50"></a>c.pending = make(map[uint64]*RPC);

    <a id="L52"></a><span class="comment">// service discovery request</span>
    <a id="L53"></a>m := &amp;msg{
        <a id="L54"></a>protocol: protocol,
        <a id="L55"></a>isReq: true,
        <a id="L56"></a>Ret: []interface{}{[]byte(nil)},
        <a id="L57"></a>Size: []int{4000},
    <a id="L58"></a>};
    <a id="L59"></a>m.packRequest();
    <a id="L60"></a>c.s.send(m);
    <a id="L61"></a>m, err = c.r.recv();
    <a id="L62"></a>if err != nil {
        <a id="L63"></a>return nil, err
    <a id="L64"></a>}
    <a id="L65"></a>m.unpackResponse();
    <a id="L66"></a>if m.status != OK {
        <a id="L67"></a>log.Stderrf(&#34;NewClient service_discovery: %s&#34;, m.status);
        <a id="L68"></a>return nil, m.status;
    <a id="L69"></a>}
    <a id="L70"></a>for n, line := range bytes.Split(m.Ret[0].([]byte), []byte{&#39;\n&#39;}, 0) {
        <a id="L71"></a>i := bytes.Index(line, []byte{&#39;:&#39;});
        <a id="L72"></a>if i &lt; 0 {
            <a id="L73"></a>continue
        <a id="L74"></a>}
        <a id="L75"></a>c.service[string(line[0:i])] = srv{uint32(n), string(line[i+1 : len(line)])};
    <a id="L76"></a>}

    <a id="L78"></a>c.out = make(chan *msg);
    <a id="L79"></a>go c.input();
    <a id="L80"></a>go c.output();
    <a id="L81"></a>return c, nil;
<a id="L82"></a>}

<a id="L84"></a>func (c *Client) input() {
    <a id="L85"></a>for {
        <a id="L86"></a>m, err := c.r.recv();
        <a id="L87"></a>if err != nil {
            <a id="L88"></a>log.Exitf(&#34;client recv: %s&#34;, err)
        <a id="L89"></a>}
        <a id="L90"></a>if m.unpackResponse(); m.status != OK {
            <a id="L91"></a>log.Stderrf(&#34;invalid message: %s&#34;, m.status);
            <a id="L92"></a>continue;
        <a id="L93"></a>}
        <a id="L94"></a>c.mu.Lock();
        <a id="L95"></a>rpc, ok := c.pending[m.requestId];
        <a id="L96"></a>if ok {
            <a id="L97"></a>c.pending[m.requestId] = nil, false
        <a id="L98"></a>}
        <a id="L99"></a>c.mu.Unlock();
        <a id="L100"></a>if !ok {
            <a id="L101"></a>log.Stderrf(&#34;unexpected response&#34;);
            <a id="L102"></a>continue;
        <a id="L103"></a>}
        <a id="L104"></a>rpc.Ret = m.Ret;
        <a id="L105"></a>rpc.Done &lt;- rpc;
    <a id="L106"></a>}
<a id="L107"></a>}

<a id="L109"></a>func (c *Client) output() {
    <a id="L110"></a>for m := range c.out {
        <a id="L111"></a>c.s.send(m)
    <a id="L112"></a>}
<a id="L113"></a>}

<a id="L115"></a><span class="comment">// NewRPC creates a new RPC on the client connection.</span>
<a id="L116"></a>func (c *Client) NewRPC(done chan *RPC) *RPC {
    <a id="L117"></a>if done == nil {
        <a id="L118"></a>done = make(chan *RPC)
    <a id="L119"></a>}
    <a id="L120"></a>c.mu.Lock();
    <a id="L121"></a>id := c.idGen;
    <a id="L122"></a>c.idGen++;
    <a id="L123"></a>c.mu.Unlock();
    <a id="L124"></a>return &amp;RPC{nil, done, OK, c, id};
<a id="L125"></a>}

<a id="L127"></a><span class="comment">// Start issues an RPC request for method name with the given arguments.</span>
<a id="L128"></a><span class="comment">// The RPC r must not be in use for another pending request.</span>
<a id="L129"></a><span class="comment">// To wait for the RPC to finish, receive from r.Done and then</span>
<a id="L130"></a><span class="comment">// inspect r.Ret and r.Errno.</span>
<a id="L131"></a>func (r *RPC) Start(name string, arg []interface{}) {
    <a id="L132"></a>var m msg;

    <a id="L134"></a>r.Errno = OK;
    <a id="L135"></a>r.c.mu.Lock();
    <a id="L136"></a>srv, ok := r.c.service[name];
    <a id="L137"></a>if !ok {
        <a id="L138"></a>r.c.mu.Unlock();
        <a id="L139"></a>r.Errno = ErrBadRPCNumber;
        <a id="L140"></a>r.Done &lt;- r;
        <a id="L141"></a>return;
    <a id="L142"></a>}
    <a id="L143"></a>r.c.pending[r.id] = r;
    <a id="L144"></a>r.c.mu.Unlock();

    <a id="L146"></a>m.protocol = protocol;
    <a id="L147"></a>m.requestId = r.id;
    <a id="L148"></a>m.isReq = true;
    <a id="L149"></a>m.rpcNumber = srv.num;
    <a id="L150"></a>m.Arg = arg;

    <a id="L152"></a><span class="comment">// Fill in the return values and sizes to generate</span>
    <a id="L153"></a><span class="comment">// the right type chars.  We&#39;ll take most any size.</span>

    <a id="L155"></a><span class="comment">// Skip over input arguments.</span>
    <a id="L156"></a><span class="comment">// We could check them against arg, but the server</span>
    <a id="L157"></a><span class="comment">// will do that anyway.</span>
    <a id="L158"></a>i := 0;
    <a id="L159"></a>for srv.fmt[i] != &#39;:&#39; {
        <a id="L160"></a>i++
    <a id="L161"></a>}
    <a id="L162"></a>fmt := srv.fmt[i+1 : len(srv.fmt)];

    <a id="L164"></a><span class="comment">// Now the return prototypes.</span>
    <a id="L165"></a>m.Ret = make([]interface{}, len(fmt)-i);
    <a id="L166"></a>m.Size = make([]int, len(fmt)-i);
    <a id="L167"></a>for i := 0; i &lt; len(fmt); i++ {
        <a id="L168"></a>switch fmt[i] {
        <a id="L169"></a>default:
            <a id="L170"></a>log.Exitf(&#34;unexpected service type %c&#34;, fmt[i])
        <a id="L171"></a>case &#39;b&#39;:
            <a id="L172"></a>m.Ret[i] = false
        <a id="L173"></a>case &#39;C&#39;:
            <a id="L174"></a>m.Ret[i] = []byte(nil);
            <a id="L175"></a>m.Size[i] = 1 &lt;&lt; 30;
        <a id="L176"></a>case &#39;d&#39;:
            <a id="L177"></a>m.Ret[i] = float64(0)
        <a id="L178"></a>case &#39;D&#39;:
            <a id="L179"></a>m.Ret[i] = []float64(nil);
            <a id="L180"></a>m.Size[i] = 1 &lt;&lt; 30;
        <a id="L181"></a>case &#39;h&#39;:
            <a id="L182"></a>m.Ret[i] = int(-1)
        <a id="L183"></a>case &#39;i&#39;:
            <a id="L184"></a>m.Ret[i] = int32(0)
        <a id="L185"></a>case &#39;I&#39;:
            <a id="L186"></a>m.Ret[i] = []int32(nil);
            <a id="L187"></a>m.Size[i] = 1 &lt;&lt; 30;
        <a id="L188"></a>case &#39;s&#39;:
            <a id="L189"></a>m.Ret[i] = &#34;&#34;;
            <a id="L190"></a>m.Size[i] = 1 &lt;&lt; 30;
        <a id="L191"></a>}
    <a id="L192"></a>}

    <a id="L194"></a>m.packRequest();
    <a id="L195"></a>r.c.out &lt;- &amp;m;
<a id="L196"></a>}

<a id="L198"></a><span class="comment">// Call is a convenient wrapper that starts the RPC request,</span>
<a id="L199"></a><span class="comment">// waits for it to finish, and then returns the results.</span>
<a id="L200"></a><span class="comment">// Its implementation is:</span>
<a id="L201"></a><span class="comment">//</span>
<a id="L202"></a><span class="comment">//	r.Start(name, arg);</span>
<a id="L203"></a><span class="comment">//	&lt;-r.Done;</span>
<a id="L204"></a><span class="comment">//	return r.Ret, r.Errno;</span>
<a id="L205"></a><span class="comment">//</span>
<a id="L206"></a>func (r *RPC) Call(name string, arg []interface{}) (ret []interface{}, err Errno) {
    <a id="L207"></a>r.Start(name, arg);
    <a id="L208"></a>&lt;-r.Done;
    <a id="L209"></a>return r.Ret, r.Errno;
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
