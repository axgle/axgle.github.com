<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/unicode/letter.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/unicode/letter.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package provides data and functions to test some properties of Unicode code points.</span>
<a id="L6"></a>package unicode

<a id="L8"></a>const (
    <a id="L9"></a>MaxRune         = 0x10FFFF; <span class="comment">// Maximum valid Unicode code point.</span>
    <a id="L10"></a>ReplacementChar = 0xFFFD;   <span class="comment">// Represents invalid code points.</span>
<a id="L11"></a>)


<a id="L14"></a><span class="comment">// The representation of a range of Unicode code points.  The range runs from Lo to Hi</span>
<a id="L15"></a><span class="comment">// inclusive and has the specified stride.</span>
<a id="L16"></a>type Range struct {
    <a id="L17"></a>Lo     int;
    <a id="L18"></a>Hi     int;
    <a id="L19"></a>Stride int;
<a id="L20"></a>}

<a id="L22"></a><span class="comment">// The representation of a range of Unicode code points for case conversion.</span>
<a id="L23"></a><span class="comment">// The range runs from Lo to Hi inclusive, with a fixed stride of 1.  Deltas</span>
<a id="L24"></a><span class="comment">// are the number to add to the code point to reach the code point for a</span>
<a id="L25"></a><span class="comment">// different case for that character.  They may be negative.  If zero, it</span>
<a id="L26"></a><span class="comment">// means the character is in the corresponding case. There is a special</span>
<a id="L27"></a><span class="comment">// case representing sequences of alternating corresponding Upper and Lower</span>
<a id="L28"></a><span class="comment">// pairs.  It appears with a fixed Delta of</span>
<a id="L29"></a><span class="comment">//	{UpperLower, UpperLower, UpperLower}</span>
<a id="L30"></a><span class="comment">// The constant UpperLower has an otherwise impossible delta value.</span>
<a id="L31"></a>type CaseRange struct {
    <a id="L32"></a>Lo    int;
    <a id="L33"></a>Hi    int;
    <a id="L34"></a>Delta d;
<a id="L35"></a>}

<a id="L37"></a><span class="comment">// Indices into the Delta arrays inside CaseRanges for case mapping.</span>
<a id="L38"></a>const (
    <a id="L39"></a>UpperCase = iota;
    <a id="L40"></a>LowerCase;
    <a id="L41"></a>TitleCase;
    <a id="L42"></a>MaxCase;
<a id="L43"></a>)

<a id="L45"></a>type d [MaxCase]int32 <span class="comment">// to make the CaseRanges text shorter</span>

<a id="L47"></a><span class="comment">// If the Delta field of a CaseRange is UpperLower or LowerUpper, it means</span>
<a id="L48"></a><span class="comment">// this CaseRange represents a sequence of the form (say)</span>
<a id="L49"></a><span class="comment">// Upper Lower Upper Lower.</span>
<a id="L50"></a>const (
    <a id="L51"></a>UpperLower = MaxRune + 1; <span class="comment">// (Cannot be a valid delta.)</span>
<a id="L52"></a>)

<a id="L54"></a><span class="comment">// Is tests whether rune is in the specified table of ranges.</span>
<a id="L55"></a>func Is(ranges []Range, rune int) bool {
    <a id="L56"></a><span class="comment">// common case: rune is ASCII or Latin-1</span>
    <a id="L57"></a>if rune &lt; 0x100 {
        <a id="L58"></a>for _, r := range ranges {
            <a id="L59"></a>if rune &gt; r.Hi {
                <a id="L60"></a>continue
            <a id="L61"></a>}
            <a id="L62"></a>if rune &lt; r.Lo {
                <a id="L63"></a>return false
            <a id="L64"></a>}
            <a id="L65"></a>return (rune-r.Lo)%r.Stride == 0;
        <a id="L66"></a>}
        <a id="L67"></a>return false;
    <a id="L68"></a>}

    <a id="L70"></a><span class="comment">// binary search over ranges</span>
    <a id="L71"></a>lo := 0;
    <a id="L72"></a>hi := len(ranges);
    <a id="L73"></a>for lo &lt; hi {
        <a id="L74"></a>m := lo + (hi-lo)/2;
        <a id="L75"></a>r := ranges[m];
        <a id="L76"></a>if r.Lo &lt;= rune &amp;&amp; rune &lt;= r.Hi {
            <a id="L77"></a>return (rune-r.Lo)%r.Stride == 0
        <a id="L78"></a>}
        <a id="L79"></a>if rune &lt; r.Lo {
            <a id="L80"></a>hi = m
        <a id="L81"></a>} else {
            <a id="L82"></a>lo = m + 1
        <a id="L83"></a>}
    <a id="L84"></a>}
    <a id="L85"></a>return false;
<a id="L86"></a>}

