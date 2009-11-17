<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/net.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/net/net.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The net package provides a portable interface to Unix</span>
<a id="L6"></a><span class="comment">// networks sockets, including TCP/IP, UDP, domain name</span>
<a id="L7"></a><span class="comment">// resolution, and Unix domain sockets.</span>
<a id="L8"></a>package net

<a id="L10"></a><span class="comment">// TODO(rsc):</span>
<a id="L11"></a><span class="comment">//	support for raw IP sockets</span>
<a id="L12"></a><span class="comment">//	support for raw ethernet sockets</span>

<a id="L14"></a>import &#34;os&#34;

<a id="L16"></a><span class="comment">// Addr represents a network end point address.</span>
<a id="L17"></a>type Addr interface {
    <a id="L18"></a>Network() string; <span class="comment">// name of the network</span>
    <a id="L19"></a>String() string;  <span class="comment">// string form of address</span>
<a id="L20"></a>}

<a id="L22"></a><span class="comment">// Conn is a generic stream-oriented network connection.</span>
<a id="L23"></a>type Conn interface {
    <a id="L24"></a><span class="comment">// Read reads data from the connection.</span>
    <a id="L25"></a><span class="comment">// Read can be made to time out and return err == os.EAGAIN</span>
    <a id="L26"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
    <a id="L27"></a>Read(b []byte) (n int, err os.Error);

    <a id="L29"></a><span class="comment">// Write writes data to the connection.</span>
    <a id="L30"></a><span class="comment">// Write can be made to time out and return err == os.EAGAIN</span>
    <a id="L31"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
    <a id="L32"></a>Write(b []byte) (n int, err os.Error);

    <a id="L34"></a><span class="comment">// Close closes the connection.</span>
    <a id="L35"></a>Close() os.Error;

    <a id="L37"></a><span class="comment">// LocalAddr returns the local network address.</span>
    <a id="L38"></a>LocalAddr() Addr;

    <a id="L40"></a><span class="comment">// RemoteAddr returns the remote network address.</span>
    <a id="L41"></a>RemoteAddr() Addr;

    <a id="L43"></a><span class="comment">// SetTimeout sets the read and write deadlines associated</span>
    <a id="L44"></a><span class="comment">// with the connection.</span>
    <a id="L45"></a>SetTimeout(nsec int64) os.Error;

    <a id="L47"></a><span class="comment">// SetReadTimeout sets the time (in nanoseconds) that</span>
    <a id="L48"></a><span class="comment">// Read will wait for data before returning os.EAGAIN.</span>
    <a id="L49"></a><span class="comment">// Setting nsec == 0 (the default) disables the deadline.</span>
    <a id="L50"></a>SetReadTimeout(nsec int64) os.Error;

    <a id="L52"></a><span class="comment">// SetWriteTimeout sets the time (in nanoseconds) that</span>
    <a id="L53"></a><span class="comment">// Write will wait to send its data before returning os.EAGAIN.</span>
    <a id="L54"></a><span class="comment">// Setting nsec == 0 (the default) disables the deadline.</span>
    <a id="L55"></a><span class="comment">// Even if write times out, it may return n &gt; 0, indicating that</span>
    <a id="L56"></a><span class="comment">// some of the data was successfully written.</span>
    <a id="L57"></a>SetWriteTimeout(nsec int64) os.Error;
<a id="L58"></a>}

