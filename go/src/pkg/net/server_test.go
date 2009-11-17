<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/server_test.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/net/server_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package net

<a id="L7"></a>import (
    <a id="L8"></a>&#34;io&#34;;
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>&#34;strings&#34;;
    <a id="L11"></a>&#34;syscall&#34;;
    <a id="L12"></a>&#34;testing&#34;;
<a id="L13"></a>)

<a id="L15"></a>func runEcho(fd io.ReadWriter, done chan&lt;- int) {
    <a id="L16"></a>var buf [1024]byte;

    <a id="L18"></a>for {
        <a id="L19"></a>n, err := fd.Read(&amp;buf);
        <a id="L20"></a>if err != nil || n == 0 {
            <a id="L21"></a>break
        <a id="L22"></a>}
        <a id="L23"></a>fd.Write(buf[0:n]);
    <a id="L24"></a>}
    <a id="L25"></a>done &lt;- 1;
<a id="L26"></a>}

<a id="L28"></a>func runServe(t *testing.T, network, addr string, listening chan&lt;- string, done chan&lt;- int) {
    <a id="L29"></a>l, err := Listen(network, addr);
    <a id="L30"></a>if err != nil {
        <a id="L31"></a>t.Fatalf(&#34;net.Listen(%q, %q) = _, %v&#34;, network, addr, err)
    <a id="L32"></a>}
    <a id="L33"></a>listening &lt;- l.Addr().String();

    <a id="L35"></a>for {
        <a id="L36"></a>fd, err := l.Accept();
        <a id="L37"></a>if err != nil {
            <a id="L38"></a>break
        <a id="L39"></a>}
        <a id="L40"></a>echodone := make(chan int);
        <a id="L41"></a>go runEcho(fd, echodone);
        <a id="L42"></a>&lt;-echodone; <span class="comment">// make sure Echo stops</span>
        <a id="L43"></a>l.Close();
    <a id="L44"></a>}
    <a id="L45"></a>done &lt;- 1;
<a id="L46"></a>}

<a id="L48"></a>func connect(t *testing.T, network, addr string) {
    <a id="L49"></a>var laddr string;
    <a id="L50"></a>if network == &#34;unixgram&#34; {
        <a id="L51"></a>laddr = addr + &#34;.local&#34;
    <a id="L52"></a>}
    <a id="L53"></a>fd, err := Dial(network, laddr, addr);
    <a id="L54"></a>if err != nil {
        <a id="L55"></a>t.Fatalf(&#34;net.Dial(%q, %q, %q) = _, %v&#34;, network, laddr, addr, err)
    <a id="L56"></a>}

    <a id="L58"></a>b := strings.Bytes(&#34;hello, world\n&#34;);
    <a id="L59"></a>var b1 [100]byte;

    <a id="L61"></a>n, errno := fd.Write(b);
    <a id="L62"></a>if n != len(b) {
        <a id="L63"></a>t.Fatalf(&#34;fd.Write(%q) = %d, %v&#34;, b, n, errno)
    <a id="L64"></a>}

    <a id="L66"></a>n, errno = fd.Read(&amp;b1);
    <a id="L67"></a>if n != len(b) {
        <a id="L68"></a>t.Fatalf(&#34;fd.Read() = %d, %v&#34;, n, errno)
    <a id="L69"></a>}
    <a id="L70"></a>fd.Close();
<a id="L71"></a>}

<a id="L73"></a>func doTest(t *testing.T, network, listenaddr, dialaddr string) {
    <a id="L74"></a>t.Logf(&#34;Test %s %s %s\n&#34;, network, listenaddr, dialaddr);
    <a id="L75"></a>listening := make(chan string);
    <a id="L76"></a>done := make(chan int);
    <a id="L77"></a>if network == &#34;tcp&#34; {
        <a id="L78"></a>listenaddr += &#34;:0&#34; <span class="comment">// any available port</span>
    <a id="L79"></a>}
    <a id="L80"></a>go runServe(t, network, listenaddr, listening, done);
    <a id="L81"></a>addr := &lt;-listening; <span class="comment">// wait for server to start</span>
    <a id="L82"></a>if network == &#34;tcp&#34; {
        <a id="L83"></a>dialaddr += addr[strings.LastIndex(addr, &#34;:&#34;):len(addr)]
    <a id="L84"></a>}
    <a id="L85"></a>connect(t, network, dialaddr);
    <a id="L86"></a>&lt;-done; <span class="comment">// make sure server stopped</span>
<a id="L87"></a>}

<a id="L89"></a>func TestTCPServer(t *testing.T) {
    <a id="L90"></a>doTest(t, &#34;tcp&#34;, &#34;0.0.0.0&#34;, &#34;127.0.0.1&#34;);
    <a id="L91"></a>doTest(t, &#34;tcp&#34;, &#34;&#34;, &#34;127.0.0.1&#34;);
    <a id="L92"></a>if kernelSupportsIPv6() {
        <a id="L93"></a>doTest(t, &#34;tcp&#34;, &#34;[::]&#34;, &#34;[::ffff:127.0.0.1]&#34;);
        <a id="L94"></a>doTest(t, &#34;tcp&#34;, &#34;[::]&#34;, &#34;127.0.0.1&#34;);
        <a id="L95"></a>doTest(t, &#34;tcp&#34;, &#34;0.0.0.0&#34;, &#34;[::ffff:127.0.0.1]&#34;);
    <a id="L96"></a>}
<a id="L97"></a>}

<a id="L99"></a>func TestUnixServer(t *testing.T) {
    <a id="L100"></a>os.Remove(&#34;/tmp/gotest.net&#34;);
    <a id="L101"></a>doTest(t, &#34;unix&#34;, &#34;/tmp/gotest.net&#34;, &#34;/tmp/gotest.net&#34;);
    <a id="L102"></a>os.Remove(&#34;/tmp/gotest.net&#34;);
    <a id="L103"></a>if syscall.OS == &#34;linux&#34; {
        <a id="L104"></a><span class="comment">// Test abstract unix domain socket, a Linux-ism</span>
        <a id="L105"></a>doTest(t, &#34;unix&#34;, &#34;@gotest/net&#34;, &#34;@gotest/net&#34;)
    <a id="L106"></a>}
<a id="L107"></a>}

<a id="L109"></a>func runPacket(t *testing.T, network, addr string, listening chan&lt;- string, done chan&lt;- int) {
    <a id="L110"></a>c, err := ListenPacket(network, addr);
    <a id="L111"></a>if err != nil {
        <a id="L112"></a>t.Fatalf(&#34;net.ListenPacket(%q, %q) = _, %v&#34;, network, addr, err)
    <a id="L113"></a>}
    <a id="L114"></a>listening &lt;- c.LocalAddr().String();
    <a id="L115"></a>c.SetReadTimeout(10e6); <span class="comment">// 10ms</span>
    <a id="L116"></a>var buf [1000]byte;
    <a id="L117"></a>for {
        <a id="L118"></a>n, addr, err := c.ReadFrom(&amp;buf);
        <a id="L119"></a>if err == os.EAGAIN {
            <a id="L120"></a>if done &lt;- 1 {
                <a id="L121"></a>break
            <a id="L122"></a>}
            <a id="L123"></a>continue;
        <a id="L124"></a>}
        <a id="L125"></a>if err != nil {
            <a id="L126"></a>break
        <a id="L127"></a>}
        <a id="L128"></a>if _, err = c.WriteTo(buf[0:n], addr); err != nil {
            <a id="L129"></a>t.Fatalf(&#34;WriteTo %v: %v&#34;, addr, err)
        <a id="L130"></a>}
    <a id="L131"></a>}
    <a id="L132"></a>c.Close();
    <a id="L133"></a>done &lt;- 1;
<a id="L134"></a>}

<a id="L136"></a>func doTestPacket(t *testing.T, network, listenaddr, dialaddr string) {
    <a id="L137"></a>t.Logf(&#34;TestPacket %s %s %s\n&#34;, network, listenaddr, dialaddr);
    <a id="L138"></a>listening := make(chan string);
    <a id="L139"></a>done := make(chan int);
    <a id="L140"></a>if network == &#34;udp&#34; {
        <a id="L141"></a>listenaddr += &#34;:0&#34; <span class="comment">// any available port</span>
    <a id="L142"></a>}
    <a id="L143"></a>go runPacket(t, network, listenaddr, listening, done);
    <a id="L144"></a>addr := &lt;-listening; <span class="comment">// wait for server to start</span>
    <a id="L145"></a>if network == &#34;udp&#34; {
        <a id="L146"></a>dialaddr += addr[strings.LastIndex(addr, &#34;:&#34;):len(addr)]
    <a id="L147"></a>}
    <a id="L148"></a>connect(t, network, dialaddr);
    <a id="L149"></a>&lt;-done; <span class="comment">// tell server to stop</span>
    <a id="L150"></a>&lt;-done; <span class="comment">// wait for stop</span>
<a id="L151"></a>}

<a id="L153"></a>func TestUDPServer(t *testing.T) {
    <a id="L154"></a>doTestPacket(t, &#34;udp&#34;, &#34;0.0.0.0&#34;, &#34;127.0.0.1&#34;);
    <a id="L155"></a>doTestPacket(t, &#34;udp&#34;, &#34;&#34;, &#34;127.0.0.1&#34;);
    <a id="L156"></a>if kernelSupportsIPv6() {
        <a id="L157"></a>doTestPacket(t, &#34;udp&#34;, &#34;[::]&#34;, &#34;[::ffff:127.0.0.1]&#34;);
        <a id="L158"></a>doTestPacket(t, &#34;udp&#34;, &#34;[::]&#34;, &#34;127.0.0.1&#34;);
        <a id="L159"></a>doTestPacket(t, &#34;udp&#34;, &#34;0.0.0.0&#34;, &#34;[::ffff:127.0.0.1]&#34;);
    <a id="L160"></a>}
<a id="L161"></a>}

<a id="L163"></a>func TestUnixDatagramServer(t *testing.T) {
    <a id="L164"></a>os.Remove(&#34;/tmp/gotest1.net&#34;);
    <a id="L165"></a>os.Remove(&#34;/tmp/gotest1.net.local&#34;);
    <a id="L166"></a>doTestPacket(t, &#34;unixgram&#34;, &#34;/tmp/gotest1.net&#34;, &#34;/tmp/gotest1.net&#34;);
    <a id="L167"></a>os.Remove(&#34;/tmp/gotest1.net&#34;);
    <a id="L168"></a>os.Remove(&#34;/tmp/gotest1.net.local&#34;);
    <a id="L169"></a>if syscall.OS == &#34;linux&#34; {
        <a id="L170"></a><span class="comment">// Test abstract unix domain socket, a Linux-ism</span>
        <a id="L171"></a>doTestPacket(t, &#34;unixgram&#34;, &#34;@gotest1/net&#34;, &#34;@gotest1/net&#34;)
    <a id="L172"></a>}
<a id="L173"></a>}
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
