<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/os/env.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/os/env.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Environment variables.</span>

<a id="L7"></a>package os

<a id="L9"></a>import (
    <a id="L10"></a>&#34;once&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">// ENOENV is the Error indicating that an environment variable does not exist.</span>
<a id="L14"></a>var ENOENV = NewError(&#34;no such environment variable&#34;)

<a id="L16"></a>var env map[string]string


<a id="L19"></a>func copyenv() {
    <a id="L20"></a>env = make(map[string]string);
    <a id="L21"></a>for _, s := range Envs {
        <a id="L22"></a>for j := 0; j &lt; len(s); j++ {
            <a id="L23"></a>if s[j] == &#39;=&#39; {
                <a id="L24"></a>env[s[0:j]] = s[j+1 : len(s)];
                <a id="L25"></a>break;
            <a id="L26"></a>}
        <a id="L27"></a>}
    <a id="L28"></a>}
<a id="L29"></a>}

<a id="L31"></a><span class="comment">// Getenverror retrieves the value of the environment variable named by the key.</span>
<a id="L32"></a><span class="comment">// It returns the value and an error, if any.</span>
<a id="L33"></a>func Getenverror(key string) (value string, err Error) {
    <a id="L34"></a>once.Do(copyenv);

    <a id="L36"></a>if len(key) == 0 {
        <a id="L37"></a>return &#34;&#34;, EINVAL
    <a id="L38"></a>}
    <a id="L39"></a>v, ok := env[key];
    <a id="L40"></a>if !ok {
        <a id="L41"></a>return &#34;&#34;, ENOENV
    <a id="L42"></a>}
    <a id="L43"></a>return v, nil;
<a id="L44"></a>}

<a id="L46"></a><span class="comment">// Getenv retrieves the value of the environment variable named by the key.</span>
<a id="L47"></a><span class="comment">// It returns the value, which will be empty if the variable is not present.</span>
<a id="L48"></a>func Getenv(key string) string {
    <a id="L49"></a>v, _ := Getenverror(key);
    <a id="L50"></a>return v;
<a id="L51"></a>}

<a id="L53"></a><span class="comment">// Setenv sets the value of the environment variable named by the key.</span>
<a id="L54"></a><span class="comment">// It returns an Error, if any.</span>
<a id="L55"></a>func Setenv(key, value string) Error {
    <a id="L56"></a>once.Do(copyenv);

    <a id="L58"></a>if len(key) == 0 {
        <a id="L59"></a>return EINVAL
    <a id="L60"></a>}
    <a id="L61"></a>env[key] = value;
    <a id="L62"></a>return nil;
<a id="L63"></a>}

<a id="L65"></a><span class="comment">// Clearenv deletes all environment variables.</span>
<a id="L66"></a>func Clearenv() {
    <a id="L67"></a>once.Do(copyenv); <span class="comment">// prevent copyenv in Getenv/Setenv</span>
    <a id="L68"></a>env = make(map[string]string);
<a id="L69"></a>}

<a id="L71"></a><span class="comment">// Environ returns an array of strings representing the environment,</span>
<a id="L72"></a><span class="comment">// in the form &#34;key=value&#34;.</span>
<a id="L73"></a>func Environ() []string {
    <a id="L74"></a>once.Do(copyenv);
    <a id="L75"></a>a := make([]string, len(env));
    <a id="L76"></a>i := 0;
    <a id="L77"></a>for k, v := range env {
        <a id="L78"></a><span class="comment">// check i &lt; len(a) for safety,</span>
        <a id="L79"></a><span class="comment">// in case env is changing underfoot.</span>
        <a id="L80"></a>if i &lt; len(a) {
            <a id="L81"></a>a[i] = k + &#34;=&#34; + v;
            <a id="L82"></a>i++;
        <a id="L83"></a>}
    <a id="L84"></a>}
    <a id="L85"></a>return a[0:i];
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
