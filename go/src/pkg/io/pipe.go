<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/io/pipe.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/io/pipe.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Pipe adapter to connect code expecting an io.Read</span>
<a id="L6"></a><span class="comment">// with code expecting an io.Write.</span>

<a id="L8"></a>package io

<a id="L10"></a>import (
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;sync&#34;;
<a id="L13"></a>)

<a id="L15"></a>type pipeReturn struct {
    <a id="L16"></a>n   int;
    <a id="L17"></a>err os.Error;
<a id="L18"></a>}

<a id="L20"></a><span class="comment">// Shared pipe structure.</span>
<a id="L21"></a>type pipe struct {
    <a id="L22"></a>rclosed bool;            <span class="comment">// Read end closed?</span>
    <a id="L23"></a>rerr    os.Error;        <span class="comment">// Error supplied to CloseReader</span>
    <a id="L24"></a>wclosed bool;            <span class="comment">// Write end closed?</span>
    <a id="L25"></a>werr    os.Error;        <span class="comment">// Error supplied to CloseWriter</span>
    <a id="L26"></a>wpend   []byte;          <span class="comment">// Written data waiting to be read.</span>
    <a id="L27"></a>wtot    int;             <span class="comment">// Bytes consumed so far in current write.</span>
    <a id="L28"></a>cr      chan []byte;     <span class="comment">// Write sends data here...</span>
    <a id="L29"></a>cw      chan pipeReturn; <span class="comment">// ... and reads the n, err back from here.</span>
<a id="L30"></a>}

<a id="L32"></a>func (p *pipe) Read(data []byte) (n int, err os.Error) {
    <a id="L33"></a>if p == nil || p.rclosed {
        <a id="L34"></a>return 0, os.EINVAL
    <a id="L35"></a>}

    <a id="L37"></a><span class="comment">// Wait for next write block if necessary.</span>
    <a id="L38"></a>if p.wpend == nil {
        <a id="L39"></a>if !p.wclosed {
            <a id="L40"></a>p.wpend = &lt;-p.cr
        <a id="L41"></a>}
        <a id="L42"></a>if p.wpend == nil {
            <a id="L43"></a>return 0, p.werr
        <a id="L44"></a>}
        <a id="L45"></a>p.wtot = 0;
    <a id="L46"></a>}

    <a id="L48"></a><span class="comment">// Read from current write block.</span>
    <a id="L49"></a>n = len(data);
    <a id="L50"></a>if n &gt; len(p.wpend) {
        <a id="L51"></a>n = len(p.wpend)
    <a id="L52"></a>}
    <a id="L53"></a>for i := 0; i &lt; n; i++ {
        <a id="L54"></a>data[i] = p.wpend[i]
    <a id="L55"></a>}
    <a id="L56"></a>p.wtot += n;
    <a id="L57"></a>p.wpend = p.wpend[n:len(p.wpend)];

    <a id="L59"></a><span class="comment">// If write block is done, finish the write.</span>
    <a id="L60"></a>if len(p.wpend) == 0 {
        <a id="L61"></a>p.wpend = nil;
        <a id="L62"></a>p.cw &lt;- pipeReturn{p.wtot, nil};
        <a id="L63"></a>p.wtot = 0;
    <a id="L64"></a>}

    <a id="L66"></a>return n, nil;
<a id="L67"></a>}

<a id="L69"></a>func (p *pipe) Write(data []byte) (n int, err os.Error) {
    <a id="L70"></a>if p == nil || p.wclosed {
        <a id="L71"></a>return 0, os.EINVAL
    <a id="L72"></a>}
    <a id="L73"></a>if p.rclosed {
        <a id="L74"></a>return 0, p.rerr
    <a id="L75"></a>}

    <a id="L77"></a><span class="comment">// Send data to reader.</span>
    <a id="L78"></a>p.cr &lt;- data;

    <a id="L80"></a><span class="comment">// Wait for reader to finish copying it.</span>
    <a id="L81"></a>res := &lt;-p.cw;
    <a id="L82"></a>return res.n, res.err;
<a id="L83"></a>}

