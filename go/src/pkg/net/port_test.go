<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/port_test.go</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/net/port_test.go</h1>

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
    <a id="L8"></a>&#34;testing&#34;;
<a id="L9"></a>)

<a id="L11"></a>type portTest struct {
    <a id="L12"></a>netw string;
    <a id="L13"></a>name string;
    <a id="L14"></a>port int;
    <a id="L15"></a>ok   bool;
<a id="L16"></a>}

<a id="L18"></a>var porttests = []portTest{
    <a id="L19"></a>portTest{&#34;tcp&#34;, &#34;echo&#34;, 7, true},
    <a id="L20"></a>portTest{&#34;tcp&#34;, &#34;discard&#34;, 9, true},
    <a id="L21"></a>portTest{&#34;tcp&#34;, &#34;systat&#34;, 11, true},
    <a id="L22"></a>portTest{&#34;tcp&#34;, &#34;daytime&#34;, 13, true},
    <a id="L23"></a>portTest{&#34;tcp&#34;, &#34;chargen&#34;, 19, true},
    <a id="L24"></a>portTest{&#34;tcp&#34;, &#34;ftp-data&#34;, 20, true},
    <a id="L25"></a>portTest{&#34;tcp&#34;, &#34;ftp&#34;, 21, true},
    <a id="L26"></a>portTest{&#34;tcp&#34;, &#34;ssh&#34;, 22, true},
    <a id="L27"></a>portTest{&#34;tcp&#34;, &#34;telnet&#34;, 23, true},
    <a id="L28"></a>portTest{&#34;tcp&#34;, &#34;smtp&#34;, 25, true},
    <a id="L29"></a>portTest{&#34;tcp&#34;, &#34;time&#34;, 37, true},
    <a id="L30"></a>portTest{&#34;tcp&#34;, &#34;domain&#34;, 53, true},
    <a id="L31"></a>portTest{&#34;tcp&#34;, &#34;gopher&#34;, 70, true},
    <a id="L32"></a>portTest{&#34;tcp&#34;, &#34;finger&#34;, 79, true},
    <a id="L33"></a>portTest{&#34;tcp&#34;, &#34;http&#34;, 80, true},

    <a id="L35"></a>portTest{&#34;udp&#34;, &#34;echo&#34;, 7, true},
    <a id="L36"></a>portTest{&#34;udp&#34;, &#34;tacacs&#34;, 49, true},
    <a id="L37"></a>portTest{&#34;udp&#34;, &#34;tftp&#34;, 69, true},
    <a id="L38"></a>portTest{&#34;udp&#34;, &#34;bootpc&#34;, 68, true},
    <a id="L39"></a>portTest{&#34;udp&#34;, &#34;bootps&#34;, 67, true},
    <a id="L40"></a>portTest{&#34;udp&#34;, &#34;domain&#34;, 53, true},
    <a id="L41"></a>portTest{&#34;udp&#34;, &#34;ntp&#34;, 123, true},
    <a id="L42"></a>portTest{&#34;udp&#34;, &#34;snmp&#34;, 161, true},
    <a id="L43"></a>portTest{&#34;udp&#34;, &#34;syslog&#34;, 514, true},
    <a id="L44"></a>portTest{&#34;udp&#34;, &#34;nfs&#34;, 2049, true},

    <a id="L46"></a>portTest{&#34;--badnet--&#34;, &#34;zzz&#34;, 0, false},
    <a id="L47"></a>portTest{&#34;tcp&#34;, &#34;--badport--&#34;, 0, false},
<a id="L48"></a>}

<a id="L50"></a>func TestLookupPort(t *testing.T) {
    <a id="L51"></a>for i := 0; i &lt; len(porttests); i++ {
        <a id="L52"></a>tt := porttests[i];
        <a id="L53"></a>if port, err := LookupPort(tt.netw, tt.name); port != tt.port || (err == nil) != tt.ok {
            <a id="L54"></a>t.Errorf(&#34;LookupPort(%q, %q) = %v, %s; want %v&#34;,
                <a id="L55"></a>tt.netw, tt.name, port, err, tt.port)
        <a id="L56"></a>}
    <a id="L57"></a>}
<a id="L58"></a>}
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
