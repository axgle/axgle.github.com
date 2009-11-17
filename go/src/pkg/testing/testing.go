<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/testing/testing.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/testing/testing.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The testing package provides support for automated testing of Go packages.</span>
<a id="L6"></a><span class="comment">// It is intended to be used in concert with the ``gotest&#39;&#39; utility, which automates</span>
<a id="L7"></a><span class="comment">// execution of any function of the form</span>
<a id="L8"></a><span class="comment">//     func TestXxx(*testing.T)</span>
<a id="L9"></a><span class="comment">// where Xxx can by any alphanumeric string (but the first letter must not be in</span>
<a id="L10"></a><span class="comment">// [a-z]) and serves to identify the test routine.</span>
<a id="L11"></a><span class="comment">// These TestXxx routines should be declared within the package they are testing.</span>
<a id="L12"></a>package testing

<a id="L14"></a>import (
    <a id="L15"></a>&#34;flag&#34;;
    <a id="L16"></a>&#34;fmt&#34;;
    <a id="L17"></a>&#34;os&#34;;
    <a id="L18"></a>&#34;runtime&#34;;
<a id="L19"></a>)

<a id="L21"></a><span class="comment">// Report as tests are run; default is silent for success.</span>
<a id="L22"></a>var chatty = flag.Bool(&#34;v&#34;, false, &#34;verbose: print additional output&#34;)
<a id="L23"></a>var match = flag.String(&#34;match&#34;, &#34;&#34;, &#34;regular expression to select tests to run&#34;)


<a id="L26"></a><span class="comment">// Insert final newline if needed and tabs after internal newlines.</span>
<a id="L27"></a>func tabify(s string) string {
    <a id="L28"></a>n := len(s);
    <a id="L29"></a>if n &gt; 0 &amp;&amp; s[n-1] != &#39;\n&#39; {
        <a id="L30"></a>s += &#34;\n&#34;;
        <a id="L31"></a>n++;
    <a id="L32"></a>}
    <a id="L33"></a>for i := 0; i &lt; n-1; i++ { <span class="comment">// -1 to avoid final newline</span>
        <a id="L34"></a>if s[i] == &#39;\n&#39; {
            <a id="L35"></a>return s[0:i+1] + &#34;\t&#34; + tabify(s[i+1:n])
        <a id="L36"></a>}
    <a id="L37"></a>}
    <a id="L38"></a>return s;
<a id="L39"></a>}

<a id="L41"></a><span class="comment">// T is a type passed to Test functions to manage test state and support formatted test logs.</span>
<a id="L42"></a><span class="comment">// Logs are accumulated during execution and dumped to standard error when done.</span>
<a id="L43"></a>type T struct {
    <a id="L44"></a>errors string;
    <a id="L45"></a>failed bool;
    <a id="L46"></a>ch     chan *T;
<a id="L47"></a>}

<a id="L49"></a><span class="comment">// Fail marks the Test function as having failed but continues execution.</span>
<a id="L50"></a>func (t *T) Fail() { t.failed = true }

<a id="L52"></a><span class="comment">// Failed returns whether the Test function has failed.</span>
<a id="L53"></a>func (t *T) Failed() bool { return t.failed }

<a id="L55"></a><span class="comment">// FailNow marks the Test function as having failed and stops its execution.</span>
<a id="L56"></a><span class="comment">// Execution will continue at the next Test.</span>
<a id="L57"></a>func (t *T) FailNow() {
    <a id="L58"></a>t.Fail();
    <a id="L59"></a>t.ch &lt;- t;
    <a id="L60"></a>runtime.Goexit();
<a id="L61"></a>}

