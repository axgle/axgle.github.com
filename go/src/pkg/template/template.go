<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/template/template.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/template/template.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">/*</span>
<a id="L6"></a><span class="comment">	Data-driven templates for generating textual output such as</span>
<a id="L7"></a><span class="comment">	HTML. See</span>
<a id="L8"></a><span class="comment">		http://code.google.com/p/json-template/wiki/Reference</span>
<a id="L9"></a><span class="comment">	for full documentation of the template language. A summary:</span>

<a id="L11"></a><span class="comment">	Templates are executed by applying them to a data structure.</span>
<a id="L12"></a><span class="comment">	Annotations in the template refer to elements of the data</span>
<a id="L13"></a><span class="comment">	structure (typically a field of a struct) to control execution</span>
<a id="L14"></a><span class="comment">	and derive values to be displayed.  The template walks the</span>
<a id="L15"></a><span class="comment">	structure as it executes and the &#34;cursor&#34; @ represents the</span>
<a id="L16"></a><span class="comment">	value at the current location in the structure.</span>

<a id="L18"></a><span class="comment">	Data items may be values or pointers; the interface hides the</span>
<a id="L19"></a><span class="comment">	indirection.</span>

<a id="L21"></a><span class="comment">	Major constructs ({} are metacharacters; [] marks optional elements):</span>

