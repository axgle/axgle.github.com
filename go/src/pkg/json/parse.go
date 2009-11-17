<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/json/parse.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/json/parse.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// JSON (JavaScript Object Notation) parser.</span>
<a id="L6"></a><span class="comment">// See http://www.json.org/</span>

<a id="L8"></a><span class="comment">// The json package implements a simple parser and</span>
<a id="L9"></a><span class="comment">// representation for JSON (JavaScript Object Notation),</span>
<a id="L10"></a><span class="comment">// as defined at http://www.json.org/.</span>
<a id="L11"></a>package json

<a id="L13"></a>import (
    <a id="L14"></a>&#34;bytes&#34;;
    <a id="L15"></a>&#34;strconv&#34;;
    <a id="L16"></a>&#34;utf8&#34;;
<a id="L17"></a>)

<a id="L19"></a><span class="comment">// Strings</span>
<a id="L20"></a><span class="comment">//</span>
<a id="L21"></a><span class="comment">//   Double quoted with escapes: \&#34; \\ \/ \b \f \n \r \t \uXXXX.</span>
<a id="L22"></a><span class="comment">//   No literal control characters, supposedly.</span>
<a id="L23"></a><span class="comment">//   Have also seen \&#39; and embedded newlines.</span>

<a id="L25"></a>func _UnHex(p string, r, l int) (v int, ok bool) {
    <a id="L26"></a>v = 0;
    <a id="L27"></a>for i := r; i &lt; l; i++ {
        <a id="L28"></a>if i &gt;= len(p) {
            <a id="L29"></a>return 0, false
        <a id="L30"></a>}
        <a id="L31"></a>v *= 16;
        <a id="L32"></a>switch {
        <a id="L33"></a>case &#39;0&#39; &lt;= p[i] &amp;&amp; p[i] &lt;= &#39;9&#39;:
            <a id="L34"></a>v += int(p[i] - &#39;0&#39;)
        <a id="L35"></a>case &#39;a&#39; &lt;= p[i] &amp;&amp; p[i] &lt;= &#39;f&#39;:
            <a id="L36"></a>v += int(p[i] - &#39;a&#39; + 10)
        <a id="L37"></a>case &#39;A&#39; &lt;= p[i] &amp;&amp; p[i] &lt;= &#39;F&#39;:
            <a id="L38"></a>v += int(p[i] - &#39;A&#39; + 10)
        <a id="L39"></a>default:
            <a id="L40"></a>return 0, false
        <a id="L41"></a>}
    <a id="L42"></a>}
    <a id="L43"></a>return v, true;
<a id="L44"></a>}

<a id="L46"></a>func _ToHex(b []byte, rune int) {
    <a id="L47"></a>const hexDigits = &#34;0123456789abcdef&#34;;
    <a id="L48"></a>b[0] = hexDigits[rune&gt;&gt;12&amp;0xf];
    <a id="L49"></a>b[1] = hexDigits[rune&gt;&gt;8&amp;0xf];
    <a id="L50"></a>b[2] = hexDigits[rune&gt;&gt;4&amp;0xf];
    <a id="L51"></a>b[3] = hexDigits[rune&amp;0xf];
<a id="L52"></a>}

<a id="L54"></a><span class="comment">// Unquote unquotes the JSON-quoted string s,</span>
<a id="L55"></a><span class="comment">// returning a raw string t.  If s is not a valid</span>
<a id="L56"></a><span class="comment">// JSON-quoted string, Unquote returns with ok set to false.</span>
<a id="L57"></a>func Unquote(s string) (t string, ok bool) {
    <a id="L58"></a>if len(s) &lt; 2 || s[0] != &#39;&#34;&#39; || s[len(s)-1] != &#39;&#34;&#39; {
        <a id="L59"></a>return
    <a id="L60"></a>}
    <a id="L61"></a>b := make([]byte, len(s));
    <a id="L62"></a>w := 0;
    <a id="L63"></a>for r := 1; r &lt; len(s)-1; {
        <a id="L64"></a>switch {
        <a id="L65"></a>case s[r] == &#39;\\&#39;:
            <a id="L66"></a>r++;
            <a id="L67"></a>if r &gt;= len(s)-1 {
                <a id="L68"></a>return
            <a id="L69"></a>}
            <a id="L70"></a>switch s[r] {
            <a id="L71"></a>default:
                <a id="L72"></a>return
            <a id="L73"></a>case &#39;&#34;&#39;, &#39;\\&#39;, &#39;/&#39;, &#39;\&#39;&#39;:
                <a id="L74"></a>b[w] = s[r];
                <a id="L75"></a>r++;
                <a id="L76"></a>w++;
            <a id="L77"></a>case &#39;b&#39;:
                <a id="L78"></a>b[w] = &#39;\b&#39;;
                <a id="L79"></a>r++;
                <a id="L80"></a>w++;
            <a id="L81"></a>case &#39;f&#39;:
                <a id="L82"></a>b[w] = &#39;\f&#39;;
                <a id="L83"></a>r++;
                <a id="L84"></a>w++;
            <a id="L85"></a>case &#39;n&#39;:
                <a id="L86"></a>b[w] = &#39;\n&#39;;
                <a id="L87"></a>r++;
                <a id="L88"></a>w++;
            <a id="L89"></a>case &#39;r&#39;:
                <a id="L90"></a>b[w] = &#39;\r&#39;;
                <a id="L91"></a>r++;
                <a id="L92"></a>w++;
            <a id="L93"></a>case &#39;t&#39;:
                <a id="L94"></a>b[w] = &#39;\t&#39;;
                <a id="L95"></a>r++;
                <a id="L96"></a>w++;
            <a id="L97"></a>case &#39;u&#39;:
                <a id="L98"></a>r++;
                <a id="L99"></a>rune, ok := _UnHex(s, r, r+4);
                <a id="L100"></a>if !ok {
                    <a id="L101"></a>return
                <a id="L102"></a>}
                <a id="L103"></a>r += 4;
                <a id="L104"></a>w += utf8.EncodeRune(rune, b[w:len(b)]);
            <a id="L105"></a>}
        <a id="L106"></a><span class="comment">// Control characters are invalid, but we&#39;ve seen raw \n.</span>
        <a id="L107"></a>case s[r] &lt; &#39; &#39; &amp;&amp; s[r] != &#39;\n&#39;:
            <a id="L108"></a>if s[r] == &#39;\n&#39; {
                <a id="L109"></a>b[w] = &#39;\n&#39;;
                <a id="L110"></a>r++;
                <a id="L111"></a>w++;
                <a id="L112"></a>break;
            <a id="L113"></a>}
            <a id="L114"></a>return;
        <a id="L115"></a><span class="comment">// ASCII</span>
        <a id="L116"></a>case s[r] &lt; utf8.RuneSelf:
            <a id="L117"></a>b[w] = s[r];
            <a id="L118"></a>r++;
            <a id="L119"></a>w++;
        <a id="L120"></a><span class="comment">// Coerce to well-formed UTF-8.</span>
        <a id="L121"></a>default:
            <a id="L122"></a>rune, size := utf8.DecodeRuneInString(s[r:len(s)]);
            <a id="L123"></a>r += size;
            <a id="L124"></a>w += utf8.EncodeRune(rune, b[w:len(b)]);
        <a id="L125"></a>}
    <a id="L126"></a>}
    <a id="L127"></a>return string(b[0:w]), true;
<a id="L128"></a>}

<a id="L130"></a><span class="comment">// Quote quotes the raw string s using JSON syntax,</span>
<a id="L131"></a><span class="comment">// so that Unquote(Quote(s)) = s, true.</span>
<a id="L132"></a>func Quote(s string) string {
    <a id="L133"></a>chr := make([]byte, 6);
    <a id="L134"></a>chr0 := chr[0:1];
    <a id="L135"></a>b := new(bytes.Buffer);
    <a id="L136"></a>chr[0] = &#39;&#34;&#39;;
    <a id="L137"></a>b.Write(chr0);

    <a id="L139"></a>for _, rune := range s {
        <a id="L140"></a>switch {
        <a id="L141"></a>case rune == &#39;&#34;&#39; || rune == &#39;\\&#39;:
            <a id="L142"></a>chr[0] = &#39;\\&#39;;
            <a id="L143"></a>chr[1] = byte(rune);
            <a id="L144"></a>b.Write(chr[0:2]);

        <a id="L146"></a>case rune == &#39;\b&#39;:
            <a id="L147"></a>chr[0] = &#39;\\&#39;;
            <a id="L148"></a>chr[1] = &#39;b&#39;;
            <a id="L149"></a>b.Write(chr[0:2]);

        <a id="L151"></a>case rune == &#39;\f&#39;:
            <a id="L152"></a>chr[0] = &#39;\\&#39;;
            <a id="L153"></a>chr[1] = &#39;f&#39;;
            <a id="L154"></a>b.Write(chr[0:2]);

        <a id="L156"></a>case rune == &#39;\n&#39;:
            <a id="L157"></a>chr[0] = &#39;\\&#39;;
            <a id="L158"></a>chr[1] = &#39;n&#39;;
            <a id="L159"></a>b.Write(chr[0:2]);

        <a id="L161"></a>case rune == &#39;\r&#39;:
            <a id="L162"></a>chr[0] = &#39;\\&#39;;
            <a id="L163"></a>chr[1] = &#39;r&#39;;
            <a id="L164"></a>b.Write(chr[0:2]);

        <a id="L166"></a>case rune == &#39;\t&#39;:
            <a id="L167"></a>chr[0] = &#39;\\&#39;;
            <a id="L168"></a>chr[1] = &#39;t&#39;;
            <a id="L169"></a>b.Write(chr[0:2]);

        <a id="L171"></a>case 0x20 &lt;= rune &amp;&amp; rune &lt; utf8.RuneSelf:
            <a id="L172"></a>chr[0] = byte(rune);
            <a id="L173"></a>b.Write(chr0);

        <a id="L175"></a>default:
            <a id="L176"></a>chr[0] = &#39;\\&#39;;
            <a id="L177"></a>chr[1] = &#39;u&#39;;
            <a id="L178"></a>_ToHex(chr[2:6], rune);
            <a id="L179"></a>b.Write(chr);
        <a id="L180"></a>}
    <a id="L181"></a>}
    <a id="L182"></a>chr[0] = &#39;&#34;&#39;;
    <a id="L183"></a>b.Write(chr0);
    <a id="L184"></a>return b.String();
<a id="L185"></a>}


<a id="L188"></a><span class="comment">// _Lexer</span>

<a id="L190"></a>type _Lexer struct {
    <a id="L191"></a>s     string;
    <a id="L192"></a>i     int;
    <a id="L193"></a>kind  int;
    <a id="L194"></a>token string;
<a id="L195"></a>}

<a id="L197"></a>func punct(c byte) bool {
    <a id="L198"></a>return c == &#39;&#34;&#39; || c == &#39;[&#39; || c == &#39;]&#39; || c == &#39;:&#39; || c == &#39;{&#39; || c == &#39;}&#39; || c == &#39;,&#39;
<a id="L199"></a>}

<a id="L201"></a>func white(c byte) bool { return c == &#39; &#39; || c == &#39;\t&#39; || c == &#39;\n&#39; || c == &#39;\v&#39; }

<a id="L203"></a>func skipwhite(p string, i int) int {
    <a id="L204"></a>for i &lt; len(p) &amp;&amp; white(p[i]) {
        <a id="L205"></a>i++
    <a id="L206"></a>}
    <a id="L207"></a>return i;
<a id="L208"></a>}

<a id="L210"></a>func skiptoken(p string, i int) int {
    <a id="L211"></a>for i &lt; len(p) &amp;&amp; !punct(p[i]) &amp;&amp; !white(p[i]) {
        <a id="L212"></a>i++
    <a id="L213"></a>}
    <a id="L214"></a>return i;
<a id="L215"></a>}

<a id="L217"></a>func skipstring(p string, i int) int {
    <a id="L218"></a>for i++; i &lt; len(p) &amp;&amp; p[i] != &#39;&#34;&#39;; i++ {
        <a id="L219"></a>if p[i] == &#39;\\&#39; {
            <a id="L220"></a>i++
        <a id="L221"></a>}
    <a id="L222"></a>}
    <a id="L223"></a>if i &gt;= len(p) {
        <a id="L224"></a>return i
    <a id="L225"></a>}
    <a id="L226"></a>return i + 1;
<a id="L227"></a>}

