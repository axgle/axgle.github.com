<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/ipsock.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/net/ipsock.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// IP sockets</span>

<a id="L7"></a>package net

<a id="L9"></a>import (
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;syscall&#34;;
<a id="L12"></a>)

<a id="L14"></a><span class="comment">// Should we try to use the IPv4 socket interface if we&#39;re</span>
<a id="L15"></a><span class="comment">// only dealing with IPv4 sockets?  As long as the host system</span>
<a id="L16"></a><span class="comment">// understands IPv6, it&#39;s okay to pass IPv4 addresses to the IPv6</span>
<a id="L17"></a><span class="comment">// interface.  That simplifies our code and is most general.</span>
<a id="L18"></a><span class="comment">// Unfortunately, we need to run on kernels built without IPv6 support too.</span>
<a id="L19"></a><span class="comment">// So probe the kernel to figure it out.</span>
<a id="L20"></a>func kernelSupportsIPv6() bool {
    <a id="L21"></a>fd, e := syscall.Socket(syscall.AF_INET6, syscall.SOCK_STREAM, syscall.IPPROTO_TCP);
    <a id="L22"></a>if fd &gt;= 0 {
        <a id="L23"></a>syscall.Close(fd)
    <a id="L24"></a>}
    <a id="L25"></a>return e == 0;
<a id="L26"></a>}

<a id="L28"></a>var preferIPv4 = !kernelSupportsIPv6()

<a id="L30"></a><span class="comment">// TODO(rsc): if syscall.OS == &#34;linux&#34;, we&#39;re supposd to read</span>
<a id="L31"></a><span class="comment">// /proc/sys/net/core/somaxconn,</span>
<a id="L32"></a><span class="comment">// to take advantage of kernels that have raised the limit.</span>
<a id="L33"></a>func listenBacklog() int { return syscall.SOMAXCONN }

<a id="L35"></a><span class="comment">// Internet sockets (TCP, UDP)</span>

<a id="L37"></a><span class="comment">// A sockaddr represents a TCP or UDP network address that can</span>
<a id="L38"></a><span class="comment">// be converted into a syscall.Sockaddr.</span>
<a id="L39"></a>type sockaddr interface {
    <a id="L40"></a>Addr;
    <a id="L41"></a>sockaddr(family int) (syscall.Sockaddr, os.Error);
    <a id="L42"></a>family() int;
<a id="L43"></a>}

<a id="L45"></a>func internetSocket(net string, laddr, raddr sockaddr, proto int, mode string, toAddr func(syscall.Sockaddr) Addr) (fd *netFD, err os.Error) {
    <a id="L46"></a><span class="comment">// Figure out IP version.</span>
    <a id="L47"></a><span class="comment">// If network has a suffix like &#34;tcp4&#34;, obey it.</span>
    <a id="L48"></a>family := syscall.AF_INET6;
    <a id="L49"></a>switch net[len(net)-1] {
    <a id="L50"></a>case &#39;4&#39;:
        <a id="L51"></a>family = syscall.AF_INET
    <a id="L52"></a>case &#39;6&#39;:
        <a id="L53"></a><span class="comment">// nothing to do</span>
    <a id="L54"></a>default:
        <a id="L55"></a><span class="comment">// Otherwise, guess.</span>
        <a id="L56"></a><span class="comment">// If the addresses are IPv4 and we prefer IPv4, use 4; else 6.</span>
        <a id="L57"></a>if preferIPv4 &amp;&amp;
            <a id="L58"></a>(laddr == nil || laddr.family() == syscall.AF_INET) &amp;&amp;
            <a id="L59"></a>(raddr == nil || raddr.family() == syscall.AF_INET) {
            <a id="L60"></a>family = syscall.AF_INET
        <a id="L61"></a>}
    <a id="L62"></a>}

    <a id="L64"></a>var la, ra syscall.Sockaddr;
    <a id="L65"></a>if laddr != nil {
        <a id="L66"></a>if la, err = laddr.sockaddr(family); err != nil {
            <a id="L67"></a>goto Error
        <a id="L68"></a>}
    <a id="L69"></a>}
    <a id="L70"></a>if raddr != nil {
        <a id="L71"></a>if ra, err = raddr.sockaddr(family); err != nil {
            <a id="L72"></a>goto Error
        <a id="L73"></a>}
    <a id="L74"></a>}
    <a id="L75"></a>fd, err = socket(net, family, proto, 0, la, ra, toAddr);
    <a id="L76"></a>if err != nil {
        <a id="L77"></a>goto Error
    <a id="L78"></a>}
    <a id="L79"></a>return fd, nil;

<a id="L81"></a>Error:
    <a id="L82"></a>addr := raddr;
    <a id="L83"></a>if mode == &#34;listen&#34; {
        <a id="L84"></a>addr = laddr
    <a id="L85"></a>}
    <a id="L86"></a>return nil, &amp;OpError{mode, net, addr, err};
