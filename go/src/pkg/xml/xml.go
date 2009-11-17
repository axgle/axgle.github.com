<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/xml/xml.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/xml/xml.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Package xml implements a simple XML 1.0 parser that</span>
<a id="L6"></a><span class="comment">// understands XML name spaces.</span>
<a id="L7"></a>package xml

<a id="L9"></a><span class="comment">// References:</span>
<a id="L10"></a><span class="comment">//    Annotated XML spec: http://www.xml.com/axml/testaxml.htm</span>
<a id="L11"></a><span class="comment">//    XML name spaces: http://www.w3.org/TR/REC-xml-names/</span>

<a id="L13"></a><span class="comment">// TODO(rsc):</span>
<a id="L14"></a><span class="comment">//	Test error handling.</span>
<a id="L15"></a><span class="comment">//	Expose parser line number in errors.</span>

<a id="L17"></a>import (
    <a id="L18"></a>&#34;bufio&#34;;
    <a id="L19"></a>&#34;bytes&#34;;
    <a id="L20"></a>&#34;io&#34;;
    <a id="L21"></a>&#34;os&#34;;
    <a id="L22"></a>&#34;strconv&#34;;
    <a id="L23"></a>&#34;strings&#34;;
    <a id="L24"></a>&#34;unicode&#34;;
    <a id="L25"></a>&#34;utf8&#34;;
<a id="L26"></a>)

<a id="L28"></a><span class="comment">// A SyntaxError represents a syntax error in the XML input stream.</span>
<a id="L29"></a>type SyntaxError string

