<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/patch/patch.go</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/patch/patch.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Package patch implements parsing and execution of the textual and</span>
<a id="L6"></a><span class="comment">// binary patch descriptions used by version control tools such as</span>
<a id="L7"></a><span class="comment">// CVS, Git, Mercurial, and Subversion.</span>
<a id="L8"></a>package patch

<a id="L10"></a>import (
    <a id="L11"></a>&#34;bytes&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;path&#34;;
    <a id="L14"></a>&#34;strings&#34;;
<a id="L15"></a>)

<a id="L17"></a><span class="comment">// A Set represents a set of patches to be applied as a single atomic unit.</span>
<a id="L18"></a><span class="comment">// Patch sets are often preceded by a descriptive header.</span>
<a id="L19"></a>type Set struct {
    <a id="L20"></a>Header string; <span class="comment">// free-form text</span>
    <a id="L21"></a>File   []*File;
<a id="L22"></a>}

<a id="L24"></a><span class="comment">// A File represents a collection of changes to be made to a single file.</span>
<a id="L25"></a>type File struct {
    <a id="L26"></a>Verb             Verb;
    <a id="L27"></a>Src              string; <span class="comment">// source for Verb == Copy, Verb == Rename</span>
    <a id="L28"></a>Dst              string;
    <a id="L29"></a>OldMode, NewMode int; <span class="comment">// 0 indicates not used</span>
    <a id="L30"></a>Diff;                 <span class="comment">// changes to data; == NoDiff if operation does not edit file</span>
<a id="L31"></a>}

<a id="L33"></a><span class="comment">// A Verb is an action performed on a file.</span>
<a id="L34"></a>type Verb string

<a id="L36"></a>const (
    <a id="L37"></a>Add    Verb = &#34;add&#34;;
    <a id="L38"></a>Copy   Verb = &#34;copy&#34;;
    <a id="L39"></a>Delete Verb = &#34;delete&#34;;
    <a id="L40"></a>Edit   Verb = &#34;edit&#34;;
    <a id="L41"></a>Rename Verb = &#34;rename&#34;;
<a id="L42"></a>)

<a id="L44"></a><span class="comment">// A Diff is any object that describes changes to transform</span>
<a id="L45"></a><span class="comment">// an old byte stream to a new one.</span>
<a id="L46"></a>type Diff interface {
    <a id="L47"></a><span class="comment">// Apply applies the changes listed in the diff</span>
    <a id="L48"></a><span class="comment">// to the string s, returning the new version of the string.</span>
    <a id="L49"></a><span class="comment">// Note that the string s need not be a text string.</span>
    <a id="L50"></a>Apply(old []byte) (new []byte, err os.Error);
<a id="L51"></a>}

<a id="L53"></a><span class="comment">// NoDiff is a no-op Diff implementation: it passes the</span>
<a id="L54"></a><span class="comment">// old data through unchanged.</span>
<a id="L55"></a>var NoDiff Diff = noDiffType(0)

<a id="L57"></a>type noDiffType int

<a id="L59"></a>func (noDiffType) Apply(old []byte) ([]byte, os.Error) {
    <a id="L60"></a>return old, nil
<a id="L61"></a>}

<a id="L63"></a><span class="comment">// A SyntaxError represents a syntax error encountered while parsing a patch.</span>
<a id="L64"></a>type SyntaxError string

<a id="L66"></a>func (e SyntaxError) String() string { return string(e) }