<a id="L229"></a>func (t *_Lexer) Next() {
    <a id="L230"></a>i, s := t.i, t.s;
    <a id="L231"></a>i = skipwhite(s, i);
    <a id="L232"></a>if i &gt;= len(s) {
        <a id="L233"></a>t.kind = 0;
        <a id="L234"></a>t.token = &#34;&#34;;
        <a id="L235"></a>t.i = len(s);
        <a id="L236"></a>return;
    <a id="L237"></a>}

    <a id="L239"></a>c := s[i];
    <a id="L240"></a>switch {
    <a id="L241"></a>case c == &#39;-&#39; || &#39;0&#39; &lt;= c &amp;&amp; c &lt;= &#39;9&#39;:
        <a id="L242"></a>j := skiptoken(s, i);
        <a id="L243"></a>t.kind = &#39;1&#39;;
        <a id="L244"></a>t.token = s[i:j];
        <a id="L245"></a>i = j;

    <a id="L247"></a>case &#39;a&#39; &lt;= c &amp;&amp; c &lt;= &#39;z&#39; || &#39;A&#39; &lt;= c &amp;&amp; c &lt;= &#39;Z&#39;:
        <a id="L248"></a>j := skiptoken(s, i);
        <a id="L249"></a>t.kind = &#39;a&#39;;
        <a id="L250"></a>t.token = s[i:j];
        <a id="L251"></a>i = j;

    <a id="L253"></a>case c == &#39;&#34;&#39;:
        <a id="L254"></a>j := skipstring(s, i);
        <a id="L255"></a>t.kind = &#39;&#34;&#39;;
        <a id="L256"></a>t.token = s[i:j];
        <a id="L257"></a>i = j;

    <a id="L259"></a>case c == &#39;[&#39;, c == &#39;]&#39;, c == &#39;:&#39;, c == &#39;{&#39;, c == &#39;}&#39;, c == &#39;,&#39;:
        <a id="L260"></a>t.kind = int(c);
        <a id="L261"></a>t.token = s[i : i+1];
        <a id="L262"></a>i++;

    <a id="L264"></a>default:
        <a id="L265"></a>t.kind = &#39;?&#39;;
        <a id="L266"></a>t.token = s[i : i+1];
    <a id="L267"></a>}

    <a id="L269"></a>t.i = i;
<a id="L270"></a>}