<a id="L87"></a>}

<a id="L89"></a>func getip(fd int, remote bool) (ip []byte, port int, ok bool) {
    <a id="L90"></a><span class="comment">// No attempt at error reporting because</span>
    <a id="L91"></a><span class="comment">// there are no possible errors, and the</span>
    <a id="L92"></a><span class="comment">// caller won&#39;t report them anyway.</span>
    <a id="L93"></a>var sa syscall.Sockaddr;
    <a id="L94"></a>if remote {
        <a id="L95"></a>sa, _ = syscall.Getpeername(fd)
    <a id="L96"></a>} else {
        <a id="L97"></a>sa, _ = syscall.Getsockname(fd)
    <a id="L98"></a>}
    <a id="L99"></a>switch sa := sa.(type) {
    <a id="L100"></a>case *syscall.SockaddrInet4:
        <a id="L101"></a>return &amp;sa.Addr, sa.Port, true
    <a id="L102"></a>case *syscall.SockaddrInet6:
        <a id="L103"></a>return &amp;sa.Addr, sa.Port, true
    <a id="L104"></a>}
    <a id="L105"></a>return;
<a id="L106"></a>}

<a id="L108"></a>func ipToSockaddr(family int, ip IP, port int) (syscall.Sockaddr, os.Error) {
    <a id="L109"></a>switch family {
    <a id="L110"></a>case syscall.AF_INET:
        <a id="L111"></a>if len(ip) == 0 {
            <a id="L112"></a>ip = IPv4zero
        <a id="L113"></a>}
        <a id="L114"></a>if ip = ip.To4(); ip == nil {
            <a id="L115"></a>return nil, os.EINVAL
        <a id="L116"></a>}
        <a id="L117"></a>s := new(syscall.SockaddrInet4);
        <a id="L118"></a>for i := 0; i &lt; IPv4len; i++ {
            <a id="L119"></a>s.Addr[i] = ip[i]
        <a id="L120"></a>}
        <a id="L121"></a>s.Port = port;
        <a id="L122"></a>return s, nil;
    <a id="L123"></a>case syscall.AF_INET6:
        <a id="L124"></a>if len(ip) == 0 {
            <a id="L125"></a>ip = IPzero
        <a id="L126"></a>}
        <a id="L127"></a><span class="comment">// IPv4 callers use 0.0.0.0 to mean &#34;announce on any available address&#34;.</span>
        <a id="L128"></a><span class="comment">// In IPv6 mode, Linux treats that as meaning &#34;announce on 0.0.0.0&#34;,</span>
        <a id="L129"></a><span class="comment">// which it refuses to do.  Rewrite to the IPv6 all zeros.</span>
        <a id="L130"></a>if p4 := ip.To4(); p4 != nil &amp;&amp; p4[0] == 0 &amp;&amp; p4[1] == 0 &amp;&amp; p4[2] == 0 &amp;&amp; p4[3] == 0 {
            <a id="L131"></a>ip = IPzero
        <a id="L132"></a>}
        <a id="L133"></a>if ip = ip.To16(); ip == nil {
            <a id="L134"></a>return nil, os.EINVAL
        <a id="L135"></a>}
        <a id="L136"></a>s := new(syscall.SockaddrInet6);
        <a id="L137"></a>for i := 0; i &lt; IPv6len; i++ {
            <a id="L138"></a>s.Addr[i] = ip[i]
        <a id="L139"></a>}
        <a id="L140"></a>s.Port = port;
        <a id="L141"></a>return s, nil;
    <a id="L142"></a>}
    <a id="L143"></a>return nil, os.EINVAL;
<a id="L144"></a>}

