<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/godoc/index.go</title>

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
  <h1 id="generatedHeader">Source file /src/cmd/godoc/index.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This file contains the infrastructure to create an</span>
<a id="L6"></a><span class="comment">// (identifier) index for a set of Go files.</span>
<a id="L7"></a><span class="comment">//</span>
<a id="L8"></a><span class="comment">// Basic indexing algorithm:</span>
<a id="L9"></a><span class="comment">// - traverse all .go files of the file tree specified by root</span>
<a id="L10"></a><span class="comment">// - for each word (identifier) encountered, collect all occurences (spots)</span>
<a id="L11"></a><span class="comment">//   into a list; this produces a list of spots for each word</span>
<a id="L12"></a><span class="comment">// - reduce the lists: from a list of spots to a list of FileRuns,</span>
<a id="L13"></a><span class="comment">//   and from a list of FileRuns into a list of PakRuns</span>
<a id="L14"></a><span class="comment">// - make a HitList from the PakRuns</span>
<a id="L15"></a><span class="comment">//</span>
<a id="L16"></a><span class="comment">// Details:</span>
<a id="L17"></a><span class="comment">// - keep two lists per word: one containing package-level declarations</span>
<a id="L18"></a><span class="comment">//   that have snippets, and one containing all other spots</span>
<a id="L19"></a><span class="comment">// - keep the snippets in a separate table indexed by snippet index</span>
<a id="L20"></a><span class="comment">//   and store the snippet index in place of the line number in a SpotInfo</span>
<a id="L21"></a><span class="comment">//   (the line number for spots with snippets is stored in the snippet)</span>
<a id="L22"></a><span class="comment">// - at the end, create lists of alternative spellings for a given</span>
<a id="L23"></a><span class="comment">//   word</span>

<a id="L25"></a>package main

<a id="L27"></a>import (
    <a id="L28"></a>&#34;container/vector&#34;;
    <a id="L29"></a>&#34;go/ast&#34;;
    <a id="L30"></a>&#34;go/parser&#34;;
    <a id="L31"></a>&#34;go/token&#34;;
    <a id="L32"></a>&#34;go/scanner&#34;;
    <a id="L33"></a>&#34;os&#34;;
    <a id="L34"></a>pathutil &#34;path&#34;;
    <a id="L35"></a>&#34;sort&#34;;
    <a id="L36"></a>&#34;strings&#34;;
<a id="L37"></a>)


<a id="L40"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L41"></a><span class="comment">// RunList</span>

<a id="L43"></a><span class="comment">// A RunList is a vector of entries that can be sorted according to some</span>
<a id="L44"></a><span class="comment">// criteria. A RunList may be compressed by grouping &#34;runs&#34; of entries</span>
<a id="L45"></a><span class="comment">// which are equal (according to the sort critera) into a new RunList of</span>
<a id="L46"></a><span class="comment">// runs. For instance, a RunList containing pairs (x, y) may be compressed</span>
<a id="L47"></a><span class="comment">// into a RunList containing pair runs (x, {y}) where each run consists of</span>
<a id="L48"></a><span class="comment">// a list of y&#39;s with the same x.</span>
<a id="L49"></a>type RunList struct {
    <a id="L50"></a>vector.Vector;
    <a id="L51"></a>less func(x, y interface{}) bool;
<a id="L52"></a>}

<a id="L54"></a>func (h *RunList) Less(i, j int) bool { return h.less(h.At(i), h.At(j)) }


<a id="L57"></a>func (h *RunList) sort(less func(x, y interface{}) bool) {
    <a id="L58"></a>h.less = less;
    <a id="L59"></a>sort.Sort(h);
<a id="L60"></a>}


<a id="L63"></a><span class="comment">// Compress entries which are the same according to a sort criteria</span>
<a id="L64"></a><span class="comment">// (specified by less) into &#34;runs&#34;.</span>
<a id="L65"></a>func (h *RunList) reduce(less func(x, y interface{}) bool, newRun func(h *RunList, i, j int) interface{}) *RunList {
    <a id="L66"></a><span class="comment">// create runs of entries with equal values</span>
    <a id="L67"></a>h.sort(less);

    <a id="L69"></a><span class="comment">// for each run, make a new run object and collect them in a new RunList</span>
    <a id="L70"></a>var hh RunList;
    <a id="L71"></a>i := 0;
    <a id="L72"></a>for j := 0; j &lt; h.Len(); j++ {
        <a id="L73"></a>if less(h.At(i), h.At(j)) {
            <a id="L74"></a>hh.Push(newRun(h, i, j));
            <a id="L75"></a>i = j; <span class="comment">// start a new run</span>
        <a id="L76"></a>}
    <a id="L77"></a>}
    <a id="L78"></a><span class="comment">// add final run, if any</span>
    <a id="L79"></a>if i &lt; h.Len() {
        <a id="L80"></a>hh.Push(newRun(h, i, h.Len()))
    <a id="L81"></a>}

    <a id="L83"></a>return &amp;hh;
<a id="L84"></a>}


