<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/os/os_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/os/os_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package os_test

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;fmt&#34;;
    <a id="L10"></a>&#34;io&#34;;
    <a id="L11"></a>. &#34;os&#34;;
    <a id="L12"></a>&#34;strings&#34;;
    <a id="L13"></a>&#34;testing&#34;;
<a id="L14"></a>)

<a id="L16"></a>var dot = []string{
    <a id="L17"></a>&#34;dir_darwin.go&#34;,
    <a id="L18"></a>&#34;dir_linux.go&#34;,
    <a id="L19"></a>&#34;env.go&#34;,
    <a id="L20"></a>&#34;error.go&#34;,
    <a id="L21"></a>&#34;file.go&#34;,
    <a id="L22"></a>&#34;os_test.go&#34;,
    <a id="L23"></a>&#34;time.go&#34;,
    <a id="L24"></a>&#34;types.go&#34;,
    <a id="L25"></a>&#34;stat_darwin.go&#34;,
    <a id="L26"></a>&#34;stat_linux.go&#34;,
<a id="L27"></a>}

<a id="L29"></a>var etc = []string{
    <a id="L30"></a>&#34;group&#34;,
    <a id="L31"></a>&#34;hosts&#34;,
    <a id="L32"></a>&#34;passwd&#34;,
<a id="L33"></a>}

