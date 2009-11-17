<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/tabwriter/tabwriter_test.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/tabwriter/tabwriter_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package tabwriter

<a id="L7"></a>import (
    <a id="L8"></a>&#34;io&#34;;
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>&#34;testing&#34;;
<a id="L11"></a>)


<a id="L14"></a>type buffer struct {
    <a id="L15"></a>a []byte;
<a id="L16"></a>}


<a id="L19"></a>func (b *buffer) init(n int) { b.a = make([]byte, n)[0:0] }


<a id="L22"></a>func (b *buffer) clear() { b.a = b.a[0:0] }


<a id="L25"></a>func (b *buffer) Write(buf []byte) (written int, err os.Error) {
    <a id="L26"></a>n := len(b.a);
    <a id="L27"></a>m := len(buf);
    <a id="L28"></a>if n+m &lt;= cap(b.a) {
        <a id="L29"></a>b.a = b.a[0 : n+m];
        <a id="L30"></a>for i := 0; i &lt; m; i++ {
            <a id="L31"></a>b.a[n+i] = buf[i]
        <a id="L32"></a>}
    <a id="L33"></a>} else {
        <a id="L34"></a>panicln(&#34;buffer.Write: buffer too small&#34;, n, m, cap(b.a))
    <a id="L35"></a>}
    <a id="L36"></a>return len(buf), nil;
<a id="L37"></a>}


<a id="L40"></a>func (b *buffer) String() string { return string(b.a) }


<a id="L43"></a>func write(t *testing.T, testname string, w *Writer, src string) {
    <a id="L44"></a>written, err := io.WriteString(w, src);
    <a id="L45"></a>if err != nil {
        <a id="L46"></a>t.Errorf(&#34;--- test: %s\n--- src:\n%s\n--- write error: %v\n&#34;, testname, src, err)
    <a id="L47"></a>}
    <a id="L48"></a>if written != len(src) {
        <a id="L49"></a>t.Errorf(&#34;--- test: %s\n--- src:\n%s\n--- written = %d, len(src) = %d\n&#34;, testname, src, written, len(src))
    <a id="L50"></a>}
<a id="L51"></a>}


<a id="L54"></a>func verify(t *testing.T, testname string, w *Writer, b *buffer, src, expected string) {
    <a id="L55"></a>err := w.Flush();
    <a id="L56"></a>if err != nil {
        <a id="L57"></a>t.Errorf(&#34;--- test: %s\n--- src:\n%s\n--- flush error: %v\n&#34;, testname, src, err)
    <a id="L58"></a>}

    <a id="L60"></a>res := b.String();
    <a id="L61"></a>if res != expected {
        <a id="L62"></a>t.Errorf(&#34;--- test: %s\n--- src:\n%s\n--- found:\n%s\n--- expected:\n%s\n&#34;, testname, src, res, expected)
    <a id="L63"></a>}
<a id="L64"></a>}


<a id="L67"></a>func check(t *testing.T, testname string, tabwidth, padding int, padchar byte, flags uint, src, expected string) {
    <a id="L68"></a>var b buffer;
    <a id="L69"></a>b.init(1000);

    <a id="L71"></a>var w Writer;
    <a id="L72"></a>w.Init(&amp;b, tabwidth, padding, padchar, flags);

    <a id="L74"></a><span class="comment">// write all at once</span>
    <a id="L75"></a>b.clear();
    <a id="L76"></a>write(t, testname, &amp;w, src);
    <a id="L77"></a>verify(t, testname, &amp;w, &amp;b, src, expected);

    <a id="L79"></a><span class="comment">// write byte-by-byte</span>
    <a id="L80"></a>b.clear();
    <a id="L81"></a>for i := 0; i &lt; len(src); i++ {
        <a id="L82"></a>write(t, testname, &amp;w, src[i:i+1])
    <a id="L83"></a>}
    <a id="L84"></a>verify(t, testname, &amp;w, &amp;b, src, expected);

    <a id="L86"></a><span class="comment">// write using Fibonacci slice sizes</span>
    <a id="L87"></a>b.clear();
    <a id="L88"></a>for i, d := 0, 0; i &lt; len(src); {
        <a id="L89"></a>write(t, testname, &amp;w, src[i:i+d]);
        <a id="L90"></a>i, d = i+d, d+1;
        <a id="L91"></a>if i+d &gt; len(src) {
            <a id="L92"></a>d = len(src) - i
        <a id="L93"></a>}
    <a id="L94"></a>}
    <a id="L95"></a>verify(t, testname, &amp;w, &amp;b, src, expected);
<a id="L96"></a>}


<a id="L99"></a>type entry struct {
    <a id="L100"></a>testname          string;
    <a id="L101"></a>tabwidth, padding int;
    <a id="L102"></a>padchar           byte;
    <a id="L103"></a>flags             uint;
    <a id="L104"></a>src, expected     string;
<a id="L105"></a>}


<a id="L108"></a>var tests = []entry{
    <a id="L109"></a>entry{
        <a id="L110"></a>&#34;1a&#34;,
        <a id="L111"></a>8, 1, &#39;.&#39;, 0,
        <a id="L112"></a>&#34;&#34;,
        <a id="L113"></a>&#34;&#34;,
    <a id="L114"></a>},

    <a id="L116"></a>entry{
        <a id="L117"></a>&#34;1a debug&#34;,
        <a id="L118"></a>8, 1, &#39;.&#39;, Debug,
        <a id="L119"></a>&#34;&#34;,
        <a id="L120"></a>&#34;&#34;,
    <a id="L121"></a>},

    <a id="L123"></a>entry{
        <a id="L124"></a>&#34;1b esc&#34;,
        <a id="L125"></a>8, 1, &#39;.&#39;, 0,
        <a id="L126"></a>&#34;\xff\xff&#34;,
        <a id="L127"></a>&#34;&#34;,
    <a id="L128"></a>},

    <a id="L130"></a>entry{
        <a id="L131"></a>&#34;1c esc&#34;,
        <a id="L132"></a>8, 1, &#39;.&#39;, 0,
        <a id="L133"></a>&#34;\xff\t\xff&#34;,
        <a id="L134"></a>&#34;\t&#34;,
    <a id="L135"></a>},

    <a id="L137"></a>entry{
        <a id="L138"></a>&#34;1d esc&#34;,
        <a id="L139"></a>8, 1, &#39;.&#39;, 0,
        <a id="L140"></a>&#34;\xff\&#34;foo\t\n\tbar\&#34;\xff&#34;,
        <a id="L141"></a>&#34;\&#34;foo\t\n\tbar\&#34;&#34;,
    <a id="L142"></a>},

    <a id="L144"></a>entry{
        <a id="L145"></a>&#34;1e esc&#34;,
        <a id="L146"></a>8, 1, &#39;.&#39;, 0,
        <a id="L147"></a>&#34;abc\xff\tdef&#34;, <span class="comment">// unterminated escape</span>
        <a id="L148"></a>&#34;abc\tdef&#34;,
    <a id="L149"></a>},

    <a id="L151"></a>entry{
        <a id="L152"></a>&#34;2&#34;,
        <a id="L153"></a>8, 1, &#39;.&#39;, 0,
        <a id="L154"></a>&#34;\n\n\n&#34;,
        <a id="L155"></a>&#34;\n\n\n&#34;,
    <a id="L156"></a>},

    <a id="L158"></a>entry{
        <a id="L159"></a>&#34;3&#34;,
        <a id="L160"></a>8, 1, &#39;.&#39;, 0,
        <a id="L161"></a>&#34;a\nb\nc&#34;,
        <a id="L162"></a>&#34;a\nb\nc&#34;,
    <a id="L163"></a>},

    <a id="L165"></a>entry{
        <a id="L166"></a>&#34;4a&#34;,
        <a id="L167"></a>8, 1, &#39;.&#39;, 0,
        <a id="L168"></a>&#34;\t&#34;, <span class="comment">// &#39;\t&#39; terminates an empty cell on last line - nothing to print</span>
        <a id="L169"></a>&#34;&#34;,
    <a id="L170"></a>},

    <a id="L172"></a>entry{
        <a id="L173"></a>&#34;4b&#34;,
        <a id="L174"></a>8, 1, &#39;.&#39;, AlignRight,
        <a id="L175"></a>&#34;\t&#34;, <span class="comment">// &#39;\t&#39; terminates an empty cell on last line - nothing to print</span>
        <a id="L176"></a>&#34;&#34;,
    <a id="L177"></a>},

    <a id="L179"></a>entry{
        <a id="L180"></a>&#34;5&#34;,
        <a id="L181"></a>8, 1, &#39;.&#39;, 0,
        <a id="L182"></a>&#34;*\t*&#34;,
        <a id="L183"></a>&#34;*.......*&#34;,
    <a id="L184"></a>},

    <a id="L186"></a>entry{
        <a id="L187"></a>&#34;5b&#34;,
        <a id="L188"></a>8, 1, &#39;.&#39;, 0,
        <a id="L189"></a>&#34;*\t*\n&#34;,
        <a id="L190"></a>&#34;*.......*\n&#34;,
    <a id="L191"></a>},

    <a id="L193"></a>entry{
        <a id="L194"></a>&#34;5c&#34;,
        <a id="L195"></a>8, 1, &#39;.&#39;, 0,
        <a id="L196"></a>&#34;*\t*\t&#34;,
        <a id="L197"></a>&#34;*.......*&#34;,
    <a id="L198"></a>},

    <a id="L200"></a>entry{
        <a id="L201"></a>&#34;5c debug&#34;,
        <a id="L202"></a>8, 1, &#39;.&#39;, Debug,
        <a id="L203"></a>&#34;*\t*\t&#34;,
        <a id="L204"></a>&#34;*.......|*&#34;,
    <a id="L205"></a>},

    <a id="L207"></a>entry{
        <a id="L208"></a>&#34;5d&#34;,
        <a id="L209"></a>8, 1, &#39;.&#39;, AlignRight,
        <a id="L210"></a>&#34;*\t*\t&#34;,
        <a id="L211"></a>&#34;.......**&#34;,
    <a id="L212"></a>},

    <a id="L214"></a>entry{
        <a id="L215"></a>&#34;6&#34;,
        <a id="L216"></a>8, 1, &#39;.&#39;, 0,
        <a id="L217"></a>&#34;\t\n&#34;,
        <a id="L218"></a>&#34;........\n&#34;,
    <a id="L219"></a>},

    <a id="L221"></a>entry{
        <a id="L222"></a>&#34;7a&#34;,
        <a id="L223"></a>8, 1, &#39;.&#39;, 0,
        <a id="L224"></a>&#34;a) foo&#34;,
        <a id="L225"></a>&#34;a) foo&#34;,
    <a id="L226"></a>},

    <a id="L228"></a>entry{
        <a id="L229"></a>&#34;7b&#34;,
        <a id="L230"></a>8, 1, &#39; &#39;, 0,
        <a id="L231"></a>&#34;b) foo\tbar&#34;,
        <a id="L232"></a>&#34;b) foo  bar&#34;,
    <a id="L233"></a>},

    <a id="L235"></a>entry{
        <a id="L236"></a>&#34;7c&#34;,
        <a id="L237"></a>8, 1, &#39;.&#39;, 0,
        <a id="L238"></a>&#34;c) foo\tbar\t&#34;,
        <a id="L239"></a>&#34;c) foo..bar&#34;,
    <a id="L240"></a>},

    <a id="L242"></a>entry{
        <a id="L243"></a>&#34;7d&#34;,
        <a id="L244"></a>8, 1, &#39;.&#39;, 0,
        <a id="L245"></a>&#34;d) foo\tbar\n&#34;,
        <a id="L246"></a>&#34;d) foo..bar\n&#34;,
    <a id="L247"></a>},

    <a id="L249"></a>entry{
        <a id="L250"></a>&#34;7e&#34;,
        <a id="L251"></a>8, 1, &#39;.&#39;, 0,
        <a id="L252"></a>&#34;e) foo\tbar\t\n&#34;,
        <a id="L253"></a>&#34;e) foo..bar.....\n&#34;,
    <a id="L254"></a>},

    <a id="L256"></a>entry{
        <a id="L257"></a>&#34;7f&#34;,
        <a id="L258"></a>8, 1, &#39;.&#39;, FilterHTML,
        <a id="L259"></a>&#34;f) f&amp;lt;o\t&lt;b&gt;bar&lt;/b&gt;\t\n&#34;,
        <a id="L260"></a>&#34;f) f&amp;lt;o..&lt;b&gt;bar&lt;/b&gt;.....\n&#34;,
    <a id="L261"></a>},

    <a id="L263"></a>entry{
        <a id="L264"></a>&#34;7g&#34;,
        <a id="L265"></a>8, 1, &#39;.&#39;, FilterHTML,
        <a id="L266"></a>&#34;g) f&amp;lt;o\t&lt;b&gt;bar&lt;/b&gt;\t non-terminated entity &amp;amp&#34;,
        <a id="L267"></a>&#34;g) f&amp;lt;o..&lt;b&gt;bar&lt;/b&gt;..... non-terminated entity &amp;amp&#34;,
    <a id="L268"></a>},

    <a id="L270"></a>entry{
        <a id="L271"></a>&#34;7g debug&#34;,
        <a id="L272"></a>8, 1, &#39;.&#39;, FilterHTML | Debug,
        <a id="L273"></a>&#34;g) f&amp;lt;o\t&lt;b&gt;bar&lt;/b&gt;\t non-terminated entity &amp;amp&#34;,
        <a id="L274"></a>&#34;g) f&amp;lt;o..|&lt;b&gt;bar&lt;/b&gt;.....| non-terminated entity &amp;amp&#34;,
    <a id="L275"></a>},

    <a id="L277"></a>entry{
        <a id="L278"></a>&#34;8&#34;,
        <a id="L279"></a>8, 1, &#39;*&#39;, 0,
        <a id="L280"></a>&#34;Hello, world!\n&#34;,
        <a id="L281"></a>&#34;Hello, world!\n&#34;,
    <a id="L282"></a>},

    <a id="L284"></a>entry{
        <a id="L285"></a>&#34;9a&#34;,
        <a id="L286"></a>1, 0, &#39;.&#39;, 0,
        <a id="L287"></a>&#34;1\t2\t3\t4\n&#34;
            <a id="L288"></a>&#34;11\t222\t3333\t44444\n&#34;,

        <a id="L290"></a>&#34;1.2..3...4\n&#34;
            <a id="L291"></a>&#34;11222333344444\n&#34;,
    <a id="L292"></a>},

    <a id="L294"></a>entry{
        <a id="L295"></a>&#34;9b&#34;,
        <a id="L296"></a>1, 0, &#39;.&#39;, FilterHTML,
        <a id="L297"></a>&#34;1\t2&lt;!---\f---&gt;\t3\t4\n&#34; <span class="comment">// \f inside HTML is ignored</span>
            <a id="L298"></a>&#34;11\t222\t3333\t44444\n&#34;,

        <a id="L300"></a>&#34;1.2&lt;!---\f---&gt;..3...4\n&#34;
            <a id="L301"></a>&#34;11222333344444\n&#34;,
    <a id="L302"></a>},

    <a id="L304"></a>entry{
        <a id="L305"></a>&#34;9c&#34;,
        <a id="L306"></a>1, 0, &#39;.&#39;, 0,
        <a id="L307"></a>&#34;1\t2\t3\t4\f&#34; <span class="comment">// \f causes a newline and flush</span>
            <a id="L308"></a>&#34;11\t222\t3333\t44444\n&#34;,

        <a id="L310"></a>&#34;1234\n&#34;
            <a id="L311"></a>&#34;11222333344444\n&#34;,
    <a id="L312"></a>},

    <a id="L314"></a>entry{
        <a id="L315"></a>&#34;9c debug&#34;,
        <a id="L316"></a>1, 0, &#39;.&#39;, Debug,
        <a id="L317"></a>&#34;1\t2\t3\t4\f&#34; <span class="comment">// \f causes a newline and flush</span>
            <a id="L318"></a>&#34;11\t222\t3333\t44444\n&#34;,

        <a id="L320"></a>&#34;1|2|3|4\n&#34;
            <a id="L321"></a>&#34;11|222|3333|44444\n&#34;,
    <a id="L322"></a>},

    <a id="L324"></a>entry{
        <a id="L325"></a>&#34;10a&#34;,
        <a id="L326"></a>5, 0, &#39;.&#39;, 0,
        <a id="L327"></a>&#34;1\t2\t3\t4\n&#34;,
        <a id="L328"></a>&#34;1....2....3....4\n&#34;,
    <a id="L329"></a>},

    <a id="L331"></a>entry{
        <a id="L332"></a>&#34;10b&#34;,
        <a id="L333"></a>5, 0, &#39;.&#39;, 0,
        <a id="L334"></a>&#34;1\t2\t3\t4\t\n&#34;,
        <a id="L335"></a>&#34;1....2....3....4....\n&#34;,
    <a id="L336"></a>},

    <a id="L338"></a>entry{
        <a id="L339"></a>&#34;11&#34;,
        <a id="L340"></a>8, 1, &#39;.&#39;, 0,
        <a id="L341"></a>&#34;本\tb\tc\n&#34;
            <a id="L342"></a>&#34;aa\t\u672c\u672c\u672c\tcccc\tddddd\n&#34;
            <a id="L343"></a>&#34;aaa\tbbbb\n&#34;,

        <a id="L345"></a>&#34;本.......b.......c\n&#34;
            <a id="L346"></a>&#34;aa......本本本.....cccc....ddddd\n&#34;
            <a id="L347"></a>&#34;aaa.....bbbb\n&#34;,
    <a id="L348"></a>},

    <a id="L350"></a>entry{
        <a id="L351"></a>&#34;12a&#34;,
        <a id="L352"></a>8, 1, &#39; &#39;, AlignRight,
        <a id="L353"></a>&#34;a\tè\tc\t\n&#34;
            <a id="L354"></a>&#34;aa\tèèè\tcccc\tddddd\t\n&#34;
            <a id="L355"></a>&#34;aaa\tèèèè\t\n&#34;,

        <a id="L357"></a>&#34;       a       è       c\n&#34;
            <a id="L358"></a>&#34;      aa     èèè    cccc   ddddd\n&#34;
            <a id="L359"></a>&#34;     aaa    èèèè\n&#34;,
    <a id="L360"></a>},

    <a id="L362"></a>entry{
        <a id="L363"></a>&#34;12b&#34;,
        <a id="L364"></a>2, 0, &#39; &#39;, 0,
        <a id="L365"></a>&#34;a\tb\tc\n&#34;
            <a id="L366"></a>&#34;aa\tbbb\tcccc\n&#34;
            <a id="L367"></a>&#34;aaa\tbbbb\n&#34;,

        <a id="L369"></a>&#34;a  b  c\n&#34;
            <a id="L370"></a>&#34;aa bbbcccc\n&#34;
            <a id="L371"></a>&#34;aaabbbb\n&#34;,
    <a id="L372"></a>},

    <a id="L374"></a>entry{
        <a id="L375"></a>&#34;12c&#34;,
        <a id="L376"></a>8, 1, &#39;_&#39;, 0,
        <a id="L377"></a>&#34;a\tb\tc\n&#34;
            <a id="L378"></a>&#34;aa\tbbb\tcccc\n&#34;
            <a id="L379"></a>&#34;aaa\tbbbb\n&#34;,

        <a id="L381"></a>&#34;a_______b_______c\n&#34;
            <a id="L382"></a>&#34;aa______bbb_____cccc\n&#34;
            <a id="L383"></a>&#34;aaa_____bbbb\n&#34;,
    <a id="L384"></a>},

    <a id="L386"></a>entry{
        <a id="L387"></a>&#34;13a&#34;,
        <a id="L388"></a>4, 1, &#39;-&#39;, 0,
        <a id="L389"></a>&#34;4444\t日本語\t22\t1\t333\n&#34;
            <a id="L390"></a>&#34;999999999\t22\n&#34;
            <a id="L391"></a>&#34;7\t22\n&#34;
            <a id="L392"></a>&#34;\t\t\t88888888\n&#34;
            <a id="L393"></a>&#34;\n&#34;
            <a id="L394"></a>&#34;666666\t666666\t666666\t4444\n&#34;
            <a id="L395"></a>&#34;1\t1\t999999999\t0000000000\n&#34;,

        <a id="L397"></a>&#34;4444------日本語-22--1---333\n&#34;
            <a id="L398"></a>&#34;999999999-22\n&#34;
            <a id="L399"></a>&#34;7---------22\n&#34;
            <a id="L400"></a>&#34;------------------88888888\n&#34;
            <a id="L401"></a>&#34;\n&#34;
            <a id="L402"></a>&#34;666666-666666-666666----4444\n&#34;
            <a id="L403"></a>&#34;1------1------999999999-0000000000\n&#34;,
    <a id="L404"></a>},

    <a id="L406"></a>entry{
        <a id="L407"></a>&#34;13b&#34;,
        <a id="L408"></a>4, 3, &#39;.&#39;, 0,
        <a id="L409"></a>&#34;4444\t333\t22\t1\t333\n&#34;
            <a id="L410"></a>&#34;999999999\t22\n&#34;
            <a id="L411"></a>&#34;7\t22\n&#34;
            <a id="L412"></a>&#34;\t\t\t88888888\n&#34;
            <a id="L413"></a>&#34;\n&#34;
            <a id="L414"></a>&#34;666666\t666666\t666666\t4444\n&#34;
            <a id="L415"></a>&#34;1\t1\t999999999\t0000000000\n&#34;,

        <a id="L417"></a>&#34;4444........333...22...1...333\n&#34;
            <a id="L418"></a>&#34;999999999...22\n&#34;
            <a id="L419"></a>&#34;7...........22\n&#34;
            <a id="L420"></a>&#34;....................88888888\n&#34;
            <a id="L421"></a>&#34;\n&#34;
            <a id="L422"></a>&#34;666666...666666...666666......4444\n&#34;
            <a id="L423"></a>&#34;1........1........999999999...0000000000\n&#34;,
    <a id="L424"></a>},

    <a id="L426"></a>entry{
        <a id="L427"></a>&#34;13c&#34;,
        <a id="L428"></a>8, 1, &#39;\t&#39;, FilterHTML,
        <a id="L429"></a>&#34;4444\t333\t22\t1\t333\n&#34;
            <a id="L430"></a>&#34;999999999\t22\n&#34;
            <a id="L431"></a>&#34;7\t22\n&#34;
            <a id="L432"></a>&#34;\t\t\t88888888\n&#34;
            <a id="L433"></a>&#34;\n&#34;
            <a id="L434"></a>&#34;666666\t666666\t666666\t4444\n&#34;
            <a id="L435"></a>&#34;1\t1\t&lt;font color=red attr=日本語&gt;999999999&lt;/font&gt;\t0000000000\n&#34;,

        <a id="L437"></a>&#34;4444\t\t333\t22\t1\t333\n&#34;
            <a id="L438"></a>&#34;999999999\t22\n&#34;
            <a id="L439"></a>&#34;7\t\t22\n&#34;
            <a id="L440"></a>&#34;\t\t\t\t88888888\n&#34;
            <a id="L441"></a>&#34;\n&#34;
            <a id="L442"></a>&#34;666666\t666666\t666666\t\t4444\n&#34;
            <a id="L443"></a>&#34;1\t1\t&lt;font color=red attr=日本語&gt;999999999&lt;/font&gt;\t0000000000\n&#34;,
    <a id="L444"></a>},

    <a id="L446"></a>entry{
        <a id="L447"></a>&#34;14&#34;,
        <a id="L448"></a>1, 2, &#39; &#39;, AlignRight,
        <a id="L449"></a>&#34;.0\t.3\t2.4\t-5.1\t\n&#34;
            <a id="L450"></a>&#34;23.0\t12345678.9\t2.4\t-989.4\t\n&#34;
            <a id="L451"></a>&#34;5.1\t12.0\t2.4\t-7.0\t\n&#34;
            <a id="L452"></a>&#34;.0\t0.0\t332.0\t8908.0\t\n&#34;
            <a id="L453"></a>&#34;.0\t-.3\t456.4\t22.1\t\n&#34;
            <a id="L454"></a>&#34;.0\t1.2\t44.4\t-13.3\t\t&#34;,

        <a id="L456"></a>&#34;    .0          .3    2.4    -5.1\n&#34;
            <a id="L457"></a>&#34;  23.0  12345678.9    2.4  -989.4\n&#34;
            <a id="L458"></a>&#34;   5.1        12.0    2.4    -7.0\n&#34;
            <a id="L459"></a>&#34;    .0         0.0  332.0  8908.0\n&#34;
            <a id="L460"></a>&#34;    .0         -.3  456.4    22.1\n&#34;
            <a id="L461"></a>&#34;    .0         1.2   44.4   -13.3&#34;,
    <a id="L462"></a>},

    <a id="L464"></a>entry{
        <a id="L465"></a>&#34;14 debug&#34;,
        <a id="L466"></a>1, 2, &#39; &#39;, AlignRight | Debug,
        <a id="L467"></a>&#34;.0\t.3\t2.4\t-5.1\t\n&#34;
            <a id="L468"></a>&#34;23.0\t12345678.9\t2.4\t-989.4\t\n&#34;
            <a id="L469"></a>&#34;5.1\t12.0\t2.4\t-7.0\t\n&#34;
            <a id="L470"></a>&#34;.0\t0.0\t332.0\t8908.0\t\n&#34;
            <a id="L471"></a>&#34;.0\t-.3\t456.4\t22.1\t\n&#34;
            <a id="L472"></a>&#34;.0\t1.2\t44.4\t-13.3\t\t&#34;,

        <a id="L474"></a>&#34;    .0|          .3|    2.4|    -5.1|\n&#34;
            <a id="L475"></a>&#34;  23.0|  12345678.9|    2.4|  -989.4|\n&#34;
            <a id="L476"></a>&#34;   5.1|        12.0|    2.4|    -7.0|\n&#34;
            <a id="L477"></a>&#34;    .0|         0.0|  332.0|  8908.0|\n&#34;
            <a id="L478"></a>&#34;    .0|         -.3|  456.4|    22.1|\n&#34;
            <a id="L479"></a>&#34;    .0|         1.2|   44.4|   -13.3|&#34;,
    <a id="L480"></a>},

    <a id="L482"></a>entry{
        <a id="L483"></a>&#34;15a&#34;,
        <a id="L484"></a>4, 0, &#39;.&#39;, 0,
        <a id="L485"></a>&#34;a\t\tb&#34;,
        <a id="L486"></a>&#34;a.......b&#34;,
    <a id="L487"></a>},

    <a id="L489"></a>entry{
        <a id="L490"></a>&#34;15b&#34;,
        <a id="L491"></a>4, 0, &#39;.&#39;, DiscardEmptyColumns,
        <a id="L492"></a>&#34;a\t\tb&#34;, <span class="comment">// htabs - do not discard column</span>
        <a id="L493"></a>&#34;a.......b&#34;,
    <a id="L494"></a>},

    <a id="L496"></a>entry{
        <a id="L497"></a>&#34;15c&#34;,
        <a id="L498"></a>4, 0, &#39;.&#39;, DiscardEmptyColumns,
        <a id="L499"></a>&#34;a\v\vb&#34;,
        <a id="L500"></a>&#34;a...b&#34;,
    <a id="L501"></a>},

    <a id="L503"></a>entry{
        <a id="L504"></a>&#34;15d&#34;,
        <a id="L505"></a>4, 0, &#39;.&#39;, AlignRight | DiscardEmptyColumns,
        <a id="L506"></a>&#34;a\v\vb&#34;,
        <a id="L507"></a>&#34;...ab&#34;,
    <a id="L508"></a>},

    <a id="L510"></a>entry{
        <a id="L511"></a>&#34;16a&#34;,
        <a id="L512"></a>100, 0, &#39;\t&#39;, 0,
        <a id="L513"></a>&#34;a\tb\t\td\n&#34;
            <a id="L514"></a>&#34;a\tb\t\td\te\n&#34;
            <a id="L515"></a>&#34;a\n&#34;
            <a id="L516"></a>&#34;a\tb\tc\td\n&#34;
            <a id="L517"></a>&#34;a\tb\tc\td\te\n&#34;,

        <a id="L519"></a>&#34;a\tb\t\td\n&#34;
            <a id="L520"></a>&#34;a\tb\t\td\te\n&#34;
            <a id="L521"></a>&#34;a\n&#34;
            <a id="L522"></a>&#34;a\tb\tc\td\n&#34;
            <a id="L523"></a>&#34;a\tb\tc\td\te\n&#34;,
    <a id="L524"></a>},

    <a id="L526"></a>entry{
        <a id="L527"></a>&#34;16b&#34;,
        <a id="L528"></a>100, 0, &#39;\t&#39;, DiscardEmptyColumns,
        <a id="L529"></a>&#34;a\vb\v\vd\n&#34;
            <a id="L530"></a>&#34;a\vb\v\vd\ve\n&#34;
            <a id="L531"></a>&#34;a\n&#34;
            <a id="L532"></a>&#34;a\vb\vc\vd\n&#34;
            <a id="L533"></a>&#34;a\vb\vc\vd\ve\n&#34;,

        <a id="L535"></a>&#34;a\tb\td\n&#34;
            <a id="L536"></a>&#34;a\tb\td\te\n&#34;
            <a id="L537"></a>&#34;a\n&#34;
            <a id="L538"></a>&#34;a\tb\tc\td\n&#34;
            <a id="L539"></a>&#34;a\tb\tc\td\te\n&#34;,
    <a id="L540"></a>},

    <a id="L542"></a>entry{
        <a id="L543"></a>&#34;16b debug&#34;,
        <a id="L544"></a>100, 0, &#39;\t&#39;, DiscardEmptyColumns | Debug,
        <a id="L545"></a>&#34;a\vb\v\vd\n&#34;
            <a id="L546"></a>&#34;a\vb\v\vd\ve\n&#34;
            <a id="L547"></a>&#34;a\n&#34;
            <a id="L548"></a>&#34;a\vb\vc\vd\n&#34;
            <a id="L549"></a>&#34;a\vb\vc\vd\ve\n&#34;,

        <a id="L551"></a>&#34;a\t|b\t||d\n&#34;
            <a id="L552"></a>&#34;a\t|b\t||d\t|e\n&#34;
            <a id="L553"></a>&#34;a\n&#34;
            <a id="L554"></a>&#34;a\t|b\t|c\t|d\n&#34;
            <a id="L555"></a>&#34;a\t|b\t|c\t|d\t|e\n&#34;,
    <a id="L556"></a>},

    <a id="L558"></a>entry{
        <a id="L559"></a>&#34;16c&#34;,
        <a id="L560"></a>100, 0, &#39;\t&#39;, DiscardEmptyColumns,
        <a id="L561"></a>&#34;a\tb\t\td\n&#34; <span class="comment">// hard tabs - do not discard column</span>
            <a id="L562"></a>&#34;a\tb\t\td\te\n&#34;
            <a id="L563"></a>&#34;a\n&#34;
            <a id="L564"></a>&#34;a\tb\tc\td\n&#34;
            <a id="L565"></a>&#34;a\tb\tc\td\te\n&#34;,

        <a id="L567"></a>&#34;a\tb\t\td\n&#34;
            <a id="L568"></a>&#34;a\tb\t\td\te\n&#34;
            <a id="L569"></a>&#34;a\n&#34;
            <a id="L570"></a>&#34;a\tb\tc\td\n&#34;
            <a id="L571"></a>&#34;a\tb\tc\td\te\n&#34;,
    <a id="L572"></a>},

    <a id="L574"></a>entry{
        <a id="L575"></a>&#34;16c debug&#34;,
        <a id="L576"></a>100, 0, &#39;\t&#39;, DiscardEmptyColumns | Debug,
        <a id="L577"></a>&#34;a\tb\t\td\n&#34; <span class="comment">// hard tabs - do not discard column</span>
            <a id="L578"></a>&#34;a\tb\t\td\te\n&#34;
            <a id="L579"></a>&#34;a\n&#34;
            <a id="L580"></a>&#34;a\tb\tc\td\n&#34;
            <a id="L581"></a>&#34;a\tb\tc\td\te\n&#34;,

        <a id="L583"></a>&#34;a\t|b\t|\t|d\n&#34;
            <a id="L584"></a>&#34;a\t|b\t|\t|d\t|e\n&#34;
            <a id="L585"></a>&#34;a\n&#34;
            <a id="L586"></a>&#34;a\t|b\t|c\t|d\n&#34;
            <a id="L587"></a>&#34;a\t|b\t|c\t|d\t|e\n&#34;,
    <a id="L588"></a>},
<a id="L589"></a>}


<a id="L592"></a>func Test(t *testing.T) {
    <a id="L593"></a>for _, e := range tests {
        <a id="L594"></a>check(t, e.testname, e.tabwidth, e.padding, e.padchar, e.flags, e.src, e.expected)
    <a id="L595"></a>}
<a id="L596"></a>}
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
