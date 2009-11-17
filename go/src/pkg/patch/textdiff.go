<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/patch/textdiff.go</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/patch/textdiff.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a>package patch

<a id="L3"></a>import (
    <a id="L4"></a>&#34;bytes&#34;;
    <a id="L5"></a>&#34;os&#34;;
<a id="L6"></a>)

<a id="L8"></a>type TextDiff []TextChunk

<a id="L10"></a><span class="comment">// A TextChunk specifies an edit to a section of a file:</span>
<a id="L11"></a><span class="comment">// the text beginning at Line, which should be exactly Old,</span>
<a id="L12"></a><span class="comment">// is to be replaced with New.</span>
<a id="L13"></a>type TextChunk struct {
    <a id="L14"></a>Line int;
    <a id="L15"></a>Old  []byte;
    <a id="L16"></a>New  []byte;
<a id="L17"></a>}

<a id="L19"></a>func ParseTextDiff(raw []byte) (TextDiff, os.Error) {
    <a id="L20"></a><span class="comment">// Copy raw so it is safe to keep references to slices.</span>
    <a id="L21"></a>_, chunks := sections(raw, &#34;@@ -&#34;);
    <a id="L22"></a>delta := 0;
    <a id="L23"></a>diff := make(TextDiff, len(chunks));
    <a id="L24"></a>for i, raw := range chunks {
        <a id="L25"></a>c := &amp;diff[i];

        <a id="L27"></a><span class="comment">// Parse start line: @@ -oldLine,oldCount +newLine,newCount @@ junk</span>
        <a id="L28"></a>chunk := splitLines(raw);
        <a id="L29"></a>chunkHeader := chunk[0];
        <a id="L30"></a>var ok bool;
        <a id="L31"></a>var oldLine, oldCount, newLine, newCount int;
        <a id="L32"></a>s := chunkHeader;
        <a id="L33"></a>if oldLine, s, ok = atoi(s, &#34;@@ -&#34;, 10); !ok {
        <a id="L34"></a>ErrChunkHdr:
            <a id="L35"></a>return nil, SyntaxError(&#34;unexpected chunk header line: &#34; + string(chunkHeader))
        <a id="L36"></a>}
        <a id="L37"></a>if len(s) == 0 || s[0] != &#39;,&#39; {
            <a id="L38"></a>oldCount = 1
        <a id="L39"></a>} else if oldCount, s, ok = atoi(s, &#34;,&#34;, 10); !ok {
            <a id="L40"></a>goto ErrChunkHdr
        <a id="L41"></a>}
        <a id="L42"></a>if newLine, s, ok = atoi(s, &#34; +&#34;, 10); !ok {
            <a id="L43"></a>goto ErrChunkHdr
        <a id="L44"></a>}
        <a id="L45"></a>if len(s) == 0 || s[0] != &#39;,&#39; {
            <a id="L46"></a>newCount = 1
        <a id="L47"></a>} else if newCount, s, ok = atoi(s, &#34;,&#34;, 10); !ok {
            <a id="L48"></a>goto ErrChunkHdr
        <a id="L49"></a>}
        <a id="L50"></a>if !hasPrefix(s, &#34; @@&#34;) {
            <a id="L51"></a>goto ErrChunkHdr
        <a id="L52"></a>}

        <a id="L54"></a><span class="comment">// Special case: for created or deleted files, the empty half</span>
        <a id="L55"></a><span class="comment">// is given as starting at line 0.  Translate to line 1.</span>
        <a id="L56"></a>if oldCount == 0 &amp;&amp; oldLine == 0 {
            <a id="L57"></a>oldLine = 1
        <a id="L58"></a>}
        <a id="L59"></a>if newCount == 0 &amp;&amp; newLine == 0 {
            <a id="L60"></a>newLine = 1
        <a id="L61"></a>}

        <a id="L63"></a><span class="comment">// Count lines in text</span>
        <a id="L64"></a>var dropOldNL, dropNewNL bool;
        <a id="L65"></a>var nold, nnew int;
        <a id="L66"></a>var lastch byte;
        <a id="L67"></a>chunk = chunk[1:len(chunk)];
        <a id="L68"></a>for _, l := range chunk {
            <a id="L69"></a>if nold == oldCount &amp;&amp; nnew == newCount &amp;&amp; (len(l) == 0 || l[0] != &#39;\\&#39;) {
                <a id="L70"></a>if len(bytes.TrimSpace(l)) != 0 {
                    <a id="L71"></a>return nil, SyntaxError(&#34;too many chunk lines&#34;)
                <a id="L72"></a>}
                <a id="L73"></a>continue;
            <a id="L74"></a>}
            <a id="L75"></a>if len(l) == 0 {
                <a id="L76"></a>return nil, SyntaxError(&#34;empty chunk line&#34;)
            <a id="L77"></a>}
            <a id="L78"></a>switch l[0] {
            <a id="L79"></a>case &#39;+&#39;:
                <a id="L80"></a>nnew++
            <a id="L81"></a>case &#39;-&#39;:
                <a id="L82"></a>nold++
            <a id="L83"></a>case &#39; &#39;:
                <a id="L84"></a>nnew++;
                <a id="L85"></a>nold++;
            <a id="L86"></a>case &#39;\\&#39;:
                <a id="L87"></a>if _, ok := skip(l, &#34;\\ No newline at end of file&#34;); ok {
                    <a id="L88"></a>switch lastch {
                    <a id="L89"></a>case &#39;-&#39;:
                        <a id="L90"></a>dropOldNL = true
                    <a id="L91"></a>case &#39;+&#39;:
                        <a id="L92"></a>dropNewNL = true
                    <a id="L93"></a>case &#39; &#39;:
                        <a id="L94"></a>dropOldNL = true;
                        <a id="L95"></a>dropNewNL = true;
                    <a id="L96"></a>default:
                        <a id="L97"></a>return nil, SyntaxError(&#34;message `\\ No newline at end of file&#39; out of context&#34;)
                    <a id="L98"></a>}
                    <a id="L99"></a>break;
                <a id="L100"></a>}
                <a id="L101"></a>fallthrough;
            <a id="L102"></a>default:
                <a id="L103"></a>return nil, SyntaxError(&#34;unexpected chunk line: &#34; + string(l))
            <a id="L104"></a>}
            <a id="L105"></a>lastch = l[0];
        <a id="L106"></a>}

        <a id="L108"></a><span class="comment">// Does it match the header?</span>
        <a id="L109"></a>if nold != oldCount || nnew != newCount {
            <a id="L110"></a>return nil, SyntaxError(&#34;chunk header does not match line count: &#34; + string(chunkHeader))
        <a id="L111"></a>}
        <a id="L112"></a>if oldLine+delta != newLine {
            <a id="L113"></a>return nil, SyntaxError(&#34;chunk delta is out of sync with previous chunks&#34;)
        <a id="L114"></a>}
        <a id="L115"></a>delta += nnew - nold;
        <a id="L116"></a>c.Line = oldLine;

        <a id="L118"></a>var old, new bytes.Buffer;
        <a id="L119"></a>nold = 0;
        <a id="L120"></a>nnew = 0;
        <a id="L121"></a>for _, l := range chunk {
            <a id="L122"></a>if nold == oldCount &amp;&amp; nnew == newCount {
                <a id="L123"></a>break
            <a id="L124"></a>}
            <a id="L125"></a>ch, l := l[0], l[1:len(l)];
            <a id="L126"></a>if ch == &#39;\\&#39; {
                <a id="L127"></a>continue
            <a id="L128"></a>}
            <a id="L129"></a>if ch != &#39;+&#39; {
                <a id="L130"></a>old.Write(l);
                <a id="L131"></a>nold++;
            <a id="L132"></a>}
            <a id="L133"></a>if ch != &#39;-&#39; {
                <a id="L134"></a>new.Write(l);
                <a id="L135"></a>nnew++;
            <a id="L136"></a>}
        <a id="L137"></a>}
        <a id="L138"></a>c.Old = old.Bytes();
        <a id="L139"></a>c.New = new.Bytes();
        <a id="L140"></a>if dropOldNL {
            <a id="L141"></a>c.Old = c.Old[0 : len(c.Old)-1]
        <a id="L142"></a>}
        <a id="L143"></a>if dropNewNL {
            <a id="L144"></a>c.New = c.New[0 : len(c.New)-1]
        <a id="L145"></a>}
    <a id="L146"></a>}
    <a id="L147"></a>return diff, nil;
<a id="L148"></a>}

<a id="L150"></a>var ErrPatchFailure = os.NewError(&#34;patch did not apply cleanly&#34;)

<a id="L152"></a><span class="comment">// Apply applies the changes listed in the diff</span>
<a id="L153"></a><span class="comment">// to the data, returning the new version.</span>
<a id="L154"></a>func (d TextDiff) Apply(data []byte) ([]byte, os.Error) {
    <a id="L155"></a>var buf bytes.Buffer;
    <a id="L156"></a>line := 1;
    <a id="L157"></a>for _, c := range d {
        <a id="L158"></a>var ok bool;
        <a id="L159"></a>var prefix []byte;
        <a id="L160"></a>prefix, data, ok = getLine(data, c.Line-line);
        <a id="L161"></a>if !ok || !bytes.HasPrefix(data, c.Old) {
            <a id="L162"></a>return nil, ErrPatchFailure
        <a id="L163"></a>}
        <a id="L164"></a>buf.Write(prefix);
        <a id="L165"></a>data = data[len(c.Old):len(data)];
        <a id="L166"></a>buf.Write(c.New);
        <a id="L167"></a>line = c.Line + bytes.Count(c.Old, newline);
    <a id="L168"></a>}
    <a id="L169"></a>buf.Write(data);
    <a id="L170"></a>return buf.Bytes(), nil;
<a id="L171"></a>}
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
