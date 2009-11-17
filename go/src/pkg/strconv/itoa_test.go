<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strconv/itoa_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/strconv/itoa_test.go</h1>

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
    <a id="L8"></a>. &#34;strconv&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a>type itob64Test struct {
    <a id="L13"></a>in   int64;
    <a id="L14"></a>base uint;
    <a id="L15"></a>out  string;
<a id="L16"></a>}

<a id="L18"></a>var itob64tests = []itob64Test{
    <a id="L19"></a>itob64Test{0, 10, &#34;0&#34;},
    <a id="L20"></a>itob64Test{1, 10, &#34;1&#34;},
    <a id="L21"></a>itob64Test{-1, 10, &#34;-1&#34;},
    <a id="L22"></a>itob64Test{12345678, 10, &#34;12345678&#34;},
    <a id="L23"></a>itob64Test{-987654321, 10, &#34;-987654321&#34;},
    <a id="L24"></a>itob64Test{1&lt;&lt;31 - 1, 10, &#34;2147483647&#34;},
    <a id="L25"></a>itob64Test{-1&lt;&lt;31 + 1, 10, &#34;-2147483647&#34;},
    <a id="L26"></a>itob64Test{1 &lt;&lt; 31, 10, &#34;2147483648&#34;},
    <a id="L27"></a>itob64Test{-1 &lt;&lt; 31, 10, &#34;-2147483648&#34;},
    <a id="L28"></a>itob64Test{1&lt;&lt;31 + 1, 10, &#34;2147483649&#34;},
    <a id="L29"></a>itob64Test{-1&lt;&lt;31 - 1, 10, &#34;-2147483649&#34;},
    <a id="L30"></a>itob64Test{1&lt;&lt;32 - 1, 10, &#34;4294967295&#34;},
    <a id="L31"></a>itob64Test{-1&lt;&lt;32 + 1, 10, &#34;-4294967295&#34;},
    <a id="L32"></a>itob64Test{1 &lt;&lt; 32, 10, &#34;4294967296&#34;},
    <a id="L33"></a>itob64Test{-1 &lt;&lt; 32, 10, &#34;-4294967296&#34;},
    <a id="L34"></a>itob64Test{1&lt;&lt;32 + 1, 10, &#34;4294967297&#34;},
    <a id="L35"></a>itob64Test{-1&lt;&lt;32 - 1, 10, &#34;-4294967297&#34;},
    <a id="L36"></a>itob64Test{1 &lt;&lt; 50, 10, &#34;1125899906842624&#34;},
    <a id="L37"></a>itob64Test{1&lt;&lt;63 - 1, 10, &#34;9223372036854775807&#34;},
    <a id="L38"></a>itob64Test{-1&lt;&lt;63 + 1, 10, &#34;-9223372036854775807&#34;},
    <a id="L39"></a>itob64Test{-1 &lt;&lt; 63, 10, &#34;-9223372036854775808&#34;},

    <a id="L41"></a>itob64Test{0, 2, &#34;0&#34;},
    <a id="L42"></a>itob64Test{10, 2, &#34;1010&#34;},
    <a id="L43"></a>itob64Test{-1, 2, &#34;-1&#34;},
    <a id="L44"></a>itob64Test{1 &lt;&lt; 15, 2, &#34;1000000000000000&#34;},

    <a id="L46"></a>itob64Test{-8, 8, &#34;-10&#34;},
    <a id="L47"></a>itob64Test{057635436545, 8, &#34;57635436545&#34;},
    <a id="L48"></a>itob64Test{1 &lt;&lt; 24, 8, &#34;100000000&#34;},

    <a id="L50"></a>itob64Test{16, 16, &#34;10&#34;},
    <a id="L51"></a>itob64Test{-0x123456789abcdef, 16, &#34;-123456789abcdef&#34;},
    <a id="L52"></a>itob64Test{1&lt;&lt;63 - 1, 16, &#34;7fffffffffffffff&#34;},

    <a id="L54"></a>itob64Test{16, 17, &#34;g&#34;},
    <a id="L55"></a>itob64Test{25, 25, &#34;10&#34;},
    <a id="L56"></a>itob64Test{(((((17*35+24)*35+21)*35+34)*35+12)*35+24)*35 + 32, 35, &#34;holycow&#34;},
    <a id="L57"></a>itob64Test{(((((17*36+24)*36+21)*36+34)*36+12)*36+24)*36 + 32, 36, &#34;holycow&#34;},
<a id="L58"></a>}

<a id="L60"></a>func TestItoa(t *testing.T) {
    <a id="L61"></a>for _, test := range itob64tests {
        <a id="L62"></a>s := Itob64(test.in, test.base);
        <a id="L63"></a>if s != test.out {
            <a id="L64"></a>t.Errorf(&#34;Itob64(%v, %v) = %v want %v\n&#34;,
                <a id="L65"></a>test.in, test.base, s, test.out)
        <a id="L66"></a>}

        <a id="L68"></a>if test.in &gt;= 0 {
            <a id="L69"></a>s := Uitob64(uint64(test.in), test.base);
            <a id="L70"></a>if s != test.out {
                <a id="L71"></a>t.Errorf(&#34;Uitob64(%v, %v) = %v want %v\n&#34;,
                    <a id="L72"></a>test.in, test.base, s, test.out)
            <a id="L73"></a>}
        <a id="L74"></a>}

        <a id="L76"></a>if int64(int(test.in)) == test.in {
            <a id="L77"></a>s := Itob(int(test.in), test.base);
            <a id="L78"></a>if s != test.out {
                <a id="L79"></a>t.Errorf(&#34;Itob(%v, %v) = %v want %v\n&#34;,
                    <a id="L80"></a>test.in, test.base, s, test.out)
            <a id="L81"></a>}

            <a id="L83"></a>if test.in &gt;= 0 {
                <a id="L84"></a>s := Uitob(uint(test.in), test.base);
                <a id="L85"></a>if s != test.out {
                    <a id="L86"></a>t.Errorf(&#34;Uitob(%v, %v) = %v want %v\n&#34;,
                        <a id="L87"></a>test.in, test.base, s, test.out)
                <a id="L88"></a>}
            <a id="L89"></a>}
        <a id="L90"></a>}

        <a id="L92"></a>if test.base == 10 {
            <a id="L93"></a>s := Itoa64(test.in);
            <a id="L94"></a>if s != test.out {
                <a id="L95"></a>t.Errorf(&#34;Itoa64(%v) = %v want %v\n&#34;,
                    <a id="L96"></a>test.in, s, test.out)
            <a id="L97"></a>}

            <a id="L99"></a>if test.in &gt;= 0 {
                <a id="L100"></a>s := Uitob64(uint64(test.in), test.base);
                <a id="L101"></a>if s != test.out {
                    <a id="L102"></a>t.Errorf(&#34;Uitob64(%v, %v) = %v want %v\n&#34;,
                        <a id="L103"></a>test.in, test.base, s, test.out)
                <a id="L104"></a>}
            <a id="L105"></a>}

            <a id="L107"></a>if int64(int(test.in)) == test.in {
                <a id="L108"></a>s := Itoa(int(test.in));
                <a id="L109"></a>if s != test.out {
                    <a id="L110"></a>t.Errorf(&#34;Itoa(%v) = %v want %v\n&#34;,
                        <a id="L111"></a>test.in, s, test.out)
                <a id="L112"></a>}

                <a id="L114"></a>if test.in &gt;= 0 {
                    <a id="L115"></a>s := Uitoa(uint(test.in));
                    <a id="L116"></a>if s != test.out {
                        <a id="L117"></a>t.Errorf(&#34;Uitoa(%v) = %v want %v\n&#34;,
                            <a id="L118"></a>test.in, s, test.out)
                    <a id="L119"></a>}
                <a id="L120"></a>}
            <a id="L121"></a>}
        <a id="L122"></a>}
    <a id="L123"></a>}
<a id="L124"></a>}

<a id="L126"></a>type uitob64Test struct {
    <a id="L127"></a>in   uint64;
    <a id="L128"></a>base uint;
    <a id="L129"></a>out  string;
<a id="L130"></a>}

<a id="L132"></a>var uitob64tests = []uitob64Test{
    <a id="L133"></a>uitob64Test{1&lt;&lt;63 - 1, 10, &#34;9223372036854775807&#34;},
    <a id="L134"></a>uitob64Test{1 &lt;&lt; 63, 10, &#34;9223372036854775808&#34;},
    <a id="L135"></a>uitob64Test{1&lt;&lt;63 + 1, 10, &#34;9223372036854775809&#34;},
    <a id="L136"></a>uitob64Test{1&lt;&lt;64 - 2, 10, &#34;18446744073709551614&#34;},
    <a id="L137"></a>uitob64Test{1&lt;&lt;64 - 1, 10, &#34;18446744073709551615&#34;},
<a id="L138"></a>}

<a id="L140"></a>func TestUitoa(t *testing.T) {
    <a id="L141"></a>for _, test := range uitob64tests {
        <a id="L142"></a>s := Uitob64(test.in, test.base);
        <a id="L143"></a>if s != test.out {
            <a id="L144"></a>t.Errorf(&#34;Uitob64(%v, %v) = %v want %v\n&#34;,
                <a id="L145"></a>test.in, test.base, s, test.out)
        <a id="L146"></a>}

        <a id="L148"></a>if uint64(uint(test.in)) == test.in {
            <a id="L149"></a>s := Uitob(uint(test.in), test.base);
            <a id="L150"></a>if s != test.out {
                <a id="L151"></a>t.Errorf(&#34;Uitob(%v, %v) = %v want %v\n&#34;,
                    <a id="L152"></a>test.in, test.base, s, test.out)
            <a id="L153"></a>}
        <a id="L154"></a>}

        <a id="L156"></a>if test.base == 10 {
            <a id="L157"></a>s := Uitoa64(test.in);
            <a id="L158"></a>if s != test.out {
                <a id="L159"></a>t.Errorf(&#34;Uitoa64(%v) = %v want %v\n&#34;,
                    <a id="L160"></a>test.in, s, test.out)
            <a id="L161"></a>}

            <a id="L163"></a>if uint64(uint(test.in)) == test.in {
                <a id="L164"></a>s := Uitoa(uint(test.in));
                <a id="L165"></a>if s != test.out {
                    <a id="L166"></a>t.Errorf(&#34;Uitoa(%v) = %v want %v\n&#34;,
                        <a id="L167"></a>test.in, s, test.out)
                <a id="L168"></a>}
            <a id="L169"></a>}
        <a id="L170"></a>}
    <a id="L171"></a>}
<a id="L172"></a>}
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
