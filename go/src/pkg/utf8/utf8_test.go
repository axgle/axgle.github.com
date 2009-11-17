<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/utf8/utf8_test.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/utf8/utf8_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package utf8_test

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;strings&#34;;
    <a id="L10"></a>&#34;testing&#34;;
    <a id="L11"></a>. &#34;utf8&#34;;
<a id="L12"></a>)

<a id="L14"></a>type Utf8Map struct {
    <a id="L15"></a>rune int;
    <a id="L16"></a>str  string;
<a id="L17"></a>}

<a id="L19"></a>var utf8map = []Utf8Map{
    <a id="L20"></a>Utf8Map{0x0000, &#34;\x00&#34;},
    <a id="L21"></a>Utf8Map{0x0001, &#34;\x01&#34;},
    <a id="L22"></a>Utf8Map{0x007e, &#34;\x7e&#34;},
    <a id="L23"></a>Utf8Map{0x007f, &#34;\x7f&#34;},
    <a id="L24"></a>Utf8Map{0x0080, &#34;\xc2\x80&#34;},
    <a id="L25"></a>Utf8Map{0x0081, &#34;\xc2\x81&#34;},
    <a id="L26"></a>Utf8Map{0x00bf, &#34;\xc2\xbf&#34;},
    <a id="L27"></a>Utf8Map{0x00c0, &#34;\xc3\x80&#34;},
    <a id="L28"></a>Utf8Map{0x00c1, &#34;\xc3\x81&#34;},
    <a id="L29"></a>Utf8Map{0x00c8, &#34;\xc3\x88&#34;},
    <a id="L30"></a>Utf8Map{0x00d0, &#34;\xc3\x90&#34;},
    <a id="L31"></a>Utf8Map{0x00e0, &#34;\xc3\xa0&#34;},
    <a id="L32"></a>Utf8Map{0x00f0, &#34;\xc3\xb0&#34;},
    <a id="L33"></a>Utf8Map{0x00f8, &#34;\xc3\xb8&#34;},
    <a id="L34"></a>Utf8Map{0x00ff, &#34;\xc3\xbf&#34;},
    <a id="L35"></a>Utf8Map{0x0100, &#34;\xc4\x80&#34;},
    <a id="L36"></a>Utf8Map{0x07ff, &#34;\xdf\xbf&#34;},
    <a id="L37"></a>Utf8Map{0x0800, &#34;\xe0\xa0\x80&#34;},
    <a id="L38"></a>Utf8Map{0x0801, &#34;\xe0\xa0\x81&#34;},
    <a id="L39"></a>Utf8Map{0xfffe, &#34;\xef\xbf\xbe&#34;},
    <a id="L40"></a>Utf8Map{0xffff, &#34;\xef\xbf\xbf&#34;},
    <a id="L41"></a>Utf8Map{0x10000, &#34;\xf0\x90\x80\x80&#34;},
    <a id="L42"></a>Utf8Map{0x10001, &#34;\xf0\x90\x80\x81&#34;},
    <a id="L43"></a>Utf8Map{0x10fffe, &#34;\xf4\x8f\xbf\xbe&#34;},
    <a id="L44"></a>Utf8Map{0x10ffff, &#34;\xf4\x8f\xbf\xbf&#34;},
<a id="L45"></a>}

<a id="L47"></a><span class="comment">// strings.Bytes with one extra byte at end</span>
<a id="L48"></a>func makeBytes(s string) []byte {
    <a id="L49"></a>s += &#34;\x00&#34;;
    <a id="L50"></a>b := strings.Bytes(s);
    <a id="L51"></a>return b[0 : len(s)-1];
<a id="L52"></a>}

<a id="L54"></a>func TestFullRune(t *testing.T) {
    <a id="L55"></a>for i := 0; i &lt; len(utf8map); i++ {
        <a id="L56"></a>m := utf8map[i];
        <a id="L57"></a>b := makeBytes(m.str);
        <a id="L58"></a>if !FullRune(b) {
            <a id="L59"></a>t.Errorf(&#34;FullRune(%q) (rune %04x) = false, want true&#34;, b, m.rune)
        <a id="L60"></a>}
        <a id="L61"></a>s := m.str;
        <a id="L62"></a>if !FullRuneInString(s) {
            <a id="L63"></a>t.Errorf(&#34;FullRuneInString(%q) (rune %04x) = false, want true&#34;, s, m.rune)
        <a id="L64"></a>}
        <a id="L65"></a>b1 := b[0 : len(b)-1];
        <a id="L66"></a>if FullRune(b1) {
            <a id="L67"></a>t.Errorf(&#34;FullRune(%q) = true, want false&#34;, b1)
        <a id="L68"></a>}
        <a id="L69"></a>s1 := string(b1);
        <a id="L70"></a>if FullRuneInString(s1) {
            <a id="L71"></a>t.Errorf(&#34;FullRune(%q) = true, want false&#34;, s1)
        <a id="L72"></a>}
    <a id="L73"></a>}
<a id="L74"></a>}

<a id="L76"></a>func TestEncodeRune(t *testing.T) {
    <a id="L77"></a>for i := 0; i &lt; len(utf8map); i++ {
        <a id="L78"></a>m := utf8map[i];
        <a id="L79"></a>b := makeBytes(m.str);
        <a id="L80"></a>var buf [10]byte;
        <a id="L81"></a>n := EncodeRune(m.rune, &amp;buf);
        <a id="L82"></a>b1 := buf[0:n];
        <a id="L83"></a>if !bytes.Equal(b, b1) {
            <a id="L84"></a>t.Errorf(&#34;EncodeRune(0x%04x) = %q want %q&#34;, m.rune, b1, b)
        <a id="L85"></a>}
    <a id="L86"></a>}
<a id="L87"></a>}

<a id="L89"></a>func TestDecodeRune(t *testing.T) {
    <a id="L90"></a>for i := 0; i &lt; len(utf8map); i++ {
        <a id="L91"></a>m := utf8map[i];
        <a id="L92"></a>b := makeBytes(m.str);
        <a id="L93"></a>rune, size := DecodeRune(b);
        <a id="L94"></a>if rune != m.rune || size != len(b) {
            <a id="L95"></a>t.Errorf(&#34;DecodeRune(%q) = 0x%04x, %d want 0x%04x, %d&#34;, b, rune, size, m.rune, len(b))
        <a id="L96"></a>}
        <a id="L97"></a>s := m.str;
        <a id="L98"></a>rune, size = DecodeRuneInString(s);
        <a id="L99"></a>if rune != m.rune || size != len(b) {
            <a id="L100"></a>t.Errorf(&#34;DecodeRune(%q) = 0x%04x, %d want 0x%04x, %d&#34;, s, rune, size, m.rune, len(b))
        <a id="L101"></a>}

        <a id="L103"></a><span class="comment">// there&#39;s an extra byte that bytes left behind - make sure trailing byte works</span>
        <a id="L104"></a>rune, size = DecodeRune(b[0:cap(b)]);
        <a id="L105"></a>if rune != m.rune || size != len(b) {
            <a id="L106"></a>t.Errorf(&#34;DecodeRune(%q) = 0x%04x, %d want 0x%04x, %d&#34;, b, rune, size, m.rune, len(b))
        <a id="L107"></a>}
        <a id="L108"></a>s = m.str + &#34;\x00&#34;;
        <a id="L109"></a>rune, size = DecodeRuneInString(s);
        <a id="L110"></a>if rune != m.rune || size != len(b) {
            <a id="L111"></a>t.Errorf(&#34;DecodeRuneInString(%q) = 0x%04x, %d want 0x%04x, %d&#34;, s, rune, size, m.rune, len(b))
        <a id="L112"></a>}

        <a id="L114"></a><span class="comment">// make sure missing bytes fail</span>
        <a id="L115"></a>wantsize := 1;
        <a id="L116"></a>if wantsize &gt;= len(b) {
            <a id="L117"></a>wantsize = 0
        <a id="L118"></a>}
        <a id="L119"></a>rune, size = DecodeRune(b[0 : len(b)-1]);
        <a id="L120"></a>if rune != RuneError || size != wantsize {
            <a id="L121"></a>t.Errorf(&#34;DecodeRune(%q) = 0x%04x, %d want 0x%04x, %d&#34;, b[0:len(b)-1], rune, size, RuneError, wantsize)
        <a id="L122"></a>}
        <a id="L123"></a>s = m.str[0 : len(m.str)-1];
        <a id="L124"></a>rune, size = DecodeRuneInString(s);
        <a id="L125"></a>if rune != RuneError || size != wantsize {
            <a id="L126"></a>t.Errorf(&#34;DecodeRuneInString(%q) = 0x%04x, %d want 0x%04x, %d&#34;, s, rune, size, RuneError, wantsize)
        <a id="L127"></a>}

        <a id="L129"></a><span class="comment">// make sure bad sequences fail</span>
        <a id="L130"></a>if len(b) == 1 {
            <a id="L131"></a>b[0] = 0x80
        <a id="L132"></a>} else {
            <a id="L133"></a>b[len(b)-1] = 0x7F
        <a id="L134"></a>}
        <a id="L135"></a>rune, size = DecodeRune(b);
        <a id="L136"></a>if rune != RuneError || size != 1 {
            <a id="L137"></a>t.Errorf(&#34;DecodeRune(%q) = 0x%04x, %d want 0x%04x, %d&#34;, b, rune, size, RuneError, 1)
        <a id="L138"></a>}
        <a id="L139"></a>s = string(b);
        <a id="L140"></a>rune, size = DecodeRune(b);
        <a id="L141"></a>if rune != RuneError || size != 1 {
            <a id="L142"></a>t.Errorf(&#34;DecodeRuneInString(%q) = 0x%04x, %d want 0x%04x, %d&#34;, s, rune, size, RuneError, 1)
        <a id="L143"></a>}
    <a id="L144"></a>}
<a id="L145"></a>}

<a id="L147"></a>type RuneCountTest struct {
    <a id="L148"></a>in  string;
    <a id="L149"></a>out int;
<a id="L150"></a>}

<a id="L152"></a>var runecounttests = []RuneCountTest{
    <a id="L153"></a>RuneCountTest{&#34;abcd&#34;, 4},
    <a id="L154"></a>RuneCountTest{&#34;☺☻☹&#34;, 3},
    <a id="L155"></a>RuneCountTest{&#34;1,2,3,4&#34;, 7},
    <a id="L156"></a>RuneCountTest{&#34;\xe2\x00&#34;, 2},
<a id="L157"></a>}

<a id="L159"></a>func TestRuneCount(t *testing.T) {
    <a id="L160"></a>for i := 0; i &lt; len(runecounttests); i++ {
        <a id="L161"></a>tt := runecounttests[i];
        <a id="L162"></a>if out := RuneCountInString(tt.in); out != tt.out {
            <a id="L163"></a>t.Errorf(&#34;RuneCountInString(%q) = %d, want %d&#34;, tt.in, out, tt.out)
        <a id="L164"></a>}
        <a id="L165"></a>if out := RuneCount(makeBytes(tt.in)); out != tt.out {
            <a id="L166"></a>t.Errorf(&#34;RuneCount(%q) = %d, want %d&#34;, tt.in, out, tt.out)
        <a id="L167"></a>}
    <a id="L168"></a>}
<a id="L169"></a>}
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
