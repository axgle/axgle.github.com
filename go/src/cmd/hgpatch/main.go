<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/hgpatch/main.go</title>

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
  <h1 id="generatedHeader">Source file /src/cmd/hgpatch/main.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package main

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;container/vector&#34;;
    <a id="L10"></a>&#34;exec&#34;;
    <a id="L11"></a>&#34;flag&#34;;
    <a id="L12"></a>&#34;fmt&#34;;
    <a id="L13"></a>&#34;io&#34;;
    <a id="L14"></a>&#34;os&#34;;
    <a id="L15"></a>&#34;patch&#34;;
    <a id="L16"></a>&#34;path&#34;;
    <a id="L17"></a>&#34;sort&#34;;
    <a id="L18"></a>&#34;strings&#34;;
<a id="L19"></a>)

<a id="L21"></a>var checkSync = flag.Bool(&#34;checksync&#34;, true, &#34;check whether repository is out of sync&#34;)

<a id="L23"></a>func usage() {
    <a id="L24"></a>fmt.Fprintf(os.Stderr, &#34;usage: hgpatch [options] [patchfile]\n&#34;);
    <a id="L25"></a>flag.PrintDefaults();
    <a id="L26"></a>os.Exit(2);
<a id="L27"></a>}

<a id="L29"></a>func main() {
    <a id="L30"></a>flag.Usage = usage;
    <a id="L31"></a>flag.Parse();

    <a id="L33"></a>args := flag.Args();
    <a id="L34"></a>var data []byte;
    <a id="L35"></a>var err os.Error;
    <a id="L36"></a>switch len(args) {
    <a id="L37"></a>case 0:
        <a id="L38"></a>data, err = io.ReadAll(os.Stdin)
    <a id="L39"></a>case 1:
        <a id="L40"></a>data, err = io.ReadFile(args[0])
    <a id="L41"></a>default:
        <a id="L42"></a>usage()
    <a id="L43"></a>}
    <a id="L44"></a>chk(err);

    <a id="L46"></a>pset, err := patch.Parse(data);
    <a id="L47"></a>chk(err);

    <a id="L49"></a><span class="comment">// Change to hg root directory, because</span>
    <a id="L50"></a><span class="comment">// patch paths are relative to root.</span>
    <a id="L51"></a>root, err := hgRoot();
    <a id="L52"></a>chk(err);
    <a id="L53"></a>chk(os.Chdir(root));

    <a id="L55"></a><span class="comment">// Make sure there are no pending changes on the server.</span>
    <a id="L56"></a>if *checkSync &amp;&amp; hgIncoming() {
        <a id="L57"></a>fmt.Fprintf(os.Stderr, &#34;incoming changes waiting; run hg sync first\n&#34;);
        <a id="L58"></a>os.Exit(2);
    <a id="L59"></a>}

    <a id="L61"></a><span class="comment">// Make sure we won&#39;t be editing files with local pending changes.</span>
    <a id="L62"></a>dirtylist, err := hgModified();
    <a id="L63"></a>chk(err);
    <a id="L64"></a>dirty := make(map[string]int);
    <a id="L65"></a>for _, f := range dirtylist {
        <a id="L66"></a>dirty[f] = 1
    <a id="L67"></a>}
    <a id="L68"></a>conflict := make(map[string]int);
    <a id="L69"></a>for _, f := range pset.File {
        <a id="L70"></a>if f.Verb == patch.Delete || f.Verb == patch.Rename {
            <a id="L71"></a>if _, ok := dirty[f.Src]; ok {
                <a id="L72"></a>conflict[f.Src] = 1
            <a id="L73"></a>}
        <a id="L74"></a>}
        <a id="L75"></a>if f.Verb != patch.Delete {
            <a id="L76"></a>if _, ok := dirty[f.Dst]; ok {
                <a id="L77"></a>conflict[f.Dst] = 1
            <a id="L78"></a>}
        <a id="L79"></a>}
    <a id="L80"></a>}
    <a id="L81"></a>if len(conflict) &gt; 0 {
        <a id="L82"></a>fmt.Fprintf(os.Stderr, &#34;cannot apply patch to locally modified files:\n&#34;);
        <a id="L83"></a>for name := range conflict {
            <a id="L84"></a>fmt.Fprintf(os.Stderr, &#34;\t%s\n&#34;, name)
        <a id="L85"></a>}
        <a id="L86"></a>os.Exit(2);
    <a id="L87"></a>}

    <a id="L89"></a><span class="comment">// Apply changes in memory.</span>
    <a id="L90"></a>op, err := pset.Apply(io.ReadFile);
    <a id="L91"></a>chk(err);

    <a id="L93"></a><span class="comment">// Write changes to disk copy: order of commands matters.</span>
    <a id="L94"></a><span class="comment">// Accumulate undo log as we go, in case there is an error.</span>
    <a id="L95"></a><span class="comment">// Also accumulate list of modified files to print at end.</span>
    <a id="L96"></a>changed := make(map[string]int);

    <a id="L98"></a><span class="comment">// Copy, Rename create the destination file, so they</span>
    <a id="L99"></a><span class="comment">// must happen before we write the data out.</span>
    <a id="L100"></a><span class="comment">// A single patch may have a Copy and a Rename</span>
    <a id="L101"></a><span class="comment">// with the same source, so we have to run all the</span>
    <a id="L102"></a><span class="comment">// Copy in one pass, then all the Rename.</span>
    <a id="L103"></a>for i := range op {
        <a id="L104"></a>o := &amp;op[i];
        <a id="L105"></a>if o.Verb == patch.Copy {
            <a id="L106"></a>makeParent(o.Dst);
            <a id="L107"></a>chk(hgCopy(o.Dst, o.Src));
            <a id="L108"></a>undoRevert(o.Dst);
            <a id="L109"></a>changed[o.Dst] = 1;
        <a id="L110"></a>}
    <a id="L111"></a>}
    <a id="L112"></a>for i := range op {
        <a id="L113"></a>o := &amp;op[i];
        <a id="L114"></a>if o.Verb == patch.Rename {
            <a id="L115"></a>makeParent(o.Dst);
            <a id="L116"></a>chk(hgRename(o.Dst, o.Src));
            <a id="L117"></a>undoRevert(o.Dst);
            <a id="L118"></a>undoRevert(o.Src);
            <a id="L119"></a>changed[o.Src] = 1;
            <a id="L120"></a>changed[o.Dst] = 1;
        <a id="L121"></a>}
    <a id="L122"></a>}

    <a id="L124"></a><span class="comment">// Run Delete before writing to files in case one of the</span>
    <a id="L125"></a><span class="comment">// deleted paths is becoming a directory.</span>
    <a id="L126"></a>for i := range op {
        <a id="L127"></a>o := &amp;op[i];
        <a id="L128"></a>if o.Verb == patch.Delete {
            <a id="L129"></a>chk(hgRemove(o.Src));
            <a id="L130"></a>undoRevert(o.Src);
            <a id="L131"></a>changed[o.Src] = 1;
        <a id="L132"></a>}
    <a id="L133"></a>}

    <a id="L135"></a><span class="comment">// Write files.</span>
    <a id="L136"></a>for i := range op {
        <a id="L137"></a>o := &amp;op[i];
        <a id="L138"></a>if o.Verb == patch.Delete {
            <a id="L139"></a>continue
        <a id="L140"></a>}
        <a id="L141"></a>if o.Verb == patch.Add {
            <a id="L142"></a>makeParent(o.Dst);
            <a id="L143"></a>changed[o.Dst] = 1;
        <a id="L144"></a>}
        <a id="L145"></a>if o.Data != nil {
            <a id="L146"></a>chk(io.WriteFile(o.Dst, o.Data, 0644));
            <a id="L147"></a>if o.Verb == patch.Add {
                <a id="L148"></a>undoRm(o.Dst)
            <a id="L149"></a>} else {
                <a id="L150"></a>undoRevert(o.Dst)
            <a id="L151"></a>}
            <a id="L152"></a>changed[o.Dst] = 1;
        <a id="L153"></a>}
        <a id="L154"></a>if o.Mode != 0 {
            <a id="L155"></a>chk(os.Chmod(o.Dst, o.Mode&amp;0755));
            <a id="L156"></a>undoRevert(o.Dst);
            <a id="L157"></a>changed[o.Dst] = 1;
        <a id="L158"></a>}
    <a id="L159"></a>}

    <a id="L161"></a><span class="comment">// hg add looks at the destination file, so it must happen</span>
    <a id="L162"></a><span class="comment">// after we write the data out.</span>
    <a id="L163"></a>for i := range op {
        <a id="L164"></a>o := &amp;op[i];
        <a id="L165"></a>if o.Verb == patch.Add {
            <a id="L166"></a>chk(hgAdd(o.Dst));
            <a id="L167"></a>undoRevert(o.Dst);
            <a id="L168"></a>changed[o.Dst] = 1;
        <a id="L169"></a>}
    <a id="L170"></a>}

    <a id="L172"></a><span class="comment">// Finished editing files.  Write the list of changed files to stdout.</span>
    <a id="L173"></a>list := make([]string, len(changed));
    <a id="L174"></a>i := 0;
    <a id="L175"></a>for f := range changed {
        <a id="L176"></a>list[i] = f;
        <a id="L177"></a>i++;
    <a id="L178"></a>}
    <a id="L179"></a>sort.SortStrings(list);
    <a id="L180"></a>for _, f := range list {
        <a id="L181"></a>fmt.Printf(&#34;%s\n&#34;, f)
    <a id="L182"></a>}
<a id="L183"></a>}


<a id="L186"></a><span class="comment">// make parent directory for name, if necessary</span>
<a id="L187"></a>func makeParent(name string) {
    <a id="L188"></a>parent, _ := path.Split(name);
    <a id="L189"></a>chk(mkdirAll(parent, 0755));
<a id="L190"></a>}

<a id="L192"></a><span class="comment">// Copy of os.MkdirAll but adds to undo log after</span>
<a id="L193"></a><span class="comment">// creating a directory.</span>
<a id="L194"></a>func mkdirAll(path string, perm int) os.Error {
    <a id="L195"></a>dir, err := os.Lstat(path);
    <a id="L196"></a>if err == nil {
        <a id="L197"></a>if dir.IsDirectory() {
            <a id="L198"></a>return nil
        <a id="L199"></a>}
        <a id="L200"></a>return &amp;os.PathError{&#34;mkdir&#34;, path, os.ENOTDIR};
    <a id="L201"></a>}

    <a id="L203"></a>i := len(path);
    <a id="L204"></a>for i &gt; 0 &amp;&amp; path[i-1] == &#39;/&#39; { <span class="comment">// Skip trailing slashes.</span>
        <a id="L205"></a>i--
    <a id="L206"></a>}

    <a id="L208"></a>j := i;
    <a id="L209"></a>for j &gt; 0 &amp;&amp; path[j-1] != &#39;/&#39; { <span class="comment">// Scan backward over element.</span>
        <a id="L210"></a>j--
    <a id="L211"></a>}

    <a id="L213"></a>if j &gt; 0 {
        <a id="L214"></a>err = mkdirAll(path[0:j-1], perm);
        <a id="L215"></a>if err != nil {
            <a id="L216"></a>return err
        <a id="L217"></a>}
    <a id="L218"></a>}

    <a id="L220"></a>err = os.Mkdir(path, perm);
    <a id="L221"></a>if err != nil {
        <a id="L222"></a><span class="comment">// Handle arguments like &#34;foo/.&#34; by</span>
        <a id="L223"></a><span class="comment">// double-checking that directory doesn&#39;t exist.</span>
        <a id="L224"></a>dir, err1 := os.Lstat(path);
        <a id="L225"></a>if err1 == nil &amp;&amp; dir.IsDirectory() {
            <a id="L226"></a>return nil
        <a id="L227"></a>}
        <a id="L228"></a>return err;
    <a id="L229"></a>}
    <a id="L230"></a>undoRm(path);
    <a id="L231"></a>return nil;
<a id="L232"></a>}

<a id="L234"></a><span class="comment">// If err != nil, process the undo log and exit.</span>
<a id="L235"></a>func chk(err os.Error) {
    <a id="L236"></a>if err != nil {
        <a id="L237"></a>fmt.Fprintf(os.Stderr, &#34;%s\n&#34;, err);
        <a id="L238"></a>runUndo();
        <a id="L239"></a>os.Exit(2);
    <a id="L240"></a>}
<a id="L241"></a>}


<a id="L244"></a><span class="comment">// Undo log</span>
<a id="L245"></a>type undo func() os.Error

<a id="L247"></a>var undoLog vector.Vector <span class="comment">// vector of undo</span>

<a id="L249"></a>func undoRevert(name string) { undoLog.Push(undo(func() os.Error { return hgRevert(name) })) }

<a id="L251"></a>func undoRm(name string) { undoLog.Push(undo(func() os.Error { return os.Remove(name) })) }

<a id="L253"></a>func runUndo() {
    <a id="L254"></a>for i := undoLog.Len() - 1; i &gt;= 0; i-- {
        <a id="L255"></a>if err := undoLog.At(i).(undo)(); err != nil {
            <a id="L256"></a>fmt.Fprintf(os.Stderr, &#34;%s\n&#34;, err)
        <a id="L257"></a>}
    <a id="L258"></a>}
<a id="L259"></a>}


<a id="L262"></a><span class="comment">// hgRoot returns the root directory of the repository.</span>
<a id="L263"></a>func hgRoot() (string, os.Error) {
    <a id="L264"></a>out, err := run([]string{&#34;hg&#34;, &#34;root&#34;}, nil);
    <a id="L265"></a>if err != nil {
        <a id="L266"></a>return &#34;&#34;, err
    <a id="L267"></a>}
    <a id="L268"></a>return strings.TrimSpace(out), nil;
<a id="L269"></a>}

<a id="L271"></a><span class="comment">// hgIncoming returns true if hg sync will pull in changes.</span>
<a id="L272"></a>func hgIncoming() bool {
    <a id="L273"></a><span class="comment">// hg -q incoming exits 0 when there is nothing incoming, 1 otherwise.</span>
    <a id="L274"></a>_, err := run([]string{&#34;hg&#34;, &#34;-q&#34;, &#34;incoming&#34;}, nil);
    <a id="L275"></a>return err == nil;
<a id="L276"></a>}

<a id="L278"></a><span class="comment">// hgModified returns a list of the modified files in the</span>
<a id="L279"></a><span class="comment">// repository.</span>
<a id="L280"></a>func hgModified() ([]string, os.Error) {
    <a id="L281"></a>out, err := run([]string{&#34;hg&#34;, &#34;status&#34;, &#34;-n&#34;}, nil);
    <a id="L282"></a>if err != nil {
        <a id="L283"></a>return nil, err
    <a id="L284"></a>}
    <a id="L285"></a>return strings.Split(strings.TrimSpace(out), &#34;\n&#34;, 0), nil;
<a id="L286"></a>}

<a id="L288"></a><span class="comment">// hgAdd adds name to the repository.</span>
<a id="L289"></a>func hgAdd(name string) os.Error {
    <a id="L290"></a>_, err := run([]string{&#34;hg&#34;, &#34;add&#34;, name}, nil);
    <a id="L291"></a>return err;
<a id="L292"></a>}

<a id="L294"></a><span class="comment">// hgRemove removes name from the repository.</span>
<a id="L295"></a>func hgRemove(name string) os.Error {
    <a id="L296"></a>_, err := run([]string{&#34;hg&#34;, &#34;rm&#34;, name}, nil);
    <a id="L297"></a>return err;
<a id="L298"></a>}

<a id="L300"></a><span class="comment">// hgRevert reverts name.</span>
<a id="L301"></a>func hgRevert(name string) os.Error {
    <a id="L302"></a>_, err := run([]string{&#34;hg&#34;, &#34;revert&#34;, name}, nil);
    <a id="L303"></a>return err;
<a id="L304"></a>}

<a id="L306"></a><span class="comment">// hgCopy copies src to dst in the repository.</span>
<a id="L307"></a><span class="comment">// Note that the argument order matches io.Copy, not &#34;hg cp&#34;.</span>
<a id="L308"></a>func hgCopy(dst, src string) os.Error {
    <a id="L309"></a>_, err := run([]string{&#34;hg&#34;, &#34;cp&#34;, src, dst}, nil);
    <a id="L310"></a>return err;
<a id="L311"></a>}

<a id="L313"></a><span class="comment">// hgRename renames src to dst in the repository.</span>
<a id="L314"></a><span class="comment">// Note that the argument order matches io.Copy, not &#34;hg mv&#34;.</span>
<a id="L315"></a>func hgRename(dst, src string) os.Error {
    <a id="L316"></a>_, err := run([]string{&#34;hg&#34;, &#34;mv&#34;, src, dst}, nil);
    <a id="L317"></a>return err;
<a id="L318"></a>}

<a id="L320"></a>func copy(a []string) []string {
    <a id="L321"></a>b := make([]string, len(a));
    <a id="L322"></a>for i, s := range a {
        <a id="L323"></a>b[i] = s
    <a id="L324"></a>}
    <a id="L325"></a>return b;
<a id="L326"></a>}

<a id="L328"></a>var lookPathCache = make(map[string]string)

<a id="L330"></a><span class="comment">// run runs the command argv, resolving argv[0] if necessary by searching $PATH.</span>
<a id="L331"></a><span class="comment">// It provides input on standard input to the command.</span>
<a id="L332"></a>func run(argv []string, input []byte) (out string, err os.Error) {
    <a id="L333"></a>if len(argv) &lt; 1 {
        <a id="L334"></a>err = os.EINVAL;
        <a id="L335"></a>goto Error;
    <a id="L336"></a>}
    <a id="L337"></a>prog, ok := lookPathCache[argv[0]];
    <a id="L338"></a>if !ok {
        <a id="L339"></a>prog, err = exec.LookPath(argv[0]);
        <a id="L340"></a>if err != nil {
            <a id="L341"></a>goto Error
        <a id="L342"></a>}
        <a id="L343"></a>lookPathCache[argv[0]] = prog;
    <a id="L344"></a>}
    <a id="L345"></a><span class="comment">// fmt.Fprintf(os.Stderr, &#34;%v\n&#34;, argv);</span>
    <a id="L346"></a>var cmd *exec.Cmd;
    <a id="L347"></a>if len(input) == 0 {
        <a id="L348"></a>cmd, err = exec.Run(prog, argv, os.Environ(), exec.DevNull, exec.Pipe, exec.MergeWithStdout);
        <a id="L349"></a>if err != nil {
            <a id="L350"></a>goto Error
        <a id="L351"></a>}
    <a id="L352"></a>} else {
        <a id="L353"></a>cmd, err = exec.Run(prog, argv, os.Environ(), exec.Pipe, exec.Pipe, exec.MergeWithStdout);
        <a id="L354"></a>if err != nil {
            <a id="L355"></a>goto Error
        <a id="L356"></a>}
        <a id="L357"></a>go func() {
            <a id="L358"></a>cmd.Stdin.Write(input);
            <a id="L359"></a>cmd.Stdin.Close();
        <a id="L360"></a>}();
    <a id="L361"></a>}
    <a id="L362"></a>defer cmd.Close();
    <a id="L363"></a>var buf bytes.Buffer;
    <a id="L364"></a>_, err = io.Copy(&amp;buf, cmd.Stdout);
    <a id="L365"></a>out = buf.String();
    <a id="L366"></a>if err != nil {
        <a id="L367"></a>cmd.Wait(0);
        <a id="L368"></a>goto Error;
    <a id="L369"></a>}
    <a id="L370"></a>w, err := cmd.Wait(0);
    <a id="L371"></a>if err != nil {
        <a id="L372"></a>goto Error
    <a id="L373"></a>}
    <a id="L374"></a>if !w.Exited() || w.ExitStatus() != 0 {
        <a id="L375"></a>err = w;
        <a id="L376"></a>goto Error;
    <a id="L377"></a>}
    <a id="L378"></a>return;

<a id="L380"></a>Error:
    <a id="L381"></a>err = &amp;runError{copy(argv), err};
    <a id="L382"></a>return;
<a id="L383"></a>}

<a id="L385"></a><span class="comment">// A runError represents an error that occurred while running a command.</span>
<a id="L386"></a>type runError struct {
    <a id="L387"></a>cmd []string;
    <a id="L388"></a>err os.Error;
<a id="L389"></a>}

<a id="L391"></a>func (e *runError) String() string { return strings.Join(e.cmd, &#34; &#34;) + &#34;: &#34; + e.err.String() }
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