<a id="L68"></a>var newline = []byte{&#39;\n&#39;}

<a id="L70"></a><span class="comment">// Parse patches the patch text to create a patch Set.</span>
<a id="L71"></a><span class="comment">// The patch text typically comprises a textual header and a sequence</span>
<a id="L72"></a><span class="comment">// of file patches, as would be generated by CVS, Subversion,</span>
<a id="L73"></a><span class="comment">// Mercurial, or Git.</span>
<a id="L74"></a>func Parse(text []byte) (*Set, os.Error) {
    <a id="L75"></a><span class="comment">// Split text into files.</span>
    <a id="L76"></a><span class="comment">// CVS and Subversion begin new files with</span>
    <a id="L77"></a><span class="comment">//	Index: file name.</span>
    <a id="L78"></a><span class="comment">//	==================</span>
    <a id="L79"></a><span class="comment">//	diff -u blah blah</span>
    <a id="L80"></a><span class="comment">//</span>
    <a id="L81"></a><span class="comment">// Mercurial and Git use</span>
    <a id="L82"></a><span class="comment">//	diff [--git] a/file/path b/file/path.</span>
    <a id="L83"></a><span class="comment">//</span>
    <a id="L84"></a><span class="comment">// First look for Index: lines.  If none, fall back on diff lines.</span>
    <a id="L85"></a>text, files := sections(text, &#34;Index: &#34;);
    <a id="L86"></a>if len(files) == 0 {
        <a id="L87"></a>text, files = sections(text, &#34;diff &#34;)
    <a id="L88"></a>}

    <a id="L90"></a>set := &amp;Set{string(text), make([]*File, len(files))};

    <a id="L92"></a><span class="comment">// Parse file header and then</span>
    <a id="L93"></a><span class="comment">// parse files into patch chunks.</span>
    <a id="L94"></a><span class="comment">// Each chunk begins with @@.</span>
    <a id="L95"></a>for i, raw := range files {
        <a id="L96"></a>p := new(File);
        <a id="L97"></a>set.File[i] = p;

        <a id="L99"></a><span class="comment">// First line of hdr is the Index: that</span>
        <a id="L100"></a><span class="comment">// begins the section.  After that is the file name.</span>
        <a id="L101"></a>s, raw, _ := getLine(raw, 1);
        <a id="L102"></a>if hasPrefix(s, &#34;Index: &#34;) {
            <a id="L103"></a>p.Dst = string(bytes.TrimSpace(s[7:len(s)]));
            <a id="L104"></a>goto HaveName;
        <a id="L105"></a>} else if hasPrefix(s, &#34;diff &#34;) {
            <a id="L106"></a>str := string(bytes.TrimSpace(s));
            <a id="L107"></a>i := strings.LastIndex(str, &#34; b/&#34;);
            <a id="L108"></a>if i &gt;= 0 {
                <a id="L109"></a>p.Dst = str[i+3 : len(str)];
                <a id="L110"></a>goto HaveName;
            <a id="L111"></a>}
        <a id="L112"></a>}
        <a id="L113"></a>return nil, SyntaxError(&#34;unexpected patch header line: &#34; + string(s));
    <a id="L114"></a>HaveName:
        <a id="L115"></a>p.Dst = path.Clean(p.Dst);
        <a id="L116"></a>if strings.HasPrefix(p.Dst, &#34;../&#34;) || strings.HasPrefix(p.Dst, &#34;/&#34;) {
            <a id="L117"></a>return nil, SyntaxError(&#34;invalid path: &#34; + p.Dst)
        <a id="L118"></a>}

        <a id="L120"></a><span class="comment">// Parse header lines giving file information:</span>
        <a id="L121"></a><span class="comment">//	new file mode %o	- file created</span>
        <a id="L122"></a><span class="comment">//	deleted file mode %o	- file deleted</span>
        <a id="L123"></a><span class="comment">//	old file mode %o	- file mode changed</span>
        <a id="L124"></a><span class="comment">//	new file mode %o	- file mode changed</span>
        <a id="L125"></a><span class="comment">//	rename from %s	- file renamed from other file</span>
        <a id="L126"></a><span class="comment">//	rename to %s</span>
        <a id="L127"></a><span class="comment">//	copy from %s		- file copied from other file</span>
        <a id="L128"></a><span class="comment">//	copy to %s</span>
        <a id="L129"></a>p.Verb = Edit;
        <a id="L130"></a>for len(raw) &gt; 0 {
            <a id="L131"></a>oldraw := raw;
            <a id="L132"></a>var l []byte;
            <a id="L133"></a>l, raw, _ = getLine(raw, 1);
            <a id="L134"></a>l = bytes.TrimSpace(l);
            <a id="L135"></a>if m, s, ok := atoi(l, &#34;new file mode &#34;, 8); ok &amp;&amp; len(s) == 0 {
                <a id="L136"></a>p.NewMode = m;
                <a id="L137"></a>p.Verb = Add;
                <a id="L138"></a>continue;
            <a id="L139"></a>}
            <a id="L140"></a>if m, s, ok := atoi(l, &#34;deleted file mode &#34;, 8); ok &amp;&amp; len(s) == 0 {
                <a id="L141"></a>p.OldMode = m;
                <a id="L142"></a>p.Verb = Delete;
                <a id="L143"></a>p.Src = p.Dst;
                <a id="L144"></a>p.Dst = &#34;&#34;;
                <a id="L145"></a>continue;
            <a id="L146"></a>}
            <a id="L147"></a>if m, s, ok := atoi(l, &#34;old file mode &#34;, 8); ok &amp;&amp; len(s) == 0 {
                <a id="L148"></a><span class="comment">// usually implies p.Verb = &#34;rename&#34; or &#34;copy&#34;</span>
                <a id="L149"></a><span class="comment">// but we&#39;ll get that from the rename or copy line.</span>
                <a id="L150"></a>p.OldMode = m;
                <a id="L151"></a>continue;
            <a id="L152"></a>}
            <a id="L153"></a>if m, s, ok := atoi(l, &#34;old mode &#34;, 8); ok &amp;&amp; len(s) == 0 {
                <a id="L154"></a>p.OldMode = m;
                <a id="L155"></a>continue;
            <a id="L156"></a>}
            <a id="L157"></a>if m, s, ok := atoi(l, &#34;new mode &#34;, 8); ok &amp;&amp; len(s) == 0 {
                <a id="L158"></a>p.NewMode = m;
                <a id="L159"></a>continue;
            <a id="L160"></a>}
            <a id="L161"></a>if s, ok := skip(l, &#34;rename from &#34;); ok &amp;&amp; len(s) &gt; 0 {
                <a id="L162"></a>p.Src = string(s);
                <a id="L163"></a>p.Verb = Rename;
                <a id="L164"></a>continue;
            <a id="L165"></a>}
            <a id="L166"></a>if s, ok := skip(l, &#34;rename to &#34;); ok &amp;&amp; len(s) &gt; 0 {
                <a id="L167"></a>p.Verb = Rename;
                <a id="L168"></a>continue;
            <a id="L169"></a>}
            <a id="L170"></a>if s, ok := skip(l, &#34;copy from &#34;); ok &amp;&amp; len(s) &gt; 0 {
                <a id="L171"></a>p.Src = string(s);
                <a id="L172"></a>p.Verb = Copy;
                <a id="L173"></a>continue;
            <a id="L174"></a>}
            <a id="L175"></a>if s, ok := skip(l, &#34;copy to &#34;); ok &amp;&amp; len(s) &gt; 0 {
                <a id="L176"></a>p.Verb = Copy;
                <a id="L177"></a>continue;
            <a id="L178"></a>}
            <a id="L179"></a>if s, ok := skip(l, &#34;Binary file &#34;); ok &amp;&amp; len(s) &gt; 0 {
                <a id="L180"></a><span class="comment">// Hg prints</span>
                <a id="L181"></a><span class="comment">//	Binary file foo has changed</span>
                <a id="L182"></a><span class="comment">// when deleting a binary file.</span>
                <a id="L183"></a>continue
            <a id="L184"></a>}
            <a id="L185"></a>if s, ok := skip(l, &#34;RCS file: &#34;); ok &amp;&amp; len(s) &gt; 0 {
                <a id="L186"></a><span class="comment">// CVS prints</span>
                <a id="L187"></a><span class="comment">//	RCS file: /cvs/plan9/bin/yesterday,v</span>
                <a id="L188"></a><span class="comment">//	retrieving revision 1.1</span>
                <a id="L189"></a><span class="comment">// for each file.</span>
                <a id="L190"></a>continue
            <a id="L191"></a>}
            <a id="L192"></a>if s, ok := skip(l, &#34;retrieving revision &#34;); ok &amp;&amp; len(s) &gt; 0 {
                <a id="L193"></a><span class="comment">// CVS prints</span>
                <a id="L194"></a><span class="comment">//	RCS file: /cvs/plan9/bin/yesterday,v</span>
                <a id="L195"></a><span class="comment">//	retrieving revision 1.1</span>
                <a id="L196"></a><span class="comment">// for each file.</span>
                <a id="L197"></a>continue
            <a id="L198"></a>}
            <a id="L199"></a>if hasPrefix(l, &#34;===&#34;) || hasPrefix(l, &#34;---&#34;) || hasPrefix(l, &#34;+++&#34;) || hasPrefix(l, &#34;diff &#34;) {
                <a id="L200"></a>continue
            <a id="L201"></a>}
            <a id="L202"></a>if hasPrefix(l, &#34;@@ -&#34;) {
                <a id="L203"></a>diff, err := ParseTextDiff(oldraw);
                <a id="L204"></a>if err != nil {
                    <a id="L205"></a>return nil, err
                <a id="L206"></a>}
                <a id="L207"></a>p.Diff = diff;
                <a id="L208"></a>break;
            <a id="L209"></a>}
            <a id="L210"></a>if hasPrefix(l, &#34;index &#34;) || hasPrefix(l, &#34;GIT binary patch&#34;) {
                <a id="L211"></a>diff, err := ParseGitBinary(oldraw);
                <a id="L212"></a>if err != nil {
                    <a id="L213"></a>return nil, err
                <a id="L214"></a>}
                <a id="L215"></a>p.Diff = diff;
                <a id="L216"></a>break;
            <a id="L217"></a>}
            <a id="L218"></a>return nil, SyntaxError(&#34;unexpected patch header line: &#34; + string(l));
        <a id="L219"></a>}
        <a id="L220"></a>if p.Diff == nil {
            <a id="L221"></a>p.Diff = NoDiff
        <a id="L222"></a>}
        <a id="L223"></a>if p.Verb == Edit {
            <a id="L224"></a>p.Src = p.Dst
        <a id="L225"></a>}
    <a id="L226"></a>}

    <a id="L228"></a>return set, nil;
<a id="L229"></a>}

<a id="L231"></a><span class="comment">// getLine returns the first n lines of data and the remainder.</span>
<a id="L232"></a><span class="comment">// If data has no newline, getLine returns data, nil, false</span>
<a id="L233"></a>func getLine(data []byte, n int) (first []byte, rest []byte, ok bool) {
    <a id="L234"></a>rest = data;
    <a id="L235"></a>ok = true;
    <a id="L236"></a>for ; n &gt; 0; n-- {
        <a id="L237"></a>nl := bytes.Index(rest, newline);
        <a id="L238"></a>if nl &lt; 0 {
            <a id="L239"></a>rest = nil;
            <a id="L240"></a>ok = false;
            <a id="L241"></a>break;
        <a id="L242"></a>}
        <a id="L243"></a>rest = rest[nl+1 : len(rest)];
    <a id="L244"></a>}
    <a id="L245"></a>first = data[0 : len(data)-len(rest)];
    <a id="L246"></a>return;
<a id="L247"></a>}

<a id="L249"></a><span class="comment">// sections returns a collection of file sections,</span>
<a id="L250"></a><span class="comment">// each of which begins with a line satisfying prefix.</span>
<a id="L251"></a><span class="comment">// text before the first instance of such a line is</span>
<a id="L252"></a><span class="comment">// returned separately.</span>
<a id="L253"></a>func sections(text []byte, prefix string) ([]byte, [][]byte) {
    <a id="L254"></a>n := 0;
    <a id="L255"></a>for b := text; ; {
        <a id="L256"></a>if hasPrefix(b, prefix) {
            <a id="L257"></a>n++
        <a id="L258"></a>}
        <a id="L259"></a>nl := bytes.Index(b, newline);
        <a id="L260"></a>if nl &lt; 0 {
            <a id="L261"></a>break
        <a id="L262"></a>}
        <a id="L263"></a>b = b[nl+1 : len(b)];
    <a id="L264"></a>}

    <a id="L266"></a>sect := make([][]byte, n+1);
    <a id="L267"></a>n = 0;
    <a id="L268"></a>for b := text; ; {
        <a id="L269"></a>if hasPrefix(b, prefix) {
            <a id="L270"></a>sect[n] = text[0 : len(text)-len(b)];
            <a id="L271"></a>n++;
            <a id="L272"></a>text = b;
        <a id="L273"></a>}
        <a id="L274"></a>nl := bytes.Index(b, newline);
        <a id="L275"></a>if nl &lt; 0 {
            <a id="L276"></a>sect[n] = text;
            <a id="L277"></a>break;
        <a id="L278"></a>}
        <a id="L279"></a>b = b[nl+1 : len(b)];
    <a id="L280"></a>}
    <a id="L281"></a>return sect[0], sect[1:len(sect)];
<a id="L282"></a>}

<a id="L284"></a><span class="comment">// if s begins with the prefix t, skip returns</span>
<a id="L285"></a><span class="comment">// s with that prefix removed and ok == true.</span>
<a id="L286"></a>func skip(s []byte, t string) (ss []byte, ok bool) {
    <a id="L287"></a>if len(s) &lt; len(t) || string(s[0:len(t)]) != t {
        <a id="L288"></a>return nil, false
    <a id="L289"></a>}
    <a id="L290"></a>return s[len(t):len(s)], true;
<a id="L291"></a>}

<a id="L293"></a><span class="comment">// if s begins with the prefix t and then is a sequence</span>
<a id="L294"></a><span class="comment">// of digits in the given base, atoi returns the number</span>
<a id="L295"></a><span class="comment">// represented by the digits and s with the</span>
<a id="L296"></a><span class="comment">// prefix and the digits removed.</span>
<a id="L297"></a>func atoi(s []byte, t string, base int) (n int, ss []byte, ok bool) {
    <a id="L298"></a>if s, ok = skip(s, t); !ok {
        <a id="L299"></a>return
    <a id="L300"></a>}
    <a id="L301"></a>var i int;
    <a id="L302"></a>for i = 0; i &lt; len(s) &amp;&amp; &#39;0&#39; &lt;= s[i] &amp;&amp; s[i] &lt;= byte(&#39;0&#39;+base-1); i++ {
        <a id="L303"></a>n = n*base + int(s[i]-&#39;0&#39;)
    <a id="L304"></a>}
    <a id="L305"></a>if i == 0 {
        <a id="L306"></a>return
    <a id="L307"></a>}
    <a id="L308"></a>return n, s[i:len(s)], true;
<a id="L309"></a>}

<a id="L311"></a><span class="comment">// hasPrefix returns true if s begins with t.</span>
<a id="L312"></a>func hasPrefix(s []byte, t string) bool {
    <a id="L313"></a>_, ok := skip(s, t);
    <a id="L314"></a>return ok;
<a id="L315"></a>}

<a id="L317"></a><span class="comment">// splitLines returns the result of splitting s into lines.</span>
<a id="L318"></a><span class="comment">// The \n on each line is preserved.</span>
<a id="L319"></a>func splitLines(s []byte) [][]byte { return bytes.SplitAfter(s, newline, 0) }
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
