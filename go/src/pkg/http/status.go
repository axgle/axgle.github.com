<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/http/status.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/http/status.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package http

<a id="L7"></a><span class="comment">// HTTP status codes, defined in RFC 2616.</span>
<a id="L8"></a>const (
    <a id="L9"></a>StatusContinue           = 100;
    <a id="L10"></a>StatusSwitchingProtocols = 101;

    <a id="L12"></a>StatusOK                   = 200;
    <a id="L13"></a>StatusCreated              = 201;
    <a id="L14"></a>StatusAccepted             = 202;
    <a id="L15"></a>StatusNonAuthoritativeInfo = 203;
    <a id="L16"></a>StatusNoContent            = 204;
    <a id="L17"></a>StatusResetContent         = 205;
    <a id="L18"></a>StatusPartialContent       = 206;

    <a id="L20"></a>StatusMultipleChoices   = 300;
    <a id="L21"></a>StatusMovedPermanently  = 301;
    <a id="L22"></a>StatusFound             = 302;
    <a id="L23"></a>StatusSeeOther          = 303;
    <a id="L24"></a>StatusNotModified       = 304;
    <a id="L25"></a>StatusUseProxy          = 305;
    <a id="L26"></a>StatusTemporaryRedirect = 307;

    <a id="L28"></a>StatusBadRequest                   = 400;
    <a id="L29"></a>StatusUnauthorized                 = 401;
    <a id="L30"></a>StatusPaymentRequired              = 402;
    <a id="L31"></a>StatusForbidden                    = 403;
    <a id="L32"></a>StatusNotFound                     = 404;
    <a id="L33"></a>StatusMethodNotAllowed             = 405;
    <a id="L34"></a>StatusNotAcceptable                = 406;
    <a id="L35"></a>StatusProxyAuthRequired            = 407;
    <a id="L36"></a>StatusRequestTimeout               = 408;
    <a id="L37"></a>StatusConflict                     = 409;
    <a id="L38"></a>StatusGone                         = 410;
    <a id="L39"></a>StatusLengthRequired               = 411;
    <a id="L40"></a>StatusPreconditionFailed           = 412;
    <a id="L41"></a>StatusRequestEntityTooLarge        = 413;
    <a id="L42"></a>StatusRequestURITooLong            = 414;
    <a id="L43"></a>StatusUnsupportedMediaType         = 415;
    <a id="L44"></a>StatusRequestedRangeNotSatisfiable = 416;
    <a id="L45"></a>StatusExpectationFailed            = 417;

    <a id="L47"></a>StatusInternalServerError     = 500;
    <a id="L48"></a>StatusNotImplemented          = 501;
    <a id="L49"></a>StatusBadGateway              = 502;
    <a id="L50"></a>StatusServiceUnavailable      = 503;
    <a id="L51"></a>StatusGatewayTimeout          = 504;
    <a id="L52"></a>StatusHTTPVersionNotSupported = 505;
<a id="L53"></a>)

<a id="L55"></a>var statusText = map[int]string{
    <a id="L56"></a>StatusContinue: &#34;Continue&#34;,
    <a id="L57"></a>StatusSwitchingProtocols: &#34;Switching Protocols&#34;,

    <a id="L59"></a>StatusOK: &#34;OK&#34;,
    <a id="L60"></a>StatusCreated: &#34;Created&#34;,
    <a id="L61"></a>StatusAccepted: &#34;Accepted&#34;,
    <a id="L62"></a>StatusNonAuthoritativeInfo: &#34;Non-Authoritative Information&#34;,
    <a id="L63"></a>StatusNoContent: &#34;No Content&#34;,
    <a id="L64"></a>StatusResetContent: &#34;Reset Content&#34;,
    <a id="L65"></a>StatusPartialContent: &#34;Partial Content&#34;,

    <a id="L67"></a>StatusMultipleChoices: &#34;Multiple Choices&#34;,
    <a id="L68"></a>StatusMovedPermanently: &#34;Moved Permanently&#34;,
    <a id="L69"></a>StatusFound: &#34;Found&#34;,
    <a id="L70"></a>StatusSeeOther: &#34;See Other&#34;,
    <a id="L71"></a>StatusNotModified: &#34;Not Modified&#34;,
    <a id="L72"></a>StatusUseProxy: &#34;Use Proxy&#34;,
    <a id="L73"></a>StatusTemporaryRedirect: &#34;Temporary Redirect&#34;,

    <a id="L75"></a>StatusBadRequest: &#34;Bad Request&#34;,
    <a id="L76"></a>StatusUnauthorized: &#34;Unauthorized&#34;,
    <a id="L77"></a>StatusPaymentRequired: &#34;Payment Required&#34;,
    <a id="L78"></a>StatusForbidden: &#34;Forbidden&#34;,
    <a id="L79"></a>StatusNotFound: &#34;Not Found&#34;,
    <a id="L80"></a>StatusMethodNotAllowed: &#34;Method Not Allowed&#34;,
    <a id="L81"></a>StatusNotAcceptable: &#34;Not Acceptable&#34;,
    <a id="L82"></a>StatusProxyAuthRequired: &#34;Proxy Authentication Required&#34;,
    <a id="L83"></a>StatusRequestTimeout: &#34;Request Timeout&#34;,
    <a id="L84"></a>StatusConflict: &#34;Conflict&#34;,
    <a id="L85"></a>StatusGone: &#34;Gone&#34;,
    <a id="L86"></a>StatusLengthRequired: &#34;Length Required&#34;,
    <a id="L87"></a>StatusPreconditionFailed: &#34;Precondition Failed&#34;,
    <a id="L88"></a>StatusRequestEntityTooLarge: &#34;Request Entity Too Large&#34;,
    <a id="L89"></a>StatusRequestURITooLong: &#34;Request URI Too Long&#34;,
    <a id="L90"></a>StatusUnsupportedMediaType: &#34;Unsupported Media Type&#34;,
    <a id="L91"></a>StatusRequestedRangeNotSatisfiable: &#34;Requested Range Not Satisfiable&#34;,
    <a id="L92"></a>StatusExpectationFailed: &#34;Expectation Failed&#34;,

    <a id="L94"></a>StatusInternalServerError: &#34;Internal Server Error&#34;,
    <a id="L95"></a>StatusNotImplemented: &#34;Not Implemented&#34;,
    <a id="L96"></a>StatusBadGateway: &#34;Bad Gateway&#34;,
    <a id="L97"></a>StatusServiceUnavailable: &#34;Service Unavailable&#34;,
    <a id="L98"></a>StatusGatewayTimeout: &#34;Gateway Timeout&#34;,
    <a id="L99"></a>StatusHTTPVersionNotSupported: &#34;HTTP Version Not Supported&#34;,
<a id="L100"></a>}
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