<a id="L60"></a><span class="comment">// PacketConn is a generic packet-oriented network connection.</span>
<a id="L61"></a>type PacketConn interface {
    <a id="L62"></a><span class="comment">// ReadFrom reads a packet from the connection,</span>
    <a id="L63"></a><span class="comment">// copying the payload into b.  It returns the number of</span>
    <a id="L64"></a><span class="comment">// bytes copied into b and the return address that</span>
    <a id="L65"></a><span class="comment">// was on the packet.</span>
    <a id="L66"></a><span class="comment">// ReadFrom can be made to time out and return err == os.EAGAIN</span>
    <a id="L67"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetReadTimeout.</span>
    <a id="L68"></a>ReadFrom(b []byte) (n int, addr Addr, err os.Error);

    <a id="L70"></a><span class="comment">// WriteTo writes a packet with payload b to addr.</span>
    <a id="L71"></a><span class="comment">// WriteTo can be made to time out and return err == os.EAGAIN</span>
    <a id="L72"></a><span class="comment">// after a fixed time limit; see SetTimeout and SetWriteTimeout.</span>
    <a id="L73"></a><span class="comment">// On packet-oriented connections, write timeouts are rare.</span>
    <a id="L74"></a>WriteTo(b []byte, addr Addr) (n int, err os.Error);

    <a id="L76"></a><span class="comment">// Close closes the connection.</span>
    <a id="L77"></a>Close() os.Error;

    <a id="L79"></a><span class="comment">// LocalAddr returns the local network address.</span>
    <a id="L80"></a>LocalAddr() Addr;

    <a id="L82"></a><span class="comment">// SetTimeout sets the read and write deadlines associated</span>
    <a id="L83"></a><span class="comment">// with the connection.</span>
    <a id="L84"></a>SetTimeout(nsec int64) os.Error;

    <a id="L86"></a><span class="comment">// SetReadTimeout sets the time (in nanoseconds) that</span>
    <a id="L87"></a><span class="comment">// Read will wait for data before returning os.EAGAIN.</span>
    <a id="L88"></a><span class="comment">// Setting nsec == 0 (the default) disables the deadline.</span>
    <a id="L89"></a>SetReadTimeout(nsec int64) os.Error;

    <a id="L91"></a><span class="comment">// SetWriteTimeout sets the time (in nanoseconds) that</span>
    <a id="L92"></a><span class="comment">// Write will wait to send its data before returning os.EAGAIN.</span>
    <a id="L93"></a><span class="comment">// Setting nsec == 0 (the default) disables the deadline.</span>
    <a id="L94"></a><span class="comment">// Even if write times out, it may return n &gt; 0, indicating that</span>
    <a id="L95"></a><span class="comment">// some of the data was successfully written.</span>
    <a id="L96"></a>SetWriteTimeout(nsec int64) os.Error;
<a id="L97"></a>}

<a id="L99"></a><span class="comment">// A Listener is a generic network listener for stream-oriented protocols.</span>
<a id="L100"></a><span class="comment">// Accept waits for the next connection and Close closes the connection.</span>
<a id="L101"></a>type Listener interface {
    <a id="L102"></a>Accept() (c Conn, err os.Error);
    <a id="L103"></a>Close() os.Error;
    <a id="L104"></a>Addr() Addr; <span class="comment">// Listener&#39;s network address</span>
<a id="L105"></a>}

