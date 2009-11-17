<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/tabwriter/tabwriter.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/tabwriter/tabwriter.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The tabwriter package implements a write filter (tabwriter.Writer)</span>
<a id="L6"></a><span class="comment">// that translates tabbed columns in input into properly aligned text.</span>
<a id="L7"></a><span class="comment">//</span>
<a id="L8"></a><span class="comment">// The package is using the Elastic Tabstops algorithm described at</span>
<a id="L9"></a><span class="comment">// http://nickgravgaard.com/elastictabstops/index.html.</span>
<a id="L10"></a><span class="comment">//</span>
<a id="L11"></a>package tabwriter

<a id="L13"></a>import (
    <a id="L14"></a>&#34;bytes&#34;;
    <a id="L15"></a>&#34;container/vector&#34;;
    <a id="L16"></a>&#34;io&#34;;
    <a id="L17"></a>&#34;os&#34;;
    <a id="L18"></a>&#34;utf8&#34;;
<a id="L19"></a>)


<a id="L22"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L23"></a><span class="comment">// Filter implementation</span>

<a id="L25"></a><span class="comment">// A cell represents a segment of text delineated by tabs, formfeed,</span>
<a id="L26"></a><span class="comment">// or newline chars. The text itself is stored in a separate buffer;</span>
<a id="L27"></a><span class="comment">// cell only describes the segment&#39;s size in bytes, its width in runes,</span>
<a id="L28"></a><span class="comment">// and whether it&#39;s an htab (&#39;\t&#39;) or vtab (&#39;\v&#39;) terminated call.</span>
<a id="L29"></a><span class="comment">//</span>
<a id="L30"></a>type cell struct {
    <a id="L31"></a>size  int;  <span class="comment">// cell size in bytes</span>
    <a id="L32"></a>width int;  <span class="comment">// cell width in runes</span>
    <a id="L33"></a>htab  bool; <span class="comment">// true if the cell is terminated by an htab (&#39;\t&#39;)</span>
<a id="L34"></a>}