<a id="L273"></a><span class="comment">// Parser</span>
<a id="L274"></a><span class="comment">//</span>
<a id="L275"></a><span class="comment">// Implements parsing but not the actions.  Those are</span>
<a id="L276"></a><span class="comment">// carried out by the implementation of the Builder interface.</span>
<a id="L277"></a><span class="comment">// A Builder represents the object being created.</span>
<a id="L278"></a><span class="comment">// Calling a method like Int64(i) sets that object to i.</span>
<a id="L279"></a><span class="comment">// Calling a method like Elem(i) or Key(s) creates a</span>
<a id="L280"></a><span class="comment">// new builder for a subpiece of the object (logically,</span>
<a id="L281"></a><span class="comment">// an array element or a map key).</span>
<a id="L282"></a><span class="comment">//</span>
<a id="L283"></a><span class="comment">// There are two Builders, in other files.</span>
<a id="L284"></a><span class="comment">// The JsonBuilder builds a generic Json structure</span>
<a id="L285"></a><span class="comment">// in which maps are maps.</span>
<a id="L286"></a><span class="comment">// The StructBuilder copies data into a possibly</span>
<a id="L287"></a><span class="comment">// nested data structure, using the &#34;map keys&#34;</span>
<a id="L288"></a><span class="comment">// as struct field names.</span>

