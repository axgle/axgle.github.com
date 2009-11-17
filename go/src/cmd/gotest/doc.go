<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/gotest/doc.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/gotest/doc.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">/*</span>

<a id="L7"></a><span class="comment">Gotest is an automated testing tool for Go packages.</span>

<a id="L9"></a><span class="comment">Normally a Go package is compiled without its test files.  Gotest</span>
<a id="L10"></a><span class="comment">is a simple script that recompiles the package along with any files</span>
<a id="L11"></a><span class="comment">named *_test.go.  Functions in the test sources named TestXXX</span>
<a id="L12"></a><span class="comment">(where XXX is any alphanumeric string starting with an upper case</span>
<a id="L13"></a><span class="comment">letter) will be run when the binary is executed.  Gotest requires</span>
<a id="L14"></a><span class="comment">that the package have a standard package Makefile, one that</span>
<a id="L15"></a><span class="comment">includes go/src/Make.pkg.</span>

<a id="L17"></a><span class="comment">The test functions are run in the order they appear in the source.</span>
<a id="L18"></a><span class="comment">They should have signature</span>

<a id="L20"></a><span class="comment">	func TestXXX(t *testing.T) { ... }</span>

<a id="L22"></a><span class="comment">See the documentation of the testing package for more information.</span>

<a id="L24"></a><span class="comment">By default, gotest needs no arguments.  It compiles all the .go files</span>
<a id="L25"></a><span class="comment">in the directory, including tests, and runs the tests.  If file names</span>
<a id="L26"></a><span class="comment">are given, only those test files are added to the package.</span>
<a id="L27"></a><span class="comment">(The non-test files are always compiled.)</span>

<a id="L29"></a><span class="comment">The package is built in a special subdirectory so it does not</span>
<a id="L30"></a><span class="comment">interfere with the non-test installation.</span>

<a id="L32"></a><span class="comment">Usage:</span>
<a id="L33"></a><span class="comment">	gotest [pkg_test.go ...]</span>

<a id="L35"></a><span class="comment">The resulting binary, called (for amd64) 6.out, has a couple of</span>
<a id="L36"></a><span class="comment">arguments.</span>

<a id="L38"></a><span class="comment">Usage:</span>
<a id="L39"></a><span class="comment">	6.out [-v] [-match pattern]</span>

<a id="L41"></a><span class="comment">The -v flag causes the tests to be logged as they run.  The --match</span>
<a id="L42"></a><span class="comment">flag causes only those tests whose names match the regular expression</span>
<a id="L43"></a><span class="comment">pattern to be run. By default all tests are run silently.  If all</span>
<a id="L44"></a><span class="comment">the specified test pass, 6.out prints PASS and exits with a 0 exit</span>
<a id="L45"></a><span class="comment">code.  If any tests fail, it prints FAIL and exits with a non-zero</span>
<a id="L46"></a><span class="comment">code.</span>

<a id="L48"></a><span class="comment">*/</span>
<a id="L49"></a>package documentation
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
