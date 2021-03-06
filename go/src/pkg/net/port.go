<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/port.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/net/port.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Read system port mappings from /etc/services</span>

<a id="L7"></a>package net

<a id="L9"></a>import (
    <a id="L10"></a>&#34;once&#34;;
    <a id="L11"></a>&#34;os&#34;;
<a id="L12"></a>)

<a id="L14"></a>var services map[string]map[string]int
<a id="L15"></a>var servicesError os.Error

<a id="L17"></a>func readServices() {
    <a id="L18"></a>services = make(map[string]map[string]int);
    <a id="L19"></a>var file *file;
    <a id="L20"></a>file, servicesError = open(&#34;/etc/services&#34;);
    <a id="L21"></a>for line, ok := file.readLine(); ok; line, ok = file.readLine() {
        <a id="L22"></a><span class="comment">// &#34;http 80/tcp www www-http # World Wide Web HTTP&#34;</span>
        <a id="L23"></a>if i := byteIndex(line, &#39;#&#39;); i &gt;= 0 {
            <a id="L24"></a>line = line[0:i]
        <a id="L25"></a>}
        <a id="L26"></a>f := getFields(line);
        <a id="L27"></a>if len(f) &lt; 2 {
            <a id="L28"></a>continue
        <a id="L29"></a>}
        <a id="L30"></a>portnet := f[1]; <span class="comment">// &#34;tcp/80&#34;</span>
        <a id="L31"></a>port, j, ok := dtoi(portnet, 0);
        <a id="L32"></a>if !ok || port &lt;= 0 || j &gt;= len(portnet) || portnet[j] != &#39;/&#39; {
            <a id="L33"></a>continue
        <a id="L34"></a>}
        <a id="L35"></a>netw := portnet[j+1 : len(portnet)]; <span class="comment">// &#34;tcp&#34;</span>
        <a id="L36"></a>m, ok1 := services[netw];
        <a id="L37"></a>if !ok1 {
            <a id="L38"></a>m = make(map[string]int);
            <a id="L39"></a>services[netw] = m;
        <a id="L40"></a>}
        <a id="L41"></a>for i := 0; i &lt; len(f); i++ {
            <a id="L42"></a>if i != 1 { <span class="comment">// f[1] was port/net</span>
                <a id="L43"></a>m[f[i]] = port
            <a id="L44"></a>}
        <a id="L45"></a>}
    <a id="L46"></a>}
    <a id="L47"></a>file.close();
<a id="L48"></a>}

<a id="L50"></a><span class="comment">// LookupPort looks up the port for the given network and service.</span>
<a id="L51"></a>func LookupPort(network, service string) (port int, err os.Error) {
    <a id="L52"></a>once.Do(readServices);

    <a id="L54"></a>switch network {
    <a id="L55"></a>case &#34;tcp4&#34;, &#34;tcp6&#34;:
        <a id="L56"></a>network = &#34;tcp&#34;
    <a id="L57"></a>case &#34;udp4&#34;, &#34;udp6&#34;:
        <a id="L58"></a>network = &#34;udp&#34;
    <a id="L59"></a>}

    <a id="L61"></a>if m, ok := services[network]; ok {
        <a id="L62"></a>if port, ok = m[service]; ok {
            <a id="L63"></a>return
        <a id="L64"></a>}
    <a id="L65"></a>}
    <a id="L66"></a>return 0, &amp;AddrError{&#34;unknown port&#34;, network + &#34;/&#34; + service};
<a id="L67"></a>}
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
