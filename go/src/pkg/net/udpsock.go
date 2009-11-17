<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/udpsock.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/net/udpsock.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// UDP sockets</span>

<a id="L7"></a>package net

<a id="L9"></a>import (
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;syscall&#34;;
<a id="L12"></a>)

<a id="L14"></a>func sockaddrToUDP(sa syscall.Sockaddr) Addr {
    <a id="L15"></a>switch sa := sa.(type) {
    <a id="L16"></a>case *syscall.SockaddrInet4:
        <a id="L17"></a>return &amp;UDPAddr{&amp;sa.Addr, sa.Port}
    <a id="L18"></a>case *syscall.SockaddrInet6:
        <a id="L19"></a>return &amp;UDPAddr{&amp;sa.Addr, sa.Port}
    <a id="L20"></a>}
    <a id="L21"></a>return nil;
<a id="L22"></a>}

<a id="L24"></a><span class="comment">// UDPAddr represents the address of a UDP end point.</span>
<a id="L25"></a>type UDPAddr struct {
    <a id="L26"></a>IP   IP;
    <a id="L27"></a>Port int;
<a id="L28"></a>}

<a id="L30"></a><span class="comment">// Network returns the address&#39;s network name, &#34;udp&#34;.</span>
<a id="L31"></a>func (a *UDPAddr) Network() string { return &#34;udp&#34; }

<a id="L33"></a>func (a *UDPAddr) String() string { return joinHostPort(a.IP.String(), itoa(a.Port)) }

<a id="L35"></a>func (a *UDPAddr) family() int {
    <a id="L36"></a>if a == nil || len(a.IP) &lt;= 4 {
        <a id="L37"></a>return syscall.AF_INET
    <a id="L38"></a>}
    <a id="L39"></a>if ip := a.IP.To4(); ip != nil {
        <a id="L40"></a>return syscall.AF_INET
    <a id="L41"></a>}
    <a id="L42"></a>return syscall.AF_INET6;
<a id="L43"></a>}

<a id="L45"></a>func (a *UDPAddr) sockaddr(family int) (syscall.Sockaddr, os.Error) {
    <a id="L46"></a>return ipToSockaddr(family, a.IP, a.Port)
<a id="L47"></a>}

<a id="L49"></a>func (a *UDPAddr) toAddr() sockaddr {
    <a id="L50"></a>if a == nil { <span class="comment">// nil *UDPAddr</span>
        <a id="L51"></a>return nil <span class="comment">// nil interface</span>
    <a id="L52"></a>}
    <a id="L53"></a>return a;
<a id="L54"></a>}

<a id="L56"></a><span class="comment">// ResolveUDPAddr parses addr as a UDP address of the form</span>
<a id="L57"></a><span class="comment">// host:port and resolves domain names or port names to</span>
<a id="L58"></a><span class="comment">// numeric addresses.  A literal IPv6 host address must be</span>
<a id="L59"></a><span class="comment">// enclosed in square brackets, as in &#34;[::]:80&#34;.</span>
<a id="L60"></a>func ResolveUDPAddr(addr string) (*UDPAddr, os.Error) {
    <a id="L61"></a>ip, port, err := hostPortToIP(&#34;udp&#34;, addr);
    <a id="L62"></a>if err != nil {
        <a id="L63"></a>return nil, err
    <a id="L64"></a>}
    <a id="L65"></a>return &amp;UDPAddr{ip, port}, nil;
<a id="L66"></a>}

<a id="L68"></a><span class="comment">// UDPConn is the implementation of the Conn and PacketConn</span>
<a id="L69"></a><span class="comment">// interfaces for UDP network connections.</span>
<a id="L70"></a>type UDPConn struct {
    <a id="L71"></a>fd *netFD;
<a id="L72"></a>}

<a id="L74"></a>func newUDPConn(fd *netFD) *UDPConn { return &amp;UDPConn{fd} }

<a id="L76"></a>func (c *UDPConn) ok() bool { return c != nil &amp;&amp; c.fd != nil }

<a id="L78"></a><span class="comment">// Implementation of the Conn interface - see Conn for documentation.</span>

<a id="L80"></a><span class="comment">// Read reads data from a single UDP packet on the connection.</span>
<a id="L81"></a><span class="comment">// If the slice b is smaller than the arriving packet,</span>
<a id="L82"></a><span class="comment">// the excess packet data may be discarded.</span>
<a id="L83"></a><span class="comment">//</span>
<a id="L84"></a><span class="comment">// Read can be made to time out and return err == os.EAGAIN</span>
<a id="L85"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
<a id="L86"></a>func (c *UDPConn) Read(b []byte) (n int, err os.Error) {
    <a id="L87"></a>if !c.ok() {
        <a id="L88"></a>return 0, os.EINVAL
    <a id="L89"></a>}
    <a id="L90"></a>return c.fd.Read(b);
<a id="L91"></a>}

<a id="L93"></a><span class="comment">// Write writes data to the connection as a single UDP packet.</span>
<a id="L94"></a><span class="comment">//</span>
<a id="L95"></a><span class="comment">// Write can be made to time out and return err == os.EAGAIN</span>
<a id="L96"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
<a id="L97"></a>func (c *UDPConn) Write(b []byte) (n int, err os.Error) {
    <a id="L98"></a>if !c.ok() {
        <a id="L99"></a>return 0, os.EINVAL
    <a id="L100"></a>}
    <a id="L101"></a>return c.fd.Write(b);
<a id="L102"></a>}

<a id="L104"></a><span class="comment">// Close closes the UDP connection.</span>
<a id="L105"></a>func (c *UDPConn) Close() os.Error {
    <a id="L106"></a>if !c.ok() {
        <a id="L107"></a>return os.EINVAL
    <a id="L108"></a>}
    <a id="L109"></a>err := c.fd.Close();
    <a id="L110"></a>c.fd = nil;
    <a id="L111"></a>return err;
<a id="L112"></a>}

<a id="L114"></a><span class="comment">// LocalAddr returns the local network address.</span>
<a id="L115"></a>func (c *UDPConn) LocalAddr() Addr {
    <a id="L116"></a>if !c.ok() {
        <a id="L117"></a>return nil
    <a id="L118"></a>}
    <a id="L119"></a>return c.fd.laddr;
<a id="L120"></a>}

<a id="L122"></a><span class="comment">// RemoteAddr returns the remote network address, a *UDPAddr.</span>
<a id="L123"></a>func (c *UDPConn) RemoteAddr() Addr {
    <a id="L124"></a>if !c.ok() {
        <a id="L125"></a>return nil
    <a id="L126"></a>}
    <a id="L127"></a>return c.fd.raddr;
<a id="L128"></a>}

<a id="L130"></a><span class="comment">// SetTimeout sets the read and write deadlines associated</span>
<a id="L131"></a><span class="comment">// with the connection.</span>
<a id="L132"></a>func (c *UDPConn) SetTimeout(nsec int64) os.Error {
    <a id="L133"></a>if !c.ok() {
        <a id="L134"></a>return os.EINVAL
    <a id="L135"></a>}
    <a id="L136"></a>return setTimeout(c.fd, nsec);
<a id="L137"></a>}

<a id="L139"></a><span class="comment">// SetReadTimeout sets the time (in nanoseconds) that</span>
<a id="L140"></a><span class="comment">// Read will wait for data before returning os.EAGAIN.</span>
<a id="L141"></a><span class="comment">// Setting nsec == 0 (the default) disables the deadline.</span>
<a id="L142"></a>func (c *UDPConn) SetReadTimeout(nsec int64) os.Error {
    <a id="L143"></a>if !c.ok() {
        <a id="L144"></a>return os.EINVAL
    <a id="L145"></a>}
    <a id="L146"></a>return setReadTimeout(c.fd, nsec);
<a id="L147"></a>}

<a id="L149"></a><span class="comment">// SetWriteTimeout sets the time (in nanoseconds) that</span>
<a id="L150"></a><span class="comment">// Write will wait to send its data before returning os.EAGAIN.</span>
<a id="L151"></a><span class="comment">// Setting nsec == 0 (the default) disables the deadline.</span>
<a id="L152"></a><span class="comment">// Even if write times out, it may return n &gt; 0, indicating that</span>
<a id="L153"></a><span class="comment">// some of the data was successfully written.</span>
<a id="L154"></a>func (c *UDPConn) SetWriteTimeout(nsec int64) os.Error {
    <a id="L155"></a>if !c.ok() {
        <a id="L156"></a>return os.EINVAL
    <a id="L157"></a>}
    <a id="L158"></a>return setWriteTimeout(c.fd, nsec);
<a id="L159"></a>}

<a id="L161"></a><span class="comment">// SetReadBuffer sets the size of the operating system&#39;s</span>
<a id="L162"></a><span class="comment">// receive buffer associated with the connection.</span>
<a id="L163"></a>func (c *UDPConn) SetReadBuffer(bytes int) os.Error {
    <a id="L164"></a>if !c.ok() {
        <a id="L165"></a>return os.EINVAL
    <a id="L166"></a>}
    <a id="L167"></a>return setReadBuffer(c.fd, bytes);
<a id="L168"></a>}

<a id="L170"></a><span class="comment">// SetWriteBuffer sets the size of the operating system&#39;s</span>
<a id="L171"></a><span class="comment">// transmit buffer associated with the connection.</span>
<a id="L172"></a>func (c *UDPConn) SetWriteBuffer(bytes int) os.Error {
    <a id="L173"></a>if !c.ok() {
        <a id="L174"></a>return os.EINVAL
    <a id="L175"></a>}
    <a id="L176"></a>return setWriteBuffer(c.fd, bytes);
<a id="L177"></a>}

<a id="L179"></a><span class="comment">// UDP-specific methods.</span>

<a id="L181"></a><span class="comment">// ReadFromUDP reads a UDP packet from c, copying the payload into b.</span>
<a id="L182"></a><span class="comment">// It returns the number of bytes copied into b and the return address</span>
<a id="L183"></a><span class="comment">// that was on the packet.</span>
<a id="L184"></a><span class="comment">//</span>
<a id="L185"></a><span class="comment">// ReadFromUDP can be made to time out and return err == os.EAGAIN</span>
<a id="L186"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
<a id="L187"></a>func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err os.Error) {
    <a id="L188"></a>if !c.ok() {
        <a id="L189"></a>return 0, nil, os.EINVAL
    <a id="L190"></a>}
    <a id="L191"></a>n, sa, errno := syscall.Recvfrom(c.fd.fd, b, 0);
    <a id="L192"></a>if errno != 0 {
        <a id="L193"></a>err = os.Errno(errno)
    <a id="L194"></a>}
    <a id="L195"></a>switch sa := sa.(type) {
    <a id="L196"></a>case *syscall.SockaddrInet4:
        <a id="L197"></a>addr = &amp;UDPAddr{&amp;sa.Addr, sa.Port}
    <a id="L198"></a>case *syscall.SockaddrInet6:
        <a id="L199"></a>addr = &amp;UDPAddr{&amp;sa.Addr, sa.Port}
    <a id="L200"></a>}
    <a id="L201"></a>return;
<a id="L202"></a>}

<a id="L204"></a><span class="comment">// ReadFrom reads a UDP packet from c, copying the payload into b.</span>
<a id="L205"></a><span class="comment">// It returns the number of bytes copied into b and the return address</span>
<a id="L206"></a><span class="comment">// that was on the packet.</span>
<a id="L207"></a><span class="comment">//</span>
<a id="L208"></a><span class="comment">// ReadFrom can be made to time out and return err == os.EAGAIN</span>
<a id="L209"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
<a id="L210"></a>func (c *UDPConn) ReadFrom(b []byte) (n int, addr Addr, err os.Error) {
    <a id="L211"></a>if !c.ok() {
        <a id="L212"></a>return 0, nil, os.EINVAL
    <a id="L213"></a>}
    <a id="L214"></a>n, uaddr, err := c.ReadFromUDP(b);
    <a id="L215"></a>return n, uaddr.toAddr(), err;
<a id="L216"></a>}

<a id="L218"></a><span class="comment">// WriteToUDP writes a UDP packet to addr via c, copying the payload from b.</span>
<a id="L219"></a><span class="comment">//</span>
<a id="L220"></a><span class="comment">// WriteToUDP can be made to time out and return err == os.EAGAIN</span>
<a id="L221"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetWriteTimeout.</span>
<a id="L222"></a><span class="comment">// On packet-oriented connections such as UDP, write timeouts are rare.</span>
<a id="L223"></a>func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (n int, err os.Error) {
    <a id="L224"></a>if !c.ok() {
        <a id="L225"></a>return 0, os.EINVAL
    <a id="L226"></a>}
    <a id="L227"></a>sa, err := addr.sockaddr(c.fd.family);
    <a id="L228"></a>if err != nil {
        <a id="L229"></a>return 0, err
    <a id="L230"></a>}
    <a id="L231"></a>if errno := syscall.Sendto(c.fd.fd, b, 0, sa); errno != 0 {
        <a id="L232"></a>return 0, os.Errno(errno)
    <a id="L233"></a>}
    <a id="L234"></a>return len(b), nil;
<a id="L235"></a>}

<a id="L237"></a><span class="comment">// WriteTo writes a UDP packet with payload b to addr via c.</span>
<a id="L238"></a><span class="comment">//</span>
<a id="L239"></a><span class="comment">// WriteTo can be made to time out and return err == os.EAGAIN</span>
<a id="L240"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetWriteTimeout.</span>
<a id="L241"></a><span class="comment">// On packet-oriented connections such as UDP, write timeouts are rare.</span>
<a id="L242"></a>func (c *UDPConn) WriteTo(b []byte, addr Addr) (n int, err os.Error) {
    <a id="L243"></a>if !c.ok() {
        <a id="L244"></a>return 0, os.EINVAL
    <a id="L245"></a>}
    <a id="L246"></a>a, ok := addr.(*UDPAddr);
    <a id="L247"></a>if !ok {
        <a id="L248"></a>return 0, &amp;OpError{&#34;writeto&#34;, &#34;udp&#34;, addr, os.EINVAL}
    <a id="L249"></a>}
    <a id="L250"></a>return c.WriteToUDP(b, a);
<a id="L251"></a>}

<a id="L253"></a><span class="comment">// DialUDP connects to the remote address raddr on the network net,</span>
<a id="L254"></a><span class="comment">// which must be &#34;udp&#34;, &#34;udp4&#34;, or &#34;udp6&#34;.  If laddr is not nil, it is used</span>
<a id="L255"></a><span class="comment">// as the local address for the connection.</span>
<a id="L256"></a>func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err os.Error) {
    <a id="L257"></a>switch net {
    <a id="L258"></a>case &#34;udp&#34;, &#34;udp4&#34;, &#34;udp6&#34;:
    <a id="L259"></a>default:
        <a id="L260"></a>return nil, UnknownNetworkError(net)
    <a id="L261"></a>}
    <a id="L262"></a>if raddr == nil {
        <a id="L263"></a>return nil, &amp;OpError{&#34;dial&#34;, &#34;udp&#34;, nil, errMissingAddress}
    <a id="L264"></a>}
    <a id="L265"></a>fd, e := internetSocket(net, laddr.toAddr(), raddr.toAddr(), syscall.SOCK_DGRAM, &#34;dial&#34;, sockaddrToUDP);
    <a id="L266"></a>if e != nil {
        <a id="L267"></a>return nil, e
    <a id="L268"></a>}
    <a id="L269"></a>return newUDPConn(fd), nil;
<a id="L270"></a>}

<a id="L272"></a><span class="comment">// ListenUDP listens for incoming UDP packets addressed to the</span>
<a id="L273"></a><span class="comment">// local address laddr.  The returned connection c&#39;s ReadFrom</span>
<a id="L274"></a><span class="comment">// and WriteTo methods can be used to receive and send UDP</span>
<a id="L275"></a><span class="comment">// packets with per-packet addressing.</span>
<a id="L276"></a>func ListenUDP(net string, laddr *UDPAddr) (c *UDPConn, err os.Error) {
    <a id="L277"></a>switch net {
    <a id="L278"></a>case &#34;udp&#34;, &#34;udp4&#34;, &#34;udp6&#34;:
    <a id="L279"></a>default:
        <a id="L280"></a>return nil, UnknownNetworkError(net)
    <a id="L281"></a>}
    <a id="L282"></a>if laddr == nil {
        <a id="L283"></a>return nil, &amp;OpError{&#34;listen&#34;, &#34;udp&#34;, nil, errMissingAddress}
    <a id="L284"></a>}
    <a id="L285"></a>fd, e := internetSocket(net, laddr.toAddr(), nil, syscall.SOCK_DGRAM, &#34;dial&#34;, sockaddrToUDP);
    <a id="L286"></a>if e != nil {
        <a id="L287"></a>return nil, e
    <a id="L288"></a>}
    <a id="L289"></a>return newUDPConn(fd), nil;
<a id="L290"></a>}
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