<a id="L37"></a><span class="comment">// A Writer is a filter that inserts padding around</span>
<a id="L38"></a><span class="comment">// tab-delimited columns in its input to align them</span>
<a id="L39"></a><span class="comment">// in the output.</span>
<a id="L40"></a><span class="comment">//</span>
<a id="L41"></a><span class="comment">// The Writer treats incoming bytes as UTF-8 encoded text</span>
<a id="L42"></a><span class="comment">// consisting of tab-terminated cells. Cells in adjacent lines</span>
<a id="L43"></a><span class="comment">// constitute a column. The Writer inserts padding as needed</span>
<a id="L44"></a><span class="comment">// to make all cells in a column have the same width, effectively</span>
<a id="L45"></a><span class="comment">// aligning the columns. Note that cells are tab-terminated,</span>
<a id="L46"></a><span class="comment">// not tab-separated: trailing non-tab text at the end of a line</span>
<a id="L47"></a><span class="comment">// is not part of any cell.</span>
<a id="L48"></a><span class="comment">//</span>
<a id="L49"></a><span class="comment">// Horizontal and vertical tabs may be used to terminate a cell.</span>
<a id="L50"></a><span class="comment">// If DiscardEmptyColumns is set, empty columns that are terminated</span>
<a id="L51"></a><span class="comment">// entirely by vertical (or &#34;soft&#34;) tabs are discarded. Columns</span>
<a id="L52"></a><span class="comment">// terminated by horizontal (or &#34;hard&#34;) tabs are not affected by</span>
<a id="L53"></a><span class="comment">// this flag.</span>
<a id="L54"></a><span class="comment">//</span>
<a id="L55"></a><span class="comment">// A segment of text may be escaped by bracketing it with Escape</span>
<a id="L56"></a><span class="comment">// characters. The tabwriter strips the Escape characters but otherwise</span>
<a id="L57"></a><span class="comment">// passes escaped text segments through unchanged. In particular, it</span>
<a id="L58"></a><span class="comment">// does not interpret any tabs or line breaks within the segment.</span>
<a id="L59"></a><span class="comment">//</span>
<a id="L60"></a><span class="comment">// The Writer assumes that all characters have the same width;</span>
<a id="L61"></a><span class="comment">// this may not be true in some fonts, especially with certain</span>
<a id="L62"></a><span class="comment">// UTF-8 characters.</span>
<a id="L63"></a><span class="comment">//</span>
<a id="L64"></a><span class="comment">// If a Writer is configured to filter HTML, HTML tags and entities</span>
<a id="L65"></a><span class="comment">// are simply passed through. The widths of tags and entities are</span>
<a id="L66"></a><span class="comment">// assumed to be zero (tags) and one (entities) for formatting purposes.</span>
<a id="L67"></a><span class="comment">//</span>
<a id="L68"></a><span class="comment">// The formfeed character (&#39;\f&#39;) acts like a newline but it also</span>
<a id="L69"></a><span class="comment">// terminates all columns in the current line (effectively calling</span>
<a id="L70"></a><span class="comment">// Flush). Cells in the next line start new columns. Unless found</span>
<a id="L71"></a><span class="comment">// inside an HTML tag or inside an escaped text segment, formfeed</span>
<a id="L72"></a><span class="comment">// characters appear as newlines in the output.</span>
<a id="L73"></a><span class="comment">//</span>
<a id="L74"></a><span class="comment">// The Writer must buffer input internally, because proper spacing</span>
<a id="L75"></a><span class="comment">// of one line may depend on the cells in future lines. Clients must</span>
<a id="L76"></a><span class="comment">// call Flush when done calling Write.</span>
<a id="L77"></a><span class="comment">//</span>
<a id="L78"></a>type Writer struct {
    <a id="L79"></a><span class="comment">// configuration</span>
    <a id="L80"></a>output    io.Writer;
    <a id="L81"></a>cellwidth int;
    <a id="L82"></a>padding   int;
    <a id="L83"></a>padbytes  [8]byte;
    <a id="L84"></a>flags     uint;

    <a id="L86"></a><span class="comment">// current state</span>
    <a id="L87"></a>buf     bytes.Buffer;     <span class="comment">// collected text w/o tabs, newlines, or formfeed chars</span>
    <a id="L88"></a>pos     int;              <span class="comment">// buffer position up to which width of incomplete cell has been computed</span>
    <a id="L89"></a>cell    cell;             <span class="comment">// current incomplete cell; cell.width is up to buf[pos] w/o ignored sections</span>
    <a id="L90"></a>endChar byte;             <span class="comment">// terminating char of escaped sequence (Escape for escapes, &#39;&gt;&#39;, &#39;;&#39; for HTML tags/entities, or 0)</span>
    <a id="L91"></a>lines   vector.Vector;    <span class="comment">// list if lines; each line is a list of cells</span>
    <a id="L92"></a>widths  vector.IntVector; <span class="comment">// list of column widths in runes - re-used during formatting</span>
<a id="L93"></a>}


<a id="L96"></a>func (b *Writer) addLine() { b.lines.Push(vector.New(0)) }


<a id="L99"></a>func (b *Writer) line(i int) *vector.Vector { return b.lines.At(i).(*vector.Vector) }


<a id="L102"></a><span class="comment">// Reset the current state.</span>
<a id="L103"></a>func (b *Writer) reset() {
    <a id="L104"></a>b.buf.Reset();
    <a id="L105"></a>b.pos = 0;
    <a id="L106"></a>b.cell = cell{};
    <a id="L107"></a>b.endChar = 0;
    <a id="L108"></a>b.lines.Init(0);
    <a id="L109"></a>b.widths.Init(0);
    <a id="L110"></a>b.addLine();
<a id="L111"></a>}


