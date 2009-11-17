<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/testing/quick/quick_test.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/testing/quick/quick_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package quick

<a id="L7"></a>import (
    <a id="L8"></a>&#34;rand&#34;;
    <a id="L9"></a>&#34;reflect&#34;;
    <a id="L10"></a>&#34;testing&#34;;
    <a id="L11"></a>&#34;os&#34;;
<a id="L12"></a>)

<a id="L14"></a>func fBool(a bool) bool { return a }

<a id="L16"></a>func fFloat32(a float32) float32 { return a }

<a id="L18"></a>func fFloat64(a float64) float64 { return a }

<a id="L20"></a>func fFloat(a float) float { return a }

<a id="L22"></a>func fInt16(a int16) int16 { return a }

<a id="L24"></a>func fInt32(a int32) int32 { return a }

<a id="L26"></a>func fInt64(a int64) int64 { return a }

<a id="L28"></a>func fInt8(a int8) int8 { return a }

<a id="L30"></a>func fInt(a int) int { return a }

<a id="L32"></a>func fUInt8(a uint8) uint8 { return a }

<a id="L34"></a>func fMap(a map[int]int) map[int]int { return a }

<a id="L36"></a>func fSlice(a []byte) []byte { return a }

<a id="L38"></a>func fString(a string) string { return a }

<a id="L40"></a>type TestStruct struct {
    <a id="L41"></a>A   int;
    <a id="L42"></a>B   string;
<a id="L43"></a>}

<a id="L45"></a>func fStruct(a TestStruct) TestStruct { return a }

<a id="L47"></a>func fUint16(a uint16) uint16 { return a }

<a id="L49"></a>func fUint32(a uint32) uint32 { return a }

<a id="L51"></a>func fUint64(a uint64) uint64 { return a }

<a id="L53"></a>func fUint8(a uint8) uint8 { return a }

<a id="L55"></a>func fUint(a uint) uint { return a }

<a id="L57"></a>func fUintptr(a uintptr) uintptr { return a }

<a id="L59"></a>func fIntptr(a *int) *int {
    <a id="L60"></a>b := *a;
    <a id="L61"></a>return &amp;b;
<a id="L62"></a>}