<a id="L146"></a><span class="comment">// Split &#34;host:port&#34; into &#34;host&#34; and &#34;port&#34;.</span>
<a id="L147"></a><span class="comment">// Host cannot contain colons unless it is bracketed.</span>
<a id="L148"></a>func splitHostPort(hostport string) (host, port string, err os.Error) {
    <a id="L149"></a><span class="comment">// The port starts after the last colon.</span>
    <a id="L150"></a>i := last(hostport, &#39;:&#39;);
    <a id="L151"></a>if i &lt; 0 {
        <a id="L152"></a>err = &amp;AddrError{&#34;missing port in address&#34;, hostport};
        <a id="L153"></a>return;
    <a id="L154"></a>}

    <a id="L156"></a>host, port = hostport[0:i], hostport[i+1:len(hostport)];

    <a id="L158"></a><span class="comment">// Can put brackets around host ...</span>
    <a id="L159"></a>if len(host) &gt; 0 &amp;&amp; host[0] == &#39;[&#39; &amp;&amp; host[len(host)-1] == &#39;]&#39; {
        <a id="L160"></a>host = host[1 : len(host)-1]
    <a id="L161"></a>} else {
        <a id="L162"></a><span class="comment">// ... but if there are no brackets, no colons.</span>
        <a id="L163"></a>if byteIndex(host, &#39;:&#39;) &gt;= 0 {
            <a id="L164"></a>err = &amp;AddrError{&#34;too many colons in address&#34;, hostport};
            <a id="L165"></a>return;
        <a id="L166"></a>}
    <a id="L167"></a>}
    <a id="L168"></a>return;
<a id="L169"></a>}

<a id="L171"></a><span class="comment">// Join &#34;host&#34; and &#34;port&#34; into &#34;host:port&#34;.</span>
<a id="L172"></a><span class="comment">// If host contains colons, will join into &#34;[host]:port&#34;.</span>
<a id="L173"></a>func joinHostPort(host, port string) string {
    <a id="L174"></a><span class="comment">// If host has colons, have to bracket it.</span>
    <a id="L175"></a>if byteIndex(host, &#39;:&#39;) &gt;= 0 {
        <a id="L176"></a>return &#34;[&#34; + host + &#34;]:&#34; + port
    <a id="L177"></a>}
    <a id="L178"></a>return host + &#34;:&#34; + port;
<a id="L179"></a>}

<a id="L181"></a><span class="comment">// Convert &#34;host:port&#34; into IP address and port.</span>
<a id="L182"></a>func hostPortToIP(net, hostport string) (ip IP, iport int, err os.Error) {
    <a id="L183"></a>host, port, err := splitHostPort(hostport);
    <a id="L184"></a>if err != nil {
        <a id="L185"></a>goto Error
    <a id="L186"></a>}

    <a id="L188"></a>var addr IP;
    <a id="L189"></a>if host != &#34;&#34; {
        <a id="L190"></a><span class="comment">// Try as an IP address.</span>
        <a id="L191"></a>addr = ParseIP(host);
        <a id="L192"></a>if addr == nil {
            <a id="L193"></a><span class="comment">// Not an IP address.  Try as a DNS name.</span>
            <a id="L194"></a>_, addrs, err1 := LookupHost(host);
            <a id="L195"></a>if err1 != nil {
                <a id="L196"></a>err = err1;
                <a id="L197"></a>goto Error;
            <a id="L198"></a>}
            <a id="L199"></a>addr = ParseIP(addrs[0]);
            <a id="L200"></a>if addr == nil {
                <a id="L201"></a><span class="comment">// should not happen</span>
                <a id="L202"></a>err = &amp;AddrError{&#34;LookupHost returned invalid address&#34;, addrs[0]};
                <a id="L203"></a>goto Error;
            <a id="L204"></a>}
        <a id="L205"></a>}
    <a id="L206"></a>}

    <a id="L208"></a>p, i, ok := dtoi(port, 0);
    <a id="L209"></a>if !ok || i != len(port) {
        <a id="L210"></a>p, err = LookupPort(net, port);
        <a id="L211"></a>if err != nil {
            <a id="L212"></a>goto Error
        <a id="L213"></a>}
    <a id="L214"></a>}
    <a id="L215"></a>if p &lt; 0 || p &gt; 0xFFFF {
        <a id="L216"></a>err = &amp;AddrError{&#34;invalid port&#34;, port};
        <a id="L217"></a>goto Error;
    <a id="L218"></a>}

    <a id="L220"></a>return addr, p, nil;

<a id="L222"></a>Error:
    <a id="L223"></a>return nil, 0, err;
<a id="L224"></a>}
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