<a id="L114"></a><span class="comment">// Internal representation (current state):</span>
<a id="L115"></a><span class="comment">//</span>
<a id="L116"></a><span class="comment">// - all text written is appended to buf; formfeed chars, tabs and newlines are stripped away</span>
<a id="L117"></a><span class="comment">// - at any given time there is a (possibly empty) incomplete cell at the end</span>
<a id="L118"></a><span class="comment">//   (the cell starts after a tab, formfeed, or newline)</span>
<a id="L119"></a><span class="comment">// - cell.size is the number of bytes belonging to the cell so far</span>
<a id="L120"></a><span class="comment">// - cell.width is text width in runes of that cell from the start of the cell to</span>
<a id="L121"></a><span class="comment">//   position pos; html tags and entities are excluded from this width if html</span>
<a id="L122"></a><span class="comment">//   filtering is enabled</span>
<a id="L123"></a><span class="comment">// - the sizes and widths of processed text are kept in the lines vector</span>
<a id="L124"></a><span class="comment">//   which contains a vector of cells for each line</span>
<a id="L125"></a><span class="comment">// - the widths vector is a temporary vector with current widths used during</span>
<a id="L126"></a><span class="comment">//   formatting; it is kept in Writer because it&#39;s re-used</span>
<a id="L127"></a><span class="comment">//</span>
<a id="L128"></a><span class="comment">//                    |&lt;---------- size ----------&gt;|</span>
<a id="L129"></a><span class="comment">//                    |                            |</span>
<a id="L130"></a><span class="comment">//                    |&lt;- width -&gt;|&lt;- ignored -&gt;|  |</span>
<a id="L131"></a><span class="comment">//                    |           |             |  |</span>
<a id="L132"></a><span class="comment">// [---processed---tab------------&lt;tag&gt;...&lt;/tag&gt;...]</span>
<a id="L133"></a><span class="comment">// ^                  ^                         ^</span>
<a id="L134"></a><span class="comment">// |                  |                         |</span>
<a id="L135"></a><span class="comment">// buf                start of incomplete cell  pos</span>


<a id="L138"></a><span class="comment">// Formatting can be controlled with these flags.</span>
<a id="L139"></a>const (
    <a id="L140"></a><span class="comment">// Ignore html tags and treat entities (starting with &#39;&amp;&#39;</span>
    <a id="L141"></a><span class="comment">// and ending in &#39;;&#39;) as single characters (width = 1).</span>
    <a id="L142"></a>FilterHTML uint = 1 &lt;&lt; iota;

    <a id="L144"></a><span class="comment">// Force right-alignment of cell content.</span>
    <a id="L145"></a><span class="comment">// Default is left-alignment.</span>
    <a id="L146"></a>AlignRight;

    <a id="L148"></a><span class="comment">// Handle empty columns as if they were not present in</span>
    <a id="L149"></a><span class="comment">// the input in the first place.</span>
    <a id="L150"></a>DiscardEmptyColumns;

    <a id="L152"></a><span class="comment">// Print a vertical bar (&#39;|&#39;) between columns (after formatting).</span>
    <a id="L153"></a><span class="comment">// Discarded colums appear as zero-width columns (&#34;||&#34;).</span>
    <a id="L154"></a>Debug;
<a id="L155"></a>)