<a id="L87"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L88"></a><span class="comment">// SpotInfo</span>

<a id="L90"></a><span class="comment">// A SpotInfo value describes a particular identifier spot in a given file;</span>
<a id="L91"></a><span class="comment">// It encodes three values: the SpotKind (declaration or use), a line or</span>
<a id="L92"></a><span class="comment">// snippet index &#34;lori&#34;, and whether it&#39;s a line or index.</span>
<a id="L93"></a><span class="comment">//</span>
<a id="L94"></a><span class="comment">// The following encoding is used:</span>
<a id="L95"></a><span class="comment">//</span>
<a id="L96"></a><span class="comment">//   bits    32   4    1       0</span>
<a id="L97"></a><span class="comment">//   value    [lori|kind|isIndex]</span>
<a id="L98"></a><span class="comment">//</span>
<a id="L99"></a>type SpotInfo uint32

<a id="L101"></a><span class="comment">// SpotKind describes whether an identifier is declared (and what kind of</span>
<a id="L102"></a><span class="comment">// declaration) or used.</span>
<a id="L103"></a>type SpotKind uint32

<a id="L105"></a>const (
    <a id="L106"></a>PackageClause SpotKind = iota;
    <a id="L107"></a>ImportDecl;
    <a id="L108"></a>ConstDecl;
    <a id="L109"></a>TypeDecl;
    <a id="L110"></a>VarDecl;
    <a id="L111"></a>FuncDecl;
    <a id="L112"></a>MethodDecl;
    <a id="L113"></a>Use;
    <a id="L114"></a>nKinds;
<a id="L115"></a>)


<a id="L118"></a>func init() {
    <a id="L119"></a><span class="comment">// sanity check: if nKinds is too large, the SpotInfo</span>
    <a id="L120"></a><span class="comment">// accessor functions may need to be updated</span>
    <a id="L121"></a>if nKinds &gt; 8 {
        <a id="L122"></a>panic()
    <a id="L123"></a>}
<a id="L124"></a>}


<a id="L127"></a><span class="comment">// makeSpotInfo makes a SpotInfo.</span>
<a id="L128"></a>func makeSpotInfo(kind SpotKind, lori int, isIndex bool) SpotInfo {
    <a id="L129"></a><span class="comment">// encode lori: bits [4..32)</span>
    <a id="L130"></a>x := SpotInfo(lori) &lt;&lt; 4;
    <a id="L131"></a>if int(x&gt;&gt;4) != lori {
        <a id="L132"></a><span class="comment">// lori value doesn&#39;t fit - since snippet indices are</span>
        <a id="L133"></a><span class="comment">// most certainly always smaller then 1&lt;&lt;28, this can</span>
        <a id="L134"></a><span class="comment">// only happen for line numbers; give it no line number (= 0)</span>
        <a id="L135"></a>x = 0
    <a id="L136"></a>}
    <a id="L137"></a><span class="comment">// encode kind: bits [1..4)</span>
    <a id="L138"></a>x |= SpotInfo(kind) &lt;&lt; 1;
    <a id="L139"></a><span class="comment">// encode isIndex: bit 0</span>
    <a id="L140"></a>if isIndex {
        <a id="L141"></a>x |= 1
    <a id="L142"></a>}
    <a id="L143"></a>return x;
<a id="L144"></a>}


<a id="L147"></a>func (x SpotInfo) Kind() SpotKind { return SpotKind(x &gt;&gt; 1 &amp; 7) }
<a id="L148"></a>func (x SpotInfo) Lori() int      { return int(x &gt;&gt; 4) }
<a id="L149"></a>func (x SpotInfo) IsIndex() bool  { return x&amp;1 != 0 }


<a id="L152"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L153"></a><span class="comment">// KindRun</span>

<a id="L155"></a><span class="comment">// Debugging support. Disable to see multiple entries per line.</span>
<a id="L156"></a>const removeDuplicates = true

<a id="L158"></a><span class="comment">// A KindRun is a run of SpotInfos of the same kind in a given file.</span>
<a id="L159"></a>type KindRun struct {
    <a id="L160"></a>Kind  SpotKind;
    <a id="L161"></a>Infos []SpotInfo;
<a id="L162"></a>}


<a id="L165"></a><span class="comment">// KindRuns are sorted by line number or index. Since the isIndex bit</span>
<a id="L166"></a><span class="comment">// is always the same for all infos in one list we can compare lori&#39;s.</span>
<a id="L167"></a>func (f *KindRun) Len() int           { return len(f.Infos) }
<a id="L168"></a>func (f *KindRun) Less(i, j int) bool { return f.Infos[i].Lori() &lt; f.Infos[j].Lori() }
<a id="L169"></a>func (f *KindRun) Swap(i, j int)      { f.Infos[i], f.Infos[j] = f.Infos[j], f.Infos[i] }


