<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/os/path_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/os/path_test.go</h1>

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
    <a id="L8"></a>. &#34;os&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a>func TestMkdirAll(t *testing.T) {
    <a id="L13"></a><span class="comment">// Create new dir, in _obj so it will get</span>
    <a id="L14"></a><span class="comment">// cleaned up by make if not by us.</span>
    <a id="L15"></a>path := &#34;_obj/_TestMkdirAll_/dir/./dir2&#34;;
    <a id="L16"></a>err := MkdirAll(path, 0777);
    <a id="L17"></a>if err != nil {
        <a id="L18"></a>t.Fatalf(&#34;MkdirAll %q: %s&#34;, path, err)
    <a id="L19"></a>}

    <a id="L21"></a><span class="comment">// Already exists, should succeed.</span>
    <a id="L22"></a>err = MkdirAll(path, 0777);
    <a id="L23"></a>if err != nil {
        <a id="L24"></a>t.Fatalf(&#34;MkdirAll %q (second time): %s&#34;, path, err)
    <a id="L25"></a>}

    <a id="L27"></a><span class="comment">// Make file.</span>
    <a id="L28"></a>fpath := path + &#34;/file&#34;;
    <a id="L29"></a>_, err = Open(fpath, O_WRONLY|O_CREAT, 0666);
    <a id="L30"></a>if err != nil {
        <a id="L31"></a>t.Fatalf(&#34;create %q: %s&#34;, fpath, err)
    <a id="L32"></a>}

    <a id="L34"></a><span class="comment">// Can&#39;t make directory named after file.</span>
    <a id="L35"></a>err = MkdirAll(fpath, 0777);
    <a id="L36"></a>if err == nil {
        <a id="L37"></a>t.Fatalf(&#34;MkdirAll %q: no error&#34;)
    <a id="L38"></a>}
    <a id="L39"></a>perr, ok := err.(*PathError);
    <a id="L40"></a>if !ok {
        <a id="L41"></a>t.Fatalf(&#34;MkdirAll %q returned %T, not *PathError&#34;, fpath, err)
    <a id="L42"></a>}
    <a id="L43"></a>if perr.Path != fpath {
        <a id="L44"></a>t.Fatalf(&#34;MkdirAll %q returned wrong error path: %q not %q&#34;, fpath, perr.Path, fpath)
    <a id="L45"></a>}

    <a id="L47"></a><span class="comment">// Can&#39;t make subdirectory of file.</span>
    <a id="L48"></a>ffpath := fpath + &#34;/subdir&#34;;
    <a id="L49"></a>err = MkdirAll(ffpath, 0777);
    <a id="L50"></a>if err == nil {
        <a id="L51"></a>t.Fatalf(&#34;MkdirAll %q: no error&#34;)
    <a id="L52"></a>}
    <a id="L53"></a>perr, ok = err.(*PathError);
    <a id="L54"></a>if !ok {
        <a id="L55"></a>t.Fatalf(&#34;MkdirAll %q returned %T, not *PathError&#34;, ffpath, err)
    <a id="L56"></a>}
    <a id="L57"></a>if perr.Path != fpath {
        <a id="L58"></a>t.Fatalf(&#34;MkdirAll %q returned wrong error path: %q not %q&#34;, ffpath, perr.Path, fpath)
    <a id="L59"></a>}

    <a id="L61"></a>RemoveAll(&#34;_obj/_TestMkdirAll_&#34;);
<a id="L62"></a>}

<a id="L64"></a>func TestRemoveAll(t *testing.T) {
    <a id="L65"></a><span class="comment">// Work directory.</span>
    <a id="L66"></a>path := &#34;_obj/_TestRemoveAll_&#34;;
    <a id="L67"></a>fpath := path + &#34;/file&#34;;
    <a id="L68"></a>dpath := path + &#34;/dir&#34;;

    <a id="L70"></a><span class="comment">// Make directory with 1 file and remove.</span>
    <a id="L71"></a>if err := MkdirAll(path, 0777); err != nil {
        <a id="L72"></a>t.Fatalf(&#34;MkdirAll %q: %s&#34;, path, err)
    <a id="L73"></a>}
    <a id="L74"></a>fd, err := Open(fpath, O_WRONLY|O_CREAT, 0666);
    <a id="L75"></a>if err != nil {
        <a id="L76"></a>t.Fatalf(&#34;create %q: %s&#34;, fpath, err)
    <a id="L77"></a>}
    <a id="L78"></a>fd.Close();
    <a id="L79"></a>if err = RemoveAll(path); err != nil {
        <a id="L80"></a>t.Fatalf(&#34;RemoveAll %q (first): %s&#34;, path, err)
    <a id="L81"></a>}
    <a id="L82"></a>if _, err := Lstat(path); err == nil {
        <a id="L83"></a>t.Fatalf(&#34;Lstat %q succeeded after RemoveAll (first)&#34;, path)
    <a id="L84"></a>}

    <a id="L86"></a><span class="comment">// Make directory with file and subdirectory and remove.</span>
    <a id="L87"></a>if err = MkdirAll(dpath, 0777); err != nil {
        <a id="L88"></a>t.Fatalf(&#34;MkdirAll %q: %s&#34;, dpath, err)
    <a id="L89"></a>}
    <a id="L90"></a>fd, err = Open(fpath, O_WRONLY|O_CREAT, 0666);
    <a id="L91"></a>if err != nil {
        <a id="L92"></a>t.Fatalf(&#34;create %q: %s&#34;, fpath, err)
    <a id="L93"></a>}
    <a id="L94"></a>fd.Close();
    <a id="L95"></a>fd, err = Open(dpath+&#34;/file&#34;, O_WRONLY|O_CREAT, 0666);
    <a id="L96"></a>if err != nil {
        <a id="L97"></a>t.Fatalf(&#34;create %q: %s&#34;, fpath, err)
    <a id="L98"></a>}
    <a id="L99"></a>fd.Close();
    <a id="L100"></a>if err = RemoveAll(path); err != nil {
        <a id="L101"></a>t.Fatalf(&#34;RemoveAll %q (second): %s&#34;, path, err)
    <a id="L102"></a>}
    <a id="L103"></a>if _, err := Lstat(path); err == nil {
        <a id="L104"></a>t.Fatalf(&#34;Lstat %q succeeded after RemoveAll (second)&#34;, path)
    <a id="L105"></a>}

    <a id="L107"></a>if Getuid() != 0 { <span class="comment">// Test fails as root</span>
        <a id="L108"></a><span class="comment">// Make directory with file and subdirectory and trigger error.</span>
        <a id="L109"></a>if err = MkdirAll(dpath, 0777); err != nil {
            <a id="L110"></a>t.Fatalf(&#34;MkdirAll %q: %s&#34;, dpath, err)
        <a id="L111"></a>}

        <a id="L113"></a>for _, s := range []string{fpath, dpath + &#34;/file1&#34;, path + &#34;/zzz&#34;} {
            <a id="L114"></a>fd, err = Open(s, O_WRONLY|O_CREAT, 0666);
            <a id="L115"></a>if err != nil {
                <a id="L116"></a>t.Fatalf(&#34;create %q: %s&#34;, s, err)
            <a id="L117"></a>}
            <a id="L118"></a>fd.Close();
        <a id="L119"></a>}
        <a id="L120"></a>if err = Chmod(dpath, 0); err != nil {
            <a id="L121"></a>t.Fatalf(&#34;Chmod %q 0: %s&#34;, dpath, err)
        <a id="L122"></a>}
        <a id="L123"></a>if err = RemoveAll(path); err == nil {
            <a id="L124"></a>_, err := Lstat(path);
            <a id="L125"></a>if err == nil {
                <a id="L126"></a>t.Errorf(&#34;Can lstat %q after supposed RemoveAll&#34;, path)
            <a id="L127"></a>}
            <a id="L128"></a>t.Fatalf(&#34;RemoveAll %q succeeded with chmod 0 subdirectory&#34;, path, err);
        <a id="L129"></a>}
        <a id="L130"></a>perr, ok := err.(*PathError);
        <a id="L131"></a>if !ok {
            <a id="L132"></a>t.Fatalf(&#34;RemoveAll %q returned %T not *PathError&#34;, path, err)
        <a id="L133"></a>}
        <a id="L134"></a>if perr.Path != dpath {
            <a id="L135"></a>t.Fatalf(&#34;RemoveAll %q failed at %q not %q&#34;, path, perr.Path, dpath)
        <a id="L136"></a>}
        <a id="L137"></a>if err = Chmod(dpath, 0777); err != nil {
            <a id="L138"></a>t.Fatalf(&#34;Chmod %q 0777: %s&#34;, dpath, err)
        <a id="L139"></a>}
        <a id="L140"></a>for _, s := range []string{fpath, path + &#34;/zzz&#34;} {
            <a id="L141"></a>if _, err := Lstat(s); err == nil {
                <a id="L142"></a>t.Fatalf(&#34;Lstat %q succeeded after partial RemoveAll&#34;, s)
            <a id="L143"></a>}
        <a id="L144"></a>}
    <a id="L145"></a>}
    <a id="L146"></a>if err = RemoveAll(path); err != nil {
        <a id="L147"></a>t.Fatalf(&#34;RemoveAll %q after partial RemoveAll: %s&#34;, path, err)
    <a id="L148"></a>}
    <a id="L149"></a>if _, err := Lstat(path); err == nil {
        <a id="L150"></a>t.Fatalf(&#34;Lstat %q succeeded after RemoveAll (final)&#34;, path)
    <a id="L151"></a>}
<a id="L152"></a>}
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