<a id="L290"></a>type _Value interface{}

<a id="L292"></a><span class="comment">// BUG(rsc): The json Builder interface needs to be</span>
<a id="L293"></a><span class="comment">// reconciled with the xml Builder interface.</span>

<a id="L295"></a><span class="comment">// A Builder is an interface implemented by clients and passed</span>
<a id="L296"></a><span class="comment">// to the JSON parser.  It gives clients full control over the</span>
<a id="L297"></a><span class="comment">// eventual representation returned by the parser.</span>
<a id="L298"></a>type Builder interface {
    <a id="L299"></a><span class="comment">// Set value</span>
    <a id="L300"></a>Int64(i int64);
    <a id="L301"></a>Uint64(i uint64);
    <a id="L302"></a>Float64(f float64);
    <a id="L303"></a>String(s string);
    <a id="L304"></a>Bool(b bool);
    <a id="L305"></a>Null();
    <a id="L306"></a>Array();
    <a id="L307"></a>Map();

    <a id="L309"></a><span class="comment">// Create sub-Builders</span>
    <a id="L310"></a>Elem(i int) Builder;
    <a id="L311"></a>Key(s string) Builder;

    <a id="L313"></a><span class="comment">// Flush changes to parent Builder if necessary.</span>
    <a id="L314"></a>Flush();
<a id="L315"></a>}