<a id="L172"></a><span class="comment">// FileRun contents are sorted by Kind for the reduction into KindRuns.</span>
<a id="L173"></a>func lessKind(x, y interface{}) bool { return x.(SpotInfo).Kind() &lt; y.(SpotInfo).Kind() }


<a id="L176"></a><span class="comment">// newKindRun allocates a new KindRun from the SpotInfo run [i, j) in h.</span>
<a id="L177"></a>func newKindRun(h *RunList, i, j int) interface{} {
    <a id="L178"></a>kind := h.At(i).(SpotInfo).Kind();
    <a id="L179"></a>infos := make([]SpotInfo, j-i);
    <a id="L180"></a>k := 0;
    <a id="L181"></a>for ; i &lt; j; i++ {
        <a id="L182"></a>infos[k] = h.At(i).(SpotInfo);
        <a id="L183"></a>k++;
    <a id="L184"></a>}
    <a id="L185"></a>run := &amp;KindRun{kind, infos};

    <a id="L187"></a><span class="comment">// Spots were sorted by file and kind to create this run.</span>
    <a id="L188"></a><span class="comment">// Within this run, sort them by line number or index.</span>
    <a id="L189"></a>sort.Sort(run);

    <a id="L191"></a>if removeDuplicates {
        <a id="L192"></a><span class="comment">// Since both the lori and kind field must be</span>
        <a id="L193"></a><span class="comment">// same for duplicates, and since the isIndex</span>
        <a id="L194"></a><span class="comment">// bit is always the same for all infos in one</span>
        <a id="L195"></a><span class="comment">// list we can simply compare the entire info.</span>
        <a id="L196"></a>k := 0;
        <a id="L197"></a>var prev SpotInfo;
        <a id="L198"></a>for i, x := range infos {
            <a id="L199"></a>if x != prev || i == 0 {
                <a id="L200"></a>infos[k] = x;
                <a id="L201"></a>k++;
                <a id="L202"></a>prev = x;
            <a id="L203"></a>}
        <a id="L204"></a>}
        <a id="L205"></a>run.Infos = infos[0:k];
    <a id="L206"></a>}

    <a id="L208"></a>return run;
<a id="L209"></a>}


<a id="L212"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L213"></a><span class="comment">// FileRun</span>

<a id="L215"></a><span class="comment">// A Pak describes a Go package.</span>
<a id="L216"></a>type Pak struct {
    <a id="L217"></a>Path string; <span class="comment">// path of directory containing the package</span>
    <a id="L218"></a>Name string; <span class="comment">// package name as declared by package clause</span>
<a id="L219"></a>}

<a id="L221"></a><span class="comment">// Paks are sorted by name (primary key) and by import path (secondary key).</span>
<a id="L222"></a>func (p *Pak) less(q *Pak) bool {
    <a id="L223"></a>return p.Name &lt; q.Name || p.Name == q.Name &amp;&amp; p.Path &lt; q.Path
<a id="L224"></a>}


<a id="L227"></a><span class="comment">// A File describes a Go file.</span>
<a id="L228"></a>type File struct {
    <a id="L229"></a>Path string; <span class="comment">// complete file name</span>
    <a id="L230"></a>Pak  Pak;    <span class="comment">// the package to which the file belongs</span>
<a id="L231"></a>}


<a id="L234"></a><span class="comment">// A Spot describes a single occurence of a word.</span>
<a id="L235"></a>type Spot struct {
    <a id="L236"></a>File *File;
    <a id="L237"></a>Info SpotInfo;
<a id="L238"></a>}


<a id="L241"></a><span class="comment">// A FileRun is a list of KindRuns belonging to the same file.</span>
<a id="L242"></a>type FileRun struct {
    <a id="L243"></a>File   *File;
    <a id="L244"></a>Groups []*KindRun;
<a id="L245"></a>}


<a id="L248"></a><span class="comment">// Spots are sorted by path for the reduction into FileRuns.</span>
<a id="L249"></a>func lessSpot(x, y interface{}) bool { return x.(Spot).File.Path &lt; y.(Spot).File.Path }


<a id="L252"></a><span class="comment">// newFileRun allocates a new FileRun from the Spot run [i, j) in h.</span>
<a id="L253"></a>func newFileRun(h0 *RunList, i, j int) interface{} {
    <a id="L254"></a>file := h0.At(i).(Spot).File;

    <a id="L256"></a><span class="comment">// reduce the list of Spots into a list of KindRuns</span>
    <a id="L257"></a>var h1 RunList;
    <a id="L258"></a>h1.Vector.Init(j - i);
    <a id="L259"></a>k := 0;
    <a id="L260"></a>for ; i &lt; j; i++ {
        <a id="L261"></a>h1.Set(k, h0.At(i).(Spot).Info);
        <a id="L262"></a>k++;
    <a id="L263"></a>}
    <a id="L264"></a>h2 := h1.reduce(lessKind, newKindRun);

    <a id="L266"></a><span class="comment">// create the FileRun</span>
    <a id="L267"></a>groups := make([]*KindRun, h2.Len());
    <a id="L268"></a>for i := 0; i &lt; h2.Len(); i++ {
        <a id="L269"></a>groups[i] = h2.At(i).(*KindRun)
    <a id="L270"></a>}
    <a id="L271"></a>return &amp;FileRun{file, groups};
<a id="L272"></a>}