<a id="L85"></a>func (p *pipe) CloseReader(rerr os.Error) os.Error {
    <a id="L86"></a>if p == nil || p.rclosed {
        <a id="L87"></a>return os.EINVAL
    <a id="L88"></a>}

    <a id="L90"></a><span class="comment">// Stop any future writes.</span>
    <a id="L91"></a>p.rclosed = true;
    <a id="L92"></a>if rerr == nil {
        <a id="L93"></a>rerr = os.EPIPE
    <a id="L94"></a>}
    <a id="L95"></a>p.rerr = rerr;

    <a id="L97"></a><span class="comment">// Stop the current write.</span>
    <a id="L98"></a>if !p.wclosed {
        <a id="L99"></a>p.cw &lt;- pipeReturn{p.wtot, rerr}
    <a id="L100"></a>}

    <a id="L102"></a>return nil;
<a id="L103"></a>}

<a id="L105"></a>func (p *pipe) CloseWriter(werr os.Error) os.Error {
    <a id="L106"></a>if werr == nil {
        <a id="L107"></a>werr = os.EOF
    <a id="L108"></a>}
    <a id="L109"></a>if p == nil || p.wclosed {
        <a id="L110"></a>return os.EINVAL
    <a id="L111"></a>}

    <a id="L113"></a><span class="comment">// Stop any future reads.</span>
    <a id="L114"></a>p.wclosed = true;
    <a id="L115"></a>p.werr = werr;

    <a id="L117"></a><span class="comment">// Stop the current read.</span>
    <a id="L118"></a>if !p.rclosed {
        <a id="L119"></a>p.cr &lt;- nil
    <a id="L120"></a>}

    <a id="L122"></a>return nil;
<a id="L123"></a>}

<a id="L125"></a><span class="comment">// Read/write halves of the pipe.</span>
<a id="L126"></a><span class="comment">// They are separate structures for two reasons:</span>
<a id="L127"></a><span class="comment">//  1.  If one end becomes garbage without being Closed,</span>
<a id="L128"></a><span class="comment">//      its finisher can Close so that the other end</span>
<a id="L129"></a><span class="comment">//      does not hang indefinitely.</span>
<a id="L130"></a><span class="comment">//  2.  Clients cannot use interface conversions on the</span>
<a id="L131"></a><span class="comment">//      read end to find the Write method, and vice versa.</span>

<a id="L133"></a><span class="comment">// A PipeReader is the read half of a pipe.</span>
<a id="L134"></a>type PipeReader struct {
    <a id="L135"></a>lock sync.Mutex;
    <a id="L136"></a>p    *pipe;
<a id="L137"></a>}

<a id="L139"></a><span class="comment">// Read implements the standard Read interface:</span>
<a id="L140"></a><span class="comment">// it reads data from the pipe, blocking until a writer</span>
<a id="L141"></a><span class="comment">// arrives or the write end is closed.</span>
<a id="L142"></a><span class="comment">// If the write end is closed with an error, that error is</span>
<a id="L143"></a><span class="comment">// returned as err; otherwise err is nil.</span>
<a id="L144"></a>func (r *PipeReader) Read(data []byte) (n int, err os.Error) {
    <a id="L145"></a>r.lock.Lock();
    <a id="L146"></a>defer r.lock.Unlock();

    <a id="L148"></a>return r.p.Read(data);
<a id="L149"></a>}

<a id="L151"></a><span class="comment">// Close closes the reader; subsequent writes to the</span>
<a id="L152"></a><span class="comment">// write half of the pipe will return the error os.EPIPE.</span>
<a id="L153"></a>func (r *PipeReader) Close() os.Error {
    <a id="L154"></a>r.lock.Lock();
    <a id="L155"></a>defer r.lock.Unlock();

    <a id="L157"></a>return r.p.CloseReader(nil);
<a id="L158"></a>}