<a id="L317"></a>func parse(lex *_Lexer, build Builder) bool {
    <a id="L318"></a>ok := false;
<a id="L319"></a>Switch:
    <a id="L320"></a>switch lex.kind {
    <a id="L321"></a>case 0:
        <a id="L322"></a>break
    <a id="L323"></a>case &#39;1&#39;:
        <a id="L324"></a><span class="comment">// If the number is exactly an integer, use that.</span>
        <a id="L325"></a>if i, err := strconv.Atoi64(lex.token); err == nil {
            <a id="L326"></a>build.Int64(i);
            <a id="L327"></a>ok = true;
        <a id="L328"></a>} else if i, err := strconv.Atoui64(lex.token); err == nil {
            <a id="L329"></a>build.Uint64(i);
            <a id="L330"></a>ok = true;
        <a id="L331"></a>} else
        <a id="L332"></a><span class="comment">// Fall back to floating point.</span>
        <a id="L333"></a>if f, err := strconv.Atof64(lex.token); err == nil {
            <a id="L334"></a>build.Float64(f);
            <a id="L335"></a>ok = true;
        <a id="L336"></a>}

    <a id="L338"></a>case &#39;a&#39;:
        <a id="L339"></a>switch lex.token {
        <a id="L340"></a>case &#34;true&#34;:
            <a id="L341"></a>build.Bool(true);
            <a id="L342"></a>ok = true;
        <a id="L343"></a>case &#34;false&#34;:
            <a id="L344"></a>build.Bool(false);
            <a id="L345"></a>ok = true;
        <a id="L346"></a>case &#34;null&#34;:
            <a id="L347"></a>build.Null();
            <a id="L348"></a>ok = true;
        <a id="L349"></a>}

    <a id="L351"></a>case &#39;&#34;&#39;:
        <a id="L352"></a>if str, ok1 := Unquote(lex.token); ok1 {
            <a id="L353"></a>build.String(str);
            <a id="L354"></a>ok = true;
        <a id="L355"></a>}

    <a id="L357"></a>case &#39;[&#39;:
        <a id="L358"></a><span class="comment">// array</span>
        <a id="L359"></a>build.Array();
        <a id="L360"></a>lex.Next();
        <a id="L361"></a>n := 0;
        <a id="L362"></a>for lex.kind != &#39;]&#39; {
            <a id="L363"></a>if n &gt; 0 {
                <a id="L364"></a>if lex.kind != &#39;,&#39; {
                    <a id="L365"></a>break Switch
                <a id="L366"></a>}
                <a id="L367"></a>lex.Next();
            <a id="L368"></a>}
            <a id="L369"></a>if !parse(lex, build.Elem(n)) {
                <a id="L370"></a>break Switch
            <a id="L371"></a>}
            <a id="L372"></a>n++;
        <a id="L373"></a>}
        <a id="L374"></a>ok = true;

    <a id="L376"></a>case &#39;{&#39;:
        <a id="L377"></a><span class="comment">// map</span>
        <a id="L378"></a>lex.Next();
        <a id="L379"></a>build.Map();
        <a id="L380"></a>n := 0;
        <a id="L381"></a>for lex.kind != &#39;}&#39; {
            <a id="L382"></a>if n &gt; 0 {
                <a id="L383"></a>if lex.kind != &#39;,&#39; {
                    <a id="L384"></a>break Switch
                <a id="L385"></a>}
                <a id="L386"></a>lex.Next();
            <a id="L387"></a>}
            <a id="L388"></a>if lex.kind != &#39;&#34;&#39; {
                <a id="L389"></a>break Switch
            <a id="L390"></a>}
            <a id="L391"></a>key, ok := Unquote(lex.token);
            <a id="L392"></a>if !ok {
                <a id="L393"></a>break Switch
            <a id="L394"></a>}
            <a id="L395"></a>lex.Next();
            <a id="L396"></a>if lex.kind != &#39;:&#39; {
                <a id="L397"></a>break Switch
            <a id="L398"></a>}
            <a id="L399"></a>lex.Next();
            <a id="L400"></a>if !parse(lex, build.Key(key)) {
                <a id="L401"></a>break Switch
            <a id="L402"></a>}
            <a id="L403"></a>n++;
        <a id="L404"></a>}
        <a id="L405"></a>ok = true;
    <a id="L406"></a>}

    <a id="L408"></a>if ok {
        <a id="L409"></a>lex.Next()
    <a id="L410"></a>}
    <a id="L411"></a>build.Flush();
    <a id="L412"></a>return ok;
<a id="L413"></a>}

<a id="L415"></a><span class="comment">// Parse parses the JSON syntax string s and makes calls to</span>
<a id="L416"></a><span class="comment">// the builder to construct a parsed representation.</span>
<a id="L417"></a><span class="comment">// On success, it returns with ok set to true.</span>
<a id="L418"></a><span class="comment">// On error, it returns with ok set to false, errindx set</span>
<a id="L419"></a><span class="comment">// to the byte index in s where a syntax error occurred,</span>
<a id="L420"></a><span class="comment">// and errtok set to the offending token.</span>
<a id="L421"></a>func Parse(s string, builder Builder) (ok bool, errindx int, errtok string) {
    <a id="L422"></a>lex := new(_Lexer);
    <a id="L423"></a>lex.s = s;
    <a id="L424"></a>lex.Next();
    <a id="L425"></a>if parse(lex, builder) {
        <a id="L426"></a>if lex.kind == 0 { <span class="comment">// EOF</span>
            <a id="L427"></a>return true, 0, &#34;&#34;
        <a id="L428"></a>}
    <a id="L429"></a>}
    <a id="L430"></a>return false, lex.i, lex.token;
<a id="L431"></a>}
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
