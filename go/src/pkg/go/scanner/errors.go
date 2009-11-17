<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/scanner/errors.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/go/scanner/errors.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package scanner

<a id="L7"></a>import (
    <a id="L8"></a>&#34;container/vector&#34;;
    <a id="L9"></a>&#34;fmt&#34;;
    <a id="L10"></a>&#34;go/token&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;sort&#34;;
<a id="L14"></a>)


<a id="L17"></a><span class="comment">// An implementation of an ErrorHandler may be provided to the Scanner.</span>
<a id="L18"></a><span class="comment">// If a syntax error is encountered and a handler was installed, Error</span>
<a id="L19"></a><span class="comment">// is called with a position and an error message. The position points</span>
<a id="L20"></a><span class="comment">// to the beginning of the offending token.</span>
<a id="L21"></a><span class="comment">//</span>
<a id="L22"></a>type ErrorHandler interface {
    <a id="L23"></a>Error(pos token.Position, msg string);
<a id="L24"></a>}


<a id="L27"></a><span class="comment">// ErrorVector implements the ErrorHandler interface. It must be</span>
<a id="L28"></a><span class="comment">// initialized with Init(). It maintains a list of errors which can</span>
<a id="L29"></a><span class="comment">// be retrieved with GetErrorList and GetError.</span>
<a id="L30"></a><span class="comment">//</span>
<a id="L31"></a><span class="comment">// A common usage pattern is to embed an ErrorVector alongside a</span>
<a id="L32"></a><span class="comment">// scanner in a data structure that uses the scanner. By passing a</span>
<a id="L33"></a><span class="comment">// reference to an ErrorVector to the scanner&#39;s Init call, default</span>
<a id="L34"></a><span class="comment">// error handling is obtained.</span>
<a id="L35"></a><span class="comment">//</span>
<a id="L36"></a>type ErrorVector struct {
    <a id="L37"></a>errors vector.Vector;
<a id="L38"></a>}


<a id="L41"></a><span class="comment">// Init initializes an ErrorVector.</span>
<a id="L42"></a>func (h *ErrorVector) Init() { h.errors.Init(0) }


<a id="L45"></a><span class="comment">// NewErrorVector creates a new ErrorVector.</span>
<a id="L46"></a>func NewErrorVector() *ErrorVector {
    <a id="L47"></a>h := new(ErrorVector);
    <a id="L48"></a>h.Init();
    <a id="L49"></a>return h;
<a id="L50"></a>}


<a id="L53"></a><span class="comment">// ErrorCount returns the number of errors collected.</span>
<a id="L54"></a>func (h *ErrorVector) ErrorCount() int { return h.errors.Len() }


<a id="L57"></a><span class="comment">// Within ErrorVector, an error is represented by an Error node. The</span>
<a id="L58"></a><span class="comment">// position Pos, if valid, points to the beginning of the offending</span>
<a id="L59"></a><span class="comment">// token, and the error condition is described by Msg.</span>
<a id="L60"></a><span class="comment">//</span>
<a id="L61"></a>type Error struct {
    <a id="L62"></a>Pos token.Position;
    <a id="L63"></a>Msg string;
<a id="L64"></a>}


<a id="L67"></a>func (e *Error) String() string {
    <a id="L68"></a>if e.Pos.Filename != &#34;&#34; || e.Pos.IsValid() {
        <a id="L69"></a><span class="comment">// don&#39;t print &#34;&lt;unknown position&gt;&#34;</span>
        <a id="L70"></a><span class="comment">// TODO(gri) reconsider the semantics of Position.IsValid</span>
        <a id="L71"></a>return e.Pos.String() + &#34;: &#34; + e.Msg
    <a id="L72"></a>}
    <a id="L73"></a>return e.Msg;
<a id="L74"></a>}


<a id="L77"></a><span class="comment">// An ErrorList is a (possibly sorted) list of Errors.</span>
<a id="L78"></a>type ErrorList []*Error


<a id="L81"></a><span class="comment">// ErrorList implements the sort Interface.</span>
<a id="L82"></a>func (p ErrorList) Len() int      { return len(p) }
<a id="L83"></a>func (p ErrorList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }


<a id="L86"></a>func (p ErrorList) Less(i, j int) bool {
    <a id="L87"></a>e := &amp;p[i].Pos;
    <a id="L88"></a>f := &amp;p[j].Pos;
    <a id="L89"></a><span class="comment">// Note that it is not sufficient to simply compare file offsets because</span>
    <a id="L90"></a><span class="comment">// the offsets do not reflect modified line information (through //line</span>
    <a id="L91"></a><span class="comment">// comments).</span>
    <a id="L92"></a>if e.Filename &lt; f.Filename {
        <a id="L93"></a>return true
    <a id="L94"></a>}
    <a id="L95"></a>if e.Filename == f.Filename {
        <a id="L96"></a>if e.Line &lt; f.Line {
            <a id="L97"></a>return true
        <a id="L98"></a>}
        <a id="L99"></a>if e.Line == f.Line {
            <a id="L100"></a>return e.Column &lt; f.Column
        <a id="L101"></a>}
    <a id="L102"></a>}
    <a id="L103"></a>return false;
<a id="L104"></a>}


