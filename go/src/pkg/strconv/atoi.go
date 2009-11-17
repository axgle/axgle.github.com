<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strconv/atoi.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/strconv/atoi.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package strconv

<a id="L7"></a>import &#34;os&#34;

<a id="L9"></a>type NumError struct {
    <a id="L10"></a>Num   string;
    <a id="L11"></a>Error os.Error;
<a id="L12"></a>}

<a id="L14"></a>func (e *NumError) String() string { return &#34;parsing &#34; + e.Num + &#34;: &#34; + e.Error.String() }


<a id="L17"></a>func computeIntsize() uint {
    <a id="L18"></a>siz := uint(8);
    <a id="L19"></a>for 1&lt;&lt;siz != 0 {
        <a id="L20"></a>siz *= 2
    <a id="L21"></a>}
    <a id="L22"></a>return siz;
<a id="L23"></a>}

<a id="L25"></a>var IntSize = computeIntsize()

<a id="L27"></a><span class="comment">// Return the first number n such that n*base &gt;= 1&lt;&lt;64.</span>
<a id="L28"></a>func cutoff64(base int) uint64 {
    <a id="L29"></a>if base &lt; 2 {
        <a id="L30"></a>return 0
    <a id="L31"></a>}
    <a id="L32"></a>return (1&lt;&lt;64-1)/uint64(base) + 1;
<a id="L33"></a>}

<a id="L35"></a><span class="comment">// Btoui64 interprets a string s in an arbitrary base b (2 to 36)</span>
<a id="L36"></a><span class="comment">// and returns the corresponding value n.  If b == 0, the base</span>
<a id="L37"></a><span class="comment">// is taken from the string prefix: base 16 for &#34;0x&#34;, base 8 for &#34;0&#34;,</span>
<a id="L38"></a><span class="comment">// and base 10 otherwise.</span>
<a id="L39"></a><span class="comment">//</span>
<a id="L40"></a><span class="comment">// The errors that Btoui64 returns have concrete type *NumError</span>
<a id="L41"></a><span class="comment">// and include err.Num = s.  If s is empty or contains invalid</span>
<a id="L42"></a><span class="comment">// digits, err.Error = os.EINVAL; if the value corresponding</span>
<a id="L43"></a><span class="comment">// to s cannot be represented by a uint64, err.Error = os.ERANGE.</span>
<a id="L44"></a>func Btoui64(s string, b int) (n uint64, err os.Error) {
    <a id="L45"></a>s0 := s;
    <a id="L46"></a>switch {
    <a id="L47"></a>case len(s) &lt; 1:
        <a id="L48"></a>err = os.EINVAL;
        <a id="L49"></a>goto Error;

    <a id="L51"></a>case 2 &lt;= b &amp;&amp; b &lt;= 36:
        <a id="L52"></a><span class="comment">// valid base; nothing to do</span>

    <a id="L54"></a>case b == 0:
        <a id="L55"></a><span class="comment">// Look for octal, hex prefix.</span>
        <a id="L56"></a>switch {
        <a id="L57"></a>case s[0] == &#39;0&#39; &amp;&amp; len(s) &gt; 1 &amp;&amp; (s[1] == &#39;x&#39; || s[1] == &#39;X&#39;):
            <a id="L58"></a>b = 16;
            <a id="L59"></a>s = s[2:len(s)];
            <a id="L60"></a>if len(s) &lt; 1 {
                <a id="L61"></a>err = os.EINVAL;
                <a id="L62"></a>goto Error;
            <a id="L63"></a>}
        <a id="L64"></a>case s[0] == &#39;0&#39;:
            <a id="L65"></a>b = 8
        <a id="L66"></a>default:
            <a id="L67"></a>b = 10
        <a id="L68"></a>}

    <a id="L70"></a>default:
        <a id="L71"></a>err = os.ErrorString(&#34;invalid base &#34; + Itoa(b));
        <a id="L72"></a>goto Error;
    <a id="L73"></a>}

    <a id="L75"></a>n = 0;
    <a id="L76"></a>cutoff := cutoff64(b);

    <a id="L78"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L79"></a>var v byte;
        <a id="L80"></a>switch {
        <a id="L81"></a>case &#39;0&#39; &lt;= s[i] &amp;&amp; s[i] &lt;= &#39;9&#39;:
            <a id="L82"></a>v = s[i] - &#39;0&#39;
        <a id="L83"></a>case &#39;a&#39; &lt;= s[i] &amp;&amp; s[i] &lt;= &#39;z&#39;:
            <a id="L84"></a>v = s[i] - &#39;a&#39; + 10
        <a id="L85"></a>case &#39;A&#39; &lt;= s[i] &amp;&amp; s[i] &lt;= &#39;Z&#39;:
            <a id="L86"></a>v = s[i] - &#39;A&#39; + 10
        <a id="L87"></a>default:
            <a id="L88"></a>n = 0;
            <a id="L89"></a>err = os.EINVAL;
            <a id="L90"></a>goto Error;
        <a id="L91"></a>}
        <a id="L92"></a>if int(v) &gt;= b {
            <a id="L93"></a>n = 0;
            <a id="L94"></a>err = os.EINVAL;
            <a id="L95"></a>goto Error;
        <a id="L96"></a>}

        <a id="L98"></a>if n &gt;= cutoff {
            <a id="L99"></a><span class="comment">// n*b overflows</span>
            <a id="L100"></a>n = 1&lt;&lt;64 - 1;
            <a id="L101"></a>err = os.ERANGE;
            <a id="L102"></a>goto Error;
        <a id="L103"></a>}
        <a id="L104"></a>n *= uint64(b);

        <a id="L106"></a>n1 := n + uint64(v);
        <a id="L107"></a>if n1 &lt; n {
            <a id="L108"></a><span class="comment">// n+v overflows</span>
            <a id="L109"></a>n = 1&lt;&lt;64 - 1;
            <a id="L110"></a>err = os.ERANGE;
            <a id="L111"></a>goto Error;
        <a id="L112"></a>}
        <a id="L113"></a>n = n1;
    <a id="L114"></a>}

    <a id="L116"></a>return n, nil;

<a id="L118"></a>Error:
    <a id="L119"></a>return n, &amp;NumError{s0, err};
<a id="L120"></a>}

<a id="L122"></a><span class="comment">// Atoui64 interprets a string s as a decimal number and</span>
<a id="L123"></a><span class="comment">// returns the corresponding value n.</span>
<a id="L124"></a><span class="comment">//</span>
<a id="L125"></a><span class="comment">// Atoui64 returns err == os.EINVAL if s is empty or contains invalid digits.</span>
<a id="L126"></a><span class="comment">// It returns err == os.ERANGE if s cannot be represented by a uint64.</span>
<a id="L127"></a>func Atoui64(s string) (n uint64, err os.Error) {
    <a id="L128"></a>return Btoui64(s, 10)
<a id="L129"></a>}

<a id="L131"></a><span class="comment">// Btoi64 is like Btoui64 but allows signed numbers and</span>
<a id="L132"></a><span class="comment">// returns its result in an int64.</span>
<a id="L133"></a>func Btoi64(s string, base int) (i int64, err os.Error) {
    <a id="L134"></a><span class="comment">// Empty string bad.</span>
    <a id="L135"></a>if len(s) == 0 {
        <a id="L136"></a>return 0, &amp;NumError{s, os.EINVAL}
    <a id="L137"></a>}

    <a id="L139"></a><span class="comment">// Pick off leading sign.</span>
    <a id="L140"></a>s0 := s;
    <a id="L141"></a>neg := false;
    <a id="L142"></a>if s[0] == &#39;+&#39; {
        <a id="L143"></a>s = s[1:len(s)]
    <a id="L144"></a>} else if s[0] == &#39;-&#39; {
        <a id="L145"></a>neg = true;
        <a id="L146"></a>s = s[1:len(s)];
    <a id="L147"></a>}

    <a id="L149"></a><span class="comment">// Convert unsigned and check range.</span>
    <a id="L150"></a>var un uint64;
    <a id="L151"></a>un, err = Btoui64(s, base);
    <a id="L152"></a>if err != nil &amp;&amp; err.(*NumError).Error != os.ERANGE {
        <a id="L153"></a>err.(*NumError).Num = s0;
        <a id="L154"></a>return 0, err;
    <a id="L155"></a>}
    <a id="L156"></a>if !neg &amp;&amp; un &gt;= 1&lt;&lt;63 {
        <a id="L157"></a>return 1&lt;&lt;63 - 1, &amp;NumError{s0, os.ERANGE}
    <a id="L158"></a>}
    <a id="L159"></a>if neg &amp;&amp; un &gt; 1&lt;&lt;63 {
        <a id="L160"></a>return -1 &lt;&lt; 63, &amp;NumError{s0, os.ERANGE}
    <a id="L161"></a>}
    <a id="L162"></a>n := int64(un);
    <a id="L163"></a>if neg {
        <a id="L164"></a>n = -n
    <a id="L165"></a>}
    <a id="L166"></a>return n, nil;
