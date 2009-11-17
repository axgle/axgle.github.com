<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strconv/fp_test.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/strconv/fp_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package strconv_test

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bufio&#34;;
    <a id="L9"></a>&#34;fmt&#34;;
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;strconv&#34;;
    <a id="L12"></a>&#34;strings&#34;;
    <a id="L13"></a>&#34;testing&#34;;
<a id="L14"></a>)

<a id="L16"></a>func pow2(i int) float64 {
    <a id="L17"></a>switch {
    <a id="L18"></a>case i &lt; 0:
        <a id="L19"></a>return 1 / pow2(-i)
    <a id="L20"></a>case i == 0:
        <a id="L21"></a>return 1
    <a id="L22"></a>case i == 1:
        <a id="L23"></a>return 2
    <a id="L24"></a>}
    <a id="L25"></a>return pow2(i/2) * pow2(i-i/2);
<a id="L26"></a>}

<a id="L28"></a><span class="comment">// Wrapper around strconv.Atof64.  Handles dddddp+ddd (binary exponent)</span>
<a id="L29"></a><span class="comment">// itself, passes the rest on to strconv.Atof64.</span>
<a id="L30"></a>func myatof64(s string) (f float64, ok bool) {
    <a id="L31"></a>a := strings.Split(s, &#34;p&#34;, 2);
    <a id="L32"></a>if len(a) == 2 {
        <a id="L33"></a>n, err := strconv.Atoi64(a[0]);
        <a id="L34"></a>if err != nil {
            <a id="L35"></a>return 0, false
        <a id="L36"></a>}
        <a id="L37"></a>e, err1 := strconv.Atoi(a[1]);
        <a id="L38"></a>if err1 != nil {
            <a id="L39"></a>println(&#34;bad e&#34;, a[1]);
            <a id="L40"></a>return 0, false;
        <a id="L41"></a>}
        <a id="L42"></a>v := float64(n);
        <a id="L43"></a><span class="comment">// We expect that v*pow2(e) fits in a float64,</span>
        <a id="L44"></a><span class="comment">// but pow2(e) by itself may not.  Be careful.</span>
        <a id="L45"></a>if e &lt;= -1000 {
            <a id="L46"></a>v *= pow2(-1000);
            <a id="L47"></a>e += 1000;
            <a id="L48"></a>for e &lt; 0 {
                <a id="L49"></a>v /= 2;
                <a id="L50"></a>e++;
            <a id="L51"></a>}
            <a id="L52"></a>return v, true;
        <a id="L53"></a>}
        <a id="L54"></a>if e &gt;= 1000 {
            <a id="L55"></a>v *= pow2(1000);
            <a id="L56"></a>e -= 1000;
            <a id="L57"></a>for e &gt; 0 {
                <a id="L58"></a>v *= 2;
                <a id="L59"></a>e--;
            <a id="L60"></a>}
            <a id="L61"></a>return v, true;
        <a id="L62"></a>}
        <a id="L63"></a>return v * pow2(e), true;
    <a id="L64"></a>}
    <a id="L65"></a>f1, err := strconv.Atof64(s);
    <a id="L66"></a>if err != nil {
        <a id="L67"></a>return 0, false
    <a id="L68"></a>}
    <a id="L69"></a>return f1, true;
<a id="L70"></a>}

<a id="L72"></a><span class="comment">// Wrapper around strconv.Atof32.  Handles dddddp+ddd (binary exponent)</span>
<a id="L73"></a><span class="comment">// itself, passes the rest on to strconv.Atof32.</span>
<a id="L74"></a>func myatof32(s string) (f float32, ok bool) {
    <a id="L75"></a>a := strings.Split(s, &#34;p&#34;, 2);
    <a id="L76"></a>if len(a) == 2 {
        <a id="L77"></a>n, err := strconv.Atoi(a[0]);
        <a id="L78"></a>if err != nil {
            <a id="L79"></a>println(&#34;bad n&#34;, a[0]);
            <a id="L80"></a>return 0, false;
        <a id="L81"></a>}
        <a id="L82"></a>e, err1 := strconv.Atoi(a[1]);
        <a id="L83"></a>if err1 != nil {
            <a id="L84"></a>println(&#34;bad p&#34;, a[1]);
            <a id="L85"></a>return 0, false;
        <a id="L86"></a>}
        <a id="L87"></a>return float32(float64(n) * pow2(e)), true;
    <a id="L88"></a>}
    <a id="L89"></a>f1, err1 := strconv.Atof32(s);
    <a id="L90"></a>if err1 != nil {
        <a id="L91"></a>return 0, false
    <a id="L92"></a>}
    <a id="L93"></a>return f1, true;
<a id="L94"></a>}

<a id="L96"></a>func TestFp(t *testing.T) {
    <a id="L97"></a>f, err := os.Open(&#34;testfp.txt&#34;, os.O_RDONLY, 0);
    <a id="L98"></a>if err != nil {
        <a id="L99"></a>panicln(&#34;testfp: open testfp.txt:&#34;, err.String())
    <a id="L100"></a>}
    <a id="L101"></a>defer f.Close();

    <a id="L103"></a>b := bufio.NewReader(f);

    <a id="L105"></a>lineno := 0;
    <a id="L106"></a>for {
        <a id="L107"></a>line, err2 := b.ReadString(&#39;\n&#39;);
        <a id="L108"></a>if err2 == os.EOF {
            <a id="L109"></a>break
        <a id="L110"></a>}
        <a id="L111"></a>if err2 != nil {
            <a id="L112"></a>panicln(&#34;testfp: read testfp.txt:&#34;, err2.String())
        <a id="L113"></a>}
        <a id="L114"></a>line = line[0 : len(line)-1];
        <a id="L115"></a>lineno++;
        <a id="L116"></a>if len(line) == 0 || line[0] == &#39;#&#39; {
            <a id="L117"></a>continue
        <a id="L118"></a>}
        <a id="L119"></a>a := strings.Split(line, &#34; &#34;, 0);
        <a id="L120"></a>if len(a) != 4 {
            <a id="L121"></a>t.Error(&#34;testfp.txt:&#34;, lineno, &#34;: wrong field count\n&#34;);
            <a id="L122"></a>continue;
        <a id="L123"></a>}
        <a id="L124"></a>var s string;
        <a id="L125"></a>var v float64;
        <a id="L126"></a>switch a[0] {
        <a id="L127"></a>case &#34;float64&#34;:
            <a id="L128"></a>var ok bool;
            <a id="L129"></a>v, ok = myatof64(a[2]);
            <a id="L130"></a>if !ok {
                <a id="L131"></a>t.Error(&#34;testfp.txt:&#34;, lineno, &#34;: cannot atof64 &#34;, a[2]);
                <a id="L132"></a>continue;
            <a id="L133"></a>}
            <a id="L134"></a>s = fmt.Sprintf(a[1], v);
        <a id="L135"></a>case &#34;float32&#34;:
            <a id="L136"></a>v1, ok := myatof32(a[2]);
            <a id="L137"></a>if !ok {
                <a id="L138"></a>t.Error(&#34;testfp.txt:&#34;, lineno, &#34;: cannot atof32 &#34;, a[2]);
                <a id="L139"></a>continue;
            <a id="L140"></a>}
            <a id="L141"></a>s = fmt.Sprintf(a[1], v1);
            <a id="L142"></a>v = float64(v1);
        <a id="L143"></a>}
        <a id="L144"></a>if s != a[3] {
            <a id="L145"></a>t.Error(&#34;testfp.txt:&#34;, lineno, &#34;: &#34;, a[0], &#34; &#34;, a[1], &#34; &#34;, a[2], &#34; (&#34;, v, &#34;) &#34;,
                <a id="L146"></a>&#34;want &#34;, a[3], &#34; got &#34;, s)
        <a id="L147"></a>}
    <a id="L148"></a>}
<a id="L149"></a>}
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
