<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/log/log.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/log/log.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Rudimentary logging package. Defines a type, Logger, with simple</span>
<a id="L6"></a><span class="comment">// methods for formatting output to one or two destinations. Also has</span>
<a id="L7"></a><span class="comment">// predefined Loggers accessible through helper functions Stdout[f],</span>
<a id="L8"></a><span class="comment">// Stderr[f], Exit[f], and Crash[f], which are easier to use than creating</span>
<a id="L9"></a><span class="comment">// a Logger manually.</span>
<a id="L10"></a><span class="comment">// Exit exits when written to.</span>
<a id="L11"></a><span class="comment">// Crash causes a crash when written to.</span>
<a id="L12"></a>package log

<a id="L14"></a>import (
    <a id="L15"></a>&#34;fmt&#34;;
    <a id="L16"></a>&#34;io&#34;;
    <a id="L17"></a>&#34;runtime&#34;;
    <a id="L18"></a>&#34;os&#34;;
    <a id="L19"></a>&#34;time&#34;;
<a id="L20"></a>)

<a id="L22"></a><span class="comment">// These flags define the properties of the Logger and the output they produce.</span>
<a id="L23"></a>const (
    <a id="L24"></a><span class="comment">// Flags</span>
    <a id="L25"></a>Lok     = iota;
    <a id="L26"></a>Lexit;  <span class="comment">// terminate execution when written</span>
    <a id="L27"></a>Lcrash; <span class="comment">// crash (panic) when written</span>
    <a id="L28"></a><span class="comment">// Bits or&#39;ed together to control what&#39;s printed. There is no control over the</span>
    <a id="L29"></a><span class="comment">// order they appear (the order listed here) or the format they present (as</span>
    <a id="L30"></a><span class="comment">// described in the comments).  A colon appears after these items:</span>
    <a id="L31"></a><span class="comment">//	2009/0123 01:23:23.123123 /a/b/c/d.go:23: message</span>
    <a id="L32"></a>Ldate          = 1 &lt;&lt; iota; <span class="comment">// the date: 2009/0123</span>
    <a id="L33"></a>Ltime;         <span class="comment">// the time: 01:23:23</span>
    <a id="L34"></a>Lmicroseconds; <span class="comment">// microsecond resolution: 01:23:23.123123.  assumes Ltime.</span>
    <a id="L35"></a>Llongfile;     <span class="comment">// full file name and line number: /a/b/c/d.go:23</span>
    <a id="L36"></a>Lshortfile;    <span class="comment">// final file name element and line number: d.go:23. overrides Llongfile</span>
    <a id="L37"></a>lAllBits       = Ldate | Ltime | Lmicroseconds | Llongfile | Lshortfile;
<a id="L38"></a>)

<a id="L40"></a><span class="comment">// Logger represents an active logging object.</span>
<a id="L41"></a>type Logger struct {
    <a id="L42"></a>out0   io.Writer; <span class="comment">// first destination for output</span>
    <a id="L43"></a>out1   io.Writer; <span class="comment">// second destination for output; may be nil</span>
    <a id="L44"></a>prefix string;    <span class="comment">// prefix to write at beginning of each line</span>
    <a id="L45"></a>flag   int;       <span class="comment">// properties</span>
<a id="L46"></a>}

<a id="L48"></a><span class="comment">// New creates a new Logger.   The out0 and out1 variables set the</span>
<a id="L49"></a><span class="comment">// destinations to which log data will be written; out1 may be nil.</span>
<a id="L50"></a><span class="comment">// The prefix appears at the beginning of each generated log line.</span>
<a id="L51"></a><span class="comment">// The flag argument defines the logging properties.</span>
<a id="L52"></a>func New(out0, out1 io.Writer, prefix string, flag int) *Logger {
    <a id="L53"></a>return &amp;Logger{out0, out1, prefix, flag}
<a id="L54"></a>}

