<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/compress/flate/token.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/compress/flate/token.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package flate

<a id="L7"></a>const (
    <a id="L8"></a><span class="comment">// 2 bits:   type   0 = literal  1=EOF  2=Match   3=Unused</span>
    <a id="L9"></a><span class="comment">// 8 bits:   xlength = length - MIN_MATCH_LENGTH</span>
    <a id="L10"></a><span class="comment">// 22 bits   xoffset = offset - MIN_OFFSET_SIZE, or literal</span>
    <a id="L11"></a>lengthShift = 22;
    <a id="L12"></a>offsetMask  = 1&lt;&lt;lengthShift - 1;
    <a id="L13"></a>typeMask    = 3 &lt;&lt; 30;
    <a id="L14"></a>literalType = 0 &lt;&lt; 30;
    <a id="L15"></a>matchType   = 1 &lt;&lt; 30;
<a id="L16"></a>)

<a id="L18"></a><span class="comment">// The length code for length X (MIN_MATCH_LENGTH &lt;= X &lt;= MAX_MATCH_LENGTH)</span>
<a id="L19"></a><span class="comment">// is lengthCodes[length - MIN_MATCH_LENGTH]</span>
<a id="L20"></a>var lengthCodes = [...]uint32{
    <a id="L21"></a>0, 1, 2, 3, 4, 5, 6, 7, 8, 8,
    <a id="L22"></a>9, 9, 10, 10, 11, 11, 12, 12, 12, 12,
    <a id="L23"></a>13, 13, 13, 13, 14, 14, 14, 14, 15, 15,
    <a id="L24"></a>15, 15, 16, 16, 16, 16, 16, 16, 16, 16,
    <a id="L25"></a>17, 17, 17, 17, 17, 17, 17, 17, 18, 18,
    <a id="L26"></a>18, 18, 18, 18, 18, 18, 19, 19, 19, 19,
    <a id="L27"></a>19, 19, 19, 19, 20, 20, 20, 20, 20, 20,
    <a id="L28"></a>20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
    <a id="L29"></a>21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
    <a id="L30"></a>21, 21, 21, 21, 21, 21, 22, 22, 22, 22,
    <a id="L31"></a>22, 22, 22, 22, 22, 22, 22, 22, 22, 22,
    <a id="L32"></a>22, 22, 23, 23, 23, 23, 23, 23, 23, 23,
    <a id="L33"></a>23, 23, 23, 23, 23, 23, 23, 23, 24, 24,
    <a id="L34"></a>24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
    <a id="L35"></a>24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
    <a id="L36"></a>24, 24, 24, 24, 24, 24, 24, 24, 24, 24,
    <a id="L37"></a>25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
    <a id="L38"></a>25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
    <a id="L39"></a>25, 25, 25, 25, 25, 25, 25, 25, 25, 25,
    <a id="L40"></a>25, 25, 26, 26, 26, 26, 26, 26, 26, 26,
    <a id="L41"></a>26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
    <a id="L42"></a>26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
    <a id="L43"></a>26, 26, 26, 26, 27, 27, 27, 27, 27, 27,
    <a id="L44"></a>27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
    <a id="L45"></a>27, 27, 27, 27, 27, 27, 27, 27, 27, 27,
    <a id="L46"></a>27, 27, 27, 27, 27, 28,
<a id="L47"></a>}

<a id="L49"></a>var offsetCodes = [...]uint32{
    <a id="L50"></a>0, 1, 2, 3, 4, 4, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7,
    <a id="L51"></a>8, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 9,
    <a id="L52"></a>10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
    <a id="L53"></a>11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
    <a id="L54"></a>12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
    <a id="L55"></a>12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
    <a id="L56"></a>13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
    <a id="L57"></a>13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
    <a id="L58"></a>14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
    <a id="L59"></a>14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
    <a id="L60"></a>14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
    <a id="L61"></a>14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
    <a id="L62"></a>15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
    <a id="L63"></a>15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
    <a id="L64"></a>15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
    <a id="L65"></a>15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
<a id="L66"></a>}

<a id="L68"></a>type token uint32

<a id="L70"></a><span class="comment">// Convert a literal into a literal token.</span>
<a id="L71"></a>func literalToken(literal uint32) token { return token(literalType + literal) }

<a id="L73"></a><span class="comment">// Convert a &lt; xlength, xoffset &gt; pair into a match token.</span>
<a id="L74"></a>func matchToken(xlength uint32, xoffset uint32) token {
    <a id="L75"></a>return token(matchType + xlength&lt;&lt;lengthShift + xoffset)
<a id="L76"></a>}

<a id="L78"></a><span class="comment">// Returns the type of a token</span>
<a id="L79"></a>func (t token) typ() uint32 { return uint32(t) &amp; typeMask }

<a id="L81"></a><span class="comment">// Returns the literal of a literal token</span>
<a id="L82"></a>func (t token) literal() uint32 { return uint32(t - literalType) }

<a id="L84"></a><span class="comment">// Returns the extra offset of a match token</span>
<a id="L85"></a>func (t token) offset() uint32 { return uint32(t) &amp; offsetMask }

<a id="L87"></a>func (t token) length() uint32 { return uint32((t - matchType) &gt;&gt; lengthShift) }

<a id="L89"></a>func lengthCode(len uint32) uint32 { return lengthCodes[len] }

<a id="L91"></a><span class="comment">// Returns the offset code corresponding to a specific offset</span>
<a id="L92"></a>func offsetCode(off uint32) uint32 {
    <a id="L93"></a>const n = uint32(len(offsetCodes));
    <a id="L94"></a>switch {
    <a id="L95"></a>case off &lt; n:
        <a id="L96"></a>return offsetCodes[off]
    <a id="L97"></a>case off&gt;&gt;7 &lt; n:
        <a id="L98"></a>return offsetCodes[off&gt;&gt;7] + 14
    <a id="L99"></a>default:
        <a id="L100"></a>return offsetCodes[off&gt;&gt;14] + 28
    <a id="L101"></a>}
    <a id="L102"></a>panic(&#34;unreachable&#34;);
<a id="L103"></a>}
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