<a id="L107"></a><span class="comment">// Dial connects to the remote address raddr on the network net.</span>
<a id="L108"></a><span class="comment">// If the string laddr is not empty, it is used as the local address</span>
<a id="L109"></a><span class="comment">// for the connection.</span>
<a id="L110"></a><span class="comment">//</span>
<a id="L111"></a><span class="comment">// Known networks are &#34;tcp&#34;, &#34;tcp4&#34; (IPv4-only), &#34;tcp6&#34; (IPv6-only),</span>
<a id="L112"></a><span class="comment">// &#34;udp&#34;, &#34;udp4&#34; (IPv4-only), and &#34;udp6&#34; (IPv6-only).</span>
<a id="L113"></a><span class="comment">//</span>
<a id="L114"></a><span class="comment">// For IP networks, addresses have the form host:port.  If host is</span>
<a id="L115"></a><span class="comment">// a literal IPv6 address, it must be enclosed in square brackets.</span>
<a id="L116"></a><span class="comment">//</span>
<a id="L117"></a><span class="comment">// Examples:</span>
<a id="L118"></a><span class="comment">//	Dial(&#34;tcp&#34;, &#34;&#34;, &#34;12.34.56.78:80&#34;)</span>
<a id="L119"></a><span class="comment">//	Dial(&#34;tcp&#34;, &#34;&#34;, &#34;google.com:80&#34;)</span>
<a id="L120"></a><span class="comment">//	Dial(&#34;tcp&#34;, &#34;&#34;, &#34;[de:ad:be:ef::ca:fe]:80&#34;)</span>
<a id="L121"></a><span class="comment">//	Dial(&#34;tcp&#34;, &#34;127.0.0.1:123&#34;, &#34;127.0.0.1:88&#34;)</span>
<a id="L122"></a><span class="comment">//</span>
<a id="L123"></a>func Dial(net, laddr, raddr string) (c Conn, err os.Error) {
    <a id="L124"></a>switch net {
    <a id="L125"></a>case &#34;tcp&#34;, &#34;tcp4&#34;, &#34;tcp6&#34;:
        <a id="L126"></a>var la, ra *TCPAddr;
        <a id="L127"></a>if laddr != &#34;&#34; {
            <a id="L128"></a>if la, err = ResolveTCPAddr(laddr); err != nil {
                <a id="L129"></a>goto Error
            <a id="L130"></a>}
        <a id="L131"></a>}
        <a id="L132"></a>if raddr != &#34;&#34; {
            <a id="L133"></a>if ra, err = ResolveTCPAddr(raddr); err != nil {
                <a id="L134"></a>goto Error
            <a id="L135"></a>}
        <a id="L136"></a>}
        <a id="L137"></a>return DialTCP(net, la, ra);
    <a id="L138"></a>case &#34;udp&#34;, &#34;udp4&#34;, &#34;upd6&#34;:
        <a id="L139"></a>var la, ra *UDPAddr;
        <a id="L140"></a>if laddr != &#34;&#34; {
            <a id="L141"></a>if la, err = ResolveUDPAddr(laddr); err != nil {
                <a id="L142"></a>goto Error
            <a id="L143"></a>}
        <a id="L144"></a>}
        <a id="L145"></a>if raddr != &#34;&#34; {
            <a id="L146"></a>if ra, err = ResolveUDPAddr(raddr); err != nil {
                <a id="L147"></a>goto Error
            <a id="L148"></a>}
        <a id="L149"></a>}
        <a id="L150"></a>return DialUDP(net, la, ra);
    <a id="L151"></a>case &#34;unix&#34;, &#34;unixgram&#34;:
        <a id="L152"></a>var la, ra *UnixAddr;
        <a id="L153"></a>if raddr != &#34;&#34; {
            <a id="L154"></a>if ra, err = ResolveUnixAddr(net, raddr); err != nil {
                <a id="L155"></a>goto Error
            <a id="L156"></a>}
        <a id="L157"></a>}
        <a id="L158"></a>if laddr != &#34;&#34; {
            <a id="L159"></a>if la, err = ResolveUnixAddr(net, laddr); err != nil {
                <a id="L160"></a>goto Error
            <a id="L161"></a>}
        <a id="L162"></a>}
        <a id="L163"></a>return DialUnix(net, la, ra);
    <a id="L164"></a>}
    <a id="L165"></a>err = UnknownNetworkError(net);
<a id="L166"></a>Error:
    <a id="L167"></a>return nil, &amp;OpError{&#34;dial&#34;, net + &#34; &#34; + raddr, nil, err};
<a id="L168"></a>}

<a id="L170"></a><span class="comment">// Listen announces on the local network address laddr.</span>
<a id="L171"></a><span class="comment">// The network string net must be a stream-oriented</span>
<a id="L172"></a><span class="comment">// network: &#34;tcp&#34;, &#34;tcp4&#34;, &#34;tcp6&#34;, or &#34;unix&#34;.</span>
<a id="L173"></a>func Listen(net, laddr string) (l Listener, err os.Error) {
    <a id="L174"></a>switch net {
    <a id="L175"></a>case &#34;tcp&#34;, &#34;tcp4&#34;, &#34;tcp6&#34;:
        <a id="L176"></a>var la *TCPAddr;
        <a id="L177"></a>if laddr != &#34;&#34; {
            <a id="L178"></a>if la, err = ResolveTCPAddr(laddr); err != nil {
                <a id="L179"></a>return nil, err
            <a id="L180"></a>}
        <a id="L181"></a>}
        <a id="L182"></a>l, err := ListenTCP(net, la);
        <a id="L183"></a>if err != nil {
            <a id="L184"></a>return nil, err
        <a id="L185"></a>}
        <a id="L186"></a>return l, nil;
    <a id="L187"></a>case &#34;unix&#34;:
        <a id="L188"></a>var la *UnixAddr;
        <a id="L189"></a>if laddr != &#34;&#34; {
            <a id="L190"></a>if la, err = ResolveUnixAddr(net, laddr); err != nil {
                <a id="L191"></a>return nil, err
            <a id="L192"></a>}
        <a id="L193"></a>}
        <a id="L194"></a>l, err := ListenUnix(net, la);
        <a id="L195"></a>if err != nil {
            <a id="L196"></a>return nil, err
        <a id="L197"></a>}
        <a id="L198"></a>return l, nil;
    <a id="L199"></a>}
    <a id="L200"></a>return nil, UnknownNetworkError(net);