<a id="L88"></a><span class="comment">// IsUpper reports whether the rune is an upper case letter.</span>
<a id="L89"></a>func IsUpper(rune int) bool {
    <a id="L90"></a>if rune &lt; 0x80 { <span class="comment">// quick ASCII check</span>
        <a id="L91"></a>return &#39;A&#39; &lt;= rune &amp;&amp; rune &lt;= &#39;Z&#39;
    <a id="L92"></a>}
    <a id="L93"></a>return Is(Upper, rune);
<a id="L94"></a>}

<a id="L96"></a><span class="comment">// IsLower reports whether the rune is a lower case letter.</span>
<a id="L97"></a>func IsLower(rune int) bool {
    <a id="L98"></a>if rune &lt; 0x80 { <span class="comment">// quick ASCII check</span>
        <a id="L99"></a>return &#39;a&#39; &lt;= rune &amp;&amp; rune &lt;= &#39;z&#39;
    <a id="L100"></a>}
    <a id="L101"></a>return Is(Lower, rune);
<a id="L102"></a>}

<a id="L104"></a><span class="comment">// IsTitle reports whether the rune is a title case letter.</span>
<a id="L105"></a>func IsTitle(rune int) bool {
    <a id="L106"></a>if rune &lt; 0x80 { <span class="comment">// quick ASCII check</span>
        <a id="L107"></a>return false
    <a id="L108"></a>}
    <a id="L109"></a>return Is(Title, rune);
<a id="L110"></a>}

<a id="L112"></a><span class="comment">// IsLetter reports whether the rune is a letter.</span>
<a id="L113"></a>func IsLetter(rune int) bool {
    <a id="L114"></a>if rune &lt; 0x80 { <span class="comment">// quick ASCII check</span>
        <a id="L115"></a>rune &amp;^= &#39;a&#39; - &#39;A&#39;;
        <a id="L116"></a>return &#39;A&#39; &lt;= rune &amp;&amp; rune &lt;= &#39;Z&#39;;
    <a id="L117"></a>}
    <a id="L118"></a>return Is(Letter, rune);
<a id="L119"></a>}

<a id="L121"></a><span class="comment">// IsSpace reports whether the rune is a white space character.</span>
<a id="L122"></a>func IsSpace(rune int) bool {
    <a id="L123"></a>if rune &lt;= 0xFF { <span class="comment">// quick Latin-1 check</span>
        <a id="L124"></a>switch rune {
        <a id="L125"></a>case &#39;\t&#39;, &#39;\n&#39;, &#39;\v&#39;, &#39;\f&#39;, &#39;\r&#39;, &#39; &#39;, 0x85, 0xA0:
            <a id="L126"></a>return true
        <a id="L127"></a>}
        <a id="L128"></a>return false;
    <a id="L129"></a>}
    <a id="L130"></a>return Is(White_Space, rune);
<a id="L131"></a>}