<a id="L31"></a>func (e SyntaxError) String() string { return &#34;XML syntax error: &#34; + string(e) }

<a id="L33"></a><span class="comment">// A Name represents an XML name (Local) annotated</span>
<a id="L34"></a><span class="comment">// with a name space identifier (Space).</span>
<a id="L35"></a><span class="comment">// In tokens returned by Parser.Token, the Space identifier</span>
<a id="L36"></a><span class="comment">// is given as a canonical URL, not the short prefix used</span>
<a id="L37"></a><span class="comment">// in the document being parsed.</span>
<a id="L38"></a>type Name struct {
    <a id="L39"></a>Space, Local string;
<a id="L40"></a>}

<a id="L42"></a><span class="comment">// An Attr represents an attribute in an XML element (Name=Value).</span>
<a id="L43"></a>type Attr struct {
    <a id="L44"></a>Name  Name;
    <a id="L45"></a>Value string;
<a id="L46"></a>}

<a id="L48"></a><span class="comment">// A Token is an interface holding one of the token types:</span>
<a id="L49"></a><span class="comment">// StartElement, EndElement, CharData, Comment, ProcInst, or Directive.</span>
<a id="L50"></a>type Token interface{}

<a id="L52"></a><span class="comment">// A StartElement represents an XML start element.</span>
<a id="L53"></a>type StartElement struct {
    <a id="L54"></a>Name Name;
    <a id="L55"></a>Attr []Attr;
<a id="L56"></a>}

<a id="L58"></a><span class="comment">// An EndElement represents an XML end element.</span>
<a id="L59"></a>type EndElement struct {
    <a id="L60"></a>Name Name;
<a id="L61"></a>}

<a id="L63"></a><span class="comment">// A CharData represents XML character data (raw text),</span>
<a id="L64"></a><span class="comment">// in which XML escape sequences have been replaced by</span>
<a id="L65"></a><span class="comment">// the characters they represent.</span>
<a id="L66"></a>type CharData []byte

<a id="L68"></a>func copy(b []byte) []byte {
    <a id="L69"></a>b1 := make([]byte, len(b));
    <a id="L70"></a>bytes.Copy(b1, b);
    <a id="L71"></a>return b1;
<a id="L72"></a>}

<a id="L74"></a>func (c CharData) Copy() CharData { return CharData(copy(c)) }

<a id="L76"></a><span class="comment">// A Comment represents an XML comment of the form &lt;!--comment--&gt;.</span>
<a id="L77"></a><span class="comment">// The bytes do not include the &lt;!-- and --&gt; comment markers.</span>
<a id="L78"></a>type Comment []byte

<a id="L80"></a>func (c Comment) Copy() Comment { return Comment(copy(c)) }

<a id="L82"></a><span class="comment">// A ProcInst represents an XML processing instruction of the form &lt;?target inst?&gt;</span>
<a id="L83"></a>type ProcInst struct {
    <a id="L84"></a>Target string;
    <a id="L85"></a>Inst   []byte;
<a id="L86"></a>}

<a id="L88"></a>func (p ProcInst) Copy() ProcInst {
    <a id="L89"></a>p.Inst = copy(p.Inst);
    <a id="L90"></a>return p;
<a id="L91"></a>}

<a id="L93"></a><span class="comment">// A Directive represents an XML directive of the form &lt;!text&gt;.</span>
<a id="L94"></a><span class="comment">// The bytes do not include the &lt;! and &gt; markers.</span>
<a id="L95"></a>type Directive []byte

<a id="L97"></a>func (d Directive) Copy() Directive { return Directive(copy(d)) }

<a id="L99"></a>type readByter interface {
    <a id="L100"></a>ReadByte() (b byte, err os.Error);
<a id="L101"></a>}

<a id="L103"></a><span class="comment">// A Parser represents an XML parser reading a particular input stream.</span>
<a id="L104"></a><span class="comment">// The parser assumes that its input is encoded in UTF-8.</span>
<a id="L105"></a>type Parser struct {
    <a id="L106"></a><span class="comment">// Strict defaults to true, enforcing the requirements</span>
    <a id="L107"></a><span class="comment">// of the XML specification.</span>
    <a id="L108"></a><span class="comment">// If set to false, the parser allows input containing common</span>
    <a id="L109"></a><span class="comment">// mistakes:</span>
    <a id="L110"></a><span class="comment">//	* If an element is missing an end tag, the parser invents</span>
    <a id="L111"></a><span class="comment">//	  end tags as necessary to keep the return values from Token</span>
    <a id="L112"></a><span class="comment">//	  properly balanced.</span>
    <a id="L113"></a><span class="comment">//	* In attribute values and character data, unknown or malformed</span>
    <a id="L114"></a><span class="comment">//	  character entities (sequences beginning with &amp;) are left alone.</span>
    <a id="L115"></a><span class="comment">//</span>
    <a id="L116"></a><span class="comment">// Setting:</span>
    <a id="L117"></a><span class="comment">//</span>
    <a id="L118"></a><span class="comment">//	p.Strict = false;</span>
    <a id="L119"></a><span class="comment">//	p.AutoClose = HTMLAutoClose;</span>
    <a id="L120"></a><span class="comment">//	p.Entity = HTMLEntity</span>
    <a id="L121"></a><span class="comment">//</span>
    <a id="L122"></a><span class="comment">// creates a parser that can handle typical HTML.</span>
    <a id="L123"></a>Strict bool;

    <a id="L125"></a><span class="comment">// When Strict == false, AutoClose indicates a set of elements to</span>
    <a id="L126"></a><span class="comment">// consider closed immediately after they are opened, regardless</span>
    <a id="L127"></a><span class="comment">// of whether an end element is present.</span>
    <a id="L128"></a>AutoClose []string;

    <a id="L130"></a><span class="comment">// Entity can be used to map non-standard entity names to string replacements.</span>
    <a id="L131"></a><span class="comment">// The parser behaves as if these standard mappings are present in the map,</span>
    <a id="L132"></a><span class="comment">// regardless of the actual map content:</span>
    <a id="L133"></a><span class="comment">//</span>
    <a id="L134"></a><span class="comment">//	&#34;lt&#34;: &#34;&lt;&#34;,</span>
    <a id="L135"></a><span class="comment">//	&#34;gt&#34;: &#34;&gt;&#34;,</span>
    <a id="L136"></a><span class="comment">//	&#34;amp&#34;: &#34;&amp;&#34;,</span>
    <a id="L137"></a><span class="comment">//	&#34;pos&#34;: &#34;&#39;&#34;,</span>
    <a id="L138"></a><span class="comment">//	&#34;quot&#34;: `&#34;`,</span>
    <a id="L139"></a><span class="comment">//</span>
    <a id="L140"></a>Entity map[string]string;

    <a id="L142"></a>r         readByter;
    <a id="L143"></a>buf       bytes.Buffer;
    <a id="L144"></a>stk       *stack;
    <a id="L145"></a>free      *stack;
    <a id="L146"></a>needClose bool;
    <a id="L147"></a>toClose   Name;
    <a id="L148"></a>nextToken Token;
    <a id="L149"></a>nextByte  int;
    <a id="L150"></a>ns        map[string]string;
    <a id="L151"></a>err       os.Error;
    <a id="L152"></a>line      int;
    <a id="L153"></a>tmp       [32]byte;
<a id="L154"></a>}

<a id="L156"></a><span class="comment">// NewParser creates a new XML parser reading from r.</span>
<a id="L157"></a>func NewParser(r io.Reader) *Parser {
    <a id="L158"></a>p := &amp;Parser{
        <a id="L159"></a>ns: make(map[string]string),
        <a id="L160"></a>nextByte: -1,
        <a id="L161"></a>line: 1,
        <a id="L162"></a>Strict: true,
    <a id="L163"></a>};

    <a id="L165"></a><span class="comment">// Get efficient byte at a time reader.</span>
    <a id="L166"></a><span class="comment">// Assume that if reader has its own</span>
    <a id="L167"></a><span class="comment">// ReadByte, it&#39;s efficient enough.</span>
    <a id="L168"></a><span class="comment">// Otherwise, use bufio.</span>
    <a id="L169"></a>if rb, ok := r.(readByter); ok {
        <a id="L170"></a>p.r = rb
    <a id="L171"></a>} else {
        <a id="L172"></a>p.r = bufio.NewReader(r)
    <a id="L173"></a>}

    <a id="L175"></a>return p;
<a id="L176"></a>}

<a id="L178"></a><span class="comment">// Token returns the next XML token in the input stream.</span>
<a id="L179"></a><span class="comment">// At the end of the input stream, Token returns nil, os.EOF.</span>
<a id="L180"></a><span class="comment">//</span>
<a id="L181"></a><span class="comment">// Slices of bytes in the returned token data refer to the</span>
<a id="L182"></a><span class="comment">// parser&#39;s internal buffer and remain valid only until the next</span>
<a id="L183"></a><span class="comment">// call to Token.  To acquire a copy of the bytes, call the token&#39;s</span>
<a id="L184"></a><span class="comment">// Copy method.</span>
<a id="L185"></a><span class="comment">//</span>
<a id="L186"></a><span class="comment">// Token expands self-closing elements such as &lt;br/&gt;</span>
<a id="L187"></a><span class="comment">// into separate start and end elements returned by successive calls.</span>
<a id="L188"></a><span class="comment">//</span>
<a id="L189"></a><span class="comment">// Token guarantees that the StartElement and EndElement</span>
<a id="L190"></a><span class="comment">// tokens it returns are properly nested and matched:</span>
<a id="L191"></a><span class="comment">// if Token encounters an unexpected end element,</span>
<a id="L192"></a><span class="comment">// it will return an error.</span>
<a id="L193"></a><span class="comment">//</span>
<a id="L194"></a><span class="comment">// Token implements XML name spaces as described by</span>
<a id="L195"></a><span class="comment">// http://www.w3.org/TR/REC-xml-names/.  Each of the</span>
<a id="L196"></a><span class="comment">// Name structures contained in the Token has the Space</span>
<a id="L197"></a><span class="comment">// set to the URL identifying its name space when known.</span>
<a id="L198"></a><span class="comment">// If Token encounters an unrecognized name space prefix,</span>
<a id="L199"></a><span class="comment">// it uses the prefix as the Space rather than report an error.</span>
<a id="L200"></a><span class="comment">//</span>
<a id="L201"></a>func (p *Parser) Token() (t Token, err os.Error) {
    <a id="L202"></a>if p.nextToken != nil {
        <a id="L203"></a>t = p.nextToken;
        <a id="L204"></a>p.nextToken = nil;
    <a id="L205"></a>} else if t, err = p.RawToken(); err != nil {
        <a id="L206"></a>return
    <a id="L207"></a>}

    <a id="L209"></a>if !p.Strict {
        <a id="L210"></a>if t1, ok := p.autoClose(t); ok {
            <a id="L211"></a>p.nextToken = t;
            <a id="L212"></a>t = t1;
        <a id="L213"></a>}
    <a id="L214"></a>}
    <a id="L215"></a>switch t1 := t.(type) {
    <a id="L216"></a>case StartElement:
        <a id="L217"></a><span class="comment">// In XML name spaces, the translations listed in the</span>
        <a id="L218"></a><span class="comment">// attributes apply to the element name and</span>
        <a id="L219"></a><span class="comment">// to the other attribute names, so process</span>
        <a id="L220"></a><span class="comment">// the translations first.</span>
        <a id="L221"></a>for _, a := range t1.Attr {
            <a id="L222"></a>if a.Name.Space == &#34;xmlns&#34; {
                <a id="L223"></a>v, ok := p.ns[a.Name.Local];
                <a id="L224"></a>p.pushNs(a.Name.Local, v, ok);
                <a id="L225"></a>p.ns[a.Name.Local] = a.Value;
            <a id="L226"></a>}
            <a id="L227"></a>if a.Name.Space == &#34;&#34; &amp;&amp; a.Name.Local == &#34;xmlns&#34; {
                <a id="L228"></a><span class="comment">// Default space for untagged names</span>
                <a id="L229"></a>v, ok := p.ns[&#34;&#34;];
                <a id="L230"></a>p.pushNs(&#34;&#34;, v, ok);
                <a id="L231"></a>p.ns[&#34;&#34;] = a.Value;
            <a id="L232"></a>}
        <a id="L233"></a>}

        <a id="L235"></a>p.translate(&amp;t1.Name, true);
        <a id="L236"></a>for i := range t1.Attr {
            <a id="L237"></a>p.translate(&amp;t1.Attr[i].Name, false)
        <a id="L238"></a>}
        <a id="L239"></a>p.pushElement(t1.Name);
        <a id="L240"></a>t = t1;

    <a id="L242"></a>case EndElement:
        <a id="L243"></a>p.translate(&amp;t1.Name, true);
        <a id="L244"></a>if !p.popElement(&amp;t1) {
            <a id="L245"></a>return nil, p.err
        <a id="L246"></a>}
        <a id="L247"></a>t = t1;
    <a id="L248"></a>}
    <a id="L249"></a>return;
<a id="L250"></a>}

<a id="L252"></a><span class="comment">// Apply name space translation to name n.</span>
<a id="L253"></a><span class="comment">// The default name space (for Space==&#34;&#34;)</span>
<a id="L254"></a><span class="comment">// applies only to element names, not to attribute names.</span>
<a id="L255"></a>func (p *Parser) translate(n *Name, isElementName bool) {
    <a id="L256"></a>switch {
    <a id="L257"></a>case n.Space == &#34;xmlns&#34;:
        <a id="L258"></a>return
    <a id="L259"></a>case n.Space == &#34;&#34; &amp;&amp; !isElementName:
        <a id="L260"></a>return
    <a id="L261"></a>case n.Space == &#34;&#34; &amp;&amp; n.Local == &#34;xmlns&#34;:
        <a id="L262"></a>return
    <a id="L263"></a>}
    <a id="L264"></a>if v, ok := p.ns[n.Space]; ok {
        <a id="L265"></a>n.Space = v
    <a id="L266"></a>}
<a id="L267"></a>}

<a id="L269"></a><span class="comment">// Parsing state - stack holds old name space translations</span>
<a id="L270"></a><span class="comment">// and the current set of open elements.  The translations to pop when</span>
<a id="L271"></a><span class="comment">// ending a given tag are *below* it on the stack, which is</span>
<a id="L272"></a><span class="comment">// more work but forced on us by XML.</span>
<a id="L273"></a>type stack struct {
    <a id="L274"></a>next *stack;
    <a id="L275"></a>kind int;
    <a id="L276"></a>name Name;
    <a id="L277"></a>ok   bool;
<a id="L278"></a>}

<a id="L280"></a>const (
    <a id="L281"></a>stkStart = iota;
    <a id="L282"></a>stkNs;
<a id="L283"></a>)

<a id="L285"></a>func (p *Parser) push(kind int) *stack {
    <a id="L286"></a>s := p.free;
    <a id="L287"></a>if s != nil {
        <a id="L288"></a>p.free = s.next
    <a id="L289"></a>} else {
        <a id="L290"></a>s = new(stack)
    <a id="L291"></a>}
    <a id="L292"></a>s.next = p.stk;
    <a id="L293"></a>s.kind = kind;
    <a id="L294"></a>p.stk = s;
    <a id="L295"></a>return s;
<a id="L296"></a>}

<a id="L298"></a>func (p *Parser) pop() *stack {
    <a id="L299"></a>s := p.stk;
    <a id="L300"></a>if s != nil {
        <a id="L301"></a>p.stk = s.next;
        <a id="L302"></a>s.next = p.free;
        <a id="L303"></a>p.free = s;
    <a id="L304"></a>}
    <a id="L305"></a>return s;
<a id="L306"></a>}

<a id="L308"></a><span class="comment">// Record that we are starting an element with the given name.</span>
<a id="L309"></a>func (p *Parser) pushElement(name Name) {
    <a id="L310"></a>s := p.push(stkStart);
    <a id="L311"></a>s.name = name;
<a id="L312"></a>}

<a id="L314"></a><span class="comment">// Record that we are changing the value of ns[local].</span>
<a id="L315"></a><span class="comment">// The old value is url, ok.</span>
<a id="L316"></a>func (p *Parser) pushNs(local string, url string, ok bool) {
    <a id="L317"></a>s := p.push(stkNs);
    <a id="L318"></a>s.name.Local = local;
    <a id="L319"></a>s.name.Space = url;
    <a id="L320"></a>s.ok = ok;
<a id="L321"></a>}

<a id="L323"></a><span class="comment">// Record that we are ending an element with the given name.</span>
<a id="L324"></a><span class="comment">// The name must match the record at the top of the stack,</span>
<a id="L325"></a><span class="comment">// which must be a pushElement record.</span>
<a id="L326"></a><span class="comment">// After popping the element, apply any undo records from</span>
<a id="L327"></a><span class="comment">// the stack to restore the name translations that existed</span>
<a id="L328"></a><span class="comment">// before we saw this element.</span>
<a id="L329"></a>func (p *Parser) popElement(t *EndElement) bool {
    <a id="L330"></a>s := p.pop();
    <a id="L331"></a>name := t.Name;
    <a id="L332"></a>switch {
    <a id="L333"></a>case s == nil || s.kind != stkStart:
        <a id="L334"></a>p.err = SyntaxError(&#34;unexpected end element &lt;/&#34; + name.Local + &#34;&gt;&#34;);
        <a id="L335"></a>return false;
    <a id="L336"></a>case s.name.Local != name.Local:
        <a id="L337"></a>if !p.Strict {
            <a id="L338"></a>p.needClose = true;
            <a id="L339"></a>p.toClose = t.Name;
            <a id="L340"></a>t.Name = s.name;
            <a id="L341"></a>return true;
        <a id="L342"></a>}
        <a id="L343"></a>p.err = SyntaxError(&#34;element &lt;&#34; + s.name.Local + &#34;&gt; closed by &lt;/&#34; + name.Local + &#34;&gt;&#34;);
        <a id="L344"></a>return false;
    <a id="L345"></a>case s.name.Space != name.Space:
        <a id="L346"></a>p.err = SyntaxError(&#34;element &lt;&#34; + s.name.Local + &#34;&gt; in space &#34; + s.name.Space +
            <a id="L347"></a>&#34;closed by &lt;/&#34; + name.Local + &#34;&gt; in space &#34; + name.Space);
        <a id="L348"></a>return false;
    <a id="L349"></a>}

    <a id="L351"></a><span class="comment">// Pop stack until a Start is on the top, undoing the</span>
    <a id="L352"></a><span class="comment">// translations that were associated with the element we just closed.</span>
    <a id="L353"></a>for p.stk != nil &amp;&amp; p.stk.kind != stkStart {
        <a id="L354"></a>s := p.pop();
        <a id="L355"></a>p.ns[s.name.Local] = s.name.Space, s.ok;
    <a id="L356"></a>}

    <a id="L358"></a>return true;
<a id="L359"></a>}

<a id="L361"></a><span class="comment">// If the top element on the stack is autoclosing and</span>
<a id="L362"></a><span class="comment">// t is not the end tag, invent the end tag.</span>
<a id="L363"></a>func (p *Parser) autoClose(t Token) (Token, bool) {
    <a id="L364"></a>if p.stk == nil || p.stk.kind != stkStart {
        <a id="L365"></a>return nil, false
    <a id="L366"></a>}
    <a id="L367"></a>name := strings.ToLower(p.stk.name.Local);
    <a id="L368"></a>for _, s := range p.AutoClose {
        <a id="L369"></a>if strings.ToLower(s) == name {
            <a id="L370"></a><span class="comment">// This one should be auto closed if t doesn&#39;t close it.</span>
            <a id="L371"></a>et, ok := t.(EndElement);
            <a id="L372"></a>if !ok || et.Name.Local != name {
                <a id="L373"></a>return EndElement{p.stk.name}, true
            <a id="L374"></a>}
            <a id="L375"></a>break;
        <a id="L376"></a>}
    <a id="L377"></a>}
    <a id="L378"></a>return nil, false;
<a id="L379"></a>}


<a id="L382"></a><span class="comment">// RawToken is like Token but does not verify that</span>
<a id="L383"></a><span class="comment">// start and end elements match and does not translate</span>
<a id="L384"></a><span class="comment">// name space prefixes to their corresponding URLs.</span>
<a id="L385"></a>func (p *Parser) RawToken() (Token, os.Error) {
    <a id="L386"></a>if p.err != nil {
        <a id="L387"></a>return nil, p.err
    <a id="L388"></a>}
    <a id="L389"></a>if p.needClose {
        <a id="L390"></a><span class="comment">// The last element we read was self-closing and</span>
        <a id="L391"></a><span class="comment">// we returned just the StartElement half.</span>
        <a id="L392"></a><span class="comment">// Return the EndElement half now.</span>
        <a id="L393"></a>p.needClose = false;
        <a id="L394"></a>return EndElement{p.toClose}, nil;
    <a id="L395"></a>}

    <a id="L397"></a>b, ok := p.getc();
    <a id="L398"></a>if !ok {
        <a id="L399"></a>return nil, p.err
    <a id="L400"></a>}

    <a id="L402"></a>if b != &#39;&lt;&#39; {
        <a id="L403"></a><span class="comment">// Text section.</span>
        <a id="L404"></a>p.ungetc(b);
        <a id="L405"></a>data := p.text(-1, false);
        <a id="L406"></a>if data == nil {
            <a id="L407"></a>return nil, p.err
        <a id="L408"></a>}
        <a id="L409"></a>return CharData(data), nil;
    <a id="L410"></a>}

    <a id="L412"></a>if b, ok = p.getc(); !ok {
        <a id="L413"></a>return nil, p.err
    <a id="L414"></a>}
    <a id="L415"></a>switch b {
    <a id="L416"></a>case &#39;/&#39;:
        <a id="L417"></a><span class="comment">// &lt;/: End element</span>
        <a id="L418"></a>var name Name;
        <a id="L419"></a>if name, ok = p.nsname(); !ok {
            <a id="L420"></a>if p.err == nil {
                <a id="L421"></a>p.err = SyntaxError(&#34;expected element name after &lt;/&#34;)
            <a id="L422"></a>}
            <a id="L423"></a>return nil, p.err;
        <a id="L424"></a>}
        <a id="L425"></a>p.space();
        <a id="L426"></a>if b, ok = p.getc(); !ok {
            <a id="L427"></a>return nil, p.err
        <a id="L428"></a>}
        <a id="L429"></a>if b != &#39;&gt;&#39; {
            <a id="L430"></a>p.err = SyntaxError(&#34;invalid characters between &lt;/&#34; + name.Local + &#34; and &gt;&#34;);
            <a id="L431"></a>return nil, p.err;
        <a id="L432"></a>}
        <a id="L433"></a>return EndElement{name}, nil;

    <a id="L435"></a>case &#39;?&#39;:
        <a id="L436"></a><span class="comment">// &lt;?: Processing instruction.</span>
        <a id="L437"></a><span class="comment">// TODO(rsc): Should parse the &lt;?xml declaration to make sure</span>
        <a id="L438"></a><span class="comment">// the version is 1.0 and the encoding is UTF-8.</span>
        <a id="L439"></a>var target string;
        <a id="L440"></a>if target, ok = p.name(); !ok {
            <a id="L441"></a>return nil, p.err
        <a id="L442"></a>}
        <a id="L443"></a>p.space();
        <a id="L444"></a>p.buf.Reset();
        <a id="L445"></a>var b0 byte;
        <a id="L446"></a>for {
            <a id="L447"></a>if b, ok = p.getc(); !ok {
                <a id="L448"></a>if p.err == os.EOF {
                    <a id="L449"></a>p.err = SyntaxError(&#34;unterminated &lt;? directive&#34;)
                <a id="L450"></a>}
                <a id="L451"></a>return nil, p.err;
            <a id="L452"></a>}
            <a id="L453"></a>p.buf.WriteByte(b);
            <a id="L454"></a>if b0 == &#39;?&#39; &amp;&amp; b == &#39;&gt;&#39; {
                <a id="L455"></a>break
            <a id="L456"></a>}
            <a id="L457"></a>b0 = b;
        <a id="L458"></a>}
        <a id="L459"></a>data := p.buf.Bytes();
        <a id="L460"></a>data = data[0 : len(data)-2]; <span class="comment">// chop ?&gt;</span>
        <a id="L461"></a>return ProcInst{target, data}, nil;

    <a id="L463"></a>case &#39;!&#39;:
        <a id="L464"></a><span class="comment">// &lt;!: Maybe comment, maybe CDATA.</span>
        <a id="L465"></a>if b, ok = p.getc(); !ok {
            <a id="L466"></a>return nil, p.err
        <a id="L467"></a>}
        <a id="L468"></a>switch b {
        <a id="L469"></a>case &#39;-&#39;: <span class="comment">// &lt;!-</span>
            <a id="L470"></a><span class="comment">// Probably &lt;!-- for a comment.</span>
            <a id="L471"></a>if b, ok = p.getc(); !ok {
                <a id="L472"></a>return nil, p.err
            <a id="L473"></a>}
            <a id="L474"></a>if b != &#39;-&#39; {
                <a id="L475"></a>p.err = SyntaxError(&#34;invalid sequence &lt;!- not part of &lt;!--&#34;);
                <a id="L476"></a>return nil, p.err;
            <a id="L477"></a>}
            <a id="L478"></a><span class="comment">// Look for terminator.</span>
            <a id="L479"></a>p.buf.Reset();
            <a id="L480"></a>var b0, b1 byte;
            <a id="L481"></a>for {
                <a id="L482"></a>if b, ok = p.getc(); !ok {
                    <a id="L483"></a>if p.err == os.EOF {
                        <a id="L484"></a>p.err = SyntaxError(&#34;unterminated &lt;!-- comment&#34;)
                    <a id="L485"></a>}
                    <a id="L486"></a>return nil, p.err;
                <a id="L487"></a>}
                <a id="L488"></a>p.buf.WriteByte(b);
                <a id="L489"></a>if b0 == &#39;-&#39; &amp;&amp; b1 == &#39;-&#39; &amp;&amp; b == &#39;&gt;&#39; {
                    <a id="L490"></a>break
                <a id="L491"></a>}
                <a id="L492"></a>b0, b1 = b1, b;
            <a id="L493"></a>}
            <a id="L494"></a>data := p.buf.Bytes();
            <a id="L495"></a>data = data[0 : len(data)-3]; <span class="comment">// chop --&gt;</span>
            <a id="L496"></a>return Comment(data), nil;

        <a id="L498"></a>case &#39;[&#39;: <span class="comment">// &lt;![</span>
            <a id="L499"></a><span class="comment">// Probably &lt;![CDATA[.</span>
            <a id="L500"></a>for i := 0; i &lt; 7; i++ {
                <a id="L501"></a>if b, ok = p.getc(); !ok {
                    <a id="L502"></a>return nil, p.err
                <a id="L503"></a>}
                <a id="L504"></a>if b != &#34;[CDATA[&#34;[i] {
                    <a id="L505"></a>p.err = SyntaxError(&#34;invalid &lt;![ sequence&#34;);
                    <a id="L506"></a>return nil, p.err;
                <a id="L507"></a>}
            <a id="L508"></a>}
            <a id="L509"></a><span class="comment">// Have &lt;![CDATA[.  Read text until ]]&gt;.</span>
            <a id="L510"></a>data := p.text(-1, true);
            <a id="L511"></a>if data == nil {
                <a id="L512"></a>return nil, p.err
            <a id="L513"></a>}
            <a id="L514"></a>return CharData(data), nil;
        <a id="L515"></a>}

        <a id="L517"></a><span class="comment">// Probably a directive: &lt;!DOCTYPE ...&gt;, &lt;!ENTITY ...&gt;, etc.</span>
        <a id="L518"></a><span class="comment">// We don&#39;t care, but accumulate for caller.</span>
        <a id="L519"></a>p.buf.Reset();
        <a id="L520"></a>p.buf.WriteByte(b);
        <a id="L521"></a>for {
            <a id="L522"></a>if b, ok = p.getc(); !ok {
                <a id="L523"></a>return nil, p.err
            <a id="L524"></a>}
            <a id="L525"></a>if b == &#39;&gt;&#39; {
                <a id="L526"></a>break
            <a id="L527"></a>}
            <a id="L528"></a>p.buf.WriteByte(b);
        <a id="L529"></a>}
        <a id="L530"></a>return Directive(p.buf.Bytes()), nil;
    <a id="L531"></a>}

    <a id="L533"></a><span class="comment">// Must be an open element like &lt;a href=&#34;foo&#34;&gt;</span>
    <a id="L534"></a>p.ungetc(b);

    <a id="L536"></a>var (
        <a id="L537"></a>name  Name;
        <a id="L538"></a>empty bool;
        <a id="L539"></a>attr  []Attr;
    <a id="L540"></a>)
    <a id="L541"></a>if name, ok = p.nsname(); !ok {
        <a id="L542"></a>if p.err == nil {
            <a id="L543"></a>p.err = SyntaxError(&#34;expected element name after &lt;&#34;)
        <a id="L544"></a>}
        <a id="L545"></a>return nil, p.err;
    <a id="L546"></a>}

    <a id="L548"></a>attr = make([]Attr, 0, 4);
    <a id="L549"></a>for {
        <a id="L550"></a>p.space();
        <a id="L551"></a>if b, ok = p.getc(); !ok {
            <a id="L552"></a>return nil, p.err
        <a id="L553"></a>}
        <a id="L554"></a>if b == &#39;/&#39; {
            <a id="L555"></a>empty = true;
            <a id="L556"></a>if b, ok = p.getc(); !ok {
                <a id="L557"></a>return nil, p.err
            <a id="L558"></a>}
            <a id="L559"></a>if b != &#39;&gt;&#39; {
                <a id="L560"></a>p.err = SyntaxError(&#34;expected /&gt; in element&#34;);
                <a id="L561"></a>return nil, p.err;
            <a id="L562"></a>}
            <a id="L563"></a>break;
        <a id="L564"></a>}
        <a id="L565"></a>if b == &#39;&gt;&#39; {
            <a id="L566"></a>break
        <a id="L567"></a>}
        <a id="L568"></a>p.ungetc(b);

        <a id="L570"></a>n := len(attr);
        <a id="L571"></a>if n &gt;= cap(attr) {
            <a id="L572"></a>nattr := make([]Attr, n, 2*cap(attr));
            <a id="L573"></a>for i, a := range attr {
                <a id="L574"></a>nattr[i] = a
            <a id="L575"></a>}
            <a id="L576"></a>attr = nattr;
        <a id="L577"></a>}
        <a id="L578"></a>attr = attr[0 : n+1];
        <a id="L579"></a>a := &amp;attr[n];
        <a id="L580"></a>if a.Name, ok = p.nsname(); !ok {
            <a id="L581"></a>if p.err == nil {
                <a id="L582"></a>p.err = SyntaxError(&#34;expected attribute name in element&#34;)
            <a id="L583"></a>}
            <a id="L584"></a>return nil, p.err;
        <a id="L585"></a>}
        <a id="L586"></a>p.space();
        <a id="L587"></a>if b, ok = p.getc(); !ok {
            <a id="L588"></a>return nil, p.err
        <a id="L589"></a>}
        <a id="L590"></a>if b != &#39;=&#39; {
            <a id="L591"></a>p.err = SyntaxError(&#34;attribute name without = in element&#34;);
            <a id="L592"></a>return nil, p.err;
        <a id="L593"></a>}
        <a id="L594"></a>p.space();
        <a id="L595"></a>if b, ok = p.getc(); !ok {
            <a id="L596"></a>return nil, p.err
        <a id="L597"></a>}
        <a id="L598"></a>if b != &#39;&#34;&#39; &amp;&amp; b != &#39;\&#39;&#39; {
            <a id="L599"></a>p.err = SyntaxError(&#34;unquoted or missing attribute value in element&#34;);
            <a id="L600"></a>return nil, p.err;
        <a id="L601"></a>}
        <a id="L602"></a>data := p.text(int(b), false);
        <a id="L603"></a>if data == nil {
            <a id="L604"></a>return nil, p.err
        <a id="L605"></a>}
        <a id="L606"></a>a.Value = string(data);
    <a id="L607"></a>}

    <a id="L609"></a>if empty {
        <a id="L610"></a>p.needClose = true;
        <a id="L611"></a>p.toClose = name;
    <a id="L612"></a>}
    <a id="L613"></a>return StartElement{name, attr}, nil;
<a id="L614"></a>}

<a id="L616"></a><span class="comment">// Skip spaces if any</span>
<a id="L617"></a>func (p *Parser) space() {
    <a id="L618"></a>for {
        <a id="L619"></a>b, ok := p.getc();
        <a id="L620"></a>if !ok {
            <a id="L621"></a>return
        <a id="L622"></a>}
        <a id="L623"></a>switch b {
        <a id="L624"></a>case &#39; &#39;, &#39;\r&#39;, &#39;\n&#39;, &#39;\t&#39;:
        <a id="L625"></a>default:
            <a id="L626"></a>p.ungetc(b);
            <a id="L627"></a>return;
        <a id="L628"></a>}
    <a id="L629"></a>}
<a id="L630"></a>}

<a id="L632"></a><span class="comment">// Read a single byte.</span>
<a id="L633"></a><span class="comment">// If there is no byte to read, return ok==false</span>
<a id="L634"></a><span class="comment">// and leave the error in p.err.</span>
<a id="L635"></a><span class="comment">// Maintain line number.</span>
<a id="L636"></a>func (p *Parser) getc() (b byte, ok bool) {
    <a id="L637"></a>if p.err != nil {
        <a id="L638"></a>return 0, false
    <a id="L639"></a>}
    <a id="L640"></a>if p.nextByte &gt;= 0 {
        <a id="L641"></a>b = byte(p.nextByte);
        <a id="L642"></a>p.nextByte = -1;
    <a id="L643"></a>} else {
        <a id="L644"></a>b, p.err = p.r.ReadByte();
        <a id="L645"></a>if p.err != nil {
            <a id="L646"></a>return 0, false
        <a id="L647"></a>}
    <a id="L648"></a>}
    <a id="L649"></a>if b == &#39;\n&#39; {
        <a id="L650"></a>p.line++
    <a id="L651"></a>}
    <a id="L652"></a>return b, true;
<a id="L653"></a>}

<a id="L655"></a><span class="comment">// Unread a single byte.</span>
<a id="L656"></a>func (p *Parser) ungetc(b byte) {
    <a id="L657"></a>if b == &#39;\n&#39; {
        <a id="L658"></a>p.line--
    <a id="L659"></a>}
    <a id="L660"></a>p.nextByte = int(b);
<a id="L661"></a>}

<a id="L663"></a>var entity = map[string]int{
    <a id="L664"></a>&#34;lt&#34;: &#39;&lt;&#39;,
    <a id="L665"></a>&#34;gt&#34;: &#39;&gt;&#39;,
    <a id="L666"></a>&#34;amp&#34;: &#39;&amp;&#39;,
    <a id="L667"></a>&#34;apos&#34;: &#39;\&#39;&#39;,
    <a id="L668"></a>&#34;quot&#34;: &#39;&#34;&#39;,
<a id="L669"></a>}

<a id="L671"></a><span class="comment">// Read plain text section (XML calls it character data).</span>
<a id="L672"></a><span class="comment">// If quote &gt;= 0, we are in a quoted string and need to find the matching quote.</span>
<a id="L673"></a><span class="comment">// If cdata == true, we are in a &lt;![CDATA[ section and need to find ]]&gt;.</span>
<a id="L674"></a><span class="comment">// On failure return nil and leave the error in p.err.</span>
<a id="L675"></a>func (p *Parser) text(quote int, cdata bool) []byte {
    <a id="L676"></a>var b0, b1 byte;
    <a id="L677"></a>var trunc int;
    <a id="L678"></a>p.buf.Reset();
<a id="L679"></a>Input:
    <a id="L680"></a>for {
        <a id="L681"></a>b, ok := p.getc();
        <a id="L682"></a>if !ok {
            <a id="L683"></a>return nil
        <a id="L684"></a>}

        <a id="L686"></a><span class="comment">// &lt;![CDATA[ section ends with ]]&gt;.</span>
        <a id="L687"></a><span class="comment">// It is an error for ]]&gt; to appear in ordinary text.</span>
        <a id="L688"></a>if b0 == &#39;]&#39; &amp;&amp; b1 == &#39;]&#39; &amp;&amp; b == &#39;&gt;&#39; {
            <a id="L689"></a>if cdata {
                <a id="L690"></a>trunc = 2;
                <a id="L691"></a>break Input;
            <a id="L692"></a>}
            <a id="L693"></a>p.err = SyntaxError(&#34;unescaped ]]&gt; not in CDATA section&#34;);
            <a id="L694"></a>return nil;
        <a id="L695"></a>}

        <a id="L697"></a><span class="comment">// Stop reading text if we see a &lt;.</span>
        <a id="L698"></a>if b == &#39;&lt;&#39; &amp;&amp; !cdata {
            <a id="L699"></a>if quote &gt;= 0 {
                <a id="L700"></a>p.err = SyntaxError(&#34;unescaped &lt; inside quoted string&#34;);
                <a id="L701"></a>return nil;
            <a id="L702"></a>}
            <a id="L703"></a>p.ungetc(&#39;&lt;&#39;);
            <a id="L704"></a>break Input;
        <a id="L705"></a>}
        <a id="L706"></a>if quote &gt;= 0 &amp;&amp; b == byte(quote) {
            <a id="L707"></a>break Input
        <a id="L708"></a>}
        <a id="L709"></a>if b == &#39;&amp;&#39; {
            <a id="L710"></a><span class="comment">// Read escaped character expression up to semicolon.</span>
            <a id="L711"></a><span class="comment">// XML in all its glory allows a document to define and use</span>
            <a id="L712"></a><span class="comment">// its own character names with &lt;!ENTITY ...&gt; directives.</span>
            <a id="L713"></a><span class="comment">// Parsers are required to recognize lt, gt, amp, apos, and quot</span>
            <a id="L714"></a><span class="comment">// even if they have not been declared.  That&#39;s all we allow.</span>
            <a id="L715"></a>var i int;
        <a id="L716"></a>CharLoop:
            <a id="L717"></a>for i = 0; i &lt; len(p.tmp); i++ {
                <a id="L718"></a>p.tmp[i], p.err = p.r.ReadByte();
                <a id="L719"></a>if p.err != nil {
                    <a id="L720"></a>return nil
                <a id="L721"></a>}
                <a id="L722"></a>c := p.tmp[i];
                <a id="L723"></a>if c == &#39;;&#39; {
                    <a id="L724"></a>break
                <a id="L725"></a>}
                <a id="L726"></a>if &#39;a&#39; &lt;= c &amp;&amp; c &lt;= &#39;z&#39; ||
                    <a id="L727"></a>&#39;A&#39; &lt;= c &amp;&amp; c &lt;= &#39;Z&#39; ||
                    <a id="L728"></a>&#39;0&#39; &lt;= c &amp;&amp; c &lt;= &#39;9&#39; ||
                    <a id="L729"></a>c == &#39;_&#39; || c == &#39;#&#39; {
                    <a id="L730"></a>continue
                <a id="L731"></a>}
                <a id="L732"></a>p.ungetc(c);
                <a id="L733"></a>break;
            <a id="L734"></a>}
            <a id="L735"></a>s := string(p.tmp[0:i]);
            <a id="L736"></a>if i &gt;= len(p.tmp) {
                <a id="L737"></a>if !p.Strict {
                    <a id="L738"></a>b0, b1 = 0, 0;
                    <a id="L739"></a>p.buf.WriteByte(&#39;&amp;&#39;);
                    <a id="L740"></a>p.buf.Write(p.tmp[0:i]);
                    <a id="L741"></a>continue Input;
                <a id="L742"></a>}
                <a id="L743"></a>p.err = SyntaxError(&#34;character entity expression &amp;&#34; + s + &#34;... too long&#34;);
                <a id="L744"></a>return nil;
            <a id="L745"></a>}
            <a id="L746"></a>var haveText bool;
            <a id="L747"></a>var text string;
            <a id="L748"></a>if i &gt;= 2 &amp;&amp; s[0] == &#39;#&#39; {
                <a id="L749"></a>var n uint64;
                <a id="L750"></a>var err os.Error;
                <a id="L751"></a>if i &gt;= 3 &amp;&amp; s[1] == &#39;x&#39; {
                    <a id="L752"></a>n, err = strconv.Btoui64(s[2:len(s)], 16)
                <a id="L753"></a>} else {
                    <a id="L754"></a>n, err = strconv.Btoui64(s[1:len(s)], 10)
                <a id="L755"></a>}
                <a id="L756"></a>if err == nil &amp;&amp; n &lt;= unicode.MaxRune {
                    <a id="L757"></a>text = string(n);
                    <a id="L758"></a>haveText = true;
                <a id="L759"></a>}
            <a id="L760"></a>} else {
                <a id="L761"></a>if r, ok := entity[s]; ok {
                    <a id="L762"></a>text = string(r);
                    <a id="L763"></a>haveText = true;
                <a id="L764"></a>} else {
                    <a id="L765"></a>text, haveText = p.Entity[s]
                <a id="L766"></a>}
            <a id="L767"></a>}
            <a id="L768"></a>if !haveText {
                <a id="L769"></a>if !p.Strict {
                    <a id="L770"></a>b0, b1 = 0, 0;
                    <a id="L771"></a>p.buf.WriteByte(&#39;&amp;&#39;);
                    <a id="L772"></a>p.buf.Write(p.tmp[0:i]);
                    <a id="L773"></a>continue Input;
                <a id="L774"></a>}
                <a id="L775"></a>p.err = SyntaxError(&#34;invalid character entity &amp;&#34; + s + &#34;;&#34;);
                <a id="L776"></a>return nil;
            <a id="L777"></a>}
            <a id="L778"></a>p.buf.Write(strings.Bytes(text));
            <a id="L779"></a>b0, b1 = 0, 0;
            <a id="L780"></a>continue Input;
        <a id="L781"></a>}
        <a id="L782"></a>p.buf.WriteByte(b);
        <a id="L783"></a>b0, b1 = b1, b;
    <a id="L784"></a>}
    <a id="L785"></a>data := p.buf.Bytes();
    <a id="L786"></a>data = data[0 : len(data)-trunc];

    <a id="L788"></a><span class="comment">// Must rewrite \r and \r\n into \n.</span>
    <a id="L789"></a>w := 0;
    <a id="L790"></a>for r := 0; r &lt; len(data); r++ {
        <a id="L791"></a>b := data[r];
        <a id="L792"></a>if b == &#39;\r&#39; {
            <a id="L793"></a>if r+1 &lt; len(data) &amp;&amp; data[r+1] == &#39;\n&#39; {
                <a id="L794"></a>continue
            <a id="L795"></a>}
            <a id="L796"></a>b = &#39;\n&#39;;
        <a id="L797"></a>}
        <a id="L798"></a>data[w] = b;
        <a id="L799"></a>w++;
    <a id="L800"></a>}
    <a id="L801"></a>return data[0:w];
<a id="L802"></a>}

