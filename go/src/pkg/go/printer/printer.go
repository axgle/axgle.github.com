<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/printer/printer.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/go/printer/printer.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The printer package implements printing of AST nodes.</span>
<a id="L6"></a>package printer

<a id="L8"></a>import (
    <a id="L9"></a>&#34;bytes&#34;;
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;go/ast&#34;;
    <a id="L12"></a>&#34;go/token&#34;;
    <a id="L13"></a>&#34;io&#34;;
    <a id="L14"></a>&#34;os&#34;;
    <a id="L15"></a>&#34;reflect&#34;;
    <a id="L16"></a>&#34;runtime&#34;;
    <a id="L17"></a>&#34;strings&#34;;
    <a id="L18"></a>&#34;tabwriter&#34;;
<a id="L19"></a>)


<a id="L22"></a>const (
    <a id="L23"></a>debug       = false; <span class="comment">// enable for debugging</span>
    <a id="L24"></a>maxNewlines = 3;     <span class="comment">// maximum vertical white space</span>
<a id="L25"></a>)


<a id="L28"></a>type whiteSpace int

<a id="L30"></a>const (
    <a id="L31"></a>ignore   = whiteSpace(0);
    <a id="L32"></a>blank    = whiteSpace(&#39; &#39;);
    <a id="L33"></a>vtab     = whiteSpace(&#39;\v&#39;);
    <a id="L34"></a>newline  = whiteSpace(&#39;\n&#39;);
    <a id="L35"></a>formfeed = whiteSpace(&#39;\f&#39;);
    <a id="L36"></a>indent   = whiteSpace(&#39;&gt;&#39;);
    <a id="L37"></a>unindent = whiteSpace(&#39;&lt;&#39;);
<a id="L38"></a>)


<a id="L41"></a>var (
    <a id="L42"></a>esc       = []byte{tabwriter.Escape};
    <a id="L43"></a>htab      = []byte{&#39;\t&#39;};
    <a id="L44"></a>htabs     = [...]byte{&#39;\t&#39;, &#39;\t&#39;, &#39;\t&#39;, &#39;\t&#39;, &#39;\t&#39;, &#39;\t&#39;, &#39;\t&#39;, &#39;\t&#39;};
    <a id="L45"></a>newlines  = [...]byte{&#39;\n&#39;, &#39;\n&#39;, &#39;\n&#39;, &#39;\n&#39;, &#39;\n&#39;, &#39;\n&#39;, &#39;\n&#39;, &#39;\n&#39;}; <span class="comment">// more than maxNewlines</span>
    <a id="L46"></a>formfeeds = [...]byte{&#39;\f&#39;, &#39;\f&#39;, &#39;\f&#39;, &#39;\f&#39;, &#39;\f&#39;, &#39;\f&#39;, &#39;\f&#39;, &#39;\f&#39;}; <span class="comment">// more than maxNewlines</span>

    <a id="L48"></a>esc_quot = strings.Bytes(&#34;&amp;#34;&#34;); <span class="comment">// shorter than &#34;&amp;quot;&#34;</span>
    <a id="L49"></a>esc_apos = strings.Bytes(&#34;&amp;#39;&#34;); <span class="comment">// shorter than &#34;&amp;apos;&#34;</span>
    <a id="L50"></a>esc_amp  = strings.Bytes(&#34;&amp;amp;&#34;);
    <a id="L51"></a>esc_lt   = strings.Bytes(&#34;&amp;lt;&#34;);
    <a id="L52"></a>esc_gt   = strings.Bytes(&#34;&amp;gt;&#34;);
<a id="L53"></a>)


<a id="L56"></a><span class="comment">// Use noPos when a position is needed but not known.</span>
<a id="L57"></a>var noPos token.Position


<a id="L60"></a><span class="comment">// Use ignoreMultiLine if the multiLine information is not important.</span>
<a id="L61"></a>var ignoreMultiLine = new(bool)


<a id="L64"></a>type printer struct {
    <a id="L65"></a><span class="comment">// Configuration (does not change after initialization)</span>
    <a id="L66"></a>output io.Writer;
    <a id="L67"></a>Config;
    <a id="L68"></a>errors chan os.Error;

    <a id="L70"></a><span class="comment">// Current state</span>
    <a id="L71"></a>written int;  <span class="comment">// number of bytes written</span>
    <a id="L72"></a>indent  int;  <span class="comment">// current indentation</span>
    <a id="L73"></a>escape  bool; <span class="comment">// true if in escape sequence</span>

    <a id="L75"></a><span class="comment">// Buffered whitespace</span>
    <a id="L76"></a>buffer []whiteSpace;

    <a id="L78"></a><span class="comment">// The (possibly estimated) position in the generated output;</span>
    <a id="L79"></a><span class="comment">// in AST space (i.e., pos is set whenever a token position is</span>
    <a id="L80"></a><span class="comment">// known accurately, and updated dependending on what has been</span>
    <a id="L81"></a><span class="comment">// written)</span>
    <a id="L82"></a>pos token.Position;

    <a id="L84"></a><span class="comment">// The value of pos immediately after the last item has been</span>
    <a id="L85"></a><span class="comment">// written using writeItem.</span>
    <a id="L86"></a>last token.Position;

    <a id="L88"></a><span class="comment">// HTML support</span>
    <a id="L89"></a>lastTaggedLine int; <span class="comment">// last line for which a line tag was written</span>

    <a id="L91"></a><span class="comment">// The list of comments; or nil.</span>
    <a id="L92"></a>comment *ast.CommentGroup;
<a id="L93"></a>}


<a id="L96"></a>func (p *printer) init(output io.Writer, cfg *Config) {
    <a id="L97"></a>p.output = output;
    <a id="L98"></a>p.Config = *cfg;
    <a id="L99"></a>p.errors = make(chan os.Error);
    <a id="L100"></a>p.buffer = make([]whiteSpace, 0, 16); <span class="comment">// whitespace sequences are short</span>
<a id="L101"></a>}


<a id="L104"></a>func (p *printer) internalError(msg ...) {
    <a id="L105"></a>if debug {
        <a id="L106"></a>fmt.Print(p.pos.String() + &#34;: &#34;);
        <a id="L107"></a>fmt.Println(msg);
        <a id="L108"></a>panic();
    <a id="L109"></a>}
<a id="L110"></a>}


<a id="L113"></a><span class="comment">// write0 writes raw (uninterpreted) data to p.output and handles errors.</span>
<a id="L114"></a><span class="comment">// write0 does not indent after newlines, and does not HTML-escape or update p.pos.</span>
<a id="L115"></a><span class="comment">//</span>
<a id="L116"></a>func (p *printer) write0(data []byte) {
    <a id="L117"></a>n, err := p.output.Write(data);
    <a id="L118"></a>p.written += n;
    <a id="L119"></a>if err != nil {
        <a id="L120"></a>p.errors &lt;- err;
        <a id="L121"></a>runtime.Goexit();
    <a id="L122"></a>}
<a id="L123"></a>}


<a id="L126"></a><span class="comment">// write interprets data and writes it to p.output. It inserts indentation</span>
<a id="L127"></a><span class="comment">// after a line break unless in a tabwriter escape sequence, and it HTML-</span>
<a id="L128"></a><span class="comment">// escapes characters if GenHTML is set. It updates p.pos as a side-effect.</span>
<a id="L129"></a><span class="comment">//</span>
<a id="L130"></a>func (p *printer) write(data []byte) {
    <a id="L131"></a>i0 := 0;
    <a id="L132"></a>for i, b := range data {
        <a id="L133"></a>switch b {
        <a id="L134"></a>case &#39;\n&#39;, &#39;\f&#39;:
            <a id="L135"></a><span class="comment">// write segment ending in b</span>
            <a id="L136"></a>p.write0(data[i0 : i+1]);

            <a id="L138"></a><span class="comment">// update p.pos</span>
            <a id="L139"></a>p.pos.Offset += i + 1 - i0;
            <a id="L140"></a>p.pos.Line++;
            <a id="L141"></a>p.pos.Column = 1;

            <a id="L143"></a>if !p.escape {
                <a id="L144"></a><span class="comment">// write indentation</span>
                <a id="L145"></a><span class="comment">// use &#34;hard&#34; htabs - indentation columns</span>
                <a id="L146"></a><span class="comment">// must not be discarded by the tabwriter</span>
                <a id="L147"></a>j := p.indent;
                <a id="L148"></a>for ; j &gt; len(htabs); j -= len(htabs) {
                    <a id="L149"></a>p.write0(&amp;htabs)
                <a id="L150"></a>}
                <a id="L151"></a>p.write0(htabs[0:j]);

                <a id="L153"></a><span class="comment">// update p.pos</span>
                <a id="L154"></a>p.pos.Offset += p.indent;
                <a id="L155"></a>p.pos.Column += p.indent;
            <a id="L156"></a>}

            <a id="L158"></a><span class="comment">// next segment start</span>
            <a id="L159"></a>i0 = i + 1;

        <a id="L161"></a>case &#39;&#34;&#39;, &#39;\&#39;&#39;, &#39;&amp;&#39;, &#39;&lt;&#39;, &#39;&gt;&#39;:
            <a id="L162"></a>if p.Mode&amp;GenHTML != 0 {
                <a id="L163"></a><span class="comment">// write segment ending in b</span>
                <a id="L164"></a>p.write0(data[i0:i]);

                <a id="L166"></a><span class="comment">// write HTML-escaped b</span>
                <a id="L167"></a>var esc []byte;
                <a id="L168"></a>switch b {
                <a id="L169"></a>case &#39;&#34;&#39;:
                    <a id="L170"></a>esc = esc_quot
                <a id="L171"></a>case &#39;\&#39;&#39;:
                    <a id="L172"></a>esc = esc_apos
                <a id="L173"></a>case &#39;&amp;&#39;:
                    <a id="L174"></a>esc = esc_amp
                <a id="L175"></a>case &#39;&lt;&#39;:
                    <a id="L176"></a>esc = esc_lt
                <a id="L177"></a>case &#39;&gt;&#39;:
                    <a id="L178"></a>esc = esc_gt
                <a id="L179"></a>}
                <a id="L180"></a>p.write0(esc);

                <a id="L182"></a><span class="comment">// update p.pos</span>
                <a id="L183"></a>d := i + 1 - i0;
                <a id="L184"></a>p.pos.Offset += d;
                <a id="L185"></a>p.pos.Column += d;

                <a id="L187"></a><span class="comment">// next segment start</span>
                <a id="L188"></a>i0 = i + 1;
            <a id="L189"></a>}

        <a id="L191"></a>case tabwriter.Escape:
            <a id="L192"></a>p.escape = !p.escape
        <a id="L193"></a>}
    <a id="L194"></a>}

    <a id="L196"></a><span class="comment">// write remaining segment</span>
    <a id="L197"></a>p.write0(data[i0:len(data)]);

    <a id="L199"></a><span class="comment">// update p.pos</span>
    <a id="L200"></a>d := len(data) - i0;
    <a id="L201"></a>p.pos.Offset += d;
    <a id="L202"></a>p.pos.Column += d;
<a id="L203"></a>}


<a id="L206"></a>func (p *printer) writeNewlines(n int) {
    <a id="L207"></a>if n &gt; 0 {
        <a id="L208"></a>if n &gt; maxNewlines {
            <a id="L209"></a>n = maxNewlines
        <a id="L210"></a>}
        <a id="L211"></a>p.write(newlines[0:n]);
    <a id="L212"></a>}
<a id="L213"></a>}


<a id="L216"></a>func (p *printer) writeFormfeeds(n int) {
    <a id="L217"></a>if n &gt; 0 {
        <a id="L218"></a>if n &gt; maxNewlines {
            <a id="L219"></a>n = maxNewlines
        <a id="L220"></a>}
        <a id="L221"></a>p.write(formfeeds[0:n]);
    <a id="L222"></a>}
<a id="L223"></a>}


<a id="L226"></a>func (p *printer) writeTaggedItem(data []byte, tag HTMLTag) {
    <a id="L227"></a><span class="comment">// write start tag, if any</span>
    <a id="L228"></a><span class="comment">// (no html-escaping and no p.pos update for tags - use write0)</span>
    <a id="L229"></a>if tag.Start != &#34;&#34; {
        <a id="L230"></a>p.write0(strings.Bytes(tag.Start))
    <a id="L231"></a>}
    <a id="L232"></a>p.write(data);
    <a id="L233"></a><span class="comment">// write end tag, if any</span>
    <a id="L234"></a>if tag.End != &#34;&#34; {
        <a id="L235"></a>p.write0(strings.Bytes(tag.End))
    <a id="L236"></a>}
<a id="L237"></a>}


<a id="L240"></a><span class="comment">// writeItem writes data at position pos. data is the text corresponding to</span>
<a id="L241"></a><span class="comment">// a single lexical token, but may also be comment text. pos is the actual</span>
<a id="L242"></a><span class="comment">// (or at least very accurately estimated) position of the data in the original</span>
<a id="L243"></a><span class="comment">// source text. If tags are present and GenHTML is set, the tags are written</span>
<a id="L244"></a><span class="comment">// before and after the data. writeItem updates p.last to the position</span>
<a id="L245"></a><span class="comment">// immediately following the data.</span>
<a id="L246"></a><span class="comment">//</span>
<a id="L247"></a>func (p *printer) writeItem(pos token.Position, data []byte, tag HTMLTag) {
    <a id="L248"></a>p.pos = pos;
    <a id="L249"></a>if debug {
        <a id="L250"></a><span class="comment">// do not update p.pos - use write0</span>
        <a id="L251"></a>p.write0(strings.Bytes(fmt.Sprintf(&#34;[%d:%d]&#34;, pos.Line, pos.Column)))
    <a id="L252"></a>}
    <a id="L253"></a>if p.Mode&amp;GenHTML != 0 {
        <a id="L254"></a><span class="comment">// write line tag if on a new line</span>
        <a id="L255"></a><span class="comment">// TODO(gri): should write line tags on each line at the start</span>
        <a id="L256"></a><span class="comment">//            will be more useful (e.g. to show line numbers)</span>
        <a id="L257"></a>if p.Styler != nil &amp;&amp; pos.Line &gt; p.lastTaggedLine {
            <a id="L258"></a>p.writeTaggedItem(p.Styler.LineTag(pos.Line));
            <a id="L259"></a>p.lastTaggedLine = pos.Line;
        <a id="L260"></a>}
        <a id="L261"></a>p.writeTaggedItem(data, tag);
    <a id="L262"></a>} else {
        <a id="L263"></a>p.write(data)
    <a id="L264"></a>}
    <a id="L265"></a>p.last = p.pos;
<a id="L266"></a>}


<a id="L269"></a><span class="comment">// writeCommentPrefix writes the whitespace before a comment.</span>
<a id="L270"></a><span class="comment">// If there is any pending whitespace, it consumes as much of</span>
<a id="L271"></a><span class="comment">// it as is likely to help the comment position properly.</span>
<a id="L272"></a><span class="comment">// pos is the comment position, next the position of the item</span>
<a id="L273"></a><span class="comment">// after all pending comments, isFirst indicates if this is the</span>
<a id="L274"></a><span class="comment">// first comment in a group of comments, and isKeyword indicates</span>
<a id="L275"></a><span class="comment">// if the next item is a keyword.</span>
<a id="L276"></a><span class="comment">//</span>
<a id="L277"></a>func (p *printer) writeCommentPrefix(pos, next token.Position, isFirst, isKeyword bool) {
    <a id="L278"></a>if !p.last.IsValid() {
        <a id="L279"></a><span class="comment">// there was no preceeding item and the comment is the</span>
        <a id="L280"></a><span class="comment">// first item to be printed - don&#39;t write any whitespace</span>
        <a id="L281"></a>return
    <a id="L282"></a>}

    <a id="L284"></a>if pos.Line == p.last.Line {
        <a id="L285"></a><span class="comment">// comment on the same line as last item:</span>
        <a id="L286"></a><span class="comment">// separate with at least one separator</span>
        <a id="L287"></a>hasSep := false;
        <a id="L288"></a>if isFirst {
            <a id="L289"></a>j := 0;
            <a id="L290"></a>for i, ch := range p.buffer {
                <a id="L291"></a>switch ch {
                <a id="L292"></a>case blank:
                    <a id="L293"></a><span class="comment">// ignore any blanks before a comment</span>
                    <a id="L294"></a>p.buffer[i] = ignore;
                    <a id="L295"></a>continue;
                <a id="L296"></a>case vtab:
                    <a id="L297"></a><span class="comment">// respect existing tabs - important</span>
                    <a id="L298"></a><span class="comment">// for proper formatting of commented structs</span>
                    <a id="L299"></a>hasSep = true;
                    <a id="L300"></a>continue;
                <a id="L301"></a>case indent:
                    <a id="L302"></a><span class="comment">// apply pending indentation</span>
                    <a id="L303"></a>continue
                <a id="L304"></a>}
                <a id="L305"></a>j = i;
                <a id="L306"></a>break;
            <a id="L307"></a>}
            <a id="L308"></a>p.writeWhitespace(j);
        <a id="L309"></a>}
        <a id="L310"></a><span class="comment">// make sure there is at least one separator</span>
        <a id="L311"></a>if !hasSep {
            <a id="L312"></a>if pos.Line == next.Line {
                <a id="L313"></a><span class="comment">// next item is on the same line as the comment</span>
                <a id="L314"></a><span class="comment">// (which must be a /*-style comment): separate</span>
                <a id="L315"></a><span class="comment">// with a blank instead of a tab</span>
                <a id="L316"></a>p.write([]byte{&#39; &#39;})
            <a id="L317"></a>} else {
                <a id="L318"></a>p.write(htab)
            <a id="L319"></a>}
        <a id="L320"></a>}

    <a id="L322"></a>} else {
        <a id="L323"></a><span class="comment">// comment on a different line:</span>
        <a id="L324"></a><span class="comment">// separate with at least one line break</span>
        <a id="L325"></a>if isFirst {
            <a id="L326"></a>j := 0;
            <a id="L327"></a>for i, ch := range p.buffer {
                <a id="L328"></a>switch ch {
                <a id="L329"></a>case blank, vtab:
                    <a id="L330"></a><span class="comment">// ignore any horizontal whitespace before line breaks</span>
                    <a id="L331"></a>p.buffer[i] = ignore;
                    <a id="L332"></a>continue;
                <a id="L333"></a>case indent:
                    <a id="L334"></a><span class="comment">// apply pending indentation</span>
                    <a id="L335"></a>continue
                <a id="L336"></a>case unindent:
                    <a id="L337"></a><span class="comment">// if the next token is a keyword, apply the outdent</span>
                    <a id="L338"></a><span class="comment">// if it appears that the comment is aligned with the</span>
                    <a id="L339"></a><span class="comment">// keyword; otherwise assume the outdent is part of a</span>
                    <a id="L340"></a><span class="comment">// closing block and stop (this scenario appears with</span>
                    <a id="L341"></a><span class="comment">// comments before a case label where the comments</span>
                    <a id="L342"></a><span class="comment">// apply to the next case instead of the current one)</span>
                    <a id="L343"></a>if isKeyword &amp;&amp; pos.Column == next.Column {
                        <a id="L344"></a>continue
                    <a id="L345"></a>}
                <a id="L346"></a>case newline, formfeed:
                    <a id="L347"></a><span class="comment">// TODO(gri): may want to keep formfeed info in some cases</span>
                    <a id="L348"></a>p.buffer[i] = ignore
                <a id="L349"></a>}
                <a id="L350"></a>j = i;
                <a id="L351"></a>break;
            <a id="L352"></a>}
            <a id="L353"></a>p.writeWhitespace(j);
        <a id="L354"></a>}
        <a id="L355"></a><span class="comment">// use formfeeds to break columns before a comment;</span>
        <a id="L356"></a><span class="comment">// this is analogous to using formfeeds to separate</span>
        <a id="L357"></a><span class="comment">// individual lines of /*-style comments</span>
        <a id="L358"></a>p.writeFormfeeds(pos.Line - p.last.Line);
    <a id="L359"></a>}
<a id="L360"></a>}


<a id="L363"></a>func (p *printer) writeCommentLine(comment *ast.Comment, pos token.Position, line []byte) {
    <a id="L364"></a><span class="comment">// line must pass through unchanged, bracket it with tabwriter.Escape</span>
    <a id="L365"></a>esc := []byte{tabwriter.Escape};
    <a id="L366"></a>line = bytes.Join([][]byte{esc, line, esc}, nil);

    <a id="L368"></a><span class="comment">// apply styler, if any</span>
    <a id="L369"></a>var tag HTMLTag;
    <a id="L370"></a>if p.Styler != nil {
        <a id="L371"></a>line, tag = p.Styler.Comment(comment, line)
    <a id="L372"></a>}

    <a id="L374"></a>p.writeItem(pos, line, tag);
<a id="L375"></a>}


<a id="L378"></a><span class="comment">// TODO(gri): Similar (but not quite identical) functionality for</span>
<a id="L379"></a><span class="comment">//            comment processing can be found in go/doc/comment.go.</span>
<a id="L380"></a><span class="comment">//            Perhaps this can be factored eventually.</span>

<a id="L382"></a><span class="comment">// Split comment text into lines</span>
<a id="L383"></a>func split(text []byte) [][]byte {
    <a id="L384"></a><span class="comment">// count lines (comment text never ends in a newline)</span>
    <a id="L385"></a>n := 1;
    <a id="L386"></a>for _, c := range text {
        <a id="L387"></a>if c == &#39;\n&#39; {
            <a id="L388"></a>n++
        <a id="L389"></a>}
    <a id="L390"></a>}

    <a id="L392"></a><span class="comment">// split</span>
    <a id="L393"></a>lines := make([][]byte, n);
    <a id="L394"></a>n = 0;
    <a id="L395"></a>i := 0;
    <a id="L396"></a>for j, c := range text {
        <a id="L397"></a>if c == &#39;\n&#39; {
            <a id="L398"></a>lines[n] = text[i:j]; <span class="comment">// exclude newline</span>
            <a id="L399"></a>i = j + 1;            <span class="comment">// discard newline</span>
            <a id="L400"></a>n++;
        <a id="L401"></a>}
    <a id="L402"></a>}
    <a id="L403"></a>lines[n] = text[i:len(text)];

    <a id="L405"></a>return lines;
<a id="L406"></a>}


<a id="L409"></a>func isBlank(s []byte) bool {
    <a id="L410"></a>for _, b := range s {
        <a id="L411"></a>if b &gt; &#39; &#39; {
            <a id="L412"></a>return false
        <a id="L413"></a>}
    <a id="L414"></a>}
    <a id="L415"></a>return true;
<a id="L416"></a>}


<a id="L419"></a>func commonPrefix(a, b []byte) []byte {
    <a id="L420"></a>i := 0;
    <a id="L421"></a>for i &lt; len(a) &amp;&amp; i &lt; len(b) &amp;&amp; a[i] == b[i] &amp;&amp; (a[i] &lt;= &#39; &#39; || a[i] == &#39;*&#39;) {
        <a id="L422"></a>i++
    <a id="L423"></a>}
    <a id="L424"></a>return a[0:i];
<a id="L425"></a>}


<a id="L428"></a>func stripCommonPrefix(lines [][]byte) {
    <a id="L429"></a>if len(lines) &lt; 2 {
        <a id="L430"></a>return <span class="comment">// at most one line - nothing to do</span>
    <a id="L431"></a>}

    <a id="L433"></a><span class="comment">// The heuristic in this function tries to handle a few</span>
    <a id="L434"></a><span class="comment">// common patterns of /*-style comments: Comments where</span>
    <a id="L435"></a><span class="comment">// the opening /* and closing */ are aligned and the</span>
    <a id="L436"></a><span class="comment">// rest of the comment text is aligned and indented with</span>
    <a id="L437"></a><span class="comment">// blanks or tabs, cases with a vertical &#34;line of stars&#34;</span>
    <a id="L438"></a><span class="comment">// on the left, and cases where the closing */ is on the</span>
    <a id="L439"></a><span class="comment">// same line as the last comment text.</span>

    <a id="L441"></a><span class="comment">// Compute maximum common white prefix of all but the first,</span>
    <a id="L442"></a><span class="comment">// last, and blank lines, and replace blank lines with empty</span>
    <a id="L443"></a><span class="comment">// lines (the first line starts with /* and has no prefix).</span>
    <a id="L444"></a>var prefix []byte;
    <a id="L445"></a>for i, line := range lines {
        <a id="L446"></a>switch {
        <a id="L447"></a>case i == 0 || i == len(lines)-1:
            <a id="L448"></a><span class="comment">// ignore</span>
        <a id="L449"></a>case isBlank(line):
            <a id="L450"></a>lines[i] = nil
        <a id="L451"></a>case prefix == nil:
            <a id="L452"></a>prefix = commonPrefix(line, line)
        <a id="L453"></a>default:
            <a id="L454"></a>prefix = commonPrefix(prefix, line)
        <a id="L455"></a>}
    <a id="L456"></a>}

    <a id="L458"></a><span class="comment">/*</span>
    <a id="L459"></a><span class="comment"> * Check for vertical &#34;line of stars&#34; and correct prefix accordingly.</span>
    <a id="L460"></a><span class="comment"> */</span>
    <a id="L461"></a>lineOfStars := false;
    <a id="L462"></a>if i := bytes.Index(prefix, []byte{&#39;*&#39;}); i &gt;= 0 {
        <a id="L463"></a><span class="comment">// Line of stars present.</span>
        <a id="L464"></a>if i &gt; 0 &amp;&amp; prefix[i-1] == &#39; &#39; {
            <a id="L465"></a>i-- <span class="comment">// remove trailing blank from prefix so stars remain aligned</span>
        <a id="L466"></a>}
        <a id="L467"></a>prefix = prefix[0:i];
        <a id="L468"></a>lineOfStars = true;
    <a id="L469"></a>} else {
        <a id="L470"></a><span class="comment">// No line of stars present.</span>
        <a id="L471"></a><span class="comment">// Determine the white space on the first line after the /*</span>
        <a id="L472"></a><span class="comment">// and before the beginning of the comment text, assume two</span>
        <a id="L473"></a><span class="comment">// blanks instead of the /* unless the first character after</span>
        <a id="L474"></a><span class="comment">// the /* is a tab. If the first comment line is empty but</span>
        <a id="L475"></a><span class="comment">// for the opening /*, assume up to 3 blanks or a tab. This</span>
        <a id="L476"></a><span class="comment">// whitespace may be found as suffix in the common prefix.</span>
        <a id="L477"></a>first := lines[0];
        <a id="L478"></a>if isBlank(first[2:len(first)]) {
            <a id="L479"></a><span class="comment">// no comment text on the first line:</span>
            <a id="L480"></a><span class="comment">// reduce prefix by up to 3 blanks or a tab</span>
            <a id="L481"></a><span class="comment">// if present - this keeps comment text indented</span>
            <a id="L482"></a><span class="comment">// relative to the /* and */&#39;s if it was indented</span>
            <a id="L483"></a><span class="comment">// in the first place</span>
            <a id="L484"></a>i := len(prefix);
            <a id="L485"></a>for n := 0; n &lt; 3 &amp;&amp; i &gt; 0 &amp;&amp; prefix[i-1] == &#39; &#39;; n++ {
                <a id="L486"></a>i--
            <a id="L487"></a>}
            <a id="L488"></a>if i == len(prefix) &amp;&amp; i &gt; 0 &amp;&amp; prefix[i-1] == &#39;\t&#39; {
                <a id="L489"></a>i--
            <a id="L490"></a>}
            <a id="L491"></a>prefix = prefix[0:i];
        <a id="L492"></a>} else {
            <a id="L493"></a><span class="comment">// comment text on the first line</span>
            <a id="L494"></a>suffix := make([]byte, len(first));
            <a id="L495"></a>n := 2;
            <a id="L496"></a>for n &lt; len(first) &amp;&amp; first[n] &lt;= &#39; &#39; {
                <a id="L497"></a>suffix[n] = first[n];
                <a id="L498"></a>n++;
            <a id="L499"></a>}
            <a id="L500"></a>if n &gt; 2 &amp;&amp; suffix[2] == &#39;\t&#39; {
                <a id="L501"></a><span class="comment">// assume the &#39;\t&#39; compensates for the /*</span>
                <a id="L502"></a>suffix = suffix[2:n]
            <a id="L503"></a>} else {
                <a id="L504"></a><span class="comment">// otherwise assume two blanks</span>
                <a id="L505"></a>suffix[0], suffix[1] = &#39; &#39;, &#39; &#39;;
                <a id="L506"></a>suffix = suffix[0:n];
            <a id="L507"></a>}
            <a id="L508"></a><span class="comment">// Shorten the computed common prefix by the length of</span>
            <a id="L509"></a><span class="comment">// suffix, if it is found as suffix of the prefix.</span>
            <a id="L510"></a>if bytes.HasSuffix(prefix, suffix) {
                <a id="L511"></a>prefix = prefix[0 : len(prefix)-len(suffix)]
            <a id="L512"></a>}
        <a id="L513"></a>}
    <a id="L514"></a>}

    <a id="L516"></a><span class="comment">// Handle last line: If it only contains a closing */, align it</span>
    <a id="L517"></a><span class="comment">// with the opening /*, otherwise align the text with the other</span>
    <a id="L518"></a><span class="comment">// lines.</span>
    <a id="L519"></a>last := lines[len(lines)-1];
    <a id="L520"></a>closing := []byte{&#39;*&#39;, &#39;/&#39;};
    <a id="L521"></a>i := bytes.Index(last, closing);
    <a id="L522"></a>if isBlank(last[0:i]) {
        <a id="L523"></a><span class="comment">// last line only contains closing */</span>
        <a id="L524"></a>var sep []byte;
        <a id="L525"></a>if lineOfStars {
            <a id="L526"></a><span class="comment">// insert an aligning blank</span>
            <a id="L527"></a>sep = []byte{&#39; &#39;}
        <a id="L528"></a>}
        <a id="L529"></a>lines[len(lines)-1] = bytes.Join([][]byte{prefix, closing}, sep);
    <a id="L530"></a>} else {
        <a id="L531"></a><span class="comment">// last line contains more comment text - assume</span>
        <a id="L532"></a><span class="comment">// it is aligned like the other lines</span>
        <a id="L533"></a>prefix = commonPrefix(prefix, last)
    <a id="L534"></a>}

    <a id="L536"></a><span class="comment">// Remove the common prefix from all but the first and empty lines.</span>
    <a id="L537"></a>for i, line := range lines {
        <a id="L538"></a>if i &gt; 0 &amp;&amp; len(line) != 0 {
            <a id="L539"></a>lines[i] = line[len(prefix):len(line)]
        <a id="L540"></a>}
    <a id="L541"></a>}
<a id="L542"></a>}


<a id="L545"></a>func (p *printer) writeComment(comment *ast.Comment) {
    <a id="L546"></a>text := comment.Text;

    <a id="L548"></a><span class="comment">// shortcut common case of //-style comments</span>
    <a id="L549"></a>if text[1] == &#39;/&#39; {
        <a id="L550"></a>p.writeCommentLine(comment, comment.Pos(), text);
        <a id="L551"></a>return;
    <a id="L552"></a>}

    <a id="L554"></a><span class="comment">// for /*-style comments, print line by line and let the</span>
    <a id="L555"></a><span class="comment">// write function take care of the proper indentation</span>
    <a id="L556"></a>lines := split(text);
    <a id="L557"></a>stripCommonPrefix(lines);

    <a id="L559"></a><span class="comment">// write comment lines, separated by formfeed,</span>
    <a id="L560"></a><span class="comment">// without a line break after the last line</span>
    <a id="L561"></a>linebreak := formfeeds[0:1];
    <a id="L562"></a>pos := comment.Pos();
    <a id="L563"></a>for i, line := range lines {
        <a id="L564"></a>if i &gt; 0 {
            <a id="L565"></a>p.write(linebreak);
            <a id="L566"></a>pos = p.pos;
        <a id="L567"></a>}
        <a id="L568"></a>if len(line) &gt; 0 {
            <a id="L569"></a>p.writeCommentLine(comment, pos, line)
        <a id="L570"></a>}
    <a id="L571"></a>}
<a id="L572"></a>}


<a id="L575"></a><span class="comment">// writeCommentSuffix writes a line break after a comment if indicated</span>
<a id="L576"></a><span class="comment">// and processes any leftover indentation information. If a line break</span>
<a id="L577"></a><span class="comment">// is needed, the kind of break (newline vs formfeed) depends on the</span>
<a id="L578"></a><span class="comment">// pending whitespace.</span>
<a id="L579"></a><span class="comment">//</span>
<a id="L580"></a>func (p *printer) writeCommentSuffix(needsLinebreak bool) {
    <a id="L581"></a>for i, ch := range p.buffer {
        <a id="L582"></a>switch ch {
        <a id="L583"></a>case blank, vtab:
            <a id="L584"></a><span class="comment">// ignore trailing whitespace</span>
            <a id="L585"></a>p.buffer[i] = ignore
        <a id="L586"></a>case indent, unindent:
            <a id="L587"></a><span class="comment">// don&#39;t loose indentation information</span>
        <a id="L588"></a>case newline, formfeed:
            <a id="L589"></a><span class="comment">// if we need a line break, keep exactly one</span>
            <a id="L590"></a>if needsLinebreak {
                <a id="L591"></a>needsLinebreak = false
            <a id="L592"></a>} else {
                <a id="L593"></a>p.buffer[i] = ignore
            <a id="L594"></a>}
        <a id="L595"></a>}
    <a id="L596"></a>}
    <a id="L597"></a>p.writeWhitespace(len(p.buffer));

    <a id="L599"></a><span class="comment">// make sure we have a line break</span>
    <a id="L600"></a>if needsLinebreak {
        <a id="L601"></a>p.write([]byte{&#39;\n&#39;})
    <a id="L602"></a>}
<a id="L603"></a>}


<a id="L606"></a><span class="comment">// intersperseComments consumes all comments that appear before the next token</span>
<a id="L607"></a><span class="comment">// and prints it together with the buffered whitespace (i.e., the whitespace</span>
<a id="L608"></a><span class="comment">// that needs to be written before the next token). A heuristic is used to mix</span>
<a id="L609"></a><span class="comment">// the comments and whitespace. The isKeyword parameter indicates if the next</span>
<a id="L610"></a><span class="comment">// token is a keyword or not.</span>
<a id="L611"></a><span class="comment">//</span>
<a id="L612"></a>func (p *printer) intersperseComments(next token.Position, isKeyword bool) {
    <a id="L613"></a>isFirst := true;
    <a id="L614"></a>needsLinebreak := false;
    <a id="L615"></a>var last *ast.Comment;
    <a id="L616"></a>for ; p.commentBefore(next); p.comment = p.comment.Next {
        <a id="L617"></a>for _, c := range p.comment.List {
            <a id="L618"></a>p.writeCommentPrefix(c.Pos(), next, isFirst, isKeyword);
            <a id="L619"></a>isFirst = false;
            <a id="L620"></a>p.writeComment(c);
            <a id="L621"></a>needsLinebreak = c.Text[1] == &#39;/&#39;;
            <a id="L622"></a>last = c;
        <a id="L623"></a>}
    <a id="L624"></a>}
    <a id="L625"></a>if last != nil &amp;&amp; !needsLinebreak &amp;&amp; last.Pos().Line == next.Line {
        <a id="L626"></a><span class="comment">// the last comment is a /*-style comment and the next item</span>
        <a id="L627"></a><span class="comment">// follows on the same line: separate with an extra blank</span>
        <a id="L628"></a>p.write([]byte{&#39; &#39;})
    <a id="L629"></a>}
    <a id="L630"></a>p.writeCommentSuffix(needsLinebreak);
<a id="L631"></a>}


<a id="L634"></a><span class="comment">// whiteWhitespace writes the first n whitespace entries.</span>
<a id="L635"></a>func (p *printer) writeWhitespace(n int) {
    <a id="L636"></a><span class="comment">// write entries</span>
    <a id="L637"></a>var data [1]byte;
    <a id="L638"></a>for i := 0; i &lt; n; i++ {
        <a id="L639"></a>switch ch := p.buffer[i]; ch {
        <a id="L640"></a>case ignore:
            <a id="L641"></a><span class="comment">// ignore!</span>
        <a id="L642"></a>case indent:
            <a id="L643"></a>p.indent++
        <a id="L644"></a>case unindent:
            <a id="L645"></a>p.indent--;
            <a id="L646"></a>if p.indent &lt; 0 {
                <a id="L647"></a>p.internalError(&#34;negative indentation:&#34;, p.indent);
                <a id="L648"></a>p.indent = 0;
            <a id="L649"></a>}
        <a id="L650"></a>case newline, formfeed:
            <a id="L651"></a><span class="comment">// A line break immediately followed by a &#34;correcting&#34;</span>
            <a id="L652"></a><span class="comment">// unindent is swapped with the unindent - this permits</span>
            <a id="L653"></a><span class="comment">// proper label positioning. If a comment is between</span>
            <a id="L654"></a><span class="comment">// the line break and the label, the unindent is not</span>
            <a id="L655"></a><span class="comment">// part of the comment whitespace prefix and the comment</span>
            <a id="L656"></a><span class="comment">// will be positioned correctly indented.</span>
            <a id="L657"></a>if i+1 &lt; n &amp;&amp; p.buffer[i+1] == unindent {
                <a id="L658"></a><span class="comment">// Use a formfeed to terminate the current section.</span>
                <a id="L659"></a><span class="comment">// Otherwise, a long label name on the next line leading</span>
                <a id="L660"></a><span class="comment">// to a wide column may increase the indentation column</span>
                <a id="L661"></a><span class="comment">// of lines before the label; effectively leading to wrong</span>
                <a id="L662"></a><span class="comment">// indentation.</span>
                <a id="L663"></a>p.buffer[i], p.buffer[i+1] = unindent, formfeed;
                <a id="L664"></a>i--; <span class="comment">// do it again</span>
                <a id="L665"></a>continue;
            <a id="L666"></a>}
            <a id="L667"></a>fallthrough;
        <a id="L668"></a>default:
            <a id="L669"></a>data[0] = byte(ch);
            <a id="L670"></a>p.write(&amp;data);
        <a id="L671"></a>}
    <a id="L672"></a>}

    <a id="L674"></a><span class="comment">// shift remaining entries down</span>
    <a id="L675"></a>i := 0;
    <a id="L676"></a>for ; n &lt; len(p.buffer); n++ {
        <a id="L677"></a>p.buffer[i] = p.buffer[n];
        <a id="L678"></a>i++;
    <a id="L679"></a>}
    <a id="L680"></a>p.buffer = p.buffer[0:i];
<a id="L681"></a>}


<a id="L684"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L685"></a><span class="comment">// Printing interface</span>

<a id="L687"></a><span class="comment">// print prints a list of &#34;items&#34; (roughly corresponding to syntactic</span>
<a id="L688"></a><span class="comment">// tokens, but also including whitespace and formatting information).</span>
<a id="L689"></a><span class="comment">// It is the only print function that should be called directly from</span>
<a id="L690"></a><span class="comment">// any of the AST printing functions in nodes.go.</span>
<a id="L691"></a><span class="comment">//</span>
<a id="L692"></a><span class="comment">// Whitespace is accumulated until a non-whitespace token appears. Any</span>
<a id="L693"></a><span class="comment">// comments that need to appear before that token are printed first,</span>
<a id="L694"></a><span class="comment">// taking into account the amount and structure of any pending white-</span>
<a id="L695"></a><span class="comment">// space for best comment placement. Then, any leftover whitespace is</span>
<a id="L696"></a><span class="comment">// printed, followed by the actual token.</span>
<a id="L697"></a><span class="comment">//</span>
<a id="L698"></a>func (p *printer) print(args ...) {
    <a id="L699"></a>v := reflect.NewValue(args).(*reflect.StructValue);
    <a id="L700"></a>for i := 0; i &lt; v.NumField(); i++ {
        <a id="L701"></a>f := v.Field(i);

        <a id="L703"></a>next := p.pos; <span class="comment">// estimated position of next item</span>
        <a id="L704"></a>var data []byte;
        <a id="L705"></a>var tag HTMLTag;
        <a id="L706"></a>isKeyword := false;
        <a id="L707"></a>switch x := f.Interface().(type) {
        <a id="L708"></a>case whiteSpace:
            <a id="L709"></a>if x == ignore {
                <a id="L710"></a><span class="comment">// don&#39;t add ignore&#39;s to the buffer; they</span>
                <a id="L711"></a><span class="comment">// may screw up &#34;correcting&#34; unindents (see</span>
                <a id="L712"></a><span class="comment">// LabeledStmt)</span>
                <a id="L713"></a>break
            <a id="L714"></a>}
            <a id="L715"></a>i := len(p.buffer);
            <a id="L716"></a>if i == cap(p.buffer) {
                <a id="L717"></a><span class="comment">// Whitespace sequences are very short so this should</span>
                <a id="L718"></a><span class="comment">// never happen. Handle gracefully (but possibly with</span>
                <a id="L719"></a><span class="comment">// bad comment placement) if it does happen.</span>
                <a id="L720"></a>p.writeWhitespace(i);
                <a id="L721"></a>i = 0;
            <a id="L722"></a>}
            <a id="L723"></a>p.buffer = p.buffer[0 : i+1];
            <a id="L724"></a>p.buffer[i] = x;
        <a id="L725"></a>case []byte:
            <a id="L726"></a><span class="comment">// TODO(gri): remove this case once commentList</span>
            <a id="L727"></a><span class="comment">//            handles comments correctly</span>
            <a id="L728"></a>data = x
        <a id="L729"></a>case string:
            <a id="L730"></a><span class="comment">// TODO(gri): remove this case once fieldList</span>
            <a id="L731"></a><span class="comment">//            handles comments correctly</span>
            <a id="L732"></a>data = strings.Bytes(x)
        <a id="L733"></a>case *ast.Ident:
            <a id="L734"></a>if p.Styler != nil {
                <a id="L735"></a>data, tag = p.Styler.Ident(x)
            <a id="L736"></a>} else {
                <a id="L737"></a>data = strings.Bytes(x.Value)
            <a id="L738"></a>}
        <a id="L739"></a>case *ast.BasicLit:
            <a id="L740"></a>if p.Styler != nil {
                <a id="L741"></a>data, tag = p.Styler.BasicLit(x)
            <a id="L742"></a>} else {
                <a id="L743"></a>data = x.Value
            <a id="L744"></a>}
            <a id="L745"></a><span class="comment">// escape all literals so they pass through unchanged</span>
            <a id="L746"></a><span class="comment">// (note that valid Go programs cannot contain esc (&#39;\xff&#39;)</span>
            <a id="L747"></a><span class="comment">// bytes since they do not appear in legal UTF-8 sequences)</span>
            <a id="L748"></a><span class="comment">// TODO(gri): this this more efficiently.</span>
            <a id="L749"></a>data = strings.Bytes(&#34;\xff&#34; + string(data) + &#34;\xff&#34;);
        <a id="L750"></a>case token.Token:
            <a id="L751"></a>if p.Styler != nil {
                <a id="L752"></a>data, tag = p.Styler.Token(x)
            <a id="L753"></a>} else {
                <a id="L754"></a>data = strings.Bytes(x.String())
            <a id="L755"></a>}
            <a id="L756"></a>isKeyword = x.IsKeyword();
        <a id="L757"></a>case token.Position:
            <a id="L758"></a>if x.IsValid() {
                <a id="L759"></a>next = x <span class="comment">// accurate position of next item</span>
            <a id="L760"></a>}
        <a id="L761"></a>default:
            <a id="L762"></a>panicln(&#34;print: unsupported argument type&#34;, f.Type().String())
        <a id="L763"></a>}
        <a id="L764"></a>p.pos = next;

        <a id="L766"></a>if data != nil {
            <a id="L767"></a>p.flush(next, isKeyword);

            <a id="L769"></a><span class="comment">// intersperse extra newlines if present in the source</span>
            <a id="L770"></a><span class="comment">// (don&#39;t do this in flush as it will cause extra newlines</span>
            <a id="L771"></a><span class="comment">// at the end of a file)</span>
            <a id="L772"></a>p.writeNewlines(next.Line - p.pos.Line);

            <a id="L774"></a>p.writeItem(next, data, tag);
        <a id="L775"></a>}
    <a id="L776"></a>}
<a id="L777"></a>}


<a id="L780"></a><span class="comment">// commentBefore returns true iff the current comment occurs</span>
<a id="L781"></a><span class="comment">// before the next position in the source code.</span>
<a id="L782"></a><span class="comment">//</span>
<a id="L783"></a>func (p *printer) commentBefore(next token.Position) bool {
    <a id="L784"></a>return p.comment != nil &amp;&amp; p.comment.List[0].Pos().Offset &lt; next.Offset
<a id="L785"></a>}


<a id="L788"></a><span class="comment">// Flush prints any pending comments and whitespace occuring</span>
<a id="L789"></a><span class="comment">// textually before the position of the next item.</span>
<a id="L790"></a><span class="comment">//</span>
<a id="L791"></a>func (p *printer) flush(next token.Position, isKeyword bool) {
    <a id="L792"></a><span class="comment">// if there are comments before the next item, intersperse them</span>
    <a id="L793"></a>if p.commentBefore(next) {
        <a id="L794"></a>p.intersperseComments(next, isKeyword)
    <a id="L795"></a>}
    <a id="L796"></a><span class="comment">// write any leftover whitespace</span>
    <a id="L797"></a>p.writeWhitespace(len(p.buffer));
<a id="L798"></a>}


<a id="L801"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L802"></a><span class="comment">// Trimmer</span>

<a id="L804"></a><span class="comment">// A trimmer is an io.Writer filter for stripping tabwriter.Escape</span>
<a id="L805"></a><span class="comment">// characters, trailing blanks and tabs, and for converting formfeed</span>
<a id="L806"></a><span class="comment">// and vtab characters into newlines and htabs (in case no tabwriter</span>
<a id="L807"></a><span class="comment">// is used).</span>
<a id="L808"></a><span class="comment">//</span>
<a id="L809"></a>type trimmer struct {
    <a id="L810"></a>output io.Writer;
    <a id="L811"></a>buf    bytes.Buffer;
<a id="L812"></a>}


<a id="L815"></a><span class="comment">// Design note: It is tempting to eliminate extra blanks occuring in</span>
<a id="L816"></a><span class="comment">//              whitespace in this function as it could simplify some</span>
<a id="L817"></a><span class="comment">//              of the blanks logic in the node printing functions.</span>
<a id="L818"></a><span class="comment">//              However, this would mess up any formatting done by</span>
<a id="L819"></a><span class="comment">//              the tabwriter.</span>

<a id="L821"></a>func (p *trimmer) Write(data []byte) (n int, err os.Error) {
    <a id="L822"></a><span class="comment">// m &lt; 0: no unwritten data except for whitespace</span>
    <a id="L823"></a><span class="comment">// m &gt;= 0: data[m:n] unwritten and no whitespace</span>
    <a id="L824"></a>m := 0;
    <a id="L825"></a>if p.buf.Len() &gt; 0 {
        <a id="L826"></a>m = -1
    <a id="L827"></a>}

    <a id="L829"></a>var b byte;
    <a id="L830"></a>for n, b = range data {
        <a id="L831"></a>switch b {
        <a id="L832"></a>default:
            <a id="L833"></a><span class="comment">// write any pending whitespace</span>
            <a id="L834"></a>if m &lt; 0 {
                <a id="L835"></a>if _, err = p.output.Write(p.buf.Bytes()); err != nil {
                    <a id="L836"></a>return
                <a id="L837"></a>}
                <a id="L838"></a>p.buf.Reset();
                <a id="L839"></a>m = n;
            <a id="L840"></a>}

        <a id="L842"></a>case &#39;\v&#39;:
            <a id="L843"></a>b = &#39;\t&#39;; <span class="comment">// convert to htab</span>
            <a id="L844"></a>fallthrough;

        <a id="L846"></a>case &#39;\t&#39;, &#39; &#39;, tabwriter.Escape:
            <a id="L847"></a><span class="comment">// write any pending (non-whitespace) data</span>
            <a id="L848"></a>if m &gt;= 0 {
                <a id="L849"></a>if _, err = p.output.Write(data[m:n]); err != nil {
                    <a id="L850"></a>return
                <a id="L851"></a>}
                <a id="L852"></a>m = -1;
            <a id="L853"></a>}
            <a id="L854"></a><span class="comment">// collect whitespace but discard tabrwiter.Escapes.</span>
            <a id="L855"></a>if b != tabwriter.Escape {
                <a id="L856"></a>p.buf.WriteByte(b) <span class="comment">// WriteByte returns no errors</span>
            <a id="L857"></a>}

        <a id="L859"></a>case &#39;\f&#39;, &#39;\n&#39;:
            <a id="L860"></a><span class="comment">// discard whitespace</span>
            <a id="L861"></a>p.buf.Reset();
            <a id="L862"></a><span class="comment">// write any pending (non-whitespace) data</span>
            <a id="L863"></a>if m &gt;= 0 {
                <a id="L864"></a>if _, err = p.output.Write(data[m:n]); err != nil {
                    <a id="L865"></a>return
                <a id="L866"></a>}
                <a id="L867"></a>m = -1;
            <a id="L868"></a>}
            <a id="L869"></a><span class="comment">// convert formfeed into newline</span>
            <a id="L870"></a>if _, err = p.output.Write(newlines[0:1]); err != nil {
                <a id="L871"></a>return
            <a id="L872"></a>}
        <a id="L873"></a>}
    <a id="L874"></a>}
    <a id="L875"></a>n = len(data);

    <a id="L877"></a><span class="comment">// write any pending non-whitespace</span>
    <a id="L878"></a>if m &gt;= 0 {
        <a id="L879"></a>if _, err = p.output.Write(data[m:n]); err != nil {
            <a id="L880"></a>return
        <a id="L881"></a>}
    <a id="L882"></a>}

    <a id="L884"></a>return;
<a id="L885"></a>}


<a id="L888"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L889"></a><span class="comment">// Public interface</span>

<a id="L891"></a><span class="comment">// General printing is controlled with these Config.Mode flags.</span>
<a id="L892"></a>const (
    <a id="L893"></a>GenHTML    uint = 1 &lt;&lt; iota; <span class="comment">// generate HTML</span>
    <a id="L894"></a>RawFormat;      <span class="comment">// do not use a tabwriter; if set, UseSpaces is ignored</span>
    <a id="L895"></a>UseSpaces;      <span class="comment">// use spaces instead of tabs for indentation and alignment</span>
<a id="L896"></a>)


<a id="L899"></a><span class="comment">// An HTMLTag specifies a start and end tag.</span>
<a id="L900"></a>type HTMLTag struct {
    <a id="L901"></a>Start, End string; <span class="comment">// empty if tags are absent</span>
<a id="L902"></a>}


<a id="L905"></a><span class="comment">// A Styler specifies formatting of line tags and elementary Go words.</span>
<a id="L906"></a><span class="comment">// A format consists of text and a (possibly empty) surrounding HTML tag.</span>
<a id="L907"></a><span class="comment">//</span>
<a id="L908"></a>type Styler interface {
    <a id="L909"></a>LineTag(line int) ([]byte, HTMLTag);
    <a id="L910"></a>Comment(c *ast.Comment, line []byte) ([]byte, HTMLTag);
    <a id="L911"></a>BasicLit(x *ast.BasicLit) ([]byte, HTMLTag);
    <a id="L912"></a>Ident(id *ast.Ident) ([]byte, HTMLTag);
    <a id="L913"></a>Token(tok token.Token) ([]byte, HTMLTag);
<a id="L914"></a>}


<a id="L917"></a><span class="comment">// A Config node controls the output of Fprint.</span>
<a id="L918"></a>type Config struct {
    <a id="L919"></a>Mode     uint;   <span class="comment">// default: 0</span>
    <a id="L920"></a>Tabwidth int;    <span class="comment">// default: 8</span>
    <a id="L921"></a>Styler   Styler; <span class="comment">// default: nil</span>
<a id="L922"></a>}


<a id="L925"></a><span class="comment">// Fprint &#34;pretty-prints&#34; an AST node to output and returns the number</span>
<a id="L926"></a><span class="comment">// of bytes written and an error (if any) for a given configuration cfg.</span>
<a id="L927"></a><span class="comment">// The node type must be *ast.File, or assignment-compatible to ast.Expr,</span>
<a id="L928"></a><span class="comment">// ast.Decl, or ast.Stmt.</span>
<a id="L929"></a><span class="comment">//</span>
<a id="L930"></a>func (cfg *Config) Fprint(output io.Writer, node interface{}) (int, os.Error) {
    <a id="L931"></a><span class="comment">// redirect output through a trimmer to eliminate trailing whitespace</span>
    <a id="L932"></a><span class="comment">// (Input to a tabwriter must be untrimmed since trailing tabs provide</span>
    <a id="L933"></a><span class="comment">// formatting information. The tabwriter could provide trimming</span>
    <a id="L934"></a><span class="comment">// functionality but no tabwriter is used when RawFormat is set.)</span>
    <a id="L935"></a>output = &amp;trimmer{output: output};

    <a id="L937"></a><span class="comment">// setup tabwriter if needed and redirect output</span>
    <a id="L938"></a>var tw *tabwriter.Writer;
    <a id="L939"></a>if cfg.Mode&amp;RawFormat == 0 {
        <a id="L940"></a>padchar := byte(&#39;\t&#39;);
        <a id="L941"></a>if cfg.Mode&amp;UseSpaces != 0 {
            <a id="L942"></a>padchar = &#39; &#39;
        <a id="L943"></a>}
        <a id="L944"></a>twmode := tabwriter.DiscardEmptyColumns;
        <a id="L945"></a>if cfg.Mode&amp;GenHTML != 0 {
            <a id="L946"></a>twmode |= tabwriter.FilterHTML
        <a id="L947"></a>}
        <a id="L948"></a>tw = tabwriter.NewWriter(output, cfg.Tabwidth, 1, padchar, twmode);
        <a id="L949"></a>output = tw;
    <a id="L950"></a>}

    <a id="L952"></a><span class="comment">// setup printer and print node</span>
    <a id="L953"></a>var p printer;
    <a id="L954"></a>p.init(output, cfg);
    <a id="L955"></a>go func() {
        <a id="L956"></a>switch n := node.(type) {
        <a id="L957"></a>case ast.Expr:
            <a id="L958"></a>p.expr(n, ignoreMultiLine)
        <a id="L959"></a>case ast.Stmt:
            <a id="L960"></a>p.stmt(n, ignoreMultiLine)
        <a id="L961"></a>case ast.Decl:
            <a id="L962"></a>p.decl(n, atTop, ignoreMultiLine)
        <a id="L963"></a>case *ast.File:
            <a id="L964"></a>p.comment = n.Comments;
            <a id="L965"></a>p.file(n);
        <a id="L966"></a>default:
            <a id="L967"></a>p.errors &lt;- os.NewError(fmt.Sprintf(&#34;printer.Fprint: unsupported node type %T&#34;, n));
            <a id="L968"></a>runtime.Goexit();
        <a id="L969"></a>}
        <a id="L970"></a>p.flush(token.Position{Offset: 1 &lt;&lt; 30, Line: 1 &lt;&lt; 30}, false); <span class="comment">// flush to &#34;infinity&#34;</span>
        <a id="L971"></a>p.errors &lt;- nil;                                                <span class="comment">// no errors</span>
    <a id="L972"></a>}();
    <a id="L973"></a>err := &lt;-p.errors; <span class="comment">// wait for completion of goroutine</span>

    <a id="L975"></a><span class="comment">// flush tabwriter, if any</span>
    <a id="L976"></a>if tw != nil {
        <a id="L977"></a>tw.Flush() <span class="comment">// ignore errors</span>
    <a id="L978"></a>}

    <a id="L980"></a>return p.written, err;
<a id="L981"></a>}


<a id="L984"></a><span class="comment">// Fprint &#34;pretty-prints&#34; an AST node to output.</span>
<a id="L985"></a><span class="comment">// It calls Config.Fprint with default settings.</span>
<a id="L986"></a><span class="comment">//</span>
<a id="L987"></a>func Fprint(output io.Writer, node interface{}) os.Error {
    <a id="L988"></a>_, err := (&amp;Config{Tabwidth: 8}).Fprint(output, node); <span class="comment">// don&#39;t care about number of bytes written</span>
    <a id="L989"></a>return err;
<a id="L990"></a>}
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
