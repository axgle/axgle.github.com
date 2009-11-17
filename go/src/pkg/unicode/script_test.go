<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/unicode/script_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/unicode/script_test.go</h1>

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

<a id="L12"></a>type T struct {
    <a id="L13"></a>rune   int;
    <a id="L14"></a>script string;
<a id="L15"></a>}

<a id="L17"></a><span class="comment">// Hand-chosen tests from Unicode 5.1.0, mostly to discover when new</span>
<a id="L18"></a><span class="comment">// scripts and categories arise.</span>
<a id="L19"></a>var inTest = []T{
    <a id="L20"></a>T{0x06e2, &#34;Arabic&#34;},
    <a id="L21"></a>T{0x0567, &#34;Armenian&#34;},
    <a id="L22"></a>T{0x1b37, &#34;Balinese&#34;},
    <a id="L23"></a>T{0x09c2, &#34;Bengali&#34;},
    <a id="L24"></a>T{0x3115, &#34;Bopomofo&#34;},
    <a id="L25"></a>T{0x282d, &#34;Braille&#34;},
    <a id="L26"></a>T{0x1a1a, &#34;Buginese&#34;},
    <a id="L27"></a>T{0x1747, &#34;Buhid&#34;},
    <a id="L28"></a>T{0x156d, &#34;Canadian_Aboriginal&#34;},
    <a id="L29"></a>T{0x102a9, &#34;Carian&#34;},
    <a id="L30"></a>T{0xaa4d, &#34;Cham&#34;},
    <a id="L31"></a>T{0x13c2, &#34;Cherokee&#34;},
    <a id="L32"></a>T{0x0020, &#34;Common&#34;},
    <a id="L33"></a>T{0x1d4a5, &#34;Common&#34;},
    <a id="L34"></a>T{0x2cfc, &#34;Coptic&#34;},
    <a id="L35"></a>T{0x12420, &#34;Cuneiform&#34;},
    <a id="L36"></a>T{0x1080c, &#34;Cypriot&#34;},
    <a id="L37"></a>T{0xa663, &#34;Cyrillic&#34;},
    <a id="L38"></a>T{0x10430, &#34;Deseret&#34;},
    <a id="L39"></a>T{0x094a, &#34;Devanagari&#34;},
    <a id="L40"></a>T{0x1271, &#34;Ethiopic&#34;},
    <a id="L41"></a>T{0x10fc, &#34;Georgian&#34;},
    <a id="L42"></a>T{0x2c40, &#34;Glagolitic&#34;},
    <a id="L43"></a>T{0x10347, &#34;Gothic&#34;},
    <a id="L44"></a>T{0x03ae, &#34;Greek&#34;},
    <a id="L45"></a>T{0x0abf, &#34;Gujarati&#34;},
    <a id="L46"></a>T{0x0a24, &#34;Gurmukhi&#34;},
    <a id="L47"></a>T{0x3028, &#34;Han&#34;},
    <a id="L48"></a>T{0x11b8, &#34;Hangul&#34;},
    <a id="L49"></a>T{0x1727, &#34;Hanunoo&#34;},
    <a id="L50"></a>T{0x05a0, &#34;Hebrew&#34;},
    <a id="L51"></a>T{0x3058, &#34;Hiragana&#34;},
    <a id="L52"></a>T{0x20e6, &#34;Inherited&#34;},
    <a id="L53"></a>T{0x0cbd, &#34;Kannada&#34;},
    <a id="L54"></a>T{0x30a6, &#34;Katakana&#34;},
    <a id="L55"></a>T{0xa928, &#34;Kayah_Li&#34;},
    <a id="L56"></a>T{0x10a11, &#34;Kharoshthi&#34;},
    <a id="L57"></a>T{0x17c6, &#34;Khmer&#34;},
    <a id="L58"></a>T{0x0eaa, &#34;Lao&#34;},
    <a id="L59"></a>T{0x1d79, &#34;Latin&#34;},
    <a id="L60"></a>T{0x1c10, &#34;Lepcha&#34;},
    <a id="L61"></a>T{0x1930, &#34;Limbu&#34;},
    <a id="L62"></a>T{0x1003c, &#34;Linear_B&#34;},
    <a id="L63"></a>T{0x10290, &#34;Lycian&#34;},
    <a id="L64"></a>T{0x10930, &#34;Lydian&#34;},
    <a id="L65"></a>T{0x0d42, &#34;Malayalam&#34;},
    <a id="L66"></a>T{0x1822, &#34;Mongolian&#34;},
    <a id="L67"></a>T{0x104c, &#34;Myanmar&#34;},
    <a id="L68"></a>T{0x19c3, &#34;New_Tai_Lue&#34;},
    <a id="L69"></a>T{0x07f8, &#34;Nko&#34;},
    <a id="L70"></a>T{0x169b, &#34;Ogham&#34;},
    <a id="L71"></a>T{0x1c6a, &#34;Ol_Chiki&#34;},
    <a id="L72"></a>T{0x10310, &#34;Old_Italic&#34;},
    <a id="L73"></a>T{0x103c9, &#34;Old_Persian&#34;},
    <a id="L74"></a>T{0x0b3e, &#34;Oriya&#34;},
    <a id="L75"></a>T{0x10491, &#34;Osmanya&#34;},
    <a id="L76"></a>T{0xa860, &#34;Phags_Pa&#34;},
    <a id="L77"></a>T{0x10918, &#34;Phoenician&#34;},
    <a id="L78"></a>T{0xa949, &#34;Rejang&#34;},
    <a id="L79"></a>T{0x16c0, &#34;Runic&#34;},
    <a id="L80"></a>T{0xa892, &#34;Saurashtra&#34;},
    <a id="L81"></a>T{0x10463, &#34;Shavian&#34;},
    <a id="L82"></a>T{0x0dbd, &#34;Sinhala&#34;},
    <a id="L83"></a>T{0x1ba3, &#34;Sundanese&#34;},
    <a id="L84"></a>T{0xa803, &#34;Syloti_Nagri&#34;},
    <a id="L85"></a>T{0x070f, &#34;Syriac&#34;},
    <a id="L86"></a>T{0x170f, &#34;Tagalog&#34;},
    <a id="L87"></a>T{0x176f, &#34;Tagbanwa&#34;},
    <a id="L88"></a>T{0x1972, &#34;Tai_Le&#34;},
    <a id="L89"></a>T{0x0bbf, &#34;Tamil&#34;},
    <a id="L90"></a>T{0x0c55, &#34;Telugu&#34;},
    <a id="L91"></a>T{0x07a7, &#34;Thaana&#34;},
    <a id="L92"></a>T{0x0e46, &#34;Thai&#34;},
    <a id="L93"></a>T{0x0f36, &#34;Tibetan&#34;},
    <a id="L94"></a>T{0x2d55, &#34;Tifinagh&#34;},
    <a id="L95"></a>T{0x10388, &#34;Ugaritic&#34;},
    <a id="L96"></a>T{0xa60e, &#34;Vai&#34;},
    <a id="L97"></a>T{0xa216, &#34;Yi&#34;},
<a id="L98"></a>}

<a id="L100"></a>var outTest = []T{ <span class="comment">// not really worth being thorough</span>
    <a id="L101"></a>T{0x20, &#34;Telugu&#34;},
<a id="L102"></a>}

<a id="L104"></a>var inCategoryTest = []T{
    <a id="L105"></a>T{0x0081, &#34;Cc&#34;},
    <a id="L106"></a>T{0x17b4, &#34;Cf&#34;},
    <a id="L107"></a>T{0xf0000, &#34;Co&#34;},
    <a id="L108"></a>T{0xdb80, &#34;Cs&#34;},
    <a id="L109"></a>T{0x0236, &#34;Ll&#34;},
    <a id="L110"></a>T{0x1d9d, &#34;Lm&#34;},
    <a id="L111"></a>T{0x07cf, &#34;Lo&#34;},
    <a id="L112"></a>T{0x1f8a, &#34;Lt&#34;},
    <a id="L113"></a>T{0x03ff, &#34;Lu&#34;},
    <a id="L114"></a>T{0x0bc1, &#34;Mc&#34;},
    <a id="L115"></a>T{0x20df, &#34;Me&#34;},
    <a id="L116"></a>T{0x07f0, &#34;Mn&#34;},
    <a id="L117"></a>T{0x1bb2, &#34;Nd&#34;},
    <a id="L118"></a>T{0x10147, &#34;Nl&#34;},
    <a id="L119"></a>T{0x2478, &#34;No&#34;},
    <a id="L120"></a>T{0xfe33, &#34;Pc&#34;},
    <a id="L121"></a>T{0x2011, &#34;Pd&#34;},
    <a id="L122"></a>T{0x301e, &#34;Pe&#34;},
    <a id="L123"></a>T{0x2e03, &#34;Pf&#34;},
    <a id="L124"></a>T{0x2e02, &#34;Pi&#34;},
    <a id="L125"></a>T{0x0022, &#34;Po&#34;},
    <a id="L126"></a>T{0x2770, &#34;Ps&#34;},
    <a id="L127"></a>T{0x00a4, &#34;Sc&#34;},
    <a id="L128"></a>T{0xa711, &#34;Sk&#34;},
    <a id="L129"></a>T{0x25f9, &#34;Sm&#34;},
    <a id="L130"></a>T{0x2108, &#34;So&#34;},
    <a id="L131"></a>T{0x2028, &#34;Zl&#34;},
    <a id="L132"></a>T{0x2029, &#34;Zp&#34;},
    <a id="L133"></a>T{0x202f, &#34;Zs&#34;},
    <a id="L134"></a>T{0x04aa, &#34;letter&#34;},
<a id="L135"></a>}

<a id="L137"></a>var inPropTest = []T{
    <a id="L138"></a>T{0x0046, &#34;ASCII_Hex_Digit&#34;},
    <a id="L139"></a>T{0x200F, &#34;Bidi_Control&#34;},
    <a id="L140"></a>T{0x2212, &#34;Dash&#34;},
    <a id="L141"></a>T{0xE0001, &#34;Deprecated&#34;},
    <a id="L142"></a>T{0x00B7, &#34;Diacritic&#34;},
    <a id="L143"></a>T{0x30FE, &#34;Extender&#34;},
    <a id="L144"></a>T{0xFF46, &#34;Hex_Digit&#34;},
    <a id="L145"></a>T{0x2E17, &#34;Hyphen&#34;},
    <a id="L146"></a>T{0x2FFB, &#34;IDS_Binary_Operator&#34;},
    <a id="L147"></a>T{0x2FF3, &#34;IDS_Trinary_Operator&#34;},
    <a id="L148"></a>T{0xFA6A, &#34;Ideographic&#34;},
    <a id="L149"></a>T{0x200D, &#34;Join_Control&#34;},
    <a id="L150"></a>T{0x0EC4, &#34;Logical_Order_Exception&#34;},
    <a id="L151"></a>T{0x2FFFF, &#34;Noncharacter_Code_Point&#34;},
    <a id="L152"></a>T{0x065E, &#34;Other_Alphabetic&#34;},
    <a id="L153"></a>T{0x2069, &#34;Other_Default_Ignorable_Code_Point&#34;},
    <a id="L154"></a>T{0x0BD7, &#34;Other_Grapheme_Extend&#34;},
    <a id="L155"></a>T{0x0387, &#34;Other_ID_Continue&#34;},
    <a id="L156"></a>T{0x212E, &#34;Other_ID_Start&#34;},
    <a id="L157"></a>T{0x2094, &#34;Other_Lowercase&#34;},
    <a id="L158"></a>T{0x2040, &#34;Other_Math&#34;},
    <a id="L159"></a>T{0x216F, &#34;Other_Uppercase&#34;},
    <a id="L160"></a>T{0x0027, &#34;Pattern_Syntax&#34;},
    <a id="L161"></a>T{0x0020, &#34;Pattern_White_Space&#34;},
    <a id="L162"></a>T{0x300D, &#34;Quotation_Mark&#34;},
    <a id="L163"></a>T{0x2EF3, &#34;Radical&#34;},
    <a id="L164"></a>T{0x061F, &#34;STerm&#34;},
    <a id="L165"></a>T{0x2071, &#34;Soft_Dotted&#34;},
    <a id="L166"></a>T{0x003A, &#34;Terminal_Punctuation&#34;},
    <a id="L167"></a>T{0x9FC3, &#34;Unified_Ideograph&#34;},
    <a id="L168"></a>T{0xFE0F, &#34;Variation_Selector&#34;},
    <a id="L169"></a>T{0x0020, &#34;White_Space&#34;},
<a id="L170"></a>}

<a id="L172"></a>func TestScripts(t *testing.T) {
    <a id="L173"></a>notTested := make(map[string]bool);
    <a id="L174"></a>for k := range Scripts {
        <a id="L175"></a>notTested[k] = true
    <a id="L176"></a>}
    <a id="L177"></a>for _, test := range inTest {
        <a id="L178"></a>if _, ok := Scripts[test.script]; !ok {
            <a id="L179"></a>t.Fatal(test.script, &#34;not a known script&#34;)
        <a id="L180"></a>}
        <a id="L181"></a>if !Is(Scripts[test.script], test.rune) {
            <a id="L182"></a>t.Errorf(&#34;IsScript(%#x, %s) = false, want true\n&#34;, test.rune, test.script)
        <a id="L183"></a>}
        <a id="L184"></a>notTested[test.script] = false, false;
    <a id="L185"></a>}
    <a id="L186"></a>for _, test := range outTest {
        <a id="L187"></a>if Is(Scripts[test.script], test.rune) {
            <a id="L188"></a>t.Errorf(&#34;IsScript(%#x, %s) = true, want false\n&#34;, test.rune, test.script)
        <a id="L189"></a>}
    <a id="L190"></a>}
    <a id="L191"></a>for k := range notTested {
        <a id="L192"></a>t.Error(&#34;not tested:&#34;, k)
    <a id="L193"></a>}
<a id="L194"></a>}

<a id="L196"></a>func TestCategories(t *testing.T) {
    <a id="L197"></a>notTested := make(map[string]bool);
    <a id="L198"></a>for k := range Categories {
        <a id="L199"></a>notTested[k] = true
    <a id="L200"></a>}
    <a id="L201"></a>for _, test := range inCategoryTest {
        <a id="L202"></a>if _, ok := Categories[test.script]; !ok {
            <a id="L203"></a>t.Fatal(test.script, &#34;not a known category&#34;)
        <a id="L204"></a>}
        <a id="L205"></a>if !Is(Categories[test.script], test.rune) {
            <a id="L206"></a>t.Errorf(&#34;IsCategory(%#x, %s) = false, want true\n&#34;, test.rune, test.script)
        <a id="L207"></a>}
        <a id="L208"></a>notTested[test.script] = false, false;
    <a id="L209"></a>}
    <a id="L210"></a>for k := range notTested {
        <a id="L211"></a>t.Error(&#34;not tested:&#34;, k)
    <a id="L212"></a>}
<a id="L213"></a>}

<a id="L215"></a>func TestProperties(t *testing.T) {
    <a id="L216"></a>notTested := make(map[string]bool);
    <a id="L217"></a>for k := range Properties {
        <a id="L218"></a>notTested[k] = true
    <a id="L219"></a>}
    <a id="L220"></a>for _, test := range inPropTest {
        <a id="L221"></a>if _, ok := Properties[test.script]; !ok {
            <a id="L222"></a>t.Fatal(test.script, &#34;not a known prop&#34;)
        <a id="L223"></a>}
        <a id="L224"></a>if !Is(Properties[test.script], test.rune) {
            <a id="L225"></a>t.Errorf(&#34;IsCategory(%#x, %s) = false, want true\n&#34;, test.rune, test.script)
        <a id="L226"></a>}
        <a id="L227"></a>notTested[test.script] = false, false;
    <a id="L228"></a>}
    <a id="L229"></a>for k := range notTested {
        <a id="L230"></a>t.Error(&#34;not tested:&#34;, k)
    <a id="L231"></a>}
<a id="L232"></a>}
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
