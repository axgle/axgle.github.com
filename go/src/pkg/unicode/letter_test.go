<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/unicode/letter_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/unicode/letter_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package unicode_test

<a id="L7"></a>import (
    <a id="L8"></a>&#34;testing&#34;;
    <a id="L9"></a>. &#34;unicode&#34;;
<a id="L10"></a>)

<a id="L12"></a>var upperTest = []int{
    <a id="L13"></a>0x41,
    <a id="L14"></a>0xc0,
    <a id="L15"></a>0xd8,
    <a id="L16"></a>0x100,
    <a id="L17"></a>0x139,
    <a id="L18"></a>0x14a,
    <a id="L19"></a>0x178,
    <a id="L20"></a>0x181,
    <a id="L21"></a>0x376,
    <a id="L22"></a>0x3cf,
    <a id="L23"></a>0x1f2a,
    <a id="L24"></a>0x2102,
    <a id="L25"></a>0x2c00,
    <a id="L26"></a>0x2c10,
    <a id="L27"></a>0x2c20,
    <a id="L28"></a>0xa650,
    <a id="L29"></a>0xa722,
    <a id="L30"></a>0xff3a,
    <a id="L31"></a>0x10400,
    <a id="L32"></a>0x1d400,
    <a id="L33"></a>0x1d7ca,
<a id="L34"></a>}

<a id="L36"></a>var notupperTest = []int{
    <a id="L37"></a>0x40,
    <a id="L38"></a>0x5b,
    <a id="L39"></a>0x61,
    <a id="L40"></a>0x185,
    <a id="L41"></a>0x1b0,
    <a id="L42"></a>0x377,
    <a id="L43"></a>0x387,
    <a id="L44"></a>0x2150,
    <a id="L45"></a>0xffff,
    <a id="L46"></a>0x10000,
<a id="L47"></a>}

<a id="L49"></a>var letterTest = []int{
    <a id="L50"></a>0x41,
    <a id="L51"></a>0x61,
    <a id="L52"></a>0xaa,
    <a id="L53"></a>0xba,
    <a id="L54"></a>0xc8,
    <a id="L55"></a>0xdb,
    <a id="L56"></a>0xf9,
    <a id="L57"></a>0x2ec,
    <a id="L58"></a>0x535,
    <a id="L59"></a>0x6e6,
    <a id="L60"></a>0x93d,
    <a id="L61"></a>0xa15,
    <a id="L62"></a>0xb99,
    <a id="L63"></a>0xdc0,
    <a id="L64"></a>0xedd,
    <a id="L65"></a>0x1000,
    <a id="L66"></a>0x1200,
    <a id="L67"></a>0x1312,
    <a id="L68"></a>0x1401,
    <a id="L69"></a>0x1885,
    <a id="L70"></a>0x2c00,
    <a id="L71"></a>0xa800,
    <a id="L72"></a>0xf900,
    <a id="L73"></a>0xfa30,
    <a id="L74"></a>0xffda,
    <a id="L75"></a>0xffdc,
    <a id="L76"></a>0x10000,
    <a id="L77"></a>0x10300,
    <a id="L78"></a>0x10400,
    <a id="L79"></a>0x20000,
    <a id="L80"></a>0x2f800,
    <a id="L81"></a>0x2fa1d,
<a id="L82"></a>}

<a id="L84"></a>var notletterTest = []int{
    <a id="L85"></a>0x20,
    <a id="L86"></a>0x35,
    <a id="L87"></a>0x375,
    <a id="L88"></a>0x620,
    <a id="L89"></a>0x700,
    <a id="L90"></a>0xfffe,
    <a id="L91"></a>0x1ffff,
    <a id="L92"></a>0x10ffff,
<a id="L93"></a>}

<a id="L95"></a><span class="comment">// Contains all the special cased Latin-1 chars.</span>
<a id="L96"></a>var spaceTest = []int{
    <a id="L97"></a>0x09,
    <a id="L98"></a>0x0a,
    <a id="L99"></a>0x0b,
    <a id="L100"></a>0x0c,
    <a id="L101"></a>0x0d,
    <a id="L102"></a>0x20,
    <a id="L103"></a>0x85,
    <a id="L104"></a>0xA0,
    <a id="L105"></a>0x2000,
    <a id="L106"></a>0x3000,
<a id="L107"></a>}

<a id="L109"></a>type caseT struct {
    <a id="L110"></a>cas, in, out int;
<a id="L111"></a>}