<a id="L201"></a>}

<a id="L203"></a><span class="comment">// ListenPacket announces on the local network address laddr.</span>
<a id="L204"></a><span class="comment">// The network string net must be a packet-oriented network:</span>
<a id="L205"></a><span class="comment">// &#34;udp&#34;, &#34;udp4&#34;, &#34;udp6&#34;, or &#34;unixgram&#34;.</span>
<a id="L206"></a>func ListenPacket(net, laddr string) (c PacketConn, err os.Error) {
    <a id="L207"></a>switch net {
    <a id="L208"></a>case &#34;udp&#34;, &#34;udp4&#34;, &#34;udp6&#34;:
        <a id="L209"></a>var la *UDPAddr;
        <a id="L210"></a>if laddr != &#34;&#34; {
            <a id="L211"></a>if la, err = ResolveUDPAddr(laddr); err != nil {
                <a id="L212"></a>return nil, err
            <a id="L213"></a>}
        <a id="L214"></a>}
        <a id="L215"></a>c, err := ListenUDP(net, la);
        <a id="L216"></a>if err != nil {
            <a id="L217"></a>return nil, err
        <a id="L218"></a>}
        <a id="L219"></a>return c, nil;
    <a id="L220"></a>case &#34;unixgram&#34;:
        <a id="L221"></a>var la *UnixAddr;
        <a id="L222"></a>if laddr != &#34;&#34; {
            <a id="L223"></a>if la, err = ResolveUnixAddr(net, laddr); err != nil {
                <a id="L224"></a>return nil, err
            <a id="L225"></a>}
        <a id="L226"></a>}
        <a id="L227"></a>c, err := DialUnix(net, la, nil);
        <a id="L228"></a>if err != nil {
            <a id="L229"></a>return nil, err
        <a id="L230"></a>}
        <a id="L231"></a>return c, nil;
    <a id="L232"></a>}
    <a id="L233"></a>return nil, UnknownNetworkError(net);
<a id="L234"></a>}

<a id="L236"></a>var errMissingAddress = os.ErrorString(&#34;missing address&#34;)

<a id="L238"></a>type OpError struct {
    <a id="L239"></a>Op    string;
    <a id="L240"></a>Net   string;
    <a id="L241"></a>Addr  Addr;
    <a id="L242"></a>Error os.Error;
<a id="L243"></a>}

<a id="L245"></a>func (e *OpError) String() string {
    <a id="L246"></a>s := e.Op;
    <a id="L247"></a>if e.Net != &#34;&#34; {
        <a id="L248"></a>s += &#34; &#34; + e.Net
    <a id="L249"></a>}
    <a id="L250"></a>if e.Addr != nil {
        <a id="L251"></a>s += &#34; &#34; + e.Addr.String()
    <a id="L252"></a>}
    <a id="L253"></a>s += &#34;: &#34; + e.Error.String();
    <a id="L254"></a>return s;
<a id="L255"></a>}

<a id="L257"></a>type AddrError struct {
    <a id="L258"></a>Error string;
    <a id="L259"></a>Addr  string;
<a id="L260"></a>}

<a id="L262"></a>func (e *AddrError) String() string {
    <a id="L263"></a>s := e.Error;
    <a id="L264"></a>if e.Addr != &#34;&#34; {
        <a id="L265"></a>s += &#34; &#34; + e.Addr
    <a id="L266"></a>}
    <a id="L267"></a>return s;
<a id="L268"></a>}

<a id="L270"></a>type UnknownNetworkError string

<a id="L272"></a>func (e UnknownNetworkError) String() string { return &#34;unknown network &#34; + string(e) }
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