<a id="L158"></a><span class="comment">// A Writer must be initialized with a call to Init. The first parameter (output)</span>
<a id="L159"></a><span class="comment">// specifies the filter output. The remaining parameters control the formatting:</span>
<a id="L160"></a><span class="comment">//</span>
<a id="L161"></a><span class="comment">//	cellwidth	minimal cell width</span>
<a id="L162"></a><span class="comment">//	padding		cell padding added to cell before computing its width</span>
<a id="L163"></a><span class="comment">//	padchar		ASCII char used for padding</span>
<a id="L164"></a><span class="comment">//			if padchar == &#39;\t&#39;, the Writer will assume that the</span>
<a id="L165"></a><span class="comment">//			width of a &#39;\t&#39; in the formatted output is cellwidth,</span>
<a id="L166"></a><span class="comment">//			and cells are left-aligned independent of align_left</span>
<a id="L167"></a><span class="comment">//			(for correct-looking results, cellwidth must correspond</span>
<a id="L168"></a><span class="comment">//			to the tab width in the viewer displaying the result)</span>
<a id="L169"></a><span class="comment">//	flags		formatting control</span>
<a id="L170"></a><span class="comment">//</span>
<a id="L171"></a><span class="comment">// To format in tab-separated columns with a tab stop of 8:</span>
<a id="L172"></a><span class="comment">//	b.Init(w, 8, 1, &#39;\t&#39;, 0);</span>
<a id="L173"></a><span class="comment">//</span>
<a id="L174"></a><span class="comment">// To format in space-separated columns with at least 4 spaces between columns:</span>
<a id="L175"></a><span class="comment">//	b.Init(w, 1, 4, &#39; &#39;, 0);</span>
<a id="L176"></a><span class="comment">//</span>
<a id="L177"></a>func (b *Writer) Init(output io.Writer, cellwidth, padding int, padchar byte, flags uint) *Writer {
    <a id="L178"></a>if cellwidth &lt; 0 {
        <a id="L179"></a>panic(&#34;negative cellwidth&#34;)
    <a id="L180"></a>}
    <a id="L181"></a>if padding &lt; 0 {
        <a id="L182"></a>panic(&#34;negative padding&#34;)
    <a id="L183"></a>}
    <a id="L184"></a>b.output = output;
    <a id="L185"></a>b.cellwidth = cellwidth;
    <a id="L186"></a>b.padding = padding;
    <a id="L187"></a>for i := len(b.padbytes) - 1; i &gt;= 0; i-- {
        <a id="L188"></a>b.padbytes[i] = padchar
    <a id="L189"></a>}
    <a id="L190"></a>if padchar == &#39;\t&#39; {
        <a id="L191"></a><span class="comment">// tab enforces left-alignment</span>
        <a id="L192"></a>flags &amp;^= AlignRight
    <a id="L193"></a>}
    <a id="L194"></a>b.flags = flags;

    <a id="L196"></a>b.reset();

    <a id="L198"></a>return b;
<a id="L199"></a>}


<a id="L202"></a><span class="comment">// debugging support (keep code around)</span>
<a id="L203"></a><span class="comment">/*</span>
<a id="L204"></a><span class="comment">func (b *Writer) dump() {</span>
<a id="L205"></a><span class="comment">	pos := 0;</span>
<a id="L206"></a><span class="comment">	for i := 0; i &lt; b.lines_size.Len(); i++ {</span>
<a id="L207"></a><span class="comment">		line_size, line_width := b.line(i);</span>
<a id="L208"></a><span class="comment">		print(&#34;(&#34;, i, &#34;) &#34;);</span>
<a id="L209"></a><span class="comment">		for j := 0; j &lt; line_size.Len(); j++ {</span>
<a id="L210"></a><span class="comment">			s := line_size.At(j);</span>
<a id="L211"></a><span class="comment">			print(&#34;[&#34;, string(b.buf.slice(pos, pos + s)), &#34;]&#34;);</span>
<a id="L212"></a><span class="comment">			pos += s;</span>
<a id="L213"></a><span class="comment">		}</span>
<a id="L214"></a><span class="comment">		print(&#34;\n&#34;);</span>
<a id="L215"></a><span class="comment">	}</span>
<a id="L216"></a><span class="comment">	print(&#34;\n&#34;);</span>
<a id="L217"></a><span class="comment">}</span>
<a id="L218"></a><span class="comment">*/</span>


<a id="L221"></a>func (b *Writer) write0(buf []byte) os.Error {
    <a id="L222"></a>n, err := b.output.Write(buf);
    <a id="L223"></a>if n != len(buf) &amp;&amp; err == nil {
        <a id="L224"></a>err = os.EIO
    <a id="L225"></a>}
    <a id="L226"></a>return err;
<a id="L227"></a>}


<a id="L230"></a>var newline = []byte{&#39;\n&#39;}

<a id="L232"></a>func (b *Writer) writePadding(textw, cellw int) os.Error {
    <a id="L233"></a>if b.cellwidth == 0 {
        <a id="L234"></a>return nil
    <a id="L235"></a>}

    <a id="L237"></a>if b.padbytes[0] == &#39;\t&#39; {
        <a id="L238"></a><span class="comment">// make cell width a multiple of cellwidth</span>
        <a id="L239"></a>cellw = ((cellw + b.cellwidth - 1) / b.cellwidth) * b.cellwidth
    <a id="L240"></a>}

    <a id="L242"></a>n := cellw - textw;
    <a id="L243"></a>if n &lt; 0 {
        <a id="L244"></a>panic(&#34;internal error&#34;)
    <a id="L245"></a>}

    <a id="L247"></a>if b.padbytes[0] == &#39;\t&#39; {
        <a id="L248"></a>n = (n + b.cellwidth - 1) / b.cellwidth
    <a id="L249"></a>}

    <a id="L251"></a>for n &gt; len(b.padbytes) {
        <a id="L252"></a>if err := b.write0(&amp;b.padbytes); err != nil {
            <a id="L253"></a>return err
        <a id="L254"></a>}
        <a id="L255"></a>n -= len(b.padbytes);
    <a id="L256"></a>}

    <a id="L258"></a>return b.write0(b.padbytes[0:n]);
<a id="L259"></a>}


<a id="L262"></a>var vbar = []byte{&#39;|&#39;}

<a id="L264"></a>func (b *Writer) writeLines(pos0 int, line0, line1 int) (pos int, err os.Error) {
    <a id="L265"></a>pos = pos0;
    <a id="L266"></a>for i := line0; i &lt; line1; i++ {
        <a id="L267"></a>line := b.line(i);
        <a id="L268"></a>for j := 0; j &lt; line.Len(); j++ {
            <a id="L269"></a>c := line.At(j).(cell);

            <a id="L271"></a>if j &gt; 0 &amp;&amp; b.flags&amp;Debug != 0 {
                <a id="L272"></a>if err = b.write0(vbar); err != nil {
                    <a id="L273"></a>return
                <a id="L274"></a>}
            <a id="L275"></a>}
            <a id="L276"></a>switch {
            <a id="L277"></a>default: <span class="comment">// align left</span>

                <a id="L279"></a>if err = b.write0(b.buf.Bytes()[pos : pos+c.size]); err != nil {
                    <a id="L280"></a>return
                <a id="L281"></a>}
                <a id="L282"></a>pos += c.size;
                <a id="L283"></a>if j &lt; b.widths.Len() {
                    <a id="L284"></a>if err = b.writePadding(c.width, b.widths.At(j)); err != nil {
                        <a id="L285"></a>return
                    <a id="L286"></a>}
                <a id="L287"></a>}

            <a id="L289"></a>case b.flags&amp;AlignRight != 0: <span class="comment">// align right</span>

                <a id="L291"></a>if j &lt; b.widths.Len() {
                    <a id="L292"></a>if err = b.writePadding(c.width, b.widths.At(j)); err != nil {
                        <a id="L293"></a>return
                    <a id="L294"></a>}
                <a id="L295"></a>}
                <a id="L296"></a>if err = b.write0(b.buf.Bytes()[pos : pos+c.size]); err != nil {
                    <a id="L297"></a>return
                <a id="L298"></a>}
                <a id="L299"></a>pos += c.size;
            <a id="L300"></a>}
        <a id="L301"></a>}

        <a id="L303"></a>if i+1 == b.lines.Len() {
            <a id="L304"></a><span class="comment">// last buffered line - we don&#39;t have a newline, so just write</span>
            <a id="L305"></a><span class="comment">// any outstanding buffered data</span>
            <a id="L306"></a>if err = b.write0(b.buf.Bytes()[pos : pos+b.cell.size]); err != nil {
                <a id="L307"></a>return
            <a id="L308"></a>}
            <a id="L309"></a>pos += b.cell.size;
        <a id="L310"></a>} else {
            <a id="L311"></a><span class="comment">// not the last line - write newline</span>
            <a id="L312"></a>if err = b.write0(newline); err != nil {
                <a id="L313"></a>return
            <a id="L314"></a>}
        <a id="L315"></a>}
    <a id="L316"></a>}
    <a id="L317"></a>return;
<a id="L318"></a>}


<a id="L321"></a><span class="comment">// Format the text between line0 and line1 (excluding line1); pos</span>
<a id="L322"></a><span class="comment">// is the buffer position corresponding to the beginning of line0.</span>
<a id="L323"></a><span class="comment">// Returns the buffer position corresponding to the beginning of</span>
<a id="L324"></a><span class="comment">// line1 and an error, if any.</span>
<a id="L325"></a><span class="comment">//</span>
<a id="L326"></a>func (b *Writer) format(pos0 int, line0, line1 int) (pos int, err os.Error) {
    <a id="L327"></a>pos = pos0;
    <a id="L328"></a>column := b.widths.Len();
    <a id="L329"></a>for this := line0; this &lt; line1; this++ {
        <a id="L330"></a>line := b.line(this);

        <a id="L332"></a>if column &lt; line.Len()-1 {
            <a id="L333"></a><span class="comment">// cell exists in this column =&gt; this line</span>
            <a id="L334"></a><span class="comment">// has more cells than the previous line</span>
            <a id="L335"></a><span class="comment">// (the last cell per line is ignored because cells are</span>
            <a id="L336"></a><span class="comment">// tab-terminated; the last cell per line describes the</span>
            <a id="L337"></a><span class="comment">// text before the newline/formfeed and does not belong</span>
            <a id="L338"></a><span class="comment">// to a column)</span>

            <a id="L340"></a><span class="comment">// print unprinted lines until beginning of block</span>
            <a id="L341"></a>if pos, err = b.writeLines(pos, line0, this); err != nil {
                <a id="L342"></a>return
            <a id="L343"></a>}
            <a id="L344"></a>line0 = this;

            <a id="L346"></a><span class="comment">// column block begin</span>
            <a id="L347"></a>width := b.cellwidth; <span class="comment">// minimal column width</span>
            <a id="L348"></a>discardable := true;  <span class="comment">// true if all cells in this column are empty and &#34;soft&#34;</span>
            <a id="L349"></a>for ; this &lt; line1; this++ {
                <a id="L350"></a>line = b.line(this);
                <a id="L351"></a>if column &lt; line.Len()-1 {
                    <a id="L352"></a><span class="comment">// cell exists in this column</span>
                    <a id="L353"></a>c := line.At(column).(cell);
                    <a id="L354"></a><span class="comment">// update width</span>
                    <a id="L355"></a>if w := c.width + b.padding; w &gt; width {
                        <a id="L356"></a>width = w
                    <a id="L357"></a>}
                    <a id="L358"></a><span class="comment">// update discardable</span>
                    <a id="L359"></a>if c.width &gt; 0 || c.htab {
                        <a id="L360"></a>discardable = false
                    <a id="L361"></a>}
                <a id="L362"></a>} else {
                    <a id="L363"></a>break
                <a id="L364"></a>}
            <a id="L365"></a>}
            <a id="L366"></a><span class="comment">// column block end</span>

            <a id="L368"></a><span class="comment">// discard empty columns if necessary</span>
            <a id="L369"></a>if discardable &amp;&amp; b.flags&amp;DiscardEmptyColumns != 0 {
                <a id="L370"></a>width = 0
            <a id="L371"></a>}

            <a id="L373"></a><span class="comment">// format and print all columns to the right of this column</span>
            <a id="L374"></a><span class="comment">// (we know the widths of this column and all columns to the left)</span>
            <a id="L375"></a>b.widths.Push(width);
            <a id="L376"></a>pos, err = b.format(pos, line0, this);
            <a id="L377"></a>b.widths.Pop();
            <a id="L378"></a>line0 = this;
        <a id="L379"></a>}
    <a id="L380"></a>}

    <a id="L382"></a><span class="comment">// print unprinted lines until end</span>
    <a id="L383"></a>return b.writeLines(pos, line0, line1);
<a id="L384"></a>}


<a id="L387"></a><span class="comment">// Append text to current cell.</span>
<a id="L388"></a>func (b *Writer) append(text []byte) {
    <a id="L389"></a>b.buf.Write(text);
    <a id="L390"></a>b.cell.size += len(text);
<a id="L391"></a>}


<a id="L394"></a><span class="comment">// Update the cell width.</span>
<a id="L395"></a>func (b *Writer) updateWidth() {
    <a id="L396"></a>b.cell.width += utf8.RuneCount(b.buf.Bytes()[b.pos:b.buf.Len()]);
    <a id="L397"></a>b.pos = b.buf.Len();
<a id="L398"></a>}


<a id="L401"></a><span class="comment">// To escape a text segment, bracket it with Escape characters.</span>
<a id="L402"></a><span class="comment">// For instance, the tab in this string &#34;Ignore this tab: \xff\t\xff&#34;</span>
<a id="L403"></a><span class="comment">// does not terminate a cell and constitutes a single character of</span>
<a id="L404"></a><span class="comment">// width one for formatting purposes.</span>
<a id="L405"></a><span class="comment">//</span>
<a id="L406"></a><span class="comment">// The value 0xff was chosen because it cannot appear in a valid UTF-8 sequence.</span>
<a id="L407"></a><span class="comment">//</span>
<a id="L408"></a>const Escape = &#39;\xff&#39;


<a id="L411"></a><span class="comment">// Start escaped mode.</span>
<a id="L412"></a>func (b *Writer) startEscape(ch byte) {
    <a id="L413"></a>switch ch {
    <a id="L414"></a>case Escape:
        <a id="L415"></a>b.endChar = Escape
    <a id="L416"></a>case &#39;&lt;&#39;:
        <a id="L417"></a>b.endChar = &#39;&gt;&#39;
    <a id="L418"></a>case &#39;&amp;&#39;:
        <a id="L419"></a>b.endChar = &#39;;&#39;
    <a id="L420"></a>}
<a id="L421"></a>}


<a id="L424"></a><span class="comment">// Terminate escaped mode. If the escaped text was an HTML tag, its width</span>
<a id="L425"></a><span class="comment">// is assumed to be zero for formatting purposes; if it was an HTML entity,</span>
<a id="L426"></a><span class="comment">// its width is assumed to be one. In all other cases, the width is the</span>
<a id="L427"></a><span class="comment">// unicode width of the text.</span>
<a id="L428"></a><span class="comment">//</span>
<a id="L429"></a>func (b *Writer) endEscape() {
    <a id="L430"></a>switch b.endChar {
    <a id="L431"></a>case Escape:
        <a id="L432"></a>b.updateWidth()
    <a id="L433"></a>case &#39;&gt;&#39;: <span class="comment">// tag of zero width</span>
    <a id="L434"></a>case &#39;;&#39;:
        <a id="L435"></a>b.cell.width++ <span class="comment">// entity, count as one rune</span>
    <a id="L436"></a>}
    <a id="L437"></a>b.pos = b.buf.Len();
    <a id="L438"></a>b.endChar = 0;
<a id="L439"></a>}


<a id="L442"></a><span class="comment">// Terminate the current cell by adding it to the list of cells of the</span>
<a id="L443"></a><span class="comment">// current line. Returns the number of cells in that line.</span>
<a id="L444"></a><span class="comment">//</span>
<a id="L445"></a>func (b *Writer) terminateCell(htab bool) int {
    <a id="L446"></a>b.cell.htab = htab;
    <a id="L447"></a>line := b.line(b.lines.Len() - 1);
    <a id="L448"></a>line.Push(b.cell);
    <a id="L449"></a>b.cell = cell{};
    <a id="L450"></a>return line.Len();
<a id="L451"></a>}


<a id="L454"></a><span class="comment">// Flush should be called after the last call to Write to ensure</span>
<a id="L455"></a><span class="comment">// that any data buffered in the Writer is written to output. Any</span>
<a id="L456"></a><span class="comment">// incomplete escape sequence at the end is simply considered</span>
<a id="L457"></a><span class="comment">// complete for formatting purposes.</span>
<a id="L458"></a><span class="comment">//</span>
<a id="L459"></a>func (b *Writer) Flush() os.Error {
    <a id="L460"></a><span class="comment">// add current cell if not empty</span>
    <a id="L461"></a>if b.cell.size &gt; 0 {
        <a id="L462"></a>if b.endChar != 0 {
            <a id="L463"></a><span class="comment">// inside escape - terminate it even if incomplete</span>
            <a id="L464"></a>b.endEscape()
        <a id="L465"></a>}
        <a id="L466"></a>b.terminateCell(false);
    <a id="L467"></a>}

    <a id="L469"></a><span class="comment">// format contents of buffer</span>
    <a id="L470"></a>_, err := b.format(0, 0, b.lines.Len());

    <a id="L472"></a><span class="comment">// reset, even in the presence of errors</span>
    <a id="L473"></a>b.reset();

    <a id="L475"></a>return err;
<a id="L476"></a>}


<a id="L479"></a><span class="comment">// Write writes buf to the writer b.</span>
<a id="L480"></a><span class="comment">// The only errors returned are ones encountered</span>
<a id="L481"></a><span class="comment">// while writing to the underlying output stream.</span>
<a id="L482"></a><span class="comment">//</span>
<a id="L483"></a>func (b *Writer) Write(buf []byte) (n int, err os.Error) {
    <a id="L484"></a><span class="comment">// split text into cells</span>
    <a id="L485"></a>n = 0;
    <a id="L486"></a>for i, ch := range buf {
        <a id="L487"></a>if b.endChar == 0 {
            <a id="L488"></a><span class="comment">// outside escape</span>
            <a id="L489"></a>switch ch {
            <a id="L490"></a>case &#39;\t&#39;, &#39;\v&#39;, &#39;\n&#39;, &#39;\f&#39;:
                <a id="L491"></a><span class="comment">// end of cell</span>
                <a id="L492"></a>b.append(buf[n:i]);
                <a id="L493"></a>b.updateWidth();
                <a id="L494"></a>n = i + 1; <span class="comment">// ch consumed</span>
                <a id="L495"></a>ncells := b.terminateCell(ch == &#39;\t&#39;);
                <a id="L496"></a>if ch == &#39;\n&#39; || ch == &#39;\f&#39; {
                    <a id="L497"></a><span class="comment">// terminate line</span>
                    <a id="L498"></a>b.addLine();
                    <a id="L499"></a>if ch == &#39;\f&#39; || ncells == 1 {
                        <a id="L500"></a><span class="comment">// A &#39;\f&#39; always forces a flush. Otherwise, if the previous</span>
                        <a id="L501"></a><span class="comment">// line has only one cell which does not have an impact on</span>
                        <a id="L502"></a><span class="comment">// the formatting of the following lines (the last cell per</span>
                        <a id="L503"></a><span class="comment">// line is ignored by format()), thus we can flush the</span>
                        <a id="L504"></a><span class="comment">// Writer contents.</span>
                        <a id="L505"></a>if err = b.Flush(); err != nil {
                            <a id="L506"></a>return
                        <a id="L507"></a>}
                    <a id="L508"></a>}
                <a id="L509"></a>}

            <a id="L511"></a>case Escape:
                <a id="L512"></a><span class="comment">// start of escaped sequence</span>
                <a id="L513"></a>b.append(buf[n:i]);
                <a id="L514"></a>b.updateWidth();
                <a id="L515"></a>n = i + 1; <span class="comment">// exclude Escape</span>
                <a id="L516"></a>b.startEscape(Escape);

            <a id="L518"></a>case &#39;&lt;&#39;, &#39;&amp;&#39;:
                <a id="L519"></a><span class="comment">// possibly an html tag/entity</span>
                <a id="L520"></a>if b.flags&amp;FilterHTML != 0 {
                    <a id="L521"></a><span class="comment">// begin of tag/entity</span>
                    <a id="L522"></a>b.append(buf[n:i]);
                    <a id="L523"></a>b.updateWidth();
                    <a id="L524"></a>n = i;
                    <a id="L525"></a>b.startEscape(ch);
                <a id="L526"></a>}
            <a id="L527"></a>}

        <a id="L529"></a>} else {
            <a id="L530"></a><span class="comment">// inside escape</span>
            <a id="L531"></a>if ch == b.endChar {
                <a id="L532"></a><span class="comment">// end of tag/entity</span>
                <a id="L533"></a>j := i + 1;
                <a id="L534"></a>if ch == Escape {
                    <a id="L535"></a>j = i <span class="comment">// exclude Escape</span>
                <a id="L536"></a>}
                <a id="L537"></a>b.append(buf[n:j]);
                <a id="L538"></a>n = i + 1; <span class="comment">// ch consumed</span>
                <a id="L539"></a>b.endEscape();
            <a id="L540"></a>}
        <a id="L541"></a>}
    <a id="L542"></a>}

    <a id="L544"></a><span class="comment">// append leftover text</span>
    <a id="L545"></a>b.append(buf[n:len(buf)]);
    <a id="L546"></a>n = len(buf);
    <a id="L547"></a>return;
<a id="L548"></a>}


<a id="L551"></a><span class="comment">// NewWriter allocates and initializes a new tabwriter.Writer.</span>
<a id="L552"></a><span class="comment">// The parameters are the same as for the the Init function.</span>
<a id="L553"></a><span class="comment">//</span>
<a id="L554"></a>func NewWriter(output io.Writer, cellwidth, padding int, padchar byte, flags uint) *Writer {
    <a id="L555"></a>return new(Writer).Init(output, cellwidth, padding, padchar, flags)
<a id="L556"></a>}
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