<a id="L56"></a>var (
    <a id="L57"></a>stdout = New(os.Stdout, nil, &#34;&#34;, Lok|Ldate|Ltime);
    <a id="L58"></a>stderr = New(os.Stderr, nil, &#34;&#34;, Lok|Ldate|Ltime);
    <a id="L59"></a>exit   = New(os.Stderr, nil, &#34;&#34;, Lexit|Ldate|Ltime);
    <a id="L60"></a>crash  = New(os.Stderr, nil, &#34;&#34;, Lcrash|Ldate|Ltime);
<a id="L61"></a>)

<a id="L63"></a>var shortnames = make(map[string]string) <span class="comment">// cache of short names to avoid allocation.</span>

<a id="L65"></a><span class="comment">// Cheap integer to fixed-width decimal ASCII.  Use a negative width to avoid zero-padding</span>
<a id="L66"></a>func itoa(i int, wid int) string {
    <a id="L67"></a>var u uint = uint(i);
    <a id="L68"></a>if u == 0 &amp;&amp; wid &lt;= 1 {
        <a id="L69"></a>return &#34;0&#34;
    <a id="L70"></a>}

    <a id="L72"></a><span class="comment">// Assemble decimal in reverse order.</span>
    <a id="L73"></a>var b [32]byte;
    <a id="L74"></a>bp := len(b);
    <a id="L75"></a>for ; u &gt; 0 || wid &gt; 0; u /= 10 {
        <a id="L76"></a>bp--;
        <a id="L77"></a>wid--;
        <a id="L78"></a>b[bp] = byte(u%10) + &#39;0&#39;;
    <a id="L79"></a>}

    <a id="L81"></a>return string(b[bp:len(b)]);
<a id="L82"></a>}

<a id="L84"></a>func (l *Logger) formatHeader(ns int64, calldepth int) string {
    <a id="L85"></a>h := l.prefix;
    <a id="L86"></a>if l.flag&amp;(Ldate|Ltime|Lmicroseconds) != 0 {
        <a id="L87"></a>t := time.SecondsToLocalTime(ns / 1e9);
        <a id="L88"></a>if l.flag&amp;(Ldate) != 0 {
            <a id="L89"></a>h += itoa(int(t.Year), 4) + &#34;/&#34; + itoa(t.Month, 2) + &#34;/&#34; + itoa(t.Day, 2) + &#34; &#34;
        <a id="L90"></a>}
        <a id="L91"></a>if l.flag&amp;(Ltime|Lmicroseconds) != 0 {
            <a id="L92"></a>h += itoa(t.Hour, 2) + &#34;:&#34; + itoa(t.Minute, 2) + &#34;:&#34; + itoa(t.Second, 2);
            <a id="L93"></a>if l.flag&amp;Lmicroseconds != 0 {
                <a id="L94"></a>h += &#34;.&#34; + itoa(int(ns%1e9)/1e3, 6)
            <a id="L95"></a>}
            <a id="L96"></a>h += &#34; &#34;;
        <a id="L97"></a>}
    <a id="L98"></a>}
    <a id="L99"></a>if l.flag&amp;(Lshortfile|Llongfile) != 0 {
        <a id="L100"></a>_, file, line, ok := runtime.Caller(calldepth);
        <a id="L101"></a>if ok {
            <a id="L102"></a>if l.flag&amp;Lshortfile != 0 {
                <a id="L103"></a>short, ok := shortnames[file];
                <a id="L104"></a>if !ok {
                    <a id="L105"></a>short = file;
                    <a id="L106"></a>for i := len(file) - 1; i &gt; 0; i-- {
                        <a id="L107"></a>if file[i] == &#39;/&#39; {
                            <a id="L108"></a>short = file[i+1 : len(file)];
                            <a id="L109"></a>break;
                        <a id="L110"></a>}
                    <a id="L111"></a>}
                    <a id="L112"></a>shortnames[file] = short;
                <a id="L113"></a>}
                <a id="L114"></a>file = short;
            <a id="L115"></a>}
        <a id="L116"></a>} else {
            <a id="L117"></a>file = &#34;???&#34;;
            <a id="L118"></a>line = 0;
        <a id="L119"></a>}
        <a id="L120"></a>h += file + &#34;:&#34; + itoa(line, -1) + &#34;: &#34;;
    <a id="L121"></a>}
    <a id="L122"></a>return h;
<a id="L123"></a>}

<a id="L125"></a><span class="comment">// Output writes the output for a logging event.  The string s contains the text to print after</span>
<a id="L126"></a><span class="comment">// the time stamp;  calldepth is used to recover the PC.  It is provided for generality, although</span>
<a id="L127"></a><span class="comment">// at the moment on all pre-defined paths it will be 2.</span>
<a id="L128"></a>func (l *Logger) Output(calldepth int, s string) {
    <a id="L129"></a>now := time.Nanoseconds(); <span class="comment">// get this early.</span>
    <a id="L130"></a>newline := &#34;\n&#34;;
    <a id="L131"></a>if len(s) &gt; 0 &amp;&amp; s[len(s)-1] == &#39;\n&#39; {
        <a id="L132"></a>newline = &#34;&#34;
    <a id="L133"></a>}
    <a id="L134"></a>s = l.formatHeader(now, calldepth+1) + s + newline;
    <a id="L135"></a>io.WriteString(l.out0, s);
    <a id="L136"></a>if l.out1 != nil {
        <a id="L137"></a>io.WriteString(l.out1, s)
    <a id="L138"></a>}
    <a id="L139"></a>switch l.flag &amp; ^lAllBits {
    <a id="L140"></a>case Lcrash:
        <a id="L141"></a>panic(&#34;log: fatal error&#34;)
    <a id="L142"></a>case Lexit:
        <a id="L143"></a>os.Exit(1)
    <a id="L144"></a>}
<a id="L145"></a>}

<a id="L147"></a><span class="comment">// Logf is analogous to Printf() for a Logger.</span>
<a id="L148"></a>func (l *Logger) Logf(format string, v ...) { l.Output(2, fmt.Sprintf(format, v)) }

<a id="L150"></a><span class="comment">// Log is analogous to Print() for a Logger.</span>
<a id="L151"></a>func (l *Logger) Log(v ...) { l.Output(2, fmt.Sprintln(v)) }

<a id="L153"></a><span class="comment">// Stdout is a helper function for easy logging to stdout. It is analogous to Print().</span>
<a id="L154"></a>func Stdout(v ...) { stdout.Output(2, fmt.Sprint(v)) }

<a id="L156"></a><span class="comment">// Stderr is a helper function for easy logging to stderr. It is analogous to Fprint(os.Stderr).</span>
<a id="L157"></a>func Stderr(v ...) { stderr.Output(2, fmt.Sprintln(v)) }

<a id="L159"></a><span class="comment">// Stdoutf is a helper functions for easy formatted logging to stdout. It is analogous to Printf().</span>
<a id="L160"></a>func Stdoutf(format string, v ...) { stdout.Output(2, fmt.Sprintf(format, v)) }

<a id="L162"></a><span class="comment">// Stderrf is a helper function for easy formatted logging to stderr. It is analogous to Fprintf(os.Stderr).</span>
<a id="L163"></a>func Stderrf(format string, v ...) { stderr.Output(2, fmt.Sprintf(format, v)) }

<a id="L165"></a><span class="comment">// Exit is equivalent to Stderr() followed by a call to os.Exit(1).</span>
<a id="L166"></a>func Exit(v ...) { exit.Output(2, fmt.Sprintln(v)) }

<a id="L168"></a><span class="comment">// Exitf is equivalent to Stderrf() followed by a call to os.Exit(1).</span>
<a id="L169"></a>func Exitf(format string, v ...) { exit.Output(2, fmt.Sprintf(format, v)) }

<a id="L171"></a><span class="comment">// Crash is equivalent to Stderr() followed by a call to panic().</span>
<a id="L172"></a>func Crash(v ...) { crash.Output(2, fmt.Sprintln(v)) }

<a id="L174"></a><span class="comment">// Crashf is equivalent to Stderrf() followed by a call to panic().</span>
<a id="L175"></a>func Crashf(format string, v ...) { crash.Output(2, fmt.Sprintf(format, v)) }
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
