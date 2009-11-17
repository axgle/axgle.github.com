<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/rpc/debug.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/rpc/debug.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package rpc

<a id="L7"></a><span class="comment">/*</span>
<a id="L8"></a><span class="comment">	Some HTML presented at http://machine:port/debug/rpc</span>
<a id="L9"></a><span class="comment">	Lists services, their methods, and some statistics, still rudimentary.</span>
<a id="L10"></a><span class="comment">*/</span>

<a id="L12"></a>import (
    <a id="L13"></a>&#34;fmt&#34;;
    <a id="L14"></a>&#34;http&#34;;
    <a id="L15"></a>&#34;sort&#34;;
    <a id="L16"></a>&#34;template&#34;;
<a id="L17"></a>)

<a id="L19"></a>const debugText = `&lt;html&gt;
	&lt;body&gt;
	&lt;title&gt;Services&lt;/title&gt;
	{.repeated section @}
	&lt;hr&gt;
	Service {name}
	&lt;hr&gt;
		&lt;table&gt;
		&lt;th align=center&gt;Method&lt;/th&gt;&lt;th align=center&gt;Calls&lt;/th&gt;
		{.repeated section meth}
			&lt;tr&gt;
			&lt;td align=left font=fixed&gt;{name}({m.argType}, {m.replyType}) os.Error&lt;/td&gt;
			&lt;td align=center&gt;{m.numCalls}&lt;/td&gt;
			&lt;/tr&gt;
		{.end}
		&lt;/table&gt;
	{.end}
	&lt;/body&gt;
	&lt;/html&gt;`

<a id="L39"></a>var debug = template.MustParse(debugText, nil)

<a id="L41"></a>type debugMethod struct {
    <a id="L42"></a>m    *methodType;
    <a id="L43"></a>name string;
<a id="L44"></a>}

<a id="L46"></a>type methodArray []debugMethod

<a id="L48"></a>type debugService struct {
    <a id="L49"></a>s    *service;
    <a id="L50"></a>name string;
    <a id="L51"></a>meth methodArray;
<a id="L52"></a>}

<a id="L54"></a>type serviceArray []debugService

<a id="L56"></a>func (s serviceArray) Len() int           { return len(s) }
<a id="L57"></a>func (s serviceArray) Less(i, j int) bool { return s[i].name &lt; s[j].name }
<a id="L58"></a>func (s serviceArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

<a id="L60"></a>func (m methodArray) Len() int           { return len(m) }
<a id="L61"></a>func (m methodArray) Less(i, j int) bool { return m[i].name &lt; m[j].name }
<a id="L62"></a>func (m methodArray) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

<a id="L64"></a><span class="comment">// Runs at /debug/rpc</span>
<a id="L65"></a>func debugHTTP(c *http.Conn, req *http.Request) {
    <a id="L66"></a><span class="comment">// Build a sorted version of the data.</span>
    <a id="L67"></a>var services = make(serviceArray, len(server.serviceMap));
    <a id="L68"></a>i := 0;
    <a id="L69"></a>server.Lock();
    <a id="L70"></a>for sname, service := range server.serviceMap {
        <a id="L71"></a>services[i] = debugService{service, sname, make(methodArray, len(service.method))};
        <a id="L72"></a>j := 0;
        <a id="L73"></a>for mname, method := range service.method {
            <a id="L74"></a>services[i].meth[j] = debugMethod{method, mname};
            <a id="L75"></a>j++;
        <a id="L76"></a>}
        <a id="L77"></a>sort.Sort(services[i].meth);
        <a id="L78"></a>i++;
    <a id="L79"></a>}
    <a id="L80"></a>server.Unlock();
    <a id="L81"></a>sort.Sort(services);
    <a id="L82"></a>err := debug.Execute(services, c);
    <a id="L83"></a>if err != nil {
        <a id="L84"></a>fmt.Fprintln(c, &#34;rpc: error executing template:&#34;, err.String())
    <a id="L85"></a>}
<a id="L86"></a>}
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
