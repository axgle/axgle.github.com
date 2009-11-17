<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/dialgoogle_test.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/net/dialgoogle_test.go</h1>

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
    <a id="L8"></a>&#34;flag&#34;;
    <a id="L9"></a>&#34;io&#34;;
    <a id="L10"></a>&#34;strings&#34;;
    <a id="L11"></a>&#34;syscall&#34;;
    <a id="L12"></a>&#34;testing&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">// If an IPv6 tunnel is running (see go/stubl), we can try dialing a real IPv6 address.</span>
<a id="L16"></a>var ipv6 = flag.Bool(&#34;ipv6&#34;, false, &#34;assume ipv6 tunnel is present&#34;)

<a id="L18"></a><span class="comment">// fd is already connected to the destination, port 80.</span>
<a id="L19"></a><span class="comment">// Run an HTTP request to fetch the appropriate page.</span>
<a id="L20"></a>func fetchGoogle(t *testing.T, fd Conn, network, addr string) {
    <a id="L21"></a>req := strings.Bytes(&#34;GET /intl/en/privacy.html HTTP/1.0\r\nHost: www.google.com\r\n\r\n&#34;);
    <a id="L22"></a>n, err := fd.Write(req);

    <a id="L24"></a>buf := make([]byte, 1000);
    <a id="L25"></a>n, err = io.ReadFull(fd, buf);

    <a id="L27"></a>if n &lt; 1000 {
        <a id="L28"></a>t.Errorf(&#34;fetchGoogle: short HTTP read from %s %s - %v&#34;, network, addr, err);
        <a id="L29"></a>return;
    <a id="L30"></a>}
<a id="L31"></a>}

<a id="L33"></a>func doDial(t *testing.T, network, addr string) {
    <a id="L34"></a>fd, err := Dial(network, &#34;&#34;, addr);
    <a id="L35"></a>if err != nil {
        <a id="L36"></a>t.Errorf(&#34;Dial(%q, %q, %q) = _, %v&#34;, network, &#34;&#34;, addr, err);
        <a id="L37"></a>return;
    <a id="L38"></a>}
    <a id="L39"></a>fetchGoogle(t, fd, network, addr);
    <a id="L40"></a>fd.Close();
<a id="L41"></a>}

<a id="L43"></a>var googleaddrs = []string{
    <a id="L44"></a>&#34;74.125.19.99:80&#34;,
    <a id="L45"></a>&#34;www.google.com:80&#34;,
    <a id="L46"></a>&#34;74.125.19.99:http&#34;,
    <a id="L47"></a>&#34;www.google.com:http&#34;,
    <a id="L48"></a>&#34;074.125.019.099:0080&#34;,
    <a id="L49"></a>&#34;[::ffff:74.125.19.99]:80&#34;,
    <a id="L50"></a>&#34;[::ffff:4a7d:1363]:80&#34;,
    <a id="L51"></a>&#34;[0:0:0:0:0000:ffff:74.125.19.99]:80&#34;,
    <a id="L52"></a>&#34;[0:0:0:0:000000:ffff:74.125.19.99]:80&#34;,
    <a id="L53"></a>&#34;[0:0:0:0:0:ffff::74.125.19.99]:80&#34;,
    <a id="L54"></a>&#34;[2001:4860:0:2001::68]:80&#34; <span class="comment">// ipv6.google.com; removed if ipv6 flag not set</span>
    ,
<a id="L56"></a>}

<a id="L58"></a>func TestDialGoogle(t *testing.T) {
    <a id="L59"></a><span class="comment">// If no ipv6 tunnel, don&#39;t try the last address.</span>
    <a id="L60"></a>if !*ipv6 {
        <a id="L61"></a>googleaddrs[len(googleaddrs)-1] = &#34;&#34;
    <a id="L62"></a>}

    <a id="L64"></a>for i := 0; i &lt; len(googleaddrs); i++ {
        <a id="L65"></a>addr := googleaddrs[i];
        <a id="L66"></a>if addr == &#34;&#34; {
            <a id="L67"></a>continue
        <a id="L68"></a>}
        <a id="L69"></a>t.Logf(&#34;-- %s --&#34;, addr);
        <a id="L70"></a>doDial(t, &#34;tcp&#34;, addr);
        <a id="L71"></a>if addr[0] != &#39;[&#39; {
            <a id="L72"></a>doDial(t, &#34;tcp4&#34;, addr);

            <a id="L74"></a>if !preferIPv4 {
                <a id="L75"></a><span class="comment">// make sure preferIPv4 flag works.</span>
                <a id="L76"></a>preferIPv4 = true;
                <a id="L77"></a>syscall.SocketDisableIPv6 = true;
                <a id="L78"></a>doDial(t, &#34;tcp4&#34;, addr);
                <a id="L79"></a>syscall.SocketDisableIPv6 = false;
                <a id="L80"></a>preferIPv4 = false;
            <a id="L81"></a>}
        <a id="L82"></a>}

        <a id="L84"></a><span class="comment">// Only run tcp6 if the kernel will take it.</span>
        <a id="L85"></a>if kernelSupportsIPv6() {
            <a id="L86"></a>doDial(t, &#34;tcp6&#34;, addr)
        <a id="L87"></a>}
    <a id="L88"></a>}
<a id="L89"></a>}
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