<a id="L275"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L276"></a><span class="comment">// PakRun</span>

<a id="L278"></a><span class="comment">// A PakRun describes a run of *FileRuns of a package.</span>
<a id="L279"></a>type PakRun struct {
    <a id="L280"></a>Pak   Pak;
    <a id="L281"></a>Files []*FileRun;
<a id="L282"></a>}

<a id="L284"></a><span class="comment">// Sorting support for files within a PakRun.</span>
<a id="L285"></a>func (p *PakRun) Len() int           { return len(p.Files) }
<a id="L286"></a>func (p *PakRun) Less(i, j int) bool { return p.Files[i].File.Path &lt; p.Files[j].File.Path }
<a id="L287"></a>func (p *PakRun) Swap(i, j int)      { p.Files[i], p.Files[j] = p.Files[j], p.Files[i] }


<a id="L290"></a><span class="comment">// FileRuns are sorted by package for the reduction into PakRuns.</span>
<a id="L291"></a>func lessFileRun(x, y interface{}) bool {
    <a id="L292"></a>return x.(*FileRun).File.Pak.less(&amp;y.(*FileRun).File.Pak)
<a id="L293"></a>}


<a id="L296"></a><span class="comment">// newPakRun allocates a new PakRun from the *FileRun run [i, j) in h.</span>
<a id="L297"></a>func newPakRun(h *RunList, i, j int) interface{} {
    <a id="L298"></a>pak := h.At(i).(*FileRun).File.Pak;
    <a id="L299"></a>files := make([]*FileRun, j-i);
    <a id="L300"></a>k := 0;
    <a id="L301"></a>for ; i &lt; j; i++ {
        <a id="L302"></a>files[k] = h.At(i).(*FileRun);
        <a id="L303"></a>k++;
    <a id="L304"></a>}
    <a id="L305"></a>run := &amp;PakRun{pak, files};
    <a id="L306"></a>sort.Sort(run); <span class="comment">// files were sorted by package; sort them by file now</span>
    <a id="L307"></a>return run;
<a id="L308"></a>}


<a id="L311"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L312"></a><span class="comment">// HitList</span>

<a id="L314"></a><span class="comment">// A HitList describes a list of PakRuns.</span>
<a id="L315"></a>type HitList []*PakRun


<a id="L318"></a><span class="comment">// PakRuns are sorted by package.</span>
<a id="L319"></a>func lessPakRun(x, y interface{}) bool { return x.(*PakRun).Pak.less(&amp;y.(*PakRun).Pak) }


<a id="L322"></a>func reduce(h0 *RunList) HitList {
    <a id="L323"></a><span class="comment">// reduce a list of Spots into a list of FileRuns</span>
    <a id="L324"></a>h1 := h0.reduce(lessSpot, newFileRun);
    <a id="L325"></a><span class="comment">// reduce a list of FileRuns into a list of PakRuns</span>
    <a id="L326"></a>h2 := h1.reduce(lessFileRun, newPakRun);
    <a id="L327"></a><span class="comment">// sort the list of PakRuns by package</span>
    <a id="L328"></a>h2.sort(lessPakRun);
    <a id="L329"></a><span class="comment">// create a HitList</span>
    <a id="L330"></a>h := make(HitList, h2.Len());
    <a id="L331"></a>for i := 0; i &lt; h2.Len(); i++ {
        <a id="L332"></a>h[i] = h2.At(i).(*PakRun)
    <a id="L333"></a>}
    <a id="L334"></a>return h;
<a id="L335"></a>}


<a id="L338"></a>func (h HitList) filter(pakname string) HitList {
    <a id="L339"></a><span class="comment">// determine number of matching packages (most of the time just one)</span>
    <a id="L340"></a>n := 0;
    <a id="L341"></a>for _, p := range h {
        <a id="L342"></a>if p.Pak.Name == pakname {
            <a id="L343"></a>n++
        <a id="L344"></a>}
    <a id="L345"></a>}
    <a id="L346"></a><span class="comment">// create filtered HitList</span>
    <a id="L347"></a>hh := make(HitList, n);
    <a id="L348"></a>i := 0;
    <a id="L349"></a>for _, p := range h {
        <a id="L350"></a>if p.Pak.Name == pakname {
            <a id="L351"></a>hh[i] = p;
            <a id="L352"></a>i++;
        <a id="L353"></a>}
    <a id="L354"></a>}
    <a id="L355"></a>return hh;
<a id="L356"></a>}


<a id="L359"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L360"></a><span class="comment">// AltWords</span>

<a id="L362"></a>type wordPair struct {
    <a id="L363"></a>canon string; <span class="comment">// canonical word spelling (all lowercase)</span>
    <a id="L364"></a>alt   string; <span class="comment">// alternative spelling</span>
<a id="L365"></a>}


