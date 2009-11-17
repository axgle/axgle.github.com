<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/dwarf/entry.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/debug/dwarf/entry.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// DWARF debug information entry parser.</span>
<a id="L6"></a><span class="comment">// An entry is a sequence of data items of a given format.</span>
<a id="L7"></a><span class="comment">// The first word in the entry is an index into what DWARF</span>
<a id="L8"></a><span class="comment">// calls the ``abbreviation table.&#39;&#39;  An abbreviation is really</span>
<a id="L9"></a><span class="comment">// just a type descriptor: it&#39;s an array of attribute tag/value format pairs.</span>

<a id="L11"></a>package dwarf

<a id="L13"></a>import &#34;os&#34;

<a id="L15"></a><span class="comment">// a single entry&#39;s description: a sequence of attributes</span>
<a id="L16"></a>type abbrev struct {
    <a id="L17"></a>tag      Tag;
    <a id="L18"></a>children bool;
    <a id="L19"></a>field    []afield;
<a id="L20"></a>}

<a id="L22"></a>type afield struct {
    <a id="L23"></a>attr Attr;
    <a id="L24"></a>fmt  format;
<a id="L25"></a>}

<a id="L27"></a><span class="comment">// a map from entry format ids to their descriptions</span>
<a id="L28"></a>type abbrevTable map[uint32]abbrev

