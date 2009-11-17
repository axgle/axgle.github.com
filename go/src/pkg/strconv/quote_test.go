<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strconv/quote_test.go</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/strconv/quote_test.go</h1>

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
    <a id="L8"></a>&#34;os&#34;;
    <a id="L9"></a>. &#34;strconv&#34;;
    <a id="L10"></a>&#34;testing&#34;;
<a id="L11"></a>)

<a id="L13"></a>type quoteTest struct {
    <a id="L14"></a>in  string;
    <a id="L15"></a>out string;
<a id="L16"></a>}

<a id="L18"></a>var quotetests = []quoteTest{
    <a id="L19"></a>quoteTest{&#34;\a\b\f\r\n\t\v&#34;, `&#34;\a\b\f\r\n\t\v&#34;`},
    <a id="L20"></a>quoteTest{&#34;\\&#34;, `&#34;\\&#34;`},
    <a id="L21"></a>quoteTest{&#34;abc\xffdef&#34;, `&#34;abc\xffdef&#34;`},
    <a id="L22"></a>quoteTest{&#34;\u263a&#34;, `&#34;\u263a&#34;`},
    <a id="L23"></a>quoteTest{&#34;\U0010ffff&#34;, `&#34;\U0010ffff&#34;`},
    <a id="L24"></a>quoteTest{&#34;\x04&#34;, `&#34;\x04&#34;`},
<a id="L25"></a>}

<a id="L27"></a>func TestQuote(t *testing.T) {
    <a id="L28"></a>for i := 0; i &lt; len(quotetests); i++ {
        <a id="L29"></a>tt := quotetests[i];
        <a id="L30"></a>if out := Quote(tt.in); out != tt.out {
            <a id="L31"></a>t.Errorf(&#34;Quote(%s) = %s, want %s&#34;, tt.in, out, tt.out)
        <a id="L32"></a>}
    <a id="L33"></a>}
<a id="L34"></a>}

<a id="L36"></a>type canBackquoteTest struct {
    <a id="L37"></a>in  string;
    <a id="L38"></a>out bool;
<a id="L39"></a>}

<a id="L41"></a>var canbackquotetests = []canBackquoteTest{
    <a id="L42"></a>canBackquoteTest{&#34;`&#34;, false},
    <a id="L43"></a>canBackquoteTest{string(0), false},
    <a id="L44"></a>canBackquoteTest{string(1), false},
    <a id="L45"></a>canBackquoteTest{string(2), false},
    <a id="L46"></a>canBackquoteTest{string(3), false},
    <a id="L47"></a>canBackquoteTest{string(4), false},
    <a id="L48"></a>canBackquoteTest{string(5), false},
    <a id="L49"></a>canBackquoteTest{string(6), false},
    <a id="L50"></a>canBackquoteTest{string(7), false},
    <a id="L51"></a>canBackquoteTest{string(8), false},
    <a id="L52"></a>canBackquoteTest{string(9), true}, <span class="comment">// \t</span>
    <a id="L53"></a>canBackquoteTest{string(10), false},
    <a id="L54"></a>canBackquoteTest{string(11), false},
    <a id="L55"></a>canBackquoteTest{string(12), false},
    <a id="L56"></a>canBackquoteTest{string(13), false},
    <a id="L57"></a>canBackquoteTest{string(14), false},
    <a id="L58"></a>canBackquoteTest{string(15), false},
    <a id="L59"></a>canBackquoteTest{string(16), false},
    <a id="L60"></a>canBackquoteTest{string(17), false},
    <a id="L61"></a>canBackquoteTest{string(18), false},
    <a id="L62"></a>canBackquoteTest{string(19), false},
    <a id="L63"></a>canBackquoteTest{string(20), false},
    <a id="L64"></a>canBackquoteTest{string(21), false},
    <a id="L65"></a>canBackquoteTest{string(22), false},
    <a id="L66"></a>canBackquoteTest{string(23), false},
    <a id="L67"></a>canBackquoteTest{string(24), false},
    <a id="L68"></a>canBackquoteTest{string(25), false},
    <a id="L69"></a>canBackquoteTest{string(26), false},
    <a id="L70"></a>canBackquoteTest{string(27), false},
    <a id="L71"></a>canBackquoteTest{string(28), false},
    <a id="L72"></a>canBackquoteTest{string(29), false},
    <a id="L73"></a>canBackquoteTest{string(30), false},
    <a id="L74"></a>canBackquoteTest{string(31), false},
    <a id="L75"></a>canBackquoteTest{`&#39; !&#34;#$%&amp;&#39;()*+,-./:;&lt;=&gt;?@[\]^_{|}~`, true},
    <a id="L76"></a>canBackquoteTest{`0123456789`, true},
    <a id="L77"></a>canBackquoteTest{`ABCDEFGHIJKLMNOPQRSTUVWXYZ`, true},
    <a id="L78"></a>canBackquoteTest{`abcdefghijklmnopqrstuvwxyz`, true},
    <a id="L79"></a>canBackquoteTest{`☺`, true},
<a id="L80"></a>}

<a id="L82"></a>func TestCanBackquote(t *testing.T) {
    <a id="L83"></a>for i := 0; i &lt; len(canbackquotetests); i++ {
        <a id="L84"></a>tt := canbackquotetests[i];
        <a id="L85"></a>if out := CanBackquote(tt.in); out != tt.out {
            <a id="L86"></a>t.Errorf(&#34;CanBackquote(%q) = %v, want %v&#34;, tt.in, out, tt.out)
        <a id="L87"></a>}
    <a id="L88"></a>}
<a id="L89"></a>}

<a id="L91"></a>var unquotetests = []quoteTest{
    <a id="L92"></a>quoteTest{`&#34;&#34;`, &#34;&#34;},
    <a id="L93"></a>quoteTest{`&#34;a&#34;`, &#34;a&#34;},
    <a id="L94"></a>quoteTest{`&#34;abc&#34;`, &#34;abc&#34;},
    <a id="L95"></a>quoteTest{`&#34;☺&#34;`, &#34;☺&#34;},
    <a id="L96"></a>quoteTest{`&#34;hello world&#34;`, &#34;hello world&#34;},
    <a id="L97"></a>quoteTest{`&#34;\xFF&#34;`, &#34;\xFF&#34;},
    <a id="L98"></a>quoteTest{`&#34;\377&#34;`, &#34;\377&#34;},
    <a id="L99"></a>quoteTest{`&#34;\u1234&#34;`, &#34;\u1234&#34;},
    <a id="L100"></a>quoteTest{`&#34;\U00010111&#34;`, &#34;\U00010111&#34;},
    <a id="L101"></a>quoteTest{`&#34;\U0001011111&#34;`, &#34;\U0001011111&#34;},
    <a id="L102"></a>quoteTest{`&#34;\a\b\f\n\r\t\v\\\&#34;&#34;`, &#34;\a\b\f\n\r\t\v\\\&#34;&#34;},
    <a id="L103"></a>quoteTest{`&#34;&#39;&#34;`, &#34;&#39;&#34;},

    <a id="L105"></a>quoteTest{`&#39;a&#39;`, &#34;a&#34;},
    <a id="L106"></a>quoteTest{`&#39;☹&#39;`, &#34;☹&#34;},
    <a id="L107"></a>quoteTest{`&#39;\a&#39;`, &#34;\a&#34;},
    <a id="L108"></a>quoteTest{`&#39;\x10&#39;`, &#34;\x10&#34;},
    <a id="L109"></a>quoteTest{`&#39;\377&#39;`, &#34;\377&#34;},
    <a id="L110"></a>quoteTest{`&#39;\u1234&#39;`, &#34;\u1234&#34;},
    <a id="L111"></a>quoteTest{`&#39;\U00010111&#39;`, &#34;\U00010111&#34;},
    <a id="L112"></a>quoteTest{`&#39;\t&#39;`, &#34;\t&#34;},
    <a id="L113"></a>quoteTest{`&#39; &#39;`, &#34; &#34;},
    <a id="L114"></a>quoteTest{`&#39;\&#39;&#39;`, &#34;&#39;&#34;},
    <a id="L115"></a>quoteTest{`&#39;&#34;&#39;`, &#34;\&#34;&#34;},

    <a id="L117"></a>quoteTest{&#34;``&#34;, ``},
    <a id="L118"></a>quoteTest{&#34;`a`&#34;, `a`},
    <a id="L119"></a>quoteTest{&#34;`abc`&#34;, `abc`},
    <a id="L120"></a>quoteTest{&#34;`☺`&#34;, `☺`},
    <a id="L121"></a>quoteTest{&#34;`hello world`&#34;, `hello world`},
    <a id="L122"></a>quoteTest{&#34;`\\xFF`&#34;, `\xFF`},
    <a id="L123"></a>quoteTest{&#34;`\\377`&#34;, `\377`},
    <a id="L124"></a>quoteTest{&#34;`\\`&#34;, `\`},
    <a id="L125"></a>quoteTest{&#34;`	`&#34;, `	`},
    <a id="L126"></a>quoteTest{&#34;` `&#34;, ` `},
<a id="L127"></a>}

<a id="L129"></a>var misquoted = []string{
    <a id="L130"></a>``,
    <a id="L131"></a>`&#34;`,
    <a id="L132"></a>`&#34;a`,
    <a id="L133"></a>`&#34;&#39;`,
    <a id="L134"></a>`b&#34;`,
    <a id="L135"></a>`&#34;\&#34;`,
    <a id="L136"></a>`&#39;\&#39;`,
    <a id="L137"></a>`&#39;ab&#39;`,
    <a id="L138"></a>`&#34;\x1!&#34;`,
    <a id="L139"></a>`&#34;\U12345678&#34;`,
    <a id="L140"></a>`&#34;\z&#34;`,
    <a id="L141"></a>&#34;`&#34;,
    <a id="L142"></a>&#34;`xxx&#34;,
    <a id="L143"></a>&#34;`\&#34;&#34;,
    <a id="L144"></a>`&#34;\&#39;&#34;`,
    <a id="L145"></a>`&#39;\&#34;&#39;`,
<a id="L146"></a>}

<a id="L148"></a>func TestUnquote(t *testing.T) {
    <a id="L149"></a>for i := 0; i &lt; len(unquotetests); i++ {
        <a id="L150"></a>tt := unquotetests[i];
        <a id="L151"></a>if out, err := Unquote(tt.in); err != nil &amp;&amp; out != tt.out {
            <a id="L152"></a>t.Errorf(&#34;Unquote(%#q) = %q, %v want %q, nil&#34;, tt.in, out, err, tt.out)
        <a id="L153"></a>}
    <a id="L154"></a>}

    <a id="L156"></a><span class="comment">// run the quote tests too, backward</span>
    <a id="L157"></a>for i := 0; i &lt; len(quotetests); i++ {
        <a id="L158"></a>tt := quotetests[i];
        <a id="L159"></a>if in, err := Unquote(tt.out); in != tt.in {
            <a id="L160"></a>t.Errorf(&#34;Unquote(%#q) = %q, %v, want %q, nil&#34;, tt.out, in, err, tt.in)
        <a id="L161"></a>}
    <a id="L162"></a>}

    <a id="L164"></a>for i := 0; i &lt; len(misquoted); i++ {
        <a id="L165"></a>s := misquoted[i];
        <a id="L166"></a>if out, err := Unquote(s); out != &#34;&#34; || err != os.EINVAL {
            <a id="L167"></a>t.Errorf(&#34;Unquote(%#q) = %q, %v want %q, %v&#34;, s, out, err, &#34;&#34;, os.EINVAL)
        <a id="L168"></a>}
    <a id="L169"></a>}
<a id="L170"></a>}
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
