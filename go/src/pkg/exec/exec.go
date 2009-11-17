<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exec/exec.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exec/exec.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The exec package runs external commands.</span>
<a id="L6"></a>package exec

<a id="L8"></a>import (
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>&#34;strings&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">// Arguments to Run.</span>
<a id="L14"></a>const (
    <a id="L15"></a>DevNull = iota;
    <a id="L16"></a>PassThrough;
    <a id="L17"></a>Pipe;
    <a id="L18"></a>MergeWithStdout;
<a id="L19"></a>)

<a id="L21"></a><span class="comment">// A Cmd represents a running command.</span>
<a id="L22"></a><span class="comment">// Stdin, Stdout, and Stderr are Files representing pipes</span>
<a id="L23"></a><span class="comment">// connected to the running command&#39;s standard input, output, and error,</span>
<a id="L24"></a><span class="comment">// or else nil, depending on the arguments to Run.</span>
<a id="L25"></a><span class="comment">// Pid is the running command&#39;s operating system process ID.</span>
<a id="L26"></a>type Cmd struct {
    <a id="L27"></a>Stdin  *os.File;
    <a id="L28"></a>Stdout *os.File;
    <a id="L29"></a>Stderr *os.File;
    <a id="L30"></a>Pid    int;
<a id="L31"></a>}

<a id="L33"></a><span class="comment">// Given mode (DevNull, etc), return file for child</span>
<a id="L34"></a><span class="comment">// and file to record in Cmd structure.</span>
<a id="L35"></a>func modeToFiles(mode, fd int) (*os.File, *os.File, os.Error) {
    <a id="L36"></a>switch mode {
    <a id="L37"></a>case DevNull:
        <a id="L38"></a>rw := os.O_WRONLY;
        <a id="L39"></a>if fd == 0 {
            <a id="L40"></a>rw = os.O_RDONLY
        <a id="L41"></a>}
        <a id="L42"></a>f, err := os.Open(&#34;/dev/null&#34;, rw, 0);
        <a id="L43"></a>return f, nil, err;
    <a id="L44"></a>case PassThrough:
        <a id="L45"></a>switch fd {
        <a id="L46"></a>case 0:
            <a id="L47"></a>return os.Stdin, nil, nil
        <a id="L48"></a>case 1:
            <a id="L49"></a>return os.Stdout, nil, nil
        <a id="L50"></a>case 2:
            <a id="L51"></a>return os.Stderr, nil, nil
        <a id="L52"></a>}
    <a id="L53"></a>case Pipe:
        <a id="L54"></a>r, w, err := os.Pipe();
        <a id="L55"></a>if err != nil {
            <a id="L56"></a>return nil, nil, err
        <a id="L57"></a>}
        <a id="L58"></a>if fd == 0 {
            <a id="L59"></a>return r, w, nil
        <a id="L60"></a>}
        <a id="L61"></a>return w, r, nil;
    <a id="L62"></a>}
    <a id="L63"></a>return nil, nil, os.EINVAL;
<a id="L64"></a>}

<a id="L66"></a><span class="comment">// Run starts the binary prog running with</span>
<a id="L67"></a><span class="comment">// arguments argv and environment envv.</span>
<a id="L68"></a><span class="comment">// It returns a pointer to a new Cmd representing</span>
<a id="L69"></a><span class="comment">// the command or an error.</span>
<a id="L70"></a><span class="comment">//</span>
<a id="L71"></a><span class="comment">// The parameters stdin, stdout, and stderr</span>
<a id="L72"></a><span class="comment">// specify how to handle standard input, output, and error.</span>
<a id="L73"></a><span class="comment">// The choices are DevNull (connect to /dev/null),</span>
<a id="L74"></a><span class="comment">// PassThrough (connect to the current process&#39;s standard stream),</span>
<a id="L75"></a><span class="comment">// Pipe (connect to an operating system pipe), and</span>
<a id="L76"></a><span class="comment">// MergeWithStdout (only for standard error; use the same</span>
<a id="L77"></a><span class="comment">// file descriptor as was used for standard output).</span>
<a id="L78"></a><span class="comment">// If a parameter is Pipe, then the corresponding field (Stdin, Stdout, Stderr)</span>
<a id="L79"></a><span class="comment">// of the returned Cmd is the other end of the pipe.</span>
<a id="L80"></a><span class="comment">// Otherwise the field in Cmd is nil.</span>
<a id="L81"></a>func Run(argv0 string, argv, envv []string, stdin, stdout, stderr int) (p *Cmd, err os.Error) {
    <a id="L82"></a>p = new(Cmd);
    <a id="L83"></a>var fd [3]*os.File;

    <a id="L85"></a>if fd[0], p.Stdin, err = modeToFiles(stdin, 0); err != nil {
        <a id="L86"></a>goto Error
    <a id="L87"></a>}
    <a id="L88"></a>if fd[1], p.Stdout, err = modeToFiles(stdout, 1); err != nil {
        <a id="L89"></a>goto Error
    <a id="L90"></a>}
    <a id="L91"></a>if stderr == MergeWithStdout {
        <a id="L92"></a>p.Stderr = p.Stdout
    <a id="L93"></a>} else if fd[2], p.Stderr, err = modeToFiles(stderr, 2); err != nil {
        <a id="L94"></a>goto Error
    <a id="L95"></a>}

    <a id="L97"></a><span class="comment">// Run command.</span>
    <a id="L98"></a>p.Pid, err = os.ForkExec(argv0, argv, envv, &#34;&#34;, &amp;fd);
    <a id="L99"></a>if err != nil {
        <a id="L100"></a>goto Error
    <a id="L101"></a>}
    <a id="L102"></a>if fd[0] != os.Stdin {
        <a id="L103"></a>fd[0].Close()
    <a id="L104"></a>}
    <a id="L105"></a>if fd[1] != os.Stdout {
        <a id="L106"></a>fd[1].Close()
    <a id="L107"></a>}
    <a id="L108"></a>if fd[2] != os.Stderr &amp;&amp; fd[2] != fd[1] {
        <a id="L109"></a>fd[2].Close()
    <a id="L110"></a>}
    <a id="L111"></a>return p, nil;

<a id="L113"></a>Error:
    <a id="L114"></a>if fd[0] != os.Stdin &amp;&amp; fd[0] != nil {
        <a id="L115"></a>fd[0].Close()
    <a id="L116"></a>}
    <a id="L117"></a>if fd[1] != os.Stdout &amp;&amp; fd[1] != nil {
        <a id="L118"></a>fd[1].Close()
    <a id="L119"></a>}
    <a id="L120"></a>if fd[2] != os.Stderr &amp;&amp; fd[2] != nil &amp;&amp; fd[2] != fd[1] {
        <a id="L121"></a>fd[2].Close()
    <a id="L122"></a>}
    <a id="L123"></a>if p.Stdin != nil {
        <a id="L124"></a>p.Stdin.Close()
    <a id="L125"></a>}
    <a id="L126"></a>if p.Stdout != nil {
        <a id="L127"></a>p.Stdout.Close()
    <a id="L128"></a>}
    <a id="L129"></a>if p.Stderr != nil {
        <a id="L130"></a>p.Stderr.Close()
    <a id="L131"></a>}
    <a id="L132"></a>return nil, err;
<a id="L133"></a>}

<a id="L135"></a><span class="comment">// Wait waits for the running command p,</span>
<a id="L136"></a><span class="comment">// returning the Waitmsg returned by os.Wait and an error.</span>
<a id="L137"></a><span class="comment">// The options are passed through to os.Wait.</span>
<a id="L138"></a><span class="comment">// Setting options to 0 waits for p to exit;</span>
<a id="L139"></a><span class="comment">// other options cause Wait to return for other</span>
<a id="L140"></a><span class="comment">// process events; see package os for details.</span>
<a id="L141"></a>func (p *Cmd) Wait(options int) (*os.Waitmsg, os.Error) {
    <a id="L142"></a>if p.Pid &lt;= 0 {
        <a id="L143"></a>return nil, os.ErrorString(&#34;exec: invalid use of Cmd.Wait&#34;)
    <a id="L144"></a>}
    <a id="L145"></a>w, err := os.Wait(p.Pid, options);
    <a id="L146"></a>if w != nil &amp;&amp; (w.Exited() || w.Signaled()) {
        <a id="L147"></a>p.Pid = -1
    <a id="L148"></a>}
    <a id="L149"></a>return w, err;
<a id="L150"></a>}

<a id="L152"></a><span class="comment">// Close waits for the running command p to exit,</span>
<a id="L153"></a><span class="comment">// if it hasn&#39;t already, and then closes the non-nil file descriptors</span>
<a id="L154"></a><span class="comment">// p.Stdin, p.Stdout, and p.Stderr.</span>
<a id="L155"></a>func (p *Cmd) Close() os.Error {
    <a id="L156"></a>if p.Pid &gt; 0 {
        <a id="L157"></a><span class="comment">// Loop on interrupt, but</span>
        <a id="L158"></a><span class="comment">// ignore other errors -- maybe</span>
        <a id="L159"></a><span class="comment">// caller has already waited for pid.</span>
        <a id="L160"></a>_, err := p.Wait(0);
        <a id="L161"></a>for err == os.EINTR {
            <a id="L162"></a>_, err = p.Wait(0)
        <a id="L163"></a>}
    <a id="L164"></a>}

    <a id="L166"></a><span class="comment">// Close the FDs that are still open.</span>
    <a id="L167"></a>var err os.Error;
    <a id="L168"></a>if p.Stdin != nil &amp;&amp; p.Stdin.Fd() &gt;= 0 {
        <a id="L169"></a>if err1 := p.Stdin.Close(); err1 != nil {
            <a id="L170"></a>err = err1
        <a id="L171"></a>}
    <a id="L172"></a>}
    <a id="L173"></a>if p.Stdout != nil &amp;&amp; p.Stdout.Fd() &gt;= 0 {
        <a id="L174"></a>if err1 := p.Stdout.Close(); err1 != nil &amp;&amp; err != nil {
            <a id="L175"></a>err = err1
        <a id="L176"></a>}
    <a id="L177"></a>}
    <a id="L178"></a>if p.Stderr != nil &amp;&amp; p.Stderr != p.Stdout &amp;&amp; p.Stderr.Fd() &gt;= 0 {
        <a id="L179"></a>if err1 := p.Stderr.Close(); err1 != nil &amp;&amp; err != nil {
            <a id="L180"></a>err = err1
        <a id="L181"></a>}
    <a id="L182"></a>}
    <a id="L183"></a>return err;
<a id="L184"></a>}

<a id="L186"></a>func canExec(file string) bool {
    <a id="L187"></a>d, err := os.Stat(file);
    <a id="L188"></a>if err != nil {
        <a id="L189"></a>return false
    <a id="L190"></a>}
    <a id="L191"></a>return d.IsRegular() &amp;&amp; d.Permission()&amp;0111 != 0;
<a id="L192"></a>}

<a id="L194"></a><span class="comment">// LookPath searches for an executable binary named file</span>
<a id="L195"></a><span class="comment">// in the directories named by the PATH environment variable.</span>
<a id="L196"></a><span class="comment">// If file contains a slash, it is tried directly and the PATH is not consulted.</span>
<a id="L197"></a><span class="comment">//</span>
<a id="L198"></a><span class="comment">// TODO(rsc): Does LookPath belong in os instead?</span>
<a id="L199"></a>func LookPath(file string) (string, os.Error) {
    <a id="L200"></a><span class="comment">// NOTE(rsc): I wish we could use the Plan 9 behavior here</span>
    <a id="L201"></a><span class="comment">// (only bypass the path if file begins with / or ./ or ../)</span>
    <a id="L202"></a><span class="comment">// but that would not match all the Unix shells.</span>

    <a id="L204"></a>if strings.Index(file, &#34;/&#34;) &gt;= 0 {
        <a id="L205"></a>if canExec(file) {
            <a id="L206"></a>return file, nil
        <a id="L207"></a>}
        <a id="L208"></a>return &#34;&#34;, os.ENOENT;
    <a id="L209"></a>}
    <a id="L210"></a>pathenv := os.Getenv(&#34;PATH&#34;);
    <a id="L211"></a>for _, dir := range strings.Split(pathenv, &#34;:&#34;, 0) {
        <a id="L212"></a>if dir == &#34;&#34; {
            <a id="L213"></a><span class="comment">// Unix shell semantics: path element &#34;&#34; means &#34;.&#34;</span>
            <a id="L214"></a>dir = &#34;.&#34;
        <a id="L215"></a>}
        <a id="L216"></a>if canExec(dir + &#34;/&#34; + file) {
            <a id="L217"></a>return dir + &#34;/&#34; + file, nil
        <a id="L218"></a>}
    <a id="L219"></a>}
    <a id="L220"></a>return &#34;&#34;, os.ENOENT;
<a id="L221"></a>}
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
