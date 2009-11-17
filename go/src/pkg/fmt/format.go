<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/fmt/format.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/fmt/format.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package fmt

<a id="L7"></a>import (
    <a id="L8"></a>&#34;strconv&#34;;
<a id="L9"></a>)


<a id="L12"></a>const nByte = 64
<a id="L13"></a>const nPows10 = 160

<a id="L15"></a>var ldigits string = &#34;0123456789abcdef&#34; <span class="comment">// var not const because we take its address</span>
<a id="L16"></a>var udigits string = &#34;0123456789ABCDEF&#34;

<a id="L18"></a><span class="comment">/*</span>
<a id="L19"></a><span class="comment">	Fmt is the raw formatter used by Printf etc.  Not meant for normal use.</span>
<a id="L20"></a><span class="comment">	See print.go for a more palatable interface.</span>

<a id="L22"></a><span class="comment">	The model is to accumulate operands into an internal buffer and then</span>
<a id="L23"></a><span class="comment">	retrieve the buffer in one hit using Str(), Putnl(), etc.  The formatting</span>
<a id="L24"></a><span class="comment">	methods return ``self&#39;&#39; so the operations can be chained.</span>

<a id="L26"></a><span class="comment">		f := fmt.New();</span>
<a id="L27"></a><span class="comment">		print(f.Fmt_d(1234).Fmt_s(&#34;\n&#34;).Str());  // create string, print it</span>
<a id="L28"></a><span class="comment">		f.Fmt_d(-1234).Fmt_s(&#34;\n&#34;).Put();  // print string</span>
<a id="L29"></a><span class="comment">		f.Fmt_ud(1&lt;&lt;63).Putnl();  // print string with automatic newline</span>
<a id="L30"></a><span class="comment">*/</span>
<a id="L31"></a>type Fmt struct {
    <a id="L32"></a>buf          string;
    <a id="L33"></a>wid          int;
    <a id="L34"></a>wid_present  bool;
    <a id="L35"></a>prec         int;
    <a id="L36"></a>prec_present bool;
    <a id="L37"></a><span class="comment">// flags</span>
    <a id="L38"></a>minus bool;
    <a id="L39"></a>plus  bool;
    <a id="L40"></a>sharp bool;
    <a id="L41"></a>space bool;
    <a id="L42"></a>zero  bool;
<a id="L43"></a>}

<a id="L45"></a>func (f *Fmt) clearflags() {
    <a id="L46"></a>f.wid = 0;
    <a id="L47"></a>f.wid_present = false;
    <a id="L48"></a>f.prec = 0;
    <a id="L49"></a>f.prec_present = false;
    <a id="L50"></a>f.minus = false;
    <a id="L51"></a>f.plus = false;
    <a id="L52"></a>f.sharp = false;
    <a id="L53"></a>f.space = false;
    <a id="L54"></a>f.zero = false;
<a id="L55"></a>}

