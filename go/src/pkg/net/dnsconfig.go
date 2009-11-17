<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/dnsconfig.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/net/dnsconfig.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Read system DNS config from /etc/resolv.conf</span>

<a id="L7"></a>package net

<a id="L9"></a>import &#34;os&#34;

<a id="L11"></a>type _DNS_Config struct {
    <a id="L12"></a>servers  []string; <span class="comment">// servers to use</span>
    <a id="L13"></a>search   []string; <span class="comment">// suffixes to append to local name</span>
    <a id="L14"></a>ndots    int;      <span class="comment">// number of dots in name to trigger absolute lookup</span>
    <a id="L15"></a>timeout  int;      <span class="comment">// seconds before giving up on packet</span>
    <a id="L16"></a>attempts int;      <span class="comment">// lost packets before giving up on server</span>
    <a id="L17"></a>rotate   bool;     <span class="comment">// round robin among servers</span>
<a id="L18"></a>}

<a id="L20"></a>var _DNS_configError os.Error

<a id="L22"></a><span class="comment">// See resolv.conf(5) on a Linux machine.</span>
<a id="L23"></a><span class="comment">// TODO(rsc): Supposed to call uname() and chop the beginning</span>
<a id="L24"></a><span class="comment">// of the host name to get the default search domain.</span>
<a id="L25"></a><span class="comment">// We assume it&#39;s in resolv.conf anyway.</span>
<a id="L26"></a>func _DNS_ReadConfig() (*_DNS_Config, os.Error) {
    <a id="L27"></a>file, err := open(&#34;/etc/resolv.conf&#34;);
    <a id="L28"></a>if err != nil {
        <a id="L29"></a>return nil, err
    <a id="L30"></a>}
    <a id="L31"></a>conf := new(_DNS_Config);
    <a id="L32"></a>conf.servers = make([]string, 3)[0:0]; <span class="comment">// small, but the standard limit</span>
    <a id="L33"></a>conf.search = make([]string, 0);
    <a id="L34"></a>conf.ndots = 1;
    <a id="L35"></a>conf.timeout = 1;
    <a id="L36"></a>conf.attempts = 1;
    <a id="L37"></a>conf.rotate = false;
    <a id="L38"></a>for line, ok := file.readLine(); ok; line, ok = file.readLine() {
        <a id="L39"></a>f := getFields(line);
        <a id="L40"></a>if len(f) &lt; 1 {
            <a id="L41"></a>continue
        <a id="L42"></a>}
        <a id="L43"></a>switch f[0] {
        <a id="L44"></a>case &#34;nameserver&#34;: <span class="comment">// add one name server</span>
            <a id="L45"></a>a := conf.servers;
            <a id="L46"></a>n := len(a);
            <a id="L47"></a>if len(f) &gt; 1 &amp;&amp; n &lt; cap(a) {
                <a id="L48"></a><span class="comment">// One more check: make sure server name is</span>
                <a id="L49"></a><span class="comment">// just an IP address.  Otherwise we need DNS</span>
                <a id="L50"></a><span class="comment">// to look it up.</span>
                <a id="L51"></a>name := f[1];
                <a id="L52"></a>if len(ParseIP(name)) != 0 {
                    <a id="L53"></a>a = a[0 : n+1];
                    <a id="L54"></a>a[n] = name;
                    <a id="L55"></a>conf.servers = a;
                <a id="L56"></a>}
            <a id="L57"></a>}

        <a id="L59"></a>case &#34;domain&#34;: <span class="comment">// set search path to just this domain</span>
            <a id="L60"></a>if len(f) &gt; 1 {
                <a id="L61"></a>conf.search = make([]string, 1);
                <a id="L62"></a>conf.search[0] = f[1];
            <a id="L63"></a>} else {
                <a id="L64"></a>conf.search = make([]string, 0)
            <a id="L65"></a>}

        <a id="L67"></a>case &#34;search&#34;: <span class="comment">// set search path to given servers</span>
            <a id="L68"></a>conf.search = make([]string, len(f)-1);
            <a id="L69"></a>for i := 0; i &lt; len(conf.search); i++ {
                <a id="L70"></a>conf.search[i] = f[i+1]
            <a id="L71"></a>}

        <a id="L73"></a>case &#34;options&#34;: <span class="comment">// magic options</span>
            <a id="L74"></a>for i := 1; i &lt; len(f); i++ {
                <a id="L75"></a>s := f[i];
                <a id="L76"></a>switch {
                <a id="L77"></a>case len(s) &gt;= 6 &amp;&amp; s[0:6] == &#34;ndots:&#34;:
                    <a id="L78"></a>n, _, _ := dtoi(s, 6);
                    <a id="L79"></a>if n &lt; 1 {
                        <a id="L80"></a>n = 1
                    <a id="L81"></a>}
                    <a id="L82"></a>conf.ndots = n;
                <a id="L83"></a>case len(s) &gt;= 8 &amp;&amp; s[0:8] == &#34;timeout:&#34;:
                    <a id="L84"></a>n, _, _ := dtoi(s, 8);
                    <a id="L85"></a>if n &lt; 1 {
                        <a id="L86"></a>n = 1
                    <a id="L87"></a>}
                    <a id="L88"></a>conf.timeout = n;
                <a id="L89"></a>case len(s) &gt;= 8 &amp;&amp; s[0:9] == &#34;attempts:&#34;:
                    <a id="L90"></a>n, _, _ := dtoi(s, 9);
                    <a id="L91"></a>if n &lt; 1 {
                        <a id="L92"></a>n = 1
                    <a id="L93"></a>}
                    <a id="L94"></a>conf.attempts = n;
                <a id="L95"></a>case s == &#34;rotate&#34;:
                    <a id="L96"></a>conf.rotate = true
                <a id="L97"></a>}
            <a id="L98"></a>}
        <a id="L99"></a>}
    <a id="L100"></a>}
    <a id="L101"></a>file.close();

    <a id="L103"></a>return conf, nil;
<a id="L104"></a>}
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