<a id="L160"></a><span class="comment">// CloseWithError closes the reader; subsequent writes</span>
<a id="L161"></a><span class="comment">// to the write half of the pipe will return the error rerr.</span>
<a id="L162"></a>func (r *PipeReader) CloseWithError(rerr os.Error) os.Error {
    <a id="L163"></a>r.lock.Lock();
    <a id="L164"></a>defer r.lock.Unlock();

    <a id="L166"></a>return r.p.CloseReader(rerr);
<a id="L167"></a>}

<a id="L169"></a>func (r *PipeReader) finish() { r.Close() }

<a id="L171"></a><span class="comment">// Write half of pipe.</span>
<a id="L172"></a>type PipeWriter struct {
    <a id="L173"></a>lock sync.Mutex;
    <a id="L174"></a>p    *pipe;
<a id="L175"></a>}

<a id="L177"></a><span class="comment">// Write implements the standard Write interface:</span>
<a id="L178"></a><span class="comment">// it writes data to the pipe, blocking until readers</span>
<a id="L179"></a><span class="comment">// have consumed all the data or the read end is closed.</span>
<a id="L180"></a><span class="comment">// If the read end is closed with an error, that err is</span>
<a id="L181"></a><span class="comment">// returned as err; otherwise err is os.EPIPE.</span>
<a id="L182"></a>func (w *PipeWriter) Write(data []byte) (n int, err os.Error) {
    <a id="L183"></a>w.lock.Lock();
    <a id="L184"></a>defer w.lock.Unlock();

    <a id="L186"></a>return w.p.Write(data);
<a id="L187"></a>}

<a id="L189"></a><span class="comment">// Close closes the writer; subsequent reads from the</span>
<a id="L190"></a><span class="comment">// read half of the pipe will return no bytes and a nil error.</span>
<a id="L191"></a>func (w *PipeWriter) Close() os.Error {
    <a id="L192"></a>w.lock.Lock();
    <a id="L193"></a>defer w.lock.Unlock();

    <a id="L195"></a>return w.p.CloseWriter(nil);
<a id="L196"></a>}

<a id="L198"></a><span class="comment">// CloseWithError closes the writer; subsequent reads from the</span>
<a id="L199"></a><span class="comment">// read half of the pipe will return no bytes and the error werr.</span>
<a id="L200"></a>func (w *PipeWriter) CloseWithError(werr os.Error) os.Error {
    <a id="L201"></a>w.lock.Lock();
    <a id="L202"></a>defer w.lock.Unlock();

    <a id="L204"></a>return w.p.CloseWriter(werr);
<a id="L205"></a>}

<a id="L207"></a>func (w *PipeWriter) finish() { w.Close() }

<a id="L209"></a><span class="comment">// Pipe creates a synchronous in-memory pipe.</span>
<a id="L210"></a><span class="comment">// It can be used to connect code expecting an io.Reader</span>
<a id="L211"></a><span class="comment">// with code expecting an io.Writer.</span>
<a id="L212"></a><span class="comment">// Reads on one end are matched with writes on the other,</span>
<a id="L213"></a><span class="comment">// copying data directly between the two; there is no internal buffering.</span>
<a id="L214"></a>func Pipe() (*PipeReader, *PipeWriter) {
    <a id="L215"></a>p := new(pipe);
    <a id="L216"></a>p.cr = make(chan []byte, 1);
    <a id="L217"></a>p.cw = make(chan pipeReturn, 1);
    <a id="L218"></a>r := new(PipeReader);
    <a id="L219"></a>r.p = p;
    <a id="L220"></a>w := new(PipeWriter);
    <a id="L221"></a>w.p = p;
    <a id="L222"></a>return r, w;
<a id="L223"></a>}
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