<a id="L35"></a>func size(name string, t *testing.T) uint64 {
    <a id="L36"></a>file, err := Open(name, O_RDONLY, 0);
    <a id="L37"></a>defer file.Close();
    <a id="L38"></a>if err != nil {
        <a id="L39"></a>t.Fatal(&#34;open failed:&#34;, err)
    <a id="L40"></a>}
    <a id="L41"></a>var buf [100]byte;
    <a id="L42"></a>len := 0;
    <a id="L43"></a>for {
        <a id="L44"></a>n, e := file.Read(&amp;buf);
        <a id="L45"></a>len += n;
        <a id="L46"></a>if e == EOF {
            <a id="L47"></a>break
        <a id="L48"></a>}
        <a id="L49"></a>if e != nil {
            <a id="L50"></a>t.Fatal(&#34;read failed:&#34;, err)
        <a id="L51"></a>}
    <a id="L52"></a>}
    <a id="L53"></a>return uint64(len);
<a id="L54"></a>}

<a id="L56"></a>func TestStat(t *testing.T) {
    <a id="L57"></a>dir, err := Stat(&#34;/etc/passwd&#34;);
    <a id="L58"></a>if err != nil {
        <a id="L59"></a>t.Fatal(&#34;stat failed:&#34;, err)
    <a id="L60"></a>}
    <a id="L61"></a>if dir.Name != &#34;passwd&#34; {
        <a id="L62"></a>t.Error(&#34;name should be passwd; is&#34;, dir.Name)
    <a id="L63"></a>}
    <a id="L64"></a>filesize := size(&#34;/etc/passwd&#34;, t);
    <a id="L65"></a>if dir.Size != filesize {
        <a id="L66"></a>t.Error(&#34;size should be&#34;, filesize, &#34;; is&#34;, dir.Size)
    <a id="L67"></a>}
<a id="L68"></a>}

<a id="L70"></a>func TestFstat(t *testing.T) {
    <a id="L71"></a>file, err1 := Open(&#34;/etc/passwd&#34;, O_RDONLY, 0);
    <a id="L72"></a>defer file.Close();
    <a id="L73"></a>if err1 != nil {
        <a id="L74"></a>t.Fatal(&#34;open failed:&#34;, err1)
    <a id="L75"></a>}
    <a id="L76"></a>dir, err2 := file.Stat();
    <a id="L77"></a>if err2 != nil {
        <a id="L78"></a>t.Fatal(&#34;fstat failed:&#34;, err2)
    <a id="L79"></a>}
    <a id="L80"></a>if dir.Name != &#34;passwd&#34; {
        <a id="L81"></a>t.Error(&#34;name should be passwd; is&#34;, dir.Name)
    <a id="L82"></a>}
    <a id="L83"></a>filesize := size(&#34;/etc/passwd&#34;, t);
    <a id="L84"></a>if dir.Size != filesize {
        <a id="L85"></a>t.Error(&#34;size should be&#34;, filesize, &#34;; is&#34;, dir.Size)
    <a id="L86"></a>}
<a id="L87"></a>}

<a id="L89"></a>func TestLstat(t *testing.T) {
    <a id="L90"></a>dir, err := Lstat(&#34;/etc/passwd&#34;);
    <a id="L91"></a>if err != nil {
        <a id="L92"></a>t.Fatal(&#34;lstat failed:&#34;, err)
    <a id="L93"></a>}
    <a id="L94"></a>if dir.Name != &#34;passwd&#34; {
        <a id="L95"></a>t.Error(&#34;name should be passwd; is&#34;, dir.Name)
    <a id="L96"></a>}
    <a id="L97"></a>filesize := size(&#34;/etc/passwd&#34;, t);
    <a id="L98"></a>if dir.Size != filesize {
        <a id="L99"></a>t.Error(&#34;size should be&#34;, filesize, &#34;; is&#34;, dir.Size)
    <a id="L100"></a>}
<a id="L101"></a>}

<a id="L103"></a>func testReaddirnames(dir string, contents []string, t *testing.T) {
    <a id="L104"></a>file, err := Open(dir, O_RDONLY, 0);
    <a id="L105"></a>defer file.Close();
    <a id="L106"></a>if err != nil {
        <a id="L107"></a>t.Fatalf(&#34;open %q failed: %v&#34;, dir, err)
    <a id="L108"></a>}
    <a id="L109"></a>s, err2 := file.Readdirnames(-1);
    <a id="L110"></a>if err2 != nil {
        <a id="L111"></a>t.Fatalf(&#34;readdirnames %q failed: %v&#34;, err2)
    <a id="L112"></a>}
    <a id="L113"></a>for _, m := range contents {
        <a id="L114"></a>found := false;
        <a id="L115"></a>for _, n := range s {
            <a id="L116"></a>if n == &#34;.&#34; || n == &#34;..&#34; {
                <a id="L117"></a>t.Errorf(&#34;got %s in directory&#34;, n)
            <a id="L118"></a>}
            <a id="L119"></a>if m == n {
                <a id="L120"></a>if found {
                    <a id="L121"></a>t.Error(&#34;present twice:&#34;, m)
                <a id="L122"></a>}
                <a id="L123"></a>found = true;
            <a id="L124"></a>}
        <a id="L125"></a>}
        <a id="L126"></a>if !found {
            <a id="L127"></a>t.Error(&#34;could not find&#34;, m)
        <a id="L128"></a>}
    <a id="L129"></a>}
<a id="L130"></a>}

<a id="L132"></a>func testReaddir(dir string, contents []string, t *testing.T) {
    <a id="L133"></a>file, err := Open(dir, O_RDONLY, 0);
    <a id="L134"></a>defer file.Close();
    <a id="L135"></a>if err != nil {
        <a id="L136"></a>t.Fatalf(&#34;open %q failed: %v&#34;, dir, err)
    <a id="L137"></a>}
    <a id="L138"></a>s, err2 := file.Readdir(-1);
    <a id="L139"></a>if err2 != nil {
        <a id="L140"></a>t.Fatalf(&#34;readdir %q failed: %v&#34;, dir, err2)
    <a id="L141"></a>}
    <a id="L142"></a>for _, m := range contents {
        <a id="L143"></a>found := false;
        <a id="L144"></a>for _, n := range s {
            <a id="L145"></a>if m == n.Name {
                <a id="L146"></a>if found {
                    <a id="L147"></a>t.Error(&#34;present twice:&#34;, m)
                <a id="L148"></a>}
                <a id="L149"></a>found = true;
            <a id="L150"></a>}
        <a id="L151"></a>}
        <a id="L152"></a>if !found {
            <a id="L153"></a>t.Error(&#34;could not find&#34;, m)
        <a id="L154"></a>}
    <a id="L155"></a>}
<a id="L156"></a>}

<a id="L158"></a>func TestReaddirnames(t *testing.T) {
    <a id="L159"></a>testReaddirnames(&#34;.&#34;, dot, t);
    <a id="L160"></a>testReaddirnames(&#34;/etc&#34;, etc, t);
<a id="L161"></a>}

<a id="L163"></a>func TestReaddir(t *testing.T) {
    <a id="L164"></a>testReaddir(&#34;.&#34;, dot, t);
    <a id="L165"></a>testReaddir(&#34;/etc&#34;, etc, t);
<a id="L166"></a>}

<a id="L168"></a><span class="comment">// Read the directory one entry at a time.</span>
<a id="L169"></a>func smallReaddirnames(file *File, length int, t *testing.T) []string {
    <a id="L170"></a>names := make([]string, length);
    <a id="L171"></a>count := 0;
    <a id="L172"></a>for {
        <a id="L173"></a>d, err := file.Readdirnames(1);
        <a id="L174"></a>if err != nil {
            <a id="L175"></a>t.Fatalf(&#34;readdir %q failed: %v&#34;, file.Name(), err)
        <a id="L176"></a>}
        <a id="L177"></a>if len(d) == 0 {
            <a id="L178"></a>break
        <a id="L179"></a>}
        <a id="L180"></a>names[count] = d[0];
        <a id="L181"></a>count++;
    <a id="L182"></a>}
    <a id="L183"></a>return names[0:count];
<a id="L184"></a>}

<a id="L186"></a><span class="comment">// Check that reading a directory one entry at a time gives the same result</span>
<a id="L187"></a><span class="comment">// as reading it all at once.</span>
<a id="L188"></a>func TestReaddirnamesOneAtATime(t *testing.T) {
    <a id="L189"></a>dir := &#34;/usr/bin&#34;; <span class="comment">// big directory that doesn&#39;t change often.</span>
    <a id="L190"></a>file, err := Open(dir, O_RDONLY, 0);
    <a id="L191"></a>defer file.Close();
    <a id="L192"></a>if err != nil {
        <a id="L193"></a>t.Fatalf(&#34;open %q failed: %v&#34;, dir, err)
    <a id="L194"></a>}
    <a id="L195"></a>all, err1 := file.Readdirnames(-1);
    <a id="L196"></a>if err1 != nil {
        <a id="L197"></a>t.Fatalf(&#34;readdirnames %q failed: %v&#34;, dir, err1)
    <a id="L198"></a>}
    <a id="L199"></a>file1, err2 := Open(dir, O_RDONLY, 0);
    <a id="L200"></a>if err2 != nil {
        <a id="L201"></a>t.Fatalf(&#34;open %q failed: %v&#34;, dir, err2)
    <a id="L202"></a>}
    <a id="L203"></a>small := smallReaddirnames(file1, len(all)+100, t); <span class="comment">// +100 in case we screw up</span>
    <a id="L204"></a>for i, n := range all {
        <a id="L205"></a>if small[i] != n {
            <a id="L206"></a>t.Errorf(&#34;small read %q %q mismatch: %v&#34;, small[i], n)
        <a id="L207"></a>}
    <a id="L208"></a>}
<a id="L209"></a>}

<a id="L211"></a>func TestHardLink(t *testing.T) {
    <a id="L212"></a>from, to := &#34;hardlinktestfrom&#34;, &#34;hardlinktestto&#34;;
    <a id="L213"></a>Remove(from); <span class="comment">// Just in case.</span>
    <a id="L214"></a>file, err := Open(to, O_CREAT|O_WRONLY, 0666);
    <a id="L215"></a>if err != nil {
        <a id="L216"></a>t.Fatalf(&#34;open %q failed: %v&#34;, to, err)
    <a id="L217"></a>}
    <a id="L218"></a>defer Remove(to);
    <a id="L219"></a>if err = file.Close(); err != nil {
        <a id="L220"></a>t.Errorf(&#34;close %q failed: %v&#34;, to, err)
    <a id="L221"></a>}
    <a id="L222"></a>err = Link(to, from);
    <a id="L223"></a>if err != nil {
        <a id="L224"></a>t.Fatalf(&#34;link %q, %q failed: %v&#34;, to, from, err)
    <a id="L225"></a>}
    <a id="L226"></a>defer Remove(from);
    <a id="L227"></a>tostat, err := Stat(to);
    <a id="L228"></a>if err != nil {
        <a id="L229"></a>t.Fatalf(&#34;stat %q failed: %v&#34;, to, err)
    <a id="L230"></a>}
    <a id="L231"></a>fromstat, err := Stat(from);
    <a id="L232"></a>if err != nil {
        <a id="L233"></a>t.Fatalf(&#34;stat %q failed: %v&#34;, from, err)
    <a id="L234"></a>}
    <a id="L235"></a>if tostat.Dev != fromstat.Dev || tostat.Ino != fromstat.Ino {
        <a id="L236"></a>t.Errorf(&#34;link %q, %q did not create hard link&#34;, to, from)
    <a id="L237"></a>}
<a id="L238"></a>}

<a id="L240"></a>func TestSymLink(t *testing.T) {
    <a id="L241"></a>from, to := &#34;symlinktestfrom&#34;, &#34;symlinktestto&#34;;
    <a id="L242"></a>Remove(from); <span class="comment">// Just in case.</span>
    <a id="L243"></a>file, err := Open(to, O_CREAT|O_WRONLY, 0666);
    <a id="L244"></a>if err != nil {
        <a id="L245"></a>t.Fatalf(&#34;open %q failed: %v&#34;, to, err)
    <a id="L246"></a>}
    <a id="L247"></a>defer Remove(to);
    <a id="L248"></a>if err = file.Close(); err != nil {
        <a id="L249"></a>t.Errorf(&#34;close %q failed: %v&#34;, to, err)
    <a id="L250"></a>}
    <a id="L251"></a>err = Symlink(to, from);
    <a id="L252"></a>if err != nil {
        <a id="L253"></a>t.Fatalf(&#34;symlink %q, %q failed: %v&#34;, to, from, err)
    <a id="L254"></a>}
    <a id="L255"></a>defer Remove(from);
    <a id="L256"></a>tostat, err := Stat(to);
    <a id="L257"></a>if err != nil {
        <a id="L258"></a>t.Fatalf(&#34;stat %q failed: %v&#34;, to, err)
    <a id="L259"></a>}
    <a id="L260"></a>if tostat.FollowedSymlink {
        <a id="L261"></a>t.Fatalf(&#34;stat %q claims to have followed a symlink&#34;, to)
    <a id="L262"></a>}
    <a id="L263"></a>fromstat, err := Stat(from);
    <a id="L264"></a>if err != nil {
        <a id="L265"></a>t.Fatalf(&#34;stat %q failed: %v&#34;, from, err)
    <a id="L266"></a>}
    <a id="L267"></a>if tostat.Dev != fromstat.Dev || tostat.Ino != fromstat.Ino {
        <a id="L268"></a>t.Errorf(&#34;symlink %q, %q did not create symlink&#34;, to, from)
    <a id="L269"></a>}
    <a id="L270"></a>fromstat, err = Lstat(from);
    <a id="L271"></a>if err != nil {
        <a id="L272"></a>t.Fatalf(&#34;lstat %q failed: %v&#34;, from, err)
    <a id="L273"></a>}
    <a id="L274"></a>if !fromstat.IsSymlink() {
        <a id="L275"></a>t.Fatalf(&#34;symlink %q, %q did not create symlink&#34;, to, from)
    <a id="L276"></a>}
    <a id="L277"></a>fromstat, err = Stat(from);
    <a id="L278"></a>if err != nil {
        <a id="L279"></a>t.Fatalf(&#34;stat %q failed: %v&#34;, from, err)
    <a id="L280"></a>}
    <a id="L281"></a>if !fromstat.FollowedSymlink {
        <a id="L282"></a>t.Fatalf(&#34;stat %q did not follow symlink&#34;)
    <a id="L283"></a>}
    <a id="L284"></a>s, err := Readlink(from);
    <a id="L285"></a>if err != nil {
        <a id="L286"></a>t.Fatalf(&#34;readlink %q failed: %v&#34;, from, err)
    <a id="L287"></a>}
    <a id="L288"></a>if s != to {
        <a id="L289"></a>t.Fatalf(&#34;after symlink %q != %q&#34;, s, to)
    <a id="L290"></a>}
    <a id="L291"></a>file, err = Open(from, O_RDONLY, 0);
    <a id="L292"></a>if err != nil {
        <a id="L293"></a>t.Fatalf(&#34;open %q failed: %v&#34;, from, err)
    <a id="L294"></a>}
    <a id="L295"></a>file.Close();
<a id="L296"></a>}

<a id="L298"></a>func TestLongSymlink(t *testing.T) {
    <a id="L299"></a>s := &#34;0123456789abcdef&#34;;
    <a id="L300"></a><span class="comment">// Long, but not too long: a common limit is 255.</span>
    <a id="L301"></a>s = s + s + s + s + s + s + s + s + s + s + s + s + s + s + s;
    <a id="L302"></a>from := &#34;longsymlinktestfrom&#34;;
    <a id="L303"></a>err := Symlink(s, from);
    <a id="L304"></a>if err != nil {
        <a id="L305"></a>t.Fatalf(&#34;symlink %q, %q failed: %v&#34;, s, from, err)
    <a id="L306"></a>}
    <a id="L307"></a>defer Remove(from);
    <a id="L308"></a>r, err := Readlink(from);
    <a id="L309"></a>if err != nil {
        <a id="L310"></a>t.Fatalf(&#34;readlink %q failed: %v&#34;, from, err)
    <a id="L311"></a>}
    <a id="L312"></a>if r != s {
        <a id="L313"></a>t.Fatalf(&#34;after symlink %q != %q&#34;, r, s)
    <a id="L314"></a>}
<a id="L315"></a>}

<a id="L317"></a>func TestForkExec(t *testing.T) {
    <a id="L318"></a>r, w, err := Pipe();
    <a id="L319"></a>if err != nil {
        <a id="L320"></a>t.Fatalf(&#34;Pipe: %v&#34;, err)
    <a id="L321"></a>}
    <a id="L322"></a>pid, err := ForkExec(&#34;/bin/pwd&#34;, []string{&#34;pwd&#34;}, nil, &#34;/&#34;, []*File{nil, w, Stderr});
    <a id="L323"></a>if err != nil {
        <a id="L324"></a>t.Fatalf(&#34;ForkExec: %v&#34;, err)
    <a id="L325"></a>}
    <a id="L326"></a>w.Close();

    <a id="L328"></a>var b bytes.Buffer;
    <a id="L329"></a>io.Copy(&amp;b, r);
    <a id="L330"></a>output := b.String();
    <a id="L331"></a>expect := &#34;/\n&#34;;
    <a id="L332"></a>if output != expect {
        <a id="L333"></a>t.Errorf(&#34;exec /bin/pwd returned %q wanted %q&#34;, output, expect)
    <a id="L334"></a>}
    <a id="L335"></a>Wait(pid, 0);
<a id="L336"></a>}

<a id="L338"></a>func checkMode(t *testing.T, path string, mode uint32) {
    <a id="L339"></a>dir, err := Stat(path);
    <a id="L340"></a>if err != nil {
        <a id="L341"></a>t.Fatalf(&#34;Stat %q (looking for mode %#o): %s&#34;, path, mode, err)
    <a id="L342"></a>}
    <a id="L343"></a>if dir.Mode&amp;0777 != mode {
        <a id="L344"></a>t.Errorf(&#34;Stat %q: mode %#o want %#o&#34;, path, dir.Mode, 0777)
    <a id="L345"></a>}
<a id="L346"></a>}

<a id="L348"></a>func TestChmod(t *testing.T) {
    <a id="L349"></a>MkdirAll(&#34;_obj&#34;, 0777);
    <a id="L350"></a>const Path = &#34;_obj/_TestChmod_&#34;;
    <a id="L351"></a>fd, err := Open(Path, O_WRONLY|O_CREAT, 0666);
    <a id="L352"></a>if err != nil {
        <a id="L353"></a>t.Fatalf(&#34;create %s: %s&#34;, Path, err)
    <a id="L354"></a>}

    <a id="L356"></a>if err = Chmod(Path, 0456); err != nil {
        <a id="L357"></a>t.Fatalf(&#34;chmod %s 0456: %s&#34;, Path, err)
    <a id="L358"></a>}
    <a id="L359"></a>checkMode(t, Path, 0456);

    <a id="L361"></a>if err = fd.Chmod(0123); err != nil {
        <a id="L362"></a>t.Fatalf(&#34;fchmod %s 0123: %s&#34;, Path, err)
    <a id="L363"></a>}
    <a id="L364"></a>checkMode(t, Path, 0123);

    <a id="L366"></a>fd.Close();
    <a id="L367"></a>Remove(Path);
<a id="L368"></a>}

<a id="L370"></a>func checkUidGid(t *testing.T, path string, uid, gid int) {
    <a id="L371"></a>dir, err := Stat(path);
    <a id="L372"></a>if err != nil {
        <a id="L373"></a>t.Fatalf(&#34;Stat %q (looking for uid/gid %d/%d): %s&#34;, path, uid, gid, err)
    <a id="L374"></a>}
    <a id="L375"></a>if dir.Uid != uint32(uid) {
        <a id="L376"></a>t.Errorf(&#34;Stat %q: uid %d want %d&#34;, path, dir.Uid, uid)
    <a id="L377"></a>}
    <a id="L378"></a>if dir.Gid != uint32(gid) {
        <a id="L379"></a>t.Errorf(&#34;Stat %q: gid %d want %d&#34;, path, dir.Gid, gid)
    <a id="L380"></a>}
<a id="L381"></a>}

<a id="L383"></a>func TestChown(t *testing.T) {
    <a id="L384"></a><span class="comment">// Use /tmp, not _obj, to make sure we&#39;re on a local file system,</span>
    <a id="L385"></a><span class="comment">// so that the group ids returned by Getgroups will be allowed</span>
    <a id="L386"></a><span class="comment">// on the file.  If _obj is on NFS, the Getgroups groups are</span>
    <a id="L387"></a><span class="comment">// basically useless.</span>

    <a id="L389"></a>const Path = &#34;/tmp/_TestChown_&#34;;
    <a id="L390"></a>fd, err := Open(Path, O_WRONLY|O_CREAT, 0666);
    <a id="L391"></a>if err != nil {
        <a id="L392"></a>t.Fatalf(&#34;create %s: %s&#34;, Path, err)
    <a id="L393"></a>}
    <a id="L394"></a>dir, err := fd.Stat();
    <a id="L395"></a>if err != nil {
        <a id="L396"></a>t.Fatalf(&#34;fstat %s: %s&#34;, Path, err)
    <a id="L397"></a>}
    <a id="L398"></a>defer fd.Close();
    <a id="L399"></a>defer Remove(Path);

    <a id="L401"></a><span class="comment">// Can&#39;t change uid unless root, but can try</span>
    <a id="L402"></a><span class="comment">// changing the group id.  First try our current group.</span>
    <a id="L403"></a>gid := Getgid();
    <a id="L404"></a>t.Log(&#34;gid:&#34;, gid);
    <a id="L405"></a>if err = Chown(Path, -1, gid); err != nil {
        <a id="L406"></a>t.Fatalf(&#34;chown %s -1 %d: %s&#34;, Path, gid, err)
    <a id="L407"></a>}
    <a id="L408"></a>checkUidGid(t, Path, int(dir.Uid), gid);

    <a id="L410"></a><span class="comment">// Then try all the auxiliary groups.</span>
    <a id="L411"></a>groups, err := Getgroups();
    <a id="L412"></a>if err != nil {
        <a id="L413"></a>t.Fatalf(&#34;getgroups: %s&#34;, err)
    <a id="L414"></a>}
    <a id="L415"></a>t.Log(&#34;groups: &#34;, groups);
    <a id="L416"></a>for _, g := range groups {
        <a id="L417"></a>if err = Chown(Path, -1, g); err != nil {
            <a id="L418"></a>t.Fatalf(&#34;chown %s -1 %d: %s&#34;, Path, g, err)
        <a id="L419"></a>}
        <a id="L420"></a>checkUidGid(t, Path, int(dir.Uid), g);

        <a id="L422"></a><span class="comment">// change back to gid to test fd.Chown</span>
        <a id="L423"></a>if err = fd.Chown(-1, gid); err != nil {
            <a id="L424"></a>t.Fatalf(&#34;fchown %s -1 %d: %s&#34;, Path, gid, err)
        <a id="L425"></a>}
        <a id="L426"></a>checkUidGid(t, Path, int(dir.Uid), gid);
    <a id="L427"></a>}
<a id="L428"></a>}

<a id="L430"></a>func checkSize(t *testing.T, path string, size uint64) {
    <a id="L431"></a>dir, err := Stat(path);
    <a id="L432"></a>if err != nil {
        <a id="L433"></a>t.Fatalf(&#34;Stat %q (looking for size %d): %s&#34;, path, size, err)
    <a id="L434"></a>}
    <a id="L435"></a>if dir.Size != size {
        <a id="L436"></a>t.Errorf(&#34;Stat %q: size %d want %d&#34;, path, dir.Size, size)
    <a id="L437"></a>}
<a id="L438"></a>}

<a id="L440"></a>func TestTruncate(t *testing.T) {
    <a id="L441"></a>MkdirAll(&#34;_obj&#34;, 0777);
    <a id="L442"></a>const Path = &#34;_obj/_TestTruncate_&#34;;
    <a id="L443"></a>fd, err := Open(Path, O_WRONLY|O_CREAT, 0666);
    <a id="L444"></a>if err != nil {
        <a id="L445"></a>t.Fatalf(&#34;create %s: %s&#34;, Path, err)
    <a id="L446"></a>}

    <a id="L448"></a>checkSize(t, Path, 0);
    <a id="L449"></a>fd.Write(strings.Bytes(&#34;hello, world\n&#34;));
    <a id="L450"></a>checkSize(t, Path, 13);
    <a id="L451"></a>fd.Truncate(10);
    <a id="L452"></a>checkSize(t, Path, 10);
    <a id="L453"></a>fd.Truncate(1024);
    <a id="L454"></a>checkSize(t, Path, 1024);
    <a id="L455"></a>fd.Truncate(0);
    <a id="L456"></a>checkSize(t, Path, 0);
    <a id="L457"></a>fd.Write(strings.Bytes(&#34;surprise!&#34;));
    <a id="L458"></a>checkSize(t, Path, 13+9); <span class="comment">// wrote at offset past where hello, world was.</span>
    <a id="L459"></a>fd.Close();
    <a id="L460"></a>Remove(Path);
<a id="L461"></a>}

<a id="L463"></a>func TestChdirAndGetwd(t *testing.T) {
    <a id="L464"></a>fd, err := Open(&#34;.&#34;, O_RDONLY, 0);
    <a id="L465"></a>if err != nil {
        <a id="L466"></a>t.Fatalf(&#34;Open .: %s&#34;, err)
    <a id="L467"></a>}
    <a id="L468"></a><span class="comment">// These are chosen carefully not to be symlinks on a Mac</span>
    <a id="L469"></a><span class="comment">// (unlike, say, /var, /etc, and /tmp).</span>
    <a id="L470"></a>dirs := []string{&#34;/bin&#34;, &#34;/&#34;, &#34;/usr/bin&#34;};
    <a id="L471"></a>for mode := 0; mode &lt; 2; mode++ {
        <a id="L472"></a>for _, d := range dirs {
            <a id="L473"></a>if mode == 0 {
                <a id="L474"></a>err = Chdir(d)
            <a id="L475"></a>} else {
                <a id="L476"></a>fd1, err := Open(d, O_RDONLY, 0);
                <a id="L477"></a>if err != nil {
                    <a id="L478"></a>t.Errorf(&#34;Open %s: %s&#34;, d, err);
                    <a id="L479"></a>continue;
                <a id="L480"></a>}
                <a id="L481"></a>err = fd1.Chdir();
                <a id="L482"></a>fd1.Close();
            <a id="L483"></a>}
            <a id="L484"></a>pwd, err1 := Getwd();
            <a id="L485"></a>err2 := fd.Chdir();
            <a id="L486"></a>if err2 != nil {
                <a id="L487"></a><span class="comment">// We changed the current directory and cannot go back.</span>
                <a id="L488"></a><span class="comment">// Don&#39;t let the tests continue; they&#39;ll scribble</span>
                <a id="L489"></a><span class="comment">// all over some other directory.</span>
                <a id="L490"></a>fmt.Fprintf(Stderr, &#34;fchdir back to dot failed: %s\n&#34;, err2);
                <a id="L491"></a>Exit(1);
            <a id="L492"></a>}
            <a id="L493"></a>if err != nil {
                <a id="L494"></a>fd.Close();
                <a id="L495"></a>t.Fatalf(&#34;Chdir %s: %s&#34;, d, err);
            <a id="L496"></a>}
            <a id="L497"></a>if err1 != nil {
                <a id="L498"></a>fd.Close();
                <a id="L499"></a>t.Fatalf(&#34;Getwd in %s: %s&#34;, d, err1);
            <a id="L500"></a>}
            <a id="L501"></a>if pwd != d {
                <a id="L502"></a>fd.Close();
                <a id="L503"></a>t.Fatalf(&#34;Getwd returned %q want %q&#34;, pwd, d);
            <a id="L504"></a>}
        <a id="L505"></a>}
    <a id="L506"></a>}
    <a id="L507"></a>fd.Close();
<a id="L508"></a>}

<a id="L510"></a>func TestTime(t *testing.T) {
    <a id="L511"></a><span class="comment">// Just want to check that Time() is getting something.</span>
    <a id="L512"></a><span class="comment">// A common failure mode on Darwin is to get 0, 0,</span>
    <a id="L513"></a><span class="comment">// because it returns the time in registers instead of</span>
    <a id="L514"></a><span class="comment">// filling in the structure passed to the system call.</span>
    <a id="L515"></a><span class="comment">// Too bad the compiler doesn&#39;t know that</span>
    <a id="L516"></a><span class="comment">// 365.24*86400 is an integer.</span>
    <a id="L517"></a>sec, nsec, err := Time();
    <a id="L518"></a>if sec &lt; (2009-1970)*36524*864 {
        <a id="L519"></a>t.Errorf(&#34;Time() = %d, %d, %s; not plausible&#34;, sec, nsec, err)
    <a id="L520"></a>}
<a id="L521"></a>}

<a id="L523"></a>func TestSeek(t *testing.T) {
    <a id="L524"></a>f, err := Open(&#34;_obj/seektest&#34;, O_CREAT|O_RDWR|O_TRUNC, 0666);
    <a id="L525"></a>if err != nil {
        <a id="L526"></a>t.Fatalf(&#34;open _obj/seektest: %s&#34;, err)
    <a id="L527"></a>}

    <a id="L529"></a>const data = &#34;hello, world\n&#34;;
    <a id="L530"></a>io.WriteString(f, data);

    <a id="L532"></a>type test struct {
        <a id="L533"></a>in     int64;
        <a id="L534"></a>whence int;
        <a id="L535"></a>out    int64;
    <a id="L536"></a>}
    <a id="L537"></a>var tests = []test{
        <a id="L538"></a>test{0, 1, int64(len(data))},
        <a id="L539"></a>test{0, 0, 0},
        <a id="L540"></a>test{5, 0, 5},
        <a id="L541"></a>test{0, 2, int64(len(data))},
        <a id="L542"></a>test{0, 0, 0},
        <a id="L543"></a>test{-1, 2, int64(len(data)) - 1},
        <a id="L544"></a>test{1 &lt;&lt; 33, 0, 1 &lt;&lt; 33},
        <a id="L545"></a>test{1 &lt;&lt; 33, 2, 1&lt;&lt;33 + int64(len(data))},
    <a id="L546"></a>};
    <a id="L547"></a>for i, tt := range tests {
        <a id="L548"></a>off, err := f.Seek(tt.in, tt.whence);
        <a id="L549"></a>if off != tt.out || err != nil {
            <a id="L550"></a>if e, ok := err.(*PathError); ok &amp;&amp; e.Error == EINVAL &amp;&amp; tt.out &gt; 1&lt;&lt;32 {
                <a id="L551"></a><span class="comment">// Reiserfs rejects the big seeks.</span>
                <a id="L552"></a><span class="comment">// http://code.google.com/p/go/issues/detail?id=91</span>
                <a id="L553"></a>break
            <a id="L554"></a>}
            <a id="L555"></a>t.Errorf(&#34;#%d: Seek(%v, %v) = %v, %v want %v, nil&#34;, i, tt.in, tt.whence, off, err, tt.out);
        <a id="L556"></a>}
    <a id="L557"></a>}
    <a id="L558"></a>f.Close();
<a id="L559"></a>}

<a id="L561"></a>type openErrorTest struct {
    <a id="L562"></a>path  string;
    <a id="L563"></a>mode  int;
    <a id="L564"></a>error string;
<a id="L565"></a>}

<a id="L567"></a>var openErrorTests = []openErrorTest{
    <a id="L568"></a>openErrorTest{
        <a id="L569"></a>&#34;/etc/no-such-file&#34;,
        <a id="L570"></a>O_RDONLY,
        <a id="L571"></a>&#34;open /etc/no-such-file: no such file or directory&#34;,
    <a id="L572"></a>},
    <a id="L573"></a>openErrorTest{
        <a id="L574"></a>&#34;/etc&#34;,
        <a id="L575"></a>O_WRONLY,
        <a id="L576"></a>&#34;open /etc: is a directory&#34;,
    <a id="L577"></a>},
    <a id="L578"></a>openErrorTest{
        <a id="L579"></a>&#34;/etc/passwd/group&#34;,
        <a id="L580"></a>O_WRONLY,
        <a id="L581"></a>&#34;open /etc/passwd/group: not a directory&#34;,
    <a id="L582"></a>},
<a id="L583"></a>}

<a id="L585"></a>func TestOpenError(t *testing.T) {
    <a id="L586"></a>for _, tt := range openErrorTests {
        <a id="L587"></a>f, err := Open(tt.path, tt.mode, 0);
        <a id="L588"></a>if err == nil {
            <a id="L589"></a>t.Errorf(&#34;Open(%q, %d) succeeded&#34;, tt.path, tt.mode);
            <a id="L590"></a>f.Close();
            <a id="L591"></a>continue;
        <a id="L592"></a>}
        <a id="L593"></a>if s := err.String(); s != tt.error {
            <a id="L594"></a>t.Errorf(&#34;Open(%q, %d) = _, %q; want %q&#34;, tt.path, tt.mode, s, tt.error)
        <a id="L595"></a>}
    <a id="L596"></a>}
<a id="L597"></a>}

<a id="L599"></a>func run(t *testing.T, cmd []string) string {
    <a id="L600"></a><span class="comment">// Run /bin/hostname and collect output.</span>
    <a id="L601"></a>r, w, err := Pipe();
    <a id="L602"></a>if err != nil {
        <a id="L603"></a>t.Fatal(err)
    <a id="L604"></a>}
    <a id="L605"></a>pid, err := ForkExec(&#34;/bin/hostname&#34;, []string{&#34;hostname&#34;}, nil, &#34;/&#34;, []*File{nil, w, Stderr});
    <a id="L606"></a>if err != nil {
        <a id="L607"></a>t.Fatal(err)
    <a id="L608"></a>}
    <a id="L609"></a>w.Close();

    <a id="L611"></a>var b bytes.Buffer;
    <a id="L612"></a>io.Copy(&amp;b, r);
    <a id="L613"></a>Wait(pid, 0);
    <a id="L614"></a>output := b.String();
    <a id="L615"></a>if n := len(output); n &gt; 0 &amp;&amp; output[n-1] == &#39;\n&#39; {
        <a id="L616"></a>output = output[0 : n-1]
    <a id="L617"></a>}
    <a id="L618"></a>if output == &#34;&#34; {
        <a id="L619"></a>t.Fatalf(&#34;%v produced no output&#34;, cmd)
    <a id="L620"></a>}

    <a id="L622"></a>return output;
<a id="L623"></a>}


<a id="L626"></a>func TestHostname(t *testing.T) {
    <a id="L627"></a><span class="comment">// Check internal Hostname() against the output of /bin/hostname.</span>
    <a id="L628"></a>hostname, err := Hostname();
    <a id="L629"></a>if err != nil {
        <a id="L630"></a>t.Fatalf(&#34;%v&#34;, err)
    <a id="L631"></a>}
    <a id="L632"></a>want := run(t, []string{&#34;/bin/hostname&#34;});
    <a id="L633"></a>if hostname != want {
        <a id="L634"></a>t.Errorf(&#34;Hostname() = %q, want %q&#34;, hostname, want)
    <a id="L635"></a>}
<a id="L636"></a>}

<a id="L638"></a>func TestReadAt(t *testing.T) {
    <a id="L639"></a>f, err := Open(&#34;_obj/readtest&#34;, O_CREAT|O_RDWR|O_TRUNC, 0666);
    <a id="L640"></a>if err != nil {
        <a id="L641"></a>t.Fatalf(&#34;open _obj/readtest: %s&#34;, err)
    <a id="L642"></a>}
    <a id="L643"></a>const data = &#34;hello, world\n&#34;;
    <a id="L644"></a>io.WriteString(f, data);

    <a id="L646"></a>b := make([]byte, 5);
    <a id="L647"></a>n, err := f.ReadAt(b, 7);
    <a id="L648"></a>if err != nil || n != len(b) {
        <a id="L649"></a>t.Fatalf(&#34;ReadAt 7: %d, %r&#34;, n, err)
    <a id="L650"></a>}
    <a id="L651"></a>if string(b) != &#34;world&#34; {
        <a id="L652"></a>t.Fatalf(&#34;ReadAt 7: have %q want %q&#34;, string(b), &#34;world&#34;)
    <a id="L653"></a>}
<a id="L654"></a>}

<a id="L656"></a>func TestWriteAt(t *testing.T) {
    <a id="L657"></a>f, err := Open(&#34;_obj/writetest&#34;, O_CREAT|O_RDWR|O_TRUNC, 0666);
    <a id="L658"></a>if err != nil {
        <a id="L659"></a>t.Fatalf(&#34;open _obj/writetest: %s&#34;, err)
    <a id="L660"></a>}
    <a id="L661"></a>const data = &#34;hello, world\n&#34;;
    <a id="L662"></a>io.WriteString(f, data);

    <a id="L664"></a>n, err := f.WriteAt(strings.Bytes(&#34;WORLD&#34;), 7);
    <a id="L665"></a>if err != nil || n != 5 {
        <a id="L666"></a>t.Fatalf(&#34;WriteAt 7: %d, %v&#34;, n, err)
    <a id="L667"></a>}

    <a id="L669"></a>b, err := io.ReadFile(&#34;_obj/writetest&#34;);
    <a id="L670"></a>if err != nil {
        <a id="L671"></a>t.Fatalf(&#34;ReadFile _obj/writetest: %v&#34;, err)
    <a id="L672"></a>}
    <a id="L673"></a>if string(b) != &#34;hello, WORLD\n&#34; {
        <a id="L674"></a>t.Fatalf(&#34;after write: have %q want %q&#34;, string(b), &#34;hello, WORLD\n&#34;)
    <a id="L675"></a>}
<a id="L676"></a>}
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
