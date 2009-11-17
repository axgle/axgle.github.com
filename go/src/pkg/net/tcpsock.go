<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/tcpsock.go</title>

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
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/net/tcpsock.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// TCP sockets</span>

<a id="L7"></a>package net

<a id="L9"></a>import (
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;syscall&#34;;
<a id="L12"></a>)

<a id="L14"></a>func sockaddrToTCP(sa syscall.Sockaddr) Addr {
    <a id="L15"></a>switch sa := sa.(type) {
    <a id="L16"></a>case *syscall.SockaddrInet4:
        <a id="L17"></a>return &amp;TCPAddr{&amp;sa.Addr, sa.Port}
    <a id="L18"></a>case *syscall.SockaddrInet6:
        <a id="L19"></a>return &amp;TCPAddr{&amp;sa.Addr, sa.Port}
    <a id="L20"></a>}
    <a id="L21"></a>return nil;
<a id="L22"></a>}

<a id="L24"></a><span class="comment">// TCPAddr represents the address of a TCP end point.</span>
<a id="L25"></a>type TCPAddr struct {
    <a id="L26"></a>IP   IP;
    <a id="L27"></a>Port int;
<a id="L28"></a>}

<a id="L30"></a><span class="comment">// Network returns the address&#39;s network name, &#34;tcp&#34;.</span>
<a id="L31"></a>func (a *TCPAddr) Network() string { return &#34;tcp&#34; }

<a id="L33"></a>func (a *TCPAddr) String() string { return joinHostPort(a.IP.String(), itoa(a.Port)) }

<a id="L35"></a>func (a *TCPAddr) family() int {
    <a id="L36"></a>if a == nil || len(a.IP) &lt;= 4 {
        <a id="L37"></a>return syscall.AF_INET
    <a id="L38"></a>}
    <a id="L39"></a>if ip := a.IP.To4(); ip != nil {
        <a id="L40"></a>return syscall.AF_INET
    <a id="L41"></a>}
    <a id="L42"></a>return syscall.AF_INET6;
<a id="L43"></a>}

<a id="L45"></a>func (a *TCPAddr) sockaddr(family int) (syscall.Sockaddr, os.Error) {
    <a id="L46"></a>return ipToSockaddr(family, a.IP, a.Port)
<a id="L47"></a>}

<a id="L49"></a>func (a *TCPAddr) toAddr() sockaddr {
    <a id="L50"></a>if a == nil { <span class="comment">// nil *TCPAddr</span>
        <a id="L51"></a>return nil <span class="comment">// nil interface</span>
    <a id="L52"></a>}
    <a id="L53"></a>return a;
<a id="L54"></a>}

<a id="L56"></a><span class="comment">// ResolveTCPAddr parses addr as a TCP address of the form</span>
<a id="L57"></a><span class="comment">// host:port and resolves domain names or port names to</span>
<a id="L58"></a><span class="comment">// numeric addresses.  A literal IPv6 host address must be</span>
<a id="L59"></a><span class="comment">// enclosed in square brackets, as in &#34;[::]:80&#34;.</span>
<a id="L60"></a>func ResolveTCPAddr(addr string) (*TCPAddr, os.Error) {
    <a id="L61"></a>ip, port, err := hostPortToIP(&#34;tcp&#34;, addr);
    <a id="L62"></a>if err != nil {
        <a id="L63"></a>return nil, err
    <a id="L64"></a>}
    <a id="L65"></a>return &amp;TCPAddr{ip, port}, nil;
<a id="L66"></a>}

<a id="L68"></a><span class="comment">// TCPConn is an implementation of the Conn interface</span>
<a id="L69"></a><span class="comment">// for TCP network connections.</span>
<a id="L70"></a>type TCPConn struct {
    <a id="L71"></a>fd *netFD;
<a id="L72"></a>}

<a id="L74"></a>func newTCPConn(fd *netFD) *TCPConn {
    <a id="L75"></a>c := &amp;TCPConn{fd};
    <a id="L76"></a>setsockoptInt(fd.fd, syscall.IPPROTO_TCP, syscall.TCP_NODELAY, 1);
    <a id="L77"></a>return c;
<a id="L78"></a>}

<a id="L80"></a>func (c *TCPConn) ok() bool { return c != nil &amp;&amp; c.fd != nil }

<a id="L82"></a><span class="comment">// Implementation of the Conn interface - see Conn for documentation.</span>

<a id="L84"></a><span class="comment">// Read reads data from the TCP connection.</span>
<a id="L85"></a><span class="comment">//</span>
<a id="L86"></a><span class="comment">// Read can be made to time out and return err == os.EAGAIN</span>
<a id="L87"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
<a id="L88"></a>func (c *TCPConn) Read(b []byte) (n int, err os.Error) {
    <a id="L89"></a>if !c.ok() {
        <a id="L90"></a>return 0, os.EINVAL
    <a id="L91"></a>}
    <a id="L92"></a>return c.fd.Read(b);
<a id="L93"></a>}

<a id="L95"></a><span class="comment">// Write writes data to the TCP connection.</span>
<a id="L96"></a><span class="comment">//</span>
<a id="L97"></a><span class="comment">// Write can be made to time out and return err == os.EAGAIN</span>
<a id="L98"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
<a id="L99"></a>func (c *TCPConn) Write(b []byte) (n int, err os.Error) {
    <a id="L100"></a>if !c.ok() {
        <a id="L101"></a>return 0, os.EINVAL
    <a id="L102"></a>}
    <a id="L103"></a>return c.fd.Write(b);
<a id="L104"></a>}

<a id="L106"></a><span class="comment">// Close closes the TCP connection.</span>
<a id="L107"></a>func (c *TCPConn) Close() os.Error {
    <a id="L108"></a>if !c.ok() {
        <a id="L109"></a>return os.EINVAL
    <a id="L110"></a>}
    <a id="L111"></a>err := c.fd.Close();
    <a id="L112"></a>c.fd = nil;
    <a id="L113"></a>return err;
<a id="L114"></a>}

<a id="L116"></a><span class="comment">// LocalAddr returns the local network address, a *TCPAddr.</span>
<a id="L117"></a>func (c *TCPConn) LocalAddr() Addr {
    <a id="L118"></a>if !c.ok() {
        <a id="L119"></a>return nil
    <a id="L120"></a>}
    <a id="L121"></a>return c.fd.laddr;
<a id="L122"></a>}

<a id="L124"></a><span class="comment">// RemoteAddr returns the remote network address, a *TCPAddr.</span>
<a id="L125"></a>func (c *TCPConn) RemoteAddr() Addr {
    <a id="L126"></a>if !c.ok() {
        <a id="L127"></a>return nil
    <a id="L128"></a>}
    <a id="L129"></a>return c.fd.raddr;
<a id="L130"></a>}

<a id="L132"></a><span class="comment">// SetTimeout sets the read and write deadlines associated</span>
<a id="L133"></a><span class="comment">// with the connection.</span>
<a id="L134"></a>func (c *TCPConn) SetTimeout(nsec int64) os.Error {
    <a id="L135"></a>if !c.ok() {
        <a id="L136"></a>return os.EINVAL
    <a id="L137"></a>}
    <a id="L138"></a>return setTimeout(c.fd, nsec);
<a id="L139"></a>}

<a id="L141"></a><span class="comment">// SetReadTimeout sets the time (in nanoseconds) that</span>
<a id="L142"></a><span class="comment">// Read will wait for data before returning os.EAGAIN.</span>
<a id="L143"></a><span class="comment">// Setting nsec == 0 (the default) disables the deadline.</span>
<a id="L144"></a>func (c *TCPConn) SetReadTimeout(nsec int64) os.Error {
    <a id="L145"></a>if !c.ok() {
        <a id="L146"></a>return os.EINVAL
    <a id="L147"></a>}
    <a id="L148"></a>return setReadTimeout(c.fd, nsec);
<a id="L149"></a>}

<a id="L151"></a><span class="comment">// SetWriteTimeout sets the time (in nanoseconds) that</span>
<a id="L152"></a><span class="comment">// Write will wait to send its data before returning os.EAGAIN.</span>
<a id="L153"></a><span class="comment">// Setting nsec == 0 (the default) disables the deadline.</span>
<a id="L154"></a><span class="comment">// Even if write times out, it may return n &gt; 0, indicating that</span>
<a id="L155"></a><span class="comment">// some of the data was successfully written.</span>
<a id="L156"></a>func (c *TCPConn) SetWriteTimeout(nsec int64) os.Error {
    <a id="L157"></a>if !c.ok() {
        <a id="L158"></a>return os.EINVAL
    <a id="L159"></a>}
    <a id="L160"></a>return setWriteTimeout(c.fd, nsec);
<a id="L161"></a>}

<a id="L163"></a><span class="comment">// SetReadBuffer sets the size of the operating system&#39;s</span>
<a id="L164"></a><span class="comment">// receive buffer associated with the connection.</span>
<a id="L165"></a>func (c *TCPConn) SetReadBuffer(bytes int) os.Error {
    <a id="L166"></a>if !c.ok() {
        <a id="L167"></a>return os.EINVAL
    <a id="L168"></a>}
    <a id="L169"></a>return setReadBuffer(c.fd, bytes);
<a id="L170"></a>}

<a id="L172"></a><span class="comment">// SetWriteBuffer sets the size of the operating system&#39;s</span>
<a id="L173"></a><span class="comment">// transmit buffer associated with the connection.</span>
<a id="L174"></a>func (c *TCPConn) SetWriteBuffer(bytes int) os.Error {
    <a id="L175"></a>if !c.ok() {
        <a id="L176"></a>return os.EINVAL
    <a id="L177"></a>}
    <a id="L178"></a>return setWriteBuffer(c.fd, bytes);
<a id="L179"></a>}

<a id="L181"></a><span class="comment">// SetLinger sets the behavior of Close() on a connection</span>
<a id="L182"></a><span class="comment">// which still has data waiting to be sent or to be acknowledged.</span>
<a id="L183"></a><span class="comment">//</span>
<a id="L184"></a><span class="comment">// If sec &lt; 0 (the default), Close returns immediately and</span>
<a id="L185"></a><span class="comment">// the operating system finishes sending the data in the background.</span>
<a id="L186"></a><span class="comment">//</span>
<a id="L187"></a><span class="comment">// If sec == 0, Close returns immediately and the operating system</span>
<a id="L188"></a><span class="comment">// discards any unsent or unacknowledged data.</span>
<a id="L189"></a><span class="comment">//</span>
<a id="L190"></a><span class="comment">// If sec &gt; 0, Close blocks for at most sec seconds waiting for</span>
<a id="L191"></a><span class="comment">// data to be sent and acknowledged.</span>
<a id="L192"></a>func (c *TCPConn) SetLinger(sec int) os.Error {
    <a id="L193"></a>if !c.ok() {
        <a id="L194"></a>return os.EINVAL
    <a id="L195"></a>}
    <a id="L196"></a>return setLinger(c.fd, sec);
<a id="L197"></a>}

<a id="L199"></a><span class="comment">// SetKeepAlive sets whether the operating system should send</span>
<a id="L200"></a><span class="comment">// keepalive messages on the connection.</span>
<a id="L201"></a>func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error {
    <a id="L202"></a>if !c.ok() {
        <a id="L203"></a>return os.EINVAL
    <a id="L204"></a>}
    <a id="L205"></a>return setKeepAlive(c.fd, keepalive);
<a id="L206"></a>}

<a id="L208"></a><span class="comment">// DialTCP is like Dial but can only connect to TCP networks</span>
<a id="L209"></a><span class="comment">// and returns a TCPConn structure.</span>
<a id="L210"></a>func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error) {
    <a id="L211"></a>if raddr == nil {
        <a id="L212"></a>return nil, &amp;OpError{&#34;dial&#34;, &#34;tcp&#34;, nil, errMissingAddress}
    <a id="L213"></a>}
    <a id="L214"></a>fd, e := internetSocket(net, laddr.toAddr(), raddr.toAddr(), syscall.SOCK_STREAM, &#34;dial&#34;, sockaddrToTCP);
    <a id="L215"></a>if e != nil {
        <a id="L216"></a>return nil, e
    <a id="L217"></a>}
    <a id="L218"></a>return newTCPConn(fd), nil;
<a id="L219"></a>}

<a id="L221"></a><span class="comment">// TCPListener is a TCP network listener.</span>
<a id="L222"></a><span class="comment">// Clients should typically use variables of type Listener</span>
<a id="L223"></a><span class="comment">// instead of assuming TCP.</span>
<a id="L224"></a>type TCPListener struct {
    <a id="L225"></a>fd *netFD;
<a id="L226"></a>}

<a id="L228"></a><span class="comment">// ListenTCP announces on the TCP address laddr and returns a TCP listener.</span>
<a id="L229"></a><span class="comment">// Net must be &#34;tcp&#34;, &#34;tcp4&#34;, or &#34;tcp6&#34;.</span>
<a id="L230"></a><span class="comment">// If laddr has a port of 0, it means to listen on some available port.</span>
<a id="L231"></a><span class="comment">// The caller can use l.Addr() to retrieve the chosen address.</span>
<a id="L232"></a>func ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error) {
    <a id="L233"></a>fd, err := internetSocket(net, laddr.toAddr(), nil, syscall.SOCK_STREAM, &#34;listen&#34;, sockaddrToTCP);
    <a id="L234"></a>if err != nil {
        <a id="L235"></a>return nil, err
    <a id="L236"></a>}
    <a id="L237"></a>errno := syscall.Listen(fd.fd, listenBacklog());
    <a id="L238"></a>if errno != 0 {
        <a id="L239"></a>syscall.Close(fd.fd);
        <a id="L240"></a>return nil, &amp;OpError{&#34;listen&#34;, &#34;tcp&#34;, laddr, os.Errno(errno)};
    <a id="L241"></a>}
    <a id="L242"></a>l = new(TCPListener);
    <a id="L243"></a>l.fd = fd;
    <a id="L244"></a>return l, nil;
<a id="L245"></a>}

<a id="L247"></a><span class="comment">// AcceptTCP accepts the next incoming call and returns the new connection</span>
<a id="L248"></a><span class="comment">// and the remote address.</span>
<a id="L249"></a>func (l *TCPListener) AcceptTCP() (c *TCPConn, err os.Error) {
    <a id="L250"></a>if l == nil || l.fd == nil || l.fd.fd &lt; 0 {
        <a id="L251"></a>return nil, os.EINVAL
    <a id="L252"></a>}
    <a id="L253"></a>fd, err := l.fd.accept(sockaddrToTCP);
    <a id="L254"></a>if err != nil {
        <a id="L255"></a>return nil, err
    <a id="L256"></a>}
    <a id="L257"></a>return newTCPConn(fd), nil;
<a id="L258"></a>}

<a id="L260"></a><span class="comment">// Accept implements the Accept method in the Listener interface;</span>
<a id="L261"></a><span class="comment">// it waits for the next call and returns a generic Conn.</span>
<a id="L262"></a>func (l *TCPListener) Accept() (c Conn, err os.Error) {
    <a id="L263"></a>c1, err := l.AcceptTCP();
    <a id="L264"></a>if err != nil {
        <a id="L265"></a>return nil, err
    <a id="L266"></a>}
    <a id="L267"></a>return c1, nil;
<a id="L268"></a>}

<a id="L270"></a><span class="comment">// Close stops listening on the TCP address.</span>
<a id="L271"></a><span class="comment">// Already Accepted connections are not closed.</span>
<a id="L272"></a>func (l *TCPListener) Close() os.Error {
    <a id="L273"></a>if l == nil || l.fd == nil {
        <a id="L274"></a>return os.EINVAL
    <a id="L275"></a>}
    <a id="L276"></a>return l.fd.Close();
<a id="L277"></a>}

<a id="L279"></a><span class="comment">// Addr returns the listener&#39;s network address, a *TCPAddr.</span>
<a id="L280"></a>func (l *TCPListener) Addr() Addr { return l.fd.laddr }
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