<a id="L63"></a><span class="comment">// Log formats its arguments using default formatting, analogous to Print(),</span>
<a id="L64"></a><span class="comment">// and records the text in the error log.</span>
<a id="L65"></a>func (t *T) Log(args ...) { t.errors += &#34;\t&#34; + tabify(fmt.Sprintln(args)) }

<a id="L67"></a><span class="comment">// Log formats its arguments according to the format, analogous to Printf(),</span>
<a id="L68"></a><span class="comment">// and records the text in the error log.</span>
<a id="L69"></a>func (t *T) Logf(format string, args ...) {
    <a id="L70"></a>t.errors += &#34;\t&#34; + tabify(fmt.Sprintf(format, args))
<a id="L71"></a>}

<a id="L73"></a><span class="comment">// Error is equivalent to Log() followed by Fail().</span>
<a id="L74"></a>func (t *T) Error(args ...) {
    <a id="L75"></a>t.Log(args);
    <a id="L76"></a>t.Fail();
<a id="L77"></a>}

<a id="L79"></a><span class="comment">// Errorf is equivalent to Logf() followed by Fail().</span>
<a id="L80"></a>func (t *T) Errorf(format string, args ...) {
    <a id="L81"></a>t.Logf(format, args);
    <a id="L82"></a>t.Fail();
<a id="L83"></a>}

<a id="L85"></a><span class="comment">// Fatal is equivalent to Log() followed by FailNow().</span>
<a id="L86"></a>func (t *T) Fatal(args ...) {
    <a id="L87"></a>t.Log(args);
    <a id="L88"></a>t.FailNow();
<a id="L89"></a>}

<a id="L91"></a><span class="comment">// Fatalf is equivalent to Logf() followed by FailNow().</span>
<a id="L92"></a>func (t *T) Fatalf(format string, args ...) {
    <a id="L93"></a>t.Logf(format, args);
    <a id="L94"></a>t.FailNow();
<a id="L95"></a>}

<a id="L97"></a><span class="comment">// An internal type but exported because it is cross-package; part of the implementation</span>
<a id="L98"></a><span class="comment">// of gotest.</span>
<a id="L99"></a>type Test struct {
    <a id="L100"></a>Name string;
    <a id="L101"></a>F    func(*T);
<a id="L102"></a>}

<a id="L104"></a>func tRunner(t *T, test *Test) {
    <a id="L105"></a>test.F(t);
    <a id="L106"></a>t.ch &lt;- t;
<a id="L107"></a>}

<a id="L109"></a><span class="comment">// An internal function but exported because it is cross-package; part of the implementation</span>
<a id="L110"></a><span class="comment">// of gotest.</span>
<a id="L111"></a>func Main(tests []Test) {
    <a id="L112"></a>flag.Parse();
    <a id="L113"></a>ok := true;
    <a id="L114"></a>if len(tests) == 0 {
        <a id="L115"></a>println(&#34;testing: warning: no tests to run&#34;)
    <a id="L116"></a>}
    <a id="L117"></a>re, err := CompileRegexp(*match);
    <a id="L118"></a>if err != &#34;&#34; {
        <a id="L119"></a>println(&#34;invalid regexp for -match:&#34;, err);
        <a id="L120"></a>os.Exit(1);
    <a id="L121"></a>}
    <a id="L122"></a>for i := 0; i &lt; len(tests); i++ {
        <a id="L123"></a>if !re.MatchString(tests[i].Name) {
            <a id="L124"></a>continue
        <a id="L125"></a>}
        <a id="L126"></a>if *chatty {
            <a id="L127"></a>println(&#34;=== RUN &#34;, tests[i].Name)
        <a id="L128"></a>}
        <a id="L129"></a>t := new(T);
        <a id="L130"></a>t.ch = make(chan *T);
        <a id="L131"></a>go tRunner(t, &amp;tests[i]);
        <a id="L132"></a>&lt;-t.ch;
        <a id="L133"></a>if t.failed {
            <a id="L134"></a>println(&#34;--- FAIL:&#34;, tests[i].Name);
            <a id="L135"></a>print(t.errors);
            <a id="L136"></a>ok = false;
        <a id="L137"></a>} else if *chatty {
            <a id="L138"></a>println(&#34;--- PASS:&#34;, tests[i].Name);
            <a id="L139"></a>print(t.errors);
        <a id="L140"></a>}
    <a id="L141"></a>}
    <a id="L142"></a>if !ok {
        <a id="L143"></a>println(&#34;FAIL&#34;);
        <a id="L144"></a>os.Exit(1);
    <a id="L145"></a>}
    <a id="L146"></a>println(&#34;PASS&#34;);
<a id="L147"></a>}
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
