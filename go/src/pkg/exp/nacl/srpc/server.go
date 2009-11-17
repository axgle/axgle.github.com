<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/nacl/srpc/server.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/nacl/srpc/server.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// SRPC server</span>

<a id="L7"></a>package srpc

<a id="L9"></a>import (
    <a id="L10"></a>&#34;bytes&#34;;
    <a id="L11"></a>&#34;log&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;syscall&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// TODO(rsc): I&#39;d prefer to make this</span>
<a id="L17"></a><span class="comment">//	type Handler func(m *msg) Errno</span>
<a id="L18"></a><span class="comment">// but NaCl can&#39;t use closures.</span>
<a id="L19"></a><span class="comment">// The explicit interface is a way to attach state.</span>

<a id="L21"></a><span class="comment">// A Handler is a handler for an SRPC method.</span>
<a id="L22"></a><span class="comment">// It reads arguments from arg, checks size for array limits,</span>
<a id="L23"></a><span class="comment">// writes return values to ret, and returns an Errno status code.</span>
<a id="L24"></a>type Handler interface {
    <a id="L25"></a>Run(arg, ret []interface{}, size []int) Errno;
<a id="L26"></a>}

<a id="L28"></a>type method struct {
    <a id="L29"></a>name    string;
    <a id="L30"></a>fmt     string;
    <a id="L31"></a>handler Handler;
<a id="L32"></a>}

<a id="L34"></a>var rpcMethod []method

<a id="L36"></a><span class="comment">// BUG(rsc): Add&#39;s format string should be replaced by analyzing the</span>
<a id="L37"></a><span class="comment">// type of an arbitrary func passed in an interface{} using reflection.</span>

<a id="L39"></a><span class="comment">// Add registers a handler for the named method.</span>
<a id="L40"></a><span class="comment">// Fmt is a Native Client format string, a sequence of</span>
<a id="L41"></a><span class="comment">// alphabetic characters representing the types of the parameter values,</span>
<a id="L42"></a><span class="comment">// a colon, and then a sequence of alphabetic characters</span>
<a id="L43"></a><span class="comment">// representing the types of the returned values.</span>
<a id="L44"></a><span class="comment">// The format characters and corresponding dynamic types are:</span>
<a id="L45"></a><span class="comment">//</span>
<a id="L46"></a><span class="comment">//	b	bool</span>
<a id="L47"></a><span class="comment">//	C	[]byte</span>
<a id="L48"></a><span class="comment">//	d	float64</span>
<a id="L49"></a><span class="comment">//	D	[]float64</span>
<a id="L50"></a><span class="comment">//	h	int	// a file descriptor (aka handle)</span>
<a id="L51"></a><span class="comment">//	i	int32</span>
<a id="L52"></a><span class="comment">//	I	[]int32</span>
<a id="L53"></a><span class="comment">//	s	string</span>
<a id="L54"></a><span class="comment">//</span>
<a id="L55"></a>func Add(name, fmt string, handler Handler) {
    <a id="L56"></a>n := len(rpcMethod);
    <a id="L57"></a>if n &gt;= cap(rpcMethod) {
        <a id="L58"></a>a := make([]method, n, (n+4)*2);
        <a id="L59"></a>for i := range a {
            <a id="L60"></a>a[i] = rpcMethod[i]
        <a id="L61"></a>}
        <a id="L62"></a>rpcMethod = a;
    <a id="L63"></a>}
    <a id="L64"></a>rpcMethod = rpcMethod[0 : n+1];
    <a id="L65"></a>rpcMethod[n] = method{name, fmt, handler};
<a id="L66"></a>}

<a id="L68"></a><span class="comment">// Serve accepts new SRPC connections from the file descriptor fd</span>
<a id="L69"></a><span class="comment">// and answers RPCs issued on those connections.</span>
<a id="L70"></a><span class="comment">// It closes fd and returns an error if the imc_accept system call fails.</span>
<a id="L71"></a>func Serve(fd int) os.Error {
    <a id="L72"></a>defer syscall.Close(fd);

    <a id="L74"></a>for {
        <a id="L75"></a>cfd, _, e := syscall.Syscall(syscall.SYS_IMC_ACCEPT, uintptr(fd), 0, 0);
        <a id="L76"></a>if e != 0 {
            <a id="L77"></a>return os.NewSyscallError(&#34;imc_accept&#34;, int(e))
        <a id="L78"></a>}
        <a id="L79"></a>go serveLoop(int(cfd));
    <a id="L80"></a>}
    <a id="L81"></a>panic(&#34;unreachable&#34;);
<a id="L82"></a>}

<a id="L84"></a>func serveLoop(fd int) {
    <a id="L85"></a>c := make(chan *msg);
    <a id="L86"></a>go sendLoop(fd, c);

    <a id="L88"></a>var r msgReceiver;
    <a id="L89"></a>r.fd = fd;
    <a id="L90"></a>for {
        <a id="L91"></a>m, err := r.recv();
        <a id="L92"></a>if err != nil {
            <a id="L93"></a>break
        <a id="L94"></a>}
        <a id="L95"></a>m.unpackRequest();
        <a id="L96"></a>if !m.gotHeader {
            <a id="L97"></a>log.Stderrf(&#34;cannot unpack header: %s&#34;, m.status);
            <a id="L98"></a>continue;
        <a id="L99"></a>}
        <a id="L100"></a><span class="comment">// log.Stdoutf(&#34;&lt;- %#v&#34;, m);</span>
        <a id="L101"></a>m.isReq = false; <span class="comment">// set up for response</span>
        <a id="L102"></a>go serveMsg(m, c);
    <a id="L103"></a>}
    <a id="L104"></a>close(c);
<a id="L105"></a>}

<a id="L107"></a>func sendLoop(fd int, c &lt;-chan *msg) {
    <a id="L108"></a>var s msgSender;
    <a id="L109"></a>s.fd = fd;
    <a id="L110"></a>for m := range c {
        <a id="L111"></a><span class="comment">// log.Stdoutf(&#34;-&gt; %#v&#34;, m);</span>
        <a id="L112"></a>m.packResponse();
        <a id="L113"></a>s.send(m);
    <a id="L114"></a>}
    <a id="L115"></a>syscall.Close(fd);
<a id="L116"></a>}

<a id="L118"></a>func serveMsg(m *msg, c chan&lt;- *msg) {
    <a id="L119"></a>if m.status != OK {
        <a id="L120"></a>c &lt;- m;
        <a id="L121"></a>return;
    <a id="L122"></a>}
    <a id="L123"></a>if m.rpcNumber &gt;= uint32(len(rpcMethod)) {
        <a id="L124"></a>m.status = ErrBadRPCNumber;
        <a id="L125"></a>c &lt;- m;
        <a id="L126"></a>return;
    <a id="L127"></a>}

    <a id="L129"></a>meth := &amp;rpcMethod[m.rpcNumber];
    <a id="L130"></a>if meth.fmt != m.fmt {
        <a id="L131"></a>switch {
        <a id="L132"></a>case len(m.fmt) &lt; len(meth.fmt):
            <a id="L133"></a>m.status = ErrTooFewArgs
        <a id="L134"></a>case len(m.fmt) &gt; len(meth.fmt):
            <a id="L135"></a>m.status = ErrTooManyArgs
        <a id="L136"></a>default:
            <a id="L137"></a><span class="comment">// There&#39;s a type mismatch.</span>
            <a id="L138"></a><span class="comment">// It&#39;s an in-arg mismatch if the mismatch happens</span>
            <a id="L139"></a><span class="comment">// before the colon; otherwise it&#39;s an out-arg mismatch.</span>
            <a id="L140"></a>m.status = ErrInArgTypeMismatch;
            <a id="L141"></a>for i := 0; i &lt; len(m.fmt) &amp;&amp; m.fmt[i] == meth.fmt[i]; i++ {
                <a id="L142"></a>if m.fmt[i] == &#39;:&#39; {
                    <a id="L143"></a>m.status = ErrOutArgTypeMismatch;
                    <a id="L144"></a>break;
                <a id="L145"></a>}
            <a id="L146"></a>}
        <a id="L147"></a>}
        <a id="L148"></a>c &lt;- m;
        <a id="L149"></a>return;
    <a id="L150"></a>}

    <a id="L152"></a>m.status = meth.handler.Run(m.Arg, m.Ret, m.Size);
    <a id="L153"></a>c &lt;- m;
<a id="L154"></a>}

<a id="L156"></a><span class="comment">// ServeRuntime serves RPCs issued by the Native Client embedded runtime.</span>
<a id="L157"></a><span class="comment">// This should be called by main once all methods have been registered using Add.</span>
<a id="L158"></a>func ServeRuntime() os.Error {
    <a id="L159"></a><span class="comment">// Call getFd to check that we are running embedded.</span>
    <a id="L160"></a>if _, err := getFd(); err != nil {
        <a id="L161"></a>return err
    <a id="L162"></a>}

    <a id="L164"></a><span class="comment">// We are running embedded.</span>
    <a id="L165"></a><span class="comment">// The fd returned by getFd is a red herring.</span>
    <a id="L166"></a><span class="comment">// Accept connections on magic fd 3.</span>
    <a id="L167"></a>return Serve(3);
<a id="L168"></a>}

<a id="L170"></a><span class="comment">// getFd runs the srpc_get_fd system call.</span>
<a id="L171"></a>func getFd() (fd int, err os.Error) {
    <a id="L172"></a>r1, _, e := syscall.Syscall(syscall.SYS_SRPC_GET_FD, 0, 0, 0);
    <a id="L173"></a>return int(r1), os.NewSyscallError(&#34;srpc_get_fd&#34;, int(e));
<a id="L174"></a>}

<a id="L176"></a><span class="comment">// Enabled returns true if SRPC is enabled in the Native Client runtime.</span>
<a id="L177"></a>func Enabled() bool {
    <a id="L178"></a>_, err := getFd();
    <a id="L179"></a>return err == nil;
<a id="L180"></a>}

<a id="L182"></a><span class="comment">// Service #0, service_discovery, returns a list of the other services</span>
<a id="L183"></a><span class="comment">// and their argument formats.</span>
<a id="L184"></a>type serviceDiscovery struct{}

<a id="L186"></a>func (serviceDiscovery) Run(arg, ret []interface{}, size []int) Errno {
    <a id="L187"></a>var b bytes.Buffer;
    <a id="L188"></a>for _, m := range rpcMethod {
        <a id="L189"></a>b.WriteString(m.name);
        <a id="L190"></a>b.WriteByte(&#39;:&#39;);
        <a id="L191"></a>b.WriteString(m.fmt);
        <a id="L192"></a>b.WriteByte(&#39;\n&#39;);
    <a id="L193"></a>}
    <a id="L194"></a>if b.Len() &gt; size[0] {
        <a id="L195"></a>return ErrNoMemory
    <a id="L196"></a>}
    <a id="L197"></a>ret[0] = b.Bytes();
    <a id="L198"></a>return OK;
<a id="L199"></a>}

<a id="L201"></a>func init() { Add(&#34;service_discovery&#34;, &#34;:C&#34;, serviceDiscovery{}) }
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