<a id="L368"></a><span class="comment">// An AltWords describes a list of alternative spellings for a</span>
<a id="L369"></a><span class="comment">// canonical (all lowercase) spelling of a word.</span>
<a id="L370"></a>type AltWords struct {
    <a id="L371"></a>Canon string;   <span class="comment">// canonical word spelling (all lowercase)</span>
    <a id="L372"></a>Alts  []string; <span class="comment">// alternative spelling for the same word</span>
<a id="L373"></a>}


<a id="L376"></a><span class="comment">// wordPairs are sorted by their canonical spelling.</span>
<a id="L377"></a>func lessWordPair(x, y interface{}) bool { return x.(*wordPair).canon &lt; y.(*wordPair).canon }


<a id="L380"></a><span class="comment">// newAltWords allocates a new AltWords from the *wordPair run [i, j) in h.</span>
<a id="L381"></a>func newAltWords(h *RunList, i, j int) interface{} {
    <a id="L382"></a>canon := h.At(i).(*wordPair).canon;
    <a id="L383"></a>alts := make([]string, j-i);
    <a id="L384"></a>k := 0;
    <a id="L385"></a>for ; i &lt; j; i++ {
        <a id="L386"></a>alts[k] = h.At(i).(*wordPair).alt;
        <a id="L387"></a>k++;
    <a id="L388"></a>}
    <a id="L389"></a>return &amp;AltWords{canon, alts};
<a id="L390"></a>}


<a id="L393"></a>func (a *AltWords) filter(s string) *AltWords {
    <a id="L394"></a>if len(a.Alts) == 1 &amp;&amp; a.Alts[0] == s {
        <a id="L395"></a><span class="comment">// there are no different alternatives</span>
        <a id="L396"></a>return nil
    <a id="L397"></a>}

    <a id="L399"></a><span class="comment">// make a new AltWords with the current spelling removed</span>
    <a id="L400"></a>alts := make([]string, len(a.Alts));
    <a id="L401"></a>i := 0;
    <a id="L402"></a>for _, w := range a.Alts {
        <a id="L403"></a>if w != s {
            <a id="L404"></a>alts[i] = w;
            <a id="L405"></a>i++;
        <a id="L406"></a>}
    <a id="L407"></a>}
    <a id="L408"></a>return &amp;AltWords{a.Canon, alts[0:i]};
<a id="L409"></a>}


<a id="L412"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L413"></a><span class="comment">// Indexer</span>

<a id="L415"></a><span class="comment">// Adjust these flags as seems best.</span>
<a id="L416"></a>const excludeMainPackages = false
<a id="L417"></a>const excludeTestFiles = false


<a id="L420"></a>type IndexResult struct {
    <a id="L421"></a>Decls  RunList; <span class="comment">// package-level declarations (with snippets)</span>
    <a id="L422"></a>Others RunList; <span class="comment">// all other occurences</span>
<a id="L423"></a>}


<a id="L426"></a><span class="comment">// An Indexer maintains the data structures and provides the machinery</span>
<a id="L427"></a><span class="comment">// for indexing .go files under a file tree. It implements the path.Visitor</span>
<a id="L428"></a><span class="comment">// interface for walking file trees, and the ast.Visitor interface for</span>
<a id="L429"></a><span class="comment">// walking Go ASTs.</span>
<a id="L430"></a>type Indexer struct {
    <a id="L431"></a>words    map[string]*IndexResult; <span class="comment">// RunLists of Spots</span>
    <a id="L432"></a>snippets vector.Vector;           <span class="comment">// vector of *Snippets, indexed by snippet indices</span>
    <a id="L433"></a>file     *File;                   <span class="comment">// current file</span>
    <a id="L434"></a>decl     ast.Decl;                <span class="comment">// current decl</span>
    <a id="L435"></a>nspots   int;                     <span class="comment">// number of spots encountered</span>
<a id="L436"></a>}


<a id="L439"></a>func (x *Indexer) addSnippet(s *Snippet) int {
    <a id="L440"></a>index := x.snippets.Len();
    <a id="L441"></a>x.snippets.Push(s);
    <a id="L442"></a>return index;
<a id="L443"></a>}


<a id="L446"></a>func (x *Indexer) visitComment(c *ast.CommentGroup) {
    <a id="L447"></a>if c != nil {
        <a id="L448"></a>ast.Walk(x, c)
    <a id="L449"></a>}
<a id="L450"></a>}


