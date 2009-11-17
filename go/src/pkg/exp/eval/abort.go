<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/abort.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

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
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/abort.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package eval

<a id="L7"></a>import (
    <a id="L8"></a>&#34;fmt&#34;;
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>&#34;runtime&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">// Abort aborts the thread&#39;s current computation,</span>
<a id="L14"></a><span class="comment">// causing the innermost Try to return err.</span>
<a id="L15"></a>func (t *Thread) Abort(err os.Error) {
    <a id="L16"></a>if t.abort == nil {
        <a id="L17"></a>panicln(&#34;abort:&#34;, err.String())
    <a id="L18"></a>}
    <a id="L19"></a>t.abort &lt;- err;
    <a id="L20"></a>runtime.Goexit();
<a id="L21"></a>}

<a id="L23"></a><span class="comment">// Try executes a computation; if the computation</span>
<a id="L24"></a><span class="comment">// Aborts, Try returns the error passed to abort.</span>
<a id="L25"></a>func (t *Thread) Try(f func(t *Thread)) os.Error {
    <a id="L26"></a>oc := t.abort;
    <a id="L27"></a>c := make(chan os.Error);
    <a id="L28"></a>t.abort = c;
    <a id="L29"></a>go func() {
        <a id="L30"></a>f(t);
        <a id="L31"></a>c &lt;- nil;
    <a id="L32"></a>}();
    <a id="L33"></a>err := &lt;-c;
    <a id="L34"></a>t.abort = oc;
    <a id="L35"></a>return err;
<a id="L36"></a>}

<a id="L38"></a>type DivByZeroError struct{}

<a id="L40"></a>func (DivByZeroError) String() string { return &#34;divide by zero&#34; }

<a id="L42"></a>type NilPointerError struct{}

<a id="L44"></a>func (NilPointerError) String() string { return &#34;nil pointer dereference&#34; }

<a id="L46"></a>type IndexError struct {
    <a id="L47"></a>Idx, Len int64;
<a id="L48"></a>}

<a id="L50"></a>func (e IndexError) String() string {
    <a id="L51"></a>if e.Idx &lt; 0 {
        <a id="L52"></a>return fmt.Sprintf(&#34;negative index: %d&#34;, e.Idx)
    <a id="L53"></a>}
    <a id="L54"></a>return fmt.Sprintf(&#34;index %d exceeds length %d&#34;, e.Idx, e.Len);
<a id="L55"></a>}

<a id="L57"></a>type SliceError struct {
    <a id="L58"></a>Lo, Hi, Cap int64;
<a id="L59"></a>}

<a id="L61"></a>func (e SliceError) String() string {
    <a id="L62"></a>return fmt.Sprintf(&#34;slice [%d:%d]; cap %d&#34;, e.Lo, e.Hi, e.Cap)
<a id="L63"></a>}

<a id="L65"></a>type KeyError struct {
    <a id="L66"></a>Key interface{};
<a id="L67"></a>}

<a id="L69"></a>func (e KeyError) String() string { return fmt.Sprintf(&#34;key &#39;%v&#39; not found in map&#34;, e.Key) }

<a id="L71"></a>type NegativeLengthError struct {
    <a id="L72"></a>Len int64;
<a id="L73"></a>}

<a id="L75"></a>func (e NegativeLengthError) String() string {
    <a id="L76"></a>return fmt.Sprintf(&#34;negative length: %d&#34;, e.Len)
<a id="L77"></a>}

<a id="L79"></a>type NegativeCapacityError struct {
    <a id="L80"></a>Len int64;
<a id="L81"></a>}

<a id="L83"></a>func (e NegativeCapacityError) String() string {
    <a id="L84"></a>return fmt.Sprintf(&#34;negative capacity: %d&#34;, e.Len)
<a id="L85"></a>}
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