<a id="L64"></a>func reportError(property string, err os.Error, t *testing.T) {
    <a id="L65"></a>if err != nil {
        <a id="L66"></a>t.Errorf(&#34;%s: %s&#34;, property, err)
    <a id="L67"></a>}
<a id="L68"></a>}

<a id="L70"></a>func TestCheckEqual(t *testing.T) {
    <a id="L71"></a>reportError(&#34;fBool&#34;, CheckEqual(fBool, fBool, nil), t);
    <a id="L72"></a>reportError(&#34;fFloat32&#34;, CheckEqual(fFloat32, fFloat32, nil), t);
    <a id="L73"></a>reportError(&#34;fFloat64&#34;, CheckEqual(fFloat64, fFloat64, nil), t);
    <a id="L74"></a>reportError(&#34;fFloat&#34;, CheckEqual(fFloat, fFloat, nil), t);
    <a id="L75"></a>reportError(&#34;fInt16&#34;, CheckEqual(fInt16, fInt16, nil), t);
    <a id="L76"></a>reportError(&#34;fInt32&#34;, CheckEqual(fInt32, fInt32, nil), t);
    <a id="L77"></a>reportError(&#34;fInt64&#34;, CheckEqual(fInt64, fInt64, nil), t);
    <a id="L78"></a>reportError(&#34;fInt8&#34;, CheckEqual(fInt8, fInt8, nil), t);
    <a id="L79"></a>reportError(&#34;fInt&#34;, CheckEqual(fInt, fInt, nil), t);
    <a id="L80"></a>reportError(&#34;fUInt8&#34;, CheckEqual(fUInt8, fUInt8, nil), t);
    <a id="L81"></a>reportError(&#34;fInt32&#34;, CheckEqual(fInt32, fInt32, nil), t);
    <a id="L82"></a>reportError(&#34;fMap&#34;, CheckEqual(fMap, fMap, nil), t);
    <a id="L83"></a>reportError(&#34;fSlice&#34;, CheckEqual(fSlice, fSlice, nil), t);
    <a id="L84"></a>reportError(&#34;fString&#34;, CheckEqual(fString, fString, nil), t);
    <a id="L85"></a>reportError(&#34;fStruct&#34;, CheckEqual(fStruct, fStruct, nil), t);
    <a id="L86"></a>reportError(&#34;fUint16&#34;, CheckEqual(fUint16, fUint16, nil), t);
    <a id="L87"></a>reportError(&#34;fUint32&#34;, CheckEqual(fUint32, fUint32, nil), t);
    <a id="L88"></a>reportError(&#34;fUint64&#34;, CheckEqual(fUint64, fUint64, nil), t);
    <a id="L89"></a>reportError(&#34;fUint8&#34;, CheckEqual(fUint8, fUint8, nil), t);
    <a id="L90"></a>reportError(&#34;fUint&#34;, CheckEqual(fUint, fUint, nil), t);
    <a id="L91"></a>reportError(&#34;fUintptr&#34;, CheckEqual(fUintptr, fUintptr, nil), t);
    <a id="L92"></a>reportError(&#34;fIntptr&#34;, CheckEqual(fIntptr, fIntptr, nil), t);
<a id="L93"></a>}

<a id="L95"></a><span class="comment">// This tests that ArbitraryValue is working by checking that all the arbitrary</span>
<a id="L96"></a><span class="comment">// values of type MyStruct have x = 42.</span>
<a id="L97"></a>type myStruct struct {
    <a id="L98"></a>x int;
<a id="L99"></a>}

<a id="L101"></a>func (m myStruct) Generate(r *rand.Rand, _ int) reflect.Value {
    <a id="L102"></a>return reflect.NewValue(myStruct{x: 42})
<a id="L103"></a>}

<a id="L105"></a>func myStructProperty(in myStruct) bool { return in.x == 42 }

<a id="L107"></a>func TestCheckProperty(t *testing.T) {
    <a id="L108"></a>reportError(&#34;myStructProperty&#34;, Check(myStructProperty, nil), t)
<a id="L109"></a>}

<a id="L111"></a>func TestFailure(t *testing.T) {
    <a id="L112"></a>f := func(x int) bool { return false };
    <a id="L113"></a>err := Check(f, nil);
    <a id="L114"></a>if err == nil {
        <a id="L115"></a>t.Errorf(&#34;Check didn&#39;t return an error&#34;)
    <a id="L116"></a>}
    <a id="L117"></a>if _, ok := err.(*CheckError); !ok {
        <a id="L118"></a>t.Errorf(&#34;Error was not a CheckError: %s&#34;, err)
    <a id="L119"></a>}

    <a id="L121"></a>err = CheckEqual(fUint, fUint32, nil);
    <a id="L122"></a>if err == nil {
        <a id="L123"></a>t.Errorf(&#34;#1 CheckEqual didn&#39;t return an error&#34;)
    <a id="L124"></a>}
    <a id="L125"></a>if _, ok := err.(SetupError); !ok {
        <a id="L126"></a>t.Errorf(&#34;#1 Error was not a SetupError: %s&#34;, err)
    <a id="L127"></a>}

    <a id="L129"></a>err = CheckEqual(func(x, y int) {}, func(x int) {}, nil);
    <a id="L130"></a>if err == nil {
        <a id="L131"></a>t.Errorf(&#34;#2 CheckEqual didn&#39;t return an error&#34;)
    <a id="L132"></a>}
    <a id="L133"></a>if _, ok := err.(SetupError); !ok {
        <a id="L134"></a>t.Errorf(&#34;#2 Error was not a SetupError: %s&#34;, err)
    <a id="L135"></a>}

    <a id="L137"></a>err = CheckEqual(func(x int) int { return 0 }, func(x int) int32 { return 0 }, nil);
    <a id="L138"></a>if err == nil {
        <a id="L139"></a>t.Errorf(&#34;#3 CheckEqual didn&#39;t return an error&#34;)
    <a id="L140"></a>}
    <a id="L141"></a>if _, ok := err.(SetupError); !ok {
        <a id="L142"></a>t.Errorf(&#34;#3 Error was not a SetupError: %s&#34;, err)
    <a id="L143"></a>}
<a id="L144"></a>}
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