<a id="L804"></a><span class="comment">// Get name space name: name with a : stuck in the middle.</span>
<a id="L805"></a><span class="comment">// The part before the : is the name space identifier.</span>
<a id="L806"></a>func (p *Parser) nsname() (name Name, ok bool) {
    <a id="L807"></a>s, ok := p.name();
    <a id="L808"></a>if !ok {
        <a id="L809"></a>return
    <a id="L810"></a>}
    <a id="L811"></a>i := strings.Index(s, &#34;:&#34;);
    <a id="L812"></a>if i &lt; 0 {
        <a id="L813"></a>name.Local = s
    <a id="L814"></a>} else {
        <a id="L815"></a>name.Space = s[0:i];
        <a id="L816"></a>name.Local = s[i+1 : len(s)];
    <a id="L817"></a>}
    <a id="L818"></a>return name, true;
<a id="L819"></a>}

<a id="L821"></a><span class="comment">// Get name: /first(first|second)*/</span>
<a id="L822"></a><span class="comment">// Do not set p.err if the name is missing: let the caller provide better context.</span>
<a id="L823"></a>func (p *Parser) name() (s string, ok bool) {
    <a id="L824"></a>var b byte;
    <a id="L825"></a>if b, ok = p.getc(); !ok {
        <a id="L826"></a>return
    <a id="L827"></a>}

    <a id="L829"></a><span class="comment">// As a first approximation, we gather the bytes [A-Za-z_:.-\x80-\xFF]*</span>
    <a id="L830"></a>if b &lt; utf8.RuneSelf &amp;&amp; !isNameByte(b) {
        <a id="L831"></a>p.ungetc(b);
        <a id="L832"></a>return;
    <a id="L833"></a>}
    <a id="L834"></a>p.buf.Reset();
    <a id="L835"></a>p.buf.WriteByte(b);
    <a id="L836"></a>for {
        <a id="L837"></a>if b, ok = p.getc(); !ok {
            <a id="L838"></a>return
        <a id="L839"></a>}
        <a id="L840"></a>if b &lt; utf8.RuneSelf &amp;&amp; !isNameByte(b) {
            <a id="L841"></a>p.ungetc(b);
            <a id="L842"></a>break;
        <a id="L843"></a>}
        <a id="L844"></a>p.buf.WriteByte(b);
    <a id="L845"></a>}

    <a id="L847"></a><span class="comment">// Then we check the characters.</span>
    <a id="L848"></a>s = p.buf.String();
    <a id="L849"></a>for i, c := range s {
        <a id="L850"></a>if !unicode.Is(first, c) &amp;&amp; (i == 0 || !unicode.Is(second, c)) {
            <a id="L851"></a>p.err = SyntaxError(&#34;invalid XML name: &#34; + s);
            <a id="L852"></a>return &#34;&#34;, false;
        <a id="L853"></a>}
    <a id="L854"></a>}
    <a id="L855"></a>return s, true;
<a id="L856"></a>}

<a id="L858"></a>func isNameByte(c byte) bool {
    <a id="L859"></a>return &#39;A&#39; &lt;= c &amp;&amp; c &lt;= &#39;Z&#39; ||
        <a id="L860"></a>&#39;a&#39; &lt;= c &amp;&amp; c &lt;= &#39;z&#39; ||
        <a id="L861"></a>&#39;0&#39; &lt;= c &amp;&amp; c &lt;= &#39;9&#39; ||
        <a id="L862"></a>c == &#39;_&#39; || c == &#39;:&#39; || c == &#39;.&#39; || c == &#39;-&#39;
<a id="L863"></a>}

<a id="L865"></a><span class="comment">// These tables were generated by cut and paste from Appendix B of</span>
<a id="L866"></a><span class="comment">// the XML spec at http://www.xml.com/axml/testaxml.htm</span>
<a id="L867"></a><span class="comment">// and then reformatting.  First corresponds to (Letter | &#39;_&#39; | &#39;:&#39;)</span>
<a id="L868"></a><span class="comment">// and second corresponds to NameChar.</span>

<a id="L870"></a>var first = []unicode.Range{
    <a id="L871"></a>unicode.Range{0x003A, 0x003A, 1},
    <a id="L872"></a>unicode.Range{0x0041, 0x005A, 1},
    <a id="L873"></a>unicode.Range{0x005F, 0x005F, 1},
    <a id="L874"></a>unicode.Range{0x0061, 0x007A, 1},
    <a id="L875"></a>unicode.Range{0x00C0, 0x00D6, 1},
    <a id="L876"></a>unicode.Range{0x00D8, 0x00F6, 1},
    <a id="L877"></a>unicode.Range{0x00F8, 0x00FF, 1},
    <a id="L878"></a>unicode.Range{0x0100, 0x0131, 1},
    <a id="L879"></a>unicode.Range{0x0134, 0x013E, 1},
    <a id="L880"></a>unicode.Range{0x0141, 0x0148, 1},
    <a id="L881"></a>unicode.Range{0x014A, 0x017E, 1},
    <a id="L882"></a>unicode.Range{0x0180, 0x01C3, 1},
    <a id="L883"></a>unicode.Range{0x01CD, 0x01F0, 1},
    <a id="L884"></a>unicode.Range{0x01F4, 0x01F5, 1},
    <a id="L885"></a>unicode.Range{0x01FA, 0x0217, 1},
    <a id="L886"></a>unicode.Range{0x0250, 0x02A8, 1},
    <a id="L887"></a>unicode.Range{0x02BB, 0x02C1, 1},
    <a id="L888"></a>unicode.Range{0x0386, 0x0386, 1},
    <a id="L889"></a>unicode.Range{0x0388, 0x038A, 1},
    <a id="L890"></a>unicode.Range{0x038C, 0x038C, 1},
    <a id="L891"></a>unicode.Range{0x038E, 0x03A1, 1},
    <a id="L892"></a>unicode.Range{0x03A3, 0x03CE, 1},
    <a id="L893"></a>unicode.Range{0x03D0, 0x03D6, 1},
    <a id="L894"></a>unicode.Range{0x03DA, 0x03E0, 2},
    <a id="L895"></a>unicode.Range{0x03E2, 0x03F3, 1},
    <a id="L896"></a>unicode.Range{0x0401, 0x040C, 1},
    <a id="L897"></a>unicode.Range{0x040E, 0x044F, 1},
    <a id="L898"></a>unicode.Range{0x0451, 0x045C, 1},
    <a id="L899"></a>unicode.Range{0x045E, 0x0481, 1},
    <a id="L900"></a>unicode.Range{0x0490, 0x04C4, 1},
    <a id="L901"></a>unicode.Range{0x04C7, 0x04C8, 1},
    <a id="L902"></a>unicode.Range{0x04CB, 0x04CC, 1},
    <a id="L903"></a>unicode.Range{0x04D0, 0x04EB, 1},
    <a id="L904"></a>unicode.Range{0x04EE, 0x04F5, 1},
    <a id="L905"></a>unicode.Range{0x04F8, 0x04F9, 1},
    <a id="L906"></a>unicode.Range{0x0531, 0x0556, 1},
    <a id="L907"></a>unicode.Range{0x0559, 0x0559, 1},
    <a id="L908"></a>unicode.Range{0x0561, 0x0586, 1},
    <a id="L909"></a>unicode.Range{0x05D0, 0x05EA, 1},
    <a id="L910"></a>unicode.Range{0x05F0, 0x05F2, 1},
    <a id="L911"></a>unicode.Range{0x0621, 0x063A, 1},
    <a id="L912"></a>unicode.Range{0x0641, 0x064A, 1},
    <a id="L913"></a>unicode.Range{0x0671, 0x06B7, 1},
    <a id="L914"></a>unicode.Range{0x06BA, 0x06BE, 1},
    <a id="L915"></a>unicode.Range{0x06C0, 0x06CE, 1},
    <a id="L916"></a>unicode.Range{0x06D0, 0x06D3, 1},
    <a id="L917"></a>unicode.Range{0x06D5, 0x06D5, 1},
    <a id="L918"></a>unicode.Range{0x06E5, 0x06E6, 1},
    <a id="L919"></a>unicode.Range{0x0905, 0x0939, 1},
    <a id="L920"></a>unicode.Range{0x093D, 0x093D, 1},
    <a id="L921"></a>unicode.Range{0x0958, 0x0961, 1},
    <a id="L922"></a>unicode.Range{0x0985, 0x098C, 1},
    <a id="L923"></a>unicode.Range{0x098F, 0x0990, 1},
    <a id="L924"></a>unicode.Range{0x0993, 0x09A8, 1},
    <a id="L925"></a>unicode.Range{0x09AA, 0x09B0, 1},
    <a id="L926"></a>unicode.Range{0x09B2, 0x09B2, 1},
    <a id="L927"></a>unicode.Range{0x09B6, 0x09B9, 1},
    <a id="L928"></a>unicode.Range{0x09DC, 0x09DD, 1},
    <a id="L929"></a>unicode.Range{0x09DF, 0x09E1, 1},
    <a id="L930"></a>unicode.Range{0x09F0, 0x09F1, 1},
    <a id="L931"></a>unicode.Range{0x0A05, 0x0A0A, 1},
    <a id="L932"></a>unicode.Range{0x0A0F, 0x0A10, 1},
    <a id="L933"></a>unicode.Range{0x0A13, 0x0A28, 1},
    <a id="L934"></a>unicode.Range{0x0A2A, 0x0A30, 1},
    <a id="L935"></a>unicode.Range{0x0A32, 0x0A33, 1},
    <a id="L936"></a>unicode.Range{0x0A35, 0x0A36, 1},
    <a id="L937"></a>unicode.Range{0x0A38, 0x0A39, 1},
    <a id="L938"></a>unicode.Range{0x0A59, 0x0A5C, 1},
    <a id="L939"></a>unicode.Range{0x0A5E, 0x0A5E, 1},
    <a id="L940"></a>unicode.Range{0x0A72, 0x0A74, 1},
    <a id="L941"></a>unicode.Range{0x0A85, 0x0A8B, 1},
    <a id="L942"></a>unicode.Range{0x0A8D, 0x0A8D, 1},
    <a id="L943"></a>unicode.Range{0x0A8F, 0x0A91, 1},
    <a id="L944"></a>unicode.Range{0x0A93, 0x0AA8, 1},
    <a id="L945"></a>unicode.Range{0x0AAA, 0x0AB0, 1},
    <a id="L946"></a>unicode.Range{0x0AB2, 0x0AB3, 1},
    <a id="L947"></a>unicode.Range{0x0AB5, 0x0AB9, 1},
    <a id="L948"></a>unicode.Range{0x0ABD, 0x0AE0, 0x23},
    <a id="L949"></a>unicode.Range{0x0B05, 0x0B0C, 1},
    <a id="L950"></a>unicode.Range{0x0B0F, 0x0B10, 1},
    <a id="L951"></a>unicode.Range{0x0B13, 0x0B28, 1},
    <a id="L952"></a>unicode.Range{0x0B2A, 0x0B30, 1},
    <a id="L953"></a>unicode.Range{0x0B32, 0x0B33, 1},
    <a id="L954"></a>unicode.Range{0x0B36, 0x0B39, 1},
    <a id="L955"></a>unicode.Range{0x0B3D, 0x0B3D, 1},
    <a id="L956"></a>unicode.Range{0x0B5C, 0x0B5D, 1},
    <a id="L957"></a>unicode.Range{0x0B5F, 0x0B61, 1},
    <a id="L958"></a>unicode.Range{0x0B85, 0x0B8A, 1},
    <a id="L959"></a>unicode.Range{0x0B8E, 0x0B90, 1},
    <a id="L960"></a>unicode.Range{0x0B92, 0x0B95, 1},
    <a id="L961"></a>unicode.Range{0x0B99, 0x0B9A, 1},
    <a id="L962"></a>unicode.Range{0x0B9C, 0x0B9C, 1},
    <a id="L963"></a>unicode.Range{0x0B9E, 0x0B9F, 1},
    <a id="L964"></a>unicode.Range{0x0BA3, 0x0BA4, 1},
    <a id="L965"></a>unicode.Range{0x0BA8, 0x0BAA, 1},
    <a id="L966"></a>unicode.Range{0x0BAE, 0x0BB5, 1},
    <a id="L967"></a>unicode.Range{0x0BB7, 0x0BB9, 1},
    <a id="L968"></a>unicode.Range{0x0C05, 0x0C0C, 1},
    <a id="L969"></a>unicode.Range{0x0C0E, 0x0C10, 1},
    <a id="L970"></a>unicode.Range{0x0C12, 0x0C28, 1},
    <a id="L971"></a>unicode.Range{0x0C2A, 0x0C33, 1},
    <a id="L972"></a>unicode.Range{0x0C35, 0x0C39, 1},
    <a id="L973"></a>unicode.Range{0x0C60, 0x0C61, 1},
    <a id="L974"></a>unicode.Range{0x0C85, 0x0C8C, 1},
    <a id="L975"></a>unicode.Range{0x0C8E, 0x0C90, 1},
    <a id="L976"></a>unicode.Range{0x0C92, 0x0CA8, 1},
    <a id="L977"></a>unicode.Range{0x0CAA, 0x0CB3, 1},
    <a id="L978"></a>unicode.Range{0x0CB5, 0x0CB9, 1},
    <a id="L979"></a>unicode.Range{0x0CDE, 0x0CDE, 1},
    <a id="L980"></a>unicode.Range{0x0CE0, 0x0CE1, 1},
    <a id="L981"></a>unicode.Range{0x0D05, 0x0D0C, 1},
    <a id="L982"></a>unicode.Range{0x0D0E, 0x0D10, 1},
    <a id="L983"></a>unicode.Range{0x0D12, 0x0D28, 1},
    <a id="L984"></a>unicode.Range{0x0D2A, 0x0D39, 1},
    <a id="L985"></a>unicode.Range{0x0D60, 0x0D61, 1},
    <a id="L986"></a>unicode.Range{0x0E01, 0x0E2E, 1},
    <a id="L987"></a>unicode.Range{0x0E30, 0x0E30, 1},
    <a id="L988"></a>unicode.Range{0x0E32, 0x0E33, 1},
    <a id="L989"></a>unicode.Range{0x0E40, 0x0E45, 1},
    <a id="L990"></a>unicode.Range{0x0E81, 0x0E82, 1},
    <a id="L991"></a>unicode.Range{0x0E84, 0x0E84, 1},
    <a id="L992"></a>unicode.Range{0x0E87, 0x0E88, 1},
    <a id="L993"></a>unicode.Range{0x0E8A, 0x0E8D, 3},
    <a id="L994"></a>unicode.Range{0x0E94, 0x0E97, 1},
    <a id="L995"></a>unicode.Range{0x0E99, 0x0E9F, 1},
    <a id="L996"></a>unicode.Range{0x0EA1, 0x0EA3, 1},
    <a id="L997"></a>unicode.Range{0x0EA5, 0x0EA7, 2},
    <a id="L998"></a>unicode.Range{0x0EAA, 0x0EAB, 1},
    <a id="L999"></a>unicode.Range{0x0EAD, 0x0EAE, 1},
    <a id="L1000"></a>unicode.Range{0x0EB0, 0x0EB0, 1},
    <a id="L1001"></a>unicode.Range{0x0EB2, 0x0EB3, 1},
    <a id="L1002"></a>unicode.Range{0x0EBD, 0x0EBD, 1},
    <a id="L1003"></a>unicode.Range{0x0EC0, 0x0EC4, 1},
    <a id="L1004"></a>unicode.Range{0x0F40, 0x0F47, 1},
    <a id="L1005"></a>unicode.Range{0x0F49, 0x0F69, 1},
    <a id="L1006"></a>unicode.Range{0x10A0, 0x10C5, 1},
    <a id="L1007"></a>unicode.Range{0x10D0, 0x10F6, 1},
    <a id="L1008"></a>unicode.Range{0x1100, 0x1100, 1},
    <a id="L1009"></a>unicode.Range{0x1102, 0x1103, 1},
    <a id="L1010"></a>unicode.Range{0x1105, 0x1107, 1},
    <a id="L1011"></a>unicode.Range{0x1109, 0x1109, 1},
    <a id="L1012"></a>unicode.Range{0x110B, 0x110C, 1},
    <a id="L1013"></a>unicode.Range{0x110E, 0x1112, 1},
    <a id="L1014"></a>unicode.Range{0x113C, 0x1140, 2},
    <a id="L1015"></a>unicode.Range{0x114C, 0x1150, 2},
    <a id="L1016"></a>unicode.Range{0x1154, 0x1155, 1},
    <a id="L1017"></a>unicode.Range{0x1159, 0x1159, 1},
    <a id="L1018"></a>unicode.Range{0x115F, 0x1161, 1},
    <a id="L1019"></a>unicode.Range{0x1163, 0x1169, 2},
    <a id="L1020"></a>unicode.Range{0x116D, 0x116E, 1},
    <a id="L1021"></a>unicode.Range{0x1172, 0x1173, 1},
    <a id="L1022"></a>unicode.Range{0x1175, 0x119E, 0x119E - 0x1175},
    <a id="L1023"></a>unicode.Range{0x11A8, 0x11AB, 0x11AB - 0x11A8},
    <a id="L1024"></a>unicode.Range{0x11AE, 0x11AF, 1},
    <a id="L1025"></a>unicode.Range{0x11B7, 0x11B8, 1},
    <a id="L1026"></a>unicode.Range{0x11BA, 0x11BA, 1},
    <a id="L1027"></a>unicode.Range{0x11BC, 0x11C2, 1},
    <a id="L1028"></a>unicode.Range{0x11EB, 0x11F0, 0x11F0 - 0x11EB},
    <a id="L1029"></a>unicode.Range{0x11F9, 0x11F9, 1},
    <a id="L1030"></a>unicode.Range{0x1E00, 0x1E9B, 1},
    <a id="L1031"></a>unicode.Range{0x1EA0, 0x1EF9, 1},
    <a id="L1032"></a>unicode.Range{0x1F00, 0x1F15, 1},
    <a id="L1033"></a>unicode.Range{0x1F18, 0x1F1D, 1},
    <a id="L1034"></a>unicode.Range{0x1F20, 0x1F45, 1},
    <a id="L1035"></a>unicode.Range{0x1F48, 0x1F4D, 1},
    <a id="L1036"></a>unicode.Range{0x1F50, 0x1F57, 1},
    <a id="L1037"></a>unicode.Range{0x1F59, 0x1F5B, 0x1F5B - 0x1F59},
    <a id="L1038"></a>unicode.Range{0x1F5D, 0x1F5D, 1},
    <a id="L1039"></a>unicode.Range{0x1F5F, 0x1F7D, 1},
    <a id="L1040"></a>unicode.Range{0x1F80, 0x1FB4, 1},
    <a id="L1041"></a>unicode.Range{0x1FB6, 0x1FBC, 1},
    <a id="L1042"></a>unicode.Range{0x1FBE, 0x1FBE, 1},
    <a id="L1043"></a>unicode.Range{0x1FC2, 0x1FC4, 1},
    <a id="L1044"></a>unicode.Range{0x1FC6, 0x1FCC, 1},
    <a id="L1045"></a>unicode.Range{0x1FD0, 0x1FD3, 1},
    <a id="L1046"></a>unicode.Range{0x1FD6, 0x1FDB, 1},
    <a id="L1047"></a>unicode.Range{0x1FE0, 0x1FEC, 1},
    <a id="L1048"></a>unicode.Range{0x1FF2, 0x1FF4, 1},
    <a id="L1049"></a>unicode.Range{0x1FF6, 0x1FFC, 1},
    <a id="L1050"></a>unicode.Range{0x2126, 0x2126, 1},
    <a id="L1051"></a>unicode.Range{0x212A, 0x212B, 1},
    <a id="L1052"></a>unicode.Range{0x212E, 0x212E, 1},
    <a id="L1053"></a>unicode.Range{0x2180, 0x2182, 1},
    <a id="L1054"></a>unicode.Range{0x3007, 0x3007, 1},
    <a id="L1055"></a>unicode.Range{0x3021, 0x3029, 1},
    <a id="L1056"></a>unicode.Range{0x3041, 0x3094, 1},
    <a id="L1057"></a>unicode.Range{0x30A1, 0x30FA, 1},
    <a id="L1058"></a>unicode.Range{0x3105, 0x312C, 1},
    <a id="L1059"></a>unicode.Range{0x4E00, 0x9FA5, 1},
    <a id="L1060"></a>unicode.Range{0xAC00, 0xD7A3, 1},
<a id="L1061"></a>}

<a id="L1063"></a>var second = []unicode.Range{
    <a id="L1064"></a>unicode.Range{0x002D, 0x002E, 1},
    <a id="L1065"></a>unicode.Range{0x0030, 0x0039, 1},
    <a id="L1066"></a>unicode.Range{0x00B7, 0x00B7, 1},
    <a id="L1067"></a>unicode.Range{0x02D0, 0x02D1, 1},
    <a id="L1068"></a>unicode.Range{0x0300, 0x0345, 1},
    <a id="L1069"></a>unicode.Range{0x0360, 0x0361, 1},
    <a id="L1070"></a>unicode.Range{0x0387, 0x0387, 1},
    <a id="L1071"></a>unicode.Range{0x0483, 0x0486, 1},
    <a id="L1072"></a>unicode.Range{0x0591, 0x05A1, 1},
    <a id="L1073"></a>unicode.Range{0x05A3, 0x05B9, 1},
    <a id="L1074"></a>unicode.Range{0x05BB, 0x05BD, 1},
    <a id="L1075"></a>unicode.Range{0x05BF, 0x05BF, 1},
    <a id="L1076"></a>unicode.Range{0x05C1, 0x05C2, 1},
    <a id="L1077"></a>unicode.Range{0x05C4, 0x0640, 0x0640 - 0x05C4},
    <a id="L1078"></a>unicode.Range{0x064B, 0x0652, 1},
    <a id="L1079"></a>unicode.Range{0x0660, 0x0669, 1},
    <a id="L1080"></a>unicode.Range{0x0670, 0x0670, 1},
    <a id="L1081"></a>unicode.Range{0x06D6, 0x06DC, 1},
    <a id="L1082"></a>unicode.Range{0x06DD, 0x06DF, 1},
    <a id="L1083"></a>unicode.Range{0x06E0, 0x06E4, 1},
    <a id="L1084"></a>unicode.Range{0x06E7, 0x06E8, 1},
    <a id="L1085"></a>unicode.Range{0x06EA, 0x06ED, 1},
    <a id="L1086"></a>unicode.Range{0x06F0, 0x06F9, 1},
    <a id="L1087"></a>unicode.Range{0x0901, 0x0903, 1},
    <a id="L1088"></a>unicode.Range{0x093C, 0x093C, 1},
    <a id="L1089"></a>unicode.Range{0x093E, 0x094C, 1},
    <a id="L1090"></a>unicode.Range{0x094D, 0x094D, 1},
    <a id="L1091"></a>unicode.Range{0x0951, 0x0954, 1},
    <a id="L1092"></a>unicode.Range{0x0962, 0x0963, 1},
    <a id="L1093"></a>unicode.Range{0x0966, 0x096F, 1},
    <a id="L1094"></a>unicode.Range{0x0981, 0x0983, 1},
    <a id="L1095"></a>unicode.Range{0x09BC, 0x09BC, 1},
    <a id="L1096"></a>unicode.Range{0x09BE, 0x09BF, 1},
    <a id="L1097"></a>unicode.Range{0x09C0, 0x09C4, 1},
    <a id="L1098"></a>unicode.Range{0x09C7, 0x09C8, 1},
    <a id="L1099"></a>unicode.Range{0x09CB, 0x09CD, 1},
    <a id="L1100"></a>unicode.Range{0x09D7, 0x09D7, 1},
    <a id="L1101"></a>unicode.Range{0x09E2, 0x09E3, 1},
    <a id="L1102"></a>unicode.Range{0x09E6, 0x09EF, 1},
    <a id="L1103"></a>unicode.Range{0x0A02, 0x0A3C, 0x3A},
    <a id="L1104"></a>unicode.Range{0x0A3E, 0x0A3F, 1},
    <a id="L1105"></a>unicode.Range{0x0A40, 0x0A42, 1},
    <a id="L1106"></a>unicode.Range{0x0A47, 0x0A48, 1},
    <a id="L1107"></a>unicode.Range{0x0A4B, 0x0A4D, 1},
    <a id="L1108"></a>unicode.Range{0x0A66, 0x0A6F, 1},
    <a id="L1109"></a>unicode.Range{0x0A70, 0x0A71, 1},
    <a id="L1110"></a>unicode.Range{0x0A81, 0x0A83, 1},
    <a id="L1111"></a>unicode.Range{0x0ABC, 0x0ABC, 1},
    <a id="L1112"></a>unicode.Range{0x0ABE, 0x0AC5, 1},
    <a id="L1113"></a>unicode.Range{0x0AC7, 0x0AC9, 1},
    <a id="L1114"></a>unicode.Range{0x0ACB, 0x0ACD, 1},
    <a id="L1115"></a>unicode.Range{0x0AE6, 0x0AEF, 1},
    <a id="L1116"></a>unicode.Range{0x0B01, 0x0B03, 1},
    <a id="L1117"></a>unicode.Range{0x0B3C, 0x0B3C, 1},
    <a id="L1118"></a>unicode.Range{0x0B3E, 0x0B43, 1},
    <a id="L1119"></a>unicode.Range{0x0B47, 0x0B48, 1},
    <a id="L1120"></a>unicode.Range{0x0B4B, 0x0B4D, 1},
    <a id="L1121"></a>unicode.Range{0x0B56, 0x0B57, 1},
    <a id="L1122"></a>unicode.Range{0x0B66, 0x0B6F, 1},
    <a id="L1123"></a>unicode.Range{0x0B82, 0x0B83, 1},
    <a id="L1124"></a>unicode.Range{0x0BBE, 0x0BC2, 1},
    <a id="L1125"></a>unicode.Range{0x0BC6, 0x0BC8, 1},
    <a id="L1126"></a>unicode.Range{0x0BCA, 0x0BCD, 1},
    <a id="L1127"></a>unicode.Range{0x0BD7, 0x0BD7, 1},
    <a id="L1128"></a>unicode.Range{0x0BE7, 0x0BEF, 1},
    <a id="L1129"></a>unicode.Range{0x0C01, 0x0C03, 1},
    <a id="L1130"></a>unicode.Range{0x0C3E, 0x0C44, 1},
    <a id="L1131"></a>unicode.Range{0x0C46, 0x0C48, 1},
    <a id="L1132"></a>unicode.Range{0x0C4A, 0x0C4D, 1},
    <a id="L1133"></a>unicode.Range{0x0C55, 0x0C56, 1},
    <a id="L1134"></a>unicode.Range{0x0C66, 0x0C6F, 1},
    <a id="L1135"></a>unicode.Range{0x0C82, 0x0C83, 1},
    <a id="L1136"></a>unicode.Range{0x0CBE, 0x0CC4, 1},
    <a id="L1137"></a>unicode.Range{0x0CC6, 0x0CC8, 1},
    <a id="L1138"></a>unicode.Range{0x0CCA, 0x0CCD, 1},
    <a id="L1139"></a>unicode.Range{0x0CD5, 0x0CD6, 1},
    <a id="L1140"></a>unicode.Range{0x0CE6, 0x0CEF, 1},
    <a id="L1141"></a>unicode.Range{0x0D02, 0x0D03, 1},
    <a id="L1142"></a>unicode.Range{0x0D3E, 0x0D43, 1},
    <a id="L1143"></a>unicode.Range{0x0D46, 0x0D48, 1},
    <a id="L1144"></a>unicode.Range{0x0D4A, 0x0D4D, 1},
    <a id="L1145"></a>unicode.Range{0x0D57, 0x0D57, 1},
    <a id="L1146"></a>unicode.Range{0x0D66, 0x0D6F, 1},
    <a id="L1147"></a>unicode.Range{0x0E31, 0x0E31, 1},
    <a id="L1148"></a>unicode.Range{0x0E34, 0x0E3A, 1},
    <a id="L1149"></a>unicode.Range{0x0E46, 0x0E46, 1},
    <a id="L1150"></a>unicode.Range{0x0E47, 0x0E4E, 1},
    <a id="L1151"></a>unicode.Range{0x0E50, 0x0E59, 1},
    <a id="L1152"></a>unicode.Range{0x0EB1, 0x0EB1, 1},
    <a id="L1153"></a>unicode.Range{0x0EB4, 0x0EB9, 1},
    <a id="L1154"></a>unicode.Range{0x0EBB, 0x0EBC, 1},
    <a id="L1155"></a>unicode.Range{0x0EC6, 0x0EC6, 1},
    <a id="L1156"></a>unicode.Range{0x0EC8, 0x0ECD, 1},
    <a id="L1157"></a>unicode.Range{0x0ED0, 0x0ED9, 1},
    <a id="L1158"></a>unicode.Range{0x0F18, 0x0F19, 1},
    <a id="L1159"></a>unicode.Range{0x0F20, 0x0F29, 1},
    <a id="L1160"></a>unicode.Range{0x0F35, 0x0F39, 2},
    <a id="L1161"></a>unicode.Range{0x0F3E, 0x0F3F, 1},
    <a id="L1162"></a>unicode.Range{0x0F71, 0x0F84, 1},
    <a id="L1163"></a>unicode.Range{0x0F86, 0x0F8B, 1},
    <a id="L1164"></a>unicode.Range{0x0F90, 0x0F95, 1},
    <a id="L1165"></a>unicode.Range{0x0F97, 0x0F97, 1},
    <a id="L1166"></a>unicode.Range{0x0F99, 0x0FAD, 1},
    <a id="L1167"></a>unicode.Range{0x0FB1, 0x0FB7, 1},
    <a id="L1168"></a>unicode.Range{0x0FB9, 0x0FB9, 1},
    <a id="L1169"></a>unicode.Range{0x20D0, 0x20DC, 1},
    <a id="L1170"></a>unicode.Range{0x20E1, 0x3005, 0x3005 - 0x20E1},
    <a id="L1171"></a>unicode.Range{0x302A, 0x302F, 1},
    <a id="L1172"></a>unicode.Range{0x3031, 0x3035, 1},
    <a id="L1173"></a>unicode.Range{0x3099, 0x309A, 1},
    <a id="L1174"></a>unicode.Range{0x309D, 0x309E, 1},
    <a id="L1175"></a>unicode.Range{0x30FC, 0x30FE, 1},
<a id="L1176"></a>}

<a id="L1178"></a><span class="comment">// HTMLEntity is an entity map containing translations for the</span>
<a id="L1179"></a><span class="comment">// standard HTML entity characters.</span>
<a id="L1180"></a>var HTMLEntity = htmlEntity

<a id="L1182"></a>var htmlEntity = map[string]string{
    <a id="L1183"></a><span class="comment">/*</span>
    <a id="L1184"></a><span class="comment">	hget http://www.w3.org/TR/html4/sgml/entities.html |</span>
    <a id="L1185"></a><span class="comment">	ssam &#39;</span>
    <a id="L1186"></a><span class="comment">		,y /\&amp;gt;/ x/\&amp;lt;(.|\n)+/ s/\n/ /g</span>
    <a id="L1187"></a><span class="comment">		,x v/^\&amp;lt;!ENTITY/d</span>
    <a id="L1188"></a><span class="comment">		,s/\&amp;lt;!ENTITY ([^ ]+) .*U\+([0-9A-F][0-9A-F][0-9A-F][0-9A-F]) .+/	&#34;\1&#34;: &#34;\\u\2&#34;,/g</span>
    <a id="L1189"></a><span class="comment">	&#39;</span>
    <a id="L1190"></a><span class="comment">*/</span>
    <a id="L1191"></a>&#34;nbsp&#34;: &#34;\u00A0&#34;,
    <a id="L1192"></a>&#34;iexcl&#34;: &#34;\u00A1&#34;,
    <a id="L1193"></a>&#34;cent&#34;: &#34;\u00A2&#34;,
    <a id="L1194"></a>&#34;pound&#34;: &#34;\u00A3&#34;,
    <a id="L1195"></a>&#34;curren&#34;: &#34;\u00A4&#34;,
    <a id="L1196"></a>&#34;yen&#34;: &#34;\u00A5&#34;,
    <a id="L1197"></a>&#34;brvbar&#34;: &#34;\u00A6&#34;,
    <a id="L1198"></a>&#34;sect&#34;: &#34;\u00A7&#34;,
    <a id="L1199"></a>&#34;uml&#34;: &#34;\u00A8&#34;,
    <a id="L1200"></a>&#34;copy&#34;: &#34;\u00A9&#34;,
    <a id="L1201"></a>&#34;ordf&#34;: &#34;\u00AA&#34;,
    <a id="L1202"></a>&#34;laquo&#34;: &#34;\u00AB&#34;,
    <a id="L1203"></a>&#34;not&#34;: &#34;\u00AC&#34;,
    <a id="L1204"></a>&#34;shy&#34;: &#34;\u00AD&#34;,
    <a id="L1205"></a>&#34;reg&#34;: &#34;\u00AE&#34;,
    <a id="L1206"></a>&#34;macr&#34;: &#34;\u00AF&#34;,
    <a id="L1207"></a>&#34;deg&#34;: &#34;\u00B0&#34;,
    <a id="L1208"></a>&#34;plusmn&#34;: &#34;\u00B1&#34;,
    <a id="L1209"></a>&#34;sup2&#34;: &#34;\u00B2&#34;,
    <a id="L1210"></a>&#34;sup3&#34;: &#34;\u00B3&#34;,
    <a id="L1211"></a>&#34;acute&#34;: &#34;\u00B4&#34;,
    <a id="L1212"></a>&#34;micro&#34;: &#34;\u00B5&#34;,
    <a id="L1213"></a>&#34;para&#34;: &#34;\u00B6&#34;,
    <a id="L1214"></a>&#34;middot&#34;: &#34;\u00B7&#34;,
    <a id="L1215"></a>&#34;cedil&#34;: &#34;\u00B8&#34;,
    <a id="L1216"></a>&#34;sup1&#34;: &#34;\u00B9&#34;,
    <a id="L1217"></a>&#34;ordm&#34;: &#34;\u00BA&#34;,
    <a id="L1218"></a>&#34;raquo&#34;: &#34;\u00BB&#34;,
    <a id="L1219"></a>&#34;frac14&#34;: &#34;\u00BC&#34;,
    <a id="L1220"></a>&#34;frac12&#34;: &#34;\u00BD&#34;,
    <a id="L1221"></a>&#34;frac34&#34;: &#34;\u00BE&#34;,
    <a id="L1222"></a>&#34;iquest&#34;: &#34;\u00BF&#34;,
    <a id="L1223"></a>&#34;Agrave&#34;: &#34;\u00C0&#34;,
    <a id="L1224"></a>&#34;Aacute&#34;: &#34;\u00C1&#34;,
    <a id="L1225"></a>&#34;Acirc&#34;: &#34;\u00C2&#34;,
    <a id="L1226"></a>&#34;Atilde&#34;: &#34;\u00C3&#34;,
    <a id="L1227"></a>&#34;Auml&#34;: &#34;\u00C4&#34;,
    <a id="L1228"></a>&#34;Aring&#34;: &#34;\u00C5&#34;,
    <a id="L1229"></a>&#34;AElig&#34;: &#34;\u00C6&#34;,
    <a id="L1230"></a>&#34;Ccedil&#34;: &#34;\u00C7&#34;,
    <a id="L1231"></a>&#34;Egrave&#34;: &#34;\u00C8&#34;,
    <a id="L1232"></a>&#34;Eacute&#34;: &#34;\u00C9&#34;,
    <a id="L1233"></a>&#34;Ecirc&#34;: &#34;\u00CA&#34;,
    <a id="L1234"></a>&#34;Euml&#34;: &#34;\u00CB&#34;,
    <a id="L1235"></a>&#34;Igrave&#34;: &#34;\u00CC&#34;,
    <a id="L1236"></a>&#34;Iacute&#34;: &#34;\u00CD&#34;,
    <a id="L1237"></a>&#34;Icirc&#34;: &#34;\u00CE&#34;,
    <a id="L1238"></a>&#34;Iuml&#34;: &#34;\u00CF&#34;,
    <a id="L1239"></a>&#34;ETH&#34;: &#34;\u00D0&#34;,
    <a id="L1240"></a>&#34;Ntilde&#34;: &#34;\u00D1&#34;,
    <a id="L1241"></a>&#34;Ograve&#34;: &#34;\u00D2&#34;,
    <a id="L1242"></a>&#34;Oacute&#34;: &#34;\u00D3&#34;,
    <a id="L1243"></a>&#34;Ocirc&#34;: &#34;\u00D4&#34;,
    <a id="L1244"></a>&#34;Otilde&#34;: &#34;\u00D5&#34;,
    <a id="L1245"></a>&#34;Ouml&#34;: &#34;\u00D6&#34;,
    <a id="L1246"></a>&#34;times&#34;: &#34;\u00D7&#34;,
    <a id="L1247"></a>&#34;Oslash&#34;: &#34;\u00D8&#34;,
    <a id="L1248"></a>&#34;Ugrave&#34;: &#34;\u00D9&#34;,
    <a id="L1249"></a>&#34;Uacute&#34;: &#34;\u00DA&#34;,
    <a id="L1250"></a>&#34;Ucirc&#34;: &#34;\u00DB&#34;,
    <a id="L1251"></a>&#34;Uuml&#34;: &#34;\u00DC&#34;,
    <a id="L1252"></a>&#34;Yacute&#34;: &#34;\u00DD&#34;,
    <a id="L1253"></a>&#34;THORN&#34;: &#34;\u00DE&#34;,
    <a id="L1254"></a>&#34;szlig&#34;: &#34;\u00DF&#34;,
    <a id="L1255"></a>&#34;agrave&#34;: &#34;\u00E0&#34;,
    <a id="L1256"></a>&#34;aacute&#34;: &#34;\u00E1&#34;,
    <a id="L1257"></a>&#34;acirc&#34;: &#34;\u00E2&#34;,
    <a id="L1258"></a>&#34;atilde&#34;: &#34;\u00E3&#34;,
    <a id="L1259"></a>&#34;auml&#34;: &#34;\u00E4&#34;,
    <a id="L1260"></a>&#34;aring&#34;: &#34;\u00E5&#34;,
    <a id="L1261"></a>&#34;aelig&#34;: &#34;\u00E6&#34;,
    <a id="L1262"></a>&#34;ccedil&#34;: &#34;\u00E7&#34;,
    <a id="L1263"></a>&#34;egrave&#34;: &#34;\u00E8&#34;,
    <a id="L1264"></a>&#34;eacute&#34;: &#34;\u00E9&#34;,
    <a id="L1265"></a>&#34;ecirc&#34;: &#34;\u00EA&#34;,
    <a id="L1266"></a>&#34;euml&#34;: &#34;\u00EB&#34;,
    <a id="L1267"></a>&#34;igrave&#34;: &#34;\u00EC&#34;,
    <a id="L1268"></a>&#34;iacute&#34;: &#34;\u00ED&#34;,
    <a id="L1269"></a>&#34;icirc&#34;: &#34;\u00EE&#34;,
    <a id="L1270"></a>&#34;iuml&#34;: &#34;\u00EF&#34;,
    <a id="L1271"></a>&#34;eth&#34;: &#34;\u00F0&#34;,
    <a id="L1272"></a>&#34;ntilde&#34;: &#34;\u00F1&#34;,
    <a id="L1273"></a>&#34;ograve&#34;: &#34;\u00F2&#34;,
    <a id="L1274"></a>&#34;oacute&#34;: &#34;\u00F3&#34;,
    <a id="L1275"></a>&#34;ocirc&#34;: &#34;\u00F4&#34;,
    <a id="L1276"></a>&#34;otilde&#34;: &#34;\u00F5&#34;,
    <a id="L1277"></a>&#34;ouml&#34;: &#34;\u00F6&#34;,
    <a id="L1278"></a>&#34;divide&#34;: &#34;\u00F7&#34;,
    <a id="L1279"></a>&#34;oslash&#34;: &#34;\u00F8&#34;,
    <a id="L1280"></a>&#34;ugrave&#34;: &#34;\u00F9&#34;,
    <a id="L1281"></a>&#34;uacute&#34;: &#34;\u00FA&#34;,
    <a id="L1282"></a>&#34;ucirc&#34;: &#34;\u00FB&#34;,
    <a id="L1283"></a>&#34;uuml&#34;: &#34;\u00FC&#34;,
    <a id="L1284"></a>&#34;yacute&#34;: &#34;\u00FD&#34;,
    <a id="L1285"></a>&#34;thorn&#34;: &#34;\u00FE&#34;,
    <a id="L1286"></a>&#34;yuml&#34;: &#34;\u00FF&#34;,
    <a id="L1287"></a>&#34;fnof&#34;: &#34;\u0192&#34;,
    <a id="L1288"></a>&#34;Alpha&#34;: &#34;\u0391&#34;,
    <a id="L1289"></a>&#34;Beta&#34;: &#34;\u0392&#34;,
    <a id="L1290"></a>&#34;Gamma&#34;: &#34;\u0393&#34;,
    <a id="L1291"></a>&#34;Delta&#34;: &#34;\u0394&#34;,
    <a id="L1292"></a>&#34;Epsilon&#34;: &#34;\u0395&#34;,
    <a id="L1293"></a>&#34;Zeta&#34;: &#34;\u0396&#34;,
    <a id="L1294"></a>&#34;Eta&#34;: &#34;\u0397&#34;,
    <a id="L1295"></a>&#34;Theta&#34;: &#34;\u0398&#34;,
    <a id="L1296"></a>&#34;Iota&#34;: &#34;\u0399&#34;,
    <a id="L1297"></a>&#34;Kappa&#34;: &#34;\u039A&#34;,
    <a id="L1298"></a>&#34;Lambda&#34;: &#34;\u039B&#34;,
    <a id="L1299"></a>&#34;Mu&#34;: &#34;\u039C&#34;,
    <a id="L1300"></a>&#34;Nu&#34;: &#34;\u039D&#34;,
    <a id="L1301"></a>&#34;Xi&#34;: &#34;\u039E&#34;,
    <a id="L1302"></a>&#34;Omicron&#34;: &#34;\u039F&#34;,
    <a id="L1303"></a>&#34;Pi&#34;: &#34;\u03A0&#34;,
    <a id="L1304"></a>&#34;Rho&#34;: &#34;\u03A1&#34;,
    <a id="L1305"></a>&#34;Sigma&#34;: &#34;\u03A3&#34;,
    <a id="L1306"></a>&#34;Tau&#34;: &#34;\u03A4&#34;,
    <a id="L1307"></a>&#34;Upsilon&#34;: &#34;\u03A5&#34;,
    <a id="L1308"></a>&#34;Phi&#34;: &#34;\u03A6&#34;,
    <a id="L1309"></a>&#34;Chi&#34;: &#34;\u03A7&#34;,
    <a id="L1310"></a>&#34;Psi&#34;: &#34;\u03A8&#34;,
    <a id="L1311"></a>&#34;Omega&#34;: &#34;\u03A9&#34;,
    <a id="L1312"></a>&#34;alpha&#34;: &#34;\u03B1&#34;,
    <a id="L1313"></a>&#34;beta&#34;: &#34;\u03B2&#34;,
    <a id="L1314"></a>&#34;gamma&#34;: &#34;\u03B3&#34;,
    <a id="L1315"></a>&#34;delta&#34;: &#34;\u03B4&#34;,
    <a id="L1316"></a>&#34;epsilon&#34;: &#34;\u03B5&#34;,
    <a id="L1317"></a>&#34;zeta&#34;: &#34;\u03B6&#34;,
    <a id="L1318"></a>&#34;eta&#34;: &#34;\u03B7&#34;,
    <a id="L1319"></a>&#34;theta&#34;: &#34;\u03B8&#34;,
    <a id="L1320"></a>&#34;iota&#34;: &#34;\u03B9&#34;,
    <a id="L1321"></a>&#34;kappa&#34;: &#34;\u03BA&#34;,
    <a id="L1322"></a>&#34;lambda&#34;: &#34;\u03BB&#34;,
    <a id="L1323"></a>&#34;mu&#34;: &#34;\u03BC&#34;,
    <a id="L1324"></a>&#34;nu&#34;: &#34;\u03BD&#34;,
    <a id="L1325"></a>&#34;xi&#34;: &#34;\u03BE&#34;,
    <a id="L1326"></a>&#34;omicron&#34;: &#34;\u03BF&#34;,
    <a id="L1327"></a>&#34;pi&#34;: &#34;\u03C0&#34;,
    <a id="L1328"></a>&#34;rho&#34;: &#34;\u03C1&#34;,
    <a id="L1329"></a>&#34;sigmaf&#34;: &#34;\u03C2&#34;,
    <a id="L1330"></a>&#34;sigma&#34;: &#34;\u03C3&#34;,
    <a id="L1331"></a>&#34;tau&#34;: &#34;\u03C4&#34;,
    <a id="L1332"></a>&#34;upsilon&#34;: &#34;\u03C5&#34;,
    <a id="L1333"></a>&#34;phi&#34;: &#34;\u03C6&#34;,
    <a id="L1334"></a>&#34;chi&#34;: &#34;\u03C7&#34;,
    <a id="L1335"></a>&#34;psi&#34;: &#34;\u03C8&#34;,
    <a id="L1336"></a>&#34;omega&#34;: &#34;\u03C9&#34;,
    <a id="L1337"></a>&#34;thetasym&#34;: &#34;\u03D1&#34;,
    <a id="L1338"></a>&#34;upsih&#34;: &#34;\u03D2&#34;,
    <a id="L1339"></a>&#34;piv&#34;: &#34;\u03D6&#34;,
    <a id="L1340"></a>&#34;bull&#34;: &#34;\u2022&#34;,
    <a id="L1341"></a>&#34;hellip&#34;: &#34;\u2026&#34;,
    <a id="L1342"></a>&#34;prime&#34;: &#34;\u2032&#34;,
    <a id="L1343"></a>&#34;Prime&#34;: &#34;\u2033&#34;,
    <a id="L1344"></a>&#34;oline&#34;: &#34;\u203E&#34;,
    <a id="L1345"></a>&#34;frasl&#34;: &#34;\u2044&#34;,
    <a id="L1346"></a>&#34;weierp&#34;: &#34;\u2118&#34;,
    <a id="L1347"></a>&#34;image&#34;: &#34;\u2111&#34;,
    <a id="L1348"></a>&#34;real&#34;: &#34;\u211C&#34;,
    <a id="L1349"></a>&#34;trade&#34;: &#34;\u2122&#34;,
    <a id="L1350"></a>&#34;alefsym&#34;: &#34;\u2135&#34;,
    <a id="L1351"></a>&#34;larr&#34;: &#34;\u2190&#34;,
    <a id="L1352"></a>&#34;uarr&#34;: &#34;\u2191&#34;,
    <a id="L1353"></a>&#34;rarr&#34;: &#34;\u2192&#34;,
    <a id="L1354"></a>&#34;darr&#34;: &#34;\u2193&#34;,
    <a id="L1355"></a>&#34;harr&#34;: &#34;\u2194&#34;,
    <a id="L1356"></a>&#34;crarr&#34;: &#34;\u21B5&#34;,
    <a id="L1357"></a>&#34;lArr&#34;: &#34;\u21D0&#34;,
    <a id="L1358"></a>&#34;uArr&#34;: &#34;\u21D1&#34;,
    <a id="L1359"></a>&#34;rArr&#34;: &#34;\u21D2&#34;,
    <a id="L1360"></a>&#34;dArr&#34;: &#34;\u21D3&#34;,
    <a id="L1361"></a>&#34;hArr&#34;: &#34;\u21D4&#34;,
    <a id="L1362"></a>&#34;forall&#34;: &#34;\u2200&#34;,
    <a id="L1363"></a>&#34;part&#34;: &#34;\u2202&#34;,
    <a id="L1364"></a>&#34;exist&#34;: &#34;\u2203&#34;,
    <a id="L1365"></a>&#34;empty&#34;: &#34;\u2205&#34;,
    <a id="L1366"></a>&#34;nabla&#34;: &#34;\u2207&#34;,
    <a id="L1367"></a>&#34;isin&#34;: &#34;\u2208&#34;,
    <a id="L1368"></a>&#34;notin&#34;: &#34;\u2209&#34;,
    <a id="L1369"></a>&#34;ni&#34;: &#34;\u220B&#34;,
    <a id="L1370"></a>&#34;prod&#34;: &#34;\u220F&#34;,
    <a id="L1371"></a>&#34;sum&#34;: &#34;\u2211&#34;,
    <a id="L1372"></a>&#34;minus&#34;: &#34;\u2212&#34;,
    <a id="L1373"></a>&#34;lowast&#34;: &#34;\u2217&#34;,
    <a id="L1374"></a>&#34;radic&#34;: &#34;\u221A&#34;,
    <a id="L1375"></a>&#34;prop&#34;: &#34;\u221D&#34;,
    <a id="L1376"></a>&#34;infin&#34;: &#34;\u221E&#34;,
    <a id="L1377"></a>&#34;ang&#34;: &#34;\u2220&#34;,
    <a id="L1378"></a>&#34;and&#34;: &#34;\u2227&#34;,
    <a id="L1379"></a>&#34;or&#34;: &#34;\u2228&#34;,
    <a id="L1380"></a>&#34;cap&#34;: &#34;\u2229&#34;,
    <a id="L1381"></a>&#34;cup&#34;: &#34;\u222A&#34;,
    <a id="L1382"></a>&#34;int&#34;: &#34;\u222B&#34;,
    <a id="L1383"></a>&#34;there4&#34;: &#34;\u2234&#34;,
    <a id="L1384"></a>&#34;sim&#34;: &#34;\u223C&#34;,
    <a id="L1385"></a>&#34;cong&#34;: &#34;\u2245&#34;,
    <a id="L1386"></a>&#34;asymp&#34;: &#34;\u2248&#34;,
    <a id="L1387"></a>&#34;ne&#34;: &#34;\u2260&#34;,
    <a id="L1388"></a>&#34;equiv&#34;: &#34;\u2261&#34;,
    <a id="L1389"></a>&#34;le&#34;: &#34;\u2264&#34;,
    <a id="L1390"></a>&#34;ge&#34;: &#34;\u2265&#34;,
    <a id="L1391"></a>&#34;sub&#34;: &#34;\u2282&#34;,
    <a id="L1392"></a>&#34;sup&#34;: &#34;\u2283&#34;,
    <a id="L1393"></a>&#34;nsub&#34;: &#34;\u2284&#34;,
    <a id="L1394"></a>&#34;sube&#34;: &#34;\u2286&#34;,
    <a id="L1395"></a>&#34;supe&#34;: &#34;\u2287&#34;,
    <a id="L1396"></a>&#34;oplus&#34;: &#34;\u2295&#34;,
    <a id="L1397"></a>&#34;otimes&#34;: &#34;\u2297&#34;,
    <a id="L1398"></a>&#34;perp&#34;: &#34;\u22A5&#34;,
    <a id="L1399"></a>&#34;sdot&#34;: &#34;\u22C5&#34;,
    <a id="L1400"></a>&#34;lceil&#34;: &#34;\u2308&#34;,
    <a id="L1401"></a>&#34;rceil&#34;: &#34;\u2309&#34;,
    <a id="L1402"></a>&#34;lfloor&#34;: &#34;\u230A&#34;,
    <a id="L1403"></a>&#34;rfloor&#34;: &#34;\u230B&#34;,
    <a id="L1404"></a>&#34;lang&#34;: &#34;\u2329&#34;,
    <a id="L1405"></a>&#34;rang&#34;: &#34;\u232A&#34;,
    <a id="L1406"></a>&#34;loz&#34;: &#34;\u25CA&#34;,
    <a id="L1407"></a>&#34;spades&#34;: &#34;\u2660&#34;,
    <a id="L1408"></a>&#34;clubs&#34;: &#34;\u2663&#34;,
    <a id="L1409"></a>&#34;hearts&#34;: &#34;\u2665&#34;,
    <a id="L1410"></a>&#34;diams&#34;: &#34;\u2666&#34;,
    <a id="L1411"></a>&#34;quot&#34;: &#34;\u0022&#34;,
    <a id="L1412"></a>&#34;amp&#34;: &#34;\u0026&#34;,
    <a id="L1413"></a>&#34;lt&#34;: &#34;\u003C&#34;,
    <a id="L1414"></a>&#34;gt&#34;: &#34;\u003E&#34;,
    <a id="L1415"></a>&#34;OElig&#34;: &#34;\u0152&#34;,
    <a id="L1416"></a>&#34;oelig&#34;: &#34;\u0153&#34;,
    <a id="L1417"></a>&#34;Scaron&#34;: &#34;\u0160&#34;,
    <a id="L1418"></a>&#34;scaron&#34;: &#34;\u0161&#34;,
    <a id="L1419"></a>&#34;Yuml&#34;: &#34;\u0178&#34;,
    <a id="L1420"></a>&#34;circ&#34;: &#34;\u02C6&#34;,
    <a id="L1421"></a>&#34;tilde&#34;: &#34;\u02DC&#34;,
    <a id="L1422"></a>&#34;ensp&#34;: &#34;\u2002&#34;,
    <a id="L1423"></a>&#34;emsp&#34;: &#34;\u2003&#34;,
    <a id="L1424"></a>&#34;thinsp&#34;: &#34;\u2009&#34;,
    <a id="L1425"></a>&#34;zwnj&#34;: &#34;\u200C&#34;,
    <a id="L1426"></a>&#34;zwj&#34;: &#34;\u200D&#34;,
    <a id="L1427"></a>&#34;lrm&#34;: &#34;\u200E&#34;,
    <a id="L1428"></a>&#34;rlm&#34;: &#34;\u200F&#34;,
    <a id="L1429"></a>&#34;ndash&#34;: &#34;\u2013&#34;,
    <a id="L1430"></a>&#34;mdash&#34;: &#34;\u2014&#34;,
    <a id="L1431"></a>&#34;lsquo&#34;: &#34;\u2018&#34;,
    <a id="L1432"></a>&#34;rsquo&#34;: &#34;\u2019&#34;,
    <a id="L1433"></a>&#34;sbquo&#34;: &#34;\u201A&#34;,
    <a id="L1434"></a>&#34;ldquo&#34;: &#34;\u201C&#34;,
    <a id="L1435"></a>&#34;rdquo&#34;: &#34;\u201D&#34;,
    <a id="L1436"></a>&#34;bdquo&#34;: &#34;\u201E&#34;,
    <a id="L1437"></a>&#34;dagger&#34;: &#34;\u2020&#34;,
    <a id="L1438"></a>&#34;Dagger&#34;: &#34;\u2021&#34;,
    <a id="L1439"></a>&#34;permil&#34;: &#34;\u2030&#34;,
    <a id="L1440"></a>&#34;lsaquo&#34;: &#34;\u2039&#34;,
    <a id="L1441"></a>&#34;rsaquo&#34;: &#34;\u203A&#34;,
    <a id="L1442"></a>&#34;euro&#34;: &#34;\u20AC&#34;,
<a id="L1443"></a>}

<a id="L1445"></a><span class="comment">// HTMLAutoClose is the set of HTML elements that</span>
<a id="L1446"></a><span class="comment">// should be considered to close automatically.</span>
<a id="L1447"></a>var HTMLAutoClose = htmlAutoClose

<a id="L1449"></a>var htmlAutoClose = []string{
    <a id="L1450"></a><span class="comment">/*</span>
    <a id="L1451"></a><span class="comment">	hget http://www.w3.org/TR/html4/loose.dtd |</span>
    <a id="L1452"></a><span class="comment">	9 sed -n &#39;s/&lt;!ELEMENT (.*) - O EMPTY.+/	&#34;\1&#34;,/p&#39; | tr A-Z a-z</span>
    <a id="L1453"></a><span class="comment">*/</span>
    <a id="L1454"></a>&#34;basefont&#34;,
    <a id="L1455"></a>&#34;br&#34;,
    <a id="L1456"></a>&#34;area&#34;,
    <a id="L1457"></a>&#34;link&#34;,
    <a id="L1458"></a>&#34;img&#34;,
    <a id="L1459"></a>&#34;param&#34;,
    <a id="L1460"></a>&#34;hr&#34;,
    <a id="L1461"></a>&#34;input&#34;,
    <a id="L1462"></a>&#34;col     &#34;,
    <a id="L1463"></a>&#34;frame&#34;,
    <a id="L1464"></a>&#34;isindex&#34;,
    <a id="L1465"></a>&#34;base&#34;,
    <a id="L1466"></a>&#34;meta&#34;,
<a id="L1467"></a>}
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