<a id="L453"></a>func (x *Indexer) visitIdent(kind SpotKind, id *ast.Ident) {
    <a id="L454"></a>if id != nil {
        <a id="L455"></a>lists, found := x.words[id.Value];
        <a id="L456"></a>if !found {
            <a id="L457"></a>lists = new(IndexResult);
            <a id="L458"></a>x.words[id.Value] = lists;
        <a id="L459"></a>}

        <a id="L461"></a>if kind == Use || x.decl == nil {
            <a id="L462"></a><span class="comment">// not a declaration or no snippet required</span>
            <a id="L463"></a>info := makeSpotInfo(kind, id.Pos().Line, false);
            <a id="L464"></a>lists.Others.Push(Spot{x.file, info});
        <a id="L465"></a>} else {
            <a id="L466"></a><span class="comment">// a declaration with snippet</span>
            <a id="L467"></a>index := x.addSnippet(NewSnippet(x.decl, id));
            <a id="L468"></a>info := makeSpotInfo(kind, index, true);
            <a id="L469"></a>lists.Decls.Push(Spot{x.file, info});
        <a id="L470"></a>}

        <a id="L472"></a>x.nspots++;
    <a id="L473"></a>}
<a id="L474"></a>}


<a id="L477"></a>func (x *Indexer) visitSpec(spec ast.Spec, isVarDecl bool) {
    <a id="L478"></a>switch n := spec.(type) {
    <a id="L479"></a>case *ast.ImportSpec:
        <a id="L480"></a>x.visitComment(n.Doc);
        <a id="L481"></a>x.visitIdent(ImportDecl, n.Name);
        <a id="L482"></a>for _, s := range n.Path {
            <a id="L483"></a>ast.Walk(x, s)
        <a id="L484"></a>}
        <a id="L485"></a>x.visitComment(n.Comment);

    <a id="L487"></a>case *ast.ValueSpec:
        <a id="L488"></a>x.visitComment(n.Doc);
        <a id="L489"></a>kind := ConstDecl;
        <a id="L490"></a>if isVarDecl {
            <a id="L491"></a>kind = VarDecl
        <a id="L492"></a>}
        <a id="L493"></a>for _, n := range n.Names {
            <a id="L494"></a>x.visitIdent(kind, n)
        <a id="L495"></a>}
        <a id="L496"></a>ast.Walk(x, n.Type);
        <a id="L497"></a>for _, v := range n.Values {
            <a id="L498"></a>ast.Walk(x, v)
        <a id="L499"></a>}
        <a id="L500"></a>x.visitComment(n.Comment);

    <a id="L502"></a>case *ast.TypeSpec:
        <a id="L503"></a>x.visitComment(n.Doc);
        <a id="L504"></a>x.visitIdent(TypeDecl, n.Name);
        <a id="L505"></a>ast.Walk(x, n.Type);
        <a id="L506"></a>x.visitComment(n.Comment);
    <a id="L507"></a>}
<a id="L508"></a>}


<a id="L511"></a>func (x *Indexer) Visit(node interface{}) bool {
    <a id="L512"></a><span class="comment">// TODO(gri): methods in interface types are categorized as VarDecl</span>
    <a id="L513"></a>switch n := node.(type) {
    <a id="L514"></a>case *ast.Ident:
        <a id="L515"></a>x.visitIdent(Use, n)

    <a id="L517"></a>case *ast.Field:
        <a id="L518"></a>x.decl = nil; <span class="comment">// no snippets for fields</span>
        <a id="L519"></a>x.visitComment(n.Doc);
        <a id="L520"></a>for _, m := range n.Names {
            <a id="L521"></a>x.visitIdent(VarDecl, m)
        <a id="L522"></a>}
        <a id="L523"></a>ast.Walk(x, n.Type);
        <a id="L524"></a>for _, s := range n.Tag {
            <a id="L525"></a>ast.Walk(x, s)
        <a id="L526"></a>}
        <a id="L527"></a>x.visitComment(n.Comment);

    <a id="L529"></a>case *ast.DeclStmt:
        <a id="L530"></a>if decl, ok := n.Decl.(*ast.GenDecl); ok {
            <a id="L531"></a><span class="comment">// local declarations can only be *ast.GenDecls</span>
            <a id="L532"></a>x.decl = nil; <span class="comment">// no snippets for local declarations</span>
            <a id="L533"></a>x.visitComment(decl.Doc);
            <a id="L534"></a>for _, s := range decl.Specs {
                <a id="L535"></a>x.visitSpec(s, decl.Tok == token.VAR)
            <a id="L536"></a>}
        <a id="L537"></a>} else {
            <a id="L538"></a><span class="comment">// handle error case gracefully</span>
            <a id="L539"></a>ast.Walk(x, n.Decl)
        <a id="L540"></a>}

    <a id="L542"></a>case *ast.GenDecl:
        <a id="L543"></a>x.decl = n;
        <a id="L544"></a>x.visitComment(n.Doc);
        <a id="L545"></a>for _, s := range n.Specs {
            <a id="L546"></a>x.visitSpec(s, n.Tok == token.VAR)
        <a id="L547"></a>}

    <a id="L549"></a>case *ast.FuncDecl:
        <a id="L550"></a>x.visitComment(n.Doc);
        <a id="L551"></a>kind := FuncDecl;
        <a id="L552"></a>if n.Recv != nil {
            <a id="L553"></a>kind = MethodDecl;
            <a id="L554"></a>ast.Walk(x, n.Recv);
        <a id="L555"></a>}
        <a id="L556"></a>x.decl = n;
        <a id="L557"></a>x.visitIdent(kind, n.Name);
        <a id="L558"></a>ast.Walk(x, n.Type);
        <a id="L559"></a>if n.Body != nil {
            <a id="L560"></a>ast.Walk(x, n.Body)
        <a id="L561"></a>}

    <a id="L563"></a>case *ast.File:
        <a id="L564"></a>x.visitComment(n.Doc);
        <a id="L565"></a>x.decl = nil;
        <a id="L566"></a>x.visitIdent(PackageClause, n.Name);
        <a id="L567"></a>for _, d := range n.Decls {
            <a id="L568"></a>ast.Walk(x, d)
        <a id="L569"></a>}
        <a id="L570"></a><span class="comment">// don&#39;t visit package level comments for now</span>
        <a id="L571"></a><span class="comment">// to avoid duplicate visiting from individual</span>
        <a id="L572"></a><span class="comment">// nodes</span>

    <a id="L574"></a>default:
        <a id="L575"></a>return true
    <a id="L576"></a>}

    <a id="L578"></a>return false;
<a id="L579"></a>}