<a id="L113"></a>var caseTest = []caseT{
    <a id="L114"></a><span class="comment">// errors</span>
    <a id="L115"></a>caseT{-1, &#39;\n&#39;, 0xFFFD},
    <a id="L116"></a>caseT{UpperCase, -1, -1},
    <a id="L117"></a>caseT{UpperCase, 1 &lt;&lt; 30, 1 &lt;&lt; 30},

    <a id="L119"></a><span class="comment">// ASCII (special-cased so test carefully)</span>
    <a id="L120"></a>caseT{UpperCase, &#39;\n&#39;, &#39;\n&#39;},
    <a id="L121"></a>caseT{UpperCase, &#39;a&#39;, &#39;A&#39;},
    <a id="L122"></a>caseT{UpperCase, &#39;A&#39;, &#39;A&#39;},
    <a id="L123"></a>caseT{UpperCase, &#39;7&#39;, &#39;7&#39;},
    <a id="L124"></a>caseT{LowerCase, &#39;\n&#39;, &#39;\n&#39;},
    <a id="L125"></a>caseT{LowerCase, &#39;a&#39;, &#39;a&#39;},
    <a id="L126"></a>caseT{LowerCase, &#39;A&#39;, &#39;a&#39;},
    <a id="L127"></a>caseT{LowerCase, &#39;7&#39;, &#39;7&#39;},
    <a id="L128"></a>caseT{TitleCase, &#39;\n&#39;, &#39;\n&#39;},
    <a id="L129"></a>caseT{TitleCase, &#39;a&#39;, &#39;A&#39;},
    <a id="L130"></a>caseT{TitleCase, &#39;A&#39;, &#39;A&#39;},
    <a id="L131"></a>caseT{TitleCase, &#39;7&#39;, &#39;7&#39;},

    <a id="L133"></a><span class="comment">// Latin-1: easy to read the tests!</span>
    <a id="L134"></a>caseT{UpperCase, 0x80, 0x80},
    <a id="L135"></a>caseT{UpperCase, &#39;Å&#39;, &#39;Å&#39;},
    <a id="L136"></a>caseT{UpperCase, &#39;å&#39;, &#39;Å&#39;},
    <a id="L137"></a>caseT{LowerCase, 0x80, 0x80},
    <a id="L138"></a>caseT{LowerCase, &#39;Å&#39;, &#39;å&#39;},
    <a id="L139"></a>caseT{LowerCase, &#39;å&#39;, &#39;å&#39;},
    <a id="L140"></a>caseT{TitleCase, 0x80, 0x80},
    <a id="L141"></a>caseT{TitleCase, &#39;Å&#39;, &#39;Å&#39;},
    <a id="L142"></a>caseT{TitleCase, &#39;å&#39;, &#39;Å&#39;},

    <a id="L144"></a><span class="comment">// 0131;LATIN SMALL LETTER DOTLESS I;Ll;0;L;;;;;N;;;0049;;0049</span>
    <a id="L145"></a>caseT{UpperCase, 0x0131, &#39;I&#39;},
    <a id="L146"></a>caseT{LowerCase, 0x0131, 0x0131},
    <a id="L147"></a>caseT{TitleCase, 0x0131, &#39;I&#39;},

    <a id="L149"></a><span class="comment">// 0133;LATIN SMALL LIGATURE IJ;Ll;0;L;&lt;compat&gt; 0069 006A;;;;N;LATIN SMALL LETTER I J;;0132;;0132</span>
    <a id="L150"></a>caseT{UpperCase, 0x0133, 0x0132},
    <a id="L151"></a>caseT{LowerCase, 0x0133, 0x0133},
    <a id="L152"></a>caseT{TitleCase, 0x0133, 0x0132},

    <a id="L154"></a><span class="comment">// 212A;KELVIN SIGN;Lu;0;L;004B;;;;N;DEGREES KELVIN;;;006B;</span>
    <a id="L155"></a>caseT{UpperCase, 0x212A, 0x212A},
    <a id="L156"></a>caseT{LowerCase, 0x212A, &#39;k&#39;},
    <a id="L157"></a>caseT{TitleCase, 0x212A, 0x212A},

    <a id="L159"></a><span class="comment">// From an UpperLower sequence</span>
    <a id="L160"></a><span class="comment">// A640;CYRILLIC CAPITAL LETTER ZEMLYA;Lu;0;L;;;;;N;;;;A641;</span>
    <a id="L161"></a>caseT{UpperCase, 0xA640, 0xA640},
    <a id="L162"></a>caseT{LowerCase, 0xA640, 0xA641},
    <a id="L163"></a>caseT{TitleCase, 0xA640, 0xA640},
    <a id="L164"></a><span class="comment">// A641;CYRILLIC SMALL LETTER ZEMLYA;Ll;0;L;;;;;N;;;A640;;A640</span>
    <a id="L165"></a>caseT{UpperCase, 0xA641, 0xA640},
    <a id="L166"></a>caseT{LowerCase, 0xA641, 0xA641},
    <a id="L167"></a>caseT{TitleCase, 0xA641, 0xA640},
    <a id="L168"></a><span class="comment">// A64E;CYRILLIC CAPITAL LETTER NEUTRAL YER;Lu;0;L;;;;;N;;;;A64F;</span>
    <a id="L169"></a>caseT{UpperCase, 0xA64E, 0xA64E},
    <a id="L170"></a>caseT{LowerCase, 0xA64E, 0xA64F},
    <a id="L171"></a>caseT{TitleCase, 0xA64E, 0xA64E},
    <a id="L172"></a><span class="comment">// A65F;CYRILLIC SMALL LETTER YN;Ll;0;L;;;;;N;;;A65E;;A65E</span>
    <a id="L173"></a>caseT{UpperCase, 0xA65F, 0xA65E},
    <a id="L174"></a>caseT{LowerCase, 0xA65F, 0xA65F},
    <a id="L175"></a>caseT{TitleCase, 0xA65F, 0xA65E},

    <a id="L177"></a><span class="comment">// From another UpperLower sequence</span>
    <a id="L178"></a><span class="comment">// 0139;LATIN CAPITAL LETTER L WITH ACUTE;Lu;0;L;004C 0301;;;;N;LATIN CAPITAL LETTER L ACUTE;;;013A;</span>
    <a id="L179"></a>caseT{UpperCase, 0x0139, 0x0139},
    <a id="L180"></a>caseT{LowerCase, 0x0139, 0x013A},
    <a id="L181"></a>caseT{TitleCase, 0x0139, 0x0139},
    <a id="L182"></a><span class="comment">// 013F;LATIN CAPITAL LETTER L WITH MIDDLE DOT;Lu;0;L;&lt;compat&gt; 004C 00B7;;;;N;;;;0140;</span>
    <a id="L183"></a>caseT{UpperCase, 0x013f, 0x013f},
    <a id="L184"></a>caseT{LowerCase, 0x013f, 0x0140},
    <a id="L185"></a>caseT{TitleCase, 0x013f, 0x013f},
    <a id="L186"></a><span class="comment">// 0148;LATIN SMALL LETTER N WITH CARON;Ll;0;L;006E 030C;;;;N;LATIN SMALL LETTER N HACEK;;0147;;0147</span>
    <a id="L187"></a>caseT{UpperCase, 0x0148, 0x0147},
    <a id="L188"></a>caseT{LowerCase, 0x0148, 0x0148},
    <a id="L189"></a>caseT{TitleCase, 0x0148, 0x0147},

    <a id="L191"></a><span class="comment">// Last block in the 5.1.0 table</span>
    <a id="L192"></a><span class="comment">// 10400;DESERET CAPITAL LETTER LONG I;Lu;0;L;;;;;N;;;;10428;</span>
    <a id="L193"></a>caseT{UpperCase, 0x10400, 0x10400},
    <a id="L194"></a>caseT{LowerCase, 0x10400, 0x10428},
    <a id="L195"></a>caseT{TitleCase, 0x10400, 0x10400},
    <a id="L196"></a><span class="comment">// 10427;DESERET CAPITAL LETTER EW;Lu;0;L;;;;;N;;;;1044F;</span>
    <a id="L197"></a>caseT{UpperCase, 0x10427, 0x10427},
    <a id="L198"></a>caseT{LowerCase, 0x10427, 0x1044F},
    <a id="L199"></a>caseT{TitleCase, 0x10427, 0x10427},
    <a id="L200"></a><span class="comment">// 10428;DESERET SMALL LETTER LONG I;Ll;0;L;;;;;N;;;10400;;10400</span>
    <a id="L201"></a>caseT{UpperCase, 0x10428, 0x10400},
    <a id="L202"></a>caseT{LowerCase, 0x10428, 0x10428},
    <a id="L203"></a>caseT{TitleCase, 0x10428, 0x10400},
    <a id="L204"></a><span class="comment">// 1044F;DESERET SMALL LETTER EW;Ll;0;L;;;;;N;;;10427;;10427</span>
    <a id="L205"></a>caseT{UpperCase, 0x1044F, 0x10427},
    <a id="L206"></a>caseT{LowerCase, 0x1044F, 0x1044F},
    <a id="L207"></a>caseT{TitleCase, 0x1044F, 0x10427},

    <a id="L209"></a><span class="comment">// First one not in the 5.1.0 table</span>
    <a id="L210"></a><span class="comment">// 10450;SHAVIAN LETTER PEEP;Lo;0;L;;;;;N;;;;;</span>
    <a id="L211"></a>caseT{UpperCase, 0x10450, 0x10450},
    <a id="L212"></a>caseT{LowerCase, 0x10450, 0x10450},
    <a id="L213"></a>caseT{TitleCase, 0x10450, 0x10450},
<a id="L214"></a>}

<a id="L216"></a>func TestIsLetter(t *testing.T) {
    <a id="L217"></a>for _, r := range upperTest {
        <a id="L218"></a>if !IsLetter(r) {
            <a id="L219"></a>t.Errorf(&#34;IsLetter(U+%04X) = false, want true\n&#34;, r)
        <a id="L220"></a>}
    <a id="L221"></a>}
    <a id="L222"></a>for _, r := range letterTest {
        <a id="L223"></a>if !IsLetter(r) {
            <a id="L224"></a>t.Errorf(&#34;IsLetter(U+%04X) = false, want true\n&#34;, r)
        <a id="L225"></a>}
    <a id="L226"></a>}
    <a id="L227"></a>for _, r := range notletterTest {
        <a id="L228"></a>if IsLetter(r) {
            <a id="L229"></a>t.Errorf(&#34;IsLetter(U+%04X) = true, want false\n&#34;, r)
        <a id="L230"></a>}
    <a id="L231"></a>}
<a id="L232"></a>}

<a id="L234"></a>func TestIsUpper(t *testing.T) {
    <a id="L235"></a>for _, r := range upperTest {
        <a id="L236"></a>if !IsUpper(r) {
            <a id="L237"></a>t.Errorf(&#34;IsUpper(U+%04X) = false, want true\n&#34;, r)
        <a id="L238"></a>}
    <a id="L239"></a>}
    <a id="L240"></a>for _, r := range notupperTest {
        <a id="L241"></a>if IsUpper(r) {
            <a id="L242"></a>t.Errorf(&#34;IsUpper(U+%04X) = true, want false\n&#34;, r)
        <a id="L243"></a>}
    <a id="L244"></a>}
    <a id="L245"></a>for _, r := range notletterTest {
        <a id="L246"></a>if IsUpper(r) {
            <a id="L247"></a>t.Errorf(&#34;IsUpper(U+%04X) = true, want false\n&#34;, r)
        <a id="L248"></a>}
    <a id="L249"></a>}
<a id="L250"></a>}

<a id="L252"></a>func caseString(c int) string {
    <a id="L253"></a>switch c {
    <a id="L254"></a>case UpperCase:
        <a id="L255"></a>return &#34;UpperCase&#34;
    <a id="L256"></a>case LowerCase:
        <a id="L257"></a>return &#34;LowerCase&#34;
    <a id="L258"></a>case TitleCase:
        <a id="L259"></a>return &#34;TitleCase&#34;
    <a id="L260"></a>}
    <a id="L261"></a>return &#34;ErrorCase&#34;;
<a id="L262"></a>}

<a id="L264"></a>func TestTo(t *testing.T) {
    <a id="L265"></a>for _, c := range caseTest {
        <a id="L266"></a>r := To(c.cas, c.in);
        <a id="L267"></a>if c.out != r {
            <a id="L268"></a>t.Errorf(&#34;To(U+%04X, %s) = U+%04X want U+%04X\n&#34;, c.in, caseString(c.cas), r, c.out)
        <a id="L269"></a>}
    <a id="L270"></a>}
<a id="L271"></a>}

<a id="L273"></a>func TestToUpperCase(t *testing.T) {
    <a id="L274"></a>for _, c := range caseTest {
        <a id="L275"></a>if c.cas != UpperCase {
            <a id="L276"></a>continue
        <a id="L277"></a>}
        <a id="L278"></a>r := ToUpper(c.in);
        <a id="L279"></a>if c.out != r {
            <a id="L280"></a>t.Errorf(&#34;ToUpper(U+%04X) = U+%04X want U+%04X\n&#34;, c.in, r, c.out)
        <a id="L281"></a>}
    <a id="L282"></a>}
<a id="L283"></a>}

<a id="L285"></a>func TestToLowerCase(t *testing.T) {
    <a id="L286"></a>for _, c := range caseTest {
        <a id="L287"></a>if c.cas != LowerCase {
            <a id="L288"></a>continue
        <a id="L289"></a>}
        <a id="L290"></a>r := ToLower(c.in);
        <a id="L291"></a>if c.out != r {
            <a id="L292"></a>t.Errorf(&#34;ToLower(U+%04X) = U+%04X want U+%04X\n&#34;, c.in, r, c.out)
        <a id="L293"></a>}
    <a id="L294"></a>}
<a id="L295"></a>}

<a id="L297"></a>func TestToTitleCase(t *testing.T) {
    <a id="L298"></a>for _, c := range caseTest {
        <a id="L299"></a>if c.cas != TitleCase {
            <a id="L300"></a>continue
        <a id="L301"></a>}
        <a id="L302"></a>r := ToTitle(c.in);
        <a id="L303"></a>if c.out != r {
            <a id="L304"></a>t.Errorf(&#34;ToTitle(U+%04X) = U+%04X want U+%04X\n&#34;, c.in, r, c.out)
        <a id="L305"></a>}
    <a id="L306"></a>}
<a id="L307"></a>}

<a id="L309"></a>func TestIsSpace(t *testing.T) {
    <a id="L310"></a>for _, c := range spaceTest {
        <a id="L311"></a>if !IsSpace(c) {
            <a id="L312"></a>t.Errorf(&#34;IsSpace(U+%04X) = false; want true&#34;, c)
        <a id="L313"></a>}
    <a id="L314"></a>}
    <a id="L315"></a>for _, c := range letterTest {
        <a id="L316"></a>if IsSpace(c) {
            <a id="L317"></a>t.Errorf(&#34;IsSpace(U+%04X) = true; want false&#34;, c)
        <a id="L318"></a>}
    <a id="L319"></a>}
<a id="L320"></a>}

<a id="L322"></a><span class="comment">// Check that the optimizations for IsLetter etc. agree with the tables.</span>
<a id="L323"></a><span class="comment">// We only need to check the Latin-1 range.</span>
<a id="L324"></a>func TestLetterOptimizations(t *testing.T) {
    <a id="L325"></a>for i := 0; i &lt; 0x100; i++ {
        <a id="L326"></a>if Is(Letter, i) != IsLetter(i) {
            <a id="L327"></a>t.Errorf(&#34;IsLetter(U+%04X) disagrees with Is(Letter)&#34;, i)
        <a id="L328"></a>}
        <a id="L329"></a>if Is(Upper, i) != IsUpper(i) {
            <a id="L330"></a>t.Errorf(&#34;IsUpper(U+%04X) disagrees with Is(Upper)&#34;, i)
        <a id="L331"></a>}
        <a id="L332"></a>if Is(Lower, i) != IsLower(i) {
            <a id="L333"></a>t.Errorf(&#34;IsLower(U+%04X) disagrees with Is(Lower)&#34;, i)
        <a id="L334"></a>}
        <a id="L335"></a>if Is(Title, i) != IsTitle(i) {
            <a id="L336"></a>t.Errorf(&#34;IsTitle(U+%04X) disagrees with Is(Title)&#34;, i)
        <a id="L337"></a>}
        <a id="L338"></a>if Is(White_Space, i) != IsSpace(i) {
            <a id="L339"></a>t.Errorf(&#34;IsSpace(U+%04X) disagrees with Is(White_Space)&#34;, i)
        <a id="L340"></a>}
        <a id="L341"></a>if To(UpperCase, i) != ToUpper(i) {
            <a id="L342"></a>t.Errorf(&#34;ToUpper(U+%04X) disagrees with To(Upper)&#34;, i)
        <a id="L343"></a>}
        <a id="L344"></a>if To(LowerCase, i) != ToLower(i) {
            <a id="L345"></a>t.Errorf(&#34;ToLower(U+%04X) disagrees with To(Lower)&#34;, i)
        <a id="L346"></a>}
        <a id="L347"></a>if To(TitleCase, i) != ToTitle(i) {
            <a id="L348"></a>t.Errorf(&#34;ToTitle(U+%04X) disagrees with To(Title)&#34;, i)
        <a id="L349"></a>}
    <a id="L350"></a>}
<a id="L351"></a>}
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