<a id="L23"></a><span class="comment">		{# comment }</span>

<a id="L25"></a><span class="comment">	A one-line comment.</span>

<a id="L27"></a><span class="comment">		{.section field} XXX [ {.or} YYY ] {.end}</span>

<a id="L29"></a><span class="comment">	Set @ to the value of the field.  It may be an explicit @</span>
<a id="L30"></a><span class="comment">	to stay at the same point in the data. If the field is nil</span>
<a id="L31"></a><span class="comment">	or empty, execute YYY; otherwise execute XXX.</span>

<a id="L33"></a><span class="comment">		{.repeated section field} XXX [ {.alternates with} ZZZ ] [ {.or} YYY ] {.end}</span>

<a id="L35"></a><span class="comment">	Like .section, but field must be an array or slice.  XXX</span>
<a id="L36"></a><span class="comment">	is executed for each element.  If the array is nil or empty,</span>
<a id="L37"></a><span class="comment">	YYY is executed instead.  If the {.alternates with} marker</span>
<a id="L38"></a><span class="comment">	is present, ZZZ is executed between iterations of XXX.</span>

<a id="L40"></a><span class="comment">		{field}</span>
<a id="L41"></a><span class="comment">		{field|formatter}</span>

<a id="L43"></a><span class="comment">	Insert the value of the field into the output. Field is</span>
<a id="L44"></a><span class="comment">	first looked for in the cursor, as in .section and .repeated.</span>
<a id="L45"></a><span class="comment">	If it is not found, the search continues in outer sections</span>
<a id="L46"></a><span class="comment">	until the top level is reached.</span>

<a id="L48"></a><span class="comment">	If a formatter is specified, it must be named in the formatter</span>
<a id="L49"></a><span class="comment">	map passed to the template set up routines or in the default</span>
<a id="L50"></a><span class="comment">	set (&#34;html&#34;,&#34;str&#34;,&#34;&#34;) and is used to process the data for</span>
<a id="L51"></a><span class="comment">	output.  The formatter function has signature</span>
<a id="L52"></a><span class="comment">		func(wr io.Write, data interface{}, formatter string)</span>
<a id="L53"></a><span class="comment">	where wr is the destination for output, data is the field</span>
<a id="L54"></a><span class="comment">	value, and formatter is its name at the invocation site.</span>
<a id="L55"></a><span class="comment">*/</span>
<a id="L56"></a>package template

<a id="L58"></a>import (
    <a id="L59"></a>&#34;container/vector&#34;;
    <a id="L60"></a>&#34;fmt&#34;;
    <a id="L61"></a>&#34;io&#34;;
    <a id="L62"></a>&#34;os&#34;;
    <a id="L63"></a>&#34;reflect&#34;;
    <a id="L64"></a>&#34;runtime&#34;;
    <a id="L65"></a>&#34;strings&#34;;
<a id="L66"></a>)

<a id="L68"></a><span class="comment">// Errors returned during parsing and execution.  Users may extract the information and reformat</span>
<a id="L69"></a><span class="comment">// if they desire.</span>
<a id="L70"></a>type Error struct {
    <a id="L71"></a>Line int;
    <a id="L72"></a>Msg  string;
<a id="L73"></a>}

<a id="L75"></a>func (e *Error) String() string { return fmt.Sprintf(&#34;line %d: %s&#34;, e.Line, e.Msg) }

<a id="L77"></a><span class="comment">// Most of the literals are aces.</span>
<a id="L78"></a>var lbrace = []byte{&#39;{&#39;}
<a id="L79"></a>var rbrace = []byte{&#39;}&#39;}
<a id="L80"></a>var space = []byte{&#39; &#39;}
<a id="L81"></a>var tab = []byte{&#39;\t&#39;}

<a id="L83"></a><span class="comment">// The various types of &#34;tokens&#34;, which are plain text or (usually) brace-delimited descriptors</span>
<a id="L84"></a>const (
    <a id="L85"></a>tokAlternates = iota;
    <a id="L86"></a>tokComment;
    <a id="L87"></a>tokEnd;
    <a id="L88"></a>tokLiteral;
    <a id="L89"></a>tokOr;
    <a id="L90"></a>tokRepeated;
    <a id="L91"></a>tokSection;
    <a id="L92"></a>tokText;
    <a id="L93"></a>tokVariable;
<a id="L94"></a>)

<a id="L96"></a><span class="comment">// FormatterMap is the type describing the mapping from formatter</span>
<a id="L97"></a><span class="comment">// names to the functions that implement them.</span>
<a id="L98"></a>type FormatterMap map[string]func(io.Writer, interface{}, string)

<a id="L100"></a><span class="comment">// Built-in formatters.</span>
<a id="L101"></a>var builtins = FormatterMap{
    <a id="L102"></a>&#34;html&#34;: HTMLFormatter,
    <a id="L103"></a>&#34;str&#34;: StringFormatter,
    <a id="L104"></a>&#34;&#34;: StringFormatter,
<a id="L105"></a>}

<a id="L107"></a><span class="comment">// The parsed state of a template is a vector of xxxElement structs.</span>
<a id="L108"></a><span class="comment">// Sections have line numbers so errors can be reported better during execution.</span>

<a id="L110"></a><span class="comment">// Plain text.</span>
<a id="L111"></a>type textElement struct {
    <a id="L112"></a>text []byte;
<a id="L113"></a>}

<a id="L115"></a><span class="comment">// A literal such as .meta-left or .meta-right</span>
<a id="L116"></a>type literalElement struct {
    <a id="L117"></a>text []byte;
<a id="L118"></a>}

<a id="L120"></a><span class="comment">// A variable to be evaluated</span>
<a id="L121"></a>type variableElement struct {
    <a id="L122"></a>linenum   int;
    <a id="L123"></a>name      string;
    <a id="L124"></a>formatter string; <span class="comment">// TODO(r): implement pipelines</span>
<a id="L125"></a>}

<a id="L127"></a><span class="comment">// A .section block, possibly with a .or</span>
<a id="L128"></a>type sectionElement struct {
    <a id="L129"></a>linenum int;    <span class="comment">// of .section itself</span>
    <a id="L130"></a>field   string; <span class="comment">// cursor field for this block</span>
    <a id="L131"></a>start   int;    <span class="comment">// first element</span>
    <a id="L132"></a>or      int;    <span class="comment">// first element of .or block</span>
    <a id="L133"></a>end     int;    <span class="comment">// one beyond last element</span>
<a id="L134"></a>}

<a id="L136"></a><span class="comment">// A .repeated block, possibly with a .or and a .alternates</span>
<a id="L137"></a>type repeatedElement struct {
    <a id="L138"></a>sectionElement;      <span class="comment">// It has the same structure...</span>
    <a id="L139"></a>altstart        int; <span class="comment">// ... except for alternates</span>
    <a id="L140"></a>altend          int;
<a id="L141"></a>}

<a id="L143"></a><span class="comment">// Template is the type that represents a template definition.</span>
<a id="L144"></a><span class="comment">// It is unchanged after parsing.</span>
<a id="L145"></a>type Template struct {
    <a id="L146"></a>fmap FormatterMap; <span class="comment">// formatters for variables</span>
    <a id="L147"></a><span class="comment">// Used during parsing:</span>
    <a id="L148"></a>ldelim, rdelim []byte;   <span class="comment">// delimiters; default {}</span>
    <a id="L149"></a>buf            []byte;   <span class="comment">// input text to process</span>
    <a id="L150"></a>p              int;      <span class="comment">// position in buf</span>
    <a id="L151"></a>linenum        int;      <span class="comment">// position in input</span>
    <a id="L152"></a>error          os.Error; <span class="comment">// error during parsing (only)</span>
    <a id="L153"></a><span class="comment">// Parsed results:</span>
    <a id="L154"></a>elems *vector.Vector;
<a id="L155"></a>}

<a id="L157"></a><span class="comment">// Internal state for executing a Template.  As we evaluate the struct,</span>
<a id="L158"></a><span class="comment">// the data item descends into the fields associated with sections, etc.</span>
<a id="L159"></a><span class="comment">// Parent is used to walk upwards to find variables higher in the tree.</span>
<a id="L160"></a>type state struct {
    <a id="L161"></a>parent *state;        <span class="comment">// parent in hierarchy</span>
    <a id="L162"></a>data   reflect.Value; <span class="comment">// the driver data for this section etc.</span>
    <a id="L163"></a>wr     io.Writer;     <span class="comment">// where to send output</span>
    <a id="L164"></a>errors chan os.Error; <span class="comment">// for reporting errors during execute</span>
<a id="L165"></a>}

<a id="L167"></a>func (parent *state) clone(data reflect.Value) *state {
    <a id="L168"></a>return &amp;state{parent, data, parent.wr, parent.errors}
<a id="L169"></a>}

<a id="L171"></a><span class="comment">// New creates a new template with the specified formatter map (which</span>
<a id="L172"></a><span class="comment">// may be nil) to define auxiliary functions for formatting variables.</span>
<a id="L173"></a>func New(fmap FormatterMap) *Template {
    <a id="L174"></a>t := new(Template);
    <a id="L175"></a>t.fmap = fmap;
    <a id="L176"></a>t.ldelim = lbrace;
    <a id="L177"></a>t.rdelim = rbrace;
    <a id="L178"></a>t.elems = vector.New(0);
    <a id="L179"></a>return t;
<a id="L180"></a>}

<a id="L182"></a><span class="comment">// Report error and stop executing.  The line number must be provided explicitly.</span>
<a id="L183"></a>func (t *Template) execError(st *state, line int, err string, args ...) {
    <a id="L184"></a>st.errors &lt;- &amp;Error{line, fmt.Sprintf(err, args)};
    <a id="L185"></a>runtime.Goexit();
<a id="L186"></a>}

<a id="L188"></a><span class="comment">// Report error, save in Template to terminate parsing.</span>
<a id="L189"></a><span class="comment">// The line number comes from the template state.</span>
<a id="L190"></a>func (t *Template) parseError(err string, args ...) {
    <a id="L191"></a>t.error = &amp;Error{t.linenum, fmt.Sprintf(err, args)}
<a id="L192"></a>}

<a id="L194"></a><span class="comment">// -- Lexical analysis</span>

<a id="L196"></a><span class="comment">// Is c a white space character?</span>
<a id="L197"></a>func white(c uint8) bool { return c == &#39; &#39; || c == &#39;\t&#39; || c == &#39;\r&#39; || c == &#39;\n&#39; }

<a id="L199"></a><span class="comment">// Safely, does s[n:n+len(t)] == t?</span>
<a id="L200"></a>func equal(s []byte, n int, t []byte) bool {
    <a id="L201"></a>b := s[n:len(s)];
    <a id="L202"></a>if len(t) &gt; len(b) { <span class="comment">// not enough space left for a match.</span>
        <a id="L203"></a>return false
    <a id="L204"></a>}
    <a id="L205"></a>for i, c := range t {
        <a id="L206"></a>if c != b[i] {
            <a id="L207"></a>return false
        <a id="L208"></a>}
    <a id="L209"></a>}
    <a id="L210"></a>return true;
<a id="L211"></a>}

<a id="L213"></a><span class="comment">// nextItem returns the next item from the input buffer.  If the returned</span>
<a id="L214"></a><span class="comment">// item is empty, we are at EOF.  The item will be either a</span>
<a id="L215"></a><span class="comment">// delimited string or a non-empty string between delimited</span>
<a id="L216"></a><span class="comment">// strings. Tokens stop at (but include, if plain text) a newline.</span>
<a id="L217"></a><span class="comment">// Action tokens on a line by themselves drop the white space on</span>
<a id="L218"></a><span class="comment">// either side, up to and including the newline.</span>
<a id="L219"></a>func (t *Template) nextItem() []byte {
    <a id="L220"></a>sawLeft := false; <span class="comment">// are we waiting for an opening delimiter?</span>
    <a id="L221"></a>special := false; <span class="comment">// is this a {.foo} directive, which means trim white space?</span>
    <a id="L222"></a><span class="comment">// Delete surrounding white space if this {.foo} is the only thing on the line.</span>
    <a id="L223"></a>trim_white := t.p == 0 || t.buf[t.p-1] == &#39;\n&#39;;
    <a id="L224"></a>only_white := true; <span class="comment">// we have seen only white space so far</span>
    <a id="L225"></a>var i int;
    <a id="L226"></a>start := t.p;
<a id="L227"></a>Loop:
    <a id="L228"></a>for i = t.p; i &lt; len(t.buf); i++ {
        <a id="L229"></a>switch {
        <a id="L230"></a>case t.buf[i] == &#39;\n&#39;:
            <a id="L231"></a>t.linenum++;
            <a id="L232"></a>i++;
            <a id="L233"></a>break Loop;
        <a id="L234"></a>case white(t.buf[i]):
            <a id="L235"></a><span class="comment">// white space, do nothing</span>
        <a id="L236"></a>case !sawLeft &amp;&amp; equal(t.buf, i, t.ldelim): <span class="comment">// sawLeft checked because delims may be equal</span>
            <a id="L237"></a><span class="comment">// anything interesting already on the line?</span>
            <a id="L238"></a>if !only_white {
                <a id="L239"></a>break Loop
            <a id="L240"></a>}
            <a id="L241"></a><span class="comment">// is it a directive or comment?</span>
            <a id="L242"></a>j := i + len(t.ldelim); <span class="comment">// position after delimiter</span>
            <a id="L243"></a>if j+1 &lt; len(t.buf) &amp;&amp; (t.buf[j] == &#39;.&#39; || t.buf[j] == &#39;#&#39;) {
                <a id="L244"></a>special = true;
                <a id="L245"></a>if trim_white &amp;&amp; only_white {
                    <a id="L246"></a>start = i
                <a id="L247"></a>}
            <a id="L248"></a>} else if i &gt; t.p { <span class="comment">// have some text accumulated so stop before delimiter</span>
                <a id="L249"></a>break Loop
            <a id="L250"></a>}
            <a id="L251"></a>sawLeft = true;
            <a id="L252"></a>i = j - 1;
        <a id="L253"></a>case equal(t.buf, i, t.rdelim):
            <a id="L254"></a>if !sawLeft {
                <a id="L255"></a>t.parseError(&#34;unmatched closing delimiter&#34;);
                <a id="L256"></a>return nil;
            <a id="L257"></a>}
            <a id="L258"></a>sawLeft = false;
            <a id="L259"></a>i += len(t.rdelim);
            <a id="L260"></a>break Loop;
        <a id="L261"></a>default:
            <a id="L262"></a>only_white = false
        <a id="L263"></a>}
    <a id="L264"></a>}
    <a id="L265"></a>if sawLeft {
        <a id="L266"></a>t.parseError(&#34;unmatched opening delimiter&#34;);
        <a id="L267"></a>return nil;
    <a id="L268"></a>}
    <a id="L269"></a>item := t.buf[start:i];
    <a id="L270"></a>if special &amp;&amp; trim_white {
        <a id="L271"></a><span class="comment">// consume trailing white space</span>
        <a id="L272"></a>for ; i &lt; len(t.buf) &amp;&amp; white(t.buf[i]); i++ {
            <a id="L273"></a>if t.buf[i] == &#39;\n&#39; {
                <a id="L274"></a>t.linenum++;
                <a id="L275"></a>i++;
                <a id="L276"></a>break; <span class="comment">// stop after newline</span>
            <a id="L277"></a>}
        <a id="L278"></a>}
    <a id="L279"></a>}
    <a id="L280"></a>t.p = i;
    <a id="L281"></a>return item;
<a id="L282"></a>}

<a id="L284"></a><span class="comment">// Turn a byte array into a white-space-split array of strings.</span>
<a id="L285"></a>func words(buf []byte) []string {
    <a id="L286"></a>s := make([]string, 0, 5);
    <a id="L287"></a>p := 0; <span class="comment">// position in buf</span>
    <a id="L288"></a><span class="comment">// one word per loop</span>
    <a id="L289"></a>for i := 0; ; i++ {
        <a id="L290"></a><span class="comment">// skip white space</span>
        <a id="L291"></a>for ; p &lt; len(buf) &amp;&amp; white(buf[p]); p++ {
        <a id="L292"></a>}
        <a id="L293"></a><span class="comment">// grab word</span>
        <a id="L294"></a>start := p;
        <a id="L295"></a>for ; p &lt; len(buf) &amp;&amp; !white(buf[p]); p++ {
        <a id="L296"></a>}
        <a id="L297"></a>if start == p { <span class="comment">// no text left</span>
            <a id="L298"></a>break
        <a id="L299"></a>}
        <a id="L300"></a>if i == cap(s) {
            <a id="L301"></a>ns := make([]string, 2*cap(s));
            <a id="L302"></a>for j := range s {
                <a id="L303"></a>ns[j] = s[j]
            <a id="L304"></a>}
            <a id="L305"></a>s = ns;
        <a id="L306"></a>}
        <a id="L307"></a>s = s[0 : i+1];
        <a id="L308"></a>s[i] = string(buf[start:p]);
    <a id="L309"></a>}
    <a id="L310"></a>return s;
<a id="L311"></a>}

<a id="L313"></a><span class="comment">// Analyze an item and return its token type and, if it&#39;s an action item, an array of</span>
<a id="L314"></a><span class="comment">// its constituent words.</span>
<a id="L315"></a>func (t *Template) analyze(item []byte) (tok int, w []string) {
    <a id="L316"></a><span class="comment">// item is known to be non-empty</span>
    <a id="L317"></a>if !equal(item, 0, t.ldelim) { <span class="comment">// doesn&#39;t start with left delimiter</span>
        <a id="L318"></a>tok = tokText;
        <a id="L319"></a>return;
    <a id="L320"></a>}
    <a id="L321"></a>if !equal(item, len(item)-len(t.rdelim), t.rdelim) { <span class="comment">// doesn&#39;t end with right delimiter</span>
        <a id="L322"></a>t.parseError(&#34;internal error: unmatched opening delimiter&#34;); <span class="comment">// lexing should prevent this</span>
        <a id="L323"></a>return;
    <a id="L324"></a>}
    <a id="L325"></a>if len(item) &lt;= len(t.ldelim)+len(t.rdelim) { <span class="comment">// no contents</span>
        <a id="L326"></a>t.parseError(&#34;empty directive&#34;);
        <a id="L327"></a>return;
    <a id="L328"></a>}
    <a id="L329"></a><span class="comment">// Comment</span>
    <a id="L330"></a>if item[len(t.ldelim)] == &#39;#&#39; {
        <a id="L331"></a>tok = tokComment;
        <a id="L332"></a>return;
    <a id="L333"></a>}
    <a id="L334"></a><span class="comment">// Split into words</span>
    <a id="L335"></a>w = words(item[len(t.ldelim) : len(item)-len(t.rdelim)]); <span class="comment">// drop final delimiter</span>
    <a id="L336"></a>if len(w) == 0 {
        <a id="L337"></a>t.parseError(&#34;empty directive&#34;);
        <a id="L338"></a>return;
    <a id="L339"></a>}
    <a id="L340"></a>if len(w) == 1 &amp;&amp; w[0][0] != &#39;.&#39; {
        <a id="L341"></a>tok = tokVariable;
        <a id="L342"></a>return;
    <a id="L343"></a>}
    <a id="L344"></a>switch w[0] {
    <a id="L345"></a>case &#34;.meta-left&#34;, &#34;.meta-right&#34;, &#34;.space&#34;, &#34;.tab&#34;:
        <a id="L346"></a>tok = tokLiteral;
        <a id="L347"></a>return;
    <a id="L348"></a>case &#34;.or&#34;:
        <a id="L349"></a>tok = tokOr;
        <a id="L350"></a>return;
    <a id="L351"></a>case &#34;.end&#34;:
        <a id="L352"></a>tok = tokEnd;
        <a id="L353"></a>return;
    <a id="L354"></a>case &#34;.section&#34;:
        <a id="L355"></a>if len(w) != 2 {
            <a id="L356"></a>t.parseError(&#34;incorrect fields for .section: %s&#34;, item);
            <a id="L357"></a>return;
        <a id="L358"></a>}
        <a id="L359"></a>tok = tokSection;
        <a id="L360"></a>return;
    <a id="L361"></a>case &#34;.repeated&#34;:
        <a id="L362"></a>if len(w) != 3 || w[1] != &#34;section&#34; {
            <a id="L363"></a>t.parseError(&#34;incorrect fields for .repeated: %s&#34;, item);
            <a id="L364"></a>return;
        <a id="L365"></a>}
        <a id="L366"></a>tok = tokRepeated;
        <a id="L367"></a>return;
    <a id="L368"></a>case &#34;.alternates&#34;:
        <a id="L369"></a>if len(w) != 2 || w[1] != &#34;with&#34; {
            <a id="L370"></a>t.parseError(&#34;incorrect fields for .alternates: %s&#34;, item);
            <a id="L371"></a>return;
        <a id="L372"></a>}
        <a id="L373"></a>tok = tokAlternates;
        <a id="L374"></a>return;
    <a id="L375"></a>}
    <a id="L376"></a>t.parseError(&#34;bad directive: %s&#34;, item);
    <a id="L377"></a>return;
<a id="L378"></a>}

<a id="L380"></a><span class="comment">// -- Parsing</span>

<a id="L382"></a><span class="comment">// Allocate a new variable-evaluation element.</span>
<a id="L383"></a>func (t *Template) newVariable(name_formatter string) (v *variableElement) {
    <a id="L384"></a>name := name_formatter;
    <a id="L385"></a>formatter := &#34;&#34;;
    <a id="L386"></a>bar := strings.Index(name_formatter, &#34;|&#34;);
    <a id="L387"></a>if bar &gt;= 0 {
        <a id="L388"></a>name = name_formatter[0:bar];
        <a id="L389"></a>formatter = name_formatter[bar+1 : len(name_formatter)];
    <a id="L390"></a>}
    <a id="L391"></a><span class="comment">// Probably ok, so let&#39;s build it.</span>
    <a id="L392"></a>v = &amp;variableElement{t.linenum, name, formatter};

    <a id="L394"></a><span class="comment">// We could remember the function address here and avoid the lookup later,</span>
    <a id="L395"></a><span class="comment">// but it&#39;s more dynamic to let the user change the map contents underfoot.</span>
    <a id="L396"></a><span class="comment">// We do require the name to be present, though.</span>

    <a id="L398"></a><span class="comment">// Is it in user-supplied map?</span>
    <a id="L399"></a>if t.fmap != nil {
        <a id="L400"></a>if _, ok := t.fmap[formatter]; ok {
            <a id="L401"></a>return
        <a id="L402"></a>}
    <a id="L403"></a>}
    <a id="L404"></a><span class="comment">// Is it in builtin map?</span>
    <a id="L405"></a>if _, ok := builtins[formatter]; ok {
        <a id="L406"></a>return
    <a id="L407"></a>}
    <a id="L408"></a>t.parseError(&#34;unknown formatter: %s&#34;, formatter);
    <a id="L409"></a>return;
<a id="L410"></a>}

<a id="L412"></a><span class="comment">// Grab the next item.  If it&#39;s simple, just append it to the template.</span>
<a id="L413"></a><span class="comment">// Otherwise return its details.</span>
<a id="L414"></a>func (t *Template) parseSimple(item []byte) (done bool, tok int, w []string) {
    <a id="L415"></a>tok, w = t.analyze(item);
    <a id="L416"></a>if t.error != nil {
        <a id="L417"></a>return
    <a id="L418"></a>}
    <a id="L419"></a>done = true; <span class="comment">// assume for simplicity</span>
    <a id="L420"></a>switch tok {
    <a id="L421"></a>case tokComment:
        <a id="L422"></a>return
    <a id="L423"></a>case tokText:
        <a id="L424"></a>t.elems.Push(&amp;textElement{item});
        <a id="L425"></a>return;
    <a id="L426"></a>case tokLiteral:
        <a id="L427"></a>switch w[0] {
        <a id="L428"></a>case &#34;.meta-left&#34;:
            <a id="L429"></a>t.elems.Push(&amp;literalElement{t.ldelim})
        <a id="L430"></a>case &#34;.meta-right&#34;:
            <a id="L431"></a>t.elems.Push(&amp;literalElement{t.rdelim})
        <a id="L432"></a>case &#34;.space&#34;:
            <a id="L433"></a>t.elems.Push(&amp;literalElement{space})
        <a id="L434"></a>case &#34;.tab&#34;:
            <a id="L435"></a>t.elems.Push(&amp;literalElement{tab})
        <a id="L436"></a>default:
            <a id="L437"></a>t.parseError(&#34;internal error: unknown literal: %s&#34;, w[0]);
            <a id="L438"></a>return;
        <a id="L439"></a>}
        <a id="L440"></a>return;
    <a id="L441"></a>case tokVariable:
        <a id="L442"></a>t.elems.Push(t.newVariable(w[0]));
        <a id="L443"></a>return;
    <a id="L444"></a>}
    <a id="L445"></a>return false, tok, w;
<a id="L446"></a>}

<a id="L448"></a><span class="comment">// parseRepeated and parseSection are mutually recursive</span>

<a id="L450"></a>func (t *Template) parseRepeated(words []string) *repeatedElement {
    <a id="L451"></a>r := new(repeatedElement);
    <a id="L452"></a>t.elems.Push(r);
    <a id="L453"></a>r.linenum = t.linenum;
    <a id="L454"></a>r.field = words[2];
    <a id="L455"></a><span class="comment">// Scan section, collecting true and false (.or) blocks.</span>
    <a id="L456"></a>r.start = t.elems.Len();
    <a id="L457"></a>r.or = -1;
    <a id="L458"></a>r.altstart = -1;
    <a id="L459"></a>r.altend = -1;
<a id="L460"></a>Loop:
    <a id="L461"></a>for t.error == nil {
        <a id="L462"></a>item := t.nextItem();
        <a id="L463"></a>if t.error != nil {
            <a id="L464"></a>break
        <a id="L465"></a>}
        <a id="L466"></a>if len(item) == 0 {
            <a id="L467"></a>t.parseError(&#34;missing .end for .repeated section&#34;);
            <a id="L468"></a>break;
        <a id="L469"></a>}
        <a id="L470"></a>done, tok, w := t.parseSimple(item);
        <a id="L471"></a>if t.error != nil {
            <a id="L472"></a>break
        <a id="L473"></a>}
        <a id="L474"></a>if done {
            <a id="L475"></a>continue
        <a id="L476"></a>}
        <a id="L477"></a>switch tok {
        <a id="L478"></a>case tokEnd:
            <a id="L479"></a>break Loop
        <a id="L480"></a>case tokOr:
            <a id="L481"></a>if r.or &gt;= 0 {
                <a id="L482"></a>t.parseError(&#34;extra .or in .repeated section&#34;);
                <a id="L483"></a>break Loop;
            <a id="L484"></a>}
            <a id="L485"></a>r.altend = t.elems.Len();
            <a id="L486"></a>r.or = t.elems.Len();
        <a id="L487"></a>case tokSection:
            <a id="L488"></a>t.parseSection(w)
        <a id="L489"></a>case tokRepeated:
            <a id="L490"></a>t.parseRepeated(w)
        <a id="L491"></a>case tokAlternates:
            <a id="L492"></a>if r.altstart &gt;= 0 {
                <a id="L493"></a>t.parseError(&#34;extra .alternates in .repeated section&#34;);
                <a id="L494"></a>break Loop;
            <a id="L495"></a>}
            <a id="L496"></a>if r.or &gt;= 0 {
                <a id="L497"></a>t.parseError(&#34;.alternates inside .or block in .repeated section&#34;);
                <a id="L498"></a>break Loop;
            <a id="L499"></a>}
            <a id="L500"></a>r.altstart = t.elems.Len();
        <a id="L501"></a>default:
            <a id="L502"></a>t.parseError(&#34;internal error: unknown repeated section item: %s&#34;, item);
            <a id="L503"></a>break Loop;
        <a id="L504"></a>}
    <a id="L505"></a>}
    <a id="L506"></a>if t.error != nil {
        <a id="L507"></a>return nil
    <a id="L508"></a>}
    <a id="L509"></a>if r.altend &lt; 0 {
        <a id="L510"></a>r.altend = t.elems.Len()
    <a id="L511"></a>}
    <a id="L512"></a>r.end = t.elems.Len();
    <a id="L513"></a>return r;
<a id="L514"></a>}

<a id="L516"></a>func (t *Template) parseSection(words []string) *sectionElement {
    <a id="L517"></a>s := new(sectionElement);
    <a id="L518"></a>t.elems.Push(s);
    <a id="L519"></a>s.linenum = t.linenum;
    <a id="L520"></a>s.field = words[1];
    <a id="L521"></a><span class="comment">// Scan section, collecting true and false (.or) blocks.</span>
    <a id="L522"></a>s.start = t.elems.Len();
    <a id="L523"></a>s.or = -1;
<a id="L524"></a>Loop:
    <a id="L525"></a>for t.error == nil {
        <a id="L526"></a>item := t.nextItem();
        <a id="L527"></a>if t.error != nil {
            <a id="L528"></a>break
        <a id="L529"></a>}
        <a id="L530"></a>if len(item) == 0 {
            <a id="L531"></a>t.parseError(&#34;missing .end for .section&#34;);
            <a id="L532"></a>break;
        <a id="L533"></a>}
        <a id="L534"></a>done, tok, w := t.parseSimple(item);
        <a id="L535"></a>if t.error != nil {
            <a id="L536"></a>break
        <a id="L537"></a>}
        <a id="L538"></a>if done {
            <a id="L539"></a>continue
        <a id="L540"></a>}
        <a id="L541"></a>switch tok {
        <a id="L542"></a>case tokEnd:
            <a id="L543"></a>break Loop
        <a id="L544"></a>case tokOr:
            <a id="L545"></a>if s.or &gt;= 0 {
                <a id="L546"></a>t.parseError(&#34;extra .or in .section&#34;);
                <a id="L547"></a>break Loop;
            <a id="L548"></a>}
            <a id="L549"></a>s.or = t.elems.Len();
        <a id="L550"></a>case tokSection:
            <a id="L551"></a>t.parseSection(w)
        <a id="L552"></a>case tokRepeated:
            <a id="L553"></a>t.parseRepeated(w)
        <a id="L554"></a>case tokAlternates:
            <a id="L555"></a>t.parseError(&#34;.alternates not in .repeated&#34;)
        <a id="L556"></a>default:
            <a id="L557"></a>t.parseError(&#34;internal error: unknown section item: %s&#34;, item)
        <a id="L558"></a>}
    <a id="L559"></a>}
    <a id="L560"></a>if t.error != nil {
        <a id="L561"></a>return nil
    <a id="L562"></a>}
    <a id="L563"></a>s.end = t.elems.Len();
    <a id="L564"></a>return s;
<a id="L565"></a>}

<a id="L567"></a>func (t *Template) parse() {
    <a id="L568"></a>for t.error == nil {
        <a id="L569"></a>item := t.nextItem();
        <a id="L570"></a>if t.error != nil {
            <a id="L571"></a>break
        <a id="L572"></a>}
        <a id="L573"></a>if len(item) == 0 {
            <a id="L574"></a>break
        <a id="L575"></a>}
        <a id="L576"></a>done, tok, w := t.parseSimple(item);
        <a id="L577"></a>if done {
            <a id="L578"></a>continue
        <a id="L579"></a>}
        <a id="L580"></a>switch tok {
        <a id="L581"></a>case tokOr, tokEnd, tokAlternates:
            <a id="L582"></a>t.parseError(&#34;unexpected %s&#34;, w[0])
        <a id="L583"></a>case tokSection:
            <a id="L584"></a>t.parseSection(w)
        <a id="L585"></a>case tokRepeated:
            <a id="L586"></a>t.parseRepeated(w)
        <a id="L587"></a>default:
            <a id="L588"></a>t.parseError(&#34;internal error: bad directive in parse: %s&#34;, item)
        <a id="L589"></a>}
    <a id="L590"></a>}
<a id="L591"></a>}

<a id="L593"></a><span class="comment">// -- Execution</span>

<a id="L595"></a><span class="comment">// If the data for this template is a struct, find the named variable.</span>
<a id="L596"></a><span class="comment">// Names of the form a.b.c are walked down the data tree.</span>
<a id="L597"></a><span class="comment">// The special name &#34;@&#34; (the &#34;cursor&#34;) denotes the current data.</span>
<a id="L598"></a><span class="comment">// The value coming in (st.data) might need indirecting to reach</span>
<a id="L599"></a><span class="comment">// a struct while the return value is not indirected - that is,</span>
<a id="L600"></a><span class="comment">// it represents the actual named field.</span>
<a id="L601"></a>func (st *state) findVar(s string) reflect.Value {
    <a id="L602"></a>if s == &#34;@&#34; {
        <a id="L603"></a>return st.data
    <a id="L604"></a>}
    <a id="L605"></a>data := st.data;
    <a id="L606"></a>elems := strings.Split(s, &#34;.&#34;, 0);
    <a id="L607"></a>for i := 0; i &lt; len(elems); i++ {
        <a id="L608"></a><span class="comment">// Look up field; data must be a struct.</span>
        <a id="L609"></a>data = reflect.Indirect(data);
        <a id="L610"></a>if data == nil {
            <a id="L611"></a>return nil
        <a id="L612"></a>}
        <a id="L613"></a>typ, ok := data.Type().(*reflect.StructType);
        <a id="L614"></a>if !ok {
            <a id="L615"></a>return nil
        <a id="L616"></a>}
        <a id="L617"></a>field, ok := typ.FieldByName(elems[i]);
        <a id="L618"></a>if !ok {
            <a id="L619"></a>return nil
        <a id="L620"></a>}
        <a id="L621"></a>data = data.(*reflect.StructValue).FieldByIndex(field.Index);
    <a id="L622"></a>}
    <a id="L623"></a>return data;
<a id="L624"></a>}

<a id="L626"></a><span class="comment">// Is there no data to look at?</span>
<a id="L627"></a>func empty(v reflect.Value) bool {
    <a id="L628"></a>v = reflect.Indirect(v);
    <a id="L629"></a>if v == nil {
        <a id="L630"></a>return true
    <a id="L631"></a>}
    <a id="L632"></a>switch v := v.(type) {
    <a id="L633"></a>case *reflect.BoolValue:
        <a id="L634"></a>return v.Get() == false
    <a id="L635"></a>case *reflect.StringValue:
        <a id="L636"></a>return v.Get() == &#34;&#34;
    <a id="L637"></a>case *reflect.StructValue:
        <a id="L638"></a>return false
    <a id="L639"></a>case *reflect.ArrayValue:
        <a id="L640"></a>return v.Len() == 0
    <a id="L641"></a>case *reflect.SliceValue:
        <a id="L642"></a>return v.Len() == 0
    <a id="L643"></a>}
    <a id="L644"></a>return true;
<a id="L645"></a>}

<a id="L647"></a><span class="comment">// Look up a variable, up through the parent if necessary.</span>
<a id="L648"></a>func (t *Template) varValue(name string, st *state) reflect.Value {
    <a id="L649"></a>field := st.findVar(name);
    <a id="L650"></a>if field == nil {
        <a id="L651"></a>if st.parent == nil {
            <a id="L652"></a>t.execError(st, t.linenum, &#34;name not found: %s&#34;, name)
        <a id="L653"></a>}
        <a id="L654"></a>return t.varValue(name, st.parent);
    <a id="L655"></a>}
    <a id="L656"></a>return field;
<a id="L657"></a>}

<a id="L659"></a><span class="comment">// Evaluate a variable, looking up through the parent if necessary.</span>
<a id="L660"></a><span class="comment">// If it has a formatter attached ({var|formatter}) run that too.</span>
<a id="L661"></a>func (t *Template) writeVariable(v *variableElement, st *state) {
    <a id="L662"></a>formatter := v.formatter;
    <a id="L663"></a>val := t.varValue(v.name, st).Interface();
    <a id="L664"></a><span class="comment">// is it in user-supplied map?</span>
    <a id="L665"></a>if t.fmap != nil {
        <a id="L666"></a>if fn, ok := t.fmap[formatter]; ok {
            <a id="L667"></a>fn(st.wr, val, formatter);
            <a id="L668"></a>return;
        <a id="L669"></a>}
    <a id="L670"></a>}
    <a id="L671"></a><span class="comment">// is it in builtin map?</span>
    <a id="L672"></a>if fn, ok := builtins[formatter]; ok {
        <a id="L673"></a>fn(st.wr, val, formatter);
        <a id="L674"></a>return;
    <a id="L675"></a>}
    <a id="L676"></a>t.execError(st, v.linenum, &#34;missing formatter %s for variable %s&#34;, formatter, v.name);
<a id="L677"></a>}

<a id="L679"></a><span class="comment">// Execute element i.  Return next index to execute.</span>
<a id="L680"></a>func (t *Template) executeElement(i int, st *state) int {
    <a id="L681"></a>switch elem := t.elems.At(i).(type) {
    <a id="L682"></a>case *textElement:
        <a id="L683"></a>st.wr.Write(elem.text);
        <a id="L684"></a>return i + 1;
    <a id="L685"></a>case *literalElement:
        <a id="L686"></a>st.wr.Write(elem.text);
        <a id="L687"></a>return i + 1;
    <a id="L688"></a>case *variableElement:
        <a id="L689"></a>t.writeVariable(elem, st);
        <a id="L690"></a>return i + 1;
    <a id="L691"></a>case *sectionElement:
        <a id="L692"></a>t.executeSection(elem, st);
        <a id="L693"></a>return elem.end;
    <a id="L694"></a>case *repeatedElement:
        <a id="L695"></a>t.executeRepeated(elem, st);
        <a id="L696"></a>return elem.end;
    <a id="L697"></a>}
    <a id="L698"></a>e := t.elems.At(i);
    <a id="L699"></a>t.execError(st, 0, &#34;internal error: bad directive in execute: %v %T\n&#34;, reflect.NewValue(e).Interface(), e);
    <a id="L700"></a>return 0;
<a id="L701"></a>}

<a id="L703"></a><span class="comment">// Execute the template.</span>
<a id="L704"></a>func (t *Template) execute(start, end int, st *state) {
    <a id="L705"></a>for i := start; i &lt; end; {
        <a id="L706"></a>i = t.executeElement(i, st)
    <a id="L707"></a>}
<a id="L708"></a>}

<a id="L710"></a><span class="comment">// Execute a .section</span>
<a id="L711"></a>func (t *Template) executeSection(s *sectionElement, st *state) {
    <a id="L712"></a><span class="comment">// Find driver data for this section.  It must be in the current struct.</span>
    <a id="L713"></a>field := t.varValue(s.field, st);
    <a id="L714"></a>if field == nil {
        <a id="L715"></a>t.execError(st, s.linenum, &#34;.section: cannot find field %s in %s&#34;, s.field, reflect.Indirect(st.data).Type())
    <a id="L716"></a>}
    <a id="L717"></a>st = st.clone(field);
    <a id="L718"></a>start, end := s.start, s.or;
    <a id="L719"></a>if !empty(field) {
        <a id="L720"></a><span class="comment">// Execute the normal block.</span>
        <a id="L721"></a>if end &lt; 0 {
            <a id="L722"></a>end = s.end
        <a id="L723"></a>}
    <a id="L724"></a>} else {
        <a id="L725"></a><span class="comment">// Execute the .or block.  If it&#39;s missing, do nothing.</span>
        <a id="L726"></a>start, end = s.or, s.end;
        <a id="L727"></a>if start &lt; 0 {
            <a id="L728"></a>return
        <a id="L729"></a>}
    <a id="L730"></a>}
    <a id="L731"></a>for i := start; i &lt; end; {
        <a id="L732"></a>i = t.executeElement(i, st)
    <a id="L733"></a>}
<a id="L734"></a>}

<a id="L736"></a><span class="comment">// Return the result of calling the Iter method on v, or nil.</span>
<a id="L737"></a>func iter(v reflect.Value) *reflect.ChanValue {
    <a id="L738"></a>for j := 0; j &lt; v.Type().NumMethod(); j++ {
        <a id="L739"></a>mth := v.Type().Method(j);
        <a id="L740"></a>fv := v.Method(j);
        <a id="L741"></a>ft := fv.Type().(*reflect.FuncType);
        <a id="L742"></a><span class="comment">// TODO(rsc): NumIn() should return 0 here, because ft is from a curried FuncValue.</span>
        <a id="L743"></a>if mth.Name != &#34;Iter&#34; || ft.NumIn() != 1 || ft.NumOut() != 1 {
            <a id="L744"></a>continue
        <a id="L745"></a>}
        <a id="L746"></a>ct, ok := ft.Out(0).(*reflect.ChanType);
        <a id="L747"></a>if !ok || ct.Dir()&amp;reflect.RecvDir == 0 {
            <a id="L748"></a>continue
        <a id="L749"></a>}
        <a id="L750"></a>return fv.Call(nil)[0].(*reflect.ChanValue);
    <a id="L751"></a>}
    <a id="L752"></a>return nil;
<a id="L753"></a>}

<a id="L755"></a><span class="comment">// Execute a .repeated section</span>
<a id="L756"></a>func (t *Template) executeRepeated(r *repeatedElement, st *state) {
    <a id="L757"></a><span class="comment">// Find driver data for this section.  It must be in the current struct.</span>
    <a id="L758"></a>field := t.varValue(r.field, st);
    <a id="L759"></a>if field == nil {
        <a id="L760"></a>t.execError(st, r.linenum, &#34;.repeated: cannot find field %s in %s&#34;, r.field, reflect.Indirect(st.data).Type())
    <a id="L761"></a>}

    <a id="L763"></a>start, end := r.start, r.or;
    <a id="L764"></a>if end &lt; 0 {
        <a id="L765"></a>end = r.end
    <a id="L766"></a>}
    <a id="L767"></a>if r.altstart &gt;= 0 {
        <a id="L768"></a>end = r.altstart
    <a id="L769"></a>}
    <a id="L770"></a>first := true;

    <a id="L772"></a>if array, ok := field.(reflect.ArrayOrSliceValue); ok {
        <a id="L773"></a>for j := 0; j &lt; array.Len(); j++ {
            <a id="L774"></a>newst := st.clone(array.Elem(j));

            <a id="L776"></a><span class="comment">// .alternates between elements</span>
            <a id="L777"></a>if !first &amp;&amp; r.altstart &gt;= 0 {
                <a id="L778"></a>for i := r.altstart; i &lt; r.altend; {
                    <a id="L779"></a>i = t.executeElement(i, newst)
                <a id="L780"></a>}
            <a id="L781"></a>}
            <a id="L782"></a>first = false;

            <a id="L784"></a>for i := start; i &lt; end; {
                <a id="L785"></a>i = t.executeElement(i, newst)
            <a id="L786"></a>}
        <a id="L787"></a>}
    <a id="L788"></a>} else if ch := iter(field); ch != nil {
        <a id="L789"></a>for {
            <a id="L790"></a>e := ch.Recv();
            <a id="L791"></a>if ch.Closed() {
                <a id="L792"></a>break
            <a id="L793"></a>}
            <a id="L794"></a>newst := st.clone(e);

            <a id="L796"></a><span class="comment">// .alternates between elements</span>
            <a id="L797"></a>if !first &amp;&amp; r.altstart &gt;= 0 {
                <a id="L798"></a>for i := r.altstart; i &lt; r.altend; {
                    <a id="L799"></a>i = t.executeElement(i, newst)
                <a id="L800"></a>}
            <a id="L801"></a>}
            <a id="L802"></a>first = false;

            <a id="L804"></a>for i := start; i &lt; end; {
                <a id="L805"></a>i = t.executeElement(i, newst)
            <a id="L806"></a>}
        <a id="L807"></a>}
    <a id="L808"></a>} else {
        <a id="L809"></a>t.execError(st, r.linenum, &#34;.repeated: cannot repeat %s (type %s)&#34;,
            <a id="L810"></a>r.field, field.Type())
    <a id="L811"></a>}

    <a id="L813"></a>if first {
        <a id="L814"></a><span class="comment">// Empty. Execute the .or block, once.  If it&#39;s missing, do nothing.</span>
        <a id="L815"></a>start, end := r.or, r.end;
        <a id="L816"></a>if start &gt;= 0 {
            <a id="L817"></a>newst := st.clone(field);
            <a id="L818"></a>for i := start; i &lt; end; {
                <a id="L819"></a>i = t.executeElement(i, newst)
            <a id="L820"></a>}
        <a id="L821"></a>}
        <a id="L822"></a>return;
    <a id="L823"></a>}
<a id="L824"></a>}

<a id="L826"></a><span class="comment">// A valid delimiter must contain no white space and be non-empty.</span>
<a id="L827"></a>func validDelim(d []byte) bool {
    <a id="L828"></a>if len(d) == 0 {
        <a id="L829"></a>return false
    <a id="L830"></a>}
    <a id="L831"></a>for _, c := range d {
        <a id="L832"></a>if white(c) {
            <a id="L833"></a>return false
        <a id="L834"></a>}
    <a id="L835"></a>}
    <a id="L836"></a>return true;
<a id="L837"></a>}

<a id="L839"></a><span class="comment">// -- Public interface</span>

<a id="L841"></a><span class="comment">// Parse initializes a Template by parsing its definition.  The string</span>
<a id="L842"></a><span class="comment">// s contains the template text.  If any errors occur, Parse returns</span>
<a id="L843"></a><span class="comment">// the error.</span>
<a id="L844"></a>func (t *Template) Parse(s string) os.Error {
    <a id="L845"></a>if !validDelim(t.ldelim) || !validDelim(t.rdelim) {
        <a id="L846"></a>return &amp;Error{1, fmt.Sprintf(&#34;bad delimiter strings %q %q&#34;, t.ldelim, t.rdelim)}
    <a id="L847"></a>}
    <a id="L848"></a>t.buf = strings.Bytes(s);
    <a id="L849"></a>t.p = 0;
    <a id="L850"></a>t.linenum = 1;
    <a id="L851"></a>t.parse();
    <a id="L852"></a>return t.error;
<a id="L853"></a>}

<a id="L855"></a><span class="comment">// Execute applies a parsed template to the specified data object,</span>
<a id="L856"></a><span class="comment">// generating output to wr.</span>
<a id="L857"></a>func (t *Template) Execute(data interface{}, wr io.Writer) os.Error {
    <a id="L858"></a><span class="comment">// Extract the driver data.</span>
    <a id="L859"></a>val := reflect.NewValue(data);
    <a id="L860"></a>errors := make(chan os.Error);
    <a id="L861"></a>go func() {
        <a id="L862"></a>t.p = 0;
        <a id="L863"></a>t.execute(0, t.elems.Len(), &amp;state{nil, val, wr, errors});
        <a id="L864"></a>errors &lt;- nil; <span class="comment">// clean return;</span>
    <a id="L865"></a>}();
    <a id="L866"></a>return &lt;-errors;
<a id="L867"></a>}

<a id="L869"></a><span class="comment">// SetDelims sets the left and right delimiters for operations in the</span>
<a id="L870"></a><span class="comment">// template.  They are validated during parsing.  They could be</span>
<a id="L871"></a><span class="comment">// validated here but it&#39;s better to keep the routine simple.  The</span>
<a id="L872"></a><span class="comment">// delimiters are very rarely invalid and Parse has the necessary</span>
<a id="L873"></a><span class="comment">// error-handling interface already.</span>
<a id="L874"></a>func (t *Template) SetDelims(left, right string) {
    <a id="L875"></a>t.ldelim = strings.Bytes(left);
    <a id="L876"></a>t.rdelim = strings.Bytes(right);
<a id="L877"></a>}

<a id="L879"></a><span class="comment">// Parse creates a Template with default parameters (such as {} for</span>
<a id="L880"></a><span class="comment">// metacharacters).  The string s contains the template text while</span>
<a id="L881"></a><span class="comment">// the formatter map fmap, which may be nil, defines auxiliary functions</span>
<a id="L882"></a><span class="comment">// for formatting variables.  The template is returned. If any errors</span>
<a id="L883"></a><span class="comment">// occur, err will be non-nil.</span>
<a id="L884"></a>func Parse(s string, fmap FormatterMap) (t *Template, err os.Error) {
    <a id="L885"></a>t = New(fmap);
    <a id="L886"></a>err = t.Parse(s);
    <a id="L887"></a>if err != nil {
        <a id="L888"></a>t = nil
    <a id="L889"></a>}
    <a id="L890"></a>return;
<a id="L891"></a>}

<a id="L893"></a><span class="comment">// MustParse is like Parse but panics if the template cannot be parsed.</span>
<a id="L894"></a>func MustParse(s string, fmap FormatterMap) *Template {
    <a id="L895"></a>t, err := Parse(s, fmap);
    <a id="L896"></a>if err != nil {
        <a id="L897"></a>panic(&#34;template parse error: &#34;, err.String())
    <a id="L898"></a>}
    <a id="L899"></a>return t;
<a id="L900"></a>}
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