<a id="L582"></a>func (x *Indexer) VisitDir(path string, d *os.Dir) bool {
    <a id="L583"></a>return true
<a id="L584"></a>}


<a id="L587"></a>func (x *Indexer) VisitFile(path string, d *os.Dir) {
    <a id="L588"></a>if !isGoFile(d) {
        <a id="L589"></a>return
    <a id="L590"></a>}

    <a id="L592"></a>if excludeTestFiles &amp;&amp; (!isPkgFile(d) || strings.HasPrefix(path, &#34;test/&#34;)) {
        <a id="L593"></a>return
    <a id="L594"></a>}

    <a id="L596"></a>if excludeMainPackages &amp;&amp; pkgName(path) == &#34;main&#34; {
        <a id="L597"></a>return
    <a id="L598"></a>}

    <a id="L600"></a>file, err := parser.ParseFile(path, nil, parser.ParseComments);
    <a id="L601"></a>if err != nil {
        <a id="L602"></a>return <span class="comment">// ignore files with (parse) errors</span>
    <a id="L603"></a>}

    <a id="L605"></a>dir, _ := pathutil.Split(path);
    <a id="L606"></a>pak := Pak{dir, file.Name.Value};
    <a id="L607"></a>x.file = &amp;File{path, pak};
    <a id="L608"></a>ast.Walk(x, file);
<a id="L609"></a>}


<a id="L612"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L613"></a><span class="comment">// Index</span>

<a id="L615"></a>type LookupResult struct {
    <a id="L616"></a>Decls  HitList; <span class="comment">// package-level declarations (with snippets)</span>
    <a id="L617"></a>Others HitList; <span class="comment">// all other occurences</span>
<a id="L618"></a>}


<a id="L621"></a>type Index struct {
    <a id="L622"></a>words    map[string]*LookupResult; <span class="comment">// maps words to hit lists</span>
    <a id="L623"></a>alts     map[string]*AltWords;     <span class="comment">// maps canonical(words) to lists of alternative spellings</span>
    <a id="L624"></a>snippets []*Snippet;               <span class="comment">// all snippets, indexed by snippet index</span>
    <a id="L625"></a>nspots   int;                      <span class="comment">// number of spots indexed (a measure of the index size)</span>
<a id="L626"></a>}


<a id="L629"></a>func canonical(w string) string { return strings.ToLower(w) }


<a id="L632"></a><span class="comment">// NewIndex creates a new index for the file tree rooted at root.</span>
<a id="L633"></a>func NewIndex(root string) *Index {
    <a id="L634"></a>var x Indexer;

    <a id="L636"></a><span class="comment">// initialize Indexer</span>
    <a id="L637"></a>x.words = make(map[string]*IndexResult);

    <a id="L639"></a><span class="comment">// collect all Spots</span>
    <a id="L640"></a>pathutil.Walk(root, &amp;x, nil);

    <a id="L642"></a><span class="comment">// for each word, reduce the RunLists into a LookupResult;</span>
    <a id="L643"></a><span class="comment">// also collect the word with its canonical spelling in a</span>
    <a id="L644"></a><span class="comment">// word list for later computation of alternative spellings</span>
    <a id="L645"></a>words := make(map[string]*LookupResult);
    <a id="L646"></a>var wlist RunList;
    <a id="L647"></a>for w, h := range x.words {
        <a id="L648"></a>decls := reduce(&amp;h.Decls);
        <a id="L649"></a>others := reduce(&amp;h.Others);
        <a id="L650"></a>words[w] = &amp;LookupResult{
            <a id="L651"></a>Decls: decls,
            <a id="L652"></a>Others: others,
        <a id="L653"></a>};
        <a id="L654"></a>wlist.Push(&amp;wordPair{canonical(w), w});
    <a id="L655"></a>}

    <a id="L657"></a><span class="comment">// reduce the word list {canonical(w), w} into</span>
    <a id="L658"></a><span class="comment">// a list of AltWords runs {canonical(w), {w}}</span>
    <a id="L659"></a>alist := wlist.reduce(lessWordPair, newAltWords);

    <a id="L661"></a><span class="comment">// convert alist into a map of alternative spellings</span>
    <a id="L662"></a>alts := make(map[string]*AltWords);
    <a id="L663"></a>for i := 0; i &lt; alist.Len(); i++ {
        <a id="L664"></a>a := alist.At(i).(*AltWords);
        <a id="L665"></a>alts[a.Canon] = a;
    <a id="L666"></a>}

    <a id="L668"></a><span class="comment">// convert snippet vector into a list</span>
    <a id="L669"></a>snippets := make([]*Snippet, x.snippets.Len());
    <a id="L670"></a>for i := 0; i &lt; x.snippets.Len(); i++ {
        <a id="L671"></a>snippets[i] = x.snippets.At(i).(*Snippet)
    <a id="L672"></a>}

    <a id="L674"></a>return &amp;Index{words, alts, snippets, x.nspots};
<a id="L675"></a>}