<a id="L57"></a>func (f *Fmt) clearbuf() { f.buf = &#34;&#34; }

<a id="L59"></a>func (f *Fmt) init() {
    <a id="L60"></a>f.clearbuf();
    <a id="L61"></a>f.clearflags();
<a id="L62"></a>}

<a id="L64"></a><span class="comment">// New returns a new initialized Fmt</span>
<a id="L65"></a>func New() *Fmt {
    <a id="L66"></a>f := new(Fmt);
    <a id="L67"></a>f.init();
    <a id="L68"></a>return f;
<a id="L69"></a>}

<a id="L71"></a><span class="comment">// Str returns the buffered contents as a string and resets the Fmt.</span>
<a id="L72"></a>func (f *Fmt) Str() string {
    <a id="L73"></a>s := f.buf;
    <a id="L74"></a>f.clearbuf();
    <a id="L75"></a>f.clearflags();
    <a id="L76"></a>f.buf = &#34;&#34;;
    <a id="L77"></a>return s;
<a id="L78"></a>}

<a id="L80"></a><span class="comment">// Put writes the buffered contents to stdout and resets the Fmt.</span>
<a id="L81"></a>func (f *Fmt) Put() {
    <a id="L82"></a>print(f.buf);
    <a id="L83"></a>f.clearbuf();
    <a id="L84"></a>f.clearflags();
<a id="L85"></a>}

<a id="L87"></a><span class="comment">// Putnl writes the buffered contents to stdout, followed by a newline, and resets the Fmt.</span>
<a id="L88"></a>func (f *Fmt) Putnl() {
    <a id="L89"></a>print(f.buf, &#34;\n&#34;);
    <a id="L90"></a>f.clearbuf();
    <a id="L91"></a>f.clearflags();
<a id="L92"></a>}

<a id="L94"></a><span class="comment">// Wp sets the width and precision for formatting the next item.</span>
<a id="L95"></a>func (f *Fmt) Wp(w, p int) *Fmt {
    <a id="L96"></a>f.wid_present = true;
    <a id="L97"></a>f.wid = w;
    <a id="L98"></a>f.prec_present = true;
    <a id="L99"></a>f.prec = p;
    <a id="L100"></a>return f;
<a id="L101"></a>}

<a id="L103"></a><span class="comment">// P sets the precision for formatting the next item.</span>
<a id="L104"></a>func (f *Fmt) P(p int) *Fmt {
    <a id="L105"></a>f.prec_present = true;
    <a id="L106"></a>f.prec = p;
    <a id="L107"></a>return f;
<a id="L108"></a>}

<a id="L110"></a><span class="comment">// W sets the width for formatting the next item.</span>
<a id="L111"></a>func (f *Fmt) W(x int) *Fmt {
    <a id="L112"></a>f.wid_present = true;
    <a id="L113"></a>f.wid = x;
    <a id="L114"></a>return f;
<a id="L115"></a>}

<a id="L117"></a><span class="comment">// append s to buf, padded on left (w &gt; 0) or right (w &lt; 0 or f.minus)</span>
<a id="L118"></a><span class="comment">// padding is in bytes, not characters (agrees with ANSIC C, not Plan 9 C)</span>
<a id="L119"></a>func (f *Fmt) pad(s string) {
    <a id="L120"></a>if f.wid_present &amp;&amp; f.wid != 0 {
        <a id="L121"></a>left := !f.minus;
        <a id="L122"></a>w := f.wid;
        <a id="L123"></a>if w &lt; 0 {
            <a id="L124"></a>left = false;
            <a id="L125"></a>w = -w;
        <a id="L126"></a>}
        <a id="L127"></a>w -= len(s);
        <a id="L128"></a>padchar := byte(&#39; &#39;);
        <a id="L129"></a>if left &amp;&amp; f.zero {
            <a id="L130"></a>padchar = &#39;0&#39;
        <a id="L131"></a>}
        <a id="L132"></a>if w &gt; 0 {
            <a id="L133"></a>if w &gt; nByte {
                <a id="L134"></a>w = nByte
            <a id="L135"></a>}
            <a id="L136"></a>buf := make([]byte, w);
            <a id="L137"></a>for i := 0; i &lt; w; i++ {
                <a id="L138"></a>buf[i] = padchar
            <a id="L139"></a>}
            <a id="L140"></a>if left {
                <a id="L141"></a>s = string(buf) + s
            <a id="L142"></a>} else {
                <a id="L143"></a>s = s + string(buf)
            <a id="L144"></a>}
        <a id="L145"></a>}
    <a id="L146"></a>}
    <a id="L147"></a>f.buf += s;
<a id="L148"></a>}

<a id="L150"></a><span class="comment">// format val into buf, ending at buf[i].  (printing is easier right-to-left;</span>
<a id="L151"></a><span class="comment">// that&#39;s why the bidi languages are right-to-left except for numbers. wait,</span>
<a id="L152"></a><span class="comment">// never mind.)  val is known to be unsigned.  we could make things maybe</span>
<a id="L153"></a><span class="comment">// marginally faster by splitting the 32-bit case out into a separate function</span>
<a id="L154"></a><span class="comment">// but it&#39;s not worth the duplication, so val has 64 bits.</span>
<a id="L155"></a>func putint(buf []byte, base, val uint64, digits string) int {
    <a id="L156"></a>i := len(buf) - 1;
    <a id="L157"></a>for val &gt;= base {
        <a id="L158"></a>buf[i] = digits[val%base];
        <a id="L159"></a>i--;
        <a id="L160"></a>val /= base;
    <a id="L161"></a>}
    <a id="L162"></a>buf[i] = digits[val];
    <a id="L163"></a>return i - 1;
<a id="L164"></a>}

<a id="L166"></a><span class="comment">// Fmt_boolean formats a boolean.</span>
<a id="L167"></a>func (f *Fmt) Fmt_boolean(v bool) *Fmt {
    <a id="L168"></a>if v {
        <a id="L169"></a>f.pad(&#34;true&#34;)
    <a id="L170"></a>} else {
        <a id="L171"></a>f.pad(&#34;false&#34;)
    <a id="L172"></a>}
    <a id="L173"></a>f.clearflags();
    <a id="L174"></a>return f;
<a id="L175"></a>}

<a id="L177"></a><span class="comment">// integer; interprets prec but not wid.</span>
<a id="L178"></a>func (f *Fmt) integer(a int64, base uint, is_signed bool, digits string) string {
    <a id="L179"></a>var buf [nByte]byte;
    <a id="L180"></a>negative := is_signed &amp;&amp; a &lt; 0;
    <a id="L181"></a>if negative {
        <a id="L182"></a>a = -a
    <a id="L183"></a>}

    <a id="L185"></a><span class="comment">// two ways to ask for extra leading zero digits: %.3d or %03d.</span>
    <a id="L186"></a><span class="comment">// apparently the first cancels the second.</span>
    <a id="L187"></a>prec := 0;
    <a id="L188"></a>if f.prec_present {
        <a id="L189"></a>prec = f.prec;
        <a id="L190"></a>f.zero = false;
    <a id="L191"></a>} else if f.zero &amp;&amp; f.wid_present &amp;&amp; !f.minus &amp;&amp; f.wid &gt; 0 {
        <a id="L192"></a>prec = f.wid;
        <a id="L193"></a>if negative || f.plus || f.space {
            <a id="L194"></a>prec-- <span class="comment">// leave room for sign</span>
        <a id="L195"></a>}
    <a id="L196"></a>}

    <a id="L198"></a>i := putint(&amp;buf, uint64(base), uint64(a), digits);
    <a id="L199"></a>for i &gt; 0 &amp;&amp; prec &gt; (nByte-1-i) {
        <a id="L200"></a>buf[i] = &#39;0&#39;;
        <a id="L201"></a>i--;
    <a id="L202"></a>}

    <a id="L204"></a>if f.sharp {
        <a id="L205"></a>switch base {
        <a id="L206"></a>case 8:
            <a id="L207"></a>if buf[i+1] != &#39;0&#39; {
                <a id="L208"></a>buf[i] = &#39;0&#39;;
                <a id="L209"></a>i--;
            <a id="L210"></a>}
        <a id="L211"></a>case 16:
            <a id="L212"></a>buf[i] = &#39;x&#39; + digits[10] - &#39;a&#39;;
            <a id="L213"></a>i--;
            <a id="L214"></a>buf[i] = &#39;0&#39;;
            <a id="L215"></a>i--;
        <a id="L216"></a>}
    <a id="L217"></a>}

    <a id="L219"></a>if negative {
        <a id="L220"></a>buf[i] = &#39;-&#39;;
        <a id="L221"></a>i--;
    <a id="L222"></a>} else if f.plus {
        <a id="L223"></a>buf[i] = &#39;+&#39;;
        <a id="L224"></a>i--;
    <a id="L225"></a>} else if f.space {
        <a id="L226"></a>buf[i] = &#39; &#39;;
        <a id="L227"></a>i--;
    <a id="L228"></a>}
    <a id="L229"></a>return string(buf[i+1 : nByte]);
<a id="L230"></a>}

<a id="L232"></a><span class="comment">// Fmt_d64 formats an int64 in decimal.</span>
<a id="L233"></a>func (f *Fmt) Fmt_d64(v int64) *Fmt {
    <a id="L234"></a>f.pad(f.integer(v, 10, true, ldigits));
    <a id="L235"></a>f.clearflags();
    <a id="L236"></a>return f;
<a id="L237"></a>}

<a id="L239"></a><span class="comment">// Fmt_d32 formats an int32 in decimal.</span>
<a id="L240"></a>func (f *Fmt) Fmt_d32(v int32) *Fmt { return f.Fmt_d64(int64(v)) }

<a id="L242"></a><span class="comment">// Fmt_d formats an int in decimal.</span>
<a id="L243"></a>func (f *Fmt) Fmt_d(v int) *Fmt { return f.Fmt_d64(int64(v)) }

<a id="L245"></a><span class="comment">// Fmt_ud64 formats a uint64 in decimal.</span>
<a id="L246"></a>func (f *Fmt) Fmt_ud64(v uint64) *Fmt {
    <a id="L247"></a>f.pad(f.integer(int64(v), 10, false, ldigits));
    <a id="L248"></a>f.clearflags();
    <a id="L249"></a>return f;
<a id="L250"></a>}

<a id="L252"></a><span class="comment">// Fmt_ud32 formats a uint32 in decimal.</span>
<a id="L253"></a>func (f *Fmt) Fmt_ud32(v uint32) *Fmt { return f.Fmt_ud64(uint64(v)) }

<a id="L255"></a><span class="comment">// Fmt_ud formats a uint in decimal.</span>
<a id="L256"></a>func (f *Fmt) Fmt_ud(v uint) *Fmt { return f.Fmt_ud64(uint64(v)) }

<a id="L258"></a><span class="comment">// Fmt_x64 formats an int64 in hexadecimal.</span>
<a id="L259"></a>func (f *Fmt) Fmt_x64(v int64) *Fmt {
    <a id="L260"></a>f.pad(f.integer(v, 16, true, ldigits));
    <a id="L261"></a>f.clearflags();
    <a id="L262"></a>return f;
<a id="L263"></a>}

<a id="L265"></a><span class="comment">// Fmt_x32 formats an int32 in hexadecimal.</span>
<a id="L266"></a>func (f *Fmt) Fmt_x32(v int32) *Fmt { return f.Fmt_x64(int64(v)) }

<a id="L268"></a><span class="comment">// Fmt_x formats an int in hexadecimal.</span>
<a id="L269"></a>func (f *Fmt) Fmt_x(v int) *Fmt { return f.Fmt_x64(int64(v)) }

<a id="L271"></a><span class="comment">// Fmt_ux64 formats a uint64 in hexadecimal.</span>
<a id="L272"></a>func (f *Fmt) Fmt_ux64(v uint64) *Fmt {
    <a id="L273"></a>f.pad(f.integer(int64(v), 16, false, ldigits));
    <a id="L274"></a>f.clearflags();
    <a id="L275"></a>return f;
<a id="L276"></a>}

<a id="L278"></a><span class="comment">// Fmt_ux32 formats a uint32 in hexadecimal.</span>
<a id="L279"></a>func (f *Fmt) Fmt_ux32(v uint32) *Fmt { return f.Fmt_ux64(uint64(v)) }

<a id="L281"></a><span class="comment">// Fmt_ux formats a uint in hexadecimal.</span>
<a id="L282"></a>func (f *Fmt) Fmt_ux(v uint) *Fmt { return f.Fmt_ux64(uint64(v)) }

<a id="L284"></a><span class="comment">// Fmt_X64 formats an int64 in upper case hexadecimal.</span>
<a id="L285"></a>func (f *Fmt) Fmt_X64(v int64) *Fmt {
    <a id="L286"></a>f.pad(f.integer(v, 16, true, udigits));
    <a id="L287"></a>f.clearflags();
    <a id="L288"></a>return f;
<a id="L289"></a>}

<a id="L291"></a><span class="comment">// Fmt_X32 formats an int32 in upper case hexadecimal.</span>
<a id="L292"></a>func (f *Fmt) Fmt_X32(v int32) *Fmt { return f.Fmt_X64(int64(v)) }

<a id="L294"></a><span class="comment">// Fmt_X formats an int in upper case hexadecimal.</span>
<a id="L295"></a>func (f *Fmt) Fmt_X(v int) *Fmt { return f.Fmt_X64(int64(v)) }

<a id="L297"></a><span class="comment">// Fmt_uX64 formats a uint64 in upper case hexadecimal.</span>
<a id="L298"></a>func (f *Fmt) Fmt_uX64(v uint64) *Fmt {
    <a id="L299"></a>f.pad(f.integer(int64(v), 16, false, udigits));
    <a id="L300"></a>f.clearflags();
    <a id="L301"></a>return f;
<a id="L302"></a>}

<a id="L304"></a><span class="comment">// Fmt_uX32 formats a uint32 in upper case hexadecimal.</span>
<a id="L305"></a>func (f *Fmt) Fmt_uX32(v uint32) *Fmt { return f.Fmt_uX64(uint64(v)) }

<a id="L307"></a><span class="comment">// Fmt_uX formats a uint in upper case hexadecimal.</span>
<a id="L308"></a>func (f *Fmt) Fmt_uX(v uint) *Fmt { return f.Fmt_uX64(uint64(v)) }

<a id="L310"></a><span class="comment">// Fmt_o64 formats an int64 in octal.</span>
<a id="L311"></a>func (f *Fmt) Fmt_o64(v int64) *Fmt {
    <a id="L312"></a>f.pad(f.integer(v, 8, true, ldigits));
    <a id="L313"></a>f.clearflags();
    <a id="L314"></a>return f;
<a id="L315"></a>}

<a id="L317"></a><span class="comment">// Fmt_o32 formats an int32 in octal.</span>
<a id="L318"></a>func (f *Fmt) Fmt_o32(v int32) *Fmt { return f.Fmt_o64(int64(v)) }

<a id="L320"></a><span class="comment">// Fmt_o formats an int in octal.</span>
<a id="L321"></a>func (f *Fmt) Fmt_o(v int) *Fmt { return f.Fmt_o64(int64(v)) }

<a id="L323"></a><span class="comment">// Fmt_uo64 formats a uint64 in octal.</span>
<a id="L324"></a>func (f *Fmt) Fmt_uo64(v uint64) *Fmt {
    <a id="L325"></a>f.pad(f.integer(int64(v), 8, false, ldigits));
    <a id="L326"></a>f.clearflags();
    <a id="L327"></a>return f;
<a id="L328"></a>}

<a id="L330"></a><span class="comment">// Fmt_uo32 formats a uint32 in octal.</span>
<a id="L331"></a>func (f *Fmt) Fmt_uo32(v uint32) *Fmt { return f.Fmt_uo64(uint64(v)) }

<a id="L333"></a><span class="comment">// Fmt_uo formats a uint in octal.</span>
<a id="L334"></a>func (f *Fmt) Fmt_uo(v uint) *Fmt { return f.Fmt_uo64(uint64(v)) }

<a id="L336"></a><span class="comment">// Fmt_b64 formats a uint64 in binary.</span>
<a id="L337"></a>func (f *Fmt) Fmt_b64(v uint64) *Fmt {
    <a id="L338"></a>f.pad(f.integer(int64(v), 2, false, ldigits));
    <a id="L339"></a>f.clearflags();
    <a id="L340"></a>return f;
<a id="L341"></a>}

<a id="L343"></a><span class="comment">// Fmt_b32 formats a uint32 in binary.</span>
<a id="L344"></a>func (f *Fmt) Fmt_b32(v uint32) *Fmt { return f.Fmt_b64(uint64(v)) }

<a id="L346"></a><span class="comment">// Fmt_b formats a uint in binary.</span>
<a id="L347"></a>func (f *Fmt) Fmt_b(v uint) *Fmt { return f.Fmt_b64(uint64(v)) }

<a id="L349"></a><span class="comment">// Fmt_c formats a Unicode character.</span>
<a id="L350"></a>func (f *Fmt) Fmt_c(v int) *Fmt {
    <a id="L351"></a>f.pad(string(v));
    <a id="L352"></a>f.clearflags();
    <a id="L353"></a>return f;
<a id="L354"></a>}

<a id="L356"></a><span class="comment">// Fmt_s formats a string.</span>
<a id="L357"></a>func (f *Fmt) Fmt_s(s string) *Fmt {
    <a id="L358"></a>if f.prec_present {
        <a id="L359"></a>if f.prec &lt; len(s) {
            <a id="L360"></a>s = s[0:f.prec]
        <a id="L361"></a>}
    <a id="L362"></a>}
    <a id="L363"></a>f.pad(s);
    <a id="L364"></a>f.clearflags();
    <a id="L365"></a>return f;
<a id="L366"></a>}

<a id="L368"></a><span class="comment">// Fmt_sx formats a string as a hexadecimal encoding of its bytes.</span>
<a id="L369"></a>func (f *Fmt) Fmt_sx(s string) *Fmt {
    <a id="L370"></a>t := &#34;&#34;;
    <a id="L371"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L372"></a>if i &gt; 0 &amp;&amp; f.space {
            <a id="L373"></a>t += &#34; &#34;
        <a id="L374"></a>}
        <a id="L375"></a>v := s[i];
        <a id="L376"></a>t += string(ldigits[v&gt;&gt;4]);
        <a id="L377"></a>t += string(ldigits[v&amp;0xF]);
    <a id="L378"></a>}
    <a id="L379"></a>f.pad(t);
    <a id="L380"></a>f.clearflags();
    <a id="L381"></a>return f;
<a id="L382"></a>}

<a id="L384"></a><span class="comment">// Fmt_sX formats a string as an uppercase hexadecimal encoding of its bytes.</span>
<a id="L385"></a>func (f *Fmt) Fmt_sX(s string) *Fmt {
    <a id="L386"></a>t := &#34;&#34;;
    <a id="L387"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L388"></a>v := s[i];
        <a id="L389"></a>t += string(udigits[v&gt;&gt;4]);
        <a id="L390"></a>t += string(udigits[v&amp;0xF]);
    <a id="L391"></a>}
    <a id="L392"></a>f.pad(t);
    <a id="L393"></a>f.clearflags();
    <a id="L394"></a>return f;
<a id="L395"></a>}

<a id="L397"></a><span class="comment">// Fmt_q formats a string as a double-quoted, escaped Go string constant.</span>
<a id="L398"></a>func (f *Fmt) Fmt_q(s string) *Fmt {
    <a id="L399"></a>var quoted string;
    <a id="L400"></a>if f.sharp &amp;&amp; strconv.CanBackquote(s) {
        <a id="L401"></a>quoted = &#34;`&#34; + s + &#34;`&#34;
    <a id="L402"></a>} else {
        <a id="L403"></a>quoted = strconv.Quote(s)
    <a id="L404"></a>}
    <a id="L405"></a>f.pad(quoted);
    <a id="L406"></a>f.clearflags();
    <a id="L407"></a>return f;
<a id="L408"></a>}

<a id="L410"></a><span class="comment">// floating-point</span>

<a id="L412"></a>func doPrec(f *Fmt, def int) int {
    <a id="L413"></a>if f.prec_present {
        <a id="L414"></a>return f.prec
    <a id="L415"></a>}
    <a id="L416"></a>return def;
<a id="L417"></a>}

<a id="L419"></a>func fmtString(f *Fmt, s string) *Fmt {
    <a id="L420"></a>f.pad(s);
    <a id="L421"></a>f.clearflags();
    <a id="L422"></a>return f;
<a id="L423"></a>}

<a id="L425"></a><span class="comment">// Fmt_e64 formats a float64 in the form -1.23e+12.</span>
<a id="L426"></a>func (f *Fmt) Fmt_e64(v float64) *Fmt {
    <a id="L427"></a>return fmtString(f, strconv.Ftoa64(v, &#39;e&#39;, doPrec(f, 6)))
<a id="L428"></a>}

<a id="L430"></a><span class="comment">// Fmt_E64 formats a float64 in the form -1.23E+12.</span>
<a id="L431"></a>func (f *Fmt) Fmt_E64(v float64) *Fmt {
    <a id="L432"></a>return fmtString(f, strconv.Ftoa64(v, &#39;E&#39;, doPrec(f, 6)))
<a id="L433"></a>}

<a id="L435"></a><span class="comment">// Fmt_f64 formats a float64 in the form -1.23.</span>
<a id="L436"></a>func (f *Fmt) Fmt_f64(v float64) *Fmt {
    <a id="L437"></a>return fmtString(f, strconv.Ftoa64(v, &#39;f&#39;, doPrec(f, 6)))
<a id="L438"></a>}

<a id="L440"></a><span class="comment">// Fmt_g64 formats a float64 in the &#39;f&#39; or &#39;e&#39; form according to size.</span>
<a id="L441"></a>func (f *Fmt) Fmt_g64(v float64) *Fmt {
    <a id="L442"></a>return fmtString(f, strconv.Ftoa64(v, &#39;g&#39;, doPrec(f, -1)))
<a id="L443"></a>}

<a id="L445"></a><span class="comment">// Fmt_g64 formats a float64 in the &#39;f&#39; or &#39;E&#39; form according to size.</span>
<a id="L446"></a>func (f *Fmt) Fmt_G64(v float64) *Fmt {
    <a id="L447"></a>return fmtString(f, strconv.Ftoa64(v, &#39;G&#39;, doPrec(f, -1)))
<a id="L448"></a>}

<a id="L450"></a><span class="comment">// Fmt_fb64 formats a float64 in the form -123p3 (exponent is power of 2).</span>
<a id="L451"></a>func (f *Fmt) Fmt_fb64(v float64) *Fmt { return fmtString(f, strconv.Ftoa64(v, &#39;b&#39;, 0)) }

<a id="L453"></a><span class="comment">// float32</span>
<a id="L454"></a><span class="comment">// cannot defer to float64 versions</span>
<a id="L455"></a><span class="comment">// because it will get rounding wrong in corner cases.</span>

<a id="L457"></a><span class="comment">// Fmt_e32 formats a float32 in the form -1.23e+12.</span>
<a id="L458"></a>func (f *Fmt) Fmt_e32(v float32) *Fmt {
    <a id="L459"></a>return fmtString(f, strconv.Ftoa32(v, &#39;e&#39;, doPrec(f, 6)))
<a id="L460"></a>}

<a id="L462"></a><span class="comment">// Fmt_E32 formats a float32 in the form -1.23E+12.</span>
<a id="L463"></a>func (f *Fmt) Fmt_E32(v float32) *Fmt {
    <a id="L464"></a>return fmtString(f, strconv.Ftoa32(v, &#39;e&#39;, doPrec(f, 6)))
<a id="L465"></a>}

<a id="L467"></a><span class="comment">// Fmt_f32 formats a float32 in the form -1.23.</span>
<a id="L468"></a>func (f *Fmt) Fmt_f32(v float32) *Fmt {
    <a id="L469"></a>return fmtString(f, strconv.Ftoa32(v, &#39;f&#39;, doPrec(f, 6)))
<a id="L470"></a>}

<a id="L472"></a><span class="comment">// Fmt_g32 formats a float32 in the &#39;f&#39; or &#39;e&#39; form according to size.</span>
<a id="L473"></a>func (f *Fmt) Fmt_g32(v float32) *Fmt {
    <a id="L474"></a>return fmtString(f, strconv.Ftoa32(v, &#39;g&#39;, doPrec(f, -1)))
<a id="L475"></a>}

<a id="L477"></a><span class="comment">// Fmt_G32 formats a float32 in the &#39;f&#39; or &#39;E&#39; form according to size.</span>
<a id="L478"></a>func (f *Fmt) Fmt_G32(v float32) *Fmt {
    <a id="L479"></a>return fmtString(f, strconv.Ftoa32(v, &#39;G&#39;, doPrec(f, -1)))
<a id="L480"></a>}

<a id="L482"></a><span class="comment">// Fmt_fb32 formats a float32 in the form -123p3 (exponent is power of 2).</span>
<a id="L483"></a>func (f *Fmt) Fmt_fb32(v float32) *Fmt { return fmtString(f, strconv.Ftoa32(v, &#39;b&#39;, 0)) }

<a id="L485"></a><span class="comment">// float</span>
<a id="L486"></a>func (x *Fmt) f(a float) *Fmt {
    <a id="L487"></a>if strconv.FloatSize == 32 {
        <a id="L488"></a>return x.Fmt_f32(float32(a))
    <a id="L489"></a>}
    <a id="L490"></a>return x.Fmt_f64(float64(a));
<a id="L491"></a>}

<a id="L493"></a>func (x *Fmt) e(a float) *Fmt {
    <a id="L494"></a>if strconv.FloatSize == 32 {
        <a id="L495"></a>return x.Fmt_e32(float32(a))
    <a id="L496"></a>}
    <a id="L497"></a>return x.Fmt_e64(float64(a));
<a id="L498"></a>}

<a id="L500"></a>func (x *Fmt) g(a float) *Fmt {
    <a id="L501"></a>if strconv.FloatSize == 32 {
        <a id="L502"></a>return x.Fmt_g32(float32(a))
    <a id="L503"></a>}
    <a id="L504"></a>return x.Fmt_g64(float64(a));
<a id="L505"></a>}

<a id="L507"></a>func (x *Fmt) fb(a float) *Fmt {
    <a id="L508"></a>if strconv.FloatSize == 32 {
        <a id="L509"></a>return x.Fmt_fb32(float32(a))
    <a id="L510"></a>}
    <a id="L511"></a>return x.Fmt_fb64(float64(a));
<a id="L512"></a>}
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