<a id="L133"></a><span class="comment">// To maps the rune to the specified case: UpperCase, LowerCase, or TitleCase</span>
<a id="L134"></a>func To(_case int, rune int) int {
    <a id="L135"></a>if _case &lt; 0 || MaxCase &lt;= _case {
        <a id="L136"></a>return ReplacementChar <span class="comment">// as reasonable an error as any</span>
    <a id="L137"></a>}
    <a id="L138"></a><span class="comment">// binary search over ranges</span>
    <a id="L139"></a>lo := 0;
    <a id="L140"></a>hi := len(CaseRanges);
    <a id="L141"></a>for lo &lt; hi {
        <a id="L142"></a>m := lo + (hi-lo)/2;
        <a id="L143"></a>r := CaseRanges[m];
        <a id="L144"></a>if r.Lo &lt;= rune &amp;&amp; rune &lt;= r.Hi {
            <a id="L145"></a>delta := int(r.Delta[_case]);
            <a id="L146"></a>if delta &gt; MaxRune {
                <a id="L147"></a><span class="comment">// In an Upper-Lower sequence, which always starts with</span>
                <a id="L148"></a><span class="comment">// an UpperCase letter, the real deltas always look like:</span>
                <a id="L149"></a><span class="comment">//	{0, 1, 0}    UpperCase (Lower is next)</span>
                <a id="L150"></a><span class="comment">//	{-1, 0, -1}  LowerCase (Upper, Title are previous)</span>
                <a id="L151"></a><span class="comment">// The characters at even offsets from the beginning of the</span>
                <a id="L152"></a><span class="comment">// sequence are upper case; the ones at odd offsets are lower.</span>
                <a id="L153"></a><span class="comment">// The correct mapping can be done by clearing or setting the low</span>
                <a id="L154"></a><span class="comment">// bit in the sequence offset.</span>
                <a id="L155"></a><span class="comment">// The constants UpperCase and TitleCase are even while LowerCase</span>
                <a id="L156"></a><span class="comment">// is odd so we take the low bit from _case.</span>
                <a id="L157"></a>return r.Lo + ((rune-r.Lo)&amp;^1 | _case&amp;1)
            <a id="L158"></a>}
            <a id="L159"></a>return rune + delta;
        <a id="L160"></a>}
        <a id="L161"></a>if rune &lt; r.Lo {
            <a id="L162"></a>hi = m
        <a id="L163"></a>} else {
            <a id="L164"></a>lo = m + 1
        <a id="L165"></a>}
    <a id="L166"></a>}
    <a id="L167"></a>return rune;
<a id="L168"></a>}

<a id="L170"></a><span class="comment">// ToUpper maps the rune to upper case</span>
<a id="L171"></a>func ToUpper(rune int) int {
    <a id="L172"></a>if rune &lt; 0x80 { <span class="comment">// quick ASCII check</span>
        <a id="L173"></a>if &#39;a&#39; &lt;= rune &amp;&amp; rune &lt;= &#39;z&#39; {
            <a id="L174"></a>rune -= &#39;a&#39; - &#39;A&#39;
        <a id="L175"></a>}
        <a id="L176"></a>return rune;
    <a id="L177"></a>}
    <a id="L178"></a>return To(UpperCase, rune);
<a id="L179"></a>}

<a id="L181"></a><span class="comment">// ToLower maps the rune to lower case</span>
<a id="L182"></a>func ToLower(rune int) int {
    <a id="L183"></a>if rune &lt; 0x80 { <span class="comment">// quick ASCII check</span>
        <a id="L184"></a>if &#39;A&#39; &lt;= rune &amp;&amp; rune &lt;= &#39;Z&#39; {
            <a id="L185"></a>rune += &#39;a&#39; - &#39;A&#39;
        <a id="L186"></a>}
        <a id="L187"></a>return rune;
    <a id="L188"></a>}
    <a id="L189"></a>return To(LowerCase, rune);
<a id="L190"></a>}

<a id="L192"></a><span class="comment">// ToTitle maps the rune to title case</span>
<a id="L193"></a>func ToTitle(rune int) int {
    <a id="L194"></a>if rune &lt; 0x80 { <span class="comment">// quick ASCII check</span>
        <a id="L195"></a>if &#39;a&#39; &lt;= rune &amp;&amp; rune &lt;= &#39;z&#39; { <span class="comment">// title case is upper case for ASCII</span>
            <a id="L196"></a>rune -= &#39;a&#39; - &#39;A&#39;
        <a id="L197"></a>}
        <a id="L198"></a>return rune;
    <a id="L199"></a>}
    <a id="L200"></a>return To(TitleCase, rune);
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