<a id="L107"></a>func (p ErrorList) String() string {
    <a id="L108"></a>switch len(p) {
    <a id="L109"></a>case 0:
        <a id="L110"></a>return &#34;unspecified error&#34;
    <a id="L111"></a>case 1:
        <a id="L112"></a>return p[0].String()
    <a id="L113"></a>}
    <a id="L114"></a>return fmt.Sprintf(&#34;%s (and %d more errors)&#34;, p[0].String(), len(p)-1);
<a id="L115"></a>}


<a id="L118"></a><span class="comment">// These constants control the construction of the ErrorList</span>
<a id="L119"></a><span class="comment">// returned by GetErrors.</span>
<a id="L120"></a><span class="comment">//</span>
<a id="L121"></a>const (
    <a id="L122"></a>Raw          = iota; <span class="comment">// leave error list unchanged</span>
    <a id="L123"></a>Sorted;      <span class="comment">// sort error list by file, line, and column number</span>
    <a id="L124"></a>NoMultiples; <span class="comment">// sort error list and leave only the first error per line</span>
<a id="L125"></a>)


<a id="L128"></a><span class="comment">// GetErrorList returns the list of errors collected by an ErrorVector.</span>
<a id="L129"></a><span class="comment">// The construction of the ErrorList returned is controlled by the mode</span>
<a id="L130"></a><span class="comment">// parameter. If there are no errors, the result is nil.</span>
<a id="L131"></a><span class="comment">//</span>
<a id="L132"></a>func (h *ErrorVector) GetErrorList(mode int) ErrorList {
    <a id="L133"></a>if h.errors.Len() == 0 {
        <a id="L134"></a>return nil
    <a id="L135"></a>}

    <a id="L137"></a>list := make(ErrorList, h.errors.Len());
    <a id="L138"></a>for i := 0; i &lt; h.errors.Len(); i++ {
        <a id="L139"></a>list[i] = h.errors.At(i).(*Error)
    <a id="L140"></a>}

    <a id="L142"></a>if mode &gt;= Sorted {
        <a id="L143"></a>sort.Sort(list)
    <a id="L144"></a>}

    <a id="L146"></a>if mode &gt;= NoMultiples {
        <a id="L147"></a>var last token.Position; <span class="comment">// initial last.Line is != any legal error line</span>
        <a id="L148"></a>i := 0;
        <a id="L149"></a>for _, e := range list {
            <a id="L150"></a>if e.Pos.Filename != last.Filename || e.Pos.Line != last.Line {
                <a id="L151"></a>last = e.Pos;
                <a id="L152"></a>list[i] = e;
                <a id="L153"></a>i++;
            <a id="L154"></a>}
        <a id="L155"></a>}
        <a id="L156"></a>list = list[0:i];
    <a id="L157"></a>}

    <a id="L159"></a>return list;
<a id="L160"></a>}


<a id="L163"></a><span class="comment">// GetError is like GetErrorList, but it returns an os.Error instead</span>
<a id="L164"></a><span class="comment">// so that a nil result can be assigned to an os.Error variable and</span>
<a id="L165"></a><span class="comment">// remains nil.</span>
<a id="L166"></a><span class="comment">//</span>
<a id="L167"></a>func (h *ErrorVector) GetError(mode int) os.Error {
    <a id="L168"></a>if h.errors.Len() == 0 {
        <a id="L169"></a>return nil
    <a id="L170"></a>}

    <a id="L172"></a>return h.GetErrorList(mode);
<a id="L173"></a>}


<a id="L176"></a><span class="comment">// ErrorVector implements the ErrorHandler interface.</span>
<a id="L177"></a>func (h *ErrorVector) Error(pos token.Position, msg string) {
    <a id="L178"></a>h.errors.Push(&amp;Error{pos, msg})
<a id="L179"></a>}


<a id="L182"></a><span class="comment">// PrintError is a utility function that prints a list of errors to w,</span>
<a id="L183"></a><span class="comment">// one error per line, if the err parameter is an ErrorList. Otherwise</span>
<a id="L184"></a><span class="comment">// it prints the err string.</span>
<a id="L185"></a><span class="comment">//</span>
<a id="L186"></a>func PrintError(w io.Writer, err os.Error) {
    <a id="L187"></a>if list, ok := err.(ErrorList); ok {
        <a id="L188"></a>for _, e := range list {
            <a id="L189"></a>fmt.Fprintf(w, &#34;%s\n&#34;, e)
        <a id="L190"></a>}
    <a id="L191"></a>} else {
        <a id="L192"></a>fmt.Fprintf(w, &#34;%s\n&#34;, err)
    <a id="L193"></a>}
<a id="L194"></a>}
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
