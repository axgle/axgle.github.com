<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/unixsock.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/net/unixsock.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Unix domain sockets</span>

<a id="L7"></a>package net

<a id="L9"></a>import (
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;syscall&#34;;
<a id="L12"></a>)

<a id="L14"></a>func unixSocket(net string, laddr, raddr *UnixAddr, mode string) (fd *netFD, err os.Error) {
    <a id="L15"></a>var proto int;
    <a id="L16"></a>switch net {
    <a id="L17"></a>default:
        <a id="L18"></a>return nil, UnknownNetworkError(net)
    <a id="L19"></a>case &#34;unix&#34;:
        <a id="L20"></a>proto = syscall.SOCK_STREAM
    <a id="L21"></a>case &#34;unixgram&#34;:
        <a id="L22"></a>proto = syscall.SOCK_DGRAM
    <a id="L23"></a>}

    <a id="L25"></a>var la, ra syscall.Sockaddr;
    <a id="L26"></a>switch mode {
    <a id="L27"></a>default:
        <a id="L28"></a>panic(&#34;unixSocket&#34;, mode)

    <a id="L30"></a>case &#34;dial&#34;:
        <a id="L31"></a>if laddr != nil {
            <a id="L32"></a>la = &amp;syscall.SockaddrUnix{Name: laddr.Name}
        <a id="L33"></a>}
        <a id="L34"></a>if raddr != nil {
            <a id="L35"></a>ra = &amp;syscall.SockaddrUnix{Name: raddr.Name}
        <a id="L36"></a>} else if proto != syscall.SOCK_DGRAM || laddr == nil {
            <a id="L37"></a>return nil, &amp;OpError{mode, net, nil, errMissingAddress}
        <a id="L38"></a>}

    <a id="L40"></a>case &#34;listen&#34;:
        <a id="L41"></a>if laddr == nil {
            <a id="L42"></a>return nil, &amp;OpError{mode, net, nil, errMissingAddress}
        <a id="L43"></a>}
        <a id="L44"></a>la = &amp;syscall.SockaddrUnix{Name: laddr.Name};
        <a id="L45"></a>if raddr != nil {
            <a id="L46"></a>return nil, &amp;OpError{mode, net, raddr, &amp;AddrError{&#34;unexpected remote address&#34;, raddr.String()}}
        <a id="L47"></a>}
    <a id="L48"></a>}

    <a id="L50"></a>f := sockaddrToUnix;
    <a id="L51"></a>if proto != syscall.SOCK_STREAM {
        <a id="L52"></a>f = sockaddrToUnixgram
    <a id="L53"></a>}
    <a id="L54"></a>fd, err = socket(net, syscall.AF_UNIX, proto, 0, la, ra, f);
    <a id="L55"></a>if err != nil {
        <a id="L56"></a>goto Error
    <a id="L57"></a>}
    <a id="L58"></a>return fd, nil;

<a id="L60"></a>Error:
    <a id="L61"></a>addr := raddr;
    <a id="L62"></a>if mode == &#34;listen&#34; {
        <a id="L63"></a>addr = laddr
    <a id="L64"></a>}
    <a id="L65"></a>return nil, &amp;OpError{mode, net, addr, err};
<a id="L66"></a>}

<a id="L68"></a><span class="comment">// UnixAddr represents the address of a Unix domain socket end point.</span>
<a id="L69"></a>type UnixAddr struct {
    <a id="L70"></a>Name     string;
    <a id="L71"></a>Datagram bool;
<a id="L72"></a>}

<a id="L74"></a>func sockaddrToUnix(sa syscall.Sockaddr) Addr {
    <a id="L75"></a>if s, ok := sa.(*syscall.SockaddrUnix); ok {
        <a id="L76"></a>return &amp;UnixAddr{s.Name, false}
    <a id="L77"></a>}
    <a id="L78"></a>return nil;
<a id="L79"></a>}

<a id="L81"></a>func sockaddrToUnixgram(sa syscall.Sockaddr) Addr {
    <a id="L82"></a>if s, ok := sa.(*syscall.SockaddrUnix); ok {
        <a id="L83"></a>return &amp;UnixAddr{s.Name, true}
    <a id="L84"></a>}
    <a id="L85"></a>return nil;
<a id="L86"></a>}

<a id="L88"></a><span class="comment">// Network returns the address&#39;s network name, &#34;unix&#34; or &#34;unixgram&#34;.</span>
<a id="L89"></a>func (a *UnixAddr) Network() string {
    <a id="L90"></a>if a == nil || !a.Datagram {
        <a id="L91"></a>return &#34;unix&#34;
    <a id="L92"></a>}
    <a id="L93"></a>return &#34;unixgram&#34;;
<a id="L94"></a>}

<a id="L96"></a>func (a *UnixAddr) String() string {
    <a id="L97"></a>if a == nil {
        <a id="L98"></a>return &#34;&lt;nil&gt;&#34;
    <a id="L99"></a>}
    <a id="L100"></a>return a.Name;
<a id="L101"></a>}

<a id="L103"></a>func (a *UnixAddr) toAddr() Addr {
    <a id="L104"></a>if a == nil { <span class="comment">// nil *UnixAddr</span>
        <a id="L105"></a>return nil <span class="comment">// nil interface</span>
    <a id="L106"></a>}
    <a id="L107"></a>return a;
<a id="L108"></a>}

<a id="L110"></a><span class="comment">// ResolveUnixAddr parses addr as a Unix domain socket address.</span>
<a id="L111"></a><span class="comment">// The string net gives the network name, &#34;unix&#34; or &#34;unixgram&#34;.</span>
<a id="L112"></a>func ResolveUnixAddr(net, addr string) (*UnixAddr, os.Error) {
    <a id="L113"></a>var datagram bool;
    <a id="L114"></a>switch net {
    <a id="L115"></a>case &#34;unix&#34;:
    <a id="L116"></a>case &#34;unixgram&#34;:
        <a id="L117"></a>datagram = true
    <a id="L118"></a>default:
        <a id="L119"></a>return nil, UnknownNetworkError(net)
    <a id="L120"></a>}
    <a id="L121"></a>return &amp;UnixAddr{addr, datagram}, nil;
<a id="L122"></a>}

<a id="L124"></a><span class="comment">// UnixConn is an implementation of the Conn interface</span>
<a id="L125"></a><span class="comment">// for connections to Unix domain sockets.</span>
<a id="L126"></a>type UnixConn struct {
    <a id="L127"></a>fd *netFD;
<a id="L128"></a>}

<a id="L130"></a>func newUnixConn(fd *netFD) *UnixConn { return &amp;UnixConn{fd} }

<a id="L132"></a>func (c *UnixConn) ok() bool { return c != nil &amp;&amp; c.fd != nil }

<a id="L134"></a><span class="comment">// Implementation of the Conn interface - see Conn for documentation.</span>

<a id="L136"></a><span class="comment">// Read reads data from the Unix domain connection.</span>
<a id="L137"></a><span class="comment">//</span>
<a id="L138"></a><span class="comment">// Read can be made to time out and return err == os.EAGAIN</span>
<a id="L139"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
<a id="L140"></a>func (c *UnixConn) Read(b []byte) (n int, err os.Error) {
    <a id="L141"></a>if !c.ok() {
        <a id="L142"></a>return 0, os.EINVAL
    <a id="L143"></a>}
    <a id="L144"></a>return c.fd.Read(b);
<a id="L145"></a>}

<a id="L147"></a><span class="comment">// Write writes data to the Unix domain connection.</span>
<a id="L148"></a><span class="comment">//</span>
<a id="L149"></a><span class="comment">// Write can be made to time out and return err == os.EAGAIN</span>
<a id="L150"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
<a id="L151"></a>func (c *UnixConn) Write(b []byte) (n int, err os.Error) {
    <a id="L152"></a>if !c.ok() {
        <a id="L153"></a>return 0, os.EINVAL
    <a id="L154"></a>}
    <a id="L155"></a>return c.fd.Write(b);
<a id="L156"></a>}

<a id="L158"></a><span class="comment">// Close closes the Unix domain connection.</span>
<a id="L159"></a>func (c *UnixConn) Close() os.Error {
    <a id="L160"></a>if !c.ok() {
        <a id="L161"></a>return os.EINVAL
    <a id="L162"></a>}
    <a id="L163"></a>err := c.fd.Close();
    <a id="L164"></a>c.fd = nil;
    <a id="L165"></a>return err;
<a id="L166"></a>}

<a id="L168"></a><span class="comment">// LocalAddr returns the local network address, a *UnixAddr.</span>
<a id="L169"></a><span class="comment">// Unlike in other protocols, LocalAddr is usually nil for dialed connections.</span>
<a id="L170"></a>func (c *UnixConn) LocalAddr() Addr {
    <a id="L171"></a>if !c.ok() {
        <a id="L172"></a>return nil
    <a id="L173"></a>}
    <a id="L174"></a>return c.fd.laddr;
<a id="L175"></a>}

<a id="L177"></a><span class="comment">// RemoteAddr returns the remote network address, a *UnixAddr.</span>
<a id="L178"></a><span class="comment">// Unlike in other protocols, RemoteAddr is usually nil for connections</span>
<a id="L179"></a><span class="comment">// accepted by a listener.</span>
<a id="L180"></a>func (c *UnixConn) RemoteAddr() Addr {
    <a id="L181"></a>if !c.ok() {
        <a id="L182"></a>return nil
    <a id="L183"></a>}
    <a id="L184"></a>return c.fd.raddr;
<a id="L185"></a>}

<a id="L187"></a><span class="comment">// SetTimeout sets the read and write deadlines associated</span>
<a id="L188"></a><span class="comment">// with the connection.</span>
<a id="L189"></a>func (c *UnixConn) SetTimeout(nsec int64) os.Error {
    <a id="L190"></a>if !c.ok() {
        <a id="L191"></a>return os.EINVAL
    <a id="L192"></a>}
    <a id="L193"></a>return setTimeout(c.fd, nsec);
<a id="L194"></a>}

<a id="L196"></a><span class="comment">// SetReadTimeout sets the time (in nanoseconds) that</span>
<a id="L197"></a><span class="comment">// Read will wait for data before returning os.EAGAIN.</span>
<a id="L198"></a><span class="comment">// Setting nsec == 0 (the default) disables the deadline.</span>
<a id="L199"></a>func (c *UnixConn) SetReadTimeout(nsec int64) os.Error {
    <a id="L200"></a>if !c.ok() {
        <a id="L201"></a>return os.EINVAL
    <a id="L202"></a>}
    <a id="L203"></a>return setReadTimeout(c.fd, nsec);
<a id="L204"></a>}

<a id="L206"></a><span class="comment">// SetWriteTimeout sets the time (in nanoseconds) that</span>
<a id="L207"></a><span class="comment">// Write will wait to send its data before returning os.EAGAIN.</span>
<a id="L208"></a><span class="comment">// Setting nsec == 0 (the default) disables the deadline.</span>
<a id="L209"></a><span class="comment">// Even if write times out, it may return n &gt; 0, indicating that</span>
<a id="L210"></a><span class="comment">// some of the data was successfully written.</span>
<a id="L211"></a>func (c *UnixConn) SetWriteTimeout(nsec int64) os.Error {
    <a id="L212"></a>if !c.ok() {
        <a id="L213"></a>return os.EINVAL
    <a id="L214"></a>}
    <a id="L215"></a>return setWriteTimeout(c.fd, nsec);
<a id="L216"></a>}

<a id="L218"></a><span class="comment">// SetReadBuffer sets the size of the operating system&#39;s</span>
<a id="L219"></a><span class="comment">// receive buffer associated with the connection.</span>
<a id="L220"></a>func (c *UnixConn) SetReadBuffer(bytes int) os.Error {
    <a id="L221"></a>if !c.ok() {
        <a id="L222"></a>return os.EINVAL
    <a id="L223"></a>}
    <a id="L224"></a>return setReadBuffer(c.fd, bytes);
<a id="L225"></a>}

<a id="L227"></a><span class="comment">// SetWriteBuffer sets the size of the operating system&#39;s</span>
<a id="L228"></a><span class="comment">// transmit buffer associated with the connection.</span>
<a id="L229"></a>func (c *UnixConn) SetWriteBuffer(bytes int) os.Error {
    <a id="L230"></a>if !c.ok() {
        <a id="L231"></a>return os.EINVAL
    <a id="L232"></a>}
    <a id="L233"></a>return setWriteBuffer(c.fd, bytes);
<a id="L234"></a>}

<a id="L236"></a><span class="comment">// ReadFromUnix reads a packet from c, copying the payload into b.</span>
<a id="L237"></a><span class="comment">// It returns the number of bytes copied into b and the return address</span>
<a id="L238"></a><span class="comment">// that was on the packet.</span>
<a id="L239"></a><span class="comment">//</span>
<a id="L240"></a><span class="comment">// ReadFromUnix can be made to time out and return err == os.EAGAIN</span>
<a id="L241"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
<a id="L242"></a>func (c *UnixConn) ReadFromUnix(b []byte) (n int, addr *UnixAddr, err os.Error) {
    <a id="L243"></a>if !c.ok() {
        <a id="L244"></a>return 0, nil, os.EINVAL
    <a id="L245"></a>}
    <a id="L246"></a>n, sa, errno := syscall.Recvfrom(c.fd.fd, b, 0);
    <a id="L247"></a>if errno != 0 {
        <a id="L248"></a>err = os.Errno(errno)
    <a id="L249"></a>}
    <a id="L250"></a>switch sa := sa.(type) {
    <a id="L251"></a>case *syscall.SockaddrUnix:
        <a id="L252"></a>addr = &amp;UnixAddr{sa.Name, c.fd.proto == syscall.SOCK_DGRAM}
    <a id="L253"></a>}
    <a id="L254"></a>return;
<a id="L255"></a>}

<a id="L257"></a><span class="comment">// ReadFrom reads a packet from c, copying the payload into b.</span>
<a id="L258"></a><span class="comment">// It returns the number of bytes copied into b and the return address</span>
<a id="L259"></a><span class="comment">// that was on the packet.</span>
<a id="L260"></a><span class="comment">//</span>
<a id="L261"></a><span class="comment">// ReadFrom can be made to time out and return err == os.EAGAIN</span>
<a id="L262"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
<a id="L263"></a>func (c *UnixConn) ReadFrom(b []byte) (n int, addr Addr, err os.Error) {
    <a id="L264"></a>if !c.ok() {
        <a id="L265"></a>return 0, nil, os.EINVAL
    <a id="L266"></a>}
    <a id="L267"></a>n, uaddr, err := c.ReadFromUnix(b);
    <a id="L268"></a>return n, uaddr.toAddr(), err;
<a id="L269"></a>}

<a id="L271"></a><span class="comment">// WriteToUnix writes a packet to addr via c, copying the payload from b.</span>
<a id="L272"></a><span class="comment">//</span>
<a id="L273"></a><span class="comment">// WriteToUnix can be made to time out and return err == os.EAGAIN</span>
<a id="L274"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetWriteTimeout.</span>
<a id="L275"></a><span class="comment">// On packet-oriented connections such as UDP, write timeouts are rare.</span>
<a id="L276"></a>func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (n int, err os.Error) {
    <a id="L277"></a>if !c.ok() {
        <a id="L278"></a>return 0, os.EINVAL
    <a id="L279"></a>}
    <a id="L280"></a>if addr.Datagram != (c.fd.proto == syscall.SOCK_DGRAM) {
        <a id="L281"></a>return 0, os.EAFNOSUPPORT
    <a id="L282"></a>}
    <a id="L283"></a>sa := &amp;syscall.SockaddrUnix{Name: addr.Name};
    <a id="L284"></a>if errno := syscall.Sendto(c.fd.fd, b, 0, sa); errno != 0 {
        <a id="L285"></a>return 0, os.Errno(errno)
    <a id="L286"></a>}
    <a id="L287"></a>return len(b), nil;
<a id="L288"></a>}

<a id="L290"></a><span class="comment">// WriteTo writes a packet to addr via c, copying the payload from b.</span>
<a id="L291"></a><span class="comment">//</span>
<a id="L292"></a><span class="comment">// WriteTo can be made to time out and return err == os.EAGAIN</span>
<a id="L293"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetWriteTimeout.</span>
<a id="L294"></a><span class="comment">// On packet-oriented connections such as UDP, write timeouts are rare.</span>
<a id="L295"></a>func (c *UnixConn) WriteTo(b []byte, addr Addr) (n int, err os.Error) {
    <a id="L296"></a>if !c.ok() {
        <a id="L297"></a>return 0, os.EINVAL
    <a id="L298"></a>}
    <a id="L299"></a>a, ok := addr.(*UnixAddr);
    <a id="L300"></a>if !ok {
        <a id="L301"></a>return 0, &amp;OpError{&#34;writeto&#34;, &#34;unix&#34;, addr, os.EINVAL}
    <a id="L302"></a>}
    <a id="L303"></a>return c.WriteToUnix(b, a);
<a id="L304"></a>}

<a id="L306"></a><span class="comment">// DialUDP connects to the remote address raddr on the network net,</span>
<a id="L307"></a><span class="comment">// which must be &#34;unix&#34; or &#34;unixdgram&#34;.  If laddr is not nil, it is used</span>
<a id="L308"></a><span class="comment">// as the local address for the connection.</span>
<a id="L309"></a>func DialUnix(net string, laddr, raddr *UnixAddr) (c *UnixConn, err os.Error) {
    <a id="L310"></a>fd, e := unixSocket(net, laddr, raddr, &#34;dial&#34;);
    <a id="L311"></a>if e != nil {
        <a id="L312"></a>return nil, e
    <a id="L313"></a>}
    <a id="L314"></a>return newUnixConn(fd), nil;
<a id="L315"></a>}

<a id="L317"></a><span class="comment">// UnixListener is a Unix domain socket listener.</span>
<a id="L318"></a><span class="comment">// Clients should typically use variables of type Listener</span>
<a id="L319"></a><span class="comment">// instead of assuming Unix domain sockets.</span>
<a id="L320"></a>type UnixListener struct {
    <a id="L321"></a>fd   *netFD;
    <a id="L322"></a>path string;
<a id="L323"></a>}

<a id="L325"></a><span class="comment">// ListenUnix announces on the Unix domain socket laddr and returns a Unix listener.</span>
<a id="L326"></a><span class="comment">// Net must be &#34;unix&#34; (stream sockets).</span>
<a id="L327"></a>func ListenUnix(net string, laddr *UnixAddr) (l *UnixListener, err os.Error) {
    <a id="L328"></a>if net != &#34;unix&#34; &amp;&amp; net != &#34;unixgram&#34; {
        <a id="L329"></a>return nil, UnknownNetworkError(net)
    <a id="L330"></a>}
    <a id="L331"></a>if laddr != nil {
        <a id="L332"></a>laddr = &amp;UnixAddr{laddr.Name, net == &#34;unixgram&#34;} <span class="comment">// make our own copy</span>
    <a id="L333"></a>}
    <a id="L334"></a>fd, e := unixSocket(net, laddr, nil, &#34;listen&#34;);
    <a id="L335"></a>if e != nil {
        <a id="L336"></a>if pe, ok := e.(*os.PathError); ok {
            <a id="L337"></a>e = pe.Error
        <a id="L338"></a>}
        <a id="L339"></a>return nil, e;
    <a id="L340"></a>}
    <a id="L341"></a>e1 := syscall.Listen(fd.fd, 8); <span class="comment">// listenBacklog());</span>
    <a id="L342"></a>if e1 != 0 {
        <a id="L343"></a>syscall.Close(fd.fd);
        <a id="L344"></a>return nil, &amp;OpError{&#34;listen&#34;, &#34;unix&#34;, laddr, os.Errno(e1)};
    <a id="L345"></a>}
    <a id="L346"></a>return &amp;UnixListener{fd, laddr.Name}, nil;
<a id="L347"></a>}

<a id="L349"></a><span class="comment">// AcceptUnix accepts the next incoming call and returns the new connection</span>
<a id="L350"></a><span class="comment">// and the remote address.</span>
<a id="L351"></a>func (l *UnixListener) AcceptUnix() (c *UnixConn, err os.Error) {
    <a id="L352"></a>if l == nil || l.fd == nil || l.fd.fd &lt; 0 {
        <a id="L353"></a>return nil, os.EINVAL
    <a id="L354"></a>}
    <a id="L355"></a>fd, e := l.fd.accept(sockaddrToUnix);
    <a id="L356"></a>if e != nil {
        <a id="L357"></a>return nil, e
    <a id="L358"></a>}
    <a id="L359"></a>c = newUnixConn(fd);
    <a id="L360"></a>return c, nil;
<a id="L361"></a>}

<a id="L363"></a><span class="comment">// Accept implements the Accept method in the Listener interface;</span>
<a id="L364"></a><span class="comment">// it waits for the next call and returns a generic Conn.</span>
<a id="L365"></a>func (l *UnixListener) Accept() (c Conn, err os.Error) {
    <a id="L366"></a>c1, err := l.AcceptUnix();
    <a id="L367"></a>if err != nil {
        <a id="L368"></a>return nil, err
    <a id="L369"></a>}
    <a id="L370"></a>return c1, nil;
<a id="L371"></a>}

<a id="L373"></a><span class="comment">// Close stops listening on the Unix address.</span>
<a id="L374"></a><span class="comment">// Already accepted connections are not closed.</span>
<a id="L375"></a>func (l *UnixListener) Close() os.Error {
    <a id="L376"></a>if l == nil || l.fd == nil {
        <a id="L377"></a>return os.EINVAL
    <a id="L378"></a>}

    <a id="L380"></a><span class="comment">// The operating system doesn&#39;t clean up</span>
    <a id="L381"></a><span class="comment">// the file that announcing created, so</span>
    <a id="L382"></a><span class="comment">// we have to clean it up ourselves.</span>
    <a id="L383"></a><span class="comment">// There&#39;s a race here--we can&#39;t know for</span>
    <a id="L384"></a><span class="comment">// sure whether someone else has come along</span>
    <a id="L385"></a><span class="comment">// and replaced our socket name already--</span>
    <a id="L386"></a><span class="comment">// but this sequence (remove then close)</span>
    <a id="L387"></a><span class="comment">// is at least compatible with the auto-remove</span>
    <a id="L388"></a><span class="comment">// sequence in ListenUnix.  It&#39;s only non-Go</span>
    <a id="L389"></a><span class="comment">// programs that can mess us up.</span>
    <a id="L390"></a>if l.path[0] != &#39;@&#39; {
        <a id="L391"></a>syscall.Unlink(l.path)
    <a id="L392"></a>}
    <a id="L393"></a>err := l.fd.Close();
    <a id="L394"></a>l.fd = nil;
    <a id="L395"></a>return err;
<a id="L396"></a>}

<a id="L398"></a><span class="comment">// Addr returns the listener&#39;s network address.</span>
<a id="L399"></a>func (l *UnixListener) Addr() Addr { return l.fd.laddr }

<a id="L401"></a><span class="comment">// ListenUnixgram listens for incoming Unix datagram packets addressed to the</span>
<a id="L402"></a><span class="comment">// local address laddr.  The returned connection c&#39;s ReadFrom</span>
<a id="L403"></a><span class="comment">// and WriteTo methods can be used to receive and send UDP</span>
<a id="L404"></a><span class="comment">// packets with per-packet addressing.  The network net must be &#34;unixgram&#34;.</span>
<a id="L405"></a>func ListenUnixgram(net string, laddr *UnixAddr) (c *UDPConn, err os.Error) {
    <a id="L406"></a>switch net {
    <a id="L407"></a>case &#34;unixgram&#34;:
    <a id="L408"></a>default:
        <a id="L409"></a>return nil, UnknownNetworkError(net)
    <a id="L410"></a>}
    <a id="L411"></a>if laddr == nil {
        <a id="L412"></a>return nil, &amp;OpError{&#34;listen&#34;, &#34;unixgram&#34;, nil, errMissingAddress}
    <a id="L413"></a>}
    <a id="L414"></a>fd, e := unixSocket(net, laddr, nil, &#34;listen&#34;);
    <a id="L415"></a>if e != nil {
        <a id="L416"></a>return nil, e
    <a id="L417"></a>}
    <a id="L418"></a>return newUDPConn(fd), nil;
<a id="L419"></a>}
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