<a id="L678"></a><span class="comment">// Size returns the number of different words and</span>
<a id="L679"></a><span class="comment">// spots indexed as a measure for the index size.</span>
<a id="L680"></a>func (x *Index) Size() (nwords int, nspots int) {
    <a id="L681"></a>return len(x.words), x.nspots
<a id="L682"></a>}


<a id="L685"></a>func (x *Index) LookupWord(w string) (match *LookupResult, alt *AltWords) {
    <a id="L686"></a>match, _ = x.words[w];
    <a id="L687"></a>alt, _ = x.alts[canonical(w)];
    <a id="L688"></a><span class="comment">// remove current spelling from alternatives</span>
    <a id="L689"></a><span class="comment">// (if there is no match, the alternatives do</span>
    <a id="L690"></a><span class="comment">// not contain the current spelling)</span>
    <a id="L691"></a>if match != nil &amp;&amp; alt != nil {
        <a id="L692"></a>alt = alt.filter(w)
    <a id="L693"></a>}
    <a id="L694"></a>return;
<a id="L695"></a>}


<a id="L698"></a>func isIdentifier(s string) bool {
    <a id="L699"></a>var S scanner.Scanner;
    <a id="L700"></a>S.Init(&#34;&#34;, strings.Bytes(s), nil, 0);
    <a id="L701"></a>if _, tok, _ := S.Scan(); tok == token.IDENT {
        <a id="L702"></a>_, tok, _ := S.Scan();
        <a id="L703"></a>return tok == token.EOF;
    <a id="L704"></a>}
    <a id="L705"></a>return false;
<a id="L706"></a>}


<a id="L709"></a><span class="comment">// For a given query, which is either a single identifier or a qualified</span>
<a id="L710"></a><span class="comment">// identifier, Lookup returns a LookupResult, and a list of alternative</span>
<a id="L711"></a><span class="comment">// spellings, if any. If the query syntax is wrong, illegal is set.</span>
<a id="L712"></a>func (x *Index) Lookup(query string) (match *LookupResult, alt *AltWords, illegal bool) {
    <a id="L713"></a>ss := strings.Split(query, &#34;.&#34;, 0);

    <a id="L715"></a><span class="comment">// check query syntax</span>
    <a id="L716"></a>for _, s := range ss {
        <a id="L717"></a>if !isIdentifier(s) {
            <a id="L718"></a>illegal = true;
            <a id="L719"></a>return;
        <a id="L720"></a>}
    <a id="L721"></a>}

    <a id="L723"></a>switch len(ss) {
    <a id="L724"></a>case 1:
        <a id="L725"></a>match, alt = x.LookupWord(ss[0])

    <a id="L727"></a>case 2:
        <a id="L728"></a>pakname := ss[0];
        <a id="L729"></a>match, alt = x.LookupWord(ss[1]);
        <a id="L730"></a>if match != nil {
            <a id="L731"></a><span class="comment">// found a match - filter by package name</span>
            <a id="L732"></a>decls := match.Decls.filter(pakname);
            <a id="L733"></a>others := match.Others.filter(pakname);
            <a id="L734"></a>match = &amp;LookupResult{decls, others};
        <a id="L735"></a>}

    <a id="L737"></a>default:
        <a id="L738"></a>illegal = true
    <a id="L739"></a>}

    <a id="L741"></a>return;
<a id="L742"></a>}


<a id="L745"></a>func (x *Index) Snippet(i int) *Snippet {
    <a id="L746"></a><span class="comment">// handle illegal snippet indices gracefully</span>
    <a id="L747"></a>if 0 &lt;= i &amp;&amp; i &lt; len(x.snippets) {
        <a id="L748"></a>return x.snippets[i]
    <a id="L749"></a>}
    <a id="L750"></a>return nil;
<a id="L751"></a>}
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