<a id="L30"></a><span class="comment">// ParseAbbrev returns the abbreviation table that starts at byte off</span>
<a id="L31"></a><span class="comment">// in the .debug_abbrev section.</span>
<a id="L32"></a>func (d *Data) parseAbbrev(off uint32) (abbrevTable, os.Error) {
    <a id="L33"></a>if m, ok := d.abbrevCache[off]; ok {
        <a id="L34"></a>return m, nil
    <a id="L35"></a>}

    <a id="L37"></a>data := d.abbrev;
    <a id="L38"></a>if off &gt; uint32(len(data)) {
        <a id="L39"></a>data = nil
    <a id="L40"></a>} else {
        <a id="L41"></a>data = data[off:len(data)]
    <a id="L42"></a>}
    <a id="L43"></a>b := makeBuf(d, &#34;abbrev&#34;, 0, data, 0);

    <a id="L45"></a><span class="comment">// Error handling is simplified by the buf getters</span>
    <a id="L46"></a><span class="comment">// returning an endless stream of 0s after an error.</span>
    <a id="L47"></a>m := make(abbrevTable);
    <a id="L48"></a>for {
        <a id="L49"></a><span class="comment">// Table ends with id == 0.</span>
        <a id="L50"></a>id := uint32(b.uint());
        <a id="L51"></a>if id == 0 {
            <a id="L52"></a>break
        <a id="L53"></a>}

        <a id="L55"></a><span class="comment">// Walk over attributes, counting.</span>
        <a id="L56"></a>n := 0;
        <a id="L57"></a>b1 := b; <span class="comment">// Read from copy of b.</span>
        <a id="L58"></a>b1.uint();
        <a id="L59"></a>b1.uint8();
        <a id="L60"></a>for {
            <a id="L61"></a>tag := b1.uint();
            <a id="L62"></a>fmt := b1.uint();
            <a id="L63"></a>if tag == 0 &amp;&amp; fmt == 0 {
                <a id="L64"></a>break
            <a id="L65"></a>}
            <a id="L66"></a>n++;
        <a id="L67"></a>}
        <a id="L68"></a>if b1.err != nil {
            <a id="L69"></a>return nil, b1.err
        <a id="L70"></a>}

        <a id="L72"></a><span class="comment">// Walk over attributes again, this time writing them down.</span>
        <a id="L73"></a>var a abbrev;
        <a id="L74"></a>a.tag = Tag(b.uint());
        <a id="L75"></a>a.children = b.uint8() != 0;
        <a id="L76"></a>a.field = make([]afield, n);
        <a id="L77"></a>for i := range a.field {
            <a id="L78"></a>a.field[i].attr = Attr(b.uint());
            <a id="L79"></a>a.field[i].fmt = format(b.uint());
        <a id="L80"></a>}
        <a id="L81"></a>b.uint();
        <a id="L82"></a>b.uint();

        <a id="L84"></a>m[id] = a;
    <a id="L85"></a>}
    <a id="L86"></a>if b.err != nil {
        <a id="L87"></a>return nil, b.err
    <a id="L88"></a>}
    <a id="L89"></a>d.abbrevCache[off] = m;
    <a id="L90"></a>return m, nil;
<a id="L91"></a>}

<a id="L93"></a><span class="comment">// An entry is a sequence of attribute/value pairs.</span>
<a id="L94"></a>type Entry struct {
    <a id="L95"></a>Offset   Offset; <span class="comment">// offset of Entry in DWARF info</span>
    <a id="L96"></a>Tag      Tag;    <span class="comment">// tag (kind of Entry)</span>
    <a id="L97"></a>Children bool;   <span class="comment">// whether Entry is followed by children</span>
    <a id="L98"></a>Field    []Field;
<a id="L99"></a>}

<a id="L101"></a><span class="comment">// A Field is a single attribute/value pair in an Entry.</span>
<a id="L102"></a>type Field struct {
    <a id="L103"></a>Attr Attr;
    <a id="L104"></a>Val  interface{};
<a id="L105"></a>}

<a id="L107"></a><span class="comment">// Val returns the value associated with attribute Attr in Entry,</span>
<a id="L108"></a><span class="comment">// or nil if there is no such attribute.</span>
<a id="L109"></a><span class="comment">//</span>
<a id="L110"></a><span class="comment">// A common idiom is to merge the check for nil return with</span>
<a id="L111"></a><span class="comment">// the check that the value has the expected dynamic type, as in:</span>
<a id="L112"></a><span class="comment">//	v, ok := e.Val(AttrSibling).(int64);</span>
<a id="L113"></a><span class="comment">//</span>
<a id="L114"></a>func (e *Entry) Val(a Attr) interface{} {
    <a id="L115"></a>for _, f := range e.Field {
        <a id="L116"></a>if f.Attr == a {
            <a id="L117"></a>return f.Val
        <a id="L118"></a>}
    <a id="L119"></a>}
    <a id="L120"></a>return nil;
<a id="L121"></a>}

<a id="L123"></a><span class="comment">// An Offset represents the location of an Entry within the DWARF info.</span>
<a id="L124"></a><span class="comment">// (See Reader.Seek.)</span>
<a id="L125"></a>type Offset uint32

<a id="L127"></a><span class="comment">// Entry reads a single entry from buf, decoding</span>
<a id="L128"></a><span class="comment">// according to the given abbreviation table.</span>
<a id="L129"></a>func (b *buf) entry(atab abbrevTable, ubase Offset) *Entry {
    <a id="L130"></a>off := b.off;
    <a id="L131"></a>id := uint32(b.uint());
    <a id="L132"></a>if id == 0 {
        <a id="L133"></a>return &amp;Entry{}
    <a id="L134"></a>}
    <a id="L135"></a>a, ok := atab[id];
    <a id="L136"></a>if !ok {
        <a id="L137"></a>b.error(&#34;unknown abbreviation table index&#34;);
        <a id="L138"></a>return nil;
    <a id="L139"></a>}
    <a id="L140"></a>e := &amp;Entry{
        <a id="L141"></a>Offset: off,
        <a id="L142"></a>Tag: a.tag,
        <a id="L143"></a>Children: a.children,
        <a id="L144"></a>Field: make([]Field, len(a.field)),
    <a id="L145"></a>};
    <a id="L146"></a>for i := range e.Field {
        <a id="L147"></a>e.Field[i].Attr = a.field[i].attr;
        <a id="L148"></a>fmt := a.field[i].fmt;
        <a id="L149"></a>if fmt == formIndirect {
            <a id="L150"></a>fmt = format(b.uint())
        <a id="L151"></a>}
        <a id="L152"></a>var val interface{}
        <a id="L153"></a>switch fmt {
        <a id="L154"></a>default:
            <a id="L155"></a>b.error(&#34;unknown entry attr format&#34;)

        <a id="L157"></a><span class="comment">// address</span>
        <a id="L158"></a>case formAddr:
            <a id="L159"></a>val = b.addr()

        <a id="L161"></a><span class="comment">// block</span>
        <a id="L162"></a>case formDwarfBlock1:
            <a id="L163"></a>val = b.bytes(int(b.uint8()))
        <a id="L164"></a>case formDwarfBlock2:
            <a id="L165"></a>val = b.bytes(int(b.uint16()))
        <a id="L166"></a>case formDwarfBlock4:
            <a id="L167"></a>val = b.bytes(int(b.uint32()))
        <a id="L168"></a>case formDwarfBlock:
            <a id="L169"></a>val = b.bytes(int(b.uint()))

        <a id="L171"></a><span class="comment">// constant</span>
        <a id="L172"></a>case formData1:
            <a id="L173"></a>val = int64(b.uint8())
        <a id="L174"></a>case formData2:
            <a id="L175"></a>val = int64(b.uint16())
        <a id="L176"></a>case formData4:
            <a id="L177"></a>val = int64(b.uint32())
        <a id="L178"></a>case formData8:
            <a id="L179"></a>val = int64(b.uint64())
        <a id="L180"></a>case formSdata:
            <a id="L181"></a>val = int64(b.int())
        <a id="L182"></a>case formUdata:
            <a id="L183"></a>val = int64(b.uint())

        <a id="L185"></a><span class="comment">// flag</span>
        <a id="L186"></a>case formFlag:
            <a id="L187"></a>val = b.uint8() == 1

        <a id="L189"></a><span class="comment">// reference to other entry</span>
        <a id="L190"></a>case formRefAddr:
            <a id="L191"></a>val = Offset(b.addr())
        <a id="L192"></a>case formRef1:
            <a id="L193"></a>val = Offset(b.uint8()) + ubase
        <a id="L194"></a>case formRef2:
            <a id="L195"></a>val = Offset(b.uint16()) + ubase
        <a id="L196"></a>case formRef4:
            <a id="L197"></a>val = Offset(b.uint32()) + ubase
        <a id="L198"></a>case formRef8:
            <a id="L199"></a>val = Offset(b.uint64()) + ubase
        <a id="L200"></a>case formRefUdata:
            <a id="L201"></a>val = Offset(b.uint()) + ubase

        <a id="L203"></a><span class="comment">// string</span>
        <a id="L204"></a>case formString:
            <a id="L205"></a>val = b.string()
        <a id="L206"></a>case formStrp:
            <a id="L207"></a>off := b.uint32(); <span class="comment">// offset into .debug_str</span>
            <a id="L208"></a>if b.err != nil {
                <a id="L209"></a>return nil
            <a id="L210"></a>}
            <a id="L211"></a>b1 := makeBuf(b.dwarf, &#34;str&#34;, 0, b.dwarf.str, 0);
            <a id="L212"></a>b1.skip(int(off));
            <a id="L213"></a>val = b1.string();
            <a id="L214"></a>if b1.err != nil {
                <a id="L215"></a>b.err = b1.err;
                <a id="L216"></a>return nil;
            <a id="L217"></a>}
        <a id="L218"></a>}
        <a id="L219"></a>e.Field[i].Val = val;
    <a id="L220"></a>}
    <a id="L221"></a>if b.err != nil {
        <a id="L222"></a>return nil
    <a id="L223"></a>}
    <a id="L224"></a>return e;
<a id="L225"></a>}

<a id="L227"></a><span class="comment">// A Reader allows reading Entry structures from a DWARF ``info&#39;&#39; section.</span>
<a id="L228"></a><span class="comment">// The Entry structures are arranged in a tree.  The Reader&#39;s Next function</span>
<a id="L229"></a><span class="comment">// return successive entries from a pre-order traversal of the tree.</span>
<a id="L230"></a><span class="comment">// If an entry has children, its Children field will be true, and the children</span>
<a id="L231"></a><span class="comment">// follow, terminated by an Entry with Tag 0.</span>
<a id="L232"></a>type Reader struct {
    <a id="L233"></a>b            buf;
    <a id="L234"></a>d            *Data;
    <a id="L235"></a>err          os.Error;
    <a id="L236"></a>unit         int;
    <a id="L237"></a>lastChildren bool;   <span class="comment">// .Children of last entry returned by Next</span>
    <a id="L238"></a>lastSibling  Offset; <span class="comment">// .Val(AttrSibling) of last entry returned by Next</span>
<a id="L239"></a>}

<a id="L241"></a><span class="comment">// Reader returns a new Reader for Data.</span>
<a id="L242"></a><span class="comment">// The reader is positioned at byte offset 0 in the DWARF ``info&#39;&#39; section.</span>
<a id="L243"></a>func (d *Data) Reader() *Reader {
    <a id="L244"></a>r := &amp;Reader{d: d};
    <a id="L245"></a>r.Seek(0);
    <a id="L246"></a>return r;
<a id="L247"></a>}

<a id="L249"></a><span class="comment">// Seek positions the Reader at offset off in the encoded entry stream.</span>
<a id="L250"></a><span class="comment">// Offset 0 can be used to denote the first entry.</span>
<a id="L251"></a>func (r *Reader) Seek(off Offset) {
    <a id="L252"></a>d := r.d;
    <a id="L253"></a>r.err = nil;
    <a id="L254"></a>r.lastChildren = false;
    <a id="L255"></a>if off == 0 {
        <a id="L256"></a>if len(d.unit) == 0 {
            <a id="L257"></a>return
        <a id="L258"></a>}
        <a id="L259"></a>u := &amp;d.unit[0];
        <a id="L260"></a>r.unit = 0;
        <a id="L261"></a>r.b = makeBuf(r.d, &#34;info&#34;, u.off, u.data, u.addrsize);
        <a id="L262"></a>return;
    <a id="L263"></a>}

    <a id="L265"></a><span class="comment">// TODO(rsc): binary search (maybe a new package)</span>
    <a id="L266"></a>var i int;
    <a id="L267"></a>var u *unit;
    <a id="L268"></a>for i = range d.unit {
        <a id="L269"></a>u = &amp;d.unit[i];
        <a id="L270"></a>if u.off &lt;= off &amp;&amp; off &lt; u.off+Offset(len(u.data)) {
            <a id="L271"></a>r.unit = i;
            <a id="L272"></a>r.b = makeBuf(r.d, &#34;info&#34;, off, u.data[off-u.off:len(u.data)], u.addrsize);
            <a id="L273"></a>return;
        <a id="L274"></a>}
    <a id="L275"></a>}
    <a id="L276"></a>r.err = os.NewError(&#34;offset out of range&#34;);
<a id="L277"></a>}

<a id="L279"></a><span class="comment">// maybeNextUnit advances to the next unit if this one is finished.</span>
<a id="L280"></a>func (r *Reader) maybeNextUnit() {
    <a id="L281"></a>for len(r.b.data) == 0 &amp;&amp; r.unit+1 &lt; len(r.d.unit) {
        <a id="L282"></a>r.unit++;
        <a id="L283"></a>u := &amp;r.d.unit[r.unit];
        <a id="L284"></a>r.b = makeBuf(r.d, &#34;info&#34;, u.off, u.data, u.addrsize);
    <a id="L285"></a>}
<a id="L286"></a>}

<a id="L288"></a><span class="comment">// Next reads the next entry from the encoded entry stream.</span>
<a id="L289"></a><span class="comment">// It returns nil, nil when it reaches the end of the section.</span>
<a id="L290"></a><span class="comment">// It returns an error if the current offset is invalid or the data at the</span>
<a id="L291"></a><span class="comment">// offset cannot be decoded as a valid Entry.</span>
<a id="L292"></a>func (r *Reader) Next() (*Entry, os.Error) {
    <a id="L293"></a>if r.err != nil {
        <a id="L294"></a>return nil, r.err
    <a id="L295"></a>}
    <a id="L296"></a>r.maybeNextUnit();
    <a id="L297"></a>if len(r.b.data) == 0 {
        <a id="L298"></a>return nil, nil
    <a id="L299"></a>}
    <a id="L300"></a>u := &amp;r.d.unit[r.unit];
    <a id="L301"></a>e := r.b.entry(u.atable, u.base);
    <a id="L302"></a>if r.b.err != nil {
        <a id="L303"></a>r.err = r.b.err;
        <a id="L304"></a>return nil, r.err;
    <a id="L305"></a>}
    <a id="L306"></a>if e != nil {
        <a id="L307"></a>r.lastChildren = e.Children;
        <a id="L308"></a>if r.lastChildren {
            <a id="L309"></a>r.lastSibling, _ = e.Val(AttrSibling).(Offset)
        <a id="L310"></a>}
    <a id="L311"></a>} else {
        <a id="L312"></a>r.lastChildren = false
    <a id="L313"></a>}
    <a id="L314"></a>return e, nil;
<a id="L315"></a>}

<a id="L317"></a><span class="comment">// SkipChildren skips over the child entries associated with</span>
<a id="L318"></a><span class="comment">// the last Entry returned by Next.  If that Entry did not have</span>
<a id="L319"></a><span class="comment">// children or Next has not been called, SkipChildren is a no-op.</span>
<a id="L320"></a>func (r *Reader) SkipChildren() {
    <a id="L321"></a>if r.err != nil || !r.lastChildren {
        <a id="L322"></a>return
    <a id="L323"></a>}

    <a id="L325"></a><span class="comment">// If the last entry had a sibling attribute,</span>
    <a id="L326"></a><span class="comment">// that attribute gives the offset of the next</span>
    <a id="L327"></a><span class="comment">// sibling, so we can avoid decoding the</span>
    <a id="L328"></a><span class="comment">// child subtrees.</span>
    <a id="L329"></a>if r.lastSibling &gt;= r.b.off {
        <a id="L330"></a>r.Seek(r.lastSibling);
        <a id="L331"></a>return;
    <a id="L332"></a>}

    <a id="L334"></a>for {
        <a id="L335"></a>e, err := r.Next();
        <a id="L336"></a>if err != nil || e == nil || e.Tag == 0 {
            <a id="L337"></a>break
        <a id="L338"></a>}
        <a id="L339"></a>if e.Children {
            <a id="L340"></a>r.SkipChildren()
        <a id="L341"></a>}
    <a id="L342"></a>}
<a id="L343"></a>}
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