<a id="L167"></a>}

<a id="L169"></a><span class="comment">// Atoi64 is like Atoui64 but allows signed numbers and</span>
<a id="L170"></a><span class="comment">// returns its result in an int64.</span>
<a id="L171"></a>func Atoi64(s string) (i int64, err os.Error) { return Btoi64(s, 10) }


<a id="L174"></a><span class="comment">// Atoui is like Atoui64 but returns its result as a uint.</span>
<a id="L175"></a>func Atoui(s string) (i uint, err os.Error) {
    <a id="L176"></a>i1, e1 := Atoui64(s);
    <a id="L177"></a>if e1 != nil &amp;&amp; e1.(*NumError).Error != os.ERANGE {
        <a id="L178"></a>return 0, e1
    <a id="L179"></a>}
    <a id="L180"></a>i = uint(i1);
    <a id="L181"></a>if uint64(i) != i1 {
        <a id="L182"></a>return ^uint(0), &amp;NumError{s, os.ERANGE}
    <a id="L183"></a>}
    <a id="L184"></a>return i, nil;
<a id="L185"></a>}

<a id="L187"></a><span class="comment">// Atoi is like Atoi64 but returns its result as an int.</span>
<a id="L188"></a>func Atoi(s string) (i int, err os.Error) {
    <a id="L189"></a>i1, e1 := Atoi64(s);
    <a id="L190"></a>if e1 != nil &amp;&amp; e1.(*NumError).Error != os.ERANGE {
        <a id="L191"></a>return 0, e1
    <a id="L192"></a>}
    <a id="L193"></a>i = int(i1);
    <a id="L194"></a>if int64(i) != i1 {
        <a id="L195"></a>if i1 &lt; 0 {
            <a id="L196"></a>return -1 &lt;&lt; (IntSize - 1), &amp;NumError{s, os.ERANGE}
        <a id="L197"></a>}
        <a id="L198"></a>return 1&lt;&lt;(IntSize-1) - 1, &amp;NumError{s, os.ERANGE};
    <a id="L199"></a>}
    <a id="L200"></a>return i, nil;
<a id="L201"></a>}
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
